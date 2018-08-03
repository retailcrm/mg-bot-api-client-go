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

func (c *MgClient) Bots(request BotsRequest) (BotsResponse, int, error) {
	var resp BotsResponse
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

func (c *MgClient) Channels(request ChannelsRequest) (ChannelsResponse, int, error) {
	var resp ChannelsResponse
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

func (c *MgClient) Managers(request ManagersRequest) (ManagersResponse, int, error) {
	var resp ManagersResponse
	var b []byte
	outgoing, _ := query.Values(request)

	data, status, err := c.GetRequest(fmt.Sprintf("/managers?%s", outgoing.Encode()), b)
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

func (c *MgClient) Customers(request CustomersRequest) (CustomersResponse, int, error) {
	var resp CustomersResponse
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

func (c *MgClient) Chats(request ChatsRequest) (ChatsResponse, int, error) {
	var resp ChatsResponse
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

func (c *MgClient) Members(request MembersRequest) (MembersResponse, int, error) {
	var resp MembersResponse
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

func (c *MgClient) Dialogs(request DialogsRequest) (DialogsResponse, int, error) {
	var resp DialogsResponse
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

	data, status, err := c.PostRequest(fmt.Sprintf("/dialogs/%s/assign", request.ID), []byte(outgoing))
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

func (c *MgClient) DialogClose(request DialogCloseRequest) (DialogCloseResponse, int, error) {
	var resp DialogCloseResponse
	outgoing, _ := json.Marshal(&request)

	data, status, err := c.PostRequest(fmt.Sprintf("/dialogs/%s/close", request.ID), []byte(outgoing))
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

func (c *MgClient) Messages(request MessagesRequest) (MessagesResponse, int, error) {
	var resp MessagesResponse
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

func (c *MgClient) MessageSend(request MessageSendRequest) (MessageResponse, int, error) {
	var resp MessageResponse
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

func (c *MgClient) MessageEdit(request MessageEditRequest) (MessageResponse, int, error) {
	var resp MessageResponse
	outgoing, _ := json.Marshal(&request)

	data, status, err := c.PatchRequest(fmt.Sprintf("/messages/%s", request.ID), []byte(outgoing))
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

func (c *MgClient) MessageDelete(request MessageDeleteRequest) (MessageResponse, int, error) {
	var resp MessageResponse
	outgoing, _ := json.Marshal(&request)

	data, status, err := c.DeleteRequest(fmt.Sprintf("/messages/%s", request.ID), []byte(outgoing))
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

func (c *MgClient) Info(request InfoRequest) (InfoResponse, int, error) {
	var resp InfoResponse
	outgoing, _ := json.Marshal(&request)

	data, status, err := c.PatchRequest("/messages/info", []byte(outgoing))
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

func (c *MgClient) Commands(request CommandsRequest) (CommandsResponse, int, error) {
	var resp CommandsResponse
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

func (c *MgClient) CommandEdit(request CommandEditRequest) (CommandEditResponse, int, error) {
	var resp CommandEditResponse
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

func (c *MgClient) CommandDelete(request CommandDeleteRequest) (CommandDeleteResponse, int, error) {
	var resp CommandDeleteResponse
	outgoing, _ := json.Marshal(&request)

	data, status, err := c.DeleteRequest(fmt.Sprintf("/my/commands/%s", request.Name), []byte(outgoing))
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
