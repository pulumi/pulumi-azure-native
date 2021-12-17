


package v20170401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActivityLogAlert struct {
	pulumi.CustomResourceState

	Actions     ActivityLogAlertActionListResponseOutput     `pulumi:"actions"`
	Condition   ActivityLogAlertAllOfConditionResponseOutput `pulumi:"condition"`
	Description pulumi.StringPtrOutput                       `pulumi:"description"`
	Enabled     pulumi.BoolPtrOutput                         `pulumi:"enabled"`
	Location    pulumi.StringOutput                          `pulumi:"location"`
	Name        pulumi.StringOutput                          `pulumi:"name"`
	Scopes      pulumi.StringArrayOutput                     `pulumi:"scopes"`
	Tags        pulumi.StringMapOutput                       `pulumi:"tags"`
	Type        pulumi.StringOutput                          `pulumi:"type"`
}


func NewActivityLogAlert(ctx *pulumi.Context,
	name string, args *ActivityLogAlertArgs, opts ...pulumi.ResourceOption) (*ActivityLogAlert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.Condition == nil {
		return nil, errors.New("invalid value for required argument 'Condition'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	if isZero(args.Enabled) {
		args.Enabled = pulumi.BoolPtr(true)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:ActivityLogAlert"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20201001:ActivityLogAlert"),
		},
	})
	opts = append(opts, aliases)
	var resource ActivityLogAlert
	err := ctx.RegisterResource("azure-native:insights/v20170401:ActivityLogAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetActivityLogAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActivityLogAlertState, opts ...pulumi.ResourceOption) (*ActivityLogAlert, error) {
	var resource ActivityLogAlert
	err := ctx.ReadResource("azure-native:insights/v20170401:ActivityLogAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type activityLogAlertState struct {
}

type ActivityLogAlertState struct {
}

func (ActivityLogAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*activityLogAlertState)(nil)).Elem()
}

type activityLogAlertArgs struct {
	Actions              ActivityLogAlertActionList     `pulumi:"actions"`
	ActivityLogAlertName *string                        `pulumi:"activityLogAlertName"`
	Condition            ActivityLogAlertAllOfCondition `pulumi:"condition"`
	Description          *string                        `pulumi:"description"`
	Enabled              *bool                          `pulumi:"enabled"`
	Location             *string                        `pulumi:"location"`
	ResourceGroupName    string                         `pulumi:"resourceGroupName"`
	Scopes               []string                       `pulumi:"scopes"`
	Tags                 map[string]string              `pulumi:"tags"`
}


type ActivityLogAlertArgs struct {
	Actions              ActivityLogAlertActionListInput
	ActivityLogAlertName pulumi.StringPtrInput
	Condition            ActivityLogAlertAllOfConditionInput
	Description          pulumi.StringPtrInput
	Enabled              pulumi.BoolPtrInput
	Location             pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Scopes               pulumi.StringArrayInput
	Tags                 pulumi.StringMapInput
}

func (ActivityLogAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activityLogAlertArgs)(nil)).Elem()
}

type ActivityLogAlertInput interface {
	pulumi.Input

	ToActivityLogAlertOutput() ActivityLogAlertOutput
	ToActivityLogAlertOutputWithContext(ctx context.Context) ActivityLogAlertOutput
}

func (*ActivityLogAlert) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivityLogAlert)(nil)).Elem()
}

func (i *ActivityLogAlert) ToActivityLogAlertOutput() ActivityLogAlertOutput {
	return i.ToActivityLogAlertOutputWithContext(context.Background())
}

func (i *ActivityLogAlert) ToActivityLogAlertOutputWithContext(ctx context.Context) ActivityLogAlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityLogAlertOutput)
}

type ActivityLogAlertOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivityLogAlert)(nil)).Elem()
}

func (o ActivityLogAlertOutput) ToActivityLogAlertOutput() ActivityLogAlertOutput {
	return o
}

func (o ActivityLogAlertOutput) ToActivityLogAlertOutputWithContext(ctx context.Context) ActivityLogAlertOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ActivityLogAlertOutput{})
}
