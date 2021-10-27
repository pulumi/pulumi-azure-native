


package v20200804preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HealthAlert struct {
	pulumi.CustomResourceState

	Actions         HealthAlertActionResponseArrayOutput `pulumi:"actions"`
	Criteria        HealthAlertCriteriaResponseOutput    `pulumi:"criteria"`
	Description     pulumi.StringOutput                  `pulumi:"description"`
	Enabled         pulumi.BoolOutput                    `pulumi:"enabled"`
	LastUpdatedTime pulumi.StringOutput                  `pulumi:"lastUpdatedTime"`
	Location        pulumi.StringOutput                  `pulumi:"location"`
	Name            pulumi.StringOutput                  `pulumi:"name"`
	Scopes          pulumi.StringArrayOutput             `pulumi:"scopes"`
	Tags            pulumi.StringMapOutput               `pulumi:"tags"`
	Type            pulumi.StringOutput                  `pulumi:"type"`
}


func NewHealthAlert(ctx *pulumi.Context,
	name string, args *HealthAlertArgs, opts ...pulumi.ResourceOption) (*HealthAlert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:alertsmanagement/v20200804preview:HealthAlert"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement:HealthAlert"),
		},
		{
			Type: pulumi.String("azure-nextgen:alertsmanagement:HealthAlert"),
		},
	})
	opts = append(opts, aliases)
	var resource HealthAlert
	err := ctx.RegisterResource("azure-native:alertsmanagement/v20200804preview:HealthAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHealthAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HealthAlertState, opts ...pulumi.ResourceOption) (*HealthAlert, error) {
	var resource HealthAlert
	err := ctx.ReadResource("azure-native:alertsmanagement/v20200804preview:HealthAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type healthAlertState struct {
}

type HealthAlertState struct {
}

func (HealthAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*healthAlertState)(nil)).Elem()
}

type healthAlertArgs struct {
	Actions           []HealthAlertAction `pulumi:"actions"`
	Criteria          HealthAlertCriteria `pulumi:"criteria"`
	Description       string              `pulumi:"description"`
	Enabled           bool                `pulumi:"enabled"`
	Location          *string             `pulumi:"location"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	RuleName          *string             `pulumi:"ruleName"`
	Scopes            []string            `pulumi:"scopes"`
	Tags              map[string]string   `pulumi:"tags"`
}


type HealthAlertArgs struct {
	Actions           HealthAlertActionArrayInput
	Criteria          HealthAlertCriteriaInput
	Description       pulumi.StringInput
	Enabled           pulumi.BoolInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RuleName          pulumi.StringPtrInput
	Scopes            pulumi.StringArrayInput
	Tags              pulumi.StringMapInput
}

func (HealthAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*healthAlertArgs)(nil)).Elem()
}

type HealthAlertInput interface {
	pulumi.Input

	ToHealthAlertOutput() HealthAlertOutput
	ToHealthAlertOutputWithContext(ctx context.Context) HealthAlertOutput
}

func (*HealthAlert) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlert)(nil))
}

func (i *HealthAlert) ToHealthAlertOutput() HealthAlertOutput {
	return i.ToHealthAlertOutputWithContext(context.Background())
}

func (i *HealthAlert) ToHealthAlertOutputWithContext(ctx context.Context) HealthAlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertOutput)
}

type HealthAlertOutput struct{ *pulumi.OutputState }

func (HealthAlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlert)(nil))
}

func (o HealthAlertOutput) ToHealthAlertOutput() HealthAlertOutput {
	return o
}

func (o HealthAlertOutput) ToHealthAlertOutputWithContext(ctx context.Context) HealthAlertOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HealthAlertInput)(nil)).Elem(), &HealthAlert{})
	pulumi.RegisterOutputType(HealthAlertOutput{})
}
