package provider

import (
	"context"
	"testing"

	az "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

// mockListStream captures every rpc.ListResponse sent through it.
type mockListStream struct {
	sent []*rpc.ListResponse
}

func (s *mockListStream) Send(resp *rpc.ListResponse) error {
	s.sent = append(s.sent, resp)
	return nil
}

// grpc.ServerStream no-ops.
func (s *mockListStream) SetHeader(metadata.MD) error  { return nil }
func (s *mockListStream) SendHeader(metadata.MD) error { return nil }
func (s *mockListStream) SetTrailer(metadata.MD)       {}
func (s *mockListStream) Context() context.Context     { return context.Background() }
func (s *mockListStream) SendMsg(any) error            { return nil }
func (s *mockListStream) RecvMsg(any) error            { return nil }

// results returns only the Result-typed responses, in order.
func (s *mockListStream) results() []*rpc.ListResponse_Result {
	var out []*rpc.ListResponse_Result
	for _, r := range s.sent {
		if v := r.GetResult(); v != nil {
			out = append(out, v)
		}
	}
	return out
}

// continuationToken returns the encoded token from the first Continuation response, or "".
func (s *mockListStream) continuationToken() string {
	for _, r := range s.sent {
		if v := r.GetContinuation(); v != nil {
			return v.ContinuationToken
		}
	}
	return ""
}

// sequentialAzureClient returns a different Get response on each call, in order.
type sequentialAzureClient struct {
	az.MockAzureClient
	getResponses []map[string]any
	callIndex    int
}

func (c *sequentialAzureClient) Get(ctx context.Context, id, apiVersion string, queryParams map[string]any) (map[string]any, error) {
	c.GetIds = append(c.GetIds, id)
	c.GetApiVersions = append(c.GetApiVersions, apiVersion)
	c.GetQueryParams = append(c.GetQueryParams, queryParams)
	idx := c.callIndex
	c.callIndex++
	if idx < len(c.getResponses) {
		return c.getResponses[idx], nil
	}
	return map[string]any{}, nil
}

// --- test helpers ---

const listTok = "azure-native:test/v20230101:Widget"

// baseListParams are the minimal path parameters for a list operation.
var baseListParams = []resources.AzureAPIParameter{
	{Name: "subscriptionId", Location: "path", IsRequired: true, Value: &resources.AzureAPIProperty{Type: "string"}},
	{Name: "resourceGroupName", Location: "path", IsRequired: true, Value: &resources.AzureAPIProperty{Type: "string"}},
}

func widgetResource(extraParams ...resources.AzureAPIParameter) resources.AzureAPIResource {
	params := append(baseListParams, extraParams...)
	return resources.AzureAPIResource{
		APIVersion: "2023-01-01",
		Path:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Test/widgets/{widgetName}",
		PutParameters: []resources.AzureAPIParameter{
			{Name: "subscriptionId", Location: "path", IsRequired: true, Value: &resources.AzureAPIProperty{Type: "string"}},
		},
		Response: map[string]resources.AzureAPIProperty{},
		ListMetadata: &resources.AzureAPIListMetadata{
			Parameters:    params,
			OperationPath: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Test/widgets",
			Method:        "GET",
			NextLinkName:  "nextLink",
		},
	}
}

func providerWithListResource(client az.AzureClient, res resources.AzureAPIResource) *azureNativeProvider {
	rm := &resources.APIMetadata{
		Types:     resources.GoMap[resources.AzureAPIType]{},
		Resources: resources.GoMap[resources.AzureAPIResource]{listTok: res},
		Invokes:   resources.GoMap[resources.AzureAPIInvoke]{},
	}
	return &azureNativeProvider{azureClient: client, resourceMap: rm}
}

func makeListReq(continuationToken string, limit, pageSize int64, extraQuery resource.PropertyMap) *rpc.ListRequest {
	q := resource.PropertyMap{
		"subscriptionId":    resource.NewStringProperty("sub-123"),
		"resourceGroupName": resource.NewStringProperty("rg-test"),
	}
	for k, v := range extraQuery {
		q[k] = v
	}
	props, _ := plugin.MarshalProperties(q, plugin.MarshalOptions{SkipNulls: true})
	return &rpc.ListRequest{
		Token:             listTok,
		Query:             props,
		ContinuationToken: continuationToken,
		Limit:             limit,
		PageSize:          pageSize,
	}
}

func pageOf(ids []string, nextLink string) map[string]any {
	items := make([]any, len(ids))
	for i, id := range ids {
		items[i] = map[string]any{"id": id, "name": "widget-" + id}
	}
	page := map[string]any{"value": items}
	if nextLink != "" {
		page["nextLink"] = nextLink
	}
	return page
}

// --- tests ---

func TestList_ResourceNotFound(t *testing.T) {
	p := providerWithListResource(&az.MockAzureClient{}, widgetResource())
	stream := &mockListStream{}
	err := p.List(&rpc.ListRequest{Token: "azure-native:test/v20230101:DoesNotExist"}, stream)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestList_NotListable(t *testing.T) {
	res := resources.AzureAPIResource{
		APIVersion:    "2023-01-01",
		Path:          "/some/path",
		PutParameters: []resources.AzureAPIParameter{},
		Response:      map[string]resources.AzureAPIProperty{},
		ListMetadata:  nil,
	}
	p := providerWithListResource(&az.MockAzureClient{}, res)
	stream := &mockListStream{}
	err := p.List(&rpc.ListRequest{Token: listTok}, stream)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not listable")
}

func TestList_SinglePage(t *testing.T) {
	client := &az.MockAzureClient{
		GetResponse: pageOf([]string{"/res/a", "/res/b", "/res/c"}, ""),
	}
	p := providerWithListResource(client, widgetResource())
	stream := &mockListStream{}

	err := p.List(makeListReq("", 0, 0, nil), stream)
	require.NoError(t, err)

	results := stream.results()
	require.Len(t, results, 3)
	assert.Equal(t, "/res/a", results[0].Id)
	assert.Equal(t, "/res/b", results[1].Id)
	assert.Equal(t, "/res/c", results[2].Id)
	assert.Empty(t, stream.continuationToken(), "no continuation token when there is no nextLink")
}

func TestList_NoNextLinkMeansNoContinuation(t *testing.T) {
	client := &az.MockAzureClient{
		GetResponse: pageOf([]string{"/res/a"}, ""),
	}
	p := providerWithListResource(client, widgetResource())
	stream := &mockListStream{}

	require.NoError(t, p.List(makeListReq("", 0, 0, nil), stream))
	assert.Empty(t, stream.continuationToken())
}

func TestList_NextLinkProducesContinuationToken(t *testing.T) {
	nextLink := "https://management.azure.com/subscriptions/sub-123/providers/Microsoft.Test/widgets?api-version=2023-01-01&%24skipToken=abc"
	client := &az.MockAzureClient{
		GetResponse: pageOf([]string{"/res/a"}, nextLink),
	}
	p := providerWithListResource(client, widgetResource())
	stream := &mockListStream{}

	require.NoError(t, p.List(makeListReq("", 0, 0, nil), stream))

	ct := stream.continuationToken()
	require.NotEmpty(t, ct)

	// Decode and verify the embedded nextLink.
	decoded, err := decodeListContinuationToken(ct)
	require.NoError(t, err)
	assert.Equal(t, nextLink, decoded.NextLink)
	assert.Nil(t, decoded.Remaining, "no limit means Remaining should be nil")
}

func TestList_ContinuationTokenFollowsNextLink(t *testing.T) {
	nextLink := "https://management.azure.com/subscriptions/sub-123/providers/Microsoft.Test/widgets?skipToken=page2"
	client := &sequentialAzureClient{
		getResponses: []map[string]any{
			pageOf([]string{"/res/a", "/res/b"}, nextLink),
			pageOf([]string{"/res/c", "/res/d"}, ""),
		},
	}
	p := providerWithListResource(client, widgetResource())

	// First call: get page 1 and capture continuation token.
	stream1 := &mockListStream{}
	require.NoError(t, p.List(makeListReq("", 0, 0, nil), stream1))
	assert.Len(t, stream1.results(), 2)

	ct := stream1.continuationToken()
	require.NotEmpty(t, ct)

	// Second call: supply the continuation token to get page 2.
	stream2 := &mockListStream{}
	require.NoError(t, p.List(makeListReq(ct, 0, 0, nil), stream2))

	results2 := stream2.results()
	require.Len(t, results2, 2)
	assert.Equal(t, "/res/c", results2[0].Id)
	assert.Equal(t, "/res/d", results2[1].Id)
	assert.Empty(t, stream2.continuationToken(), "last page has no nextLink")

	// Second call must have used the nextLink URL, not the original path.
	// The skip token travels in the query params map (not in the id) because
	// initRequest overwrites RawQuery from queryParams, so we split path and query.
	require.Len(t, client.GetIds, 2)
	assert.Equal(t, "/subscriptions/sub-123/providers/Microsoft.Test/widgets", client.GetIds[1],
		"second Get must use the nextLink path")
	assert.Equal(t, "page2", client.GetQueryParams[1]["skipToken"],
		"second Get must carry the skip token from the nextLink query")
}

func TestList_LimitStopsStreaming(t *testing.T) {
	client := &az.MockAzureClient{
		GetResponse: pageOf([]string{"/res/a", "/res/b", "/res/c", "/res/d", "/res/e"}, ""),
	}
	p := providerWithListResource(client, widgetResource())
	stream := &mockListStream{}

	require.NoError(t, p.List(makeListReq("", 2, 0, nil), stream))

	assert.Len(t, stream.results(), 2)
	assert.Empty(t, stream.continuationToken(), "limit reached, no more pages needed")
}

func TestList_LimitExactlyMet_NoSpuriousContinuation(t *testing.T) {
	// When the page has exactly as many items as the limit, and Azure still has
	// a nextLink, we must NOT send a continuation token.
	nextLink := "https://management.azure.com/nextPage"
	client := &az.MockAzureClient{
		GetResponse: pageOf([]string{"/res/a", "/res/b", "/res/c"}, nextLink),
	}
	p := providerWithListResource(client, widgetResource())
	stream := &mockListStream{}

	require.NoError(t, p.List(makeListReq("", 3, 0, nil), stream))

	assert.Len(t, stream.results(), 3)
	assert.Empty(t, stream.continuationToken(), "limit exactly met must not produce a continuation")
}

func TestList_LimitCarriedAcrossPages(t *testing.T) {
	nextLink := "https://management.azure.com/nextPage"
	client := &sequentialAzureClient{
		getResponses: []map[string]any{
			pageOf([]string{"/res/a", "/res/b"}, nextLink),
			pageOf([]string{"/res/c", "/res/d", "/res/e"}, ""),
		},
	}
	p := providerWithListResource(client, widgetResource())

	// limit=4: page 1 returns 2, so 2 remain for page 2.
	stream1 := &mockListStream{}
	require.NoError(t, p.List(makeListReq("", 4, 0, nil), stream1))
	assert.Len(t, stream1.results(), 2)

	ct := stream1.continuationToken()
	require.NotEmpty(t, ct)

	decoded, err := decodeListContinuationToken(ct)
	require.NoError(t, err)
	require.NotNil(t, decoded.Remaining)
	assert.Equal(t, int64(2), *decoded.Remaining, "2 items still owed after page 1")

	// Page 2: only 2 of the 3 available items should be streamed.
	stream2 := &mockListStream{}
	require.NoError(t, p.List(makeListReq(ct, 4, 0, nil), stream2))
	assert.Len(t, stream2.results(), 2)
	assert.Empty(t, stream2.continuationToken(), "limit reached on page 2")
}

func TestList_PageSizeInjectedAsTopParam(t *testing.T) {
	topParam := resources.AzureAPIParameter{
		Name:     "$top",
		Location: "query",
		Value:    &resources.AzureAPIProperty{Type: "integer", SdkName: "top"},
	}
	client := &az.MockAzureClient{
		GetResponse: pageOf([]string{"/res/a"}, ""),
	}
	p := providerWithListResource(client, widgetResource(topParam))
	stream := &mockListStream{}

	require.NoError(t, p.List(makeListReq("", 0, 10, nil), stream))

	require.Len(t, client.GetQueryParams, 1)
	assert.Equal(t, float64(10), client.GetQueryParams[0]["$top"],
		"pageSize must be forwarded as $top query parameter")
}

func TestList_PageSizeClampedByLimit(t *testing.T) {
	topParam := resources.AzureAPIParameter{
		Name:     "$top",
		Location: "query",
		Value:    &resources.AzureAPIProperty{Type: "integer", SdkName: "top"},
	}
	client := &az.MockAzureClient{
		GetResponse: pageOf([]string{"/res/a", "/res/b"}, ""),
	}
	p := providerWithListResource(client, widgetResource(topParam))
	stream := &mockListStream{}

	// pageSize=100 but limit=5 — should clamp $top down to 5.
	require.NoError(t, p.List(makeListReq("", 5, 100, nil), stream))

	require.Len(t, client.GetQueryParams, 1)
	assert.Equal(t, float64(5), client.GetQueryParams[0]["$top"],
		"$top must be clamped to limit when limit < pageSize")
}

func TestList_LimitFallsBackToPageSizeWhenPageSizeUnset(t *testing.T) {
	topParam := resources.AzureAPIParameter{
		Name:     "$top",
		Location: "query",
		Value:    &resources.AzureAPIProperty{Type: "integer", SdkName: "top"},
	}
	client := &az.MockAzureClient{
		GetResponse: pageOf([]string{"/res/a"}, ""),
	}
	p := providerWithListResource(client, widgetResource(topParam))
	stream := &mockListStream{}

	// pageSize=0 but limit=7 — should use limit as the page size.
	require.NoError(t, p.List(makeListReq("", 7, 0, nil), stream))

	require.Len(t, client.GetQueryParams, 1)
	assert.Equal(t, float64(7), client.GetQueryParams[0]["$top"],
		"limit should be used as $top when pageSize is not set")
}

func TestList_SubscriptionIDFromProvider(t *testing.T) {
	client := &az.MockAzureClient{
		GetResponse: pageOf([]string{"/res/a"}, ""),
	}
	p := providerWithListResource(client, widgetResource())
	p.subscriptionID = "provider-sub-456"

	// Send a query without a subscriptionId — it should be filled from the provider config.
	q, _ := plugin.MarshalProperties(resource.PropertyMap{
		"resourceGroupName": resource.NewStringProperty("rg-test"),
	}, plugin.MarshalOptions{SkipNulls: true})
	stream := &mockListStream{}

	err := p.List(&rpc.ListRequest{Token: listTok, Query: q}, stream)
	require.NoError(t, err)

	require.Len(t, client.GetIds, 1)
	assert.Contains(t, client.GetIds[0], "provider-sub-456",
		"subscriptionId from provider config must be used when not in query")
}
