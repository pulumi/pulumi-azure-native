


package recoveryservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceGuardProxy struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput               `pulumi:"eTag"`
	Location   pulumi.StringPtrOutput               `pulumi:"location"`
	Name       pulumi.StringOutput                  `pulumi:"name"`
	Properties ResourceGuardProxyBaseResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput               `pulumi:"tags"`
	Type       pulumi.StringOutput                  `pulumi:"type"`
}


func NewResourceGuardProxy(ctx *pulumi.Context,
	name string, args *ResourceGuardProxyArgs, opts ...pulumi.ResourceOption) (*ResourceGuardProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201preview:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220601preview:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220901preview:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230101:ResourceGuardProxy"),
		},
	})
	opts = append(opts, aliases)
	var resource ResourceGuardProxy
	err := ctx.RegisterResource("azure-native:recoveryservices:ResourceGuardProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetResourceGuardProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGuardProxyState, opts ...pulumi.ResourceOption) (*ResourceGuardProxy, error) {
	var resource ResourceGuardProxy
	err := ctx.ReadResource("azure-native:recoveryservices:ResourceGuardProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type resourceGuardProxyState struct {
}

type ResourceGuardProxyState struct {
}

func (ResourceGuardProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGuardProxyState)(nil)).Elem()
}

type resourceGuardProxyArgs struct {
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	ResourceGuardProxyName *string `pulumi:"resourceGuardProxyName"`
	VaultName              string  `pulumi:"vaultName"`
}


type ResourceGuardProxyArgs struct {
	ResourceGroupName      pulumi.StringInput
	ResourceGuardProxyName pulumi.StringPtrInput
	VaultName              pulumi.StringInput
}

func (ResourceGuardProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGuardProxyArgs)(nil)).Elem()
}

type ResourceGuardProxyInput interface {
	pulumi.Input

	ToResourceGuardProxyOutput() ResourceGuardProxyOutput
	ToResourceGuardProxyOutputWithContext(ctx context.Context) ResourceGuardProxyOutput
}

func (*ResourceGuardProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGuardProxy)(nil)).Elem()
}

func (i *ResourceGuardProxy) ToResourceGuardProxyOutput() ResourceGuardProxyOutput {
	return i.ToResourceGuardProxyOutputWithContext(context.Background())
}

func (i *ResourceGuardProxy) ToResourceGuardProxyOutputWithContext(ctx context.Context) ResourceGuardProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardProxyOutput)
}

type ResourceGuardProxyOutput struct{ *pulumi.OutputState }

func (ResourceGuardProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGuardProxy)(nil)).Elem()
}

func (o ResourceGuardProxyOutput) ToResourceGuardProxyOutput() ResourceGuardProxyOutput {
	return o
}

func (o ResourceGuardProxyOutput) ToResourceGuardProxyOutputWithContext(ctx context.Context) ResourceGuardProxyOutput {
	return o
}

func (o ResourceGuardProxyOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuardProxy) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o ResourceGuardProxyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuardProxy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResourceGuardProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGuardProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceGuardProxyOutput) Properties() ResourceGuardProxyBaseResponseOutput {
	return o.ApplyT(func(v *ResourceGuardProxy) ResourceGuardProxyBaseResponseOutput { return v.Properties }).(ResourceGuardProxyBaseResponseOutput)
}

func (o ResourceGuardProxyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceGuardProxy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ResourceGuardProxyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGuardProxy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceGuardProxyOutput{})
}
