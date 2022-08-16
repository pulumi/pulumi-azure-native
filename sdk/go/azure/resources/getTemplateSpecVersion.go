


package resources

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTemplateSpecVersion(ctx *pulumi.Context, args *LookupTemplateSpecVersionArgs, opts ...pulumi.InvokeOption) (*LookupTemplateSpecVersionResult, error) {
	var rv LookupTemplateSpecVersionResult
	err := ctx.Invoke("azure-native:resources:getTemplateSpecVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateSpecVersionArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TemplateSpecName    string `pulumi:"templateSpecName"`
	TemplateSpecVersion string `pulumi:"templateSpecVersion"`
}


type LookupTemplateSpecVersionResult struct {
	Description      *string                          `pulumi:"description"`
	Id               string                           `pulumi:"id"`
	LinkedTemplates  []LinkedTemplateArtifactResponse `pulumi:"linkedTemplates"`
	Location         string                           `pulumi:"location"`
	MainTemplate     interface{}                      `pulumi:"mainTemplate"`
	Metadata         interface{}                      `pulumi:"metadata"`
	Name             string                           `pulumi:"name"`
	SystemData       SystemDataResponse               `pulumi:"systemData"`
	Tags             map[string]string                `pulumi:"tags"`
	Type             string                           `pulumi:"type"`
	UiFormDefinition interface{}                      `pulumi:"uiFormDefinition"`
}

func LookupTemplateSpecVersionOutput(ctx *pulumi.Context, args LookupTemplateSpecVersionOutputArgs, opts ...pulumi.InvokeOption) LookupTemplateSpecVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTemplateSpecVersionResult, error) {
			args := v.(LookupTemplateSpecVersionArgs)
			r, err := LookupTemplateSpecVersion(ctx, &args, opts...)
			var s LookupTemplateSpecVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTemplateSpecVersionResultOutput)
}

type LookupTemplateSpecVersionOutputArgs struct {
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	TemplateSpecName    pulumi.StringInput `pulumi:"templateSpecName"`
	TemplateSpecVersion pulumi.StringInput `pulumi:"templateSpecVersion"`
}

func (LookupTemplateSpecVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateSpecVersionArgs)(nil)).Elem()
}


type LookupTemplateSpecVersionResultOutput struct{ *pulumi.OutputState }

func (LookupTemplateSpecVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateSpecVersionResult)(nil)).Elem()
}

func (o LookupTemplateSpecVersionResultOutput) ToLookupTemplateSpecVersionResultOutput() LookupTemplateSpecVersionResultOutput {
	return o
}

func (o LookupTemplateSpecVersionResultOutput) ToLookupTemplateSpecVersionResultOutputWithContext(ctx context.Context) LookupTemplateSpecVersionResultOutput {
	return o
}

func (o LookupTemplateSpecVersionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupTemplateSpecVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTemplateSpecVersionResultOutput) LinkedTemplates() LinkedTemplateArtifactResponseArrayOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) []LinkedTemplateArtifactResponse { return v.LinkedTemplates }).(LinkedTemplateArtifactResponseArrayOutput)
}

func (o LookupTemplateSpecVersionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupTemplateSpecVersionResultOutput) MainTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) interface{} { return v.MainTemplate }).(pulumi.AnyOutput)
}

func (o LookupTemplateSpecVersionResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupTemplateSpecVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTemplateSpecVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTemplateSpecVersionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupTemplateSpecVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupTemplateSpecVersionResultOutput) UiFormDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) interface{} { return v.UiFormDefinition }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemplateSpecVersionResultOutput{})
}
