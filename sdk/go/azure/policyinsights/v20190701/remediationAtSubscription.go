


package v20190701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RemediationAtSubscription struct {
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


func NewRemediationAtSubscription(ctx *pulumi.Context,
	name string, args *RemediationAtSubscriptionArgs, opts ...pulumi.ResourceOption) (*RemediationAtSubscription, error) {
	if args == nil {
		args = &RemediationAtSubscriptionArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:policyinsights:RemediationAtSubscription"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20180701preview:RemediationAtSubscription"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20211001:RemediationAtSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource RemediationAtSubscription
	err := ctx.RegisterResource("azure-native:policyinsights/v20190701:RemediationAtSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRemediationAtSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationAtSubscriptionState, opts ...pulumi.ResourceOption) (*RemediationAtSubscription, error) {
	var resource RemediationAtSubscription
	err := ctx.ReadResource("azure-native:policyinsights/v20190701:RemediationAtSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type remediationAtSubscriptionState struct {
}

type RemediationAtSubscriptionState struct {
}

func (RemediationAtSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtSubscriptionState)(nil)).Elem()
}

type remediationAtSubscriptionArgs struct {
	Filters                     *RemediationFilters `pulumi:"filters"`
	PolicyAssignmentId          *string             `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string             `pulumi:"policyDefinitionReferenceId"`
	RemediationName             *string             `pulumi:"remediationName"`
	ResourceDiscoveryMode       *string             `pulumi:"resourceDiscoveryMode"`
}


type RemediationAtSubscriptionArgs struct {
	Filters                     RemediationFiltersPtrInput
	PolicyAssignmentId          pulumi.StringPtrInput
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	RemediationName             pulumi.StringPtrInput
	ResourceDiscoveryMode       pulumi.StringPtrInput
}

func (RemediationAtSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtSubscriptionArgs)(nil)).Elem()
}

type RemediationAtSubscriptionInput interface {
	pulumi.Input

	ToRemediationAtSubscriptionOutput() RemediationAtSubscriptionOutput
	ToRemediationAtSubscriptionOutputWithContext(ctx context.Context) RemediationAtSubscriptionOutput
}

func (*RemediationAtSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationAtSubscription)(nil)).Elem()
}

func (i *RemediationAtSubscription) ToRemediationAtSubscriptionOutput() RemediationAtSubscriptionOutput {
	return i.ToRemediationAtSubscriptionOutputWithContext(context.Background())
}

func (i *RemediationAtSubscription) ToRemediationAtSubscriptionOutputWithContext(ctx context.Context) RemediationAtSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationAtSubscriptionOutput)
}

type RemediationAtSubscriptionOutput struct{ *pulumi.OutputState }

func (RemediationAtSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationAtSubscription)(nil)).Elem()
}

func (o RemediationAtSubscriptionOutput) ToRemediationAtSubscriptionOutput() RemediationAtSubscriptionOutput {
	return o
}

func (o RemediationAtSubscriptionOutput) ToRemediationAtSubscriptionOutputWithContext(ctx context.Context) RemediationAtSubscriptionOutput {
	return o
}

func (o RemediationAtSubscriptionOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtSubscription) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o RemediationAtSubscriptionOutput) DeploymentStatus() RemediationDeploymentSummaryResponseOutput {
	return o.ApplyT(func(v *RemediationAtSubscription) RemediationDeploymentSummaryResponseOutput {
		return v.DeploymentStatus
	}).(RemediationDeploymentSummaryResponseOutput)
}

func (o RemediationAtSubscriptionOutput) Filters() RemediationFiltersResponsePtrOutput {
	return o.ApplyT(func(v *RemediationAtSubscription) RemediationFiltersResponsePtrOutput { return v.Filters }).(RemediationFiltersResponsePtrOutput)
}

func (o RemediationAtSubscriptionOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtSubscription) pulumi.StringOutput { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

func (o RemediationAtSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RemediationAtSubscriptionOutput) PolicyAssignmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtSubscription) pulumi.StringPtrOutput { return v.PolicyAssignmentId }).(pulumi.StringPtrOutput)
}

func (o RemediationAtSubscriptionOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtSubscription) pulumi.StringPtrOutput { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

func (o RemediationAtSubscriptionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtSubscription) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RemediationAtSubscriptionOutput) ResourceDiscoveryMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtSubscription) pulumi.StringPtrOutput { return v.ResourceDiscoveryMode }).(pulumi.StringPtrOutput)
}

func (o RemediationAtSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RemediationAtSubscriptionOutput{})
}
