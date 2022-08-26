


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DefenderSettingsPropertiesMdeIntegration struct {
	Status string `pulumi:"status"`
}





type DefenderSettingsPropertiesMdeIntegrationInput interface {
	pulumi.Input

	ToDefenderSettingsPropertiesMdeIntegrationOutput() DefenderSettingsPropertiesMdeIntegrationOutput
	ToDefenderSettingsPropertiesMdeIntegrationOutputWithContext(context.Context) DefenderSettingsPropertiesMdeIntegrationOutput
}

type DefenderSettingsPropertiesMdeIntegrationArgs struct {
	Status pulumi.StringInput `pulumi:"status"`
}

func (DefenderSettingsPropertiesMdeIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderSettingsPropertiesMdeIntegration)(nil)).Elem()
}

func (i DefenderSettingsPropertiesMdeIntegrationArgs) ToDefenderSettingsPropertiesMdeIntegrationOutput() DefenderSettingsPropertiesMdeIntegrationOutput {
	return i.ToDefenderSettingsPropertiesMdeIntegrationOutputWithContext(context.Background())
}

func (i DefenderSettingsPropertiesMdeIntegrationArgs) ToDefenderSettingsPropertiesMdeIntegrationOutputWithContext(ctx context.Context) DefenderSettingsPropertiesMdeIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderSettingsPropertiesMdeIntegrationOutput)
}

type DefenderSettingsPropertiesMdeIntegrationOutput struct{ *pulumi.OutputState }

func (DefenderSettingsPropertiesMdeIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderSettingsPropertiesMdeIntegration)(nil)).Elem()
}

func (o DefenderSettingsPropertiesMdeIntegrationOutput) ToDefenderSettingsPropertiesMdeIntegrationOutput() DefenderSettingsPropertiesMdeIntegrationOutput {
	return o
}

func (o DefenderSettingsPropertiesMdeIntegrationOutput) ToDefenderSettingsPropertiesMdeIntegrationOutputWithContext(ctx context.Context) DefenderSettingsPropertiesMdeIntegrationOutput {
	return o
}

func (o DefenderSettingsPropertiesMdeIntegrationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderSettingsPropertiesMdeIntegration) string { return v.Status }).(pulumi.StringOutput)
}

type DefenderSettingsPropertiesResponseMdeIntegration struct {
	Status string `pulumi:"status"`
}

type DefenderSettingsPropertiesResponseMdeIntegrationOutput struct{ *pulumi.OutputState }

func (DefenderSettingsPropertiesResponseMdeIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderSettingsPropertiesResponseMdeIntegration)(nil)).Elem()
}

func (o DefenderSettingsPropertiesResponseMdeIntegrationOutput) ToDefenderSettingsPropertiesResponseMdeIntegrationOutput() DefenderSettingsPropertiesResponseMdeIntegrationOutput {
	return o
}

func (o DefenderSettingsPropertiesResponseMdeIntegrationOutput) ToDefenderSettingsPropertiesResponseMdeIntegrationOutputWithContext(ctx context.Context) DefenderSettingsPropertiesResponseMdeIntegrationOutput {
	return o
}

func (o DefenderSettingsPropertiesResponseMdeIntegrationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderSettingsPropertiesResponseMdeIntegration) string { return v.Status }).(pulumi.StringOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DefenderSettingsPropertiesMdeIntegrationOutput{})
	pulumi.RegisterOutputType(DefenderSettingsPropertiesResponseMdeIntegrationOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
