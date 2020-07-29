package gourd

import (
	"context"
	"github.com/andersfylling/disgord"
)

type Handler struct {
	Commands []*Command
	Modules  []*Module
}

func (handler *Handler) addCommand(cmd *Command) *Command {
	handler.Commands = append(handler.Commands, cmd)
	return cmd
}

func (handler *Handler) GetCommandMap() map[string]*Command {
	cmdMap := make(map[string]*Command)

	for _, cmd := range handler.Commands {
		cmdMap[cmd.Name] = cmd
	}

	return cmdMap
}

func (handler *Handler) AddModule(mdl *Module) *Module {
	handler.Modules = append(handler.Modules, mdl)

	for _, cmd := range mdl.Commands {
		handler.addCommand(cmd)
	}

	return mdl
}

func (handler *Handler) GetModuleMap() map[string]*Module {
	mdlMap := make(map[string]*Module)

	for _, mdl := range handler.Modules {
		mdlMap[mdl.Name] = mdl
	}

	return mdlMap
}

// handleNilInhibitor allows usage no matter what
func (handler *Handler) handleNilInhibitor() bool {
	return true
}

// handleRoleInhibitor looks for the supplied role in the user's roles
// Will return false in case of Direct Message -> roles not possible
func (handler *Handler) handleRoleInhibitor(ctx CommandContext) bool {
	if ctx.IsPrivate() {
		return false
	}

	inhibitor := ctx.Command.Module.Inhibitor.(RoleInhibitor)
	requiredRole := inhibitor.Value // This will be the ID
	userRoles := ctx.AuthorMember().Roles

	for _, role := range userRoles {
		if role.String() == requiredRole {
			return true
		}
	}

	if inhibitor.Response != nil {
		_, err := ctx.Reply(inhibitor.Response)
		if err != nil {
			return false
		}
	}
	return false
}

// handlePermissionInhibitor compares user permission bit with required permission bit
// Will return false in case of Direct Message -> permissions not possible
func (handler *Handler) handlePermissionInhibitor(ctx CommandContext) bool {
	if ctx.IsPrivate() {
		return false
	}

	guild, err := ctx.Guild()
	if err != nil {
		return false
	}

	inhibitor := ctx.Command.Module.Inhibitor.(PermissionInhibitor)
	requiredPerm := inhibitor.Value
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

	if inhibitor.Response != nil {
		_, err := ctx.Reply(inhibitor.Response)
		if err != nil {
			return false
		}
	}
	return false
}

// handleKeywordInhibitor looks for assigned keywords on the user
// Works in direct messages; is not guild-dependant
func (handler *Handler) handleKeywordInhibitor(ctx CommandContext) bool {
	inhibitor := ctx.Command.Module.Inhibitor.(KeywordInhibitor)

	if ctx.Gourd.HasKeyword(ctx.Author().ID.String(), inhibitor.Value) {
		return true
	}

	if inhibitor.Response != nil {
		_, err := ctx.Reply(inhibitor.Response)
		if err != nil {
			return false
		}
	}
	return false
}

func (handler *Handler) handleOwnerInhibitor(ctx CommandContext) bool {
	inhibitor := ctx.Command.Module.Inhibitor.(OwnerInhibitor)

	if ctx.IsAuthorOwner() {
		return true
	}

	if inhibitor.Response != nil {
		_, err := ctx.Reply(inhibitor.Response)
		if err != nil {
			return false
		}
	}
	return false
}
