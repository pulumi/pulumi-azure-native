


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTemplateSpec(ctx *pulumi.Context, args *LookupTemplateSpecArgs, opts ...pulumi.InvokeOption) (*LookupTemplateSpecResult, error) {
	var rv LookupTemplateSpecResult
	err := ctx.Invoke("azure-native:resources/v20210301preview:getTemplateSpec", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateSpecArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	TemplateSpecName  string  `pulumi:"templateSpecName"`
}


type LookupTemplateSpecResult struct {
	Description *string                                    `pulumi:"description"`
	DisplayName *string                                    `pulumi:"displayName"`
	Id          string                                     `pulumi:"id"`
	Location    string                                     `pulumi:"location"`
	Metadata    interface{}                                `pulumi:"metadata"`
	Name        string                                     `pulumi:"name"`
	SystemData  SystemDataResponse                         `pulumi:"systemData"`
	Tags        map[string]string                          `pulumi:"tags"`
	Type        string                                     `pulumi:"type"`
	Versions    map[string]TemplateSpecVersionInfoResponse `pulumi:"versions"`
}

func LookupTemplateSpecOutput(ctx *pulumi.Context, args LookupTemplateSpecOutputArgs, opts ...pulumi.InvokeOption) LookupTemplateSpecResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTemplateSpecResult, error) {
			args := v.(LookupTemplateSpecArgs)
			r, err := LookupTemplateSpec(ctx, &args, opts...)
			var s LookupTemplateSpecResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTemplateSpecResultOutput)
}

type LookupTemplateSpecOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	TemplateSpecName  pulumi.StringInput    `pulumi:"templateSpecName"`
}

func (LookupTemplateSpecOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateSpecArgs)(nil)).Elem()
}


type LookupTemplateSpecResultOutput struct{ *pulumi.OutputState }

func (LookupTemplateSpecResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateSpecResult)(nil)).Elem()
}

func (o LookupTemplateSpecResultOutput) ToLookupTemplateSpecResultOutput() LookupTemplateSpecResultOutput {
	return o
}

func (o LookupTemplateSpecResultOutput) ToLookupTemplateSpecResultOutputWithContext(ctx context.Context) LookupTemplateSpecResultOutput {
	return o
}

func (o LookupTemplateSpecResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateSpecResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupTemplateSpecResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateSpecResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupTemplateSpecResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTemplateSpecResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupTemplateSpecResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTemplateSpecResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupTemplateSpecResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTemplateSpecResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTemplateSpecResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTemplateSpecResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTemplateSpecResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupTemplateSpecResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupTemplateSpecResultOutput) Versions() TemplateSpecVersionInfoResponseMapOutput {
	return o.ApplyT(func(v LookupTemplateSpecResult) map[string]TemplateSpecVersionInfoResponse { return v.Versions }).(TemplateSpecVersionInfoResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemplateSpecResultOutput{})
}
