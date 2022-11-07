


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkGroup struct {
	pulumi.CustomResourceState

	Description       pulumi.StringPtrOutput   `pulumi:"description"`
	Etag              pulumi.StringOutput      `pulumi:"etag"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewNetworkGroup(ctx *pulumi.Context,
	name string, args *NetworkGroupArgs, opts ...pulumi.ResourceOption) (*NetworkGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:NetworkGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkGroup
	err := ctx.RegisterResource("azure-native:network/v20220501:NetworkGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkGroupState, opts ...pulumi.ResourceOption) (*NetworkGroup, error) {
	var resource NetworkGroup
	err := ctx.ReadResource("azure-native:network/v20220501:NetworkGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkGroupState struct {
}

type NetworkGroupState struct {
}

func (NetworkGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkGroupState)(nil)).Elem()
}

type networkGroupArgs struct {
	Description        *string `pulumi:"description"`
	NetworkGroupName   *string `pulumi:"networkGroupName"`
	NetworkManagerName string  `pulumi:"networkManagerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type NetworkGroupArgs struct {
	Description        pulumi.StringPtrInput
	NetworkGroupName   pulumi.StringPtrInput
	NetworkManagerName pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
}

func (NetworkGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkGroupArgs)(nil)).Elem()
}

type NetworkGroupInput interface {
	pulumi.Input

	ToNetworkGroupOutput() NetworkGroupOutput
	ToNetworkGroupOutputWithContext(ctx context.Context) NetworkGroupOutput
}

func (*NetworkGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkGroup)(nil)).Elem()
}

func (i *NetworkGroup) ToNetworkGroupOutput() NetworkGroupOutput {
	return i.ToNetworkGroupOutputWithContext(context.Background())
}

func (i *NetworkGroup) ToNetworkGroupOutputWithContext(ctx context.Context) NetworkGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkGroupOutput)
}

type NetworkGroupOutput struct{ *pulumi.OutputState }

func (NetworkGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkGroup)(nil)).Elem()
}

func (o NetworkGroupOutput) ToNetworkGroupOutput() NetworkGroupOutput {
	return o
}

func (o NetworkGroupOutput) ToNetworkGroupOutputWithContext(ctx context.Context) NetworkGroupOutput {
	return o
}

func (o NetworkGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o NetworkGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NetworkGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o NetworkGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkGroupOutput{})
}
