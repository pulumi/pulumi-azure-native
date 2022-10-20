


package v20211001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CommitmentPlan struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput                    `pulumi:"etag"`
	Name       pulumi.StringOutput                    `pulumi:"name"`
	Properties CommitmentPlanPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput               `pulumi:"systemData"`
	Type       pulumi.StringOutput                    `pulumi:"type"`
}


func NewCommitmentPlan(ctx *pulumi.Context,
	name string, args *CommitmentPlanArgs, opts ...pulumi.ResourceOption) (*CommitmentPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cognitiveservices:CommitmentPlan"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20220301:CommitmentPlan"),
		},
	})
	opts = append(opts, aliases)
	var resource CommitmentPlan
	err := ctx.RegisterResource("azure-native:cognitiveservices/v20211001:CommitmentPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCommitmentPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommitmentPlanState, opts ...pulumi.ResourceOption) (*CommitmentPlan, error) {
	var resource CommitmentPlan
	err := ctx.ReadResource("azure-native:cognitiveservices/v20211001:CommitmentPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type commitmentPlanState struct {
}

type CommitmentPlanState struct {
}

func (CommitmentPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*commitmentPlanState)(nil)).Elem()
}

type commitmentPlanArgs struct {
	AccountName        string                    `pulumi:"accountName"`
	CommitmentPlanName *string                   `pulumi:"commitmentPlanName"`
	Properties         *CommitmentPlanProperties `pulumi:"properties"`
	ResourceGroupName  string                    `pulumi:"resourceGroupName"`
}


type CommitmentPlanArgs struct {
	AccountName        pulumi.StringInput
	CommitmentPlanName pulumi.StringPtrInput
	Properties         CommitmentPlanPropertiesPtrInput
	ResourceGroupName  pulumi.StringInput
}

func (CommitmentPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commitmentPlanArgs)(nil)).Elem()
}

type CommitmentPlanInput interface {
	pulumi.Input

	ToCommitmentPlanOutput() CommitmentPlanOutput
	ToCommitmentPlanOutputWithContext(ctx context.Context) CommitmentPlanOutput
}

func (*CommitmentPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlan)(nil)).Elem()
}

func (i *CommitmentPlan) ToCommitmentPlanOutput() CommitmentPlanOutput {
	return i.ToCommitmentPlanOutputWithContext(context.Background())
}

func (i *CommitmentPlan) ToCommitmentPlanOutputWithContext(ctx context.Context) CommitmentPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanOutput)
}

type CommitmentPlanOutput struct{ *pulumi.OutputState }

func (CommitmentPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlan)(nil)).Elem()
}

func (o CommitmentPlanOutput) ToCommitmentPlanOutput() CommitmentPlanOutput {
	return o
}

func (o CommitmentPlanOutput) ToCommitmentPlanOutputWithContext(ctx context.Context) CommitmentPlanOutput {
	return o
}

func (o CommitmentPlanOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o CommitmentPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CommitmentPlanOutput) Properties() CommitmentPlanPropertiesResponseOutput {
	return o.ApplyT(func(v *CommitmentPlan) CommitmentPlanPropertiesResponseOutput { return v.Properties }).(CommitmentPlanPropertiesResponseOutput)
}

func (o CommitmentPlanOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CommitmentPlan) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o CommitmentPlanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CommitmentPlan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CommitmentPlanOutput{})
}
