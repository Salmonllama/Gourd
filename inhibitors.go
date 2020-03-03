package gourd

import (
	"github.com/andersfylling/disgord"
)

/*
 * Inhibitors
 * All inhibitors except NilInhibitor, OwnerInhibitor have a Value and a Response
 * Value is the Inhibitor's 'comparator'; what it will use to decide command usage or not
 * Response is the message that will get sent back when the command is 'inhibited'. No message will be sent if nil.
 */

/*
 * Nil Inhibitor
 * Allows command usage no matter what
 * Value is neither present nor necessary
 * Response is neither present nor necessary
 */
type NilInhibitor struct {
}

func (inhibitor *NilInhibitor) String() string {
	return "This command requires no permissions."
}

/*
 * Role Inhibitor
 * Allows command usage if the user has the role (value)
 * Value is the string ID of the role, **NOT the snowflake!**
 */
type RoleInhibitor struct {
	Value    string
	Response interface{}
}

/*
 * Permission Inhibitor
 * Allows command usage if the user has the permission bit (value)
 * Value is the disgord.PermissionBit. It is recommended to use disgord.PermissionBlahBlah
 */
type PermissionInhibitor struct {
	Value    disgord.PermissionBit
	Response interface{}
}

func (inhibitor *PermissionInhibitor) GetValue() disgord.PermissionBit {
	return inhibitor.Value
}

/*
 * Keyword Inhibitor
 * Allows command usage if the user has the given keyword
 * Value is the keyword they should be assigned to.
 * See <wiki link> for keyword reference
 */
type KeywordInhibitor struct {
	Value    string
	Response interface{}
}

/*
 * Owner Inhibitor
 * Allows command usage only if the user is the bot owner
 * value means nothing; The owner ID is supplied in Gourd initialization
 */
type OwnerInhibitor struct {
	Value    string
	Response interface{}
}
