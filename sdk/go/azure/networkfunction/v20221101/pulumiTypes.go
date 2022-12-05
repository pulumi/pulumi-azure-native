


package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EmissionPoliciesPropertiesFormat struct {
	EmissionDestinations []EmissionPolicyDestination `pulumi:"emissionDestinations"`
	EmissionType         *string                     `pulumi:"emissionType"`
}





type EmissionPoliciesPropertiesFormatInput interface {
	pulumi.Input

	ToEmissionPoliciesPropertiesFormatOutput() EmissionPoliciesPropertiesFormatOutput
	ToEmissionPoliciesPropertiesFormatOutputWithContext(context.Context) EmissionPoliciesPropertiesFormatOutput
}

type EmissionPoliciesPropertiesFormatArgs struct {
	EmissionDestinations EmissionPolicyDestinationArrayInput `pulumi:"emissionDestinations"`
	EmissionType         pulumi.StringPtrInput               `pulumi:"emissionType"`
}

func (EmissionPoliciesPropertiesFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmissionPoliciesPropertiesFormat)(nil)).Elem()
}

func (i EmissionPoliciesPropertiesFormatArgs) ToEmissionPoliciesPropertiesFormatOutput() EmissionPoliciesPropertiesFormatOutput {
	return i.ToEmissionPoliciesPropertiesFormatOutputWithContext(context.Background())
}

func (i EmissionPoliciesPropertiesFormatArgs) ToEmissionPoliciesPropertiesFormatOutputWithContext(ctx context.Context) EmissionPoliciesPropertiesFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmissionPoliciesPropertiesFormatOutput)
}





type EmissionPoliciesPropertiesFormatArrayInput interface {
	pulumi.Input

	ToEmissionPoliciesPropertiesFormatArrayOutput() EmissionPoliciesPropertiesFormatArrayOutput
	ToEmissionPoliciesPropertiesFormatArrayOutputWithContext(context.Context) EmissionPoliciesPropertiesFormatArrayOutput
}

type EmissionPoliciesPropertiesFormatArray []EmissionPoliciesPropertiesFormatInput

func (EmissionPoliciesPropertiesFormatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmissionPoliciesPropertiesFormat)(nil)).Elem()
}

func (i EmissionPoliciesPropertiesFormatArray) ToEmissionPoliciesPropertiesFormatArrayOutput() EmissionPoliciesPropertiesFormatArrayOutput {
	return i.ToEmissionPoliciesPropertiesFormatArrayOutputWithContext(context.Background())
}

func (i EmissionPoliciesPropertiesFormatArray) ToEmissionPoliciesPropertiesFormatArrayOutputWithContext(ctx context.Context) EmissionPoliciesPropertiesFormatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmissionPoliciesPropertiesFormatArrayOutput)
}

type EmissionPoliciesPropertiesFormatOutput struct{ *pulumi.OutputState }

func (EmissionPoliciesPropertiesFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmissionPoliciesPropertiesFormat)(nil)).Elem()
}

func (o EmissionPoliciesPropertiesFormatOutput) ToEmissionPoliciesPropertiesFormatOutput() EmissionPoliciesPropertiesFormatOutput {
	return o
}

func (o EmissionPoliciesPropertiesFormatOutput) ToEmissionPoliciesPropertiesFormatOutputWithContext(ctx context.Context) EmissionPoliciesPropertiesFormatOutput {
	return o
}

func (o EmissionPoliciesPropertiesFormatOutput) EmissionDestinations() EmissionPolicyDestinationArrayOutput {
	return o.ApplyT(func(v EmissionPoliciesPropertiesFormat) []EmissionPolicyDestination { return v.EmissionDestinations }).(EmissionPolicyDestinationArrayOutput)
}

func (o EmissionPoliciesPropertiesFormatOutput) EmissionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmissionPoliciesPropertiesFormat) *string { return v.EmissionType }).(pulumi.StringPtrOutput)
}

type EmissionPoliciesPropertiesFormatArrayOutput struct{ *pulumi.OutputState }

func (EmissionPoliciesPropertiesFormatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmissionPoliciesPropertiesFormat)(nil)).Elem()
}

func (o EmissionPoliciesPropertiesFormatArrayOutput) ToEmissionPoliciesPropertiesFormatArrayOutput() EmissionPoliciesPropertiesFormatArrayOutput {
	return o
}

func (o EmissionPoliciesPropertiesFormatArrayOutput) ToEmissionPoliciesPropertiesFormatArrayOutputWithContext(ctx context.Context) EmissionPoliciesPropertiesFormatArrayOutput {
	return o
}

func (o EmissionPoliciesPropertiesFormatArrayOutput) Index(i pulumi.IntInput) EmissionPoliciesPropertiesFormatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EmissionPoliciesPropertiesFormat {
		return vs[0].([]EmissionPoliciesPropertiesFormat)[vs[1].(int)]
	}).(EmissionPoliciesPropertiesFormatOutput)
}

type EmissionPoliciesPropertiesFormatResponse struct {
	EmissionDestinations []EmissionPolicyDestinationResponse `pulumi:"emissionDestinations"`
	EmissionType         *string                             `pulumi:"emissionType"`
}

type EmissionPoliciesPropertiesFormatResponseOutput struct{ *pulumi.OutputState }

func (EmissionPoliciesPropertiesFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmissionPoliciesPropertiesFormatResponse)(nil)).Elem()
}

func (o EmissionPoliciesPropertiesFormatResponseOutput) ToEmissionPoliciesPropertiesFormatResponseOutput() EmissionPoliciesPropertiesFormatResponseOutput {
	return o
}

func (o EmissionPoliciesPropertiesFormatResponseOutput) ToEmissionPoliciesPropertiesFormatResponseOutputWithContext(ctx context.Context) EmissionPoliciesPropertiesFormatResponseOutput {
	return o
}

func (o EmissionPoliciesPropertiesFormatResponseOutput) EmissionDestinations() EmissionPolicyDestinationResponseArrayOutput {
	return o.ApplyT(func(v EmissionPoliciesPropertiesFormatResponse) []EmissionPolicyDestinationResponse {
		return v.EmissionDestinations
	}).(EmissionPolicyDestinationResponseArrayOutput)
}

func (o EmissionPoliciesPropertiesFormatResponseOutput) EmissionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmissionPoliciesPropertiesFormatResponse) *string { return v.EmissionType }).(pulumi.StringPtrOutput)
}

type EmissionPoliciesPropertiesFormatResponseArrayOutput struct{ *pulumi.OutputState }

func (EmissionPoliciesPropertiesFormatResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmissionPoliciesPropertiesFormatResponse)(nil)).Elem()
}

func (o EmissionPoliciesPropertiesFormatResponseArrayOutput) ToEmissionPoliciesPropertiesFormatResponseArrayOutput() EmissionPoliciesPropertiesFormatResponseArrayOutput {
	return o
}

func (o EmissionPoliciesPropertiesFormatResponseArrayOutput) ToEmissionPoliciesPropertiesFormatResponseArrayOutputWithContext(ctx context.Context) EmissionPoliciesPropertiesFormatResponseArrayOutput {
	return o
}

func (o EmissionPoliciesPropertiesFormatResponseArrayOutput) Index(i pulumi.IntInput) EmissionPoliciesPropertiesFormatResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EmissionPoliciesPropertiesFormatResponse {
		return vs[0].([]EmissionPoliciesPropertiesFormatResponse)[vs[1].(int)]
	}).(EmissionPoliciesPropertiesFormatResponseOutput)
}

type EmissionPolicyDestination struct {
	DestinationType *string `pulumi:"destinationType"`
}





type EmissionPolicyDestinationInput interface {
	pulumi.Input

	ToEmissionPolicyDestinationOutput() EmissionPolicyDestinationOutput
	ToEmissionPolicyDestinationOutputWithContext(context.Context) EmissionPolicyDestinationOutput
}

type EmissionPolicyDestinationArgs struct {
	DestinationType pulumi.StringPtrInput `pulumi:"destinationType"`
}

func (EmissionPolicyDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmissionPolicyDestination)(nil)).Elem()
}

func (i EmissionPolicyDestinationArgs) ToEmissionPolicyDestinationOutput() EmissionPolicyDestinationOutput {
	return i.ToEmissionPolicyDestinationOutputWithContext(context.Background())
}

func (i EmissionPolicyDestinationArgs) ToEmissionPolicyDestinationOutputWithContext(ctx context.Context) EmissionPolicyDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmissionPolicyDestinationOutput)
}





type EmissionPolicyDestinationArrayInput interface {
	pulumi.Input

	ToEmissionPolicyDestinationArrayOutput() EmissionPolicyDestinationArrayOutput
	ToEmissionPolicyDestinationArrayOutputWithContext(context.Context) EmissionPolicyDestinationArrayOutput
}

type EmissionPolicyDestinationArray []EmissionPolicyDestinationInput

func (EmissionPolicyDestinationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmissionPolicyDestination)(nil)).Elem()
}

func (i EmissionPolicyDestinationArray) ToEmissionPolicyDestinationArrayOutput() EmissionPolicyDestinationArrayOutput {
	return i.ToEmissionPolicyDestinationArrayOutputWithContext(context.Background())
}

func (i EmissionPolicyDestinationArray) ToEmissionPolicyDestinationArrayOutputWithContext(ctx context.Context) EmissionPolicyDestinationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmissionPolicyDestinationArrayOutput)
}

type EmissionPolicyDestinationOutput struct{ *pulumi.OutputState }

func (EmissionPolicyDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmissionPolicyDestination)(nil)).Elem()
}

func (o EmissionPolicyDestinationOutput) ToEmissionPolicyDestinationOutput() EmissionPolicyDestinationOutput {
	return o
}

func (o EmissionPolicyDestinationOutput) ToEmissionPolicyDestinationOutputWithContext(ctx context.Context) EmissionPolicyDestinationOutput {
	return o
}

func (o EmissionPolicyDestinationOutput) DestinationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmissionPolicyDestination) *string { return v.DestinationType }).(pulumi.StringPtrOutput)
}

type EmissionPolicyDestinationArrayOutput struct{ *pulumi.OutputState }

func (EmissionPolicyDestinationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmissionPolicyDestination)(nil)).Elem()
}

func (o EmissionPolicyDestinationArrayOutput) ToEmissionPolicyDestinationArrayOutput() EmissionPolicyDestinationArrayOutput {
	return o
}

func (o EmissionPolicyDestinationArrayOutput) ToEmissionPolicyDestinationArrayOutputWithContext(ctx context.Context) EmissionPolicyDestinationArrayOutput {
	return o
}

func (o EmissionPolicyDestinationArrayOutput) Index(i pulumi.IntInput) EmissionPolicyDestinationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EmissionPolicyDestination {
		return vs[0].([]EmissionPolicyDestination)[vs[1].(int)]
	}).(EmissionPolicyDestinationOutput)
}

type EmissionPolicyDestinationResponse struct {
	DestinationType *string `pulumi:"destinationType"`
}

type EmissionPolicyDestinationResponseOutput struct{ *pulumi.OutputState }

func (EmissionPolicyDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmissionPolicyDestinationResponse)(nil)).Elem()
}

func (o EmissionPolicyDestinationResponseOutput) ToEmissionPolicyDestinationResponseOutput() EmissionPolicyDestinationResponseOutput {
	return o
}

func (o EmissionPolicyDestinationResponseOutput) ToEmissionPolicyDestinationResponseOutputWithContext(ctx context.Context) EmissionPolicyDestinationResponseOutput {
	return o
}

func (o EmissionPolicyDestinationResponseOutput) DestinationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmissionPolicyDestinationResponse) *string { return v.DestinationType }).(pulumi.StringPtrOutput)
}

type EmissionPolicyDestinationResponseArrayOutput struct{ *pulumi.OutputState }

func (EmissionPolicyDestinationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmissionPolicyDestinationResponse)(nil)).Elem()
}

func (o EmissionPolicyDestinationResponseArrayOutput) ToEmissionPolicyDestinationResponseArrayOutput() EmissionPolicyDestinationResponseArrayOutput {
	return o
}

func (o EmissionPolicyDestinationResponseArrayOutput) ToEmissionPolicyDestinationResponseArrayOutputWithContext(ctx context.Context) EmissionPolicyDestinationResponseArrayOutput {
	return o
}

func (o EmissionPolicyDestinationResponseArrayOutput) Index(i pulumi.IntInput) EmissionPolicyDestinationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EmissionPolicyDestinationResponse {
		return vs[0].([]EmissionPolicyDestinationResponse)[vs[1].(int)]
	}).(EmissionPolicyDestinationResponseOutput)
}

type IngestionPolicyPropertiesFormat struct {
	IngestionSources []IngestionSourcesPropertiesFormat `pulumi:"ingestionSources"`
	IngestionType    *string                            `pulumi:"ingestionType"`
}





type IngestionPolicyPropertiesFormatInput interface {
	pulumi.Input

	ToIngestionPolicyPropertiesFormatOutput() IngestionPolicyPropertiesFormatOutput
	ToIngestionPolicyPropertiesFormatOutputWithContext(context.Context) IngestionPolicyPropertiesFormatOutput
}

type IngestionPolicyPropertiesFormatArgs struct {
	IngestionSources IngestionSourcesPropertiesFormatArrayInput `pulumi:"ingestionSources"`
	IngestionType    pulumi.StringPtrInput                      `pulumi:"ingestionType"`
}

func (IngestionPolicyPropertiesFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestionPolicyPropertiesFormat)(nil)).Elem()
}

func (i IngestionPolicyPropertiesFormatArgs) ToIngestionPolicyPropertiesFormatOutput() IngestionPolicyPropertiesFormatOutput {
	return i.ToIngestionPolicyPropertiesFormatOutputWithContext(context.Background())
}

func (i IngestionPolicyPropertiesFormatArgs) ToIngestionPolicyPropertiesFormatOutputWithContext(ctx context.Context) IngestionPolicyPropertiesFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionPolicyPropertiesFormatOutput)
}

func (i IngestionPolicyPropertiesFormatArgs) ToIngestionPolicyPropertiesFormatPtrOutput() IngestionPolicyPropertiesFormatPtrOutput {
	return i.ToIngestionPolicyPropertiesFormatPtrOutputWithContext(context.Background())
}

func (i IngestionPolicyPropertiesFormatArgs) ToIngestionPolicyPropertiesFormatPtrOutputWithContext(ctx context.Context) IngestionPolicyPropertiesFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionPolicyPropertiesFormatOutput).ToIngestionPolicyPropertiesFormatPtrOutputWithContext(ctx)
}









type IngestionPolicyPropertiesFormatPtrInput interface {
	pulumi.Input

	ToIngestionPolicyPropertiesFormatPtrOutput() IngestionPolicyPropertiesFormatPtrOutput
	ToIngestionPolicyPropertiesFormatPtrOutputWithContext(context.Context) IngestionPolicyPropertiesFormatPtrOutput
}

type ingestionPolicyPropertiesFormatPtrType IngestionPolicyPropertiesFormatArgs

func IngestionPolicyPropertiesFormatPtr(v *IngestionPolicyPropertiesFormatArgs) IngestionPolicyPropertiesFormatPtrInput {
	return (*ingestionPolicyPropertiesFormatPtrType)(v)
}

func (*ingestionPolicyPropertiesFormatPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IngestionPolicyPropertiesFormat)(nil)).Elem()
}

func (i *ingestionPolicyPropertiesFormatPtrType) ToIngestionPolicyPropertiesFormatPtrOutput() IngestionPolicyPropertiesFormatPtrOutput {
	return i.ToIngestionPolicyPropertiesFormatPtrOutputWithContext(context.Background())
}

func (i *ingestionPolicyPropertiesFormatPtrType) ToIngestionPolicyPropertiesFormatPtrOutputWithContext(ctx context.Context) IngestionPolicyPropertiesFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionPolicyPropertiesFormatPtrOutput)
}

type IngestionPolicyPropertiesFormatOutput struct{ *pulumi.OutputState }

func (IngestionPolicyPropertiesFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestionPolicyPropertiesFormat)(nil)).Elem()
}

func (o IngestionPolicyPropertiesFormatOutput) ToIngestionPolicyPropertiesFormatOutput() IngestionPolicyPropertiesFormatOutput {
	return o
}

func (o IngestionPolicyPropertiesFormatOutput) ToIngestionPolicyPropertiesFormatOutputWithContext(ctx context.Context) IngestionPolicyPropertiesFormatOutput {
	return o
}

func (o IngestionPolicyPropertiesFormatOutput) ToIngestionPolicyPropertiesFormatPtrOutput() IngestionPolicyPropertiesFormatPtrOutput {
	return o.ToIngestionPolicyPropertiesFormatPtrOutputWithContext(context.Background())
}

func (o IngestionPolicyPropertiesFormatOutput) ToIngestionPolicyPropertiesFormatPtrOutputWithContext(ctx context.Context) IngestionPolicyPropertiesFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IngestionPolicyPropertiesFormat) *IngestionPolicyPropertiesFormat {
		return &v
	}).(IngestionPolicyPropertiesFormatPtrOutput)
}

func (o IngestionPolicyPropertiesFormatOutput) IngestionSources() IngestionSourcesPropertiesFormatArrayOutput {
	return o.ApplyT(func(v IngestionPolicyPropertiesFormat) []IngestionSourcesPropertiesFormat { return v.IngestionSources }).(IngestionSourcesPropertiesFormatArrayOutput)
}

func (o IngestionPolicyPropertiesFormatOutput) IngestionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestionPolicyPropertiesFormat) *string { return v.IngestionType }).(pulumi.StringPtrOutput)
}

type IngestionPolicyPropertiesFormatPtrOutput struct{ *pulumi.OutputState }

func (IngestionPolicyPropertiesFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngestionPolicyPropertiesFormat)(nil)).Elem()
}

func (o IngestionPolicyPropertiesFormatPtrOutput) ToIngestionPolicyPropertiesFormatPtrOutput() IngestionPolicyPropertiesFormatPtrOutput {
	return o
}

func (o IngestionPolicyPropertiesFormatPtrOutput) ToIngestionPolicyPropertiesFormatPtrOutputWithContext(ctx context.Context) IngestionPolicyPropertiesFormatPtrOutput {
	return o
}

func (o IngestionPolicyPropertiesFormatPtrOutput) Elem() IngestionPolicyPropertiesFormatOutput {
	return o.ApplyT(func(v *IngestionPolicyPropertiesFormat) IngestionPolicyPropertiesFormat {
		if v != nil {
			return *v
		}
		var ret IngestionPolicyPropertiesFormat
		return ret
	}).(IngestionPolicyPropertiesFormatOutput)
}

func (o IngestionPolicyPropertiesFormatPtrOutput) IngestionSources() IngestionSourcesPropertiesFormatArrayOutput {
	return o.ApplyT(func(v *IngestionPolicyPropertiesFormat) []IngestionSourcesPropertiesFormat {
		if v == nil {
			return nil
		}
		return v.IngestionSources
	}).(IngestionSourcesPropertiesFormatArrayOutput)
}

func (o IngestionPolicyPropertiesFormatPtrOutput) IngestionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngestionPolicyPropertiesFormat) *string {
		if v == nil {
			return nil
		}
		return v.IngestionType
	}).(pulumi.StringPtrOutput)
}

type IngestionPolicyPropertiesFormatResponse struct {
	IngestionSources []IngestionSourcesPropertiesFormatResponse `pulumi:"ingestionSources"`
	IngestionType    *string                                    `pulumi:"ingestionType"`
}

type IngestionPolicyPropertiesFormatResponseOutput struct{ *pulumi.OutputState }

func (IngestionPolicyPropertiesFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestionPolicyPropertiesFormatResponse)(nil)).Elem()
}

func (o IngestionPolicyPropertiesFormatResponseOutput) ToIngestionPolicyPropertiesFormatResponseOutput() IngestionPolicyPropertiesFormatResponseOutput {
	return o
}

func (o IngestionPolicyPropertiesFormatResponseOutput) ToIngestionPolicyPropertiesFormatResponseOutputWithContext(ctx context.Context) IngestionPolicyPropertiesFormatResponseOutput {
	return o
}

func (o IngestionPolicyPropertiesFormatResponseOutput) IngestionSources() IngestionSourcesPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v IngestionPolicyPropertiesFormatResponse) []IngestionSourcesPropertiesFormatResponse {
		return v.IngestionSources
	}).(IngestionSourcesPropertiesFormatResponseArrayOutput)
}

func (o IngestionPolicyPropertiesFormatResponseOutput) IngestionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestionPolicyPropertiesFormatResponse) *string { return v.IngestionType }).(pulumi.StringPtrOutput)
}

type IngestionPolicyPropertiesFormatResponsePtrOutput struct{ *pulumi.OutputState }

func (IngestionPolicyPropertiesFormatResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngestionPolicyPropertiesFormatResponse)(nil)).Elem()
}

func (o IngestionPolicyPropertiesFormatResponsePtrOutput) ToIngestionPolicyPropertiesFormatResponsePtrOutput() IngestionPolicyPropertiesFormatResponsePtrOutput {
	return o
}

func (o IngestionPolicyPropertiesFormatResponsePtrOutput) ToIngestionPolicyPropertiesFormatResponsePtrOutputWithContext(ctx context.Context) IngestionPolicyPropertiesFormatResponsePtrOutput {
	return o
}

func (o IngestionPolicyPropertiesFormatResponsePtrOutput) Elem() IngestionPolicyPropertiesFormatResponseOutput {
	return o.ApplyT(func(v *IngestionPolicyPropertiesFormatResponse) IngestionPolicyPropertiesFormatResponse {
		if v != nil {
			return *v
		}
		var ret IngestionPolicyPropertiesFormatResponse
		return ret
	}).(IngestionPolicyPropertiesFormatResponseOutput)
}

func (o IngestionPolicyPropertiesFormatResponsePtrOutput) IngestionSources() IngestionSourcesPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v *IngestionPolicyPropertiesFormatResponse) []IngestionSourcesPropertiesFormatResponse {
		if v == nil {
			return nil
		}
		return v.IngestionSources
	}).(IngestionSourcesPropertiesFormatResponseArrayOutput)
}

func (o IngestionPolicyPropertiesFormatResponsePtrOutput) IngestionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngestionPolicyPropertiesFormatResponse) *string {
		if v == nil {
			return nil
		}
		return v.IngestionType
	}).(pulumi.StringPtrOutput)
}

type IngestionSourcesPropertiesFormat struct {
	ResourceId *string `pulumi:"resourceId"`
	SourceType *string `pulumi:"sourceType"`
}





type IngestionSourcesPropertiesFormatInput interface {
	pulumi.Input

	ToIngestionSourcesPropertiesFormatOutput() IngestionSourcesPropertiesFormatOutput
	ToIngestionSourcesPropertiesFormatOutputWithContext(context.Context) IngestionSourcesPropertiesFormatOutput
}

type IngestionSourcesPropertiesFormatArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	SourceType pulumi.StringPtrInput `pulumi:"sourceType"`
}

func (IngestionSourcesPropertiesFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestionSourcesPropertiesFormat)(nil)).Elem()
}

func (i IngestionSourcesPropertiesFormatArgs) ToIngestionSourcesPropertiesFormatOutput() IngestionSourcesPropertiesFormatOutput {
	return i.ToIngestionSourcesPropertiesFormatOutputWithContext(context.Background())
}

func (i IngestionSourcesPropertiesFormatArgs) ToIngestionSourcesPropertiesFormatOutputWithContext(ctx context.Context) IngestionSourcesPropertiesFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionSourcesPropertiesFormatOutput)
}





type IngestionSourcesPropertiesFormatArrayInput interface {
	pulumi.Input

	ToIngestionSourcesPropertiesFormatArrayOutput() IngestionSourcesPropertiesFormatArrayOutput
	ToIngestionSourcesPropertiesFormatArrayOutputWithContext(context.Context) IngestionSourcesPropertiesFormatArrayOutput
}

type IngestionSourcesPropertiesFormatArray []IngestionSourcesPropertiesFormatInput

func (IngestionSourcesPropertiesFormatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IngestionSourcesPropertiesFormat)(nil)).Elem()
}

func (i IngestionSourcesPropertiesFormatArray) ToIngestionSourcesPropertiesFormatArrayOutput() IngestionSourcesPropertiesFormatArrayOutput {
	return i.ToIngestionSourcesPropertiesFormatArrayOutputWithContext(context.Background())
}

func (i IngestionSourcesPropertiesFormatArray) ToIngestionSourcesPropertiesFormatArrayOutputWithContext(ctx context.Context) IngestionSourcesPropertiesFormatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionSourcesPropertiesFormatArrayOutput)
}

type IngestionSourcesPropertiesFormatOutput struct{ *pulumi.OutputState }

func (IngestionSourcesPropertiesFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestionSourcesPropertiesFormat)(nil)).Elem()
}

func (o IngestionSourcesPropertiesFormatOutput) ToIngestionSourcesPropertiesFormatOutput() IngestionSourcesPropertiesFormatOutput {
	return o
}

func (o IngestionSourcesPropertiesFormatOutput) ToIngestionSourcesPropertiesFormatOutputWithContext(ctx context.Context) IngestionSourcesPropertiesFormatOutput {
	return o
}

func (o IngestionSourcesPropertiesFormatOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestionSourcesPropertiesFormat) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o IngestionSourcesPropertiesFormatOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestionSourcesPropertiesFormat) *string { return v.SourceType }).(pulumi.StringPtrOutput)
}

type IngestionSourcesPropertiesFormatArrayOutput struct{ *pulumi.OutputState }

func (IngestionSourcesPropertiesFormatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IngestionSourcesPropertiesFormat)(nil)).Elem()
}

func (o IngestionSourcesPropertiesFormatArrayOutput) ToIngestionSourcesPropertiesFormatArrayOutput() IngestionSourcesPropertiesFormatArrayOutput {
	return o
}

func (o IngestionSourcesPropertiesFormatArrayOutput) ToIngestionSourcesPropertiesFormatArrayOutputWithContext(ctx context.Context) IngestionSourcesPropertiesFormatArrayOutput {
	return o
}

func (o IngestionSourcesPropertiesFormatArrayOutput) Index(i pulumi.IntInput) IngestionSourcesPropertiesFormatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IngestionSourcesPropertiesFormat {
		return vs[0].([]IngestionSourcesPropertiesFormat)[vs[1].(int)]
	}).(IngestionSourcesPropertiesFormatOutput)
}

type IngestionSourcesPropertiesFormatResponse struct {
	ResourceId *string `pulumi:"resourceId"`
	SourceType *string `pulumi:"sourceType"`
}

type IngestionSourcesPropertiesFormatResponseOutput struct{ *pulumi.OutputState }

func (IngestionSourcesPropertiesFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestionSourcesPropertiesFormatResponse)(nil)).Elem()
}

func (o IngestionSourcesPropertiesFormatResponseOutput) ToIngestionSourcesPropertiesFormatResponseOutput() IngestionSourcesPropertiesFormatResponseOutput {
	return o
}

func (o IngestionSourcesPropertiesFormatResponseOutput) ToIngestionSourcesPropertiesFormatResponseOutputWithContext(ctx context.Context) IngestionSourcesPropertiesFormatResponseOutput {
	return o
}

func (o IngestionSourcesPropertiesFormatResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestionSourcesPropertiesFormatResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o IngestionSourcesPropertiesFormatResponseOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngestionSourcesPropertiesFormatResponse) *string { return v.SourceType }).(pulumi.StringPtrOutput)
}

type IngestionSourcesPropertiesFormatResponseArrayOutput struct{ *pulumi.OutputState }

func (IngestionSourcesPropertiesFormatResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IngestionSourcesPropertiesFormatResponse)(nil)).Elem()
}

func (o IngestionSourcesPropertiesFormatResponseArrayOutput) ToIngestionSourcesPropertiesFormatResponseArrayOutput() IngestionSourcesPropertiesFormatResponseArrayOutput {
	return o
}

func (o IngestionSourcesPropertiesFormatResponseArrayOutput) ToIngestionSourcesPropertiesFormatResponseArrayOutputWithContext(ctx context.Context) IngestionSourcesPropertiesFormatResponseArrayOutput {
	return o
}

func (o IngestionSourcesPropertiesFormatResponseArrayOutput) Index(i pulumi.IntInput) IngestionSourcesPropertiesFormatResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IngestionSourcesPropertiesFormatResponse {
		return vs[0].([]IngestionSourcesPropertiesFormatResponse)[vs[1].(int)]
	}).(IngestionSourcesPropertiesFormatResponseOutput)
}

type ResourceReferenceResponse struct {
	Id string `pulumi:"id"`
}

type ResourceReferenceResponseOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ResourceReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponsePtrOutput) ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput {
	return o
}

func (o ResourceReferenceResponsePtrOutput) ToResourceReferenceResponsePtrOutputWithContext(ctx context.Context) ResourceReferenceResponsePtrOutput {
	return o
}

func (o ResourceReferenceResponsePtrOutput) Elem() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) ResourceReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ResourceReferenceResponse
		return ret
	}).(ResourceReferenceResponseOutput)
}

func (o ResourceReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ResourceReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseArrayOutput) ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput {
	return o
}

func (o ResourceReferenceResponseArrayOutput) ToResourceReferenceResponseArrayOutputWithContext(ctx context.Context) ResourceReferenceResponseArrayOutput {
	return o
}

func (o ResourceReferenceResponseArrayOutput) Index(i pulumi.IntInput) ResourceReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceReferenceResponse {
		return vs[0].([]ResourceReferenceResponse)[vs[1].(int)]
	}).(ResourceReferenceResponseOutput)
}

type TrackedResourceResponseSystemData struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type TrackedResourceResponseSystemDataOutput struct{ *pulumi.OutputState }

func (TrackedResourceResponseSystemDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackedResourceResponseSystemData)(nil)).Elem()
}

func (o TrackedResourceResponseSystemDataOutput) ToTrackedResourceResponseSystemDataOutput() TrackedResourceResponseSystemDataOutput {
	return o
}

func (o TrackedResourceResponseSystemDataOutput) ToTrackedResourceResponseSystemDataOutputWithContext(ctx context.Context) TrackedResourceResponseSystemDataOutput {
	return o
}

func (o TrackedResourceResponseSystemDataOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrackedResourceResponseSystemData) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o TrackedResourceResponseSystemDataOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrackedResourceResponseSystemData) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o TrackedResourceResponseSystemDataOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrackedResourceResponseSystemData) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o TrackedResourceResponseSystemDataOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrackedResourceResponseSystemData) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o TrackedResourceResponseSystemDataOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrackedResourceResponseSystemData) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EmissionPoliciesPropertiesFormatOutput{})
	pulumi.RegisterOutputType(EmissionPoliciesPropertiesFormatArrayOutput{})
	pulumi.RegisterOutputType(EmissionPoliciesPropertiesFormatResponseOutput{})
	pulumi.RegisterOutputType(EmissionPoliciesPropertiesFormatResponseArrayOutput{})
	pulumi.RegisterOutputType(EmissionPolicyDestinationOutput{})
	pulumi.RegisterOutputType(EmissionPolicyDestinationArrayOutput{})
	pulumi.RegisterOutputType(EmissionPolicyDestinationResponseOutput{})
	pulumi.RegisterOutputType(EmissionPolicyDestinationResponseArrayOutput{})
	pulumi.RegisterOutputType(IngestionPolicyPropertiesFormatOutput{})
	pulumi.RegisterOutputType(IngestionPolicyPropertiesFormatPtrOutput{})
	pulumi.RegisterOutputType(IngestionPolicyPropertiesFormatResponseOutput{})
	pulumi.RegisterOutputType(IngestionPolicyPropertiesFormatResponsePtrOutput{})
	pulumi.RegisterOutputType(IngestionSourcesPropertiesFormatOutput{})
	pulumi.RegisterOutputType(IngestionSourcesPropertiesFormatArrayOutput{})
	pulumi.RegisterOutputType(IngestionSourcesPropertiesFormatResponseOutput{})
	pulumi.RegisterOutputType(IngestionSourcesPropertiesFormatResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(TrackedResourceResponseSystemDataOutput{})
}
