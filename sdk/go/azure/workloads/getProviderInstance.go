


package workloads

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProviderInstance(ctx *pulumi.Context, args *LookupProviderInstanceArgs, opts ...pulumi.InvokeOption) (*LookupProviderInstanceResult, error) {
	var rv LookupProviderInstanceResult
	err := ctx.Invoke("azure-native:workloads:getProviderInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProviderInstanceArgs struct {
	MonitorName          string `pulumi:"monitorName"`
	ProviderInstanceName string `pulumi:"providerInstanceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupProviderInstanceResult struct {
	Errors            ProviderInstancePropertiesResponseErrors `pulumi:"errors"`
	Id                string                                   `pulumi:"id"`
	Identity          *UserAssignedServiceIdentityResponse     `pulumi:"identity"`
	Name              string                                   `pulumi:"name"`
	ProviderSettings  interface{}                              `pulumi:"providerSettings"`
	ProvisioningState string                                   `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                       `pulumi:"systemData"`
	Type              string                                   `pulumi:"type"`
}

func LookupProviderInstanceOutput(ctx *pulumi.Context, args LookupProviderInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupProviderInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProviderInstanceResult, error) {
			args := v.(LookupProviderInstanceArgs)
			r, err := LookupProviderInstance(ctx, &args, opts...)
			var s LookupProviderInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProviderInstanceResultOutput)
}

type LookupProviderInstanceOutputArgs struct {
	MonitorName          pulumi.StringInput `pulumi:"monitorName"`
	ProviderInstanceName pulumi.StringInput `pulumi:"providerInstanceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProviderInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderInstanceArgs)(nil)).Elem()
}


type LookupProviderInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupProviderInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderInstanceResult)(nil)).Elem()
}

func (o LookupProviderInstanceResultOutput) ToLookupProviderInstanceResultOutput() LookupProviderInstanceResultOutput {
	return o
}

func (o LookupProviderInstanceResultOutput) ToLookupProviderInstanceResultOutputWithContext(ctx context.Context) LookupProviderInstanceResultOutput {
	return o
}

func (o LookupProviderInstanceResultOutput) Errors() ProviderInstancePropertiesResponseErrorsOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) ProviderInstancePropertiesResponseErrors { return v.Errors }).(ProviderInstancePropertiesResponseErrorsOutput)
}

func (o LookupProviderInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProviderInstanceResultOutput) Identity() UserAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) *UserAssignedServiceIdentityResponse { return v.Identity }).(UserAssignedServiceIdentityResponsePtrOutput)
}

func (o LookupProviderInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProviderInstanceResultOutput) ProviderSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) interface{} { return v.ProviderSettings }).(pulumi.AnyOutput)
}

func (o LookupProviderInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupProviderInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupProviderInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProviderInstanceResultOutput{})
}
