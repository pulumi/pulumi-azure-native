


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerTrustGroup struct {
	pulumi.CustomResourceState

	GroupMembers ServerInfoResponseArrayOutput `pulumi:"groupMembers"`
	Name         pulumi.StringOutput           `pulumi:"name"`
	TrustScopes  pulumi.StringArrayOutput      `pulumi:"trustScopes"`
	Type         pulumi.StringOutput           `pulumi:"type"`
}


func NewServerTrustGroup(ctx *pulumi.Context,
	name string, args *ServerTrustGroupArgs, opts ...pulumi.ResourceOption) (*ServerTrustGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupMembers == nil {
		return nil, errors.New("invalid value for required argument 'GroupMembers'")
	}
	if args.LocationName == nil {
		return nil, errors.New("invalid value for required argument 'LocationName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TrustScopes == nil {
		return nil, errors.New("invalid value for required argument 'TrustScopes'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ServerTrustGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerTrustGroup
	err := ctx.RegisterResource("azure-native:sql/v20211101preview:ServerTrustGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerTrustGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerTrustGroupState, opts ...pulumi.ResourceOption) (*ServerTrustGroup, error) {
	var resource ServerTrustGroup
	err := ctx.ReadResource("azure-native:sql/v20211101preview:ServerTrustGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverTrustGroupState struct {
}

type ServerTrustGroupState struct {
}

func (ServerTrustGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverTrustGroupState)(nil)).Elem()
}

type serverTrustGroupArgs struct {
	GroupMembers         []ServerInfo `pulumi:"groupMembers"`
	LocationName         string       `pulumi:"locationName"`
	ResourceGroupName    string       `pulumi:"resourceGroupName"`
	ServerTrustGroupName *string      `pulumi:"serverTrustGroupName"`
	TrustScopes          []string     `pulumi:"trustScopes"`
}


type ServerTrustGroupArgs struct {
	GroupMembers         ServerInfoArrayInput
	LocationName         pulumi.StringInput
	ResourceGroupName    pulumi.StringInput
	ServerTrustGroupName pulumi.StringPtrInput
	TrustScopes          pulumi.StringArrayInput
}

func (ServerTrustGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverTrustGroupArgs)(nil)).Elem()
}

type ServerTrustGroupInput interface {
	pulumi.Input

	ToServerTrustGroupOutput() ServerTrustGroupOutput
	ToServerTrustGroupOutputWithContext(ctx context.Context) ServerTrustGroupOutput
}

func (*ServerTrustGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerTrustGroup)(nil)).Elem()
}

func (i *ServerTrustGroup) ToServerTrustGroupOutput() ServerTrustGroupOutput {
	return i.ToServerTrustGroupOutputWithContext(context.Background())
}

func (i *ServerTrustGroup) ToServerTrustGroupOutputWithContext(ctx context.Context) ServerTrustGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerTrustGroupOutput)
}

type ServerTrustGroupOutput struct{ *pulumi.OutputState }

func (ServerTrustGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerTrustGroup)(nil)).Elem()
}

func (o ServerTrustGroupOutput) ToServerTrustGroupOutput() ServerTrustGroupOutput {
	return o
}

func (o ServerTrustGroupOutput) ToServerTrustGroupOutputWithContext(ctx context.Context) ServerTrustGroupOutput {
	return o
}

func (o ServerTrustGroupOutput) GroupMembers() ServerInfoResponseArrayOutput {
	return o.ApplyT(func(v *ServerTrustGroup) ServerInfoResponseArrayOutput { return v.GroupMembers }).(ServerInfoResponseArrayOutput)
}

func (o ServerTrustGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTrustGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerTrustGroupOutput) TrustScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerTrustGroup) pulumi.StringArrayOutput { return v.TrustScopes }).(pulumi.StringArrayOutput)
}

func (o ServerTrustGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTrustGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerTrustGroupOutput{})
}
