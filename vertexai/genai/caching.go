package genai

import (
	"fmt"
	"time"

	aiplatform "github.com/qclaogui/gaip/genproto/aiplatform/apiv1beta1"
	pb "github.com/qclaogui/gaip/genproto/aiplatform/apiv1beta1/aiplatformpb"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type cacheClient = aiplatform.GenAiCacheClient

var (
	newCacheClient     = aiplatform.NewGenAiCacheClient
	newCacheRESTClient = aiplatform.NewGenAiCacheRESTClient
)

// GenerativeModelFromCachedContent returns a [GenerativeModel] that uses the given [CachedContent].
// The argument should come from a call to [Client.CreateCachedContent] or [Client.GetCachedContent].
func (c *Client) GenerativeModelFromCachedContent(cc *CachedContent) *GenerativeModel {
	return &GenerativeModel{
		c:                 c,
		name:              cc.Model,
		fullName:          inferFullModelName(c.projectID, c.location, cc.Model),
		CachedContentName: cc.Name,
	}
}

// CachedContentToUpdate specifies which fields of a CachedContent to modify in a call to
// [Client.UpdateCachedContent].
type CachedContentToUpdate struct {
	// If non-nil, update the expire time or TTL.
	Expiration *ExpireTimeOrTTL
}

// ExpireTimeOrTTL describes the time when a resource expires.
// If ExpireTime is non-zero, it is the expiration time.
// Otherwise, the expiration time is the value of TTL ("time to live") added
// to the current time.
type ExpireTimeOrTTL struct {
	ExpireTime time.Time
	TTL        time.Duration
}

// populateCachedContentTo populates some fields of p from v.
func populateCachedContentTo(p *pb.CachedContent, v *CachedContent) {
	exp := v.Expiration
	if !exp.ExpireTime.IsZero() {
		p.Expiration = &pb.CachedContent_ExpireTime{
			ExpireTime: timestamppb.New(exp.ExpireTime),
		}
	} else if exp.TTL != 0 {
		p.Expiration = &pb.CachedContent_Ttl{
			Ttl: durationpb.New(exp.TTL),
		}
	}
	// If both fields of v.Expiration are zero, leave p.Expiration unset.
}

// populateCachedContentFrom populates some fields of v from p.
func populateCachedContentFrom(v *CachedContent, p *pb.CachedContent) {
	if p.Expiration == nil {
		return
	}
	switch e := p.Expiration.(type) {
	case *pb.CachedContent_ExpireTime:
		v.Expiration.ExpireTime = pvTimeFromProto(e.ExpireTime)
	case *pb.CachedContent_Ttl:
		v.Expiration.TTL = e.Ttl.AsDuration()
	default:
		panic(fmt.Sprintf("unknown type of CachedContent.Expiration: %T", p.Expiration))
	}
}
