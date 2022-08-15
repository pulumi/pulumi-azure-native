


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppPublishingCredentialsSlot(ctx *pulumi.Context, args *ListWebAppPublishingCredentialsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppPublishingCredentialsSlotResult, error) {
	var rv ListWebAppPublishingCredentialsSlotResult
	err := ctx.Invoke("azure-native:web/v20180201:listWebAppPublishingCredentialsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppPublishingCredentialsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppPublishingCredentialsSlotResult struct {
	Id                         string  `pulumi:"id"`
	Kind                       *string `pulumi:"kind"`
	Name                       string  `pulumi:"name"`
	PublishingPassword         *string `pulumi:"publishingPassword"`
	PublishingPasswordHash     *string `pulumi:"publishingPasswordHash"`
	PublishingPasswordHashSalt *string `pulumi:"publishingPasswordHashSalt"`
	PublishingUserName         string  `pulumi:"publishingUserName"`
	ScmUri                     *string `pulumi:"scmUri"`
	Type                       string  `pulumi:"type"`
}

func ListWebAppPublishingCredentialsSlotOutput(ctx *pulumi.Context, args ListWebAppPublishingCredentialsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppPublishingCredentialsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppPublishingCredentialsSlotResult, error) {
			args := v.(ListWebAppPublishingCredentialsSlotArgs)
			r, err := ListWebAppPublishingCredentialsSlot(ctx, &args, opts...)
			var s ListWebAppPublishingCredentialsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppPublishingCredentialsSlotResultOutput)
}

type ListWebAppPublishingCredentialsSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppPublishingCredentialsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppPublishingCredentialsSlotArgs)(nil)).Elem()
}


type ListWebAppPublishingCredentialsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppPublishingCredentialsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppPublishingCredentialsSlotResult)(nil)).Elem()
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) ToListWebAppPublishingCredentialsSlotResultOutput() ListWebAppPublishingCredentialsSlotResultOutput {
	return o
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) ToListWebAppPublishingCredentialsSlotResultOutputWithContext(ctx context.Context) ListWebAppPublishingCredentialsSlotResultOutput {
	return o
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) PublishingPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) *string { return v.PublishingPassword }).(pulumi.StringPtrOutput)
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) PublishingPasswordHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) *string { return v.PublishingPasswordHash }).(pulumi.StringPtrOutput)
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) PublishingPasswordHashSalt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) *string { return v.PublishingPasswordHashSalt }).(pulumi.StringPtrOutput)
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) PublishingUserName() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) string { return v.PublishingUserName }).(pulumi.StringOutput)
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) ScmUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) *string { return v.ScmUri }).(pulumi.StringPtrOutput)
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppPublishingCredentialsSlotResultOutput{})
}
