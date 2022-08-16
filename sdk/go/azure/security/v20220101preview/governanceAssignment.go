


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GovernanceAssignment struct {
	pulumi.CustomResourceState

	AdditionalData              GovernanceAssignmentAdditionalDataResponsePtrOutput `pulumi:"additionalData"`
	GovernanceEmailNotification GovernanceEmailNotificationResponsePtrOutput        `pulumi:"governanceEmailNotification"`
	IsGracePeriod               pulumi.BoolPtrOutput                                `pulumi:"isGracePeriod"`
	Name                        pulumi.StringOutput                                 `pulumi:"name"`
	Owner                       pulumi.StringPtrOutput                              `pulumi:"owner"`
	RemediationDueDate          pulumi.StringOutput                                 `pulumi:"remediationDueDate"`
	RemediationEta              RemediationEtaResponsePtrOutput                     `pulumi:"remediationEta"`
	Type                        pulumi.StringOutput                                 `pulumi:"type"`
}


func NewGovernanceAssignment(ctx *pulumi.Context,
	name string, args *GovernanceAssignmentArgs, opts ...pulumi.ResourceOption) (*GovernanceAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssessmentName == nil {
		return nil, errors.New("invalid value for required argument 'AssessmentName'")
	}
	if args.RemediationDueDate == nil {
		return nil, errors.New("invalid value for required argument 'RemediationDueDate'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	var resource GovernanceAssignment
	err := ctx.RegisterResource("azure-native:security/v20220101preview:GovernanceAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGovernanceAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GovernanceAssignmentState, opts ...pulumi.ResourceOption) (*GovernanceAssignment, error) {
	var resource GovernanceAssignment
	err := ctx.ReadResource("azure-native:security/v20220101preview:GovernanceAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type governanceAssignmentState struct {
}

type GovernanceAssignmentState struct {
}

func (GovernanceAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*governanceAssignmentState)(nil)).Elem()
}

type governanceAssignmentArgs struct {
	AdditionalData              *GovernanceAssignmentAdditionalData `pulumi:"additionalData"`
	AssessmentName              string                              `pulumi:"assessmentName"`
	AssignmentKey               *string                             `pulumi:"assignmentKey"`
	GovernanceEmailNotification *GovernanceEmailNotification        `pulumi:"governanceEmailNotification"`
	IsGracePeriod               *bool                               `pulumi:"isGracePeriod"`
	Owner                       *string                             `pulumi:"owner"`
	RemediationDueDate          string                              `pulumi:"remediationDueDate"`
	RemediationEta              *RemediationEta                     `pulumi:"remediationEta"`
	Scope                       string                              `pulumi:"scope"`
}


type GovernanceAssignmentArgs struct {
	AdditionalData              GovernanceAssignmentAdditionalDataPtrInput
	AssessmentName              pulumi.StringInput
	AssignmentKey               pulumi.StringPtrInput
	GovernanceEmailNotification GovernanceEmailNotificationPtrInput
	IsGracePeriod               pulumi.BoolPtrInput
	Owner                       pulumi.StringPtrInput
	RemediationDueDate          pulumi.StringInput
	RemediationEta              RemediationEtaPtrInput
	Scope                       pulumi.StringInput
}

func (GovernanceAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*governanceAssignmentArgs)(nil)).Elem()
}

type GovernanceAssignmentInput interface {
	pulumi.Input

	ToGovernanceAssignmentOutput() GovernanceAssignmentOutput
	ToGovernanceAssignmentOutputWithContext(ctx context.Context) GovernanceAssignmentOutput
}

func (*GovernanceAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceAssignment)(nil)).Elem()
}

func (i *GovernanceAssignment) ToGovernanceAssignmentOutput() GovernanceAssignmentOutput {
	return i.ToGovernanceAssignmentOutputWithContext(context.Background())
}

func (i *GovernanceAssignment) ToGovernanceAssignmentOutputWithContext(ctx context.Context) GovernanceAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GovernanceAssignmentOutput)
}

type GovernanceAssignmentOutput struct{ *pulumi.OutputState }

func (GovernanceAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceAssignment)(nil)).Elem()
}

func (o GovernanceAssignmentOutput) ToGovernanceAssignmentOutput() GovernanceAssignmentOutput {
	return o
}

func (o GovernanceAssignmentOutput) ToGovernanceAssignmentOutputWithContext(ctx context.Context) GovernanceAssignmentOutput {
	return o
}

func (o GovernanceAssignmentOutput) AdditionalData() GovernanceAssignmentAdditionalDataResponsePtrOutput {
	return o.ApplyT(func(v *GovernanceAssignment) GovernanceAssignmentAdditionalDataResponsePtrOutput {
		return v.AdditionalData
	}).(GovernanceAssignmentAdditionalDataResponsePtrOutput)
}

func (o GovernanceAssignmentOutput) GovernanceEmailNotification() GovernanceEmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v *GovernanceAssignment) GovernanceEmailNotificationResponsePtrOutput {
		return v.GovernanceEmailNotification
	}).(GovernanceEmailNotificationResponsePtrOutput)
}

func (o GovernanceAssignmentOutput) IsGracePeriod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GovernanceAssignment) pulumi.BoolPtrOutput { return v.IsGracePeriod }).(pulumi.BoolPtrOutput)
}

func (o GovernanceAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GovernanceAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GovernanceAssignmentOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GovernanceAssignment) pulumi.StringPtrOutput { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o GovernanceAssignmentOutput) RemediationDueDate() pulumi.StringOutput {
	return o.ApplyT(func(v *GovernanceAssignment) pulumi.StringOutput { return v.RemediationDueDate }).(pulumi.StringOutput)
}

func (o GovernanceAssignmentOutput) RemediationEta() RemediationEtaResponsePtrOutput {
	return o.ApplyT(func(v *GovernanceAssignment) RemediationEtaResponsePtrOutput { return v.RemediationEta }).(RemediationEtaResponsePtrOutput)
}

func (o GovernanceAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GovernanceAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GovernanceAssignmentOutput{})
}
