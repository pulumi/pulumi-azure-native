


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LinkedService struct {
	pulumi.CustomResourceState

	Identity   IdentityResponsePtrOutput        `pulumi:"identity"`
	Location   pulumi.StringPtrOutput           `pulumi:"location"`
	Name       pulumi.StringOutput              `pulumi:"name"`
	Properties LinkedServicePropsResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput              `pulumi:"type"`
}


func NewLinkedService(ctx *pulumi.Context,
	name string, args *LinkedServiceArgs, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:LinkedService"),
		},
	})
	opts = append(opts, aliases)
	var resource LinkedService
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20200901preview:LinkedService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLinkedService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceState, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	var resource LinkedService
	err := ctx.ReadResource("azure-native:machinelearningservices/v20200901preview:LinkedService", name, id, state, &resource, opts...)
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
	Identity          *Identity           `pulumi:"identity"`
	LinkName          *string             `pulumi:"linkName"`
	Location          *string             `pulumi:"location"`
	Name              *string             `pulumi:"name"`
	Properties        *LinkedServiceProps `pulumi:"properties"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	WorkspaceName     string              `pulumi:"workspaceName"`
}


type LinkedServiceArgs struct {
	Identity          IdentityPtrInput
	LinkName          pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	Properties        LinkedServicePropsPtrInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
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
	return reflect.TypeOf((**LinkedService)(nil)).Elem()
}

func (i *LinkedService) ToLinkedServiceOutput() LinkedServiceOutput {
	return i.ToLinkedServiceOutputWithContext(context.Background())
}

func (i *LinkedService) ToLinkedServiceOutputWithContext(ctx context.Context) LinkedServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceOutput)
}

type LinkedServiceOutput struct{ *pulumi.OutputState }

func (LinkedServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedService)(nil)).Elem()
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
