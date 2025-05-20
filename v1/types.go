package v1

import (
	"encoding/json"
	"net/http"
	"time"
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

	WsEventMessageNew        string = "message_new"
	WsEventMessageUpdated    string = "message_updated"
	WsEventMessageDeleted    string = "message_deleted"
	WsEventDialogOpened      string = "dialog_opened"
	WsEventDialogClosed      string = "dialog_closed"
	WsEventDialogAssign      string = "dialog_assign"
	WsEventChatCreated       string = "chat_created"
	WsEventChatUpdated       string = "chat_updated"
	WsEventChatUnreadUpdated string = "chat_unread_updated"
	WsEventUserOnlineUpdated string = "user_online_updated"
	WsEventUserJoined        string = "user_joined_chat"
	WsEventUserLeave         string = "user_left_chat"
	WsEventUserUpdated       string = "user_updated"
	WsCustomerUpdated        string = "customer_updated"
	WsBotUpdated             string = "bot_updated"
	WsEventChannelUpdated    string = "channel_updated"
	WsEventSettingsUpdated   string = "settings_updated"
	WsEventChatsDeleted      string = "chats_deleted"

	WsOptionIncludeMassCommunication WsOption = "include_mass_communication"

	ChannelFeatureNone    string = "none"
	ChannelFeatureReceive string = "receive"
	ChannelFeatureSend    string = "send"
	ChannelFeatureBoth    string = "both"

	BotRoleDistributor string = "distributor"
	BotRoleResponsible string = "responsible"
	BotRoleHidden      string = "hidden"

	MsgTypeText    string = "text"
	MsgTypeSystem  string = "system"
	MsgTypeCommand string = "command"
	MsgTypeOrder   string = "order"
	MsgTypeProduct string = "product"
	MsgTypeFile    string = "file"
	MsgTypeImage   string = "image"

	MsgOrderStatusCodeNew        = "new"
	MsgOrderStatusCodeApproval   = "approval"
	MsgOrderStatusCodeAssembling = "assembling"
	MsgOrderStatusCodeDelivery   = "delivery"
	MsgOrderStatusCodeComplete   = "complete"
	MsgOrderStatusCodeCancel     = "cancel"

	MsgCurrencyRub = "rub"
	MsgCurrencyUah = "uah"
	MsgCurrencyByr = "byr"
	MsgCurrencyKzt = "kzt"
	MsgCurrencyUsd = "usd"
	MsgCurrencyEur = "eur"

	SuggestionTypeText  = "text"
	SuggestionTypeEmail = "email"
	SuggestionTypePhone = "phone"

	ColorLightRed         = "light-red"
	ColorLightBlue        = "light-blue"
	ColorLightGreen       = "light-green"
	ColorLightOrange      = "light-orange"
	ColorLightGray        = "light-gray"
	ColorLightGrayishBlue = "light-grayish-blue"
	ColorRed              = "red"
	ColorBlue             = "blue"
	ColorGreen            = "green"
	ColorOrange           = "orange"
	ColorGray             = "gray"
	ColorGrayishBlue      = "grayish-blue"

	WaitingLevelNone    = "none"
	WaitingLevelWarning = "warning"
	WaitingLevelDanger  = "danger"
)

// MgClient type
type MgClient struct {
	URL        string `json:"url"`
	Token      string `json:"token"`
	Debug      bool   `json:"debug"`
	httpClient *http.Client
	logger     BasicLogger `json:"-"`
}

// Request types
type (
	BotsRequest struct {
		ID      uint64 `url:"id,omitempty"`
		Active  uint8  `url:"active,omitempty"`
		Self    uint8  `url:"self,omitempty"`
		Role    string `url:"role,omitempty"`
		Since   string `url:"since,omitempty"`
		Until   string `url:"until,omitempty"`
		SinceID uint64 `url:"since_id,omitempty"`
		UntilID uint64 `url:"until_id,omitempty"`
		Limit   int    `url:"limit,omitempty"`
	}

	ChannelsRequest struct {
		ID      uint64   `url:"id,omitempty"`
		Types   []string `url:"types,omitempty"`
		Active  uint8    `url:"active,omitempty"`
		Since   string   `url:"since,omitempty"`
		Until   string   `url:"until,omitempty"`
		SinceID uint64   `url:"since_id,omitempty"`
		UntilID uint64   `url:"until_id,omitempty"`
		Limit   int      `url:"limit,omitempty"`
	}

	UsersRequest struct {
		ID         uint64 `url:"id,omitempty"`
		ExternalID string `url:"external_id,omitempty" json:"external_id"`
		Online     uint8  `url:"online,omitempty"`
		Active     uint8  `url:"active,omitempty"`
		Since      string `url:"since,omitempty"`
		Until      string `url:"until,omitempty"`
		SinceID    uint64 `url:"since_id,omitempty"`
		UntilID    uint64 `url:"until_id,omitempty"`
		Limit      int    `url:"limit,omitempty"`
	}

	CustomersRequest struct {
		ID          uint64 `url:"id,omitempty"`
		ChannelID   uint64 `url:"channel_id,omitempty" json:"channel_id"`
		ChannelType string `url:"channel_type,omitempty" json:"channel_type"`
		ExternalID  string `url:"external_id,omitempty" json:"external_id"`
		Since       string `url:"since,omitempty"`
		Until       string `url:"until,omitempty"`
		SinceID     uint64 `url:"since_id,omitempty"`
		UntilID     uint64 `url:"until_id,omitempty"`
		Limit       int    `url:"limit,omitempty"`
	}

	ChatsRequest struct {
		ID                       uint64 `url:"id,omitempty"`
		ChannelID                uint64 `url:"channel_id,omitempty" json:"channel_id"`
		ChannelType              string `url:"channel_type,omitempty" json:"channel_type"`
		CustomerID               uint64 `url:"customer_id,omitempty" json:"customer_id"`
		CustomerExternalID       string `url:"customer_external_id,omitempty" json:"customer_external_id"`
		Since                    string `url:"since,omitempty"`
		Until                    string `url:"until,omitempty"`
		Limit                    int    `url:"limit,omitempty"`
		SinceID                  int    `url:"since_id,omitempty"`
		UntilID                  int    `url:"until_id,omitempty"`
		IncludeMassCommunication uint8  `url:"include_mass_communication,omitempty"`
	}

	MembersRequest struct {
		ChatID  uint64 `url:"chat_id,omitempty" json:"chat_id"`
		UserID  string `url:"user_id,omitempty" json:"user_id"`
		State   string `url:"state,omitempty"`
		Since   string `url:"since,omitempty"`
		Until   string `url:"until,omitempty"`
		SinceID uint64 `url:"since_id,omitempty"`
		UntilID uint64 `url:"until_id,omitempty"`
		Limit   int    `url:"limit,omitempty"`
	}

	DialogsRequest struct {
		ID                       uint64 `url:"id,omitempty"`
		ChatID                   string `url:"chat_id,omitempty" json:"chat_id"`
		UserID                   string `url:"user_id,omitempty" json:"user_id"`
		BotID                    string `url:"bot_id,omitempty" json:"bot_id"`
		Assign                   uint8  `url:"assign,omitempty"`
		Active                   uint8  `url:"active,omitempty"`
		Since                    string `url:"since,omitempty"`
		SinceID                  int    `url:"since_id,omitempty"`
		Until                    string `url:"until,omitempty"`
		UntilID                  int    `url:"until_id,omitempty"`
		Limit                    int    `url:"limit,omitempty"`
		IncludeMassCommunication uint8  `url:"include_mass_communication,omitempty"`
	}

	DialogAssignRequest struct {
		DialogID uint64 `url:"dialog_id,omitempty" json:"dialog_id"`
		UserID   uint64 `url:"user_id,omitempty" json:"user_id"`
		BotID    uint64 `url:"bot_id,omitempty" json:"bot_id"`
	}

	DialogTagsAddRequest struct {
		DialogID uint64    `url:"dialog_id,omitempty"`
		Tags     []TagsAdd `json:"tags"`
	}

	TagsAdd struct {
		Name      string  `json:"name"`
		ColorCode *string `json:"color_code"`
	}

	DialogTagsDeleteRequest struct {
		DialogID uint64       `url:"dialog_id,omitempty"`
		Tags     []TagsDelete `json:"tags"`
	}

	TagsDelete struct {
		Name string `json:"name"`
	}

	MessagesRequest struct {
		ID                       []int  `url:"id,omitempty"`
		ChatID                   uint64 `url:"chat_id,omitempty" json:"chat_id"`
		DialogID                 uint64 `url:"dialog_id,omitempty" json:"dialog_id"`
		UserID                   uint64 `url:"user_id,omitempty" json:"user_id"`
		CustomerID               uint64 `url:"customer_id,omitempty" json:"customer_id"`
		BotID                    uint64 `url:"bot_id,omitempty" json:"bot_id"`
		ChannelID                uint64 `url:"channel_id,omitempty" json:"channel_id"`
		ChannelType              string `url:"channel_type,omitempty" json:"channel_type"`
		Scope                    string `url:"scope,omitempty"`
		Type                     string `url:"type,omitempty"`
		Since                    string `url:"since,omitempty"`
		Until                    string `url:"until,omitempty"`
		SinceID                  int    `url:"since_id,omitempty"`
		UntilID                  int    `url:"until_id,omitempty"`
		Limit                    int    `url:"limit,omitempty"`
		IncludeMassCommunication uint8  `url:"include_mass_communication,omitempty"`
	}

	MessageSendRequest struct {
		Type                 string                `url:"type,omitempty" json:"type"`
		Content              string                `url:"content,omitempty" json:"content"`
		Product              *MessageProduct       `url:"product,omitempty" json:"product"`
		Order                *MessageOrder         `url:"order,omitempty" json:"order"`
		Items                []Item                `url:"order,omitempty" json:"items"`
		Scope                string                `url:"scope,omitempty" json:"scope"`
		ChatID               uint64                `url:"chat_id,omitempty" json:"chat_id"`
		QuoteMessageId       uint64                `url:"quote_message_id,omitempty" json:"quote_message_id"`
		TransportAttachments *TransportAttachments `url:"transport_attachments,omitempty" json:"transport_attachments"`
	}

	TransportAttachments struct {
		Suggestions []Suggestion `url:"suggestions,omitempty" json:"suggestions"`
	}

	Suggestion struct {
		Type  string `url:"type,omitempty" json:"type"`
		Title string `url:"title,omitempty" json:"title"`
	}

	MessageEditRequest struct {
		ID      uint64 `url:"id,omitempty"`
		Content string `url:"content,omitempty" json:"content"`
	}

	InfoRequest struct {
		Name   string   `url:"name,omitempty" json:"name"`
		Avatar string   `url:"avatar_url,omitempty" json:"avatar_url,omitempty"`
		Roles  []string `url:"roles,omitempty" json:"roles"`
	}

	CommandsRequest struct {
		ID      uint64 `url:"id,omitempty"`
		Name    string `url:"name,omitempty"`
		Since   string `url:"since,omitempty"`
		Until   string `url:"until,omitempty"`
		SinceID uint64 `url:"since_id,omitempty"`
		UntilID uint64 `url:"until_id,omitempty"`
		Limit   int    `url:"limit,omitempty"`
	}

	CommandEditRequest struct {
		Name        string `url:"name,omitempty" json:"name"`
		Description string `url:"description,omitempty" json:"description"`
	}

	UploadFileByUrlRequest struct {
		Url string `json:"url"`
	}

	UpdateFileMetadataRequest struct {
		ID            string `json:"-"`
		Transcription string `json:"transcription,omitempty"`
		// Current status of transcription process. Available values: "in_progress", "ready", "error"
		TranscriptionStatus string `json:"transcription_status,omitempty"`
	}
)

// Response types
type (
	BotsResponseItem struct {
		ID            uint64   `json:"id"`
		Name          string   `json:"name"`
		ClientID      string   `json:"client_id,omitempty"`
		AvatarUrl     string   `json:"avatar_url,omitempty"`
		CreatedAt     string   `json:"created_at,omitempty"`
		UpdatedAt     string   `json:"updated_at,omitempty"`
		DeactivatedAt string   `json:"deactivated_at,omitempty"`
		IsActive      bool     `json:"is_active"`
		IsSelf        bool     `json:"is_self"`
		Roles         []string `json:"roles,omitempty"`
	}

	ChannelResponseItem struct {
		ID            uint64          `json:"id"`
		Type          string          `json:"type"`
		Name          string          `json:"name"`
		Settings      ChannelSettings `json:"settings"`
		CreatedAt     string          `json:"created_at"`
		UpdatedAt     string          `json:"updated_at"`
		ActivatedAt   string          `json:"activated_at"`
		DeactivatedAt string          `json:"deactivated_at"`
		IsActive      bool            `json:"is_active"`
	}

	UsersResponseItem struct {
		ID                 uint64 `json:"id"`
		ExternalID         string `json:"external_id,omitempty"`
		Username           string `json:"username,omitempty"`
		FirstName          string `json:"first_name,omitempty"`
		LastName           string `json:"last_name,omitempty"`
		CreatedAt          string `json:"created_at"`
		UpdatedAt          string `json:"updated_at,omitempty"`
		RevokedAt          string `json:"revoked_at,omitempty"`
		Available          bool   `json:"available"`
		IsOnline           bool   `json:"is_online"`
		Connected          bool   `json:"connected"`
		IsActive           bool   `json:"is_active"`
		IsTechnicalAccount bool   `json:"is_technical_account"`
		Avatar             string `json:"avatar_url,omitempty"`
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
		Utm        *Utm   `json:"utm,omitempty"`
	}

	ChatResponseItem struct {
		ID              uint64    `json:"id"`
		Avatar          string    `json:"avatar"`
		Name            string    `json:"name"`
		Channel         Channel   `json:"channel,omitempty"`
		Customer        UserRef   `json:"customer"`
		AuthorID        uint64    `json:"author_id"`
		LastMessage     Message   `json:"last_message"`
		LastUserMessage MessageID `json:"last_user_message"`
		LastActivity    string    `json:"last_activity"`
		LastDialog      Dialog    `json:"last_dialog"`
		CreatedAt       string    `json:"created_at"`
		UpdatedAt       string    `json:"updated_at"`
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
		Utm             *Utm        `json:"utm,omitempty"`
	}

	DialogAssignResponse struct {
		Responsible         Responsible `json:"responsible"`
		PreviousResponsible Responsible `json:"previous_responsible,omitempty"`
		LeftUserID          uint64      `json:"left_user_id,omitempty"`
		IsReAssign          bool        `json:"is_reassign"`
	}

	DialogUnassignResponse struct {
		PreviousResponsible Responsible `json:"previous_responsible,omitempty"`
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

	FullFileResponse struct {
		ID   string `json:"id,omitempty"`
		Type string `json:"type,omitempty"`
		Size int    `json:"size,omitempty"`
		Url  string `json:"url,omitempty"`
	}

	UploadFileResponse struct {
		ID        string    `json:"id"`
		Hash      string    `json:"hash"`
		Type      string    `json:"type"`
		Meta      FileMeta  `json:"meta"`
		MimeType  string    `json:"mime_type"`
		Size      int       `json:"size"`
		Url       *string   `json:"source_url"`
		CreatedAt time.Time `json:"created_at"`
	}
)

// WS event types
type (
	WsEvent struct {
		Type  string          `json:"type"`
		Meta  EventMeta       `json:"meta"`
		AppID uint            `json:"app_id"`
		Data  json.RawMessage `json:"data"`
	}

	EventMeta struct {
		Timestamp int64 `json:"timestamp"`
	}
)

// WS options.
type (
	WsOption string
)

// Single entity types
type (
	Message struct {
		ID      uint64          `json:"id"`
		Time    string          `json:"time"`
		Type    string          `json:"type"`
		Scope   string          `json:"scope"`
		ChatID  uint64          `json:"chat_id"`
		IsRead  bool            `json:"is_read"`
		IsEdit  bool            `json:"is_edit"`
		Status  string          `json:"status"`
		Chat    *Chat           `json:"chat"`
		From    *UserRef        `json:"from"`
		Product *MessageProduct `json:"product,omitempty"`
		Order   *MessageOrder   `json:"order,omitempty"`
		Dialog  *MessageDialog  `json:"dialog,omitempty"`
		*TextMessage
		*SystemMessage
		*AttachmentList
	}

	MessageID struct {
		ID uint64 `json:"id"`
	}

	TextMessage struct {
		Content string        `json:"content"`
		Quote   *QuoteMessage `json:"quote"`
		Actions []string      `json:"actions"`
	}

	SystemMessage struct {
		Action string `json:"action"`
		// Deprecated: Use Message.Dialog.ID instead.
		Dialog      *SystemMessageDialog `json:"dialog,omitempty"`
		User        *UserRef             `json:"user,omitempty"`
		Responsible *UserRef             `json:"responsible,omitempty"`
	}

	// Deprecated: Use MessageDialog instead.
	SystemMessageDialog struct {
		ID uint64 `json:"id"`
	}

	MessageDialog struct {
		ID uint64 `json:"id"`
	}

	QuoteMessage struct {
		ID      uint64   `json:"id"`
		Content string   `json:"content"`
		Time    string   `json:"time"`
		From    *UserRef `json:"from"`
	}

	AttachmentList struct {
		Items []Attachment `json:"items"`
		Note  string       `json:"note"`
	}

	Attachment struct {
		File

		Caption string `json:"caption"`
	}

	File struct {
		PreviewURL    *string `json:"preview_url,omitempty"`
		Height        *uint64 `json:"height,omitempty"`
		Width         *uint64 `json:"width,omitempty"`
		Transcription string  `json:"transcription,omitempty"`
		ID            string  `json:"id"`
		Mime          string  `json:"type"`
		Type          string  `json:"kind"`
		Duration      int     `json:"duration,omitempty"`
		Size          uint64  `json:"size"`
	}

	MessageProduct struct {
		ID       uint64                `json:"id"`
		Name     string                `json:"name"`
		Article  string                `json:"article,omitempty"`
		Url      string                `json:"url,omitempty"`
		Img      string                `json:"img,omitempty"`
		Cost     *MessageOrderCost     `json:"cost,omitempty"`
		Quantity *MessageOrderQuantity `json:"quantity,omitempty"`
	}

	MessageOrder struct {
		Number   string                `json:"number"`
		Url      string                `json:"url,omitempty"`
		Date     string                `json:"date,omitempty"`
		Cost     *MessageOrderCost     `json:"cost,omitempty"`
		Status   *MessageOrderStatus   `json:"status,omitempty"`
		Delivery *MessageOrderDelivery `json:"delivery"`
		Payments []MessageOrderPayment `json:"payment"`
		Items    []MessageOrderItem    `json:"items,omitempty"`
	}

	MessageOrderStatus struct {
		Code string `json:"code,omitempty"`
		Name string `json:"name,omitempty"`
	}

	MessageOrderItem struct {
		Name     string                `json:"name,omitempty"`
		Url      string                `json:"url,omitempty"`
		Img      string                `json:"img,omitempty"`
		Quantity *MessageOrderQuantity `json:"quantity,omitempty"`
		Price    *MessageOrderCost     `json:"price,omitempty"`
	}

	MessageOrderCost struct {
		Value    float32 `json:"value,omitempty"`
		Currency string  `json:"currency"`
	}

	MessageOrderQuantity struct {
		Value float32 `json:"value"`
		Unit  string  `json:"unit"`
	}

	MessageOrderPayment struct {
		Name   string                     `json:"name"`
		Status *MessageOrderPaymentStatus `json:"status"`
		Amount *MessageOrderCost          `json:"amount"`
	}

	MessageOrderPaymentStatus struct {
		Name  string `json:"name"`
		Payed bool   `json:"payed"`
	}

	MessageOrderDelivery struct {
		Name    string            `json:"name"`
		Price   *MessageOrderCost `json:"price"`
		Address string            `json:"address"`
		Comment string            `json:"comment,omitempty"`
	}

	UserRef struct {
		ID                 uint64 `json:"id"`
		ExternalID         string `json:"external_id"`
		Avatar             string `json:"avatar"`
		Type               string `json:"type"`
		Name               string `json:"name"`
		FirstName          string `json:"first_name,omitempty"`
		LastName           string `json:"last_name,omitempty"`
		Phone              string `json:"phone,omitempty"`
		Email              string `json:"email,omitempty"`
		IsAdmin            bool   `json:"is_admin"`
		Available          bool   `json:"available"`
		IsTechnicalAccount bool   `json:"is_technical_account"`
		IsSystem           bool   `json:"is_system"`
	}

	Channel struct {
		ID          uint64          `json:"id"`
		TransportID uint64          `json:"transport_id"`
		Type        string          `json:"type"`
		Name        string          `json:"name"`
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

	Chat struct {
		ID              uint64     `json:"id"`
		Avatar          string     `json:"avatar"`
		Name            string     `json:"name"`
		Channel         *Channel   `json:"channel,omitempty"`
		Members         []Member   `json:"members"`
		Customer        *UserRef   `json:"customer"`
		AuthorID        uint64     `json:"author_id"`
		LastMessage     *Message   `json:"last_message"`
		LastUserMessage *MessageID `json:"last_user_message"`
		LastActivity    string     `json:"last_activity"`
	}

	WaitingChat struct {
		Chat

		WaitingLevel string `json:"waiting_level"`
	}

	Member struct {
		IsAuthor bool     `json:"is_author"`
		State    string   `json:"state"`
		User     *UserRef `json:"user"`
	}

	Dialog struct {
		ID              uint64       `json:"id"`
		BeginMessageID  *uint64      `json:"begin_message_id"`
		EndingMessageID *uint64      `json:"ending_message_id"`
		Chat            *Chat        `json:"chat"`
		Responsible     *Responsible `json:"responsible"`
		CreatedAt       string       `json:"created_at"`
		ClosedAt        *string      `json:"closed_at"`
		Utm             *Utm         `json:"utm,omitempty"`
	}

	FileMeta struct {
		Width  *int `json:"width,omitempty"`
		Height *int `json:"height,omitempty"`
	}

	Item struct {
		ID      string `json:"id"`
		Caption string `json:"caption"`
	}

	Utm struct {
		Source   string `json:"source"`
		Medium   string `json:"medium"`
		Campaign string `json:"campaign"`
		Term     string `json:"term"`
		Content  string `json:"content"`
	}
)

// Channel settings
type (
	CRUDChannelSettings struct {
		Creating string `json:"creating"` // none, receive, send, both
		Editing  string `json:"editing"`  // none, receive, send, both
		Deleting string `json:"deleting"` // none, receive, send, both
	}

	ChannelSettingsText struct {
		CRUDChannelSettings
		Quoting       string `json:"quoting"` // none, receive, send, both
		MaxCharsCount uint16 `json:"max_chars_count"`
	}

	ChannelSettings struct {
		CustomerExternalID string `json:"customer_external_id"`

		SendingPolicy struct {
			NewCustomer       string `json:"new_customer"`
			AfterReplyTimeout string `json:"after_reply_timeout"`
		} `json:"sending_policy"`

		Status struct {
			Delivered string `json:"delivered"`
			Read      string `json:"read"`
		} `json:"status"`

		Text ChannelSettingsText `json:"text"`

		Product struct {
			CRUDChannelSettings
		} `json:"product"`

		Order struct {
			CRUDChannelSettings
		} `json:"order"`

		Image struct {
			CRUDChannelSettings
			Quoting string `json:"quoting"`

			MaxItemsCount     int    `json:"max_items_count"`
			NoteMaxCharsCount uint16 `json:"note_max_chars_count"`
		} `json:"image"`

		File struct {
			CRUDChannelSettings
			Quoting string `json:"quoting"`

			MaxItemsCount     int    `json:"max_items_count"`
			NoteMaxCharsCount uint16 `json:"note_max_chars_count"`
		} `json:"file"`

		Audio struct {
			Creating string `json:"creating"`
			Quoting  string `json:"quoting"`
			Deleting string `json:"deleting"`

			MaxItemsCount int `json:"max_items_count"`
		} `json:"audio"`

		Suggestions struct {
			Text  string `json:"text"`
			Phone string `json:"phone"`
			Email string `json:"email"`
		} `json:"suggestions"`
	}
)

// Events
type (
	WsEventMessageNewData struct {
		Message *Message `json:"message"`
	}

	WsEventMessageUpdatedData struct {
		Message *Message `json:"message"`
	}

	WsEventMessageDeletedData struct {
		Message *Message `json:"message"`
	}

	// Deprecated: Use WsEventWaitingChatCreatedData instead.
	WsEventChatCreatedData struct {
		Chat *Chat `json:"chat"`
	}

	// Deprecated: Use WsEventWaitingChatUpdatedData instead.
	WsEventChatUpdatedData struct {
		Chat *Chat `json:"chat"`
	}

	WsEventDialogOpenedData struct {
		Dialog *Dialog `json:"dialog"`
	}

	WsEventDialogClosedData struct {
		Dialog *Dialog `json:"dialog"`
	}

	WsEventUserLeaveData struct {
		Reason string `json:"reason"`
		Chat   struct {
			ID uint64 `json:"id"`
		} `json:"chat"`
		User struct {
			ID uint64 `json:"id"`
		} `json:"user"`
	}

	WsEventUserUpdatedData struct {
		*UserRef
		IsActive bool `json:"is_active"`
	}

	WsEventCustomerUpdatedData struct {
		*UserRef
	}

	WsEventBotUpdatedData struct {
		*UserRef
	}

	WsEventDialogAssignData struct {
		Dialog *Dialog `json:"dialog"`
		Chat   *Chat   `json:"chat"`
	}

	EventUserJoinedChatData struct {
		Chat *Chat    `json:"chat"`
		User *UserRef `json:"user"`
	}

	WsEventUserOnlineUpdatedData struct {
		User      *UserRef `json:"user"`
		Online    bool     `json:"online"`
		Connected bool     `json:"connected"`
	}

	WsEventChatsDeletedData struct {
		ChatIds []int64 `json:"chat_ids"`
	}

	WsEventChannelUpdatedData struct {
		Channel *ChannelResponseItem `json:"channel"`
	}

	WsEventWaitingChatCreatedData struct {
		Chat *WaitingChat `json:"chat"`
	}

	WsEventWaitingChatUpdatedData struct {
		Chat *WaitingChat `json:"chat"`
	}
)
