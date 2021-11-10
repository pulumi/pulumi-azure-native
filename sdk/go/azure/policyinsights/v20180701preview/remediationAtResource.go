


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RemediationAtResource struct {
	pulumi.CustomResourceState

	CreatedOn                   pulumi.StringOutput                           `pulumi:"createdOn"`
	DeploymentStatus            RemediationDeploymentSummaryResponsePtrOutput `pulumi:"deploymentStatus"`
	Filters                     RemediationFiltersResponsePtrOutput           `pulumi:"filters"`
	LastUpdatedOn               pulumi.StringOutput                           `pulumi:"lastUpdatedOn"`
	Name                        pulumi.StringOutput                           `pulumi:"name"`
	PolicyAssignmentId          pulumi.StringPtrOutput                        `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId pulumi.StringPtrOutput                        `pulumi:"policyDefinitionReferenceId"`
	ProvisioningState           pulumi.StringOutput                           `pulumi:"provisioningState"`
	Type                        pulumi.StringOutput                           `pulumi:"type"`
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
			Type: pulumi.String("azure-native:policyinsights:RemediationAtResource"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20190701:RemediationAtResource"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20211001:RemediationAtResource"),
		},
	})
	opts = append(opts, aliases)
	var resource RemediationAtResource
	err := ctx.RegisterResource("azure-native:policyinsights/v20180701preview:RemediationAtResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRemediationAtResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationAtResourceState, opts ...pulumi.ResourceOption) (*RemediationAtResource, error) {
	var resource RemediationAtResource
	err := ctx.ReadResource("azure-native:policyinsights/v20180701preview:RemediationAtResource", name, id, state, &resource, opts...)
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
	DeploymentStatus            *RemediationDeploymentSummary `pulumi:"deploymentStatus"`
	Filters                     *RemediationFilters           `pulumi:"filters"`
	PolicyAssignmentId          *string                       `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string                       `pulumi:"policyDefinitionReferenceId"`
	RemediationName             *string                       `pulumi:"remediationName"`
	ResourceId                  string                        `pulumi:"resourceId"`
}


type RemediationAtResourceArgs struct {
	DeploymentStatus            RemediationDeploymentSummaryPtrInput
	Filters                     RemediationFiltersPtrInput
	PolicyAssignmentId          pulumi.StringPtrInput
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	RemediationName             pulumi.StringPtrInput
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
