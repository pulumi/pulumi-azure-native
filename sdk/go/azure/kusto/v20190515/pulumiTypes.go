


package v20190515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureSku struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
	Tier     string `pulumi:"tier"`
}





type AzureSkuInput interface {
	pulumi.Input

	ToAzureSkuOutput() AzureSkuOutput
	ToAzureSkuOutputWithContext(context.Context) AzureSkuOutput
}

type AzureSkuArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	Name     pulumi.StringInput `pulumi:"name"`
	Tier     pulumi.StringInput `pulumi:"tier"`
}

func (AzureSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSku)(nil)).Elem()
}

func (i AzureSkuArgs) ToAzureSkuOutput() AzureSkuOutput {
	return i.ToAzureSkuOutputWithContext(context.Background())
}

func (i AzureSkuArgs) ToAzureSkuOutputWithContext(ctx context.Context) AzureSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuOutput)
}

type AzureSkuOutput struct{ *pulumi.OutputState }

func (AzureSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSku)(nil)).Elem()
}

func (o AzureSkuOutput) ToAzureSkuOutput() AzureSkuOutput {
	return o
}

func (o AzureSkuOutput) ToAzureSkuOutputWithContext(ctx context.Context) AzureSkuOutput {
	return o
}

func (o AzureSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o AzureSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o AzureSkuOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSku) string { return v.Tier }).(pulumi.StringOutput)
}

type AzureSkuResponse struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
	Tier     string `pulumi:"tier"`
}

type AzureSkuResponseOutput struct{ *pulumi.OutputState }

func (AzureSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuResponse)(nil)).Elem()
}

func (o AzureSkuResponseOutput) ToAzureSkuResponseOutput() AzureSkuResponseOutput {
	return o
}

func (o AzureSkuResponseOutput) ToAzureSkuResponseOutputWithContext(ctx context.Context) AzureSkuResponseOutput {
	return o
}

func (o AzureSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o AzureSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AzureSkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type DatabasePrincipalResponse struct {
	AppId      *string `pulumi:"appId"`
	Email      *string `pulumi:"email"`
	Fqn        *string `pulumi:"fqn"`
	Name       string  `pulumi:"name"`
	Role       string  `pulumi:"role"`
	TenantName string  `pulumi:"tenantName"`
	Type       string  `pulumi:"type"`
}

type DatabaseStatisticsResponse struct {
	Size *float64 `pulumi:"size"`
}

type DatabaseStatisticsResponseOutput struct{ *pulumi.OutputState }

func (DatabaseStatisticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseStatisticsResponse)(nil)).Elem()
}

func (o DatabaseStatisticsResponseOutput) ToDatabaseStatisticsResponseOutput() DatabaseStatisticsResponseOutput {
	return o
}

func (o DatabaseStatisticsResponseOutput) ToDatabaseStatisticsResponseOutputWithContext(ctx context.Context) DatabaseStatisticsResponseOutput {
	return o
}

func (o DatabaseStatisticsResponseOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DatabaseStatisticsResponse) *float64 { return v.Size }).(pulumi.Float64PtrOutput)
}

type OptimizedAutoscale struct {
	IsEnabled bool `pulumi:"isEnabled"`
	Maximum   int  `pulumi:"maximum"`
	Minimum   int  `pulumi:"minimum"`
	Version   int  `pulumi:"version"`
}





type OptimizedAutoscaleInput interface {
	pulumi.Input

	ToOptimizedAutoscaleOutput() OptimizedAutoscaleOutput
	ToOptimizedAutoscaleOutputWithContext(context.Context) OptimizedAutoscaleOutput
}

type OptimizedAutoscaleArgs struct {
	IsEnabled pulumi.BoolInput `pulumi:"isEnabled"`
	Maximum   pulumi.IntInput  `pulumi:"maximum"`
	Minimum   pulumi.IntInput  `pulumi:"minimum"`
	Version   pulumi.IntInput  `pulumi:"version"`
}

func (OptimizedAutoscaleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OptimizedAutoscale)(nil)).Elem()
}

func (i OptimizedAutoscaleArgs) ToOptimizedAutoscaleOutput() OptimizedAutoscaleOutput {
	return i.ToOptimizedAutoscaleOutputWithContext(context.Background())
}

func (i OptimizedAutoscaleArgs) ToOptimizedAutoscaleOutputWithContext(ctx context.Context) OptimizedAutoscaleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptimizedAutoscaleOutput)
}

func (i OptimizedAutoscaleArgs) ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput {
	return i.ToOptimizedAutoscalePtrOutputWithContext(context.Background())
}

func (i OptimizedAutoscaleArgs) ToOptimizedAutoscalePtrOutputWithContext(ctx context.Context) OptimizedAutoscalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptimizedAutoscaleOutput).ToOptimizedAutoscalePtrOutputWithContext(ctx)
}









type OptimizedAutoscalePtrInput interface {
	pulumi.Input

	ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput
	ToOptimizedAutoscalePtrOutputWithContext(context.Context) OptimizedAutoscalePtrOutput
}

type optimizedAutoscalePtrType OptimizedAutoscaleArgs

func OptimizedAutoscalePtr(v *OptimizedAutoscaleArgs) OptimizedAutoscalePtrInput {
	return (*optimizedAutoscalePtrType)(v)
}

func (*optimizedAutoscalePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OptimizedAutoscale)(nil)).Elem()
}

func (i *optimizedAutoscalePtrType) ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput {
	return i.ToOptimizedAutoscalePtrOutputWithContext(context.Background())
}

func (i *optimizedAutoscalePtrType) ToOptimizedAutoscalePtrOutputWithContext(ctx context.Context) OptimizedAutoscalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptimizedAutoscalePtrOutput)
}

type OptimizedAutoscaleOutput struct{ *pulumi.OutputState }

func (OptimizedAutoscaleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OptimizedAutoscale)(nil)).Elem()
}

func (o OptimizedAutoscaleOutput) ToOptimizedAutoscaleOutput() OptimizedAutoscaleOutput {
	return o
}

func (o OptimizedAutoscaleOutput) ToOptimizedAutoscaleOutputWithContext(ctx context.Context) OptimizedAutoscaleOutput {
	return o
}

func (o OptimizedAutoscaleOutput) ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput {
	return o.ToOptimizedAutoscalePtrOutputWithContext(context.Background())
}

func (o OptimizedAutoscaleOutput) ToOptimizedAutoscalePtrOutputWithContext(ctx context.Context) OptimizedAutoscalePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OptimizedAutoscale) *OptimizedAutoscale {
		return &v
	}).(OptimizedAutoscalePtrOutput)
}

func (o OptimizedAutoscaleOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v OptimizedAutoscale) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o OptimizedAutoscaleOutput) Maximum() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscale) int { return v.Maximum }).(pulumi.IntOutput)
}

func (o OptimizedAutoscaleOutput) Minimum() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscale) int { return v.Minimum }).(pulumi.IntOutput)
}

func (o OptimizedAutoscaleOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscale) int { return v.Version }).(pulumi.IntOutput)
}

type OptimizedAutoscalePtrOutput struct{ *pulumi.OutputState }

func (OptimizedAutoscalePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OptimizedAutoscale)(nil)).Elem()
}

func (o OptimizedAutoscalePtrOutput) ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput {
	return o
}

func (o OptimizedAutoscalePtrOutput) ToOptimizedAutoscalePtrOutputWithContext(ctx context.Context) OptimizedAutoscalePtrOutput {
	return o
}

func (o OptimizedAutoscalePtrOutput) Elem() OptimizedAutoscaleOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) OptimizedAutoscale {
		if v != nil {
			return *v
		}
		var ret OptimizedAutoscale
		return ret
	}).(OptimizedAutoscaleOutput)
}

func (o OptimizedAutoscalePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o OptimizedAutoscalePtrOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) *int {
		if v == nil {
			return nil
		}
		return &v.Maximum
	}).(pulumi.IntPtrOutput)
}

func (o OptimizedAutoscalePtrOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) *int {
		if v == nil {
			return nil
		}
		return &v.Minimum
	}).(pulumi.IntPtrOutput)
}

func (o OptimizedAutoscalePtrOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) *int {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.IntPtrOutput)
}

type OptimizedAutoscaleResponse struct {
	IsEnabled bool `pulumi:"isEnabled"`
	Maximum   int  `pulumi:"maximum"`
	Minimum   int  `pulumi:"minimum"`
	Version   int  `pulumi:"version"`
}

type OptimizedAutoscaleResponseOutput struct{ *pulumi.OutputState }

func (OptimizedAutoscaleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OptimizedAutoscaleResponse)(nil)).Elem()
}

func (o OptimizedAutoscaleResponseOutput) ToOptimizedAutoscaleResponseOutput() OptimizedAutoscaleResponseOutput {
	return o
}

func (o OptimizedAutoscaleResponseOutput) ToOptimizedAutoscaleResponseOutputWithContext(ctx context.Context) OptimizedAutoscaleResponseOutput {
	return o
}

func (o OptimizedAutoscaleResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v OptimizedAutoscaleResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o OptimizedAutoscaleResponseOutput) Maximum() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscaleResponse) int { return v.Maximum }).(pulumi.IntOutput)
}

func (o OptimizedAutoscaleResponseOutput) Minimum() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscaleResponse) int { return v.Minimum }).(pulumi.IntOutput)
}

func (o OptimizedAutoscaleResponseOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscaleResponse) int { return v.Version }).(pulumi.IntOutput)
}

type OptimizedAutoscaleResponsePtrOutput struct{ *pulumi.OutputState }

func (OptimizedAutoscaleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OptimizedAutoscaleResponse)(nil)).Elem()
}

func (o OptimizedAutoscaleResponsePtrOutput) ToOptimizedAutoscaleResponsePtrOutput() OptimizedAutoscaleResponsePtrOutput {
	return o
}

func (o OptimizedAutoscaleResponsePtrOutput) ToOptimizedAutoscaleResponsePtrOutputWithContext(ctx context.Context) OptimizedAutoscaleResponsePtrOutput {
	return o
}

func (o OptimizedAutoscaleResponsePtrOutput) Elem() OptimizedAutoscaleResponseOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) OptimizedAutoscaleResponse {
		if v != nil {
			return *v
		}
		var ret OptimizedAutoscaleResponse
		return ret
	}).(OptimizedAutoscaleResponseOutput)
}

func (o OptimizedAutoscaleResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o OptimizedAutoscaleResponsePtrOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Maximum
	}).(pulumi.IntPtrOutput)
}

func (o OptimizedAutoscaleResponsePtrOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Minimum
	}).(pulumi.IntPtrOutput)
}

func (o OptimizedAutoscaleResponsePtrOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.IntPtrOutput)
}

type TrustedExternalTenant struct {
	Value *string `pulumi:"value"`
}





type TrustedExternalTenantInput interface {
	pulumi.Input

	ToTrustedExternalTenantOutput() TrustedExternalTenantOutput
	ToTrustedExternalTenantOutputWithContext(context.Context) TrustedExternalTenantOutput
}

type TrustedExternalTenantArgs struct {
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (TrustedExternalTenantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedExternalTenant)(nil)).Elem()
}

func (i TrustedExternalTenantArgs) ToTrustedExternalTenantOutput() TrustedExternalTenantOutput {
	return i.ToTrustedExternalTenantOutputWithContext(context.Background())
}

func (i TrustedExternalTenantArgs) ToTrustedExternalTenantOutputWithContext(ctx context.Context) TrustedExternalTenantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedExternalTenantOutput)
}





type TrustedExternalTenantArrayInput interface {
	pulumi.Input

	ToTrustedExternalTenantArrayOutput() TrustedExternalTenantArrayOutput
	ToTrustedExternalTenantArrayOutputWithContext(context.Context) TrustedExternalTenantArrayOutput
}

type TrustedExternalTenantArray []TrustedExternalTenantInput

func (TrustedExternalTenantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustedExternalTenant)(nil)).Elem()
}

func (i TrustedExternalTenantArray) ToTrustedExternalTenantArrayOutput() TrustedExternalTenantArrayOutput {
	return i.ToTrustedExternalTenantArrayOutputWithContext(context.Background())
}

func (i TrustedExternalTenantArray) ToTrustedExternalTenantArrayOutputWithContext(ctx context.Context) TrustedExternalTenantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedExternalTenantArrayOutput)
}

type TrustedExternalTenantOutput struct{ *pulumi.OutputState }

func (TrustedExternalTenantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedExternalTenant)(nil)).Elem()
}

func (o TrustedExternalTenantOutput) ToTrustedExternalTenantOutput() TrustedExternalTenantOutput {
	return o
}

func (o TrustedExternalTenantOutput) ToTrustedExternalTenantOutputWithContext(ctx context.Context) TrustedExternalTenantOutput {
	return o
}

func (o TrustedExternalTenantOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustedExternalTenant) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type TrustedExternalTenantArrayOutput struct{ *pulumi.OutputState }

func (TrustedExternalTenantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustedExternalTenant)(nil)).Elem()
}

func (o TrustedExternalTenantArrayOutput) ToTrustedExternalTenantArrayOutput() TrustedExternalTenantArrayOutput {
	return o
}

func (o TrustedExternalTenantArrayOutput) ToTrustedExternalTenantArrayOutputWithContext(ctx context.Context) TrustedExternalTenantArrayOutput {
	return o
}

func (o TrustedExternalTenantArrayOutput) Index(i pulumi.IntInput) TrustedExternalTenantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrustedExternalTenant {
		return vs[0].([]TrustedExternalTenant)[vs[1].(int)]
	}).(TrustedExternalTenantOutput)
}

type TrustedExternalTenantResponse struct {
	Value *string `pulumi:"value"`
}

type TrustedExternalTenantResponseOutput struct{ *pulumi.OutputState }

func (TrustedExternalTenantResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedExternalTenantResponse)(nil)).Elem()
}

func (o TrustedExternalTenantResponseOutput) ToTrustedExternalTenantResponseOutput() TrustedExternalTenantResponseOutput {
	return o
}

func (o TrustedExternalTenantResponseOutput) ToTrustedExternalTenantResponseOutputWithContext(ctx context.Context) TrustedExternalTenantResponseOutput {
	return o
}

func (o TrustedExternalTenantResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustedExternalTenantResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type TrustedExternalTenantResponseArrayOutput struct{ *pulumi.OutputState }

func (TrustedExternalTenantResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustedExternalTenantResponse)(nil)).Elem()
}

func (o TrustedExternalTenantResponseArrayOutput) ToTrustedExternalTenantResponseArrayOutput() TrustedExternalTenantResponseArrayOutput {
	return o
}

func (o TrustedExternalTenantResponseArrayOutput) ToTrustedExternalTenantResponseArrayOutputWithContext(ctx context.Context) TrustedExternalTenantResponseArrayOutput {
	return o
}

func (o TrustedExternalTenantResponseArrayOutput) Index(i pulumi.IntInput) TrustedExternalTenantResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrustedExternalTenantResponse {
		return vs[0].([]TrustedExternalTenantResponse)[vs[1].(int)]
	}).(TrustedExternalTenantResponseOutput)
}

type VirtualNetworkConfiguration struct {
	DataManagementPublicIpId string `pulumi:"dataManagementPublicIpId"`
	EnginePublicIpId         string `pulumi:"enginePublicIpId"`
	SubnetId                 string `pulumi:"subnetId"`
}





type VirtualNetworkConfigurationInput interface {
	pulumi.Input

	ToVirtualNetworkConfigurationOutput() VirtualNetworkConfigurationOutput
	ToVirtualNetworkConfigurationOutputWithContext(context.Context) VirtualNetworkConfigurationOutput
}

type VirtualNetworkConfigurationArgs struct {
	DataManagementPublicIpId pulumi.StringInput `pulumi:"dataManagementPublicIpId"`
	EnginePublicIpId         pulumi.StringInput `pulumi:"enginePublicIpId"`
	SubnetId                 pulumi.StringInput `pulumi:"subnetId"`
}

func (VirtualNetworkConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfiguration)(nil)).Elem()
}

func (i VirtualNetworkConfigurationArgs) ToVirtualNetworkConfigurationOutput() VirtualNetworkConfigurationOutput {
	return i.ToVirtualNetworkConfigurationOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigurationArgs) ToVirtualNetworkConfigurationOutputWithContext(ctx context.Context) VirtualNetworkConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigurationOutput)
}

func (i VirtualNetworkConfigurationArgs) ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput {
	return i.ToVirtualNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigurationArgs) ToVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigurationOutput).ToVirtualNetworkConfigurationPtrOutputWithContext(ctx)
}









type VirtualNetworkConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput
	ToVirtualNetworkConfigurationPtrOutputWithContext(context.Context) VirtualNetworkConfigurationPtrOutput
}

type virtualNetworkConfigurationPtrType VirtualNetworkConfigurationArgs

func VirtualNetworkConfigurationPtr(v *VirtualNetworkConfigurationArgs) VirtualNetworkConfigurationPtrInput {
	return (*virtualNetworkConfigurationPtrType)(v)
}

func (*virtualNetworkConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfiguration)(nil)).Elem()
}

func (i *virtualNetworkConfigurationPtrType) ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput {
	return i.ToVirtualNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkConfigurationPtrType) ToVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigurationPtrOutput)
}

type VirtualNetworkConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfiguration)(nil)).Elem()
}

func (o VirtualNetworkConfigurationOutput) ToVirtualNetworkConfigurationOutput() VirtualNetworkConfigurationOutput {
	return o
}

func (o VirtualNetworkConfigurationOutput) ToVirtualNetworkConfigurationOutputWithContext(ctx context.Context) VirtualNetworkConfigurationOutput {
	return o
}

func (o VirtualNetworkConfigurationOutput) ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput {
	return o.ToVirtualNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkConfigurationOutput) ToVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkConfiguration) *VirtualNetworkConfiguration {
		return &v
	}).(VirtualNetworkConfigurationPtrOutput)
}

func (o VirtualNetworkConfigurationOutput) DataManagementPublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfiguration) string { return v.DataManagementPublicIpId }).(pulumi.StringOutput)
}

func (o VirtualNetworkConfigurationOutput) EnginePublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfiguration) string { return v.EnginePublicIpId }).(pulumi.StringOutput)
}

func (o VirtualNetworkConfigurationOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfiguration) string { return v.SubnetId }).(pulumi.StringOutput)
}

type VirtualNetworkConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfiguration)(nil)).Elem()
}

func (o VirtualNetworkConfigurationPtrOutput) ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput {
	return o
}

func (o VirtualNetworkConfigurationPtrOutput) ToVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationPtrOutput {
	return o
}

func (o VirtualNetworkConfigurationPtrOutput) Elem() VirtualNetworkConfigurationOutput {
	return o.ApplyT(func(v *VirtualNetworkConfiguration) VirtualNetworkConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkConfiguration
		return ret
	}).(VirtualNetworkConfigurationOutput)
}

func (o VirtualNetworkConfigurationPtrOutput) DataManagementPublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.DataManagementPublicIpId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationPtrOutput) EnginePublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.EnginePublicIpId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationPtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkConfigurationResponse struct {
	DataManagementPublicIpId string `pulumi:"dataManagementPublicIpId"`
	EnginePublicIpId         string `pulumi:"enginePublicIpId"`
	SubnetId                 string `pulumi:"subnetId"`
}

type VirtualNetworkConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfigurationResponse)(nil)).Elem()
}

func (o VirtualNetworkConfigurationResponseOutput) ToVirtualNetworkConfigurationResponseOutput() VirtualNetworkConfigurationResponseOutput {
	return o
}

func (o VirtualNetworkConfigurationResponseOutput) ToVirtualNetworkConfigurationResponseOutputWithContext(ctx context.Context) VirtualNetworkConfigurationResponseOutput {
	return o
}

func (o VirtualNetworkConfigurationResponseOutput) DataManagementPublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfigurationResponse) string { return v.DataManagementPublicIpId }).(pulumi.StringOutput)
}

func (o VirtualNetworkConfigurationResponseOutput) EnginePublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfigurationResponse) string { return v.EnginePublicIpId }).(pulumi.StringOutput)
}

func (o VirtualNetworkConfigurationResponseOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfigurationResponse) string { return v.SubnetId }).(pulumi.StringOutput)
}

type VirtualNetworkConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfigurationResponse)(nil)).Elem()
}

func (o VirtualNetworkConfigurationResponsePtrOutput) ToVirtualNetworkConfigurationResponsePtrOutput() VirtualNetworkConfigurationResponsePtrOutput {
	return o
}

func (o VirtualNetworkConfigurationResponsePtrOutput) ToVirtualNetworkConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationResponsePtrOutput {
	return o
}

func (o VirtualNetworkConfigurationResponsePtrOutput) Elem() VirtualNetworkConfigurationResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigurationResponse) VirtualNetworkConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkConfigurationResponse
		return ret
	}).(VirtualNetworkConfigurationResponseOutput)
}

func (o VirtualNetworkConfigurationResponsePtrOutput) DataManagementPublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DataManagementPublicIpId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationResponsePtrOutput) EnginePublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EnginePublicIpId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationResponsePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SubnetId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureSkuOutput{})
	pulumi.RegisterOutputType(AzureSkuResponseOutput{})
	pulumi.RegisterOutputType(DatabaseStatisticsResponseOutput{})
	pulumi.RegisterOutputType(OptimizedAutoscaleOutput{})
	pulumi.RegisterOutputType(OptimizedAutoscalePtrOutput{})
	pulumi.RegisterOutputType(OptimizedAutoscaleResponseOutput{})
	pulumi.RegisterOutputType(OptimizedAutoscaleResponsePtrOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantArrayOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantResponseOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationResponsePtrOutput{})
}
