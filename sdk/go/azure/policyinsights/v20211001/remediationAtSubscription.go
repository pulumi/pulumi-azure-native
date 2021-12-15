


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RemediationAtSubscription struct {
	pulumi.CustomResourceState

	CorrelationId               pulumi.StringOutput                                    `pulumi:"correlationId"`
	CreatedOn                   pulumi.StringOutput                                    `pulumi:"createdOn"`
	DeploymentStatus            RemediationDeploymentSummaryResponseOutput             `pulumi:"deploymentStatus"`
	FailureThreshold            RemediationPropertiesResponseFailureThresholdPtrOutput `pulumi:"failureThreshold"`
	Filters                     RemediationFiltersResponsePtrOutput                    `pulumi:"filters"`
	LastUpdatedOn               pulumi.StringOutput                                    `pulumi:"lastUpdatedOn"`
	Name                        pulumi.StringOutput                                    `pulumi:"name"`
	ParallelDeployments         pulumi.IntPtrOutput                                    `pulumi:"parallelDeployments"`
	PolicyAssignmentId          pulumi.StringPtrOutput                                 `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId pulumi.StringPtrOutput                                 `pulumi:"policyDefinitionReferenceId"`
	ProvisioningState           pulumi.StringOutput                                    `pulumi:"provisioningState"`
	ResourceCount               pulumi.IntPtrOutput                                    `pulumi:"resourceCount"`
	ResourceDiscoveryMode       pulumi.StringPtrOutput                                 `pulumi:"resourceDiscoveryMode"`
	StatusMessage               pulumi.StringOutput                                    `pulumi:"statusMessage"`
	SystemData                  SystemDataResponseOutput                               `pulumi:"systemData"`
	Type                        pulumi.StringOutput                                    `pulumi:"type"`
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
			Type: pulumi.String("azure-native:policyinsights/v20190701:RemediationAtSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource RemediationAtSubscription
	err := ctx.RegisterResource("azure-native:policyinsights/v20211001:RemediationAtSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRemediationAtSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationAtSubscriptionState, opts ...pulumi.ResourceOption) (*RemediationAtSubscription, error) {
	var resource RemediationAtSubscription
	err := ctx.ReadResource("azure-native:policyinsights/v20211001:RemediationAtSubscription", name, id, state, &resource, opts...)
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
	FailureThreshold            *RemediationPropertiesFailureThreshold `pulumi:"failureThreshold"`
	Filters                     *RemediationFilters                    `pulumi:"filters"`
	ParallelDeployments         *int                                   `pulumi:"parallelDeployments"`
	PolicyAssignmentId          *string                                `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string                                `pulumi:"policyDefinitionReferenceId"`
	RemediationName             *string                                `pulumi:"remediationName"`
	ResourceCount               *int                                   `pulumi:"resourceCount"`
	ResourceDiscoveryMode       *string                                `pulumi:"resourceDiscoveryMode"`
}


type RemediationAtSubscriptionArgs struct {
	FailureThreshold            RemediationPropertiesFailureThresholdPtrInput
	Filters                     RemediationFiltersPtrInput
	ParallelDeployments         pulumi.IntPtrInput
	PolicyAssignmentId          pulumi.StringPtrInput
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	RemediationName             pulumi.StringPtrInput
	ResourceCount               pulumi.IntPtrInput
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

func init() {
	pulumi.RegisterOutputType(RemediationAtSubscriptionOutput{})
}
