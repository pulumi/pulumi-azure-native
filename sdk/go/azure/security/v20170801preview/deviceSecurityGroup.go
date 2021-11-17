


package v20170801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DeviceSecurityGroup struct {
	pulumi.CustomResourceState

	AllowlistRules  AllowlistCustomAlertRuleResponseArrayOutput  `pulumi:"allowlistRules"`
	DenylistRules   DenylistCustomAlertRuleResponseArrayOutput   `pulumi:"denylistRules"`
	Name            pulumi.StringOutput                          `pulumi:"name"`
	ThresholdRules  ThresholdCustomAlertRuleResponseArrayOutput  `pulumi:"thresholdRules"`
	TimeWindowRules TimeWindowCustomAlertRuleResponseArrayOutput `pulumi:"timeWindowRules"`
	Type            pulumi.StringOutput                          `pulumi:"type"`
}


func NewDeviceSecurityGroup(ctx *pulumi.Context,
	name string, args *DeviceSecurityGroupArgs, opts ...pulumi.ResourceOption) (*DeviceSecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:DeviceSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:security/v20190801:DeviceSecurityGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource DeviceSecurityGroup
	err := ctx.RegisterResource("azure-native:security/v20170801preview:DeviceSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDeviceSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceSecurityGroupState, opts ...pulumi.ResourceOption) (*DeviceSecurityGroup, error) {
	var resource DeviceSecurityGroup
	err := ctx.ReadResource("azure-native:security/v20170801preview:DeviceSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type deviceSecurityGroupState struct {
}

type DeviceSecurityGroupState struct {
}

func (DeviceSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceSecurityGroupState)(nil)).Elem()
}

type deviceSecurityGroupArgs struct {
	AllowlistRules          []AllowlistCustomAlertRule  `pulumi:"allowlistRules"`
	DenylistRules           []DenylistCustomAlertRule   `pulumi:"denylistRules"`
	DeviceSecurityGroupName *string                     `pulumi:"deviceSecurityGroupName"`
	ResourceId              string                      `pulumi:"resourceId"`
	ThresholdRules          []ThresholdCustomAlertRule  `pulumi:"thresholdRules"`
	TimeWindowRules         []TimeWindowCustomAlertRule `pulumi:"timeWindowRules"`
}


type DeviceSecurityGroupArgs struct {
	AllowlistRules          AllowlistCustomAlertRuleArrayInput
	DenylistRules           DenylistCustomAlertRuleArrayInput
	DeviceSecurityGroupName pulumi.StringPtrInput
	ResourceId              pulumi.StringInput
	ThresholdRules          ThresholdCustomAlertRuleArrayInput
	TimeWindowRules         TimeWindowCustomAlertRuleArrayInput
}

func (DeviceSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceSecurityGroupArgs)(nil)).Elem()
}

type DeviceSecurityGroupInput interface {
	pulumi.Input

	ToDeviceSecurityGroupOutput() DeviceSecurityGroupOutput
	ToDeviceSecurityGroupOutputWithContext(ctx context.Context) DeviceSecurityGroupOutput
}

func (*DeviceSecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceSecurityGroup)(nil))
}

func (i *DeviceSecurityGroup) ToDeviceSecurityGroupOutput() DeviceSecurityGroupOutput {
	return i.ToDeviceSecurityGroupOutputWithContext(context.Background())
}

func (i *DeviceSecurityGroup) ToDeviceSecurityGroupOutputWithContext(ctx context.Context) DeviceSecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceSecurityGroupOutput)
}

type DeviceSecurityGroupOutput struct{ *pulumi.OutputState }

func (DeviceSecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceSecurityGroup)(nil))
}

func (o DeviceSecurityGroupOutput) ToDeviceSecurityGroupOutput() DeviceSecurityGroupOutput {
	return o
}

func (o DeviceSecurityGroupOutput) ToDeviceSecurityGroupOutputWithContext(ctx context.Context) DeviceSecurityGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DeviceSecurityGroupOutput{})
}
