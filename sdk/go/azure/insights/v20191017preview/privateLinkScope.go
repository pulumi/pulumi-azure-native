


package v20191017preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkScope struct {
	pulumi.CustomResourceState

	Location                   pulumi.StringOutput                          `pulumi:"location"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
}


func NewPrivateLinkScope(ctx *pulumi.Context,
	name string, args *PrivateLinkScopeArgs, opts ...pulumi.ResourceOption) (*PrivateLinkScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210701preview:PrivateLinkScope"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkScope
	err := ctx.RegisterResource("azure-native:insights/v20191017preview:PrivateLinkScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkScopeState, opts ...pulumi.ResourceOption) (*PrivateLinkScope, error) {
	var resource PrivateLinkScope
	err := ctx.ReadResource("azure-native:insights/v20191017preview:PrivateLinkScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkScopeState struct {
}

type PrivateLinkScopeState struct {
}

func (PrivateLinkScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkScopeState)(nil)).Elem()
}

type privateLinkScopeArgs struct {
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ScopeName         *string           `pulumi:"scopeName"`
	Tags              map[string]string `pulumi:"tags"`
}


type PrivateLinkScopeArgs struct {
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ScopeName         pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (PrivateLinkScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkScopeArgs)(nil)).Elem()
}

type PrivateLinkScopeInput interface {
	pulumi.Input

	ToPrivateLinkScopeOutput() PrivateLinkScopeOutput
	ToPrivateLinkScopeOutputWithContext(ctx context.Context) PrivateLinkScopeOutput
}

func (*PrivateLinkScope) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkScope)(nil)).Elem()
}

func (i *PrivateLinkScope) ToPrivateLinkScopeOutput() PrivateLinkScopeOutput {
	return i.ToPrivateLinkScopeOutputWithContext(context.Background())
}

func (i *PrivateLinkScope) ToPrivateLinkScopeOutputWithContext(ctx context.Context) PrivateLinkScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkScopeOutput)
}

type PrivateLinkScopeOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkScope)(nil)).Elem()
}

func (o PrivateLinkScopeOutput) ToPrivateLinkScopeOutput() PrivateLinkScopeOutput {
	return o
}

func (o PrivateLinkScopeOutput) ToPrivateLinkScopeOutputWithContext(ctx context.Context) PrivateLinkScopeOutput {
	return o
}

func (o PrivateLinkScopeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrivateLinkScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateLinkScopeOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *PrivateLinkScope) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o PrivateLinkScopeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateLinkScopeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateLinkScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkScopeOutput{})
}
