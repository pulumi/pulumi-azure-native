


package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleAssignmentArtifact(ctx *pulumi.Context, args *LookupRoleAssignmentArtifactArgs, opts ...pulumi.InvokeOption) (*LookupRoleAssignmentArtifactResult, error) {
	var rv LookupRoleAssignmentArtifactResult
	err := ctx.Invoke("azure-native:blueprint/v20181101preview:getRoleAssignmentArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleAssignmentArtifactArgs struct {
	ArtifactName  string `pulumi:"artifactName"`
	BlueprintName string `pulumi:"blueprintName"`
	ResourceScope string `pulumi:"resourceScope"`
}


type LookupRoleAssignmentArtifactResult struct {
	DependsOn        []string    `pulumi:"dependsOn"`
	Description      *string     `pulumi:"description"`
	DisplayName      *string     `pulumi:"displayName"`
	Id               string      `pulumi:"id"`
	Kind             string      `pulumi:"kind"`
	Name             string      `pulumi:"name"`
	PrincipalIds     interface{} `pulumi:"principalIds"`
	ResourceGroup    *string     `pulumi:"resourceGroup"`
	RoleDefinitionId string      `pulumi:"roleDefinitionId"`
	Type             string      `pulumi:"type"`
}

func LookupRoleAssignmentArtifactOutput(ctx *pulumi.Context, args LookupRoleAssignmentArtifactOutputArgs, opts ...pulumi.InvokeOption) LookupRoleAssignmentArtifactResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoleAssignmentArtifactResult, error) {
			args := v.(LookupRoleAssignmentArtifactArgs)
			r, err := LookupRoleAssignmentArtifact(ctx, &args, opts...)
			var s LookupRoleAssignmentArtifactResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoleAssignmentArtifactResultOutput)
}

type LookupRoleAssignmentArtifactOutputArgs struct {
	ArtifactName  pulumi.StringInput `pulumi:"artifactName"`
	BlueprintName pulumi.StringInput `pulumi:"blueprintName"`
	ResourceScope pulumi.StringInput `pulumi:"resourceScope"`
}

func (LookupRoleAssignmentArtifactOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleAssignmentArtifactArgs)(nil)).Elem()
}


type LookupRoleAssignmentArtifactResultOutput struct{ *pulumi.OutputState }

func (LookupRoleAssignmentArtifactResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleAssignmentArtifactResult)(nil)).Elem()
}

func (o LookupRoleAssignmentArtifactResultOutput) ToLookupRoleAssignmentArtifactResultOutput() LookupRoleAssignmentArtifactResultOutput {
	return o
}

func (o LookupRoleAssignmentArtifactResultOutput) ToLookupRoleAssignmentArtifactResultOutputWithContext(ctx context.Context) LookupRoleAssignmentArtifactResultOutput {
	return o
}

func (o LookupRoleAssignmentArtifactResultOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRoleAssignmentArtifactResult) []string { return v.DependsOn }).(pulumi.StringArrayOutput)
}

func (o LookupRoleAssignmentArtifactResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentArtifactResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentArtifactResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentArtifactResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentArtifactResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentArtifactResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentArtifactResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentArtifactResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentArtifactResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentArtifactResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentArtifactResultOutput) PrincipalIds() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupRoleAssignmentArtifactResult) interface{} { return v.PrincipalIds }).(pulumi.AnyOutput)
}

func (o LookupRoleAssignmentArtifactResultOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentArtifactResult) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentArtifactResultOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentArtifactResult) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentArtifactResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentArtifactResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleAssignmentArtifactResultOutput{})
}
