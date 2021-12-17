


package appplatform

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApp(ctx *pulumi.Context, args *LookupAppArgs, opts ...pulumi.InvokeOption) (*LookupAppResult, error) {
	var rv LookupAppResult
	err := ctx.Invoke("azure-native:appplatform:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAppArgs struct {
	AppName           string  `pulumi:"appName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	SyncStatus        *string `pulumi:"syncStatus"`
}


type LookupAppResult struct {
	Id         string                             `pulumi:"id"`
	Identity   *ManagedIdentityPropertiesResponse `pulumi:"identity"`
	Location   *string                            `pulumi:"location"`
	Name       string                             `pulumi:"name"`
	Properties AppResourcePropertiesResponse      `pulumi:"properties"`
	Type       string                             `pulumi:"type"`
}


func (val *LookupAppResult) Defaults() *LookupAppResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
