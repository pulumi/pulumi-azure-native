


package v20160301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFunction(ctx *pulumi.Context, args *LookupFunctionArgs, opts ...pulumi.InvokeOption) (*LookupFunctionResult, error) {
	var rv LookupFunctionResult
	err := ctx.Invoke("azure-native:streamanalytics/v20160301:getFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFunctionArgs struct {
	FunctionName      string `pulumi:"functionName"`
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFunctionResult struct {
	Id         string                           `pulumi:"id"`
	Name       *string                          `pulumi:"name"`
	Properties ScalarFunctionPropertiesResponse `pulumi:"properties"`
	Type       string                           `pulumi:"type"`
}
