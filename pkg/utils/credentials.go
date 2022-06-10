package utils

import (
	"fmt"

	"github.com/millbj92/synctl/pkg/repository"
)

// GetCredentialsByRole func for getting credentials from a role name.
func GetCredentialsByRole(role string) ([]string, error) {
	// Define credentials variable.
	var credentials []string

	// Switch given role.
	switch role {
	case repository.Admin:
		// Admin credentials (all access).
		credentials = []string{
			repository.ConnectionsAdd,
			repository.ConnectionsList,
			repository.ConnectionsRemove,
			repository.ConnectionsUpdate,
			repository.LinksAdd,
			repository.LinksList,
			repository.LinksRemove,
			repository.LinksUpdate,
			repository.PermissionsAdd,
			repository.PermissionsList,
			repository.PermissionsRemove,
			repository.PermissionsUpdate,
			repository.RemoteConfigAdd,
			repository.RemoteConfigList,
			repository.RemoteConfigRemove,
			repository.RemoteConfigUpdate,
			repository.RolesAdd,
			repository.RolesList,
			repository.RolesRemove,
			repository.RolesUpdate,
			repository.SettingsAdd,
			repository.SettingsList,
			repository.SettingsUpdate,
			repository.TasksAdd,
			repository.TasksList,
			repository.TasksRemove,
			repository.TasksRun,
			repository.TasksStop,
			repository.TasksRestart,
			repository.TasksPause,
			repository.TasksResume,
			repository.TasksLogs,
			repository.TasksStatus,
			repository.TasksLogs_Download,
			repository.UsersAdd,
			repository.UsersList,
			repository.UsersRemove,
			repository.UsersUpdate,
		}
	case repository.Editor:
		// Moderator credentials (only book creation and update).
		credentials = []string{
			repository.ConnectionsAdd,
			repository.ConnectionsList,
			repository.ConnectionsUpdate,
			repository.LinksAdd,
			repository.LinksList,
			repository.LinksUpdate,
			repository.RemoteConfigAdd,
			repository.RemoteConfigList,
			repository.RemoteConfigUpdate,
			repository.TasksAdd,
			repository.TasksList,
			repository.TasksUpdate,
			repository.TasksRun,
			repository.TasksResume,
			repository.TasksRestart,
			repository.TasksLogs,
			repository.TasksStatus,
			repository.TasksLogs_Download,
			repository.SettingsList,
		}
	case repository.Viewer:
		// Simple user credentials (only book creation).
		credentials = []string{
			repository.ConnectionsList,
			repository.LinksList,
			repository.RemoteConfigList,
			repository.TasksList,
			repository.UsersList,
		}
	case repository.Disabled:
		// Disabled user credentials (no access).
		credentials = []string{}
	default:
		// Return error message.
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
