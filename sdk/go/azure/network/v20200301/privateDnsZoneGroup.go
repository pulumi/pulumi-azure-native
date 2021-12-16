


package v20200301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateDnsZoneGroup struct {
	pulumi.CustomResourceState

	Etag                  pulumi.StringOutput                     `pulumi:"etag"`
	Name                  pulumi.StringPtrOutput                  `pulumi:"name"`
	PrivateDnsZoneConfigs PrivateDnsZoneConfigResponseArrayOutput `pulumi:"privateDnsZoneConfigs"`
	ProvisioningState     pulumi.StringOutput                     `pulumi:"provisioningState"`
}


func NewPrivateDnsZoneGroup(ctx *pulumi.Context,
	name string, args *PrivateDnsZoneGroupArgs, opts ...pulumi.ResourceOption) (*PrivateDnsZoneGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateEndpointName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:PrivateDnsZoneGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateDnsZoneGroup
	err := ctx.RegisterResource("azure-native:network/v20200301:PrivateDnsZoneGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateDnsZoneGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateDnsZoneGroupState, opts ...pulumi.ResourceOption) (*PrivateDnsZoneGroup, error) {
	var resource PrivateDnsZoneGroup
	err := ctx.ReadResource("azure-native:network/v20200301:PrivateDnsZoneGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateDnsZoneGroupState struct {
}

type PrivateDnsZoneGroupState struct {
}

func (PrivateDnsZoneGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDnsZoneGroupState)(nil)).Elem()
}

type privateDnsZoneGroupArgs struct {
	Id                      *string                `pulumi:"id"`
	Name                    *string                `pulumi:"name"`
	PrivateDnsZoneConfigs   []PrivateDnsZoneConfig `pulumi:"privateDnsZoneConfigs"`
	PrivateDnsZoneGroupName *string                `pulumi:"privateDnsZoneGroupName"`
	PrivateEndpointName     string                 `pulumi:"privateEndpointName"`
	ResourceGroupName       string                 `pulumi:"resourceGroupName"`
}


type PrivateDnsZoneGroupArgs struct {
	Id                      pulumi.StringPtrInput
	Name                    pulumi.StringPtrInput
	PrivateDnsZoneConfigs   PrivateDnsZoneConfigArrayInput
	PrivateDnsZoneGroupName pulumi.StringPtrInput
	PrivateEndpointName     pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
}

func (PrivateDnsZoneGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDnsZoneGroupArgs)(nil)).Elem()
}

type PrivateDnsZoneGroupInput interface {
	pulumi.Input

	ToPrivateDnsZoneGroupOutput() PrivateDnsZoneGroupOutput
	ToPrivateDnsZoneGroupOutputWithContext(ctx context.Context) PrivateDnsZoneGroupOutput
}

func (*PrivateDnsZoneGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDnsZoneGroup)(nil)).Elem()
}

func (i *PrivateDnsZoneGroup) ToPrivateDnsZoneGroupOutput() PrivateDnsZoneGroupOutput {
	return i.ToPrivateDnsZoneGroupOutputWithContext(context.Background())
}

func (i *PrivateDnsZoneGroup) ToPrivateDnsZoneGroupOutputWithContext(ctx context.Context) PrivateDnsZoneGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDnsZoneGroupOutput)
}

type PrivateDnsZoneGroupOutput struct{ *pulumi.OutputState }

func (PrivateDnsZoneGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDnsZoneGroup)(nil)).Elem()
}

func (o PrivateDnsZoneGroupOutput) ToPrivateDnsZoneGroupOutput() PrivateDnsZoneGroupOutput {
	return o
}

func (o PrivateDnsZoneGroupOutput) ToPrivateDnsZoneGroupOutputWithContext(ctx context.Context) PrivateDnsZoneGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateDnsZoneGroupOutput{})
}
