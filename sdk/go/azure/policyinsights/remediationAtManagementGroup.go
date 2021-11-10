


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
	return reflect.TypeOf((*RemediationAtManagementGroup)(nil))
}

func (i *RemediationAtManagementGroup) ToRemediationAtManagementGroupOutput() RemediationAtManagementGroupOutput {
	return i.ToRemediationAtManagementGroupOutputWithContext(context.Background())
}

func (i *RemediationAtManagementGroup) ToRemediationAtManagementGroupOutputWithContext(ctx context.Context) RemediationAtManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationAtManagementGroupOutput)
}

type RemediationAtManagementGroupOutput struct{ *pulumi.OutputState }

func (RemediationAtManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationAtManagementGroup)(nil))
}

func (o RemediationAtManagementGroupOutput) ToRemediationAtManagementGroupOutput() RemediationAtManagementGroupOutput {
	return o
}

func (o RemediationAtManagementGroupOutput) ToRemediationAtManagementGroupOutputWithContext(ctx context.Context) RemediationAtManagementGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RemediationAtManagementGroupOutput{})
}
