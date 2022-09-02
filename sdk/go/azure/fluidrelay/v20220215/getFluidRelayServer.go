


package v20220215

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFluidRelayServer(ctx *pulumi.Context, args *LookupFluidRelayServerArgs, opts ...pulumi.InvokeOption) (*LookupFluidRelayServerResult, error) {
	var rv LookupFluidRelayServerResult
	err := ctx.Invoke("azure-native:fluidrelay/v20220215:getFluidRelayServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFluidRelayServerArgs struct {
	FluidRelayServerName string `pulumi:"fluidRelayServerName"`
	ResourceGroup        string `pulumi:"resourceGroup"`
}


type LookupFluidRelayServerResult struct {
	FluidRelayEndpoints FluidRelayEndpointsResponse `pulumi:"fluidRelayEndpoints"`
	FrsTenantId         string                      `pulumi:"frsTenantId"`
	Id                  string                      `pulumi:"id"`
	Identity            *IdentityResponse           `pulumi:"identity"`
	Location            string                      `pulumi:"location"`
	Name                string                      `pulumi:"name"`
	ProvisioningState   *string                     `pulumi:"provisioningState"`
	SystemData          SystemDataResponse          `pulumi:"systemData"`
	Tags                map[string]string           `pulumi:"tags"`
	Type                string                      `pulumi:"type"`
}

func LookupFluidRelayServerOutput(ctx *pulumi.Context, args LookupFluidRelayServerOutputArgs, opts ...pulumi.InvokeOption) LookupFluidRelayServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFluidRelayServerResult, error) {
			args := v.(LookupFluidRelayServerArgs)
			r, err := LookupFluidRelayServer(ctx, &args, opts...)
			var s LookupFluidRelayServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFluidRelayServerResultOutput)
}

type LookupFluidRelayServerOutputArgs struct {
	FluidRelayServerName pulumi.StringInput `pulumi:"fluidRelayServerName"`
	ResourceGroup        pulumi.StringInput `pulumi:"resourceGroup"`
}

func (LookupFluidRelayServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFluidRelayServerArgs)(nil)).Elem()
}


type LookupFluidRelayServerResultOutput struct{ *pulumi.OutputState }

func (LookupFluidRelayServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFluidRelayServerResult)(nil)).Elem()
}

func (o LookupFluidRelayServerResultOutput) ToLookupFluidRelayServerResultOutput() LookupFluidRelayServerResultOutput {
	return o
}

func (o LookupFluidRelayServerResultOutput) ToLookupFluidRelayServerResultOutputWithContext(ctx context.Context) LookupFluidRelayServerResultOutput {
	return o
}

func (o LookupFluidRelayServerResultOutput) FluidRelayEndpoints() FluidRelayEndpointsResponseOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) FluidRelayEndpointsResponse { return v.FluidRelayEndpoints }).(FluidRelayEndpointsResponseOutput)
}

func (o LookupFluidRelayServerResultOutput) FrsTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) string { return v.FrsTenantId }).(pulumi.StringOutput)
}

func (o LookupFluidRelayServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFluidRelayServerResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupFluidRelayServerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupFluidRelayServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFluidRelayServerResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupFluidRelayServerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupFluidRelayServerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupFluidRelayServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFluidRelayServerResultOutput{})
}
