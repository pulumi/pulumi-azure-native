


package v20200401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IpGroup struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput            `pulumi:"etag"`
	Firewalls         SubResourceResponseArrayOutput `pulumi:"firewalls"`
	IpAddresses       pulumi.StringArrayOutput       `pulumi:"ipAddresses"`
	Location          pulumi.StringPtrOutput         `pulumi:"location"`
	Name              pulumi.StringOutput            `pulumi:"name"`
	ProvisioningState pulumi.StringOutput            `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput         `pulumi:"tags"`
	Type              pulumi.StringOutput            `pulumi:"type"`
}


func NewIpGroup(ctx *pulumi.Context,
	name string, args *IpGroupArgs, opts ...pulumi.ResourceOption) (*IpGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:IpGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:IpGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource IpGroup
	err := ctx.RegisterResource("azure-native:network/v20200401:IpGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIpGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpGroupState, opts ...pulumi.ResourceOption) (*IpGroup, error) {
	var resource IpGroup
	err := ctx.ReadResource("azure-native:network/v20200401:IpGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ipGroupState struct {
}

type IpGroupState struct {
}

func (IpGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipGroupState)(nil)).Elem()
}

type ipGroupArgs struct {
	Id                *string           `pulumi:"id"`
	IpAddresses       []string          `pulumi:"ipAddresses"`
	IpGroupsName      *string           `pulumi:"ipGroupsName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type IpGroupArgs struct {
	Id                pulumi.StringPtrInput
	IpAddresses       pulumi.StringArrayInput
	IpGroupsName      pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (IpGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipGroupArgs)(nil)).Elem()
}

type IpGroupInput interface {
	pulumi.Input

	ToIpGroupOutput() IpGroupOutput
	ToIpGroupOutputWithContext(ctx context.Context) IpGroupOutput
}

func (*IpGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*IpGroup)(nil))
}

func (i *IpGroup) ToIpGroupOutput() IpGroupOutput {
	return i.ToIpGroupOutputWithContext(context.Background())
}

func (i *IpGroup) ToIpGroupOutputWithContext(ctx context.Context) IpGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpGroupOutput)
}

type IpGroupOutput struct{ *pulumi.OutputState }

func (IpGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpGroup)(nil))
}

func (o IpGroupOutput) ToIpGroupOutput() IpGroupOutput {
	return o
}

func (o IpGroupOutput) ToIpGroupOutputWithContext(ctx context.Context) IpGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IpGroupOutput{})
}
