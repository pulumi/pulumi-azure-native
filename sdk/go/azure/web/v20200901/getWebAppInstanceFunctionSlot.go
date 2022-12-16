


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppInstanceFunctionSlot(ctx *pulumi.Context, args *LookupWebAppInstanceFunctionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppInstanceFunctionSlotResult, error) {
	var rv LookupWebAppInstanceFunctionSlotResult
	err := ctx.Invoke("azure-native:web/v20200901:getWebAppInstanceFunctionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppInstanceFunctionSlotArgs struct {
	FunctionName      string `pulumi:"functionName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppInstanceFunctionSlotResult struct {
	Config             interface{}        `pulumi:"config"`
	ConfigHref         *string            `pulumi:"configHref"`
	Files              map[string]string  `pulumi:"files"`
	FunctionAppId      *string            `pulumi:"functionAppId"`
	Href               *string            `pulumi:"href"`
	Id                 string             `pulumi:"id"`
	InvokeUrlTemplate  *string            `pulumi:"invokeUrlTemplate"`
	IsDisabled         *bool              `pulumi:"isDisabled"`
	Kind               *string            `pulumi:"kind"`
	Language           *string            `pulumi:"language"`
	Name               string             `pulumi:"name"`
	ScriptHref         *string            `pulumi:"scriptHref"`
	ScriptRootPathHref *string            `pulumi:"scriptRootPathHref"`
	SecretsFileHref    *string            `pulumi:"secretsFileHref"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	TestData           *string            `pulumi:"testData"`
	TestDataHref       *string            `pulumi:"testDataHref"`
	Type               string             `pulumi:"type"`
}

func LookupWebAppInstanceFunctionSlotOutput(ctx *pulumi.Context, args LookupWebAppInstanceFunctionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppInstanceFunctionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppInstanceFunctionSlotResult, error) {
			args := v.(LookupWebAppInstanceFunctionSlotArgs)
			r, err := LookupWebAppInstanceFunctionSlot(ctx, &args, opts...)
			var s LookupWebAppInstanceFunctionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppInstanceFunctionSlotResultOutput)
}

type LookupWebAppInstanceFunctionSlotOutputArgs struct {
	FunctionName      pulumi.StringInput `pulumi:"functionName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppInstanceFunctionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppInstanceFunctionSlotArgs)(nil)).Elem()
}


type LookupWebAppInstanceFunctionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppInstanceFunctionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppInstanceFunctionSlotResult)(nil)).Elem()
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) ToLookupWebAppInstanceFunctionSlotResultOutput() LookupWebAppInstanceFunctionSlotResultOutput {
	return o
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) ToLookupWebAppInstanceFunctionSlotResultOutputWithContext(ctx context.Context) LookupWebAppInstanceFunctionSlotResultOutput {
	return o
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) Config() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) interface{} { return v.Config }).(pulumi.AnyOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) ConfigHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.ConfigHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) Files() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) map[string]string { return v.Files }).(pulumi.StringMapOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) FunctionAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.FunctionAppId }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) Href() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.Href }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) InvokeUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.InvokeUrlTemplate }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.Language }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) ScriptHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.ScriptHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) ScriptRootPathHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.ScriptRootPathHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) SecretsFileHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.SecretsFileHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) TestData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.TestData }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) TestDataHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.TestDataHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppInstanceFunctionSlotResultOutput{})
}
