


package hybridcompute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkScopedResource struct {
	pulumi.CustomResourceState

	LinkedResourceId  pulumi.StringPtrOutput `pulumi:"linkedResourceId"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewPrivateLinkScopedResource(ctx *pulumi.Context,
	name string, args *PrivateLinkScopedResourceArgs, opts ...pulumi.ResourceOption) (*PrivateLinkScopedResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScopeName == nil {
		return nil, errors.New("invalid value for required argument 'ScopeName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200815preview:PrivateLinkScopedResource"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkScopedResource
	err := ctx.RegisterResource("azure-native:hybridcompute:PrivateLinkScopedResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkScopedResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkScopedResourceState, opts ...pulumi.ResourceOption) (*PrivateLinkScopedResource, error) {
	var resource PrivateLinkScopedResource
	err := ctx.ReadResource("azure-native:hybridcompute:PrivateLinkScopedResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkScopedResourceState struct {
}

type PrivateLinkScopedResourceState struct {
}

func (PrivateLinkScopedResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkScopedResourceState)(nil)).Elem()
}

type privateLinkScopedResourceArgs struct {
	LinkedResourceId  *string `pulumi:"linkedResourceId"`
	Name              *string `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ScopeName         string  `pulumi:"scopeName"`
}


type PrivateLinkScopedResourceArgs struct {
	LinkedResourceId  pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ScopeName         pulumi.StringInput
}

func (PrivateLinkScopedResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkScopedResourceArgs)(nil)).Elem()
}

type PrivateLinkScopedResourceInput interface {
	pulumi.Input

	ToPrivateLinkScopedResourceOutput() PrivateLinkScopedResourceOutput
	ToPrivateLinkScopedResourceOutputWithContext(ctx context.Context) PrivateLinkScopedResourceOutput
}

func (*PrivateLinkScopedResource) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkScopedResource)(nil))
}

func (i *PrivateLinkScopedResource) ToPrivateLinkScopedResourceOutput() PrivateLinkScopedResourceOutput {
	return i.ToPrivateLinkScopedResourceOutputWithContext(context.Background())
}

func (i *PrivateLinkScopedResource) ToPrivateLinkScopedResourceOutputWithContext(ctx context.Context) PrivateLinkScopedResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkScopedResourceOutput)
}

type PrivateLinkScopedResourceOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkScopedResource)(nil))
}

func (o PrivateLinkScopedResourceOutput) ToPrivateLinkScopedResourceOutput() PrivateLinkScopedResourceOutput {
	return o
}

func (o PrivateLinkScopedResourceOutput) ToPrivateLinkScopedResourceOutputWithContext(ctx context.Context) PrivateLinkScopedResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkScopedResourceOutput{})
}
