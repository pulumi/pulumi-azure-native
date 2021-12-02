


package v20150801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConnectionKeys(ctx *pulumi.Context, args *ListConnectionKeysArgs, opts ...pulumi.InvokeOption) (*ListConnectionKeysResult, error) {
	var rv ListConnectionKeysResult
	err := ctx.Invoke("azure-native:web/v20150801preview:listConnectionKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectionKeysArgs struct {
	ConnectionName    string            `pulumi:"connectionName"`
	Id                *string           `pulumi:"id"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
	ValidityTimeSpan  *string           `pulumi:"validityTimeSpan"`
}

type ListConnectionKeysResult struct {
	ConnectionKey   *string                `pulumi:"connectionKey"`
	ParameterValues map[string]interface{} `pulumi:"parameterValues"`
}
