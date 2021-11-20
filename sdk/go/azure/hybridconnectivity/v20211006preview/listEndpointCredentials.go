


package v20211006preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEndpointCredentials(ctx *pulumi.Context, args *ListEndpointCredentialsArgs, opts ...pulumi.InvokeOption) (*ListEndpointCredentialsResult, error) {
	var rv ListEndpointCredentialsResult
	err := ctx.Invoke("azure-native:hybridconnectivity/v20211006preview:listEndpointCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEndpointCredentialsArgs struct {
	EndpointName string `pulumi:"endpointName"`
	Expiresin    *int   `pulumi:"expiresin"`
	ResourceUri  string `pulumi:"resourceUri"`
}


type ListEndpointCredentialsResult struct {
	AccessKey            string   `pulumi:"accessKey"`
	ExpiresOn            *float64 `pulumi:"expiresOn"`
	HybridConnectionName string   `pulumi:"hybridConnectionName"`
	NamespaceName        string   `pulumi:"namespaceName"`
	NamespaceNameSuffix  string   `pulumi:"namespaceNameSuffix"`
}
