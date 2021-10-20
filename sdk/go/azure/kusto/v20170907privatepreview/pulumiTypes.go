


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

func (i AzureSkuArgs) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return i.ToAzureSkuPtrOutputWithContext(context.Background())
}

func (i AzureSkuArgs) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuOutput).ToAzureSkuPtrOutputWithContext(ctx)
}









type AzureSkuPtrInput interface {
	pulumi.Input

	ToAzureSkuPtrOutput() AzureSkuPtrOutput
	ToAzureSkuPtrOutputWithContext(context.Context) AzureSkuPtrOutput
}

type azureSkuPtrType AzureSkuArgs

func AzureSkuPtr(v *AzureSkuArgs) AzureSkuPtrInput {
	return (*azureSkuPtrType)(v)
}

func (*azureSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSku)(nil)).Elem()
}

func (i *azureSkuPtrType) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return i.ToAzureSkuPtrOutputWithContext(context.Background())
}

func (i *azureSkuPtrType) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuPtrOutput)
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

func (o AzureSkuOutput) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return o.ToAzureSkuPtrOutputWithContext(context.Background())
}

func (o AzureSkuOutput) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSku) *AzureSku {
		return &v
	}).(AzureSkuPtrOutput)
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

type AzureSkuPtrOutput struct{ *pulumi.OutputState }

func (AzureSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSku)(nil)).Elem()
}

func (o AzureSkuPtrOutput) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return o
}

func (o AzureSkuPtrOutput) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return o
}

func (o AzureSkuPtrOutput) Elem() AzureSkuOutput {
	return o.ApplyT(func(v *AzureSku) AzureSku {
		if v != nil {
			return *v
		}
		var ret AzureSku
		return ret
	}).(AzureSkuOutput)
}

func (o AzureSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o AzureSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o AzureSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSku) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type AzureSkuResponse struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
	Tier     string `pulumi:"tier"`
}





type AzureSkuResponseInput interface {
	pulumi.Input

	ToAzureSkuResponseOutput() AzureSkuResponseOutput
	ToAzureSkuResponseOutputWithContext(context.Context) AzureSkuResponseOutput
}

type AzureSkuResponseArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	Name     pulumi.StringInput `pulumi:"name"`
	Tier     pulumi.StringInput `pulumi:"tier"`
}

func (AzureSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuResponse)(nil)).Elem()
}

func (i AzureSkuResponseArgs) ToAzureSkuResponseOutput() AzureSkuResponseOutput {
	return i.ToAzureSkuResponseOutputWithContext(context.Background())
}

func (i AzureSkuResponseArgs) ToAzureSkuResponseOutputWithContext(ctx context.Context) AzureSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuResponseOutput)
}

func (i AzureSkuResponseArgs) ToAzureSkuResponsePtrOutput() AzureSkuResponsePtrOutput {
	return i.ToAzureSkuResponsePtrOutputWithContext(context.Background())
}

func (i AzureSkuResponseArgs) ToAzureSkuResponsePtrOutputWithContext(ctx context.Context) AzureSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuResponseOutput).ToAzureSkuResponsePtrOutputWithContext(ctx)
}









type AzureSkuResponsePtrInput interface {
	pulumi.Input

	ToAzureSkuResponsePtrOutput() AzureSkuResponsePtrOutput
	ToAzureSkuResponsePtrOutputWithContext(context.Context) AzureSkuResponsePtrOutput
}

type azureSkuResponsePtrType AzureSkuResponseArgs

func AzureSkuResponsePtr(v *AzureSkuResponseArgs) AzureSkuResponsePtrInput {
	return (*azureSkuResponsePtrType)(v)
}

func (*azureSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSkuResponse)(nil)).Elem()
}

func (i *azureSkuResponsePtrType) ToAzureSkuResponsePtrOutput() AzureSkuResponsePtrOutput {
	return i.ToAzureSkuResponsePtrOutputWithContext(context.Background())
}

func (i *azureSkuResponsePtrType) ToAzureSkuResponsePtrOutputWithContext(ctx context.Context) AzureSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuResponsePtrOutput)
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

func (o AzureSkuResponseOutput) ToAzureSkuResponsePtrOutput() AzureSkuResponsePtrOutput {
	return o.ToAzureSkuResponsePtrOutputWithContext(context.Background())
}

func (o AzureSkuResponseOutput) ToAzureSkuResponsePtrOutputWithContext(ctx context.Context) AzureSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSkuResponse) *AzureSkuResponse {
		return &v
	}).(AzureSkuResponsePtrOutput)
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

type AzureSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSkuResponse)(nil)).Elem()
}

func (o AzureSkuResponsePtrOutput) ToAzureSkuResponsePtrOutput() AzureSkuResponsePtrOutput {
	return o
}

func (o AzureSkuResponsePtrOutput) ToAzureSkuResponsePtrOutputWithContext(ctx context.Context) AzureSkuResponsePtrOutput {
	return o
}

func (o AzureSkuResponsePtrOutput) Elem() AzureSkuResponseOutput {
	return o.ApplyT(func(v *AzureSkuResponse) AzureSkuResponse {
		if v != nil {
			return *v
		}
		var ret AzureSkuResponse
		return ret
	}).(AzureSkuResponseOutput)
}

func (o AzureSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o AzureSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o AzureSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type DatabasePrincipalResponse struct {
	AppId *string `pulumi:"appId"`
	Email *string `pulumi:"email"`
	Fqn   *string `pulumi:"fqn"`
	Name  string  `pulumi:"name"`
	Role  string  `pulumi:"role"`
	Type  string  `pulumi:"type"`
}





type DatabasePrincipalResponseInput interface {
	pulumi.Input

	ToDatabasePrincipalResponseOutput() DatabasePrincipalResponseOutput
	ToDatabasePrincipalResponseOutputWithContext(context.Context) DatabasePrincipalResponseOutput
}

type DatabasePrincipalResponseArgs struct {
	AppId pulumi.StringPtrInput `pulumi:"appId"`
	Email pulumi.StringPtrInput `pulumi:"email"`
	Fqn   pulumi.StringPtrInput `pulumi:"fqn"`
	Name  pulumi.StringInput    `pulumi:"name"`
	Role  pulumi.StringInput    `pulumi:"role"`
	Type  pulumi.StringInput    `pulumi:"type"`
}

func (DatabasePrincipalResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePrincipalResponse)(nil)).Elem()
}

func (i DatabasePrincipalResponseArgs) ToDatabasePrincipalResponseOutput() DatabasePrincipalResponseOutput {
	return i.ToDatabasePrincipalResponseOutputWithContext(context.Background())
}

func (i DatabasePrincipalResponseArgs) ToDatabasePrincipalResponseOutputWithContext(ctx context.Context) DatabasePrincipalResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrincipalResponseOutput)
}





type DatabasePrincipalResponseArrayInput interface {
	pulumi.Input

	ToDatabasePrincipalResponseArrayOutput() DatabasePrincipalResponseArrayOutput
	ToDatabasePrincipalResponseArrayOutputWithContext(context.Context) DatabasePrincipalResponseArrayOutput
}

type DatabasePrincipalResponseArray []DatabasePrincipalResponseInput

func (DatabasePrincipalResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabasePrincipalResponse)(nil)).Elem()
}

func (i DatabasePrincipalResponseArray) ToDatabasePrincipalResponseArrayOutput() DatabasePrincipalResponseArrayOutput {
	return i.ToDatabasePrincipalResponseArrayOutputWithContext(context.Background())
}

func (i DatabasePrincipalResponseArray) ToDatabasePrincipalResponseArrayOutputWithContext(ctx context.Context) DatabasePrincipalResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrincipalResponseArrayOutput)
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





type DatabaseStatisticsResponseInput interface {
	pulumi.Input

	ToDatabaseStatisticsResponseOutput() DatabaseStatisticsResponseOutput
	ToDatabaseStatisticsResponseOutputWithContext(context.Context) DatabaseStatisticsResponseOutput
}

type DatabaseStatisticsResponseArgs struct {
	Size pulumi.Float64PtrInput `pulumi:"size"`
}

func (DatabaseStatisticsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseStatisticsResponse)(nil)).Elem()
}

func (i DatabaseStatisticsResponseArgs) ToDatabaseStatisticsResponseOutput() DatabaseStatisticsResponseOutput {
	return i.ToDatabaseStatisticsResponseOutputWithContext(context.Background())
}

func (i DatabaseStatisticsResponseArgs) ToDatabaseStatisticsResponseOutputWithContext(ctx context.Context) DatabaseStatisticsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseStatisticsResponseOutput)
}

func (i DatabaseStatisticsResponseArgs) ToDatabaseStatisticsResponsePtrOutput() DatabaseStatisticsResponsePtrOutput {
	return i.ToDatabaseStatisticsResponsePtrOutputWithContext(context.Background())
}

func (i DatabaseStatisticsResponseArgs) ToDatabaseStatisticsResponsePtrOutputWithContext(ctx context.Context) DatabaseStatisticsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseStatisticsResponseOutput).ToDatabaseStatisticsResponsePtrOutputWithContext(ctx)
}









type DatabaseStatisticsResponsePtrInput interface {
	pulumi.Input

	ToDatabaseStatisticsResponsePtrOutput() DatabaseStatisticsResponsePtrOutput
	ToDatabaseStatisticsResponsePtrOutputWithContext(context.Context) DatabaseStatisticsResponsePtrOutput
}

type databaseStatisticsResponsePtrType DatabaseStatisticsResponseArgs

func DatabaseStatisticsResponsePtr(v *DatabaseStatisticsResponseArgs) DatabaseStatisticsResponsePtrInput {
	return (*databaseStatisticsResponsePtrType)(v)
}

func (*databaseStatisticsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseStatisticsResponse)(nil)).Elem()
}

func (i *databaseStatisticsResponsePtrType) ToDatabaseStatisticsResponsePtrOutput() DatabaseStatisticsResponsePtrOutput {
	return i.ToDatabaseStatisticsResponsePtrOutputWithContext(context.Background())
}

func (i *databaseStatisticsResponsePtrType) ToDatabaseStatisticsResponsePtrOutputWithContext(ctx context.Context) DatabaseStatisticsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseStatisticsResponsePtrOutput)
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

func (o DatabaseStatisticsResponseOutput) ToDatabaseStatisticsResponsePtrOutput() DatabaseStatisticsResponsePtrOutput {
	return o.ToDatabaseStatisticsResponsePtrOutputWithContext(context.Background())
}

func (o DatabaseStatisticsResponseOutput) ToDatabaseStatisticsResponsePtrOutputWithContext(ctx context.Context) DatabaseStatisticsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseStatisticsResponse) *DatabaseStatisticsResponse {
		return &v
	}).(DatabaseStatisticsResponsePtrOutput)
}

func (o DatabaseStatisticsResponseOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DatabaseStatisticsResponse) *float64 { return v.Size }).(pulumi.Float64PtrOutput)
}

type DatabaseStatisticsResponsePtrOutput struct{ *pulumi.OutputState }

func (DatabaseStatisticsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseStatisticsResponse)(nil)).Elem()
}

func (o DatabaseStatisticsResponsePtrOutput) ToDatabaseStatisticsResponsePtrOutput() DatabaseStatisticsResponsePtrOutput {
	return o
}

func (o DatabaseStatisticsResponsePtrOutput) ToDatabaseStatisticsResponsePtrOutputWithContext(ctx context.Context) DatabaseStatisticsResponsePtrOutput {
	return o
}

func (o DatabaseStatisticsResponsePtrOutput) Elem() DatabaseStatisticsResponseOutput {
	return o.ApplyT(func(v *DatabaseStatisticsResponse) DatabaseStatisticsResponse {
		if v != nil {
			return *v
		}
		var ret DatabaseStatisticsResponse
		return ret
	}).(DatabaseStatisticsResponseOutput)
}

func (o DatabaseStatisticsResponsePtrOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DatabaseStatisticsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.Float64PtrOutput)
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





type TrustedExternalTenantResponseInput interface {
	pulumi.Input

	ToTrustedExternalTenantResponseOutput() TrustedExternalTenantResponseOutput
	ToTrustedExternalTenantResponseOutputWithContext(context.Context) TrustedExternalTenantResponseOutput
}

type TrustedExternalTenantResponseArgs struct {
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (TrustedExternalTenantResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedExternalTenantResponse)(nil)).Elem()
}

func (i TrustedExternalTenantResponseArgs) ToTrustedExternalTenantResponseOutput() TrustedExternalTenantResponseOutput {
	return i.ToTrustedExternalTenantResponseOutputWithContext(context.Background())
}

func (i TrustedExternalTenantResponseArgs) ToTrustedExternalTenantResponseOutputWithContext(ctx context.Context) TrustedExternalTenantResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedExternalTenantResponseOutput)
}





type TrustedExternalTenantResponseArrayInput interface {
	pulumi.Input

	ToTrustedExternalTenantResponseArrayOutput() TrustedExternalTenantResponseArrayOutput
	ToTrustedExternalTenantResponseArrayOutputWithContext(context.Context) TrustedExternalTenantResponseArrayOutput
}

type TrustedExternalTenantResponseArray []TrustedExternalTenantResponseInput

func (TrustedExternalTenantResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustedExternalTenantResponse)(nil)).Elem()
}

func (i TrustedExternalTenantResponseArray) ToTrustedExternalTenantResponseArrayOutput() TrustedExternalTenantResponseArrayOutput {
	return i.ToTrustedExternalTenantResponseArrayOutputWithContext(context.Background())
}

func (i TrustedExternalTenantResponseArray) ToTrustedExternalTenantResponseArrayOutputWithContext(ctx context.Context) TrustedExternalTenantResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedExternalTenantResponseArrayOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSkuInput)(nil)).Elem(), AzureSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSkuPtrInput)(nil)).Elem(), AzureSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSkuResponseInput)(nil)).Elem(), AzureSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSkuResponsePtrInput)(nil)).Elem(), AzureSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePrincipalResponseInput)(nil)).Elem(), DatabasePrincipalResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePrincipalResponseArrayInput)(nil)).Elem(), DatabasePrincipalResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseStatisticsResponseInput)(nil)).Elem(), DatabaseStatisticsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseStatisticsResponsePtrInput)(nil)).Elem(), DatabaseStatisticsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustedExternalTenantInput)(nil)).Elem(), TrustedExternalTenantArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustedExternalTenantArrayInput)(nil)).Elem(), TrustedExternalTenantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustedExternalTenantResponseInput)(nil)).Elem(), TrustedExternalTenantResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustedExternalTenantResponseArrayInput)(nil)).Elem(), TrustedExternalTenantResponseArray{})
	pulumi.RegisterOutputType(AzureSkuOutput{})
	pulumi.RegisterOutputType(AzureSkuPtrOutput{})
	pulumi.RegisterOutputType(AzureSkuResponseOutput{})
	pulumi.RegisterOutputType(AzureSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(DatabasePrincipalResponseOutput{})
	pulumi.RegisterOutputType(DatabasePrincipalResponseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseStatisticsResponseOutput{})
	pulumi.RegisterOutputType(DatabaseStatisticsResponsePtrOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantArrayOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantResponseOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantResponseArrayOutput{})
}
