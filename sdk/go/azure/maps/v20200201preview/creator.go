


package v20200201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Creator struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput             `pulumi:"location"`
	Name       pulumi.StringOutput             `pulumi:"name"`
	Properties CreatorPropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput          `pulumi:"tags"`
	Type       pulumi.StringOutput             `pulumi:"type"`
}


func NewCreator(ctx *pulumi.Context,
	name string, args *CreatorArgs, opts ...pulumi.ResourceOption) (*Creator, error) {
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
			Type: pulumi.String("azure-native:maps:Creator"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20210201:Creator"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20210701preview:Creator"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20211201preview:Creator"),
		},
	})
	opts = append(opts, aliases)
	var resource Creator
	err := ctx.RegisterResource("azure-native:maps/v20200201preview:Creator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCreator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CreatorState, opts ...pulumi.ResourceOption) (*Creator, error) {
	var resource Creator
	err := ctx.ReadResource("azure-native:maps/v20200201preview:Creator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type creatorState struct {
}

type CreatorState struct {
}

func (CreatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*creatorState)(nil)).Elem()
}

type creatorArgs struct {
	AccountName       string            `pulumi:"accountName"`
	CreatorName       *string           `pulumi:"creatorName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type CreatorArgs struct {
	AccountName       pulumi.StringInput
	CreatorName       pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (CreatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*creatorArgs)(nil)).Elem()
}

type CreatorInput interface {
	pulumi.Input

	ToCreatorOutput() CreatorOutput
	ToCreatorOutputWithContext(ctx context.Context) CreatorOutput
}

func (*Creator) ElementType() reflect.Type {
	return reflect.TypeOf((*Creator)(nil))
}

func (i *Creator) ToCreatorOutput() CreatorOutput {
	return i.ToCreatorOutputWithContext(context.Background())
}

func (i *Creator) ToCreatorOutputWithContext(ctx context.Context) CreatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatorOutput)
}

type CreatorOutput struct{ *pulumi.OutputState }

func (CreatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Creator)(nil))
}

func (o CreatorOutput) ToCreatorOutput() CreatorOutput {
	return o
}

func (o CreatorOutput) ToCreatorOutputWithContext(ctx context.Context) CreatorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CreatorOutput{})
}
