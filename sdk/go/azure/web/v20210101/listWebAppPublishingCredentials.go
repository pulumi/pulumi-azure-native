


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppPublishingCredentials(ctx *pulumi.Context, args *ListWebAppPublishingCredentialsArgs, opts ...pulumi.InvokeOption) (*ListWebAppPublishingCredentialsResult, error) {
	var rv ListWebAppPublishingCredentialsResult
	err := ctx.Invoke("azure-native:web/v20210101:listWebAppPublishingCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppPublishingCredentialsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppPublishingCredentialsResult struct {
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

func ListWebAppPublishingCredentialsOutput(ctx *pulumi.Context, args ListWebAppPublishingCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppPublishingCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppPublishingCredentialsResult, error) {
			args := v.(ListWebAppPublishingCredentialsArgs)
			r, err := ListWebAppPublishingCredentials(ctx, &args, opts...)
			var s ListWebAppPublishingCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppPublishingCredentialsResultOutput)
}

type ListWebAppPublishingCredentialsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppPublishingCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppPublishingCredentialsArgs)(nil)).Elem()
}


type ListWebAppPublishingCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppPublishingCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppPublishingCredentialsResult)(nil)).Elem()
}

func (o ListWebAppPublishingCredentialsResultOutput) ToListWebAppPublishingCredentialsResultOutput() ListWebAppPublishingCredentialsResultOutput {
	return o
}

func (o ListWebAppPublishingCredentialsResultOutput) ToListWebAppPublishingCredentialsResultOutputWithContext(ctx context.Context) ListWebAppPublishingCredentialsResultOutput {
	return o
}

func (o ListWebAppPublishingCredentialsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppPublishingCredentialsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppPublishingCredentialsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppPublishingCredentialsResultOutput) PublishingPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) *string { return v.PublishingPassword }).(pulumi.StringPtrOutput)
}

func (o ListWebAppPublishingCredentialsResultOutput) PublishingPasswordHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) *string { return v.PublishingPasswordHash }).(pulumi.StringPtrOutput)
}

func (o ListWebAppPublishingCredentialsResultOutput) PublishingPasswordHashSalt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) *string { return v.PublishingPasswordHashSalt }).(pulumi.StringPtrOutput)
}

func (o ListWebAppPublishingCredentialsResultOutput) PublishingUserName() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) string { return v.PublishingUserName }).(pulumi.StringOutput)
}

func (o ListWebAppPublishingCredentialsResultOutput) ScmUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) *string { return v.ScmUri }).(pulumi.StringPtrOutput)
}

func (o ListWebAppPublishingCredentialsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppPublishingCredentialsResultOutput{})
}
