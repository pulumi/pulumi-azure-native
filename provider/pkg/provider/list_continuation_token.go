package provider

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type listContinuationToken struct {
	// NextLink is the URL to fetch the next page of results, if any.
	// This is an Azure URL returned by pageable operations used for listing subsequent pages of results.
	// When it is non-empty, the provider will use this URL to fetch the next page of results instead of
	// making a new list request with the original parameters.
	NextLink string `json:"nextLink,omitempty"`
	// Remaining is the number of items remaining to be listed on a per-session basis.
	Remaining *int64 `json:"remaining,omitempty"`
}

func isEmptyListContinuationToken(token *listContinuationToken) bool {
	if token == nil {
		return true
	}
	if token.NextLink == "" && token.Remaining == nil {
		return true
	}

	if token.Remaining != nil && *token.Remaining <= 0 {
		return true
	}

	return false
}

func encodeListContinuationToken(token *listContinuationToken) (string, error) {
	if isEmptyListContinuationToken(token) {
		return "", nil
	}

	b, err := json.Marshal(token)
	if err != nil {
		return "", fmt.Errorf("marshalling list continuation token: %w", err)
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func decodeListContinuationToken(encoded string) (*listContinuationToken, error) {
	if encoded == "" {
		return &listContinuationToken{
			NextLink:  "",
			Remaining: nil,
		}, nil
	}

	b, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("decoding list continuation token: %w", err)
	}

	var token listContinuationToken
	if err := json.Unmarshal(b, &token); err != nil {
		return nil, fmt.Errorf("unmarshalling list continuation token: %w", err)
	}

	return &token, nil
}
