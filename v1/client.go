package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// New initialize client
func New(url string, token string) *MgClient {
	return &MgClient{
		URL:        url,
		Token:      token,
		httpClient: &http.Client{Timeout: 20 * time.Second},
	}
}

// Bots get all available bots
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	data, status, err := client.Bots()
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
//
// 	for _, bot := range data {
// 		fmt.Printf("%v %v\n", bot.Name, bot.CreatedAt)
// 	}
func (c *MgClient) Bots(request BotsRequest) ([]BotsResponseItem, int, error) {
	var resp []BotsResponseItem
	var b []byte
	outgoing, _ := query.Values(request)

	data, status, err := c.GetRequest(fmt.Sprintf("/bots?%s", outgoing.Encode()), b)
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// Channels get all available channels
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	data, status, err := client.Channels()
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
//
// 	for _, channel := range data {
// 		fmt.Printf("%v %v\n", channel.Type, channel.CreatedAt)
// 	}
func (c *MgClient) Channels(request ChannelsRequest) ([]ChannelResponseItem, int, error) {
	var resp []ChannelResponseItem
	var b []byte
	outgoing, _ := query.Values(request)

	data, status, err := c.GetRequest(fmt.Sprintf("/channels?%s", outgoing.Encode()), b)
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// Users get all available users
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	data, status, err := client.Users(UsersRequest:{Active:1})
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
//
// 	for _, user := range data {
// 		fmt.Printf("%v %v\n", user.FirstName, user.IsOnline)
// 	}
func (c *MgClient) Users(request UsersRequest) ([]UsersResponseItem, int, error) {
	var resp []UsersResponseItem
	var b []byte
	outgoing, _ := query.Values(request)

	data, status, err := c.GetRequest(fmt.Sprintf("/users?%s", outgoing.Encode()), b)
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// Customers get all available customers
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	data, status, err := client.Customers()
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
//
// 	for _, customer := range data {
// 		fmt.Printf("%v %v\n", customer.FirstName, customer.Avatar)
// 	}
func (c *MgClient) Customers(request CustomersRequest) ([]CustomersResponseItem, int, error) {
	var resp []CustomersResponseItem
	var b []byte
	outgoing, _ := query.Values(request)

	data, status, err := c.GetRequest(fmt.Sprintf("/customers?%s", outgoing.Encode()), b)
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// Chats get all available chats
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	data, status, err := client.Chats(ChatsRequest{ChannelType:ChannelTypeWhatsapp})
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
//
// 	for _, chat := range data {
// 		fmt.Printf("%v %v\n", chat.Customer, chat.LastMessage)
// 	}
func (c *MgClient) Chats(request ChatsRequest) ([]ChatResponseItem, int, error) {
	var resp []ChatResponseItem
	var b []byte
	outgoing, _ := query.Values(request)

	data, status, err := c.GetRequest(fmt.Sprintf("/chats?%s", outgoing.Encode()), b)
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// Members get all available chat members
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	data, status, err := client.Members(MembersRequest{State:ChatMemberStateActive})
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
//
// 	for _, member := range data {
// 		fmt.Printf("%v\n", member.CreatedAt)
// 	}
func (c *MgClient) Members(request MembersRequest) ([]MemberResponseItem, int, error) {
	var resp []MemberResponseItem
	var b []byte
	outgoing, _ := query.Values(request)

	data, status, err := c.GetRequest(fmt.Sprintf("/members?%s", outgoing.Encode()), b)
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// Dialogs get all available dialogs
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	data, status, err := client.Dialogs(DialogsRequest{Active:1})
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
//
// 	for _, dialog := range data {
// 		fmt.Printf("%v %v\n", dialog.ChatID, dialog.CreatedAt)
// 	}
func (c *MgClient) Dialogs(request DialogsRequest) ([]DialogResponseItem, int, error) {
	var resp []DialogResponseItem
	var b []byte
	outgoing, _ := query.Values(request)

	data, status, err := c.GetRequest(fmt.Sprintf("/dialogs?%s", outgoing.Encode()), b)
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

func (c *MgClient) DialogAssign(request DialogAssignRequest) (DialogAssignResponse, int, error) {
	var resp DialogAssignResponse
	outgoing, _ := json.Marshal(&request)

	data, status, err := c.PatchRequest(fmt.Sprintf("/dialogs/%d/assign", request.DialogID), []byte(outgoing))
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// DialogClose close selected dialog
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	_, status, err := client.DialogClose(123)
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
func (c *MgClient) DialogClose(request uint64) (map[string]interface{}, int, error) {
	var resp map[string]interface{}
	outgoing, _ := json.Marshal(&request)

	data, status, err := c.DeleteRequest(fmt.Sprintf("/dialogs/%d/close", request), []byte(outgoing))
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// Messages get all available messages
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	data, status, err := client.Messages(MessagesRequest{ManagerID:5})
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
//
// 	for _, message := range data {
// 		fmt.Printf("%v %v %v\n", message.ChatID, message.CreatedAt, message.CustomerID)
// 	}
func (c *MgClient) Messages(request MessagesRequest) ([]MessagesResponseItem, int, error) {
	var resp []MessagesResponseItem
	var b []byte
	outgoing, _ := query.Values(request)

	data, status, err := c.GetRequest(fmt.Sprintf("/messages?%s", outgoing.Encode()), b)
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// MessageSend send message
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	data, status, err := client.MessageSend(MessageSendRequest{
// 		Scope:   "public",
// 		Content: "test",
// 		ChatID:  i,
// 	})
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
//
// fmt.Printf("%v \n", data.MessageID, data.Time)
func (c *MgClient) MessageSend(request MessageSendRequest) (MessageSendResponse, int, error) {
	var resp MessageSendResponse
	outgoing, _ := json.Marshal(&request)

	data, status, err := c.PostRequest("/messages", []byte(outgoing))
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// MessageEdit update selected message
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	_, status, err := client.MessageEdit(MessageEditRequest{
// 		ID:      123,
// 		Content: "test",
// 	})
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
func (c *MgClient) MessageEdit(request MessageEditRequest) (map[string]interface{}, int, error) {
	var resp map[string]interface{}
	outgoing, _ := json.Marshal(&request)

	data, status, err := c.PatchRequest(fmt.Sprintf("/messages/%d", request.ID), []byte(outgoing))
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// MessageDelete delete selected message
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	_, status, err := client.MessageDelete(123)
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
func (c *MgClient) MessageDelete(request uint64) (map[string]interface{}, int, error) {
	var resp map[string]interface{}
	outgoing, _ := json.Marshal(&request)

	data, status, err := c.DeleteRequest(fmt.Sprintf("/messages/%d", request), []byte(outgoing))
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// Info updates bot information
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	_, status, err := client.Info(InfoRequest{Name: "AWESOME", Avatar: "https://example.com/logo.svg"})
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
func (c *MgClient) Info(request InfoRequest) (map[string]interface{}, int, error) {
	var resp map[string]interface{}
	outgoing, _ := json.Marshal(&request)

	data, status, err := c.PatchRequest("/my/info", []byte(outgoing))
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// Commands get all available commands for bot
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	data, status, err := client.Commands()
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
//
// 	for _, command := range data {
// 		fmt.Printf("%v %v\n", command.Name, command.Description)
// 	}
func (c *MgClient) Commands(request CommandsRequest) ([]CommandsResponseItem, int, error) {
	var resp []CommandsResponseItem
	var b []byte
	outgoing, _ := query.Values(request)

	data, status, err := c.GetRequest(fmt.Sprintf("/my/commands?%s", outgoing.Encode()), b)
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// CommandEdit create or change command for bot
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	data, status, err := client.CommandEdit(CommandEditRequest{
// 		BotID:       1,
// 		Name:        "show_payment_types",
// 		Description: "Get available payment types",
// 	})
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
//
// 	fmt.Printf("%v %v\n", data.Name, data.Description)
func (c *MgClient) CommandEdit(request CommandEditRequest) (CommandsResponseItem, int, error) {
	var resp CommandsResponseItem
	outgoing, _ := json.Marshal(&request)

	data, status, err := c.PutRequest(fmt.Sprintf("/my/commands/%s", request.Name), []byte(outgoing))
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

// CommandDelete delete selected command for bot
//
// Example:
//
// 	var client = v1.New("https://demo.url", "09jIJ")
//
// 	_, status, err := client.CommandDelete(show_payment_types)
//
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	if status >= http.StatusBadRequest {
// 		fmt.Printf("%v", err)
// 	}
func (c *MgClient) CommandDelete(request string) (map[string]interface{}, int, error) {
	var resp map[string]interface{}
	outgoing, _ := json.Marshal(&request)

	data, status, err := c.DeleteRequest(fmt.Sprintf("/my/commands/%s", request), []byte(outgoing))
	if err != nil {
		return resp, status, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, status, err
	}

	if status > http.StatusCreated || status < http.StatusOK {
		return resp, status, c.Error(data)
	}

	return resp, status, err
}

func (c *MgClient) Error(info []byte) error {
	var data map[string]interface{}

	if err := json.Unmarshal(info, &data); err != nil {
		return err
	}

	values := data["errors"].([]interface{})

	return errors.New(values[0].(string))
}
