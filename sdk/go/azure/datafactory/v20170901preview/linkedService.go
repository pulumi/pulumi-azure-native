


package v20170901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LinkedService struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput `pulumi:"etag"`
	Name       pulumi.StringOutput `pulumi:"name"`
	Properties pulumi.AnyOutput    `pulumi:"properties"`
	Type       pulumi.StringOutput `pulumi:"type"`
}


func NewLinkedService(ctx *pulumi.Context,
	name string, args *LinkedServiceArgs, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FactoryName == nil {
		return nil, errors.New("invalid value for required argument 'FactoryName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datafactory/v20170901preview:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:datafactory:LinkedService"),
		},
		{
			Type: pulumi.String("azure-nextgen:datafactory:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:datafactory/v20180601:LinkedService"),
		},
		{
			Type: pulumi.String("azure-nextgen:datafactory/v20180601:LinkedService"),
		},
	})
	opts = append(opts, aliases)
	var resource LinkedService
	err := ctx.RegisterResource("azure-native:datafactory/v20170901preview:LinkedService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLinkedService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceState, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	var resource LinkedService
	err := ctx.ReadResource("azure-native:datafactory/v20170901preview:LinkedService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type linkedServiceState struct {
}

type LinkedServiceState struct {
}

func (LinkedServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceState)(nil)).Elem()
}

type linkedServiceArgs struct {
	FactoryName       string      `pulumi:"factoryName"`
	LinkedServiceName *string     `pulumi:"linkedServiceName"`
	Properties        interface{} `pulumi:"properties"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
}


type LinkedServiceArgs struct {
	FactoryName       pulumi.StringInput
	LinkedServiceName pulumi.StringPtrInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
}

func (LinkedServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceArgs)(nil)).Elem()
}

type LinkedServiceInput interface {
	pulumi.Input

	ToLinkedServiceOutput() LinkedServiceOutput
	ToLinkedServiceOutputWithContext(ctx context.Context) LinkedServiceOutput
}

func (*LinkedService) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedService)(nil))
}

func (i *LinkedService) ToLinkedServiceOutput() LinkedServiceOutput {
	return i.ToLinkedServiceOutputWithContext(context.Background())
}

func (i *LinkedService) ToLinkedServiceOutputWithContext(ctx context.Context) LinkedServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceOutput)
}

type LinkedServiceOutput struct{ *pulumi.OutputState }

func (LinkedServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedService)(nil))
}

func (o LinkedServiceOutput) ToLinkedServiceOutput() LinkedServiceOutput {
	return o
}

func (o LinkedServiceOutput) ToLinkedServiceOutputWithContext(ctx context.Context) LinkedServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceOutput{})
}
