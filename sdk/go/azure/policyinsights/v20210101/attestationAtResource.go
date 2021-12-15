


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AttestationAtResource struct {
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


func NewAttestationAtResource(ctx *pulumi.Context,
	name string, args *AttestationAtResourceArgs, opts ...pulumi.ResourceOption) (*AttestationAtResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyAssignmentId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyAssignmentId'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:policyinsights:AttestationAtResource"),
		},
	})
	opts = append(opts, aliases)
	var resource AttestationAtResource
	err := ctx.RegisterResource("azure-native:policyinsights/v20210101:AttestationAtResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAttestationAtResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestationAtResourceState, opts ...pulumi.ResourceOption) (*AttestationAtResource, error) {
	var resource AttestationAtResource
	err := ctx.ReadResource("azure-native:policyinsights/v20210101:AttestationAtResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type attestationAtResourceState struct {
}

type AttestationAtResourceState struct {
}

func (AttestationAtResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestationAtResourceState)(nil)).Elem()
}

type attestationAtResourceArgs struct {
	AttestationName             *string               `pulumi:"attestationName"`
	Comments                    *string               `pulumi:"comments"`
	ComplianceState             *string               `pulumi:"complianceState"`
	Evidence                    []AttestationEvidence `pulumi:"evidence"`
	ExpiresOn                   *string               `pulumi:"expiresOn"`
	Owner                       *string               `pulumi:"owner"`
	PolicyAssignmentId          string                `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string               `pulumi:"policyDefinitionReferenceId"`
	ResourceId                  string                `pulumi:"resourceId"`
}


type AttestationAtResourceArgs struct {
	AttestationName             pulumi.StringPtrInput
	Comments                    pulumi.StringPtrInput
	ComplianceState             pulumi.StringPtrInput
	Evidence                    AttestationEvidenceArrayInput
	ExpiresOn                   pulumi.StringPtrInput
	Owner                       pulumi.StringPtrInput
	PolicyAssignmentId          pulumi.StringInput
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	ResourceId                  pulumi.StringInput
}

func (AttestationAtResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestationAtResourceArgs)(nil)).Elem()
}

type AttestationAtResourceInput interface {
	pulumi.Input

	ToAttestationAtResourceOutput() AttestationAtResourceOutput
	ToAttestationAtResourceOutputWithContext(ctx context.Context) AttestationAtResourceOutput
}

func (*AttestationAtResource) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestationAtResource)(nil)).Elem()
}

func (i *AttestationAtResource) ToAttestationAtResourceOutput() AttestationAtResourceOutput {
	return i.ToAttestationAtResourceOutputWithContext(context.Background())
}

func (i *AttestationAtResource) ToAttestationAtResourceOutputWithContext(ctx context.Context) AttestationAtResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationAtResourceOutput)
}

type AttestationAtResourceOutput struct{ *pulumi.OutputState }

func (AttestationAtResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestationAtResource)(nil)).Elem()
}

func (o AttestationAtResourceOutput) ToAttestationAtResourceOutput() AttestationAtResourceOutput {
	return o
}

func (o AttestationAtResourceOutput) ToAttestationAtResourceOutputWithContext(ctx context.Context) AttestationAtResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AttestationAtResourceOutput{})
}
