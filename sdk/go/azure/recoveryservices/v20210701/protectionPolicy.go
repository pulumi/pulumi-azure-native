


package v20210701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProtectionPolicy struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput `pulumi:"eTag"`
	Location   pulumi.StringPtrOutput `pulumi:"location"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Properties pulumi.AnyOutput       `pulumi:"properties"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewProtectionPolicy(ctx *pulumi.Context,
	name string, args *ProtectionPolicyArgs, opts ...pulumi.ResourceOption) (*ProtectionPolicy, error) {
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
			Type: pulumi.String("azure-native:recoveryservices:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160601:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201001:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201201:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210101:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201preview:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ProtectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220601preview:ProtectionPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ProtectionPolicy
	err := ctx.RegisterResource("azure-native:recoveryservices/v20210701:ProtectionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProtectionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectionPolicyState, opts ...pulumi.ResourceOption) (*ProtectionPolicy, error) {
	var resource ProtectionPolicy
	err := ctx.ReadResource("azure-native:recoveryservices/v20210701:ProtectionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type protectionPolicyState struct {
}

type ProtectionPolicyState struct {
}

func (ProtectionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionPolicyState)(nil)).Elem()
}

type protectionPolicyArgs struct {
	ETag              *string           `pulumi:"eTag"`
	Location          *string           `pulumi:"location"`
	PolicyName        *string           `pulumi:"policyName"`
	Properties        interface{}       `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VaultName         string            `pulumi:"vaultName"`
}


type ProtectionPolicyArgs struct {
	ETag              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	PolicyName        pulumi.StringPtrInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VaultName         pulumi.StringInput
}

func (ProtectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionPolicyArgs)(nil)).Elem()
}

type ProtectionPolicyInput interface {
	pulumi.Input

	ToProtectionPolicyOutput() ProtectionPolicyOutput
	ToProtectionPolicyOutputWithContext(ctx context.Context) ProtectionPolicyOutput
}

func (*ProtectionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionPolicy)(nil)).Elem()
}

func (i *ProtectionPolicy) ToProtectionPolicyOutput() ProtectionPolicyOutput {
	return i.ToProtectionPolicyOutputWithContext(context.Background())
}

func (i *ProtectionPolicy) ToProtectionPolicyOutputWithContext(ctx context.Context) ProtectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionPolicyOutput)
}

type ProtectionPolicyOutput struct{ *pulumi.OutputState }

func (ProtectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionPolicy)(nil)).Elem()
}

func (o ProtectionPolicyOutput) ToProtectionPolicyOutput() ProtectionPolicyOutput {
	return o
}

func (o ProtectionPolicyOutput) ToProtectionPolicyOutputWithContext(ctx context.Context) ProtectionPolicyOutput {
	return o
}

func (o ProtectionPolicyOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionPolicy) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o ProtectionPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ProtectionPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectionPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProtectionPolicyOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ProtectionPolicy) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o ProtectionPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProtectionPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ProtectionPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectionPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProtectionPolicyOutput{})
}
