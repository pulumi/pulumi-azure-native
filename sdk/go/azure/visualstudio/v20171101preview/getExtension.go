


package v20171101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExtension(ctx *pulumi.Context, args *LookupExtensionArgs, opts ...pulumi.InvokeOption) (*LookupExtensionResult, error) {
	var rv LookupExtensionResult
	err := ctx.Invoke("azure-native:visualstudio/v20171101preview:getExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtensionArgs struct {
	AccountResourceName   string `pulumi:"accountResourceName"`
	ExtensionResourceName string `pulumi:"extensionResourceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupExtensionResult struct {
	Id         string                         `pulumi:"id"`
	Location   *string                        `pulumi:"location"`
	Name       string                         `pulumi:"name"`
	Plan       *ExtensionResourcePlanResponse `pulumi:"plan"`
	Properties map[string]string              `pulumi:"properties"`
	Tags       map[string]string              `pulumi:"tags"`
	Type       string                         `pulumi:"type"`
}
