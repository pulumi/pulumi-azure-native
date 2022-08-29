


package v20170426

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleAssignment(ctx *pulumi.Context, args *LookupRoleAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleAssignmentResult, error) {
	var rv LookupRoleAssignmentResult
	err := ctx.Invoke("azure-native:customerinsights/v20170426:getRoleAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleAssignmentArgs struct {
	AssignmentName    string `pulumi:"assignmentName"`
	HubName           string `pulumi:"hubName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRoleAssignmentResult struct {
	AssignmentName     string                          `pulumi:"assignmentName"`
	ConflationPolicies *ResourceSetDescriptionResponse `pulumi:"conflationPolicies"`
	Connectors         *ResourceSetDescriptionResponse `pulumi:"connectors"`
	Description        map[string]string               `pulumi:"description"`
	DisplayName        map[string]string               `pulumi:"displayName"`
	Id                 string                          `pulumi:"id"`
	Interactions       *ResourceSetDescriptionResponse `pulumi:"interactions"`
	Kpis               *ResourceSetDescriptionResponse `pulumi:"kpis"`
	Links              *ResourceSetDescriptionResponse `pulumi:"links"`
	Name               string                          `pulumi:"name"`
	Principals         []AssignmentPrincipalResponse   `pulumi:"principals"`
	Profiles           *ResourceSetDescriptionResponse `pulumi:"profiles"`
	ProvisioningState  string                          `pulumi:"provisioningState"`
	RelationshipLinks  *ResourceSetDescriptionResponse `pulumi:"relationshipLinks"`
	Relationships      *ResourceSetDescriptionResponse `pulumi:"relationships"`
	Role               string                          `pulumi:"role"`
	RoleAssignments    *ResourceSetDescriptionResponse `pulumi:"roleAssignments"`
	SasPolicies        *ResourceSetDescriptionResponse `pulumi:"sasPolicies"`
	Segments           *ResourceSetDescriptionResponse `pulumi:"segments"`
	TenantId           string                          `pulumi:"tenantId"`
	Type               string                          `pulumi:"type"`
	Views              *ResourceSetDescriptionResponse `pulumi:"views"`
	WidgetTypes        *ResourceSetDescriptionResponse `pulumi:"widgetTypes"`
}

func LookupRoleAssignmentOutput(ctx *pulumi.Context, args LookupRoleAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupRoleAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoleAssignmentResult, error) {
			args := v.(LookupRoleAssignmentArgs)
			r, err := LookupRoleAssignment(ctx, &args, opts...)
			var s LookupRoleAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoleAssignmentResultOutput)
}

type LookupRoleAssignmentOutputArgs struct {
	AssignmentName    pulumi.StringInput `pulumi:"assignmentName"`
	HubName           pulumi.StringInput `pulumi:"hubName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRoleAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleAssignmentArgs)(nil)).Elem()
}


type LookupRoleAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupRoleAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleAssignmentResult)(nil)).Elem()
}

func (o LookupRoleAssignmentResultOutput) ToLookupRoleAssignmentResultOutput() LookupRoleAssignmentResultOutput {
	return o
}

func (o LookupRoleAssignmentResultOutput) ToLookupRoleAssignmentResultOutputWithContext(ctx context.Context) LookupRoleAssignmentResultOutput {
	return o
}

func (o LookupRoleAssignmentResultOutput) AssignmentName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.AssignmentName }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) ConflationPolicies() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *ResourceSetDescriptionResponse { return v.ConflationPolicies }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Connectors() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *ResourceSetDescriptionResponse { return v.Connectors }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) map[string]string { return v.Description }).(pulumi.StringMapOutput)
}

func (o LookupRoleAssignmentResultOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o LookupRoleAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) Interactions() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *ResourceSetDescriptionResponse { return v.Interactions }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Kpis() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *ResourceSetDescriptionResponse { return v.Kpis }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Links() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *ResourceSetDescriptionResponse { return v.Links }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) Principals() AssignmentPrincipalResponseArrayOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) []AssignmentPrincipalResponse { return v.Principals }).(AssignmentPrincipalResponseArrayOutput)
}

func (o LookupRoleAssignmentResultOutput) Profiles() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *ResourceSetDescriptionResponse { return v.Profiles }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o LookupRoleAssignmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) RelationshipLinks() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *ResourceSetDescriptionResponse { return v.RelationshipLinks }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Relationships() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *ResourceSetDescriptionResponse { return v.Relationships }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Role }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) RoleAssignments() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *ResourceSetDescriptionResponse { return v.RoleAssignments }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o LookupRoleAssignmentResultOutput) SasPolicies() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *ResourceSetDescriptionResponse { return v.SasPolicies }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Segments() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *ResourceSetDescriptionResponse { return v.Segments }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o LookupRoleAssignmentResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) Views() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *ResourceSetDescriptionResponse { return v.Views }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o LookupRoleAssignmentResultOutput) WidgetTypes() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *ResourceSetDescriptionResponse { return v.WidgetTypes }).(ResourceSetDescriptionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleAssignmentResultOutput{})
}
