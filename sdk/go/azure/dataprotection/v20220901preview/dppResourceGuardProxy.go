


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DppResourceGuardProxy struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                  `pulumi:"name"`
	Properties ResourceGuardProxyBaseResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput             `pulumi:"systemData"`
	Type       pulumi.StringOutput                  `pulumi:"type"`
}


func NewDppResourceGuardProxy(ctx *pulumi.Context,
	name string, args *DppResourceGuardProxyArgs, opts ...pulumi.ResourceOption) (*DppResourceGuardProxy, error) {
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
			Type: pulumi.String("azure-native:dataprotection/v20221001preview:DppResourceGuardProxy"),
		},
	})
	opts = append(opts, aliases)
	var resource DppResourceGuardProxy
	err := ctx.RegisterResource("azure-native:dataprotection/v20220901preview:DppResourceGuardProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDppResourceGuardProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DppResourceGuardProxyState, opts ...pulumi.ResourceOption) (*DppResourceGuardProxy, error) {
	var resource DppResourceGuardProxy
	err := ctx.ReadResource("azure-native:dataprotection/v20220901preview:DppResourceGuardProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dppResourceGuardProxyState struct {
}

type DppResourceGuardProxyState struct {
}

func (DppResourceGuardProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dppResourceGuardProxyState)(nil)).Elem()
}

type dppResourceGuardProxyArgs struct {
	Properties             *ResourceGuardProxyBase `pulumi:"properties"`
	ResourceGroupName      string                  `pulumi:"resourceGroupName"`
	ResourceGuardProxyName *string                 `pulumi:"resourceGuardProxyName"`
	VaultName              string                  `pulumi:"vaultName"`
}


type DppResourceGuardProxyArgs struct {
	Properties             ResourceGuardProxyBasePtrInput
	ResourceGroupName      pulumi.StringInput
	ResourceGuardProxyName pulumi.StringPtrInput
	VaultName              pulumi.StringInput
}

func (DppResourceGuardProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dppResourceGuardProxyArgs)(nil)).Elem()
}

type DppResourceGuardProxyInput interface {
	pulumi.Input

	ToDppResourceGuardProxyOutput() DppResourceGuardProxyOutput
	ToDppResourceGuardProxyOutputWithContext(ctx context.Context) DppResourceGuardProxyOutput
}

func (*DppResourceGuardProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**DppResourceGuardProxy)(nil)).Elem()
}

func (i *DppResourceGuardProxy) ToDppResourceGuardProxyOutput() DppResourceGuardProxyOutput {
	return i.ToDppResourceGuardProxyOutputWithContext(context.Background())
}

func (i *DppResourceGuardProxy) ToDppResourceGuardProxyOutputWithContext(ctx context.Context) DppResourceGuardProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DppResourceGuardProxyOutput)
}

type DppResourceGuardProxyOutput struct{ *pulumi.OutputState }

func (DppResourceGuardProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DppResourceGuardProxy)(nil)).Elem()
}

func (o DppResourceGuardProxyOutput) ToDppResourceGuardProxyOutput() DppResourceGuardProxyOutput {
	return o
}

func (o DppResourceGuardProxyOutput) ToDppResourceGuardProxyOutputWithContext(ctx context.Context) DppResourceGuardProxyOutput {
	return o
}

func (o DppResourceGuardProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DppResourceGuardProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DppResourceGuardProxyOutput) Properties() ResourceGuardProxyBaseResponseOutput {
	return o.ApplyT(func(v *DppResourceGuardProxy) ResourceGuardProxyBaseResponseOutput { return v.Properties }).(ResourceGuardProxyBaseResponseOutput)
}

func (o DppResourceGuardProxyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DppResourceGuardProxy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DppResourceGuardProxyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DppResourceGuardProxy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DppResourceGuardProxyOutput{})
}
