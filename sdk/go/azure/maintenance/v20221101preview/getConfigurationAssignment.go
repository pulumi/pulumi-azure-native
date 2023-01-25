


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationAssignment(ctx *pulumi.Context, args *LookupConfigurationAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationAssignmentResult, error) {
	var rv LookupConfigurationAssignmentResult
	err := ctx.Invoke("azure-native:maintenance/v20221101preview:getConfigurationAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationAssignmentArgs struct {
	ConfigurationAssignmentName string `pulumi:"configurationAssignmentName"`
	ProviderName                string `pulumi:"providerName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	ResourceName                string `pulumi:"resourceName"`
	ResourceType                string `pulumi:"resourceType"`
}


type LookupConfigurationAssignmentResult struct {
	Id                         string             `pulumi:"id"`
	Location                   *string            `pulumi:"location"`
	MaintenanceConfigurationId *string            `pulumi:"maintenanceConfigurationId"`
	Name                       string             `pulumi:"name"`
	ResourceId                 *string            `pulumi:"resourceId"`
	SystemData                 SystemDataResponse `pulumi:"systemData"`
	Type                       string             `pulumi:"type"`
}

func LookupConfigurationAssignmentOutput(ctx *pulumi.Context, args LookupConfigurationAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationAssignmentResult, error) {
			args := v.(LookupConfigurationAssignmentArgs)
			r, err := LookupConfigurationAssignment(ctx, &args, opts...)
			var s LookupConfigurationAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationAssignmentResultOutput)
}

type LookupConfigurationAssignmentOutputArgs struct {
	ConfigurationAssignmentName pulumi.StringInput `pulumi:"configurationAssignmentName"`
	ProviderName                pulumi.StringInput `pulumi:"providerName"`
	ResourceGroupName           pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                pulumi.StringInput `pulumi:"resourceName"`
	ResourceType                pulumi.StringInput `pulumi:"resourceType"`
}

func (LookupConfigurationAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationAssignmentArgs)(nil)).Elem()
}


type LookupConfigurationAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationAssignmentResult)(nil)).Elem()
}

func (o LookupConfigurationAssignmentResultOutput) ToLookupConfigurationAssignmentResultOutput() LookupConfigurationAssignmentResultOutput {
	return o
}

func (o LookupConfigurationAssignmentResultOutput) ToLookupConfigurationAssignmentResultOutputWithContext(ctx context.Context) LookupConfigurationAssignmentResultOutput {
	return o
}

func (o LookupConfigurationAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigurationAssignmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationAssignmentResultOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentResult) *string { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigurationAssignmentResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConfigurationAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationAssignmentResultOutput{})
}
