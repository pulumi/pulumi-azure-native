


package v20220602preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFleet(ctx *pulumi.Context, args *LookupFleetArgs, opts ...pulumi.InvokeOption) (*LookupFleetResult, error) {
	var rv LookupFleetResult
	err := ctx.Invoke("azure-native:containerservice/v20220602preview:getFleet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFleetArgs struct {
	FleetName         string `pulumi:"fleetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFleetResult struct {
	Etag              string                   `pulumi:"etag"`
	HubProfile        *FleetHubProfileResponse `pulumi:"hubProfile"`
	Id                string                   `pulumi:"id"`
	Location          string                   `pulumi:"location"`
	Name              string                   `pulumi:"name"`
	ProvisioningState string                   `pulumi:"provisioningState"`
	SystemData        SystemDataResponse       `pulumi:"systemData"`
	Tags              map[string]string        `pulumi:"tags"`
	Type              string                   `pulumi:"type"`
}

func LookupFleetOutput(ctx *pulumi.Context, args LookupFleetOutputArgs, opts ...pulumi.InvokeOption) LookupFleetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFleetResult, error) {
			args := v.(LookupFleetArgs)
			r, err := LookupFleet(ctx, &args, opts...)
			var s LookupFleetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFleetResultOutput)
}

type LookupFleetOutputArgs struct {
	FleetName         pulumi.StringInput `pulumi:"fleetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFleetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetArgs)(nil)).Elem()
}


type LookupFleetResultOutput struct{ *pulumi.OutputState }

func (LookupFleetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetResult)(nil)).Elem()
}

func (o LookupFleetResultOutput) ToLookupFleetResultOutput() LookupFleetResultOutput {
	return o
}

func (o LookupFleetResultOutput) ToLookupFleetResultOutputWithContext(ctx context.Context) LookupFleetResultOutput {
	return o
}

func (o LookupFleetResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupFleetResultOutput) HubProfile() FleetHubProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *FleetHubProfileResponse { return v.HubProfile }).(FleetHubProfileResponsePtrOutput)
}

func (o LookupFleetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFleetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupFleetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFleetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupFleetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFleetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupFleetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFleetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupFleetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFleetResultOutput{})
}
