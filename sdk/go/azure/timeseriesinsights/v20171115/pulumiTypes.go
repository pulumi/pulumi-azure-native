


package v20171115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvironmentStateDetailsResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}





type EnvironmentStateDetailsResponseInput interface {
	pulumi.Input

	ToEnvironmentStateDetailsResponseOutput() EnvironmentStateDetailsResponseOutput
	ToEnvironmentStateDetailsResponseOutputWithContext(context.Context) EnvironmentStateDetailsResponseOutput
}

type EnvironmentStateDetailsResponseArgs struct {
	Code    pulumi.StringPtrInput `pulumi:"code"`
	Message pulumi.StringPtrInput `pulumi:"message"`
}

func (EnvironmentStateDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentStateDetailsResponse)(nil)).Elem()
}

func (i EnvironmentStateDetailsResponseArgs) ToEnvironmentStateDetailsResponseOutput() EnvironmentStateDetailsResponseOutput {
	return i.ToEnvironmentStateDetailsResponseOutputWithContext(context.Background())
}

func (i EnvironmentStateDetailsResponseArgs) ToEnvironmentStateDetailsResponseOutputWithContext(ctx context.Context) EnvironmentStateDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentStateDetailsResponseOutput)
}

func (i EnvironmentStateDetailsResponseArgs) ToEnvironmentStateDetailsResponsePtrOutput() EnvironmentStateDetailsResponsePtrOutput {
	return i.ToEnvironmentStateDetailsResponsePtrOutputWithContext(context.Background())
}

func (i EnvironmentStateDetailsResponseArgs) ToEnvironmentStateDetailsResponsePtrOutputWithContext(ctx context.Context) EnvironmentStateDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentStateDetailsResponseOutput).ToEnvironmentStateDetailsResponsePtrOutputWithContext(ctx)
}









type EnvironmentStateDetailsResponsePtrInput interface {
	pulumi.Input

	ToEnvironmentStateDetailsResponsePtrOutput() EnvironmentStateDetailsResponsePtrOutput
	ToEnvironmentStateDetailsResponsePtrOutputWithContext(context.Context) EnvironmentStateDetailsResponsePtrOutput
}

type environmentStateDetailsResponsePtrType EnvironmentStateDetailsResponseArgs

func EnvironmentStateDetailsResponsePtr(v *EnvironmentStateDetailsResponseArgs) EnvironmentStateDetailsResponsePtrInput {
	return (*environmentStateDetailsResponsePtrType)(v)
}

func (*environmentStateDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentStateDetailsResponse)(nil)).Elem()
}

func (i *environmentStateDetailsResponsePtrType) ToEnvironmentStateDetailsResponsePtrOutput() EnvironmentStateDetailsResponsePtrOutput {
	return i.ToEnvironmentStateDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *environmentStateDetailsResponsePtrType) ToEnvironmentStateDetailsResponsePtrOutputWithContext(ctx context.Context) EnvironmentStateDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentStateDetailsResponsePtrOutput)
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

func (o EnvironmentStateDetailsResponseOutput) ToEnvironmentStateDetailsResponsePtrOutput() EnvironmentStateDetailsResponsePtrOutput {
	return o.ToEnvironmentStateDetailsResponsePtrOutputWithContext(context.Background())
}

func (o EnvironmentStateDetailsResponseOutput) ToEnvironmentStateDetailsResponsePtrOutputWithContext(ctx context.Context) EnvironmentStateDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentStateDetailsResponse) *EnvironmentStateDetailsResponse {
		return &v
	}).(EnvironmentStateDetailsResponsePtrOutput)
}

func (o EnvironmentStateDetailsResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentStateDetailsResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o EnvironmentStateDetailsResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentStateDetailsResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type EnvironmentStateDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (EnvironmentStateDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentStateDetailsResponse)(nil)).Elem()
}

func (o EnvironmentStateDetailsResponsePtrOutput) ToEnvironmentStateDetailsResponsePtrOutput() EnvironmentStateDetailsResponsePtrOutput {
	return o
}

func (o EnvironmentStateDetailsResponsePtrOutput) ToEnvironmentStateDetailsResponsePtrOutputWithContext(ctx context.Context) EnvironmentStateDetailsResponsePtrOutput {
	return o
}

func (o EnvironmentStateDetailsResponsePtrOutput) Elem() EnvironmentStateDetailsResponseOutput {
	return o.ApplyT(func(v *EnvironmentStateDetailsResponse) EnvironmentStateDetailsResponse {
		if v != nil {
			return *v
		}
		var ret EnvironmentStateDetailsResponse
		return ret
	}).(EnvironmentStateDetailsResponseOutput)
}

func (o EnvironmentStateDetailsResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentStateDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o EnvironmentStateDetailsResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentStateDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type EnvironmentStatusResponse struct {
	Ingress IngressEnvironmentStatusResponse `pulumi:"ingress"`
}





type EnvironmentStatusResponseInput interface {
	pulumi.Input

	ToEnvironmentStatusResponseOutput() EnvironmentStatusResponseOutput
	ToEnvironmentStatusResponseOutputWithContext(context.Context) EnvironmentStatusResponseOutput
}

type EnvironmentStatusResponseArgs struct {
	Ingress IngressEnvironmentStatusResponseInput `pulumi:"ingress"`
}

func (EnvironmentStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentStatusResponse)(nil)).Elem()
}

func (i EnvironmentStatusResponseArgs) ToEnvironmentStatusResponseOutput() EnvironmentStatusResponseOutput {
	return i.ToEnvironmentStatusResponseOutputWithContext(context.Background())
}

func (i EnvironmentStatusResponseArgs) ToEnvironmentStatusResponseOutputWithContext(ctx context.Context) EnvironmentStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentStatusResponseOutput)
}

func (i EnvironmentStatusResponseArgs) ToEnvironmentStatusResponsePtrOutput() EnvironmentStatusResponsePtrOutput {
	return i.ToEnvironmentStatusResponsePtrOutputWithContext(context.Background())
}

func (i EnvironmentStatusResponseArgs) ToEnvironmentStatusResponsePtrOutputWithContext(ctx context.Context) EnvironmentStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentStatusResponseOutput).ToEnvironmentStatusResponsePtrOutputWithContext(ctx)
}









type EnvironmentStatusResponsePtrInput interface {
	pulumi.Input

	ToEnvironmentStatusResponsePtrOutput() EnvironmentStatusResponsePtrOutput
	ToEnvironmentStatusResponsePtrOutputWithContext(context.Context) EnvironmentStatusResponsePtrOutput
}

type environmentStatusResponsePtrType EnvironmentStatusResponseArgs

func EnvironmentStatusResponsePtr(v *EnvironmentStatusResponseArgs) EnvironmentStatusResponsePtrInput {
	return (*environmentStatusResponsePtrType)(v)
}

func (*environmentStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentStatusResponse)(nil)).Elem()
}

func (i *environmentStatusResponsePtrType) ToEnvironmentStatusResponsePtrOutput() EnvironmentStatusResponsePtrOutput {
	return i.ToEnvironmentStatusResponsePtrOutputWithContext(context.Background())
}

func (i *environmentStatusResponsePtrType) ToEnvironmentStatusResponsePtrOutputWithContext(ctx context.Context) EnvironmentStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentStatusResponsePtrOutput)
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

func (o EnvironmentStatusResponseOutput) ToEnvironmentStatusResponsePtrOutput() EnvironmentStatusResponsePtrOutput {
	return o.ToEnvironmentStatusResponsePtrOutputWithContext(context.Background())
}

func (o EnvironmentStatusResponseOutput) ToEnvironmentStatusResponsePtrOutputWithContext(ctx context.Context) EnvironmentStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentStatusResponse) *EnvironmentStatusResponse {
		return &v
	}).(EnvironmentStatusResponsePtrOutput)
}

func (o EnvironmentStatusResponseOutput) Ingress() IngressEnvironmentStatusResponseOutput {
	return o.ApplyT(func(v EnvironmentStatusResponse) IngressEnvironmentStatusResponse { return v.Ingress }).(IngressEnvironmentStatusResponseOutput)
}

type EnvironmentStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (EnvironmentStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentStatusResponse)(nil)).Elem()
}

func (o EnvironmentStatusResponsePtrOutput) ToEnvironmentStatusResponsePtrOutput() EnvironmentStatusResponsePtrOutput {
	return o
}

func (o EnvironmentStatusResponsePtrOutput) ToEnvironmentStatusResponsePtrOutputWithContext(ctx context.Context) EnvironmentStatusResponsePtrOutput {
	return o
}

func (o EnvironmentStatusResponsePtrOutput) Elem() EnvironmentStatusResponseOutput {
	return o.ApplyT(func(v *EnvironmentStatusResponse) EnvironmentStatusResponse {
		if v != nil {
			return *v
		}
		var ret EnvironmentStatusResponse
		return ret
	}).(EnvironmentStatusResponseOutput)
}

func (o EnvironmentStatusResponsePtrOutput) Ingress() IngressEnvironmentStatusResponsePtrOutput {
	return o.ApplyT(func(v *EnvironmentStatusResponse) *IngressEnvironmentStatusResponse {
		if v == nil {
			return nil
		}
		return &v.Ingress
	}).(IngressEnvironmentStatusResponsePtrOutput)
}

type IngressEnvironmentStatusResponse struct {
	State        *string                         `pulumi:"state"`
	StateDetails EnvironmentStateDetailsResponse `pulumi:"stateDetails"`
}





type IngressEnvironmentStatusResponseInput interface {
	pulumi.Input

	ToIngressEnvironmentStatusResponseOutput() IngressEnvironmentStatusResponseOutput
	ToIngressEnvironmentStatusResponseOutputWithContext(context.Context) IngressEnvironmentStatusResponseOutput
}

type IngressEnvironmentStatusResponseArgs struct {
	State        pulumi.StringPtrInput                `pulumi:"state"`
	StateDetails EnvironmentStateDetailsResponseInput `pulumi:"stateDetails"`
}

func (IngressEnvironmentStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressEnvironmentStatusResponse)(nil)).Elem()
}

func (i IngressEnvironmentStatusResponseArgs) ToIngressEnvironmentStatusResponseOutput() IngressEnvironmentStatusResponseOutput {
	return i.ToIngressEnvironmentStatusResponseOutputWithContext(context.Background())
}

func (i IngressEnvironmentStatusResponseArgs) ToIngressEnvironmentStatusResponseOutputWithContext(ctx context.Context) IngressEnvironmentStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressEnvironmentStatusResponseOutput)
}

func (i IngressEnvironmentStatusResponseArgs) ToIngressEnvironmentStatusResponsePtrOutput() IngressEnvironmentStatusResponsePtrOutput {
	return i.ToIngressEnvironmentStatusResponsePtrOutputWithContext(context.Background())
}

func (i IngressEnvironmentStatusResponseArgs) ToIngressEnvironmentStatusResponsePtrOutputWithContext(ctx context.Context) IngressEnvironmentStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressEnvironmentStatusResponseOutput).ToIngressEnvironmentStatusResponsePtrOutputWithContext(ctx)
}









type IngressEnvironmentStatusResponsePtrInput interface {
	pulumi.Input

	ToIngressEnvironmentStatusResponsePtrOutput() IngressEnvironmentStatusResponsePtrOutput
	ToIngressEnvironmentStatusResponsePtrOutputWithContext(context.Context) IngressEnvironmentStatusResponsePtrOutput
}

type ingressEnvironmentStatusResponsePtrType IngressEnvironmentStatusResponseArgs

func IngressEnvironmentStatusResponsePtr(v *IngressEnvironmentStatusResponseArgs) IngressEnvironmentStatusResponsePtrInput {
	return (*ingressEnvironmentStatusResponsePtrType)(v)
}

func (*ingressEnvironmentStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressEnvironmentStatusResponse)(nil)).Elem()
}

func (i *ingressEnvironmentStatusResponsePtrType) ToIngressEnvironmentStatusResponsePtrOutput() IngressEnvironmentStatusResponsePtrOutput {
	return i.ToIngressEnvironmentStatusResponsePtrOutputWithContext(context.Background())
}

func (i *ingressEnvironmentStatusResponsePtrType) ToIngressEnvironmentStatusResponsePtrOutputWithContext(ctx context.Context) IngressEnvironmentStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressEnvironmentStatusResponsePtrOutput)
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

func (o IngressEnvironmentStatusResponseOutput) ToIngressEnvironmentStatusResponsePtrOutput() IngressEnvironmentStatusResponsePtrOutput {
	return o.ToIngressEnvironmentStatusResponsePtrOutputWithContext(context.Background())
}

func (o IngressEnvironmentStatusResponseOutput) ToIngressEnvironmentStatusResponsePtrOutputWithContext(ctx context.Context) IngressEnvironmentStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IngressEnvironmentStatusResponse) *IngressEnvironmentStatusResponse {
		return &v
	}).(IngressEnvironmentStatusResponsePtrOutput)
}

func (o IngressEnvironmentStatusResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngressEnvironmentStatusResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o IngressEnvironmentStatusResponseOutput) StateDetails() EnvironmentStateDetailsResponseOutput {
	return o.ApplyT(func(v IngressEnvironmentStatusResponse) EnvironmentStateDetailsResponse { return v.StateDetails }).(EnvironmentStateDetailsResponseOutput)
}

type IngressEnvironmentStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (IngressEnvironmentStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressEnvironmentStatusResponse)(nil)).Elem()
}

func (o IngressEnvironmentStatusResponsePtrOutput) ToIngressEnvironmentStatusResponsePtrOutput() IngressEnvironmentStatusResponsePtrOutput {
	return o
}

func (o IngressEnvironmentStatusResponsePtrOutput) ToIngressEnvironmentStatusResponsePtrOutputWithContext(ctx context.Context) IngressEnvironmentStatusResponsePtrOutput {
	return o
}

func (o IngressEnvironmentStatusResponsePtrOutput) Elem() IngressEnvironmentStatusResponseOutput {
	return o.ApplyT(func(v *IngressEnvironmentStatusResponse) IngressEnvironmentStatusResponse {
		if v != nil {
			return *v
		}
		var ret IngressEnvironmentStatusResponse
		return ret
	}).(IngressEnvironmentStatusResponseOutput)
}

func (o IngressEnvironmentStatusResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngressEnvironmentStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

func (o IngressEnvironmentStatusResponsePtrOutput) StateDetails() EnvironmentStateDetailsResponsePtrOutput {
	return o.ApplyT(func(v *IngressEnvironmentStatusResponse) *EnvironmentStateDetailsResponse {
		if v == nil {
			return nil
		}
		return &v.StateDetails
	}).(EnvironmentStateDetailsResponsePtrOutput)
}

type PartitionKeyProperty struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type PartitionKeyPropertyInput interface {
	pulumi.Input

	ToPartitionKeyPropertyOutput() PartitionKeyPropertyOutput
	ToPartitionKeyPropertyOutputWithContext(context.Context) PartitionKeyPropertyOutput
}

type PartitionKeyPropertyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PartitionKeyPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionKeyProperty)(nil)).Elem()
}

func (i PartitionKeyPropertyArgs) ToPartitionKeyPropertyOutput() PartitionKeyPropertyOutput {
	return i.ToPartitionKeyPropertyOutputWithContext(context.Background())
}

func (i PartitionKeyPropertyArgs) ToPartitionKeyPropertyOutputWithContext(ctx context.Context) PartitionKeyPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionKeyPropertyOutput)
}





type PartitionKeyPropertyArrayInput interface {
	pulumi.Input

	ToPartitionKeyPropertyArrayOutput() PartitionKeyPropertyArrayOutput
	ToPartitionKeyPropertyArrayOutputWithContext(context.Context) PartitionKeyPropertyArrayOutput
}

type PartitionKeyPropertyArray []PartitionKeyPropertyInput

func (PartitionKeyPropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartitionKeyProperty)(nil)).Elem()
}

func (i PartitionKeyPropertyArray) ToPartitionKeyPropertyArrayOutput() PartitionKeyPropertyArrayOutput {
	return i.ToPartitionKeyPropertyArrayOutputWithContext(context.Background())
}

func (i PartitionKeyPropertyArray) ToPartitionKeyPropertyArrayOutputWithContext(ctx context.Context) PartitionKeyPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionKeyPropertyArrayOutput)
}

type PartitionKeyPropertyOutput struct{ *pulumi.OutputState }

func (PartitionKeyPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionKeyProperty)(nil)).Elem()
}

func (o PartitionKeyPropertyOutput) ToPartitionKeyPropertyOutput() PartitionKeyPropertyOutput {
	return o
}

func (o PartitionKeyPropertyOutput) ToPartitionKeyPropertyOutputWithContext(ctx context.Context) PartitionKeyPropertyOutput {
	return o
}

func (o PartitionKeyPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartitionKeyProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PartitionKeyPropertyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartitionKeyProperty) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PartitionKeyPropertyArrayOutput struct{ *pulumi.OutputState }

func (PartitionKeyPropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartitionKeyProperty)(nil)).Elem()
}

func (o PartitionKeyPropertyArrayOutput) ToPartitionKeyPropertyArrayOutput() PartitionKeyPropertyArrayOutput {
	return o
}

func (o PartitionKeyPropertyArrayOutput) ToPartitionKeyPropertyArrayOutputWithContext(ctx context.Context) PartitionKeyPropertyArrayOutput {
	return o
}

func (o PartitionKeyPropertyArrayOutput) Index(i pulumi.IntInput) PartitionKeyPropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PartitionKeyProperty {
		return vs[0].([]PartitionKeyProperty)[vs[1].(int)]
	}).(PartitionKeyPropertyOutput)
}

type PartitionKeyPropertyResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type PartitionKeyPropertyResponseInput interface {
	pulumi.Input

	ToPartitionKeyPropertyResponseOutput() PartitionKeyPropertyResponseOutput
	ToPartitionKeyPropertyResponseOutputWithContext(context.Context) PartitionKeyPropertyResponseOutput
}

type PartitionKeyPropertyResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PartitionKeyPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionKeyPropertyResponse)(nil)).Elem()
}

func (i PartitionKeyPropertyResponseArgs) ToPartitionKeyPropertyResponseOutput() PartitionKeyPropertyResponseOutput {
	return i.ToPartitionKeyPropertyResponseOutputWithContext(context.Background())
}

func (i PartitionKeyPropertyResponseArgs) ToPartitionKeyPropertyResponseOutputWithContext(ctx context.Context) PartitionKeyPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionKeyPropertyResponseOutput)
}





type PartitionKeyPropertyResponseArrayInput interface {
	pulumi.Input

	ToPartitionKeyPropertyResponseArrayOutput() PartitionKeyPropertyResponseArrayOutput
	ToPartitionKeyPropertyResponseArrayOutputWithContext(context.Context) PartitionKeyPropertyResponseArrayOutput
}

type PartitionKeyPropertyResponseArray []PartitionKeyPropertyResponseInput

func (PartitionKeyPropertyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartitionKeyPropertyResponse)(nil)).Elem()
}

func (i PartitionKeyPropertyResponseArray) ToPartitionKeyPropertyResponseArrayOutput() PartitionKeyPropertyResponseArrayOutput {
	return i.ToPartitionKeyPropertyResponseArrayOutputWithContext(context.Background())
}

func (i PartitionKeyPropertyResponseArray) ToPartitionKeyPropertyResponseArrayOutputWithContext(ctx context.Context) PartitionKeyPropertyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionKeyPropertyResponseArrayOutput)
}

type PartitionKeyPropertyResponseOutput struct{ *pulumi.OutputState }

func (PartitionKeyPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionKeyPropertyResponse)(nil)).Elem()
}

func (o PartitionKeyPropertyResponseOutput) ToPartitionKeyPropertyResponseOutput() PartitionKeyPropertyResponseOutput {
	return o
}

func (o PartitionKeyPropertyResponseOutput) ToPartitionKeyPropertyResponseOutputWithContext(ctx context.Context) PartitionKeyPropertyResponseOutput {
	return o
}

func (o PartitionKeyPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartitionKeyPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PartitionKeyPropertyResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartitionKeyPropertyResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PartitionKeyPropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (PartitionKeyPropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartitionKeyPropertyResponse)(nil)).Elem()
}

func (o PartitionKeyPropertyResponseArrayOutput) ToPartitionKeyPropertyResponseArrayOutput() PartitionKeyPropertyResponseArrayOutput {
	return o
}

func (o PartitionKeyPropertyResponseArrayOutput) ToPartitionKeyPropertyResponseArrayOutputWithContext(ctx context.Context) PartitionKeyPropertyResponseArrayOutput {
	return o
}

func (o PartitionKeyPropertyResponseArrayOutput) Index(i pulumi.IntInput) PartitionKeyPropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PartitionKeyPropertyResponse {
		return vs[0].([]PartitionKeyPropertyResponse)[vs[1].(int)]
	}).(PartitionKeyPropertyResponseOutput)
}

type ReferenceDataSetKeyProperty struct {
	Name *string                       `pulumi:"name"`
	Type *ReferenceDataKeyPropertyType `pulumi:"type"`
}





type ReferenceDataSetKeyPropertyInput interface {
	pulumi.Input

	ToReferenceDataSetKeyPropertyOutput() ReferenceDataSetKeyPropertyOutput
	ToReferenceDataSetKeyPropertyOutputWithContext(context.Context) ReferenceDataSetKeyPropertyOutput
}

type ReferenceDataSetKeyPropertyArgs struct {
	Name pulumi.StringPtrInput                `pulumi:"name"`
	Type ReferenceDataKeyPropertyTypePtrInput `pulumi:"type"`
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

func (o ReferenceDataSetKeyPropertyOutput) Type() ReferenceDataKeyPropertyTypePtrOutput {
	return o.ApplyT(func(v ReferenceDataSetKeyProperty) *ReferenceDataKeyPropertyType { return v.Type }).(ReferenceDataKeyPropertyTypePtrOutput)
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





type ReferenceDataSetKeyPropertyResponseInput interface {
	pulumi.Input

	ToReferenceDataSetKeyPropertyResponseOutput() ReferenceDataSetKeyPropertyResponseOutput
	ToReferenceDataSetKeyPropertyResponseOutputWithContext(context.Context) ReferenceDataSetKeyPropertyResponseOutput
}

type ReferenceDataSetKeyPropertyResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ReferenceDataSetKeyPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceDataSetKeyPropertyResponse)(nil)).Elem()
}

func (i ReferenceDataSetKeyPropertyResponseArgs) ToReferenceDataSetKeyPropertyResponseOutput() ReferenceDataSetKeyPropertyResponseOutput {
	return i.ToReferenceDataSetKeyPropertyResponseOutputWithContext(context.Background())
}

func (i ReferenceDataSetKeyPropertyResponseArgs) ToReferenceDataSetKeyPropertyResponseOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceDataSetKeyPropertyResponseOutput)
}





type ReferenceDataSetKeyPropertyResponseArrayInput interface {
	pulumi.Input

	ToReferenceDataSetKeyPropertyResponseArrayOutput() ReferenceDataSetKeyPropertyResponseArrayOutput
	ToReferenceDataSetKeyPropertyResponseArrayOutputWithContext(context.Context) ReferenceDataSetKeyPropertyResponseArrayOutput
}

type ReferenceDataSetKeyPropertyResponseArray []ReferenceDataSetKeyPropertyResponseInput

func (ReferenceDataSetKeyPropertyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReferenceDataSetKeyPropertyResponse)(nil)).Elem()
}

func (i ReferenceDataSetKeyPropertyResponseArray) ToReferenceDataSetKeyPropertyResponseArrayOutput() ReferenceDataSetKeyPropertyResponseArrayOutput {
	return i.ToReferenceDataSetKeyPropertyResponseArrayOutputWithContext(context.Background())
}

func (i ReferenceDataSetKeyPropertyResponseArray) ToReferenceDataSetKeyPropertyResponseArrayOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceDataSetKeyPropertyResponseArrayOutput)
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
	Capacity int     `pulumi:"capacity"`
	Name     SkuName `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntInput `pulumi:"capacity"`
	Name     SkuNameInput    `pulumi:"name"`
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

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
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

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v Sku) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o SkuOutput) Name() SkuNameOutput {
	return o.ApplyT(func(v Sku) SkuName { return v.Name }).(SkuNameOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Name() SkuNamePtrOutput {
	return o.ApplyT(func(v *Sku) *SkuName {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(SkuNamePtrOutput)
}

type SkuResponse struct {
	Capacity int    `pulumi:"capacity"`
	Name     string `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntInput    `pulumi:"capacity"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
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

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v SkuResponse) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentStateDetailsResponseInput)(nil)).Elem(), EnvironmentStateDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentStateDetailsResponsePtrInput)(nil)).Elem(), EnvironmentStateDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentStatusResponseInput)(nil)).Elem(), EnvironmentStatusResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentStatusResponsePtrInput)(nil)).Elem(), EnvironmentStatusResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngressEnvironmentStatusResponseInput)(nil)).Elem(), IngressEnvironmentStatusResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngressEnvironmentStatusResponsePtrInput)(nil)).Elem(), IngressEnvironmentStatusResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionKeyPropertyInput)(nil)).Elem(), PartitionKeyPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionKeyPropertyArrayInput)(nil)).Elem(), PartitionKeyPropertyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionKeyPropertyResponseInput)(nil)).Elem(), PartitionKeyPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionKeyPropertyResponseArrayInput)(nil)).Elem(), PartitionKeyPropertyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceDataSetKeyPropertyInput)(nil)).Elem(), ReferenceDataSetKeyPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceDataSetKeyPropertyArrayInput)(nil)).Elem(), ReferenceDataSetKeyPropertyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceDataSetKeyPropertyResponseInput)(nil)).Elem(), ReferenceDataSetKeyPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceDataSetKeyPropertyResponseArrayInput)(nil)).Elem(), ReferenceDataSetKeyPropertyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterOutputType(EnvironmentStateDetailsResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentStateDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentStatusResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(IngressEnvironmentStatusResponseOutput{})
	pulumi.RegisterOutputType(IngressEnvironmentStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(PartitionKeyPropertyOutput{})
	pulumi.RegisterOutputType(PartitionKeyPropertyArrayOutput{})
	pulumi.RegisterOutputType(PartitionKeyPropertyResponseOutput{})
	pulumi.RegisterOutputType(PartitionKeyPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyArrayOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyResponseOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
