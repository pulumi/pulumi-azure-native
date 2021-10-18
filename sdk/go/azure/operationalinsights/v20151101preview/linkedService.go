


package v20151101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LinkedService struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput `pulumi:"name"`
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	Type       pulumi.StringOutput `pulumi:"type"`
}


func NewLinkedService(ctx *pulumi.Context,
	name string, args *LinkedServiceArgs, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20151101preview:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights:LinkedService"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190801preview:LinkedService"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20190801preview:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:LinkedService"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20200301preview:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:LinkedService"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20200801:LinkedService"),
		},
	})
	opts = append(opts, aliases)
	var resource LinkedService
	err := ctx.RegisterResource("azure-native:operationalinsights/v20151101preview:LinkedService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLinkedService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceState, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	var resource LinkedService
	err := ctx.ReadResource("azure-native:operationalinsights/v20151101preview:LinkedService", name, id, state, &resource, opts...)
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
	LinkedServiceName *string `pulumi:"linkedServiceName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceId        string  `pulumi:"resourceId"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type LinkedServiceArgs struct {
	LinkedServiceName pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceId        pulumi.StringInput
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
