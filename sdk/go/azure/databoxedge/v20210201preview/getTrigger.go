


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupTrigger(ctx *pulumi.Context, args *LookupTriggerArgs, opts ...pulumi.InvokeOption) (*LookupTriggerResult, error) {
	var rv LookupTriggerResult
	err := ctx.Invoke("azure-native:databoxedge/v20210201preview:getTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTriggerArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupTriggerResult struct {
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}
