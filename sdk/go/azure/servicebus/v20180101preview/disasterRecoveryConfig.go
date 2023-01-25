


package v20180101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DisasterRecoveryConfig struct {
	pulumi.CustomResourceState

	AlternateName                     pulumi.StringPtrOutput `pulumi:"alternateName"`
	Name                              pulumi.StringOutput    `pulumi:"name"`
	PartnerNamespace                  pulumi.StringPtrOutput `pulumi:"partnerNamespace"`
	PendingReplicationOperationsCount pulumi.Float64Output   `pulumi:"pendingReplicationOperationsCount"`
	ProvisioningState                 pulumi.StringOutput    `pulumi:"provisioningState"`
	Role                              pulumi.StringOutput    `pulumi:"role"`
	Type                              pulumi.StringOutput    `pulumi:"type"`
}


func NewDisasterRecoveryConfig(ctx *pulumi.Context,
	name string, args *DisasterRecoveryConfigArgs, opts ...pulumi.ResourceOption) (*DisasterRecoveryConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20211101:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20220101preview:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20221001preview:DisasterRecoveryConfig"),
		},
	})
	opts = append(opts, aliases)
	var resource DisasterRecoveryConfig
	err := ctx.RegisterResource("azure-native:servicebus/v20180101preview:DisasterRecoveryConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDisasterRecoveryConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DisasterRecoveryConfigState, opts ...pulumi.ResourceOption) (*DisasterRecoveryConfig, error) {
	var resource DisasterRecoveryConfig
	err := ctx.ReadResource("azure-native:servicebus/v20180101preview:DisasterRecoveryConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type disasterRecoveryConfigState struct {
}

type DisasterRecoveryConfigState struct {
}

func (DisasterRecoveryConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*disasterRecoveryConfigState)(nil)).Elem()
}

type disasterRecoveryConfigArgs struct {
	Alias             *string `pulumi:"alias"`
	AlternateName     *string `pulumi:"alternateName"`
	NamespaceName     string  `pulumi:"namespaceName"`
	PartnerNamespace  *string `pulumi:"partnerNamespace"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type DisasterRecoveryConfigArgs struct {
	Alias             pulumi.StringPtrInput
	AlternateName     pulumi.StringPtrInput
	NamespaceName     pulumi.StringInput
	PartnerNamespace  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (DisasterRecoveryConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*disasterRecoveryConfigArgs)(nil)).Elem()
}

type DisasterRecoveryConfigInput interface {
	pulumi.Input

	ToDisasterRecoveryConfigOutput() DisasterRecoveryConfigOutput
	ToDisasterRecoveryConfigOutputWithContext(ctx context.Context) DisasterRecoveryConfigOutput
}

func (*DisasterRecoveryConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**DisasterRecoveryConfig)(nil)).Elem()
}

func (i *DisasterRecoveryConfig) ToDisasterRecoveryConfigOutput() DisasterRecoveryConfigOutput {
	return i.ToDisasterRecoveryConfigOutputWithContext(context.Background())
}

func (i *DisasterRecoveryConfig) ToDisasterRecoveryConfigOutputWithContext(ctx context.Context) DisasterRecoveryConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisasterRecoveryConfigOutput)
}

type DisasterRecoveryConfigOutput struct{ *pulumi.OutputState }

func (DisasterRecoveryConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DisasterRecoveryConfig)(nil)).Elem()
}

func (o DisasterRecoveryConfigOutput) ToDisasterRecoveryConfigOutput() DisasterRecoveryConfigOutput {
	return o
}

func (o DisasterRecoveryConfigOutput) ToDisasterRecoveryConfigOutputWithContext(ctx context.Context) DisasterRecoveryConfigOutput {
	return o
}

func (o DisasterRecoveryConfigOutput) AlternateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.StringPtrOutput { return v.AlternateName }).(pulumi.StringPtrOutput)
}

func (o DisasterRecoveryConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DisasterRecoveryConfigOutput) PartnerNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.StringPtrOutput { return v.PartnerNamespace }).(pulumi.StringPtrOutput)
}

func (o DisasterRecoveryConfigOutput) PendingReplicationOperationsCount() pulumi.Float64Output {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.Float64Output { return v.PendingReplicationOperationsCount }).(pulumi.Float64Output)
}

func (o DisasterRecoveryConfigOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DisasterRecoveryConfigOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o DisasterRecoveryConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DisasterRecoveryConfigOutput{})
}
