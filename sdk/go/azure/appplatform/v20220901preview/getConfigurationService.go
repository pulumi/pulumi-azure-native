


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationService(ctx *pulumi.Context, args *LookupConfigurationServiceArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationServiceResult, error) {
	var rv LookupConfigurationServiceResult
	err := ctx.Invoke("azure-native:appplatform/v20220901preview:getConfigurationService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationServiceArgs struct {
	ConfigurationServiceName string `pulumi:"configurationServiceName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
	ServiceName              string `pulumi:"serviceName"`
}


type LookupConfigurationServiceResult struct {
	Id         string                                 `pulumi:"id"`
	Name       string                                 `pulumi:"name"`
	Properties ConfigurationServicePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                     `pulumi:"systemData"`
	Type       string                                 `pulumi:"type"`
}

func LookupConfigurationServiceOutput(ctx *pulumi.Context, args LookupConfigurationServiceOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationServiceResult, error) {
			args := v.(LookupConfigurationServiceArgs)
			r, err := LookupConfigurationService(ctx, &args, opts...)
			var s LookupConfigurationServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationServiceResultOutput)
}

type LookupConfigurationServiceOutputArgs struct {
	ConfigurationServiceName pulumi.StringInput `pulumi:"configurationServiceName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName              pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupConfigurationServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationServiceArgs)(nil)).Elem()
}


type LookupConfigurationServiceResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationServiceResult)(nil)).Elem()
}

func (o LookupConfigurationServiceResultOutput) ToLookupConfigurationServiceResultOutput() LookupConfigurationServiceResultOutput {
	return o
}

func (o LookupConfigurationServiceResultOutput) ToLookupConfigurationServiceResultOutputWithContext(ctx context.Context) LookupConfigurationServiceResultOutput {
	return o
}

func (o LookupConfigurationServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigurationServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigurationServiceResultOutput) Properties() ConfigurationServicePropertiesResponseOutput {
	return o.ApplyT(func(v LookupConfigurationServiceResult) ConfigurationServicePropertiesResponse { return v.Properties }).(ConfigurationServicePropertiesResponseOutput)
}

func (o LookupConfigurationServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConfigurationServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationServiceResultOutput{})
}
