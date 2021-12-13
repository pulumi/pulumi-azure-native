


package logic

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RosettaNetProcessConfiguration struct {
	pulumi.CustomResourceState

	ActivitySettings      RosettaNetPipActivitySettingsResponseOutput `pulumi:"activitySettings"`
	ChangedTime           pulumi.StringOutput                         `pulumi:"changedTime"`
	CreatedTime           pulumi.StringOutput                         `pulumi:"createdTime"`
	Description           pulumi.StringPtrOutput                      `pulumi:"description"`
	InitiatorRoleSettings RosettaNetPipRoleSettingsResponseOutput     `pulumi:"initiatorRoleSettings"`
	Location              pulumi.StringPtrOutput                      `pulumi:"location"`
	Metadata              pulumi.StringMapOutput                      `pulumi:"metadata"`
	Name                  pulumi.StringOutput                         `pulumi:"name"`
	ProcessCode           pulumi.StringOutput                         `pulumi:"processCode"`
	ProcessName           pulumi.StringOutput                         `pulumi:"processName"`
	ProcessVersion        pulumi.StringOutput                         `pulumi:"processVersion"`
	ResponderRoleSettings RosettaNetPipRoleSettingsResponseOutput     `pulumi:"responderRoleSettings"`
	Tags                  pulumi.StringMapOutput                      `pulumi:"tags"`
	Type                  pulumi.StringOutput                         `pulumi:"type"`
}


func NewRosettaNetProcessConfiguration(ctx *pulumi.Context,
	name string, args *RosettaNetProcessConfigurationArgs, opts ...pulumi.ResourceOption) (*RosettaNetProcessConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActivitySettings == nil {
		return nil, errors.New("invalid value for required argument 'ActivitySettings'")
	}
	if args.InitiatorRoleSettings == nil {
		return nil, errors.New("invalid value for required argument 'InitiatorRoleSettings'")
	}
	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ProcessCode == nil {
		return nil, errors.New("invalid value for required argument 'ProcessCode'")
	}
	if args.ProcessName == nil {
		return nil, errors.New("invalid value for required argument 'ProcessName'")
	}
	if args.ProcessVersion == nil {
		return nil, errors.New("invalid value for required argument 'ProcessVersion'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResponderRoleSettings == nil {
		return nil, errors.New("invalid value for required argument 'ResponderRoleSettings'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic/v20160601:RosettaNetProcessConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource RosettaNetProcessConfiguration
	err := ctx.RegisterResource("azure-native:logic:RosettaNetProcessConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRosettaNetProcessConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RosettaNetProcessConfigurationState, opts ...pulumi.ResourceOption) (*RosettaNetProcessConfiguration, error) {
	var resource RosettaNetProcessConfiguration
	err := ctx.ReadResource("azure-native:logic:RosettaNetProcessConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type rosettaNetProcessConfigurationState struct {
}

type RosettaNetProcessConfigurationState struct {
}

func (RosettaNetProcessConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*rosettaNetProcessConfigurationState)(nil)).Elem()
}

type rosettaNetProcessConfigurationArgs struct {
	ActivitySettings                   RosettaNetPipActivitySettings `pulumi:"activitySettings"`
	Description                        *string                       `pulumi:"description"`
	InitiatorRoleSettings              RosettaNetPipRoleSettings     `pulumi:"initiatorRoleSettings"`
	IntegrationAccountName             string                        `pulumi:"integrationAccountName"`
	Location                           *string                       `pulumi:"location"`
	Metadata                           map[string]string             `pulumi:"metadata"`
	ProcessCode                        string                        `pulumi:"processCode"`
	ProcessName                        string                        `pulumi:"processName"`
	ProcessVersion                     string                        `pulumi:"processVersion"`
	ResourceGroupName                  string                        `pulumi:"resourceGroupName"`
	ResponderRoleSettings              RosettaNetPipRoleSettings     `pulumi:"responderRoleSettings"`
	RosettaNetProcessConfigurationName *string                       `pulumi:"rosettaNetProcessConfigurationName"`
	Tags                               map[string]string             `pulumi:"tags"`
}


type RosettaNetProcessConfigurationArgs struct {
	ActivitySettings                   RosettaNetPipActivitySettingsInput
	Description                        pulumi.StringPtrInput
	InitiatorRoleSettings              RosettaNetPipRoleSettingsInput
	IntegrationAccountName             pulumi.StringInput
	Location                           pulumi.StringPtrInput
	Metadata                           pulumi.StringMapInput
	ProcessCode                        pulumi.StringInput
	ProcessName                        pulumi.StringInput
	ProcessVersion                     pulumi.StringInput
	ResourceGroupName                  pulumi.StringInput
	ResponderRoleSettings              RosettaNetPipRoleSettingsInput
	RosettaNetProcessConfigurationName pulumi.StringPtrInput
	Tags                               pulumi.StringMapInput
}

func (RosettaNetProcessConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rosettaNetProcessConfigurationArgs)(nil)).Elem()
}

type RosettaNetProcessConfigurationInput interface {
	pulumi.Input

	ToRosettaNetProcessConfigurationOutput() RosettaNetProcessConfigurationOutput
	ToRosettaNetProcessConfigurationOutputWithContext(ctx context.Context) RosettaNetProcessConfigurationOutput
}

func (*RosettaNetProcessConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**RosettaNetProcessConfiguration)(nil)).Elem()
}

func (i *RosettaNetProcessConfiguration) ToRosettaNetProcessConfigurationOutput() RosettaNetProcessConfigurationOutput {
	return i.ToRosettaNetProcessConfigurationOutputWithContext(context.Background())
}

func (i *RosettaNetProcessConfiguration) ToRosettaNetProcessConfigurationOutputWithContext(ctx context.Context) RosettaNetProcessConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RosettaNetProcessConfigurationOutput)
}

type RosettaNetProcessConfigurationOutput struct{ *pulumi.OutputState }

func (RosettaNetProcessConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RosettaNetProcessConfiguration)(nil)).Elem()
}

func (o RosettaNetProcessConfigurationOutput) ToRosettaNetProcessConfigurationOutput() RosettaNetProcessConfigurationOutput {
	return o
}

func (o RosettaNetProcessConfigurationOutput) ToRosettaNetProcessConfigurationOutputWithContext(ctx context.Context) RosettaNetProcessConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RosettaNetProcessConfigurationOutput{})
}
