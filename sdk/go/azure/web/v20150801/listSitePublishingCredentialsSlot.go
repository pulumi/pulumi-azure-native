


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSitePublishingCredentialsSlot(ctx *pulumi.Context, args *ListSitePublishingCredentialsSlotArgs, opts ...pulumi.InvokeOption) (*ListSitePublishingCredentialsSlotResult, error) {
	var rv ListSitePublishingCredentialsSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:listSitePublishingCredentialsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSitePublishingCredentialsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListSitePublishingCredentialsSlotResult struct {
	Id                 *string           `pulumi:"id"`
	Kind               *string           `pulumi:"kind"`
	Location           string            `pulumi:"location"`
	Name               *string           `pulumi:"name"`
	PublishingPassword *string           `pulumi:"publishingPassword"`
	PublishingUserName *string           `pulumi:"publishingUserName"`
	ScmUri             *string           `pulumi:"scmUri"`
	Tags               map[string]string `pulumi:"tags"`
	Type               *string           `pulumi:"type"`
}

func ListSitePublishingCredentialsSlotOutput(ctx *pulumi.Context, args ListSitePublishingCredentialsSlotOutputArgs, opts ...pulumi.InvokeOption) ListSitePublishingCredentialsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSitePublishingCredentialsSlotResult, error) {
			args := v.(ListSitePublishingCredentialsSlotArgs)
			r, err := ListSitePublishingCredentialsSlot(ctx, &args, opts...)
			var s ListSitePublishingCredentialsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSitePublishingCredentialsSlotResultOutput)
}

type ListSitePublishingCredentialsSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListSitePublishingCredentialsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSitePublishingCredentialsSlotArgs)(nil)).Elem()
}


type ListSitePublishingCredentialsSlotResultOutput struct{ *pulumi.OutputState }

func (ListSitePublishingCredentialsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSitePublishingCredentialsSlotResult)(nil)).Elem()
}

func (o ListSitePublishingCredentialsSlotResultOutput) ToListSitePublishingCredentialsSlotResultOutput() ListSitePublishingCredentialsSlotResultOutput {
	return o
}

func (o ListSitePublishingCredentialsSlotResultOutput) ToListSitePublishingCredentialsSlotResultOutputWithContext(ctx context.Context) ListSitePublishingCredentialsSlotResultOutput {
	return o
}

func (o ListSitePublishingCredentialsSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListSitePublishingCredentialsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListSitePublishingCredentialsSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListSitePublishingCredentialsSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListSitePublishingCredentialsSlotResultOutput) PublishingPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsSlotResult) *string { return v.PublishingPassword }).(pulumi.StringPtrOutput)
}

func (o ListSitePublishingCredentialsSlotResultOutput) PublishingUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsSlotResult) *string { return v.PublishingUserName }).(pulumi.StringPtrOutput)
}

func (o ListSitePublishingCredentialsSlotResultOutput) ScmUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsSlotResult) *string { return v.ScmUri }).(pulumi.StringPtrOutput)
}

func (o ListSitePublishingCredentialsSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListSitePublishingCredentialsSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSitePublishingCredentialsSlotResultOutput{})
}
