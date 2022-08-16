


package v20180815preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLongTermEnvironment(ctx *pulumi.Context, args *LookupLongTermEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupLongTermEnvironmentResult, error) {
	var rv LookupLongTermEnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20180815preview:getLongTermEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLongTermEnvironmentArgs struct {
	EnvironmentName   string  `pulumi:"environmentName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupLongTermEnvironmentResult struct {
	CreationTime           string                                     `pulumi:"creationTime"`
	DataAccessFqdn         string                                     `pulumi:"dataAccessFqdn"`
	DataAccessId           string                                     `pulumi:"dataAccessId"`
	Id                     string                                     `pulumi:"id"`
	Kind                   string                                     `pulumi:"kind"`
	Location               string                                     `pulumi:"location"`
	Name                   string                                     `pulumi:"name"`
	ProvisioningState      string                                     `pulumi:"provisioningState"`
	Sku                    SkuResponse                                `pulumi:"sku"`
	Status                 EnvironmentStatusResponse                  `pulumi:"status"`
	StorageConfiguration   LongTermStorageConfigurationOutputResponse `pulumi:"storageConfiguration"`
	Tags                   map[string]string                          `pulumi:"tags"`
	TimeSeriesIdProperties []TimeSeriesIdPropertyResponse             `pulumi:"timeSeriesIdProperties"`
	Type                   string                                     `pulumi:"type"`
	WarmStoreConfiguration *WarmStoreConfigurationPropertiesResponse  `pulumi:"warmStoreConfiguration"`
}

func LookupLongTermEnvironmentOutput(ctx *pulumi.Context, args LookupLongTermEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupLongTermEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLongTermEnvironmentResult, error) {
			args := v.(LookupLongTermEnvironmentArgs)
			r, err := LookupLongTermEnvironment(ctx, &args, opts...)
			var s LookupLongTermEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLongTermEnvironmentResultOutput)
}

type LookupLongTermEnvironmentOutputArgs struct {
	EnvironmentName   pulumi.StringInput    `pulumi:"environmentName"`
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupLongTermEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLongTermEnvironmentArgs)(nil)).Elem()
}


type LookupLongTermEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupLongTermEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLongTermEnvironmentResult)(nil)).Elem()
}

func (o LookupLongTermEnvironmentResultOutput) ToLookupLongTermEnvironmentResultOutput() LookupLongTermEnvironmentResultOutput {
	return o
}

func (o LookupLongTermEnvironmentResultOutput) ToLookupLongTermEnvironmentResultOutputWithContext(ctx context.Context) LookupLongTermEnvironmentResultOutput {
	return o
}

func (o LookupLongTermEnvironmentResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupLongTermEnvironmentResultOutput) DataAccessFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) string { return v.DataAccessFqdn }).(pulumi.StringOutput)
}

func (o LookupLongTermEnvironmentResultOutput) DataAccessId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) string { return v.DataAccessId }).(pulumi.StringOutput)
}

func (o LookupLongTermEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLongTermEnvironmentResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupLongTermEnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupLongTermEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLongTermEnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLongTermEnvironmentResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupLongTermEnvironmentResultOutput) Status() EnvironmentStatusResponseOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) EnvironmentStatusResponse { return v.Status }).(EnvironmentStatusResponseOutput)
}

func (o LookupLongTermEnvironmentResultOutput) StorageConfiguration() LongTermStorageConfigurationOutputResponseOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) LongTermStorageConfigurationOutputResponse {
		return v.StorageConfiguration
	}).(LongTermStorageConfigurationOutputResponseOutput)
}

func (o LookupLongTermEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLongTermEnvironmentResultOutput) TimeSeriesIdProperties() TimeSeriesIdPropertyResponseArrayOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) []TimeSeriesIdPropertyResponse {
		return v.TimeSeriesIdProperties
	}).(TimeSeriesIdPropertyResponseArrayOutput)
}

func (o LookupLongTermEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupLongTermEnvironmentResultOutput) WarmStoreConfiguration() WarmStoreConfigurationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupLongTermEnvironmentResult) *WarmStoreConfigurationPropertiesResponse {
		return v.WarmStoreConfiguration
	}).(WarmStoreConfigurationPropertiesResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLongTermEnvironmentResultOutput{})
}
