


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementGroupNetworkManagerConnection struct {
	pulumi.CustomResourceState

	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	Etag             pulumi.StringOutput      `pulumi:"etag"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	NetworkManagerId pulumi.StringPtrOutput   `pulumi:"networkManagerId"`
	SystemData       SystemDataResponseOutput `pulumi:"systemData"`
	Type             pulumi.StringOutput      `pulumi:"type"`
}


func NewManagementGroupNetworkManagerConnection(ctx *pulumi.Context,
	name string, args *ManagementGroupNetworkManagerConnectionArgs, opts ...pulumi.ResourceOption) (*ManagementGroupNetworkManagerConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20220101:ManagementGroupNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:ManagementGroupNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ManagementGroupNetworkManagerConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementGroupNetworkManagerConnection
	err := ctx.RegisterResource("azure-native:network/v20220201preview:ManagementGroupNetworkManagerConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementGroupNetworkManagerConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementGroupNetworkManagerConnectionState, opts ...pulumi.ResourceOption) (*ManagementGroupNetworkManagerConnection, error) {
	var resource ManagementGroupNetworkManagerConnection
	err := ctx.ReadResource("azure-native:network/v20220201preview:ManagementGroupNetworkManagerConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managementGroupNetworkManagerConnectionState struct {
}

type ManagementGroupNetworkManagerConnectionState struct {
}

func (ManagementGroupNetworkManagerConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupNetworkManagerConnectionState)(nil)).Elem()
}

type managementGroupNetworkManagerConnectionArgs struct {
	Description                  *string `pulumi:"description"`
	ManagementGroupId            string  `pulumi:"managementGroupId"`
	NetworkManagerConnectionName *string `pulumi:"networkManagerConnectionName"`
	NetworkManagerId             *string `pulumi:"networkManagerId"`
}


type ManagementGroupNetworkManagerConnectionArgs struct {
	Description                  pulumi.StringPtrInput
	ManagementGroupId            pulumi.StringInput
	NetworkManagerConnectionName pulumi.StringPtrInput
	NetworkManagerId             pulumi.StringPtrInput
}

func (ManagementGroupNetworkManagerConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupNetworkManagerConnectionArgs)(nil)).Elem()
}

type ManagementGroupNetworkManagerConnectionInput interface {
	pulumi.Input

	ToManagementGroupNetworkManagerConnectionOutput() ManagementGroupNetworkManagerConnectionOutput
	ToManagementGroupNetworkManagerConnectionOutputWithContext(ctx context.Context) ManagementGroupNetworkManagerConnectionOutput
}

func (*ManagementGroupNetworkManagerConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupNetworkManagerConnection)(nil)).Elem()
}

func (i *ManagementGroupNetworkManagerConnection) ToManagementGroupNetworkManagerConnectionOutput() ManagementGroupNetworkManagerConnectionOutput {
	return i.ToManagementGroupNetworkManagerConnectionOutputWithContext(context.Background())
}

func (i *ManagementGroupNetworkManagerConnection) ToManagementGroupNetworkManagerConnectionOutputWithContext(ctx context.Context) ManagementGroupNetworkManagerConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupNetworkManagerConnectionOutput)
}

type ManagementGroupNetworkManagerConnectionOutput struct{ *pulumi.OutputState }

func (ManagementGroupNetworkManagerConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupNetworkManagerConnection)(nil)).Elem()
}

func (o ManagementGroupNetworkManagerConnectionOutput) ToManagementGroupNetworkManagerConnectionOutput() ManagementGroupNetworkManagerConnectionOutput {
	return o
}

func (o ManagementGroupNetworkManagerConnectionOutput) ToManagementGroupNetworkManagerConnectionOutputWithContext(ctx context.Context) ManagementGroupNetworkManagerConnectionOutput {
	return o
}

func (o ManagementGroupNetworkManagerConnectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupNetworkManagerConnection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupNetworkManagerConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupNetworkManagerConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ManagementGroupNetworkManagerConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupNetworkManagerConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagementGroupNetworkManagerConnectionOutput) NetworkManagerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupNetworkManagerConnection) pulumi.StringPtrOutput { return v.NetworkManagerId }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupNetworkManagerConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagementGroupNetworkManagerConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ManagementGroupNetworkManagerConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupNetworkManagerConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementGroupNetworkManagerConnectionOutput{})
}
