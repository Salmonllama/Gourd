package gourd

import (
	"context"
	"github.com/andersfylling/disgord"
)

type Inhibitor struct {
	Value interface{}
	Response interface{}
}

type InhibitorHandler interface {
	handle(ctx CommandContext) bool
}

// Interface with a handle function -> interfacing makes it easer to work with in the main gourd logic

// Nil Inhibitor.
// Allows command usage no matter what.
// Value is neither present nor necessary.
// Response is neither present nor necessary.
type NilInhibitor struct {
	Inhibitor
}

func (nilInhibitor *NilInhibitor) handle(ctx CommandContext) bool {
	return true
}

// Role Inhibitor.
// Allows command usage if the user has the role (value).
// Value is the string ID of the role, **NOT the snowflake!**.
type RoleInhibitor struct {
	Inhibitor
}

// Will return false in the case of private messages -> no roles
func (roleInhibitor *RoleInhibitor) handle(ctx CommandContext) bool {
	if ctx.IsPrivate() {
		return false
	}

	requiredRole := roleInhibitor.Value.(string) // This will be the ID
	userRoles := ctx.AuthorMember().Roles

	for _, role := range userRoles {
		if role.String() == requiredRole {
			return true
		}
	}

	return false
}

// Permission Inhibitor.
// Allows command usage if the user has the permission bit (value).
// Value is the disgord.PermissionBit. It is recommended to use disgord.PermissionBlahBlah.
type PermissionInhibitor struct {
	Inhibitor
}

// Will return false in the case of private messages -> no discord permissions
func (permissionInhibitor *PermissionInhibitor) handle(ctx CommandContext) bool {
	if ctx.IsPrivate() {
		return false
	}

	guild, err := ctx.Guild()
	if err != nil {
		return false
	}

	requiredPerm := permissionInhibitor.Value.(disgord.PermissionBit)
	userPerm, err := ctx.Client.GetMemberPermissions(
		context.Background(),
		guild.ID,
		ctx.Author().ID,
	)
	if err != nil {
		return false
	}

	if userPerm&disgord.PermissionAdministrator > 0 { // Admin is a diff. permission
		return true
	}

	if userPerm&requiredPerm > 0 {
		return true
	}

	return false
}

// Keyword Inhibitor.
// Allows command usage if the user has the given keyword.
// Value is the string keyword they should be assigned to.
// See <wiki link> for keyword reference.
type KeywordInhibitor struct {
	Inhibitor
}

// Works in direct messages -> not dependent on a guild
func (keywordInhibitor *KeywordInhibitor) handle(ctx CommandContext) bool {
	keyword := keywordInhibitor.Value.(string)
	if ctx.Gourd.HasKeyword(ctx.Author().ID.String(), keyword) {
		return true
	}

	return false
}

// Owner Inhibitor.
// Allows command usage only if the user is the bot owner.
// value means nothing; The owner ID is supplied in Gourd initialization.
type OwnerInhibitor struct {
	Inhibitor
}

func (ownerInhibitor *OwnerInhibitor) handle(ctx CommandContext) bool {
	if ctx.IsAuthorOwner() {
		return true
	}

	return false
}
