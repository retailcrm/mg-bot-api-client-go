package v1

import (
	"bytes"
	"encoding/json"
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
	"gopkg.in/h2non/gock.v1"
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
	}

	os.Exit(m.Run())
}

var (
	mgURL    = "https://api.example.com"
	mgToken  = "test_token"
	debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
)

func client(opts ...Option) *MgClient {
	if debug != false {
		opts = append(opts, OptionDebug())
	}

	return New(mgURL, mgToken, opts...)
}

func TestMgClient_Bots(t *testing.T) {
	c := client()

	defer gock.Off()

	gock.New(mgURL).
		Get("/api/bot/v1/bots").
		Reply(200).
		BodyString(`[{"id": 1, "name": "Test Bot", "created_at": "2018-01-01T00:00:00.000000Z", "is_active": true, "is_self": true}]`)

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

	defer gock.Off()

	gock.New(mgURL).
		Get("/api/bot/v1/channels").
		Reply(200).
		BodyString(`[
			{
				"id": 1,
				"type": "custom",
				"name": "Test custom channel",
				"settings": {
				  "customer_external_id": "phone",
				  "sending_policy": {
					"new_customer": "no",
					"after_reply_timeout": "template"
				  },
				  "status": {
					"delivered": "both",
					"read": "receive"
				  },
				  "text": {
					"creating": "both",
					"editing": "receive",
					"quoting": "send",
					"deleting": "receive",
					"max_chars_count": 777
				  },
				  "product": {
					"creating": "receive",
					"editing": "receive",
					"deleting": "receive"
				  },
				  "order": {
					"creating": "receive",
					"editing": "receive",
					"deleting": "receive"
				  },
				  "image": {
					"creating": "both",
					"quoting": "receive",
					"editing": "none",
					"deleting": "receive",
					"max_items_count": 1,
					"note_max_chars_count": 777
				  },
				  "file": {
					"creating": "both",
					"quoting": "receive",
					"editing": "none",
					"deleting": "receive",
					"max_items_count": 1,
					"note_max_chars_count": 777
				  },
				  "suggestions": {
					"text": "receive",
					"email": "receive",
					"phone": "receive"
				  }
				},
				"created_at": "2018-01-01T00:00:00.000000Z",
				"updated_at": null,
				"activated_at": "2018-01-01T00:00:00.000000Z",
				"deactivated_at": null,
				"is_active": true
			  }
		]`)

	channels, status, err := c.Channels(ChannelsRequest{Active: 1})
	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Len(t, channels, 1)

	ch := channels[0]
	assert.Equal(t, uint64(1), ch.ID)
	assert.Equal(t, ChannelTypeCustom, ch.Type)
	assert.Equal(t, "Test custom channel", ch.Name)
	assert.Equal(t, "2018-01-01T00:00:00.000000Z", ch.CreatedAt)
	assert.Empty(t, ch.UpdatedAt)
	assert.Equal(t, "2018-01-01T00:00:00.000000Z", ch.ActivatedAt)
	assert.Empty(t, ch.DeactivatedAt)
	assert.True(t, ch.IsActive)

	chs := ch.Settings
	assert.Equal(t, "phone", chs.CustomerExternalID)

	assert.Equal(t, "no", chs.SendingPolicy.NewCustomer)
	assert.Equal(t, "template", chs.SendingPolicy.AfterReplyTimeout)

	assert.Equal(t, ChannelFeatureBoth, chs.Status.Delivered)
	assert.Equal(t, ChannelFeatureReceive, chs.Status.Read)

	assert.Equal(t, ChannelFeatureBoth, chs.Text.Creating)
	assert.Equal(t, ChannelFeatureReceive, chs.Text.Editing)
	assert.Equal(t, ChannelFeatureSend, chs.Text.Quoting)
	assert.Equal(t, ChannelFeatureReceive, chs.Text.Deleting)
	assert.Equal(t, uint16(777), chs.Text.MaxCharsCount)

	assert.Equal(t, ChannelFeatureReceive, chs.Product.Creating)
	assert.Equal(t, ChannelFeatureReceive, chs.Product.Editing)
	assert.Equal(t, ChannelFeatureReceive, chs.Product.Deleting)

	assert.Equal(t, ChannelFeatureReceive, chs.Order.Creating)
	assert.Equal(t, ChannelFeatureReceive, chs.Order.Editing)
	assert.Equal(t, ChannelFeatureReceive, chs.Order.Deleting)

	assert.Equal(t, ChannelFeatureBoth, chs.Image.Creating)
	assert.Equal(t, ChannelFeatureNone, chs.Image.Editing)
	assert.Equal(t, ChannelFeatureReceive, chs.Image.Quoting)
	assert.Equal(t, ChannelFeatureReceive, chs.Image.Deleting)
	assert.Equal(t, 1, chs.Image.MaxItemsCount)
	assert.Equal(t, uint16(777), chs.Image.NoteMaxCharsCount)

	assert.Equal(t, ChannelFeatureBoth, chs.File.Creating)
	assert.Equal(t, ChannelFeatureNone, chs.File.Editing)
	assert.Equal(t, ChannelFeatureReceive, chs.File.Quoting)
	assert.Equal(t, ChannelFeatureReceive, chs.File.Deleting)
	assert.Equal(t, 1, chs.File.MaxItemsCount)
	assert.Equal(t, uint16(777), chs.File.NoteMaxCharsCount)

	assert.Equal(t, ChannelFeatureReceive, chs.Suggestions.Text)
	assert.Equal(t, ChannelFeatureReceive, chs.Suggestions.Email)
	assert.Equal(t, ChannelFeatureReceive, chs.Suggestions.Phone)
}

func TestMgClient_Users(t *testing.T) {
	c := client()

	defer gock.Off()

	gock.New(mgURL).
		Get("/api/bot/v1/users").
		Reply(200).
		BodyString(`[{"id": 1, "external_id":"1", "username": "Test", "first_name":"Test", "last_name":"Test", "created_at": "2018-01-01T00:00:00.000000Z", "is_active": true, "is_online": true, "is_technical_account": true}]`)

	req := UsersRequest{Active: 1}

	data, status, err := c.Users(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	for _, user := range data {
		assert.Equal(t, uint64(1), user.ID)
		assert.Equal(t, "1", user.ExternalID)
		assert.Equal(t, "Test", user.Username)
		assert.Equal(t, "Test", user.FirstName)
		assert.Equal(t, "Test", user.LastName)
		assert.Equal(t, "2018-01-01T00:00:00.000000Z", user.CreatedAt)
		assert.Equal(t, true, user.IsActive)
		assert.Equal(t, true, user.IsOnline)
		assert.Equal(t, true, user.IsTechnicalAccount)
	}
}

func TestMgClient_Customers(t *testing.T) {
	c := client()

	defer gock.Off()

	response := `
	[
		{
			"id": 1,
			"channel_id": 1, 
			"created_at": 
			"2018-01-01T00:00:00.000000Z", 
			"utm": {
				"source": "test"
			}
		},
		{
			"id": 2,
			"channel_id": 1, 
			"created_at": 
			"2018-01-01T00:00:00.000000Z", 
			"utm": {
				"source": null
			}
		},
		{
			"id": 3,
			"channel_id": 1, 
			"created_at": 
			"2018-01-01T00:00:00.000000Z", 
			"utm": null
		},
		{
			"id": 4,
			"channel_id": 1, 
			"created_at": "2018-01-01T00:00:00.000000Z"
		}
	]`

	gock.New(mgURL).
		Get("/api/bot/v1/customers").
		Reply(200).
		BodyString(response)

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

	assert.Equal(t, "test", data[0].Utm.Source)
	assert.Equal(t, "", data[1].Utm.Source)
	assert.Nil(t, data[2].Utm)
	assert.Nil(t, data[3].Utm)
}

func TestMgClient_Chats(t *testing.T) {
	c := client()

	defer gock.Off()

	gock.New(mgURL).
		Get("/api/bot/v1/chats").
		Reply(200).
		BodyString(`[
			{"id": 2,"customer": {"id": 2, "name": "Foo"}, "created_at": "2018-01-01T00:00:00.000000Z"},
			{"id": 3,"customer": {"id": 3, "name": "Bar"}, "created_at": "2018-01-02T00:00:00.000000Z"}
		]`)

	req := ChatsRequest{
		ChannelType: ChannelTypeTelegram,
		SinceID:     1,
	}

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

	defer gock.Off()

	gock.New(mgURL).
		Get("/api/bot/v1/members").
		Reply(200).
		BodyString(`[{"id": 1,"user_id": 1, "chat_id": 1, "created_at": "2018-01-01T00:00:00.000000Z"}]`)

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

	defer gock.Off()

	gock.New(mgURL).
		Get("/api/bot/v1/dialogs").
		Reply(200).
		BodyString(`[{"id": 1, "chat_id": 1, "created_at": "2018-01-01T00:00:00.000000Z"}]`)

	req := DialogsRequest{Active: 0, SinceID: 1}

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

	d := 1
	u := 1
	req := DialogAssignRequest{DialogID: uint64(d), UserID: uint64(u)}
	r, _ := json.Marshal(req)

	defer gock.Off()

	gock.New(mgURL).
		Patch("/api/bot/v1/dialogs/1/assign").
		JSON(r).
		Reply(400).
		BodyString(`{"errors": ["dialog is not the latest in the chat"]}`)

	_, status, err := c.DialogAssign(req)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, status)
}

func TestMgClient_DialogUnassign(t *testing.T) {
	c := client()
	defer gock.Off()

	t.Run("success", func(t *testing.T) {
		gock.New(mgURL).
			Patch("/api/bot/v1/dialogs/777/unassign").
			Reply(200).
			BodyString(`{"previous_responsible": {"id": 111, "type": "bot", "assigned_at": "2020-07-14T14:11:44.000000Z"}}`)

		resp, status, err := c.DialogUnassign(777)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, status)

		assert.Equal(t, int64(111), resp.PreviousResponsible.ID)
		assert.Equal(t, "bot", resp.PreviousResponsible.Type)
		assert.Equal(t, "2020-07-14T14:11:44.000000Z", resp.PreviousResponsible.AssignAt)
	})

	t.Run("dialog not latest in chat", func(t *testing.T) {
		gock.New(mgURL).
			Patch("/api/bot/v1/dialogs/666/unassign").
			Reply(400).
			BodyString(`{"errors": ["dialog is not the latest in the chat"]}`)

		_, status, err := c.DialogUnassign(666)

		assert.Error(t, err, "dialog is not the latest in the chat")
		assert.Equal(t, http.StatusBadRequest, status)
	})

	t.Run("dialog is not assigned", func(t *testing.T) {
		gock.New(mgURL).
			Patch("/api/bot/v1/dialogs/555/unassign").
			Reply(400).
			BodyString(`{"errors": ["dialog is not assigned"]}`)

		_, status, err := c.DialogUnassign(555)

		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, status)
	})

	t.Run("dialog not found", func(t *testing.T) {
		gock.New(mgURL).
			Patch("/api/bot/v1/dialogs/444/unassign").
			Reply(404).
			BodyString(`{"errors": ["dialog #444 not found"]}`)

		_, status, err := c.DialogUnassign(444)

		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, status)
	})
}

func TestMgClient_DialogClose(t *testing.T) {
	c := client()
	i := 1

	defer gock.Off()

	gock.New(mgURL).
		Delete("/api/bot/v1/dialogs/1/close").
		Reply(400).
		BodyString(`{"errors": ["dialog #1 not found"]}`)

	_, status, err := c.DialogClose(uint64(i))

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, status)
}

func TestMgClient_DialogsTagsAdd(t *testing.T) {
	c := client()

	color := ColorBlue
	req := DialogTagsAddRequest{
		DialogID: uint64(1),
		Tags: []TagsAdd{
			{Name: "foo", ColorCode: nil},
			{Name: "bar", ColorCode: &color},
		},
	}
	r, _ := json.Marshal(req)

	defer gock.Off()

	gock.New(mgURL).
		Patch("/api/bot/v1/dialogs/1/tags/add").
		JSON(r).
		Reply(200).
		BodyString(`{}`)

	status, err := c.DialogsTagsAdd(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
}

func TestMgClient_DialogsTagsDelete(t *testing.T) {
	c := client()

	req := DialogTagsDeleteRequest{
		DialogID: uint64(1),
		Tags: []TagsDelete{
			{Name: "foo"},
			{Name: "bar"},
		},
	}
	r, _ := json.Marshal(req)

	defer gock.Off()

	gock.New(mgURL).
		Patch("/api/bot/v1/dialogs/1/tags/delete").
		JSON(r).
		Reply(200).
		BodyString(`{}`)

	status, err := c.DialogTagsDelete(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
}

func TestMgClient_Messages(t *testing.T) {
	c := client()

	defer gock.Off()

	gock.New(mgURL).
		Get("/api/bot/v1/messages").
		Reply(200).
		BodyString(`[{"id": 1, "time": "2018-01-01T00:00:00+03:00", "type": "text", "scope": "public", "chat_id": 1, "is_read": false, "is_edit": false, "status": "received", "created_at": "2018-01-01T00:00:00.000000Z"}]`)

	req := MessagesRequest{ChannelType: ChannelTypeTelegram, Scope: MessageScopePublic}

	data, status, err := c.Messages(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	for _, message := range data {
		assert.NotEmpty(t, message.ID)
	}
}

func TestMgClient_MessageSendText(t *testing.T) {
	c := client()

	i := uint64(1)
	message := MessageSendRequest{
		Type:    MsgTypeText,
		Scope:   "public",
		Content: "test",
		ChatID:  i,
	}

	defer gock.Off()

	gock.New(mgURL).
		Post("/api/bot/v1/messages").
		JSON(message).
		Reply(200).
		BodyString(`{"message_id": 1, "time": "2018-01-01T00:00:00+03:00"}`)

	data, status, err := c.MessageSend(message)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, data.MessageID)
}

func TestMgClient_MessageSendTextWithSuggestions(t *testing.T) {
	c := client()

	i := uint64(1)
	message := MessageSendRequest{
		Type:    MsgTypeText,
		Scope:   "public",
		Content: "test message with suggestions",
		ChatID:  i,
		TransportAttachments: &TransportAttachments{
			Suggestions: []Suggestion{
				{
					Type:  SuggestionTypeText,
					Title: "text suggestion",
				},
				{Type: SuggestionTypeEmail},
				{Type: SuggestionTypePhone},
			},
		},
	}

	defer gock.Off()

	gock.New(mgURL).
		Post("/api/bot/v1/messages").
		JSON(message).
		Reply(200).
		BodyString(`{"message_id": 1, "time": "2018-01-01T00:00:00+03:00"}`)

	data, status, err := c.MessageSend(message)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, data.MessageID)
}

func TestMgClient_MessageSendProduct(t *testing.T) {
	c := client()

	message := MessageSendRequest{
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
	}

	defer gock.Off()

	gock.New(mgURL).
		Post("/api/bot/v1/messages").
		JSON(message).
		Reply(200).
		BodyString(`{"message_id": 1, "time": "2018-01-01T00:00:00+03:00"}`)

	msg, _, err := c.MessageSend(message)

	if err != nil {
		t.Errorf("%v", err)
	}

	assert.NoError(t, err)
	t.Logf("%v", msg)
}

func TestMgClient_MessageSendOrder(t *testing.T) {
	c := client()

	message := MessageSendRequest{
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
				Price: &MessageOrderCost{
					Value:    1100,
					Currency: MsgCurrencyRub,
				},
			},
			Items: []MessageOrderItem{
				{
					Name: "iPhone 6",
					Url:  "https://example.com/product.html",
					Img:  "https://example.com/picture.png",
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
	}

	defer gock.Off()

	gock.New(mgURL).
		Post("/api/bot/v1/messages").
		JSON(message).
		Reply(200).
		BodyString(`{"message_id": 1, "time": "2018-01-01T00:00:00+03:00"}`)

	msg, _, err := c.MessageSend(message)

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

	message := MessageEditRequest{
		ID:      uint64(1),
		Content: "test",
	}

	defer gock.Off()

	gock.New(mgURL).
		Patch("/api/bot/v1/messages/1").
		JSON(message).
		Reply(200).
		BodyString(`{"message_id": 1, "time": "2018-01-01T00:00:00+03:00"}`)

	e, status, err := c.MessageEdit(message)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	t.Logf("Message edit: %v", e)
}

func TestMgClient_MessageDelete(t *testing.T) {
	c := client()

	defer gock.Off()

	gock.New(mgURL).
		Delete("/api/bot/v1/messages/1").
		Reply(200).
		BodyString(`{}`)

	d, status, err := c.MessageDelete(1)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	t.Logf("Message delete: %v", d)
}

func TestMgClient_Info(t *testing.T) {
	c := client()
	req := InfoRequest{Name: "AWESOME", Avatar: "https://test.com/awesome_bot_avatar"}

	defer gock.Off()

	gock.New(mgURL).
		Patch("/api/bot/v1/my/info").
		JSON(req).
		Reply(200).
		BodyString(`{}`)

	_, status, err := c.Info(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
}

func TestMgClient_Commands(t *testing.T) {
	c := client()

	defer gock.Off()

	gock.New(mgURL).
		Get("/api/bot/v1/my/commands").
		Reply(200).
		BodyString(`[{"id": 1, "name": "command_name", "description": "Command description", "created_at": "2018-01-01T00:00:00.000000Z"}]`)

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
		Name:        "test_command",
		Description: "Test command",
	}

	defer gock.Off()

	gock.New(mgURL).
		Put("/api/bot/v1/my/commands/test_command").
		JSON(req).
		Reply(200).
		BodyString(`{"id": 1, "name": "test_command", "description": "Test description"}`)

	n, status, err := c.CommandEdit(req)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, n.ID)

	gock.New(mgURL).
		Delete("/api/bot/v1/my/commands/test_command").
		Reply(200).
		BodyString(`{}`)

	d, status, err := c.CommandDelete(n.Name)
	if err != nil {
		t.Errorf("%d %v", status, err)
	}

	assert.NoError(t, err)
	t.Logf("%v", d)
}

func TestMgClient_WsMeta_With_Options(t *testing.T) {
	c := client()
	events := []string{"user_updated", "user_join_chat"}
	params := []WsParams{WsOptionIncludeMassCommunication}

	url, headers, err := c.WsMeta(events, params...)

	if err != nil {
		t.Errorf("%v", err)
	}

	resURL := "wss://api.example.com/api/bot/v1/ws?events=user_updated,user_join_chat&options=include_mass_communication"
	resToken := c.Token

	assert.Equal(t, resURL, url)
	assert.Equal(t, resToken, headers["X-Bot-Token"][0])
}

func TestMgClient_WsMeta(t *testing.T) {
	c := client()
	events := []string{"user_updated", "user_join_chat"}
	url, headers, err := c.WsMeta(events)

	if err != nil {
		t.Errorf("%v", err)
	}

	resURL := fmt.Sprintf("%s%s%s%s", strings.Replace(c.URL, "https", "wss", 1), prefix, "/ws?events=", strings.Join(events[:], ","))
	resToken := c.Token

	assert.Equal(t, resURL, url)
	assert.Equal(t, resToken, headers["X-Bot-Token"][0])
}

func TestMgClient_UploadFile(t *testing.T) {
	c := client()

	resp, err := http.Get("https://via.placeholder.com/300")
	if err != nil {
		t.Errorf("%v", err)
	}

	defer resp.Body.Close()
	defer gock.Off()

	gock.New(mgURL).
		Post("/api/bot/v1/files/upload").
		Reply(200).
		BodyString(`{"created_at": "2018-01-01T00:00:00.000000Z", "hash": "hash", "id": "1"}`)

	data, status, err := c.UploadFile(resp.Body)

	if status != http.StatusOK {
		t.Errorf("%v", err)
	}

	t.Logf("File %+v is upload", data)
}

func TestMgClient_UploadFileByUrl(t *testing.T) {
	c := client()
	file := UploadFileByUrlRequest{
		Url: "https://via.placeholder.com/300",
	}

	defer gock.Off()

	gock.New(mgURL).
		Post("/api/bot/v1/files/upload_by_url").
		JSON(file).
		Reply(200).
		BodyString(`{"created_at": "2018-01-01T00:00:00.000000Z", "hash": "hash", "id": "1"}`)

	uploadFileResponse, st, err := c.UploadFileByURL(file)

	if st != http.StatusOK {
		t.Errorf("%v", err)
	}

	t.Logf("File %+v is upload", uploadFileResponse.ID)

	assert.NoError(t, err)
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

func TestMgClient_DebugNoLogger(t *testing.T) {
	c := client()
	c.Debug = true

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	c.writeLog("Test log string")

	assert.Contains(t, buf.String(), "Test log string")
}

func TestMgClient_DebugWithLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "Custom log prefix ", 0)

	c := client(OptionDebug(), OptionLogger(logger))

	c.writeLog("Test log string")

	assert.Contains(t, buf.String(), "Custom log prefix Test log string")
}

func TestMgClient_SuccessChatsByCustomerId(t *testing.T) {
	defer gock.Off()
	customerID := uint64(191140)
	gock.New(mgURL).
		Path("/api/bot/v1/chats").
		MatchParam("customer_id", fmt.Sprintf("%d", customerID)).
		MatchHeader("X-Bot-Token", mgToken).
		Reply(http.StatusOK).
		JSON(getJSONResponseChats())

	apiClient := client()
	chatsRequest := ChatsRequest{
		CustomerID: customerID,
	}

	resp, statusCode, err := apiClient.Chats(chatsRequest)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, 1, len(resp))
	assert.Equal(t, uint64(9000), resp[0].ID)
	assert.Equal(t, uint64(8000), resp[0].Channel.ID)
	assert.Equal(t, customerID, resp[0].Customer.ID)
	assert.Equal(t, "Имя Фамилия", resp[0].Customer.Name)
	assert.Equal(t, "Имя", resp[0].Customer.FirstName)
	assert.Equal(t, "Фамилия", resp[0].Customer.LastName)
}

func getJSONResponseChats() string {
	return `[
		{
			"id": 9000,
			"channel": {
				"id": 8000,
				"avatar": "",
				"transport_id": 555,
				"type": "transport",
				"settings": {
					"status": {
						"delivered": "send"
					},
					"text": {
						"creating": "both",
						"editing": "both",
						"quoting": "both",
						"deleting": "receive",
						"max_chars_count": 4096
					},
					"product": {
						"creating": "receive",
						"editing": "receive"
					},
					"order": {
						"creating": "receive",
						"editing": "receive"
					},
					"image": {
						"creating": "both",
						"editing": "both",
						"quoting": "both",
						"deleting": "receive",
						"max_items_count": 10
					},
					"file": {
						"creating": "both",
						"editing": "both",
						"quoting": "both",
						"deleting": "receive",
						"max_items_count": 1
					},
					"audio": {
						"creating": "both",
						"quoting": "both",
						"deleting": "receive",
						"max_items_count": 1
					},
					"suggestions": {
						"text": "both",
						"email": "both",
						"phone": "both"
					}
				},
				"name": "@test_bot123",
				"is_active": false
			},
			"customer": {
				"id": 191140,
				"external_id": "",
				"type": "customer",
				"avatar": "",
				"name": "Имя Фамилия",
				"username": "Имя",
				"first_name": "Имя",
				"last_name": "Фамилия"
			},
			"last_activity": "2022-10-28T13:17:38+03:00",
			"created_at": "2022-10-07T14:00:24.795382Z",
			"updated_at": "2022-10-28T12:19:04.834592Z"
		}
	]`
}
