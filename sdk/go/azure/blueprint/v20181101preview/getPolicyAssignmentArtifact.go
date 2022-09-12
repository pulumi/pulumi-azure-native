


package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyAssignmentArtifact(ctx *pulumi.Context, args *LookupPolicyAssignmentArtifactArgs, opts ...pulumi.InvokeOption) (*LookupPolicyAssignmentArtifactResult, error) {
	var rv LookupPolicyAssignmentArtifactResult
	err := ctx.Invoke("azure-native:blueprint/v20181101preview:getPolicyAssignmentArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyAssignmentArtifactArgs struct {
	ArtifactName  string `pulumi:"artifactName"`
	BlueprintName string `pulumi:"blueprintName"`
	ResourceScope string `pulumi:"resourceScope"`
}


type LookupPolicyAssignmentArtifactResult struct {
	DependsOn          []string                          `pulumi:"dependsOn"`
	Description        *string                           `pulumi:"description"`
	DisplayName        *string                           `pulumi:"displayName"`
	Id                 string                            `pulumi:"id"`
	Kind               string                            `pulumi:"kind"`
	Name               string                            `pulumi:"name"`
	Parameters         map[string]ParameterValueResponse `pulumi:"parameters"`
	PolicyDefinitionId string                            `pulumi:"policyDefinitionId"`
	ResourceGroup      *string                           `pulumi:"resourceGroup"`
	Type               string                            `pulumi:"type"`
}

func LookupPolicyAssignmentArtifactOutput(ctx *pulumi.Context, args LookupPolicyAssignmentArtifactOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyAssignmentArtifactResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyAssignmentArtifactResult, error) {
			args := v.(LookupPolicyAssignmentArtifactArgs)
			r, err := LookupPolicyAssignmentArtifact(ctx, &args, opts...)
			var s LookupPolicyAssignmentArtifactResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyAssignmentArtifactResultOutput)
}

type LookupPolicyAssignmentArtifactOutputArgs struct {
	ArtifactName  pulumi.StringInput `pulumi:"artifactName"`
	BlueprintName pulumi.StringInput `pulumi:"blueprintName"`
	ResourceScope pulumi.StringInput `pulumi:"resourceScope"`
}

func (LookupPolicyAssignmentArtifactOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyAssignmentArtifactArgs)(nil)).Elem()
}


type LookupPolicyAssignmentArtifactResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyAssignmentArtifactResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyAssignmentArtifactResult)(nil)).Elem()
}

func (o LookupPolicyAssignmentArtifactResultOutput) ToLookupPolicyAssignmentArtifactResultOutput() LookupPolicyAssignmentArtifactResultOutput {
	return o
}

func (o LookupPolicyAssignmentArtifactResultOutput) ToLookupPolicyAssignmentArtifactResultOutputWithContext(ctx context.Context) LookupPolicyAssignmentArtifactResultOutput {
	return o
}

func (o LookupPolicyAssignmentArtifactResultOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) []string { return v.DependsOn }).(pulumi.StringArrayOutput)
}

func (o LookupPolicyAssignmentArtifactResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyAssignmentArtifactResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyAssignmentArtifactResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicyAssignmentArtifactResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupPolicyAssignmentArtifactResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPolicyAssignmentArtifactResultOutput) Parameters() ParameterValueResponseMapOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) map[string]ParameterValueResponse { return v.Parameters }).(ParameterValueResponseMapOutput)
}

func (o LookupPolicyAssignmentArtifactResultOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) string { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

func (o LookupPolicyAssignmentArtifactResultOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyAssignmentArtifactResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyAssignmentArtifactResultOutput{})
}
