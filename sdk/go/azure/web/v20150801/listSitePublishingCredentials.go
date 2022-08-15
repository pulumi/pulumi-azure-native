


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSitePublishingCredentials(ctx *pulumi.Context, args *ListSitePublishingCredentialsArgs, opts ...pulumi.InvokeOption) (*ListSitePublishingCredentialsResult, error) {
	var rv ListSitePublishingCredentialsResult
	err := ctx.Invoke("azure-native:web/v20150801:listSitePublishingCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSitePublishingCredentialsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListSitePublishingCredentialsResult struct {
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

func ListSitePublishingCredentialsOutput(ctx *pulumi.Context, args ListSitePublishingCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListSitePublishingCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSitePublishingCredentialsResult, error) {
			args := v.(ListSitePublishingCredentialsArgs)
			r, err := ListSitePublishingCredentials(ctx, &args, opts...)
			var s ListSitePublishingCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSitePublishingCredentialsResultOutput)
}

type ListSitePublishingCredentialsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListSitePublishingCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSitePublishingCredentialsArgs)(nil)).Elem()
}


type ListSitePublishingCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListSitePublishingCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSitePublishingCredentialsResult)(nil)).Elem()
}

func (o ListSitePublishingCredentialsResultOutput) ToListSitePublishingCredentialsResultOutput() ListSitePublishingCredentialsResultOutput {
	return o
}

func (o ListSitePublishingCredentialsResultOutput) ToListSitePublishingCredentialsResultOutputWithContext(ctx context.Context) ListSitePublishingCredentialsResultOutput {
	return o
}

func (o ListSitePublishingCredentialsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListSitePublishingCredentialsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListSitePublishingCredentialsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListSitePublishingCredentialsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListSitePublishingCredentialsResultOutput) PublishingPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsResult) *string { return v.PublishingPassword }).(pulumi.StringPtrOutput)
}

func (o ListSitePublishingCredentialsResultOutput) PublishingUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsResult) *string { return v.PublishingUserName }).(pulumi.StringPtrOutput)
}

func (o ListSitePublishingCredentialsResultOutput) ScmUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsResult) *string { return v.ScmUri }).(pulumi.StringPtrOutput)
}

func (o ListSitePublishingCredentialsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListSitePublishingCredentialsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSitePublishingCredentialsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSitePublishingCredentialsResultOutput{})
}
