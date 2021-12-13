


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MoveResource struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                  `pulumi:"name"`
	Properties MoveResourcePropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                  `pulumi:"type"`
}


func NewMoveResource(ctx *pulumi.Context,
	name string, args *MoveResourceArgs, opts ...pulumi.ResourceOption) (*MoveResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MoveCollectionName == nil {
		return nil, errors.New("invalid value for required argument 'MoveCollectionName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate:MoveResource"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20191001preview:MoveResource"),
		},
	})
	opts = append(opts, aliases)
	var resource MoveResource
	err := ctx.RegisterResource("azure-native:migrate/v20210101:MoveResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMoveResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MoveResourceState, opts ...pulumi.ResourceOption) (*MoveResource, error) {
	var resource MoveResource
	err := ctx.ReadResource("azure-native:migrate/v20210101:MoveResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type moveResourceState struct {
}

type MoveResourceState struct {
}

func (MoveResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*moveResourceState)(nil)).Elem()
}

type moveResourceArgs struct {
	MoveCollectionName string                  `pulumi:"moveCollectionName"`
	MoveResourceName   *string                 `pulumi:"moveResourceName"`
	Properties         *MoveResourceProperties `pulumi:"properties"`
	ResourceGroupName  string                  `pulumi:"resourceGroupName"`
}


type MoveResourceArgs struct {
	MoveCollectionName pulumi.StringInput
	MoveResourceName   pulumi.StringPtrInput
	Properties         MoveResourcePropertiesPtrInput
	ResourceGroupName  pulumi.StringInput
}

func (MoveResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*moveResourceArgs)(nil)).Elem()
}

type MoveResourceInput interface {
	pulumi.Input

	ToMoveResourceOutput() MoveResourceOutput
	ToMoveResourceOutputWithContext(ctx context.Context) MoveResourceOutput
}

func (*MoveResource) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResource)(nil)).Elem()
}

func (i *MoveResource) ToMoveResourceOutput() MoveResourceOutput {
	return i.ToMoveResourceOutputWithContext(context.Background())
}

func (i *MoveResource) ToMoveResourceOutputWithContext(ctx context.Context) MoveResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceOutput)
}

type MoveResourceOutput struct{ *pulumi.OutputState }

func (MoveResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResource)(nil)).Elem()
}

func (o MoveResourceOutput) ToMoveResourceOutput() MoveResourceOutput {
	return o
}

func (o MoveResourceOutput) ToMoveResourceOutputWithContext(ctx context.Context) MoveResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MoveResourceOutput{})
}
