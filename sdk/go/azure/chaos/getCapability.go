


package chaos

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCapability(ctx *pulumi.Context, args *LookupCapabilityArgs, opts ...pulumi.InvokeOption) (*LookupCapabilityResult, error) {
	var rv LookupCapabilityResult
	err := ctx.Invoke("azure-native:chaos:getCapability", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCapabilityArgs struct {
	CapabilityName          string `pulumi:"capabilityName"`
	ParentProviderNamespace string `pulumi:"parentProviderNamespace"`
	ParentResourceName      string `pulumi:"parentResourceName"`
	ParentResourceType      string `pulumi:"parentResourceType"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	TargetName              string `pulumi:"targetName"`
}


type LookupCapabilityResult struct {
	Id         string                       `pulumi:"id"`
	Name       string                       `pulumi:"name"`
	Properties CapabilityPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse           `pulumi:"systemData"`
	Type       string                       `pulumi:"type"`
}
