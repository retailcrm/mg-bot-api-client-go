package v1

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func TestMain(m *testing.M) {
	if os.Getenv("DEVELOPER_NODE") == "1" {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		os.Exit(m.Run())
	}
}

var (
	mgURL   = os.Getenv("MG_URL")
	mgToken = os.Getenv("MG_BOT_TOKEN")
)

func client() *MgClient {
	return New(mgURL, mgToken)
}

func TestMgClient_Bots(t *testing.T) {
	c := client()
	req := BotsRequest{Active: 1}

	data, status, err := c.Bots(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	for _, bot := range data {
		assert.NotEmpty(t, bot.CreatedAt)
	}
}

func TestMgClient_Channels(t *testing.T) {
	c := client()
	req := ChannelsRequest{Active: 1}

	data, status, err := c.Channels(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	for _, channel := range data {
		assert.NotEmpty(t, channel.Type)
	}
}

func TestMgClient_Users(t *testing.T) {
	c := client()
	req := UsersRequest{Active: 1}

	data, status, err := c.Users(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	for _, user := range data {
		assert.NotEmpty(t, user.FirstName)
	}
}

func TestMgClient_Customers(t *testing.T) {
	c := client()
	req := CustomersRequest{}

	data, status, err := c.Customers(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	for _, customer := range data {
		assert.NotEmpty(t, customer.ChannelId)
	}
}

func TestMgClient_Chats(t *testing.T) {
	c := client()
	req := ChatsRequest{ChannelType: ChannelTypeTelegram}

	data, status, err := c.Chats(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	for _, chat := range data {
		assert.NotEmpty(t, chat.Customer.Name)
	}
}

func TestMgClient_Members(t *testing.T) {
	c := client()
	req := MembersRequest{State: ChatMemberStateLeaved}

	data, status, err := c.Members(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)

	for _, member := range data {
		assert.NotEmpty(t, member.ChatID)
	}
}

func TestMgClient_Dialogs(t *testing.T) {
	c := client()
	req := DialogsRequest{Active: 0}

	data, status, err := c.Dialogs(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	for _, dialog := range data {
		assert.NotEmpty(t, dialog.ChatID)
	}
}

func TestMgClient_DialogAssign(t *testing.T) {
	c := client()
	i, err := strconv.ParseUint(os.Getenv("MG_BOT_DIALOG"), 10, 64)
	m, err := strconv.ParseUint(os.Getenv("MG_BOT_USER"), 10, 64)
	req := DialogAssignRequest{DialogID: i, UserID: m}

	_, status, err := c.DialogAssign(req)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, status)
}

func TestMgClient_DialogClose(t *testing.T) {
	c := client()
	i, err := strconv.ParseUint(os.Getenv("MG_BOT_DIALOG"), 10, 64)
	_, status, err := c.DialogClose(i)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, status)
}

func TestMgClient_Messages(t *testing.T) {
	c := client()
	req := MessagesRequest{ChannelType: ChannelTypeTelegram, Scope: MessageScopePublic}

	data, status, err := c.Messages(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	for _, message := range data {
		assert.NotEmpty(t, message.Content)
	}
}

func TestMgClient_MessageSendText(t *testing.T) {
	c := client()
	i, err := strconv.ParseUint(os.Getenv("MG_BOT_CHAT"), 10, 64)
	message := MessageSendRequest{
		Type:    MsgTypeText,
		Scope:   "public",
		Content: "test",
		ChatID:  i,
	}

	data, status, err := c.MessageSend(message)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, data.MessageID)
}

func TestMgClient_MessageSendProduct(t *testing.T) {
	c := client()

	msg, _, err := c.MessageSend(MessageSendRequest{
		Type:   MsgTypeProduct,
		ChatID: 5,
		Scope:  "public",
		Product: &MessageProduct{
			ID:      1,
			Name:    "Some Product",
			Article: "Art-111",
			Url:     "https://example.com",
			Img:     "http://example.com/pic.jpg",
			Cost: &MessageOrderCost{
				Value:    29900,
				Currency: "rub",
			},
			Quantity: &MessageOrderQuantity{
				Value: 1,
			},
		},
	})

	if err != nil {
		t.Errorf("%v", err)
	}

	assert.NoError(t, err)
	t.Logf("%v", msg)
}

func TestMgClient_MessageSendOrder(t *testing.T) {
	c := client()

	msg, _, err := c.MessageSend(MessageSendRequest{
		Type:   MsgTypeOrder,
		ChatID: 5,
		Scope:  "public",
		Order: &MessageOrder{
			Number: RandStringBytesMaskImprSrc(7),
			Cost: &MessageOrderCost{
				Value:    29900,
				Currency: MsgCurrencyRub,
			},
			Status: &MessageOrderStatus{
				Code: MsgOrderStatusCodeNew,
				Name: "Новый",
			},
			Delivery: &MessageOrderDelivery{
				Name:    "Курьерская доставка",
				Address: "г. Москва, Проспект Мира, 9",
				Amount: &MessageOrderCost{
					Value:    1100,
					Currency: MsgCurrencyRub,
				},
			},
			Items: []MessageOrderItem{
				{
					Name: "iPhone 6",
					Price: &MessageOrderCost{
						Value:    29900,
						Currency: MsgCurrencyRub,
					},
					Quantity: &MessageOrderQuantity{
						Value: 1,
					},
				},
			},
		},
	})

	if err != nil {
		t.Errorf("%v", err)
	}

	assert.NoError(t, err)
	t.Logf("%v", msg)
}

func TestMgClient_RandomStringGenerator(t *testing.T) {
	rnd := RandStringBytesMaskImprSrc(7)
	assert.NotEmpty(t, rnd)
	t.Logf("%v", rnd)
}

func TestMgClient_MessageEdit(t *testing.T) {
	c := client()
	i, err := strconv.ParseUint(os.Getenv("MG_BOT_CHAT"), 10, 64)
	message := MessageSendRequest{
		Type:    MsgTypeText,
		Scope:   "public",
		Content: "test",
		ChatID:  i,
	}

	s, status, err := c.MessageSend(message)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, s.MessageID)

	edit := MessageEditRequest{
		ID:      s.MessageID,
		Content: "test",
	}

	e, status, err := c.MessageEdit(edit)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	t.Logf("Message edit: %v", e)
}

func TestMgClient_MessageDelete(t *testing.T) {
	c := client()
	i, err := strconv.ParseUint(os.Getenv("MG_BOT_CHAT"), 10, 64)
	message := MessageSendRequest{
		Type:    MsgTypeText,
		Scope:   "public",
		Content: "test",
		ChatID:  i,
	}

	s, status, err := c.MessageSend(message)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, s.MessageID)

	d, status, err := c.MessageDelete(s.MessageID)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	t.Logf("Message delete: %v", d)
}

func TestMgClient_Info(t *testing.T) {
	c := client()
	req := InfoRequest{Name: "AWESOME", Avatar: os.Getenv("MG_BOT_LOGO")}

	_, status, err := c.Info(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
}

func TestMgClient_Commands(t *testing.T) {
	c := client()
	req := CommandsRequest{}

	data, status, err := c.Commands(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	for _, command := range data {
		assert.NotEmpty(t, command.Description)
	}
}

func TestMgClient_CommandEditDelete(t *testing.T) {
	c := client()
	req := CommandEditRequest{
		Name:        "show_payment_types",
		Description: "Get available payment types",
	}

	n, status, err := c.CommandEdit(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, n.ID)

	d, status, err := c.CommandDelete(n.Name)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	t.Logf("%v", d)
}

func TestMgClient_WsMeta(t *testing.T) {
	c := client()
	events := []string{"user_updated", "user_join_chat"}
	url, headers, err := c.WsMeta(events)

	if err != nil {
		t.Errorf("%v", err)
	}

	resUrl := fmt.Sprintf("%s%s%s%s", strings.Replace(c.URL, "https", "wss", 1), prefix, "/ws?events=", strings.Join(events[:], ","))
	resToken := c.Token

	assert.Equal(t, resUrl, url)
	assert.Equal(t, resToken, headers["X-Bot-Token"][0])
}

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
