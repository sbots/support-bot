package models

// Chat contains information about the place a message was sent.
type Chat struct {
	// ID is a unique identifier for this chat
	ID int64 `json:"id"`
	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`
	// Title for supergroups, channels and group chats
	//
	// optional
	Title string `json:"title"`
	// UserName for private chats, supergroups and channels if available
	//
	// optional
	UserName string `json:"username"`
	// FirstName of the other party in a private chat
	//
	// optional
	FirstName string `json:"first_name"`
	// LastName of the other party in a private chat
	//
	// optional
	LastName string `json:"last_name"`
	// AllMembersAreAdmins
	//
	// optional
	AllMembersAreAdmins bool `json:"all_members_are_administrators"`
	// Photo is a chat photo
	Photo *ChatPhoto `json:"photo"`
	// Description for groups, supergroups and channel chats
	//
	// optional
	Description string `json:"description,omitempty"`
	// InviteLink is a chat invite link, for groups, supergroups and channel chats.
	// Each administrator in a chat generates their own invite links,
	// so the bot must first generate the link using exportChatInviteLink
	//
	// optional
	InviteLink string `json:"invite_link,omitempty"`
	// PinnedMessage Pinned message, for groups, supergroups and channels
	//
	// optional
	PinnedMessage *Message `json:"pinned_message"`
}
