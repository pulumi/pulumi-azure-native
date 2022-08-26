


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AttestationAtSubscription struct {
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


func NewAttestationAtSubscription(ctx *pulumi.Context,
	name string, args *AttestationAtSubscriptionArgs, opts ...pulumi.ResourceOption) (*AttestationAtSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyAssignmentId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyAssignmentId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:policyinsights:AttestationAtSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource AttestationAtSubscription
	err := ctx.RegisterResource("azure-native:policyinsights/v20210101:AttestationAtSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAttestationAtSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestationAtSubscriptionState, opts ...pulumi.ResourceOption) (*AttestationAtSubscription, error) {
	var resource AttestationAtSubscription
	err := ctx.ReadResource("azure-native:policyinsights/v20210101:AttestationAtSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type attestationAtSubscriptionState struct {
}

type AttestationAtSubscriptionState struct {
}

func (AttestationAtSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestationAtSubscriptionState)(nil)).Elem()
}

type attestationAtSubscriptionArgs struct {
	AttestationName             *string               `pulumi:"attestationName"`
	Comments                    *string               `pulumi:"comments"`
	ComplianceState             *string               `pulumi:"complianceState"`
	Evidence                    []AttestationEvidence `pulumi:"evidence"`
	ExpiresOn                   *string               `pulumi:"expiresOn"`
	Owner                       *string               `pulumi:"owner"`
	PolicyAssignmentId          string                `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string               `pulumi:"policyDefinitionReferenceId"`
}


type AttestationAtSubscriptionArgs struct {
	AttestationName             pulumi.StringPtrInput
	Comments                    pulumi.StringPtrInput
	ComplianceState             pulumi.StringPtrInput
	Evidence                    AttestationEvidenceArrayInput
	ExpiresOn                   pulumi.StringPtrInput
	Owner                       pulumi.StringPtrInput
	PolicyAssignmentId          pulumi.StringInput
	PolicyDefinitionReferenceId pulumi.StringPtrInput
}

func (AttestationAtSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestationAtSubscriptionArgs)(nil)).Elem()
}

type AttestationAtSubscriptionInput interface {
	pulumi.Input

	ToAttestationAtSubscriptionOutput() AttestationAtSubscriptionOutput
	ToAttestationAtSubscriptionOutputWithContext(ctx context.Context) AttestationAtSubscriptionOutput
}

func (*AttestationAtSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestationAtSubscription)(nil)).Elem()
}

func (i *AttestationAtSubscription) ToAttestationAtSubscriptionOutput() AttestationAtSubscriptionOutput {
	return i.ToAttestationAtSubscriptionOutputWithContext(context.Background())
}

func (i *AttestationAtSubscription) ToAttestationAtSubscriptionOutputWithContext(ctx context.Context) AttestationAtSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationAtSubscriptionOutput)
}

type AttestationAtSubscriptionOutput struct{ *pulumi.OutputState }

func (AttestationAtSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestationAtSubscription)(nil)).Elem()
}

func (o AttestationAtSubscriptionOutput) ToAttestationAtSubscriptionOutput() AttestationAtSubscriptionOutput {
	return o
}

func (o AttestationAtSubscriptionOutput) ToAttestationAtSubscriptionOutputWithContext(ctx context.Context) AttestationAtSubscriptionOutput {
	return o
}

func (o AttestationAtSubscriptionOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtSubscription) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o AttestationAtSubscriptionOutput) ComplianceState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtSubscription) pulumi.StringPtrOutput { return v.ComplianceState }).(pulumi.StringPtrOutput)
}

func (o AttestationAtSubscriptionOutput) Evidence() AttestationEvidenceResponseArrayOutput {
	return o.ApplyT(func(v *AttestationAtSubscription) AttestationEvidenceResponseArrayOutput { return v.Evidence }).(AttestationEvidenceResponseArrayOutput)
}

func (o AttestationAtSubscriptionOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtSubscription) pulumi.StringPtrOutput { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

func (o AttestationAtSubscriptionOutput) LastComplianceStateChangeAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtSubscription) pulumi.StringOutput { return v.LastComplianceStateChangeAt }).(pulumi.StringOutput)
}

func (o AttestationAtSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AttestationAtSubscriptionOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtSubscription) pulumi.StringPtrOutput { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o AttestationAtSubscriptionOutput) PolicyAssignmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtSubscription) pulumi.StringOutput { return v.PolicyAssignmentId }).(pulumi.StringOutput)
}

func (o AttestationAtSubscriptionOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationAtSubscription) pulumi.StringPtrOutput { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

func (o AttestationAtSubscriptionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtSubscription) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AttestationAtSubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AttestationAtSubscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AttestationAtSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationAtSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AttestationAtSubscriptionOutput{})
}
