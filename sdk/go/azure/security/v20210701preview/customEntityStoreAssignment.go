


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomEntityStoreAssignment struct {
	pulumi.CustomResourceState

	EntityStoreDatabaseLink pulumi.StringPtrOutput   `pulumi:"entityStoreDatabaseLink"`
	Name                    pulumi.StringOutput      `pulumi:"name"`
	Principal               pulumi.StringPtrOutput   `pulumi:"principal"`
	SystemData              SystemDataResponseOutput `pulumi:"systemData"`
	Type                    pulumi.StringOutput      `pulumi:"type"`
}


func NewCustomEntityStoreAssignment(ctx *pulumi.Context,
	name string, args *CustomEntityStoreAssignmentArgs, opts ...pulumi.ResourceOption) (*CustomEntityStoreAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:CustomEntityStoreAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomEntityStoreAssignment
	err := ctx.RegisterResource("azure-native:security/v20210701preview:CustomEntityStoreAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomEntityStoreAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomEntityStoreAssignmentState, opts ...pulumi.ResourceOption) (*CustomEntityStoreAssignment, error) {
	var resource CustomEntityStoreAssignment
	err := ctx.ReadResource("azure-native:security/v20210701preview:CustomEntityStoreAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type customEntityStoreAssignmentState struct {
}

type CustomEntityStoreAssignmentState struct {
}

func (CustomEntityStoreAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*customEntityStoreAssignmentState)(nil)).Elem()
}

type customEntityStoreAssignmentArgs struct {
	CustomEntityStoreAssignmentName *string `pulumi:"customEntityStoreAssignmentName"`
	Principal                       *string `pulumi:"principal"`
	ResourceGroupName               string  `pulumi:"resourceGroupName"`
}


type CustomEntityStoreAssignmentArgs struct {
	CustomEntityStoreAssignmentName pulumi.StringPtrInput
	Principal                       pulumi.StringPtrInput
	ResourceGroupName               pulumi.StringInput
}

func (CustomEntityStoreAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customEntityStoreAssignmentArgs)(nil)).Elem()
}

type CustomEntityStoreAssignmentInput interface {
	pulumi.Input

	ToCustomEntityStoreAssignmentOutput() CustomEntityStoreAssignmentOutput
	ToCustomEntityStoreAssignmentOutputWithContext(ctx context.Context) CustomEntityStoreAssignmentOutput
}

func (*CustomEntityStoreAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomEntityStoreAssignment)(nil)).Elem()
}

func (i *CustomEntityStoreAssignment) ToCustomEntityStoreAssignmentOutput() CustomEntityStoreAssignmentOutput {
	return i.ToCustomEntityStoreAssignmentOutputWithContext(context.Background())
}

func (i *CustomEntityStoreAssignment) ToCustomEntityStoreAssignmentOutputWithContext(ctx context.Context) CustomEntityStoreAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomEntityStoreAssignmentOutput)
}

type CustomEntityStoreAssignmentOutput struct{ *pulumi.OutputState }

func (CustomEntityStoreAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomEntityStoreAssignment)(nil)).Elem()
}

func (o CustomEntityStoreAssignmentOutput) ToCustomEntityStoreAssignmentOutput() CustomEntityStoreAssignmentOutput {
	return o
}

func (o CustomEntityStoreAssignmentOutput) ToCustomEntityStoreAssignmentOutputWithContext(ctx context.Context) CustomEntityStoreAssignmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomEntityStoreAssignmentOutput{})
}
