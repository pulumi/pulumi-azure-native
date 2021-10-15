


package iotsecurity

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

func (i DefenderSettingsPropertiesMdeIntegrationArgs) ToDefenderSettingsPropertiesMdeIntegrationPtrOutput() DefenderSettingsPropertiesMdeIntegrationPtrOutput {
	return i.ToDefenderSettingsPropertiesMdeIntegrationPtrOutputWithContext(context.Background())
}

func (i DefenderSettingsPropertiesMdeIntegrationArgs) ToDefenderSettingsPropertiesMdeIntegrationPtrOutputWithContext(ctx context.Context) DefenderSettingsPropertiesMdeIntegrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderSettingsPropertiesMdeIntegrationOutput).ToDefenderSettingsPropertiesMdeIntegrationPtrOutputWithContext(ctx)
}









type DefenderSettingsPropertiesMdeIntegrationPtrInput interface {
	pulumi.Input

	ToDefenderSettingsPropertiesMdeIntegrationPtrOutput() DefenderSettingsPropertiesMdeIntegrationPtrOutput
	ToDefenderSettingsPropertiesMdeIntegrationPtrOutputWithContext(context.Context) DefenderSettingsPropertiesMdeIntegrationPtrOutput
}

type defenderSettingsPropertiesMdeIntegrationPtrType DefenderSettingsPropertiesMdeIntegrationArgs

func DefenderSettingsPropertiesMdeIntegrationPtr(v *DefenderSettingsPropertiesMdeIntegrationArgs) DefenderSettingsPropertiesMdeIntegrationPtrInput {
	return (*defenderSettingsPropertiesMdeIntegrationPtrType)(v)
}

func (*defenderSettingsPropertiesMdeIntegrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderSettingsPropertiesMdeIntegration)(nil)).Elem()
}

func (i *defenderSettingsPropertiesMdeIntegrationPtrType) ToDefenderSettingsPropertiesMdeIntegrationPtrOutput() DefenderSettingsPropertiesMdeIntegrationPtrOutput {
	return i.ToDefenderSettingsPropertiesMdeIntegrationPtrOutputWithContext(context.Background())
}

func (i *defenderSettingsPropertiesMdeIntegrationPtrType) ToDefenderSettingsPropertiesMdeIntegrationPtrOutputWithContext(ctx context.Context) DefenderSettingsPropertiesMdeIntegrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderSettingsPropertiesMdeIntegrationPtrOutput)
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

func (o DefenderSettingsPropertiesMdeIntegrationOutput) ToDefenderSettingsPropertiesMdeIntegrationPtrOutput() DefenderSettingsPropertiesMdeIntegrationPtrOutput {
	return o.ToDefenderSettingsPropertiesMdeIntegrationPtrOutputWithContext(context.Background())
}

func (o DefenderSettingsPropertiesMdeIntegrationOutput) ToDefenderSettingsPropertiesMdeIntegrationPtrOutputWithContext(ctx context.Context) DefenderSettingsPropertiesMdeIntegrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderSettingsPropertiesMdeIntegration) *DefenderSettingsPropertiesMdeIntegration {
		return &v
	}).(DefenderSettingsPropertiesMdeIntegrationPtrOutput)
}

func (o DefenderSettingsPropertiesMdeIntegrationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderSettingsPropertiesMdeIntegration) string { return v.Status }).(pulumi.StringOutput)
}

type DefenderSettingsPropertiesMdeIntegrationPtrOutput struct{ *pulumi.OutputState }

func (DefenderSettingsPropertiesMdeIntegrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderSettingsPropertiesMdeIntegration)(nil)).Elem()
}

func (o DefenderSettingsPropertiesMdeIntegrationPtrOutput) ToDefenderSettingsPropertiesMdeIntegrationPtrOutput() DefenderSettingsPropertiesMdeIntegrationPtrOutput {
	return o
}

func (o DefenderSettingsPropertiesMdeIntegrationPtrOutput) ToDefenderSettingsPropertiesMdeIntegrationPtrOutputWithContext(ctx context.Context) DefenderSettingsPropertiesMdeIntegrationPtrOutput {
	return o
}

func (o DefenderSettingsPropertiesMdeIntegrationPtrOutput) Elem() DefenderSettingsPropertiesMdeIntegrationOutput {
	return o.ApplyT(func(v *DefenderSettingsPropertiesMdeIntegration) DefenderSettingsPropertiesMdeIntegration {
		if v != nil {
			return *v
		}
		var ret DefenderSettingsPropertiesMdeIntegration
		return ret
	}).(DefenderSettingsPropertiesMdeIntegrationOutput)
}

func (o DefenderSettingsPropertiesMdeIntegrationPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderSettingsPropertiesMdeIntegration) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type DefenderSettingsPropertiesResponseMdeIntegration struct {
	Status string `pulumi:"status"`
}





type DefenderSettingsPropertiesResponseMdeIntegrationInput interface {
	pulumi.Input

	ToDefenderSettingsPropertiesResponseMdeIntegrationOutput() DefenderSettingsPropertiesResponseMdeIntegrationOutput
	ToDefenderSettingsPropertiesResponseMdeIntegrationOutputWithContext(context.Context) DefenderSettingsPropertiesResponseMdeIntegrationOutput
}

type DefenderSettingsPropertiesResponseMdeIntegrationArgs struct {
	Status pulumi.StringInput `pulumi:"status"`
}

func (DefenderSettingsPropertiesResponseMdeIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderSettingsPropertiesResponseMdeIntegration)(nil)).Elem()
}

func (i DefenderSettingsPropertiesResponseMdeIntegrationArgs) ToDefenderSettingsPropertiesResponseMdeIntegrationOutput() DefenderSettingsPropertiesResponseMdeIntegrationOutput {
	return i.ToDefenderSettingsPropertiesResponseMdeIntegrationOutputWithContext(context.Background())
}

func (i DefenderSettingsPropertiesResponseMdeIntegrationArgs) ToDefenderSettingsPropertiesResponseMdeIntegrationOutputWithContext(ctx context.Context) DefenderSettingsPropertiesResponseMdeIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderSettingsPropertiesResponseMdeIntegrationOutput)
}

func (i DefenderSettingsPropertiesResponseMdeIntegrationArgs) ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutput() DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput {
	return i.ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutputWithContext(context.Background())
}

func (i DefenderSettingsPropertiesResponseMdeIntegrationArgs) ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutputWithContext(ctx context.Context) DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderSettingsPropertiesResponseMdeIntegrationOutput).ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutputWithContext(ctx)
}









type DefenderSettingsPropertiesResponseMdeIntegrationPtrInput interface {
	pulumi.Input

	ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutput() DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput
	ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutputWithContext(context.Context) DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput
}

type defenderSettingsPropertiesResponseMdeIntegrationPtrType DefenderSettingsPropertiesResponseMdeIntegrationArgs

func DefenderSettingsPropertiesResponseMdeIntegrationPtr(v *DefenderSettingsPropertiesResponseMdeIntegrationArgs) DefenderSettingsPropertiesResponseMdeIntegrationPtrInput {
	return (*defenderSettingsPropertiesResponseMdeIntegrationPtrType)(v)
}

func (*defenderSettingsPropertiesResponseMdeIntegrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderSettingsPropertiesResponseMdeIntegration)(nil)).Elem()
}

func (i *defenderSettingsPropertiesResponseMdeIntegrationPtrType) ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutput() DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput {
	return i.ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutputWithContext(context.Background())
}

func (i *defenderSettingsPropertiesResponseMdeIntegrationPtrType) ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutputWithContext(ctx context.Context) DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput)
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

func (o DefenderSettingsPropertiesResponseMdeIntegrationOutput) ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutput() DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput {
	return o.ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutputWithContext(context.Background())
}

func (o DefenderSettingsPropertiesResponseMdeIntegrationOutput) ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutputWithContext(ctx context.Context) DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderSettingsPropertiesResponseMdeIntegration) *DefenderSettingsPropertiesResponseMdeIntegration {
		return &v
	}).(DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput)
}

func (o DefenderSettingsPropertiesResponseMdeIntegrationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderSettingsPropertiesResponseMdeIntegration) string { return v.Status }).(pulumi.StringOutput)
}

type DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput struct{ *pulumi.OutputState }

func (DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderSettingsPropertiesResponseMdeIntegration)(nil)).Elem()
}

func (o DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput) ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutput() DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput {
	return o
}

func (o DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput) ToDefenderSettingsPropertiesResponseMdeIntegrationPtrOutputWithContext(ctx context.Context) DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput {
	return o
}

func (o DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput) Elem() DefenderSettingsPropertiesResponseMdeIntegrationOutput {
	return o.ApplyT(func(v *DefenderSettingsPropertiesResponseMdeIntegration) DefenderSettingsPropertiesResponseMdeIntegration {
		if v != nil {
			return *v
		}
		var ret DefenderSettingsPropertiesResponseMdeIntegration
		return ret
	}).(DefenderSettingsPropertiesResponseMdeIntegrationOutput)
}

func (o DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderSettingsPropertiesResponseMdeIntegration) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderSettingsPropertiesMdeIntegrationInput)(nil)).Elem(), DefenderSettingsPropertiesMdeIntegrationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderSettingsPropertiesMdeIntegrationPtrInput)(nil)).Elem(), DefenderSettingsPropertiesMdeIntegrationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderSettingsPropertiesResponseMdeIntegrationInput)(nil)).Elem(), DefenderSettingsPropertiesResponseMdeIntegrationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderSettingsPropertiesResponseMdeIntegrationPtrInput)(nil)).Elem(), DefenderSettingsPropertiesResponseMdeIntegrationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterOutputType(DefenderSettingsPropertiesMdeIntegrationOutput{})
	pulumi.RegisterOutputType(DefenderSettingsPropertiesMdeIntegrationPtrOutput{})
	pulumi.RegisterOutputType(DefenderSettingsPropertiesResponseMdeIntegrationOutput{})
	pulumi.RegisterOutputType(DefenderSettingsPropertiesResponseMdeIntegrationPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
