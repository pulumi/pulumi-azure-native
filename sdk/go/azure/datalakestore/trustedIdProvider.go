


package datalakestore

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TrustedIdProvider struct {
	pulumi.CustomResourceState

	IdProvider pulumi.StringOutput `pulumi:"idProvider"`
	Name       pulumi.StringOutput `pulumi:"name"`
	Type       pulumi.StringOutput `pulumi:"type"`
}


func NewTrustedIdProvider(ctx *pulumi.Context,
	name string, args *TrustedIdProviderArgs, opts ...pulumi.ResourceOption) (*TrustedIdProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.IdProvider == nil {
		return nil, errors.New("invalid value for required argument 'IdProvider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datalakestore/v20161101:TrustedIdProvider"),
		},
	})
	opts = append(opts, aliases)
	var resource TrustedIdProvider
	err := ctx.RegisterResource("azure-native:datalakestore:TrustedIdProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTrustedIdProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrustedIdProviderState, opts ...pulumi.ResourceOption) (*TrustedIdProvider, error) {
	var resource TrustedIdProvider
	err := ctx.ReadResource("azure-native:datalakestore:TrustedIdProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type trustedIdProviderState struct {
}

type TrustedIdProviderState struct {
}

func (TrustedIdProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*trustedIdProviderState)(nil)).Elem()
}

type trustedIdProviderArgs struct {
	AccountName           string  `pulumi:"accountName"`
	IdProvider            string  `pulumi:"idProvider"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	TrustedIdProviderName *string `pulumi:"trustedIdProviderName"`
}


type TrustedIdProviderArgs struct {
	AccountName           pulumi.StringInput
	IdProvider            pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	TrustedIdProviderName pulumi.StringPtrInput
}

func (TrustedIdProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trustedIdProviderArgs)(nil)).Elem()
}

type TrustedIdProviderInput interface {
	pulumi.Input

	ToTrustedIdProviderOutput() TrustedIdProviderOutput
	ToTrustedIdProviderOutputWithContext(ctx context.Context) TrustedIdProviderOutput
}

func (*TrustedIdProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustedIdProvider)(nil)).Elem()
}

func (i *TrustedIdProvider) ToTrustedIdProviderOutput() TrustedIdProviderOutput {
	return i.ToTrustedIdProviderOutputWithContext(context.Background())
}

func (i *TrustedIdProvider) ToTrustedIdProviderOutputWithContext(ctx context.Context) TrustedIdProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedIdProviderOutput)
}

type TrustedIdProviderOutput struct{ *pulumi.OutputState }

func (TrustedIdProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustedIdProvider)(nil)).Elem()
}

func (o TrustedIdProviderOutput) ToTrustedIdProviderOutput() TrustedIdProviderOutput {
	return o
}

func (o TrustedIdProviderOutput) ToTrustedIdProviderOutputWithContext(ctx context.Context) TrustedIdProviderOutput {
	return o
}

func (o TrustedIdProviderOutput) IdProvider() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustedIdProvider) pulumi.StringOutput { return v.IdProvider }).(pulumi.StringOutput)
}

func (o TrustedIdProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustedIdProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TrustedIdProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustedIdProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TrustedIdProviderOutput{})
}
