


package v20191201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEmailTemplate(ctx *pulumi.Context, args *LookupEmailTemplateArgs, opts ...pulumi.InvokeOption) (*LookupEmailTemplateResult, error) {
	var rv LookupEmailTemplateResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201:getEmailTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEmailTemplateArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	TemplateName      string `pulumi:"templateName"`
}


type LookupEmailTemplateResult struct {
	Body        string                                              `pulumi:"body"`
	Description *string                                             `pulumi:"description"`
	Id          string                                              `pulumi:"id"`
	IsDefault   bool                                                `pulumi:"isDefault"`
	Name        string                                              `pulumi:"name"`
	Parameters  []EmailTemplateParametersContractPropertiesResponse `pulumi:"parameters"`
	Subject     string                                              `pulumi:"subject"`
	Title       *string                                             `pulumi:"title"`
	Type        string                                              `pulumi:"type"`
}

func LookupEmailTemplateOutput(ctx *pulumi.Context, args LookupEmailTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupEmailTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEmailTemplateResult, error) {
			args := v.(LookupEmailTemplateArgs)
			r, err := LookupEmailTemplate(ctx, &args, opts...)
			var s LookupEmailTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEmailTemplateResultOutput)
}

type LookupEmailTemplateOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	TemplateName      pulumi.StringInput `pulumi:"templateName"`
}

func (LookupEmailTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEmailTemplateArgs)(nil)).Elem()
}


type LookupEmailTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupEmailTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEmailTemplateResult)(nil)).Elem()
}

func (o LookupEmailTemplateResultOutput) ToLookupEmailTemplateResultOutput() LookupEmailTemplateResultOutput {
	return o
}

func (o LookupEmailTemplateResultOutput) ToLookupEmailTemplateResultOutputWithContext(ctx context.Context) LookupEmailTemplateResultOutput {
	return o
}

func (o LookupEmailTemplateResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) string { return v.Body }).(pulumi.StringOutput)
}

func (o LookupEmailTemplateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupEmailTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEmailTemplateResultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

func (o LookupEmailTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEmailTemplateResultOutput) Parameters() EmailTemplateParametersContractPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) []EmailTemplateParametersContractPropertiesResponse {
		return v.Parameters
	}).(EmailTemplateParametersContractPropertiesResponseArrayOutput)
}

func (o LookupEmailTemplateResultOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) string { return v.Subject }).(pulumi.StringOutput)
}

func (o LookupEmailTemplateResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o LookupEmailTemplateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEmailTemplateResultOutput{})
}
