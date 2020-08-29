package gourd

import (
	"github.com/andersfylling/disgord"
)

// NilInhibitor allows command usage no matter what.
// Does not have a value or a response
type NilInhibitor struct{}

// RoleInhibitor allows command usage if the user has the role (value).
// Value is the string ID of the role, **NOT the snowflake!**.
// This inhibitor will not work in private messages as there are no roles.
type RoleInhibitor struct {
	Value    string
	Response interface{}
}

// Will return false in the case of private messages -> no roles
func (roleInhibitor RoleInhibitor) handle(userRoles []disgord.Snowflake) bool {
	requiredRole := roleInhibitor.Value // This will be the ID

	for _, role := range userRoles {
		if role.String() == requiredRole {
			return true
		}
	}

	return false
}

// PermissionInhibitor allows command usage if the user has the permission bit (value).
// Value is the disgord.PermissionBit. It is recommended to use disgord.PermissionBlahBlah.
// This inhibitor will not work in private messages as there are no permissions.
type PermissionInhibitor struct {
	Value    uint64
	Response interface{}
}

// Will return false in the case of private messages -> no discord permissions
func (permissionInhibitor PermissionInhibitor) handle(userPerm disgord.PermissionBits) bool {
	requiredPerm := permissionInhibitor.Value

	if userPerm&disgord.PermissionAdministrator > 0 { // Admin is a diff. permission
		return true
	}

	if userPerm&requiredPerm > 0 {
		return true
	}

	return false
}

// KeywordInhibitor allows command usage if the user has the given keyword.
// Value is the string keyword they should be assigned to.
// See <wiki link here> for keyword how-tos
type KeywordInhibitor struct {
	Value    string
	Response interface{}
}

// OwnerInhibitor allows command usage only if the user is the bot owner.
// Does not have a value; the owner ID is supplied in Gourd initialization.
type OwnerInhibitor struct {
	Response interface{}
}
