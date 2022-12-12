


package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupPublishedBlueprint(ctx *pulumi.Context, args *LookupPublishedBlueprintArgs, opts ...pulumi.InvokeOption) (*LookupPublishedBlueprintResult, error) {
	var rv LookupPublishedBlueprintResult
	err := ctx.Invoke("azure-native:blueprint/v20171111preview:getPublishedBlueprint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublishedBlueprintArgs struct {
	BlueprintName       string `pulumi:"blueprintName"`
	ManagementGroupName string `pulumi:"managementGroupName"`
	VersionId           string `pulumi:"versionId"`
}


type LookupPublishedBlueprintResult struct {
	BlueprintName  *string                                    `pulumi:"blueprintName"`
	ChangeNotes    *string                                    `pulumi:"changeNotes"`
	Description    *string                                    `pulumi:"description"`
	DisplayName    *string                                    `pulumi:"displayName"`
	Id             string                                     `pulumi:"id"`
	Name           string                                     `pulumi:"name"`
	Parameters     map[string]ParameterDefinitionResponse     `pulumi:"parameters"`
	ResourceGroups map[string]ResourceGroupDefinitionResponse `pulumi:"resourceGroups"`
	Status         BlueprintStatusResponse                    `pulumi:"status"`
	TargetScope    *string                                    `pulumi:"targetScope"`
	Type           string                                     `pulumi:"type"`
}

func LookupPublishedBlueprintOutput(ctx *pulumi.Context, args LookupPublishedBlueprintOutputArgs, opts ...pulumi.InvokeOption) LookupPublishedBlueprintResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPublishedBlueprintResult, error) {
			args := v.(LookupPublishedBlueprintArgs)
			r, err := LookupPublishedBlueprint(ctx, &args, opts...)
			var s LookupPublishedBlueprintResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPublishedBlueprintResultOutput)
}

type LookupPublishedBlueprintOutputArgs struct {
	BlueprintName       pulumi.StringInput `pulumi:"blueprintName"`
	ManagementGroupName pulumi.StringInput `pulumi:"managementGroupName"`
	VersionId           pulumi.StringInput `pulumi:"versionId"`
}

func (LookupPublishedBlueprintOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublishedBlueprintArgs)(nil)).Elem()
}


type LookupPublishedBlueprintResultOutput struct{ *pulumi.OutputState }

func (LookupPublishedBlueprintResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublishedBlueprintResult)(nil)).Elem()
}

func (o LookupPublishedBlueprintResultOutput) ToLookupPublishedBlueprintResultOutput() LookupPublishedBlueprintResultOutput {
	return o
}

func (o LookupPublishedBlueprintResultOutput) ToLookupPublishedBlueprintResultOutputWithContext(ctx context.Context) LookupPublishedBlueprintResultOutput {
	return o
}

func (o LookupPublishedBlueprintResultOutput) BlueprintName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublishedBlueprintResult) *string { return v.BlueprintName }).(pulumi.StringPtrOutput)
}

func (o LookupPublishedBlueprintResultOutput) ChangeNotes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublishedBlueprintResult) *string { return v.ChangeNotes }).(pulumi.StringPtrOutput)
}

func (o LookupPublishedBlueprintResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublishedBlueprintResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPublishedBlueprintResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublishedBlueprintResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupPublishedBlueprintResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublishedBlueprintResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPublishedBlueprintResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublishedBlueprintResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPublishedBlueprintResultOutput) Parameters() ParameterDefinitionResponseMapOutput {
	return o.ApplyT(func(v LookupPublishedBlueprintResult) map[string]ParameterDefinitionResponse { return v.Parameters }).(ParameterDefinitionResponseMapOutput)
}

func (o LookupPublishedBlueprintResultOutput) ResourceGroups() ResourceGroupDefinitionResponseMapOutput {
	return o.ApplyT(func(v LookupPublishedBlueprintResult) map[string]ResourceGroupDefinitionResponse {
		return v.ResourceGroups
	}).(ResourceGroupDefinitionResponseMapOutput)
}

func (o LookupPublishedBlueprintResultOutput) Status() BlueprintStatusResponseOutput {
	return o.ApplyT(func(v LookupPublishedBlueprintResult) BlueprintStatusResponse { return v.Status }).(BlueprintStatusResponseOutput)
}

func (o LookupPublishedBlueprintResultOutput) TargetScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublishedBlueprintResult) *string { return v.TargetScope }).(pulumi.StringPtrOutput)
}

func (o LookupPublishedBlueprintResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublishedBlueprintResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublishedBlueprintResultOutput{})
}
