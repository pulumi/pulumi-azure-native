


package v20221201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CommitmentPlanAssociation struct {
	pulumi.CustomResourceState

	AccountId  pulumi.StringPtrOutput   `pulumi:"accountId"`
	Etag       pulumi.StringOutput      `pulumi:"etag"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewCommitmentPlanAssociation(ctx *pulumi.Context,
	name string, args *CommitmentPlanAssociationArgs, opts ...pulumi.ResourceOption) (*CommitmentPlanAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CommitmentPlanName == nil {
		return nil, errors.New("invalid value for required argument 'CommitmentPlanName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource CommitmentPlanAssociation
	err := ctx.RegisterResource("azure-native:cognitiveservices/v20221201:CommitmentPlanAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCommitmentPlanAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommitmentPlanAssociationState, opts ...pulumi.ResourceOption) (*CommitmentPlanAssociation, error) {
	var resource CommitmentPlanAssociation
	err := ctx.ReadResource("azure-native:cognitiveservices/v20221201:CommitmentPlanAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type commitmentPlanAssociationState struct {
}

type CommitmentPlanAssociationState struct {
}

func (CommitmentPlanAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*commitmentPlanAssociationState)(nil)).Elem()
}

type commitmentPlanAssociationArgs struct {
	AccountId                     *string `pulumi:"accountId"`
	CommitmentPlanAssociationName *string `pulumi:"commitmentPlanAssociationName"`
	CommitmentPlanName            string  `pulumi:"commitmentPlanName"`
	ResourceGroupName             string  `pulumi:"resourceGroupName"`
}


type CommitmentPlanAssociationArgs struct {
	AccountId                     pulumi.StringPtrInput
	CommitmentPlanAssociationName pulumi.StringPtrInput
	CommitmentPlanName            pulumi.StringInput
	ResourceGroupName             pulumi.StringInput
}

func (CommitmentPlanAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commitmentPlanAssociationArgs)(nil)).Elem()
}

type CommitmentPlanAssociationInput interface {
	pulumi.Input

	ToCommitmentPlanAssociationOutput() CommitmentPlanAssociationOutput
	ToCommitmentPlanAssociationOutputWithContext(ctx context.Context) CommitmentPlanAssociationOutput
}

func (*CommitmentPlanAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanAssociation)(nil)).Elem()
}

func (i *CommitmentPlanAssociation) ToCommitmentPlanAssociationOutput() CommitmentPlanAssociationOutput {
	return i.ToCommitmentPlanAssociationOutputWithContext(context.Background())
}

func (i *CommitmentPlanAssociation) ToCommitmentPlanAssociationOutputWithContext(ctx context.Context) CommitmentPlanAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanAssociationOutput)
}

type CommitmentPlanAssociationOutput struct{ *pulumi.OutputState }

func (CommitmentPlanAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanAssociation)(nil)).Elem()
}

func (o CommitmentPlanAssociationOutput) ToCommitmentPlanAssociationOutput() CommitmentPlanAssociationOutput {
	return o
}

func (o CommitmentPlanAssociationOutput) ToCommitmentPlanAssociationOutputWithContext(ctx context.Context) CommitmentPlanAssociationOutput {
	return o
}

func (o CommitmentPlanAssociationOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanAssociation) pulumi.StringPtrOutput { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o CommitmentPlanAssociationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlanAssociation) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o CommitmentPlanAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlanAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CommitmentPlanAssociationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CommitmentPlanAssociation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o CommitmentPlanAssociationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlanAssociation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CommitmentPlanAssociationOutput{})
}
