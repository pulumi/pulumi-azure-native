


package policyinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AttestationAtResourceGroup struct {
	pulumi.CustomResourceState

	Comments                    pulumi.StringPtrOutput                 `pulumi:"comments"`
	ComplianceState             pulumi.StringPtrOutput                 `pulumi:"complianceState"`
	Evidence                    AttestationEvidenceResponseArrayOutput `pulumi:"evidence"`
	ExpiresOn                   pulumi.StringPtrOutput                 `pulumi:"expiresOn"`
	LastComplianceStateChangeAt pulumi.StringOutput                    `pulumi:"lastComplianceStateChangeAt"`
	Name                        pulumi.StringOutput                    `pulumi:"name"`
	Owner                       pulumi.StringPtrOutput                 `pulumi:"owner"`
	PolicyAssignmentId          pulumi.StringOutput                    `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId pulumi.StringPtrOutput                 `pulumi:"policyDefinitionReferenceId"`
	ProvisioningState           pulumi.StringOutput                    `pulumi:"provisioningState"`
	SystemData                  SystemDataResponseOutput               `pulumi:"systemData"`
	Type                        pulumi.StringOutput                    `pulumi:"type"`
}


func NewAttestationAtResourceGroup(ctx *pulumi.Context,
	name string, args *AttestationAtResourceGroupArgs, opts ...pulumi.ResourceOption) (*AttestationAtResourceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyAssignmentId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyAssignmentId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:policyinsights/v20210101:AttestationAtResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20220901:AttestationAtResourceGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource AttestationAtResourceGroup
	err := ctx.RegisterResource("azure-native:policyinsights:AttestationAtResourceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAttestationAtResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestationAtResourceGroupState, opts ...pulumi.ResourceOption) (*AttestationAtResourceGroup, error) {
	var resource AttestationAtResourceGroup
	err := ctx.ReadResource("azure-native:policyinsights:AttestationAtResourceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type attestationAtResourceGroupState struct {
}

type AttestationAtResourceGroupState struct {
}

func (AttestationAtResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestationAtResourceGroupState)(nil)).Elem()
}

type attestationAtResourceGroupArgs struct {
	AttestationName             *string               `pulumi:"attestationName"`
	Comments                    *string               `pulumi:"comments"`
	ComplianceState             *string               `pulumi:"complianceState"`
	Evidence                    []AttestationEvidence `pulumi:"evidence"`
	ExpiresOn                   *string               `pulumi:"expiresOn"`
	Owner                       *string               `pulumi:"owner"`
	PolicyAssignmentId          string                `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string               `pulumi:"policyDefinitionReferenceId"`
	ResourceGroupName           string                `pulumi:"resourceGroupName"`
}


type AttestationAtResourceGroupArgs struct {
	AttestationName             pulumi.StringPtrInput
	Comments                    pulumi.StringPtrInput
	ComplianceState             pulumi.StringPtrInput
	Evidence                    AttestationEvidenceArrayInput
	ExpiresOn                   pulumi.StringPtrInput
	Owner                       pulumi.StringPtrInput
	PolicyAssignmentId          pulumi.StringInput
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
}

func (AttestationAtResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestationAtResourceGroupArgs)(nil)).Elem()
}

type AttestationAtResourceGroupInput interface {
	pulumi.Input

	ToAttestationAtResourceGroupOutput() AttestationAtResourceGroupOutput
	ToAttestationAtResourceGroupOutputWithContext(ctx context.Context) AttestationAtResourceGroupOutput
}

func (*AttestationAtResourceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestationAtResourceGroup)(nil)).Elem()
}

func (i *AttestationAtResourceGroup) ToAttestationAtResourceGroupOutput() AttestationAtResourceGroupOutput {
	return i.ToAttestationAtResourceGroupOutputWithContext(context.Background())
}

func (i *AttestationAtResourceGroup) ToAttestationAtResourceGroupOutputWithContext(ctx context.Context) AttestationAtResourceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationAtResourceGroupOutput)
}

type AttestationAtResourceGroupOutput struct{ *pulumi.OutputState }

func (AttestationAtResourceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestationAtResourceGroup)(nil)).Elem()
}

func (o AttestationAtResourceGroupOutput) ToAttestationAtResourceGroupOutput() AttestationAtResourceGroupOutput {
	return o
}

func (o AttestationAtResourceGroupOutput) ToAttestationAtResourceGroupOutputWithContext(ctx context.Context) AttestationAtResourceGroupOutput {
	return o
}

func (o AttestationAtResourceGroupOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtResourceGroup) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o AttestationAtResourceGroupOutput) ComplianceState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtResourceGroup) pulumi.StringPtrOutput { return v.ComplianceState }).(pulumi.StringPtrOutput)
}

func (o AttestationAtResourceGroupOutput) Evidence() AttestationEvidenceResponseArrayOutput {
	return o.ApplyT(func(v *AttestationAtResourceGroup) AttestationEvidenceResponseArrayOutput { return v.Evidence }).(AttestationEvidenceResponseArrayOutput)
}

func (o AttestationAtResourceGroupOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtResourceGroup) pulumi.StringPtrOutput { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

func (o AttestationAtResourceGroupOutput) LastComplianceStateChangeAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtResourceGroup) pulumi.StringOutput { return v.LastComplianceStateChangeAt }).(pulumi.StringOutput)
}

func (o AttestationAtResourceGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtResourceGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AttestationAtResourceGroupOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtResourceGroup) pulumi.StringPtrOutput { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o AttestationAtResourceGroupOutput) PolicyAssignmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtResourceGroup) pulumi.StringOutput { return v.PolicyAssignmentId }).(pulumi.StringOutput)
}

func (o AttestationAtResourceGroupOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtResourceGroup) pulumi.StringPtrOutput { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

func (o AttestationAtResourceGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtResourceGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AttestationAtResourceGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AttestationAtResourceGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AttestationAtResourceGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtResourceGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AttestationAtResourceGroupOutput{})
}
