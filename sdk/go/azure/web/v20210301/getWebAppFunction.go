


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppFunction(ctx *pulumi.Context, args *LookupWebAppFunctionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppFunctionResult, error) {
	var rv LookupWebAppFunctionResult
	err := ctx.Invoke("azure-native:web/v20210301:getWebAppFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppFunctionArgs struct {
	FunctionName      string `pulumi:"functionName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppFunctionResult struct {
	Config             interface{}       `pulumi:"config"`
	ConfigHref         *string           `pulumi:"configHref"`
	Files              map[string]string `pulumi:"files"`
	FunctionAppId      *string           `pulumi:"functionAppId"`
	Href               *string           `pulumi:"href"`
	Id                 string            `pulumi:"id"`
	InvokeUrlTemplate  *string           `pulumi:"invokeUrlTemplate"`
	IsDisabled         *bool             `pulumi:"isDisabled"`
	Kind               *string           `pulumi:"kind"`
	Language           *string           `pulumi:"language"`
	Name               string            `pulumi:"name"`
	ScriptHref         *string           `pulumi:"scriptHref"`
	ScriptRootPathHref *string           `pulumi:"scriptRootPathHref"`
	SecretsFileHref    *string           `pulumi:"secretsFileHref"`
	TestData           *string           `pulumi:"testData"`
	TestDataHref       *string           `pulumi:"testDataHref"`
	Type               string            `pulumi:"type"`
}

func LookupWebAppFunctionOutput(ctx *pulumi.Context, args LookupWebAppFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppFunctionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppFunctionResult, error) {
			args := v.(LookupWebAppFunctionArgs)
			r, err := LookupWebAppFunction(ctx, &args, opts...)
			var s LookupWebAppFunctionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppFunctionResultOutput)
}

type LookupWebAppFunctionOutputArgs struct {
	FunctionName      pulumi.StringInput `pulumi:"functionName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppFunctionArgs)(nil)).Elem()
}


type LookupWebAppFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppFunctionResult)(nil)).Elem()
}

func (o LookupWebAppFunctionResultOutput) ToLookupWebAppFunctionResultOutput() LookupWebAppFunctionResultOutput {
	return o
}

func (o LookupWebAppFunctionResultOutput) ToLookupWebAppFunctionResultOutputWithContext(ctx context.Context) LookupWebAppFunctionResultOutput {
	return o
}

func (o LookupWebAppFunctionResultOutput) Config() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) interface{} { return v.Config }).(pulumi.AnyOutput)
}

func (o LookupWebAppFunctionResultOutput) ConfigHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.ConfigHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) Files() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) map[string]string { return v.Files }).(pulumi.StringMapOutput)
}

func (o LookupWebAppFunctionResultOutput) FunctionAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.FunctionAppId }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) Href() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.Href }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppFunctionResultOutput) InvokeUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.InvokeUrlTemplate }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.Language }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppFunctionResultOutput) ScriptHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.ScriptHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) ScriptRootPathHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.ScriptRootPathHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) SecretsFileHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.SecretsFileHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) TestData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.TestData }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) TestDataHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.TestDataHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppFunctionResultOutput{})
}
