


package v20180301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type LinkedServer struct {
	pulumi.CustomResourceState

	LinkedRedisCacheId       pulumi.StringOutput `pulumi:"linkedRedisCacheId"`
	LinkedRedisCacheLocation pulumi.StringOutput `pulumi:"linkedRedisCacheLocation"`
	Name                     pulumi.StringOutput `pulumi:"name"`
	ProvisioningState        pulumi.StringOutput `pulumi:"provisioningState"`
	ServerRole               pulumi.StringOutput `pulumi:"serverRole"`
	Type                     pulumi.StringOutput `pulumi:"type"`
}


func NewLinkedServer(ctx *pulumi.Context,
	name string, args *LinkedServerArgs, opts ...pulumi.ResourceOption) (*LinkedServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LinkedRedisCacheId == nil {
		return nil, errors.New("invalid value for required argument 'LinkedRedisCacheId'")
	}
	if args.LinkedRedisCacheLocation == nil {
		return nil, errors.New("invalid value for required argument 'LinkedRedisCacheLocation'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerRole == nil {
		return nil, errors.New("invalid value for required argument 'ServerRole'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cache:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20170201:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20171001:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20190701:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20200601:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20201201:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210601:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220501:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220601:LinkedServer"),
		},
	})
	opts = append(opts, aliases)
	var resource LinkedServer
	err := ctx.RegisterResource("azure-native:cache/v20180301:LinkedServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLinkedServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServerState, opts ...pulumi.ResourceOption) (*LinkedServer, error) {
	var resource LinkedServer
	err := ctx.ReadResource("azure-native:cache/v20180301:LinkedServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type linkedServerState struct {
}

type LinkedServerState struct {
}

func (LinkedServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServerState)(nil)).Elem()
}

type linkedServerArgs struct {
	LinkedRedisCacheId       string          `pulumi:"linkedRedisCacheId"`
	LinkedRedisCacheLocation string          `pulumi:"linkedRedisCacheLocation"`
	LinkedServerName         *string         `pulumi:"linkedServerName"`
	Name                     string          `pulumi:"name"`
	ResourceGroupName        string          `pulumi:"resourceGroupName"`
	ServerRole               ReplicationRole `pulumi:"serverRole"`
}


type LinkedServerArgs struct {
	LinkedRedisCacheId       pulumi.StringInput
	LinkedRedisCacheLocation pulumi.StringInput
	LinkedServerName         pulumi.StringPtrInput
	Name                     pulumi.StringInput
	ResourceGroupName        pulumi.StringInput
	ServerRole               ReplicationRoleInput
}

func (LinkedServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServerArgs)(nil)).Elem()
}

type LinkedServerInput interface {
	pulumi.Input

	ToLinkedServerOutput() LinkedServerOutput
	ToLinkedServerOutputWithContext(ctx context.Context) LinkedServerOutput
}

func (*LinkedServer) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServer)(nil)).Elem()
}

func (i *LinkedServer) ToLinkedServerOutput() LinkedServerOutput {
	return i.ToLinkedServerOutputWithContext(context.Background())
}

func (i *LinkedServer) ToLinkedServerOutputWithContext(ctx context.Context) LinkedServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServerOutput)
}

type LinkedServerOutput struct{ *pulumi.OutputState }

func (LinkedServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServer)(nil)).Elem()
}

func (o LinkedServerOutput) ToLinkedServerOutput() LinkedServerOutput {
	return o
}

func (o LinkedServerOutput) ToLinkedServerOutputWithContext(ctx context.Context) LinkedServerOutput {
	return o
}

func (o LinkedServerOutput) LinkedRedisCacheId() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.LinkedRedisCacheId }).(pulumi.StringOutput)
}

func (o LinkedServerOutput) LinkedRedisCacheLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.LinkedRedisCacheLocation }).(pulumi.StringOutput)
}

func (o LinkedServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LinkedServerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LinkedServerOutput) ServerRole() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.ServerRole }).(pulumi.StringOutput)
}

func (o LinkedServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServerOutput{})
}
