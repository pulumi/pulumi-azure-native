


package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountBatchConfiguration(ctx *pulumi.Context, args *LookupIntegrationAccountBatchConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountBatchConfigurationResult, error) {
	var rv LookupIntegrationAccountBatchConfigurationResult
	err := ctx.Invoke("azure-native:logic:getIntegrationAccountBatchConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountBatchConfigurationArgs struct {
	BatchConfigurationName string `pulumi:"batchConfigurationName"`
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupIntegrationAccountBatchConfigurationResult struct {
	Id         string                               `pulumi:"id"`
	Location   *string                              `pulumi:"location"`
	Name       string                               `pulumi:"name"`
	Properties BatchConfigurationPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                    `pulumi:"tags"`
	Type       string                               `pulumi:"type"`
}

func LookupIntegrationAccountBatchConfigurationOutput(ctx *pulumi.Context, args LookupIntegrationAccountBatchConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationAccountBatchConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationAccountBatchConfigurationResult, error) {
			args := v.(LookupIntegrationAccountBatchConfigurationArgs)
			r, err := LookupIntegrationAccountBatchConfiguration(ctx, &args, opts...)
			var s LookupIntegrationAccountBatchConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationAccountBatchConfigurationResultOutput)
}

type LookupIntegrationAccountBatchConfigurationOutputArgs struct {
	BatchConfigurationName pulumi.StringInput `pulumi:"batchConfigurationName"`
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIntegrationAccountBatchConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountBatchConfigurationArgs)(nil)).Elem()
}


type LookupIntegrationAccountBatchConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationAccountBatchConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountBatchConfigurationResult)(nil)).Elem()
}

func (o LookupIntegrationAccountBatchConfigurationResultOutput) ToLookupIntegrationAccountBatchConfigurationResultOutput() LookupIntegrationAccountBatchConfigurationResultOutput {
	return o
}

func (o LookupIntegrationAccountBatchConfigurationResultOutput) ToLookupIntegrationAccountBatchConfigurationResultOutputWithContext(ctx context.Context) LookupIntegrationAccountBatchConfigurationResultOutput {
	return o
}

func (o LookupIntegrationAccountBatchConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountBatchConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountBatchConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountBatchConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountBatchConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountBatchConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountBatchConfigurationResultOutput) Properties() BatchConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupIntegrationAccountBatchConfigurationResult) BatchConfigurationPropertiesResponse {
		return v.Properties
	}).(BatchConfigurationPropertiesResponseOutput)
}

func (o LookupIntegrationAccountBatchConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationAccountBatchConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIntegrationAccountBatchConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountBatchConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationAccountBatchConfigurationResultOutput{})
}
