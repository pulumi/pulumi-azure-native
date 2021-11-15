


package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssignmentPrincipal struct {
	PrincipalId       string            `pulumi:"principalId"`
	PrincipalMetadata map[string]string `pulumi:"principalMetadata"`
	PrincipalType     string            `pulumi:"principalType"`
}





type AssignmentPrincipalInput interface {
	pulumi.Input

	ToAssignmentPrincipalOutput() AssignmentPrincipalOutput
	ToAssignmentPrincipalOutputWithContext(context.Context) AssignmentPrincipalOutput
}

type AssignmentPrincipalArgs struct {
	PrincipalId       pulumi.StringInput    `pulumi:"principalId"`
	PrincipalMetadata pulumi.StringMapInput `pulumi:"principalMetadata"`
	PrincipalType     pulumi.StringInput    `pulumi:"principalType"`
}

func (AssignmentPrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentPrincipal)(nil)).Elem()
}

func (i AssignmentPrincipalArgs) ToAssignmentPrincipalOutput() AssignmentPrincipalOutput {
	return i.ToAssignmentPrincipalOutputWithContext(context.Background())
}

func (i AssignmentPrincipalArgs) ToAssignmentPrincipalOutputWithContext(ctx context.Context) AssignmentPrincipalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPrincipalOutput)
}





type AssignmentPrincipalArrayInput interface {
	pulumi.Input

	ToAssignmentPrincipalArrayOutput() AssignmentPrincipalArrayOutput
	ToAssignmentPrincipalArrayOutputWithContext(context.Context) AssignmentPrincipalArrayOutput
}

type AssignmentPrincipalArray []AssignmentPrincipalInput

func (AssignmentPrincipalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssignmentPrincipal)(nil)).Elem()
}

func (i AssignmentPrincipalArray) ToAssignmentPrincipalArrayOutput() AssignmentPrincipalArrayOutput {
	return i.ToAssignmentPrincipalArrayOutputWithContext(context.Background())
}

func (i AssignmentPrincipalArray) ToAssignmentPrincipalArrayOutputWithContext(ctx context.Context) AssignmentPrincipalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPrincipalArrayOutput)
}

type AssignmentPrincipalOutput struct{ *pulumi.OutputState }

func (AssignmentPrincipalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentPrincipal)(nil)).Elem()
}

func (o AssignmentPrincipalOutput) ToAssignmentPrincipalOutput() AssignmentPrincipalOutput {
	return o
}

func (o AssignmentPrincipalOutput) ToAssignmentPrincipalOutputWithContext(ctx context.Context) AssignmentPrincipalOutput {
	return o
}

func (o AssignmentPrincipalOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentPrincipal) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o AssignmentPrincipalOutput) PrincipalMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v AssignmentPrincipal) map[string]string { return v.PrincipalMetadata }).(pulumi.StringMapOutput)
}

func (o AssignmentPrincipalOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentPrincipal) string { return v.PrincipalType }).(pulumi.StringOutput)
}

type AssignmentPrincipalArrayOutput struct{ *pulumi.OutputState }

func (AssignmentPrincipalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssignmentPrincipal)(nil)).Elem()
}

func (o AssignmentPrincipalArrayOutput) ToAssignmentPrincipalArrayOutput() AssignmentPrincipalArrayOutput {
	return o
}

func (o AssignmentPrincipalArrayOutput) ToAssignmentPrincipalArrayOutputWithContext(ctx context.Context) AssignmentPrincipalArrayOutput {
	return o
}

func (o AssignmentPrincipalArrayOutput) Index(i pulumi.IntInput) AssignmentPrincipalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AssignmentPrincipal {
		return vs[0].([]AssignmentPrincipal)[vs[1].(int)]
	}).(AssignmentPrincipalOutput)
}

type AssignmentPrincipalResponse struct {
	PrincipalId       string            `pulumi:"principalId"`
	PrincipalMetadata map[string]string `pulumi:"principalMetadata"`
	PrincipalType     string            `pulumi:"principalType"`
}





type AssignmentPrincipalResponseInput interface {
	pulumi.Input

	ToAssignmentPrincipalResponseOutput() AssignmentPrincipalResponseOutput
	ToAssignmentPrincipalResponseOutputWithContext(context.Context) AssignmentPrincipalResponseOutput
}

type AssignmentPrincipalResponseArgs struct {
	PrincipalId       pulumi.StringInput    `pulumi:"principalId"`
	PrincipalMetadata pulumi.StringMapInput `pulumi:"principalMetadata"`
	PrincipalType     pulumi.StringInput    `pulumi:"principalType"`
}

func (AssignmentPrincipalResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentPrincipalResponse)(nil)).Elem()
}

func (i AssignmentPrincipalResponseArgs) ToAssignmentPrincipalResponseOutput() AssignmentPrincipalResponseOutput {
	return i.ToAssignmentPrincipalResponseOutputWithContext(context.Background())
}

func (i AssignmentPrincipalResponseArgs) ToAssignmentPrincipalResponseOutputWithContext(ctx context.Context) AssignmentPrincipalResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPrincipalResponseOutput)
}





type AssignmentPrincipalResponseArrayInput interface {
	pulumi.Input

	ToAssignmentPrincipalResponseArrayOutput() AssignmentPrincipalResponseArrayOutput
	ToAssignmentPrincipalResponseArrayOutputWithContext(context.Context) AssignmentPrincipalResponseArrayOutput
}

type AssignmentPrincipalResponseArray []AssignmentPrincipalResponseInput

func (AssignmentPrincipalResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssignmentPrincipalResponse)(nil)).Elem()
}

func (i AssignmentPrincipalResponseArray) ToAssignmentPrincipalResponseArrayOutput() AssignmentPrincipalResponseArrayOutput {
	return i.ToAssignmentPrincipalResponseArrayOutputWithContext(context.Background())
}

func (i AssignmentPrincipalResponseArray) ToAssignmentPrincipalResponseArrayOutputWithContext(ctx context.Context) AssignmentPrincipalResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPrincipalResponseArrayOutput)
}

type AssignmentPrincipalResponseOutput struct{ *pulumi.OutputState }

func (AssignmentPrincipalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentPrincipalResponse)(nil)).Elem()
}

func (o AssignmentPrincipalResponseOutput) ToAssignmentPrincipalResponseOutput() AssignmentPrincipalResponseOutput {
	return o
}

func (o AssignmentPrincipalResponseOutput) ToAssignmentPrincipalResponseOutputWithContext(ctx context.Context) AssignmentPrincipalResponseOutput {
	return o
}

func (o AssignmentPrincipalResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentPrincipalResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o AssignmentPrincipalResponseOutput) PrincipalMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v AssignmentPrincipalResponse) map[string]string { return v.PrincipalMetadata }).(pulumi.StringMapOutput)
}

func (o AssignmentPrincipalResponseOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentPrincipalResponse) string { return v.PrincipalType }).(pulumi.StringOutput)
}

type AssignmentPrincipalResponseArrayOutput struct{ *pulumi.OutputState }

func (AssignmentPrincipalResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssignmentPrincipalResponse)(nil)).Elem()
}

func (o AssignmentPrincipalResponseArrayOutput) ToAssignmentPrincipalResponseArrayOutput() AssignmentPrincipalResponseArrayOutput {
	return o
}

func (o AssignmentPrincipalResponseArrayOutput) ToAssignmentPrincipalResponseArrayOutputWithContext(ctx context.Context) AssignmentPrincipalResponseArrayOutput {
	return o
}

func (o AssignmentPrincipalResponseArrayOutput) Index(i pulumi.IntInput) AssignmentPrincipalResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AssignmentPrincipalResponse {
		return vs[0].([]AssignmentPrincipalResponse)[vs[1].(int)]
	}).(AssignmentPrincipalResponseOutput)
}

type ConnectorMappingAvailability struct {
	Frequency *FrequencyTypes `pulumi:"frequency"`
	Interval  int             `pulumi:"interval"`
}





type ConnectorMappingAvailabilityInput interface {
	pulumi.Input

	ToConnectorMappingAvailabilityOutput() ConnectorMappingAvailabilityOutput
	ToConnectorMappingAvailabilityOutputWithContext(context.Context) ConnectorMappingAvailabilityOutput
}

type ConnectorMappingAvailabilityArgs struct {
	Frequency FrequencyTypesPtrInput `pulumi:"frequency"`
	Interval  pulumi.IntInput        `pulumi:"interval"`
}

func (ConnectorMappingAvailabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingAvailability)(nil)).Elem()
}

func (i ConnectorMappingAvailabilityArgs) ToConnectorMappingAvailabilityOutput() ConnectorMappingAvailabilityOutput {
	return i.ToConnectorMappingAvailabilityOutputWithContext(context.Background())
}

func (i ConnectorMappingAvailabilityArgs) ToConnectorMappingAvailabilityOutputWithContext(ctx context.Context) ConnectorMappingAvailabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingAvailabilityOutput)
}

func (i ConnectorMappingAvailabilityArgs) ToConnectorMappingAvailabilityPtrOutput() ConnectorMappingAvailabilityPtrOutput {
	return i.ToConnectorMappingAvailabilityPtrOutputWithContext(context.Background())
}

func (i ConnectorMappingAvailabilityArgs) ToConnectorMappingAvailabilityPtrOutputWithContext(ctx context.Context) ConnectorMappingAvailabilityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingAvailabilityOutput).ToConnectorMappingAvailabilityPtrOutputWithContext(ctx)
}









type ConnectorMappingAvailabilityPtrInput interface {
	pulumi.Input

	ToConnectorMappingAvailabilityPtrOutput() ConnectorMappingAvailabilityPtrOutput
	ToConnectorMappingAvailabilityPtrOutputWithContext(context.Context) ConnectorMappingAvailabilityPtrOutput
}

type connectorMappingAvailabilityPtrType ConnectorMappingAvailabilityArgs

func ConnectorMappingAvailabilityPtr(v *ConnectorMappingAvailabilityArgs) ConnectorMappingAvailabilityPtrInput {
	return (*connectorMappingAvailabilityPtrType)(v)
}

func (*connectorMappingAvailabilityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingAvailability)(nil)).Elem()
}

func (i *connectorMappingAvailabilityPtrType) ToConnectorMappingAvailabilityPtrOutput() ConnectorMappingAvailabilityPtrOutput {
	return i.ToConnectorMappingAvailabilityPtrOutputWithContext(context.Background())
}

func (i *connectorMappingAvailabilityPtrType) ToConnectorMappingAvailabilityPtrOutputWithContext(ctx context.Context) ConnectorMappingAvailabilityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingAvailabilityPtrOutput)
}

type ConnectorMappingAvailabilityOutput struct{ *pulumi.OutputState }

func (ConnectorMappingAvailabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingAvailability)(nil)).Elem()
}

func (o ConnectorMappingAvailabilityOutput) ToConnectorMappingAvailabilityOutput() ConnectorMappingAvailabilityOutput {
	return o
}

func (o ConnectorMappingAvailabilityOutput) ToConnectorMappingAvailabilityOutputWithContext(ctx context.Context) ConnectorMappingAvailabilityOutput {
	return o
}

func (o ConnectorMappingAvailabilityOutput) ToConnectorMappingAvailabilityPtrOutput() ConnectorMappingAvailabilityPtrOutput {
	return o.ToConnectorMappingAvailabilityPtrOutputWithContext(context.Background())
}

func (o ConnectorMappingAvailabilityOutput) ToConnectorMappingAvailabilityPtrOutputWithContext(ctx context.Context) ConnectorMappingAvailabilityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectorMappingAvailability) *ConnectorMappingAvailability {
		return &v
	}).(ConnectorMappingAvailabilityPtrOutput)
}

func (o ConnectorMappingAvailabilityOutput) Frequency() FrequencyTypesPtrOutput {
	return o.ApplyT(func(v ConnectorMappingAvailability) *FrequencyTypes { return v.Frequency }).(FrequencyTypesPtrOutput)
}

func (o ConnectorMappingAvailabilityOutput) Interval() pulumi.IntOutput {
	return o.ApplyT(func(v ConnectorMappingAvailability) int { return v.Interval }).(pulumi.IntOutput)
}

type ConnectorMappingAvailabilityPtrOutput struct{ *pulumi.OutputState }

func (ConnectorMappingAvailabilityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingAvailability)(nil)).Elem()
}

func (o ConnectorMappingAvailabilityPtrOutput) ToConnectorMappingAvailabilityPtrOutput() ConnectorMappingAvailabilityPtrOutput {
	return o
}

func (o ConnectorMappingAvailabilityPtrOutput) ToConnectorMappingAvailabilityPtrOutputWithContext(ctx context.Context) ConnectorMappingAvailabilityPtrOutput {
	return o
}

func (o ConnectorMappingAvailabilityPtrOutput) Elem() ConnectorMappingAvailabilityOutput {
	return o.ApplyT(func(v *ConnectorMappingAvailability) ConnectorMappingAvailability {
		if v != nil {
			return *v
		}
		var ret ConnectorMappingAvailability
		return ret
	}).(ConnectorMappingAvailabilityOutput)
}

func (o ConnectorMappingAvailabilityPtrOutput) Frequency() FrequencyTypesPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingAvailability) *FrequencyTypes {
		if v == nil {
			return nil
		}
		return v.Frequency
	}).(FrequencyTypesPtrOutput)
}

func (o ConnectorMappingAvailabilityPtrOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingAvailability) *int {
		if v == nil {
			return nil
		}
		return &v.Interval
	}).(pulumi.IntPtrOutput)
}

type ConnectorMappingAvailabilityResponse struct {
	Frequency *string `pulumi:"frequency"`
	Interval  int     `pulumi:"interval"`
}





type ConnectorMappingAvailabilityResponseInput interface {
	pulumi.Input

	ToConnectorMappingAvailabilityResponseOutput() ConnectorMappingAvailabilityResponseOutput
	ToConnectorMappingAvailabilityResponseOutputWithContext(context.Context) ConnectorMappingAvailabilityResponseOutput
}

type ConnectorMappingAvailabilityResponseArgs struct {
	Frequency pulumi.StringPtrInput `pulumi:"frequency"`
	Interval  pulumi.IntInput       `pulumi:"interval"`
}

func (ConnectorMappingAvailabilityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingAvailabilityResponse)(nil)).Elem()
}

func (i ConnectorMappingAvailabilityResponseArgs) ToConnectorMappingAvailabilityResponseOutput() ConnectorMappingAvailabilityResponseOutput {
	return i.ToConnectorMappingAvailabilityResponseOutputWithContext(context.Background())
}

func (i ConnectorMappingAvailabilityResponseArgs) ToConnectorMappingAvailabilityResponseOutputWithContext(ctx context.Context) ConnectorMappingAvailabilityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingAvailabilityResponseOutput)
}

func (i ConnectorMappingAvailabilityResponseArgs) ToConnectorMappingAvailabilityResponsePtrOutput() ConnectorMappingAvailabilityResponsePtrOutput {
	return i.ToConnectorMappingAvailabilityResponsePtrOutputWithContext(context.Background())
}

func (i ConnectorMappingAvailabilityResponseArgs) ToConnectorMappingAvailabilityResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingAvailabilityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingAvailabilityResponseOutput).ToConnectorMappingAvailabilityResponsePtrOutputWithContext(ctx)
}









type ConnectorMappingAvailabilityResponsePtrInput interface {
	pulumi.Input

	ToConnectorMappingAvailabilityResponsePtrOutput() ConnectorMappingAvailabilityResponsePtrOutput
	ToConnectorMappingAvailabilityResponsePtrOutputWithContext(context.Context) ConnectorMappingAvailabilityResponsePtrOutput
}

type connectorMappingAvailabilityResponsePtrType ConnectorMappingAvailabilityResponseArgs

func ConnectorMappingAvailabilityResponsePtr(v *ConnectorMappingAvailabilityResponseArgs) ConnectorMappingAvailabilityResponsePtrInput {
	return (*connectorMappingAvailabilityResponsePtrType)(v)
}

func (*connectorMappingAvailabilityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingAvailabilityResponse)(nil)).Elem()
}

func (i *connectorMappingAvailabilityResponsePtrType) ToConnectorMappingAvailabilityResponsePtrOutput() ConnectorMappingAvailabilityResponsePtrOutput {
	return i.ToConnectorMappingAvailabilityResponsePtrOutputWithContext(context.Background())
}

func (i *connectorMappingAvailabilityResponsePtrType) ToConnectorMappingAvailabilityResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingAvailabilityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingAvailabilityResponsePtrOutput)
}

type ConnectorMappingAvailabilityResponseOutput struct{ *pulumi.OutputState }

func (ConnectorMappingAvailabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingAvailabilityResponse)(nil)).Elem()
}

func (o ConnectorMappingAvailabilityResponseOutput) ToConnectorMappingAvailabilityResponseOutput() ConnectorMappingAvailabilityResponseOutput {
	return o
}

func (o ConnectorMappingAvailabilityResponseOutput) ToConnectorMappingAvailabilityResponseOutputWithContext(ctx context.Context) ConnectorMappingAvailabilityResponseOutput {
	return o
}

func (o ConnectorMappingAvailabilityResponseOutput) ToConnectorMappingAvailabilityResponsePtrOutput() ConnectorMappingAvailabilityResponsePtrOutput {
	return o.ToConnectorMappingAvailabilityResponsePtrOutputWithContext(context.Background())
}

func (o ConnectorMappingAvailabilityResponseOutput) ToConnectorMappingAvailabilityResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingAvailabilityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectorMappingAvailabilityResponse) *ConnectorMappingAvailabilityResponse {
		return &v
	}).(ConnectorMappingAvailabilityResponsePtrOutput)
}

func (o ConnectorMappingAvailabilityResponseOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingAvailabilityResponse) *string { return v.Frequency }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingAvailabilityResponseOutput) Interval() pulumi.IntOutput {
	return o.ApplyT(func(v ConnectorMappingAvailabilityResponse) int { return v.Interval }).(pulumi.IntOutput)
}

type ConnectorMappingAvailabilityResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectorMappingAvailabilityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingAvailabilityResponse)(nil)).Elem()
}

func (o ConnectorMappingAvailabilityResponsePtrOutput) ToConnectorMappingAvailabilityResponsePtrOutput() ConnectorMappingAvailabilityResponsePtrOutput {
	return o
}

func (o ConnectorMappingAvailabilityResponsePtrOutput) ToConnectorMappingAvailabilityResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingAvailabilityResponsePtrOutput {
	return o
}

func (o ConnectorMappingAvailabilityResponsePtrOutput) Elem() ConnectorMappingAvailabilityResponseOutput {
	return o.ApplyT(func(v *ConnectorMappingAvailabilityResponse) ConnectorMappingAvailabilityResponse {
		if v != nil {
			return *v
		}
		var ret ConnectorMappingAvailabilityResponse
		return ret
	}).(ConnectorMappingAvailabilityResponseOutput)
}

func (o ConnectorMappingAvailabilityResponsePtrOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingAvailabilityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Frequency
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingAvailabilityResponsePtrOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingAvailabilityResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Interval
	}).(pulumi.IntPtrOutput)
}

type ConnectorMappingCompleteOperation struct {
	CompletionOperationType *CompletionOperationTypes `pulumi:"completionOperationType"`
	DestinationFolder       *string                   `pulumi:"destinationFolder"`
}





type ConnectorMappingCompleteOperationInput interface {
	pulumi.Input

	ToConnectorMappingCompleteOperationOutput() ConnectorMappingCompleteOperationOutput
	ToConnectorMappingCompleteOperationOutputWithContext(context.Context) ConnectorMappingCompleteOperationOutput
}

type ConnectorMappingCompleteOperationArgs struct {
	CompletionOperationType CompletionOperationTypesPtrInput `pulumi:"completionOperationType"`
	DestinationFolder       pulumi.StringPtrInput            `pulumi:"destinationFolder"`
}

func (ConnectorMappingCompleteOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingCompleteOperation)(nil)).Elem()
}

func (i ConnectorMappingCompleteOperationArgs) ToConnectorMappingCompleteOperationOutput() ConnectorMappingCompleteOperationOutput {
	return i.ToConnectorMappingCompleteOperationOutputWithContext(context.Background())
}

func (i ConnectorMappingCompleteOperationArgs) ToConnectorMappingCompleteOperationOutputWithContext(ctx context.Context) ConnectorMappingCompleteOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingCompleteOperationOutput)
}

func (i ConnectorMappingCompleteOperationArgs) ToConnectorMappingCompleteOperationPtrOutput() ConnectorMappingCompleteOperationPtrOutput {
	return i.ToConnectorMappingCompleteOperationPtrOutputWithContext(context.Background())
}

func (i ConnectorMappingCompleteOperationArgs) ToConnectorMappingCompleteOperationPtrOutputWithContext(ctx context.Context) ConnectorMappingCompleteOperationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingCompleteOperationOutput).ToConnectorMappingCompleteOperationPtrOutputWithContext(ctx)
}









type ConnectorMappingCompleteOperationPtrInput interface {
	pulumi.Input

	ToConnectorMappingCompleteOperationPtrOutput() ConnectorMappingCompleteOperationPtrOutput
	ToConnectorMappingCompleteOperationPtrOutputWithContext(context.Context) ConnectorMappingCompleteOperationPtrOutput
}

type connectorMappingCompleteOperationPtrType ConnectorMappingCompleteOperationArgs

func ConnectorMappingCompleteOperationPtr(v *ConnectorMappingCompleteOperationArgs) ConnectorMappingCompleteOperationPtrInput {
	return (*connectorMappingCompleteOperationPtrType)(v)
}

func (*connectorMappingCompleteOperationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingCompleteOperation)(nil)).Elem()
}

func (i *connectorMappingCompleteOperationPtrType) ToConnectorMappingCompleteOperationPtrOutput() ConnectorMappingCompleteOperationPtrOutput {
	return i.ToConnectorMappingCompleteOperationPtrOutputWithContext(context.Background())
}

func (i *connectorMappingCompleteOperationPtrType) ToConnectorMappingCompleteOperationPtrOutputWithContext(ctx context.Context) ConnectorMappingCompleteOperationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingCompleteOperationPtrOutput)
}

type ConnectorMappingCompleteOperationOutput struct{ *pulumi.OutputState }

func (ConnectorMappingCompleteOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingCompleteOperation)(nil)).Elem()
}

func (o ConnectorMappingCompleteOperationOutput) ToConnectorMappingCompleteOperationOutput() ConnectorMappingCompleteOperationOutput {
	return o
}

func (o ConnectorMappingCompleteOperationOutput) ToConnectorMappingCompleteOperationOutputWithContext(ctx context.Context) ConnectorMappingCompleteOperationOutput {
	return o
}

func (o ConnectorMappingCompleteOperationOutput) ToConnectorMappingCompleteOperationPtrOutput() ConnectorMappingCompleteOperationPtrOutput {
	return o.ToConnectorMappingCompleteOperationPtrOutputWithContext(context.Background())
}

func (o ConnectorMappingCompleteOperationOutput) ToConnectorMappingCompleteOperationPtrOutputWithContext(ctx context.Context) ConnectorMappingCompleteOperationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectorMappingCompleteOperation) *ConnectorMappingCompleteOperation {
		return &v
	}).(ConnectorMappingCompleteOperationPtrOutput)
}

func (o ConnectorMappingCompleteOperationOutput) CompletionOperationType() CompletionOperationTypesPtrOutput {
	return o.ApplyT(func(v ConnectorMappingCompleteOperation) *CompletionOperationTypes { return v.CompletionOperationType }).(CompletionOperationTypesPtrOutput)
}

func (o ConnectorMappingCompleteOperationOutput) DestinationFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingCompleteOperation) *string { return v.DestinationFolder }).(pulumi.StringPtrOutput)
}

type ConnectorMappingCompleteOperationPtrOutput struct{ *pulumi.OutputState }

func (ConnectorMappingCompleteOperationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingCompleteOperation)(nil)).Elem()
}

func (o ConnectorMappingCompleteOperationPtrOutput) ToConnectorMappingCompleteOperationPtrOutput() ConnectorMappingCompleteOperationPtrOutput {
	return o
}

func (o ConnectorMappingCompleteOperationPtrOutput) ToConnectorMappingCompleteOperationPtrOutputWithContext(ctx context.Context) ConnectorMappingCompleteOperationPtrOutput {
	return o
}

func (o ConnectorMappingCompleteOperationPtrOutput) Elem() ConnectorMappingCompleteOperationOutput {
	return o.ApplyT(func(v *ConnectorMappingCompleteOperation) ConnectorMappingCompleteOperation {
		if v != nil {
			return *v
		}
		var ret ConnectorMappingCompleteOperation
		return ret
	}).(ConnectorMappingCompleteOperationOutput)
}

func (o ConnectorMappingCompleteOperationPtrOutput) CompletionOperationType() CompletionOperationTypesPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingCompleteOperation) *CompletionOperationTypes {
		if v == nil {
			return nil
		}
		return v.CompletionOperationType
	}).(CompletionOperationTypesPtrOutput)
}

func (o ConnectorMappingCompleteOperationPtrOutput) DestinationFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingCompleteOperation) *string {
		if v == nil {
			return nil
		}
		return v.DestinationFolder
	}).(pulumi.StringPtrOutput)
}

type ConnectorMappingCompleteOperationResponse struct {
	CompletionOperationType *string `pulumi:"completionOperationType"`
	DestinationFolder       *string `pulumi:"destinationFolder"`
}





type ConnectorMappingCompleteOperationResponseInput interface {
	pulumi.Input

	ToConnectorMappingCompleteOperationResponseOutput() ConnectorMappingCompleteOperationResponseOutput
	ToConnectorMappingCompleteOperationResponseOutputWithContext(context.Context) ConnectorMappingCompleteOperationResponseOutput
}

type ConnectorMappingCompleteOperationResponseArgs struct {
	CompletionOperationType pulumi.StringPtrInput `pulumi:"completionOperationType"`
	DestinationFolder       pulumi.StringPtrInput `pulumi:"destinationFolder"`
}

func (ConnectorMappingCompleteOperationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingCompleteOperationResponse)(nil)).Elem()
}

func (i ConnectorMappingCompleteOperationResponseArgs) ToConnectorMappingCompleteOperationResponseOutput() ConnectorMappingCompleteOperationResponseOutput {
	return i.ToConnectorMappingCompleteOperationResponseOutputWithContext(context.Background())
}

func (i ConnectorMappingCompleteOperationResponseArgs) ToConnectorMappingCompleteOperationResponseOutputWithContext(ctx context.Context) ConnectorMappingCompleteOperationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingCompleteOperationResponseOutput)
}

func (i ConnectorMappingCompleteOperationResponseArgs) ToConnectorMappingCompleteOperationResponsePtrOutput() ConnectorMappingCompleteOperationResponsePtrOutput {
	return i.ToConnectorMappingCompleteOperationResponsePtrOutputWithContext(context.Background())
}

func (i ConnectorMappingCompleteOperationResponseArgs) ToConnectorMappingCompleteOperationResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingCompleteOperationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingCompleteOperationResponseOutput).ToConnectorMappingCompleteOperationResponsePtrOutputWithContext(ctx)
}









type ConnectorMappingCompleteOperationResponsePtrInput interface {
	pulumi.Input

	ToConnectorMappingCompleteOperationResponsePtrOutput() ConnectorMappingCompleteOperationResponsePtrOutput
	ToConnectorMappingCompleteOperationResponsePtrOutputWithContext(context.Context) ConnectorMappingCompleteOperationResponsePtrOutput
}

type connectorMappingCompleteOperationResponsePtrType ConnectorMappingCompleteOperationResponseArgs

func ConnectorMappingCompleteOperationResponsePtr(v *ConnectorMappingCompleteOperationResponseArgs) ConnectorMappingCompleteOperationResponsePtrInput {
	return (*connectorMappingCompleteOperationResponsePtrType)(v)
}

func (*connectorMappingCompleteOperationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingCompleteOperationResponse)(nil)).Elem()
}

func (i *connectorMappingCompleteOperationResponsePtrType) ToConnectorMappingCompleteOperationResponsePtrOutput() ConnectorMappingCompleteOperationResponsePtrOutput {
	return i.ToConnectorMappingCompleteOperationResponsePtrOutputWithContext(context.Background())
}

func (i *connectorMappingCompleteOperationResponsePtrType) ToConnectorMappingCompleteOperationResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingCompleteOperationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingCompleteOperationResponsePtrOutput)
}

type ConnectorMappingCompleteOperationResponseOutput struct{ *pulumi.OutputState }

func (ConnectorMappingCompleteOperationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingCompleteOperationResponse)(nil)).Elem()
}

func (o ConnectorMappingCompleteOperationResponseOutput) ToConnectorMappingCompleteOperationResponseOutput() ConnectorMappingCompleteOperationResponseOutput {
	return o
}

func (o ConnectorMappingCompleteOperationResponseOutput) ToConnectorMappingCompleteOperationResponseOutputWithContext(ctx context.Context) ConnectorMappingCompleteOperationResponseOutput {
	return o
}

func (o ConnectorMappingCompleteOperationResponseOutput) ToConnectorMappingCompleteOperationResponsePtrOutput() ConnectorMappingCompleteOperationResponsePtrOutput {
	return o.ToConnectorMappingCompleteOperationResponsePtrOutputWithContext(context.Background())
}

func (o ConnectorMappingCompleteOperationResponseOutput) ToConnectorMappingCompleteOperationResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingCompleteOperationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectorMappingCompleteOperationResponse) *ConnectorMappingCompleteOperationResponse {
		return &v
	}).(ConnectorMappingCompleteOperationResponsePtrOutput)
}

func (o ConnectorMappingCompleteOperationResponseOutput) CompletionOperationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingCompleteOperationResponse) *string { return v.CompletionOperationType }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingCompleteOperationResponseOutput) DestinationFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingCompleteOperationResponse) *string { return v.DestinationFolder }).(pulumi.StringPtrOutput)
}

type ConnectorMappingCompleteOperationResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectorMappingCompleteOperationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingCompleteOperationResponse)(nil)).Elem()
}

func (o ConnectorMappingCompleteOperationResponsePtrOutput) ToConnectorMappingCompleteOperationResponsePtrOutput() ConnectorMappingCompleteOperationResponsePtrOutput {
	return o
}

func (o ConnectorMappingCompleteOperationResponsePtrOutput) ToConnectorMappingCompleteOperationResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingCompleteOperationResponsePtrOutput {
	return o
}

func (o ConnectorMappingCompleteOperationResponsePtrOutput) Elem() ConnectorMappingCompleteOperationResponseOutput {
	return o.ApplyT(func(v *ConnectorMappingCompleteOperationResponse) ConnectorMappingCompleteOperationResponse {
		if v != nil {
			return *v
		}
		var ret ConnectorMappingCompleteOperationResponse
		return ret
	}).(ConnectorMappingCompleteOperationResponseOutput)
}

func (o ConnectorMappingCompleteOperationResponsePtrOutput) CompletionOperationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingCompleteOperationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CompletionOperationType
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingCompleteOperationResponsePtrOutput) DestinationFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingCompleteOperationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DestinationFolder
	}).(pulumi.StringPtrOutput)
}

type ConnectorMappingErrorManagement struct {
	ErrorLimit          *int                 `pulumi:"errorLimit"`
	ErrorManagementType ErrorManagementTypes `pulumi:"errorManagementType"`
}





type ConnectorMappingErrorManagementInput interface {
	pulumi.Input

	ToConnectorMappingErrorManagementOutput() ConnectorMappingErrorManagementOutput
	ToConnectorMappingErrorManagementOutputWithContext(context.Context) ConnectorMappingErrorManagementOutput
}

type ConnectorMappingErrorManagementArgs struct {
	ErrorLimit          pulumi.IntPtrInput        `pulumi:"errorLimit"`
	ErrorManagementType ErrorManagementTypesInput `pulumi:"errorManagementType"`
}

func (ConnectorMappingErrorManagementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingErrorManagement)(nil)).Elem()
}

func (i ConnectorMappingErrorManagementArgs) ToConnectorMappingErrorManagementOutput() ConnectorMappingErrorManagementOutput {
	return i.ToConnectorMappingErrorManagementOutputWithContext(context.Background())
}

func (i ConnectorMappingErrorManagementArgs) ToConnectorMappingErrorManagementOutputWithContext(ctx context.Context) ConnectorMappingErrorManagementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingErrorManagementOutput)
}

func (i ConnectorMappingErrorManagementArgs) ToConnectorMappingErrorManagementPtrOutput() ConnectorMappingErrorManagementPtrOutput {
	return i.ToConnectorMappingErrorManagementPtrOutputWithContext(context.Background())
}

func (i ConnectorMappingErrorManagementArgs) ToConnectorMappingErrorManagementPtrOutputWithContext(ctx context.Context) ConnectorMappingErrorManagementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingErrorManagementOutput).ToConnectorMappingErrorManagementPtrOutputWithContext(ctx)
}









type ConnectorMappingErrorManagementPtrInput interface {
	pulumi.Input

	ToConnectorMappingErrorManagementPtrOutput() ConnectorMappingErrorManagementPtrOutput
	ToConnectorMappingErrorManagementPtrOutputWithContext(context.Context) ConnectorMappingErrorManagementPtrOutput
}

type connectorMappingErrorManagementPtrType ConnectorMappingErrorManagementArgs

func ConnectorMappingErrorManagementPtr(v *ConnectorMappingErrorManagementArgs) ConnectorMappingErrorManagementPtrInput {
	return (*connectorMappingErrorManagementPtrType)(v)
}

func (*connectorMappingErrorManagementPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingErrorManagement)(nil)).Elem()
}

func (i *connectorMappingErrorManagementPtrType) ToConnectorMappingErrorManagementPtrOutput() ConnectorMappingErrorManagementPtrOutput {
	return i.ToConnectorMappingErrorManagementPtrOutputWithContext(context.Background())
}

func (i *connectorMappingErrorManagementPtrType) ToConnectorMappingErrorManagementPtrOutputWithContext(ctx context.Context) ConnectorMappingErrorManagementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingErrorManagementPtrOutput)
}

type ConnectorMappingErrorManagementOutput struct{ *pulumi.OutputState }

func (ConnectorMappingErrorManagementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingErrorManagement)(nil)).Elem()
}

func (o ConnectorMappingErrorManagementOutput) ToConnectorMappingErrorManagementOutput() ConnectorMappingErrorManagementOutput {
	return o
}

func (o ConnectorMappingErrorManagementOutput) ToConnectorMappingErrorManagementOutputWithContext(ctx context.Context) ConnectorMappingErrorManagementOutput {
	return o
}

func (o ConnectorMappingErrorManagementOutput) ToConnectorMappingErrorManagementPtrOutput() ConnectorMappingErrorManagementPtrOutput {
	return o.ToConnectorMappingErrorManagementPtrOutputWithContext(context.Background())
}

func (o ConnectorMappingErrorManagementOutput) ToConnectorMappingErrorManagementPtrOutputWithContext(ctx context.Context) ConnectorMappingErrorManagementPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectorMappingErrorManagement) *ConnectorMappingErrorManagement {
		return &v
	}).(ConnectorMappingErrorManagementPtrOutput)
}

func (o ConnectorMappingErrorManagementOutput) ErrorLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConnectorMappingErrorManagement) *int { return v.ErrorLimit }).(pulumi.IntPtrOutput)
}

func (o ConnectorMappingErrorManagementOutput) ErrorManagementType() ErrorManagementTypesOutput {
	return o.ApplyT(func(v ConnectorMappingErrorManagement) ErrorManagementTypes { return v.ErrorManagementType }).(ErrorManagementTypesOutput)
}

type ConnectorMappingErrorManagementPtrOutput struct{ *pulumi.OutputState }

func (ConnectorMappingErrorManagementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingErrorManagement)(nil)).Elem()
}

func (o ConnectorMappingErrorManagementPtrOutput) ToConnectorMappingErrorManagementPtrOutput() ConnectorMappingErrorManagementPtrOutput {
	return o
}

func (o ConnectorMappingErrorManagementPtrOutput) ToConnectorMappingErrorManagementPtrOutputWithContext(ctx context.Context) ConnectorMappingErrorManagementPtrOutput {
	return o
}

func (o ConnectorMappingErrorManagementPtrOutput) Elem() ConnectorMappingErrorManagementOutput {
	return o.ApplyT(func(v *ConnectorMappingErrorManagement) ConnectorMappingErrorManagement {
		if v != nil {
			return *v
		}
		var ret ConnectorMappingErrorManagement
		return ret
	}).(ConnectorMappingErrorManagementOutput)
}

func (o ConnectorMappingErrorManagementPtrOutput) ErrorLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingErrorManagement) *int {
		if v == nil {
			return nil
		}
		return v.ErrorLimit
	}).(pulumi.IntPtrOutput)
}

func (o ConnectorMappingErrorManagementPtrOutput) ErrorManagementType() ErrorManagementTypesPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingErrorManagement) *ErrorManagementTypes {
		if v == nil {
			return nil
		}
		return &v.ErrorManagementType
	}).(ErrorManagementTypesPtrOutput)
}

type ConnectorMappingErrorManagementResponse struct {
	ErrorLimit          *int   `pulumi:"errorLimit"`
	ErrorManagementType string `pulumi:"errorManagementType"`
}





type ConnectorMappingErrorManagementResponseInput interface {
	pulumi.Input

	ToConnectorMappingErrorManagementResponseOutput() ConnectorMappingErrorManagementResponseOutput
	ToConnectorMappingErrorManagementResponseOutputWithContext(context.Context) ConnectorMappingErrorManagementResponseOutput
}

type ConnectorMappingErrorManagementResponseArgs struct {
	ErrorLimit          pulumi.IntPtrInput `pulumi:"errorLimit"`
	ErrorManagementType pulumi.StringInput `pulumi:"errorManagementType"`
}

func (ConnectorMappingErrorManagementResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingErrorManagementResponse)(nil)).Elem()
}

func (i ConnectorMappingErrorManagementResponseArgs) ToConnectorMappingErrorManagementResponseOutput() ConnectorMappingErrorManagementResponseOutput {
	return i.ToConnectorMappingErrorManagementResponseOutputWithContext(context.Background())
}

func (i ConnectorMappingErrorManagementResponseArgs) ToConnectorMappingErrorManagementResponseOutputWithContext(ctx context.Context) ConnectorMappingErrorManagementResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingErrorManagementResponseOutput)
}

func (i ConnectorMappingErrorManagementResponseArgs) ToConnectorMappingErrorManagementResponsePtrOutput() ConnectorMappingErrorManagementResponsePtrOutput {
	return i.ToConnectorMappingErrorManagementResponsePtrOutputWithContext(context.Background())
}

func (i ConnectorMappingErrorManagementResponseArgs) ToConnectorMappingErrorManagementResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingErrorManagementResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingErrorManagementResponseOutput).ToConnectorMappingErrorManagementResponsePtrOutputWithContext(ctx)
}









type ConnectorMappingErrorManagementResponsePtrInput interface {
	pulumi.Input

	ToConnectorMappingErrorManagementResponsePtrOutput() ConnectorMappingErrorManagementResponsePtrOutput
	ToConnectorMappingErrorManagementResponsePtrOutputWithContext(context.Context) ConnectorMappingErrorManagementResponsePtrOutput
}

type connectorMappingErrorManagementResponsePtrType ConnectorMappingErrorManagementResponseArgs

func ConnectorMappingErrorManagementResponsePtr(v *ConnectorMappingErrorManagementResponseArgs) ConnectorMappingErrorManagementResponsePtrInput {
	return (*connectorMappingErrorManagementResponsePtrType)(v)
}

func (*connectorMappingErrorManagementResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingErrorManagementResponse)(nil)).Elem()
}

func (i *connectorMappingErrorManagementResponsePtrType) ToConnectorMappingErrorManagementResponsePtrOutput() ConnectorMappingErrorManagementResponsePtrOutput {
	return i.ToConnectorMappingErrorManagementResponsePtrOutputWithContext(context.Background())
}

func (i *connectorMappingErrorManagementResponsePtrType) ToConnectorMappingErrorManagementResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingErrorManagementResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingErrorManagementResponsePtrOutput)
}

type ConnectorMappingErrorManagementResponseOutput struct{ *pulumi.OutputState }

func (ConnectorMappingErrorManagementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingErrorManagementResponse)(nil)).Elem()
}

func (o ConnectorMappingErrorManagementResponseOutput) ToConnectorMappingErrorManagementResponseOutput() ConnectorMappingErrorManagementResponseOutput {
	return o
}

func (o ConnectorMappingErrorManagementResponseOutput) ToConnectorMappingErrorManagementResponseOutputWithContext(ctx context.Context) ConnectorMappingErrorManagementResponseOutput {
	return o
}

func (o ConnectorMappingErrorManagementResponseOutput) ToConnectorMappingErrorManagementResponsePtrOutput() ConnectorMappingErrorManagementResponsePtrOutput {
	return o.ToConnectorMappingErrorManagementResponsePtrOutputWithContext(context.Background())
}

func (o ConnectorMappingErrorManagementResponseOutput) ToConnectorMappingErrorManagementResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingErrorManagementResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectorMappingErrorManagementResponse) *ConnectorMappingErrorManagementResponse {
		return &v
	}).(ConnectorMappingErrorManagementResponsePtrOutput)
}

func (o ConnectorMappingErrorManagementResponseOutput) ErrorLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConnectorMappingErrorManagementResponse) *int { return v.ErrorLimit }).(pulumi.IntPtrOutput)
}

func (o ConnectorMappingErrorManagementResponseOutput) ErrorManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorMappingErrorManagementResponse) string { return v.ErrorManagementType }).(pulumi.StringOutput)
}

type ConnectorMappingErrorManagementResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectorMappingErrorManagementResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingErrorManagementResponse)(nil)).Elem()
}

func (o ConnectorMappingErrorManagementResponsePtrOutput) ToConnectorMappingErrorManagementResponsePtrOutput() ConnectorMappingErrorManagementResponsePtrOutput {
	return o
}

func (o ConnectorMappingErrorManagementResponsePtrOutput) ToConnectorMappingErrorManagementResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingErrorManagementResponsePtrOutput {
	return o
}

func (o ConnectorMappingErrorManagementResponsePtrOutput) Elem() ConnectorMappingErrorManagementResponseOutput {
	return o.ApplyT(func(v *ConnectorMappingErrorManagementResponse) ConnectorMappingErrorManagementResponse {
		if v != nil {
			return *v
		}
		var ret ConnectorMappingErrorManagementResponse
		return ret
	}).(ConnectorMappingErrorManagementResponseOutput)
}

func (o ConnectorMappingErrorManagementResponsePtrOutput) ErrorLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingErrorManagementResponse) *int {
		if v == nil {
			return nil
		}
		return v.ErrorLimit
	}).(pulumi.IntPtrOutput)
}

func (o ConnectorMappingErrorManagementResponsePtrOutput) ErrorManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingErrorManagementResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorManagementType
	}).(pulumi.StringPtrOutput)
}

type ConnectorMappingFormat struct {
	AcceptLanguage       *string     `pulumi:"acceptLanguage"`
	ArraySeparator       *string     `pulumi:"arraySeparator"`
	ColumnDelimiter      *string     `pulumi:"columnDelimiter"`
	FormatType           FormatTypes `pulumi:"formatType"`
	QuoteCharacter       *string     `pulumi:"quoteCharacter"`
	QuoteEscapeCharacter *string     `pulumi:"quoteEscapeCharacter"`
}





type ConnectorMappingFormatInput interface {
	pulumi.Input

	ToConnectorMappingFormatOutput() ConnectorMappingFormatOutput
	ToConnectorMappingFormatOutputWithContext(context.Context) ConnectorMappingFormatOutput
}

type ConnectorMappingFormatArgs struct {
	AcceptLanguage       pulumi.StringPtrInput `pulumi:"acceptLanguage"`
	ArraySeparator       pulumi.StringPtrInput `pulumi:"arraySeparator"`
	ColumnDelimiter      pulumi.StringPtrInput `pulumi:"columnDelimiter"`
	FormatType           FormatTypesInput      `pulumi:"formatType"`
	QuoteCharacter       pulumi.StringPtrInput `pulumi:"quoteCharacter"`
	QuoteEscapeCharacter pulumi.StringPtrInput `pulumi:"quoteEscapeCharacter"`
}

func (ConnectorMappingFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingFormat)(nil)).Elem()
}

func (i ConnectorMappingFormatArgs) ToConnectorMappingFormatOutput() ConnectorMappingFormatOutput {
	return i.ToConnectorMappingFormatOutputWithContext(context.Background())
}

func (i ConnectorMappingFormatArgs) ToConnectorMappingFormatOutputWithContext(ctx context.Context) ConnectorMappingFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingFormatOutput)
}

func (i ConnectorMappingFormatArgs) ToConnectorMappingFormatPtrOutput() ConnectorMappingFormatPtrOutput {
	return i.ToConnectorMappingFormatPtrOutputWithContext(context.Background())
}

func (i ConnectorMappingFormatArgs) ToConnectorMappingFormatPtrOutputWithContext(ctx context.Context) ConnectorMappingFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingFormatOutput).ToConnectorMappingFormatPtrOutputWithContext(ctx)
}









type ConnectorMappingFormatPtrInput interface {
	pulumi.Input

	ToConnectorMappingFormatPtrOutput() ConnectorMappingFormatPtrOutput
	ToConnectorMappingFormatPtrOutputWithContext(context.Context) ConnectorMappingFormatPtrOutput
}

type connectorMappingFormatPtrType ConnectorMappingFormatArgs

func ConnectorMappingFormatPtr(v *ConnectorMappingFormatArgs) ConnectorMappingFormatPtrInput {
	return (*connectorMappingFormatPtrType)(v)
}

func (*connectorMappingFormatPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingFormat)(nil)).Elem()
}

func (i *connectorMappingFormatPtrType) ToConnectorMappingFormatPtrOutput() ConnectorMappingFormatPtrOutput {
	return i.ToConnectorMappingFormatPtrOutputWithContext(context.Background())
}

func (i *connectorMappingFormatPtrType) ToConnectorMappingFormatPtrOutputWithContext(ctx context.Context) ConnectorMappingFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingFormatPtrOutput)
}

type ConnectorMappingFormatOutput struct{ *pulumi.OutputState }

func (ConnectorMappingFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingFormat)(nil)).Elem()
}

func (o ConnectorMappingFormatOutput) ToConnectorMappingFormatOutput() ConnectorMappingFormatOutput {
	return o
}

func (o ConnectorMappingFormatOutput) ToConnectorMappingFormatOutputWithContext(ctx context.Context) ConnectorMappingFormatOutput {
	return o
}

func (o ConnectorMappingFormatOutput) ToConnectorMappingFormatPtrOutput() ConnectorMappingFormatPtrOutput {
	return o.ToConnectorMappingFormatPtrOutputWithContext(context.Background())
}

func (o ConnectorMappingFormatOutput) ToConnectorMappingFormatPtrOutputWithContext(ctx context.Context) ConnectorMappingFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectorMappingFormat) *ConnectorMappingFormat {
		return &v
	}).(ConnectorMappingFormatPtrOutput)
}

func (o ConnectorMappingFormatOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingFormat) *string { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatOutput) ArraySeparator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingFormat) *string { return v.ArraySeparator }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatOutput) ColumnDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingFormat) *string { return v.ColumnDelimiter }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatOutput) FormatType() FormatTypesOutput {
	return o.ApplyT(func(v ConnectorMappingFormat) FormatTypes { return v.FormatType }).(FormatTypesOutput)
}

func (o ConnectorMappingFormatOutput) QuoteCharacter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingFormat) *string { return v.QuoteCharacter }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatOutput) QuoteEscapeCharacter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingFormat) *string { return v.QuoteEscapeCharacter }).(pulumi.StringPtrOutput)
}

type ConnectorMappingFormatPtrOutput struct{ *pulumi.OutputState }

func (ConnectorMappingFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingFormat)(nil)).Elem()
}

func (o ConnectorMappingFormatPtrOutput) ToConnectorMappingFormatPtrOutput() ConnectorMappingFormatPtrOutput {
	return o
}

func (o ConnectorMappingFormatPtrOutput) ToConnectorMappingFormatPtrOutputWithContext(ctx context.Context) ConnectorMappingFormatPtrOutput {
	return o
}

func (o ConnectorMappingFormatPtrOutput) Elem() ConnectorMappingFormatOutput {
	return o.ApplyT(func(v *ConnectorMappingFormat) ConnectorMappingFormat {
		if v != nil {
			return *v
		}
		var ret ConnectorMappingFormat
		return ret
	}).(ConnectorMappingFormatOutput)
}

func (o ConnectorMappingFormatPtrOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingFormat) *string {
		if v == nil {
			return nil
		}
		return v.AcceptLanguage
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatPtrOutput) ArraySeparator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingFormat) *string {
		if v == nil {
			return nil
		}
		return v.ArraySeparator
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatPtrOutput) ColumnDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingFormat) *string {
		if v == nil {
			return nil
		}
		return v.ColumnDelimiter
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatPtrOutput) FormatType() FormatTypesPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingFormat) *FormatTypes {
		if v == nil {
			return nil
		}
		return &v.FormatType
	}).(FormatTypesPtrOutput)
}

func (o ConnectorMappingFormatPtrOutput) QuoteCharacter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingFormat) *string {
		if v == nil {
			return nil
		}
		return v.QuoteCharacter
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatPtrOutput) QuoteEscapeCharacter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingFormat) *string {
		if v == nil {
			return nil
		}
		return v.QuoteEscapeCharacter
	}).(pulumi.StringPtrOutput)
}

type ConnectorMappingFormatResponse struct {
	AcceptLanguage       *string `pulumi:"acceptLanguage"`
	ArraySeparator       *string `pulumi:"arraySeparator"`
	ColumnDelimiter      *string `pulumi:"columnDelimiter"`
	FormatType           string  `pulumi:"formatType"`
	QuoteCharacter       *string `pulumi:"quoteCharacter"`
	QuoteEscapeCharacter *string `pulumi:"quoteEscapeCharacter"`
}





type ConnectorMappingFormatResponseInput interface {
	pulumi.Input

	ToConnectorMappingFormatResponseOutput() ConnectorMappingFormatResponseOutput
	ToConnectorMappingFormatResponseOutputWithContext(context.Context) ConnectorMappingFormatResponseOutput
}

type ConnectorMappingFormatResponseArgs struct {
	AcceptLanguage       pulumi.StringPtrInput `pulumi:"acceptLanguage"`
	ArraySeparator       pulumi.StringPtrInput `pulumi:"arraySeparator"`
	ColumnDelimiter      pulumi.StringPtrInput `pulumi:"columnDelimiter"`
	FormatType           pulumi.StringInput    `pulumi:"formatType"`
	QuoteCharacter       pulumi.StringPtrInput `pulumi:"quoteCharacter"`
	QuoteEscapeCharacter pulumi.StringPtrInput `pulumi:"quoteEscapeCharacter"`
}

func (ConnectorMappingFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingFormatResponse)(nil)).Elem()
}

func (i ConnectorMappingFormatResponseArgs) ToConnectorMappingFormatResponseOutput() ConnectorMappingFormatResponseOutput {
	return i.ToConnectorMappingFormatResponseOutputWithContext(context.Background())
}

func (i ConnectorMappingFormatResponseArgs) ToConnectorMappingFormatResponseOutputWithContext(ctx context.Context) ConnectorMappingFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingFormatResponseOutput)
}

func (i ConnectorMappingFormatResponseArgs) ToConnectorMappingFormatResponsePtrOutput() ConnectorMappingFormatResponsePtrOutput {
	return i.ToConnectorMappingFormatResponsePtrOutputWithContext(context.Background())
}

func (i ConnectorMappingFormatResponseArgs) ToConnectorMappingFormatResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingFormatResponseOutput).ToConnectorMappingFormatResponsePtrOutputWithContext(ctx)
}









type ConnectorMappingFormatResponsePtrInput interface {
	pulumi.Input

	ToConnectorMappingFormatResponsePtrOutput() ConnectorMappingFormatResponsePtrOutput
	ToConnectorMappingFormatResponsePtrOutputWithContext(context.Context) ConnectorMappingFormatResponsePtrOutput
}

type connectorMappingFormatResponsePtrType ConnectorMappingFormatResponseArgs

func ConnectorMappingFormatResponsePtr(v *ConnectorMappingFormatResponseArgs) ConnectorMappingFormatResponsePtrInput {
	return (*connectorMappingFormatResponsePtrType)(v)
}

func (*connectorMappingFormatResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingFormatResponse)(nil)).Elem()
}

func (i *connectorMappingFormatResponsePtrType) ToConnectorMappingFormatResponsePtrOutput() ConnectorMappingFormatResponsePtrOutput {
	return i.ToConnectorMappingFormatResponsePtrOutputWithContext(context.Background())
}

func (i *connectorMappingFormatResponsePtrType) ToConnectorMappingFormatResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingFormatResponsePtrOutput)
}

type ConnectorMappingFormatResponseOutput struct{ *pulumi.OutputState }

func (ConnectorMappingFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingFormatResponse)(nil)).Elem()
}

func (o ConnectorMappingFormatResponseOutput) ToConnectorMappingFormatResponseOutput() ConnectorMappingFormatResponseOutput {
	return o
}

func (o ConnectorMappingFormatResponseOutput) ToConnectorMappingFormatResponseOutputWithContext(ctx context.Context) ConnectorMappingFormatResponseOutput {
	return o
}

func (o ConnectorMappingFormatResponseOutput) ToConnectorMappingFormatResponsePtrOutput() ConnectorMappingFormatResponsePtrOutput {
	return o.ToConnectorMappingFormatResponsePtrOutputWithContext(context.Background())
}

func (o ConnectorMappingFormatResponseOutput) ToConnectorMappingFormatResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingFormatResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectorMappingFormatResponse) *ConnectorMappingFormatResponse {
		return &v
	}).(ConnectorMappingFormatResponsePtrOutput)
}

func (o ConnectorMappingFormatResponseOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingFormatResponse) *string { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatResponseOutput) ArraySeparator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingFormatResponse) *string { return v.ArraySeparator }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatResponseOutput) ColumnDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingFormatResponse) *string { return v.ColumnDelimiter }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatResponseOutput) FormatType() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorMappingFormatResponse) string { return v.FormatType }).(pulumi.StringOutput)
}

func (o ConnectorMappingFormatResponseOutput) QuoteCharacter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingFormatResponse) *string { return v.QuoteCharacter }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatResponseOutput) QuoteEscapeCharacter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingFormatResponse) *string { return v.QuoteEscapeCharacter }).(pulumi.StringPtrOutput)
}

type ConnectorMappingFormatResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectorMappingFormatResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingFormatResponse)(nil)).Elem()
}

func (o ConnectorMappingFormatResponsePtrOutput) ToConnectorMappingFormatResponsePtrOutput() ConnectorMappingFormatResponsePtrOutput {
	return o
}

func (o ConnectorMappingFormatResponsePtrOutput) ToConnectorMappingFormatResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingFormatResponsePtrOutput {
	return o
}

func (o ConnectorMappingFormatResponsePtrOutput) Elem() ConnectorMappingFormatResponseOutput {
	return o.ApplyT(func(v *ConnectorMappingFormatResponse) ConnectorMappingFormatResponse {
		if v != nil {
			return *v
		}
		var ret ConnectorMappingFormatResponse
		return ret
	}).(ConnectorMappingFormatResponseOutput)
}

func (o ConnectorMappingFormatResponsePtrOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingFormatResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcceptLanguage
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatResponsePtrOutput) ArraySeparator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingFormatResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArraySeparator
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatResponsePtrOutput) ColumnDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingFormatResponse) *string {
		if v == nil {
			return nil
		}
		return v.ColumnDelimiter
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatResponsePtrOutput) FormatType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingFormatResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FormatType
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatResponsePtrOutput) QuoteCharacter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingFormatResponse) *string {
		if v == nil {
			return nil
		}
		return v.QuoteCharacter
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingFormatResponsePtrOutput) QuoteEscapeCharacter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingFormatResponse) *string {
		if v == nil {
			return nil
		}
		return v.QuoteEscapeCharacter
	}).(pulumi.StringPtrOutput)
}

type ConnectorMappingProperties struct {
	Availability      ConnectorMappingAvailability      `pulumi:"availability"`
	CompleteOperation ConnectorMappingCompleteOperation `pulumi:"completeOperation"`
	ErrorManagement   ConnectorMappingErrorManagement   `pulumi:"errorManagement"`
	FileFilter        *string                           `pulumi:"fileFilter"`
	FolderPath        *string                           `pulumi:"folderPath"`
	Format            ConnectorMappingFormat            `pulumi:"format"`
	HasHeader         *bool                             `pulumi:"hasHeader"`
	Structure         []ConnectorMappingStructure       `pulumi:"structure"`
}





type ConnectorMappingPropertiesInput interface {
	pulumi.Input

	ToConnectorMappingPropertiesOutput() ConnectorMappingPropertiesOutput
	ToConnectorMappingPropertiesOutputWithContext(context.Context) ConnectorMappingPropertiesOutput
}

type ConnectorMappingPropertiesArgs struct {
	Availability      ConnectorMappingAvailabilityInput      `pulumi:"availability"`
	CompleteOperation ConnectorMappingCompleteOperationInput `pulumi:"completeOperation"`
	ErrorManagement   ConnectorMappingErrorManagementInput   `pulumi:"errorManagement"`
	FileFilter        pulumi.StringPtrInput                  `pulumi:"fileFilter"`
	FolderPath        pulumi.StringPtrInput                  `pulumi:"folderPath"`
	Format            ConnectorMappingFormatInput            `pulumi:"format"`
	HasHeader         pulumi.BoolPtrInput                    `pulumi:"hasHeader"`
	Structure         ConnectorMappingStructureArrayInput    `pulumi:"structure"`
}

func (ConnectorMappingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingProperties)(nil)).Elem()
}

func (i ConnectorMappingPropertiesArgs) ToConnectorMappingPropertiesOutput() ConnectorMappingPropertiesOutput {
	return i.ToConnectorMappingPropertiesOutputWithContext(context.Background())
}

func (i ConnectorMappingPropertiesArgs) ToConnectorMappingPropertiesOutputWithContext(ctx context.Context) ConnectorMappingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingPropertiesOutput)
}

func (i ConnectorMappingPropertiesArgs) ToConnectorMappingPropertiesPtrOutput() ConnectorMappingPropertiesPtrOutput {
	return i.ToConnectorMappingPropertiesPtrOutputWithContext(context.Background())
}

func (i ConnectorMappingPropertiesArgs) ToConnectorMappingPropertiesPtrOutputWithContext(ctx context.Context) ConnectorMappingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingPropertiesOutput).ToConnectorMappingPropertiesPtrOutputWithContext(ctx)
}









type ConnectorMappingPropertiesPtrInput interface {
	pulumi.Input

	ToConnectorMappingPropertiesPtrOutput() ConnectorMappingPropertiesPtrOutput
	ToConnectorMappingPropertiesPtrOutputWithContext(context.Context) ConnectorMappingPropertiesPtrOutput
}

type connectorMappingPropertiesPtrType ConnectorMappingPropertiesArgs

func ConnectorMappingPropertiesPtr(v *ConnectorMappingPropertiesArgs) ConnectorMappingPropertiesPtrInput {
	return (*connectorMappingPropertiesPtrType)(v)
}

func (*connectorMappingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingProperties)(nil)).Elem()
}

func (i *connectorMappingPropertiesPtrType) ToConnectorMappingPropertiesPtrOutput() ConnectorMappingPropertiesPtrOutput {
	return i.ToConnectorMappingPropertiesPtrOutputWithContext(context.Background())
}

func (i *connectorMappingPropertiesPtrType) ToConnectorMappingPropertiesPtrOutputWithContext(ctx context.Context) ConnectorMappingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingPropertiesPtrOutput)
}

type ConnectorMappingPropertiesOutput struct{ *pulumi.OutputState }

func (ConnectorMappingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingProperties)(nil)).Elem()
}

func (o ConnectorMappingPropertiesOutput) ToConnectorMappingPropertiesOutput() ConnectorMappingPropertiesOutput {
	return o
}

func (o ConnectorMappingPropertiesOutput) ToConnectorMappingPropertiesOutputWithContext(ctx context.Context) ConnectorMappingPropertiesOutput {
	return o
}

func (o ConnectorMappingPropertiesOutput) ToConnectorMappingPropertiesPtrOutput() ConnectorMappingPropertiesPtrOutput {
	return o.ToConnectorMappingPropertiesPtrOutputWithContext(context.Background())
}

func (o ConnectorMappingPropertiesOutput) ToConnectorMappingPropertiesPtrOutputWithContext(ctx context.Context) ConnectorMappingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectorMappingProperties) *ConnectorMappingProperties {
		return &v
	}).(ConnectorMappingPropertiesPtrOutput)
}

func (o ConnectorMappingPropertiesOutput) Availability() ConnectorMappingAvailabilityOutput {
	return o.ApplyT(func(v ConnectorMappingProperties) ConnectorMappingAvailability { return v.Availability }).(ConnectorMappingAvailabilityOutput)
}

func (o ConnectorMappingPropertiesOutput) CompleteOperation() ConnectorMappingCompleteOperationOutput {
	return o.ApplyT(func(v ConnectorMappingProperties) ConnectorMappingCompleteOperation { return v.CompleteOperation }).(ConnectorMappingCompleteOperationOutput)
}

func (o ConnectorMappingPropertiesOutput) ErrorManagement() ConnectorMappingErrorManagementOutput {
	return o.ApplyT(func(v ConnectorMappingProperties) ConnectorMappingErrorManagement { return v.ErrorManagement }).(ConnectorMappingErrorManagementOutput)
}

func (o ConnectorMappingPropertiesOutput) FileFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingProperties) *string { return v.FileFilter }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingPropertiesOutput) FolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingProperties) *string { return v.FolderPath }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingPropertiesOutput) Format() ConnectorMappingFormatOutput {
	return o.ApplyT(func(v ConnectorMappingProperties) ConnectorMappingFormat { return v.Format }).(ConnectorMappingFormatOutput)
}

func (o ConnectorMappingPropertiesOutput) HasHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConnectorMappingProperties) *bool { return v.HasHeader }).(pulumi.BoolPtrOutput)
}

func (o ConnectorMappingPropertiesOutput) Structure() ConnectorMappingStructureArrayOutput {
	return o.ApplyT(func(v ConnectorMappingProperties) []ConnectorMappingStructure { return v.Structure }).(ConnectorMappingStructureArrayOutput)
}

type ConnectorMappingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConnectorMappingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingProperties)(nil)).Elem()
}

func (o ConnectorMappingPropertiesPtrOutput) ToConnectorMappingPropertiesPtrOutput() ConnectorMappingPropertiesPtrOutput {
	return o
}

func (o ConnectorMappingPropertiesPtrOutput) ToConnectorMappingPropertiesPtrOutputWithContext(ctx context.Context) ConnectorMappingPropertiesPtrOutput {
	return o
}

func (o ConnectorMappingPropertiesPtrOutput) Elem() ConnectorMappingPropertiesOutput {
	return o.ApplyT(func(v *ConnectorMappingProperties) ConnectorMappingProperties {
		if v != nil {
			return *v
		}
		var ret ConnectorMappingProperties
		return ret
	}).(ConnectorMappingPropertiesOutput)
}

func (o ConnectorMappingPropertiesPtrOutput) Availability() ConnectorMappingAvailabilityPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingProperties) *ConnectorMappingAvailability {
		if v == nil {
			return nil
		}
		return &v.Availability
	}).(ConnectorMappingAvailabilityPtrOutput)
}

func (o ConnectorMappingPropertiesPtrOutput) CompleteOperation() ConnectorMappingCompleteOperationPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingProperties) *ConnectorMappingCompleteOperation {
		if v == nil {
			return nil
		}
		return &v.CompleteOperation
	}).(ConnectorMappingCompleteOperationPtrOutput)
}

func (o ConnectorMappingPropertiesPtrOutput) ErrorManagement() ConnectorMappingErrorManagementPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingProperties) *ConnectorMappingErrorManagement {
		if v == nil {
			return nil
		}
		return &v.ErrorManagement
	}).(ConnectorMappingErrorManagementPtrOutput)
}

func (o ConnectorMappingPropertiesPtrOutput) FileFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingProperties) *string {
		if v == nil {
			return nil
		}
		return v.FileFilter
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingPropertiesPtrOutput) FolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingProperties) *string {
		if v == nil {
			return nil
		}
		return v.FolderPath
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingPropertiesPtrOutput) Format() ConnectorMappingFormatPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingProperties) *ConnectorMappingFormat {
		if v == nil {
			return nil
		}
		return &v.Format
	}).(ConnectorMappingFormatPtrOutput)
}

func (o ConnectorMappingPropertiesPtrOutput) HasHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingProperties) *bool {
		if v == nil {
			return nil
		}
		return v.HasHeader
	}).(pulumi.BoolPtrOutput)
}

func (o ConnectorMappingPropertiesPtrOutput) Structure() ConnectorMappingStructureArrayOutput {
	return o.ApplyT(func(v *ConnectorMappingProperties) []ConnectorMappingStructure {
		if v == nil {
			return nil
		}
		return v.Structure
	}).(ConnectorMappingStructureArrayOutput)
}

type ConnectorMappingPropertiesResponse struct {
	Availability      ConnectorMappingAvailabilityResponse      `pulumi:"availability"`
	CompleteOperation ConnectorMappingCompleteOperationResponse `pulumi:"completeOperation"`
	ErrorManagement   ConnectorMappingErrorManagementResponse   `pulumi:"errorManagement"`
	FileFilter        *string                                   `pulumi:"fileFilter"`
	FolderPath        *string                                   `pulumi:"folderPath"`
	Format            ConnectorMappingFormatResponse            `pulumi:"format"`
	HasHeader         *bool                                     `pulumi:"hasHeader"`
	Structure         []ConnectorMappingStructureResponse       `pulumi:"structure"`
}





type ConnectorMappingPropertiesResponseInput interface {
	pulumi.Input

	ToConnectorMappingPropertiesResponseOutput() ConnectorMappingPropertiesResponseOutput
	ToConnectorMappingPropertiesResponseOutputWithContext(context.Context) ConnectorMappingPropertiesResponseOutput
}

type ConnectorMappingPropertiesResponseArgs struct {
	Availability      ConnectorMappingAvailabilityResponseInput      `pulumi:"availability"`
	CompleteOperation ConnectorMappingCompleteOperationResponseInput `pulumi:"completeOperation"`
	ErrorManagement   ConnectorMappingErrorManagementResponseInput   `pulumi:"errorManagement"`
	FileFilter        pulumi.StringPtrInput                          `pulumi:"fileFilter"`
	FolderPath        pulumi.StringPtrInput                          `pulumi:"folderPath"`
	Format            ConnectorMappingFormatResponseInput            `pulumi:"format"`
	HasHeader         pulumi.BoolPtrInput                            `pulumi:"hasHeader"`
	Structure         ConnectorMappingStructureResponseArrayInput    `pulumi:"structure"`
}

func (ConnectorMappingPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingPropertiesResponse)(nil)).Elem()
}

func (i ConnectorMappingPropertiesResponseArgs) ToConnectorMappingPropertiesResponseOutput() ConnectorMappingPropertiesResponseOutput {
	return i.ToConnectorMappingPropertiesResponseOutputWithContext(context.Background())
}

func (i ConnectorMappingPropertiesResponseArgs) ToConnectorMappingPropertiesResponseOutputWithContext(ctx context.Context) ConnectorMappingPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingPropertiesResponseOutput)
}

func (i ConnectorMappingPropertiesResponseArgs) ToConnectorMappingPropertiesResponsePtrOutput() ConnectorMappingPropertiesResponsePtrOutput {
	return i.ToConnectorMappingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ConnectorMappingPropertiesResponseArgs) ToConnectorMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingPropertiesResponseOutput).ToConnectorMappingPropertiesResponsePtrOutputWithContext(ctx)
}









type ConnectorMappingPropertiesResponsePtrInput interface {
	pulumi.Input

	ToConnectorMappingPropertiesResponsePtrOutput() ConnectorMappingPropertiesResponsePtrOutput
	ToConnectorMappingPropertiesResponsePtrOutputWithContext(context.Context) ConnectorMappingPropertiesResponsePtrOutput
}

type connectorMappingPropertiesResponsePtrType ConnectorMappingPropertiesResponseArgs

func ConnectorMappingPropertiesResponsePtr(v *ConnectorMappingPropertiesResponseArgs) ConnectorMappingPropertiesResponsePtrInput {
	return (*connectorMappingPropertiesResponsePtrType)(v)
}

func (*connectorMappingPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingPropertiesResponse)(nil)).Elem()
}

func (i *connectorMappingPropertiesResponsePtrType) ToConnectorMappingPropertiesResponsePtrOutput() ConnectorMappingPropertiesResponsePtrOutput {
	return i.ToConnectorMappingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *connectorMappingPropertiesResponsePtrType) ToConnectorMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingPropertiesResponsePtrOutput)
}

type ConnectorMappingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConnectorMappingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingPropertiesResponse)(nil)).Elem()
}

func (o ConnectorMappingPropertiesResponseOutput) ToConnectorMappingPropertiesResponseOutput() ConnectorMappingPropertiesResponseOutput {
	return o
}

func (o ConnectorMappingPropertiesResponseOutput) ToConnectorMappingPropertiesResponseOutputWithContext(ctx context.Context) ConnectorMappingPropertiesResponseOutput {
	return o
}

func (o ConnectorMappingPropertiesResponseOutput) ToConnectorMappingPropertiesResponsePtrOutput() ConnectorMappingPropertiesResponsePtrOutput {
	return o.ToConnectorMappingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ConnectorMappingPropertiesResponseOutput) ToConnectorMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectorMappingPropertiesResponse) *ConnectorMappingPropertiesResponse {
		return &v
	}).(ConnectorMappingPropertiesResponsePtrOutput)
}

func (o ConnectorMappingPropertiesResponseOutput) Availability() ConnectorMappingAvailabilityResponseOutput {
	return o.ApplyT(func(v ConnectorMappingPropertiesResponse) ConnectorMappingAvailabilityResponse { return v.Availability }).(ConnectorMappingAvailabilityResponseOutput)
}

func (o ConnectorMappingPropertiesResponseOutput) CompleteOperation() ConnectorMappingCompleteOperationResponseOutput {
	return o.ApplyT(func(v ConnectorMappingPropertiesResponse) ConnectorMappingCompleteOperationResponse {
		return v.CompleteOperation
	}).(ConnectorMappingCompleteOperationResponseOutput)
}

func (o ConnectorMappingPropertiesResponseOutput) ErrorManagement() ConnectorMappingErrorManagementResponseOutput {
	return o.ApplyT(func(v ConnectorMappingPropertiesResponse) ConnectorMappingErrorManagementResponse {
		return v.ErrorManagement
	}).(ConnectorMappingErrorManagementResponseOutput)
}

func (o ConnectorMappingPropertiesResponseOutput) FileFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingPropertiesResponse) *string { return v.FileFilter }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingPropertiesResponseOutput) FolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingPropertiesResponse) *string { return v.FolderPath }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingPropertiesResponseOutput) Format() ConnectorMappingFormatResponseOutput {
	return o.ApplyT(func(v ConnectorMappingPropertiesResponse) ConnectorMappingFormatResponse { return v.Format }).(ConnectorMappingFormatResponseOutput)
}

func (o ConnectorMappingPropertiesResponseOutput) HasHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConnectorMappingPropertiesResponse) *bool { return v.HasHeader }).(pulumi.BoolPtrOutput)
}

func (o ConnectorMappingPropertiesResponseOutput) Structure() ConnectorMappingStructureResponseArrayOutput {
	return o.ApplyT(func(v ConnectorMappingPropertiesResponse) []ConnectorMappingStructureResponse { return v.Structure }).(ConnectorMappingStructureResponseArrayOutput)
}

type ConnectorMappingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectorMappingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMappingPropertiesResponse)(nil)).Elem()
}

func (o ConnectorMappingPropertiesResponsePtrOutput) ToConnectorMappingPropertiesResponsePtrOutput() ConnectorMappingPropertiesResponsePtrOutput {
	return o
}

func (o ConnectorMappingPropertiesResponsePtrOutput) ToConnectorMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) ConnectorMappingPropertiesResponsePtrOutput {
	return o
}

func (o ConnectorMappingPropertiesResponsePtrOutput) Elem() ConnectorMappingPropertiesResponseOutput {
	return o.ApplyT(func(v *ConnectorMappingPropertiesResponse) ConnectorMappingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ConnectorMappingPropertiesResponse
		return ret
	}).(ConnectorMappingPropertiesResponseOutput)
}

func (o ConnectorMappingPropertiesResponsePtrOutput) Availability() ConnectorMappingAvailabilityResponsePtrOutput {
	return o.ApplyT(func(v *ConnectorMappingPropertiesResponse) *ConnectorMappingAvailabilityResponse {
		if v == nil {
			return nil
		}
		return &v.Availability
	}).(ConnectorMappingAvailabilityResponsePtrOutput)
}

func (o ConnectorMappingPropertiesResponsePtrOutput) CompleteOperation() ConnectorMappingCompleteOperationResponsePtrOutput {
	return o.ApplyT(func(v *ConnectorMappingPropertiesResponse) *ConnectorMappingCompleteOperationResponse {
		if v == nil {
			return nil
		}
		return &v.CompleteOperation
	}).(ConnectorMappingCompleteOperationResponsePtrOutput)
}

func (o ConnectorMappingPropertiesResponsePtrOutput) ErrorManagement() ConnectorMappingErrorManagementResponsePtrOutput {
	return o.ApplyT(func(v *ConnectorMappingPropertiesResponse) *ConnectorMappingErrorManagementResponse {
		if v == nil {
			return nil
		}
		return &v.ErrorManagement
	}).(ConnectorMappingErrorManagementResponsePtrOutput)
}

func (o ConnectorMappingPropertiesResponsePtrOutput) FileFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FileFilter
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingPropertiesResponsePtrOutput) FolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FolderPath
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingPropertiesResponsePtrOutput) Format() ConnectorMappingFormatResponsePtrOutput {
	return o.ApplyT(func(v *ConnectorMappingPropertiesResponse) *ConnectorMappingFormatResponse {
		if v == nil {
			return nil
		}
		return &v.Format
	}).(ConnectorMappingFormatResponsePtrOutput)
}

func (o ConnectorMappingPropertiesResponsePtrOutput) HasHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConnectorMappingPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.HasHeader
	}).(pulumi.BoolPtrOutput)
}

func (o ConnectorMappingPropertiesResponsePtrOutput) Structure() ConnectorMappingStructureResponseArrayOutput {
	return o.ApplyT(func(v *ConnectorMappingPropertiesResponse) []ConnectorMappingStructureResponse {
		if v == nil {
			return nil
		}
		return v.Structure
	}).(ConnectorMappingStructureResponseArrayOutput)
}

type ConnectorMappingStructure struct {
	ColumnName            string  `pulumi:"columnName"`
	CustomFormatSpecifier *string `pulumi:"customFormatSpecifier"`
	IsEncrypted           *bool   `pulumi:"isEncrypted"`
	PropertyName          string  `pulumi:"propertyName"`
}





type ConnectorMappingStructureInput interface {
	pulumi.Input

	ToConnectorMappingStructureOutput() ConnectorMappingStructureOutput
	ToConnectorMappingStructureOutputWithContext(context.Context) ConnectorMappingStructureOutput
}

type ConnectorMappingStructureArgs struct {
	ColumnName            pulumi.StringInput    `pulumi:"columnName"`
	CustomFormatSpecifier pulumi.StringPtrInput `pulumi:"customFormatSpecifier"`
	IsEncrypted           pulumi.BoolPtrInput   `pulumi:"isEncrypted"`
	PropertyName          pulumi.StringInput    `pulumi:"propertyName"`
}

func (ConnectorMappingStructureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingStructure)(nil)).Elem()
}

func (i ConnectorMappingStructureArgs) ToConnectorMappingStructureOutput() ConnectorMappingStructureOutput {
	return i.ToConnectorMappingStructureOutputWithContext(context.Background())
}

func (i ConnectorMappingStructureArgs) ToConnectorMappingStructureOutputWithContext(ctx context.Context) ConnectorMappingStructureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingStructureOutput)
}





type ConnectorMappingStructureArrayInput interface {
	pulumi.Input

	ToConnectorMappingStructureArrayOutput() ConnectorMappingStructureArrayOutput
	ToConnectorMappingStructureArrayOutputWithContext(context.Context) ConnectorMappingStructureArrayOutput
}

type ConnectorMappingStructureArray []ConnectorMappingStructureInput

func (ConnectorMappingStructureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectorMappingStructure)(nil)).Elem()
}

func (i ConnectorMappingStructureArray) ToConnectorMappingStructureArrayOutput() ConnectorMappingStructureArrayOutput {
	return i.ToConnectorMappingStructureArrayOutputWithContext(context.Background())
}

func (i ConnectorMappingStructureArray) ToConnectorMappingStructureArrayOutputWithContext(ctx context.Context) ConnectorMappingStructureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingStructureArrayOutput)
}

type ConnectorMappingStructureOutput struct{ *pulumi.OutputState }

func (ConnectorMappingStructureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingStructure)(nil)).Elem()
}

func (o ConnectorMappingStructureOutput) ToConnectorMappingStructureOutput() ConnectorMappingStructureOutput {
	return o
}

func (o ConnectorMappingStructureOutput) ToConnectorMappingStructureOutputWithContext(ctx context.Context) ConnectorMappingStructureOutput {
	return o
}

func (o ConnectorMappingStructureOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorMappingStructure) string { return v.ColumnName }).(pulumi.StringOutput)
}

func (o ConnectorMappingStructureOutput) CustomFormatSpecifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingStructure) *string { return v.CustomFormatSpecifier }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingStructureOutput) IsEncrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConnectorMappingStructure) *bool { return v.IsEncrypted }).(pulumi.BoolPtrOutput)
}

func (o ConnectorMappingStructureOutput) PropertyName() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorMappingStructure) string { return v.PropertyName }).(pulumi.StringOutput)
}

type ConnectorMappingStructureArrayOutput struct{ *pulumi.OutputState }

func (ConnectorMappingStructureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectorMappingStructure)(nil)).Elem()
}

func (o ConnectorMappingStructureArrayOutput) ToConnectorMappingStructureArrayOutput() ConnectorMappingStructureArrayOutput {
	return o
}

func (o ConnectorMappingStructureArrayOutput) ToConnectorMappingStructureArrayOutputWithContext(ctx context.Context) ConnectorMappingStructureArrayOutput {
	return o
}

func (o ConnectorMappingStructureArrayOutput) Index(i pulumi.IntInput) ConnectorMappingStructureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectorMappingStructure {
		return vs[0].([]ConnectorMappingStructure)[vs[1].(int)]
	}).(ConnectorMappingStructureOutput)
}

type ConnectorMappingStructureResponse struct {
	ColumnName            string  `pulumi:"columnName"`
	CustomFormatSpecifier *string `pulumi:"customFormatSpecifier"`
	IsEncrypted           *bool   `pulumi:"isEncrypted"`
	PropertyName          string  `pulumi:"propertyName"`
}





type ConnectorMappingStructureResponseInput interface {
	pulumi.Input

	ToConnectorMappingStructureResponseOutput() ConnectorMappingStructureResponseOutput
	ToConnectorMappingStructureResponseOutputWithContext(context.Context) ConnectorMappingStructureResponseOutput
}

type ConnectorMappingStructureResponseArgs struct {
	ColumnName            pulumi.StringInput    `pulumi:"columnName"`
	CustomFormatSpecifier pulumi.StringPtrInput `pulumi:"customFormatSpecifier"`
	IsEncrypted           pulumi.BoolPtrInput   `pulumi:"isEncrypted"`
	PropertyName          pulumi.StringInput    `pulumi:"propertyName"`
}

func (ConnectorMappingStructureResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingStructureResponse)(nil)).Elem()
}

func (i ConnectorMappingStructureResponseArgs) ToConnectorMappingStructureResponseOutput() ConnectorMappingStructureResponseOutput {
	return i.ToConnectorMappingStructureResponseOutputWithContext(context.Background())
}

func (i ConnectorMappingStructureResponseArgs) ToConnectorMappingStructureResponseOutputWithContext(ctx context.Context) ConnectorMappingStructureResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingStructureResponseOutput)
}





type ConnectorMappingStructureResponseArrayInput interface {
	pulumi.Input

	ToConnectorMappingStructureResponseArrayOutput() ConnectorMappingStructureResponseArrayOutput
	ToConnectorMappingStructureResponseArrayOutputWithContext(context.Context) ConnectorMappingStructureResponseArrayOutput
}

type ConnectorMappingStructureResponseArray []ConnectorMappingStructureResponseInput

func (ConnectorMappingStructureResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectorMappingStructureResponse)(nil)).Elem()
}

func (i ConnectorMappingStructureResponseArray) ToConnectorMappingStructureResponseArrayOutput() ConnectorMappingStructureResponseArrayOutput {
	return i.ToConnectorMappingStructureResponseArrayOutputWithContext(context.Background())
}

func (i ConnectorMappingStructureResponseArray) ToConnectorMappingStructureResponseArrayOutputWithContext(ctx context.Context) ConnectorMappingStructureResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingStructureResponseArrayOutput)
}

type ConnectorMappingStructureResponseOutput struct{ *pulumi.OutputState }

func (ConnectorMappingStructureResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMappingStructureResponse)(nil)).Elem()
}

func (o ConnectorMappingStructureResponseOutput) ToConnectorMappingStructureResponseOutput() ConnectorMappingStructureResponseOutput {
	return o
}

func (o ConnectorMappingStructureResponseOutput) ToConnectorMappingStructureResponseOutputWithContext(ctx context.Context) ConnectorMappingStructureResponseOutput {
	return o
}

func (o ConnectorMappingStructureResponseOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorMappingStructureResponse) string { return v.ColumnName }).(pulumi.StringOutput)
}

func (o ConnectorMappingStructureResponseOutput) CustomFormatSpecifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorMappingStructureResponse) *string { return v.CustomFormatSpecifier }).(pulumi.StringPtrOutput)
}

func (o ConnectorMappingStructureResponseOutput) IsEncrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConnectorMappingStructureResponse) *bool { return v.IsEncrypted }).(pulumi.BoolPtrOutput)
}

func (o ConnectorMappingStructureResponseOutput) PropertyName() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorMappingStructureResponse) string { return v.PropertyName }).(pulumi.StringOutput)
}

type ConnectorMappingStructureResponseArrayOutput struct{ *pulumi.OutputState }

func (ConnectorMappingStructureResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectorMappingStructureResponse)(nil)).Elem()
}

func (o ConnectorMappingStructureResponseArrayOutput) ToConnectorMappingStructureResponseArrayOutput() ConnectorMappingStructureResponseArrayOutput {
	return o
}

func (o ConnectorMappingStructureResponseArrayOutput) ToConnectorMappingStructureResponseArrayOutputWithContext(ctx context.Context) ConnectorMappingStructureResponseArrayOutput {
	return o
}

func (o ConnectorMappingStructureResponseArrayOutput) Index(i pulumi.IntInput) ConnectorMappingStructureResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectorMappingStructureResponse {
		return vs[0].([]ConnectorMappingStructureResponse)[vs[1].(int)]
	}).(ConnectorMappingStructureResponseOutput)
}

type DataSourcePrecedenceResponse struct {
	DataSourceReferenceId string `pulumi:"dataSourceReferenceId"`
	DataSourceType        string `pulumi:"dataSourceType"`
	Id                    int    `pulumi:"id"`
	Name                  string `pulumi:"name"`
	Precedence            *int   `pulumi:"precedence"`
	Status                string `pulumi:"status"`
}





type DataSourcePrecedenceResponseInput interface {
	pulumi.Input

	ToDataSourcePrecedenceResponseOutput() DataSourcePrecedenceResponseOutput
	ToDataSourcePrecedenceResponseOutputWithContext(context.Context) DataSourcePrecedenceResponseOutput
}

type DataSourcePrecedenceResponseArgs struct {
	DataSourceReferenceId pulumi.StringInput `pulumi:"dataSourceReferenceId"`
	DataSourceType        pulumi.StringInput `pulumi:"dataSourceType"`
	Id                    pulumi.IntInput    `pulumi:"id"`
	Name                  pulumi.StringInput `pulumi:"name"`
	Precedence            pulumi.IntPtrInput `pulumi:"precedence"`
	Status                pulumi.StringInput `pulumi:"status"`
}

func (DataSourcePrecedenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourcePrecedenceResponse)(nil)).Elem()
}

func (i DataSourcePrecedenceResponseArgs) ToDataSourcePrecedenceResponseOutput() DataSourcePrecedenceResponseOutput {
	return i.ToDataSourcePrecedenceResponseOutputWithContext(context.Background())
}

func (i DataSourcePrecedenceResponseArgs) ToDataSourcePrecedenceResponseOutputWithContext(ctx context.Context) DataSourcePrecedenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourcePrecedenceResponseOutput)
}





type DataSourcePrecedenceResponseArrayInput interface {
	pulumi.Input

	ToDataSourcePrecedenceResponseArrayOutput() DataSourcePrecedenceResponseArrayOutput
	ToDataSourcePrecedenceResponseArrayOutputWithContext(context.Context) DataSourcePrecedenceResponseArrayOutput
}

type DataSourcePrecedenceResponseArray []DataSourcePrecedenceResponseInput

func (DataSourcePrecedenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataSourcePrecedenceResponse)(nil)).Elem()
}

func (i DataSourcePrecedenceResponseArray) ToDataSourcePrecedenceResponseArrayOutput() DataSourcePrecedenceResponseArrayOutput {
	return i.ToDataSourcePrecedenceResponseArrayOutputWithContext(context.Background())
}

func (i DataSourcePrecedenceResponseArray) ToDataSourcePrecedenceResponseArrayOutputWithContext(ctx context.Context) DataSourcePrecedenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourcePrecedenceResponseArrayOutput)
}

type DataSourcePrecedenceResponseOutput struct{ *pulumi.OutputState }

func (DataSourcePrecedenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourcePrecedenceResponse)(nil)).Elem()
}

func (o DataSourcePrecedenceResponseOutput) ToDataSourcePrecedenceResponseOutput() DataSourcePrecedenceResponseOutput {
	return o
}

func (o DataSourcePrecedenceResponseOutput) ToDataSourcePrecedenceResponseOutputWithContext(ctx context.Context) DataSourcePrecedenceResponseOutput {
	return o
}

func (o DataSourcePrecedenceResponseOutput) DataSourceReferenceId() pulumi.StringOutput {
	return o.ApplyT(func(v DataSourcePrecedenceResponse) string { return v.DataSourceReferenceId }).(pulumi.StringOutput)
}

func (o DataSourcePrecedenceResponseOutput) DataSourceType() pulumi.StringOutput {
	return o.ApplyT(func(v DataSourcePrecedenceResponse) string { return v.DataSourceType }).(pulumi.StringOutput)
}

func (o DataSourcePrecedenceResponseOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v DataSourcePrecedenceResponse) int { return v.Id }).(pulumi.IntOutput)
}

func (o DataSourcePrecedenceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataSourcePrecedenceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DataSourcePrecedenceResponseOutput) Precedence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataSourcePrecedenceResponse) *int { return v.Precedence }).(pulumi.IntPtrOutput)
}

func (o DataSourcePrecedenceResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v DataSourcePrecedenceResponse) string { return v.Status }).(pulumi.StringOutput)
}

type DataSourcePrecedenceResponseArrayOutput struct{ *pulumi.OutputState }

func (DataSourcePrecedenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataSourcePrecedenceResponse)(nil)).Elem()
}

func (o DataSourcePrecedenceResponseArrayOutput) ToDataSourcePrecedenceResponseArrayOutput() DataSourcePrecedenceResponseArrayOutput {
	return o
}

func (o DataSourcePrecedenceResponseArrayOutput) ToDataSourcePrecedenceResponseArrayOutputWithContext(ctx context.Context) DataSourcePrecedenceResponseArrayOutput {
	return o
}

func (o DataSourcePrecedenceResponseArrayOutput) Index(i pulumi.IntInput) DataSourcePrecedenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataSourcePrecedenceResponse {
		return vs[0].([]DataSourcePrecedenceResponse)[vs[1].(int)]
	}).(DataSourcePrecedenceResponseOutput)
}

type HubBillingInfoFormat struct {
	MaxUnits *int    `pulumi:"maxUnits"`
	MinUnits *int    `pulumi:"minUnits"`
	SkuName  *string `pulumi:"skuName"`
}





type HubBillingInfoFormatInput interface {
	pulumi.Input

	ToHubBillingInfoFormatOutput() HubBillingInfoFormatOutput
	ToHubBillingInfoFormatOutputWithContext(context.Context) HubBillingInfoFormatOutput
}

type HubBillingInfoFormatArgs struct {
	MaxUnits pulumi.IntPtrInput    `pulumi:"maxUnits"`
	MinUnits pulumi.IntPtrInput    `pulumi:"minUnits"`
	SkuName  pulumi.StringPtrInput `pulumi:"skuName"`
}

func (HubBillingInfoFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HubBillingInfoFormat)(nil)).Elem()
}

func (i HubBillingInfoFormatArgs) ToHubBillingInfoFormatOutput() HubBillingInfoFormatOutput {
	return i.ToHubBillingInfoFormatOutputWithContext(context.Background())
}

func (i HubBillingInfoFormatArgs) ToHubBillingInfoFormatOutputWithContext(ctx context.Context) HubBillingInfoFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubBillingInfoFormatOutput)
}

func (i HubBillingInfoFormatArgs) ToHubBillingInfoFormatPtrOutput() HubBillingInfoFormatPtrOutput {
	return i.ToHubBillingInfoFormatPtrOutputWithContext(context.Background())
}

func (i HubBillingInfoFormatArgs) ToHubBillingInfoFormatPtrOutputWithContext(ctx context.Context) HubBillingInfoFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubBillingInfoFormatOutput).ToHubBillingInfoFormatPtrOutputWithContext(ctx)
}









type HubBillingInfoFormatPtrInput interface {
	pulumi.Input

	ToHubBillingInfoFormatPtrOutput() HubBillingInfoFormatPtrOutput
	ToHubBillingInfoFormatPtrOutputWithContext(context.Context) HubBillingInfoFormatPtrOutput
}

type hubBillingInfoFormatPtrType HubBillingInfoFormatArgs

func HubBillingInfoFormatPtr(v *HubBillingInfoFormatArgs) HubBillingInfoFormatPtrInput {
	return (*hubBillingInfoFormatPtrType)(v)
}

func (*hubBillingInfoFormatPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HubBillingInfoFormat)(nil)).Elem()
}

func (i *hubBillingInfoFormatPtrType) ToHubBillingInfoFormatPtrOutput() HubBillingInfoFormatPtrOutput {
	return i.ToHubBillingInfoFormatPtrOutputWithContext(context.Background())
}

func (i *hubBillingInfoFormatPtrType) ToHubBillingInfoFormatPtrOutputWithContext(ctx context.Context) HubBillingInfoFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubBillingInfoFormatPtrOutput)
}

type HubBillingInfoFormatOutput struct{ *pulumi.OutputState }

func (HubBillingInfoFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HubBillingInfoFormat)(nil)).Elem()
}

func (o HubBillingInfoFormatOutput) ToHubBillingInfoFormatOutput() HubBillingInfoFormatOutput {
	return o
}

func (o HubBillingInfoFormatOutput) ToHubBillingInfoFormatOutputWithContext(ctx context.Context) HubBillingInfoFormatOutput {
	return o
}

func (o HubBillingInfoFormatOutput) ToHubBillingInfoFormatPtrOutput() HubBillingInfoFormatPtrOutput {
	return o.ToHubBillingInfoFormatPtrOutputWithContext(context.Background())
}

func (o HubBillingInfoFormatOutput) ToHubBillingInfoFormatPtrOutputWithContext(ctx context.Context) HubBillingInfoFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HubBillingInfoFormat) *HubBillingInfoFormat {
		return &v
	}).(HubBillingInfoFormatPtrOutput)
}

func (o HubBillingInfoFormatOutput) MaxUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HubBillingInfoFormat) *int { return v.MaxUnits }).(pulumi.IntPtrOutput)
}

func (o HubBillingInfoFormatOutput) MinUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HubBillingInfoFormat) *int { return v.MinUnits }).(pulumi.IntPtrOutput)
}

func (o HubBillingInfoFormatOutput) SkuName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HubBillingInfoFormat) *string { return v.SkuName }).(pulumi.StringPtrOutput)
}

type HubBillingInfoFormatPtrOutput struct{ *pulumi.OutputState }

func (HubBillingInfoFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HubBillingInfoFormat)(nil)).Elem()
}

func (o HubBillingInfoFormatPtrOutput) ToHubBillingInfoFormatPtrOutput() HubBillingInfoFormatPtrOutput {
	return o
}

func (o HubBillingInfoFormatPtrOutput) ToHubBillingInfoFormatPtrOutputWithContext(ctx context.Context) HubBillingInfoFormatPtrOutput {
	return o
}

func (o HubBillingInfoFormatPtrOutput) Elem() HubBillingInfoFormatOutput {
	return o.ApplyT(func(v *HubBillingInfoFormat) HubBillingInfoFormat {
		if v != nil {
			return *v
		}
		var ret HubBillingInfoFormat
		return ret
	}).(HubBillingInfoFormatOutput)
}

func (o HubBillingInfoFormatPtrOutput) MaxUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HubBillingInfoFormat) *int {
		if v == nil {
			return nil
		}
		return v.MaxUnits
	}).(pulumi.IntPtrOutput)
}

func (o HubBillingInfoFormatPtrOutput) MinUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HubBillingInfoFormat) *int {
		if v == nil {
			return nil
		}
		return v.MinUnits
	}).(pulumi.IntPtrOutput)
}

func (o HubBillingInfoFormatPtrOutput) SkuName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HubBillingInfoFormat) *string {
		if v == nil {
			return nil
		}
		return v.SkuName
	}).(pulumi.StringPtrOutput)
}

type HubBillingInfoFormatResponse struct {
	MaxUnits *int    `pulumi:"maxUnits"`
	MinUnits *int    `pulumi:"minUnits"`
	SkuName  *string `pulumi:"skuName"`
}





type HubBillingInfoFormatResponseInput interface {
	pulumi.Input

	ToHubBillingInfoFormatResponseOutput() HubBillingInfoFormatResponseOutput
	ToHubBillingInfoFormatResponseOutputWithContext(context.Context) HubBillingInfoFormatResponseOutput
}

type HubBillingInfoFormatResponseArgs struct {
	MaxUnits pulumi.IntPtrInput    `pulumi:"maxUnits"`
	MinUnits pulumi.IntPtrInput    `pulumi:"minUnits"`
	SkuName  pulumi.StringPtrInput `pulumi:"skuName"`
}

func (HubBillingInfoFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HubBillingInfoFormatResponse)(nil)).Elem()
}

func (i HubBillingInfoFormatResponseArgs) ToHubBillingInfoFormatResponseOutput() HubBillingInfoFormatResponseOutput {
	return i.ToHubBillingInfoFormatResponseOutputWithContext(context.Background())
}

func (i HubBillingInfoFormatResponseArgs) ToHubBillingInfoFormatResponseOutputWithContext(ctx context.Context) HubBillingInfoFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubBillingInfoFormatResponseOutput)
}

func (i HubBillingInfoFormatResponseArgs) ToHubBillingInfoFormatResponsePtrOutput() HubBillingInfoFormatResponsePtrOutput {
	return i.ToHubBillingInfoFormatResponsePtrOutputWithContext(context.Background())
}

func (i HubBillingInfoFormatResponseArgs) ToHubBillingInfoFormatResponsePtrOutputWithContext(ctx context.Context) HubBillingInfoFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubBillingInfoFormatResponseOutput).ToHubBillingInfoFormatResponsePtrOutputWithContext(ctx)
}









type HubBillingInfoFormatResponsePtrInput interface {
	pulumi.Input

	ToHubBillingInfoFormatResponsePtrOutput() HubBillingInfoFormatResponsePtrOutput
	ToHubBillingInfoFormatResponsePtrOutputWithContext(context.Context) HubBillingInfoFormatResponsePtrOutput
}

type hubBillingInfoFormatResponsePtrType HubBillingInfoFormatResponseArgs

func HubBillingInfoFormatResponsePtr(v *HubBillingInfoFormatResponseArgs) HubBillingInfoFormatResponsePtrInput {
	return (*hubBillingInfoFormatResponsePtrType)(v)
}

func (*hubBillingInfoFormatResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HubBillingInfoFormatResponse)(nil)).Elem()
}

func (i *hubBillingInfoFormatResponsePtrType) ToHubBillingInfoFormatResponsePtrOutput() HubBillingInfoFormatResponsePtrOutput {
	return i.ToHubBillingInfoFormatResponsePtrOutputWithContext(context.Background())
}

func (i *hubBillingInfoFormatResponsePtrType) ToHubBillingInfoFormatResponsePtrOutputWithContext(ctx context.Context) HubBillingInfoFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubBillingInfoFormatResponsePtrOutput)
}

type HubBillingInfoFormatResponseOutput struct{ *pulumi.OutputState }

func (HubBillingInfoFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HubBillingInfoFormatResponse)(nil)).Elem()
}

func (o HubBillingInfoFormatResponseOutput) ToHubBillingInfoFormatResponseOutput() HubBillingInfoFormatResponseOutput {
	return o
}

func (o HubBillingInfoFormatResponseOutput) ToHubBillingInfoFormatResponseOutputWithContext(ctx context.Context) HubBillingInfoFormatResponseOutput {
	return o
}

func (o HubBillingInfoFormatResponseOutput) ToHubBillingInfoFormatResponsePtrOutput() HubBillingInfoFormatResponsePtrOutput {
	return o.ToHubBillingInfoFormatResponsePtrOutputWithContext(context.Background())
}

func (o HubBillingInfoFormatResponseOutput) ToHubBillingInfoFormatResponsePtrOutputWithContext(ctx context.Context) HubBillingInfoFormatResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HubBillingInfoFormatResponse) *HubBillingInfoFormatResponse {
		return &v
	}).(HubBillingInfoFormatResponsePtrOutput)
}

func (o HubBillingInfoFormatResponseOutput) MaxUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HubBillingInfoFormatResponse) *int { return v.MaxUnits }).(pulumi.IntPtrOutput)
}

func (o HubBillingInfoFormatResponseOutput) MinUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HubBillingInfoFormatResponse) *int { return v.MinUnits }).(pulumi.IntPtrOutput)
}

func (o HubBillingInfoFormatResponseOutput) SkuName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HubBillingInfoFormatResponse) *string { return v.SkuName }).(pulumi.StringPtrOutput)
}

type HubBillingInfoFormatResponsePtrOutput struct{ *pulumi.OutputState }

func (HubBillingInfoFormatResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HubBillingInfoFormatResponse)(nil)).Elem()
}

func (o HubBillingInfoFormatResponsePtrOutput) ToHubBillingInfoFormatResponsePtrOutput() HubBillingInfoFormatResponsePtrOutput {
	return o
}

func (o HubBillingInfoFormatResponsePtrOutput) ToHubBillingInfoFormatResponsePtrOutputWithContext(ctx context.Context) HubBillingInfoFormatResponsePtrOutput {
	return o
}

func (o HubBillingInfoFormatResponsePtrOutput) Elem() HubBillingInfoFormatResponseOutput {
	return o.ApplyT(func(v *HubBillingInfoFormatResponse) HubBillingInfoFormatResponse {
		if v != nil {
			return *v
		}
		var ret HubBillingInfoFormatResponse
		return ret
	}).(HubBillingInfoFormatResponseOutput)
}

func (o HubBillingInfoFormatResponsePtrOutput) MaxUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HubBillingInfoFormatResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxUnits
	}).(pulumi.IntPtrOutput)
}

func (o HubBillingInfoFormatResponsePtrOutput) MinUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HubBillingInfoFormatResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinUnits
	}).(pulumi.IntPtrOutput)
}

func (o HubBillingInfoFormatResponsePtrOutput) SkuName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HubBillingInfoFormatResponse) *string {
		if v == nil {
			return nil
		}
		return v.SkuName
	}).(pulumi.StringPtrOutput)
}

type KpiAlias struct {
	AliasName  string `pulumi:"aliasName"`
	Expression string `pulumi:"expression"`
}





type KpiAliasInput interface {
	pulumi.Input

	ToKpiAliasOutput() KpiAliasOutput
	ToKpiAliasOutputWithContext(context.Context) KpiAliasOutput
}

type KpiAliasArgs struct {
	AliasName  pulumi.StringInput `pulumi:"aliasName"`
	Expression pulumi.StringInput `pulumi:"expression"`
}

func (KpiAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiAlias)(nil)).Elem()
}

func (i KpiAliasArgs) ToKpiAliasOutput() KpiAliasOutput {
	return i.ToKpiAliasOutputWithContext(context.Background())
}

func (i KpiAliasArgs) ToKpiAliasOutputWithContext(ctx context.Context) KpiAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiAliasOutput)
}





type KpiAliasArrayInput interface {
	pulumi.Input

	ToKpiAliasArrayOutput() KpiAliasArrayOutput
	ToKpiAliasArrayOutputWithContext(context.Context) KpiAliasArrayOutput
}

type KpiAliasArray []KpiAliasInput

func (KpiAliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiAlias)(nil)).Elem()
}

func (i KpiAliasArray) ToKpiAliasArrayOutput() KpiAliasArrayOutput {
	return i.ToKpiAliasArrayOutputWithContext(context.Background())
}

func (i KpiAliasArray) ToKpiAliasArrayOutputWithContext(ctx context.Context) KpiAliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiAliasArrayOutput)
}

type KpiAliasOutput struct{ *pulumi.OutputState }

func (KpiAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiAlias)(nil)).Elem()
}

func (o KpiAliasOutput) ToKpiAliasOutput() KpiAliasOutput {
	return o
}

func (o KpiAliasOutput) ToKpiAliasOutputWithContext(ctx context.Context) KpiAliasOutput {
	return o
}

func (o KpiAliasOutput) AliasName() pulumi.StringOutput {
	return o.ApplyT(func(v KpiAlias) string { return v.AliasName }).(pulumi.StringOutput)
}

func (o KpiAliasOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v KpiAlias) string { return v.Expression }).(pulumi.StringOutput)
}

type KpiAliasArrayOutput struct{ *pulumi.OutputState }

func (KpiAliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiAlias)(nil)).Elem()
}

func (o KpiAliasArrayOutput) ToKpiAliasArrayOutput() KpiAliasArrayOutput {
	return o
}

func (o KpiAliasArrayOutput) ToKpiAliasArrayOutputWithContext(ctx context.Context) KpiAliasArrayOutput {
	return o
}

func (o KpiAliasArrayOutput) Index(i pulumi.IntInput) KpiAliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KpiAlias {
		return vs[0].([]KpiAlias)[vs[1].(int)]
	}).(KpiAliasOutput)
}

type KpiAliasResponse struct {
	AliasName  string `pulumi:"aliasName"`
	Expression string `pulumi:"expression"`
}





type KpiAliasResponseInput interface {
	pulumi.Input

	ToKpiAliasResponseOutput() KpiAliasResponseOutput
	ToKpiAliasResponseOutputWithContext(context.Context) KpiAliasResponseOutput
}

type KpiAliasResponseArgs struct {
	AliasName  pulumi.StringInput `pulumi:"aliasName"`
	Expression pulumi.StringInput `pulumi:"expression"`
}

func (KpiAliasResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiAliasResponse)(nil)).Elem()
}

func (i KpiAliasResponseArgs) ToKpiAliasResponseOutput() KpiAliasResponseOutput {
	return i.ToKpiAliasResponseOutputWithContext(context.Background())
}

func (i KpiAliasResponseArgs) ToKpiAliasResponseOutputWithContext(ctx context.Context) KpiAliasResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiAliasResponseOutput)
}





type KpiAliasResponseArrayInput interface {
	pulumi.Input

	ToKpiAliasResponseArrayOutput() KpiAliasResponseArrayOutput
	ToKpiAliasResponseArrayOutputWithContext(context.Context) KpiAliasResponseArrayOutput
}

type KpiAliasResponseArray []KpiAliasResponseInput

func (KpiAliasResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiAliasResponse)(nil)).Elem()
}

func (i KpiAliasResponseArray) ToKpiAliasResponseArrayOutput() KpiAliasResponseArrayOutput {
	return i.ToKpiAliasResponseArrayOutputWithContext(context.Background())
}

func (i KpiAliasResponseArray) ToKpiAliasResponseArrayOutputWithContext(ctx context.Context) KpiAliasResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiAliasResponseArrayOutput)
}

type KpiAliasResponseOutput struct{ *pulumi.OutputState }

func (KpiAliasResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiAliasResponse)(nil)).Elem()
}

func (o KpiAliasResponseOutput) ToKpiAliasResponseOutput() KpiAliasResponseOutput {
	return o
}

func (o KpiAliasResponseOutput) ToKpiAliasResponseOutputWithContext(ctx context.Context) KpiAliasResponseOutput {
	return o
}

func (o KpiAliasResponseOutput) AliasName() pulumi.StringOutput {
	return o.ApplyT(func(v KpiAliasResponse) string { return v.AliasName }).(pulumi.StringOutput)
}

func (o KpiAliasResponseOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v KpiAliasResponse) string { return v.Expression }).(pulumi.StringOutput)
}

type KpiAliasResponseArrayOutput struct{ *pulumi.OutputState }

func (KpiAliasResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiAliasResponse)(nil)).Elem()
}

func (o KpiAliasResponseArrayOutput) ToKpiAliasResponseArrayOutput() KpiAliasResponseArrayOutput {
	return o
}

func (o KpiAliasResponseArrayOutput) ToKpiAliasResponseArrayOutputWithContext(ctx context.Context) KpiAliasResponseArrayOutput {
	return o
}

func (o KpiAliasResponseArrayOutput) Index(i pulumi.IntInput) KpiAliasResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KpiAliasResponse {
		return vs[0].([]KpiAliasResponse)[vs[1].(int)]
	}).(KpiAliasResponseOutput)
}

type KpiExtract struct {
	Expression  string `pulumi:"expression"`
	ExtractName string `pulumi:"extractName"`
}





type KpiExtractInput interface {
	pulumi.Input

	ToKpiExtractOutput() KpiExtractOutput
	ToKpiExtractOutputWithContext(context.Context) KpiExtractOutput
}

type KpiExtractArgs struct {
	Expression  pulumi.StringInput `pulumi:"expression"`
	ExtractName pulumi.StringInput `pulumi:"extractName"`
}

func (KpiExtractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiExtract)(nil)).Elem()
}

func (i KpiExtractArgs) ToKpiExtractOutput() KpiExtractOutput {
	return i.ToKpiExtractOutputWithContext(context.Background())
}

func (i KpiExtractArgs) ToKpiExtractOutputWithContext(ctx context.Context) KpiExtractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiExtractOutput)
}





type KpiExtractArrayInput interface {
	pulumi.Input

	ToKpiExtractArrayOutput() KpiExtractArrayOutput
	ToKpiExtractArrayOutputWithContext(context.Context) KpiExtractArrayOutput
}

type KpiExtractArray []KpiExtractInput

func (KpiExtractArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiExtract)(nil)).Elem()
}

func (i KpiExtractArray) ToKpiExtractArrayOutput() KpiExtractArrayOutput {
	return i.ToKpiExtractArrayOutputWithContext(context.Background())
}

func (i KpiExtractArray) ToKpiExtractArrayOutputWithContext(ctx context.Context) KpiExtractArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiExtractArrayOutput)
}

type KpiExtractOutput struct{ *pulumi.OutputState }

func (KpiExtractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiExtract)(nil)).Elem()
}

func (o KpiExtractOutput) ToKpiExtractOutput() KpiExtractOutput {
	return o
}

func (o KpiExtractOutput) ToKpiExtractOutputWithContext(ctx context.Context) KpiExtractOutput {
	return o
}

func (o KpiExtractOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v KpiExtract) string { return v.Expression }).(pulumi.StringOutput)
}

func (o KpiExtractOutput) ExtractName() pulumi.StringOutput {
	return o.ApplyT(func(v KpiExtract) string { return v.ExtractName }).(pulumi.StringOutput)
}

type KpiExtractArrayOutput struct{ *pulumi.OutputState }

func (KpiExtractArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiExtract)(nil)).Elem()
}

func (o KpiExtractArrayOutput) ToKpiExtractArrayOutput() KpiExtractArrayOutput {
	return o
}

func (o KpiExtractArrayOutput) ToKpiExtractArrayOutputWithContext(ctx context.Context) KpiExtractArrayOutput {
	return o
}

func (o KpiExtractArrayOutput) Index(i pulumi.IntInput) KpiExtractOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KpiExtract {
		return vs[0].([]KpiExtract)[vs[1].(int)]
	}).(KpiExtractOutput)
}

type KpiExtractResponse struct {
	Expression  string `pulumi:"expression"`
	ExtractName string `pulumi:"extractName"`
}





type KpiExtractResponseInput interface {
	pulumi.Input

	ToKpiExtractResponseOutput() KpiExtractResponseOutput
	ToKpiExtractResponseOutputWithContext(context.Context) KpiExtractResponseOutput
}

type KpiExtractResponseArgs struct {
	Expression  pulumi.StringInput `pulumi:"expression"`
	ExtractName pulumi.StringInput `pulumi:"extractName"`
}

func (KpiExtractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiExtractResponse)(nil)).Elem()
}

func (i KpiExtractResponseArgs) ToKpiExtractResponseOutput() KpiExtractResponseOutput {
	return i.ToKpiExtractResponseOutputWithContext(context.Background())
}

func (i KpiExtractResponseArgs) ToKpiExtractResponseOutputWithContext(ctx context.Context) KpiExtractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiExtractResponseOutput)
}





type KpiExtractResponseArrayInput interface {
	pulumi.Input

	ToKpiExtractResponseArrayOutput() KpiExtractResponseArrayOutput
	ToKpiExtractResponseArrayOutputWithContext(context.Context) KpiExtractResponseArrayOutput
}

type KpiExtractResponseArray []KpiExtractResponseInput

func (KpiExtractResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiExtractResponse)(nil)).Elem()
}

func (i KpiExtractResponseArray) ToKpiExtractResponseArrayOutput() KpiExtractResponseArrayOutput {
	return i.ToKpiExtractResponseArrayOutputWithContext(context.Background())
}

func (i KpiExtractResponseArray) ToKpiExtractResponseArrayOutputWithContext(ctx context.Context) KpiExtractResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiExtractResponseArrayOutput)
}

type KpiExtractResponseOutput struct{ *pulumi.OutputState }

func (KpiExtractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiExtractResponse)(nil)).Elem()
}

func (o KpiExtractResponseOutput) ToKpiExtractResponseOutput() KpiExtractResponseOutput {
	return o
}

func (o KpiExtractResponseOutput) ToKpiExtractResponseOutputWithContext(ctx context.Context) KpiExtractResponseOutput {
	return o
}

func (o KpiExtractResponseOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v KpiExtractResponse) string { return v.Expression }).(pulumi.StringOutput)
}

func (o KpiExtractResponseOutput) ExtractName() pulumi.StringOutput {
	return o.ApplyT(func(v KpiExtractResponse) string { return v.ExtractName }).(pulumi.StringOutput)
}

type KpiExtractResponseArrayOutput struct{ *pulumi.OutputState }

func (KpiExtractResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiExtractResponse)(nil)).Elem()
}

func (o KpiExtractResponseArrayOutput) ToKpiExtractResponseArrayOutput() KpiExtractResponseArrayOutput {
	return o
}

func (o KpiExtractResponseArrayOutput) ToKpiExtractResponseArrayOutputWithContext(ctx context.Context) KpiExtractResponseArrayOutput {
	return o
}

func (o KpiExtractResponseArrayOutput) Index(i pulumi.IntInput) KpiExtractResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KpiExtractResponse {
		return vs[0].([]KpiExtractResponse)[vs[1].(int)]
	}).(KpiExtractResponseOutput)
}

type KpiGroupByMetadataResponse struct {
	DisplayName map[string]string `pulumi:"displayName"`
	FieldName   *string           `pulumi:"fieldName"`
	FieldType   *string           `pulumi:"fieldType"`
}





type KpiGroupByMetadataResponseInput interface {
	pulumi.Input

	ToKpiGroupByMetadataResponseOutput() KpiGroupByMetadataResponseOutput
	ToKpiGroupByMetadataResponseOutputWithContext(context.Context) KpiGroupByMetadataResponseOutput
}

type KpiGroupByMetadataResponseArgs struct {
	DisplayName pulumi.StringMapInput `pulumi:"displayName"`
	FieldName   pulumi.StringPtrInput `pulumi:"fieldName"`
	FieldType   pulumi.StringPtrInput `pulumi:"fieldType"`
}

func (KpiGroupByMetadataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiGroupByMetadataResponse)(nil)).Elem()
}

func (i KpiGroupByMetadataResponseArgs) ToKpiGroupByMetadataResponseOutput() KpiGroupByMetadataResponseOutput {
	return i.ToKpiGroupByMetadataResponseOutputWithContext(context.Background())
}

func (i KpiGroupByMetadataResponseArgs) ToKpiGroupByMetadataResponseOutputWithContext(ctx context.Context) KpiGroupByMetadataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiGroupByMetadataResponseOutput)
}





type KpiGroupByMetadataResponseArrayInput interface {
	pulumi.Input

	ToKpiGroupByMetadataResponseArrayOutput() KpiGroupByMetadataResponseArrayOutput
	ToKpiGroupByMetadataResponseArrayOutputWithContext(context.Context) KpiGroupByMetadataResponseArrayOutput
}

type KpiGroupByMetadataResponseArray []KpiGroupByMetadataResponseInput

func (KpiGroupByMetadataResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiGroupByMetadataResponse)(nil)).Elem()
}

func (i KpiGroupByMetadataResponseArray) ToKpiGroupByMetadataResponseArrayOutput() KpiGroupByMetadataResponseArrayOutput {
	return i.ToKpiGroupByMetadataResponseArrayOutputWithContext(context.Background())
}

func (i KpiGroupByMetadataResponseArray) ToKpiGroupByMetadataResponseArrayOutputWithContext(ctx context.Context) KpiGroupByMetadataResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiGroupByMetadataResponseArrayOutput)
}

type KpiGroupByMetadataResponseOutput struct{ *pulumi.OutputState }

func (KpiGroupByMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiGroupByMetadataResponse)(nil)).Elem()
}

func (o KpiGroupByMetadataResponseOutput) ToKpiGroupByMetadataResponseOutput() KpiGroupByMetadataResponseOutput {
	return o
}

func (o KpiGroupByMetadataResponseOutput) ToKpiGroupByMetadataResponseOutputWithContext(ctx context.Context) KpiGroupByMetadataResponseOutput {
	return o
}

func (o KpiGroupByMetadataResponseOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v KpiGroupByMetadataResponse) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o KpiGroupByMetadataResponseOutput) FieldName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KpiGroupByMetadataResponse) *string { return v.FieldName }).(pulumi.StringPtrOutput)
}

func (o KpiGroupByMetadataResponseOutput) FieldType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KpiGroupByMetadataResponse) *string { return v.FieldType }).(pulumi.StringPtrOutput)
}

type KpiGroupByMetadataResponseArrayOutput struct{ *pulumi.OutputState }

func (KpiGroupByMetadataResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiGroupByMetadataResponse)(nil)).Elem()
}

func (o KpiGroupByMetadataResponseArrayOutput) ToKpiGroupByMetadataResponseArrayOutput() KpiGroupByMetadataResponseArrayOutput {
	return o
}

func (o KpiGroupByMetadataResponseArrayOutput) ToKpiGroupByMetadataResponseArrayOutputWithContext(ctx context.Context) KpiGroupByMetadataResponseArrayOutput {
	return o
}

func (o KpiGroupByMetadataResponseArrayOutput) Index(i pulumi.IntInput) KpiGroupByMetadataResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KpiGroupByMetadataResponse {
		return vs[0].([]KpiGroupByMetadataResponse)[vs[1].(int)]
	}).(KpiGroupByMetadataResponseOutput)
}

type KpiParticipantProfilesMetadataResponse struct {
	TypeName string `pulumi:"typeName"`
}





type KpiParticipantProfilesMetadataResponseInput interface {
	pulumi.Input

	ToKpiParticipantProfilesMetadataResponseOutput() KpiParticipantProfilesMetadataResponseOutput
	ToKpiParticipantProfilesMetadataResponseOutputWithContext(context.Context) KpiParticipantProfilesMetadataResponseOutput
}

type KpiParticipantProfilesMetadataResponseArgs struct {
	TypeName pulumi.StringInput `pulumi:"typeName"`
}

func (KpiParticipantProfilesMetadataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiParticipantProfilesMetadataResponse)(nil)).Elem()
}

func (i KpiParticipantProfilesMetadataResponseArgs) ToKpiParticipantProfilesMetadataResponseOutput() KpiParticipantProfilesMetadataResponseOutput {
	return i.ToKpiParticipantProfilesMetadataResponseOutputWithContext(context.Background())
}

func (i KpiParticipantProfilesMetadataResponseArgs) ToKpiParticipantProfilesMetadataResponseOutputWithContext(ctx context.Context) KpiParticipantProfilesMetadataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiParticipantProfilesMetadataResponseOutput)
}





type KpiParticipantProfilesMetadataResponseArrayInput interface {
	pulumi.Input

	ToKpiParticipantProfilesMetadataResponseArrayOutput() KpiParticipantProfilesMetadataResponseArrayOutput
	ToKpiParticipantProfilesMetadataResponseArrayOutputWithContext(context.Context) KpiParticipantProfilesMetadataResponseArrayOutput
}

type KpiParticipantProfilesMetadataResponseArray []KpiParticipantProfilesMetadataResponseInput

func (KpiParticipantProfilesMetadataResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiParticipantProfilesMetadataResponse)(nil)).Elem()
}

func (i KpiParticipantProfilesMetadataResponseArray) ToKpiParticipantProfilesMetadataResponseArrayOutput() KpiParticipantProfilesMetadataResponseArrayOutput {
	return i.ToKpiParticipantProfilesMetadataResponseArrayOutputWithContext(context.Background())
}

func (i KpiParticipantProfilesMetadataResponseArray) ToKpiParticipantProfilesMetadataResponseArrayOutputWithContext(ctx context.Context) KpiParticipantProfilesMetadataResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiParticipantProfilesMetadataResponseArrayOutput)
}

type KpiParticipantProfilesMetadataResponseOutput struct{ *pulumi.OutputState }

func (KpiParticipantProfilesMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiParticipantProfilesMetadataResponse)(nil)).Elem()
}

func (o KpiParticipantProfilesMetadataResponseOutput) ToKpiParticipantProfilesMetadataResponseOutput() KpiParticipantProfilesMetadataResponseOutput {
	return o
}

func (o KpiParticipantProfilesMetadataResponseOutput) ToKpiParticipantProfilesMetadataResponseOutputWithContext(ctx context.Context) KpiParticipantProfilesMetadataResponseOutput {
	return o
}

func (o KpiParticipantProfilesMetadataResponseOutput) TypeName() pulumi.StringOutput {
	return o.ApplyT(func(v KpiParticipantProfilesMetadataResponse) string { return v.TypeName }).(pulumi.StringOutput)
}

type KpiParticipantProfilesMetadataResponseArrayOutput struct{ *pulumi.OutputState }

func (KpiParticipantProfilesMetadataResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiParticipantProfilesMetadataResponse)(nil)).Elem()
}

func (o KpiParticipantProfilesMetadataResponseArrayOutput) ToKpiParticipantProfilesMetadataResponseArrayOutput() KpiParticipantProfilesMetadataResponseArrayOutput {
	return o
}

func (o KpiParticipantProfilesMetadataResponseArrayOutput) ToKpiParticipantProfilesMetadataResponseArrayOutputWithContext(ctx context.Context) KpiParticipantProfilesMetadataResponseArrayOutput {
	return o
}

func (o KpiParticipantProfilesMetadataResponseArrayOutput) Index(i pulumi.IntInput) KpiParticipantProfilesMetadataResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KpiParticipantProfilesMetadataResponse {
		return vs[0].([]KpiParticipantProfilesMetadataResponse)[vs[1].(int)]
	}).(KpiParticipantProfilesMetadataResponseOutput)
}

type KpiThresholds struct {
	IncreasingKpi bool    `pulumi:"increasingKpi"`
	LowerLimit    float64 `pulumi:"lowerLimit"`
	UpperLimit    float64 `pulumi:"upperLimit"`
}





type KpiThresholdsInput interface {
	pulumi.Input

	ToKpiThresholdsOutput() KpiThresholdsOutput
	ToKpiThresholdsOutputWithContext(context.Context) KpiThresholdsOutput
}

type KpiThresholdsArgs struct {
	IncreasingKpi pulumi.BoolInput    `pulumi:"increasingKpi"`
	LowerLimit    pulumi.Float64Input `pulumi:"lowerLimit"`
	UpperLimit    pulumi.Float64Input `pulumi:"upperLimit"`
}

func (KpiThresholdsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiThresholds)(nil)).Elem()
}

func (i KpiThresholdsArgs) ToKpiThresholdsOutput() KpiThresholdsOutput {
	return i.ToKpiThresholdsOutputWithContext(context.Background())
}

func (i KpiThresholdsArgs) ToKpiThresholdsOutputWithContext(ctx context.Context) KpiThresholdsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiThresholdsOutput)
}

func (i KpiThresholdsArgs) ToKpiThresholdsPtrOutput() KpiThresholdsPtrOutput {
	return i.ToKpiThresholdsPtrOutputWithContext(context.Background())
}

func (i KpiThresholdsArgs) ToKpiThresholdsPtrOutputWithContext(ctx context.Context) KpiThresholdsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiThresholdsOutput).ToKpiThresholdsPtrOutputWithContext(ctx)
}









type KpiThresholdsPtrInput interface {
	pulumi.Input

	ToKpiThresholdsPtrOutput() KpiThresholdsPtrOutput
	ToKpiThresholdsPtrOutputWithContext(context.Context) KpiThresholdsPtrOutput
}

type kpiThresholdsPtrType KpiThresholdsArgs

func KpiThresholdsPtr(v *KpiThresholdsArgs) KpiThresholdsPtrInput {
	return (*kpiThresholdsPtrType)(v)
}

func (*kpiThresholdsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KpiThresholds)(nil)).Elem()
}

func (i *kpiThresholdsPtrType) ToKpiThresholdsPtrOutput() KpiThresholdsPtrOutput {
	return i.ToKpiThresholdsPtrOutputWithContext(context.Background())
}

func (i *kpiThresholdsPtrType) ToKpiThresholdsPtrOutputWithContext(ctx context.Context) KpiThresholdsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiThresholdsPtrOutput)
}

type KpiThresholdsOutput struct{ *pulumi.OutputState }

func (KpiThresholdsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiThresholds)(nil)).Elem()
}

func (o KpiThresholdsOutput) ToKpiThresholdsOutput() KpiThresholdsOutput {
	return o
}

func (o KpiThresholdsOutput) ToKpiThresholdsOutputWithContext(ctx context.Context) KpiThresholdsOutput {
	return o
}

func (o KpiThresholdsOutput) ToKpiThresholdsPtrOutput() KpiThresholdsPtrOutput {
	return o.ToKpiThresholdsPtrOutputWithContext(context.Background())
}

func (o KpiThresholdsOutput) ToKpiThresholdsPtrOutputWithContext(ctx context.Context) KpiThresholdsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KpiThresholds) *KpiThresholds {
		return &v
	}).(KpiThresholdsPtrOutput)
}

func (o KpiThresholdsOutput) IncreasingKpi() pulumi.BoolOutput {
	return o.ApplyT(func(v KpiThresholds) bool { return v.IncreasingKpi }).(pulumi.BoolOutput)
}

func (o KpiThresholdsOutput) LowerLimit() pulumi.Float64Output {
	return o.ApplyT(func(v KpiThresholds) float64 { return v.LowerLimit }).(pulumi.Float64Output)
}

func (o KpiThresholdsOutput) UpperLimit() pulumi.Float64Output {
	return o.ApplyT(func(v KpiThresholds) float64 { return v.UpperLimit }).(pulumi.Float64Output)
}

type KpiThresholdsPtrOutput struct{ *pulumi.OutputState }

func (KpiThresholdsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KpiThresholds)(nil)).Elem()
}

func (o KpiThresholdsPtrOutput) ToKpiThresholdsPtrOutput() KpiThresholdsPtrOutput {
	return o
}

func (o KpiThresholdsPtrOutput) ToKpiThresholdsPtrOutputWithContext(ctx context.Context) KpiThresholdsPtrOutput {
	return o
}

func (o KpiThresholdsPtrOutput) Elem() KpiThresholdsOutput {
	return o.ApplyT(func(v *KpiThresholds) KpiThresholds {
		if v != nil {
			return *v
		}
		var ret KpiThresholds
		return ret
	}).(KpiThresholdsOutput)
}

func (o KpiThresholdsPtrOutput) IncreasingKpi() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KpiThresholds) *bool {
		if v == nil {
			return nil
		}
		return &v.IncreasingKpi
	}).(pulumi.BoolPtrOutput)
}

func (o KpiThresholdsPtrOutput) LowerLimit() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KpiThresholds) *float64 {
		if v == nil {
			return nil
		}
		return &v.LowerLimit
	}).(pulumi.Float64PtrOutput)
}

func (o KpiThresholdsPtrOutput) UpperLimit() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KpiThresholds) *float64 {
		if v == nil {
			return nil
		}
		return &v.UpperLimit
	}).(pulumi.Float64PtrOutput)
}

type KpiThresholdsResponse struct {
	IncreasingKpi bool    `pulumi:"increasingKpi"`
	LowerLimit    float64 `pulumi:"lowerLimit"`
	UpperLimit    float64 `pulumi:"upperLimit"`
}





type KpiThresholdsResponseInput interface {
	pulumi.Input

	ToKpiThresholdsResponseOutput() KpiThresholdsResponseOutput
	ToKpiThresholdsResponseOutputWithContext(context.Context) KpiThresholdsResponseOutput
}

type KpiThresholdsResponseArgs struct {
	IncreasingKpi pulumi.BoolInput    `pulumi:"increasingKpi"`
	LowerLimit    pulumi.Float64Input `pulumi:"lowerLimit"`
	UpperLimit    pulumi.Float64Input `pulumi:"upperLimit"`
}

func (KpiThresholdsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiThresholdsResponse)(nil)).Elem()
}

func (i KpiThresholdsResponseArgs) ToKpiThresholdsResponseOutput() KpiThresholdsResponseOutput {
	return i.ToKpiThresholdsResponseOutputWithContext(context.Background())
}

func (i KpiThresholdsResponseArgs) ToKpiThresholdsResponseOutputWithContext(ctx context.Context) KpiThresholdsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiThresholdsResponseOutput)
}

func (i KpiThresholdsResponseArgs) ToKpiThresholdsResponsePtrOutput() KpiThresholdsResponsePtrOutput {
	return i.ToKpiThresholdsResponsePtrOutputWithContext(context.Background())
}

func (i KpiThresholdsResponseArgs) ToKpiThresholdsResponsePtrOutputWithContext(ctx context.Context) KpiThresholdsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiThresholdsResponseOutput).ToKpiThresholdsResponsePtrOutputWithContext(ctx)
}









type KpiThresholdsResponsePtrInput interface {
	pulumi.Input

	ToKpiThresholdsResponsePtrOutput() KpiThresholdsResponsePtrOutput
	ToKpiThresholdsResponsePtrOutputWithContext(context.Context) KpiThresholdsResponsePtrOutput
}

type kpiThresholdsResponsePtrType KpiThresholdsResponseArgs

func KpiThresholdsResponsePtr(v *KpiThresholdsResponseArgs) KpiThresholdsResponsePtrInput {
	return (*kpiThresholdsResponsePtrType)(v)
}

func (*kpiThresholdsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KpiThresholdsResponse)(nil)).Elem()
}

func (i *kpiThresholdsResponsePtrType) ToKpiThresholdsResponsePtrOutput() KpiThresholdsResponsePtrOutput {
	return i.ToKpiThresholdsResponsePtrOutputWithContext(context.Background())
}

func (i *kpiThresholdsResponsePtrType) ToKpiThresholdsResponsePtrOutputWithContext(ctx context.Context) KpiThresholdsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiThresholdsResponsePtrOutput)
}

type KpiThresholdsResponseOutput struct{ *pulumi.OutputState }

func (KpiThresholdsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiThresholdsResponse)(nil)).Elem()
}

func (o KpiThresholdsResponseOutput) ToKpiThresholdsResponseOutput() KpiThresholdsResponseOutput {
	return o
}

func (o KpiThresholdsResponseOutput) ToKpiThresholdsResponseOutputWithContext(ctx context.Context) KpiThresholdsResponseOutput {
	return o
}

func (o KpiThresholdsResponseOutput) ToKpiThresholdsResponsePtrOutput() KpiThresholdsResponsePtrOutput {
	return o.ToKpiThresholdsResponsePtrOutputWithContext(context.Background())
}

func (o KpiThresholdsResponseOutput) ToKpiThresholdsResponsePtrOutputWithContext(ctx context.Context) KpiThresholdsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KpiThresholdsResponse) *KpiThresholdsResponse {
		return &v
	}).(KpiThresholdsResponsePtrOutput)
}

func (o KpiThresholdsResponseOutput) IncreasingKpi() pulumi.BoolOutput {
	return o.ApplyT(func(v KpiThresholdsResponse) bool { return v.IncreasingKpi }).(pulumi.BoolOutput)
}

func (o KpiThresholdsResponseOutput) LowerLimit() pulumi.Float64Output {
	return o.ApplyT(func(v KpiThresholdsResponse) float64 { return v.LowerLimit }).(pulumi.Float64Output)
}

func (o KpiThresholdsResponseOutput) UpperLimit() pulumi.Float64Output {
	return o.ApplyT(func(v KpiThresholdsResponse) float64 { return v.UpperLimit }).(pulumi.Float64Output)
}

type KpiThresholdsResponsePtrOutput struct{ *pulumi.OutputState }

func (KpiThresholdsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KpiThresholdsResponse)(nil)).Elem()
}

func (o KpiThresholdsResponsePtrOutput) ToKpiThresholdsResponsePtrOutput() KpiThresholdsResponsePtrOutput {
	return o
}

func (o KpiThresholdsResponsePtrOutput) ToKpiThresholdsResponsePtrOutputWithContext(ctx context.Context) KpiThresholdsResponsePtrOutput {
	return o
}

func (o KpiThresholdsResponsePtrOutput) Elem() KpiThresholdsResponseOutput {
	return o.ApplyT(func(v *KpiThresholdsResponse) KpiThresholdsResponse {
		if v != nil {
			return *v
		}
		var ret KpiThresholdsResponse
		return ret
	}).(KpiThresholdsResponseOutput)
}

func (o KpiThresholdsResponsePtrOutput) IncreasingKpi() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KpiThresholdsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IncreasingKpi
	}).(pulumi.BoolPtrOutput)
}

func (o KpiThresholdsResponsePtrOutput) LowerLimit() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KpiThresholdsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.LowerLimit
	}).(pulumi.Float64PtrOutput)
}

func (o KpiThresholdsResponsePtrOutput) UpperLimit() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KpiThresholdsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.UpperLimit
	}).(pulumi.Float64PtrOutput)
}

type ParticipantPropertyReference struct {
	InteractionPropertyName string `pulumi:"interactionPropertyName"`
	ProfilePropertyName     string `pulumi:"profilePropertyName"`
}





type ParticipantPropertyReferenceInput interface {
	pulumi.Input

	ToParticipantPropertyReferenceOutput() ParticipantPropertyReferenceOutput
	ToParticipantPropertyReferenceOutputWithContext(context.Context) ParticipantPropertyReferenceOutput
}

type ParticipantPropertyReferenceArgs struct {
	InteractionPropertyName pulumi.StringInput `pulumi:"interactionPropertyName"`
	ProfilePropertyName     pulumi.StringInput `pulumi:"profilePropertyName"`
}

func (ParticipantPropertyReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParticipantPropertyReference)(nil)).Elem()
}

func (i ParticipantPropertyReferenceArgs) ToParticipantPropertyReferenceOutput() ParticipantPropertyReferenceOutput {
	return i.ToParticipantPropertyReferenceOutputWithContext(context.Background())
}

func (i ParticipantPropertyReferenceArgs) ToParticipantPropertyReferenceOutputWithContext(ctx context.Context) ParticipantPropertyReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParticipantPropertyReferenceOutput)
}





type ParticipantPropertyReferenceArrayInput interface {
	pulumi.Input

	ToParticipantPropertyReferenceArrayOutput() ParticipantPropertyReferenceArrayOutput
	ToParticipantPropertyReferenceArrayOutputWithContext(context.Context) ParticipantPropertyReferenceArrayOutput
}

type ParticipantPropertyReferenceArray []ParticipantPropertyReferenceInput

func (ParticipantPropertyReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParticipantPropertyReference)(nil)).Elem()
}

func (i ParticipantPropertyReferenceArray) ToParticipantPropertyReferenceArrayOutput() ParticipantPropertyReferenceArrayOutput {
	return i.ToParticipantPropertyReferenceArrayOutputWithContext(context.Background())
}

func (i ParticipantPropertyReferenceArray) ToParticipantPropertyReferenceArrayOutputWithContext(ctx context.Context) ParticipantPropertyReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParticipantPropertyReferenceArrayOutput)
}

type ParticipantPropertyReferenceOutput struct{ *pulumi.OutputState }

func (ParticipantPropertyReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParticipantPropertyReference)(nil)).Elem()
}

func (o ParticipantPropertyReferenceOutput) ToParticipantPropertyReferenceOutput() ParticipantPropertyReferenceOutput {
	return o
}

func (o ParticipantPropertyReferenceOutput) ToParticipantPropertyReferenceOutputWithContext(ctx context.Context) ParticipantPropertyReferenceOutput {
	return o
}

func (o ParticipantPropertyReferenceOutput) InteractionPropertyName() pulumi.StringOutput {
	return o.ApplyT(func(v ParticipantPropertyReference) string { return v.InteractionPropertyName }).(pulumi.StringOutput)
}

func (o ParticipantPropertyReferenceOutput) ProfilePropertyName() pulumi.StringOutput {
	return o.ApplyT(func(v ParticipantPropertyReference) string { return v.ProfilePropertyName }).(pulumi.StringOutput)
}

type ParticipantPropertyReferenceArrayOutput struct{ *pulumi.OutputState }

func (ParticipantPropertyReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParticipantPropertyReference)(nil)).Elem()
}

func (o ParticipantPropertyReferenceArrayOutput) ToParticipantPropertyReferenceArrayOutput() ParticipantPropertyReferenceArrayOutput {
	return o
}

func (o ParticipantPropertyReferenceArrayOutput) ToParticipantPropertyReferenceArrayOutputWithContext(ctx context.Context) ParticipantPropertyReferenceArrayOutput {
	return o
}

func (o ParticipantPropertyReferenceArrayOutput) Index(i pulumi.IntInput) ParticipantPropertyReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParticipantPropertyReference {
		return vs[0].([]ParticipantPropertyReference)[vs[1].(int)]
	}).(ParticipantPropertyReferenceOutput)
}

type ParticipantPropertyReferenceResponse struct {
	InteractionPropertyName string `pulumi:"interactionPropertyName"`
	ProfilePropertyName     string `pulumi:"profilePropertyName"`
}





type ParticipantPropertyReferenceResponseInput interface {
	pulumi.Input

	ToParticipantPropertyReferenceResponseOutput() ParticipantPropertyReferenceResponseOutput
	ToParticipantPropertyReferenceResponseOutputWithContext(context.Context) ParticipantPropertyReferenceResponseOutput
}

type ParticipantPropertyReferenceResponseArgs struct {
	InteractionPropertyName pulumi.StringInput `pulumi:"interactionPropertyName"`
	ProfilePropertyName     pulumi.StringInput `pulumi:"profilePropertyName"`
}

func (ParticipantPropertyReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParticipantPropertyReferenceResponse)(nil)).Elem()
}

func (i ParticipantPropertyReferenceResponseArgs) ToParticipantPropertyReferenceResponseOutput() ParticipantPropertyReferenceResponseOutput {
	return i.ToParticipantPropertyReferenceResponseOutputWithContext(context.Background())
}

func (i ParticipantPropertyReferenceResponseArgs) ToParticipantPropertyReferenceResponseOutputWithContext(ctx context.Context) ParticipantPropertyReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParticipantPropertyReferenceResponseOutput)
}





type ParticipantPropertyReferenceResponseArrayInput interface {
	pulumi.Input

	ToParticipantPropertyReferenceResponseArrayOutput() ParticipantPropertyReferenceResponseArrayOutput
	ToParticipantPropertyReferenceResponseArrayOutputWithContext(context.Context) ParticipantPropertyReferenceResponseArrayOutput
}

type ParticipantPropertyReferenceResponseArray []ParticipantPropertyReferenceResponseInput

func (ParticipantPropertyReferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParticipantPropertyReferenceResponse)(nil)).Elem()
}

func (i ParticipantPropertyReferenceResponseArray) ToParticipantPropertyReferenceResponseArrayOutput() ParticipantPropertyReferenceResponseArrayOutput {
	return i.ToParticipantPropertyReferenceResponseArrayOutputWithContext(context.Background())
}

func (i ParticipantPropertyReferenceResponseArray) ToParticipantPropertyReferenceResponseArrayOutputWithContext(ctx context.Context) ParticipantPropertyReferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParticipantPropertyReferenceResponseArrayOutput)
}

type ParticipantPropertyReferenceResponseOutput struct{ *pulumi.OutputState }

func (ParticipantPropertyReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParticipantPropertyReferenceResponse)(nil)).Elem()
}

func (o ParticipantPropertyReferenceResponseOutput) ToParticipantPropertyReferenceResponseOutput() ParticipantPropertyReferenceResponseOutput {
	return o
}

func (o ParticipantPropertyReferenceResponseOutput) ToParticipantPropertyReferenceResponseOutputWithContext(ctx context.Context) ParticipantPropertyReferenceResponseOutput {
	return o
}

func (o ParticipantPropertyReferenceResponseOutput) InteractionPropertyName() pulumi.StringOutput {
	return o.ApplyT(func(v ParticipantPropertyReferenceResponse) string { return v.InteractionPropertyName }).(pulumi.StringOutput)
}

func (o ParticipantPropertyReferenceResponseOutput) ProfilePropertyName() pulumi.StringOutput {
	return o.ApplyT(func(v ParticipantPropertyReferenceResponse) string { return v.ProfilePropertyName }).(pulumi.StringOutput)
}

type ParticipantPropertyReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (ParticipantPropertyReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParticipantPropertyReferenceResponse)(nil)).Elem()
}

func (o ParticipantPropertyReferenceResponseArrayOutput) ToParticipantPropertyReferenceResponseArrayOutput() ParticipantPropertyReferenceResponseArrayOutput {
	return o
}

func (o ParticipantPropertyReferenceResponseArrayOutput) ToParticipantPropertyReferenceResponseArrayOutputWithContext(ctx context.Context) ParticipantPropertyReferenceResponseArrayOutput {
	return o
}

func (o ParticipantPropertyReferenceResponseArrayOutput) Index(i pulumi.IntInput) ParticipantPropertyReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParticipantPropertyReferenceResponse {
		return vs[0].([]ParticipantPropertyReferenceResponse)[vs[1].(int)]
	}).(ParticipantPropertyReferenceResponseOutput)
}

type ProfileEnumValidValuesFormat struct {
	LocalizedValueNames map[string]string `pulumi:"localizedValueNames"`
	Value               *int              `pulumi:"value"`
}





type ProfileEnumValidValuesFormatInput interface {
	pulumi.Input

	ToProfileEnumValidValuesFormatOutput() ProfileEnumValidValuesFormatOutput
	ToProfileEnumValidValuesFormatOutputWithContext(context.Context) ProfileEnumValidValuesFormatOutput
}

type ProfileEnumValidValuesFormatArgs struct {
	LocalizedValueNames pulumi.StringMapInput `pulumi:"localizedValueNames"`
	Value               pulumi.IntPtrInput    `pulumi:"value"`
}

func (ProfileEnumValidValuesFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileEnumValidValuesFormat)(nil)).Elem()
}

func (i ProfileEnumValidValuesFormatArgs) ToProfileEnumValidValuesFormatOutput() ProfileEnumValidValuesFormatOutput {
	return i.ToProfileEnumValidValuesFormatOutputWithContext(context.Background())
}

func (i ProfileEnumValidValuesFormatArgs) ToProfileEnumValidValuesFormatOutputWithContext(ctx context.Context) ProfileEnumValidValuesFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileEnumValidValuesFormatOutput)
}





type ProfileEnumValidValuesFormatArrayInput interface {
	pulumi.Input

	ToProfileEnumValidValuesFormatArrayOutput() ProfileEnumValidValuesFormatArrayOutput
	ToProfileEnumValidValuesFormatArrayOutputWithContext(context.Context) ProfileEnumValidValuesFormatArrayOutput
}

type ProfileEnumValidValuesFormatArray []ProfileEnumValidValuesFormatInput

func (ProfileEnumValidValuesFormatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileEnumValidValuesFormat)(nil)).Elem()
}

func (i ProfileEnumValidValuesFormatArray) ToProfileEnumValidValuesFormatArrayOutput() ProfileEnumValidValuesFormatArrayOutput {
	return i.ToProfileEnumValidValuesFormatArrayOutputWithContext(context.Background())
}

func (i ProfileEnumValidValuesFormatArray) ToProfileEnumValidValuesFormatArrayOutputWithContext(ctx context.Context) ProfileEnumValidValuesFormatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileEnumValidValuesFormatArrayOutput)
}

type ProfileEnumValidValuesFormatOutput struct{ *pulumi.OutputState }

func (ProfileEnumValidValuesFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileEnumValidValuesFormat)(nil)).Elem()
}

func (o ProfileEnumValidValuesFormatOutput) ToProfileEnumValidValuesFormatOutput() ProfileEnumValidValuesFormatOutput {
	return o
}

func (o ProfileEnumValidValuesFormatOutput) ToProfileEnumValidValuesFormatOutputWithContext(ctx context.Context) ProfileEnumValidValuesFormatOutput {
	return o
}

func (o ProfileEnumValidValuesFormatOutput) LocalizedValueNames() pulumi.StringMapOutput {
	return o.ApplyT(func(v ProfileEnumValidValuesFormat) map[string]string { return v.LocalizedValueNames }).(pulumi.StringMapOutput)
}

func (o ProfileEnumValidValuesFormatOutput) Value() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProfileEnumValidValuesFormat) *int { return v.Value }).(pulumi.IntPtrOutput)
}

type ProfileEnumValidValuesFormatArrayOutput struct{ *pulumi.OutputState }

func (ProfileEnumValidValuesFormatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileEnumValidValuesFormat)(nil)).Elem()
}

func (o ProfileEnumValidValuesFormatArrayOutput) ToProfileEnumValidValuesFormatArrayOutput() ProfileEnumValidValuesFormatArrayOutput {
	return o
}

func (o ProfileEnumValidValuesFormatArrayOutput) ToProfileEnumValidValuesFormatArrayOutputWithContext(ctx context.Context) ProfileEnumValidValuesFormatArrayOutput {
	return o
}

func (o ProfileEnumValidValuesFormatArrayOutput) Index(i pulumi.IntInput) ProfileEnumValidValuesFormatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileEnumValidValuesFormat {
		return vs[0].([]ProfileEnumValidValuesFormat)[vs[1].(int)]
	}).(ProfileEnumValidValuesFormatOutput)
}

type ProfileEnumValidValuesFormatResponse struct {
	LocalizedValueNames map[string]string `pulumi:"localizedValueNames"`
	Value               *int              `pulumi:"value"`
}





type ProfileEnumValidValuesFormatResponseInput interface {
	pulumi.Input

	ToProfileEnumValidValuesFormatResponseOutput() ProfileEnumValidValuesFormatResponseOutput
	ToProfileEnumValidValuesFormatResponseOutputWithContext(context.Context) ProfileEnumValidValuesFormatResponseOutput
}

type ProfileEnumValidValuesFormatResponseArgs struct {
	LocalizedValueNames pulumi.StringMapInput `pulumi:"localizedValueNames"`
	Value               pulumi.IntPtrInput    `pulumi:"value"`
}

func (ProfileEnumValidValuesFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileEnumValidValuesFormatResponse)(nil)).Elem()
}

func (i ProfileEnumValidValuesFormatResponseArgs) ToProfileEnumValidValuesFormatResponseOutput() ProfileEnumValidValuesFormatResponseOutput {
	return i.ToProfileEnumValidValuesFormatResponseOutputWithContext(context.Background())
}

func (i ProfileEnumValidValuesFormatResponseArgs) ToProfileEnumValidValuesFormatResponseOutputWithContext(ctx context.Context) ProfileEnumValidValuesFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileEnumValidValuesFormatResponseOutput)
}





type ProfileEnumValidValuesFormatResponseArrayInput interface {
	pulumi.Input

	ToProfileEnumValidValuesFormatResponseArrayOutput() ProfileEnumValidValuesFormatResponseArrayOutput
	ToProfileEnumValidValuesFormatResponseArrayOutputWithContext(context.Context) ProfileEnumValidValuesFormatResponseArrayOutput
}

type ProfileEnumValidValuesFormatResponseArray []ProfileEnumValidValuesFormatResponseInput

func (ProfileEnumValidValuesFormatResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileEnumValidValuesFormatResponse)(nil)).Elem()
}

func (i ProfileEnumValidValuesFormatResponseArray) ToProfileEnumValidValuesFormatResponseArrayOutput() ProfileEnumValidValuesFormatResponseArrayOutput {
	return i.ToProfileEnumValidValuesFormatResponseArrayOutputWithContext(context.Background())
}

func (i ProfileEnumValidValuesFormatResponseArray) ToProfileEnumValidValuesFormatResponseArrayOutputWithContext(ctx context.Context) ProfileEnumValidValuesFormatResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileEnumValidValuesFormatResponseArrayOutput)
}

type ProfileEnumValidValuesFormatResponseOutput struct{ *pulumi.OutputState }

func (ProfileEnumValidValuesFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileEnumValidValuesFormatResponse)(nil)).Elem()
}

func (o ProfileEnumValidValuesFormatResponseOutput) ToProfileEnumValidValuesFormatResponseOutput() ProfileEnumValidValuesFormatResponseOutput {
	return o
}

func (o ProfileEnumValidValuesFormatResponseOutput) ToProfileEnumValidValuesFormatResponseOutputWithContext(ctx context.Context) ProfileEnumValidValuesFormatResponseOutput {
	return o
}

func (o ProfileEnumValidValuesFormatResponseOutput) LocalizedValueNames() pulumi.StringMapOutput {
	return o.ApplyT(func(v ProfileEnumValidValuesFormatResponse) map[string]string { return v.LocalizedValueNames }).(pulumi.StringMapOutput)
}

func (o ProfileEnumValidValuesFormatResponseOutput) Value() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProfileEnumValidValuesFormatResponse) *int { return v.Value }).(pulumi.IntPtrOutput)
}

type ProfileEnumValidValuesFormatResponseArrayOutput struct{ *pulumi.OutputState }

func (ProfileEnumValidValuesFormatResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileEnumValidValuesFormatResponse)(nil)).Elem()
}

func (o ProfileEnumValidValuesFormatResponseArrayOutput) ToProfileEnumValidValuesFormatResponseArrayOutput() ProfileEnumValidValuesFormatResponseArrayOutput {
	return o
}

func (o ProfileEnumValidValuesFormatResponseArrayOutput) ToProfileEnumValidValuesFormatResponseArrayOutputWithContext(ctx context.Context) ProfileEnumValidValuesFormatResponseArrayOutput {
	return o
}

func (o ProfileEnumValidValuesFormatResponseArrayOutput) Index(i pulumi.IntInput) ProfileEnumValidValuesFormatResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileEnumValidValuesFormatResponse {
		return vs[0].([]ProfileEnumValidValuesFormatResponse)[vs[1].(int)]
	}).(ProfileEnumValidValuesFormatResponseOutput)
}

type PropertyDefinition struct {
	ArrayValueSeparator *string                        `pulumi:"arrayValueSeparator"`
	EnumValidValues     []ProfileEnumValidValuesFormat `pulumi:"enumValidValues"`
	FieldName           string                         `pulumi:"fieldName"`
	FieldType           string                         `pulumi:"fieldType"`
	IsArray             *bool                          `pulumi:"isArray"`
	IsAvailableInGraph  *bool                          `pulumi:"isAvailableInGraph"`
	IsEnum              *bool                          `pulumi:"isEnum"`
	IsFlagEnum          *bool                          `pulumi:"isFlagEnum"`
	IsImage             *bool                          `pulumi:"isImage"`
	IsLocalizedString   *bool                          `pulumi:"isLocalizedString"`
	IsName              *bool                          `pulumi:"isName"`
	IsRequired          *bool                          `pulumi:"isRequired"`
	MaxLength           *int                           `pulumi:"maxLength"`
	PropertyId          *string                        `pulumi:"propertyId"`
	SchemaItemPropLink  *string                        `pulumi:"schemaItemPropLink"`
}





type PropertyDefinitionInput interface {
	pulumi.Input

	ToPropertyDefinitionOutput() PropertyDefinitionOutput
	ToPropertyDefinitionOutputWithContext(context.Context) PropertyDefinitionOutput
}

type PropertyDefinitionArgs struct {
	ArrayValueSeparator pulumi.StringPtrInput                  `pulumi:"arrayValueSeparator"`
	EnumValidValues     ProfileEnumValidValuesFormatArrayInput `pulumi:"enumValidValues"`
	FieldName           pulumi.StringInput                     `pulumi:"fieldName"`
	FieldType           pulumi.StringInput                     `pulumi:"fieldType"`
	IsArray             pulumi.BoolPtrInput                    `pulumi:"isArray"`
	IsAvailableInGraph  pulumi.BoolPtrInput                    `pulumi:"isAvailableInGraph"`
	IsEnum              pulumi.BoolPtrInput                    `pulumi:"isEnum"`
	IsFlagEnum          pulumi.BoolPtrInput                    `pulumi:"isFlagEnum"`
	IsImage             pulumi.BoolPtrInput                    `pulumi:"isImage"`
	IsLocalizedString   pulumi.BoolPtrInput                    `pulumi:"isLocalizedString"`
	IsName              pulumi.BoolPtrInput                    `pulumi:"isName"`
	IsRequired          pulumi.BoolPtrInput                    `pulumi:"isRequired"`
	MaxLength           pulumi.IntPtrInput                     `pulumi:"maxLength"`
	PropertyId          pulumi.StringPtrInput                  `pulumi:"propertyId"`
	SchemaItemPropLink  pulumi.StringPtrInput                  `pulumi:"schemaItemPropLink"`
}

func (PropertyDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertyDefinition)(nil)).Elem()
}

func (i PropertyDefinitionArgs) ToPropertyDefinitionOutput() PropertyDefinitionOutput {
	return i.ToPropertyDefinitionOutputWithContext(context.Background())
}

func (i PropertyDefinitionArgs) ToPropertyDefinitionOutputWithContext(ctx context.Context) PropertyDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyDefinitionOutput)
}





type PropertyDefinitionArrayInput interface {
	pulumi.Input

	ToPropertyDefinitionArrayOutput() PropertyDefinitionArrayOutput
	ToPropertyDefinitionArrayOutputWithContext(context.Context) PropertyDefinitionArrayOutput
}

type PropertyDefinitionArray []PropertyDefinitionInput

func (PropertyDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PropertyDefinition)(nil)).Elem()
}

func (i PropertyDefinitionArray) ToPropertyDefinitionArrayOutput() PropertyDefinitionArrayOutput {
	return i.ToPropertyDefinitionArrayOutputWithContext(context.Background())
}

func (i PropertyDefinitionArray) ToPropertyDefinitionArrayOutputWithContext(ctx context.Context) PropertyDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyDefinitionArrayOutput)
}

type PropertyDefinitionOutput struct{ *pulumi.OutputState }

func (PropertyDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertyDefinition)(nil)).Elem()
}

func (o PropertyDefinitionOutput) ToPropertyDefinitionOutput() PropertyDefinitionOutput {
	return o
}

func (o PropertyDefinitionOutput) ToPropertyDefinitionOutputWithContext(ctx context.Context) PropertyDefinitionOutput {
	return o
}

func (o PropertyDefinitionOutput) ArrayValueSeparator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PropertyDefinition) *string { return v.ArrayValueSeparator }).(pulumi.StringPtrOutput)
}

func (o PropertyDefinitionOutput) EnumValidValues() ProfileEnumValidValuesFormatArrayOutput {
	return o.ApplyT(func(v PropertyDefinition) []ProfileEnumValidValuesFormat { return v.EnumValidValues }).(ProfileEnumValidValuesFormatArrayOutput)
}

func (o PropertyDefinitionOutput) FieldName() pulumi.StringOutput {
	return o.ApplyT(func(v PropertyDefinition) string { return v.FieldName }).(pulumi.StringOutput)
}

func (o PropertyDefinitionOutput) FieldType() pulumi.StringOutput {
	return o.ApplyT(func(v PropertyDefinition) string { return v.FieldType }).(pulumi.StringOutput)
}

func (o PropertyDefinitionOutput) IsArray() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinition) *bool { return v.IsArray }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionOutput) IsAvailableInGraph() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinition) *bool { return v.IsAvailableInGraph }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionOutput) IsEnum() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinition) *bool { return v.IsEnum }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionOutput) IsFlagEnum() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinition) *bool { return v.IsFlagEnum }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionOutput) IsImage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinition) *bool { return v.IsImage }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionOutput) IsLocalizedString() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinition) *bool { return v.IsLocalizedString }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionOutput) IsName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinition) *bool { return v.IsName }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionOutput) IsRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinition) *bool { return v.IsRequired }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionOutput) MaxLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PropertyDefinition) *int { return v.MaxLength }).(pulumi.IntPtrOutput)
}

func (o PropertyDefinitionOutput) PropertyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PropertyDefinition) *string { return v.PropertyId }).(pulumi.StringPtrOutput)
}

func (o PropertyDefinitionOutput) SchemaItemPropLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PropertyDefinition) *string { return v.SchemaItemPropLink }).(pulumi.StringPtrOutput)
}

type PropertyDefinitionArrayOutput struct{ *pulumi.OutputState }

func (PropertyDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PropertyDefinition)(nil)).Elem()
}

func (o PropertyDefinitionArrayOutput) ToPropertyDefinitionArrayOutput() PropertyDefinitionArrayOutput {
	return o
}

func (o PropertyDefinitionArrayOutput) ToPropertyDefinitionArrayOutputWithContext(ctx context.Context) PropertyDefinitionArrayOutput {
	return o
}

func (o PropertyDefinitionArrayOutput) Index(i pulumi.IntInput) PropertyDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PropertyDefinition {
		return vs[0].([]PropertyDefinition)[vs[1].(int)]
	}).(PropertyDefinitionOutput)
}

type PropertyDefinitionResponse struct {
	ArrayValueSeparator       *string                                `pulumi:"arrayValueSeparator"`
	DataSourcePrecedenceRules []DataSourcePrecedenceResponse         `pulumi:"dataSourcePrecedenceRules"`
	EnumValidValues           []ProfileEnumValidValuesFormatResponse `pulumi:"enumValidValues"`
	FieldName                 string                                 `pulumi:"fieldName"`
	FieldType                 string                                 `pulumi:"fieldType"`
	IsArray                   *bool                                  `pulumi:"isArray"`
	IsAvailableInGraph        *bool                                  `pulumi:"isAvailableInGraph"`
	IsEnum                    *bool                                  `pulumi:"isEnum"`
	IsFlagEnum                *bool                                  `pulumi:"isFlagEnum"`
	IsImage                   *bool                                  `pulumi:"isImage"`
	IsLocalizedString         *bool                                  `pulumi:"isLocalizedString"`
	IsName                    *bool                                  `pulumi:"isName"`
	IsRequired                *bool                                  `pulumi:"isRequired"`
	MaxLength                 *int                                   `pulumi:"maxLength"`
	PropertyId                *string                                `pulumi:"propertyId"`
	SchemaItemPropLink        *string                                `pulumi:"schemaItemPropLink"`
}





type PropertyDefinitionResponseInput interface {
	pulumi.Input

	ToPropertyDefinitionResponseOutput() PropertyDefinitionResponseOutput
	ToPropertyDefinitionResponseOutputWithContext(context.Context) PropertyDefinitionResponseOutput
}

type PropertyDefinitionResponseArgs struct {
	ArrayValueSeparator       pulumi.StringPtrInput                          `pulumi:"arrayValueSeparator"`
	DataSourcePrecedenceRules DataSourcePrecedenceResponseArrayInput         `pulumi:"dataSourcePrecedenceRules"`
	EnumValidValues           ProfileEnumValidValuesFormatResponseArrayInput `pulumi:"enumValidValues"`
	FieldName                 pulumi.StringInput                             `pulumi:"fieldName"`
	FieldType                 pulumi.StringInput                             `pulumi:"fieldType"`
	IsArray                   pulumi.BoolPtrInput                            `pulumi:"isArray"`
	IsAvailableInGraph        pulumi.BoolPtrInput                            `pulumi:"isAvailableInGraph"`
	IsEnum                    pulumi.BoolPtrInput                            `pulumi:"isEnum"`
	IsFlagEnum                pulumi.BoolPtrInput                            `pulumi:"isFlagEnum"`
	IsImage                   pulumi.BoolPtrInput                            `pulumi:"isImage"`
	IsLocalizedString         pulumi.BoolPtrInput                            `pulumi:"isLocalizedString"`
	IsName                    pulumi.BoolPtrInput                            `pulumi:"isName"`
	IsRequired                pulumi.BoolPtrInput                            `pulumi:"isRequired"`
	MaxLength                 pulumi.IntPtrInput                             `pulumi:"maxLength"`
	PropertyId                pulumi.StringPtrInput                          `pulumi:"propertyId"`
	SchemaItemPropLink        pulumi.StringPtrInput                          `pulumi:"schemaItemPropLink"`
}

func (PropertyDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertyDefinitionResponse)(nil)).Elem()
}

func (i PropertyDefinitionResponseArgs) ToPropertyDefinitionResponseOutput() PropertyDefinitionResponseOutput {
	return i.ToPropertyDefinitionResponseOutputWithContext(context.Background())
}

func (i PropertyDefinitionResponseArgs) ToPropertyDefinitionResponseOutputWithContext(ctx context.Context) PropertyDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyDefinitionResponseOutput)
}





type PropertyDefinitionResponseArrayInput interface {
	pulumi.Input

	ToPropertyDefinitionResponseArrayOutput() PropertyDefinitionResponseArrayOutput
	ToPropertyDefinitionResponseArrayOutputWithContext(context.Context) PropertyDefinitionResponseArrayOutput
}

type PropertyDefinitionResponseArray []PropertyDefinitionResponseInput

func (PropertyDefinitionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PropertyDefinitionResponse)(nil)).Elem()
}

func (i PropertyDefinitionResponseArray) ToPropertyDefinitionResponseArrayOutput() PropertyDefinitionResponseArrayOutput {
	return i.ToPropertyDefinitionResponseArrayOutputWithContext(context.Background())
}

func (i PropertyDefinitionResponseArray) ToPropertyDefinitionResponseArrayOutputWithContext(ctx context.Context) PropertyDefinitionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyDefinitionResponseArrayOutput)
}

type PropertyDefinitionResponseOutput struct{ *pulumi.OutputState }

func (PropertyDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertyDefinitionResponse)(nil)).Elem()
}

func (o PropertyDefinitionResponseOutput) ToPropertyDefinitionResponseOutput() PropertyDefinitionResponseOutput {
	return o
}

func (o PropertyDefinitionResponseOutput) ToPropertyDefinitionResponseOutputWithContext(ctx context.Context) PropertyDefinitionResponseOutput {
	return o
}

func (o PropertyDefinitionResponseOutput) ArrayValueSeparator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) *string { return v.ArrayValueSeparator }).(pulumi.StringPtrOutput)
}

func (o PropertyDefinitionResponseOutput) DataSourcePrecedenceRules() DataSourcePrecedenceResponseArrayOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) []DataSourcePrecedenceResponse { return v.DataSourcePrecedenceRules }).(DataSourcePrecedenceResponseArrayOutput)
}

func (o PropertyDefinitionResponseOutput) EnumValidValues() ProfileEnumValidValuesFormatResponseArrayOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) []ProfileEnumValidValuesFormatResponse { return v.EnumValidValues }).(ProfileEnumValidValuesFormatResponseArrayOutput)
}

func (o PropertyDefinitionResponseOutput) FieldName() pulumi.StringOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) string { return v.FieldName }).(pulumi.StringOutput)
}

func (o PropertyDefinitionResponseOutput) FieldType() pulumi.StringOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) string { return v.FieldType }).(pulumi.StringOutput)
}

func (o PropertyDefinitionResponseOutput) IsArray() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) *bool { return v.IsArray }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionResponseOutput) IsAvailableInGraph() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) *bool { return v.IsAvailableInGraph }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionResponseOutput) IsEnum() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) *bool { return v.IsEnum }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionResponseOutput) IsFlagEnum() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) *bool { return v.IsFlagEnum }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionResponseOutput) IsImage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) *bool { return v.IsImage }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionResponseOutput) IsLocalizedString() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) *bool { return v.IsLocalizedString }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionResponseOutput) IsName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) *bool { return v.IsName }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionResponseOutput) IsRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) *bool { return v.IsRequired }).(pulumi.BoolPtrOutput)
}

func (o PropertyDefinitionResponseOutput) MaxLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) *int { return v.MaxLength }).(pulumi.IntPtrOutput)
}

func (o PropertyDefinitionResponseOutput) PropertyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) *string { return v.PropertyId }).(pulumi.StringPtrOutput)
}

func (o PropertyDefinitionResponseOutput) SchemaItemPropLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PropertyDefinitionResponse) *string { return v.SchemaItemPropLink }).(pulumi.StringPtrOutput)
}

type PropertyDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (PropertyDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PropertyDefinitionResponse)(nil)).Elem()
}

func (o PropertyDefinitionResponseArrayOutput) ToPropertyDefinitionResponseArrayOutput() PropertyDefinitionResponseArrayOutput {
	return o
}

func (o PropertyDefinitionResponseArrayOutput) ToPropertyDefinitionResponseArrayOutputWithContext(ctx context.Context) PropertyDefinitionResponseArrayOutput {
	return o
}

func (o PropertyDefinitionResponseArrayOutput) Index(i pulumi.IntInput) PropertyDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PropertyDefinitionResponse {
		return vs[0].([]PropertyDefinitionResponse)[vs[1].(int)]
	}).(PropertyDefinitionResponseOutput)
}

type RelationshipLinkFieldMapping struct {
	InteractionFieldName  string     `pulumi:"interactionFieldName"`
	LinkType              *LinkTypes `pulumi:"linkType"`
	RelationshipFieldName string     `pulumi:"relationshipFieldName"`
}





type RelationshipLinkFieldMappingInput interface {
	pulumi.Input

	ToRelationshipLinkFieldMappingOutput() RelationshipLinkFieldMappingOutput
	ToRelationshipLinkFieldMappingOutputWithContext(context.Context) RelationshipLinkFieldMappingOutput
}

type RelationshipLinkFieldMappingArgs struct {
	InteractionFieldName  pulumi.StringInput `pulumi:"interactionFieldName"`
	LinkType              LinkTypesPtrInput  `pulumi:"linkType"`
	RelationshipFieldName pulumi.StringInput `pulumi:"relationshipFieldName"`
}

func (RelationshipLinkFieldMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipLinkFieldMapping)(nil)).Elem()
}

func (i RelationshipLinkFieldMappingArgs) ToRelationshipLinkFieldMappingOutput() RelationshipLinkFieldMappingOutput {
	return i.ToRelationshipLinkFieldMappingOutputWithContext(context.Background())
}

func (i RelationshipLinkFieldMappingArgs) ToRelationshipLinkFieldMappingOutputWithContext(ctx context.Context) RelationshipLinkFieldMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipLinkFieldMappingOutput)
}





type RelationshipLinkFieldMappingArrayInput interface {
	pulumi.Input

	ToRelationshipLinkFieldMappingArrayOutput() RelationshipLinkFieldMappingArrayOutput
	ToRelationshipLinkFieldMappingArrayOutputWithContext(context.Context) RelationshipLinkFieldMappingArrayOutput
}

type RelationshipLinkFieldMappingArray []RelationshipLinkFieldMappingInput

func (RelationshipLinkFieldMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RelationshipLinkFieldMapping)(nil)).Elem()
}

func (i RelationshipLinkFieldMappingArray) ToRelationshipLinkFieldMappingArrayOutput() RelationshipLinkFieldMappingArrayOutput {
	return i.ToRelationshipLinkFieldMappingArrayOutputWithContext(context.Background())
}

func (i RelationshipLinkFieldMappingArray) ToRelationshipLinkFieldMappingArrayOutputWithContext(ctx context.Context) RelationshipLinkFieldMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipLinkFieldMappingArrayOutput)
}

type RelationshipLinkFieldMappingOutput struct{ *pulumi.OutputState }

func (RelationshipLinkFieldMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipLinkFieldMapping)(nil)).Elem()
}

func (o RelationshipLinkFieldMappingOutput) ToRelationshipLinkFieldMappingOutput() RelationshipLinkFieldMappingOutput {
	return o
}

func (o RelationshipLinkFieldMappingOutput) ToRelationshipLinkFieldMappingOutputWithContext(ctx context.Context) RelationshipLinkFieldMappingOutput {
	return o
}

func (o RelationshipLinkFieldMappingOutput) InteractionFieldName() pulumi.StringOutput {
	return o.ApplyT(func(v RelationshipLinkFieldMapping) string { return v.InteractionFieldName }).(pulumi.StringOutput)
}

func (o RelationshipLinkFieldMappingOutput) LinkType() LinkTypesPtrOutput {
	return o.ApplyT(func(v RelationshipLinkFieldMapping) *LinkTypes { return v.LinkType }).(LinkTypesPtrOutput)
}

func (o RelationshipLinkFieldMappingOutput) RelationshipFieldName() pulumi.StringOutput {
	return o.ApplyT(func(v RelationshipLinkFieldMapping) string { return v.RelationshipFieldName }).(pulumi.StringOutput)
}

type RelationshipLinkFieldMappingArrayOutput struct{ *pulumi.OutputState }

func (RelationshipLinkFieldMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RelationshipLinkFieldMapping)(nil)).Elem()
}

func (o RelationshipLinkFieldMappingArrayOutput) ToRelationshipLinkFieldMappingArrayOutput() RelationshipLinkFieldMappingArrayOutput {
	return o
}

func (o RelationshipLinkFieldMappingArrayOutput) ToRelationshipLinkFieldMappingArrayOutputWithContext(ctx context.Context) RelationshipLinkFieldMappingArrayOutput {
	return o
}

func (o RelationshipLinkFieldMappingArrayOutput) Index(i pulumi.IntInput) RelationshipLinkFieldMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RelationshipLinkFieldMapping {
		return vs[0].([]RelationshipLinkFieldMapping)[vs[1].(int)]
	}).(RelationshipLinkFieldMappingOutput)
}

type RelationshipLinkFieldMappingResponse struct {
	InteractionFieldName  string  `pulumi:"interactionFieldName"`
	LinkType              *string `pulumi:"linkType"`
	RelationshipFieldName string  `pulumi:"relationshipFieldName"`
}





type RelationshipLinkFieldMappingResponseInput interface {
	pulumi.Input

	ToRelationshipLinkFieldMappingResponseOutput() RelationshipLinkFieldMappingResponseOutput
	ToRelationshipLinkFieldMappingResponseOutputWithContext(context.Context) RelationshipLinkFieldMappingResponseOutput
}

type RelationshipLinkFieldMappingResponseArgs struct {
	InteractionFieldName  pulumi.StringInput    `pulumi:"interactionFieldName"`
	LinkType              pulumi.StringPtrInput `pulumi:"linkType"`
	RelationshipFieldName pulumi.StringInput    `pulumi:"relationshipFieldName"`
}

func (RelationshipLinkFieldMappingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipLinkFieldMappingResponse)(nil)).Elem()
}

func (i RelationshipLinkFieldMappingResponseArgs) ToRelationshipLinkFieldMappingResponseOutput() RelationshipLinkFieldMappingResponseOutput {
	return i.ToRelationshipLinkFieldMappingResponseOutputWithContext(context.Background())
}

func (i RelationshipLinkFieldMappingResponseArgs) ToRelationshipLinkFieldMappingResponseOutputWithContext(ctx context.Context) RelationshipLinkFieldMappingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipLinkFieldMappingResponseOutput)
}





type RelationshipLinkFieldMappingResponseArrayInput interface {
	pulumi.Input

	ToRelationshipLinkFieldMappingResponseArrayOutput() RelationshipLinkFieldMappingResponseArrayOutput
	ToRelationshipLinkFieldMappingResponseArrayOutputWithContext(context.Context) RelationshipLinkFieldMappingResponseArrayOutput
}

type RelationshipLinkFieldMappingResponseArray []RelationshipLinkFieldMappingResponseInput

func (RelationshipLinkFieldMappingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RelationshipLinkFieldMappingResponse)(nil)).Elem()
}

func (i RelationshipLinkFieldMappingResponseArray) ToRelationshipLinkFieldMappingResponseArrayOutput() RelationshipLinkFieldMappingResponseArrayOutput {
	return i.ToRelationshipLinkFieldMappingResponseArrayOutputWithContext(context.Background())
}

func (i RelationshipLinkFieldMappingResponseArray) ToRelationshipLinkFieldMappingResponseArrayOutputWithContext(ctx context.Context) RelationshipLinkFieldMappingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipLinkFieldMappingResponseArrayOutput)
}

type RelationshipLinkFieldMappingResponseOutput struct{ *pulumi.OutputState }

func (RelationshipLinkFieldMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipLinkFieldMappingResponse)(nil)).Elem()
}

func (o RelationshipLinkFieldMappingResponseOutput) ToRelationshipLinkFieldMappingResponseOutput() RelationshipLinkFieldMappingResponseOutput {
	return o
}

func (o RelationshipLinkFieldMappingResponseOutput) ToRelationshipLinkFieldMappingResponseOutputWithContext(ctx context.Context) RelationshipLinkFieldMappingResponseOutput {
	return o
}

func (o RelationshipLinkFieldMappingResponseOutput) InteractionFieldName() pulumi.StringOutput {
	return o.ApplyT(func(v RelationshipLinkFieldMappingResponse) string { return v.InteractionFieldName }).(pulumi.StringOutput)
}

func (o RelationshipLinkFieldMappingResponseOutput) LinkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RelationshipLinkFieldMappingResponse) *string { return v.LinkType }).(pulumi.StringPtrOutput)
}

func (o RelationshipLinkFieldMappingResponseOutput) RelationshipFieldName() pulumi.StringOutput {
	return o.ApplyT(func(v RelationshipLinkFieldMappingResponse) string { return v.RelationshipFieldName }).(pulumi.StringOutput)
}

type RelationshipLinkFieldMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (RelationshipLinkFieldMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RelationshipLinkFieldMappingResponse)(nil)).Elem()
}

func (o RelationshipLinkFieldMappingResponseArrayOutput) ToRelationshipLinkFieldMappingResponseArrayOutput() RelationshipLinkFieldMappingResponseArrayOutput {
	return o
}

func (o RelationshipLinkFieldMappingResponseArrayOutput) ToRelationshipLinkFieldMappingResponseArrayOutputWithContext(ctx context.Context) RelationshipLinkFieldMappingResponseArrayOutput {
	return o
}

func (o RelationshipLinkFieldMappingResponseArrayOutput) Index(i pulumi.IntInput) RelationshipLinkFieldMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RelationshipLinkFieldMappingResponse {
		return vs[0].([]RelationshipLinkFieldMappingResponse)[vs[1].(int)]
	}).(RelationshipLinkFieldMappingResponseOutput)
}

type RelationshipTypeFieldMapping struct {
	ProfileFieldName          string `pulumi:"profileFieldName"`
	RelatedProfileKeyProperty string `pulumi:"relatedProfileKeyProperty"`
}





type RelationshipTypeFieldMappingInput interface {
	pulumi.Input

	ToRelationshipTypeFieldMappingOutput() RelationshipTypeFieldMappingOutput
	ToRelationshipTypeFieldMappingOutputWithContext(context.Context) RelationshipTypeFieldMappingOutput
}

type RelationshipTypeFieldMappingArgs struct {
	ProfileFieldName          pulumi.StringInput `pulumi:"profileFieldName"`
	RelatedProfileKeyProperty pulumi.StringInput `pulumi:"relatedProfileKeyProperty"`
}

func (RelationshipTypeFieldMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipTypeFieldMapping)(nil)).Elem()
}

func (i RelationshipTypeFieldMappingArgs) ToRelationshipTypeFieldMappingOutput() RelationshipTypeFieldMappingOutput {
	return i.ToRelationshipTypeFieldMappingOutputWithContext(context.Background())
}

func (i RelationshipTypeFieldMappingArgs) ToRelationshipTypeFieldMappingOutputWithContext(ctx context.Context) RelationshipTypeFieldMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipTypeFieldMappingOutput)
}





type RelationshipTypeFieldMappingArrayInput interface {
	pulumi.Input

	ToRelationshipTypeFieldMappingArrayOutput() RelationshipTypeFieldMappingArrayOutput
	ToRelationshipTypeFieldMappingArrayOutputWithContext(context.Context) RelationshipTypeFieldMappingArrayOutput
}

type RelationshipTypeFieldMappingArray []RelationshipTypeFieldMappingInput

func (RelationshipTypeFieldMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RelationshipTypeFieldMapping)(nil)).Elem()
}

func (i RelationshipTypeFieldMappingArray) ToRelationshipTypeFieldMappingArrayOutput() RelationshipTypeFieldMappingArrayOutput {
	return i.ToRelationshipTypeFieldMappingArrayOutputWithContext(context.Background())
}

func (i RelationshipTypeFieldMappingArray) ToRelationshipTypeFieldMappingArrayOutputWithContext(ctx context.Context) RelationshipTypeFieldMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipTypeFieldMappingArrayOutput)
}

type RelationshipTypeFieldMappingOutput struct{ *pulumi.OutputState }

func (RelationshipTypeFieldMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipTypeFieldMapping)(nil)).Elem()
}

func (o RelationshipTypeFieldMappingOutput) ToRelationshipTypeFieldMappingOutput() RelationshipTypeFieldMappingOutput {
	return o
}

func (o RelationshipTypeFieldMappingOutput) ToRelationshipTypeFieldMappingOutputWithContext(ctx context.Context) RelationshipTypeFieldMappingOutput {
	return o
}

func (o RelationshipTypeFieldMappingOutput) ProfileFieldName() pulumi.StringOutput {
	return o.ApplyT(func(v RelationshipTypeFieldMapping) string { return v.ProfileFieldName }).(pulumi.StringOutput)
}

func (o RelationshipTypeFieldMappingOutput) RelatedProfileKeyProperty() pulumi.StringOutput {
	return o.ApplyT(func(v RelationshipTypeFieldMapping) string { return v.RelatedProfileKeyProperty }).(pulumi.StringOutput)
}

type RelationshipTypeFieldMappingArrayOutput struct{ *pulumi.OutputState }

func (RelationshipTypeFieldMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RelationshipTypeFieldMapping)(nil)).Elem()
}

func (o RelationshipTypeFieldMappingArrayOutput) ToRelationshipTypeFieldMappingArrayOutput() RelationshipTypeFieldMappingArrayOutput {
	return o
}

func (o RelationshipTypeFieldMappingArrayOutput) ToRelationshipTypeFieldMappingArrayOutputWithContext(ctx context.Context) RelationshipTypeFieldMappingArrayOutput {
	return o
}

func (o RelationshipTypeFieldMappingArrayOutput) Index(i pulumi.IntInput) RelationshipTypeFieldMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RelationshipTypeFieldMapping {
		return vs[0].([]RelationshipTypeFieldMapping)[vs[1].(int)]
	}).(RelationshipTypeFieldMappingOutput)
}

type RelationshipTypeFieldMappingResponse struct {
	ProfileFieldName          string `pulumi:"profileFieldName"`
	RelatedProfileKeyProperty string `pulumi:"relatedProfileKeyProperty"`
}





type RelationshipTypeFieldMappingResponseInput interface {
	pulumi.Input

	ToRelationshipTypeFieldMappingResponseOutput() RelationshipTypeFieldMappingResponseOutput
	ToRelationshipTypeFieldMappingResponseOutputWithContext(context.Context) RelationshipTypeFieldMappingResponseOutput
}

type RelationshipTypeFieldMappingResponseArgs struct {
	ProfileFieldName          pulumi.StringInput `pulumi:"profileFieldName"`
	RelatedProfileKeyProperty pulumi.StringInput `pulumi:"relatedProfileKeyProperty"`
}

func (RelationshipTypeFieldMappingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipTypeFieldMappingResponse)(nil)).Elem()
}

func (i RelationshipTypeFieldMappingResponseArgs) ToRelationshipTypeFieldMappingResponseOutput() RelationshipTypeFieldMappingResponseOutput {
	return i.ToRelationshipTypeFieldMappingResponseOutputWithContext(context.Background())
}

func (i RelationshipTypeFieldMappingResponseArgs) ToRelationshipTypeFieldMappingResponseOutputWithContext(ctx context.Context) RelationshipTypeFieldMappingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipTypeFieldMappingResponseOutput)
}





type RelationshipTypeFieldMappingResponseArrayInput interface {
	pulumi.Input

	ToRelationshipTypeFieldMappingResponseArrayOutput() RelationshipTypeFieldMappingResponseArrayOutput
	ToRelationshipTypeFieldMappingResponseArrayOutputWithContext(context.Context) RelationshipTypeFieldMappingResponseArrayOutput
}

type RelationshipTypeFieldMappingResponseArray []RelationshipTypeFieldMappingResponseInput

func (RelationshipTypeFieldMappingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RelationshipTypeFieldMappingResponse)(nil)).Elem()
}

func (i RelationshipTypeFieldMappingResponseArray) ToRelationshipTypeFieldMappingResponseArrayOutput() RelationshipTypeFieldMappingResponseArrayOutput {
	return i.ToRelationshipTypeFieldMappingResponseArrayOutputWithContext(context.Background())
}

func (i RelationshipTypeFieldMappingResponseArray) ToRelationshipTypeFieldMappingResponseArrayOutputWithContext(ctx context.Context) RelationshipTypeFieldMappingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipTypeFieldMappingResponseArrayOutput)
}

type RelationshipTypeFieldMappingResponseOutput struct{ *pulumi.OutputState }

func (RelationshipTypeFieldMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipTypeFieldMappingResponse)(nil)).Elem()
}

func (o RelationshipTypeFieldMappingResponseOutput) ToRelationshipTypeFieldMappingResponseOutput() RelationshipTypeFieldMappingResponseOutput {
	return o
}

func (o RelationshipTypeFieldMappingResponseOutput) ToRelationshipTypeFieldMappingResponseOutputWithContext(ctx context.Context) RelationshipTypeFieldMappingResponseOutput {
	return o
}

func (o RelationshipTypeFieldMappingResponseOutput) ProfileFieldName() pulumi.StringOutput {
	return o.ApplyT(func(v RelationshipTypeFieldMappingResponse) string { return v.ProfileFieldName }).(pulumi.StringOutput)
}

func (o RelationshipTypeFieldMappingResponseOutput) RelatedProfileKeyProperty() pulumi.StringOutput {
	return o.ApplyT(func(v RelationshipTypeFieldMappingResponse) string { return v.RelatedProfileKeyProperty }).(pulumi.StringOutput)
}

type RelationshipTypeFieldMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (RelationshipTypeFieldMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RelationshipTypeFieldMappingResponse)(nil)).Elem()
}

func (o RelationshipTypeFieldMappingResponseArrayOutput) ToRelationshipTypeFieldMappingResponseArrayOutput() RelationshipTypeFieldMappingResponseArrayOutput {
	return o
}

func (o RelationshipTypeFieldMappingResponseArrayOutput) ToRelationshipTypeFieldMappingResponseArrayOutputWithContext(ctx context.Context) RelationshipTypeFieldMappingResponseArrayOutput {
	return o
}

func (o RelationshipTypeFieldMappingResponseArrayOutput) Index(i pulumi.IntInput) RelationshipTypeFieldMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RelationshipTypeFieldMappingResponse {
		return vs[0].([]RelationshipTypeFieldMappingResponse)[vs[1].(int)]
	}).(RelationshipTypeFieldMappingResponseOutput)
}

type RelationshipTypeMapping struct {
	FieldMappings []RelationshipTypeFieldMapping `pulumi:"fieldMappings"`
}





type RelationshipTypeMappingInput interface {
	pulumi.Input

	ToRelationshipTypeMappingOutput() RelationshipTypeMappingOutput
	ToRelationshipTypeMappingOutputWithContext(context.Context) RelationshipTypeMappingOutput
}

type RelationshipTypeMappingArgs struct {
	FieldMappings RelationshipTypeFieldMappingArrayInput `pulumi:"fieldMappings"`
}

func (RelationshipTypeMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipTypeMapping)(nil)).Elem()
}

func (i RelationshipTypeMappingArgs) ToRelationshipTypeMappingOutput() RelationshipTypeMappingOutput {
	return i.ToRelationshipTypeMappingOutputWithContext(context.Background())
}

func (i RelationshipTypeMappingArgs) ToRelationshipTypeMappingOutputWithContext(ctx context.Context) RelationshipTypeMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipTypeMappingOutput)
}





type RelationshipTypeMappingArrayInput interface {
	pulumi.Input

	ToRelationshipTypeMappingArrayOutput() RelationshipTypeMappingArrayOutput
	ToRelationshipTypeMappingArrayOutputWithContext(context.Context) RelationshipTypeMappingArrayOutput
}

type RelationshipTypeMappingArray []RelationshipTypeMappingInput

func (RelationshipTypeMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RelationshipTypeMapping)(nil)).Elem()
}

func (i RelationshipTypeMappingArray) ToRelationshipTypeMappingArrayOutput() RelationshipTypeMappingArrayOutput {
	return i.ToRelationshipTypeMappingArrayOutputWithContext(context.Background())
}

func (i RelationshipTypeMappingArray) ToRelationshipTypeMappingArrayOutputWithContext(ctx context.Context) RelationshipTypeMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipTypeMappingArrayOutput)
}

type RelationshipTypeMappingOutput struct{ *pulumi.OutputState }

func (RelationshipTypeMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipTypeMapping)(nil)).Elem()
}

func (o RelationshipTypeMappingOutput) ToRelationshipTypeMappingOutput() RelationshipTypeMappingOutput {
	return o
}

func (o RelationshipTypeMappingOutput) ToRelationshipTypeMappingOutputWithContext(ctx context.Context) RelationshipTypeMappingOutput {
	return o
}

func (o RelationshipTypeMappingOutput) FieldMappings() RelationshipTypeFieldMappingArrayOutput {
	return o.ApplyT(func(v RelationshipTypeMapping) []RelationshipTypeFieldMapping { return v.FieldMappings }).(RelationshipTypeFieldMappingArrayOutput)
}

type RelationshipTypeMappingArrayOutput struct{ *pulumi.OutputState }

func (RelationshipTypeMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RelationshipTypeMapping)(nil)).Elem()
}

func (o RelationshipTypeMappingArrayOutput) ToRelationshipTypeMappingArrayOutput() RelationshipTypeMappingArrayOutput {
	return o
}

func (o RelationshipTypeMappingArrayOutput) ToRelationshipTypeMappingArrayOutputWithContext(ctx context.Context) RelationshipTypeMappingArrayOutput {
	return o
}

func (o RelationshipTypeMappingArrayOutput) Index(i pulumi.IntInput) RelationshipTypeMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RelationshipTypeMapping {
		return vs[0].([]RelationshipTypeMapping)[vs[1].(int)]
	}).(RelationshipTypeMappingOutput)
}

type RelationshipTypeMappingResponse struct {
	FieldMappings []RelationshipTypeFieldMappingResponse `pulumi:"fieldMappings"`
}





type RelationshipTypeMappingResponseInput interface {
	pulumi.Input

	ToRelationshipTypeMappingResponseOutput() RelationshipTypeMappingResponseOutput
	ToRelationshipTypeMappingResponseOutputWithContext(context.Context) RelationshipTypeMappingResponseOutput
}

type RelationshipTypeMappingResponseArgs struct {
	FieldMappings RelationshipTypeFieldMappingResponseArrayInput `pulumi:"fieldMappings"`
}

func (RelationshipTypeMappingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipTypeMappingResponse)(nil)).Elem()
}

func (i RelationshipTypeMappingResponseArgs) ToRelationshipTypeMappingResponseOutput() RelationshipTypeMappingResponseOutput {
	return i.ToRelationshipTypeMappingResponseOutputWithContext(context.Background())
}

func (i RelationshipTypeMappingResponseArgs) ToRelationshipTypeMappingResponseOutputWithContext(ctx context.Context) RelationshipTypeMappingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipTypeMappingResponseOutput)
}





type RelationshipTypeMappingResponseArrayInput interface {
	pulumi.Input

	ToRelationshipTypeMappingResponseArrayOutput() RelationshipTypeMappingResponseArrayOutput
	ToRelationshipTypeMappingResponseArrayOutputWithContext(context.Context) RelationshipTypeMappingResponseArrayOutput
}

type RelationshipTypeMappingResponseArray []RelationshipTypeMappingResponseInput

func (RelationshipTypeMappingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RelationshipTypeMappingResponse)(nil)).Elem()
}

func (i RelationshipTypeMappingResponseArray) ToRelationshipTypeMappingResponseArrayOutput() RelationshipTypeMappingResponseArrayOutput {
	return i.ToRelationshipTypeMappingResponseArrayOutputWithContext(context.Background())
}

func (i RelationshipTypeMappingResponseArray) ToRelationshipTypeMappingResponseArrayOutputWithContext(ctx context.Context) RelationshipTypeMappingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipTypeMappingResponseArrayOutput)
}

type RelationshipTypeMappingResponseOutput struct{ *pulumi.OutputState }

func (RelationshipTypeMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipTypeMappingResponse)(nil)).Elem()
}

func (o RelationshipTypeMappingResponseOutput) ToRelationshipTypeMappingResponseOutput() RelationshipTypeMappingResponseOutput {
	return o
}

func (o RelationshipTypeMappingResponseOutput) ToRelationshipTypeMappingResponseOutputWithContext(ctx context.Context) RelationshipTypeMappingResponseOutput {
	return o
}

func (o RelationshipTypeMappingResponseOutput) FieldMappings() RelationshipTypeFieldMappingResponseArrayOutput {
	return o.ApplyT(func(v RelationshipTypeMappingResponse) []RelationshipTypeFieldMappingResponse { return v.FieldMappings }).(RelationshipTypeFieldMappingResponseArrayOutput)
}

type RelationshipTypeMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (RelationshipTypeMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RelationshipTypeMappingResponse)(nil)).Elem()
}

func (o RelationshipTypeMappingResponseArrayOutput) ToRelationshipTypeMappingResponseArrayOutput() RelationshipTypeMappingResponseArrayOutput {
	return o
}

func (o RelationshipTypeMappingResponseArrayOutput) ToRelationshipTypeMappingResponseArrayOutputWithContext(ctx context.Context) RelationshipTypeMappingResponseArrayOutput {
	return o
}

func (o RelationshipTypeMappingResponseArrayOutput) Index(i pulumi.IntInput) RelationshipTypeMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RelationshipTypeMappingResponse {
		return vs[0].([]RelationshipTypeMappingResponse)[vs[1].(int)]
	}).(RelationshipTypeMappingResponseOutput)
}

type ResourceSetDescription struct {
	Elements   []string `pulumi:"elements"`
	Exceptions []string `pulumi:"exceptions"`
}





type ResourceSetDescriptionInput interface {
	pulumi.Input

	ToResourceSetDescriptionOutput() ResourceSetDescriptionOutput
	ToResourceSetDescriptionOutputWithContext(context.Context) ResourceSetDescriptionOutput
}

type ResourceSetDescriptionArgs struct {
	Elements   pulumi.StringArrayInput `pulumi:"elements"`
	Exceptions pulumi.StringArrayInput `pulumi:"exceptions"`
}

func (ResourceSetDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSetDescription)(nil)).Elem()
}

func (i ResourceSetDescriptionArgs) ToResourceSetDescriptionOutput() ResourceSetDescriptionOutput {
	return i.ToResourceSetDescriptionOutputWithContext(context.Background())
}

func (i ResourceSetDescriptionArgs) ToResourceSetDescriptionOutputWithContext(ctx context.Context) ResourceSetDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSetDescriptionOutput)
}

func (i ResourceSetDescriptionArgs) ToResourceSetDescriptionPtrOutput() ResourceSetDescriptionPtrOutput {
	return i.ToResourceSetDescriptionPtrOutputWithContext(context.Background())
}

func (i ResourceSetDescriptionArgs) ToResourceSetDescriptionPtrOutputWithContext(ctx context.Context) ResourceSetDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSetDescriptionOutput).ToResourceSetDescriptionPtrOutputWithContext(ctx)
}









type ResourceSetDescriptionPtrInput interface {
	pulumi.Input

	ToResourceSetDescriptionPtrOutput() ResourceSetDescriptionPtrOutput
	ToResourceSetDescriptionPtrOutputWithContext(context.Context) ResourceSetDescriptionPtrOutput
}

type resourceSetDescriptionPtrType ResourceSetDescriptionArgs

func ResourceSetDescriptionPtr(v *ResourceSetDescriptionArgs) ResourceSetDescriptionPtrInput {
	return (*resourceSetDescriptionPtrType)(v)
}

func (*resourceSetDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSetDescription)(nil)).Elem()
}

func (i *resourceSetDescriptionPtrType) ToResourceSetDescriptionPtrOutput() ResourceSetDescriptionPtrOutput {
	return i.ToResourceSetDescriptionPtrOutputWithContext(context.Background())
}

func (i *resourceSetDescriptionPtrType) ToResourceSetDescriptionPtrOutputWithContext(ctx context.Context) ResourceSetDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSetDescriptionPtrOutput)
}

type ResourceSetDescriptionOutput struct{ *pulumi.OutputState }

func (ResourceSetDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSetDescription)(nil)).Elem()
}

func (o ResourceSetDescriptionOutput) ToResourceSetDescriptionOutput() ResourceSetDescriptionOutput {
	return o
}

func (o ResourceSetDescriptionOutput) ToResourceSetDescriptionOutputWithContext(ctx context.Context) ResourceSetDescriptionOutput {
	return o
}

func (o ResourceSetDescriptionOutput) ToResourceSetDescriptionPtrOutput() ResourceSetDescriptionPtrOutput {
	return o.ToResourceSetDescriptionPtrOutputWithContext(context.Background())
}

func (o ResourceSetDescriptionOutput) ToResourceSetDescriptionPtrOutputWithContext(ctx context.Context) ResourceSetDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSetDescription) *ResourceSetDescription {
		return &v
	}).(ResourceSetDescriptionPtrOutput)
}

func (o ResourceSetDescriptionOutput) Elements() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceSetDescription) []string { return v.Elements }).(pulumi.StringArrayOutput)
}

func (o ResourceSetDescriptionOutput) Exceptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceSetDescription) []string { return v.Exceptions }).(pulumi.StringArrayOutput)
}

type ResourceSetDescriptionPtrOutput struct{ *pulumi.OutputState }

func (ResourceSetDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSetDescription)(nil)).Elem()
}

func (o ResourceSetDescriptionPtrOutput) ToResourceSetDescriptionPtrOutput() ResourceSetDescriptionPtrOutput {
	return o
}

func (o ResourceSetDescriptionPtrOutput) ToResourceSetDescriptionPtrOutputWithContext(ctx context.Context) ResourceSetDescriptionPtrOutput {
	return o
}

func (o ResourceSetDescriptionPtrOutput) Elem() ResourceSetDescriptionOutput {
	return o.ApplyT(func(v *ResourceSetDescription) ResourceSetDescription {
		if v != nil {
			return *v
		}
		var ret ResourceSetDescription
		return ret
	}).(ResourceSetDescriptionOutput)
}

func (o ResourceSetDescriptionPtrOutput) Elements() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceSetDescription) []string {
		if v == nil {
			return nil
		}
		return v.Elements
	}).(pulumi.StringArrayOutput)
}

func (o ResourceSetDescriptionPtrOutput) Exceptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceSetDescription) []string {
		if v == nil {
			return nil
		}
		return v.Exceptions
	}).(pulumi.StringArrayOutput)
}

type ResourceSetDescriptionResponse struct {
	Elements   []string `pulumi:"elements"`
	Exceptions []string `pulumi:"exceptions"`
}





type ResourceSetDescriptionResponseInput interface {
	pulumi.Input

	ToResourceSetDescriptionResponseOutput() ResourceSetDescriptionResponseOutput
	ToResourceSetDescriptionResponseOutputWithContext(context.Context) ResourceSetDescriptionResponseOutput
}

type ResourceSetDescriptionResponseArgs struct {
	Elements   pulumi.StringArrayInput `pulumi:"elements"`
	Exceptions pulumi.StringArrayInput `pulumi:"exceptions"`
}

func (ResourceSetDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSetDescriptionResponse)(nil)).Elem()
}

func (i ResourceSetDescriptionResponseArgs) ToResourceSetDescriptionResponseOutput() ResourceSetDescriptionResponseOutput {
	return i.ToResourceSetDescriptionResponseOutputWithContext(context.Background())
}

func (i ResourceSetDescriptionResponseArgs) ToResourceSetDescriptionResponseOutputWithContext(ctx context.Context) ResourceSetDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSetDescriptionResponseOutput)
}

func (i ResourceSetDescriptionResponseArgs) ToResourceSetDescriptionResponsePtrOutput() ResourceSetDescriptionResponsePtrOutput {
	return i.ToResourceSetDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i ResourceSetDescriptionResponseArgs) ToResourceSetDescriptionResponsePtrOutputWithContext(ctx context.Context) ResourceSetDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSetDescriptionResponseOutput).ToResourceSetDescriptionResponsePtrOutputWithContext(ctx)
}









type ResourceSetDescriptionResponsePtrInput interface {
	pulumi.Input

	ToResourceSetDescriptionResponsePtrOutput() ResourceSetDescriptionResponsePtrOutput
	ToResourceSetDescriptionResponsePtrOutputWithContext(context.Context) ResourceSetDescriptionResponsePtrOutput
}

type resourceSetDescriptionResponsePtrType ResourceSetDescriptionResponseArgs

func ResourceSetDescriptionResponsePtr(v *ResourceSetDescriptionResponseArgs) ResourceSetDescriptionResponsePtrInput {
	return (*resourceSetDescriptionResponsePtrType)(v)
}

func (*resourceSetDescriptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSetDescriptionResponse)(nil)).Elem()
}

func (i *resourceSetDescriptionResponsePtrType) ToResourceSetDescriptionResponsePtrOutput() ResourceSetDescriptionResponsePtrOutput {
	return i.ToResourceSetDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i *resourceSetDescriptionResponsePtrType) ToResourceSetDescriptionResponsePtrOutputWithContext(ctx context.Context) ResourceSetDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSetDescriptionResponsePtrOutput)
}

type ResourceSetDescriptionResponseOutput struct{ *pulumi.OutputState }

func (ResourceSetDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSetDescriptionResponse)(nil)).Elem()
}

func (o ResourceSetDescriptionResponseOutput) ToResourceSetDescriptionResponseOutput() ResourceSetDescriptionResponseOutput {
	return o
}

func (o ResourceSetDescriptionResponseOutput) ToResourceSetDescriptionResponseOutputWithContext(ctx context.Context) ResourceSetDescriptionResponseOutput {
	return o
}

func (o ResourceSetDescriptionResponseOutput) ToResourceSetDescriptionResponsePtrOutput() ResourceSetDescriptionResponsePtrOutput {
	return o.ToResourceSetDescriptionResponsePtrOutputWithContext(context.Background())
}

func (o ResourceSetDescriptionResponseOutput) ToResourceSetDescriptionResponsePtrOutputWithContext(ctx context.Context) ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSetDescriptionResponse) *ResourceSetDescriptionResponse {
		return &v
	}).(ResourceSetDescriptionResponsePtrOutput)
}

func (o ResourceSetDescriptionResponseOutput) Elements() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceSetDescriptionResponse) []string { return v.Elements }).(pulumi.StringArrayOutput)
}

func (o ResourceSetDescriptionResponseOutput) Exceptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceSetDescriptionResponse) []string { return v.Exceptions }).(pulumi.StringArrayOutput)
}

type ResourceSetDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSetDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSetDescriptionResponse)(nil)).Elem()
}

func (o ResourceSetDescriptionResponsePtrOutput) ToResourceSetDescriptionResponsePtrOutput() ResourceSetDescriptionResponsePtrOutput {
	return o
}

func (o ResourceSetDescriptionResponsePtrOutput) ToResourceSetDescriptionResponsePtrOutputWithContext(ctx context.Context) ResourceSetDescriptionResponsePtrOutput {
	return o
}

func (o ResourceSetDescriptionResponsePtrOutput) Elem() ResourceSetDescriptionResponseOutput {
	return o.ApplyT(func(v *ResourceSetDescriptionResponse) ResourceSetDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSetDescriptionResponse
		return ret
	}).(ResourceSetDescriptionResponseOutput)
}

func (o ResourceSetDescriptionResponsePtrOutput) Elements() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceSetDescriptionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Elements
	}).(pulumi.StringArrayOutput)
}

func (o ResourceSetDescriptionResponsePtrOutput) Exceptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceSetDescriptionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Exceptions
	}).(pulumi.StringArrayOutput)
}

type StrongId struct {
	Description      map[string]string `pulumi:"description"`
	DisplayName      map[string]string `pulumi:"displayName"`
	KeyPropertyNames []string          `pulumi:"keyPropertyNames"`
	StrongIdName     string            `pulumi:"strongIdName"`
}





type StrongIdInput interface {
	pulumi.Input

	ToStrongIdOutput() StrongIdOutput
	ToStrongIdOutputWithContext(context.Context) StrongIdOutput
}

type StrongIdArgs struct {
	Description      pulumi.StringMapInput   `pulumi:"description"`
	DisplayName      pulumi.StringMapInput   `pulumi:"displayName"`
	KeyPropertyNames pulumi.StringArrayInput `pulumi:"keyPropertyNames"`
	StrongIdName     pulumi.StringInput      `pulumi:"strongIdName"`
}

func (StrongIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StrongId)(nil)).Elem()
}

func (i StrongIdArgs) ToStrongIdOutput() StrongIdOutput {
	return i.ToStrongIdOutputWithContext(context.Background())
}

func (i StrongIdArgs) ToStrongIdOutputWithContext(ctx context.Context) StrongIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrongIdOutput)
}





type StrongIdArrayInput interface {
	pulumi.Input

	ToStrongIdArrayOutput() StrongIdArrayOutput
	ToStrongIdArrayOutputWithContext(context.Context) StrongIdArrayOutput
}

type StrongIdArray []StrongIdInput

func (StrongIdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StrongId)(nil)).Elem()
}

func (i StrongIdArray) ToStrongIdArrayOutput() StrongIdArrayOutput {
	return i.ToStrongIdArrayOutputWithContext(context.Background())
}

func (i StrongIdArray) ToStrongIdArrayOutputWithContext(ctx context.Context) StrongIdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrongIdArrayOutput)
}

type StrongIdOutput struct{ *pulumi.OutputState }

func (StrongIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StrongId)(nil)).Elem()
}

func (o StrongIdOutput) ToStrongIdOutput() StrongIdOutput {
	return o
}

func (o StrongIdOutput) ToStrongIdOutputWithContext(ctx context.Context) StrongIdOutput {
	return o
}

func (o StrongIdOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v StrongId) map[string]string { return v.Description }).(pulumi.StringMapOutput)
}

func (o StrongIdOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v StrongId) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o StrongIdOutput) KeyPropertyNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StrongId) []string { return v.KeyPropertyNames }).(pulumi.StringArrayOutput)
}

func (o StrongIdOutput) StrongIdName() pulumi.StringOutput {
	return o.ApplyT(func(v StrongId) string { return v.StrongIdName }).(pulumi.StringOutput)
}

type StrongIdArrayOutput struct{ *pulumi.OutputState }

func (StrongIdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StrongId)(nil)).Elem()
}

func (o StrongIdArrayOutput) ToStrongIdArrayOutput() StrongIdArrayOutput {
	return o
}

func (o StrongIdArrayOutput) ToStrongIdArrayOutputWithContext(ctx context.Context) StrongIdArrayOutput {
	return o
}

func (o StrongIdArrayOutput) Index(i pulumi.IntInput) StrongIdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StrongId {
		return vs[0].([]StrongId)[vs[1].(int)]
	}).(StrongIdOutput)
}

type StrongIdResponse struct {
	Description      map[string]string `pulumi:"description"`
	DisplayName      map[string]string `pulumi:"displayName"`
	KeyPropertyNames []string          `pulumi:"keyPropertyNames"`
	StrongIdName     string            `pulumi:"strongIdName"`
}





type StrongIdResponseInput interface {
	pulumi.Input

	ToStrongIdResponseOutput() StrongIdResponseOutput
	ToStrongIdResponseOutputWithContext(context.Context) StrongIdResponseOutput
}

type StrongIdResponseArgs struct {
	Description      pulumi.StringMapInput   `pulumi:"description"`
	DisplayName      pulumi.StringMapInput   `pulumi:"displayName"`
	KeyPropertyNames pulumi.StringArrayInput `pulumi:"keyPropertyNames"`
	StrongIdName     pulumi.StringInput      `pulumi:"strongIdName"`
}

func (StrongIdResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StrongIdResponse)(nil)).Elem()
}

func (i StrongIdResponseArgs) ToStrongIdResponseOutput() StrongIdResponseOutput {
	return i.ToStrongIdResponseOutputWithContext(context.Background())
}

func (i StrongIdResponseArgs) ToStrongIdResponseOutputWithContext(ctx context.Context) StrongIdResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrongIdResponseOutput)
}





type StrongIdResponseArrayInput interface {
	pulumi.Input

	ToStrongIdResponseArrayOutput() StrongIdResponseArrayOutput
	ToStrongIdResponseArrayOutputWithContext(context.Context) StrongIdResponseArrayOutput
}

type StrongIdResponseArray []StrongIdResponseInput

func (StrongIdResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StrongIdResponse)(nil)).Elem()
}

func (i StrongIdResponseArray) ToStrongIdResponseArrayOutput() StrongIdResponseArrayOutput {
	return i.ToStrongIdResponseArrayOutputWithContext(context.Background())
}

func (i StrongIdResponseArray) ToStrongIdResponseArrayOutputWithContext(ctx context.Context) StrongIdResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrongIdResponseArrayOutput)
}

type StrongIdResponseOutput struct{ *pulumi.OutputState }

func (StrongIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StrongIdResponse)(nil)).Elem()
}

func (o StrongIdResponseOutput) ToStrongIdResponseOutput() StrongIdResponseOutput {
	return o
}

func (o StrongIdResponseOutput) ToStrongIdResponseOutputWithContext(ctx context.Context) StrongIdResponseOutput {
	return o
}

func (o StrongIdResponseOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v StrongIdResponse) map[string]string { return v.Description }).(pulumi.StringMapOutput)
}

func (o StrongIdResponseOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v StrongIdResponse) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o StrongIdResponseOutput) KeyPropertyNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StrongIdResponse) []string { return v.KeyPropertyNames }).(pulumi.StringArrayOutput)
}

func (o StrongIdResponseOutput) StrongIdName() pulumi.StringOutput {
	return o.ApplyT(func(v StrongIdResponse) string { return v.StrongIdName }).(pulumi.StringOutput)
}

type StrongIdResponseArrayOutput struct{ *pulumi.OutputState }

func (StrongIdResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StrongIdResponse)(nil)).Elem()
}

func (o StrongIdResponseArrayOutput) ToStrongIdResponseArrayOutput() StrongIdResponseArrayOutput {
	return o
}

func (o StrongIdResponseArrayOutput) ToStrongIdResponseArrayOutputWithContext(ctx context.Context) StrongIdResponseArrayOutput {
	return o
}

func (o StrongIdResponseArrayOutput) Index(i pulumi.IntInput) StrongIdResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StrongIdResponse {
		return vs[0].([]StrongIdResponse)[vs[1].(int)]
	}).(StrongIdResponseOutput)
}

type TypePropertiesMapping struct {
	InteractionTypePropertyName string     `pulumi:"interactionTypePropertyName"`
	IsProfileTypeId             *bool      `pulumi:"isProfileTypeId"`
	LinkType                    *LinkTypes `pulumi:"linkType"`
	ProfileTypePropertyName     string     `pulumi:"profileTypePropertyName"`
}





type TypePropertiesMappingInput interface {
	pulumi.Input

	ToTypePropertiesMappingOutput() TypePropertiesMappingOutput
	ToTypePropertiesMappingOutputWithContext(context.Context) TypePropertiesMappingOutput
}

type TypePropertiesMappingArgs struct {
	InteractionTypePropertyName pulumi.StringInput  `pulumi:"interactionTypePropertyName"`
	IsProfileTypeId             pulumi.BoolPtrInput `pulumi:"isProfileTypeId"`
	LinkType                    LinkTypesPtrInput   `pulumi:"linkType"`
	ProfileTypePropertyName     pulumi.StringInput  `pulumi:"profileTypePropertyName"`
}

func (TypePropertiesMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TypePropertiesMapping)(nil)).Elem()
}

func (i TypePropertiesMappingArgs) ToTypePropertiesMappingOutput() TypePropertiesMappingOutput {
	return i.ToTypePropertiesMappingOutputWithContext(context.Background())
}

func (i TypePropertiesMappingArgs) ToTypePropertiesMappingOutputWithContext(ctx context.Context) TypePropertiesMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TypePropertiesMappingOutput)
}





type TypePropertiesMappingArrayInput interface {
	pulumi.Input

	ToTypePropertiesMappingArrayOutput() TypePropertiesMappingArrayOutput
	ToTypePropertiesMappingArrayOutputWithContext(context.Context) TypePropertiesMappingArrayOutput
}

type TypePropertiesMappingArray []TypePropertiesMappingInput

func (TypePropertiesMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TypePropertiesMapping)(nil)).Elem()
}

func (i TypePropertiesMappingArray) ToTypePropertiesMappingArrayOutput() TypePropertiesMappingArrayOutput {
	return i.ToTypePropertiesMappingArrayOutputWithContext(context.Background())
}

func (i TypePropertiesMappingArray) ToTypePropertiesMappingArrayOutputWithContext(ctx context.Context) TypePropertiesMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TypePropertiesMappingArrayOutput)
}

type TypePropertiesMappingOutput struct{ *pulumi.OutputState }

func (TypePropertiesMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TypePropertiesMapping)(nil)).Elem()
}

func (o TypePropertiesMappingOutput) ToTypePropertiesMappingOutput() TypePropertiesMappingOutput {
	return o
}

func (o TypePropertiesMappingOutput) ToTypePropertiesMappingOutputWithContext(ctx context.Context) TypePropertiesMappingOutput {
	return o
}

func (o TypePropertiesMappingOutput) InteractionTypePropertyName() pulumi.StringOutput {
	return o.ApplyT(func(v TypePropertiesMapping) string { return v.InteractionTypePropertyName }).(pulumi.StringOutput)
}

func (o TypePropertiesMappingOutput) IsProfileTypeId() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TypePropertiesMapping) *bool { return v.IsProfileTypeId }).(pulumi.BoolPtrOutput)
}

func (o TypePropertiesMappingOutput) LinkType() LinkTypesPtrOutput {
	return o.ApplyT(func(v TypePropertiesMapping) *LinkTypes { return v.LinkType }).(LinkTypesPtrOutput)
}

func (o TypePropertiesMappingOutput) ProfileTypePropertyName() pulumi.StringOutput {
	return o.ApplyT(func(v TypePropertiesMapping) string { return v.ProfileTypePropertyName }).(pulumi.StringOutput)
}

type TypePropertiesMappingArrayOutput struct{ *pulumi.OutputState }

func (TypePropertiesMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TypePropertiesMapping)(nil)).Elem()
}

func (o TypePropertiesMappingArrayOutput) ToTypePropertiesMappingArrayOutput() TypePropertiesMappingArrayOutput {
	return o
}

func (o TypePropertiesMappingArrayOutput) ToTypePropertiesMappingArrayOutputWithContext(ctx context.Context) TypePropertiesMappingArrayOutput {
	return o
}

func (o TypePropertiesMappingArrayOutput) Index(i pulumi.IntInput) TypePropertiesMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TypePropertiesMapping {
		return vs[0].([]TypePropertiesMapping)[vs[1].(int)]
	}).(TypePropertiesMappingOutput)
}

type TypePropertiesMappingResponse struct {
	InteractionTypePropertyName string  `pulumi:"interactionTypePropertyName"`
	IsProfileTypeId             *bool   `pulumi:"isProfileTypeId"`
	LinkType                    *string `pulumi:"linkType"`
	ProfileTypePropertyName     string  `pulumi:"profileTypePropertyName"`
}





type TypePropertiesMappingResponseInput interface {
	pulumi.Input

	ToTypePropertiesMappingResponseOutput() TypePropertiesMappingResponseOutput
	ToTypePropertiesMappingResponseOutputWithContext(context.Context) TypePropertiesMappingResponseOutput
}

type TypePropertiesMappingResponseArgs struct {
	InteractionTypePropertyName pulumi.StringInput    `pulumi:"interactionTypePropertyName"`
	IsProfileTypeId             pulumi.BoolPtrInput   `pulumi:"isProfileTypeId"`
	LinkType                    pulumi.StringPtrInput `pulumi:"linkType"`
	ProfileTypePropertyName     pulumi.StringInput    `pulumi:"profileTypePropertyName"`
}

func (TypePropertiesMappingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TypePropertiesMappingResponse)(nil)).Elem()
}

func (i TypePropertiesMappingResponseArgs) ToTypePropertiesMappingResponseOutput() TypePropertiesMappingResponseOutput {
	return i.ToTypePropertiesMappingResponseOutputWithContext(context.Background())
}

func (i TypePropertiesMappingResponseArgs) ToTypePropertiesMappingResponseOutputWithContext(ctx context.Context) TypePropertiesMappingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TypePropertiesMappingResponseOutput)
}





type TypePropertiesMappingResponseArrayInput interface {
	pulumi.Input

	ToTypePropertiesMappingResponseArrayOutput() TypePropertiesMappingResponseArrayOutput
	ToTypePropertiesMappingResponseArrayOutputWithContext(context.Context) TypePropertiesMappingResponseArrayOutput
}

type TypePropertiesMappingResponseArray []TypePropertiesMappingResponseInput

func (TypePropertiesMappingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TypePropertiesMappingResponse)(nil)).Elem()
}

func (i TypePropertiesMappingResponseArray) ToTypePropertiesMappingResponseArrayOutput() TypePropertiesMappingResponseArrayOutput {
	return i.ToTypePropertiesMappingResponseArrayOutputWithContext(context.Background())
}

func (i TypePropertiesMappingResponseArray) ToTypePropertiesMappingResponseArrayOutputWithContext(ctx context.Context) TypePropertiesMappingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TypePropertiesMappingResponseArrayOutput)
}

type TypePropertiesMappingResponseOutput struct{ *pulumi.OutputState }

func (TypePropertiesMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TypePropertiesMappingResponse)(nil)).Elem()
}

func (o TypePropertiesMappingResponseOutput) ToTypePropertiesMappingResponseOutput() TypePropertiesMappingResponseOutput {
	return o
}

func (o TypePropertiesMappingResponseOutput) ToTypePropertiesMappingResponseOutputWithContext(ctx context.Context) TypePropertiesMappingResponseOutput {
	return o
}

func (o TypePropertiesMappingResponseOutput) InteractionTypePropertyName() pulumi.StringOutput {
	return o.ApplyT(func(v TypePropertiesMappingResponse) string { return v.InteractionTypePropertyName }).(pulumi.StringOutput)
}

func (o TypePropertiesMappingResponseOutput) IsProfileTypeId() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TypePropertiesMappingResponse) *bool { return v.IsProfileTypeId }).(pulumi.BoolPtrOutput)
}

func (o TypePropertiesMappingResponseOutput) LinkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TypePropertiesMappingResponse) *string { return v.LinkType }).(pulumi.StringPtrOutput)
}

func (o TypePropertiesMappingResponseOutput) ProfileTypePropertyName() pulumi.StringOutput {
	return o.ApplyT(func(v TypePropertiesMappingResponse) string { return v.ProfileTypePropertyName }).(pulumi.StringOutput)
}

type TypePropertiesMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (TypePropertiesMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TypePropertiesMappingResponse)(nil)).Elem()
}

func (o TypePropertiesMappingResponseArrayOutput) ToTypePropertiesMappingResponseArrayOutput() TypePropertiesMappingResponseArrayOutput {
	return o
}

func (o TypePropertiesMappingResponseArrayOutput) ToTypePropertiesMappingResponseArrayOutputWithContext(ctx context.Context) TypePropertiesMappingResponseArrayOutput {
	return o
}

func (o TypePropertiesMappingResponseArrayOutput) Index(i pulumi.IntInput) TypePropertiesMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TypePropertiesMappingResponse {
		return vs[0].([]TypePropertiesMappingResponse)[vs[1].(int)]
	}).(TypePropertiesMappingResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AssignmentPrincipalOutput{})
	pulumi.RegisterOutputType(AssignmentPrincipalArrayOutput{})
	pulumi.RegisterOutputType(AssignmentPrincipalResponseOutput{})
	pulumi.RegisterOutputType(AssignmentPrincipalResponseArrayOutput{})
	pulumi.RegisterOutputType(ConnectorMappingAvailabilityOutput{})
	pulumi.RegisterOutputType(ConnectorMappingAvailabilityPtrOutput{})
	pulumi.RegisterOutputType(ConnectorMappingAvailabilityResponseOutput{})
	pulumi.RegisterOutputType(ConnectorMappingAvailabilityResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectorMappingCompleteOperationOutput{})
	pulumi.RegisterOutputType(ConnectorMappingCompleteOperationPtrOutput{})
	pulumi.RegisterOutputType(ConnectorMappingCompleteOperationResponseOutput{})
	pulumi.RegisterOutputType(ConnectorMappingCompleteOperationResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectorMappingErrorManagementOutput{})
	pulumi.RegisterOutputType(ConnectorMappingErrorManagementPtrOutput{})
	pulumi.RegisterOutputType(ConnectorMappingErrorManagementResponseOutput{})
	pulumi.RegisterOutputType(ConnectorMappingErrorManagementResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectorMappingFormatOutput{})
	pulumi.RegisterOutputType(ConnectorMappingFormatPtrOutput{})
	pulumi.RegisterOutputType(ConnectorMappingFormatResponseOutput{})
	pulumi.RegisterOutputType(ConnectorMappingFormatResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectorMappingPropertiesOutput{})
	pulumi.RegisterOutputType(ConnectorMappingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConnectorMappingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConnectorMappingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectorMappingStructureOutput{})
	pulumi.RegisterOutputType(ConnectorMappingStructureArrayOutput{})
	pulumi.RegisterOutputType(ConnectorMappingStructureResponseOutput{})
	pulumi.RegisterOutputType(ConnectorMappingStructureResponseArrayOutput{})
	pulumi.RegisterOutputType(DataSourcePrecedenceResponseOutput{})
	pulumi.RegisterOutputType(DataSourcePrecedenceResponseArrayOutput{})
	pulumi.RegisterOutputType(HubBillingInfoFormatOutput{})
	pulumi.RegisterOutputType(HubBillingInfoFormatPtrOutput{})
	pulumi.RegisterOutputType(HubBillingInfoFormatResponseOutput{})
	pulumi.RegisterOutputType(HubBillingInfoFormatResponsePtrOutput{})
	pulumi.RegisterOutputType(KpiAliasOutput{})
	pulumi.RegisterOutputType(KpiAliasArrayOutput{})
	pulumi.RegisterOutputType(KpiAliasResponseOutput{})
	pulumi.RegisterOutputType(KpiAliasResponseArrayOutput{})
	pulumi.RegisterOutputType(KpiExtractOutput{})
	pulumi.RegisterOutputType(KpiExtractArrayOutput{})
	pulumi.RegisterOutputType(KpiExtractResponseOutput{})
	pulumi.RegisterOutputType(KpiExtractResponseArrayOutput{})
	pulumi.RegisterOutputType(KpiGroupByMetadataResponseOutput{})
	pulumi.RegisterOutputType(KpiGroupByMetadataResponseArrayOutput{})
	pulumi.RegisterOutputType(KpiParticipantProfilesMetadataResponseOutput{})
	pulumi.RegisterOutputType(KpiParticipantProfilesMetadataResponseArrayOutput{})
	pulumi.RegisterOutputType(KpiThresholdsOutput{})
	pulumi.RegisterOutputType(KpiThresholdsPtrOutput{})
	pulumi.RegisterOutputType(KpiThresholdsResponseOutput{})
	pulumi.RegisterOutputType(KpiThresholdsResponsePtrOutput{})
	pulumi.RegisterOutputType(ParticipantPropertyReferenceOutput{})
	pulumi.RegisterOutputType(ParticipantPropertyReferenceArrayOutput{})
	pulumi.RegisterOutputType(ParticipantPropertyReferenceResponseOutput{})
	pulumi.RegisterOutputType(ParticipantPropertyReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(ProfileEnumValidValuesFormatOutput{})
	pulumi.RegisterOutputType(ProfileEnumValidValuesFormatArrayOutput{})
	pulumi.RegisterOutputType(ProfileEnumValidValuesFormatResponseOutput{})
	pulumi.RegisterOutputType(ProfileEnumValidValuesFormatResponseArrayOutput{})
	pulumi.RegisterOutputType(PropertyDefinitionOutput{})
	pulumi.RegisterOutputType(PropertyDefinitionArrayOutput{})
	pulumi.RegisterOutputType(PropertyDefinitionResponseOutput{})
	pulumi.RegisterOutputType(PropertyDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(RelationshipLinkFieldMappingOutput{})
	pulumi.RegisterOutputType(RelationshipLinkFieldMappingArrayOutput{})
	pulumi.RegisterOutputType(RelationshipLinkFieldMappingResponseOutput{})
	pulumi.RegisterOutputType(RelationshipLinkFieldMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(RelationshipTypeFieldMappingOutput{})
	pulumi.RegisterOutputType(RelationshipTypeFieldMappingArrayOutput{})
	pulumi.RegisterOutputType(RelationshipTypeFieldMappingResponseOutput{})
	pulumi.RegisterOutputType(RelationshipTypeFieldMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(RelationshipTypeMappingOutput{})
	pulumi.RegisterOutputType(RelationshipTypeMappingArrayOutput{})
	pulumi.RegisterOutputType(RelationshipTypeMappingResponseOutput{})
	pulumi.RegisterOutputType(RelationshipTypeMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceSetDescriptionOutput{})
	pulumi.RegisterOutputType(ResourceSetDescriptionPtrOutput{})
	pulumi.RegisterOutputType(ResourceSetDescriptionResponseOutput{})
	pulumi.RegisterOutputType(ResourceSetDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(StrongIdOutput{})
	pulumi.RegisterOutputType(StrongIdArrayOutput{})
	pulumi.RegisterOutputType(StrongIdResponseOutput{})
	pulumi.RegisterOutputType(StrongIdResponseArrayOutput{})
	pulumi.RegisterOutputType(TypePropertiesMappingOutput{})
	pulumi.RegisterOutputType(TypePropertiesMappingArrayOutput{})
	pulumi.RegisterOutputType(TypePropertiesMappingResponseOutput{})
	pulumi.RegisterOutputType(TypePropertiesMappingResponseArrayOutput{})
}
