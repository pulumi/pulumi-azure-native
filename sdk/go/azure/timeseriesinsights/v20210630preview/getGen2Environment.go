


package v20210630preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGen2Environment(ctx *pulumi.Context, args *LookupGen2EnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupGen2EnvironmentResult, error) {
	var rv LookupGen2EnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20210630preview:getGen2Environment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGen2EnvironmentArgs struct {
	EnvironmentName   string  `pulumi:"environmentName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupGen2EnvironmentResult struct {
	CreationTime               string                                    `pulumi:"creationTime"`
	DataAccessFqdn             string                                    `pulumi:"dataAccessFqdn"`
	DataAccessId               string                                    `pulumi:"dataAccessId"`
	Id                         string                                    `pulumi:"id"`
	Kind                       string                                    `pulumi:"kind"`
	Location                   string                                    `pulumi:"location"`
	Name                       string                                    `pulumi:"name"`
	ProvisioningState          string                                    `pulumi:"provisioningState"`
	Sku                        SkuResponse                               `pulumi:"sku"`
	Status                     EnvironmentStatusResponse                 `pulumi:"status"`
	StorageConfiguration       Gen2StorageConfigurationOutputResponse    `pulumi:"storageConfiguration"`
	SupportsCustomerManagedKey bool                                      `pulumi:"supportsCustomerManagedKey"`
	Tags                       map[string]string                         `pulumi:"tags"`
	TimeSeriesIdProperties     []TimeSeriesIdPropertyResponse            `pulumi:"timeSeriesIdProperties"`
	Type                       string                                    `pulumi:"type"`
	WarmStoreConfiguration     *WarmStoreConfigurationPropertiesResponse `pulumi:"warmStoreConfiguration"`
}

func LookupGen2EnvironmentOutput(ctx *pulumi.Context, args LookupGen2EnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupGen2EnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGen2EnvironmentResult, error) {
			args := v.(LookupGen2EnvironmentArgs)
			r, err := LookupGen2Environment(ctx, &args, opts...)
			var s LookupGen2EnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGen2EnvironmentResultOutput)
}

type LookupGen2EnvironmentOutputArgs struct {
	EnvironmentName   pulumi.StringInput    `pulumi:"environmentName"`
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupGen2EnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGen2EnvironmentArgs)(nil)).Elem()
}


type LookupGen2EnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupGen2EnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGen2EnvironmentResult)(nil)).Elem()
}

func (o LookupGen2EnvironmentResultOutput) ToLookupGen2EnvironmentResultOutput() LookupGen2EnvironmentResultOutput {
	return o
}

func (o LookupGen2EnvironmentResultOutput) ToLookupGen2EnvironmentResultOutputWithContext(ctx context.Context) LookupGen2EnvironmentResultOutput {
	return o
}

func (o LookupGen2EnvironmentResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupGen2EnvironmentResultOutput) DataAccessFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.DataAccessFqdn }).(pulumi.StringOutput)
}

func (o LookupGen2EnvironmentResultOutput) DataAccessId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.DataAccessId }).(pulumi.StringOutput)
}

func (o LookupGen2EnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGen2EnvironmentResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupGen2EnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGen2EnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGen2EnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupGen2EnvironmentResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupGen2EnvironmentResultOutput) Status() EnvironmentStatusResponseOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) EnvironmentStatusResponse { return v.Status }).(EnvironmentStatusResponseOutput)
}

func (o LookupGen2EnvironmentResultOutput) StorageConfiguration() Gen2StorageConfigurationOutputResponseOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) Gen2StorageConfigurationOutputResponse {
		return v.StorageConfiguration
	}).(Gen2StorageConfigurationOutputResponseOutput)
}

func (o LookupGen2EnvironmentResultOutput) SupportsCustomerManagedKey() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) bool { return v.SupportsCustomerManagedKey }).(pulumi.BoolOutput)
}

func (o LookupGen2EnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGen2EnvironmentResultOutput) TimeSeriesIdProperties() TimeSeriesIdPropertyResponseArrayOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) []TimeSeriesIdPropertyResponse { return v.TimeSeriesIdProperties }).(TimeSeriesIdPropertyResponseArrayOutput)
}

func (o LookupGen2EnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupGen2EnvironmentResultOutput) WarmStoreConfiguration() WarmStoreConfigurationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupGen2EnvironmentResult) *WarmStoreConfigurationPropertiesResponse {
		return v.WarmStoreConfiguration
	}).(WarmStoreConfigurationPropertiesResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGen2EnvironmentResultOutput{})
}
