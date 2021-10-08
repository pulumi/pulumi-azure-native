


package v20210630preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	var rv LookupEnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20210630preview:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	EnvironmentName   string  `pulumi:"environmentName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupEnvironmentResult struct {
	Id       string            `pulumi:"id"`
	Kind     string            `pulumi:"kind"`
	Location string            `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Sku      SkuResponse       `pulumi:"sku"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}
