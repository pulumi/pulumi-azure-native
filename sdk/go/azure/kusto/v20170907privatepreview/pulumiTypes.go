


package v20170907privatepreview

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
	AppId *string `pulumi:"appId"`
	Email *string `pulumi:"email"`
	Fqn   *string `pulumi:"fqn"`
	Name  string  `pulumi:"name"`
	Role  string  `pulumi:"role"`
	Type  string  `pulumi:"type"`
}

type DatabasePrincipalResponseOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePrincipalResponse)(nil)).Elem()
}

func (o DatabasePrincipalResponseOutput) ToDatabasePrincipalResponseOutput() DatabasePrincipalResponseOutput {
	return o
}

func (o DatabasePrincipalResponseOutput) ToDatabasePrincipalResponseOutputWithContext(ctx context.Context) DatabasePrincipalResponseOutput {
	return o
}

func (o DatabasePrincipalResponseOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o DatabasePrincipalResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o DatabasePrincipalResponseOutput) Fqn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) *string { return v.Fqn }).(pulumi.StringPtrOutput)
}

func (o DatabasePrincipalResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DatabasePrincipalResponseOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) string { return v.Role }).(pulumi.StringOutput)
}

func (o DatabasePrincipalResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) string { return v.Type }).(pulumi.StringOutput)
}

type DatabasePrincipalResponseArrayOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabasePrincipalResponse)(nil)).Elem()
}

func (o DatabasePrincipalResponseArrayOutput) ToDatabasePrincipalResponseArrayOutput() DatabasePrincipalResponseArrayOutput {
	return o
}

func (o DatabasePrincipalResponseArrayOutput) ToDatabasePrincipalResponseArrayOutputWithContext(ctx context.Context) DatabasePrincipalResponseArrayOutput {
	return o
}

func (o DatabasePrincipalResponseArrayOutput) Index(i pulumi.IntInput) DatabasePrincipalResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabasePrincipalResponse {
		return vs[0].([]DatabasePrincipalResponse)[vs[1].(int)]
	}).(DatabasePrincipalResponseOutput)
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

func init() {
	pulumi.RegisterOutputType(AzureSkuOutput{})
	pulumi.RegisterOutputType(AzureSkuResponseOutput{})
	pulumi.RegisterOutputType(DatabasePrincipalResponseOutput{})
	pulumi.RegisterOutputType(DatabasePrincipalResponseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseStatisticsResponseOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantArrayOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantResponseOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantResponseArrayOutput{})
}
