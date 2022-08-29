


package v20210430preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationProfileHCRPAssignment(ctx *pulumi.Context, args *LookupConfigurationProfileHCRPAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfileHCRPAssignmentResult, error) {
	var rv LookupConfigurationProfileHCRPAssignmentResult
	err := ctx.Invoke("azure-native:automanage/v20210430preview:getConfigurationProfileHCRPAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfileHCRPAssignmentArgs struct {
	ConfigurationProfileAssignmentName string `pulumi:"configurationProfileAssignmentName"`
	MachineName                        string `pulumi:"machineName"`
	ResourceGroupName                  string `pulumi:"resourceGroupName"`
}


type LookupConfigurationProfileHCRPAssignmentResult struct {
	Id         string                                           `pulumi:"id"`
	Name       string                                           `pulumi:"name"`
	Properties ConfigurationProfileAssignmentPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                               `pulumi:"systemData"`
	Type       string                                           `pulumi:"type"`
}

func LookupConfigurationProfileHCRPAssignmentOutput(ctx *pulumi.Context, args LookupConfigurationProfileHCRPAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationProfileHCRPAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationProfileHCRPAssignmentResult, error) {
			args := v.(LookupConfigurationProfileHCRPAssignmentArgs)
			r, err := LookupConfigurationProfileHCRPAssignment(ctx, &args, opts...)
			var s LookupConfigurationProfileHCRPAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationProfileHCRPAssignmentResultOutput)
}

type LookupConfigurationProfileHCRPAssignmentOutputArgs struct {
	ConfigurationProfileAssignmentName pulumi.StringInput `pulumi:"configurationProfileAssignmentName"`
	MachineName                        pulumi.StringInput `pulumi:"machineName"`
	ResourceGroupName                  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConfigurationProfileHCRPAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfileHCRPAssignmentArgs)(nil)).Elem()
}


type LookupConfigurationProfileHCRPAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationProfileHCRPAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfileHCRPAssignmentResult)(nil)).Elem()
}

func (o LookupConfigurationProfileHCRPAssignmentResultOutput) ToLookupConfigurationProfileHCRPAssignmentResultOutput() LookupConfigurationProfileHCRPAssignmentResultOutput {
	return o
}

func (o LookupConfigurationProfileHCRPAssignmentResultOutput) ToLookupConfigurationProfileHCRPAssignmentResultOutputWithContext(ctx context.Context) LookupConfigurationProfileHCRPAssignmentResultOutput {
	return o
}

func (o LookupConfigurationProfileHCRPAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileHCRPAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfileHCRPAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileHCRPAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfileHCRPAssignmentResultOutput) Properties() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfileHCRPAssignmentResult) ConfigurationProfileAssignmentPropertiesResponse {
		return v.Properties
	}).(ConfigurationProfileAssignmentPropertiesResponseOutput)
}

func (o LookupConfigurationProfileHCRPAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfileHCRPAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConfigurationProfileHCRPAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileHCRPAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationProfileHCRPAssignmentResultOutput{})
}
