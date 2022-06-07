package utils

import (
	"fmt"

	"github.com/millbj92/synctl/pkg/repository"
)

// VerifyRole func for verifying a given role.
func VerifyRole(role string) (string, error) {
	// Switch given role.
	switch role {
	case repository.Admin:
		// Nothing to do, verified successfully.
	case repository.Editor:
		// Nothing to do, verified successfully.
	case repository.Viewer:
		// Nothing to do, verified successfully.
	case repository.Disabled:
		// Nothing to do, verified successfully.
	default:
		// Return error message.
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}
