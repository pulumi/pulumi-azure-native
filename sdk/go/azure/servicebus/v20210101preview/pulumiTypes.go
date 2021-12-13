


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Action struct {
	CompatibilityLevel    *int    `pulumi:"compatibilityLevel"`
	RequiresPreprocessing *bool   `pulumi:"requiresPreprocessing"`
	SqlExpression         *string `pulumi:"sqlExpression"`
}


func (val *Action) Defaults() *Action {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RequiresPreprocessing) {
		requiresPreprocessing_ := true
		tmp.RequiresPreprocessing = &requiresPreprocessing_
	}
	return &tmp
}





type ActionInput interface {
	pulumi.Input

	ToActionOutput() ActionOutput
	ToActionOutputWithContext(context.Context) ActionOutput
}

type ActionArgs struct {
	CompatibilityLevel    pulumi.IntPtrInput    `pulumi:"compatibilityLevel"`
	RequiresPreprocessing pulumi.BoolPtrInput   `pulumi:"requiresPreprocessing"`
	SqlExpression         pulumi.StringPtrInput `pulumi:"sqlExpression"`
}

func (ActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (i ActionArgs) ToActionOutput() ActionOutput {
	return i.ToActionOutputWithContext(context.Background())
}

func (i ActionArgs) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionOutput)
}

func (i ActionArgs) ToActionPtrOutput() ActionPtrOutput {
	return i.ToActionPtrOutputWithContext(context.Background())
}

func (i ActionArgs) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionOutput).ToActionPtrOutputWithContext(ctx)
}









type ActionPtrInput interface {
	pulumi.Input

	ToActionPtrOutput() ActionPtrOutput
	ToActionPtrOutputWithContext(context.Context) ActionPtrOutput
}

type actionPtrType ActionArgs

func ActionPtr(v *ActionArgs) ActionPtrInput {
	return (*actionPtrType)(v)
}

func (*actionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (i *actionPtrType) ToActionPtrOutput() ActionPtrOutput {
	return i.ToActionPtrOutputWithContext(context.Background())
}

func (i *actionPtrType) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionPtrOutput)
}

type ActionOutput struct{ *pulumi.OutputState }

func (ActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (o ActionOutput) ToActionOutput() ActionOutput {
	return o
}

func (o ActionOutput) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return o
}

func (o ActionOutput) ToActionPtrOutput() ActionPtrOutput {
	return o.ToActionPtrOutputWithContext(context.Background())
}

func (o ActionOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Action) *Action {
		return &v
	}).(ActionPtrOutput)
}

func (o ActionOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Action) *int { return v.CompatibilityLevel }).(pulumi.IntPtrOutput)
}

func (o ActionOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Action) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

func (o ActionOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Action) *string { return v.SqlExpression }).(pulumi.StringPtrOutput)
}

type ActionPtrOutput struct{ *pulumi.OutputState }

func (ActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (o ActionPtrOutput) ToActionPtrOutput() ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) Elem() ActionOutput {
	return o.ApplyT(func(v *Action) Action {
		if v != nil {
			return *v
		}
		var ret Action
		return ret
	}).(ActionOutput)
}

func (o ActionPtrOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Action) *int {
		if v == nil {
			return nil
		}
		return v.CompatibilityLevel
	}).(pulumi.IntPtrOutput)
}

func (o ActionPtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Action) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

func (o ActionPtrOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Action) *string {
		if v == nil {
			return nil
		}
		return v.SqlExpression
	}).(pulumi.StringPtrOutput)
}

type ActionResponse struct {
	CompatibilityLevel    *int    `pulumi:"compatibilityLevel"`
	RequiresPreprocessing *bool   `pulumi:"requiresPreprocessing"`
	SqlExpression         *string `pulumi:"sqlExpression"`
}


func (val *ActionResponse) Defaults() *ActionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RequiresPreprocessing) {
		requiresPreprocessing_ := true
		tmp.RequiresPreprocessing = &requiresPreprocessing_
	}
	return &tmp
}

type ActionResponseOutput struct{ *pulumi.OutputState }

func (ActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionResponse)(nil)).Elem()
}

func (o ActionResponseOutput) ToActionResponseOutput() ActionResponseOutput {
	return o
}

func (o ActionResponseOutput) ToActionResponseOutputWithContext(ctx context.Context) ActionResponseOutput {
	return o
}

func (o ActionResponseOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ActionResponse) *int { return v.CompatibilityLevel }).(pulumi.IntPtrOutput)
}

func (o ActionResponseOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActionResponse) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

func (o ActionResponseOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionResponse) *string { return v.SqlExpression }).(pulumi.StringPtrOutput)
}

type ActionResponsePtrOutput struct{ *pulumi.OutputState }

func (ActionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionResponse)(nil)).Elem()
}

func (o ActionResponsePtrOutput) ToActionResponsePtrOutput() ActionResponsePtrOutput {
	return o
}

func (o ActionResponsePtrOutput) ToActionResponsePtrOutputWithContext(ctx context.Context) ActionResponsePtrOutput {
	return o
}

func (o ActionResponsePtrOutput) Elem() ActionResponseOutput {
	return o.ApplyT(func(v *ActionResponse) ActionResponse {
		if v != nil {
			return *v
		}
		var ret ActionResponse
		return ret
	}).(ActionResponseOutput)
}

func (o ActionResponsePtrOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ActionResponse) *int {
		if v == nil {
			return nil
		}
		return v.CompatibilityLevel
	}).(pulumi.IntPtrOutput)
}

func (o ActionResponsePtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ActionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

func (o ActionResponsePtrOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SqlExpression
	}).(pulumi.StringPtrOutput)
}

type ConnectionState struct {
	Description *string `pulumi:"description"`
	Status      *string `pulumi:"status"`
}





type ConnectionStateInput interface {
	pulumi.Input

	ToConnectionStateOutput() ConnectionStateOutput
	ToConnectionStateOutputWithContext(context.Context) ConnectionStateOutput
}

type ConnectionStateArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Status      pulumi.StringPtrInput `pulumi:"status"`
}

func (ConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionState)(nil)).Elem()
}

func (i ConnectionStateArgs) ToConnectionStateOutput() ConnectionStateOutput {
	return i.ToConnectionStateOutputWithContext(context.Background())
}

func (i ConnectionStateArgs) ToConnectionStateOutputWithContext(ctx context.Context) ConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateOutput)
}

func (i ConnectionStateArgs) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return i.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (i ConnectionStateArgs) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateOutput).ToConnectionStatePtrOutputWithContext(ctx)
}









type ConnectionStatePtrInput interface {
	pulumi.Input

	ToConnectionStatePtrOutput() ConnectionStatePtrOutput
	ToConnectionStatePtrOutputWithContext(context.Context) ConnectionStatePtrOutput
}

type connectionStatePtrType ConnectionStateArgs

func ConnectionStatePtr(v *ConnectionStateArgs) ConnectionStatePtrInput {
	return (*connectionStatePtrType)(v)
}

func (*connectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionState)(nil)).Elem()
}

func (i *connectionStatePtrType) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return i.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (i *connectionStatePtrType) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatePtrOutput)
}

type ConnectionStateOutput struct{ *pulumi.OutputState }

func (ConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionState)(nil)).Elem()
}

func (o ConnectionStateOutput) ToConnectionStateOutput() ConnectionStateOutput {
	return o
}

func (o ConnectionStateOutput) ToConnectionStateOutputWithContext(ctx context.Context) ConnectionStateOutput {
	return o
}

func (o ConnectionStateOutput) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return o.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (o ConnectionStateOutput) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionState) *ConnectionState {
		return &v
	}).(ConnectionStatePtrOutput)
}

func (o ConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionState)(nil)).Elem()
}

func (o ConnectionStatePtrOutput) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return o
}

func (o ConnectionStatePtrOutput) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return o
}

func (o ConnectionStatePtrOutput) Elem() ConnectionStateOutput {
	return o.ApplyT(func(v *ConnectionState) ConnectionState {
		if v != nil {
			return *v
		}
		var ret ConnectionState
		return ret
	}).(ConnectionStateOutput)
}

func (o ConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ConnectionStateResponse struct {
	Description *string `pulumi:"description"`
	Status      *string `pulumi:"status"`
}

type ConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (ConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStateResponse)(nil)).Elem()
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponseOutput() ConnectionStateResponseOutput {
	return o
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponseOutputWithContext(ctx context.Context) ConnectionStateResponseOutput {
	return o
}

func (o ConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStateResponse)(nil)).Elem()
}

func (o ConnectionStateResponsePtrOutput) ToConnectionStateResponsePtrOutput() ConnectionStateResponsePtrOutput {
	return o
}

func (o ConnectionStateResponsePtrOutput) ToConnectionStateResponsePtrOutputWithContext(ctx context.Context) ConnectionStateResponsePtrOutput {
	return o
}

func (o ConnectionStateResponsePtrOutput) Elem() ConnectionStateResponseOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) ConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionStateResponse
		return ret
	}).(ConnectionStateResponseOutput)
}

func (o ConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type CorrelationFilter struct {
	ContentType           *string           `pulumi:"contentType"`
	CorrelationId         *string           `pulumi:"correlationId"`
	Label                 *string           `pulumi:"label"`
	MessageId             *string           `pulumi:"messageId"`
	Properties            map[string]string `pulumi:"properties"`
	ReplyTo               *string           `pulumi:"replyTo"`
	ReplyToSessionId      *string           `pulumi:"replyToSessionId"`
	RequiresPreprocessing *bool             `pulumi:"requiresPreprocessing"`
	SessionId             *string           `pulumi:"sessionId"`
	To                    *string           `pulumi:"to"`
}


func (val *CorrelationFilter) Defaults() *CorrelationFilter {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RequiresPreprocessing) {
		requiresPreprocessing_ := true
		tmp.RequiresPreprocessing = &requiresPreprocessing_
	}
	return &tmp
}





type CorrelationFilterInput interface {
	pulumi.Input

	ToCorrelationFilterOutput() CorrelationFilterOutput
	ToCorrelationFilterOutputWithContext(context.Context) CorrelationFilterOutput
}

type CorrelationFilterArgs struct {
	ContentType           pulumi.StringPtrInput `pulumi:"contentType"`
	CorrelationId         pulumi.StringPtrInput `pulumi:"correlationId"`
	Label                 pulumi.StringPtrInput `pulumi:"label"`
	MessageId             pulumi.StringPtrInput `pulumi:"messageId"`
	Properties            pulumi.StringMapInput `pulumi:"properties"`
	ReplyTo               pulumi.StringPtrInput `pulumi:"replyTo"`
	ReplyToSessionId      pulumi.StringPtrInput `pulumi:"replyToSessionId"`
	RequiresPreprocessing pulumi.BoolPtrInput   `pulumi:"requiresPreprocessing"`
	SessionId             pulumi.StringPtrInput `pulumi:"sessionId"`
	To                    pulumi.StringPtrInput `pulumi:"to"`
}

func (CorrelationFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorrelationFilter)(nil)).Elem()
}

func (i CorrelationFilterArgs) ToCorrelationFilterOutput() CorrelationFilterOutput {
	return i.ToCorrelationFilterOutputWithContext(context.Background())
}

func (i CorrelationFilterArgs) ToCorrelationFilterOutputWithContext(ctx context.Context) CorrelationFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorrelationFilterOutput)
}

func (i CorrelationFilterArgs) ToCorrelationFilterPtrOutput() CorrelationFilterPtrOutput {
	return i.ToCorrelationFilterPtrOutputWithContext(context.Background())
}

func (i CorrelationFilterArgs) ToCorrelationFilterPtrOutputWithContext(ctx context.Context) CorrelationFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorrelationFilterOutput).ToCorrelationFilterPtrOutputWithContext(ctx)
}









type CorrelationFilterPtrInput interface {
	pulumi.Input

	ToCorrelationFilterPtrOutput() CorrelationFilterPtrOutput
	ToCorrelationFilterPtrOutputWithContext(context.Context) CorrelationFilterPtrOutput
}

type correlationFilterPtrType CorrelationFilterArgs

func CorrelationFilterPtr(v *CorrelationFilterArgs) CorrelationFilterPtrInput {
	return (*correlationFilterPtrType)(v)
}

func (*correlationFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CorrelationFilter)(nil)).Elem()
}

func (i *correlationFilterPtrType) ToCorrelationFilterPtrOutput() CorrelationFilterPtrOutput {
	return i.ToCorrelationFilterPtrOutputWithContext(context.Background())
}

func (i *correlationFilterPtrType) ToCorrelationFilterPtrOutputWithContext(ctx context.Context) CorrelationFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorrelationFilterPtrOutput)
}

type CorrelationFilterOutput struct{ *pulumi.OutputState }

func (CorrelationFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorrelationFilter)(nil)).Elem()
}

func (o CorrelationFilterOutput) ToCorrelationFilterOutput() CorrelationFilterOutput {
	return o
}

func (o CorrelationFilterOutput) ToCorrelationFilterOutputWithContext(ctx context.Context) CorrelationFilterOutput {
	return o
}

func (o CorrelationFilterOutput) ToCorrelationFilterPtrOutput() CorrelationFilterPtrOutput {
	return o.ToCorrelationFilterPtrOutputWithContext(context.Background())
}

func (o CorrelationFilterOutput) ToCorrelationFilterPtrOutputWithContext(ctx context.Context) CorrelationFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CorrelationFilter) *CorrelationFilter {
		return &v
	}).(CorrelationFilterPtrOutput)
}

func (o CorrelationFilterOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CorrelationFilter) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CorrelationFilterOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.ReplyTo }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.ReplyToSessionId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

func (o CorrelationFilterOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.SessionId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.To }).(pulumi.StringPtrOutput)
}

type CorrelationFilterPtrOutput struct{ *pulumi.OutputState }

func (CorrelationFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorrelationFilter)(nil)).Elem()
}

func (o CorrelationFilterPtrOutput) ToCorrelationFilterPtrOutput() CorrelationFilterPtrOutput {
	return o
}

func (o CorrelationFilterPtrOutput) ToCorrelationFilterPtrOutputWithContext(ctx context.Context) CorrelationFilterPtrOutput {
	return o
}

func (o CorrelationFilterPtrOutput) Elem() CorrelationFilterOutput {
	return o.ApplyT(func(v *CorrelationFilter) CorrelationFilter {
		if v != nil {
			return *v
		}
		var ret CorrelationFilter
		return ret
	}).(CorrelationFilterOutput)
}

func (o CorrelationFilterPtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.ContentType
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.CorrelationId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.MessageId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CorrelationFilter) map[string]string {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.StringMapOutput)
}

func (o CorrelationFilterPtrOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.ReplyTo
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.ReplyToSessionId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

func (o CorrelationFilterPtrOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.SessionId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type CorrelationFilterResponse struct {
	ContentType           *string           `pulumi:"contentType"`
	CorrelationId         *string           `pulumi:"correlationId"`
	Label                 *string           `pulumi:"label"`
	MessageId             *string           `pulumi:"messageId"`
	Properties            map[string]string `pulumi:"properties"`
	ReplyTo               *string           `pulumi:"replyTo"`
	ReplyToSessionId      *string           `pulumi:"replyToSessionId"`
	RequiresPreprocessing *bool             `pulumi:"requiresPreprocessing"`
	SessionId             *string           `pulumi:"sessionId"`
	To                    *string           `pulumi:"to"`
}


func (val *CorrelationFilterResponse) Defaults() *CorrelationFilterResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RequiresPreprocessing) {
		requiresPreprocessing_ := true
		tmp.RequiresPreprocessing = &requiresPreprocessing_
	}
	return &tmp
}

type CorrelationFilterResponseOutput struct{ *pulumi.OutputState }

func (CorrelationFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorrelationFilterResponse)(nil)).Elem()
}

func (o CorrelationFilterResponseOutput) ToCorrelationFilterResponseOutput() CorrelationFilterResponseOutput {
	return o
}

func (o CorrelationFilterResponseOutput) ToCorrelationFilterResponseOutputWithContext(ctx context.Context) CorrelationFilterResponseOutput {
	return o
}

func (o CorrelationFilterResponseOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CorrelationFilterResponseOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.ReplyTo }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.ReplyToSessionId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

func (o CorrelationFilterResponseOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.SessionId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.To }).(pulumi.StringPtrOutput)
}

type CorrelationFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (CorrelationFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorrelationFilterResponse)(nil)).Elem()
}

func (o CorrelationFilterResponsePtrOutput) ToCorrelationFilterResponsePtrOutput() CorrelationFilterResponsePtrOutput {
	return o
}

func (o CorrelationFilterResponsePtrOutput) ToCorrelationFilterResponsePtrOutputWithContext(ctx context.Context) CorrelationFilterResponsePtrOutput {
	return o
}

func (o CorrelationFilterResponsePtrOutput) Elem() CorrelationFilterResponseOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) CorrelationFilterResponse {
		if v != nil {
			return *v
		}
		var ret CorrelationFilterResponse
		return ret
	}).(CorrelationFilterResponseOutput)
}

func (o CorrelationFilterResponsePtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentType
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.CorrelationId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.MessageId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.StringMapOutput)
}

func (o CorrelationFilterResponsePtrOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReplyTo
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReplyToSessionId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SessionId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type DictionaryValueResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type DictionaryValueResponseOutput struct{ *pulumi.OutputState }

func (DictionaryValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DictionaryValueResponse)(nil)).Elem()
}

func (o DictionaryValueResponseOutput) ToDictionaryValueResponseOutput() DictionaryValueResponseOutput {
	return o
}

func (o DictionaryValueResponseOutput) ToDictionaryValueResponseOutputWithContext(ctx context.Context) DictionaryValueResponseOutput {
	return o
}

func (o DictionaryValueResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v DictionaryValueResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o DictionaryValueResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v DictionaryValueResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type DictionaryValueResponseMapOutput struct{ *pulumi.OutputState }

func (DictionaryValueResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DictionaryValueResponse)(nil)).Elem()
}

func (o DictionaryValueResponseMapOutput) ToDictionaryValueResponseMapOutput() DictionaryValueResponseMapOutput {
	return o
}

func (o DictionaryValueResponseMapOutput) ToDictionaryValueResponseMapOutputWithContext(ctx context.Context) DictionaryValueResponseMapOutput {
	return o
}

func (o DictionaryValueResponseMapOutput) MapIndex(k pulumi.StringInput) DictionaryValueResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DictionaryValueResponse {
		return vs[0].(map[string]DictionaryValueResponse)[vs[1].(string)]
	}).(DictionaryValueResponseOutput)
}

type Encryption struct {
	KeySource                       *KeySource           `pulumi:"keySource"`
	KeyVaultProperties              []KeyVaultProperties `pulumi:"keyVaultProperties"`
	RequireInfrastructureEncryption *bool                `pulumi:"requireInfrastructureEncryption"`
}


func (val *Encryption) Defaults() *Encryption {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		keySource_ := KeySource("Microsoft.KeyVault")
		tmp.KeySource = &keySource_
	}
	return &tmp
}





type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

type EncryptionArgs struct {
	KeySource                       KeySourcePtrInput            `pulumi:"keySource"`
	KeyVaultProperties              KeyVaultPropertiesArrayInput `pulumi:"keyVaultProperties"`
	RequireInfrastructureEncryption pulumi.BoolPtrInput          `pulumi:"requireInfrastructureEncryption"`
}

func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}









type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

type EncryptionOutput struct{ *pulumi.OutputState }

func (EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (o EncryptionOutput) ToEncryptionOutput() EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o.ToEncryptionPtrOutputWithContext(context.Background())
}

func (o EncryptionOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Encryption) *Encryption {
		return &v
	}).(EncryptionPtrOutput)
}

func (o EncryptionOutput) KeySource() KeySourcePtrOutput {
	return o.ApplyT(func(v Encryption) *KeySource { return v.KeySource }).(KeySourcePtrOutput)
}

func (o EncryptionOutput) KeyVaultProperties() KeyVaultPropertiesArrayOutput {
	return o.ApplyT(func(v Encryption) []KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesArrayOutput)
}

func (o EncryptionOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Encryption) *bool { return v.RequireInfrastructureEncryption }).(pulumi.BoolPtrOutput)
}

type EncryptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) Elem() EncryptionOutput {
	return o.ApplyT(func(v *Encryption) Encryption {
		if v != nil {
			return *v
		}
		var ret Encryption
		return ret
	}).(EncryptionOutput)
}

func (o EncryptionPtrOutput) KeySource() KeySourcePtrOutput {
	return o.ApplyT(func(v *Encryption) *KeySource {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(KeySourcePtrOutput)
}

func (o EncryptionPtrOutput) KeyVaultProperties() KeyVaultPropertiesArrayOutput {
	return o.ApplyT(func(v *Encryption) []KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesArrayOutput)
}

func (o EncryptionPtrOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Encryption) *bool {
		if v == nil {
			return nil
		}
		return v.RequireInfrastructureEncryption
	}).(pulumi.BoolPtrOutput)
}

type EncryptionResponse struct {
	KeySource                       *string                      `pulumi:"keySource"`
	KeyVaultProperties              []KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	RequireInfrastructureEncryption *bool                        `pulumi:"requireInfrastructureEncryption"`
}


func (val *EncryptionResponse) Defaults() *EncryptionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		keySource_ := "Microsoft.KeyVault"
		tmp.KeySource = &keySource_
	}
	return &tmp
}

type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponseArrayOutput {
	return o.ApplyT(func(v EncryptionResponse) []KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponseArrayOutput)
}

func (o EncryptionResponseOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *bool { return v.RequireInfrastructureEncryption }).(pulumi.BoolPtrOutput)
}

type EncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) Elem() EncryptionResponseOutput {
	return o.ApplyT(func(v *EncryptionResponse) EncryptionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionResponse
		return ret
	}).(EncryptionResponseOutput)
}

func (o EncryptionResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *EncryptionResponse) []KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponseArrayOutput)
}

func (o EncryptionResponsePtrOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequireInfrastructureEncryption
	}).(pulumi.BoolPtrOutput)
}

type Identity struct {
	Type                   *ManagedServiceIdentityType `pulumi:"type"`
	UserAssignedIdentities map[string]interface{}      `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   ManagedServiceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput                    `pulumi:"userAssignedIdentities"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ManagedServiceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ManagedServiceIdentityType { return v.Type }).(ManagedServiceIdentityTypePtrOutput)
}

func (o IdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v Identity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ManagedServiceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ManagedServiceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ManagedServiceIdentityTypePtrOutput)
}

func (o IdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *Identity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityResponse struct {
	PrincipalId            string                             `pulumi:"principalId"`
	TenantId               string                             `pulumi:"tenantId"`
	Type                   *string                            `pulumi:"type"`
	UserAssignedIdentities map[string]DictionaryValueResponse `pulumi:"userAssignedIdentities"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) UserAssignedIdentities() DictionaryValueResponseMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]DictionaryValueResponse { return v.UserAssignedIdentities }).(DictionaryValueResponseMapOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) UserAssignedIdentities() DictionaryValueResponseMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]DictionaryValueResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(DictionaryValueResponseMapOutput)
}

type KeyVaultProperties struct {
	Identity    *UserAssignedIdentityProperties `pulumi:"identity"`
	KeyName     *string                         `pulumi:"keyName"`
	KeyVaultUri *string                         `pulumi:"keyVaultUri"`
	KeyVersion  *string                         `pulumi:"keyVersion"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	Identity    UserAssignedIdentityPropertiesPtrInput `pulumi:"identity"`
	KeyName     pulumi.StringPtrInput                  `pulumi:"keyName"`
	KeyVaultUri pulumi.StringPtrInput                  `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringPtrInput                  `pulumi:"keyVersion"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}





type KeyVaultPropertiesArrayInput interface {
	pulumi.Input

	ToKeyVaultPropertiesArrayOutput() KeyVaultPropertiesArrayOutput
	ToKeyVaultPropertiesArrayOutputWithContext(context.Context) KeyVaultPropertiesArrayOutput
}

type KeyVaultPropertiesArray []KeyVaultPropertiesInput

func (KeyVaultPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArray) ToKeyVaultPropertiesArrayOutput() KeyVaultPropertiesArrayOutput {
	return i.ToKeyVaultPropertiesArrayOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArray) ToKeyVaultPropertiesArrayOutputWithContext(ctx context.Context) KeyVaultPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesArrayOutput)
}

type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) Identity() UserAssignedIdentityPropertiesPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *UserAssignedIdentityProperties { return v.Identity }).(UserAssignedIdentityPropertiesPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesArrayOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesArrayOutput) ToKeyVaultPropertiesArrayOutput() KeyVaultPropertiesArrayOutput {
	return o
}

func (o KeyVaultPropertiesArrayOutput) ToKeyVaultPropertiesArrayOutputWithContext(ctx context.Context) KeyVaultPropertiesArrayOutput {
	return o
}

func (o KeyVaultPropertiesArrayOutput) Index(i pulumi.IntInput) KeyVaultPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyVaultProperties {
		return vs[0].([]KeyVaultProperties)[vs[1].(int)]
	}).(KeyVaultPropertiesOutput)
}

type KeyVaultPropertiesResponse struct {
	Identity    *UserAssignedIdentityPropertiesResponse `pulumi:"identity"`
	KeyName     *string                                 `pulumi:"keyName"`
	KeyVaultUri *string                                 `pulumi:"keyVaultUri"`
	KeyVersion  *string                                 `pulumi:"keyVersion"`
}

type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) Identity() UserAssignedIdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *UserAssignedIdentityPropertiesResponse { return v.Identity }).(UserAssignedIdentityPropertiesResponsePtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseArrayOutput) ToKeyVaultPropertiesResponseArrayOutput() KeyVaultPropertiesResponseArrayOutput {
	return o
}

func (o KeyVaultPropertiesResponseArrayOutput) ToKeyVaultPropertiesResponseArrayOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseArrayOutput {
	return o
}

func (o KeyVaultPropertiesResponseArrayOutput) Index(i pulumi.IntInput) KeyVaultPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyVaultPropertiesResponse {
		return vs[0].([]KeyVaultPropertiesResponse)[vs[1].(int)]
	}).(KeyVaultPropertiesResponseOutput)
}

type MessageCountDetailsResponse struct {
	ActiveMessageCount             float64 `pulumi:"activeMessageCount"`
	DeadLetterMessageCount         float64 `pulumi:"deadLetterMessageCount"`
	ScheduledMessageCount          float64 `pulumi:"scheduledMessageCount"`
	TransferDeadLetterMessageCount float64 `pulumi:"transferDeadLetterMessageCount"`
	TransferMessageCount           float64 `pulumi:"transferMessageCount"`
}

type MessageCountDetailsResponseOutput struct{ *pulumi.OutputState }

func (MessageCountDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MessageCountDetailsResponse)(nil)).Elem()
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponseOutput() MessageCountDetailsResponseOutput {
	return o
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponseOutputWithContext(ctx context.Context) MessageCountDetailsResponseOutput {
	return o
}

func (o MessageCountDetailsResponseOutput) ActiveMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.ActiveMessageCount }).(pulumi.Float64Output)
}

func (o MessageCountDetailsResponseOutput) DeadLetterMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.DeadLetterMessageCount }).(pulumi.Float64Output)
}

func (o MessageCountDetailsResponseOutput) ScheduledMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.ScheduledMessageCount }).(pulumi.Float64Output)
}

func (o MessageCountDetailsResponseOutput) TransferDeadLetterMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.TransferDeadLetterMessageCount }).(pulumi.Float64Output)
}

func (o MessageCountDetailsResponseOutput) TransferMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.TransferMessageCount }).(pulumi.Float64Output)
}

type NWRuleSetIpRules struct {
	Action *string `pulumi:"action"`
	IpMask *string `pulumi:"ipMask"`
}


func (val *NWRuleSetIpRules) Defaults() *NWRuleSetIpRules {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type NWRuleSetIpRulesInput interface {
	pulumi.Input

	ToNWRuleSetIpRulesOutput() NWRuleSetIpRulesOutput
	ToNWRuleSetIpRulesOutputWithContext(context.Context) NWRuleSetIpRulesOutput
}

type NWRuleSetIpRulesArgs struct {
	Action pulumi.StringPtrInput `pulumi:"action"`
	IpMask pulumi.StringPtrInput `pulumi:"ipMask"`
}

func (NWRuleSetIpRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetIpRules)(nil)).Elem()
}

func (i NWRuleSetIpRulesArgs) ToNWRuleSetIpRulesOutput() NWRuleSetIpRulesOutput {
	return i.ToNWRuleSetIpRulesOutputWithContext(context.Background())
}

func (i NWRuleSetIpRulesArgs) ToNWRuleSetIpRulesOutputWithContext(ctx context.Context) NWRuleSetIpRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetIpRulesOutput)
}





type NWRuleSetIpRulesArrayInput interface {
	pulumi.Input

	ToNWRuleSetIpRulesArrayOutput() NWRuleSetIpRulesArrayOutput
	ToNWRuleSetIpRulesArrayOutputWithContext(context.Context) NWRuleSetIpRulesArrayOutput
}

type NWRuleSetIpRulesArray []NWRuleSetIpRulesInput

func (NWRuleSetIpRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetIpRules)(nil)).Elem()
}

func (i NWRuleSetIpRulesArray) ToNWRuleSetIpRulesArrayOutput() NWRuleSetIpRulesArrayOutput {
	return i.ToNWRuleSetIpRulesArrayOutputWithContext(context.Background())
}

func (i NWRuleSetIpRulesArray) ToNWRuleSetIpRulesArrayOutputWithContext(ctx context.Context) NWRuleSetIpRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetIpRulesArrayOutput)
}

type NWRuleSetIpRulesOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetIpRules)(nil)).Elem()
}

func (o NWRuleSetIpRulesOutput) ToNWRuleSetIpRulesOutput() NWRuleSetIpRulesOutput {
	return o
}

func (o NWRuleSetIpRulesOutput) ToNWRuleSetIpRulesOutputWithContext(ctx context.Context) NWRuleSetIpRulesOutput {
	return o
}

func (o NWRuleSetIpRulesOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRules) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NWRuleSetIpRulesOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRules) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type NWRuleSetIpRulesArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetIpRules)(nil)).Elem()
}

func (o NWRuleSetIpRulesArrayOutput) ToNWRuleSetIpRulesArrayOutput() NWRuleSetIpRulesArrayOutput {
	return o
}

func (o NWRuleSetIpRulesArrayOutput) ToNWRuleSetIpRulesArrayOutputWithContext(ctx context.Context) NWRuleSetIpRulesArrayOutput {
	return o
}

func (o NWRuleSetIpRulesArrayOutput) Index(i pulumi.IntInput) NWRuleSetIpRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetIpRules {
		return vs[0].([]NWRuleSetIpRules)[vs[1].(int)]
	}).(NWRuleSetIpRulesOutput)
}

type NWRuleSetIpRulesResponse struct {
	Action *string `pulumi:"action"`
	IpMask *string `pulumi:"ipMask"`
}


func (val *NWRuleSetIpRulesResponse) Defaults() *NWRuleSetIpRulesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}

type NWRuleSetIpRulesResponseOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetIpRulesResponse)(nil)).Elem()
}

func (o NWRuleSetIpRulesResponseOutput) ToNWRuleSetIpRulesResponseOutput() NWRuleSetIpRulesResponseOutput {
	return o
}

func (o NWRuleSetIpRulesResponseOutput) ToNWRuleSetIpRulesResponseOutputWithContext(ctx context.Context) NWRuleSetIpRulesResponseOutput {
	return o
}

func (o NWRuleSetIpRulesResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRulesResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NWRuleSetIpRulesResponseOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRulesResponse) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type NWRuleSetIpRulesResponseArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetIpRulesResponse)(nil)).Elem()
}

func (o NWRuleSetIpRulesResponseArrayOutput) ToNWRuleSetIpRulesResponseArrayOutput() NWRuleSetIpRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetIpRulesResponseArrayOutput) ToNWRuleSetIpRulesResponseArrayOutputWithContext(ctx context.Context) NWRuleSetIpRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetIpRulesResponseArrayOutput) Index(i pulumi.IntInput) NWRuleSetIpRulesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetIpRulesResponse {
		return vs[0].([]NWRuleSetIpRulesResponse)[vs[1].(int)]
	}).(NWRuleSetIpRulesResponseOutput)
}

type NWRuleSetVirtualNetworkRules struct {
	IgnoreMissingVnetServiceEndpoint *bool   `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Subnet                           *Subnet `pulumi:"subnet"`
}





type NWRuleSetVirtualNetworkRulesInput interface {
	pulumi.Input

	ToNWRuleSetVirtualNetworkRulesOutput() NWRuleSetVirtualNetworkRulesOutput
	ToNWRuleSetVirtualNetworkRulesOutputWithContext(context.Context) NWRuleSetVirtualNetworkRulesOutput
}

type NWRuleSetVirtualNetworkRulesArgs struct {
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Subnet                           SubnetPtrInput      `pulumi:"subnet"`
}

func (NWRuleSetVirtualNetworkRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (i NWRuleSetVirtualNetworkRulesArgs) ToNWRuleSetVirtualNetworkRulesOutput() NWRuleSetVirtualNetworkRulesOutput {
	return i.ToNWRuleSetVirtualNetworkRulesOutputWithContext(context.Background())
}

func (i NWRuleSetVirtualNetworkRulesArgs) ToNWRuleSetVirtualNetworkRulesOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetVirtualNetworkRulesOutput)
}





type NWRuleSetVirtualNetworkRulesArrayInput interface {
	pulumi.Input

	ToNWRuleSetVirtualNetworkRulesArrayOutput() NWRuleSetVirtualNetworkRulesArrayOutput
	ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(context.Context) NWRuleSetVirtualNetworkRulesArrayOutput
}

type NWRuleSetVirtualNetworkRulesArray []NWRuleSetVirtualNetworkRulesInput

func (NWRuleSetVirtualNetworkRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (i NWRuleSetVirtualNetworkRulesArray) ToNWRuleSetVirtualNetworkRulesArrayOutput() NWRuleSetVirtualNetworkRulesArrayOutput {
	return i.ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(context.Background())
}

func (i NWRuleSetVirtualNetworkRulesArray) ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetVirtualNetworkRulesArrayOutput)
}

type NWRuleSetVirtualNetworkRulesOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesOutput) ToNWRuleSetVirtualNetworkRulesOutput() NWRuleSetVirtualNetworkRulesOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesOutput) ToNWRuleSetVirtualNetworkRulesOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRules) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o NWRuleSetVirtualNetworkRulesOutput) Subnet() SubnetPtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRules) *Subnet { return v.Subnet }).(SubnetPtrOutput)
}

type NWRuleSetVirtualNetworkRulesArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesArrayOutput) ToNWRuleSetVirtualNetworkRulesArrayOutput() NWRuleSetVirtualNetworkRulesArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesArrayOutput) ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesArrayOutput) Index(i pulumi.IntInput) NWRuleSetVirtualNetworkRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetVirtualNetworkRules {
		return vs[0].([]NWRuleSetVirtualNetworkRules)[vs[1].(int)]
	}).(NWRuleSetVirtualNetworkRulesOutput)
}

type NWRuleSetVirtualNetworkRulesResponse struct {
	IgnoreMissingVnetServiceEndpoint *bool           `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Subnet                           *SubnetResponse `pulumi:"subnet"`
}

type NWRuleSetVirtualNetworkRulesResponseOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetVirtualNetworkRulesResponse)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) ToNWRuleSetVirtualNetworkRulesResponseOutput() NWRuleSetVirtualNetworkRulesResponseOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) ToNWRuleSetVirtualNetworkRulesResponseOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesResponseOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRulesResponse) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRulesResponse) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

type NWRuleSetVirtualNetworkRulesResponseArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetVirtualNetworkRulesResponse)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesResponseArrayOutput) ToNWRuleSetVirtualNetworkRulesResponseArrayOutput() NWRuleSetVirtualNetworkRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseArrayOutput) ToNWRuleSetVirtualNetworkRulesResponseArrayOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseArrayOutput) Index(i pulumi.IntInput) NWRuleSetVirtualNetworkRulesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetVirtualNetworkRulesResponse {
		return vs[0].([]NWRuleSetVirtualNetworkRulesResponse)[vs[1].(int)]
	}).(NWRuleSetVirtualNetworkRulesResponseOutput)
}

type PrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionType struct {
	PrivateEndpoint                   *PrivateEndpoint `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 *string          `pulumi:"provisioningState"`
}





type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	PrivateEndpoint                   PrivateEndpointPtrInput `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState ConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringPtrInput   `pulumi:"provisioningState"`
}

func (PrivateEndpointConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return i.ToPrivateEndpointConnectionTypeOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput)
}





type PrivateEndpointConnectionTypeArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput
	ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Context) PrivateEndpointConnectionTypeArrayOutput
}

type PrivateEndpointConnectionTypeArray []PrivateEndpointConnectionTypeInput

func (PrivateEndpointConnectionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return i.ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeArrayOutput)
}

type PrivateEndpointConnectionTypeOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) PrivateEndpoint() PrivateEndpointPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *PrivateEndpoint { return v.PrivateEndpoint }).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionTypeOutput) PrivateLinkServiceConnectionState() ConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *ConnectionState { return v.PrivateLinkServiceConnectionState }).(ConnectionStatePtrOutput)
}

func (o PrivateEndpointConnectionTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionTypeArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionType {
		return vs[0].([]PrivateEndpointConnectionType)[vs[1].(int)]
	}).(PrivateEndpointConnectionTypeOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                   `pulumi:"id"`
	Name                              string                   `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 *string                  `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse       `pulumi:"systemData"`
	Type                              string                   `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() ConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *ConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id *string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SBSku struct {
	Capacity *int     `pulumi:"capacity"`
	Name     SkuName  `pulumi:"name"`
	Tier     *SkuTier `pulumi:"tier"`
}





type SBSkuInput interface {
	pulumi.Input

	ToSBSkuOutput() SBSkuOutput
	ToSBSkuOutputWithContext(context.Context) SBSkuOutput
}

type SBSkuArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	Name     SkuNameInput       `pulumi:"name"`
	Tier     SkuTierPtrInput    `pulumi:"tier"`
}

func (SBSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SBSku)(nil)).Elem()
}

func (i SBSkuArgs) ToSBSkuOutput() SBSkuOutput {
	return i.ToSBSkuOutputWithContext(context.Background())
}

func (i SBSkuArgs) ToSBSkuOutputWithContext(ctx context.Context) SBSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuOutput)
}

func (i SBSkuArgs) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return i.ToSBSkuPtrOutputWithContext(context.Background())
}

func (i SBSkuArgs) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuOutput).ToSBSkuPtrOutputWithContext(ctx)
}









type SBSkuPtrInput interface {
	pulumi.Input

	ToSBSkuPtrOutput() SBSkuPtrOutput
	ToSBSkuPtrOutputWithContext(context.Context) SBSkuPtrOutput
}

type sbskuPtrType SBSkuArgs

func SBSkuPtr(v *SBSkuArgs) SBSkuPtrInput {
	return (*sbskuPtrType)(v)
}

func (*sbskuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SBSku)(nil)).Elem()
}

func (i *sbskuPtrType) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return i.ToSBSkuPtrOutputWithContext(context.Background())
}

func (i *sbskuPtrType) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuPtrOutput)
}

type SBSkuOutput struct{ *pulumi.OutputState }

func (SBSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SBSku)(nil)).Elem()
}

func (o SBSkuOutput) ToSBSkuOutput() SBSkuOutput {
	return o
}

func (o SBSkuOutput) ToSBSkuOutputWithContext(ctx context.Context) SBSkuOutput {
	return o
}

func (o SBSkuOutput) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return o.ToSBSkuPtrOutputWithContext(context.Background())
}

func (o SBSkuOutput) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SBSku) *SBSku {
		return &v
	}).(SBSkuPtrOutput)
}

func (o SBSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SBSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SBSkuOutput) Name() SkuNameOutput {
	return o.ApplyT(func(v SBSku) SkuName { return v.Name }).(SkuNameOutput)
}

func (o SBSkuOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v SBSku) *SkuTier { return v.Tier }).(SkuTierPtrOutput)
}

type SBSkuPtrOutput struct{ *pulumi.OutputState }

func (SBSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SBSku)(nil)).Elem()
}

func (o SBSkuPtrOutput) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return o
}

func (o SBSkuPtrOutput) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return o
}

func (o SBSkuPtrOutput) Elem() SBSkuOutput {
	return o.ApplyT(func(v *SBSku) SBSku {
		if v != nil {
			return *v
		}
		var ret SBSku
		return ret
	}).(SBSkuOutput)
}

func (o SBSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SBSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SBSkuPtrOutput) Name() SkuNamePtrOutput {
	return o.ApplyT(func(v *SBSku) *SkuName {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(SkuNamePtrOutput)
}

func (o SBSkuPtrOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v *SBSku) *SkuTier {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(SkuTierPtrOutput)
}

type SBSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}

type SBSkuResponseOutput struct{ *pulumi.OutputState }

func (SBSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SBSkuResponse)(nil)).Elem()
}

func (o SBSkuResponseOutput) ToSBSkuResponseOutput() SBSkuResponseOutput {
	return o
}

func (o SBSkuResponseOutput) ToSBSkuResponseOutputWithContext(ctx context.Context) SBSkuResponseOutput {
	return o
}

func (o SBSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SBSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SBSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SBSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SBSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SBSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SBSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SBSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SBSkuResponse)(nil)).Elem()
}

func (o SBSkuResponsePtrOutput) ToSBSkuResponsePtrOutput() SBSkuResponsePtrOutput {
	return o
}

func (o SBSkuResponsePtrOutput) ToSBSkuResponsePtrOutputWithContext(ctx context.Context) SBSkuResponsePtrOutput {
	return o
}

func (o SBSkuResponsePtrOutput) Elem() SBSkuResponseOutput {
	return o.ApplyT(func(v *SBSkuResponse) SBSkuResponse {
		if v != nil {
			return *v
		}
		var ret SBSkuResponse
		return ret
	}).(SBSkuResponseOutput)
}

func (o SBSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SBSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SBSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SBSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SBSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SBSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SqlFilter struct {
	CompatibilityLevel    *int    `pulumi:"compatibilityLevel"`
	RequiresPreprocessing *bool   `pulumi:"requiresPreprocessing"`
	SqlExpression         *string `pulumi:"sqlExpression"`
}


func (val *SqlFilter) Defaults() *SqlFilter {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CompatibilityLevel) {
		compatibilityLevel_ := 20
		tmp.CompatibilityLevel = &compatibilityLevel_
	}
	if isZero(tmp.RequiresPreprocessing) {
		requiresPreprocessing_ := true
		tmp.RequiresPreprocessing = &requiresPreprocessing_
	}
	return &tmp
}





type SqlFilterInput interface {
	pulumi.Input

	ToSqlFilterOutput() SqlFilterOutput
	ToSqlFilterOutputWithContext(context.Context) SqlFilterOutput
}

type SqlFilterArgs struct {
	CompatibilityLevel    pulumi.IntPtrInput    `pulumi:"compatibilityLevel"`
	RequiresPreprocessing pulumi.BoolPtrInput   `pulumi:"requiresPreprocessing"`
	SqlExpression         pulumi.StringPtrInput `pulumi:"sqlExpression"`
}

func (SqlFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlFilter)(nil)).Elem()
}

func (i SqlFilterArgs) ToSqlFilterOutput() SqlFilterOutput {
	return i.ToSqlFilterOutputWithContext(context.Background())
}

func (i SqlFilterArgs) ToSqlFilterOutputWithContext(ctx context.Context) SqlFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlFilterOutput)
}

func (i SqlFilterArgs) ToSqlFilterPtrOutput() SqlFilterPtrOutput {
	return i.ToSqlFilterPtrOutputWithContext(context.Background())
}

func (i SqlFilterArgs) ToSqlFilterPtrOutputWithContext(ctx context.Context) SqlFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlFilterOutput).ToSqlFilterPtrOutputWithContext(ctx)
}









type SqlFilterPtrInput interface {
	pulumi.Input

	ToSqlFilterPtrOutput() SqlFilterPtrOutput
	ToSqlFilterPtrOutputWithContext(context.Context) SqlFilterPtrOutput
}

type sqlFilterPtrType SqlFilterArgs

func SqlFilterPtr(v *SqlFilterArgs) SqlFilterPtrInput {
	return (*sqlFilterPtrType)(v)
}

func (*sqlFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlFilter)(nil)).Elem()
}

func (i *sqlFilterPtrType) ToSqlFilterPtrOutput() SqlFilterPtrOutput {
	return i.ToSqlFilterPtrOutputWithContext(context.Background())
}

func (i *sqlFilterPtrType) ToSqlFilterPtrOutputWithContext(ctx context.Context) SqlFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlFilterPtrOutput)
}

type SqlFilterOutput struct{ *pulumi.OutputState }

func (SqlFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlFilter)(nil)).Elem()
}

func (o SqlFilterOutput) ToSqlFilterOutput() SqlFilterOutput {
	return o
}

func (o SqlFilterOutput) ToSqlFilterOutputWithContext(ctx context.Context) SqlFilterOutput {
	return o
}

func (o SqlFilterOutput) ToSqlFilterPtrOutput() SqlFilterPtrOutput {
	return o.ToSqlFilterPtrOutputWithContext(context.Background())
}

func (o SqlFilterOutput) ToSqlFilterPtrOutputWithContext(ctx context.Context) SqlFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlFilter) *SqlFilter {
		return &v
	}).(SqlFilterPtrOutput)
}

func (o SqlFilterOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlFilter) *int { return v.CompatibilityLevel }).(pulumi.IntPtrOutput)
}

func (o SqlFilterOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlFilter) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

func (o SqlFilterOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlFilter) *string { return v.SqlExpression }).(pulumi.StringPtrOutput)
}

type SqlFilterPtrOutput struct{ *pulumi.OutputState }

func (SqlFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlFilter)(nil)).Elem()
}

func (o SqlFilterPtrOutput) ToSqlFilterPtrOutput() SqlFilterPtrOutput {
	return o
}

func (o SqlFilterPtrOutput) ToSqlFilterPtrOutputWithContext(ctx context.Context) SqlFilterPtrOutput {
	return o
}

func (o SqlFilterPtrOutput) Elem() SqlFilterOutput {
	return o.ApplyT(func(v *SqlFilter) SqlFilter {
		if v != nil {
			return *v
		}
		var ret SqlFilter
		return ret
	}).(SqlFilterOutput)
}

func (o SqlFilterPtrOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlFilter) *int {
		if v == nil {
			return nil
		}
		return v.CompatibilityLevel
	}).(pulumi.IntPtrOutput)
}

func (o SqlFilterPtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlFilter) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

func (o SqlFilterPtrOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlFilter) *string {
		if v == nil {
			return nil
		}
		return v.SqlExpression
	}).(pulumi.StringPtrOutput)
}

type SqlFilterResponse struct {
	CompatibilityLevel    *int    `pulumi:"compatibilityLevel"`
	RequiresPreprocessing *bool   `pulumi:"requiresPreprocessing"`
	SqlExpression         *string `pulumi:"sqlExpression"`
}


func (val *SqlFilterResponse) Defaults() *SqlFilterResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CompatibilityLevel) {
		compatibilityLevel_ := 20
		tmp.CompatibilityLevel = &compatibilityLevel_
	}
	if isZero(tmp.RequiresPreprocessing) {
		requiresPreprocessing_ := true
		tmp.RequiresPreprocessing = &requiresPreprocessing_
	}
	return &tmp
}

type SqlFilterResponseOutput struct{ *pulumi.OutputState }

func (SqlFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlFilterResponse)(nil)).Elem()
}

func (o SqlFilterResponseOutput) ToSqlFilterResponseOutput() SqlFilterResponseOutput {
	return o
}

func (o SqlFilterResponseOutput) ToSqlFilterResponseOutputWithContext(ctx context.Context) SqlFilterResponseOutput {
	return o
}

func (o SqlFilterResponseOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlFilterResponse) *int { return v.CompatibilityLevel }).(pulumi.IntPtrOutput)
}

func (o SqlFilterResponseOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlFilterResponse) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

func (o SqlFilterResponseOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlFilterResponse) *string { return v.SqlExpression }).(pulumi.StringPtrOutput)
}

type SqlFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlFilterResponse)(nil)).Elem()
}

func (o SqlFilterResponsePtrOutput) ToSqlFilterResponsePtrOutput() SqlFilterResponsePtrOutput {
	return o
}

func (o SqlFilterResponsePtrOutput) ToSqlFilterResponsePtrOutputWithContext(ctx context.Context) SqlFilterResponsePtrOutput {
	return o
}

func (o SqlFilterResponsePtrOutput) Elem() SqlFilterResponseOutput {
	return o.ApplyT(func(v *SqlFilterResponse) SqlFilterResponse {
		if v != nil {
			return *v
		}
		var ret SqlFilterResponse
		return ret
	}).(SqlFilterResponseOutput)
}

func (o SqlFilterResponsePtrOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlFilterResponse) *int {
		if v == nil {
			return nil
		}
		return v.CompatibilityLevel
	}).(pulumi.IntPtrOutput)
}

func (o SqlFilterResponsePtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlFilterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

func (o SqlFilterResponsePtrOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SqlExpression
	}).(pulumi.StringPtrOutput)
}

type Subnet struct {
	Id string `pulumi:"id"`
}





type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(context.Context) SubnetOutput
}

type SubnetArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (i SubnetArgs) ToSubnetOutput() SubnetOutput {
	return i.ToSubnetOutputWithContext(context.Background())
}

func (i SubnetArgs) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput)
}

func (i SubnetArgs) ToSubnetPtrOutput() SubnetPtrOutput {
	return i.ToSubnetPtrOutputWithContext(context.Background())
}

func (i SubnetArgs) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput).ToSubnetPtrOutputWithContext(ctx)
}









type SubnetPtrInput interface {
	pulumi.Input

	ToSubnetPtrOutput() SubnetPtrOutput
	ToSubnetPtrOutputWithContext(context.Context) SubnetPtrOutput
}

type subnetPtrType SubnetArgs

func SubnetPtr(v *SubnetArgs) SubnetPtrInput {
	return (*subnetPtrType)(v)
}

func (*subnetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (i *subnetPtrType) ToSubnetPtrOutput() SubnetPtrOutput {
	return i.ToSubnetPtrOutputWithContext(context.Background())
}

func (i *subnetPtrType) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPtrOutput)
}

type SubnetOutput struct{ *pulumi.OutputState }

func (SubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (o SubnetOutput) ToSubnetOutput() SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetPtrOutput() SubnetPtrOutput {
	return o.ToSubnetPtrOutputWithContext(context.Background())
}

func (o SubnetOutput) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Subnet) *Subnet {
		return &v
	}).(SubnetPtrOutput)
}

func (o SubnetOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Subnet) string { return v.Id }).(pulumi.StringOutput)
}

type SubnetPtrOutput struct{ *pulumi.OutputState }

func (SubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (o SubnetPtrOutput) ToSubnetPtrOutput() SubnetPtrOutput {
	return o
}

func (o SubnetPtrOutput) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return o
}

func (o SubnetPtrOutput) Elem() SubnetOutput {
	return o.ApplyT(func(v *Subnet) Subnet {
		if v != nil {
			return *v
		}
		var ret Subnet
		return ret
	}).(SubnetOutput)
}

func (o SubnetPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type SubnetResponse struct {
	Id string `pulumi:"id"`
}

type SubnetResponseOutput struct{ *pulumi.OutputState }

func (SubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseOutput) ToSubnetResponseOutput() SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) ToSubnetResponseOutputWithContext(ctx context.Context) SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetResponse) string { return v.Id }).(pulumi.StringOutput)
}

type SubnetResponsePtrOutput struct{ *pulumi.OutputState }

func (SubnetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetResponse)(nil)).Elem()
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutput() SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutputWithContext(ctx context.Context) SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) Elem() SubnetResponseOutput {
	return o.ApplyT(func(v *SubnetResponse) SubnetResponse {
		if v != nil {
			return *v
		}
		var ret SubnetResponse
		return ret
	}).(SubnetResponseOutput)
}

func (o SubnetResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
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

type UserAssignedIdentityProperties struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}





type UserAssignedIdentityPropertiesInput interface {
	pulumi.Input

	ToUserAssignedIdentityPropertiesOutput() UserAssignedIdentityPropertiesOutput
	ToUserAssignedIdentityPropertiesOutputWithContext(context.Context) UserAssignedIdentityPropertiesOutput
}

type UserAssignedIdentityPropertiesArgs struct {
	UserAssignedIdentity pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (UserAssignedIdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityProperties)(nil)).Elem()
}

func (i UserAssignedIdentityPropertiesArgs) ToUserAssignedIdentityPropertiesOutput() UserAssignedIdentityPropertiesOutput {
	return i.ToUserAssignedIdentityPropertiesOutputWithContext(context.Background())
}

func (i UserAssignedIdentityPropertiesArgs) ToUserAssignedIdentityPropertiesOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityPropertiesOutput)
}

func (i UserAssignedIdentityPropertiesArgs) ToUserAssignedIdentityPropertiesPtrOutput() UserAssignedIdentityPropertiesPtrOutput {
	return i.ToUserAssignedIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i UserAssignedIdentityPropertiesArgs) ToUserAssignedIdentityPropertiesPtrOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityPropertiesOutput).ToUserAssignedIdentityPropertiesPtrOutputWithContext(ctx)
}









type UserAssignedIdentityPropertiesPtrInput interface {
	pulumi.Input

	ToUserAssignedIdentityPropertiesPtrOutput() UserAssignedIdentityPropertiesPtrOutput
	ToUserAssignedIdentityPropertiesPtrOutputWithContext(context.Context) UserAssignedIdentityPropertiesPtrOutput
}

type userAssignedIdentityPropertiesPtrType UserAssignedIdentityPropertiesArgs

func UserAssignedIdentityPropertiesPtr(v *UserAssignedIdentityPropertiesArgs) UserAssignedIdentityPropertiesPtrInput {
	return (*userAssignedIdentityPropertiesPtrType)(v)
}

func (*userAssignedIdentityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedIdentityProperties)(nil)).Elem()
}

func (i *userAssignedIdentityPropertiesPtrType) ToUserAssignedIdentityPropertiesPtrOutput() UserAssignedIdentityPropertiesPtrOutput {
	return i.ToUserAssignedIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i *userAssignedIdentityPropertiesPtrType) ToUserAssignedIdentityPropertiesPtrOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityPropertiesPtrOutput)
}

type UserAssignedIdentityPropertiesOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityProperties)(nil)).Elem()
}

func (o UserAssignedIdentityPropertiesOutput) ToUserAssignedIdentityPropertiesOutput() UserAssignedIdentityPropertiesOutput {
	return o
}

func (o UserAssignedIdentityPropertiesOutput) ToUserAssignedIdentityPropertiesOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesOutput {
	return o
}

func (o UserAssignedIdentityPropertiesOutput) ToUserAssignedIdentityPropertiesPtrOutput() UserAssignedIdentityPropertiesPtrOutput {
	return o.ToUserAssignedIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (o UserAssignedIdentityPropertiesOutput) ToUserAssignedIdentityPropertiesPtrOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserAssignedIdentityProperties) *UserAssignedIdentityProperties {
		return &v
	}).(UserAssignedIdentityPropertiesPtrOutput)
}

func (o UserAssignedIdentityPropertiesOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityProperties) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedIdentityProperties)(nil)).Elem()
}

func (o UserAssignedIdentityPropertiesPtrOutput) ToUserAssignedIdentityPropertiesPtrOutput() UserAssignedIdentityPropertiesPtrOutput {
	return o
}

func (o UserAssignedIdentityPropertiesPtrOutput) ToUserAssignedIdentityPropertiesPtrOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesPtrOutput {
	return o
}

func (o UserAssignedIdentityPropertiesPtrOutput) Elem() UserAssignedIdentityPropertiesOutput {
	return o.ApplyT(func(v *UserAssignedIdentityProperties) UserAssignedIdentityProperties {
		if v != nil {
			return *v
		}
		var ret UserAssignedIdentityProperties
		return ret
	}).(UserAssignedIdentityPropertiesOutput)
}

func (o UserAssignedIdentityPropertiesPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAssignedIdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityPropertiesResponse struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}

type UserAssignedIdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserAssignedIdentityPropertiesResponseOutput) ToUserAssignedIdentityPropertiesResponseOutput() UserAssignedIdentityPropertiesResponseOutput {
	return o
}

func (o UserAssignedIdentityPropertiesResponseOutput) ToUserAssignedIdentityPropertiesResponseOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesResponseOutput {
	return o
}

func (o UserAssignedIdentityPropertiesResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityPropertiesResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserAssignedIdentityPropertiesResponsePtrOutput) ToUserAssignedIdentityPropertiesResponsePtrOutput() UserAssignedIdentityPropertiesResponsePtrOutput {
	return o
}

func (o UserAssignedIdentityPropertiesResponsePtrOutput) ToUserAssignedIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesResponsePtrOutput {
	return o
}

func (o UserAssignedIdentityPropertiesResponsePtrOutput) Elem() UserAssignedIdentityPropertiesResponseOutput {
	return o.ApplyT(func(v *UserAssignedIdentityPropertiesResponse) UserAssignedIdentityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret UserAssignedIdentityPropertiesResponse
		return ret
	}).(UserAssignedIdentityPropertiesResponseOutput)
}

func (o UserAssignedIdentityPropertiesResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAssignedIdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionPtrOutput{})
	pulumi.RegisterOutputType(ActionResponseOutput{})
	pulumi.RegisterOutputType(ActionResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStateOutput{})
	pulumi.RegisterOutputType(ConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(ConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(CorrelationFilterOutput{})
	pulumi.RegisterOutputType(CorrelationFilterPtrOutput{})
	pulumi.RegisterOutputType(CorrelationFilterResponseOutput{})
	pulumi.RegisterOutputType(CorrelationFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(DictionaryValueResponseOutput{})
	pulumi.RegisterOutputType(DictionaryValueResponseMapOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(MessageCountDetailsResponseOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesResponseOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesResponseArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesResponseOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(SBSkuOutput{})
	pulumi.RegisterOutputType(SBSkuPtrOutput{})
	pulumi.RegisterOutputType(SBSkuResponseOutput{})
	pulumi.RegisterOutputType(SBSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlFilterOutput{})
	pulumi.RegisterOutputType(SqlFilterPtrOutput{})
	pulumi.RegisterOutputType(SqlFilterResponseOutput{})
	pulumi.RegisterOutputType(SqlFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(SubnetOutput{})
	pulumi.RegisterOutputType(SubnetPtrOutput{})
	pulumi.RegisterOutputType(SubnetResponseOutput{})
	pulumi.RegisterOutputType(SubnetResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityPropertiesOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityPropertiesResponsePtrOutput{})
}
