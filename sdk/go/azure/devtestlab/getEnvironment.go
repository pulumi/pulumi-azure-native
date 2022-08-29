


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	var rv LookupEnvironmentResult
	err := ctx.Invoke("azure-native:devtestlab:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	UserName          string  `pulumi:"userName"`
}


type LookupEnvironmentResult struct {
	ArmTemplateDisplayName *string                                  `pulumi:"armTemplateDisplayName"`
	CreatedByUser          string                                   `pulumi:"createdByUser"`
	DeploymentProperties   *EnvironmentDeploymentPropertiesResponse `pulumi:"deploymentProperties"`
	Id                     string                                   `pulumi:"id"`
	Location               *string                                  `pulumi:"location"`
	Name                   string                                   `pulumi:"name"`
	ProvisioningState      string                                   `pulumi:"provisioningState"`
	ResourceGroupId        string                                   `pulumi:"resourceGroupId"`
	Tags                   map[string]string                        `pulumi:"tags"`
	Type                   string                                   `pulumi:"type"`
	UniqueIdentifier       string                                   `pulumi:"uniqueIdentifier"`
}

func LookupEnvironmentOutput(ctx *pulumi.Context, args LookupEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentResult, error) {
			args := v.(LookupEnvironmentArgs)
			r, err := LookupEnvironment(ctx, &args, opts...)
			var s LookupEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentResultOutput)
}

type LookupEnvironmentOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	UserName          pulumi.StringInput    `pulumi:"userName"`
}

func (LookupEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentArgs)(nil)).Elem()
}


type LookupEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentResult)(nil)).Elem()
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutput() LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutputWithContext(ctx context.Context) LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ArmTemplateDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.ArmTemplateDisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) CreatedByUser() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.CreatedByUser }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) DeploymentProperties() EnvironmentDeploymentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *EnvironmentDeploymentPropertiesResponse {
		return v.DeploymentProperties
	}).(EnvironmentDeploymentPropertiesResponsePtrOutput)
}

func (o LookupEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.ResourceGroupId }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentResultOutput{})
}
