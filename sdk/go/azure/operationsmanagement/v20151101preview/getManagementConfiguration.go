


package v20151101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementConfiguration(ctx *pulumi.Context, args *LookupManagementConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupManagementConfigurationResult, error) {
	var rv LookupManagementConfigurationResult
	err := ctx.Invoke("azure-native:operationsmanagement/v20151101preview:getManagementConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementConfigurationArgs struct {
	ManagementConfigurationName string `pulumi:"managementConfigurationName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
}


type LookupManagementConfigurationResult struct {
	Id         string                                    `pulumi:"id"`
	Location   *string                                   `pulumi:"location"`
	Name       string                                    `pulumi:"name"`
	Properties ManagementConfigurationPropertiesResponse `pulumi:"properties"`
	Type       string                                    `pulumi:"type"`
}

func LookupManagementConfigurationOutput(ctx *pulumi.Context, args LookupManagementConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupManagementConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementConfigurationResult, error) {
			args := v.(LookupManagementConfigurationArgs)
			r, err := LookupManagementConfiguration(ctx, &args, opts...)
			var s LookupManagementConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementConfigurationResultOutput)
}

type LookupManagementConfigurationOutputArgs struct {
	ManagementConfigurationName pulumi.StringInput `pulumi:"managementConfigurationName"`
	ResourceGroupName           pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagementConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementConfigurationArgs)(nil)).Elem()
}


type LookupManagementConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupManagementConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementConfigurationResult)(nil)).Elem()
}

func (o LookupManagementConfigurationResultOutput) ToLookupManagementConfigurationResultOutput() LookupManagementConfigurationResultOutput {
	return o
}

func (o LookupManagementConfigurationResultOutput) ToLookupManagementConfigurationResultOutputWithContext(ctx context.Context) LookupManagementConfigurationResultOutput {
	return o
}

func (o LookupManagementConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagementConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupManagementConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagementConfigurationResultOutput) Properties() ManagementConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupManagementConfigurationResult) ManagementConfigurationPropertiesResponse {
		return v.Properties
	}).(ManagementConfigurationPropertiesResponseOutput)
}

func (o LookupManagementConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementConfigurationResultOutput{})
}
