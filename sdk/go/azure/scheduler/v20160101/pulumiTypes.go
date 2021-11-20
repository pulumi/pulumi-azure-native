


package v20160101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HttpAuthentication struct {
	Type *HttpAuthenticationType `pulumi:"type"`
}





type HttpAuthenticationInput interface {
	pulumi.Input

	ToHttpAuthenticationOutput() HttpAuthenticationOutput
	ToHttpAuthenticationOutputWithContext(context.Context) HttpAuthenticationOutput
}

type HttpAuthenticationArgs struct {
	Type HttpAuthenticationTypePtrInput `pulumi:"type"`
}

func (HttpAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpAuthentication)(nil)).Elem()
}

func (i HttpAuthenticationArgs) ToHttpAuthenticationOutput() HttpAuthenticationOutput {
	return i.ToHttpAuthenticationOutputWithContext(context.Background())
}

func (i HttpAuthenticationArgs) ToHttpAuthenticationOutputWithContext(ctx context.Context) HttpAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpAuthenticationOutput)
}

func (i HttpAuthenticationArgs) ToHttpAuthenticationPtrOutput() HttpAuthenticationPtrOutput {
	return i.ToHttpAuthenticationPtrOutputWithContext(context.Background())
}

func (i HttpAuthenticationArgs) ToHttpAuthenticationPtrOutputWithContext(ctx context.Context) HttpAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpAuthenticationOutput).ToHttpAuthenticationPtrOutputWithContext(ctx)
}









type HttpAuthenticationPtrInput interface {
	pulumi.Input

	ToHttpAuthenticationPtrOutput() HttpAuthenticationPtrOutput
	ToHttpAuthenticationPtrOutputWithContext(context.Context) HttpAuthenticationPtrOutput
}

type httpAuthenticationPtrType HttpAuthenticationArgs

func HttpAuthenticationPtr(v *HttpAuthenticationArgs) HttpAuthenticationPtrInput {
	return (*httpAuthenticationPtrType)(v)
}

func (*httpAuthenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpAuthentication)(nil)).Elem()
}

func (i *httpAuthenticationPtrType) ToHttpAuthenticationPtrOutput() HttpAuthenticationPtrOutput {
	return i.ToHttpAuthenticationPtrOutputWithContext(context.Background())
}

func (i *httpAuthenticationPtrType) ToHttpAuthenticationPtrOutputWithContext(ctx context.Context) HttpAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpAuthenticationPtrOutput)
}

type HttpAuthenticationOutput struct{ *pulumi.OutputState }

func (HttpAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpAuthentication)(nil)).Elem()
}

func (o HttpAuthenticationOutput) ToHttpAuthenticationOutput() HttpAuthenticationOutput {
	return o
}

func (o HttpAuthenticationOutput) ToHttpAuthenticationOutputWithContext(ctx context.Context) HttpAuthenticationOutput {
	return o
}

func (o HttpAuthenticationOutput) ToHttpAuthenticationPtrOutput() HttpAuthenticationPtrOutput {
	return o.ToHttpAuthenticationPtrOutputWithContext(context.Background())
}

func (o HttpAuthenticationOutput) ToHttpAuthenticationPtrOutputWithContext(ctx context.Context) HttpAuthenticationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpAuthentication) *HttpAuthentication {
		return &v
	}).(HttpAuthenticationPtrOutput)
}

func (o HttpAuthenticationOutput) Type() HttpAuthenticationTypePtrOutput {
	return o.ApplyT(func(v HttpAuthentication) *HttpAuthenticationType { return v.Type }).(HttpAuthenticationTypePtrOutput)
}

type HttpAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (HttpAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpAuthentication)(nil)).Elem()
}

func (o HttpAuthenticationPtrOutput) ToHttpAuthenticationPtrOutput() HttpAuthenticationPtrOutput {
	return o
}

func (o HttpAuthenticationPtrOutput) ToHttpAuthenticationPtrOutputWithContext(ctx context.Context) HttpAuthenticationPtrOutput {
	return o
}

func (o HttpAuthenticationPtrOutput) Elem() HttpAuthenticationOutput {
	return o.ApplyT(func(v *HttpAuthentication) HttpAuthentication {
		if v != nil {
			return *v
		}
		var ret HttpAuthentication
		return ret
	}).(HttpAuthenticationOutput)
}

func (o HttpAuthenticationPtrOutput) Type() HttpAuthenticationTypePtrOutput {
	return o.ApplyT(func(v *HttpAuthentication) *HttpAuthenticationType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(HttpAuthenticationTypePtrOutput)
}

type HttpAuthenticationResponse struct {
	Type *string `pulumi:"type"`
}





type HttpAuthenticationResponseInput interface {
	pulumi.Input

	ToHttpAuthenticationResponseOutput() HttpAuthenticationResponseOutput
	ToHttpAuthenticationResponseOutputWithContext(context.Context) HttpAuthenticationResponseOutput
}

type HttpAuthenticationResponseArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (HttpAuthenticationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpAuthenticationResponse)(nil)).Elem()
}

func (i HttpAuthenticationResponseArgs) ToHttpAuthenticationResponseOutput() HttpAuthenticationResponseOutput {
	return i.ToHttpAuthenticationResponseOutputWithContext(context.Background())
}

func (i HttpAuthenticationResponseArgs) ToHttpAuthenticationResponseOutputWithContext(ctx context.Context) HttpAuthenticationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpAuthenticationResponseOutput)
}

func (i HttpAuthenticationResponseArgs) ToHttpAuthenticationResponsePtrOutput() HttpAuthenticationResponsePtrOutput {
	return i.ToHttpAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (i HttpAuthenticationResponseArgs) ToHttpAuthenticationResponsePtrOutputWithContext(ctx context.Context) HttpAuthenticationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpAuthenticationResponseOutput).ToHttpAuthenticationResponsePtrOutputWithContext(ctx)
}









type HttpAuthenticationResponsePtrInput interface {
	pulumi.Input

	ToHttpAuthenticationResponsePtrOutput() HttpAuthenticationResponsePtrOutput
	ToHttpAuthenticationResponsePtrOutputWithContext(context.Context) HttpAuthenticationResponsePtrOutput
}

type httpAuthenticationResponsePtrType HttpAuthenticationResponseArgs

func HttpAuthenticationResponsePtr(v *HttpAuthenticationResponseArgs) HttpAuthenticationResponsePtrInput {
	return (*httpAuthenticationResponsePtrType)(v)
}

func (*httpAuthenticationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpAuthenticationResponse)(nil)).Elem()
}

func (i *httpAuthenticationResponsePtrType) ToHttpAuthenticationResponsePtrOutput() HttpAuthenticationResponsePtrOutput {
	return i.ToHttpAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (i *httpAuthenticationResponsePtrType) ToHttpAuthenticationResponsePtrOutputWithContext(ctx context.Context) HttpAuthenticationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpAuthenticationResponsePtrOutput)
}

type HttpAuthenticationResponseOutput struct{ *pulumi.OutputState }

func (HttpAuthenticationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpAuthenticationResponse)(nil)).Elem()
}

func (o HttpAuthenticationResponseOutput) ToHttpAuthenticationResponseOutput() HttpAuthenticationResponseOutput {
	return o
}

func (o HttpAuthenticationResponseOutput) ToHttpAuthenticationResponseOutputWithContext(ctx context.Context) HttpAuthenticationResponseOutput {
	return o
}

func (o HttpAuthenticationResponseOutput) ToHttpAuthenticationResponsePtrOutput() HttpAuthenticationResponsePtrOutput {
	return o.ToHttpAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (o HttpAuthenticationResponseOutput) ToHttpAuthenticationResponsePtrOutputWithContext(ctx context.Context) HttpAuthenticationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpAuthenticationResponse) *HttpAuthenticationResponse {
		return &v
	}).(HttpAuthenticationResponsePtrOutput)
}

func (o HttpAuthenticationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpAuthenticationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type HttpAuthenticationResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpAuthenticationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpAuthenticationResponse)(nil)).Elem()
}

func (o HttpAuthenticationResponsePtrOutput) ToHttpAuthenticationResponsePtrOutput() HttpAuthenticationResponsePtrOutput {
	return o
}

func (o HttpAuthenticationResponsePtrOutput) ToHttpAuthenticationResponsePtrOutputWithContext(ctx context.Context) HttpAuthenticationResponsePtrOutput {
	return o
}

func (o HttpAuthenticationResponsePtrOutput) Elem() HttpAuthenticationResponseOutput {
	return o.ApplyT(func(v *HttpAuthenticationResponse) HttpAuthenticationResponse {
		if v != nil {
			return *v
		}
		var ret HttpAuthenticationResponse
		return ret
	}).(HttpAuthenticationResponseOutput)
}

func (o HttpAuthenticationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpAuthenticationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type HttpRequest struct {
	Authentication *HttpAuthentication `pulumi:"authentication"`
	Body           *string             `pulumi:"body"`
	Headers        map[string]string   `pulumi:"headers"`
	Method         *string             `pulumi:"method"`
	Uri            *string             `pulumi:"uri"`
}





type HttpRequestInput interface {
	pulumi.Input

	ToHttpRequestOutput() HttpRequestOutput
	ToHttpRequestOutputWithContext(context.Context) HttpRequestOutput
}

type HttpRequestArgs struct {
	Authentication HttpAuthenticationPtrInput `pulumi:"authentication"`
	Body           pulumi.StringPtrInput      `pulumi:"body"`
	Headers        pulumi.StringMapInput      `pulumi:"headers"`
	Method         pulumi.StringPtrInput      `pulumi:"method"`
	Uri            pulumi.StringPtrInput      `pulumi:"uri"`
}

func (HttpRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRequest)(nil)).Elem()
}

func (i HttpRequestArgs) ToHttpRequestOutput() HttpRequestOutput {
	return i.ToHttpRequestOutputWithContext(context.Background())
}

func (i HttpRequestArgs) ToHttpRequestOutputWithContext(ctx context.Context) HttpRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRequestOutput)
}

func (i HttpRequestArgs) ToHttpRequestPtrOutput() HttpRequestPtrOutput {
	return i.ToHttpRequestPtrOutputWithContext(context.Background())
}

func (i HttpRequestArgs) ToHttpRequestPtrOutputWithContext(ctx context.Context) HttpRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRequestOutput).ToHttpRequestPtrOutputWithContext(ctx)
}









type HttpRequestPtrInput interface {
	pulumi.Input

	ToHttpRequestPtrOutput() HttpRequestPtrOutput
	ToHttpRequestPtrOutputWithContext(context.Context) HttpRequestPtrOutput
}

type httpRequestPtrType HttpRequestArgs

func HttpRequestPtr(v *HttpRequestArgs) HttpRequestPtrInput {
	return (*httpRequestPtrType)(v)
}

func (*httpRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpRequest)(nil)).Elem()
}

func (i *httpRequestPtrType) ToHttpRequestPtrOutput() HttpRequestPtrOutput {
	return i.ToHttpRequestPtrOutputWithContext(context.Background())
}

func (i *httpRequestPtrType) ToHttpRequestPtrOutputWithContext(ctx context.Context) HttpRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRequestPtrOutput)
}

type HttpRequestOutput struct{ *pulumi.OutputState }

func (HttpRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRequest)(nil)).Elem()
}

func (o HttpRequestOutput) ToHttpRequestOutput() HttpRequestOutput {
	return o
}

func (o HttpRequestOutput) ToHttpRequestOutputWithContext(ctx context.Context) HttpRequestOutput {
	return o
}

func (o HttpRequestOutput) ToHttpRequestPtrOutput() HttpRequestPtrOutput {
	return o.ToHttpRequestPtrOutputWithContext(context.Background())
}

func (o HttpRequestOutput) ToHttpRequestPtrOutputWithContext(ctx context.Context) HttpRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpRequest) *HttpRequest {
		return &v
	}).(HttpRequestPtrOutput)
}

func (o HttpRequestOutput) Authentication() HttpAuthenticationPtrOutput {
	return o.ApplyT(func(v HttpRequest) *HttpAuthentication { return v.Authentication }).(HttpAuthenticationPtrOutput)
}

func (o HttpRequestOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRequest) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o HttpRequestOutput) Headers() pulumi.StringMapOutput {
	return o.ApplyT(func(v HttpRequest) map[string]string { return v.Headers }).(pulumi.StringMapOutput)
}

func (o HttpRequestOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRequest) *string { return v.Method }).(pulumi.StringPtrOutput)
}

func (o HttpRequestOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRequest) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type HttpRequestPtrOutput struct{ *pulumi.OutputState }

func (HttpRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpRequest)(nil)).Elem()
}

func (o HttpRequestPtrOutput) ToHttpRequestPtrOutput() HttpRequestPtrOutput {
	return o
}

func (o HttpRequestPtrOutput) ToHttpRequestPtrOutputWithContext(ctx context.Context) HttpRequestPtrOutput {
	return o
}

func (o HttpRequestPtrOutput) Elem() HttpRequestOutput {
	return o.ApplyT(func(v *HttpRequest) HttpRequest {
		if v != nil {
			return *v
		}
		var ret HttpRequest
		return ret
	}).(HttpRequestOutput)
}

func (o HttpRequestPtrOutput) Authentication() HttpAuthenticationPtrOutput {
	return o.ApplyT(func(v *HttpRequest) *HttpAuthentication {
		if v == nil {
			return nil
		}
		return v.Authentication
	}).(HttpAuthenticationPtrOutput)
}

func (o HttpRequestPtrOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpRequest) *string {
		if v == nil {
			return nil
		}
		return v.Body
	}).(pulumi.StringPtrOutput)
}

func (o HttpRequestPtrOutput) Headers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HttpRequest) map[string]string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringMapOutput)
}

func (o HttpRequestPtrOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpRequest) *string {
		if v == nil {
			return nil
		}
		return v.Method
	}).(pulumi.StringPtrOutput)
}

func (o HttpRequestPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpRequest) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type HttpRequestResponse struct {
	Authentication *HttpAuthenticationResponse `pulumi:"authentication"`
	Body           *string                     `pulumi:"body"`
	Headers        map[string]string           `pulumi:"headers"`
	Method         *string                     `pulumi:"method"`
	Uri            *string                     `pulumi:"uri"`
}





type HttpRequestResponseInput interface {
	pulumi.Input

	ToHttpRequestResponseOutput() HttpRequestResponseOutput
	ToHttpRequestResponseOutputWithContext(context.Context) HttpRequestResponseOutput
}

type HttpRequestResponseArgs struct {
	Authentication HttpAuthenticationResponsePtrInput `pulumi:"authentication"`
	Body           pulumi.StringPtrInput              `pulumi:"body"`
	Headers        pulumi.StringMapInput              `pulumi:"headers"`
	Method         pulumi.StringPtrInput              `pulumi:"method"`
	Uri            pulumi.StringPtrInput              `pulumi:"uri"`
}

func (HttpRequestResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRequestResponse)(nil)).Elem()
}

func (i HttpRequestResponseArgs) ToHttpRequestResponseOutput() HttpRequestResponseOutput {
	return i.ToHttpRequestResponseOutputWithContext(context.Background())
}

func (i HttpRequestResponseArgs) ToHttpRequestResponseOutputWithContext(ctx context.Context) HttpRequestResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRequestResponseOutput)
}

func (i HttpRequestResponseArgs) ToHttpRequestResponsePtrOutput() HttpRequestResponsePtrOutput {
	return i.ToHttpRequestResponsePtrOutputWithContext(context.Background())
}

func (i HttpRequestResponseArgs) ToHttpRequestResponsePtrOutputWithContext(ctx context.Context) HttpRequestResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRequestResponseOutput).ToHttpRequestResponsePtrOutputWithContext(ctx)
}









type HttpRequestResponsePtrInput interface {
	pulumi.Input

	ToHttpRequestResponsePtrOutput() HttpRequestResponsePtrOutput
	ToHttpRequestResponsePtrOutputWithContext(context.Context) HttpRequestResponsePtrOutput
}

type httpRequestResponsePtrType HttpRequestResponseArgs

func HttpRequestResponsePtr(v *HttpRequestResponseArgs) HttpRequestResponsePtrInput {
	return (*httpRequestResponsePtrType)(v)
}

func (*httpRequestResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpRequestResponse)(nil)).Elem()
}

func (i *httpRequestResponsePtrType) ToHttpRequestResponsePtrOutput() HttpRequestResponsePtrOutput {
	return i.ToHttpRequestResponsePtrOutputWithContext(context.Background())
}

func (i *httpRequestResponsePtrType) ToHttpRequestResponsePtrOutputWithContext(ctx context.Context) HttpRequestResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRequestResponsePtrOutput)
}

type HttpRequestResponseOutput struct{ *pulumi.OutputState }

func (HttpRequestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRequestResponse)(nil)).Elem()
}

func (o HttpRequestResponseOutput) ToHttpRequestResponseOutput() HttpRequestResponseOutput {
	return o
}

func (o HttpRequestResponseOutput) ToHttpRequestResponseOutputWithContext(ctx context.Context) HttpRequestResponseOutput {
	return o
}

func (o HttpRequestResponseOutput) ToHttpRequestResponsePtrOutput() HttpRequestResponsePtrOutput {
	return o.ToHttpRequestResponsePtrOutputWithContext(context.Background())
}

func (o HttpRequestResponseOutput) ToHttpRequestResponsePtrOutputWithContext(ctx context.Context) HttpRequestResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpRequestResponse) *HttpRequestResponse {
		return &v
	}).(HttpRequestResponsePtrOutput)
}

func (o HttpRequestResponseOutput) Authentication() HttpAuthenticationResponsePtrOutput {
	return o.ApplyT(func(v HttpRequestResponse) *HttpAuthenticationResponse { return v.Authentication }).(HttpAuthenticationResponsePtrOutput)
}

func (o HttpRequestResponseOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRequestResponse) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o HttpRequestResponseOutput) Headers() pulumi.StringMapOutput {
	return o.ApplyT(func(v HttpRequestResponse) map[string]string { return v.Headers }).(pulumi.StringMapOutput)
}

func (o HttpRequestResponseOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRequestResponse) *string { return v.Method }).(pulumi.StringPtrOutput)
}

func (o HttpRequestResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRequestResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type HttpRequestResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpRequestResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpRequestResponse)(nil)).Elem()
}

func (o HttpRequestResponsePtrOutput) ToHttpRequestResponsePtrOutput() HttpRequestResponsePtrOutput {
	return o
}

func (o HttpRequestResponsePtrOutput) ToHttpRequestResponsePtrOutputWithContext(ctx context.Context) HttpRequestResponsePtrOutput {
	return o
}

func (o HttpRequestResponsePtrOutput) Elem() HttpRequestResponseOutput {
	return o.ApplyT(func(v *HttpRequestResponse) HttpRequestResponse {
		if v != nil {
			return *v
		}
		var ret HttpRequestResponse
		return ret
	}).(HttpRequestResponseOutput)
}

func (o HttpRequestResponsePtrOutput) Authentication() HttpAuthenticationResponsePtrOutput {
	return o.ApplyT(func(v *HttpRequestResponse) *HttpAuthenticationResponse {
		if v == nil {
			return nil
		}
		return v.Authentication
	}).(HttpAuthenticationResponsePtrOutput)
}

func (o HttpRequestResponsePtrOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpRequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.Body
	}).(pulumi.StringPtrOutput)
}

func (o HttpRequestResponsePtrOutput) Headers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HttpRequestResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringMapOutput)
}

func (o HttpRequestResponsePtrOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpRequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.Method
	}).(pulumi.StringPtrOutput)
}

func (o HttpRequestResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpRequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type JobAction struct {
	ErrorAction            *JobErrorAction         `pulumi:"errorAction"`
	QueueMessage           *StorageQueueMessage    `pulumi:"queueMessage"`
	Request                *HttpRequest            `pulumi:"request"`
	RetryPolicy            *RetryPolicy            `pulumi:"retryPolicy"`
	ServiceBusQueueMessage *ServiceBusQueueMessage `pulumi:"serviceBusQueueMessage"`
	ServiceBusTopicMessage *ServiceBusTopicMessage `pulumi:"serviceBusTopicMessage"`
	Type                   *JobActionType          `pulumi:"type"`
}





type JobActionInput interface {
	pulumi.Input

	ToJobActionOutput() JobActionOutput
	ToJobActionOutputWithContext(context.Context) JobActionOutput
}

type JobActionArgs struct {
	ErrorAction            JobErrorActionPtrInput         `pulumi:"errorAction"`
	QueueMessage           StorageQueueMessagePtrInput    `pulumi:"queueMessage"`
	Request                HttpRequestPtrInput            `pulumi:"request"`
	RetryPolicy            RetryPolicyPtrInput            `pulumi:"retryPolicy"`
	ServiceBusQueueMessage ServiceBusQueueMessagePtrInput `pulumi:"serviceBusQueueMessage"`
	ServiceBusTopicMessage ServiceBusTopicMessagePtrInput `pulumi:"serviceBusTopicMessage"`
	Type                   JobActionTypePtrInput          `pulumi:"type"`
}

func (JobActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobAction)(nil)).Elem()
}

func (i JobActionArgs) ToJobActionOutput() JobActionOutput {
	return i.ToJobActionOutputWithContext(context.Background())
}

func (i JobActionArgs) ToJobActionOutputWithContext(ctx context.Context) JobActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobActionOutput)
}

func (i JobActionArgs) ToJobActionPtrOutput() JobActionPtrOutput {
	return i.ToJobActionPtrOutputWithContext(context.Background())
}

func (i JobActionArgs) ToJobActionPtrOutputWithContext(ctx context.Context) JobActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobActionOutput).ToJobActionPtrOutputWithContext(ctx)
}









type JobActionPtrInput interface {
	pulumi.Input

	ToJobActionPtrOutput() JobActionPtrOutput
	ToJobActionPtrOutputWithContext(context.Context) JobActionPtrOutput
}

type jobActionPtrType JobActionArgs

func JobActionPtr(v *JobActionArgs) JobActionPtrInput {
	return (*jobActionPtrType)(v)
}

func (*jobActionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobAction)(nil)).Elem()
}

func (i *jobActionPtrType) ToJobActionPtrOutput() JobActionPtrOutput {
	return i.ToJobActionPtrOutputWithContext(context.Background())
}

func (i *jobActionPtrType) ToJobActionPtrOutputWithContext(ctx context.Context) JobActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobActionPtrOutput)
}

type JobActionOutput struct{ *pulumi.OutputState }

func (JobActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobAction)(nil)).Elem()
}

func (o JobActionOutput) ToJobActionOutput() JobActionOutput {
	return o
}

func (o JobActionOutput) ToJobActionOutputWithContext(ctx context.Context) JobActionOutput {
	return o
}

func (o JobActionOutput) ToJobActionPtrOutput() JobActionPtrOutput {
	return o.ToJobActionPtrOutputWithContext(context.Background())
}

func (o JobActionOutput) ToJobActionPtrOutputWithContext(ctx context.Context) JobActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobAction) *JobAction {
		return &v
	}).(JobActionPtrOutput)
}

func (o JobActionOutput) ErrorAction() JobErrorActionPtrOutput {
	return o.ApplyT(func(v JobAction) *JobErrorAction { return v.ErrorAction }).(JobErrorActionPtrOutput)
}

func (o JobActionOutput) QueueMessage() StorageQueueMessagePtrOutput {
	return o.ApplyT(func(v JobAction) *StorageQueueMessage { return v.QueueMessage }).(StorageQueueMessagePtrOutput)
}

func (o JobActionOutput) Request() HttpRequestPtrOutput {
	return o.ApplyT(func(v JobAction) *HttpRequest { return v.Request }).(HttpRequestPtrOutput)
}

func (o JobActionOutput) RetryPolicy() RetryPolicyPtrOutput {
	return o.ApplyT(func(v JobAction) *RetryPolicy { return v.RetryPolicy }).(RetryPolicyPtrOutput)
}

func (o JobActionOutput) ServiceBusQueueMessage() ServiceBusQueueMessagePtrOutput {
	return o.ApplyT(func(v JobAction) *ServiceBusQueueMessage { return v.ServiceBusQueueMessage }).(ServiceBusQueueMessagePtrOutput)
}

func (o JobActionOutput) ServiceBusTopicMessage() ServiceBusTopicMessagePtrOutput {
	return o.ApplyT(func(v JobAction) *ServiceBusTopicMessage { return v.ServiceBusTopicMessage }).(ServiceBusTopicMessagePtrOutput)
}

func (o JobActionOutput) Type() JobActionTypePtrOutput {
	return o.ApplyT(func(v JobAction) *JobActionType { return v.Type }).(JobActionTypePtrOutput)
}

type JobActionPtrOutput struct{ *pulumi.OutputState }

func (JobActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobAction)(nil)).Elem()
}

func (o JobActionPtrOutput) ToJobActionPtrOutput() JobActionPtrOutput {
	return o
}

func (o JobActionPtrOutput) ToJobActionPtrOutputWithContext(ctx context.Context) JobActionPtrOutput {
	return o
}

func (o JobActionPtrOutput) Elem() JobActionOutput {
	return o.ApplyT(func(v *JobAction) JobAction {
		if v != nil {
			return *v
		}
		var ret JobAction
		return ret
	}).(JobActionOutput)
}

func (o JobActionPtrOutput) ErrorAction() JobErrorActionPtrOutput {
	return o.ApplyT(func(v *JobAction) *JobErrorAction {
		if v == nil {
			return nil
		}
		return v.ErrorAction
	}).(JobErrorActionPtrOutput)
}

func (o JobActionPtrOutput) QueueMessage() StorageQueueMessagePtrOutput {
	return o.ApplyT(func(v *JobAction) *StorageQueueMessage {
		if v == nil {
			return nil
		}
		return v.QueueMessage
	}).(StorageQueueMessagePtrOutput)
}

func (o JobActionPtrOutput) Request() HttpRequestPtrOutput {
	return o.ApplyT(func(v *JobAction) *HttpRequest {
		if v == nil {
			return nil
		}
		return v.Request
	}).(HttpRequestPtrOutput)
}

func (o JobActionPtrOutput) RetryPolicy() RetryPolicyPtrOutput {
	return o.ApplyT(func(v *JobAction) *RetryPolicy {
		if v == nil {
			return nil
		}
		return v.RetryPolicy
	}).(RetryPolicyPtrOutput)
}

func (o JobActionPtrOutput) ServiceBusQueueMessage() ServiceBusQueueMessagePtrOutput {
	return o.ApplyT(func(v *JobAction) *ServiceBusQueueMessage {
		if v == nil {
			return nil
		}
		return v.ServiceBusQueueMessage
	}).(ServiceBusQueueMessagePtrOutput)
}

func (o JobActionPtrOutput) ServiceBusTopicMessage() ServiceBusTopicMessagePtrOutput {
	return o.ApplyT(func(v *JobAction) *ServiceBusTopicMessage {
		if v == nil {
			return nil
		}
		return v.ServiceBusTopicMessage
	}).(ServiceBusTopicMessagePtrOutput)
}

func (o JobActionPtrOutput) Type() JobActionTypePtrOutput {
	return o.ApplyT(func(v *JobAction) *JobActionType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(JobActionTypePtrOutput)
}

type JobActionResponse struct {
	ErrorAction            *JobErrorActionResponse         `pulumi:"errorAction"`
	QueueMessage           *StorageQueueMessageResponse    `pulumi:"queueMessage"`
	Request                *HttpRequestResponse            `pulumi:"request"`
	RetryPolicy            *RetryPolicyResponse            `pulumi:"retryPolicy"`
	ServiceBusQueueMessage *ServiceBusQueueMessageResponse `pulumi:"serviceBusQueueMessage"`
	ServiceBusTopicMessage *ServiceBusTopicMessageResponse `pulumi:"serviceBusTopicMessage"`
	Type                   *string                         `pulumi:"type"`
}





type JobActionResponseInput interface {
	pulumi.Input

	ToJobActionResponseOutput() JobActionResponseOutput
	ToJobActionResponseOutputWithContext(context.Context) JobActionResponseOutput
}

type JobActionResponseArgs struct {
	ErrorAction            JobErrorActionResponsePtrInput         `pulumi:"errorAction"`
	QueueMessage           StorageQueueMessageResponsePtrInput    `pulumi:"queueMessage"`
	Request                HttpRequestResponsePtrInput            `pulumi:"request"`
	RetryPolicy            RetryPolicyResponsePtrInput            `pulumi:"retryPolicy"`
	ServiceBusQueueMessage ServiceBusQueueMessageResponsePtrInput `pulumi:"serviceBusQueueMessage"`
	ServiceBusTopicMessage ServiceBusTopicMessageResponsePtrInput `pulumi:"serviceBusTopicMessage"`
	Type                   pulumi.StringPtrInput                  `pulumi:"type"`
}

func (JobActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobActionResponse)(nil)).Elem()
}

func (i JobActionResponseArgs) ToJobActionResponseOutput() JobActionResponseOutput {
	return i.ToJobActionResponseOutputWithContext(context.Background())
}

func (i JobActionResponseArgs) ToJobActionResponseOutputWithContext(ctx context.Context) JobActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobActionResponseOutput)
}

func (i JobActionResponseArgs) ToJobActionResponsePtrOutput() JobActionResponsePtrOutput {
	return i.ToJobActionResponsePtrOutputWithContext(context.Background())
}

func (i JobActionResponseArgs) ToJobActionResponsePtrOutputWithContext(ctx context.Context) JobActionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobActionResponseOutput).ToJobActionResponsePtrOutputWithContext(ctx)
}









type JobActionResponsePtrInput interface {
	pulumi.Input

	ToJobActionResponsePtrOutput() JobActionResponsePtrOutput
	ToJobActionResponsePtrOutputWithContext(context.Context) JobActionResponsePtrOutput
}

type jobActionResponsePtrType JobActionResponseArgs

func JobActionResponsePtr(v *JobActionResponseArgs) JobActionResponsePtrInput {
	return (*jobActionResponsePtrType)(v)
}

func (*jobActionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobActionResponse)(nil)).Elem()
}

func (i *jobActionResponsePtrType) ToJobActionResponsePtrOutput() JobActionResponsePtrOutput {
	return i.ToJobActionResponsePtrOutputWithContext(context.Background())
}

func (i *jobActionResponsePtrType) ToJobActionResponsePtrOutputWithContext(ctx context.Context) JobActionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobActionResponsePtrOutput)
}

type JobActionResponseOutput struct{ *pulumi.OutputState }

func (JobActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobActionResponse)(nil)).Elem()
}

func (o JobActionResponseOutput) ToJobActionResponseOutput() JobActionResponseOutput {
	return o
}

func (o JobActionResponseOutput) ToJobActionResponseOutputWithContext(ctx context.Context) JobActionResponseOutput {
	return o
}

func (o JobActionResponseOutput) ToJobActionResponsePtrOutput() JobActionResponsePtrOutput {
	return o.ToJobActionResponsePtrOutputWithContext(context.Background())
}

func (o JobActionResponseOutput) ToJobActionResponsePtrOutputWithContext(ctx context.Context) JobActionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobActionResponse) *JobActionResponse {
		return &v
	}).(JobActionResponsePtrOutput)
}

func (o JobActionResponseOutput) ErrorAction() JobErrorActionResponsePtrOutput {
	return o.ApplyT(func(v JobActionResponse) *JobErrorActionResponse { return v.ErrorAction }).(JobErrorActionResponsePtrOutput)
}

func (o JobActionResponseOutput) QueueMessage() StorageQueueMessageResponsePtrOutput {
	return o.ApplyT(func(v JobActionResponse) *StorageQueueMessageResponse { return v.QueueMessage }).(StorageQueueMessageResponsePtrOutput)
}

func (o JobActionResponseOutput) Request() HttpRequestResponsePtrOutput {
	return o.ApplyT(func(v JobActionResponse) *HttpRequestResponse { return v.Request }).(HttpRequestResponsePtrOutput)
}

func (o JobActionResponseOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v JobActionResponse) *RetryPolicyResponse { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

func (o JobActionResponseOutput) ServiceBusQueueMessage() ServiceBusQueueMessageResponsePtrOutput {
	return o.ApplyT(func(v JobActionResponse) *ServiceBusQueueMessageResponse { return v.ServiceBusQueueMessage }).(ServiceBusQueueMessageResponsePtrOutput)
}

func (o JobActionResponseOutput) ServiceBusTopicMessage() ServiceBusTopicMessageResponsePtrOutput {
	return o.ApplyT(func(v JobActionResponse) *ServiceBusTopicMessageResponse { return v.ServiceBusTopicMessage }).(ServiceBusTopicMessageResponsePtrOutput)
}

func (o JobActionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobActionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type JobActionResponsePtrOutput struct{ *pulumi.OutputState }

func (JobActionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobActionResponse)(nil)).Elem()
}

func (o JobActionResponsePtrOutput) ToJobActionResponsePtrOutput() JobActionResponsePtrOutput {
	return o
}

func (o JobActionResponsePtrOutput) ToJobActionResponsePtrOutputWithContext(ctx context.Context) JobActionResponsePtrOutput {
	return o
}

func (o JobActionResponsePtrOutput) Elem() JobActionResponseOutput {
	return o.ApplyT(func(v *JobActionResponse) JobActionResponse {
		if v != nil {
			return *v
		}
		var ret JobActionResponse
		return ret
	}).(JobActionResponseOutput)
}

func (o JobActionResponsePtrOutput) ErrorAction() JobErrorActionResponsePtrOutput {
	return o.ApplyT(func(v *JobActionResponse) *JobErrorActionResponse {
		if v == nil {
			return nil
		}
		return v.ErrorAction
	}).(JobErrorActionResponsePtrOutput)
}

func (o JobActionResponsePtrOutput) QueueMessage() StorageQueueMessageResponsePtrOutput {
	return o.ApplyT(func(v *JobActionResponse) *StorageQueueMessageResponse {
		if v == nil {
			return nil
		}
		return v.QueueMessage
	}).(StorageQueueMessageResponsePtrOutput)
}

func (o JobActionResponsePtrOutput) Request() HttpRequestResponsePtrOutput {
	return o.ApplyT(func(v *JobActionResponse) *HttpRequestResponse {
		if v == nil {
			return nil
		}
		return v.Request
	}).(HttpRequestResponsePtrOutput)
}

func (o JobActionResponsePtrOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v *JobActionResponse) *RetryPolicyResponse {
		if v == nil {
			return nil
		}
		return v.RetryPolicy
	}).(RetryPolicyResponsePtrOutput)
}

func (o JobActionResponsePtrOutput) ServiceBusQueueMessage() ServiceBusQueueMessageResponsePtrOutput {
	return o.ApplyT(func(v *JobActionResponse) *ServiceBusQueueMessageResponse {
		if v == nil {
			return nil
		}
		return v.ServiceBusQueueMessage
	}).(ServiceBusQueueMessageResponsePtrOutput)
}

func (o JobActionResponsePtrOutput) ServiceBusTopicMessage() ServiceBusTopicMessageResponsePtrOutput {
	return o.ApplyT(func(v *JobActionResponse) *ServiceBusTopicMessageResponse {
		if v == nil {
			return nil
		}
		return v.ServiceBusTopicMessage
	}).(ServiceBusTopicMessageResponsePtrOutput)
}

func (o JobActionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type JobCollectionProperties struct {
	Quota *JobCollectionQuota     `pulumi:"quota"`
	Sku   *Sku                    `pulumi:"sku"`
	State *JobCollectionStateEnum `pulumi:"state"`
}





type JobCollectionPropertiesInput interface {
	pulumi.Input

	ToJobCollectionPropertiesOutput() JobCollectionPropertiesOutput
	ToJobCollectionPropertiesOutputWithContext(context.Context) JobCollectionPropertiesOutput
}

type JobCollectionPropertiesArgs struct {
	Quota JobCollectionQuotaPtrInput     `pulumi:"quota"`
	Sku   SkuPtrInput                    `pulumi:"sku"`
	State JobCollectionStateEnumPtrInput `pulumi:"state"`
}

func (JobCollectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobCollectionProperties)(nil)).Elem()
}

func (i JobCollectionPropertiesArgs) ToJobCollectionPropertiesOutput() JobCollectionPropertiesOutput {
	return i.ToJobCollectionPropertiesOutputWithContext(context.Background())
}

func (i JobCollectionPropertiesArgs) ToJobCollectionPropertiesOutputWithContext(ctx context.Context) JobCollectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionPropertiesOutput)
}

func (i JobCollectionPropertiesArgs) ToJobCollectionPropertiesPtrOutput() JobCollectionPropertiesPtrOutput {
	return i.ToJobCollectionPropertiesPtrOutputWithContext(context.Background())
}

func (i JobCollectionPropertiesArgs) ToJobCollectionPropertiesPtrOutputWithContext(ctx context.Context) JobCollectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionPropertiesOutput).ToJobCollectionPropertiesPtrOutputWithContext(ctx)
}









type JobCollectionPropertiesPtrInput interface {
	pulumi.Input

	ToJobCollectionPropertiesPtrOutput() JobCollectionPropertiesPtrOutput
	ToJobCollectionPropertiesPtrOutputWithContext(context.Context) JobCollectionPropertiesPtrOutput
}

type jobCollectionPropertiesPtrType JobCollectionPropertiesArgs

func JobCollectionPropertiesPtr(v *JobCollectionPropertiesArgs) JobCollectionPropertiesPtrInput {
	return (*jobCollectionPropertiesPtrType)(v)
}

func (*jobCollectionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCollectionProperties)(nil)).Elem()
}

func (i *jobCollectionPropertiesPtrType) ToJobCollectionPropertiesPtrOutput() JobCollectionPropertiesPtrOutput {
	return i.ToJobCollectionPropertiesPtrOutputWithContext(context.Background())
}

func (i *jobCollectionPropertiesPtrType) ToJobCollectionPropertiesPtrOutputWithContext(ctx context.Context) JobCollectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionPropertiesPtrOutput)
}

type JobCollectionPropertiesOutput struct{ *pulumi.OutputState }

func (JobCollectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobCollectionProperties)(nil)).Elem()
}

func (o JobCollectionPropertiesOutput) ToJobCollectionPropertiesOutput() JobCollectionPropertiesOutput {
	return o
}

func (o JobCollectionPropertiesOutput) ToJobCollectionPropertiesOutputWithContext(ctx context.Context) JobCollectionPropertiesOutput {
	return o
}

func (o JobCollectionPropertiesOutput) ToJobCollectionPropertiesPtrOutput() JobCollectionPropertiesPtrOutput {
	return o.ToJobCollectionPropertiesPtrOutputWithContext(context.Background())
}

func (o JobCollectionPropertiesOutput) ToJobCollectionPropertiesPtrOutputWithContext(ctx context.Context) JobCollectionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobCollectionProperties) *JobCollectionProperties {
		return &v
	}).(JobCollectionPropertiesPtrOutput)
}

func (o JobCollectionPropertiesOutput) Quota() JobCollectionQuotaPtrOutput {
	return o.ApplyT(func(v JobCollectionProperties) *JobCollectionQuota { return v.Quota }).(JobCollectionQuotaPtrOutput)
}

func (o JobCollectionPropertiesOutput) Sku() SkuPtrOutput {
	return o.ApplyT(func(v JobCollectionProperties) *Sku { return v.Sku }).(SkuPtrOutput)
}

func (o JobCollectionPropertiesOutput) State() JobCollectionStateEnumPtrOutput {
	return o.ApplyT(func(v JobCollectionProperties) *JobCollectionStateEnum { return v.State }).(JobCollectionStateEnumPtrOutput)
}

type JobCollectionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (JobCollectionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCollectionProperties)(nil)).Elem()
}

func (o JobCollectionPropertiesPtrOutput) ToJobCollectionPropertiesPtrOutput() JobCollectionPropertiesPtrOutput {
	return o
}

func (o JobCollectionPropertiesPtrOutput) ToJobCollectionPropertiesPtrOutputWithContext(ctx context.Context) JobCollectionPropertiesPtrOutput {
	return o
}

func (o JobCollectionPropertiesPtrOutput) Elem() JobCollectionPropertiesOutput {
	return o.ApplyT(func(v *JobCollectionProperties) JobCollectionProperties {
		if v != nil {
			return *v
		}
		var ret JobCollectionProperties
		return ret
	}).(JobCollectionPropertiesOutput)
}

func (o JobCollectionPropertiesPtrOutput) Quota() JobCollectionQuotaPtrOutput {
	return o.ApplyT(func(v *JobCollectionProperties) *JobCollectionQuota {
		if v == nil {
			return nil
		}
		return v.Quota
	}).(JobCollectionQuotaPtrOutput)
}

func (o JobCollectionPropertiesPtrOutput) Sku() SkuPtrOutput {
	return o.ApplyT(func(v *JobCollectionProperties) *Sku {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(SkuPtrOutput)
}

func (o JobCollectionPropertiesPtrOutput) State() JobCollectionStateEnumPtrOutput {
	return o.ApplyT(func(v *JobCollectionProperties) *JobCollectionStateEnum {
		if v == nil {
			return nil
		}
		return v.State
	}).(JobCollectionStateEnumPtrOutput)
}

type JobCollectionPropertiesResponse struct {
	Quota *JobCollectionQuotaResponse `pulumi:"quota"`
	Sku   *SkuResponse                `pulumi:"sku"`
	State *string                     `pulumi:"state"`
}





type JobCollectionPropertiesResponseInput interface {
	pulumi.Input

	ToJobCollectionPropertiesResponseOutput() JobCollectionPropertiesResponseOutput
	ToJobCollectionPropertiesResponseOutputWithContext(context.Context) JobCollectionPropertiesResponseOutput
}

type JobCollectionPropertiesResponseArgs struct {
	Quota JobCollectionQuotaResponsePtrInput `pulumi:"quota"`
	Sku   SkuResponsePtrInput                `pulumi:"sku"`
	State pulumi.StringPtrInput              `pulumi:"state"`
}

func (JobCollectionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobCollectionPropertiesResponse)(nil)).Elem()
}

func (i JobCollectionPropertiesResponseArgs) ToJobCollectionPropertiesResponseOutput() JobCollectionPropertiesResponseOutput {
	return i.ToJobCollectionPropertiesResponseOutputWithContext(context.Background())
}

func (i JobCollectionPropertiesResponseArgs) ToJobCollectionPropertiesResponseOutputWithContext(ctx context.Context) JobCollectionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionPropertiesResponseOutput)
}

func (i JobCollectionPropertiesResponseArgs) ToJobCollectionPropertiesResponsePtrOutput() JobCollectionPropertiesResponsePtrOutput {
	return i.ToJobCollectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i JobCollectionPropertiesResponseArgs) ToJobCollectionPropertiesResponsePtrOutputWithContext(ctx context.Context) JobCollectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionPropertiesResponseOutput).ToJobCollectionPropertiesResponsePtrOutputWithContext(ctx)
}









type JobCollectionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToJobCollectionPropertiesResponsePtrOutput() JobCollectionPropertiesResponsePtrOutput
	ToJobCollectionPropertiesResponsePtrOutputWithContext(context.Context) JobCollectionPropertiesResponsePtrOutput
}

type jobCollectionPropertiesResponsePtrType JobCollectionPropertiesResponseArgs

func JobCollectionPropertiesResponsePtr(v *JobCollectionPropertiesResponseArgs) JobCollectionPropertiesResponsePtrInput {
	return (*jobCollectionPropertiesResponsePtrType)(v)
}

func (*jobCollectionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCollectionPropertiesResponse)(nil)).Elem()
}

func (i *jobCollectionPropertiesResponsePtrType) ToJobCollectionPropertiesResponsePtrOutput() JobCollectionPropertiesResponsePtrOutput {
	return i.ToJobCollectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *jobCollectionPropertiesResponsePtrType) ToJobCollectionPropertiesResponsePtrOutputWithContext(ctx context.Context) JobCollectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionPropertiesResponsePtrOutput)
}

type JobCollectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (JobCollectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobCollectionPropertiesResponse)(nil)).Elem()
}

func (o JobCollectionPropertiesResponseOutput) ToJobCollectionPropertiesResponseOutput() JobCollectionPropertiesResponseOutput {
	return o
}

func (o JobCollectionPropertiesResponseOutput) ToJobCollectionPropertiesResponseOutputWithContext(ctx context.Context) JobCollectionPropertiesResponseOutput {
	return o
}

func (o JobCollectionPropertiesResponseOutput) ToJobCollectionPropertiesResponsePtrOutput() JobCollectionPropertiesResponsePtrOutput {
	return o.ToJobCollectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o JobCollectionPropertiesResponseOutput) ToJobCollectionPropertiesResponsePtrOutputWithContext(ctx context.Context) JobCollectionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobCollectionPropertiesResponse) *JobCollectionPropertiesResponse {
		return &v
	}).(JobCollectionPropertiesResponsePtrOutput)
}

func (o JobCollectionPropertiesResponseOutput) Quota() JobCollectionQuotaResponsePtrOutput {
	return o.ApplyT(func(v JobCollectionPropertiesResponse) *JobCollectionQuotaResponse { return v.Quota }).(JobCollectionQuotaResponsePtrOutput)
}

func (o JobCollectionPropertiesResponseOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v JobCollectionPropertiesResponse) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o JobCollectionPropertiesResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobCollectionPropertiesResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type JobCollectionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (JobCollectionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCollectionPropertiesResponse)(nil)).Elem()
}

func (o JobCollectionPropertiesResponsePtrOutput) ToJobCollectionPropertiesResponsePtrOutput() JobCollectionPropertiesResponsePtrOutput {
	return o
}

func (o JobCollectionPropertiesResponsePtrOutput) ToJobCollectionPropertiesResponsePtrOutputWithContext(ctx context.Context) JobCollectionPropertiesResponsePtrOutput {
	return o
}

func (o JobCollectionPropertiesResponsePtrOutput) Elem() JobCollectionPropertiesResponseOutput {
	return o.ApplyT(func(v *JobCollectionPropertiesResponse) JobCollectionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret JobCollectionPropertiesResponse
		return ret
	}).(JobCollectionPropertiesResponseOutput)
}

func (o JobCollectionPropertiesResponsePtrOutput) Quota() JobCollectionQuotaResponsePtrOutput {
	return o.ApplyT(func(v *JobCollectionPropertiesResponse) *JobCollectionQuotaResponse {
		if v == nil {
			return nil
		}
		return v.Quota
	}).(JobCollectionQuotaResponsePtrOutput)
}

func (o JobCollectionPropertiesResponsePtrOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *JobCollectionPropertiesResponse) *SkuResponse {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(SkuResponsePtrOutput)
}

func (o JobCollectionPropertiesResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobCollectionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type JobCollectionQuota struct {
	MaxJobCount      *int              `pulumi:"maxJobCount"`
	MaxJobOccurrence *int              `pulumi:"maxJobOccurrence"`
	MaxRecurrence    *JobMaxRecurrence `pulumi:"maxRecurrence"`
}





type JobCollectionQuotaInput interface {
	pulumi.Input

	ToJobCollectionQuotaOutput() JobCollectionQuotaOutput
	ToJobCollectionQuotaOutputWithContext(context.Context) JobCollectionQuotaOutput
}

type JobCollectionQuotaArgs struct {
	MaxJobCount      pulumi.IntPtrInput       `pulumi:"maxJobCount"`
	MaxJobOccurrence pulumi.IntPtrInput       `pulumi:"maxJobOccurrence"`
	MaxRecurrence    JobMaxRecurrencePtrInput `pulumi:"maxRecurrence"`
}

func (JobCollectionQuotaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobCollectionQuota)(nil)).Elem()
}

func (i JobCollectionQuotaArgs) ToJobCollectionQuotaOutput() JobCollectionQuotaOutput {
	return i.ToJobCollectionQuotaOutputWithContext(context.Background())
}

func (i JobCollectionQuotaArgs) ToJobCollectionQuotaOutputWithContext(ctx context.Context) JobCollectionQuotaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionQuotaOutput)
}

func (i JobCollectionQuotaArgs) ToJobCollectionQuotaPtrOutput() JobCollectionQuotaPtrOutput {
	return i.ToJobCollectionQuotaPtrOutputWithContext(context.Background())
}

func (i JobCollectionQuotaArgs) ToJobCollectionQuotaPtrOutputWithContext(ctx context.Context) JobCollectionQuotaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionQuotaOutput).ToJobCollectionQuotaPtrOutputWithContext(ctx)
}









type JobCollectionQuotaPtrInput interface {
	pulumi.Input

	ToJobCollectionQuotaPtrOutput() JobCollectionQuotaPtrOutput
	ToJobCollectionQuotaPtrOutputWithContext(context.Context) JobCollectionQuotaPtrOutput
}

type jobCollectionQuotaPtrType JobCollectionQuotaArgs

func JobCollectionQuotaPtr(v *JobCollectionQuotaArgs) JobCollectionQuotaPtrInput {
	return (*jobCollectionQuotaPtrType)(v)
}

func (*jobCollectionQuotaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCollectionQuota)(nil)).Elem()
}

func (i *jobCollectionQuotaPtrType) ToJobCollectionQuotaPtrOutput() JobCollectionQuotaPtrOutput {
	return i.ToJobCollectionQuotaPtrOutputWithContext(context.Background())
}

func (i *jobCollectionQuotaPtrType) ToJobCollectionQuotaPtrOutputWithContext(ctx context.Context) JobCollectionQuotaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionQuotaPtrOutput)
}

type JobCollectionQuotaOutput struct{ *pulumi.OutputState }

func (JobCollectionQuotaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobCollectionQuota)(nil)).Elem()
}

func (o JobCollectionQuotaOutput) ToJobCollectionQuotaOutput() JobCollectionQuotaOutput {
	return o
}

func (o JobCollectionQuotaOutput) ToJobCollectionQuotaOutputWithContext(ctx context.Context) JobCollectionQuotaOutput {
	return o
}

func (o JobCollectionQuotaOutput) ToJobCollectionQuotaPtrOutput() JobCollectionQuotaPtrOutput {
	return o.ToJobCollectionQuotaPtrOutputWithContext(context.Background())
}

func (o JobCollectionQuotaOutput) ToJobCollectionQuotaPtrOutputWithContext(ctx context.Context) JobCollectionQuotaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobCollectionQuota) *JobCollectionQuota {
		return &v
	}).(JobCollectionQuotaPtrOutput)
}

func (o JobCollectionQuotaOutput) MaxJobCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobCollectionQuota) *int { return v.MaxJobCount }).(pulumi.IntPtrOutput)
}

func (o JobCollectionQuotaOutput) MaxJobOccurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobCollectionQuota) *int { return v.MaxJobOccurrence }).(pulumi.IntPtrOutput)
}

func (o JobCollectionQuotaOutput) MaxRecurrence() JobMaxRecurrencePtrOutput {
	return o.ApplyT(func(v JobCollectionQuota) *JobMaxRecurrence { return v.MaxRecurrence }).(JobMaxRecurrencePtrOutput)
}

type JobCollectionQuotaPtrOutput struct{ *pulumi.OutputState }

func (JobCollectionQuotaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCollectionQuota)(nil)).Elem()
}

func (o JobCollectionQuotaPtrOutput) ToJobCollectionQuotaPtrOutput() JobCollectionQuotaPtrOutput {
	return o
}

func (o JobCollectionQuotaPtrOutput) ToJobCollectionQuotaPtrOutputWithContext(ctx context.Context) JobCollectionQuotaPtrOutput {
	return o
}

func (o JobCollectionQuotaPtrOutput) Elem() JobCollectionQuotaOutput {
	return o.ApplyT(func(v *JobCollectionQuota) JobCollectionQuota {
		if v != nil {
			return *v
		}
		var ret JobCollectionQuota
		return ret
	}).(JobCollectionQuotaOutput)
}

func (o JobCollectionQuotaPtrOutput) MaxJobCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobCollectionQuota) *int {
		if v == nil {
			return nil
		}
		return v.MaxJobCount
	}).(pulumi.IntPtrOutput)
}

func (o JobCollectionQuotaPtrOutput) MaxJobOccurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobCollectionQuota) *int {
		if v == nil {
			return nil
		}
		return v.MaxJobOccurrence
	}).(pulumi.IntPtrOutput)
}

func (o JobCollectionQuotaPtrOutput) MaxRecurrence() JobMaxRecurrencePtrOutput {
	return o.ApplyT(func(v *JobCollectionQuota) *JobMaxRecurrence {
		if v == nil {
			return nil
		}
		return v.MaxRecurrence
	}).(JobMaxRecurrencePtrOutput)
}

type JobCollectionQuotaResponse struct {
	MaxJobCount      *int                      `pulumi:"maxJobCount"`
	MaxJobOccurrence *int                      `pulumi:"maxJobOccurrence"`
	MaxRecurrence    *JobMaxRecurrenceResponse `pulumi:"maxRecurrence"`
}





type JobCollectionQuotaResponseInput interface {
	pulumi.Input

	ToJobCollectionQuotaResponseOutput() JobCollectionQuotaResponseOutput
	ToJobCollectionQuotaResponseOutputWithContext(context.Context) JobCollectionQuotaResponseOutput
}

type JobCollectionQuotaResponseArgs struct {
	MaxJobCount      pulumi.IntPtrInput               `pulumi:"maxJobCount"`
	MaxJobOccurrence pulumi.IntPtrInput               `pulumi:"maxJobOccurrence"`
	MaxRecurrence    JobMaxRecurrenceResponsePtrInput `pulumi:"maxRecurrence"`
}

func (JobCollectionQuotaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobCollectionQuotaResponse)(nil)).Elem()
}

func (i JobCollectionQuotaResponseArgs) ToJobCollectionQuotaResponseOutput() JobCollectionQuotaResponseOutput {
	return i.ToJobCollectionQuotaResponseOutputWithContext(context.Background())
}

func (i JobCollectionQuotaResponseArgs) ToJobCollectionQuotaResponseOutputWithContext(ctx context.Context) JobCollectionQuotaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionQuotaResponseOutput)
}

func (i JobCollectionQuotaResponseArgs) ToJobCollectionQuotaResponsePtrOutput() JobCollectionQuotaResponsePtrOutput {
	return i.ToJobCollectionQuotaResponsePtrOutputWithContext(context.Background())
}

func (i JobCollectionQuotaResponseArgs) ToJobCollectionQuotaResponsePtrOutputWithContext(ctx context.Context) JobCollectionQuotaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionQuotaResponseOutput).ToJobCollectionQuotaResponsePtrOutputWithContext(ctx)
}









type JobCollectionQuotaResponsePtrInput interface {
	pulumi.Input

	ToJobCollectionQuotaResponsePtrOutput() JobCollectionQuotaResponsePtrOutput
	ToJobCollectionQuotaResponsePtrOutputWithContext(context.Context) JobCollectionQuotaResponsePtrOutput
}

type jobCollectionQuotaResponsePtrType JobCollectionQuotaResponseArgs

func JobCollectionQuotaResponsePtr(v *JobCollectionQuotaResponseArgs) JobCollectionQuotaResponsePtrInput {
	return (*jobCollectionQuotaResponsePtrType)(v)
}

func (*jobCollectionQuotaResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCollectionQuotaResponse)(nil)).Elem()
}

func (i *jobCollectionQuotaResponsePtrType) ToJobCollectionQuotaResponsePtrOutput() JobCollectionQuotaResponsePtrOutput {
	return i.ToJobCollectionQuotaResponsePtrOutputWithContext(context.Background())
}

func (i *jobCollectionQuotaResponsePtrType) ToJobCollectionQuotaResponsePtrOutputWithContext(ctx context.Context) JobCollectionQuotaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionQuotaResponsePtrOutput)
}

type JobCollectionQuotaResponseOutput struct{ *pulumi.OutputState }

func (JobCollectionQuotaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobCollectionQuotaResponse)(nil)).Elem()
}

func (o JobCollectionQuotaResponseOutput) ToJobCollectionQuotaResponseOutput() JobCollectionQuotaResponseOutput {
	return o
}

func (o JobCollectionQuotaResponseOutput) ToJobCollectionQuotaResponseOutputWithContext(ctx context.Context) JobCollectionQuotaResponseOutput {
	return o
}

func (o JobCollectionQuotaResponseOutput) ToJobCollectionQuotaResponsePtrOutput() JobCollectionQuotaResponsePtrOutput {
	return o.ToJobCollectionQuotaResponsePtrOutputWithContext(context.Background())
}

func (o JobCollectionQuotaResponseOutput) ToJobCollectionQuotaResponsePtrOutputWithContext(ctx context.Context) JobCollectionQuotaResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobCollectionQuotaResponse) *JobCollectionQuotaResponse {
		return &v
	}).(JobCollectionQuotaResponsePtrOutput)
}

func (o JobCollectionQuotaResponseOutput) MaxJobCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobCollectionQuotaResponse) *int { return v.MaxJobCount }).(pulumi.IntPtrOutput)
}

func (o JobCollectionQuotaResponseOutput) MaxJobOccurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobCollectionQuotaResponse) *int { return v.MaxJobOccurrence }).(pulumi.IntPtrOutput)
}

func (o JobCollectionQuotaResponseOutput) MaxRecurrence() JobMaxRecurrenceResponsePtrOutput {
	return o.ApplyT(func(v JobCollectionQuotaResponse) *JobMaxRecurrenceResponse { return v.MaxRecurrence }).(JobMaxRecurrenceResponsePtrOutput)
}

type JobCollectionQuotaResponsePtrOutput struct{ *pulumi.OutputState }

func (JobCollectionQuotaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCollectionQuotaResponse)(nil)).Elem()
}

func (o JobCollectionQuotaResponsePtrOutput) ToJobCollectionQuotaResponsePtrOutput() JobCollectionQuotaResponsePtrOutput {
	return o
}

func (o JobCollectionQuotaResponsePtrOutput) ToJobCollectionQuotaResponsePtrOutputWithContext(ctx context.Context) JobCollectionQuotaResponsePtrOutput {
	return o
}

func (o JobCollectionQuotaResponsePtrOutput) Elem() JobCollectionQuotaResponseOutput {
	return o.ApplyT(func(v *JobCollectionQuotaResponse) JobCollectionQuotaResponse {
		if v != nil {
			return *v
		}
		var ret JobCollectionQuotaResponse
		return ret
	}).(JobCollectionQuotaResponseOutput)
}

func (o JobCollectionQuotaResponsePtrOutput) MaxJobCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobCollectionQuotaResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxJobCount
	}).(pulumi.IntPtrOutput)
}

func (o JobCollectionQuotaResponsePtrOutput) MaxJobOccurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobCollectionQuotaResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxJobOccurrence
	}).(pulumi.IntPtrOutput)
}

func (o JobCollectionQuotaResponsePtrOutput) MaxRecurrence() JobMaxRecurrenceResponsePtrOutput {
	return o.ApplyT(func(v *JobCollectionQuotaResponse) *JobMaxRecurrenceResponse {
		if v == nil {
			return nil
		}
		return v.MaxRecurrence
	}).(JobMaxRecurrenceResponsePtrOutput)
}

type JobErrorAction struct {
	QueueMessage           *StorageQueueMessage    `pulumi:"queueMessage"`
	Request                *HttpRequest            `pulumi:"request"`
	RetryPolicy            *RetryPolicy            `pulumi:"retryPolicy"`
	ServiceBusQueueMessage *ServiceBusQueueMessage `pulumi:"serviceBusQueueMessage"`
	ServiceBusTopicMessage *ServiceBusTopicMessage `pulumi:"serviceBusTopicMessage"`
	Type                   *JobActionType          `pulumi:"type"`
}





type JobErrorActionInput interface {
	pulumi.Input

	ToJobErrorActionOutput() JobErrorActionOutput
	ToJobErrorActionOutputWithContext(context.Context) JobErrorActionOutput
}

type JobErrorActionArgs struct {
	QueueMessage           StorageQueueMessagePtrInput    `pulumi:"queueMessage"`
	Request                HttpRequestPtrInput            `pulumi:"request"`
	RetryPolicy            RetryPolicyPtrInput            `pulumi:"retryPolicy"`
	ServiceBusQueueMessage ServiceBusQueueMessagePtrInput `pulumi:"serviceBusQueueMessage"`
	ServiceBusTopicMessage ServiceBusTopicMessagePtrInput `pulumi:"serviceBusTopicMessage"`
	Type                   JobActionTypePtrInput          `pulumi:"type"`
}

func (JobErrorActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorAction)(nil)).Elem()
}

func (i JobErrorActionArgs) ToJobErrorActionOutput() JobErrorActionOutput {
	return i.ToJobErrorActionOutputWithContext(context.Background())
}

func (i JobErrorActionArgs) ToJobErrorActionOutputWithContext(ctx context.Context) JobErrorActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobErrorActionOutput)
}

func (i JobErrorActionArgs) ToJobErrorActionPtrOutput() JobErrorActionPtrOutput {
	return i.ToJobErrorActionPtrOutputWithContext(context.Background())
}

func (i JobErrorActionArgs) ToJobErrorActionPtrOutputWithContext(ctx context.Context) JobErrorActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobErrorActionOutput).ToJobErrorActionPtrOutputWithContext(ctx)
}









type JobErrorActionPtrInput interface {
	pulumi.Input

	ToJobErrorActionPtrOutput() JobErrorActionPtrOutput
	ToJobErrorActionPtrOutputWithContext(context.Context) JobErrorActionPtrOutput
}

type jobErrorActionPtrType JobErrorActionArgs

func JobErrorActionPtr(v *JobErrorActionArgs) JobErrorActionPtrInput {
	return (*jobErrorActionPtrType)(v)
}

func (*jobErrorActionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobErrorAction)(nil)).Elem()
}

func (i *jobErrorActionPtrType) ToJobErrorActionPtrOutput() JobErrorActionPtrOutput {
	return i.ToJobErrorActionPtrOutputWithContext(context.Background())
}

func (i *jobErrorActionPtrType) ToJobErrorActionPtrOutputWithContext(ctx context.Context) JobErrorActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobErrorActionPtrOutput)
}

type JobErrorActionOutput struct{ *pulumi.OutputState }

func (JobErrorActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorAction)(nil)).Elem()
}

func (o JobErrorActionOutput) ToJobErrorActionOutput() JobErrorActionOutput {
	return o
}

func (o JobErrorActionOutput) ToJobErrorActionOutputWithContext(ctx context.Context) JobErrorActionOutput {
	return o
}

func (o JobErrorActionOutput) ToJobErrorActionPtrOutput() JobErrorActionPtrOutput {
	return o.ToJobErrorActionPtrOutputWithContext(context.Background())
}

func (o JobErrorActionOutput) ToJobErrorActionPtrOutputWithContext(ctx context.Context) JobErrorActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobErrorAction) *JobErrorAction {
		return &v
	}).(JobErrorActionPtrOutput)
}

func (o JobErrorActionOutput) QueueMessage() StorageQueueMessagePtrOutput {
	return o.ApplyT(func(v JobErrorAction) *StorageQueueMessage { return v.QueueMessage }).(StorageQueueMessagePtrOutput)
}

func (o JobErrorActionOutput) Request() HttpRequestPtrOutput {
	return o.ApplyT(func(v JobErrorAction) *HttpRequest { return v.Request }).(HttpRequestPtrOutput)
}

func (o JobErrorActionOutput) RetryPolicy() RetryPolicyPtrOutput {
	return o.ApplyT(func(v JobErrorAction) *RetryPolicy { return v.RetryPolicy }).(RetryPolicyPtrOutput)
}

func (o JobErrorActionOutput) ServiceBusQueueMessage() ServiceBusQueueMessagePtrOutput {
	return o.ApplyT(func(v JobErrorAction) *ServiceBusQueueMessage { return v.ServiceBusQueueMessage }).(ServiceBusQueueMessagePtrOutput)
}

func (o JobErrorActionOutput) ServiceBusTopicMessage() ServiceBusTopicMessagePtrOutput {
	return o.ApplyT(func(v JobErrorAction) *ServiceBusTopicMessage { return v.ServiceBusTopicMessage }).(ServiceBusTopicMessagePtrOutput)
}

func (o JobErrorActionOutput) Type() JobActionTypePtrOutput {
	return o.ApplyT(func(v JobErrorAction) *JobActionType { return v.Type }).(JobActionTypePtrOutput)
}

type JobErrorActionPtrOutput struct{ *pulumi.OutputState }

func (JobErrorActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobErrorAction)(nil)).Elem()
}

func (o JobErrorActionPtrOutput) ToJobErrorActionPtrOutput() JobErrorActionPtrOutput {
	return o
}

func (o JobErrorActionPtrOutput) ToJobErrorActionPtrOutputWithContext(ctx context.Context) JobErrorActionPtrOutput {
	return o
}

func (o JobErrorActionPtrOutput) Elem() JobErrorActionOutput {
	return o.ApplyT(func(v *JobErrorAction) JobErrorAction {
		if v != nil {
			return *v
		}
		var ret JobErrorAction
		return ret
	}).(JobErrorActionOutput)
}

func (o JobErrorActionPtrOutput) QueueMessage() StorageQueueMessagePtrOutput {
	return o.ApplyT(func(v *JobErrorAction) *StorageQueueMessage {
		if v == nil {
			return nil
		}
		return v.QueueMessage
	}).(StorageQueueMessagePtrOutput)
}

func (o JobErrorActionPtrOutput) Request() HttpRequestPtrOutput {
	return o.ApplyT(func(v *JobErrorAction) *HttpRequest {
		if v == nil {
			return nil
		}
		return v.Request
	}).(HttpRequestPtrOutput)
}

func (o JobErrorActionPtrOutput) RetryPolicy() RetryPolicyPtrOutput {
	return o.ApplyT(func(v *JobErrorAction) *RetryPolicy {
		if v == nil {
			return nil
		}
		return v.RetryPolicy
	}).(RetryPolicyPtrOutput)
}

func (o JobErrorActionPtrOutput) ServiceBusQueueMessage() ServiceBusQueueMessagePtrOutput {
	return o.ApplyT(func(v *JobErrorAction) *ServiceBusQueueMessage {
		if v == nil {
			return nil
		}
		return v.ServiceBusQueueMessage
	}).(ServiceBusQueueMessagePtrOutput)
}

func (o JobErrorActionPtrOutput) ServiceBusTopicMessage() ServiceBusTopicMessagePtrOutput {
	return o.ApplyT(func(v *JobErrorAction) *ServiceBusTopicMessage {
		if v == nil {
			return nil
		}
		return v.ServiceBusTopicMessage
	}).(ServiceBusTopicMessagePtrOutput)
}

func (o JobErrorActionPtrOutput) Type() JobActionTypePtrOutput {
	return o.ApplyT(func(v *JobErrorAction) *JobActionType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(JobActionTypePtrOutput)
}

type JobErrorActionResponse struct {
	QueueMessage           *StorageQueueMessageResponse    `pulumi:"queueMessage"`
	Request                *HttpRequestResponse            `pulumi:"request"`
	RetryPolicy            *RetryPolicyResponse            `pulumi:"retryPolicy"`
	ServiceBusQueueMessage *ServiceBusQueueMessageResponse `pulumi:"serviceBusQueueMessage"`
	ServiceBusTopicMessage *ServiceBusTopicMessageResponse `pulumi:"serviceBusTopicMessage"`
	Type                   *string                         `pulumi:"type"`
}





type JobErrorActionResponseInput interface {
	pulumi.Input

	ToJobErrorActionResponseOutput() JobErrorActionResponseOutput
	ToJobErrorActionResponseOutputWithContext(context.Context) JobErrorActionResponseOutput
}

type JobErrorActionResponseArgs struct {
	QueueMessage           StorageQueueMessageResponsePtrInput    `pulumi:"queueMessage"`
	Request                HttpRequestResponsePtrInput            `pulumi:"request"`
	RetryPolicy            RetryPolicyResponsePtrInput            `pulumi:"retryPolicy"`
	ServiceBusQueueMessage ServiceBusQueueMessageResponsePtrInput `pulumi:"serviceBusQueueMessage"`
	ServiceBusTopicMessage ServiceBusTopicMessageResponsePtrInput `pulumi:"serviceBusTopicMessage"`
	Type                   pulumi.StringPtrInput                  `pulumi:"type"`
}

func (JobErrorActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorActionResponse)(nil)).Elem()
}

func (i JobErrorActionResponseArgs) ToJobErrorActionResponseOutput() JobErrorActionResponseOutput {
	return i.ToJobErrorActionResponseOutputWithContext(context.Background())
}

func (i JobErrorActionResponseArgs) ToJobErrorActionResponseOutputWithContext(ctx context.Context) JobErrorActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobErrorActionResponseOutput)
}

func (i JobErrorActionResponseArgs) ToJobErrorActionResponsePtrOutput() JobErrorActionResponsePtrOutput {
	return i.ToJobErrorActionResponsePtrOutputWithContext(context.Background())
}

func (i JobErrorActionResponseArgs) ToJobErrorActionResponsePtrOutputWithContext(ctx context.Context) JobErrorActionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobErrorActionResponseOutput).ToJobErrorActionResponsePtrOutputWithContext(ctx)
}









type JobErrorActionResponsePtrInput interface {
	pulumi.Input

	ToJobErrorActionResponsePtrOutput() JobErrorActionResponsePtrOutput
	ToJobErrorActionResponsePtrOutputWithContext(context.Context) JobErrorActionResponsePtrOutput
}

type jobErrorActionResponsePtrType JobErrorActionResponseArgs

func JobErrorActionResponsePtr(v *JobErrorActionResponseArgs) JobErrorActionResponsePtrInput {
	return (*jobErrorActionResponsePtrType)(v)
}

func (*jobErrorActionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobErrorActionResponse)(nil)).Elem()
}

func (i *jobErrorActionResponsePtrType) ToJobErrorActionResponsePtrOutput() JobErrorActionResponsePtrOutput {
	return i.ToJobErrorActionResponsePtrOutputWithContext(context.Background())
}

func (i *jobErrorActionResponsePtrType) ToJobErrorActionResponsePtrOutputWithContext(ctx context.Context) JobErrorActionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobErrorActionResponsePtrOutput)
}

type JobErrorActionResponseOutput struct{ *pulumi.OutputState }

func (JobErrorActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorActionResponse)(nil)).Elem()
}

func (o JobErrorActionResponseOutput) ToJobErrorActionResponseOutput() JobErrorActionResponseOutput {
	return o
}

func (o JobErrorActionResponseOutput) ToJobErrorActionResponseOutputWithContext(ctx context.Context) JobErrorActionResponseOutput {
	return o
}

func (o JobErrorActionResponseOutput) ToJobErrorActionResponsePtrOutput() JobErrorActionResponsePtrOutput {
	return o.ToJobErrorActionResponsePtrOutputWithContext(context.Background())
}

func (o JobErrorActionResponseOutput) ToJobErrorActionResponsePtrOutputWithContext(ctx context.Context) JobErrorActionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobErrorActionResponse) *JobErrorActionResponse {
		return &v
	}).(JobErrorActionResponsePtrOutput)
}

func (o JobErrorActionResponseOutput) QueueMessage() StorageQueueMessageResponsePtrOutput {
	return o.ApplyT(func(v JobErrorActionResponse) *StorageQueueMessageResponse { return v.QueueMessage }).(StorageQueueMessageResponsePtrOutput)
}

func (o JobErrorActionResponseOutput) Request() HttpRequestResponsePtrOutput {
	return o.ApplyT(func(v JobErrorActionResponse) *HttpRequestResponse { return v.Request }).(HttpRequestResponsePtrOutput)
}

func (o JobErrorActionResponseOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v JobErrorActionResponse) *RetryPolicyResponse { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

func (o JobErrorActionResponseOutput) ServiceBusQueueMessage() ServiceBusQueueMessageResponsePtrOutput {
	return o.ApplyT(func(v JobErrorActionResponse) *ServiceBusQueueMessageResponse { return v.ServiceBusQueueMessage }).(ServiceBusQueueMessageResponsePtrOutput)
}

func (o JobErrorActionResponseOutput) ServiceBusTopicMessage() ServiceBusTopicMessageResponsePtrOutput {
	return o.ApplyT(func(v JobErrorActionResponse) *ServiceBusTopicMessageResponse { return v.ServiceBusTopicMessage }).(ServiceBusTopicMessageResponsePtrOutput)
}

func (o JobErrorActionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobErrorActionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type JobErrorActionResponsePtrOutput struct{ *pulumi.OutputState }

func (JobErrorActionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobErrorActionResponse)(nil)).Elem()
}

func (o JobErrorActionResponsePtrOutput) ToJobErrorActionResponsePtrOutput() JobErrorActionResponsePtrOutput {
	return o
}

func (o JobErrorActionResponsePtrOutput) ToJobErrorActionResponsePtrOutputWithContext(ctx context.Context) JobErrorActionResponsePtrOutput {
	return o
}

func (o JobErrorActionResponsePtrOutput) Elem() JobErrorActionResponseOutput {
	return o.ApplyT(func(v *JobErrorActionResponse) JobErrorActionResponse {
		if v != nil {
			return *v
		}
		var ret JobErrorActionResponse
		return ret
	}).(JobErrorActionResponseOutput)
}

func (o JobErrorActionResponsePtrOutput) QueueMessage() StorageQueueMessageResponsePtrOutput {
	return o.ApplyT(func(v *JobErrorActionResponse) *StorageQueueMessageResponse {
		if v == nil {
			return nil
		}
		return v.QueueMessage
	}).(StorageQueueMessageResponsePtrOutput)
}

func (o JobErrorActionResponsePtrOutput) Request() HttpRequestResponsePtrOutput {
	return o.ApplyT(func(v *JobErrorActionResponse) *HttpRequestResponse {
		if v == nil {
			return nil
		}
		return v.Request
	}).(HttpRequestResponsePtrOutput)
}

func (o JobErrorActionResponsePtrOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v *JobErrorActionResponse) *RetryPolicyResponse {
		if v == nil {
			return nil
		}
		return v.RetryPolicy
	}).(RetryPolicyResponsePtrOutput)
}

func (o JobErrorActionResponsePtrOutput) ServiceBusQueueMessage() ServiceBusQueueMessageResponsePtrOutput {
	return o.ApplyT(func(v *JobErrorActionResponse) *ServiceBusQueueMessageResponse {
		if v == nil {
			return nil
		}
		return v.ServiceBusQueueMessage
	}).(ServiceBusQueueMessageResponsePtrOutput)
}

func (o JobErrorActionResponsePtrOutput) ServiceBusTopicMessage() ServiceBusTopicMessageResponsePtrOutput {
	return o.ApplyT(func(v *JobErrorActionResponse) *ServiceBusTopicMessageResponse {
		if v == nil {
			return nil
		}
		return v.ServiceBusTopicMessage
	}).(ServiceBusTopicMessageResponsePtrOutput)
}

func (o JobErrorActionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobErrorActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type JobMaxRecurrence struct {
	Frequency *RecurrenceFrequency `pulumi:"frequency"`
	Interval  *int                 `pulumi:"interval"`
}





type JobMaxRecurrenceInput interface {
	pulumi.Input

	ToJobMaxRecurrenceOutput() JobMaxRecurrenceOutput
	ToJobMaxRecurrenceOutputWithContext(context.Context) JobMaxRecurrenceOutput
}

type JobMaxRecurrenceArgs struct {
	Frequency RecurrenceFrequencyPtrInput `pulumi:"frequency"`
	Interval  pulumi.IntPtrInput          `pulumi:"interval"`
}

func (JobMaxRecurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobMaxRecurrence)(nil)).Elem()
}

func (i JobMaxRecurrenceArgs) ToJobMaxRecurrenceOutput() JobMaxRecurrenceOutput {
	return i.ToJobMaxRecurrenceOutputWithContext(context.Background())
}

func (i JobMaxRecurrenceArgs) ToJobMaxRecurrenceOutputWithContext(ctx context.Context) JobMaxRecurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobMaxRecurrenceOutput)
}

func (i JobMaxRecurrenceArgs) ToJobMaxRecurrencePtrOutput() JobMaxRecurrencePtrOutput {
	return i.ToJobMaxRecurrencePtrOutputWithContext(context.Background())
}

func (i JobMaxRecurrenceArgs) ToJobMaxRecurrencePtrOutputWithContext(ctx context.Context) JobMaxRecurrencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobMaxRecurrenceOutput).ToJobMaxRecurrencePtrOutputWithContext(ctx)
}









type JobMaxRecurrencePtrInput interface {
	pulumi.Input

	ToJobMaxRecurrencePtrOutput() JobMaxRecurrencePtrOutput
	ToJobMaxRecurrencePtrOutputWithContext(context.Context) JobMaxRecurrencePtrOutput
}

type jobMaxRecurrencePtrType JobMaxRecurrenceArgs

func JobMaxRecurrencePtr(v *JobMaxRecurrenceArgs) JobMaxRecurrencePtrInput {
	return (*jobMaxRecurrencePtrType)(v)
}

func (*jobMaxRecurrencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobMaxRecurrence)(nil)).Elem()
}

func (i *jobMaxRecurrencePtrType) ToJobMaxRecurrencePtrOutput() JobMaxRecurrencePtrOutput {
	return i.ToJobMaxRecurrencePtrOutputWithContext(context.Background())
}

func (i *jobMaxRecurrencePtrType) ToJobMaxRecurrencePtrOutputWithContext(ctx context.Context) JobMaxRecurrencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobMaxRecurrencePtrOutput)
}

type JobMaxRecurrenceOutput struct{ *pulumi.OutputState }

func (JobMaxRecurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobMaxRecurrence)(nil)).Elem()
}

func (o JobMaxRecurrenceOutput) ToJobMaxRecurrenceOutput() JobMaxRecurrenceOutput {
	return o
}

func (o JobMaxRecurrenceOutput) ToJobMaxRecurrenceOutputWithContext(ctx context.Context) JobMaxRecurrenceOutput {
	return o
}

func (o JobMaxRecurrenceOutput) ToJobMaxRecurrencePtrOutput() JobMaxRecurrencePtrOutput {
	return o.ToJobMaxRecurrencePtrOutputWithContext(context.Background())
}

func (o JobMaxRecurrenceOutput) ToJobMaxRecurrencePtrOutputWithContext(ctx context.Context) JobMaxRecurrencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobMaxRecurrence) *JobMaxRecurrence {
		return &v
	}).(JobMaxRecurrencePtrOutput)
}

func (o JobMaxRecurrenceOutput) Frequency() RecurrenceFrequencyPtrOutput {
	return o.ApplyT(func(v JobMaxRecurrence) *RecurrenceFrequency { return v.Frequency }).(RecurrenceFrequencyPtrOutput)
}

func (o JobMaxRecurrenceOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobMaxRecurrence) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

type JobMaxRecurrencePtrOutput struct{ *pulumi.OutputState }

func (JobMaxRecurrencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobMaxRecurrence)(nil)).Elem()
}

func (o JobMaxRecurrencePtrOutput) ToJobMaxRecurrencePtrOutput() JobMaxRecurrencePtrOutput {
	return o
}

func (o JobMaxRecurrencePtrOutput) ToJobMaxRecurrencePtrOutputWithContext(ctx context.Context) JobMaxRecurrencePtrOutput {
	return o
}

func (o JobMaxRecurrencePtrOutput) Elem() JobMaxRecurrenceOutput {
	return o.ApplyT(func(v *JobMaxRecurrence) JobMaxRecurrence {
		if v != nil {
			return *v
		}
		var ret JobMaxRecurrence
		return ret
	}).(JobMaxRecurrenceOutput)
}

func (o JobMaxRecurrencePtrOutput) Frequency() RecurrenceFrequencyPtrOutput {
	return o.ApplyT(func(v *JobMaxRecurrence) *RecurrenceFrequency {
		if v == nil {
			return nil
		}
		return v.Frequency
	}).(RecurrenceFrequencyPtrOutput)
}

func (o JobMaxRecurrencePtrOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobMaxRecurrence) *int {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.IntPtrOutput)
}

type JobMaxRecurrenceResponse struct {
	Frequency *string `pulumi:"frequency"`
	Interval  *int    `pulumi:"interval"`
}





type JobMaxRecurrenceResponseInput interface {
	pulumi.Input

	ToJobMaxRecurrenceResponseOutput() JobMaxRecurrenceResponseOutput
	ToJobMaxRecurrenceResponseOutputWithContext(context.Context) JobMaxRecurrenceResponseOutput
}

type JobMaxRecurrenceResponseArgs struct {
	Frequency pulumi.StringPtrInput `pulumi:"frequency"`
	Interval  pulumi.IntPtrInput    `pulumi:"interval"`
}

func (JobMaxRecurrenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobMaxRecurrenceResponse)(nil)).Elem()
}

func (i JobMaxRecurrenceResponseArgs) ToJobMaxRecurrenceResponseOutput() JobMaxRecurrenceResponseOutput {
	return i.ToJobMaxRecurrenceResponseOutputWithContext(context.Background())
}

func (i JobMaxRecurrenceResponseArgs) ToJobMaxRecurrenceResponseOutputWithContext(ctx context.Context) JobMaxRecurrenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobMaxRecurrenceResponseOutput)
}

func (i JobMaxRecurrenceResponseArgs) ToJobMaxRecurrenceResponsePtrOutput() JobMaxRecurrenceResponsePtrOutput {
	return i.ToJobMaxRecurrenceResponsePtrOutputWithContext(context.Background())
}

func (i JobMaxRecurrenceResponseArgs) ToJobMaxRecurrenceResponsePtrOutputWithContext(ctx context.Context) JobMaxRecurrenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobMaxRecurrenceResponseOutput).ToJobMaxRecurrenceResponsePtrOutputWithContext(ctx)
}









type JobMaxRecurrenceResponsePtrInput interface {
	pulumi.Input

	ToJobMaxRecurrenceResponsePtrOutput() JobMaxRecurrenceResponsePtrOutput
	ToJobMaxRecurrenceResponsePtrOutputWithContext(context.Context) JobMaxRecurrenceResponsePtrOutput
}

type jobMaxRecurrenceResponsePtrType JobMaxRecurrenceResponseArgs

func JobMaxRecurrenceResponsePtr(v *JobMaxRecurrenceResponseArgs) JobMaxRecurrenceResponsePtrInput {
	return (*jobMaxRecurrenceResponsePtrType)(v)
}

func (*jobMaxRecurrenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobMaxRecurrenceResponse)(nil)).Elem()
}

func (i *jobMaxRecurrenceResponsePtrType) ToJobMaxRecurrenceResponsePtrOutput() JobMaxRecurrenceResponsePtrOutput {
	return i.ToJobMaxRecurrenceResponsePtrOutputWithContext(context.Background())
}

func (i *jobMaxRecurrenceResponsePtrType) ToJobMaxRecurrenceResponsePtrOutputWithContext(ctx context.Context) JobMaxRecurrenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobMaxRecurrenceResponsePtrOutput)
}

type JobMaxRecurrenceResponseOutput struct{ *pulumi.OutputState }

func (JobMaxRecurrenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobMaxRecurrenceResponse)(nil)).Elem()
}

func (o JobMaxRecurrenceResponseOutput) ToJobMaxRecurrenceResponseOutput() JobMaxRecurrenceResponseOutput {
	return o
}

func (o JobMaxRecurrenceResponseOutput) ToJobMaxRecurrenceResponseOutputWithContext(ctx context.Context) JobMaxRecurrenceResponseOutput {
	return o
}

func (o JobMaxRecurrenceResponseOutput) ToJobMaxRecurrenceResponsePtrOutput() JobMaxRecurrenceResponsePtrOutput {
	return o.ToJobMaxRecurrenceResponsePtrOutputWithContext(context.Background())
}

func (o JobMaxRecurrenceResponseOutput) ToJobMaxRecurrenceResponsePtrOutputWithContext(ctx context.Context) JobMaxRecurrenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobMaxRecurrenceResponse) *JobMaxRecurrenceResponse {
		return &v
	}).(JobMaxRecurrenceResponsePtrOutput)
}

func (o JobMaxRecurrenceResponseOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobMaxRecurrenceResponse) *string { return v.Frequency }).(pulumi.StringPtrOutput)
}

func (o JobMaxRecurrenceResponseOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobMaxRecurrenceResponse) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

type JobMaxRecurrenceResponsePtrOutput struct{ *pulumi.OutputState }

func (JobMaxRecurrenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobMaxRecurrenceResponse)(nil)).Elem()
}

func (o JobMaxRecurrenceResponsePtrOutput) ToJobMaxRecurrenceResponsePtrOutput() JobMaxRecurrenceResponsePtrOutput {
	return o
}

func (o JobMaxRecurrenceResponsePtrOutput) ToJobMaxRecurrenceResponsePtrOutputWithContext(ctx context.Context) JobMaxRecurrenceResponsePtrOutput {
	return o
}

func (o JobMaxRecurrenceResponsePtrOutput) Elem() JobMaxRecurrenceResponseOutput {
	return o.ApplyT(func(v *JobMaxRecurrenceResponse) JobMaxRecurrenceResponse {
		if v != nil {
			return *v
		}
		var ret JobMaxRecurrenceResponse
		return ret
	}).(JobMaxRecurrenceResponseOutput)
}

func (o JobMaxRecurrenceResponsePtrOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobMaxRecurrenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Frequency
	}).(pulumi.StringPtrOutput)
}

func (o JobMaxRecurrenceResponsePtrOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobMaxRecurrenceResponse) *int {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.IntPtrOutput)
}

type JobProperties struct {
	Action     *JobAction     `pulumi:"action"`
	Recurrence *JobRecurrence `pulumi:"recurrence"`
	StartTime  *string        `pulumi:"startTime"`
	State      *JobStateEnum  `pulumi:"state"`
}





type JobPropertiesInput interface {
	pulumi.Input

	ToJobPropertiesOutput() JobPropertiesOutput
	ToJobPropertiesOutputWithContext(context.Context) JobPropertiesOutput
}

type JobPropertiesArgs struct {
	Action     JobActionPtrInput     `pulumi:"action"`
	Recurrence JobRecurrencePtrInput `pulumi:"recurrence"`
	StartTime  pulumi.StringPtrInput `pulumi:"startTime"`
	State      JobStateEnumPtrInput  `pulumi:"state"`
}

func (JobPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobProperties)(nil)).Elem()
}

func (i JobPropertiesArgs) ToJobPropertiesOutput() JobPropertiesOutput {
	return i.ToJobPropertiesOutputWithContext(context.Background())
}

func (i JobPropertiesArgs) ToJobPropertiesOutputWithContext(ctx context.Context) JobPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPropertiesOutput)
}

func (i JobPropertiesArgs) ToJobPropertiesPtrOutput() JobPropertiesPtrOutput {
	return i.ToJobPropertiesPtrOutputWithContext(context.Background())
}

func (i JobPropertiesArgs) ToJobPropertiesPtrOutputWithContext(ctx context.Context) JobPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPropertiesOutput).ToJobPropertiesPtrOutputWithContext(ctx)
}









type JobPropertiesPtrInput interface {
	pulumi.Input

	ToJobPropertiesPtrOutput() JobPropertiesPtrOutput
	ToJobPropertiesPtrOutputWithContext(context.Context) JobPropertiesPtrOutput
}

type jobPropertiesPtrType JobPropertiesArgs

func JobPropertiesPtr(v *JobPropertiesArgs) JobPropertiesPtrInput {
	return (*jobPropertiesPtrType)(v)
}

func (*jobPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobProperties)(nil)).Elem()
}

func (i *jobPropertiesPtrType) ToJobPropertiesPtrOutput() JobPropertiesPtrOutput {
	return i.ToJobPropertiesPtrOutputWithContext(context.Background())
}

func (i *jobPropertiesPtrType) ToJobPropertiesPtrOutputWithContext(ctx context.Context) JobPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPropertiesPtrOutput)
}

type JobPropertiesOutput struct{ *pulumi.OutputState }

func (JobPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobProperties)(nil)).Elem()
}

func (o JobPropertiesOutput) ToJobPropertiesOutput() JobPropertiesOutput {
	return o
}

func (o JobPropertiesOutput) ToJobPropertiesOutputWithContext(ctx context.Context) JobPropertiesOutput {
	return o
}

func (o JobPropertiesOutput) ToJobPropertiesPtrOutput() JobPropertiesPtrOutput {
	return o.ToJobPropertiesPtrOutputWithContext(context.Background())
}

func (o JobPropertiesOutput) ToJobPropertiesPtrOutputWithContext(ctx context.Context) JobPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobProperties) *JobProperties {
		return &v
	}).(JobPropertiesPtrOutput)
}

func (o JobPropertiesOutput) Action() JobActionPtrOutput {
	return o.ApplyT(func(v JobProperties) *JobAction { return v.Action }).(JobActionPtrOutput)
}

func (o JobPropertiesOutput) Recurrence() JobRecurrencePtrOutput {
	return o.ApplyT(func(v JobProperties) *JobRecurrence { return v.Recurrence }).(JobRecurrencePtrOutput)
}

func (o JobPropertiesOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobProperties) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o JobPropertiesOutput) State() JobStateEnumPtrOutput {
	return o.ApplyT(func(v JobProperties) *JobStateEnum { return v.State }).(JobStateEnumPtrOutput)
}

type JobPropertiesPtrOutput struct{ *pulumi.OutputState }

func (JobPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobProperties)(nil)).Elem()
}

func (o JobPropertiesPtrOutput) ToJobPropertiesPtrOutput() JobPropertiesPtrOutput {
	return o
}

func (o JobPropertiesPtrOutput) ToJobPropertiesPtrOutputWithContext(ctx context.Context) JobPropertiesPtrOutput {
	return o
}

func (o JobPropertiesPtrOutput) Elem() JobPropertiesOutput {
	return o.ApplyT(func(v *JobProperties) JobProperties {
		if v != nil {
			return *v
		}
		var ret JobProperties
		return ret
	}).(JobPropertiesOutput)
}

func (o JobPropertiesPtrOutput) Action() JobActionPtrOutput {
	return o.ApplyT(func(v *JobProperties) *JobAction {
		if v == nil {
			return nil
		}
		return v.Action
	}).(JobActionPtrOutput)
}

func (o JobPropertiesPtrOutput) Recurrence() JobRecurrencePtrOutput {
	return o.ApplyT(func(v *JobProperties) *JobRecurrence {
		if v == nil {
			return nil
		}
		return v.Recurrence
	}).(JobRecurrencePtrOutput)
}

func (o JobPropertiesPtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobProperties) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o JobPropertiesPtrOutput) State() JobStateEnumPtrOutput {
	return o.ApplyT(func(v *JobProperties) *JobStateEnum {
		if v == nil {
			return nil
		}
		return v.State
	}).(JobStateEnumPtrOutput)
}

type JobPropertiesResponse struct {
	Action     *JobActionResponse     `pulumi:"action"`
	Recurrence *JobRecurrenceResponse `pulumi:"recurrence"`
	StartTime  *string                `pulumi:"startTime"`
	State      *string                `pulumi:"state"`
	Status     JobStatusResponse      `pulumi:"status"`
}





type JobPropertiesResponseInput interface {
	pulumi.Input

	ToJobPropertiesResponseOutput() JobPropertiesResponseOutput
	ToJobPropertiesResponseOutputWithContext(context.Context) JobPropertiesResponseOutput
}

type JobPropertiesResponseArgs struct {
	Action     JobActionResponsePtrInput     `pulumi:"action"`
	Recurrence JobRecurrenceResponsePtrInput `pulumi:"recurrence"`
	StartTime  pulumi.StringPtrInput         `pulumi:"startTime"`
	State      pulumi.StringPtrInput         `pulumi:"state"`
	Status     JobStatusResponseInput        `pulumi:"status"`
}

func (JobPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobPropertiesResponse)(nil)).Elem()
}

func (i JobPropertiesResponseArgs) ToJobPropertiesResponseOutput() JobPropertiesResponseOutput {
	return i.ToJobPropertiesResponseOutputWithContext(context.Background())
}

func (i JobPropertiesResponseArgs) ToJobPropertiesResponseOutputWithContext(ctx context.Context) JobPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPropertiesResponseOutput)
}

func (i JobPropertiesResponseArgs) ToJobPropertiesResponsePtrOutput() JobPropertiesResponsePtrOutput {
	return i.ToJobPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i JobPropertiesResponseArgs) ToJobPropertiesResponsePtrOutputWithContext(ctx context.Context) JobPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPropertiesResponseOutput).ToJobPropertiesResponsePtrOutputWithContext(ctx)
}









type JobPropertiesResponsePtrInput interface {
	pulumi.Input

	ToJobPropertiesResponsePtrOutput() JobPropertiesResponsePtrOutput
	ToJobPropertiesResponsePtrOutputWithContext(context.Context) JobPropertiesResponsePtrOutput
}

type jobPropertiesResponsePtrType JobPropertiesResponseArgs

func JobPropertiesResponsePtr(v *JobPropertiesResponseArgs) JobPropertiesResponsePtrInput {
	return (*jobPropertiesResponsePtrType)(v)
}

func (*jobPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobPropertiesResponse)(nil)).Elem()
}

func (i *jobPropertiesResponsePtrType) ToJobPropertiesResponsePtrOutput() JobPropertiesResponsePtrOutput {
	return i.ToJobPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *jobPropertiesResponsePtrType) ToJobPropertiesResponsePtrOutputWithContext(ctx context.Context) JobPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPropertiesResponsePtrOutput)
}

type JobPropertiesResponseOutput struct{ *pulumi.OutputState }

func (JobPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobPropertiesResponse)(nil)).Elem()
}

func (o JobPropertiesResponseOutput) ToJobPropertiesResponseOutput() JobPropertiesResponseOutput {
	return o
}

func (o JobPropertiesResponseOutput) ToJobPropertiesResponseOutputWithContext(ctx context.Context) JobPropertiesResponseOutput {
	return o
}

func (o JobPropertiesResponseOutput) ToJobPropertiesResponsePtrOutput() JobPropertiesResponsePtrOutput {
	return o.ToJobPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o JobPropertiesResponseOutput) ToJobPropertiesResponsePtrOutputWithContext(ctx context.Context) JobPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobPropertiesResponse) *JobPropertiesResponse {
		return &v
	}).(JobPropertiesResponsePtrOutput)
}

func (o JobPropertiesResponseOutput) Action() JobActionResponsePtrOutput {
	return o.ApplyT(func(v JobPropertiesResponse) *JobActionResponse { return v.Action }).(JobActionResponsePtrOutput)
}

func (o JobPropertiesResponseOutput) Recurrence() JobRecurrenceResponsePtrOutput {
	return o.ApplyT(func(v JobPropertiesResponse) *JobRecurrenceResponse { return v.Recurrence }).(JobRecurrenceResponsePtrOutput)
}

func (o JobPropertiesResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobPropertiesResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o JobPropertiesResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobPropertiesResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o JobPropertiesResponseOutput) Status() JobStatusResponseOutput {
	return o.ApplyT(func(v JobPropertiesResponse) JobStatusResponse { return v.Status }).(JobStatusResponseOutput)
}

type JobPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (JobPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobPropertiesResponse)(nil)).Elem()
}

func (o JobPropertiesResponsePtrOutput) ToJobPropertiesResponsePtrOutput() JobPropertiesResponsePtrOutput {
	return o
}

func (o JobPropertiesResponsePtrOutput) ToJobPropertiesResponsePtrOutputWithContext(ctx context.Context) JobPropertiesResponsePtrOutput {
	return o
}

func (o JobPropertiesResponsePtrOutput) Elem() JobPropertiesResponseOutput {
	return o.ApplyT(func(v *JobPropertiesResponse) JobPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret JobPropertiesResponse
		return ret
	}).(JobPropertiesResponseOutput)
}

func (o JobPropertiesResponsePtrOutput) Action() JobActionResponsePtrOutput {
	return o.ApplyT(func(v *JobPropertiesResponse) *JobActionResponse {
		if v == nil {
			return nil
		}
		return v.Action
	}).(JobActionResponsePtrOutput)
}

func (o JobPropertiesResponsePtrOutput) Recurrence() JobRecurrenceResponsePtrOutput {
	return o.ApplyT(func(v *JobPropertiesResponse) *JobRecurrenceResponse {
		if v == nil {
			return nil
		}
		return v.Recurrence
	}).(JobRecurrenceResponsePtrOutput)
}

func (o JobPropertiesResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o JobPropertiesResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

func (o JobPropertiesResponsePtrOutput) Status() JobStatusResponsePtrOutput {
	return o.ApplyT(func(v *JobPropertiesResponse) *JobStatusResponse {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(JobStatusResponsePtrOutput)
}

type JobRecurrence struct {
	Count     *int                   `pulumi:"count"`
	EndTime   *string                `pulumi:"endTime"`
	Frequency *RecurrenceFrequency   `pulumi:"frequency"`
	Interval  *int                   `pulumi:"interval"`
	Schedule  *JobRecurrenceSchedule `pulumi:"schedule"`
}





type JobRecurrenceInput interface {
	pulumi.Input

	ToJobRecurrenceOutput() JobRecurrenceOutput
	ToJobRecurrenceOutputWithContext(context.Context) JobRecurrenceOutput
}

type JobRecurrenceArgs struct {
	Count     pulumi.IntPtrInput            `pulumi:"count"`
	EndTime   pulumi.StringPtrInput         `pulumi:"endTime"`
	Frequency RecurrenceFrequencyPtrInput   `pulumi:"frequency"`
	Interval  pulumi.IntPtrInput            `pulumi:"interval"`
	Schedule  JobRecurrenceSchedulePtrInput `pulumi:"schedule"`
}

func (JobRecurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobRecurrence)(nil)).Elem()
}

func (i JobRecurrenceArgs) ToJobRecurrenceOutput() JobRecurrenceOutput {
	return i.ToJobRecurrenceOutputWithContext(context.Background())
}

func (i JobRecurrenceArgs) ToJobRecurrenceOutputWithContext(ctx context.Context) JobRecurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceOutput)
}

func (i JobRecurrenceArgs) ToJobRecurrencePtrOutput() JobRecurrencePtrOutput {
	return i.ToJobRecurrencePtrOutputWithContext(context.Background())
}

func (i JobRecurrenceArgs) ToJobRecurrencePtrOutputWithContext(ctx context.Context) JobRecurrencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceOutput).ToJobRecurrencePtrOutputWithContext(ctx)
}









type JobRecurrencePtrInput interface {
	pulumi.Input

	ToJobRecurrencePtrOutput() JobRecurrencePtrOutput
	ToJobRecurrencePtrOutputWithContext(context.Context) JobRecurrencePtrOutput
}

type jobRecurrencePtrType JobRecurrenceArgs

func JobRecurrencePtr(v *JobRecurrenceArgs) JobRecurrencePtrInput {
	return (*jobRecurrencePtrType)(v)
}

func (*jobRecurrencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobRecurrence)(nil)).Elem()
}

func (i *jobRecurrencePtrType) ToJobRecurrencePtrOutput() JobRecurrencePtrOutput {
	return i.ToJobRecurrencePtrOutputWithContext(context.Background())
}

func (i *jobRecurrencePtrType) ToJobRecurrencePtrOutputWithContext(ctx context.Context) JobRecurrencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrencePtrOutput)
}

type JobRecurrenceOutput struct{ *pulumi.OutputState }

func (JobRecurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobRecurrence)(nil)).Elem()
}

func (o JobRecurrenceOutput) ToJobRecurrenceOutput() JobRecurrenceOutput {
	return o
}

func (o JobRecurrenceOutput) ToJobRecurrenceOutputWithContext(ctx context.Context) JobRecurrenceOutput {
	return o
}

func (o JobRecurrenceOutput) ToJobRecurrencePtrOutput() JobRecurrencePtrOutput {
	return o.ToJobRecurrencePtrOutputWithContext(context.Background())
}

func (o JobRecurrenceOutput) ToJobRecurrencePtrOutputWithContext(ctx context.Context) JobRecurrencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobRecurrence) *JobRecurrence {
		return &v
	}).(JobRecurrencePtrOutput)
}

func (o JobRecurrenceOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobRecurrence) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o JobRecurrenceOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobRecurrence) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o JobRecurrenceOutput) Frequency() RecurrenceFrequencyPtrOutput {
	return o.ApplyT(func(v JobRecurrence) *RecurrenceFrequency { return v.Frequency }).(RecurrenceFrequencyPtrOutput)
}

func (o JobRecurrenceOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobRecurrence) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o JobRecurrenceOutput) Schedule() JobRecurrenceSchedulePtrOutput {
	return o.ApplyT(func(v JobRecurrence) *JobRecurrenceSchedule { return v.Schedule }).(JobRecurrenceSchedulePtrOutput)
}

type JobRecurrencePtrOutput struct{ *pulumi.OutputState }

func (JobRecurrencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobRecurrence)(nil)).Elem()
}

func (o JobRecurrencePtrOutput) ToJobRecurrencePtrOutput() JobRecurrencePtrOutput {
	return o
}

func (o JobRecurrencePtrOutput) ToJobRecurrencePtrOutputWithContext(ctx context.Context) JobRecurrencePtrOutput {
	return o
}

func (o JobRecurrencePtrOutput) Elem() JobRecurrenceOutput {
	return o.ApplyT(func(v *JobRecurrence) JobRecurrence {
		if v != nil {
			return *v
		}
		var ret JobRecurrence
		return ret
	}).(JobRecurrenceOutput)
}

func (o JobRecurrencePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobRecurrence) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o JobRecurrencePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobRecurrence) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o JobRecurrencePtrOutput) Frequency() RecurrenceFrequencyPtrOutput {
	return o.ApplyT(func(v *JobRecurrence) *RecurrenceFrequency {
		if v == nil {
			return nil
		}
		return v.Frequency
	}).(RecurrenceFrequencyPtrOutput)
}

func (o JobRecurrencePtrOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobRecurrence) *int {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.IntPtrOutput)
}

func (o JobRecurrencePtrOutput) Schedule() JobRecurrenceSchedulePtrOutput {
	return o.ApplyT(func(v *JobRecurrence) *JobRecurrenceSchedule {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(JobRecurrenceSchedulePtrOutput)
}

type JobRecurrenceResponse struct {
	Count     *int                           `pulumi:"count"`
	EndTime   *string                        `pulumi:"endTime"`
	Frequency *string                        `pulumi:"frequency"`
	Interval  *int                           `pulumi:"interval"`
	Schedule  *JobRecurrenceScheduleResponse `pulumi:"schedule"`
}





type JobRecurrenceResponseInput interface {
	pulumi.Input

	ToJobRecurrenceResponseOutput() JobRecurrenceResponseOutput
	ToJobRecurrenceResponseOutputWithContext(context.Context) JobRecurrenceResponseOutput
}

type JobRecurrenceResponseArgs struct {
	Count     pulumi.IntPtrInput                    `pulumi:"count"`
	EndTime   pulumi.StringPtrInput                 `pulumi:"endTime"`
	Frequency pulumi.StringPtrInput                 `pulumi:"frequency"`
	Interval  pulumi.IntPtrInput                    `pulumi:"interval"`
	Schedule  JobRecurrenceScheduleResponsePtrInput `pulumi:"schedule"`
}

func (JobRecurrenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobRecurrenceResponse)(nil)).Elem()
}

func (i JobRecurrenceResponseArgs) ToJobRecurrenceResponseOutput() JobRecurrenceResponseOutput {
	return i.ToJobRecurrenceResponseOutputWithContext(context.Background())
}

func (i JobRecurrenceResponseArgs) ToJobRecurrenceResponseOutputWithContext(ctx context.Context) JobRecurrenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceResponseOutput)
}

func (i JobRecurrenceResponseArgs) ToJobRecurrenceResponsePtrOutput() JobRecurrenceResponsePtrOutput {
	return i.ToJobRecurrenceResponsePtrOutputWithContext(context.Background())
}

func (i JobRecurrenceResponseArgs) ToJobRecurrenceResponsePtrOutputWithContext(ctx context.Context) JobRecurrenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceResponseOutput).ToJobRecurrenceResponsePtrOutputWithContext(ctx)
}









type JobRecurrenceResponsePtrInput interface {
	pulumi.Input

	ToJobRecurrenceResponsePtrOutput() JobRecurrenceResponsePtrOutput
	ToJobRecurrenceResponsePtrOutputWithContext(context.Context) JobRecurrenceResponsePtrOutput
}

type jobRecurrenceResponsePtrType JobRecurrenceResponseArgs

func JobRecurrenceResponsePtr(v *JobRecurrenceResponseArgs) JobRecurrenceResponsePtrInput {
	return (*jobRecurrenceResponsePtrType)(v)
}

func (*jobRecurrenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobRecurrenceResponse)(nil)).Elem()
}

func (i *jobRecurrenceResponsePtrType) ToJobRecurrenceResponsePtrOutput() JobRecurrenceResponsePtrOutput {
	return i.ToJobRecurrenceResponsePtrOutputWithContext(context.Background())
}

func (i *jobRecurrenceResponsePtrType) ToJobRecurrenceResponsePtrOutputWithContext(ctx context.Context) JobRecurrenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceResponsePtrOutput)
}

type JobRecurrenceResponseOutput struct{ *pulumi.OutputState }

func (JobRecurrenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobRecurrenceResponse)(nil)).Elem()
}

func (o JobRecurrenceResponseOutput) ToJobRecurrenceResponseOutput() JobRecurrenceResponseOutput {
	return o
}

func (o JobRecurrenceResponseOutput) ToJobRecurrenceResponseOutputWithContext(ctx context.Context) JobRecurrenceResponseOutput {
	return o
}

func (o JobRecurrenceResponseOutput) ToJobRecurrenceResponsePtrOutput() JobRecurrenceResponsePtrOutput {
	return o.ToJobRecurrenceResponsePtrOutputWithContext(context.Background())
}

func (o JobRecurrenceResponseOutput) ToJobRecurrenceResponsePtrOutputWithContext(ctx context.Context) JobRecurrenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobRecurrenceResponse) *JobRecurrenceResponse {
		return &v
	}).(JobRecurrenceResponsePtrOutput)
}

func (o JobRecurrenceResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobRecurrenceResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o JobRecurrenceResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobRecurrenceResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o JobRecurrenceResponseOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobRecurrenceResponse) *string { return v.Frequency }).(pulumi.StringPtrOutput)
}

func (o JobRecurrenceResponseOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobRecurrenceResponse) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o JobRecurrenceResponseOutput) Schedule() JobRecurrenceScheduleResponsePtrOutput {
	return o.ApplyT(func(v JobRecurrenceResponse) *JobRecurrenceScheduleResponse { return v.Schedule }).(JobRecurrenceScheduleResponsePtrOutput)
}

type JobRecurrenceResponsePtrOutput struct{ *pulumi.OutputState }

func (JobRecurrenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobRecurrenceResponse)(nil)).Elem()
}

func (o JobRecurrenceResponsePtrOutput) ToJobRecurrenceResponsePtrOutput() JobRecurrenceResponsePtrOutput {
	return o
}

func (o JobRecurrenceResponsePtrOutput) ToJobRecurrenceResponsePtrOutputWithContext(ctx context.Context) JobRecurrenceResponsePtrOutput {
	return o
}

func (o JobRecurrenceResponsePtrOutput) Elem() JobRecurrenceResponseOutput {
	return o.ApplyT(func(v *JobRecurrenceResponse) JobRecurrenceResponse {
		if v != nil {
			return *v
		}
		var ret JobRecurrenceResponse
		return ret
	}).(JobRecurrenceResponseOutput)
}

func (o JobRecurrenceResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobRecurrenceResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o JobRecurrenceResponsePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobRecurrenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o JobRecurrenceResponsePtrOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobRecurrenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Frequency
	}).(pulumi.StringPtrOutput)
}

func (o JobRecurrenceResponsePtrOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobRecurrenceResponse) *int {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.IntPtrOutput)
}

func (o JobRecurrenceResponsePtrOutput) Schedule() JobRecurrenceScheduleResponsePtrOutput {
	return o.ApplyT(func(v *JobRecurrenceResponse) *JobRecurrenceScheduleResponse {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(JobRecurrenceScheduleResponsePtrOutput)
}

type JobRecurrenceSchedule struct {
	Hours              []int                                    `pulumi:"hours"`
	Minutes            []int                                    `pulumi:"minutes"`
	MonthDays          []int                                    `pulumi:"monthDays"`
	MonthlyOccurrences []JobRecurrenceScheduleMonthlyOccurrence `pulumi:"monthlyOccurrences"`
	WeekDays           []DayOfWeek                              `pulumi:"weekDays"`
}





type JobRecurrenceScheduleInput interface {
	pulumi.Input

	ToJobRecurrenceScheduleOutput() JobRecurrenceScheduleOutput
	ToJobRecurrenceScheduleOutputWithContext(context.Context) JobRecurrenceScheduleOutput
}

type JobRecurrenceScheduleArgs struct {
	Hours              pulumi.IntArrayInput                             `pulumi:"hours"`
	Minutes            pulumi.IntArrayInput                             `pulumi:"minutes"`
	MonthDays          pulumi.IntArrayInput                             `pulumi:"monthDays"`
	MonthlyOccurrences JobRecurrenceScheduleMonthlyOccurrenceArrayInput `pulumi:"monthlyOccurrences"`
	WeekDays           DayOfWeekArrayInput                              `pulumi:"weekDays"`
}

func (JobRecurrenceScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobRecurrenceSchedule)(nil)).Elem()
}

func (i JobRecurrenceScheduleArgs) ToJobRecurrenceScheduleOutput() JobRecurrenceScheduleOutput {
	return i.ToJobRecurrenceScheduleOutputWithContext(context.Background())
}

func (i JobRecurrenceScheduleArgs) ToJobRecurrenceScheduleOutputWithContext(ctx context.Context) JobRecurrenceScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceScheduleOutput)
}

func (i JobRecurrenceScheduleArgs) ToJobRecurrenceSchedulePtrOutput() JobRecurrenceSchedulePtrOutput {
	return i.ToJobRecurrenceSchedulePtrOutputWithContext(context.Background())
}

func (i JobRecurrenceScheduleArgs) ToJobRecurrenceSchedulePtrOutputWithContext(ctx context.Context) JobRecurrenceSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceScheduleOutput).ToJobRecurrenceSchedulePtrOutputWithContext(ctx)
}









type JobRecurrenceSchedulePtrInput interface {
	pulumi.Input

	ToJobRecurrenceSchedulePtrOutput() JobRecurrenceSchedulePtrOutput
	ToJobRecurrenceSchedulePtrOutputWithContext(context.Context) JobRecurrenceSchedulePtrOutput
}

type jobRecurrenceSchedulePtrType JobRecurrenceScheduleArgs

func JobRecurrenceSchedulePtr(v *JobRecurrenceScheduleArgs) JobRecurrenceSchedulePtrInput {
	return (*jobRecurrenceSchedulePtrType)(v)
}

func (*jobRecurrenceSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobRecurrenceSchedule)(nil)).Elem()
}

func (i *jobRecurrenceSchedulePtrType) ToJobRecurrenceSchedulePtrOutput() JobRecurrenceSchedulePtrOutput {
	return i.ToJobRecurrenceSchedulePtrOutputWithContext(context.Background())
}

func (i *jobRecurrenceSchedulePtrType) ToJobRecurrenceSchedulePtrOutputWithContext(ctx context.Context) JobRecurrenceSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceSchedulePtrOutput)
}

type JobRecurrenceScheduleOutput struct{ *pulumi.OutputState }

func (JobRecurrenceScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobRecurrenceSchedule)(nil)).Elem()
}

func (o JobRecurrenceScheduleOutput) ToJobRecurrenceScheduleOutput() JobRecurrenceScheduleOutput {
	return o
}

func (o JobRecurrenceScheduleOutput) ToJobRecurrenceScheduleOutputWithContext(ctx context.Context) JobRecurrenceScheduleOutput {
	return o
}

func (o JobRecurrenceScheduleOutput) ToJobRecurrenceSchedulePtrOutput() JobRecurrenceSchedulePtrOutput {
	return o.ToJobRecurrenceSchedulePtrOutputWithContext(context.Background())
}

func (o JobRecurrenceScheduleOutput) ToJobRecurrenceSchedulePtrOutputWithContext(ctx context.Context) JobRecurrenceSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobRecurrenceSchedule) *JobRecurrenceSchedule {
		return &v
	}).(JobRecurrenceSchedulePtrOutput)
}

func (o JobRecurrenceScheduleOutput) Hours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v JobRecurrenceSchedule) []int { return v.Hours }).(pulumi.IntArrayOutput)
}

func (o JobRecurrenceScheduleOutput) Minutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v JobRecurrenceSchedule) []int { return v.Minutes }).(pulumi.IntArrayOutput)
}

func (o JobRecurrenceScheduleOutput) MonthDays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v JobRecurrenceSchedule) []int { return v.MonthDays }).(pulumi.IntArrayOutput)
}

func (o JobRecurrenceScheduleOutput) MonthlyOccurrences() JobRecurrenceScheduleMonthlyOccurrenceArrayOutput {
	return o.ApplyT(func(v JobRecurrenceSchedule) []JobRecurrenceScheduleMonthlyOccurrence { return v.MonthlyOccurrences }).(JobRecurrenceScheduleMonthlyOccurrenceArrayOutput)
}

func (o JobRecurrenceScheduleOutput) WeekDays() DayOfWeekArrayOutput {
	return o.ApplyT(func(v JobRecurrenceSchedule) []DayOfWeek { return v.WeekDays }).(DayOfWeekArrayOutput)
}

type JobRecurrenceSchedulePtrOutput struct{ *pulumi.OutputState }

func (JobRecurrenceSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobRecurrenceSchedule)(nil)).Elem()
}

func (o JobRecurrenceSchedulePtrOutput) ToJobRecurrenceSchedulePtrOutput() JobRecurrenceSchedulePtrOutput {
	return o
}

func (o JobRecurrenceSchedulePtrOutput) ToJobRecurrenceSchedulePtrOutputWithContext(ctx context.Context) JobRecurrenceSchedulePtrOutput {
	return o
}

func (o JobRecurrenceSchedulePtrOutput) Elem() JobRecurrenceScheduleOutput {
	return o.ApplyT(func(v *JobRecurrenceSchedule) JobRecurrenceSchedule {
		if v != nil {
			return *v
		}
		var ret JobRecurrenceSchedule
		return ret
	}).(JobRecurrenceScheduleOutput)
}

func (o JobRecurrenceSchedulePtrOutput) Hours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *JobRecurrenceSchedule) []int {
		if v == nil {
			return nil
		}
		return v.Hours
	}).(pulumi.IntArrayOutput)
}

func (o JobRecurrenceSchedulePtrOutput) Minutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *JobRecurrenceSchedule) []int {
		if v == nil {
			return nil
		}
		return v.Minutes
	}).(pulumi.IntArrayOutput)
}

func (o JobRecurrenceSchedulePtrOutput) MonthDays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *JobRecurrenceSchedule) []int {
		if v == nil {
			return nil
		}
		return v.MonthDays
	}).(pulumi.IntArrayOutput)
}

func (o JobRecurrenceSchedulePtrOutput) MonthlyOccurrences() JobRecurrenceScheduleMonthlyOccurrenceArrayOutput {
	return o.ApplyT(func(v *JobRecurrenceSchedule) []JobRecurrenceScheduleMonthlyOccurrence {
		if v == nil {
			return nil
		}
		return v.MonthlyOccurrences
	}).(JobRecurrenceScheduleMonthlyOccurrenceArrayOutput)
}

func (o JobRecurrenceSchedulePtrOutput) WeekDays() DayOfWeekArrayOutput {
	return o.ApplyT(func(v *JobRecurrenceSchedule) []DayOfWeek {
		if v == nil {
			return nil
		}
		return v.WeekDays
	}).(DayOfWeekArrayOutput)
}

type JobRecurrenceScheduleMonthlyOccurrence struct {
	Day        *JobScheduleDay `pulumi:"day"`
	Occurrence *int            `pulumi:"occurrence"`
}





type JobRecurrenceScheduleMonthlyOccurrenceInput interface {
	pulumi.Input

	ToJobRecurrenceScheduleMonthlyOccurrenceOutput() JobRecurrenceScheduleMonthlyOccurrenceOutput
	ToJobRecurrenceScheduleMonthlyOccurrenceOutputWithContext(context.Context) JobRecurrenceScheduleMonthlyOccurrenceOutput
}

type JobRecurrenceScheduleMonthlyOccurrenceArgs struct {
	Day        JobScheduleDayPtrInput `pulumi:"day"`
	Occurrence pulumi.IntPtrInput     `pulumi:"occurrence"`
}

func (JobRecurrenceScheduleMonthlyOccurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobRecurrenceScheduleMonthlyOccurrence)(nil)).Elem()
}

func (i JobRecurrenceScheduleMonthlyOccurrenceArgs) ToJobRecurrenceScheduleMonthlyOccurrenceOutput() JobRecurrenceScheduleMonthlyOccurrenceOutput {
	return i.ToJobRecurrenceScheduleMonthlyOccurrenceOutputWithContext(context.Background())
}

func (i JobRecurrenceScheduleMonthlyOccurrenceArgs) ToJobRecurrenceScheduleMonthlyOccurrenceOutputWithContext(ctx context.Context) JobRecurrenceScheduleMonthlyOccurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceScheduleMonthlyOccurrenceOutput)
}





type JobRecurrenceScheduleMonthlyOccurrenceArrayInput interface {
	pulumi.Input

	ToJobRecurrenceScheduleMonthlyOccurrenceArrayOutput() JobRecurrenceScheduleMonthlyOccurrenceArrayOutput
	ToJobRecurrenceScheduleMonthlyOccurrenceArrayOutputWithContext(context.Context) JobRecurrenceScheduleMonthlyOccurrenceArrayOutput
}

type JobRecurrenceScheduleMonthlyOccurrenceArray []JobRecurrenceScheduleMonthlyOccurrenceInput

func (JobRecurrenceScheduleMonthlyOccurrenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobRecurrenceScheduleMonthlyOccurrence)(nil)).Elem()
}

func (i JobRecurrenceScheduleMonthlyOccurrenceArray) ToJobRecurrenceScheduleMonthlyOccurrenceArrayOutput() JobRecurrenceScheduleMonthlyOccurrenceArrayOutput {
	return i.ToJobRecurrenceScheduleMonthlyOccurrenceArrayOutputWithContext(context.Background())
}

func (i JobRecurrenceScheduleMonthlyOccurrenceArray) ToJobRecurrenceScheduleMonthlyOccurrenceArrayOutputWithContext(ctx context.Context) JobRecurrenceScheduleMonthlyOccurrenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceScheduleMonthlyOccurrenceArrayOutput)
}

type JobRecurrenceScheduleMonthlyOccurrenceOutput struct{ *pulumi.OutputState }

func (JobRecurrenceScheduleMonthlyOccurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobRecurrenceScheduleMonthlyOccurrence)(nil)).Elem()
}

func (o JobRecurrenceScheduleMonthlyOccurrenceOutput) ToJobRecurrenceScheduleMonthlyOccurrenceOutput() JobRecurrenceScheduleMonthlyOccurrenceOutput {
	return o
}

func (o JobRecurrenceScheduleMonthlyOccurrenceOutput) ToJobRecurrenceScheduleMonthlyOccurrenceOutputWithContext(ctx context.Context) JobRecurrenceScheduleMonthlyOccurrenceOutput {
	return o
}

func (o JobRecurrenceScheduleMonthlyOccurrenceOutput) Day() JobScheduleDayPtrOutput {
	return o.ApplyT(func(v JobRecurrenceScheduleMonthlyOccurrence) *JobScheduleDay { return v.Day }).(JobScheduleDayPtrOutput)
}

func (o JobRecurrenceScheduleMonthlyOccurrenceOutput) Occurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobRecurrenceScheduleMonthlyOccurrence) *int { return v.Occurrence }).(pulumi.IntPtrOutput)
}

type JobRecurrenceScheduleMonthlyOccurrenceArrayOutput struct{ *pulumi.OutputState }

func (JobRecurrenceScheduleMonthlyOccurrenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobRecurrenceScheduleMonthlyOccurrence)(nil)).Elem()
}

func (o JobRecurrenceScheduleMonthlyOccurrenceArrayOutput) ToJobRecurrenceScheduleMonthlyOccurrenceArrayOutput() JobRecurrenceScheduleMonthlyOccurrenceArrayOutput {
	return o
}

func (o JobRecurrenceScheduleMonthlyOccurrenceArrayOutput) ToJobRecurrenceScheduleMonthlyOccurrenceArrayOutputWithContext(ctx context.Context) JobRecurrenceScheduleMonthlyOccurrenceArrayOutput {
	return o
}

func (o JobRecurrenceScheduleMonthlyOccurrenceArrayOutput) Index(i pulumi.IntInput) JobRecurrenceScheduleMonthlyOccurrenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobRecurrenceScheduleMonthlyOccurrence {
		return vs[0].([]JobRecurrenceScheduleMonthlyOccurrence)[vs[1].(int)]
	}).(JobRecurrenceScheduleMonthlyOccurrenceOutput)
}

type JobRecurrenceScheduleMonthlyOccurrenceResponse struct {
	Day        *string `pulumi:"day"`
	Occurrence *int    `pulumi:"occurrence"`
}





type JobRecurrenceScheduleMonthlyOccurrenceResponseInput interface {
	pulumi.Input

	ToJobRecurrenceScheduleMonthlyOccurrenceResponseOutput() JobRecurrenceScheduleMonthlyOccurrenceResponseOutput
	ToJobRecurrenceScheduleMonthlyOccurrenceResponseOutputWithContext(context.Context) JobRecurrenceScheduleMonthlyOccurrenceResponseOutput
}

type JobRecurrenceScheduleMonthlyOccurrenceResponseArgs struct {
	Day        pulumi.StringPtrInput `pulumi:"day"`
	Occurrence pulumi.IntPtrInput    `pulumi:"occurrence"`
}

func (JobRecurrenceScheduleMonthlyOccurrenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobRecurrenceScheduleMonthlyOccurrenceResponse)(nil)).Elem()
}

func (i JobRecurrenceScheduleMonthlyOccurrenceResponseArgs) ToJobRecurrenceScheduleMonthlyOccurrenceResponseOutput() JobRecurrenceScheduleMonthlyOccurrenceResponseOutput {
	return i.ToJobRecurrenceScheduleMonthlyOccurrenceResponseOutputWithContext(context.Background())
}

func (i JobRecurrenceScheduleMonthlyOccurrenceResponseArgs) ToJobRecurrenceScheduleMonthlyOccurrenceResponseOutputWithContext(ctx context.Context) JobRecurrenceScheduleMonthlyOccurrenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceScheduleMonthlyOccurrenceResponseOutput)
}





type JobRecurrenceScheduleMonthlyOccurrenceResponseArrayInput interface {
	pulumi.Input

	ToJobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput() JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput
	ToJobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutputWithContext(context.Context) JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput
}

type JobRecurrenceScheduleMonthlyOccurrenceResponseArray []JobRecurrenceScheduleMonthlyOccurrenceResponseInput

func (JobRecurrenceScheduleMonthlyOccurrenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobRecurrenceScheduleMonthlyOccurrenceResponse)(nil)).Elem()
}

func (i JobRecurrenceScheduleMonthlyOccurrenceResponseArray) ToJobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput() JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput {
	return i.ToJobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutputWithContext(context.Background())
}

func (i JobRecurrenceScheduleMonthlyOccurrenceResponseArray) ToJobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutputWithContext(ctx context.Context) JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput)
}

type JobRecurrenceScheduleMonthlyOccurrenceResponseOutput struct{ *pulumi.OutputState }

func (JobRecurrenceScheduleMonthlyOccurrenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobRecurrenceScheduleMonthlyOccurrenceResponse)(nil)).Elem()
}

func (o JobRecurrenceScheduleMonthlyOccurrenceResponseOutput) ToJobRecurrenceScheduleMonthlyOccurrenceResponseOutput() JobRecurrenceScheduleMonthlyOccurrenceResponseOutput {
	return o
}

func (o JobRecurrenceScheduleMonthlyOccurrenceResponseOutput) ToJobRecurrenceScheduleMonthlyOccurrenceResponseOutputWithContext(ctx context.Context) JobRecurrenceScheduleMonthlyOccurrenceResponseOutput {
	return o
}

func (o JobRecurrenceScheduleMonthlyOccurrenceResponseOutput) Day() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobRecurrenceScheduleMonthlyOccurrenceResponse) *string { return v.Day }).(pulumi.StringPtrOutput)
}

func (o JobRecurrenceScheduleMonthlyOccurrenceResponseOutput) Occurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobRecurrenceScheduleMonthlyOccurrenceResponse) *int { return v.Occurrence }).(pulumi.IntPtrOutput)
}

type JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput struct{ *pulumi.OutputState }

func (JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobRecurrenceScheduleMonthlyOccurrenceResponse)(nil)).Elem()
}

func (o JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput) ToJobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput() JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput {
	return o
}

func (o JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput) ToJobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutputWithContext(ctx context.Context) JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput {
	return o
}

func (o JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput) Index(i pulumi.IntInput) JobRecurrenceScheduleMonthlyOccurrenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobRecurrenceScheduleMonthlyOccurrenceResponse {
		return vs[0].([]JobRecurrenceScheduleMonthlyOccurrenceResponse)[vs[1].(int)]
	}).(JobRecurrenceScheduleMonthlyOccurrenceResponseOutput)
}

type JobRecurrenceScheduleResponse struct {
	Hours              []int                                            `pulumi:"hours"`
	Minutes            []int                                            `pulumi:"minutes"`
	MonthDays          []int                                            `pulumi:"monthDays"`
	MonthlyOccurrences []JobRecurrenceScheduleMonthlyOccurrenceResponse `pulumi:"monthlyOccurrences"`
	WeekDays           []string                                         `pulumi:"weekDays"`
}





type JobRecurrenceScheduleResponseInput interface {
	pulumi.Input

	ToJobRecurrenceScheduleResponseOutput() JobRecurrenceScheduleResponseOutput
	ToJobRecurrenceScheduleResponseOutputWithContext(context.Context) JobRecurrenceScheduleResponseOutput
}

type JobRecurrenceScheduleResponseArgs struct {
	Hours              pulumi.IntArrayInput                                     `pulumi:"hours"`
	Minutes            pulumi.IntArrayInput                                     `pulumi:"minutes"`
	MonthDays          pulumi.IntArrayInput                                     `pulumi:"monthDays"`
	MonthlyOccurrences JobRecurrenceScheduleMonthlyOccurrenceResponseArrayInput `pulumi:"monthlyOccurrences"`
	WeekDays           pulumi.StringArrayInput                                  `pulumi:"weekDays"`
}

func (JobRecurrenceScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobRecurrenceScheduleResponse)(nil)).Elem()
}

func (i JobRecurrenceScheduleResponseArgs) ToJobRecurrenceScheduleResponseOutput() JobRecurrenceScheduleResponseOutput {
	return i.ToJobRecurrenceScheduleResponseOutputWithContext(context.Background())
}

func (i JobRecurrenceScheduleResponseArgs) ToJobRecurrenceScheduleResponseOutputWithContext(ctx context.Context) JobRecurrenceScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceScheduleResponseOutput)
}

func (i JobRecurrenceScheduleResponseArgs) ToJobRecurrenceScheduleResponsePtrOutput() JobRecurrenceScheduleResponsePtrOutput {
	return i.ToJobRecurrenceScheduleResponsePtrOutputWithContext(context.Background())
}

func (i JobRecurrenceScheduleResponseArgs) ToJobRecurrenceScheduleResponsePtrOutputWithContext(ctx context.Context) JobRecurrenceScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceScheduleResponseOutput).ToJobRecurrenceScheduleResponsePtrOutputWithContext(ctx)
}









type JobRecurrenceScheduleResponsePtrInput interface {
	pulumi.Input

	ToJobRecurrenceScheduleResponsePtrOutput() JobRecurrenceScheduleResponsePtrOutput
	ToJobRecurrenceScheduleResponsePtrOutputWithContext(context.Context) JobRecurrenceScheduleResponsePtrOutput
}

type jobRecurrenceScheduleResponsePtrType JobRecurrenceScheduleResponseArgs

func JobRecurrenceScheduleResponsePtr(v *JobRecurrenceScheduleResponseArgs) JobRecurrenceScheduleResponsePtrInput {
	return (*jobRecurrenceScheduleResponsePtrType)(v)
}

func (*jobRecurrenceScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobRecurrenceScheduleResponse)(nil)).Elem()
}

func (i *jobRecurrenceScheduleResponsePtrType) ToJobRecurrenceScheduleResponsePtrOutput() JobRecurrenceScheduleResponsePtrOutput {
	return i.ToJobRecurrenceScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *jobRecurrenceScheduleResponsePtrType) ToJobRecurrenceScheduleResponsePtrOutputWithContext(ctx context.Context) JobRecurrenceScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobRecurrenceScheduleResponsePtrOutput)
}

type JobRecurrenceScheduleResponseOutput struct{ *pulumi.OutputState }

func (JobRecurrenceScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobRecurrenceScheduleResponse)(nil)).Elem()
}

func (o JobRecurrenceScheduleResponseOutput) ToJobRecurrenceScheduleResponseOutput() JobRecurrenceScheduleResponseOutput {
	return o
}

func (o JobRecurrenceScheduleResponseOutput) ToJobRecurrenceScheduleResponseOutputWithContext(ctx context.Context) JobRecurrenceScheduleResponseOutput {
	return o
}

func (o JobRecurrenceScheduleResponseOutput) ToJobRecurrenceScheduleResponsePtrOutput() JobRecurrenceScheduleResponsePtrOutput {
	return o.ToJobRecurrenceScheduleResponsePtrOutputWithContext(context.Background())
}

func (o JobRecurrenceScheduleResponseOutput) ToJobRecurrenceScheduleResponsePtrOutputWithContext(ctx context.Context) JobRecurrenceScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobRecurrenceScheduleResponse) *JobRecurrenceScheduleResponse {
		return &v
	}).(JobRecurrenceScheduleResponsePtrOutput)
}

func (o JobRecurrenceScheduleResponseOutput) Hours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v JobRecurrenceScheduleResponse) []int { return v.Hours }).(pulumi.IntArrayOutput)
}

func (o JobRecurrenceScheduleResponseOutput) Minutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v JobRecurrenceScheduleResponse) []int { return v.Minutes }).(pulumi.IntArrayOutput)
}

func (o JobRecurrenceScheduleResponseOutput) MonthDays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v JobRecurrenceScheduleResponse) []int { return v.MonthDays }).(pulumi.IntArrayOutput)
}

func (o JobRecurrenceScheduleResponseOutput) MonthlyOccurrences() JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput {
	return o.ApplyT(func(v JobRecurrenceScheduleResponse) []JobRecurrenceScheduleMonthlyOccurrenceResponse {
		return v.MonthlyOccurrences
	}).(JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput)
}

func (o JobRecurrenceScheduleResponseOutput) WeekDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JobRecurrenceScheduleResponse) []string { return v.WeekDays }).(pulumi.StringArrayOutput)
}

type JobRecurrenceScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (JobRecurrenceScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobRecurrenceScheduleResponse)(nil)).Elem()
}

func (o JobRecurrenceScheduleResponsePtrOutput) ToJobRecurrenceScheduleResponsePtrOutput() JobRecurrenceScheduleResponsePtrOutput {
	return o
}

func (o JobRecurrenceScheduleResponsePtrOutput) ToJobRecurrenceScheduleResponsePtrOutputWithContext(ctx context.Context) JobRecurrenceScheduleResponsePtrOutput {
	return o
}

func (o JobRecurrenceScheduleResponsePtrOutput) Elem() JobRecurrenceScheduleResponseOutput {
	return o.ApplyT(func(v *JobRecurrenceScheduleResponse) JobRecurrenceScheduleResponse {
		if v != nil {
			return *v
		}
		var ret JobRecurrenceScheduleResponse
		return ret
	}).(JobRecurrenceScheduleResponseOutput)
}

func (o JobRecurrenceScheduleResponsePtrOutput) Hours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *JobRecurrenceScheduleResponse) []int {
		if v == nil {
			return nil
		}
		return v.Hours
	}).(pulumi.IntArrayOutput)
}

func (o JobRecurrenceScheduleResponsePtrOutput) Minutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *JobRecurrenceScheduleResponse) []int {
		if v == nil {
			return nil
		}
		return v.Minutes
	}).(pulumi.IntArrayOutput)
}

func (o JobRecurrenceScheduleResponsePtrOutput) MonthDays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *JobRecurrenceScheduleResponse) []int {
		if v == nil {
			return nil
		}
		return v.MonthDays
	}).(pulumi.IntArrayOutput)
}

func (o JobRecurrenceScheduleResponsePtrOutput) MonthlyOccurrences() JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput {
	return o.ApplyT(func(v *JobRecurrenceScheduleResponse) []JobRecurrenceScheduleMonthlyOccurrenceResponse {
		if v == nil {
			return nil
		}
		return v.MonthlyOccurrences
	}).(JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput)
}

func (o JobRecurrenceScheduleResponsePtrOutput) WeekDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JobRecurrenceScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.WeekDays
	}).(pulumi.StringArrayOutput)
}

type JobStatusResponse struct {
	ExecutionCount    int    `pulumi:"executionCount"`
	FailureCount      int    `pulumi:"failureCount"`
	FaultedCount      int    `pulumi:"faultedCount"`
	LastExecutionTime string `pulumi:"lastExecutionTime"`
	NextExecutionTime string `pulumi:"nextExecutionTime"`
}





type JobStatusResponseInput interface {
	pulumi.Input

	ToJobStatusResponseOutput() JobStatusResponseOutput
	ToJobStatusResponseOutputWithContext(context.Context) JobStatusResponseOutput
}

type JobStatusResponseArgs struct {
	ExecutionCount    pulumi.IntInput    `pulumi:"executionCount"`
	FailureCount      pulumi.IntInput    `pulumi:"failureCount"`
	FaultedCount      pulumi.IntInput    `pulumi:"faultedCount"`
	LastExecutionTime pulumi.StringInput `pulumi:"lastExecutionTime"`
	NextExecutionTime pulumi.StringInput `pulumi:"nextExecutionTime"`
}

func (JobStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStatusResponse)(nil)).Elem()
}

func (i JobStatusResponseArgs) ToJobStatusResponseOutput() JobStatusResponseOutput {
	return i.ToJobStatusResponseOutputWithContext(context.Background())
}

func (i JobStatusResponseArgs) ToJobStatusResponseOutputWithContext(ctx context.Context) JobStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStatusResponseOutput)
}

func (i JobStatusResponseArgs) ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput {
	return i.ToJobStatusResponsePtrOutputWithContext(context.Background())
}

func (i JobStatusResponseArgs) ToJobStatusResponsePtrOutputWithContext(ctx context.Context) JobStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStatusResponseOutput).ToJobStatusResponsePtrOutputWithContext(ctx)
}









type JobStatusResponsePtrInput interface {
	pulumi.Input

	ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput
	ToJobStatusResponsePtrOutputWithContext(context.Context) JobStatusResponsePtrOutput
}

type jobStatusResponsePtrType JobStatusResponseArgs

func JobStatusResponsePtr(v *JobStatusResponseArgs) JobStatusResponsePtrInput {
	return (*jobStatusResponsePtrType)(v)
}

func (*jobStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStatusResponse)(nil)).Elem()
}

func (i *jobStatusResponsePtrType) ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput {
	return i.ToJobStatusResponsePtrOutputWithContext(context.Background())
}

func (i *jobStatusResponsePtrType) ToJobStatusResponsePtrOutputWithContext(ctx context.Context) JobStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStatusResponsePtrOutput)
}

type JobStatusResponseOutput struct{ *pulumi.OutputState }

func (JobStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStatusResponse)(nil)).Elem()
}

func (o JobStatusResponseOutput) ToJobStatusResponseOutput() JobStatusResponseOutput {
	return o
}

func (o JobStatusResponseOutput) ToJobStatusResponseOutputWithContext(ctx context.Context) JobStatusResponseOutput {
	return o
}

func (o JobStatusResponseOutput) ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput {
	return o.ToJobStatusResponsePtrOutputWithContext(context.Background())
}

func (o JobStatusResponseOutput) ToJobStatusResponsePtrOutputWithContext(ctx context.Context) JobStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStatusResponse) *JobStatusResponse {
		return &v
	}).(JobStatusResponsePtrOutput)
}

func (o JobStatusResponseOutput) ExecutionCount() pulumi.IntOutput {
	return o.ApplyT(func(v JobStatusResponse) int { return v.ExecutionCount }).(pulumi.IntOutput)
}

func (o JobStatusResponseOutput) FailureCount() pulumi.IntOutput {
	return o.ApplyT(func(v JobStatusResponse) int { return v.FailureCount }).(pulumi.IntOutput)
}

func (o JobStatusResponseOutput) FaultedCount() pulumi.IntOutput {
	return o.ApplyT(func(v JobStatusResponse) int { return v.FaultedCount }).(pulumi.IntOutput)
}

func (o JobStatusResponseOutput) LastExecutionTime() pulumi.StringOutput {
	return o.ApplyT(func(v JobStatusResponse) string { return v.LastExecutionTime }).(pulumi.StringOutput)
}

func (o JobStatusResponseOutput) NextExecutionTime() pulumi.StringOutput {
	return o.ApplyT(func(v JobStatusResponse) string { return v.NextExecutionTime }).(pulumi.StringOutput)
}

type JobStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (JobStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStatusResponse)(nil)).Elem()
}

func (o JobStatusResponsePtrOutput) ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput {
	return o
}

func (o JobStatusResponsePtrOutput) ToJobStatusResponsePtrOutputWithContext(ctx context.Context) JobStatusResponsePtrOutput {
	return o
}

func (o JobStatusResponsePtrOutput) Elem() JobStatusResponseOutput {
	return o.ApplyT(func(v *JobStatusResponse) JobStatusResponse {
		if v != nil {
			return *v
		}
		var ret JobStatusResponse
		return ret
	}).(JobStatusResponseOutput)
}

func (o JobStatusResponsePtrOutput) ExecutionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStatusResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ExecutionCount
	}).(pulumi.IntPtrOutput)
}

func (o JobStatusResponsePtrOutput) FailureCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStatusResponse) *int {
		if v == nil {
			return nil
		}
		return &v.FailureCount
	}).(pulumi.IntPtrOutput)
}

func (o JobStatusResponsePtrOutput) FaultedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStatusResponse) *int {
		if v == nil {
			return nil
		}
		return &v.FaultedCount
	}).(pulumi.IntPtrOutput)
}

func (o JobStatusResponsePtrOutput) LastExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastExecutionTime
	}).(pulumi.StringPtrOutput)
}

func (o JobStatusResponsePtrOutput) NextExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.NextExecutionTime
	}).(pulumi.StringPtrOutput)
}

type RetryPolicy struct {
	RetryCount    *int       `pulumi:"retryCount"`
	RetryInterval *string    `pulumi:"retryInterval"`
	RetryType     *RetryType `pulumi:"retryType"`
}





type RetryPolicyInput interface {
	pulumi.Input

	ToRetryPolicyOutput() RetryPolicyOutput
	ToRetryPolicyOutputWithContext(context.Context) RetryPolicyOutput
}

type RetryPolicyArgs struct {
	RetryCount    pulumi.IntPtrInput    `pulumi:"retryCount"`
	RetryInterval pulumi.StringPtrInput `pulumi:"retryInterval"`
	RetryType     RetryTypePtrInput     `pulumi:"retryType"`
}

func (RetryPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetryPolicy)(nil)).Elem()
}

func (i RetryPolicyArgs) ToRetryPolicyOutput() RetryPolicyOutput {
	return i.ToRetryPolicyOutputWithContext(context.Background())
}

func (i RetryPolicyArgs) ToRetryPolicyOutputWithContext(ctx context.Context) RetryPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyOutput)
}

func (i RetryPolicyArgs) ToRetryPolicyPtrOutput() RetryPolicyPtrOutput {
	return i.ToRetryPolicyPtrOutputWithContext(context.Background())
}

func (i RetryPolicyArgs) ToRetryPolicyPtrOutputWithContext(ctx context.Context) RetryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyOutput).ToRetryPolicyPtrOutputWithContext(ctx)
}









type RetryPolicyPtrInput interface {
	pulumi.Input

	ToRetryPolicyPtrOutput() RetryPolicyPtrOutput
	ToRetryPolicyPtrOutputWithContext(context.Context) RetryPolicyPtrOutput
}

type retryPolicyPtrType RetryPolicyArgs

func RetryPolicyPtr(v *RetryPolicyArgs) RetryPolicyPtrInput {
	return (*retryPolicyPtrType)(v)
}

func (*retryPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetryPolicy)(nil)).Elem()
}

func (i *retryPolicyPtrType) ToRetryPolicyPtrOutput() RetryPolicyPtrOutput {
	return i.ToRetryPolicyPtrOutputWithContext(context.Background())
}

func (i *retryPolicyPtrType) ToRetryPolicyPtrOutputWithContext(ctx context.Context) RetryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyPtrOutput)
}

type RetryPolicyOutput struct{ *pulumi.OutputState }

func (RetryPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetryPolicy)(nil)).Elem()
}

func (o RetryPolicyOutput) ToRetryPolicyOutput() RetryPolicyOutput {
	return o
}

func (o RetryPolicyOutput) ToRetryPolicyOutputWithContext(ctx context.Context) RetryPolicyOutput {
	return o
}

func (o RetryPolicyOutput) ToRetryPolicyPtrOutput() RetryPolicyPtrOutput {
	return o.ToRetryPolicyPtrOutputWithContext(context.Background())
}

func (o RetryPolicyOutput) ToRetryPolicyPtrOutputWithContext(ctx context.Context) RetryPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetryPolicy) *RetryPolicy {
		return &v
	}).(RetryPolicyPtrOutput)
}

func (o RetryPolicyOutput) RetryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetryPolicy) *int { return v.RetryCount }).(pulumi.IntPtrOutput)
}

func (o RetryPolicyOutput) RetryInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetryPolicy) *string { return v.RetryInterval }).(pulumi.StringPtrOutput)
}

func (o RetryPolicyOutput) RetryType() RetryTypePtrOutput {
	return o.ApplyT(func(v RetryPolicy) *RetryType { return v.RetryType }).(RetryTypePtrOutput)
}

type RetryPolicyPtrOutput struct{ *pulumi.OutputState }

func (RetryPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetryPolicy)(nil)).Elem()
}

func (o RetryPolicyPtrOutput) ToRetryPolicyPtrOutput() RetryPolicyPtrOutput {
	return o
}

func (o RetryPolicyPtrOutput) ToRetryPolicyPtrOutputWithContext(ctx context.Context) RetryPolicyPtrOutput {
	return o
}

func (o RetryPolicyPtrOutput) Elem() RetryPolicyOutput {
	return o.ApplyT(func(v *RetryPolicy) RetryPolicy {
		if v != nil {
			return *v
		}
		var ret RetryPolicy
		return ret
	}).(RetryPolicyOutput)
}

func (o RetryPolicyPtrOutput) RetryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetryPolicy) *int {
		if v == nil {
			return nil
		}
		return v.RetryCount
	}).(pulumi.IntPtrOutput)
}

func (o RetryPolicyPtrOutput) RetryInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetryPolicy) *string {
		if v == nil {
			return nil
		}
		return v.RetryInterval
	}).(pulumi.StringPtrOutput)
}

func (o RetryPolicyPtrOutput) RetryType() RetryTypePtrOutput {
	return o.ApplyT(func(v *RetryPolicy) *RetryType {
		if v == nil {
			return nil
		}
		return v.RetryType
	}).(RetryTypePtrOutput)
}

type RetryPolicyResponse struct {
	RetryCount    *int    `pulumi:"retryCount"`
	RetryInterval *string `pulumi:"retryInterval"`
	RetryType     *string `pulumi:"retryType"`
}





type RetryPolicyResponseInput interface {
	pulumi.Input

	ToRetryPolicyResponseOutput() RetryPolicyResponseOutput
	ToRetryPolicyResponseOutputWithContext(context.Context) RetryPolicyResponseOutput
}

type RetryPolicyResponseArgs struct {
	RetryCount    pulumi.IntPtrInput    `pulumi:"retryCount"`
	RetryInterval pulumi.StringPtrInput `pulumi:"retryInterval"`
	RetryType     pulumi.StringPtrInput `pulumi:"retryType"`
}

func (RetryPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetryPolicyResponse)(nil)).Elem()
}

func (i RetryPolicyResponseArgs) ToRetryPolicyResponseOutput() RetryPolicyResponseOutput {
	return i.ToRetryPolicyResponseOutputWithContext(context.Background())
}

func (i RetryPolicyResponseArgs) ToRetryPolicyResponseOutputWithContext(ctx context.Context) RetryPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyResponseOutput)
}

func (i RetryPolicyResponseArgs) ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput {
	return i.ToRetryPolicyResponsePtrOutputWithContext(context.Background())
}

func (i RetryPolicyResponseArgs) ToRetryPolicyResponsePtrOutputWithContext(ctx context.Context) RetryPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyResponseOutput).ToRetryPolicyResponsePtrOutputWithContext(ctx)
}









type RetryPolicyResponsePtrInput interface {
	pulumi.Input

	ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput
	ToRetryPolicyResponsePtrOutputWithContext(context.Context) RetryPolicyResponsePtrOutput
}

type retryPolicyResponsePtrType RetryPolicyResponseArgs

func RetryPolicyResponsePtr(v *RetryPolicyResponseArgs) RetryPolicyResponsePtrInput {
	return (*retryPolicyResponsePtrType)(v)
}

func (*retryPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetryPolicyResponse)(nil)).Elem()
}

func (i *retryPolicyResponsePtrType) ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput {
	return i.ToRetryPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *retryPolicyResponsePtrType) ToRetryPolicyResponsePtrOutputWithContext(ctx context.Context) RetryPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyResponsePtrOutput)
}

type RetryPolicyResponseOutput struct{ *pulumi.OutputState }

func (RetryPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetryPolicyResponse)(nil)).Elem()
}

func (o RetryPolicyResponseOutput) ToRetryPolicyResponseOutput() RetryPolicyResponseOutput {
	return o
}

func (o RetryPolicyResponseOutput) ToRetryPolicyResponseOutputWithContext(ctx context.Context) RetryPolicyResponseOutput {
	return o
}

func (o RetryPolicyResponseOutput) ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput {
	return o.ToRetryPolicyResponsePtrOutputWithContext(context.Background())
}

func (o RetryPolicyResponseOutput) ToRetryPolicyResponsePtrOutputWithContext(ctx context.Context) RetryPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetryPolicyResponse) *RetryPolicyResponse {
		return &v
	}).(RetryPolicyResponsePtrOutput)
}

func (o RetryPolicyResponseOutput) RetryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetryPolicyResponse) *int { return v.RetryCount }).(pulumi.IntPtrOutput)
}

func (o RetryPolicyResponseOutput) RetryInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetryPolicyResponse) *string { return v.RetryInterval }).(pulumi.StringPtrOutput)
}

func (o RetryPolicyResponseOutput) RetryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetryPolicyResponse) *string { return v.RetryType }).(pulumi.StringPtrOutput)
}

type RetryPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RetryPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetryPolicyResponse)(nil)).Elem()
}

func (o RetryPolicyResponsePtrOutput) ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput {
	return o
}

func (o RetryPolicyResponsePtrOutput) ToRetryPolicyResponsePtrOutputWithContext(ctx context.Context) RetryPolicyResponsePtrOutput {
	return o
}

func (o RetryPolicyResponsePtrOutput) Elem() RetryPolicyResponseOutput {
	return o.ApplyT(func(v *RetryPolicyResponse) RetryPolicyResponse {
		if v != nil {
			return *v
		}
		var ret RetryPolicyResponse
		return ret
	}).(RetryPolicyResponseOutput)
}

func (o RetryPolicyResponsePtrOutput) RetryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetryPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetryCount
	}).(pulumi.IntPtrOutput)
}

func (o RetryPolicyResponsePtrOutput) RetryInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetryPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.RetryInterval
	}).(pulumi.StringPtrOutput)
}

func (o RetryPolicyResponsePtrOutput) RetryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetryPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.RetryType
	}).(pulumi.StringPtrOutput)
}

type ServiceBusAuthentication struct {
	SasKey     *string                       `pulumi:"sasKey"`
	SasKeyName *string                       `pulumi:"sasKeyName"`
	Type       *ServiceBusAuthenticationType `pulumi:"type"`
}





type ServiceBusAuthenticationInput interface {
	pulumi.Input

	ToServiceBusAuthenticationOutput() ServiceBusAuthenticationOutput
	ToServiceBusAuthenticationOutputWithContext(context.Context) ServiceBusAuthenticationOutput
}

type ServiceBusAuthenticationArgs struct {
	SasKey     pulumi.StringPtrInput                `pulumi:"sasKey"`
	SasKeyName pulumi.StringPtrInput                `pulumi:"sasKeyName"`
	Type       ServiceBusAuthenticationTypePtrInput `pulumi:"type"`
}

func (ServiceBusAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusAuthentication)(nil)).Elem()
}

func (i ServiceBusAuthenticationArgs) ToServiceBusAuthenticationOutput() ServiceBusAuthenticationOutput {
	return i.ToServiceBusAuthenticationOutputWithContext(context.Background())
}

func (i ServiceBusAuthenticationArgs) ToServiceBusAuthenticationOutputWithContext(ctx context.Context) ServiceBusAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusAuthenticationOutput)
}

func (i ServiceBusAuthenticationArgs) ToServiceBusAuthenticationPtrOutput() ServiceBusAuthenticationPtrOutput {
	return i.ToServiceBusAuthenticationPtrOutputWithContext(context.Background())
}

func (i ServiceBusAuthenticationArgs) ToServiceBusAuthenticationPtrOutputWithContext(ctx context.Context) ServiceBusAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusAuthenticationOutput).ToServiceBusAuthenticationPtrOutputWithContext(ctx)
}









type ServiceBusAuthenticationPtrInput interface {
	pulumi.Input

	ToServiceBusAuthenticationPtrOutput() ServiceBusAuthenticationPtrOutput
	ToServiceBusAuthenticationPtrOutputWithContext(context.Context) ServiceBusAuthenticationPtrOutput
}

type serviceBusAuthenticationPtrType ServiceBusAuthenticationArgs

func ServiceBusAuthenticationPtr(v *ServiceBusAuthenticationArgs) ServiceBusAuthenticationPtrInput {
	return (*serviceBusAuthenticationPtrType)(v)
}

func (*serviceBusAuthenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusAuthentication)(nil)).Elem()
}

func (i *serviceBusAuthenticationPtrType) ToServiceBusAuthenticationPtrOutput() ServiceBusAuthenticationPtrOutput {
	return i.ToServiceBusAuthenticationPtrOutputWithContext(context.Background())
}

func (i *serviceBusAuthenticationPtrType) ToServiceBusAuthenticationPtrOutputWithContext(ctx context.Context) ServiceBusAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusAuthenticationPtrOutput)
}

type ServiceBusAuthenticationOutput struct{ *pulumi.OutputState }

func (ServiceBusAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusAuthentication)(nil)).Elem()
}

func (o ServiceBusAuthenticationOutput) ToServiceBusAuthenticationOutput() ServiceBusAuthenticationOutput {
	return o
}

func (o ServiceBusAuthenticationOutput) ToServiceBusAuthenticationOutputWithContext(ctx context.Context) ServiceBusAuthenticationOutput {
	return o
}

func (o ServiceBusAuthenticationOutput) ToServiceBusAuthenticationPtrOutput() ServiceBusAuthenticationPtrOutput {
	return o.ToServiceBusAuthenticationPtrOutputWithContext(context.Background())
}

func (o ServiceBusAuthenticationOutput) ToServiceBusAuthenticationPtrOutputWithContext(ctx context.Context) ServiceBusAuthenticationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceBusAuthentication) *ServiceBusAuthentication {
		return &v
	}).(ServiceBusAuthenticationPtrOutput)
}

func (o ServiceBusAuthenticationOutput) SasKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusAuthentication) *string { return v.SasKey }).(pulumi.StringPtrOutput)
}

func (o ServiceBusAuthenticationOutput) SasKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusAuthentication) *string { return v.SasKeyName }).(pulumi.StringPtrOutput)
}

func (o ServiceBusAuthenticationOutput) Type() ServiceBusAuthenticationTypePtrOutput {
	return o.ApplyT(func(v ServiceBusAuthentication) *ServiceBusAuthenticationType { return v.Type }).(ServiceBusAuthenticationTypePtrOutput)
}

type ServiceBusAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (ServiceBusAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusAuthentication)(nil)).Elem()
}

func (o ServiceBusAuthenticationPtrOutput) ToServiceBusAuthenticationPtrOutput() ServiceBusAuthenticationPtrOutput {
	return o
}

func (o ServiceBusAuthenticationPtrOutput) ToServiceBusAuthenticationPtrOutputWithContext(ctx context.Context) ServiceBusAuthenticationPtrOutput {
	return o
}

func (o ServiceBusAuthenticationPtrOutput) Elem() ServiceBusAuthenticationOutput {
	return o.ApplyT(func(v *ServiceBusAuthentication) ServiceBusAuthentication {
		if v != nil {
			return *v
		}
		var ret ServiceBusAuthentication
		return ret
	}).(ServiceBusAuthenticationOutput)
}

func (o ServiceBusAuthenticationPtrOutput) SasKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.SasKey
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusAuthenticationPtrOutput) SasKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.SasKeyName
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusAuthenticationPtrOutput) Type() ServiceBusAuthenticationTypePtrOutput {
	return o.ApplyT(func(v *ServiceBusAuthentication) *ServiceBusAuthenticationType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ServiceBusAuthenticationTypePtrOutput)
}

type ServiceBusAuthenticationResponse struct {
	SasKey     *string `pulumi:"sasKey"`
	SasKeyName *string `pulumi:"sasKeyName"`
	Type       *string `pulumi:"type"`
}





type ServiceBusAuthenticationResponseInput interface {
	pulumi.Input

	ToServiceBusAuthenticationResponseOutput() ServiceBusAuthenticationResponseOutput
	ToServiceBusAuthenticationResponseOutputWithContext(context.Context) ServiceBusAuthenticationResponseOutput
}

type ServiceBusAuthenticationResponseArgs struct {
	SasKey     pulumi.StringPtrInput `pulumi:"sasKey"`
	SasKeyName pulumi.StringPtrInput `pulumi:"sasKeyName"`
	Type       pulumi.StringPtrInput `pulumi:"type"`
}

func (ServiceBusAuthenticationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusAuthenticationResponse)(nil)).Elem()
}

func (i ServiceBusAuthenticationResponseArgs) ToServiceBusAuthenticationResponseOutput() ServiceBusAuthenticationResponseOutput {
	return i.ToServiceBusAuthenticationResponseOutputWithContext(context.Background())
}

func (i ServiceBusAuthenticationResponseArgs) ToServiceBusAuthenticationResponseOutputWithContext(ctx context.Context) ServiceBusAuthenticationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusAuthenticationResponseOutput)
}

func (i ServiceBusAuthenticationResponseArgs) ToServiceBusAuthenticationResponsePtrOutput() ServiceBusAuthenticationResponsePtrOutput {
	return i.ToServiceBusAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (i ServiceBusAuthenticationResponseArgs) ToServiceBusAuthenticationResponsePtrOutputWithContext(ctx context.Context) ServiceBusAuthenticationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusAuthenticationResponseOutput).ToServiceBusAuthenticationResponsePtrOutputWithContext(ctx)
}









type ServiceBusAuthenticationResponsePtrInput interface {
	pulumi.Input

	ToServiceBusAuthenticationResponsePtrOutput() ServiceBusAuthenticationResponsePtrOutput
	ToServiceBusAuthenticationResponsePtrOutputWithContext(context.Context) ServiceBusAuthenticationResponsePtrOutput
}

type serviceBusAuthenticationResponsePtrType ServiceBusAuthenticationResponseArgs

func ServiceBusAuthenticationResponsePtr(v *ServiceBusAuthenticationResponseArgs) ServiceBusAuthenticationResponsePtrInput {
	return (*serviceBusAuthenticationResponsePtrType)(v)
}

func (*serviceBusAuthenticationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusAuthenticationResponse)(nil)).Elem()
}

func (i *serviceBusAuthenticationResponsePtrType) ToServiceBusAuthenticationResponsePtrOutput() ServiceBusAuthenticationResponsePtrOutput {
	return i.ToServiceBusAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (i *serviceBusAuthenticationResponsePtrType) ToServiceBusAuthenticationResponsePtrOutputWithContext(ctx context.Context) ServiceBusAuthenticationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusAuthenticationResponsePtrOutput)
}

type ServiceBusAuthenticationResponseOutput struct{ *pulumi.OutputState }

func (ServiceBusAuthenticationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusAuthenticationResponse)(nil)).Elem()
}

func (o ServiceBusAuthenticationResponseOutput) ToServiceBusAuthenticationResponseOutput() ServiceBusAuthenticationResponseOutput {
	return o
}

func (o ServiceBusAuthenticationResponseOutput) ToServiceBusAuthenticationResponseOutputWithContext(ctx context.Context) ServiceBusAuthenticationResponseOutput {
	return o
}

func (o ServiceBusAuthenticationResponseOutput) ToServiceBusAuthenticationResponsePtrOutput() ServiceBusAuthenticationResponsePtrOutput {
	return o.ToServiceBusAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (o ServiceBusAuthenticationResponseOutput) ToServiceBusAuthenticationResponsePtrOutputWithContext(ctx context.Context) ServiceBusAuthenticationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceBusAuthenticationResponse) *ServiceBusAuthenticationResponse {
		return &v
	}).(ServiceBusAuthenticationResponsePtrOutput)
}

func (o ServiceBusAuthenticationResponseOutput) SasKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusAuthenticationResponse) *string { return v.SasKey }).(pulumi.StringPtrOutput)
}

func (o ServiceBusAuthenticationResponseOutput) SasKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusAuthenticationResponse) *string { return v.SasKeyName }).(pulumi.StringPtrOutput)
}

func (o ServiceBusAuthenticationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusAuthenticationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ServiceBusAuthenticationResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceBusAuthenticationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusAuthenticationResponse)(nil)).Elem()
}

func (o ServiceBusAuthenticationResponsePtrOutput) ToServiceBusAuthenticationResponsePtrOutput() ServiceBusAuthenticationResponsePtrOutput {
	return o
}

func (o ServiceBusAuthenticationResponsePtrOutput) ToServiceBusAuthenticationResponsePtrOutputWithContext(ctx context.Context) ServiceBusAuthenticationResponsePtrOutput {
	return o
}

func (o ServiceBusAuthenticationResponsePtrOutput) Elem() ServiceBusAuthenticationResponseOutput {
	return o.ApplyT(func(v *ServiceBusAuthenticationResponse) ServiceBusAuthenticationResponse {
		if v != nil {
			return *v
		}
		var ret ServiceBusAuthenticationResponse
		return ret
	}).(ServiceBusAuthenticationResponseOutput)
}

func (o ServiceBusAuthenticationResponsePtrOutput) SasKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusAuthenticationResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasKey
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusAuthenticationResponsePtrOutput) SasKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusAuthenticationResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasKeyName
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusAuthenticationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusAuthenticationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ServiceBusBrokeredMessageProperties struct {
	ContentType             *string `pulumi:"contentType"`
	CorrelationId           *string `pulumi:"correlationId"`
	ForcePersistence        *bool   `pulumi:"forcePersistence"`
	Label                   *string `pulumi:"label"`
	MessageId               *string `pulumi:"messageId"`
	PartitionKey            *string `pulumi:"partitionKey"`
	ReplyTo                 *string `pulumi:"replyTo"`
	ReplyToSessionId        *string `pulumi:"replyToSessionId"`
	ScheduledEnqueueTimeUtc *string `pulumi:"scheduledEnqueueTimeUtc"`
	SessionId               *string `pulumi:"sessionId"`
	TimeToLive              *string `pulumi:"timeToLive"`
	To                      *string `pulumi:"to"`
	ViaPartitionKey         *string `pulumi:"viaPartitionKey"`
}





type ServiceBusBrokeredMessagePropertiesInput interface {
	pulumi.Input

	ToServiceBusBrokeredMessagePropertiesOutput() ServiceBusBrokeredMessagePropertiesOutput
	ToServiceBusBrokeredMessagePropertiesOutputWithContext(context.Context) ServiceBusBrokeredMessagePropertiesOutput
}

type ServiceBusBrokeredMessagePropertiesArgs struct {
	ContentType             pulumi.StringPtrInput `pulumi:"contentType"`
	CorrelationId           pulumi.StringPtrInput `pulumi:"correlationId"`
	ForcePersistence        pulumi.BoolPtrInput   `pulumi:"forcePersistence"`
	Label                   pulumi.StringPtrInput `pulumi:"label"`
	MessageId               pulumi.StringPtrInput `pulumi:"messageId"`
	PartitionKey            pulumi.StringPtrInput `pulumi:"partitionKey"`
	ReplyTo                 pulumi.StringPtrInput `pulumi:"replyTo"`
	ReplyToSessionId        pulumi.StringPtrInput `pulumi:"replyToSessionId"`
	ScheduledEnqueueTimeUtc pulumi.StringPtrInput `pulumi:"scheduledEnqueueTimeUtc"`
	SessionId               pulumi.StringPtrInput `pulumi:"sessionId"`
	TimeToLive              pulumi.StringPtrInput `pulumi:"timeToLive"`
	To                      pulumi.StringPtrInput `pulumi:"to"`
	ViaPartitionKey         pulumi.StringPtrInput `pulumi:"viaPartitionKey"`
}

func (ServiceBusBrokeredMessagePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusBrokeredMessageProperties)(nil)).Elem()
}

func (i ServiceBusBrokeredMessagePropertiesArgs) ToServiceBusBrokeredMessagePropertiesOutput() ServiceBusBrokeredMessagePropertiesOutput {
	return i.ToServiceBusBrokeredMessagePropertiesOutputWithContext(context.Background())
}

func (i ServiceBusBrokeredMessagePropertiesArgs) ToServiceBusBrokeredMessagePropertiesOutputWithContext(ctx context.Context) ServiceBusBrokeredMessagePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusBrokeredMessagePropertiesOutput)
}

func (i ServiceBusBrokeredMessagePropertiesArgs) ToServiceBusBrokeredMessagePropertiesPtrOutput() ServiceBusBrokeredMessagePropertiesPtrOutput {
	return i.ToServiceBusBrokeredMessagePropertiesPtrOutputWithContext(context.Background())
}

func (i ServiceBusBrokeredMessagePropertiesArgs) ToServiceBusBrokeredMessagePropertiesPtrOutputWithContext(ctx context.Context) ServiceBusBrokeredMessagePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusBrokeredMessagePropertiesOutput).ToServiceBusBrokeredMessagePropertiesPtrOutputWithContext(ctx)
}









type ServiceBusBrokeredMessagePropertiesPtrInput interface {
	pulumi.Input

	ToServiceBusBrokeredMessagePropertiesPtrOutput() ServiceBusBrokeredMessagePropertiesPtrOutput
	ToServiceBusBrokeredMessagePropertiesPtrOutputWithContext(context.Context) ServiceBusBrokeredMessagePropertiesPtrOutput
}

type serviceBusBrokeredMessagePropertiesPtrType ServiceBusBrokeredMessagePropertiesArgs

func ServiceBusBrokeredMessagePropertiesPtr(v *ServiceBusBrokeredMessagePropertiesArgs) ServiceBusBrokeredMessagePropertiesPtrInput {
	return (*serviceBusBrokeredMessagePropertiesPtrType)(v)
}

func (*serviceBusBrokeredMessagePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusBrokeredMessageProperties)(nil)).Elem()
}

func (i *serviceBusBrokeredMessagePropertiesPtrType) ToServiceBusBrokeredMessagePropertiesPtrOutput() ServiceBusBrokeredMessagePropertiesPtrOutput {
	return i.ToServiceBusBrokeredMessagePropertiesPtrOutputWithContext(context.Background())
}

func (i *serviceBusBrokeredMessagePropertiesPtrType) ToServiceBusBrokeredMessagePropertiesPtrOutputWithContext(ctx context.Context) ServiceBusBrokeredMessagePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusBrokeredMessagePropertiesPtrOutput)
}

type ServiceBusBrokeredMessagePropertiesOutput struct{ *pulumi.OutputState }

func (ServiceBusBrokeredMessagePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusBrokeredMessageProperties)(nil)).Elem()
}

func (o ServiceBusBrokeredMessagePropertiesOutput) ToServiceBusBrokeredMessagePropertiesOutput() ServiceBusBrokeredMessagePropertiesOutput {
	return o
}

func (o ServiceBusBrokeredMessagePropertiesOutput) ToServiceBusBrokeredMessagePropertiesOutputWithContext(ctx context.Context) ServiceBusBrokeredMessagePropertiesOutput {
	return o
}

func (o ServiceBusBrokeredMessagePropertiesOutput) ToServiceBusBrokeredMessagePropertiesPtrOutput() ServiceBusBrokeredMessagePropertiesPtrOutput {
	return o.ToServiceBusBrokeredMessagePropertiesPtrOutputWithContext(context.Background())
}

func (o ServiceBusBrokeredMessagePropertiesOutput) ToServiceBusBrokeredMessagePropertiesPtrOutputWithContext(ctx context.Context) ServiceBusBrokeredMessagePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceBusBrokeredMessageProperties) *ServiceBusBrokeredMessageProperties {
		return &v
	}).(ServiceBusBrokeredMessagePropertiesPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessageProperties) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessageProperties) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesOutput) ForcePersistence() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessageProperties) *bool { return v.ForcePersistence }).(pulumi.BoolPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessageProperties) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessageProperties) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessageProperties) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessageProperties) *string { return v.ReplyTo }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessageProperties) *string { return v.ReplyToSessionId }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesOutput) ScheduledEnqueueTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessageProperties) *string { return v.ScheduledEnqueueTimeUtc }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessageProperties) *string { return v.SessionId }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesOutput) TimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessageProperties) *string { return v.TimeToLive }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessageProperties) *string { return v.To }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesOutput) ViaPartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessageProperties) *string { return v.ViaPartitionKey }).(pulumi.StringPtrOutput)
}

type ServiceBusBrokeredMessagePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ServiceBusBrokeredMessagePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusBrokeredMessageProperties)(nil)).Elem()
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) ToServiceBusBrokeredMessagePropertiesPtrOutput() ServiceBusBrokeredMessagePropertiesPtrOutput {
	return o
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) ToServiceBusBrokeredMessagePropertiesPtrOutputWithContext(ctx context.Context) ServiceBusBrokeredMessagePropertiesPtrOutput {
	return o
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) Elem() ServiceBusBrokeredMessagePropertiesOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) ServiceBusBrokeredMessageProperties {
		if v != nil {
			return *v
		}
		var ret ServiceBusBrokeredMessageProperties
		return ret
	}).(ServiceBusBrokeredMessagePropertiesOutput)
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) *string {
		if v == nil {
			return nil
		}
		return v.ContentType
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) *string {
		if v == nil {
			return nil
		}
		return v.CorrelationId
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) ForcePersistence() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) *bool {
		if v == nil {
			return nil
		}
		return v.ForcePersistence
	}).(pulumi.BoolPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) *string {
		if v == nil {
			return nil
		}
		return v.MessageId
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) *string {
		if v == nil {
			return nil
		}
		return v.PartitionKey
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) *string {
		if v == nil {
			return nil
		}
		return v.ReplyTo
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) *string {
		if v == nil {
			return nil
		}
		return v.ReplyToSessionId
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) ScheduledEnqueueTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) *string {
		if v == nil {
			return nil
		}
		return v.ScheduledEnqueueTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) *string {
		if v == nil {
			return nil
		}
		return v.SessionId
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) TimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) *string {
		if v == nil {
			return nil
		}
		return v.TimeToLive
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesPtrOutput) ViaPartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessageProperties) *string {
		if v == nil {
			return nil
		}
		return v.ViaPartitionKey
	}).(pulumi.StringPtrOutput)
}

type ServiceBusBrokeredMessagePropertiesResponse struct {
	ContentType             *string `pulumi:"contentType"`
	CorrelationId           *string `pulumi:"correlationId"`
	ForcePersistence        *bool   `pulumi:"forcePersistence"`
	Label                   *string `pulumi:"label"`
	MessageId               *string `pulumi:"messageId"`
	PartitionKey            *string `pulumi:"partitionKey"`
	ReplyTo                 *string `pulumi:"replyTo"`
	ReplyToSessionId        *string `pulumi:"replyToSessionId"`
	ScheduledEnqueueTimeUtc *string `pulumi:"scheduledEnqueueTimeUtc"`
	SessionId               *string `pulumi:"sessionId"`
	TimeToLive              *string `pulumi:"timeToLive"`
	To                      *string `pulumi:"to"`
	ViaPartitionKey         *string `pulumi:"viaPartitionKey"`
}





type ServiceBusBrokeredMessagePropertiesResponseInput interface {
	pulumi.Input

	ToServiceBusBrokeredMessagePropertiesResponseOutput() ServiceBusBrokeredMessagePropertiesResponseOutput
	ToServiceBusBrokeredMessagePropertiesResponseOutputWithContext(context.Context) ServiceBusBrokeredMessagePropertiesResponseOutput
}

type ServiceBusBrokeredMessagePropertiesResponseArgs struct {
	ContentType             pulumi.StringPtrInput `pulumi:"contentType"`
	CorrelationId           pulumi.StringPtrInput `pulumi:"correlationId"`
	ForcePersistence        pulumi.BoolPtrInput   `pulumi:"forcePersistence"`
	Label                   pulumi.StringPtrInput `pulumi:"label"`
	MessageId               pulumi.StringPtrInput `pulumi:"messageId"`
	PartitionKey            pulumi.StringPtrInput `pulumi:"partitionKey"`
	ReplyTo                 pulumi.StringPtrInput `pulumi:"replyTo"`
	ReplyToSessionId        pulumi.StringPtrInput `pulumi:"replyToSessionId"`
	ScheduledEnqueueTimeUtc pulumi.StringPtrInput `pulumi:"scheduledEnqueueTimeUtc"`
	SessionId               pulumi.StringPtrInput `pulumi:"sessionId"`
	TimeToLive              pulumi.StringPtrInput `pulumi:"timeToLive"`
	To                      pulumi.StringPtrInput `pulumi:"to"`
	ViaPartitionKey         pulumi.StringPtrInput `pulumi:"viaPartitionKey"`
}

func (ServiceBusBrokeredMessagePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusBrokeredMessagePropertiesResponse)(nil)).Elem()
}

func (i ServiceBusBrokeredMessagePropertiesResponseArgs) ToServiceBusBrokeredMessagePropertiesResponseOutput() ServiceBusBrokeredMessagePropertiesResponseOutput {
	return i.ToServiceBusBrokeredMessagePropertiesResponseOutputWithContext(context.Background())
}

func (i ServiceBusBrokeredMessagePropertiesResponseArgs) ToServiceBusBrokeredMessagePropertiesResponseOutputWithContext(ctx context.Context) ServiceBusBrokeredMessagePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusBrokeredMessagePropertiesResponseOutput)
}

func (i ServiceBusBrokeredMessagePropertiesResponseArgs) ToServiceBusBrokeredMessagePropertiesResponsePtrOutput() ServiceBusBrokeredMessagePropertiesResponsePtrOutput {
	return i.ToServiceBusBrokeredMessagePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ServiceBusBrokeredMessagePropertiesResponseArgs) ToServiceBusBrokeredMessagePropertiesResponsePtrOutputWithContext(ctx context.Context) ServiceBusBrokeredMessagePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusBrokeredMessagePropertiesResponseOutput).ToServiceBusBrokeredMessagePropertiesResponsePtrOutputWithContext(ctx)
}









type ServiceBusBrokeredMessagePropertiesResponsePtrInput interface {
	pulumi.Input

	ToServiceBusBrokeredMessagePropertiesResponsePtrOutput() ServiceBusBrokeredMessagePropertiesResponsePtrOutput
	ToServiceBusBrokeredMessagePropertiesResponsePtrOutputWithContext(context.Context) ServiceBusBrokeredMessagePropertiesResponsePtrOutput
}

type serviceBusBrokeredMessagePropertiesResponsePtrType ServiceBusBrokeredMessagePropertiesResponseArgs

func ServiceBusBrokeredMessagePropertiesResponsePtr(v *ServiceBusBrokeredMessagePropertiesResponseArgs) ServiceBusBrokeredMessagePropertiesResponsePtrInput {
	return (*serviceBusBrokeredMessagePropertiesResponsePtrType)(v)
}

func (*serviceBusBrokeredMessagePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusBrokeredMessagePropertiesResponse)(nil)).Elem()
}

func (i *serviceBusBrokeredMessagePropertiesResponsePtrType) ToServiceBusBrokeredMessagePropertiesResponsePtrOutput() ServiceBusBrokeredMessagePropertiesResponsePtrOutput {
	return i.ToServiceBusBrokeredMessagePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *serviceBusBrokeredMessagePropertiesResponsePtrType) ToServiceBusBrokeredMessagePropertiesResponsePtrOutputWithContext(ctx context.Context) ServiceBusBrokeredMessagePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusBrokeredMessagePropertiesResponsePtrOutput)
}

type ServiceBusBrokeredMessagePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ServiceBusBrokeredMessagePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusBrokeredMessagePropertiesResponse)(nil)).Elem()
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) ToServiceBusBrokeredMessagePropertiesResponseOutput() ServiceBusBrokeredMessagePropertiesResponseOutput {
	return o
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) ToServiceBusBrokeredMessagePropertiesResponseOutputWithContext(ctx context.Context) ServiceBusBrokeredMessagePropertiesResponseOutput {
	return o
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) ToServiceBusBrokeredMessagePropertiesResponsePtrOutput() ServiceBusBrokeredMessagePropertiesResponsePtrOutput {
	return o.ToServiceBusBrokeredMessagePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) ToServiceBusBrokeredMessagePropertiesResponsePtrOutputWithContext(ctx context.Context) ServiceBusBrokeredMessagePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceBusBrokeredMessagePropertiesResponse) *ServiceBusBrokeredMessagePropertiesResponse {
		return &v
	}).(ServiceBusBrokeredMessagePropertiesResponsePtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessagePropertiesResponse) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessagePropertiesResponse) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) ForcePersistence() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessagePropertiesResponse) *bool { return v.ForcePersistence }).(pulumi.BoolPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessagePropertiesResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessagePropertiesResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessagePropertiesResponse) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessagePropertiesResponse) *string { return v.ReplyTo }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessagePropertiesResponse) *string { return v.ReplyToSessionId }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) ScheduledEnqueueTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessagePropertiesResponse) *string { return v.ScheduledEnqueueTimeUtc }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessagePropertiesResponse) *string { return v.SessionId }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) TimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessagePropertiesResponse) *string { return v.TimeToLive }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessagePropertiesResponse) *string { return v.To }).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponseOutput) ViaPartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusBrokeredMessagePropertiesResponse) *string { return v.ViaPartitionKey }).(pulumi.StringPtrOutput)
}

type ServiceBusBrokeredMessagePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceBusBrokeredMessagePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusBrokeredMessagePropertiesResponse)(nil)).Elem()
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) ToServiceBusBrokeredMessagePropertiesResponsePtrOutput() ServiceBusBrokeredMessagePropertiesResponsePtrOutput {
	return o
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) ToServiceBusBrokeredMessagePropertiesResponsePtrOutputWithContext(ctx context.Context) ServiceBusBrokeredMessagePropertiesResponsePtrOutput {
	return o
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) Elem() ServiceBusBrokeredMessagePropertiesResponseOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) ServiceBusBrokeredMessagePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ServiceBusBrokeredMessagePropertiesResponse
		return ret
	}).(ServiceBusBrokeredMessagePropertiesResponseOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentType
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CorrelationId
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) ForcePersistence() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ForcePersistence
	}).(pulumi.BoolPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.MessageId
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PartitionKey
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReplyTo
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReplyToSessionId
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) ScheduledEnqueueTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScheduledEnqueueTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SessionId
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) TimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeToLive
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusBrokeredMessagePropertiesResponsePtrOutput) ViaPartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusBrokeredMessagePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ViaPartitionKey
	}).(pulumi.StringPtrOutput)
}

type ServiceBusQueueMessage struct {
	Authentication            *ServiceBusAuthentication            `pulumi:"authentication"`
	BrokeredMessageProperties *ServiceBusBrokeredMessageProperties `pulumi:"brokeredMessageProperties"`
	CustomMessageProperties   map[string]string                    `pulumi:"customMessageProperties"`
	Message                   *string                              `pulumi:"message"`
	Namespace                 *string                              `pulumi:"namespace"`
	QueueName                 *string                              `pulumi:"queueName"`
	TransportType             *ServiceBusTransportType             `pulumi:"transportType"`
}





type ServiceBusQueueMessageInput interface {
	pulumi.Input

	ToServiceBusQueueMessageOutput() ServiceBusQueueMessageOutput
	ToServiceBusQueueMessageOutputWithContext(context.Context) ServiceBusQueueMessageOutput
}

type ServiceBusQueueMessageArgs struct {
	Authentication            ServiceBusAuthenticationPtrInput            `pulumi:"authentication"`
	BrokeredMessageProperties ServiceBusBrokeredMessagePropertiesPtrInput `pulumi:"brokeredMessageProperties"`
	CustomMessageProperties   pulumi.StringMapInput                       `pulumi:"customMessageProperties"`
	Message                   pulumi.StringPtrInput                       `pulumi:"message"`
	Namespace                 pulumi.StringPtrInput                       `pulumi:"namespace"`
	QueueName                 pulumi.StringPtrInput                       `pulumi:"queueName"`
	TransportType             ServiceBusTransportTypePtrInput             `pulumi:"transportType"`
}

func (ServiceBusQueueMessageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusQueueMessage)(nil)).Elem()
}

func (i ServiceBusQueueMessageArgs) ToServiceBusQueueMessageOutput() ServiceBusQueueMessageOutput {
	return i.ToServiceBusQueueMessageOutputWithContext(context.Background())
}

func (i ServiceBusQueueMessageArgs) ToServiceBusQueueMessageOutputWithContext(ctx context.Context) ServiceBusQueueMessageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusQueueMessageOutput)
}

func (i ServiceBusQueueMessageArgs) ToServiceBusQueueMessagePtrOutput() ServiceBusQueueMessagePtrOutput {
	return i.ToServiceBusQueueMessagePtrOutputWithContext(context.Background())
}

func (i ServiceBusQueueMessageArgs) ToServiceBusQueueMessagePtrOutputWithContext(ctx context.Context) ServiceBusQueueMessagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusQueueMessageOutput).ToServiceBusQueueMessagePtrOutputWithContext(ctx)
}









type ServiceBusQueueMessagePtrInput interface {
	pulumi.Input

	ToServiceBusQueueMessagePtrOutput() ServiceBusQueueMessagePtrOutput
	ToServiceBusQueueMessagePtrOutputWithContext(context.Context) ServiceBusQueueMessagePtrOutput
}

type serviceBusQueueMessagePtrType ServiceBusQueueMessageArgs

func ServiceBusQueueMessagePtr(v *ServiceBusQueueMessageArgs) ServiceBusQueueMessagePtrInput {
	return (*serviceBusQueueMessagePtrType)(v)
}

func (*serviceBusQueueMessagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusQueueMessage)(nil)).Elem()
}

func (i *serviceBusQueueMessagePtrType) ToServiceBusQueueMessagePtrOutput() ServiceBusQueueMessagePtrOutput {
	return i.ToServiceBusQueueMessagePtrOutputWithContext(context.Background())
}

func (i *serviceBusQueueMessagePtrType) ToServiceBusQueueMessagePtrOutputWithContext(ctx context.Context) ServiceBusQueueMessagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusQueueMessagePtrOutput)
}

type ServiceBusQueueMessageOutput struct{ *pulumi.OutputState }

func (ServiceBusQueueMessageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusQueueMessage)(nil)).Elem()
}

func (o ServiceBusQueueMessageOutput) ToServiceBusQueueMessageOutput() ServiceBusQueueMessageOutput {
	return o
}

func (o ServiceBusQueueMessageOutput) ToServiceBusQueueMessageOutputWithContext(ctx context.Context) ServiceBusQueueMessageOutput {
	return o
}

func (o ServiceBusQueueMessageOutput) ToServiceBusQueueMessagePtrOutput() ServiceBusQueueMessagePtrOutput {
	return o.ToServiceBusQueueMessagePtrOutputWithContext(context.Background())
}

func (o ServiceBusQueueMessageOutput) ToServiceBusQueueMessagePtrOutputWithContext(ctx context.Context) ServiceBusQueueMessagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceBusQueueMessage) *ServiceBusQueueMessage {
		return &v
	}).(ServiceBusQueueMessagePtrOutput)
}

func (o ServiceBusQueueMessageOutput) Authentication() ServiceBusAuthenticationPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueMessage) *ServiceBusAuthentication { return v.Authentication }).(ServiceBusAuthenticationPtrOutput)
}

func (o ServiceBusQueueMessageOutput) BrokeredMessageProperties() ServiceBusBrokeredMessagePropertiesPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueMessage) *ServiceBusBrokeredMessageProperties {
		return v.BrokeredMessageProperties
	}).(ServiceBusBrokeredMessagePropertiesPtrOutput)
}

func (o ServiceBusQueueMessageOutput) CustomMessageProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ServiceBusQueueMessage) map[string]string { return v.CustomMessageProperties }).(pulumi.StringMapOutput)
}

func (o ServiceBusQueueMessageOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueMessage) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueMessageOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueMessage) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueMessageOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueMessage) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueMessageOutput) TransportType() ServiceBusTransportTypePtrOutput {
	return o.ApplyT(func(v ServiceBusQueueMessage) *ServiceBusTransportType { return v.TransportType }).(ServiceBusTransportTypePtrOutput)
}

type ServiceBusQueueMessagePtrOutput struct{ *pulumi.OutputState }

func (ServiceBusQueueMessagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusQueueMessage)(nil)).Elem()
}

func (o ServiceBusQueueMessagePtrOutput) ToServiceBusQueueMessagePtrOutput() ServiceBusQueueMessagePtrOutput {
	return o
}

func (o ServiceBusQueueMessagePtrOutput) ToServiceBusQueueMessagePtrOutputWithContext(ctx context.Context) ServiceBusQueueMessagePtrOutput {
	return o
}

func (o ServiceBusQueueMessagePtrOutput) Elem() ServiceBusQueueMessageOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessage) ServiceBusQueueMessage {
		if v != nil {
			return *v
		}
		var ret ServiceBusQueueMessage
		return ret
	}).(ServiceBusQueueMessageOutput)
}

func (o ServiceBusQueueMessagePtrOutput) Authentication() ServiceBusAuthenticationPtrOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessage) *ServiceBusAuthentication {
		if v == nil {
			return nil
		}
		return v.Authentication
	}).(ServiceBusAuthenticationPtrOutput)
}

func (o ServiceBusQueueMessagePtrOutput) BrokeredMessageProperties() ServiceBusBrokeredMessagePropertiesPtrOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessage) *ServiceBusBrokeredMessageProperties {
		if v == nil {
			return nil
		}
		return v.BrokeredMessageProperties
	}).(ServiceBusBrokeredMessagePropertiesPtrOutput)
}

func (o ServiceBusQueueMessagePtrOutput) CustomMessageProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessage) map[string]string {
		if v == nil {
			return nil
		}
		return v.CustomMessageProperties
	}).(pulumi.StringMapOutput)
}

func (o ServiceBusQueueMessagePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessage) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueMessagePtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessage) *string {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueMessagePtrOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessage) *string {
		if v == nil {
			return nil
		}
		return v.QueueName
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueMessagePtrOutput) TransportType() ServiceBusTransportTypePtrOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessage) *ServiceBusTransportType {
		if v == nil {
			return nil
		}
		return v.TransportType
	}).(ServiceBusTransportTypePtrOutput)
}

type ServiceBusQueueMessageResponse struct {
	Authentication            *ServiceBusAuthenticationResponse            `pulumi:"authentication"`
	BrokeredMessageProperties *ServiceBusBrokeredMessagePropertiesResponse `pulumi:"brokeredMessageProperties"`
	CustomMessageProperties   map[string]string                            `pulumi:"customMessageProperties"`
	Message                   *string                                      `pulumi:"message"`
	Namespace                 *string                                      `pulumi:"namespace"`
	QueueName                 *string                                      `pulumi:"queueName"`
	TransportType             *string                                      `pulumi:"transportType"`
}





type ServiceBusQueueMessageResponseInput interface {
	pulumi.Input

	ToServiceBusQueueMessageResponseOutput() ServiceBusQueueMessageResponseOutput
	ToServiceBusQueueMessageResponseOutputWithContext(context.Context) ServiceBusQueueMessageResponseOutput
}

type ServiceBusQueueMessageResponseArgs struct {
	Authentication            ServiceBusAuthenticationResponsePtrInput            `pulumi:"authentication"`
	BrokeredMessageProperties ServiceBusBrokeredMessagePropertiesResponsePtrInput `pulumi:"brokeredMessageProperties"`
	CustomMessageProperties   pulumi.StringMapInput                               `pulumi:"customMessageProperties"`
	Message                   pulumi.StringPtrInput                               `pulumi:"message"`
	Namespace                 pulumi.StringPtrInput                               `pulumi:"namespace"`
	QueueName                 pulumi.StringPtrInput                               `pulumi:"queueName"`
	TransportType             pulumi.StringPtrInput                               `pulumi:"transportType"`
}

func (ServiceBusQueueMessageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusQueueMessageResponse)(nil)).Elem()
}

func (i ServiceBusQueueMessageResponseArgs) ToServiceBusQueueMessageResponseOutput() ServiceBusQueueMessageResponseOutput {
	return i.ToServiceBusQueueMessageResponseOutputWithContext(context.Background())
}

func (i ServiceBusQueueMessageResponseArgs) ToServiceBusQueueMessageResponseOutputWithContext(ctx context.Context) ServiceBusQueueMessageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusQueueMessageResponseOutput)
}

func (i ServiceBusQueueMessageResponseArgs) ToServiceBusQueueMessageResponsePtrOutput() ServiceBusQueueMessageResponsePtrOutput {
	return i.ToServiceBusQueueMessageResponsePtrOutputWithContext(context.Background())
}

func (i ServiceBusQueueMessageResponseArgs) ToServiceBusQueueMessageResponsePtrOutputWithContext(ctx context.Context) ServiceBusQueueMessageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusQueueMessageResponseOutput).ToServiceBusQueueMessageResponsePtrOutputWithContext(ctx)
}









type ServiceBusQueueMessageResponsePtrInput interface {
	pulumi.Input

	ToServiceBusQueueMessageResponsePtrOutput() ServiceBusQueueMessageResponsePtrOutput
	ToServiceBusQueueMessageResponsePtrOutputWithContext(context.Context) ServiceBusQueueMessageResponsePtrOutput
}

type serviceBusQueueMessageResponsePtrType ServiceBusQueueMessageResponseArgs

func ServiceBusQueueMessageResponsePtr(v *ServiceBusQueueMessageResponseArgs) ServiceBusQueueMessageResponsePtrInput {
	return (*serviceBusQueueMessageResponsePtrType)(v)
}

func (*serviceBusQueueMessageResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusQueueMessageResponse)(nil)).Elem()
}

func (i *serviceBusQueueMessageResponsePtrType) ToServiceBusQueueMessageResponsePtrOutput() ServiceBusQueueMessageResponsePtrOutput {
	return i.ToServiceBusQueueMessageResponsePtrOutputWithContext(context.Background())
}

func (i *serviceBusQueueMessageResponsePtrType) ToServiceBusQueueMessageResponsePtrOutputWithContext(ctx context.Context) ServiceBusQueueMessageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusQueueMessageResponsePtrOutput)
}

type ServiceBusQueueMessageResponseOutput struct{ *pulumi.OutputState }

func (ServiceBusQueueMessageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusQueueMessageResponse)(nil)).Elem()
}

func (o ServiceBusQueueMessageResponseOutput) ToServiceBusQueueMessageResponseOutput() ServiceBusQueueMessageResponseOutput {
	return o
}

func (o ServiceBusQueueMessageResponseOutput) ToServiceBusQueueMessageResponseOutputWithContext(ctx context.Context) ServiceBusQueueMessageResponseOutput {
	return o
}

func (o ServiceBusQueueMessageResponseOutput) ToServiceBusQueueMessageResponsePtrOutput() ServiceBusQueueMessageResponsePtrOutput {
	return o.ToServiceBusQueueMessageResponsePtrOutputWithContext(context.Background())
}

func (o ServiceBusQueueMessageResponseOutput) ToServiceBusQueueMessageResponsePtrOutputWithContext(ctx context.Context) ServiceBusQueueMessageResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceBusQueueMessageResponse) *ServiceBusQueueMessageResponse {
		return &v
	}).(ServiceBusQueueMessageResponsePtrOutput)
}

func (o ServiceBusQueueMessageResponseOutput) Authentication() ServiceBusAuthenticationResponsePtrOutput {
	return o.ApplyT(func(v ServiceBusQueueMessageResponse) *ServiceBusAuthenticationResponse { return v.Authentication }).(ServiceBusAuthenticationResponsePtrOutput)
}

func (o ServiceBusQueueMessageResponseOutput) BrokeredMessageProperties() ServiceBusBrokeredMessagePropertiesResponsePtrOutput {
	return o.ApplyT(func(v ServiceBusQueueMessageResponse) *ServiceBusBrokeredMessagePropertiesResponse {
		return v.BrokeredMessageProperties
	}).(ServiceBusBrokeredMessagePropertiesResponsePtrOutput)
}

func (o ServiceBusQueueMessageResponseOutput) CustomMessageProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ServiceBusQueueMessageResponse) map[string]string { return v.CustomMessageProperties }).(pulumi.StringMapOutput)
}

func (o ServiceBusQueueMessageResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueMessageResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueMessageResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueMessageResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueMessageResponseOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueMessageResponse) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueMessageResponseOutput) TransportType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueMessageResponse) *string { return v.TransportType }).(pulumi.StringPtrOutput)
}

type ServiceBusQueueMessageResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceBusQueueMessageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusQueueMessageResponse)(nil)).Elem()
}

func (o ServiceBusQueueMessageResponsePtrOutput) ToServiceBusQueueMessageResponsePtrOutput() ServiceBusQueueMessageResponsePtrOutput {
	return o
}

func (o ServiceBusQueueMessageResponsePtrOutput) ToServiceBusQueueMessageResponsePtrOutputWithContext(ctx context.Context) ServiceBusQueueMessageResponsePtrOutput {
	return o
}

func (o ServiceBusQueueMessageResponsePtrOutput) Elem() ServiceBusQueueMessageResponseOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessageResponse) ServiceBusQueueMessageResponse {
		if v != nil {
			return *v
		}
		var ret ServiceBusQueueMessageResponse
		return ret
	}).(ServiceBusQueueMessageResponseOutput)
}

func (o ServiceBusQueueMessageResponsePtrOutput) Authentication() ServiceBusAuthenticationResponsePtrOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessageResponse) *ServiceBusAuthenticationResponse {
		if v == nil {
			return nil
		}
		return v.Authentication
	}).(ServiceBusAuthenticationResponsePtrOutput)
}

func (o ServiceBusQueueMessageResponsePtrOutput) BrokeredMessageProperties() ServiceBusBrokeredMessagePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessageResponse) *ServiceBusBrokeredMessagePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.BrokeredMessageProperties
	}).(ServiceBusBrokeredMessagePropertiesResponsePtrOutput)
}

func (o ServiceBusQueueMessageResponsePtrOutput) CustomMessageProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessageResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.CustomMessageProperties
	}).(pulumi.StringMapOutput)
}

func (o ServiceBusQueueMessageResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueMessageResponsePtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueMessageResponsePtrOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.QueueName
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueMessageResponsePtrOutput) TransportType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusQueueMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.TransportType
	}).(pulumi.StringPtrOutput)
}

type ServiceBusTopicMessage struct {
	Authentication            *ServiceBusAuthentication            `pulumi:"authentication"`
	BrokeredMessageProperties *ServiceBusBrokeredMessageProperties `pulumi:"brokeredMessageProperties"`
	CustomMessageProperties   map[string]string                    `pulumi:"customMessageProperties"`
	Message                   *string                              `pulumi:"message"`
	Namespace                 *string                              `pulumi:"namespace"`
	TopicPath                 *string                              `pulumi:"topicPath"`
	TransportType             *ServiceBusTransportType             `pulumi:"transportType"`
}





type ServiceBusTopicMessageInput interface {
	pulumi.Input

	ToServiceBusTopicMessageOutput() ServiceBusTopicMessageOutput
	ToServiceBusTopicMessageOutputWithContext(context.Context) ServiceBusTopicMessageOutput
}

type ServiceBusTopicMessageArgs struct {
	Authentication            ServiceBusAuthenticationPtrInput            `pulumi:"authentication"`
	BrokeredMessageProperties ServiceBusBrokeredMessagePropertiesPtrInput `pulumi:"brokeredMessageProperties"`
	CustomMessageProperties   pulumi.StringMapInput                       `pulumi:"customMessageProperties"`
	Message                   pulumi.StringPtrInput                       `pulumi:"message"`
	Namespace                 pulumi.StringPtrInput                       `pulumi:"namespace"`
	TopicPath                 pulumi.StringPtrInput                       `pulumi:"topicPath"`
	TransportType             ServiceBusTransportTypePtrInput             `pulumi:"transportType"`
}

func (ServiceBusTopicMessageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusTopicMessage)(nil)).Elem()
}

func (i ServiceBusTopicMessageArgs) ToServiceBusTopicMessageOutput() ServiceBusTopicMessageOutput {
	return i.ToServiceBusTopicMessageOutputWithContext(context.Background())
}

func (i ServiceBusTopicMessageArgs) ToServiceBusTopicMessageOutputWithContext(ctx context.Context) ServiceBusTopicMessageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusTopicMessageOutput)
}

func (i ServiceBusTopicMessageArgs) ToServiceBusTopicMessagePtrOutput() ServiceBusTopicMessagePtrOutput {
	return i.ToServiceBusTopicMessagePtrOutputWithContext(context.Background())
}

func (i ServiceBusTopicMessageArgs) ToServiceBusTopicMessagePtrOutputWithContext(ctx context.Context) ServiceBusTopicMessagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusTopicMessageOutput).ToServiceBusTopicMessagePtrOutputWithContext(ctx)
}









type ServiceBusTopicMessagePtrInput interface {
	pulumi.Input

	ToServiceBusTopicMessagePtrOutput() ServiceBusTopicMessagePtrOutput
	ToServiceBusTopicMessagePtrOutputWithContext(context.Context) ServiceBusTopicMessagePtrOutput
}

type serviceBusTopicMessagePtrType ServiceBusTopicMessageArgs

func ServiceBusTopicMessagePtr(v *ServiceBusTopicMessageArgs) ServiceBusTopicMessagePtrInput {
	return (*serviceBusTopicMessagePtrType)(v)
}

func (*serviceBusTopicMessagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusTopicMessage)(nil)).Elem()
}

func (i *serviceBusTopicMessagePtrType) ToServiceBusTopicMessagePtrOutput() ServiceBusTopicMessagePtrOutput {
	return i.ToServiceBusTopicMessagePtrOutputWithContext(context.Background())
}

func (i *serviceBusTopicMessagePtrType) ToServiceBusTopicMessagePtrOutputWithContext(ctx context.Context) ServiceBusTopicMessagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusTopicMessagePtrOutput)
}

type ServiceBusTopicMessageOutput struct{ *pulumi.OutputState }

func (ServiceBusTopicMessageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusTopicMessage)(nil)).Elem()
}

func (o ServiceBusTopicMessageOutput) ToServiceBusTopicMessageOutput() ServiceBusTopicMessageOutput {
	return o
}

func (o ServiceBusTopicMessageOutput) ToServiceBusTopicMessageOutputWithContext(ctx context.Context) ServiceBusTopicMessageOutput {
	return o
}

func (o ServiceBusTopicMessageOutput) ToServiceBusTopicMessagePtrOutput() ServiceBusTopicMessagePtrOutput {
	return o.ToServiceBusTopicMessagePtrOutputWithContext(context.Background())
}

func (o ServiceBusTopicMessageOutput) ToServiceBusTopicMessagePtrOutputWithContext(ctx context.Context) ServiceBusTopicMessagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceBusTopicMessage) *ServiceBusTopicMessage {
		return &v
	}).(ServiceBusTopicMessagePtrOutput)
}

func (o ServiceBusTopicMessageOutput) Authentication() ServiceBusAuthenticationPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicMessage) *ServiceBusAuthentication { return v.Authentication }).(ServiceBusAuthenticationPtrOutput)
}

func (o ServiceBusTopicMessageOutput) BrokeredMessageProperties() ServiceBusBrokeredMessagePropertiesPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicMessage) *ServiceBusBrokeredMessageProperties {
		return v.BrokeredMessageProperties
	}).(ServiceBusBrokeredMessagePropertiesPtrOutput)
}

func (o ServiceBusTopicMessageOutput) CustomMessageProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ServiceBusTopicMessage) map[string]string { return v.CustomMessageProperties }).(pulumi.StringMapOutput)
}

func (o ServiceBusTopicMessageOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicMessage) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicMessageOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicMessage) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicMessageOutput) TopicPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicMessage) *string { return v.TopicPath }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicMessageOutput) TransportType() ServiceBusTransportTypePtrOutput {
	return o.ApplyT(func(v ServiceBusTopicMessage) *ServiceBusTransportType { return v.TransportType }).(ServiceBusTransportTypePtrOutput)
}

type ServiceBusTopicMessagePtrOutput struct{ *pulumi.OutputState }

func (ServiceBusTopicMessagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusTopicMessage)(nil)).Elem()
}

func (o ServiceBusTopicMessagePtrOutput) ToServiceBusTopicMessagePtrOutput() ServiceBusTopicMessagePtrOutput {
	return o
}

func (o ServiceBusTopicMessagePtrOutput) ToServiceBusTopicMessagePtrOutputWithContext(ctx context.Context) ServiceBusTopicMessagePtrOutput {
	return o
}

func (o ServiceBusTopicMessagePtrOutput) Elem() ServiceBusTopicMessageOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessage) ServiceBusTopicMessage {
		if v != nil {
			return *v
		}
		var ret ServiceBusTopicMessage
		return ret
	}).(ServiceBusTopicMessageOutput)
}

func (o ServiceBusTopicMessagePtrOutput) Authentication() ServiceBusAuthenticationPtrOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessage) *ServiceBusAuthentication {
		if v == nil {
			return nil
		}
		return v.Authentication
	}).(ServiceBusAuthenticationPtrOutput)
}

func (o ServiceBusTopicMessagePtrOutput) BrokeredMessageProperties() ServiceBusBrokeredMessagePropertiesPtrOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessage) *ServiceBusBrokeredMessageProperties {
		if v == nil {
			return nil
		}
		return v.BrokeredMessageProperties
	}).(ServiceBusBrokeredMessagePropertiesPtrOutput)
}

func (o ServiceBusTopicMessagePtrOutput) CustomMessageProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessage) map[string]string {
		if v == nil {
			return nil
		}
		return v.CustomMessageProperties
	}).(pulumi.StringMapOutput)
}

func (o ServiceBusTopicMessagePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessage) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicMessagePtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessage) *string {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicMessagePtrOutput) TopicPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessage) *string {
		if v == nil {
			return nil
		}
		return v.TopicPath
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicMessagePtrOutput) TransportType() ServiceBusTransportTypePtrOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessage) *ServiceBusTransportType {
		if v == nil {
			return nil
		}
		return v.TransportType
	}).(ServiceBusTransportTypePtrOutput)
}

type ServiceBusTopicMessageResponse struct {
	Authentication            *ServiceBusAuthenticationResponse            `pulumi:"authentication"`
	BrokeredMessageProperties *ServiceBusBrokeredMessagePropertiesResponse `pulumi:"brokeredMessageProperties"`
	CustomMessageProperties   map[string]string                            `pulumi:"customMessageProperties"`
	Message                   *string                                      `pulumi:"message"`
	Namespace                 *string                                      `pulumi:"namespace"`
	TopicPath                 *string                                      `pulumi:"topicPath"`
	TransportType             *string                                      `pulumi:"transportType"`
}





type ServiceBusTopicMessageResponseInput interface {
	pulumi.Input

	ToServiceBusTopicMessageResponseOutput() ServiceBusTopicMessageResponseOutput
	ToServiceBusTopicMessageResponseOutputWithContext(context.Context) ServiceBusTopicMessageResponseOutput
}

type ServiceBusTopicMessageResponseArgs struct {
	Authentication            ServiceBusAuthenticationResponsePtrInput            `pulumi:"authentication"`
	BrokeredMessageProperties ServiceBusBrokeredMessagePropertiesResponsePtrInput `pulumi:"brokeredMessageProperties"`
	CustomMessageProperties   pulumi.StringMapInput                               `pulumi:"customMessageProperties"`
	Message                   pulumi.StringPtrInput                               `pulumi:"message"`
	Namespace                 pulumi.StringPtrInput                               `pulumi:"namespace"`
	TopicPath                 pulumi.StringPtrInput                               `pulumi:"topicPath"`
	TransportType             pulumi.StringPtrInput                               `pulumi:"transportType"`
}

func (ServiceBusTopicMessageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusTopicMessageResponse)(nil)).Elem()
}

func (i ServiceBusTopicMessageResponseArgs) ToServiceBusTopicMessageResponseOutput() ServiceBusTopicMessageResponseOutput {
	return i.ToServiceBusTopicMessageResponseOutputWithContext(context.Background())
}

func (i ServiceBusTopicMessageResponseArgs) ToServiceBusTopicMessageResponseOutputWithContext(ctx context.Context) ServiceBusTopicMessageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusTopicMessageResponseOutput)
}

func (i ServiceBusTopicMessageResponseArgs) ToServiceBusTopicMessageResponsePtrOutput() ServiceBusTopicMessageResponsePtrOutput {
	return i.ToServiceBusTopicMessageResponsePtrOutputWithContext(context.Background())
}

func (i ServiceBusTopicMessageResponseArgs) ToServiceBusTopicMessageResponsePtrOutputWithContext(ctx context.Context) ServiceBusTopicMessageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusTopicMessageResponseOutput).ToServiceBusTopicMessageResponsePtrOutputWithContext(ctx)
}









type ServiceBusTopicMessageResponsePtrInput interface {
	pulumi.Input

	ToServiceBusTopicMessageResponsePtrOutput() ServiceBusTopicMessageResponsePtrOutput
	ToServiceBusTopicMessageResponsePtrOutputWithContext(context.Context) ServiceBusTopicMessageResponsePtrOutput
}

type serviceBusTopicMessageResponsePtrType ServiceBusTopicMessageResponseArgs

func ServiceBusTopicMessageResponsePtr(v *ServiceBusTopicMessageResponseArgs) ServiceBusTopicMessageResponsePtrInput {
	return (*serviceBusTopicMessageResponsePtrType)(v)
}

func (*serviceBusTopicMessageResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusTopicMessageResponse)(nil)).Elem()
}

func (i *serviceBusTopicMessageResponsePtrType) ToServiceBusTopicMessageResponsePtrOutput() ServiceBusTopicMessageResponsePtrOutput {
	return i.ToServiceBusTopicMessageResponsePtrOutputWithContext(context.Background())
}

func (i *serviceBusTopicMessageResponsePtrType) ToServiceBusTopicMessageResponsePtrOutputWithContext(ctx context.Context) ServiceBusTopicMessageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusTopicMessageResponsePtrOutput)
}

type ServiceBusTopicMessageResponseOutput struct{ *pulumi.OutputState }

func (ServiceBusTopicMessageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusTopicMessageResponse)(nil)).Elem()
}

func (o ServiceBusTopicMessageResponseOutput) ToServiceBusTopicMessageResponseOutput() ServiceBusTopicMessageResponseOutput {
	return o
}

func (o ServiceBusTopicMessageResponseOutput) ToServiceBusTopicMessageResponseOutputWithContext(ctx context.Context) ServiceBusTopicMessageResponseOutput {
	return o
}

func (o ServiceBusTopicMessageResponseOutput) ToServiceBusTopicMessageResponsePtrOutput() ServiceBusTopicMessageResponsePtrOutput {
	return o.ToServiceBusTopicMessageResponsePtrOutputWithContext(context.Background())
}

func (o ServiceBusTopicMessageResponseOutput) ToServiceBusTopicMessageResponsePtrOutputWithContext(ctx context.Context) ServiceBusTopicMessageResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceBusTopicMessageResponse) *ServiceBusTopicMessageResponse {
		return &v
	}).(ServiceBusTopicMessageResponsePtrOutput)
}

func (o ServiceBusTopicMessageResponseOutput) Authentication() ServiceBusAuthenticationResponsePtrOutput {
	return o.ApplyT(func(v ServiceBusTopicMessageResponse) *ServiceBusAuthenticationResponse { return v.Authentication }).(ServiceBusAuthenticationResponsePtrOutput)
}

func (o ServiceBusTopicMessageResponseOutput) BrokeredMessageProperties() ServiceBusBrokeredMessagePropertiesResponsePtrOutput {
	return o.ApplyT(func(v ServiceBusTopicMessageResponse) *ServiceBusBrokeredMessagePropertiesResponse {
		return v.BrokeredMessageProperties
	}).(ServiceBusBrokeredMessagePropertiesResponsePtrOutput)
}

func (o ServiceBusTopicMessageResponseOutput) CustomMessageProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ServiceBusTopicMessageResponse) map[string]string { return v.CustomMessageProperties }).(pulumi.StringMapOutput)
}

func (o ServiceBusTopicMessageResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicMessageResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicMessageResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicMessageResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicMessageResponseOutput) TopicPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicMessageResponse) *string { return v.TopicPath }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicMessageResponseOutput) TransportType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicMessageResponse) *string { return v.TransportType }).(pulumi.StringPtrOutput)
}

type ServiceBusTopicMessageResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceBusTopicMessageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBusTopicMessageResponse)(nil)).Elem()
}

func (o ServiceBusTopicMessageResponsePtrOutput) ToServiceBusTopicMessageResponsePtrOutput() ServiceBusTopicMessageResponsePtrOutput {
	return o
}

func (o ServiceBusTopicMessageResponsePtrOutput) ToServiceBusTopicMessageResponsePtrOutputWithContext(ctx context.Context) ServiceBusTopicMessageResponsePtrOutput {
	return o
}

func (o ServiceBusTopicMessageResponsePtrOutput) Elem() ServiceBusTopicMessageResponseOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessageResponse) ServiceBusTopicMessageResponse {
		if v != nil {
			return *v
		}
		var ret ServiceBusTopicMessageResponse
		return ret
	}).(ServiceBusTopicMessageResponseOutput)
}

func (o ServiceBusTopicMessageResponsePtrOutput) Authentication() ServiceBusAuthenticationResponsePtrOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessageResponse) *ServiceBusAuthenticationResponse {
		if v == nil {
			return nil
		}
		return v.Authentication
	}).(ServiceBusAuthenticationResponsePtrOutput)
}

func (o ServiceBusTopicMessageResponsePtrOutput) BrokeredMessageProperties() ServiceBusBrokeredMessagePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessageResponse) *ServiceBusBrokeredMessagePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.BrokeredMessageProperties
	}).(ServiceBusBrokeredMessagePropertiesResponsePtrOutput)
}

func (o ServiceBusTopicMessageResponsePtrOutput) CustomMessageProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessageResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.CustomMessageProperties
	}).(pulumi.StringMapOutput)
}

func (o ServiceBusTopicMessageResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicMessageResponsePtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicMessageResponsePtrOutput) TopicPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.TopicPath
	}).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicMessageResponsePtrOutput) TransportType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBusTopicMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.TransportType
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name *SkuDefinition `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name SkuDefinitionPtrInput `pulumi:"name"`
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

func (o SkuOutput) Name() SkuDefinitionPtrOutput {
	return o.ApplyT(func(v Sku) *SkuDefinition { return v.Name }).(SkuDefinitionPtrOutput)
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

func (o SkuPtrOutput) Name() SkuDefinitionPtrOutput {
	return o.ApplyT(func(v *Sku) *SkuDefinition {
		if v == nil {
			return nil
		}
		return v.Name
	}).(SkuDefinitionPtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
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

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type StorageQueueMessage struct {
	Message        *string `pulumi:"message"`
	QueueName      *string `pulumi:"queueName"`
	SasToken       *string `pulumi:"sasToken"`
	StorageAccount *string `pulumi:"storageAccount"`
}





type StorageQueueMessageInput interface {
	pulumi.Input

	ToStorageQueueMessageOutput() StorageQueueMessageOutput
	ToStorageQueueMessageOutputWithContext(context.Context) StorageQueueMessageOutput
}

type StorageQueueMessageArgs struct {
	Message        pulumi.StringPtrInput `pulumi:"message"`
	QueueName      pulumi.StringPtrInput `pulumi:"queueName"`
	SasToken       pulumi.StringPtrInput `pulumi:"sasToken"`
	StorageAccount pulumi.StringPtrInput `pulumi:"storageAccount"`
}

func (StorageQueueMessageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQueueMessage)(nil)).Elem()
}

func (i StorageQueueMessageArgs) ToStorageQueueMessageOutput() StorageQueueMessageOutput {
	return i.ToStorageQueueMessageOutputWithContext(context.Background())
}

func (i StorageQueueMessageArgs) ToStorageQueueMessageOutputWithContext(ctx context.Context) StorageQueueMessageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageQueueMessageOutput)
}

func (i StorageQueueMessageArgs) ToStorageQueueMessagePtrOutput() StorageQueueMessagePtrOutput {
	return i.ToStorageQueueMessagePtrOutputWithContext(context.Background())
}

func (i StorageQueueMessageArgs) ToStorageQueueMessagePtrOutputWithContext(ctx context.Context) StorageQueueMessagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageQueueMessageOutput).ToStorageQueueMessagePtrOutputWithContext(ctx)
}









type StorageQueueMessagePtrInput interface {
	pulumi.Input

	ToStorageQueueMessagePtrOutput() StorageQueueMessagePtrOutput
	ToStorageQueueMessagePtrOutputWithContext(context.Context) StorageQueueMessagePtrOutput
}

type storageQueueMessagePtrType StorageQueueMessageArgs

func StorageQueueMessagePtr(v *StorageQueueMessageArgs) StorageQueueMessagePtrInput {
	return (*storageQueueMessagePtrType)(v)
}

func (*storageQueueMessagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageQueueMessage)(nil)).Elem()
}

func (i *storageQueueMessagePtrType) ToStorageQueueMessagePtrOutput() StorageQueueMessagePtrOutput {
	return i.ToStorageQueueMessagePtrOutputWithContext(context.Background())
}

func (i *storageQueueMessagePtrType) ToStorageQueueMessagePtrOutputWithContext(ctx context.Context) StorageQueueMessagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageQueueMessagePtrOutput)
}

type StorageQueueMessageOutput struct{ *pulumi.OutputState }

func (StorageQueueMessageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQueueMessage)(nil)).Elem()
}

func (o StorageQueueMessageOutput) ToStorageQueueMessageOutput() StorageQueueMessageOutput {
	return o
}

func (o StorageQueueMessageOutput) ToStorageQueueMessageOutputWithContext(ctx context.Context) StorageQueueMessageOutput {
	return o
}

func (o StorageQueueMessageOutput) ToStorageQueueMessagePtrOutput() StorageQueueMessagePtrOutput {
	return o.ToStorageQueueMessagePtrOutputWithContext(context.Background())
}

func (o StorageQueueMessageOutput) ToStorageQueueMessagePtrOutputWithContext(ctx context.Context) StorageQueueMessagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageQueueMessage) *StorageQueueMessage {
		return &v
	}).(StorageQueueMessagePtrOutput)
}

func (o StorageQueueMessageOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueMessage) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o StorageQueueMessageOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueMessage) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

func (o StorageQueueMessageOutput) SasToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueMessage) *string { return v.SasToken }).(pulumi.StringPtrOutput)
}

func (o StorageQueueMessageOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueMessage) *string { return v.StorageAccount }).(pulumi.StringPtrOutput)
}

type StorageQueueMessagePtrOutput struct{ *pulumi.OutputState }

func (StorageQueueMessagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageQueueMessage)(nil)).Elem()
}

func (o StorageQueueMessagePtrOutput) ToStorageQueueMessagePtrOutput() StorageQueueMessagePtrOutput {
	return o
}

func (o StorageQueueMessagePtrOutput) ToStorageQueueMessagePtrOutputWithContext(ctx context.Context) StorageQueueMessagePtrOutput {
	return o
}

func (o StorageQueueMessagePtrOutput) Elem() StorageQueueMessageOutput {
	return o.ApplyT(func(v *StorageQueueMessage) StorageQueueMessage {
		if v != nil {
			return *v
		}
		var ret StorageQueueMessage
		return ret
	}).(StorageQueueMessageOutput)
}

func (o StorageQueueMessagePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageQueueMessage) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o StorageQueueMessagePtrOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageQueueMessage) *string {
		if v == nil {
			return nil
		}
		return v.QueueName
	}).(pulumi.StringPtrOutput)
}

func (o StorageQueueMessagePtrOutput) SasToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageQueueMessage) *string {
		if v == nil {
			return nil
		}
		return v.SasToken
	}).(pulumi.StringPtrOutput)
}

func (o StorageQueueMessagePtrOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageQueueMessage) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccount
	}).(pulumi.StringPtrOutput)
}

type StorageQueueMessageResponse struct {
	Message        *string `pulumi:"message"`
	QueueName      *string `pulumi:"queueName"`
	SasToken       *string `pulumi:"sasToken"`
	StorageAccount *string `pulumi:"storageAccount"`
}





type StorageQueueMessageResponseInput interface {
	pulumi.Input

	ToStorageQueueMessageResponseOutput() StorageQueueMessageResponseOutput
	ToStorageQueueMessageResponseOutputWithContext(context.Context) StorageQueueMessageResponseOutput
}

type StorageQueueMessageResponseArgs struct {
	Message        pulumi.StringPtrInput `pulumi:"message"`
	QueueName      pulumi.StringPtrInput `pulumi:"queueName"`
	SasToken       pulumi.StringPtrInput `pulumi:"sasToken"`
	StorageAccount pulumi.StringPtrInput `pulumi:"storageAccount"`
}

func (StorageQueueMessageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQueueMessageResponse)(nil)).Elem()
}

func (i StorageQueueMessageResponseArgs) ToStorageQueueMessageResponseOutput() StorageQueueMessageResponseOutput {
	return i.ToStorageQueueMessageResponseOutputWithContext(context.Background())
}

func (i StorageQueueMessageResponseArgs) ToStorageQueueMessageResponseOutputWithContext(ctx context.Context) StorageQueueMessageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageQueueMessageResponseOutput)
}

func (i StorageQueueMessageResponseArgs) ToStorageQueueMessageResponsePtrOutput() StorageQueueMessageResponsePtrOutput {
	return i.ToStorageQueueMessageResponsePtrOutputWithContext(context.Background())
}

func (i StorageQueueMessageResponseArgs) ToStorageQueueMessageResponsePtrOutputWithContext(ctx context.Context) StorageQueueMessageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageQueueMessageResponseOutput).ToStorageQueueMessageResponsePtrOutputWithContext(ctx)
}









type StorageQueueMessageResponsePtrInput interface {
	pulumi.Input

	ToStorageQueueMessageResponsePtrOutput() StorageQueueMessageResponsePtrOutput
	ToStorageQueueMessageResponsePtrOutputWithContext(context.Context) StorageQueueMessageResponsePtrOutput
}

type storageQueueMessageResponsePtrType StorageQueueMessageResponseArgs

func StorageQueueMessageResponsePtr(v *StorageQueueMessageResponseArgs) StorageQueueMessageResponsePtrInput {
	return (*storageQueueMessageResponsePtrType)(v)
}

func (*storageQueueMessageResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageQueueMessageResponse)(nil)).Elem()
}

func (i *storageQueueMessageResponsePtrType) ToStorageQueueMessageResponsePtrOutput() StorageQueueMessageResponsePtrOutput {
	return i.ToStorageQueueMessageResponsePtrOutputWithContext(context.Background())
}

func (i *storageQueueMessageResponsePtrType) ToStorageQueueMessageResponsePtrOutputWithContext(ctx context.Context) StorageQueueMessageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageQueueMessageResponsePtrOutput)
}

type StorageQueueMessageResponseOutput struct{ *pulumi.OutputState }

func (StorageQueueMessageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQueueMessageResponse)(nil)).Elem()
}

func (o StorageQueueMessageResponseOutput) ToStorageQueueMessageResponseOutput() StorageQueueMessageResponseOutput {
	return o
}

func (o StorageQueueMessageResponseOutput) ToStorageQueueMessageResponseOutputWithContext(ctx context.Context) StorageQueueMessageResponseOutput {
	return o
}

func (o StorageQueueMessageResponseOutput) ToStorageQueueMessageResponsePtrOutput() StorageQueueMessageResponsePtrOutput {
	return o.ToStorageQueueMessageResponsePtrOutputWithContext(context.Background())
}

func (o StorageQueueMessageResponseOutput) ToStorageQueueMessageResponsePtrOutputWithContext(ctx context.Context) StorageQueueMessageResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageQueueMessageResponse) *StorageQueueMessageResponse {
		return &v
	}).(StorageQueueMessageResponsePtrOutput)
}

func (o StorageQueueMessageResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueMessageResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o StorageQueueMessageResponseOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueMessageResponse) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

func (o StorageQueueMessageResponseOutput) SasToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueMessageResponse) *string { return v.SasToken }).(pulumi.StringPtrOutput)
}

func (o StorageQueueMessageResponseOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueMessageResponse) *string { return v.StorageAccount }).(pulumi.StringPtrOutput)
}

type StorageQueueMessageResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageQueueMessageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageQueueMessageResponse)(nil)).Elem()
}

func (o StorageQueueMessageResponsePtrOutput) ToStorageQueueMessageResponsePtrOutput() StorageQueueMessageResponsePtrOutput {
	return o
}

func (o StorageQueueMessageResponsePtrOutput) ToStorageQueueMessageResponsePtrOutputWithContext(ctx context.Context) StorageQueueMessageResponsePtrOutput {
	return o
}

func (o StorageQueueMessageResponsePtrOutput) Elem() StorageQueueMessageResponseOutput {
	return o.ApplyT(func(v *StorageQueueMessageResponse) StorageQueueMessageResponse {
		if v != nil {
			return *v
		}
		var ret StorageQueueMessageResponse
		return ret
	}).(StorageQueueMessageResponseOutput)
}

func (o StorageQueueMessageResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageQueueMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o StorageQueueMessageResponsePtrOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageQueueMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.QueueName
	}).(pulumi.StringPtrOutput)
}

func (o StorageQueueMessageResponsePtrOutput) SasToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageQueueMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasToken
	}).(pulumi.StringPtrOutput)
}

func (o StorageQueueMessageResponsePtrOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageQueueMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccount
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HttpAuthenticationOutput{})
	pulumi.RegisterOutputType(HttpAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(HttpAuthenticationResponseOutput{})
	pulumi.RegisterOutputType(HttpAuthenticationResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpRequestOutput{})
	pulumi.RegisterOutputType(HttpRequestPtrOutput{})
	pulumi.RegisterOutputType(HttpRequestResponseOutput{})
	pulumi.RegisterOutputType(HttpRequestResponsePtrOutput{})
	pulumi.RegisterOutputType(JobActionOutput{})
	pulumi.RegisterOutputType(JobActionPtrOutput{})
	pulumi.RegisterOutputType(JobActionResponseOutput{})
	pulumi.RegisterOutputType(JobActionResponsePtrOutput{})
	pulumi.RegisterOutputType(JobCollectionPropertiesOutput{})
	pulumi.RegisterOutputType(JobCollectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(JobCollectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(JobCollectionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(JobCollectionQuotaOutput{})
	pulumi.RegisterOutputType(JobCollectionQuotaPtrOutput{})
	pulumi.RegisterOutputType(JobCollectionQuotaResponseOutput{})
	pulumi.RegisterOutputType(JobCollectionQuotaResponsePtrOutput{})
	pulumi.RegisterOutputType(JobErrorActionOutput{})
	pulumi.RegisterOutputType(JobErrorActionPtrOutput{})
	pulumi.RegisterOutputType(JobErrorActionResponseOutput{})
	pulumi.RegisterOutputType(JobErrorActionResponsePtrOutput{})
	pulumi.RegisterOutputType(JobMaxRecurrenceOutput{})
	pulumi.RegisterOutputType(JobMaxRecurrencePtrOutput{})
	pulumi.RegisterOutputType(JobMaxRecurrenceResponseOutput{})
	pulumi.RegisterOutputType(JobMaxRecurrenceResponsePtrOutput{})
	pulumi.RegisterOutputType(JobPropertiesOutput{})
	pulumi.RegisterOutputType(JobPropertiesPtrOutput{})
	pulumi.RegisterOutputType(JobPropertiesResponseOutput{})
	pulumi.RegisterOutputType(JobPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(JobRecurrenceOutput{})
	pulumi.RegisterOutputType(JobRecurrencePtrOutput{})
	pulumi.RegisterOutputType(JobRecurrenceResponseOutput{})
	pulumi.RegisterOutputType(JobRecurrenceResponsePtrOutput{})
	pulumi.RegisterOutputType(JobRecurrenceScheduleOutput{})
	pulumi.RegisterOutputType(JobRecurrenceSchedulePtrOutput{})
	pulumi.RegisterOutputType(JobRecurrenceScheduleMonthlyOccurrenceOutput{})
	pulumi.RegisterOutputType(JobRecurrenceScheduleMonthlyOccurrenceArrayOutput{})
	pulumi.RegisterOutputType(JobRecurrenceScheduleMonthlyOccurrenceResponseOutput{})
	pulumi.RegisterOutputType(JobRecurrenceScheduleMonthlyOccurrenceResponseArrayOutput{})
	pulumi.RegisterOutputType(JobRecurrenceScheduleResponseOutput{})
	pulumi.RegisterOutputType(JobRecurrenceScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(JobStatusResponseOutput{})
	pulumi.RegisterOutputType(JobStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(RetryPolicyOutput{})
	pulumi.RegisterOutputType(RetryPolicyPtrOutput{})
	pulumi.RegisterOutputType(RetryPolicyResponseOutput{})
	pulumi.RegisterOutputType(RetryPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceBusAuthenticationOutput{})
	pulumi.RegisterOutputType(ServiceBusAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(ServiceBusAuthenticationResponseOutput{})
	pulumi.RegisterOutputType(ServiceBusAuthenticationResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceBusBrokeredMessagePropertiesOutput{})
	pulumi.RegisterOutputType(ServiceBusBrokeredMessagePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ServiceBusBrokeredMessagePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServiceBusBrokeredMessagePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceBusQueueMessageOutput{})
	pulumi.RegisterOutputType(ServiceBusQueueMessagePtrOutput{})
	pulumi.RegisterOutputType(ServiceBusQueueMessageResponseOutput{})
	pulumi.RegisterOutputType(ServiceBusQueueMessageResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceBusTopicMessageOutput{})
	pulumi.RegisterOutputType(ServiceBusTopicMessagePtrOutput{})
	pulumi.RegisterOutputType(ServiceBusTopicMessageResponseOutput{})
	pulumi.RegisterOutputType(ServiceBusTopicMessageResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageQueueMessageOutput{})
	pulumi.RegisterOutputType(StorageQueueMessagePtrOutput{})
	pulumi.RegisterOutputType(StorageQueueMessageResponseOutput{})
	pulumi.RegisterOutputType(StorageQueueMessageResponsePtrOutput{})
}
