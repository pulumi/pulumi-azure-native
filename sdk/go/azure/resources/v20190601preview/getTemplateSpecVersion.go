


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTemplateSpecVersion(ctx *pulumi.Context, args *LookupTemplateSpecVersionArgs, opts ...pulumi.InvokeOption) (*LookupTemplateSpecVersionResult, error) {
	var rv LookupTemplateSpecVersionResult
	err := ctx.Invoke("azure-native:resources/v20190601preview:getTemplateSpecVersion", args, &rv, opts...)
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
	Artifacts   []TemplateSpecTemplateArtifactResponse `pulumi:"artifacts"`
	Description *string                                `pulumi:"description"`
	Id          string                                 `pulumi:"id"`
	Location    string                                 `pulumi:"location"`
	Name        string                                 `pulumi:"name"`
	SystemData  SystemDataResponse                     `pulumi:"systemData"`
	Tags        map[string]string                      `pulumi:"tags"`
	Template    interface{}                            `pulumi:"template"`
	Type        string                                 `pulumi:"type"`
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

func (o LookupTemplateSpecVersionResultOutput) Artifacts() TemplateSpecTemplateArtifactResponseArrayOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) []TemplateSpecTemplateArtifactResponse { return v.Artifacts }).(TemplateSpecTemplateArtifactResponseArrayOutput)
}

func (o LookupTemplateSpecVersionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupTemplateSpecVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTemplateSpecVersionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) string { return v.Location }).(pulumi.StringOutput)
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

func (o LookupTemplateSpecVersionResultOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) interface{} { return v.Template }).(pulumi.AnyOutput)
}

func (o LookupTemplateSpecVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemplateSpecVersionResultOutput{})
}
