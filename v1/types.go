package v1

import (
	"net/http"
)

const (
	ChannelTypeTelegram   string = "telegram"
	ChannelTypeFacebook   string = "fbmessenger"
	ChannelTypeViber      string = "viber"
	ChannelTypeWhatsapp   string = "whatsapp"
	ChannelTypeSkype      string = "skype"
	ChannelTypeVk         string = "vk"
	ChannelTypeInstagram  string = "instagram"
	ChannelTypeConsultant string = "consultant"
	ChannelTypeCustom     string = "custom"

	ChatMemberStateActive string = "active"
	ChatMemberStateKicked string = "kicked"
	ChatMemberStateLeaved string = "leaved"

	MessageScopePublic  string = "public"
	MessageScopePrivate string = "private"

	ChannelFeatureNone    string = "none"
	ChannelFeatureReceive string = "receive"
	ChannelFeatureSend    string = "send"
	ChannelFeatureBoth    string = "both"
)

// MgClient type
type MgClient struct {
	URL        string `json:"url"`
	Token      string `json:"token"`
	Debug      bool   `json:"debug"`
	httpClient *http.Client
}

// Request types
type (
	BotsRequest struct {
		ID     uint64 `url:"id,omitempty"`
		Active uint8  `url:"active,omitempty"`
		Since  string `url:"since,omitempty"`
		Until  string `url:"until,omitempty"`
	}

	ChannelsRequest struct {
		ID     uint64 `url:"id,omitempty"`
		Types  string `url:"types,omitempty"`
		Active uint8  `url:"active,omitempty"`
		Since  string `url:"since,omitempty"`
		Until  string `url:"until,omitempty"`
	}

	UsersRequest struct {
		ID         uint64 `url:"id,omitempty"`
		ExternalID string `url:"external_id,omitempty" json:"external_id"`
		Online     uint8  `url:"online,omitempty"`
		Active     uint8  `url:"active,omitempty"`
		Since      string `url:"since,omitempty"`
		Until      string `url:"until,omitempty"`
	}

	CustomersRequest struct {
		ID         uint64 `url:"id,omitempty"`
		ExternalID string `url:"external_id,omitempty" json:"external_id"`
		Since      string `url:"since,omitempty"`
		Until      string `url:"until,omitempty"`
	}

	ChatsRequest struct {
		ID          uint64 `url:"id,omitempty"`
		ChannelID   uint64 `url:"channel_id,omitempty" json:"channel_id"`
		ChannelType string `url:"channel_type,omitempty" json:"channel_type"`
		Since       string `url:"since,omitempty"`
		Until       string `url:"until,omitempty"`
	}

	MembersRequest struct {
		ChatID uint64 `url:"chat_id,omitempty" json:"chat_id"`
		UserID string `url:"user_id,omitempty" json:"user_id"`
		State  string `url:"state,omitempty"`
		Since  string `url:"since,omitempty"`
		Until  string `url:"until,omitempty"`
	}

	DialogsRequest struct {
		ID        uint64 `url:"id,omitempty"`
		ChatID    string `url:"chat_id,omitempty" json:"chat_id"`
		ManagerID string `url:"manager_id,omitempty" json:"manager_id"`
		BotID     string `url:"bot_id,omitempty" json:"bot_id"`
		Assign    uint8  `url:"assign,omitempty"`
		Active    uint8  `url:"active,omitempty"`
		Since     string `url:"since,omitempty"`
		Until     string `url:"until,omitempty"`
	}

	DialogAssignRequest struct {
		DialogID  uint64 `url:"dialog_id,omitempty" json:"dialog_id"`
		ManagerID uint64 `url:"manager_id,omitempty" json:"manager_id"`
		BotID     uint64 `url:"bot_id,omitempty" json:"bot_id"`
	}

	MessagesRequest struct {
		ID          uint64 `url:"id,omitempty"`
		ChatID      uint64 `url:"chat_id,omitempty" json:"chat_id"`
		DialogID    uint64 `url:"dialog_id,omitempty" json:"dialog_id"`
		ManagerID   uint64 `url:"manager_id,omitempty" json:"manager_id"`
		CustomerID  uint64 `url:"customer_id,omitempty" json:"customer_id"`
		BotID       uint64 `url:"bot_id,omitempty" json:"bot_id"`
		ChannelID   uint64 `url:"channel_id,omitempty" json:"channel_id"`
		ChannelType string `url:"channel_type,omitempty" json:"channel_type"`
		Scope       string `url:"scope,omitempty"`
		Type        string `url:"type,omitempty"`
		Since       string `url:"since,omitempty"`
		Until       string `url:"until,omitempty"`
	}

	MessageSendRequest struct {
		Content        string `url:"content,omitempty" json:"content"`
		Scope          string `url:"scope,omitempty" json:"scope"`
		ChatID         uint64 `url:"chat_id,omitempty" json:"chat_id"`
		QuoteMessageId uint64 `url:"quote_message_id,omitempty" json:"quote_message_id"`
	}

	MessageEditRequest struct {
		ID      uint64 `url:"id,omitempty"`
		Content string `url:"content,omitempty" json:"content"`
	}

	InfoRequest struct {
		Name   string `url:"name,omitempty" json:"name"`
		Avatar string `url:"avatar_url,omitempty" json:"avatar_url"`
	}

	CommandsRequest struct {
		ID    uint64 `url:"id,omitempty"`
		Name  string `url:"name,omitempty"`
		Since string `url:"since,omitempty"`
		Until string `url:"until,omitempty"`
	}

	CommandEditRequest struct {
		Name        string `url:"name,omitempty" json:"name"`
		Description string `url:"description,omitempty" json:"description"`
	}
)

// Response types
type (
	BotsResponseItem struct {
		ID            uint64 `json:"id"`
		Name          string `json:"name"`
		ClientID      string `json:"client_id,omitempty"`
		AvatarUrl     string `json:"avatar_url,omitempty"`
		CreatedAt     string `json:"created_at,omitempty"`
		UpdatedAt     string `json:"updated_at,omitempty"`
		DeactivatedAt string `json:"deactivated_at,omitempty"`
		IsActive      bool   `json:"is_active"`
		IsSelf        bool   `json:"is_self"`
	}

	ChannelResponseItem struct {
		ID            uint64          `json:"id"`
		Type          string          `json:"type"`
		Settings      ChannelSettings `json:"settings"`
		CreatedAt     string          `json:"created_at"`
		UpdatedAt     string          `json:"updated_at"`
		ActivatedAt   string          `json:"activated_at"`
		DeactivatedAt string          `json:"deactivated_at"`
		IsActive      bool            `json:"is_active"`
	}

	UsersResponseItem struct {
		ID         uint64 `json:"id"`
		ExternalID string `json:"external_id,omitempty"`
		Username   string `json:"username,omitempty"`
		FirstName  string `json:"first_name,omitempty"`
		LastName   string `json:"last_name,omitempty"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at,omitempty"`
		RevokedAt  string `json:"revoked_at,omitempty"`
		IsOnline   bool   `json:"is_online"`
		IsActive   bool   `json:"is_active"`
		Avatar     string `json:"avatar_url,omitempty"`
	}

	CustomersResponseItem struct {
		ID         uint64 `json:"id"`
		ExternalID string `json:"external_id,omitempty"`
		ChannelId  uint64 `json:"channel_id,omitempty"`
		Username   string `json:"username,omitempty"`
		FirstName  string `json:"first_name,omitempty"`
		LastName   string `json:"last_name,omitempty"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at,omitempty"`
		RevokedAt  string `json:"revoked_at,omitempty"`
		Avatar     string `json:"avatar_url,omitempty"`
		ProfileURL string `json:"profile_url,omitempty"`
		Country    string `json:"country,omitempty"`
		Language   string `json:"language,omitempty"`
		Phone      string `json:"phone,omitempty"`
		Email      string `json:"email,omitempty"`
	}

	ChatResponseItem struct {
		ID           uint64  `json:"id"`
		Avatar       string  `json:"avatar"`
		Name         string  `json:"name"`
		Channel      Channel `json:"channel,omitempty"`
		Customer     UserRef `json:"customer"`
		AuthorID     uint64  `json:"author_id"`
		LastMessage  Message `json:"last_message"`
		LastActivity string  `json:"last_activity"`
		CreatedAt    string  `json:"created_at"`
		UpdatedAt    string  `json:"updated_at"`
	}

	MemberResponseItem struct {
		ID        uint64 `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at,omitempty"`
		IsAuthor  bool   `json:"is_author"`
		State     string `json:"state"`
		ChatID    uint64 `json:"chat_id"`
		UserID    uint64 `json:"user_id"`
	}

	DialogResponseItem struct {
		ID              uint64      `json:"id"`
		ChatID          uint64      `json:"chat_id"`
		BeginMessageID  uint64      `json:"begin_message_id,omitempty"`
		EndingMessageID uint64      `json:"ending_message_id,omitempty"`
		BotID           uint64      `json:"bot_id,omitempty"`
		CreatedAt       string      `json:"created_at"`
		UpdatedAt       string      `json:"updated_at,omitempty"`
		ClosedAt        string      `json:"closed_at,omitempty"`
		IsAssigned      bool        `json:"is_assigned"`
		Responsible     Responsible `json:"responsible,omitempty"`
		IsActive        bool        `json:"is_active"`
	}

	DialogAssignResponse struct {
		Responsible         Responsible `json:"responsible"`
		PreviousResponsible Responsible `json:"previous_responsible,omitempty"`
		LeftManagerID       uint64      `json:"left_manager_id,omitempty"`
		IsReAssign          bool        `json:"is_reassign"`
	}

	MessagesResponseItem struct {
		Message
		ChannelID     uint64 `json:"channel_id,omitempty"`
		ChannelSentAt string `json:"channel_sent_at,omitempty"`
		CreatedAt     string `json:"created_at"`
		UpdatedAt     string `json:"updated_at"`
	}

	MessageSendResponse struct {
		MessageID uint64 `json:"message_id"`
		Time      string `json:"time"`
	}

	CommandsResponseItem struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at,omitempty"`
	}
)

// Single entity types
type (
	Message struct {
		ID     uint64 `json:"id"`
		Time   string `json:"time"`
		Type   string `json:"type"`
		ChatID uint64 `json:"chat_id"`
		IsRead bool   `json:"is_read"`
		Status string `json:"status"`
		TextMessage
		SystemMessage
	}

	TextMessage struct {
		Scope   string       `json:"scope"`
		Content string       `json:"content"`
		From    UserRef      `json:"from"`
		Quote   QuoteMessage `json:"quote"`
		IsEdit  bool         `json:"is_edit"`
		Actions []string     `json:"actions"`
	}

	SystemMessage struct {
		Action string              `json:"action"`
		Dialog SystemMessageDialog `json:"dialog,omitempty"`
		User   UserRef             `json:"user,omitempty"`
	}

	SystemMessageDialog struct {
		ID uint64 `json:"id"`
	}

	QuoteMessage struct {
		ID      uint64  `json:"id"`
		Content string  `json:"content"`
		Time    string  `json:"time"`
		From    UserRef `json:"from"`
	}

	UserRef struct {
		ID     uint64 `json:"id"`
		Avatar string `json:"avatar"`
		Type   string `json:"type"`
		Name   string `json:"name"`
		Phone  string `json:"phone,omitempty"`
		Email  string `json:"email,omitempty"`
	}

	Channel struct {
		ID          uint64          `json:"id"`
		TransportID uint64          `json:"transport_id"`
		Type        string          `json:"type"`
		Supports    ChannelSupports `json:"supports"`
	}

	ChannelSupports struct {
		Messages []string `json:"messages"`
		Statuses []string `json:"statuses"`
	}

	Responsible struct {
		ID       int64  `json:"id"`
		Type     string `json:"type"`
		AssignAt string `json:"assigned_at"`
	}

	Command struct {
		ID          uint64
		BotID       uint64
		Name        string
		Description string
		CreatedAt   string
		UpdatedAt   string
	}
)

// Channel settings
type (
	ChannelSettingsText struct {
		Creating string `json:"creating"`
		Editing  string `json:"editing"`
		Quoting  string `json:"quoting"`
		Deleting string `json:"deleting"`
	}

	ChannelSettings struct {
		SpamAllowed bool `json:"spam_allowed"`

		Status struct {
			Delivered string `json:"delivered"`
			Read      string `json:"read"`
		} `json:"status"`

		Text ChannelSettingsText `json:"text"`
	}
)
