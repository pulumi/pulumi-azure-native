


package v20180101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTagDescription(ctx *pulumi.Context, args *LookupTagDescriptionArgs, opts ...pulumi.InvokeOption) (*LookupTagDescriptionResult, error) {
	var rv LookupTagDescriptionResult
	err := ctx.Invoke("azure-native:apimanagement/v20180101:getTagDescription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagDescriptionArgs struct {
	ApiId             string `pulumi:"apiId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	TagId             string `pulumi:"tagId"`
}


type LookupTagDescriptionResult struct {
	Description             *string `pulumi:"description"`
	DisplayName             *string `pulumi:"displayName"`
	ExternalDocsDescription *string `pulumi:"externalDocsDescription"`
	ExternalDocsUrl         *string `pulumi:"externalDocsUrl"`
	Id                      string  `pulumi:"id"`
	Name                    string  `pulumi:"name"`
	Type                    string  `pulumi:"type"`
}

func LookupTagDescriptionOutput(ctx *pulumi.Context, args LookupTagDescriptionOutputArgs, opts ...pulumi.InvokeOption) LookupTagDescriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagDescriptionResult, error) {
			args := v.(LookupTagDescriptionArgs)
			r, err := LookupTagDescription(ctx, &args, opts...)
			var s LookupTagDescriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagDescriptionResultOutput)
}

type LookupTagDescriptionOutputArgs struct {
	ApiId             pulumi.StringInput `pulumi:"apiId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	TagId             pulumi.StringInput `pulumi:"tagId"`
}

func (LookupTagDescriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagDescriptionArgs)(nil)).Elem()
}


type LookupTagDescriptionResultOutput struct{ *pulumi.OutputState }

func (LookupTagDescriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagDescriptionResult)(nil)).Elem()
}

func (o LookupTagDescriptionResultOutput) ToLookupTagDescriptionResultOutput() LookupTagDescriptionResultOutput {
	return o
}

func (o LookupTagDescriptionResultOutput) ToLookupTagDescriptionResultOutputWithContext(ctx context.Context) LookupTagDescriptionResultOutput {
	return o
}

func (o LookupTagDescriptionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTagDescriptionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupTagDescriptionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTagDescriptionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupTagDescriptionResultOutput) ExternalDocsDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTagDescriptionResult) *string { return v.ExternalDocsDescription }).(pulumi.StringPtrOutput)
}

func (o LookupTagDescriptionResultOutput) ExternalDocsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTagDescriptionResult) *string { return v.ExternalDocsUrl }).(pulumi.StringPtrOutput)
}

func (o LookupTagDescriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagDescriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagDescriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagDescriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTagDescriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagDescriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagDescriptionResultOutput{})
}
