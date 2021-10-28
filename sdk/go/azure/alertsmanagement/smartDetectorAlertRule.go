


package alertsmanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SmartDetectorAlertRule struct {
	pulumi.CustomResourceState

	ActionGroups ActionGroupsInformationResponseOutput  `pulumi:"actionGroups"`
	Description  pulumi.StringPtrOutput                 `pulumi:"description"`
	Detector     DetectorResponseOutput                 `pulumi:"detector"`
	Frequency    pulumi.StringOutput                    `pulumi:"frequency"`
	Location     pulumi.StringPtrOutput                 `pulumi:"location"`
	Name         pulumi.StringOutput                    `pulumi:"name"`
	Scope        pulumi.StringArrayOutput               `pulumi:"scope"`
	Severity     pulumi.StringOutput                    `pulumi:"severity"`
	State        pulumi.StringOutput                    `pulumi:"state"`
	Tags         pulumi.StringMapOutput                 `pulumi:"tags"`
	Throttling   ThrottlingInformationResponsePtrOutput `pulumi:"throttling"`
	Type         pulumi.StringOutput                    `pulumi:"type"`
}


func NewSmartDetectorAlertRule(ctx *pulumi.Context,
	name string, args *SmartDetectorAlertRuleArgs, opts ...pulumi.ResourceOption) (*SmartDetectorAlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActionGroups == nil {
		return nil, errors.New("invalid value for required argument 'ActionGroups'")
	}
	if args.Detector == nil {
		return nil, errors.New("invalid value for required argument 'Detector'")
	}
	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	if args.Location == nil {
		args.Location = pulumi.StringPtr("global")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:alertsmanagement:SmartDetectorAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20190301:SmartDetectorAlertRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:alertsmanagement/v20190301:SmartDetectorAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20190601:SmartDetectorAlertRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:alertsmanagement/v20190601:SmartDetectorAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20210401:SmartDetectorAlertRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:alertsmanagement/v20210401:SmartDetectorAlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource SmartDetectorAlertRule
	err := ctx.RegisterResource("azure-native:alertsmanagement:SmartDetectorAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSmartDetectorAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SmartDetectorAlertRuleState, opts ...pulumi.ResourceOption) (*SmartDetectorAlertRule, error) {
	var resource SmartDetectorAlertRule
	err := ctx.ReadResource("azure-native:alertsmanagement:SmartDetectorAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type smartDetectorAlertRuleState struct {
}

type SmartDetectorAlertRuleState struct {
}

func (SmartDetectorAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*smartDetectorAlertRuleState)(nil)).Elem()
}

type smartDetectorAlertRuleArgs struct {
	ActionGroups      ActionGroupsInformation `pulumi:"actionGroups"`
	AlertRuleName     *string                 `pulumi:"alertRuleName"`
	Description       *string                 `pulumi:"description"`
	Detector          Detector                `pulumi:"detector"`
	Frequency         string                  `pulumi:"frequency"`
	Location          *string                 `pulumi:"location"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Scope             []string                `pulumi:"scope"`
	Severity          string                  `pulumi:"severity"`
	State             string                  `pulumi:"state"`
	Tags              map[string]string       `pulumi:"tags"`
	Throttling        *ThrottlingInformation  `pulumi:"throttling"`
}


type SmartDetectorAlertRuleArgs struct {
	ActionGroups      ActionGroupsInformationInput
	AlertRuleName     pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	Detector          DetectorInput
	Frequency         pulumi.StringInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Scope             pulumi.StringArrayInput
	Severity          pulumi.StringInput
	State             pulumi.StringInput
	Tags              pulumi.StringMapInput
	Throttling        ThrottlingInformationPtrInput
}

func (SmartDetectorAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*smartDetectorAlertRuleArgs)(nil)).Elem()
}

type SmartDetectorAlertRuleInput interface {
	pulumi.Input

	ToSmartDetectorAlertRuleOutput() SmartDetectorAlertRuleOutput
	ToSmartDetectorAlertRuleOutputWithContext(ctx context.Context) SmartDetectorAlertRuleOutput
}

func (*SmartDetectorAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((*SmartDetectorAlertRule)(nil))
}

func (i *SmartDetectorAlertRule) ToSmartDetectorAlertRuleOutput() SmartDetectorAlertRuleOutput {
	return i.ToSmartDetectorAlertRuleOutputWithContext(context.Background())
}

func (i *SmartDetectorAlertRule) ToSmartDetectorAlertRuleOutputWithContext(ctx context.Context) SmartDetectorAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmartDetectorAlertRuleOutput)
}

type SmartDetectorAlertRuleOutput struct{ *pulumi.OutputState }

func (SmartDetectorAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SmartDetectorAlertRule)(nil))
}

func (o SmartDetectorAlertRuleOutput) ToSmartDetectorAlertRuleOutput() SmartDetectorAlertRuleOutput {
	return o
}

func (o SmartDetectorAlertRuleOutput) ToSmartDetectorAlertRuleOutputWithContext(ctx context.Context) SmartDetectorAlertRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SmartDetectorAlertRuleOutput{})
}
