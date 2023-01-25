


package v20220904

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SyncIdentityProvider struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput      `pulumi:"name"`
	Resources  pulumi.StringPtrOutput   `pulumi:"resources"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewSyncIdentityProvider(ctx *pulumi.Context,
	name string, args *SyncIdentityProviderArgs, opts ...pulumi.ResourceOption) (*SyncIdentityProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	var resource SyncIdentityProvider
	err := ctx.RegisterResource("azure-native:redhatopenshift/v20220904:SyncIdentityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSyncIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncIdentityProviderState, opts ...pulumi.ResourceOption) (*SyncIdentityProvider, error) {
	var resource SyncIdentityProvider
	err := ctx.ReadResource("azure-native:redhatopenshift/v20220904:SyncIdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type syncIdentityProviderState struct {
}

type SyncIdentityProviderState struct {
}

func (SyncIdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncIdentityProviderState)(nil)).Elem()
}

type syncIdentityProviderArgs struct {
	ChildResourceName *string `pulumi:"childResourceName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceName      string  `pulumi:"resourceName"`
	Resources         *string `pulumi:"resources"`
}


type SyncIdentityProviderArgs struct {
	ChildResourceName pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
	Resources         pulumi.StringPtrInput
}

func (SyncIdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncIdentityProviderArgs)(nil)).Elem()
}

type SyncIdentityProviderInput interface {
	pulumi.Input

	ToSyncIdentityProviderOutput() SyncIdentityProviderOutput
	ToSyncIdentityProviderOutputWithContext(ctx context.Context) SyncIdentityProviderOutput
}

func (*SyncIdentityProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncIdentityProvider)(nil)).Elem()
}

func (i *SyncIdentityProvider) ToSyncIdentityProviderOutput() SyncIdentityProviderOutput {
	return i.ToSyncIdentityProviderOutputWithContext(context.Background())
}

func (i *SyncIdentityProvider) ToSyncIdentityProviderOutputWithContext(ctx context.Context) SyncIdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncIdentityProviderOutput)
}

type SyncIdentityProviderOutput struct{ *pulumi.OutputState }

func (SyncIdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncIdentityProvider)(nil)).Elem()
}

func (o SyncIdentityProviderOutput) ToSyncIdentityProviderOutput() SyncIdentityProviderOutput {
	return o
}

func (o SyncIdentityProviderOutput) ToSyncIdentityProviderOutputWithContext(ctx context.Context) SyncIdentityProviderOutput {
	return o
}

func (o SyncIdentityProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncIdentityProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SyncIdentityProviderOutput) Resources() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncIdentityProvider) pulumi.StringPtrOutput { return v.Resources }).(pulumi.StringPtrOutput)
}

func (o SyncIdentityProviderOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SyncIdentityProvider) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SyncIdentityProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncIdentityProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SyncIdentityProviderOutput{})
}
