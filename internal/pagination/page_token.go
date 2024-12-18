// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package pagination

import (
	"fmt"
)

// PageToken is a page token that uses an offset to delineate which page to fetch.
type PageToken struct {
	// Offset of the page.
	Offset int64
	// RequestChecksum is the checksum of the request that generated the page token.
	RequestChecksum uint32
}

// pageTokenChecksumMask is a random bitmask applied to offset-based page token checksums.
//
// Change the bitmask to force checksum failures when changing the page token implementation.
const pageTokenChecksumMask uint32 = 0x9acb0442

// ParsePageToken parses an offset-based page token from the provided Request.
//
// If the request does not have a page token, a page token with offset 0 will be returned.
func ParsePageToken(request Request) (_ PageToken, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("parse page token: %w", err)
		}
	}()

	requestChecksum, err := calculateRequestChecksum(request)
	requestChecksum ^= pageTokenChecksumMask // apply checksum mask for PageToken

	if request.GetPageToken() == "" {
		offset := int64(0)
		if s, ok := request.(skipRequest); ok {
			offset += int64(s.GetSkip())
		}

		return PageToken{
			Offset:          offset,
			RequestChecksum: requestChecksum,
		}, nil
	}

	var pageToken PageToken
	if err = DecodePageToken(request.GetPageToken(), &pageToken); err != nil {
		return PageToken{}, err
	}
	if pageToken.RequestChecksum != requestChecksum {
		return PageToken{}, fmt.Errorf(
			"checksum mismatch (got 0x%x but expected 0x%x)", pageToken.RequestChecksum, requestChecksum,
		)
	}

	if s, ok := request.(skipRequest); ok {
		pageToken.Offset += int64(s.GetSkip())
	}
	return pageToken, nil
}

// Next returns the next page token for the provided Request.
func (p PageToken) Next(request Request) PageToken {
	p.Offset += int64(request.GetPageSize())
	return p
}

// String returns a string representation of the page token.
func (p PageToken) String() string {
	return EncodePageToken(&p)
}
