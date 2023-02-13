package car

import (
	"context"
	"sfcar/internal/id"
	rentalpb "sfcar/rental/api/gen/v1"
)

// Manager defines a car manager.
type Manager struct {
}

// Verify the car status.
func (m *Manager) Verify(ctx context.Context, id id.CarID, location *rentalpb.Location) error {
	return nil
}

// Unlock the car, if lock failure returns the reason.
func (m *Manager) Unlock(ctx context.Context, id id.CarID) error {
	return nil
}

// Lock locks a car.
func (m *Manager) Lock(ctx context.Context, cid id.CarID) error {
	return nil
}
