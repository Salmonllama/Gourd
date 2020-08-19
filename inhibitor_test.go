package gourd

import (
	"github.com/andersfylling/disgord"
	"testing"
)

func TestRoleHandler(t *testing.T) {
	ctx := CommandContext{
		Command: &Command{
			Inhibitor: RoleInhibitor{
				Value: "472232069301534726",
			},
		},
	}

	roleId := disgord.Snowflake(472232069301534726)
	userRoles := []disgord.Snowflake{roleId}

	if ret := ctx.Command.Inhibitor.(RoleInhibitor).handle(userRoles); ret != true {
		t.Error("Expected true value from RoleInhibitor")
	}
}

func TestPermissionHandler(t *testing.T) {
	ctx := CommandContext{
		Command: &Command{
			Inhibitor: PermissionInhibitor{
				Value: disgord.PermissionAddReactions,
			},
		},
	}

	if ret := ctx.Command.Inhibitor.(PermissionInhibitor).handle(disgord.PermissionAddReactions); ret != true {
		t.Error("Expected true value from PermissionInhibitor")
	}
}
