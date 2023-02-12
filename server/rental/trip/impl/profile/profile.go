package profile

import (
	"context"
	"sfcar/internal/id"
)

// Manager defines a profile manager.
type Manager struct {
}

// Verify verifies that the account is eligible to create a trip.
func (m *Manager) Verify(context.Context, id.AccountID) (id.IdentityID, error) {
	return id.IdentityID("identity1"), nil
}
