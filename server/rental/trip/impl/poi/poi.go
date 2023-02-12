package poi

import (
	"context"
	"github.com/golang/protobuf/proto"
	"hash/fnv"
	rentalpb "sfcar/rental/api/gen/v1"
)

var poi = []string{
	"中关村",
	"天安门",
	"陆家嘴",
	"迪士尼",
	"天河体育中心",
	"广州塔",
}

// Manager defines a poi manager.
type Manager struct {
}

// Resolve resolves the given location random return the poi.
func (*Manager) Resolve(c context.Context, loc *rentalpb.Location) (string, error) {
	b, err := proto.Marshal(loc)
	if err != nil {
		return "", err
	}

	h := fnv.New32()
	_, _ = h.Write(b)
	//
	return poi[int(h.Sum32())%len(poi)], nil
}
