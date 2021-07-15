package viber

var errorCodes = map[int]string{
	0:  "ok",                           // Success
	1:  "invalidUrl",                   // The webhook URL is not valid
	2:  "invalidAuthToken",             // The authentication token is not valid
	3:  "badData",                      // There is an error in the request itself (missing comma, brackets, etc.)
	4:  "missingData",                  // Some mandatory data is missing
	5:  "receiverNotRegistered",        // The receiver is not registered to Viber
	6:  "receiverNotSubscribed",        // The receiver is not subscribed to the account
	7:  "publicAccountBlocked",         // The account is blocked
	8:  "publicAccountNotFound",        // The account associated with the token is not a account.
	9:  "publicAccountSuspended",       // The account is suspended
	10: "webhookNotSet",                // No webhook was set for the account
	11: "receiverNoSuitableDevice",     // The receiver is using a device or a Viber version that don’t support accounts
	12: "tooManyRequests",              // Rate control breach
	13: "apiVersionNotSupported",       // Maximum supported account version by all user’s devices is less than the minApiVersion in the message
	14: "incompatibleWithVersion",      // minApiVersion is not compatible to the message fields
	15: "publicAccountNotAuthorized",   // The account is not authorized
	16: "inchatReplyMessageNotAllowed", // Inline message not allowed
	17: "publicAccountIsNotInline",     // The account is not inline
	18: "noPublicChat",                 // Failed to post to public account.The bot is missing a Public Chat interface
	19: "cannotSendBroadcast",          // Cannot send broadcast message
	20: "broadcastNotAllowed",          // Attempt to send broadcast message from the bot
}
