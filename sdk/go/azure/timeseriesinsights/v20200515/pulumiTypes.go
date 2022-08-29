


package v20200515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvironmentStateDetailsResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}

type EnvironmentStateDetailsResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentStateDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentStateDetailsResponse)(nil)).Elem()
}

func (o EnvironmentStateDetailsResponseOutput) ToEnvironmentStateDetailsResponseOutput() EnvironmentStateDetailsResponseOutput {
	return o
}

func (o EnvironmentStateDetailsResponseOutput) ToEnvironmentStateDetailsResponseOutputWithContext(ctx context.Context) EnvironmentStateDetailsResponseOutput {
	return o
}

func (o EnvironmentStateDetailsResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentStateDetailsResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o EnvironmentStateDetailsResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentStateDetailsResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type EnvironmentStatusResponse struct {
	Ingress     IngressEnvironmentStatusResponse     `pulumi:"ingress"`
	WarmStorage WarmStorageEnvironmentStatusResponse `pulumi:"warmStorage"`
}

type EnvironmentStatusResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentStatusResponse)(nil)).Elem()
}

func (o EnvironmentStatusResponseOutput) ToEnvironmentStatusResponseOutput() EnvironmentStatusResponseOutput {
	return o
}

func (o EnvironmentStatusResponseOutput) ToEnvironmentStatusResponseOutputWithContext(ctx context.Context) EnvironmentStatusResponseOutput {
	return o
}

func (o EnvironmentStatusResponseOutput) Ingress() IngressEnvironmentStatusResponseOutput {
	return o.ApplyT(func(v EnvironmentStatusResponse) IngressEnvironmentStatusResponse { return v.Ingress }).(IngressEnvironmentStatusResponseOutput)
}

func (o EnvironmentStatusResponseOutput) WarmStorage() WarmStorageEnvironmentStatusResponseOutput {
	return o.ApplyT(func(v EnvironmentStatusResponse) WarmStorageEnvironmentStatusResponse { return v.WarmStorage }).(WarmStorageEnvironmentStatusResponseOutput)
}

type Gen2StorageConfigurationInput struct {
	AccountName   string `pulumi:"accountName"`
	ManagementKey string `pulumi:"managementKey"`
}





type Gen2StorageConfigurationInputInput interface {
	pulumi.Input

	ToGen2StorageConfigurationInputOutput() Gen2StorageConfigurationInputOutput
	ToGen2StorageConfigurationInputOutputWithContext(context.Context) Gen2StorageConfigurationInputOutput
}

type Gen2StorageConfigurationInputArgs struct {
	AccountName   pulumi.StringInput `pulumi:"accountName"`
	ManagementKey pulumi.StringInput `pulumi:"managementKey"`
}

func (Gen2StorageConfigurationInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Gen2StorageConfigurationInput)(nil)).Elem()
}

func (i Gen2StorageConfigurationInputArgs) ToGen2StorageConfigurationInputOutput() Gen2StorageConfigurationInputOutput {
	return i.ToGen2StorageConfigurationInputOutputWithContext(context.Background())
}

func (i Gen2StorageConfigurationInputArgs) ToGen2StorageConfigurationInputOutputWithContext(ctx context.Context) Gen2StorageConfigurationInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Gen2StorageConfigurationInputOutput)
}

type Gen2StorageConfigurationInputOutput struct{ *pulumi.OutputState }

func (Gen2StorageConfigurationInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Gen2StorageConfigurationInput)(nil)).Elem()
}

func (o Gen2StorageConfigurationInputOutput) ToGen2StorageConfigurationInputOutput() Gen2StorageConfigurationInputOutput {
	return o
}

func (o Gen2StorageConfigurationInputOutput) ToGen2StorageConfigurationInputOutputWithContext(ctx context.Context) Gen2StorageConfigurationInputOutput {
	return o
}

func (o Gen2StorageConfigurationInputOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v Gen2StorageConfigurationInput) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o Gen2StorageConfigurationInputOutput) ManagementKey() pulumi.StringOutput {
	return o.ApplyT(func(v Gen2StorageConfigurationInput) string { return v.ManagementKey }).(pulumi.StringOutput)
}

type Gen2StorageConfigurationOutputResponse struct {
	AccountName string `pulumi:"accountName"`
}

type Gen2StorageConfigurationOutputResponseOutput struct{ *pulumi.OutputState }

func (Gen2StorageConfigurationOutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Gen2StorageConfigurationOutputResponse)(nil)).Elem()
}

func (o Gen2StorageConfigurationOutputResponseOutput) ToGen2StorageConfigurationOutputResponseOutput() Gen2StorageConfigurationOutputResponseOutput {
	return o
}

func (o Gen2StorageConfigurationOutputResponseOutput) ToGen2StorageConfigurationOutputResponseOutputWithContext(ctx context.Context) Gen2StorageConfigurationOutputResponseOutput {
	return o
}

func (o Gen2StorageConfigurationOutputResponseOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v Gen2StorageConfigurationOutputResponse) string { return v.AccountName }).(pulumi.StringOutput)
}

type IngressEnvironmentStatusResponse struct {
	State        *string                         `pulumi:"state"`
	StateDetails EnvironmentStateDetailsResponse `pulumi:"stateDetails"`
}

type IngressEnvironmentStatusResponseOutput struct{ *pulumi.OutputState }

func (IngressEnvironmentStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressEnvironmentStatusResponse)(nil)).Elem()
}

func (o IngressEnvironmentStatusResponseOutput) ToIngressEnvironmentStatusResponseOutput() IngressEnvironmentStatusResponseOutput {
	return o
}

func (o IngressEnvironmentStatusResponseOutput) ToIngressEnvironmentStatusResponseOutputWithContext(ctx context.Context) IngressEnvironmentStatusResponseOutput {
	return o
}

func (o IngressEnvironmentStatusResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngressEnvironmentStatusResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o IngressEnvironmentStatusResponseOutput) StateDetails() EnvironmentStateDetailsResponseOutput {
	return o.ApplyT(func(v IngressEnvironmentStatusResponse) EnvironmentStateDetailsResponse { return v.StateDetails }).(EnvironmentStateDetailsResponseOutput)
}

type LocalTimestamp struct {
	Format         *string                       `pulumi:"format"`
	TimeZoneOffset *LocalTimestampTimeZoneOffset `pulumi:"timeZoneOffset"`
}





type LocalTimestampInput interface {
	pulumi.Input

	ToLocalTimestampOutput() LocalTimestampOutput
	ToLocalTimestampOutputWithContext(context.Context) LocalTimestampOutput
}

type LocalTimestampArgs struct {
	Format         pulumi.StringPtrInput                `pulumi:"format"`
	TimeZoneOffset LocalTimestampTimeZoneOffsetPtrInput `pulumi:"timeZoneOffset"`
}

func (LocalTimestampArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalTimestamp)(nil)).Elem()
}

func (i LocalTimestampArgs) ToLocalTimestampOutput() LocalTimestampOutput {
	return i.ToLocalTimestampOutputWithContext(context.Background())
}

func (i LocalTimestampArgs) ToLocalTimestampOutputWithContext(ctx context.Context) LocalTimestampOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTimestampOutput)
}

func (i LocalTimestampArgs) ToLocalTimestampPtrOutput() LocalTimestampPtrOutput {
	return i.ToLocalTimestampPtrOutputWithContext(context.Background())
}

func (i LocalTimestampArgs) ToLocalTimestampPtrOutputWithContext(ctx context.Context) LocalTimestampPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTimestampOutput).ToLocalTimestampPtrOutputWithContext(ctx)
}









type LocalTimestampPtrInput interface {
	pulumi.Input

	ToLocalTimestampPtrOutput() LocalTimestampPtrOutput
	ToLocalTimestampPtrOutputWithContext(context.Context) LocalTimestampPtrOutput
}

type localTimestampPtrType LocalTimestampArgs

func LocalTimestampPtr(v *LocalTimestampArgs) LocalTimestampPtrInput {
	return (*localTimestampPtrType)(v)
}

func (*localTimestampPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalTimestamp)(nil)).Elem()
}

func (i *localTimestampPtrType) ToLocalTimestampPtrOutput() LocalTimestampPtrOutput {
	return i.ToLocalTimestampPtrOutputWithContext(context.Background())
}

func (i *localTimestampPtrType) ToLocalTimestampPtrOutputWithContext(ctx context.Context) LocalTimestampPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTimestampPtrOutput)
}

type LocalTimestampOutput struct{ *pulumi.OutputState }

func (LocalTimestampOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalTimestamp)(nil)).Elem()
}

func (o LocalTimestampOutput) ToLocalTimestampOutput() LocalTimestampOutput {
	return o
}

func (o LocalTimestampOutput) ToLocalTimestampOutputWithContext(ctx context.Context) LocalTimestampOutput {
	return o
}

func (o LocalTimestampOutput) ToLocalTimestampPtrOutput() LocalTimestampPtrOutput {
	return o.ToLocalTimestampPtrOutputWithContext(context.Background())
}

func (o LocalTimestampOutput) ToLocalTimestampPtrOutputWithContext(ctx context.Context) LocalTimestampPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalTimestamp) *LocalTimestamp {
		return &v
	}).(LocalTimestampPtrOutput)
}

func (o LocalTimestampOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalTimestamp) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o LocalTimestampOutput) TimeZoneOffset() LocalTimestampTimeZoneOffsetPtrOutput {
	return o.ApplyT(func(v LocalTimestamp) *LocalTimestampTimeZoneOffset { return v.TimeZoneOffset }).(LocalTimestampTimeZoneOffsetPtrOutput)
}

type LocalTimestampPtrOutput struct{ *pulumi.OutputState }

func (LocalTimestampPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalTimestamp)(nil)).Elem()
}

func (o LocalTimestampPtrOutput) ToLocalTimestampPtrOutput() LocalTimestampPtrOutput {
	return o
}

func (o LocalTimestampPtrOutput) ToLocalTimestampPtrOutputWithContext(ctx context.Context) LocalTimestampPtrOutput {
	return o
}

func (o LocalTimestampPtrOutput) Elem() LocalTimestampOutput {
	return o.ApplyT(func(v *LocalTimestamp) LocalTimestamp {
		if v != nil {
			return *v
		}
		var ret LocalTimestamp
		return ret
	}).(LocalTimestampOutput)
}

func (o LocalTimestampPtrOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTimestamp) *string {
		if v == nil {
			return nil
		}
		return v.Format
	}).(pulumi.StringPtrOutput)
}

func (o LocalTimestampPtrOutput) TimeZoneOffset() LocalTimestampTimeZoneOffsetPtrOutput {
	return o.ApplyT(func(v *LocalTimestamp) *LocalTimestampTimeZoneOffset {
		if v == nil {
			return nil
		}
		return v.TimeZoneOffset
	}).(LocalTimestampTimeZoneOffsetPtrOutput)
}

type LocalTimestampResponse struct {
	Format         *string                               `pulumi:"format"`
	TimeZoneOffset *LocalTimestampResponseTimeZoneOffset `pulumi:"timeZoneOffset"`
}

type LocalTimestampResponseOutput struct{ *pulumi.OutputState }

func (LocalTimestampResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalTimestampResponse)(nil)).Elem()
}

func (o LocalTimestampResponseOutput) ToLocalTimestampResponseOutput() LocalTimestampResponseOutput {
	return o
}

func (o LocalTimestampResponseOutput) ToLocalTimestampResponseOutputWithContext(ctx context.Context) LocalTimestampResponseOutput {
	return o
}

func (o LocalTimestampResponseOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalTimestampResponse) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o LocalTimestampResponseOutput) TimeZoneOffset() LocalTimestampResponseTimeZoneOffsetPtrOutput {
	return o.ApplyT(func(v LocalTimestampResponse) *LocalTimestampResponseTimeZoneOffset { return v.TimeZoneOffset }).(LocalTimestampResponseTimeZoneOffsetPtrOutput)
}

type LocalTimestampResponsePtrOutput struct{ *pulumi.OutputState }

func (LocalTimestampResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalTimestampResponse)(nil)).Elem()
}

func (o LocalTimestampResponsePtrOutput) ToLocalTimestampResponsePtrOutput() LocalTimestampResponsePtrOutput {
	return o
}

func (o LocalTimestampResponsePtrOutput) ToLocalTimestampResponsePtrOutputWithContext(ctx context.Context) LocalTimestampResponsePtrOutput {
	return o
}

func (o LocalTimestampResponsePtrOutput) Elem() LocalTimestampResponseOutput {
	return o.ApplyT(func(v *LocalTimestampResponse) LocalTimestampResponse {
		if v != nil {
			return *v
		}
		var ret LocalTimestampResponse
		return ret
	}).(LocalTimestampResponseOutput)
}

func (o LocalTimestampResponsePtrOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTimestampResponse) *string {
		if v == nil {
			return nil
		}
		return v.Format
	}).(pulumi.StringPtrOutput)
}

func (o LocalTimestampResponsePtrOutput) TimeZoneOffset() LocalTimestampResponseTimeZoneOffsetPtrOutput {
	return o.ApplyT(func(v *LocalTimestampResponse) *LocalTimestampResponseTimeZoneOffset {
		if v == nil {
			return nil
		}
		return v.TimeZoneOffset
	}).(LocalTimestampResponseTimeZoneOffsetPtrOutput)
}

type LocalTimestampResponseTimeZoneOffset struct {
	PropertyName *string `pulumi:"propertyName"`
}

type LocalTimestampResponseTimeZoneOffsetOutput struct{ *pulumi.OutputState }

func (LocalTimestampResponseTimeZoneOffsetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalTimestampResponseTimeZoneOffset)(nil)).Elem()
}

func (o LocalTimestampResponseTimeZoneOffsetOutput) ToLocalTimestampResponseTimeZoneOffsetOutput() LocalTimestampResponseTimeZoneOffsetOutput {
	return o
}

func (o LocalTimestampResponseTimeZoneOffsetOutput) ToLocalTimestampResponseTimeZoneOffsetOutputWithContext(ctx context.Context) LocalTimestampResponseTimeZoneOffsetOutput {
	return o
}

func (o LocalTimestampResponseTimeZoneOffsetOutput) PropertyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalTimestampResponseTimeZoneOffset) *string { return v.PropertyName }).(pulumi.StringPtrOutput)
}

type LocalTimestampResponseTimeZoneOffsetPtrOutput struct{ *pulumi.OutputState }

func (LocalTimestampResponseTimeZoneOffsetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalTimestampResponseTimeZoneOffset)(nil)).Elem()
}

func (o LocalTimestampResponseTimeZoneOffsetPtrOutput) ToLocalTimestampResponseTimeZoneOffsetPtrOutput() LocalTimestampResponseTimeZoneOffsetPtrOutput {
	return o
}

func (o LocalTimestampResponseTimeZoneOffsetPtrOutput) ToLocalTimestampResponseTimeZoneOffsetPtrOutputWithContext(ctx context.Context) LocalTimestampResponseTimeZoneOffsetPtrOutput {
	return o
}

func (o LocalTimestampResponseTimeZoneOffsetPtrOutput) Elem() LocalTimestampResponseTimeZoneOffsetOutput {
	return o.ApplyT(func(v *LocalTimestampResponseTimeZoneOffset) LocalTimestampResponseTimeZoneOffset {
		if v != nil {
			return *v
		}
		var ret LocalTimestampResponseTimeZoneOffset
		return ret
	}).(LocalTimestampResponseTimeZoneOffsetOutput)
}

func (o LocalTimestampResponseTimeZoneOffsetPtrOutput) PropertyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTimestampResponseTimeZoneOffset) *string {
		if v == nil {
			return nil
		}
		return v.PropertyName
	}).(pulumi.StringPtrOutput)
}

type LocalTimestampTimeZoneOffset struct {
	PropertyName *string `pulumi:"propertyName"`
}





type LocalTimestampTimeZoneOffsetInput interface {
	pulumi.Input

	ToLocalTimestampTimeZoneOffsetOutput() LocalTimestampTimeZoneOffsetOutput
	ToLocalTimestampTimeZoneOffsetOutputWithContext(context.Context) LocalTimestampTimeZoneOffsetOutput
}

type LocalTimestampTimeZoneOffsetArgs struct {
	PropertyName pulumi.StringPtrInput `pulumi:"propertyName"`
}

func (LocalTimestampTimeZoneOffsetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalTimestampTimeZoneOffset)(nil)).Elem()
}

func (i LocalTimestampTimeZoneOffsetArgs) ToLocalTimestampTimeZoneOffsetOutput() LocalTimestampTimeZoneOffsetOutput {
	return i.ToLocalTimestampTimeZoneOffsetOutputWithContext(context.Background())
}

func (i LocalTimestampTimeZoneOffsetArgs) ToLocalTimestampTimeZoneOffsetOutputWithContext(ctx context.Context) LocalTimestampTimeZoneOffsetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTimestampTimeZoneOffsetOutput)
}

func (i LocalTimestampTimeZoneOffsetArgs) ToLocalTimestampTimeZoneOffsetPtrOutput() LocalTimestampTimeZoneOffsetPtrOutput {
	return i.ToLocalTimestampTimeZoneOffsetPtrOutputWithContext(context.Background())
}

func (i LocalTimestampTimeZoneOffsetArgs) ToLocalTimestampTimeZoneOffsetPtrOutputWithContext(ctx context.Context) LocalTimestampTimeZoneOffsetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTimestampTimeZoneOffsetOutput).ToLocalTimestampTimeZoneOffsetPtrOutputWithContext(ctx)
}









type LocalTimestampTimeZoneOffsetPtrInput interface {
	pulumi.Input

	ToLocalTimestampTimeZoneOffsetPtrOutput() LocalTimestampTimeZoneOffsetPtrOutput
	ToLocalTimestampTimeZoneOffsetPtrOutputWithContext(context.Context) LocalTimestampTimeZoneOffsetPtrOutput
}

type localTimestampTimeZoneOffsetPtrType LocalTimestampTimeZoneOffsetArgs

func LocalTimestampTimeZoneOffsetPtr(v *LocalTimestampTimeZoneOffsetArgs) LocalTimestampTimeZoneOffsetPtrInput {
	return (*localTimestampTimeZoneOffsetPtrType)(v)
}

func (*localTimestampTimeZoneOffsetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalTimestampTimeZoneOffset)(nil)).Elem()
}

func (i *localTimestampTimeZoneOffsetPtrType) ToLocalTimestampTimeZoneOffsetPtrOutput() LocalTimestampTimeZoneOffsetPtrOutput {
	return i.ToLocalTimestampTimeZoneOffsetPtrOutputWithContext(context.Background())
}

func (i *localTimestampTimeZoneOffsetPtrType) ToLocalTimestampTimeZoneOffsetPtrOutputWithContext(ctx context.Context) LocalTimestampTimeZoneOffsetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTimestampTimeZoneOffsetPtrOutput)
}

type LocalTimestampTimeZoneOffsetOutput struct{ *pulumi.OutputState }

func (LocalTimestampTimeZoneOffsetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalTimestampTimeZoneOffset)(nil)).Elem()
}

func (o LocalTimestampTimeZoneOffsetOutput) ToLocalTimestampTimeZoneOffsetOutput() LocalTimestampTimeZoneOffsetOutput {
	return o
}

func (o LocalTimestampTimeZoneOffsetOutput) ToLocalTimestampTimeZoneOffsetOutputWithContext(ctx context.Context) LocalTimestampTimeZoneOffsetOutput {
	return o
}

func (o LocalTimestampTimeZoneOffsetOutput) ToLocalTimestampTimeZoneOffsetPtrOutput() LocalTimestampTimeZoneOffsetPtrOutput {
	return o.ToLocalTimestampTimeZoneOffsetPtrOutputWithContext(context.Background())
}

func (o LocalTimestampTimeZoneOffsetOutput) ToLocalTimestampTimeZoneOffsetPtrOutputWithContext(ctx context.Context) LocalTimestampTimeZoneOffsetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalTimestampTimeZoneOffset) *LocalTimestampTimeZoneOffset {
		return &v
	}).(LocalTimestampTimeZoneOffsetPtrOutput)
}

func (o LocalTimestampTimeZoneOffsetOutput) PropertyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalTimestampTimeZoneOffset) *string { return v.PropertyName }).(pulumi.StringPtrOutput)
}

type LocalTimestampTimeZoneOffsetPtrOutput struct{ *pulumi.OutputState }

func (LocalTimestampTimeZoneOffsetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalTimestampTimeZoneOffset)(nil)).Elem()
}

func (o LocalTimestampTimeZoneOffsetPtrOutput) ToLocalTimestampTimeZoneOffsetPtrOutput() LocalTimestampTimeZoneOffsetPtrOutput {
	return o
}

func (o LocalTimestampTimeZoneOffsetPtrOutput) ToLocalTimestampTimeZoneOffsetPtrOutputWithContext(ctx context.Context) LocalTimestampTimeZoneOffsetPtrOutput {
	return o
}

func (o LocalTimestampTimeZoneOffsetPtrOutput) Elem() LocalTimestampTimeZoneOffsetOutput {
	return o.ApplyT(func(v *LocalTimestampTimeZoneOffset) LocalTimestampTimeZoneOffset {
		if v != nil {
			return *v
		}
		var ret LocalTimestampTimeZoneOffset
		return ret
	}).(LocalTimestampTimeZoneOffsetOutput)
}

func (o LocalTimestampTimeZoneOffsetPtrOutput) PropertyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTimestampTimeZoneOffset) *string {
		if v == nil {
			return nil
		}
		return v.PropertyName
	}).(pulumi.StringPtrOutput)
}

type ReferenceDataSetKeyProperty struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ReferenceDataSetKeyPropertyInput interface {
	pulumi.Input

	ToReferenceDataSetKeyPropertyOutput() ReferenceDataSetKeyPropertyOutput
	ToReferenceDataSetKeyPropertyOutputWithContext(context.Context) ReferenceDataSetKeyPropertyOutput
}

type ReferenceDataSetKeyPropertyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ReferenceDataSetKeyPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceDataSetKeyProperty)(nil)).Elem()
}

func (i ReferenceDataSetKeyPropertyArgs) ToReferenceDataSetKeyPropertyOutput() ReferenceDataSetKeyPropertyOutput {
	return i.ToReferenceDataSetKeyPropertyOutputWithContext(context.Background())
}

func (i ReferenceDataSetKeyPropertyArgs) ToReferenceDataSetKeyPropertyOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceDataSetKeyPropertyOutput)
}





type ReferenceDataSetKeyPropertyArrayInput interface {
	pulumi.Input

	ToReferenceDataSetKeyPropertyArrayOutput() ReferenceDataSetKeyPropertyArrayOutput
	ToReferenceDataSetKeyPropertyArrayOutputWithContext(context.Context) ReferenceDataSetKeyPropertyArrayOutput
}

type ReferenceDataSetKeyPropertyArray []ReferenceDataSetKeyPropertyInput

func (ReferenceDataSetKeyPropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReferenceDataSetKeyProperty)(nil)).Elem()
}

func (i ReferenceDataSetKeyPropertyArray) ToReferenceDataSetKeyPropertyArrayOutput() ReferenceDataSetKeyPropertyArrayOutput {
	return i.ToReferenceDataSetKeyPropertyArrayOutputWithContext(context.Background())
}

func (i ReferenceDataSetKeyPropertyArray) ToReferenceDataSetKeyPropertyArrayOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceDataSetKeyPropertyArrayOutput)
}

type ReferenceDataSetKeyPropertyOutput struct{ *pulumi.OutputState }

func (ReferenceDataSetKeyPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceDataSetKeyProperty)(nil)).Elem()
}

func (o ReferenceDataSetKeyPropertyOutput) ToReferenceDataSetKeyPropertyOutput() ReferenceDataSetKeyPropertyOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyOutput) ToReferenceDataSetKeyPropertyOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceDataSetKeyProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ReferenceDataSetKeyPropertyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceDataSetKeyProperty) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ReferenceDataSetKeyPropertyArrayOutput struct{ *pulumi.OutputState }

func (ReferenceDataSetKeyPropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReferenceDataSetKeyProperty)(nil)).Elem()
}

func (o ReferenceDataSetKeyPropertyArrayOutput) ToReferenceDataSetKeyPropertyArrayOutput() ReferenceDataSetKeyPropertyArrayOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyArrayOutput) ToReferenceDataSetKeyPropertyArrayOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyArrayOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyArrayOutput) Index(i pulumi.IntInput) ReferenceDataSetKeyPropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReferenceDataSetKeyProperty {
		return vs[0].([]ReferenceDataSetKeyProperty)[vs[1].(int)]
	}).(ReferenceDataSetKeyPropertyOutput)
}

type ReferenceDataSetKeyPropertyResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type ReferenceDataSetKeyPropertyResponseOutput struct{ *pulumi.OutputState }

func (ReferenceDataSetKeyPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceDataSetKeyPropertyResponse)(nil)).Elem()
}

func (o ReferenceDataSetKeyPropertyResponseOutput) ToReferenceDataSetKeyPropertyResponseOutput() ReferenceDataSetKeyPropertyResponseOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyResponseOutput) ToReferenceDataSetKeyPropertyResponseOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyResponseOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceDataSetKeyPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ReferenceDataSetKeyPropertyResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceDataSetKeyPropertyResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ReferenceDataSetKeyPropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (ReferenceDataSetKeyPropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReferenceDataSetKeyPropertyResponse)(nil)).Elem()
}

func (o ReferenceDataSetKeyPropertyResponseArrayOutput) ToReferenceDataSetKeyPropertyResponseArrayOutput() ReferenceDataSetKeyPropertyResponseArrayOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyResponseArrayOutput) ToReferenceDataSetKeyPropertyResponseArrayOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyResponseArrayOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyResponseArrayOutput) Index(i pulumi.IntInput) ReferenceDataSetKeyPropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReferenceDataSetKeyPropertyResponse {
		return vs[0].([]ReferenceDataSetKeyPropertyResponse)[vs[1].(int)]
	}).(ReferenceDataSetKeyPropertyResponseOutput)
}

type Sku struct {
	Capacity int    `pulumi:"capacity"`
	Name     string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntInput    `pulumi:"capacity"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v Sku) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponse struct {
	Capacity int    `pulumi:"capacity"`
	Name     string `pulumi:"name"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v SkuResponse) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type TimeSeriesIdProperty struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type TimeSeriesIdPropertyInput interface {
	pulumi.Input

	ToTimeSeriesIdPropertyOutput() TimeSeriesIdPropertyOutput
	ToTimeSeriesIdPropertyOutputWithContext(context.Context) TimeSeriesIdPropertyOutput
}

type TimeSeriesIdPropertyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (TimeSeriesIdPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesIdProperty)(nil)).Elem()
}

func (i TimeSeriesIdPropertyArgs) ToTimeSeriesIdPropertyOutput() TimeSeriesIdPropertyOutput {
	return i.ToTimeSeriesIdPropertyOutputWithContext(context.Background())
}

func (i TimeSeriesIdPropertyArgs) ToTimeSeriesIdPropertyOutputWithContext(ctx context.Context) TimeSeriesIdPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesIdPropertyOutput)
}





type TimeSeriesIdPropertyArrayInput interface {
	pulumi.Input

	ToTimeSeriesIdPropertyArrayOutput() TimeSeriesIdPropertyArrayOutput
	ToTimeSeriesIdPropertyArrayOutputWithContext(context.Context) TimeSeriesIdPropertyArrayOutput
}

type TimeSeriesIdPropertyArray []TimeSeriesIdPropertyInput

func (TimeSeriesIdPropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeSeriesIdProperty)(nil)).Elem()
}

func (i TimeSeriesIdPropertyArray) ToTimeSeriesIdPropertyArrayOutput() TimeSeriesIdPropertyArrayOutput {
	return i.ToTimeSeriesIdPropertyArrayOutputWithContext(context.Background())
}

func (i TimeSeriesIdPropertyArray) ToTimeSeriesIdPropertyArrayOutputWithContext(ctx context.Context) TimeSeriesIdPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesIdPropertyArrayOutput)
}

type TimeSeriesIdPropertyOutput struct{ *pulumi.OutputState }

func (TimeSeriesIdPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesIdProperty)(nil)).Elem()
}

func (o TimeSeriesIdPropertyOutput) ToTimeSeriesIdPropertyOutput() TimeSeriesIdPropertyOutput {
	return o
}

func (o TimeSeriesIdPropertyOutput) ToTimeSeriesIdPropertyOutputWithContext(ctx context.Context) TimeSeriesIdPropertyOutput {
	return o
}

func (o TimeSeriesIdPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimeSeriesIdProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TimeSeriesIdPropertyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimeSeriesIdProperty) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type TimeSeriesIdPropertyArrayOutput struct{ *pulumi.OutputState }

func (TimeSeriesIdPropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeSeriesIdProperty)(nil)).Elem()
}

func (o TimeSeriesIdPropertyArrayOutput) ToTimeSeriesIdPropertyArrayOutput() TimeSeriesIdPropertyArrayOutput {
	return o
}

func (o TimeSeriesIdPropertyArrayOutput) ToTimeSeriesIdPropertyArrayOutputWithContext(ctx context.Context) TimeSeriesIdPropertyArrayOutput {
	return o
}

func (o TimeSeriesIdPropertyArrayOutput) Index(i pulumi.IntInput) TimeSeriesIdPropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeSeriesIdProperty {
		return vs[0].([]TimeSeriesIdProperty)[vs[1].(int)]
	}).(TimeSeriesIdPropertyOutput)
}

type TimeSeriesIdPropertyResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type TimeSeriesIdPropertyResponseOutput struct{ *pulumi.OutputState }

func (TimeSeriesIdPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesIdPropertyResponse)(nil)).Elem()
}

func (o TimeSeriesIdPropertyResponseOutput) ToTimeSeriesIdPropertyResponseOutput() TimeSeriesIdPropertyResponseOutput {
	return o
}

func (o TimeSeriesIdPropertyResponseOutput) ToTimeSeriesIdPropertyResponseOutputWithContext(ctx context.Context) TimeSeriesIdPropertyResponseOutput {
	return o
}

func (o TimeSeriesIdPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimeSeriesIdPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TimeSeriesIdPropertyResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimeSeriesIdPropertyResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type TimeSeriesIdPropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (TimeSeriesIdPropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeSeriesIdPropertyResponse)(nil)).Elem()
}

func (o TimeSeriesIdPropertyResponseArrayOutput) ToTimeSeriesIdPropertyResponseArrayOutput() TimeSeriesIdPropertyResponseArrayOutput {
	return o
}

func (o TimeSeriesIdPropertyResponseArrayOutput) ToTimeSeriesIdPropertyResponseArrayOutputWithContext(ctx context.Context) TimeSeriesIdPropertyResponseArrayOutput {
	return o
}

func (o TimeSeriesIdPropertyResponseArrayOutput) Index(i pulumi.IntInput) TimeSeriesIdPropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeSeriesIdPropertyResponse {
		return vs[0].([]TimeSeriesIdPropertyResponse)[vs[1].(int)]
	}).(TimeSeriesIdPropertyResponseOutput)
}

type WarmStorageEnvironmentStatusResponse struct {
	CurrentCount *int    `pulumi:"currentCount"`
	MaxCount     *int    `pulumi:"maxCount"`
	State        *string `pulumi:"state"`
}

type WarmStorageEnvironmentStatusResponseOutput struct{ *pulumi.OutputState }

func (WarmStorageEnvironmentStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WarmStorageEnvironmentStatusResponse)(nil)).Elem()
}

func (o WarmStorageEnvironmentStatusResponseOutput) ToWarmStorageEnvironmentStatusResponseOutput() WarmStorageEnvironmentStatusResponseOutput {
	return o
}

func (o WarmStorageEnvironmentStatusResponseOutput) ToWarmStorageEnvironmentStatusResponseOutputWithContext(ctx context.Context) WarmStorageEnvironmentStatusResponseOutput {
	return o
}

func (o WarmStorageEnvironmentStatusResponseOutput) CurrentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WarmStorageEnvironmentStatusResponse) *int { return v.CurrentCount }).(pulumi.IntPtrOutput)
}

func (o WarmStorageEnvironmentStatusResponseOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WarmStorageEnvironmentStatusResponse) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o WarmStorageEnvironmentStatusResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WarmStorageEnvironmentStatusResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type WarmStoreConfigurationProperties struct {
	DataRetention string `pulumi:"dataRetention"`
}





type WarmStoreConfigurationPropertiesInput interface {
	pulumi.Input

	ToWarmStoreConfigurationPropertiesOutput() WarmStoreConfigurationPropertiesOutput
	ToWarmStoreConfigurationPropertiesOutputWithContext(context.Context) WarmStoreConfigurationPropertiesOutput
}

type WarmStoreConfigurationPropertiesArgs struct {
	DataRetention pulumi.StringInput `pulumi:"dataRetention"`
}

func (WarmStoreConfigurationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WarmStoreConfigurationProperties)(nil)).Elem()
}

func (i WarmStoreConfigurationPropertiesArgs) ToWarmStoreConfigurationPropertiesOutput() WarmStoreConfigurationPropertiesOutput {
	return i.ToWarmStoreConfigurationPropertiesOutputWithContext(context.Background())
}

func (i WarmStoreConfigurationPropertiesArgs) ToWarmStoreConfigurationPropertiesOutputWithContext(ctx context.Context) WarmStoreConfigurationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WarmStoreConfigurationPropertiesOutput)
}

func (i WarmStoreConfigurationPropertiesArgs) ToWarmStoreConfigurationPropertiesPtrOutput() WarmStoreConfigurationPropertiesPtrOutput {
	return i.ToWarmStoreConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (i WarmStoreConfigurationPropertiesArgs) ToWarmStoreConfigurationPropertiesPtrOutputWithContext(ctx context.Context) WarmStoreConfigurationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WarmStoreConfigurationPropertiesOutput).ToWarmStoreConfigurationPropertiesPtrOutputWithContext(ctx)
}









type WarmStoreConfigurationPropertiesPtrInput interface {
	pulumi.Input

	ToWarmStoreConfigurationPropertiesPtrOutput() WarmStoreConfigurationPropertiesPtrOutput
	ToWarmStoreConfigurationPropertiesPtrOutputWithContext(context.Context) WarmStoreConfigurationPropertiesPtrOutput
}

type warmStoreConfigurationPropertiesPtrType WarmStoreConfigurationPropertiesArgs

func WarmStoreConfigurationPropertiesPtr(v *WarmStoreConfigurationPropertiesArgs) WarmStoreConfigurationPropertiesPtrInput {
	return (*warmStoreConfigurationPropertiesPtrType)(v)
}

func (*warmStoreConfigurationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WarmStoreConfigurationProperties)(nil)).Elem()
}

func (i *warmStoreConfigurationPropertiesPtrType) ToWarmStoreConfigurationPropertiesPtrOutput() WarmStoreConfigurationPropertiesPtrOutput {
	return i.ToWarmStoreConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (i *warmStoreConfigurationPropertiesPtrType) ToWarmStoreConfigurationPropertiesPtrOutputWithContext(ctx context.Context) WarmStoreConfigurationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WarmStoreConfigurationPropertiesPtrOutput)
}

type WarmStoreConfigurationPropertiesOutput struct{ *pulumi.OutputState }

func (WarmStoreConfigurationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WarmStoreConfigurationProperties)(nil)).Elem()
}

func (o WarmStoreConfigurationPropertiesOutput) ToWarmStoreConfigurationPropertiesOutput() WarmStoreConfigurationPropertiesOutput {
	return o
}

func (o WarmStoreConfigurationPropertiesOutput) ToWarmStoreConfigurationPropertiesOutputWithContext(ctx context.Context) WarmStoreConfigurationPropertiesOutput {
	return o
}

func (o WarmStoreConfigurationPropertiesOutput) ToWarmStoreConfigurationPropertiesPtrOutput() WarmStoreConfigurationPropertiesPtrOutput {
	return o.ToWarmStoreConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (o WarmStoreConfigurationPropertiesOutput) ToWarmStoreConfigurationPropertiesPtrOutputWithContext(ctx context.Context) WarmStoreConfigurationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WarmStoreConfigurationProperties) *WarmStoreConfigurationProperties {
		return &v
	}).(WarmStoreConfigurationPropertiesPtrOutput)
}

func (o WarmStoreConfigurationPropertiesOutput) DataRetention() pulumi.StringOutput {
	return o.ApplyT(func(v WarmStoreConfigurationProperties) string { return v.DataRetention }).(pulumi.StringOutput)
}

type WarmStoreConfigurationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (WarmStoreConfigurationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WarmStoreConfigurationProperties)(nil)).Elem()
}

func (o WarmStoreConfigurationPropertiesPtrOutput) ToWarmStoreConfigurationPropertiesPtrOutput() WarmStoreConfigurationPropertiesPtrOutput {
	return o
}

func (o WarmStoreConfigurationPropertiesPtrOutput) ToWarmStoreConfigurationPropertiesPtrOutputWithContext(ctx context.Context) WarmStoreConfigurationPropertiesPtrOutput {
	return o
}

func (o WarmStoreConfigurationPropertiesPtrOutput) Elem() WarmStoreConfigurationPropertiesOutput {
	return o.ApplyT(func(v *WarmStoreConfigurationProperties) WarmStoreConfigurationProperties {
		if v != nil {
			return *v
		}
		var ret WarmStoreConfigurationProperties
		return ret
	}).(WarmStoreConfigurationPropertiesOutput)
}

func (o WarmStoreConfigurationPropertiesPtrOutput) DataRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WarmStoreConfigurationProperties) *string {
		if v == nil {
			return nil
		}
		return &v.DataRetention
	}).(pulumi.StringPtrOutput)
}

type WarmStoreConfigurationPropertiesResponse struct {
	DataRetention string `pulumi:"dataRetention"`
}

type WarmStoreConfigurationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (WarmStoreConfigurationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WarmStoreConfigurationPropertiesResponse)(nil)).Elem()
}

func (o WarmStoreConfigurationPropertiesResponseOutput) ToWarmStoreConfigurationPropertiesResponseOutput() WarmStoreConfigurationPropertiesResponseOutput {
	return o
}

func (o WarmStoreConfigurationPropertiesResponseOutput) ToWarmStoreConfigurationPropertiesResponseOutputWithContext(ctx context.Context) WarmStoreConfigurationPropertiesResponseOutput {
	return o
}

func (o WarmStoreConfigurationPropertiesResponseOutput) DataRetention() pulumi.StringOutput {
	return o.ApplyT(func(v WarmStoreConfigurationPropertiesResponse) string { return v.DataRetention }).(pulumi.StringOutput)
}

type WarmStoreConfigurationPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (WarmStoreConfigurationPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WarmStoreConfigurationPropertiesResponse)(nil)).Elem()
}

func (o WarmStoreConfigurationPropertiesResponsePtrOutput) ToWarmStoreConfigurationPropertiesResponsePtrOutput() WarmStoreConfigurationPropertiesResponsePtrOutput {
	return o
}

func (o WarmStoreConfigurationPropertiesResponsePtrOutput) ToWarmStoreConfigurationPropertiesResponsePtrOutputWithContext(ctx context.Context) WarmStoreConfigurationPropertiesResponsePtrOutput {
	return o
}

func (o WarmStoreConfigurationPropertiesResponsePtrOutput) Elem() WarmStoreConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v *WarmStoreConfigurationPropertiesResponse) WarmStoreConfigurationPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret WarmStoreConfigurationPropertiesResponse
		return ret
	}).(WarmStoreConfigurationPropertiesResponseOutput)
}

func (o WarmStoreConfigurationPropertiesResponsePtrOutput) DataRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WarmStoreConfigurationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DataRetention
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentStateDetailsResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentStatusResponseOutput{})
	pulumi.RegisterOutputType(Gen2StorageConfigurationInputOutput{})
	pulumi.RegisterOutputType(Gen2StorageConfigurationOutputResponseOutput{})
	pulumi.RegisterOutputType(IngressEnvironmentStatusResponseOutput{})
	pulumi.RegisterOutputType(LocalTimestampOutput{})
	pulumi.RegisterOutputType(LocalTimestampPtrOutput{})
	pulumi.RegisterOutputType(LocalTimestampResponseOutput{})
	pulumi.RegisterOutputType(LocalTimestampResponsePtrOutput{})
	pulumi.RegisterOutputType(LocalTimestampResponseTimeZoneOffsetOutput{})
	pulumi.RegisterOutputType(LocalTimestampResponseTimeZoneOffsetPtrOutput{})
	pulumi.RegisterOutputType(LocalTimestampTimeZoneOffsetOutput{})
	pulumi.RegisterOutputType(LocalTimestampTimeZoneOffsetPtrOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyArrayOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyResponseOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(TimeSeriesIdPropertyOutput{})
	pulumi.RegisterOutputType(TimeSeriesIdPropertyArrayOutput{})
	pulumi.RegisterOutputType(TimeSeriesIdPropertyResponseOutput{})
	pulumi.RegisterOutputType(TimeSeriesIdPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(WarmStorageEnvironmentStatusResponseOutput{})
	pulumi.RegisterOutputType(WarmStoreConfigurationPropertiesOutput{})
	pulumi.RegisterOutputType(WarmStoreConfigurationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(WarmStoreConfigurationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(WarmStoreConfigurationPropertiesResponsePtrOutput{})
}
