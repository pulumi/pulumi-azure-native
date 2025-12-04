package azure

import (
	"context"
)

type AzureDeleter interface {
	Delete(ctx context.Context, id, apiVersion, asyncStyle string, queryParams map[string]any) error
}

type AzureClient interface {
	AzureDeleter
	CanCreate(ctx context.Context, id, path, apiVersion, readMethod string, isSingletonResource, hasDefaultBody bool, isDefaultResponse func(map[string]any) bool) error
	Get(ctx context.Context, id string, apiVersion string, queryParams map[string]any) (map[string]any, error)
	Head(ctx context.Context, id string, apiVersion string) error
	Patch(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error)
	Post(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}) (any, error)
	Put(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error)
}

// MockAzureClient implements the AzureClient interface for tests.
type MockAzureClient struct {
	// Resource ids that were retrieved via Get, in order
	GetIds []string
	// API versions that were used in Get, in order
	GetApiVersions []string

	// If set, this response will be returned for all Get requests; otherwise nil.
	GetResponse    map[string]any
	GetResponseErr error

	// Resource ids that were used in Post, in order
	PostIds []string
	// Bodies that were sent via Post, in order
	PostBodies []map[string]any

	// Resource ids that were used in Put, in order
	PutIds []string
	// Bodies that were sent via Put, in order
	PutBodies []map[string]any

	PutResponseErr error

	QueryParamsOfLastDelete map[string]any
}

func (m *MockAzureClient) Delete(ctx context.Context, id, apiVersion, asyncStyle string, queryParams map[string]any) error {
	m.QueryParamsOfLastDelete = queryParams
	return nil
}
func (m *MockAzureClient) CanCreate(ctx context.Context, id, path, apiVersion, readMethod string, isSingletonResource, hasDefaultBody bool, isDefaultResponse func(map[string]any) bool) error {
	return nil
}
func (m *MockAzureClient) Get(ctx context.Context, id string, apiVersion string, queryParams map[string]any) (map[string]any, error) {
	m.GetIds = append(m.GetIds, id)
	m.GetApiVersions = append(m.GetApiVersions, apiVersion)
	return m.GetResponse, m.GetResponseErr
}
func (m *MockAzureClient) Head(ctx context.Context, id string, apiVersion string) error {
	return nil
}
func (m *MockAzureClient) Patch(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error) {
	return nil, false, nil
}
func (m *MockAzureClient) Post(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}) (any, error) {
	m.PostIds = append(m.PostIds, id)
	m.PostBodies = append(m.PostBodies, bodyProps)
	return map[string]any{}, nil
}
func (m *MockAzureClient) Put(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error) {
	m.PutIds = append(m.PutIds, id)
	m.PutBodies = append(m.PutBodies, bodyProps)
	return nil, false, m.PutResponseErr
}
func (m *MockAzureClient) IsNotFound(err error) bool {
	return false
}
