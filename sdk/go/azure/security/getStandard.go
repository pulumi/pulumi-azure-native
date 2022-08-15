


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStandard(ctx *pulumi.Context, args *LookupStandardArgs, opts ...pulumi.InvokeOption) (*LookupStandardResult, error) {
	var rv LookupStandardResult
	err := ctx.Invoke("azure-native:security:getStandard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStandardArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StandardId        string `pulumi:"standardId"`
}


type LookupStandardResult struct {
	Category        *string                               `pulumi:"category"`
	Components      []StandardComponentPropertiesResponse `pulumi:"components"`
	Description     *string                               `pulumi:"description"`
	DisplayName     *string                               `pulumi:"displayName"`
	Etag            *string                               `pulumi:"etag"`
	Id              string                                `pulumi:"id"`
	Kind            *string                               `pulumi:"kind"`
	Location        *string                               `pulumi:"location"`
	Name            string                                `pulumi:"name"`
	StandardType    string                                `pulumi:"standardType"`
	SupportedClouds []string                              `pulumi:"supportedClouds"`
	SystemData      SystemDataResponse                    `pulumi:"systemData"`
	Tags            map[string]string                     `pulumi:"tags"`
	Type            string                                `pulumi:"type"`
}

func LookupStandardOutput(ctx *pulumi.Context, args LookupStandardOutputArgs, opts ...pulumi.InvokeOption) LookupStandardResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStandardResult, error) {
			args := v.(LookupStandardArgs)
			r, err := LookupStandard(ctx, &args, opts...)
			var s LookupStandardResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStandardResultOutput)
}

type LookupStandardOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	StandardId        pulumi.StringInput `pulumi:"standardId"`
}

func (LookupStandardOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStandardArgs)(nil)).Elem()
}


type LookupStandardResultOutput struct{ *pulumi.OutputState }

func (LookupStandardResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStandardResult)(nil)).Elem()
}

func (o LookupStandardResultOutput) ToLookupStandardResultOutput() LookupStandardResultOutput {
	return o
}

func (o LookupStandardResultOutput) ToLookupStandardResultOutputWithContext(ctx context.Context) LookupStandardResultOutput {
	return o
}

func (o LookupStandardResultOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardResult) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o LookupStandardResultOutput) Components() StandardComponentPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupStandardResult) []StandardComponentPropertiesResponse { return v.Components }).(StandardComponentPropertiesResponseArrayOutput)
}

func (o LookupStandardResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupStandardResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupStandardResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupStandardResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupStandardResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupStandardResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) StandardType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.StandardType }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) SupportedClouds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStandardResult) []string { return v.SupportedClouds }).(pulumi.StringArrayOutput)
}

func (o LookupStandardResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStandardResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupStandardResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStandardResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStandardResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStandardResultOutput{})
}
