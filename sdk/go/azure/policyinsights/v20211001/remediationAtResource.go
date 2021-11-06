


package v20211001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RemediationAtResource struct {
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


func NewRemediationAtResource(ctx *pulumi.Context,
	name string, args *RemediationAtResourceArgs, opts ...pulumi.ResourceOption) (*RemediationAtResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:policyinsights/v20211001:RemediationAtResource"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights:RemediationAtResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:policyinsights:RemediationAtResource"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20180701preview:RemediationAtResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:policyinsights/v20180701preview:RemediationAtResource"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20190701:RemediationAtResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:policyinsights/v20190701:RemediationAtResource"),
		},
	})
	opts = append(opts, aliases)
	var resource RemediationAtResource
	err := ctx.RegisterResource("azure-native:policyinsights/v20211001:RemediationAtResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRemediationAtResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationAtResourceState, opts ...pulumi.ResourceOption) (*RemediationAtResource, error) {
	var resource RemediationAtResource
	err := ctx.ReadResource("azure-native:policyinsights/v20211001:RemediationAtResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type remediationAtResourceState struct {
}

type RemediationAtResourceState struct {
}

func (RemediationAtResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtResourceState)(nil)).Elem()
}

type remediationAtResourceArgs struct {
	FailureThreshold            *RemediationPropertiesFailureThreshold `pulumi:"failureThreshold"`
	Filters                     *RemediationFilters                    `pulumi:"filters"`
	ParallelDeployments         *int                                   `pulumi:"parallelDeployments"`
	PolicyAssignmentId          *string                                `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string                                `pulumi:"policyDefinitionReferenceId"`
	RemediationName             *string                                `pulumi:"remediationName"`
	ResourceCount               *int                                   `pulumi:"resourceCount"`
	ResourceDiscoveryMode       *string                                `pulumi:"resourceDiscoveryMode"`
	ResourceId                  string                                 `pulumi:"resourceId"`
}


type RemediationAtResourceArgs struct {
	FailureThreshold            RemediationPropertiesFailureThresholdPtrInput
	Filters                     RemediationFiltersPtrInput
	ParallelDeployments         pulumi.IntPtrInput
	PolicyAssignmentId          pulumi.StringPtrInput
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	RemediationName             pulumi.StringPtrInput
	ResourceCount               pulumi.IntPtrInput
	ResourceDiscoveryMode       pulumi.StringPtrInput
	ResourceId                  pulumi.StringInput
}

func (RemediationAtResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtResourceArgs)(nil)).Elem()
}

type RemediationAtResourceInput interface {
	pulumi.Input

	ToRemediationAtResourceOutput() RemediationAtResourceOutput
	ToRemediationAtResourceOutputWithContext(ctx context.Context) RemediationAtResourceOutput
}

func (*RemediationAtResource) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationAtResource)(nil))
}

func (i *RemediationAtResource) ToRemediationAtResourceOutput() RemediationAtResourceOutput {
	return i.ToRemediationAtResourceOutputWithContext(context.Background())
}

func (i *RemediationAtResource) ToRemediationAtResourceOutputWithContext(ctx context.Context) RemediationAtResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationAtResourceOutput)
}

type RemediationAtResourceOutput struct{ *pulumi.OutputState }

func (RemediationAtResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationAtResource)(nil))
}

func (o RemediationAtResourceOutput) ToRemediationAtResourceOutput() RemediationAtResourceOutput {
	return o
}

func (o RemediationAtResourceOutput) ToRemediationAtResourceOutputWithContext(ctx context.Context) RemediationAtResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RemediationAtResourceOutput{})
}
