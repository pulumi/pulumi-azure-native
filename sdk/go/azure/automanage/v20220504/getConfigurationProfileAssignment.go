


package v20220504

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationProfileAssignment(ctx *pulumi.Context, args *LookupConfigurationProfileAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfileAssignmentResult, error) {
	var rv LookupConfigurationProfileAssignmentResult
	err := ctx.Invoke("azure-native:automanage/v20220504:getConfigurationProfileAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfileAssignmentArgs struct {
	ConfigurationProfileAssignmentName string `pulumi:"configurationProfileAssignmentName"`
	ResourceGroupName                  string `pulumi:"resourceGroupName"`
	VmName                             string `pulumi:"vmName"`
}


type LookupConfigurationProfileAssignmentResult struct {
	Id         string                                           `pulumi:"id"`
	ManagedBy  string                                           `pulumi:"managedBy"`
	Name       string                                           `pulumi:"name"`
	Properties ConfigurationProfileAssignmentPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                               `pulumi:"systemData"`
	Type       string                                           `pulumi:"type"`
}

func LookupConfigurationProfileAssignmentOutput(ctx *pulumi.Context, args LookupConfigurationProfileAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationProfileAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationProfileAssignmentResult, error) {
			args := v.(LookupConfigurationProfileAssignmentArgs)
			r, err := LookupConfigurationProfileAssignment(ctx, &args, opts...)
			var s LookupConfigurationProfileAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationProfileAssignmentResultOutput)
}

type LookupConfigurationProfileAssignmentOutputArgs struct {
	ConfigurationProfileAssignmentName pulumi.StringInput `pulumi:"configurationProfileAssignmentName"`
	ResourceGroupName                  pulumi.StringInput `pulumi:"resourceGroupName"`
	VmName                             pulumi.StringInput `pulumi:"vmName"`
}

func (LookupConfigurationProfileAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfileAssignmentArgs)(nil)).Elem()
}


type LookupConfigurationProfileAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationProfileAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfileAssignmentResult)(nil)).Elem()
}

func (o LookupConfigurationProfileAssignmentResultOutput) ToLookupConfigurationProfileAssignmentResultOutput() LookupConfigurationProfileAssignmentResultOutput {
	return o
}

func (o LookupConfigurationProfileAssignmentResultOutput) ToLookupConfigurationProfileAssignmentResultOutputWithContext(ctx context.Context) LookupConfigurationProfileAssignmentResultOutput {
	return o
}

func (o LookupConfigurationProfileAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfileAssignmentResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileAssignmentResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfileAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfileAssignmentResultOutput) Properties() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfileAssignmentResult) ConfigurationProfileAssignmentPropertiesResponse {
		return v.Properties
	}).(ConfigurationProfileAssignmentPropertiesResponseOutput)
}

func (o LookupConfigurationProfileAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfileAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConfigurationProfileAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationProfileAssignmentResultOutput{})
}
