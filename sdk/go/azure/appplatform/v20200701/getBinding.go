


package v20200701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBinding(ctx *pulumi.Context, args *LookupBindingArgs, opts ...pulumi.InvokeOption) (*LookupBindingResult, error) {
	var rv LookupBindingResult
	err := ctx.Invoke("azure-native:appplatform/v20200701:getBinding", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBindingArgs struct {
	AppName           string `pulumi:"appName"`
	BindingName       string `pulumi:"bindingName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupBindingResult struct {
	Id         string                            `pulumi:"id"`
	Name       string                            `pulumi:"name"`
	Properties BindingResourcePropertiesResponse `pulumi:"properties"`
	Type       string                            `pulumi:"type"`
}
