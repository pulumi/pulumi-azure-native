


package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StaticMember struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput      `pulumi:"etag"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	ResourceId pulumi.StringPtrOutput   `pulumi:"resourceId"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewStaticMember(ctx *pulumi.Context,
	name string, args *StaticMemberArgs, opts ...pulumi.ResourceOption) (*StaticMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkGroupName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkGroupName'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210501preview:StaticMember"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:StaticMember"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:StaticMember"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:StaticMember"),
		},
	})
	opts = append(opts, aliases)
	var resource StaticMember
	err := ctx.RegisterResource("azure-native:network:StaticMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStaticMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticMemberState, opts ...pulumi.ResourceOption) (*StaticMember, error) {
	var resource StaticMember
	err := ctx.ReadResource("azure-native:network:StaticMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type staticMemberState struct {
}

type StaticMemberState struct {
}

func (StaticMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticMemberState)(nil)).Elem()
}

type staticMemberArgs struct {
	NetworkGroupName   string  `pulumi:"networkGroupName"`
	NetworkManagerName string  `pulumi:"networkManagerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	ResourceId         *string `pulumi:"resourceId"`
	StaticMemberName   *string `pulumi:"staticMemberName"`
}


type StaticMemberArgs struct {
	NetworkGroupName   pulumi.StringInput
	NetworkManagerName pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	ResourceId         pulumi.StringPtrInput
	StaticMemberName   pulumi.StringPtrInput
}

func (StaticMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticMemberArgs)(nil)).Elem()
}

type StaticMemberInput interface {
	pulumi.Input

	ToStaticMemberOutput() StaticMemberOutput
	ToStaticMemberOutputWithContext(ctx context.Context) StaticMemberOutput
}

func (*StaticMember) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticMember)(nil)).Elem()
}

func (i *StaticMember) ToStaticMemberOutput() StaticMemberOutput {
	return i.ToStaticMemberOutputWithContext(context.Background())
}

func (i *StaticMember) ToStaticMemberOutputWithContext(ctx context.Context) StaticMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticMemberOutput)
}

type StaticMemberOutput struct{ *pulumi.OutputState }

func (StaticMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticMember)(nil)).Elem()
}

func (o StaticMemberOutput) ToStaticMemberOutput() StaticMemberOutput {
	return o
}

func (o StaticMemberOutput) ToStaticMemberOutputWithContext(ctx context.Context) StaticMemberOutput {
	return o
}

func (o StaticMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o StaticMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StaticMemberOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticMember) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o StaticMemberOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StaticMember) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o StaticMemberOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticMember) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticMemberOutput{})
}
