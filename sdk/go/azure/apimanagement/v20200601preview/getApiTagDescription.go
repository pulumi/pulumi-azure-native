


package v20200601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiTagDescription(ctx *pulumi.Context, args *LookupApiTagDescriptionArgs, opts ...pulumi.InvokeOption) (*LookupApiTagDescriptionResult, error) {
	var rv LookupApiTagDescriptionResult
	err := ctx.Invoke("azure-native:apimanagement/v20200601preview:getApiTagDescription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiTagDescriptionArgs struct {
	ApiId             string `pulumi:"apiId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	TagDescriptionId  string `pulumi:"tagDescriptionId"`
}


type LookupApiTagDescriptionResult struct {
	Description             *string `pulumi:"description"`
	DisplayName             *string `pulumi:"displayName"`
	ExternalDocsDescription *string `pulumi:"externalDocsDescription"`
	ExternalDocsUrl         *string `pulumi:"externalDocsUrl"`
	Id                      string  `pulumi:"id"`
	Name                    string  `pulumi:"name"`
	TagId                   *string `pulumi:"tagId"`
	Type                    string  `pulumi:"type"`
}

func LookupApiTagDescriptionOutput(ctx *pulumi.Context, args LookupApiTagDescriptionOutputArgs, opts ...pulumi.InvokeOption) LookupApiTagDescriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiTagDescriptionResult, error) {
			args := v.(LookupApiTagDescriptionArgs)
			r, err := LookupApiTagDescription(ctx, &args, opts...)
			var s LookupApiTagDescriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiTagDescriptionResultOutput)
}

type LookupApiTagDescriptionOutputArgs struct {
	ApiId             pulumi.StringInput `pulumi:"apiId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	TagDescriptionId  pulumi.StringInput `pulumi:"tagDescriptionId"`
}

func (LookupApiTagDescriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiTagDescriptionArgs)(nil)).Elem()
}


type LookupApiTagDescriptionResultOutput struct{ *pulumi.OutputState }

func (LookupApiTagDescriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiTagDescriptionResult)(nil)).Elem()
}

func (o LookupApiTagDescriptionResultOutput) ToLookupApiTagDescriptionResultOutput() LookupApiTagDescriptionResultOutput {
	return o
}

func (o LookupApiTagDescriptionResultOutput) ToLookupApiTagDescriptionResultOutputWithContext(ctx context.Context) LookupApiTagDescriptionResultOutput {
	return o
}

func (o LookupApiTagDescriptionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApiTagDescriptionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupApiTagDescriptionResultOutput) ExternalDocsDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) *string { return v.ExternalDocsDescription }).(pulumi.StringPtrOutput)
}

func (o LookupApiTagDescriptionResultOutput) ExternalDocsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) *string { return v.ExternalDocsUrl }).(pulumi.StringPtrOutput)
}

func (o LookupApiTagDescriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiTagDescriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiTagDescriptionResultOutput) TagId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) *string { return v.TagId }).(pulumi.StringPtrOutput)
}

func (o LookupApiTagDescriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiTagDescriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiTagDescriptionResultOutput{})
}
