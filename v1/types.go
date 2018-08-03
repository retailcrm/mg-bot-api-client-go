package v1

import (
	"net/http"
)

// MgClient type
type MgClient struct {
	URL        string `json:"url"`
	Token      string `json:"token"`
	Debug      bool   `json:"debug"`
	httpClient *http.Client
}

type BotsRequest struct {
	ID     uint64
	Self   string
	Active string
	Since  string
	Until  string
}

type ChannelsRequest struct {
	ID     uint64
	Types  string
	Active string
	Since  string
	Until  string
}

type ManagersRequest struct {
	ID         uint64
	ExternalID string `json:"external_id"`
	Online     string
	Active     string
	Since      string
	Until      string
}

type CustomersRequest struct {
	ID         uint64
	ExternalID string `json:"external_id"`
	Since      string
	Until      string
}

type ChatsRequest struct {
	ID          uint64
	ChannelID   string `json:"channel_id"`
	ChannelType string `json:"channel_type"`
	Since       string
	Until       string
}

type MembersRequest struct {
	ChatID     string `json:"chat_id"`
	ManagerID  string `json:"manager_id"`
	CustomerID string `json:"customer_id"`
	Status     string
	Since      string
	Until      string
}

type DialogsRequest struct {
	ID        uint64
	ChatID    string `json:"chat_id"`
	ManagerID string `json:"manager_id"`
	BotID     string `json:"bot_id"`
	Active    string
	Assigned  string
	Since     string
	Until     string
}

type MessagesRequest struct {
	ID          uint64
	ChatID      string `json:"chat_id"`
	DialogID    string `json:"dialog_id"`
	ManagerID   string `json:"manager_id"`
	CustomerID  string `json:"customer_id"`
	BotID       string `json:"bot_id"`
	ChannelID   string `json:"channel_id"`
	ChannelType string `json:"channel_type"`
	Scope       string
	Since       string
	Until       string
}

type MessageSendRequest struct {
	Content        string `json:"content"`
	Scope          uint8  `json:"scope"`
	ChatID         uint64 `json:"chat_id"`
	QuoteMessageId uint64 `json:"omitempty,quote_message_id"`
}

type MessageEditRequest struct {
	ID      uint64
	Content string `json:"content"`
}

type CommandsRequest struct {
	ID    string
	Name  string
	Since string
	Until string
}

type BotsResponse struct {
	Bots []BotListItem
}

type BotListItem struct {
	ID            uint64   `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Events        []string `json:"events,omitempty,brackets"`
	ClientID      string   `json:"client_id"`
	AvatarUrl     string   `json:"avatar_url"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     *string  `json:"updated_at"`
	DeactivatedAt *string  `json:"deactivated_at"`
	IsActive      bool     `json:"is_active"`
	IsSelf        bool     `json:"is_self"`
}

type ChannelsResponse struct {
	Channels []ChannelListItem
}

type ChannelListItem struct {
	ID            uint64   `json:"id"`
	Type          string   `json:"type"`
	Events        []string `json:"events,omitempty,brackets"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     *string  `json:"updated_at"`
	ActivatedAt   string   `json:"activated_at"`
	DeactivatedAt *string  `json:"deactivated_at"`
	IsActive      bool     `json:"is_active"`
}

type ManagersResponse struct {
	Managers []ManagersListItem
}

type ManagersListItem struct {
	ID         uint64  `json:"id"`
	ExternalID *string `json:"external_id,omitempty"`
	Username   *string `json:"username,omitempty"`
	FirstName  *string `json:"first_name,omitempty"`
	LastName   *string `json:"last_name,omitempty"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  *string `json:"updated_at,omitempty"`
	RevokedAt  *string `json:"revoked_at,omitempty"`
	IsOnline   bool    `json:"is_online"`
	IsActive   bool    `json:"is_active"`
	Avatar     *string `json:"avatar_url,omitempty"`
}

type CustomersResponse struct {
	Customers []CustomersListItem
}

type CustomersListItem struct {
	ID         uint64  `json:"id"`
	ExternalID *string `json:"external_id,omitempty"`
	ChannelId  *uint64 `json:"channel_id,omitempty"`
	Username   *string `json:"username,omitempty"`
	FirstName  *string `json:"first_name,omitempty"`
	LastName   *string `json:"last_name,omitempty"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  *string `json:"updated_at,omitempty"`
	RevokedAt  *string `json:"revoked_at,omitempty"`
	Avatar     *string `json:"avatar_url,omitempty"`
	ProfileURL *string `json:"profile_url,omitempty"`
	Country    *string `json:"country,omitempty"`
	Language   *string `json:"language,omitempty"`
	Phone      *string `json:"phone,omitempty"`
}

type ChatsResponse struct {
	Chats []ChatListItem
}

type ChatListItem struct {
	ID        uint64  `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Channel   *string `json:"channel,omitempty"`
	ChannelId *uint64 `json:"channel_id,omitempty"`
}

type MembersResponse struct {
	Members []MemberListItem
}

type MemberListItem struct {
	ID        uint64  `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	IsAuthor  bool    `json:"is_author"`
	State     string  `json:"state"`
	ChatID    uint64  `json:"chat_id"`
	UserID    uint64  `json:"user_id"`
}

type DialogsResponse struct {
	Dialogs []DialogListItem
}

type DialogListItem struct {
	ID              uint64       `json:"id"`
	ChatID          uint64       `json:"chat_id"`
	BotID           *uint64      `json:"bot_id,omitempty"`
	BeginMessageID  *uint64      `json:"begin_message_id,omitempty"`
	EndingMessageID *uint64      `json:"ending_message_id,omitempty"`
	CreatedAt       string       `json:"created_at"`
	UpdatedAt       *string      `json:"updated_at,omitempty"`
	ClosedAt        *string      `json:"closed_at,omitempty"`
	IsAssign        bool         `json:"is_assign"`
	Responsible     *Responsible `json:"responsible,omitempty"`
	IsActive        bool         `json:"is_active"`
}

type MessagesResponse struct {
	Messages []MessagesListItem
}

type MessagesListItem struct {
	ID             uint64  `json:"id"`
	Content        string  `json:"content"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
	Scope          uint8   `json:"scope"`
	ChatID         uint64  `json:"chat_id"`
	Sender         Sender  `json:"sender"`
	ChannelID      *uint64 `json:"channel_id,omitempty"`
	ChannelSentAt  *string `json:"channel_sent_at,omitempty"`
	QuoteMessageId *uint64 `json:"quote_message_id,omitempty"`
	EditedAt       *string `json:"edited_at,omitempty"`
	DeletedAt      *string `json:"deleted_at,omitempty"`
}

type MessageResponse struct {
	ID   uint64 `json:"id"`
	Time string `json:"content"`
}

type UpdateBotRequest struct {
	Name   string   `json:"name,omitempty"`
	Avatar *string  `json:"avatar_url"`
	Events []string `json:"events,omitempty,brackets"`
}

type Responsible struct {
	Type     string `json:"type"`
	ID       int64  `json:"id"`
	AssignAt string `json:"assign_at"`
}

type DialogResponsibleRequest struct {
	ManagerID int64 `json:"manager_id"`
	BotID     int64 `json:"bot_id"`
}

type AssignResponse struct {
	Responsible         Responsible  `json:"responsible"`
	IsReAssign          bool         `json:"is_reassign"`
	PreviousResponsible *Responsible `json:"previous_responsible,omitempty"`
	LeftManagerID       *uint64      `json:"left_manager_id,omitempty"`
}

type Sender struct {
	ID   int64
	Type string
}
