


package policyinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RemediationAtManagementGroup struct {
	pulumi.CustomResourceState

	CreatedOn                   pulumi.StringOutput                        `pulumi:"createdOn"`
	DeploymentStatus            RemediationDeploymentSummaryResponseOutput `pulumi:"deploymentStatus"`
	Filters                     RemediationFiltersResponsePtrOutput        `pulumi:"filters"`
	LastUpdatedOn               pulumi.StringOutput                        `pulumi:"lastUpdatedOn"`
	Name                        pulumi.StringOutput                        `pulumi:"name"`
	PolicyAssignmentId          pulumi.StringPtrOutput                     `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId pulumi.StringPtrOutput                     `pulumi:"policyDefinitionReferenceId"`
	ProvisioningState           pulumi.StringOutput                        `pulumi:"provisioningState"`
	ResourceDiscoveryMode       pulumi.StringPtrOutput                     `pulumi:"resourceDiscoveryMode"`
	Type                        pulumi.StringOutput                        `pulumi:"type"`
}


func NewRemediationAtManagementGroup(ctx *pulumi.Context,
	name string, args *RemediationAtManagementGroupArgs, opts ...pulumi.ResourceOption) (*RemediationAtManagementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	if args.ManagementGroupsNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupsNamespace'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:policyinsights/v20180701preview:RemediationAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20190701:RemediationAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20211001:RemediationAtManagementGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource RemediationAtManagementGroup
	err := ctx.RegisterResource("azure-native:policyinsights:RemediationAtManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRemediationAtManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationAtManagementGroupState, opts ...pulumi.ResourceOption) (*RemediationAtManagementGroup, error) {
	var resource RemediationAtManagementGroup
	err := ctx.ReadResource("azure-native:policyinsights:RemediationAtManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type remediationAtManagementGroupState struct {
}

type RemediationAtManagementGroupState struct {
}

func (RemediationAtManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtManagementGroupState)(nil)).Elem()
}

type remediationAtManagementGroupArgs struct {
	Filters                     *RemediationFilters `pulumi:"filters"`
	ManagementGroupId           string              `pulumi:"managementGroupId"`
	ManagementGroupsNamespace   string              `pulumi:"managementGroupsNamespace"`
	PolicyAssignmentId          *string             `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string             `pulumi:"policyDefinitionReferenceId"`
	RemediationName             *string             `pulumi:"remediationName"`
	ResourceDiscoveryMode       *string             `pulumi:"resourceDiscoveryMode"`
}


type RemediationAtManagementGroupArgs struct {
	Filters                     RemediationFiltersPtrInput
	ManagementGroupId           pulumi.StringInput
	ManagementGroupsNamespace   pulumi.StringInput
	PolicyAssignmentId          pulumi.StringPtrInput
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	RemediationName             pulumi.StringPtrInput
	ResourceDiscoveryMode       pulumi.StringPtrInput
}

func (RemediationAtManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtManagementGroupArgs)(nil)).Elem()
}

type RemediationAtManagementGroupInput interface {
	pulumi.Input

	ToRemediationAtManagementGroupOutput() RemediationAtManagementGroupOutput
	ToRemediationAtManagementGroupOutputWithContext(ctx context.Context) RemediationAtManagementGroupOutput
}

func (*RemediationAtManagementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationAtManagementGroup)(nil)).Elem()
}

func (i *RemediationAtManagementGroup) ToRemediationAtManagementGroupOutput() RemediationAtManagementGroupOutput {
	return i.ToRemediationAtManagementGroupOutputWithContext(context.Background())
}

func (i *RemediationAtManagementGroup) ToRemediationAtManagementGroupOutputWithContext(ctx context.Context) RemediationAtManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationAtManagementGroupOutput)
}

type RemediationAtManagementGroupOutput struct{ *pulumi.OutputState }

func (RemediationAtManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationAtManagementGroup)(nil)).Elem()
}

func (o RemediationAtManagementGroupOutput) ToRemediationAtManagementGroupOutput() RemediationAtManagementGroupOutput {
	return o
}

func (o RemediationAtManagementGroupOutput) ToRemediationAtManagementGroupOutputWithContext(ctx context.Context) RemediationAtManagementGroupOutput {
	return o
}

func (o RemediationAtManagementGroupOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o RemediationAtManagementGroupOutput) DeploymentStatus() RemediationDeploymentSummaryResponseOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) RemediationDeploymentSummaryResponseOutput {
		return v.DeploymentStatus
	}).(RemediationDeploymentSummaryResponseOutput)
}

func (o RemediationAtManagementGroupOutput) Filters() RemediationFiltersResponsePtrOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) RemediationFiltersResponsePtrOutput { return v.Filters }).(RemediationFiltersResponsePtrOutput)
}

func (o RemediationAtManagementGroupOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringOutput { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

func (o RemediationAtManagementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RemediationAtManagementGroupOutput) PolicyAssignmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringPtrOutput { return v.PolicyAssignmentId }).(pulumi.StringPtrOutput)
}

func (o RemediationAtManagementGroupOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringPtrOutput { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

func (o RemediationAtManagementGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RemediationAtManagementGroupOutput) ResourceDiscoveryMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringPtrOutput { return v.ResourceDiscoveryMode }).(pulumi.StringPtrOutput)
}

func (o RemediationAtManagementGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RemediationAtManagementGroupOutput{})
}
