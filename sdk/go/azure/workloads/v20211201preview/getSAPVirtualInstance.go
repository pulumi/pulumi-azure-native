


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSAPVirtualInstance(ctx *pulumi.Context, args *LookupSAPVirtualInstanceArgs, opts ...pulumi.InvokeOption) (*LookupSAPVirtualInstanceResult, error) {
	var rv LookupSAPVirtualInstanceResult
	err := ctx.Invoke("azure-native:workloads/v20211201preview:getSAPVirtualInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSAPVirtualInstanceArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SapVirtualInstanceName string `pulumi:"sapVirtualInstanceName"`
}


type LookupSAPVirtualInstanceResult struct {
	Configuration                     interface{}                          `pulumi:"configuration"`
	Environment                       string                               `pulumi:"environment"`
	Errors                            SAPVirtualInstanceErrorResponse      `pulumi:"errors"`
	Health                            string                               `pulumi:"health"`
	Id                                string                               `pulumi:"id"`
	Identity                          *UserAssignedServiceIdentityResponse `pulumi:"identity"`
	Location                          string                               `pulumi:"location"`
	ManagedResourceGroupConfiguration *ManagedRGConfigurationResponse      `pulumi:"managedResourceGroupConfiguration"`
	Name                              string                               `pulumi:"name"`
	ProvisioningState                 string                               `pulumi:"provisioningState"`
	SapProduct                        string                               `pulumi:"sapProduct"`
	State                             string                               `pulumi:"state"`
	Status                            string                               `pulumi:"status"`
	SystemData                        SystemDataResponse                   `pulumi:"systemData"`
	Tags                              map[string]string                    `pulumi:"tags"`
	Type                              string                               `pulumi:"type"`
}

func LookupSAPVirtualInstanceOutput(ctx *pulumi.Context, args LookupSAPVirtualInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupSAPVirtualInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSAPVirtualInstanceResult, error) {
			args := v.(LookupSAPVirtualInstanceArgs)
			r, err := LookupSAPVirtualInstance(ctx, &args, opts...)
			var s LookupSAPVirtualInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSAPVirtualInstanceResultOutput)
}

type LookupSAPVirtualInstanceOutputArgs struct {
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	SapVirtualInstanceName pulumi.StringInput `pulumi:"sapVirtualInstanceName"`
}

func (LookupSAPVirtualInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSAPVirtualInstanceArgs)(nil)).Elem()
}


type LookupSAPVirtualInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupSAPVirtualInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSAPVirtualInstanceResult)(nil)).Elem()
}

func (o LookupSAPVirtualInstanceResultOutput) ToLookupSAPVirtualInstanceResultOutput() LookupSAPVirtualInstanceResultOutput {
	return o
}

func (o LookupSAPVirtualInstanceResultOutput) ToLookupSAPVirtualInstanceResultOutputWithContext(ctx context.Context) LookupSAPVirtualInstanceResultOutput {
	return o
}

func (o LookupSAPVirtualInstanceResultOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) interface{} { return v.Configuration }).(pulumi.AnyOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) string { return v.Environment }).(pulumi.StringOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) Errors() SAPVirtualInstanceErrorResponseOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) SAPVirtualInstanceErrorResponse { return v.Errors }).(SAPVirtualInstanceErrorResponseOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) string { return v.Health }).(pulumi.StringOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) Identity() UserAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) *UserAssignedServiceIdentityResponse { return v.Identity }).(UserAssignedServiceIdentityResponsePtrOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) ManagedResourceGroupConfiguration() ManagedRGConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) *ManagedRGConfigurationResponse {
		return v.ManagedResourceGroupConfiguration
	}).(ManagedRGConfigurationResponsePtrOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) SapProduct() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) string { return v.SapProduct }).(pulumi.StringOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSAPVirtualInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPVirtualInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSAPVirtualInstanceResultOutput{})
}
