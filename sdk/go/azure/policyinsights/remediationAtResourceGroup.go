


package policyinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RemediationAtResourceGroup struct {
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


func NewRemediationAtResourceGroup(ctx *pulumi.Context,
	name string, args *RemediationAtResourceGroupArgs, opts ...pulumi.ResourceOption) (*RemediationAtResourceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:policyinsights/v20180701preview:RemediationAtResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20190701:RemediationAtResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20211001:RemediationAtResourceGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource RemediationAtResourceGroup
	err := ctx.RegisterResource("azure-native:policyinsights:RemediationAtResourceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRemediationAtResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationAtResourceGroupState, opts ...pulumi.ResourceOption) (*RemediationAtResourceGroup, error) {
	var resource RemediationAtResourceGroup
	err := ctx.ReadResource("azure-native:policyinsights:RemediationAtResourceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type remediationAtResourceGroupState struct {
}

type RemediationAtResourceGroupState struct {
}

func (RemediationAtResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtResourceGroupState)(nil)).Elem()
}

type remediationAtResourceGroupArgs struct {
	Filters                     *RemediationFilters `pulumi:"filters"`
	PolicyAssignmentId          *string             `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string             `pulumi:"policyDefinitionReferenceId"`
	RemediationName             *string             `pulumi:"remediationName"`
	ResourceDiscoveryMode       *string             `pulumi:"resourceDiscoveryMode"`
	ResourceGroupName           string              `pulumi:"resourceGroupName"`
}


type RemediationAtResourceGroupArgs struct {
	Filters                     RemediationFiltersPtrInput
	PolicyAssignmentId          pulumi.StringPtrInput
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	RemediationName             pulumi.StringPtrInput
	ResourceDiscoveryMode       pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
}

func (RemediationAtResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtResourceGroupArgs)(nil)).Elem()
}

type RemediationAtResourceGroupInput interface {
	pulumi.Input

	ToRemediationAtResourceGroupOutput() RemediationAtResourceGroupOutput
	ToRemediationAtResourceGroupOutputWithContext(ctx context.Context) RemediationAtResourceGroupOutput
}

func (*RemediationAtResourceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationAtResourceGroup)(nil)).Elem()
}

func (i *RemediationAtResourceGroup) ToRemediationAtResourceGroupOutput() RemediationAtResourceGroupOutput {
	return i.ToRemediationAtResourceGroupOutputWithContext(context.Background())
}

func (i *RemediationAtResourceGroup) ToRemediationAtResourceGroupOutputWithContext(ctx context.Context) RemediationAtResourceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationAtResourceGroupOutput)
}

type RemediationAtResourceGroupOutput struct{ *pulumi.OutputState }

func (RemediationAtResourceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationAtResourceGroup)(nil)).Elem()
}

func (o RemediationAtResourceGroupOutput) ToRemediationAtResourceGroupOutput() RemediationAtResourceGroupOutput {
	return o
}

func (o RemediationAtResourceGroupOutput) ToRemediationAtResourceGroupOutputWithContext(ctx context.Context) RemediationAtResourceGroupOutput {
	return o
}

func (o RemediationAtResourceGroupOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtResourceGroup) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o RemediationAtResourceGroupOutput) DeploymentStatus() RemediationDeploymentSummaryResponseOutput {
	return o.ApplyT(func(v *RemediationAtResourceGroup) RemediationDeploymentSummaryResponseOutput {
		return v.DeploymentStatus
	}).(RemediationDeploymentSummaryResponseOutput)
}

func (o RemediationAtResourceGroupOutput) Filters() RemediationFiltersResponsePtrOutput {
	return o.ApplyT(func(v *RemediationAtResourceGroup) RemediationFiltersResponsePtrOutput { return v.Filters }).(RemediationFiltersResponsePtrOutput)
}

func (o RemediationAtResourceGroupOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtResourceGroup) pulumi.StringOutput { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

func (o RemediationAtResourceGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtResourceGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RemediationAtResourceGroupOutput) PolicyAssignmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtResourceGroup) pulumi.StringPtrOutput { return v.PolicyAssignmentId }).(pulumi.StringPtrOutput)
}

func (o RemediationAtResourceGroupOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtResourceGroup) pulumi.StringPtrOutput { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

func (o RemediationAtResourceGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtResourceGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RemediationAtResourceGroupOutput) ResourceDiscoveryMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtResourceGroup) pulumi.StringPtrOutput { return v.ResourceDiscoveryMode }).(pulumi.StringPtrOutput)
}

func (o RemediationAtResourceGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtResourceGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RemediationAtResourceGroupOutput{})
}
