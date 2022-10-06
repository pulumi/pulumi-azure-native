


package blueprint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlueprint(ctx *pulumi.Context, args *LookupBlueprintArgs, opts ...pulumi.InvokeOption) (*LookupBlueprintResult, error) {
	var rv LookupBlueprintResult
	err := ctx.Invoke("azure-native:blueprint:getBlueprint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlueprintArgs struct {
	BlueprintName string `pulumi:"blueprintName"`
	ResourceScope string `pulumi:"resourceScope"`
}


type LookupBlueprintResult struct {
	Description    *string                                    `pulumi:"description"`
	DisplayName    *string                                    `pulumi:"displayName"`
	Id             string                                     `pulumi:"id"`
	Layout         interface{}                                `pulumi:"layout"`
	Name           string                                     `pulumi:"name"`
	Parameters     map[string]ParameterDefinitionResponse     `pulumi:"parameters"`
	ResourceGroups map[string]ResourceGroupDefinitionResponse `pulumi:"resourceGroups"`
	Status         BlueprintStatusResponse                    `pulumi:"status"`
	TargetScope    string                                     `pulumi:"targetScope"`
	Type           string                                     `pulumi:"type"`
	Versions       interface{}                                `pulumi:"versions"`
}

func LookupBlueprintOutput(ctx *pulumi.Context, args LookupBlueprintOutputArgs, opts ...pulumi.InvokeOption) LookupBlueprintResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlueprintResult, error) {
			args := v.(LookupBlueprintArgs)
			r, err := LookupBlueprint(ctx, &args, opts...)
			var s LookupBlueprintResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlueprintResultOutput)
}

type LookupBlueprintOutputArgs struct {
	BlueprintName pulumi.StringInput `pulumi:"blueprintName"`
	ResourceScope pulumi.StringInput `pulumi:"resourceScope"`
}

func (LookupBlueprintOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlueprintArgs)(nil)).Elem()
}


type LookupBlueprintResultOutput struct{ *pulumi.OutputState }

func (LookupBlueprintResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlueprintResult)(nil)).Elem()
}

func (o LookupBlueprintResultOutput) ToLookupBlueprintResultOutput() LookupBlueprintResultOutput {
	return o
}

func (o LookupBlueprintResultOutput) ToLookupBlueprintResultOutputWithContext(ctx context.Context) LookupBlueprintResultOutput {
	return o
}

func (o LookupBlueprintResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlueprintResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupBlueprintResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlueprintResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupBlueprintResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlueprintResultOutput) Layout() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupBlueprintResult) interface{} { return v.Layout }).(pulumi.AnyOutput)
}

func (o LookupBlueprintResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlueprintResultOutput) Parameters() ParameterDefinitionResponseMapOutput {
	return o.ApplyT(func(v LookupBlueprintResult) map[string]ParameterDefinitionResponse { return v.Parameters }).(ParameterDefinitionResponseMapOutput)
}

func (o LookupBlueprintResultOutput) ResourceGroups() ResourceGroupDefinitionResponseMapOutput {
	return o.ApplyT(func(v LookupBlueprintResult) map[string]ResourceGroupDefinitionResponse { return v.ResourceGroups }).(ResourceGroupDefinitionResponseMapOutput)
}

func (o LookupBlueprintResultOutput) Status() BlueprintStatusResponseOutput {
	return o.ApplyT(func(v LookupBlueprintResult) BlueprintStatusResponse { return v.Status }).(BlueprintStatusResponseOutput)
}

func (o LookupBlueprintResultOutput) TargetScope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.TargetScope }).(pulumi.StringOutput)
}

func (o LookupBlueprintResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBlueprintResultOutput) Versions() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupBlueprintResult) interface{} { return v.Versions }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlueprintResultOutput{})
}
