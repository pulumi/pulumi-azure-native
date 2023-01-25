


package v20211001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProtectionIntent struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput `pulumi:"eTag"`
	Location   pulumi.StringPtrOutput `pulumi:"location"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Properties pulumi.AnyOutput       `pulumi:"properties"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewProtectionIntent(ctx *pulumi.Context,
	name string, args *ProtectionIntentArgs, opts ...pulumi.ResourceOption) (*ProtectionIntent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20170701:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201preview:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220601preview:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220901preview:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230101:ProtectionIntent"),
		},
	})
	opts = append(opts, aliases)
	var resource ProtectionIntent
	err := ctx.RegisterResource("azure-native:recoveryservices/v20211001:ProtectionIntent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProtectionIntent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectionIntentState, opts ...pulumi.ResourceOption) (*ProtectionIntent, error) {
	var resource ProtectionIntent
	err := ctx.ReadResource("azure-native:recoveryservices/v20211001:ProtectionIntent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type protectionIntentState struct {
}

type ProtectionIntentState struct {
}

func (ProtectionIntentState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionIntentState)(nil)).Elem()
}

type protectionIntentArgs struct {
	ETag              *string           `pulumi:"eTag"`
	FabricName        string            `pulumi:"fabricName"`
	IntentObjectName  *string           `pulumi:"intentObjectName"`
	Location          *string           `pulumi:"location"`
	Properties        interface{}       `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VaultName         string            `pulumi:"vaultName"`
}


type ProtectionIntentArgs struct {
	ETag              pulumi.StringPtrInput
	FabricName        pulumi.StringInput
	IntentObjectName  pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VaultName         pulumi.StringInput
}

func (ProtectionIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionIntentArgs)(nil)).Elem()
}

type ProtectionIntentInput interface {
	pulumi.Input

	ToProtectionIntentOutput() ProtectionIntentOutput
	ToProtectionIntentOutputWithContext(ctx context.Context) ProtectionIntentOutput
}

func (*ProtectionIntent) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionIntent)(nil)).Elem()
}

func (i *ProtectionIntent) ToProtectionIntentOutput() ProtectionIntentOutput {
	return i.ToProtectionIntentOutputWithContext(context.Background())
}

func (i *ProtectionIntent) ToProtectionIntentOutputWithContext(ctx context.Context) ProtectionIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionIntentOutput)
}

type ProtectionIntentOutput struct{ *pulumi.OutputState }

func (ProtectionIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionIntent)(nil)).Elem()
}

func (o ProtectionIntentOutput) ToProtectionIntentOutput() ProtectionIntentOutput {
	return o
}

func (o ProtectionIntentOutput) ToProtectionIntentOutputWithContext(ctx context.Context) ProtectionIntentOutput {
	return o
}

func (o ProtectionIntentOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionIntent) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o ProtectionIntentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionIntent) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ProtectionIntentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectionIntent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProtectionIntentOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ProtectionIntent) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o ProtectionIntentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProtectionIntent) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ProtectionIntentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectionIntent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProtectionIntentOutput{})
}
