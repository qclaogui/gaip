// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

import (
	"fmt"
	"time"

	pb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta/generativelanguagepb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Caches provides methods for managing the context caching.
// You don't need to initiate this struct. Create a client instance via NewClient, and
// then access Caches through client.Caches field.
type Caches struct {
	apiClient *apiClient
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
