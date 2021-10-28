


package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DisasterRecoveryConfiguration struct {
	pulumi.CustomResourceState

	AutoFailover             pulumi.StringOutput `pulumi:"autoFailover"`
	FailoverPolicy           pulumi.StringOutput `pulumi:"failoverPolicy"`
	Location                 pulumi.StringOutput `pulumi:"location"`
	LogicalServerName        pulumi.StringOutput `pulumi:"logicalServerName"`
	Name                     pulumi.StringOutput `pulumi:"name"`
	PartnerLogicalServerName pulumi.StringOutput `pulumi:"partnerLogicalServerName"`
	PartnerServerId          pulumi.StringOutput `pulumi:"partnerServerId"`
	Role                     pulumi.StringOutput `pulumi:"role"`
	Status                   pulumi.StringOutput `pulumi:"status"`
	Type                     pulumi.StringOutput `pulumi:"type"`
}


func NewDisasterRecoveryConfiguration(ctx *pulumi.Context,
	name string, args *DisasterRecoveryConfigurationArgs, opts ...pulumi.ResourceOption) (*DisasterRecoveryConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql:DisasterRecoveryConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:DisasterRecoveryConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20140401:DisasterRecoveryConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource DisasterRecoveryConfiguration
	err := ctx.RegisterResource("azure-native:sql:DisasterRecoveryConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDisasterRecoveryConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DisasterRecoveryConfigurationState, opts ...pulumi.ResourceOption) (*DisasterRecoveryConfiguration, error) {
	var resource DisasterRecoveryConfiguration
	err := ctx.ReadResource("azure-native:sql:DisasterRecoveryConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type disasterRecoveryConfigurationState struct {
}

type DisasterRecoveryConfigurationState struct {
}

func (DisasterRecoveryConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*disasterRecoveryConfigurationState)(nil)).Elem()
}

type disasterRecoveryConfigurationArgs struct {
	DisasterRecoveryConfigurationName *string `pulumi:"disasterRecoveryConfigurationName"`
	ResourceGroupName                 string  `pulumi:"resourceGroupName"`
	ServerName                        string  `pulumi:"serverName"`
}


type DisasterRecoveryConfigurationArgs struct {
	DisasterRecoveryConfigurationName pulumi.StringPtrInput
	ResourceGroupName                 pulumi.StringInput
	ServerName                        pulumi.StringInput
}

func (DisasterRecoveryConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*disasterRecoveryConfigurationArgs)(nil)).Elem()
}

type DisasterRecoveryConfigurationInput interface {
	pulumi.Input

	ToDisasterRecoveryConfigurationOutput() DisasterRecoveryConfigurationOutput
	ToDisasterRecoveryConfigurationOutputWithContext(ctx context.Context) DisasterRecoveryConfigurationOutput
}

func (*DisasterRecoveryConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*DisasterRecoveryConfiguration)(nil))
}

func (i *DisasterRecoveryConfiguration) ToDisasterRecoveryConfigurationOutput() DisasterRecoveryConfigurationOutput {
	return i.ToDisasterRecoveryConfigurationOutputWithContext(context.Background())
}

func (i *DisasterRecoveryConfiguration) ToDisasterRecoveryConfigurationOutputWithContext(ctx context.Context) DisasterRecoveryConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisasterRecoveryConfigurationOutput)
}

type DisasterRecoveryConfigurationOutput struct{ *pulumi.OutputState }

func (DisasterRecoveryConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DisasterRecoveryConfiguration)(nil))
}

func (o DisasterRecoveryConfigurationOutput) ToDisasterRecoveryConfigurationOutput() DisasterRecoveryConfigurationOutput {
	return o
}

func (o DisasterRecoveryConfigurationOutput) ToDisasterRecoveryConfigurationOutputWithContext(ctx context.Context) DisasterRecoveryConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DisasterRecoveryConfigurationInput)(nil)).Elem(), &DisasterRecoveryConfiguration{})
	pulumi.RegisterOutputType(DisasterRecoveryConfigurationOutput{})
}
