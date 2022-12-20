


package v20210430preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationProfileHCIAssignment(ctx *pulumi.Context, args *LookupConfigurationProfileHCIAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfileHCIAssignmentResult, error) {
	var rv LookupConfigurationProfileHCIAssignmentResult
	err := ctx.Invoke("azure-native:automanage/v20210430preview:getConfigurationProfileHCIAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfileHCIAssignmentArgs struct {
	ClusterName                        string `pulumi:"clusterName"`
	ConfigurationProfileAssignmentName string `pulumi:"configurationProfileAssignmentName"`
	ResourceGroupName                  string `pulumi:"resourceGroupName"`
}


type LookupConfigurationProfileHCIAssignmentResult struct {
	Id         string                                           `pulumi:"id"`
	Name       string                                           `pulumi:"name"`
	Properties ConfigurationProfileAssignmentPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                               `pulumi:"systemData"`
	Type       string                                           `pulumi:"type"`
}

func LookupConfigurationProfileHCIAssignmentOutput(ctx *pulumi.Context, args LookupConfigurationProfileHCIAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationProfileHCIAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationProfileHCIAssignmentResult, error) {
			args := v.(LookupConfigurationProfileHCIAssignmentArgs)
			r, err := LookupConfigurationProfileHCIAssignment(ctx, &args, opts...)
			var s LookupConfigurationProfileHCIAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationProfileHCIAssignmentResultOutput)
}

type LookupConfigurationProfileHCIAssignmentOutputArgs struct {
	ClusterName                        pulumi.StringInput `pulumi:"clusterName"`
	ConfigurationProfileAssignmentName pulumi.StringInput `pulumi:"configurationProfileAssignmentName"`
	ResourceGroupName                  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConfigurationProfileHCIAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfileHCIAssignmentArgs)(nil)).Elem()
}


type LookupConfigurationProfileHCIAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationProfileHCIAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfileHCIAssignmentResult)(nil)).Elem()
}

func (o LookupConfigurationProfileHCIAssignmentResultOutput) ToLookupConfigurationProfileHCIAssignmentResultOutput() LookupConfigurationProfileHCIAssignmentResultOutput {
	return o
}

func (o LookupConfigurationProfileHCIAssignmentResultOutput) ToLookupConfigurationProfileHCIAssignmentResultOutputWithContext(ctx context.Context) LookupConfigurationProfileHCIAssignmentResultOutput {
	return o
}

func (o LookupConfigurationProfileHCIAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileHCIAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfileHCIAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileHCIAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfileHCIAssignmentResultOutput) Properties() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfileHCIAssignmentResult) ConfigurationProfileAssignmentPropertiesResponse {
		return v.Properties
	}).(ConfigurationProfileAssignmentPropertiesResponseOutput)
}

func (o LookupConfigurationProfileHCIAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfileHCIAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConfigurationProfileHCIAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileHCIAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationProfileHCIAssignmentResultOutput{})
}
