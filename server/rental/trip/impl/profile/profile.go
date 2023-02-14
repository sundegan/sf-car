package profile

import (
	"context"
	"encoding/base64"
	"fmt"
	"google.golang.org/protobuf/proto"
	"sfcar/internal/id"
	rentalpb "sfcar/rental/api/gen/v1"
)

// Fetcher defines the interface to fetch profile.
type Fetcher interface {
	GetProfile(c context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error)
}

// Manager defines a profile manager.
type Manager struct {
	Fetcher Fetcher
}

// Verify verifies that the account is eligible to create a trip.
func (m *Manager) Verify(c context.Context, aid id.AccountID) (id.IdentityID, error) {
	nilID := id.IdentityID("")
	p, err := m.Fetcher.GetProfile(c, &rentalpb.GetProfileRequest{})
	if err != nil {
		return nilID, fmt.Errorf("cannot get profile: %v", err)
	}

	if p.IdentityStatus != rentalpb.IdentityStatus_VERIFIED {
		return nilID, fmt.Errorf("invalid identity status: %v", err)
	}

	b, err := proto.Marshal(p.Identity)
	if err != nil {
		return nilID, fmt.Errorf("cannot marshal identity: %v", err)
	}
	iID := id.IdentityID(base64.StdEncoding.EncodeToString(b))
	return iID, nil
}
