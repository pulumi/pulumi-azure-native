


package v20160129

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkspaceCollection struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput    `pulumi:"location"`
	Name       pulumi.StringPtrOutput    `pulumi:"name"`
	Properties pulumi.AnyOutput          `pulumi:"properties"`
	Sku        AzureSkuResponsePtrOutput `pulumi:"sku"`
	Tags       pulumi.StringMapOutput    `pulumi:"tags"`
	Type       pulumi.StringPtrOutput    `pulumi:"type"`
}


func NewWorkspaceCollection(ctx *pulumi.Context,
	name string, args *WorkspaceCollectionArgs, opts ...pulumi.ResourceOption) (*WorkspaceCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:powerbi/v20160129:WorkspaceCollection"),
		},
		{
			Type: pulumi.String("azure-native:powerbi:WorkspaceCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:powerbi:WorkspaceCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkspaceCollection
	err := ctx.RegisterResource("azure-native:powerbi/v20160129:WorkspaceCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspaceCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceCollectionState, opts ...pulumi.ResourceOption) (*WorkspaceCollection, error) {
	var resource WorkspaceCollection
	err := ctx.ReadResource("azure-native:powerbi/v20160129:WorkspaceCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workspaceCollectionState struct {
}

type WorkspaceCollectionState struct {
}

func (WorkspaceCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceCollectionState)(nil)).Elem()
}

type workspaceCollectionArgs struct {
	Location                *string           `pulumi:"location"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
	Sku                     *AzureSku         `pulumi:"sku"`
	Tags                    map[string]string `pulumi:"tags"`
	WorkspaceCollectionName *string           `pulumi:"workspaceCollectionName"`
}


type WorkspaceCollectionArgs struct {
	Location                pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	Sku                     AzureSkuPtrInput
	Tags                    pulumi.StringMapInput
	WorkspaceCollectionName pulumi.StringPtrInput
}

func (WorkspaceCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceCollectionArgs)(nil)).Elem()
}

type WorkspaceCollectionInput interface {
	pulumi.Input

	ToWorkspaceCollectionOutput() WorkspaceCollectionOutput
	ToWorkspaceCollectionOutputWithContext(ctx context.Context) WorkspaceCollectionOutput
}

func (*WorkspaceCollection) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCollection)(nil))
}

func (i *WorkspaceCollection) ToWorkspaceCollectionOutput() WorkspaceCollectionOutput {
	return i.ToWorkspaceCollectionOutputWithContext(context.Background())
}

func (i *WorkspaceCollection) ToWorkspaceCollectionOutputWithContext(ctx context.Context) WorkspaceCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCollectionOutput)
}

type WorkspaceCollectionOutput struct{ *pulumi.OutputState }

func (WorkspaceCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCollection)(nil))
}

func (o WorkspaceCollectionOutput) ToWorkspaceCollectionOutput() WorkspaceCollectionOutput {
	return o
}

func (o WorkspaceCollectionOutput) ToWorkspaceCollectionOutputWithContext(ctx context.Context) WorkspaceCollectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkspaceCollectionOutput{})
}
