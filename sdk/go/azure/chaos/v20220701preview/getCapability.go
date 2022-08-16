


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCapability(ctx *pulumi.Context, args *LookupCapabilityArgs, opts ...pulumi.InvokeOption) (*LookupCapabilityResult, error) {
	var rv LookupCapabilityResult
	err := ctx.Invoke("azure-native:chaos/v20220701preview:getCapability", args, &rv, opts...)
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

func LookupCapabilityOutput(ctx *pulumi.Context, args LookupCapabilityOutputArgs, opts ...pulumi.InvokeOption) LookupCapabilityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCapabilityResult, error) {
			args := v.(LookupCapabilityArgs)
			r, err := LookupCapability(ctx, &args, opts...)
			var s LookupCapabilityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCapabilityResultOutput)
}

type LookupCapabilityOutputArgs struct {
	CapabilityName          pulumi.StringInput `pulumi:"capabilityName"`
	ParentProviderNamespace pulumi.StringInput `pulumi:"parentProviderNamespace"`
	ParentResourceName      pulumi.StringInput `pulumi:"parentResourceName"`
	ParentResourceType      pulumi.StringInput `pulumi:"parentResourceType"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	TargetName              pulumi.StringInput `pulumi:"targetName"`
}

func (LookupCapabilityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapabilityArgs)(nil)).Elem()
}


type LookupCapabilityResultOutput struct{ *pulumi.OutputState }

func (LookupCapabilityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapabilityResult)(nil)).Elem()
}

func (o LookupCapabilityResultOutput) ToLookupCapabilityResultOutput() LookupCapabilityResultOutput {
	return o
}

func (o LookupCapabilityResultOutput) ToLookupCapabilityResultOutputWithContext(ctx context.Context) LookupCapabilityResultOutput {
	return o
}

func (o LookupCapabilityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapabilityResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCapabilityResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapabilityResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCapabilityResultOutput) Properties() CapabilityPropertiesResponseOutput {
	return o.ApplyT(func(v LookupCapabilityResult) CapabilityPropertiesResponse { return v.Properties }).(CapabilityPropertiesResponseOutput)
}

func (o LookupCapabilityResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCapabilityResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCapabilityResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapabilityResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCapabilityResultOutput{})
}
