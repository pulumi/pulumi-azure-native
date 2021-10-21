


package v20180630

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContentHash struct {
	Algorithm string `pulumi:"algorithm"`
	Value     string `pulumi:"value"`
}





type ContentHashInput interface {
	pulumi.Input

	ToContentHashOutput() ContentHashOutput
	ToContentHashOutputWithContext(context.Context) ContentHashOutput
}

type ContentHashArgs struct {
	Algorithm pulumi.StringInput `pulumi:"algorithm"`
	Value     pulumi.StringInput `pulumi:"value"`
}

func (ContentHashArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHash)(nil)).Elem()
}

func (i ContentHashArgs) ToContentHashOutput() ContentHashOutput {
	return i.ToContentHashOutputWithContext(context.Background())
}

func (i ContentHashArgs) ToContentHashOutputWithContext(ctx context.Context) ContentHashOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashOutput)
}

func (i ContentHashArgs) ToContentHashPtrOutput() ContentHashPtrOutput {
	return i.ToContentHashPtrOutputWithContext(context.Background())
}

func (i ContentHashArgs) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashOutput).ToContentHashPtrOutputWithContext(ctx)
}









type ContentHashPtrInput interface {
	pulumi.Input

	ToContentHashPtrOutput() ContentHashPtrOutput
	ToContentHashPtrOutputWithContext(context.Context) ContentHashPtrOutput
}

type contentHashPtrType ContentHashArgs

func ContentHashPtr(v *ContentHashArgs) ContentHashPtrInput {
	return (*contentHashPtrType)(v)
}

func (*contentHashPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHash)(nil)).Elem()
}

func (i *contentHashPtrType) ToContentHashPtrOutput() ContentHashPtrOutput {
	return i.ToContentHashPtrOutputWithContext(context.Background())
}

func (i *contentHashPtrType) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashPtrOutput)
}

type ContentHashOutput struct{ *pulumi.OutputState }

func (ContentHashOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHash)(nil)).Elem()
}

func (o ContentHashOutput) ToContentHashOutput() ContentHashOutput {
	return o
}

func (o ContentHashOutput) ToContentHashOutputWithContext(ctx context.Context) ContentHashOutput {
	return o
}

func (o ContentHashOutput) ToContentHashPtrOutput() ContentHashPtrOutput {
	return o.ToContentHashPtrOutputWithContext(context.Background())
}

func (o ContentHashOutput) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentHash) *ContentHash {
		return &v
	}).(ContentHashPtrOutput)
}

func (o ContentHashOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHash) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o ContentHashOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHash) string { return v.Value }).(pulumi.StringOutput)
}

type ContentHashPtrOutput struct{ *pulumi.OutputState }

func (ContentHashPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHash)(nil)).Elem()
}

func (o ContentHashPtrOutput) ToContentHashPtrOutput() ContentHashPtrOutput {
	return o
}

func (o ContentHashPtrOutput) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return o
}

func (o ContentHashPtrOutput) Elem() ContentHashOutput {
	return o.ApplyT(func(v *ContentHash) ContentHash {
		if v != nil {
			return *v
		}
		var ret ContentHash
		return ret
	}).(ContentHashOutput)
}

func (o ContentHashPtrOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHash) *string {
		if v == nil {
			return nil
		}
		return &v.Algorithm
	}).(pulumi.StringPtrOutput)
}

func (o ContentHashPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHash) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type ContentHashResponse struct {
	Algorithm string `pulumi:"algorithm"`
	Value     string `pulumi:"value"`
}





type ContentHashResponseInput interface {
	pulumi.Input

	ToContentHashResponseOutput() ContentHashResponseOutput
	ToContentHashResponseOutputWithContext(context.Context) ContentHashResponseOutput
}

type ContentHashResponseArgs struct {
	Algorithm pulumi.StringInput `pulumi:"algorithm"`
	Value     pulumi.StringInput `pulumi:"value"`
}

func (ContentHashResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHashResponse)(nil)).Elem()
}

func (i ContentHashResponseArgs) ToContentHashResponseOutput() ContentHashResponseOutput {
	return i.ToContentHashResponseOutputWithContext(context.Background())
}

func (i ContentHashResponseArgs) ToContentHashResponseOutputWithContext(ctx context.Context) ContentHashResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashResponseOutput)
}

func (i ContentHashResponseArgs) ToContentHashResponsePtrOutput() ContentHashResponsePtrOutput {
	return i.ToContentHashResponsePtrOutputWithContext(context.Background())
}

func (i ContentHashResponseArgs) ToContentHashResponsePtrOutputWithContext(ctx context.Context) ContentHashResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashResponseOutput).ToContentHashResponsePtrOutputWithContext(ctx)
}









type ContentHashResponsePtrInput interface {
	pulumi.Input

	ToContentHashResponsePtrOutput() ContentHashResponsePtrOutput
	ToContentHashResponsePtrOutputWithContext(context.Context) ContentHashResponsePtrOutput
}

type contentHashResponsePtrType ContentHashResponseArgs

func ContentHashResponsePtr(v *ContentHashResponseArgs) ContentHashResponsePtrInput {
	return (*contentHashResponsePtrType)(v)
}

func (*contentHashResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHashResponse)(nil)).Elem()
}

func (i *contentHashResponsePtrType) ToContentHashResponsePtrOutput() ContentHashResponsePtrOutput {
	return i.ToContentHashResponsePtrOutputWithContext(context.Background())
}

func (i *contentHashResponsePtrType) ToContentHashResponsePtrOutputWithContext(ctx context.Context) ContentHashResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashResponsePtrOutput)
}

type ContentHashResponseOutput struct{ *pulumi.OutputState }

func (ContentHashResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHashResponse)(nil)).Elem()
}

func (o ContentHashResponseOutput) ToContentHashResponseOutput() ContentHashResponseOutput {
	return o
}

func (o ContentHashResponseOutput) ToContentHashResponseOutputWithContext(ctx context.Context) ContentHashResponseOutput {
	return o
}

func (o ContentHashResponseOutput) ToContentHashResponsePtrOutput() ContentHashResponsePtrOutput {
	return o.ToContentHashResponsePtrOutputWithContext(context.Background())
}

func (o ContentHashResponseOutput) ToContentHashResponsePtrOutputWithContext(ctx context.Context) ContentHashResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentHashResponse) *ContentHashResponse {
		return &v
	}).(ContentHashResponsePtrOutput)
}

func (o ContentHashResponseOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHashResponse) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o ContentHashResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHashResponse) string { return v.Value }).(pulumi.StringOutput)
}

type ContentHashResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentHashResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHashResponse)(nil)).Elem()
}

func (o ContentHashResponsePtrOutput) ToContentHashResponsePtrOutput() ContentHashResponsePtrOutput {
	return o
}

func (o ContentHashResponsePtrOutput) ToContentHashResponsePtrOutputWithContext(ctx context.Context) ContentHashResponsePtrOutput {
	return o
}

func (o ContentHashResponsePtrOutput) Elem() ContentHashResponseOutput {
	return o.ApplyT(func(v *ContentHashResponse) ContentHashResponse {
		if v != nil {
			return *v
		}
		var ret ContentHashResponse
		return ret
	}).(ContentHashResponseOutput)
}

func (o ContentHashResponsePtrOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHashResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Algorithm
	}).(pulumi.StringPtrOutput)
}

func (o ContentHashResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHashResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type ContentLink struct {
	ContentHash *ContentHash `pulumi:"contentHash"`
	Uri         *string      `pulumi:"uri"`
	Version     *string      `pulumi:"version"`
}





type ContentLinkInput interface {
	pulumi.Input

	ToContentLinkOutput() ContentLinkOutput
	ToContentLinkOutputWithContext(context.Context) ContentLinkOutput
}

type ContentLinkArgs struct {
	ContentHash ContentHashPtrInput   `pulumi:"contentHash"`
	Uri         pulumi.StringPtrInput `pulumi:"uri"`
	Version     pulumi.StringPtrInput `pulumi:"version"`
}

func (ContentLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLink)(nil)).Elem()
}

func (i ContentLinkArgs) ToContentLinkOutput() ContentLinkOutput {
	return i.ToContentLinkOutputWithContext(context.Background())
}

func (i ContentLinkArgs) ToContentLinkOutputWithContext(ctx context.Context) ContentLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkOutput)
}

func (i ContentLinkArgs) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return i.ToContentLinkPtrOutputWithContext(context.Background())
}

func (i ContentLinkArgs) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkOutput).ToContentLinkPtrOutputWithContext(ctx)
}









type ContentLinkPtrInput interface {
	pulumi.Input

	ToContentLinkPtrOutput() ContentLinkPtrOutput
	ToContentLinkPtrOutputWithContext(context.Context) ContentLinkPtrOutput
}

type contentLinkPtrType ContentLinkArgs

func ContentLinkPtr(v *ContentLinkArgs) ContentLinkPtrInput {
	return (*contentLinkPtrType)(v)
}

func (*contentLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLink)(nil)).Elem()
}

func (i *contentLinkPtrType) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return i.ToContentLinkPtrOutputWithContext(context.Background())
}

func (i *contentLinkPtrType) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkPtrOutput)
}

type ContentLinkOutput struct{ *pulumi.OutputState }

func (ContentLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLink)(nil)).Elem()
}

func (o ContentLinkOutput) ToContentLinkOutput() ContentLinkOutput {
	return o
}

func (o ContentLinkOutput) ToContentLinkOutputWithContext(ctx context.Context) ContentLinkOutput {
	return o
}

func (o ContentLinkOutput) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return o.ToContentLinkPtrOutputWithContext(context.Background())
}

func (o ContentLinkOutput) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentLink) *ContentLink {
		return &v
	}).(ContentLinkPtrOutput)
}

func (o ContentLinkOutput) ContentHash() ContentHashPtrOutput {
	return o.ApplyT(func(v ContentLink) *ContentHash { return v.ContentHash }).(ContentHashPtrOutput)
}

func (o ContentLinkOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLink) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

func (o ContentLinkOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLink) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ContentLinkPtrOutput struct{ *pulumi.OutputState }

func (ContentLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLink)(nil)).Elem()
}

func (o ContentLinkPtrOutput) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return o
}

func (o ContentLinkPtrOutput) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return o
}

func (o ContentLinkPtrOutput) Elem() ContentLinkOutput {
	return o.ApplyT(func(v *ContentLink) ContentLink {
		if v != nil {
			return *v
		}
		var ret ContentLink
		return ret
	}).(ContentLinkOutput)
}

func (o ContentLinkPtrOutput) ContentHash() ContentHashPtrOutput {
	return o.ApplyT(func(v *ContentLink) *ContentHash {
		if v == nil {
			return nil
		}
		return v.ContentHash
	}).(ContentHashPtrOutput)
}

func (o ContentLinkPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLink) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

func (o ContentLinkPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLink) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ContentLinkResponse struct {
	ContentHash *ContentHashResponse `pulumi:"contentHash"`
	Uri         *string              `pulumi:"uri"`
	Version     *string              `pulumi:"version"`
}





type ContentLinkResponseInput interface {
	pulumi.Input

	ToContentLinkResponseOutput() ContentLinkResponseOutput
	ToContentLinkResponseOutputWithContext(context.Context) ContentLinkResponseOutput
}

type ContentLinkResponseArgs struct {
	ContentHash ContentHashResponsePtrInput `pulumi:"contentHash"`
	Uri         pulumi.StringPtrInput       `pulumi:"uri"`
	Version     pulumi.StringPtrInput       `pulumi:"version"`
}

func (ContentLinkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLinkResponse)(nil)).Elem()
}

func (i ContentLinkResponseArgs) ToContentLinkResponseOutput() ContentLinkResponseOutput {
	return i.ToContentLinkResponseOutputWithContext(context.Background())
}

func (i ContentLinkResponseArgs) ToContentLinkResponseOutputWithContext(ctx context.Context) ContentLinkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkResponseOutput)
}

func (i ContentLinkResponseArgs) ToContentLinkResponsePtrOutput() ContentLinkResponsePtrOutput {
	return i.ToContentLinkResponsePtrOutputWithContext(context.Background())
}

func (i ContentLinkResponseArgs) ToContentLinkResponsePtrOutputWithContext(ctx context.Context) ContentLinkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkResponseOutput).ToContentLinkResponsePtrOutputWithContext(ctx)
}









type ContentLinkResponsePtrInput interface {
	pulumi.Input

	ToContentLinkResponsePtrOutput() ContentLinkResponsePtrOutput
	ToContentLinkResponsePtrOutputWithContext(context.Context) ContentLinkResponsePtrOutput
}

type contentLinkResponsePtrType ContentLinkResponseArgs

func ContentLinkResponsePtr(v *ContentLinkResponseArgs) ContentLinkResponsePtrInput {
	return (*contentLinkResponsePtrType)(v)
}

func (*contentLinkResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLinkResponse)(nil)).Elem()
}

func (i *contentLinkResponsePtrType) ToContentLinkResponsePtrOutput() ContentLinkResponsePtrOutput {
	return i.ToContentLinkResponsePtrOutputWithContext(context.Background())
}

func (i *contentLinkResponsePtrType) ToContentLinkResponsePtrOutputWithContext(ctx context.Context) ContentLinkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkResponsePtrOutput)
}

type ContentLinkResponseOutput struct{ *pulumi.OutputState }

func (ContentLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLinkResponse)(nil)).Elem()
}

func (o ContentLinkResponseOutput) ToContentLinkResponseOutput() ContentLinkResponseOutput {
	return o
}

func (o ContentLinkResponseOutput) ToContentLinkResponseOutputWithContext(ctx context.Context) ContentLinkResponseOutput {
	return o
}

func (o ContentLinkResponseOutput) ToContentLinkResponsePtrOutput() ContentLinkResponsePtrOutput {
	return o.ToContentLinkResponsePtrOutputWithContext(context.Background())
}

func (o ContentLinkResponseOutput) ToContentLinkResponsePtrOutputWithContext(ctx context.Context) ContentLinkResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentLinkResponse) *ContentLinkResponse {
		return &v
	}).(ContentLinkResponsePtrOutput)
}

func (o ContentLinkResponseOutput) ContentHash() ContentHashResponsePtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *ContentHashResponse { return v.ContentHash }).(ContentHashResponsePtrOutput)
}

func (o ContentLinkResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

func (o ContentLinkResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ContentLinkResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentLinkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLinkResponse)(nil)).Elem()
}

func (o ContentLinkResponsePtrOutput) ToContentLinkResponsePtrOutput() ContentLinkResponsePtrOutput {
	return o
}

func (o ContentLinkResponsePtrOutput) ToContentLinkResponsePtrOutputWithContext(ctx context.Context) ContentLinkResponsePtrOutput {
	return o
}

func (o ContentLinkResponsePtrOutput) Elem() ContentLinkResponseOutput {
	return o.ApplyT(func(v *ContentLinkResponse) ContentLinkResponse {
		if v != nil {
			return *v
		}
		var ret ContentLinkResponse
		return ret
	}).(ContentLinkResponseOutput)
}

func (o ContentLinkResponsePtrOutput) ContentHash() ContentHashResponsePtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *ContentHashResponse {
		if v == nil {
			return nil
		}
		return v.ContentHash
	}).(ContentHashResponsePtrOutput)
}

func (o ContentLinkResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

func (o ContentLinkResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ModuleErrorInfoResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}





type ModuleErrorInfoResponseInput interface {
	pulumi.Input

	ToModuleErrorInfoResponseOutput() ModuleErrorInfoResponseOutput
	ToModuleErrorInfoResponseOutputWithContext(context.Context) ModuleErrorInfoResponseOutput
}

type ModuleErrorInfoResponseArgs struct {
	Code    pulumi.StringPtrInput `pulumi:"code"`
	Message pulumi.StringPtrInput `pulumi:"message"`
}

func (ModuleErrorInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleErrorInfoResponse)(nil)).Elem()
}

func (i ModuleErrorInfoResponseArgs) ToModuleErrorInfoResponseOutput() ModuleErrorInfoResponseOutput {
	return i.ToModuleErrorInfoResponseOutputWithContext(context.Background())
}

func (i ModuleErrorInfoResponseArgs) ToModuleErrorInfoResponseOutputWithContext(ctx context.Context) ModuleErrorInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleErrorInfoResponseOutput)
}

func (i ModuleErrorInfoResponseArgs) ToModuleErrorInfoResponsePtrOutput() ModuleErrorInfoResponsePtrOutput {
	return i.ToModuleErrorInfoResponsePtrOutputWithContext(context.Background())
}

func (i ModuleErrorInfoResponseArgs) ToModuleErrorInfoResponsePtrOutputWithContext(ctx context.Context) ModuleErrorInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleErrorInfoResponseOutput).ToModuleErrorInfoResponsePtrOutputWithContext(ctx)
}









type ModuleErrorInfoResponsePtrInput interface {
	pulumi.Input

	ToModuleErrorInfoResponsePtrOutput() ModuleErrorInfoResponsePtrOutput
	ToModuleErrorInfoResponsePtrOutputWithContext(context.Context) ModuleErrorInfoResponsePtrOutput
}

type moduleErrorInfoResponsePtrType ModuleErrorInfoResponseArgs

func ModuleErrorInfoResponsePtr(v *ModuleErrorInfoResponseArgs) ModuleErrorInfoResponsePtrInput {
	return (*moduleErrorInfoResponsePtrType)(v)
}

func (*moduleErrorInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleErrorInfoResponse)(nil)).Elem()
}

func (i *moduleErrorInfoResponsePtrType) ToModuleErrorInfoResponsePtrOutput() ModuleErrorInfoResponsePtrOutput {
	return i.ToModuleErrorInfoResponsePtrOutputWithContext(context.Background())
}

func (i *moduleErrorInfoResponsePtrType) ToModuleErrorInfoResponsePtrOutputWithContext(ctx context.Context) ModuleErrorInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleErrorInfoResponsePtrOutput)
}

type ModuleErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (ModuleErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleErrorInfoResponse)(nil)).Elem()
}

func (o ModuleErrorInfoResponseOutput) ToModuleErrorInfoResponseOutput() ModuleErrorInfoResponseOutput {
	return o
}

func (o ModuleErrorInfoResponseOutput) ToModuleErrorInfoResponseOutputWithContext(ctx context.Context) ModuleErrorInfoResponseOutput {
	return o
}

func (o ModuleErrorInfoResponseOutput) ToModuleErrorInfoResponsePtrOutput() ModuleErrorInfoResponsePtrOutput {
	return o.ToModuleErrorInfoResponsePtrOutputWithContext(context.Background())
}

func (o ModuleErrorInfoResponseOutput) ToModuleErrorInfoResponsePtrOutputWithContext(ctx context.Context) ModuleErrorInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModuleErrorInfoResponse) *ModuleErrorInfoResponse {
		return &v
	}).(ModuleErrorInfoResponsePtrOutput)
}

func (o ModuleErrorInfoResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModuleErrorInfoResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ModuleErrorInfoResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModuleErrorInfoResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type ModuleErrorInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ModuleErrorInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleErrorInfoResponse)(nil)).Elem()
}

func (o ModuleErrorInfoResponsePtrOutput) ToModuleErrorInfoResponsePtrOutput() ModuleErrorInfoResponsePtrOutput {
	return o
}

func (o ModuleErrorInfoResponsePtrOutput) ToModuleErrorInfoResponsePtrOutputWithContext(ctx context.Context) ModuleErrorInfoResponsePtrOutput {
	return o
}

func (o ModuleErrorInfoResponsePtrOutput) Elem() ModuleErrorInfoResponseOutput {
	return o.ApplyT(func(v *ModuleErrorInfoResponse) ModuleErrorInfoResponse {
		if v != nil {
			return *v
		}
		var ret ModuleErrorInfoResponse
		return ret
	}).(ModuleErrorInfoResponseOutput)
}

func (o ModuleErrorInfoResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModuleErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ModuleErrorInfoResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModuleErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type RunbookDraft struct {
	CreationTime     *string                     `pulumi:"creationTime"`
	DraftContentLink *ContentLink                `pulumi:"draftContentLink"`
	InEdit           *bool                       `pulumi:"inEdit"`
	LastModifiedTime *string                     `pulumi:"lastModifiedTime"`
	OutputTypes      []string                    `pulumi:"outputTypes"`
	Parameters       map[string]RunbookParameter `pulumi:"parameters"`
}





type RunbookDraftInput interface {
	pulumi.Input

	ToRunbookDraftOutput() RunbookDraftOutput
	ToRunbookDraftOutputWithContext(context.Context) RunbookDraftOutput
}

type RunbookDraftArgs struct {
	CreationTime     pulumi.StringPtrInput    `pulumi:"creationTime"`
	DraftContentLink ContentLinkPtrInput      `pulumi:"draftContentLink"`
	InEdit           pulumi.BoolPtrInput      `pulumi:"inEdit"`
	LastModifiedTime pulumi.StringPtrInput    `pulumi:"lastModifiedTime"`
	OutputTypes      pulumi.StringArrayInput  `pulumi:"outputTypes"`
	Parameters       RunbookParameterMapInput `pulumi:"parameters"`
}

func (RunbookDraftArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookDraft)(nil)).Elem()
}

func (i RunbookDraftArgs) ToRunbookDraftOutput() RunbookDraftOutput {
	return i.ToRunbookDraftOutputWithContext(context.Background())
}

func (i RunbookDraftArgs) ToRunbookDraftOutputWithContext(ctx context.Context) RunbookDraftOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftOutput)
}

func (i RunbookDraftArgs) ToRunbookDraftPtrOutput() RunbookDraftPtrOutput {
	return i.ToRunbookDraftPtrOutputWithContext(context.Background())
}

func (i RunbookDraftArgs) ToRunbookDraftPtrOutputWithContext(ctx context.Context) RunbookDraftPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftOutput).ToRunbookDraftPtrOutputWithContext(ctx)
}









type RunbookDraftPtrInput interface {
	pulumi.Input

	ToRunbookDraftPtrOutput() RunbookDraftPtrOutput
	ToRunbookDraftPtrOutputWithContext(context.Context) RunbookDraftPtrOutput
}

type runbookDraftPtrType RunbookDraftArgs

func RunbookDraftPtr(v *RunbookDraftArgs) RunbookDraftPtrInput {
	return (*runbookDraftPtrType)(v)
}

func (*runbookDraftPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookDraft)(nil)).Elem()
}

func (i *runbookDraftPtrType) ToRunbookDraftPtrOutput() RunbookDraftPtrOutput {
	return i.ToRunbookDraftPtrOutputWithContext(context.Background())
}

func (i *runbookDraftPtrType) ToRunbookDraftPtrOutputWithContext(ctx context.Context) RunbookDraftPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftPtrOutput)
}

type RunbookDraftOutput struct{ *pulumi.OutputState }

func (RunbookDraftOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookDraft)(nil)).Elem()
}

func (o RunbookDraftOutput) ToRunbookDraftOutput() RunbookDraftOutput {
	return o
}

func (o RunbookDraftOutput) ToRunbookDraftOutputWithContext(ctx context.Context) RunbookDraftOutput {
	return o
}

func (o RunbookDraftOutput) ToRunbookDraftPtrOutput() RunbookDraftPtrOutput {
	return o.ToRunbookDraftPtrOutputWithContext(context.Background())
}

func (o RunbookDraftOutput) ToRunbookDraftPtrOutputWithContext(ctx context.Context) RunbookDraftPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunbookDraft) *RunbookDraft {
		return &v
	}).(RunbookDraftPtrOutput)
}

func (o RunbookDraftOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookDraft) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o RunbookDraftOutput) DraftContentLink() ContentLinkPtrOutput {
	return o.ApplyT(func(v RunbookDraft) *ContentLink { return v.DraftContentLink }).(ContentLinkPtrOutput)
}

func (o RunbookDraftOutput) InEdit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RunbookDraft) *bool { return v.InEdit }).(pulumi.BoolPtrOutput)
}

func (o RunbookDraftOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookDraft) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o RunbookDraftOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RunbookDraft) []string { return v.OutputTypes }).(pulumi.StringArrayOutput)
}

func (o RunbookDraftOutput) Parameters() RunbookParameterMapOutput {
	return o.ApplyT(func(v RunbookDraft) map[string]RunbookParameter { return v.Parameters }).(RunbookParameterMapOutput)
}

type RunbookDraftPtrOutput struct{ *pulumi.OutputState }

func (RunbookDraftPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookDraft)(nil)).Elem()
}

func (o RunbookDraftPtrOutput) ToRunbookDraftPtrOutput() RunbookDraftPtrOutput {
	return o
}

func (o RunbookDraftPtrOutput) ToRunbookDraftPtrOutputWithContext(ctx context.Context) RunbookDraftPtrOutput {
	return o
}

func (o RunbookDraftPtrOutput) Elem() RunbookDraftOutput {
	return o.ApplyT(func(v *RunbookDraft) RunbookDraft {
		if v != nil {
			return *v
		}
		var ret RunbookDraft
		return ret
	}).(RunbookDraftOutput)
}

func (o RunbookDraftPtrOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookDraft) *string {
		if v == nil {
			return nil
		}
		return v.CreationTime
	}).(pulumi.StringPtrOutput)
}

func (o RunbookDraftPtrOutput) DraftContentLink() ContentLinkPtrOutput {
	return o.ApplyT(func(v *RunbookDraft) *ContentLink {
		if v == nil {
			return nil
		}
		return v.DraftContentLink
	}).(ContentLinkPtrOutput)
}

func (o RunbookDraftPtrOutput) InEdit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RunbookDraft) *bool {
		if v == nil {
			return nil
		}
		return v.InEdit
	}).(pulumi.BoolPtrOutput)
}

func (o RunbookDraftPtrOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookDraft) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedTime
	}).(pulumi.StringPtrOutput)
}

func (o RunbookDraftPtrOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RunbookDraft) []string {
		if v == nil {
			return nil
		}
		return v.OutputTypes
	}).(pulumi.StringArrayOutput)
}

func (o RunbookDraftPtrOutput) Parameters() RunbookParameterMapOutput {
	return o.ApplyT(func(v *RunbookDraft) map[string]RunbookParameter {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(RunbookParameterMapOutput)
}

type RunbookDraftResponse struct {
	CreationTime     *string                             `pulumi:"creationTime"`
	DraftContentLink *ContentLinkResponse                `pulumi:"draftContentLink"`
	InEdit           *bool                               `pulumi:"inEdit"`
	LastModifiedTime *string                             `pulumi:"lastModifiedTime"`
	OutputTypes      []string                            `pulumi:"outputTypes"`
	Parameters       map[string]RunbookParameterResponse `pulumi:"parameters"`
}





type RunbookDraftResponseInput interface {
	pulumi.Input

	ToRunbookDraftResponseOutput() RunbookDraftResponseOutput
	ToRunbookDraftResponseOutputWithContext(context.Context) RunbookDraftResponseOutput
}

type RunbookDraftResponseArgs struct {
	CreationTime     pulumi.StringPtrInput            `pulumi:"creationTime"`
	DraftContentLink ContentLinkResponsePtrInput      `pulumi:"draftContentLink"`
	InEdit           pulumi.BoolPtrInput              `pulumi:"inEdit"`
	LastModifiedTime pulumi.StringPtrInput            `pulumi:"lastModifiedTime"`
	OutputTypes      pulumi.StringArrayInput          `pulumi:"outputTypes"`
	Parameters       RunbookParameterResponseMapInput `pulumi:"parameters"`
}

func (RunbookDraftResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookDraftResponse)(nil)).Elem()
}

func (i RunbookDraftResponseArgs) ToRunbookDraftResponseOutput() RunbookDraftResponseOutput {
	return i.ToRunbookDraftResponseOutputWithContext(context.Background())
}

func (i RunbookDraftResponseArgs) ToRunbookDraftResponseOutputWithContext(ctx context.Context) RunbookDraftResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftResponseOutput)
}

func (i RunbookDraftResponseArgs) ToRunbookDraftResponsePtrOutput() RunbookDraftResponsePtrOutput {
	return i.ToRunbookDraftResponsePtrOutputWithContext(context.Background())
}

func (i RunbookDraftResponseArgs) ToRunbookDraftResponsePtrOutputWithContext(ctx context.Context) RunbookDraftResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftResponseOutput).ToRunbookDraftResponsePtrOutputWithContext(ctx)
}









type RunbookDraftResponsePtrInput interface {
	pulumi.Input

	ToRunbookDraftResponsePtrOutput() RunbookDraftResponsePtrOutput
	ToRunbookDraftResponsePtrOutputWithContext(context.Context) RunbookDraftResponsePtrOutput
}

type runbookDraftResponsePtrType RunbookDraftResponseArgs

func RunbookDraftResponsePtr(v *RunbookDraftResponseArgs) RunbookDraftResponsePtrInput {
	return (*runbookDraftResponsePtrType)(v)
}

func (*runbookDraftResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookDraftResponse)(nil)).Elem()
}

func (i *runbookDraftResponsePtrType) ToRunbookDraftResponsePtrOutput() RunbookDraftResponsePtrOutput {
	return i.ToRunbookDraftResponsePtrOutputWithContext(context.Background())
}

func (i *runbookDraftResponsePtrType) ToRunbookDraftResponsePtrOutputWithContext(ctx context.Context) RunbookDraftResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftResponsePtrOutput)
}

type RunbookDraftResponseOutput struct{ *pulumi.OutputState }

func (RunbookDraftResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookDraftResponse)(nil)).Elem()
}

func (o RunbookDraftResponseOutput) ToRunbookDraftResponseOutput() RunbookDraftResponseOutput {
	return o
}

func (o RunbookDraftResponseOutput) ToRunbookDraftResponseOutputWithContext(ctx context.Context) RunbookDraftResponseOutput {
	return o
}

func (o RunbookDraftResponseOutput) ToRunbookDraftResponsePtrOutput() RunbookDraftResponsePtrOutput {
	return o.ToRunbookDraftResponsePtrOutputWithContext(context.Background())
}

func (o RunbookDraftResponseOutput) ToRunbookDraftResponsePtrOutputWithContext(ctx context.Context) RunbookDraftResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunbookDraftResponse) *RunbookDraftResponse {
		return &v
	}).(RunbookDraftResponsePtrOutput)
}

func (o RunbookDraftResponseOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookDraftResponse) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o RunbookDraftResponseOutput) DraftContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v RunbookDraftResponse) *ContentLinkResponse { return v.DraftContentLink }).(ContentLinkResponsePtrOutput)
}

func (o RunbookDraftResponseOutput) InEdit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RunbookDraftResponse) *bool { return v.InEdit }).(pulumi.BoolPtrOutput)
}

func (o RunbookDraftResponseOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookDraftResponse) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o RunbookDraftResponseOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RunbookDraftResponse) []string { return v.OutputTypes }).(pulumi.StringArrayOutput)
}

func (o RunbookDraftResponseOutput) Parameters() RunbookParameterResponseMapOutput {
	return o.ApplyT(func(v RunbookDraftResponse) map[string]RunbookParameterResponse { return v.Parameters }).(RunbookParameterResponseMapOutput)
}

type RunbookDraftResponsePtrOutput struct{ *pulumi.OutputState }

func (RunbookDraftResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookDraftResponse)(nil)).Elem()
}

func (o RunbookDraftResponsePtrOutput) ToRunbookDraftResponsePtrOutput() RunbookDraftResponsePtrOutput {
	return o
}

func (o RunbookDraftResponsePtrOutput) ToRunbookDraftResponsePtrOutputWithContext(ctx context.Context) RunbookDraftResponsePtrOutput {
	return o
}

func (o RunbookDraftResponsePtrOutput) Elem() RunbookDraftResponseOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) RunbookDraftResponse {
		if v != nil {
			return *v
		}
		var ret RunbookDraftResponse
		return ret
	}).(RunbookDraftResponseOutput)
}

func (o RunbookDraftResponsePtrOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreationTime
	}).(pulumi.StringPtrOutput)
}

func (o RunbookDraftResponsePtrOutput) DraftContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) *ContentLinkResponse {
		if v == nil {
			return nil
		}
		return v.DraftContentLink
	}).(ContentLinkResponsePtrOutput)
}

func (o RunbookDraftResponsePtrOutput) InEdit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) *bool {
		if v == nil {
			return nil
		}
		return v.InEdit
	}).(pulumi.BoolPtrOutput)
}

func (o RunbookDraftResponsePtrOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedTime
	}).(pulumi.StringPtrOutput)
}

func (o RunbookDraftResponsePtrOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) []string {
		if v == nil {
			return nil
		}
		return v.OutputTypes
	}).(pulumi.StringArrayOutput)
}

func (o RunbookDraftResponsePtrOutput) Parameters() RunbookParameterResponseMapOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) map[string]RunbookParameterResponse {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(RunbookParameterResponseMapOutput)
}

type RunbookParameter struct {
	DefaultValue *string `pulumi:"defaultValue"`
	IsMandatory  *bool   `pulumi:"isMandatory"`
	Position     *int    `pulumi:"position"`
	Type         *string `pulumi:"type"`
}





type RunbookParameterInput interface {
	pulumi.Input

	ToRunbookParameterOutput() RunbookParameterOutput
	ToRunbookParameterOutputWithContext(context.Context) RunbookParameterOutput
}

type RunbookParameterArgs struct {
	DefaultValue pulumi.StringPtrInput `pulumi:"defaultValue"`
	IsMandatory  pulumi.BoolPtrInput   `pulumi:"isMandatory"`
	Position     pulumi.IntPtrInput    `pulumi:"position"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (RunbookParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookParameter)(nil)).Elem()
}

func (i RunbookParameterArgs) ToRunbookParameterOutput() RunbookParameterOutput {
	return i.ToRunbookParameterOutputWithContext(context.Background())
}

func (i RunbookParameterArgs) ToRunbookParameterOutputWithContext(ctx context.Context) RunbookParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookParameterOutput)
}





type RunbookParameterMapInput interface {
	pulumi.Input

	ToRunbookParameterMapOutput() RunbookParameterMapOutput
	ToRunbookParameterMapOutputWithContext(context.Context) RunbookParameterMapOutput
}

type RunbookParameterMap map[string]RunbookParameterInput

func (RunbookParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RunbookParameter)(nil)).Elem()
}

func (i RunbookParameterMap) ToRunbookParameterMapOutput() RunbookParameterMapOutput {
	return i.ToRunbookParameterMapOutputWithContext(context.Background())
}

func (i RunbookParameterMap) ToRunbookParameterMapOutputWithContext(ctx context.Context) RunbookParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookParameterMapOutput)
}

type RunbookParameterOutput struct{ *pulumi.OutputState }

func (RunbookParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookParameter)(nil)).Elem()
}

func (o RunbookParameterOutput) ToRunbookParameterOutput() RunbookParameterOutput {
	return o
}

func (o RunbookParameterOutput) ToRunbookParameterOutputWithContext(ctx context.Context) RunbookParameterOutput {
	return o
}

func (o RunbookParameterOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookParameter) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

func (o RunbookParameterOutput) IsMandatory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RunbookParameter) *bool { return v.IsMandatory }).(pulumi.BoolPtrOutput)
}

func (o RunbookParameterOutput) Position() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RunbookParameter) *int { return v.Position }).(pulumi.IntPtrOutput)
}

func (o RunbookParameterOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookParameter) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type RunbookParameterMapOutput struct{ *pulumi.OutputState }

func (RunbookParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RunbookParameter)(nil)).Elem()
}

func (o RunbookParameterMapOutput) ToRunbookParameterMapOutput() RunbookParameterMapOutput {
	return o
}

func (o RunbookParameterMapOutput) ToRunbookParameterMapOutputWithContext(ctx context.Context) RunbookParameterMapOutput {
	return o
}

func (o RunbookParameterMapOutput) MapIndex(k pulumi.StringInput) RunbookParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RunbookParameter {
		return vs[0].(map[string]RunbookParameter)[vs[1].(string)]
	}).(RunbookParameterOutput)
}

type RunbookParameterResponse struct {
	DefaultValue *string `pulumi:"defaultValue"`
	IsMandatory  *bool   `pulumi:"isMandatory"`
	Position     *int    `pulumi:"position"`
	Type         *string `pulumi:"type"`
}





type RunbookParameterResponseInput interface {
	pulumi.Input

	ToRunbookParameterResponseOutput() RunbookParameterResponseOutput
	ToRunbookParameterResponseOutputWithContext(context.Context) RunbookParameterResponseOutput
}

type RunbookParameterResponseArgs struct {
	DefaultValue pulumi.StringPtrInput `pulumi:"defaultValue"`
	IsMandatory  pulumi.BoolPtrInput   `pulumi:"isMandatory"`
	Position     pulumi.IntPtrInput    `pulumi:"position"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (RunbookParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookParameterResponse)(nil)).Elem()
}

func (i RunbookParameterResponseArgs) ToRunbookParameterResponseOutput() RunbookParameterResponseOutput {
	return i.ToRunbookParameterResponseOutputWithContext(context.Background())
}

func (i RunbookParameterResponseArgs) ToRunbookParameterResponseOutputWithContext(ctx context.Context) RunbookParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookParameterResponseOutput)
}





type RunbookParameterResponseMapInput interface {
	pulumi.Input

	ToRunbookParameterResponseMapOutput() RunbookParameterResponseMapOutput
	ToRunbookParameterResponseMapOutputWithContext(context.Context) RunbookParameterResponseMapOutput
}

type RunbookParameterResponseMap map[string]RunbookParameterResponseInput

func (RunbookParameterResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RunbookParameterResponse)(nil)).Elem()
}

func (i RunbookParameterResponseMap) ToRunbookParameterResponseMapOutput() RunbookParameterResponseMapOutput {
	return i.ToRunbookParameterResponseMapOutputWithContext(context.Background())
}

func (i RunbookParameterResponseMap) ToRunbookParameterResponseMapOutputWithContext(ctx context.Context) RunbookParameterResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookParameterResponseMapOutput)
}

type RunbookParameterResponseOutput struct{ *pulumi.OutputState }

func (RunbookParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookParameterResponse)(nil)).Elem()
}

func (o RunbookParameterResponseOutput) ToRunbookParameterResponseOutput() RunbookParameterResponseOutput {
	return o
}

func (o RunbookParameterResponseOutput) ToRunbookParameterResponseOutputWithContext(ctx context.Context) RunbookParameterResponseOutput {
	return o
}

func (o RunbookParameterResponseOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookParameterResponse) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

func (o RunbookParameterResponseOutput) IsMandatory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RunbookParameterResponse) *bool { return v.IsMandatory }).(pulumi.BoolPtrOutput)
}

func (o RunbookParameterResponseOutput) Position() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RunbookParameterResponse) *int { return v.Position }).(pulumi.IntPtrOutput)
}

func (o RunbookParameterResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookParameterResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type RunbookParameterResponseMapOutput struct{ *pulumi.OutputState }

func (RunbookParameterResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RunbookParameterResponse)(nil)).Elem()
}

func (o RunbookParameterResponseMapOutput) ToRunbookParameterResponseMapOutput() RunbookParameterResponseMapOutput {
	return o
}

func (o RunbookParameterResponseMapOutput) ToRunbookParameterResponseMapOutputWithContext(ctx context.Context) RunbookParameterResponseMapOutput {
	return o
}

func (o RunbookParameterResponseMapOutput) MapIndex(k pulumi.StringInput) RunbookParameterResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RunbookParameterResponse {
		return vs[0].(map[string]RunbookParameterResponse)[vs[1].(string)]
	}).(RunbookParameterResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContentHashInput)(nil)).Elem(), ContentHashArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentHashPtrInput)(nil)).Elem(), ContentHashArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentHashResponseInput)(nil)).Elem(), ContentHashResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentHashResponsePtrInput)(nil)).Elem(), ContentHashResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentLinkInput)(nil)).Elem(), ContentLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentLinkPtrInput)(nil)).Elem(), ContentLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentLinkResponseInput)(nil)).Elem(), ContentLinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentLinkResponsePtrInput)(nil)).Elem(), ContentLinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModuleErrorInfoResponseInput)(nil)).Elem(), ModuleErrorInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModuleErrorInfoResponsePtrInput)(nil)).Elem(), ModuleErrorInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookDraftInput)(nil)).Elem(), RunbookDraftArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookDraftPtrInput)(nil)).Elem(), RunbookDraftArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookDraftResponseInput)(nil)).Elem(), RunbookDraftResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookDraftResponsePtrInput)(nil)).Elem(), RunbookDraftResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookParameterInput)(nil)).Elem(), RunbookParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookParameterMapInput)(nil)).Elem(), RunbookParameterMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookParameterResponseInput)(nil)).Elem(), RunbookParameterResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookParameterResponseMapInput)(nil)).Elem(), RunbookParameterResponseMap{})
	pulumi.RegisterOutputType(ContentHashOutput{})
	pulumi.RegisterOutputType(ContentHashPtrOutput{})
	pulumi.RegisterOutputType(ContentHashResponseOutput{})
	pulumi.RegisterOutputType(ContentHashResponsePtrOutput{})
	pulumi.RegisterOutputType(ContentLinkOutput{})
	pulumi.RegisterOutputType(ContentLinkPtrOutput{})
	pulumi.RegisterOutputType(ContentLinkResponseOutput{})
	pulumi.RegisterOutputType(ContentLinkResponsePtrOutput{})
	pulumi.RegisterOutputType(ModuleErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(ModuleErrorInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(RunbookDraftOutput{})
	pulumi.RegisterOutputType(RunbookDraftPtrOutput{})
	pulumi.RegisterOutputType(RunbookDraftResponseOutput{})
	pulumi.RegisterOutputType(RunbookDraftResponsePtrOutput{})
	pulumi.RegisterOutputType(RunbookParameterOutput{})
	pulumi.RegisterOutputType(RunbookParameterMapOutput{})
	pulumi.RegisterOutputType(RunbookParameterResponseOutput{})
	pulumi.RegisterOutputType(RunbookParameterResponseMapOutput{})
}
