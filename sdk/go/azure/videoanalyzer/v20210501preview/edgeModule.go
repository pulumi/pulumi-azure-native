


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EdgeModule struct {
	pulumi.CustomResourceState

	EdgeModuleId pulumi.StringOutput      `pulumi:"edgeModuleId"`
	Name         pulumi.StringOutput      `pulumi:"name"`
	SystemData   SystemDataResponseOutput `pulumi:"systemData"`
	Type         pulumi.StringOutput      `pulumi:"type"`
}


func NewEdgeModule(ctx *pulumi.Context,
	name string, args *EdgeModuleArgs, opts ...pulumi.ResourceOption) (*EdgeModule, error) {
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
			Type: pulumi.String("azure-nextgen:videoanalyzer/v20210501preview:EdgeModule"),
		},
		{
			Type: pulumi.String("azure-native:videoanalyzer:EdgeModule"),
		},
		{
			Type: pulumi.String("azure-nextgen:videoanalyzer:EdgeModule"),
		},
		{
			Type: pulumi.String("azure-native:videoanalyzer/v20211101preview:EdgeModule"),
		},
		{
			Type: pulumi.String("azure-nextgen:videoanalyzer/v20211101preview:EdgeModule"),
		},
	})
	opts = append(opts, aliases)
	var resource EdgeModule
	err := ctx.RegisterResource("azure-native:videoanalyzer/v20210501preview:EdgeModule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEdgeModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeModuleState, opts ...pulumi.ResourceOption) (*EdgeModule, error) {
	var resource EdgeModule
	err := ctx.ReadResource("azure-native:videoanalyzer/v20210501preview:EdgeModule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type edgeModuleState struct {
}

type EdgeModuleState struct {
}

func (EdgeModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeModuleState)(nil)).Elem()
}

type edgeModuleArgs struct {
	AccountName       string  `pulumi:"accountName"`
	EdgeModuleName    *string `pulumi:"edgeModuleName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type EdgeModuleArgs struct {
	AccountName       pulumi.StringInput
	EdgeModuleName    pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (EdgeModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeModuleArgs)(nil)).Elem()
}

type EdgeModuleInput interface {
	pulumi.Input

	ToEdgeModuleOutput() EdgeModuleOutput
	ToEdgeModuleOutputWithContext(ctx context.Context) EdgeModuleOutput
}

func (*EdgeModule) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeModule)(nil))
}

func (i *EdgeModule) ToEdgeModuleOutput() EdgeModuleOutput {
	return i.ToEdgeModuleOutputWithContext(context.Background())
}

func (i *EdgeModule) ToEdgeModuleOutputWithContext(ctx context.Context) EdgeModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeModuleOutput)
}

type EdgeModuleOutput struct{ *pulumi.OutputState }

func (EdgeModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeModule)(nil))
}

func (o EdgeModuleOutput) ToEdgeModuleOutput() EdgeModuleOutput {
	return o
}

func (o EdgeModuleOutput) ToEdgeModuleOutputWithContext(ctx context.Context) EdgeModuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EdgeModuleOutput{})
}
