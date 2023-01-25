


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationAssignmentParent(ctx *pulumi.Context, args *LookupConfigurationAssignmentParentArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationAssignmentParentResult, error) {
	var rv LookupConfigurationAssignmentParentResult
	err := ctx.Invoke("azure-native:maintenance/v20221101preview:getConfigurationAssignmentParent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationAssignmentParentArgs struct {
	ConfigurationAssignmentName string `pulumi:"configurationAssignmentName"`
	ProviderName                string `pulumi:"providerName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	ResourceName                string `pulumi:"resourceName"`
	ResourceParentName          string `pulumi:"resourceParentName"`
	ResourceParentType          string `pulumi:"resourceParentType"`
	ResourceType                string `pulumi:"resourceType"`
}


type LookupConfigurationAssignmentParentResult struct {
	Id                         string             `pulumi:"id"`
	Location                   *string            `pulumi:"location"`
	MaintenanceConfigurationId *string            `pulumi:"maintenanceConfigurationId"`
	Name                       string             `pulumi:"name"`
	ResourceId                 *string            `pulumi:"resourceId"`
	SystemData                 SystemDataResponse `pulumi:"systemData"`
	Type                       string             `pulumi:"type"`
}

func LookupConfigurationAssignmentParentOutput(ctx *pulumi.Context, args LookupConfigurationAssignmentParentOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationAssignmentParentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationAssignmentParentResult, error) {
			args := v.(LookupConfigurationAssignmentParentArgs)
			r, err := LookupConfigurationAssignmentParent(ctx, &args, opts...)
			var s LookupConfigurationAssignmentParentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationAssignmentParentResultOutput)
}

type LookupConfigurationAssignmentParentOutputArgs struct {
	ConfigurationAssignmentName pulumi.StringInput `pulumi:"configurationAssignmentName"`
	ProviderName                pulumi.StringInput `pulumi:"providerName"`
	ResourceGroupName           pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                pulumi.StringInput `pulumi:"resourceName"`
	ResourceParentName          pulumi.StringInput `pulumi:"resourceParentName"`
	ResourceParentType          pulumi.StringInput `pulumi:"resourceParentType"`
	ResourceType                pulumi.StringInput `pulumi:"resourceType"`
}

func (LookupConfigurationAssignmentParentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationAssignmentParentArgs)(nil)).Elem()
}


type LookupConfigurationAssignmentParentResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationAssignmentParentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationAssignmentParentResult)(nil)).Elem()
}

func (o LookupConfigurationAssignmentParentResultOutput) ToLookupConfigurationAssignmentParentResultOutput() LookupConfigurationAssignmentParentResultOutput {
	return o
}

func (o LookupConfigurationAssignmentParentResultOutput) ToLookupConfigurationAssignmentParentResultOutputWithContext(ctx context.Context) LookupConfigurationAssignmentParentResultOutput {
	return o
}

func (o LookupConfigurationAssignmentParentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigurationAssignmentParentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationAssignmentParentResultOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) *string { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationAssignmentParentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigurationAssignmentParentResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationAssignmentParentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConfigurationAssignmentParentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationAssignmentParentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationAssignmentParentResultOutput{})
}
