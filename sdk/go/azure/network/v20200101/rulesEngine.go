


package v20200101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RulesEngine struct {
	pulumi.CustomResourceState

	Name          pulumi.StringOutput                `pulumi:"name"`
	ResourceState pulumi.StringOutput                `pulumi:"resourceState"`
	Rules         RulesEngineRuleResponseArrayOutput `pulumi:"rules"`
	Type          pulumi.StringOutput                `pulumi:"type"`
}


func NewRulesEngine(ctx *pulumi.Context,
	name string, args *RulesEngineArgs, opts ...pulumi.ResourceOption) (*RulesEngine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FrontDoorName == nil {
		return nil, errors.New("invalid value for required argument 'FrontDoorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:RulesEngine"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:RulesEngine"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:RulesEngine"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210601:RulesEngine"),
		},
	})
	opts = append(opts, aliases)
	var resource RulesEngine
	err := ctx.RegisterResource("azure-native:network/v20200101:RulesEngine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRulesEngine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RulesEngineState, opts ...pulumi.ResourceOption) (*RulesEngine, error) {
	var resource RulesEngine
	err := ctx.ReadResource("azure-native:network/v20200101:RulesEngine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type rulesEngineState struct {
}

type RulesEngineState struct {
}

func (RulesEngineState) ElementType() reflect.Type {
	return reflect.TypeOf((*rulesEngineState)(nil)).Elem()
}

type rulesEngineArgs struct {
	FrontDoorName     string            `pulumi:"frontDoorName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Rules             []RulesEngineRule `pulumi:"rules"`
	RulesEngineName   *string           `pulumi:"rulesEngineName"`
}


type RulesEngineArgs struct {
	FrontDoorName     pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Rules             RulesEngineRuleArrayInput
	RulesEngineName   pulumi.StringPtrInput
}

func (RulesEngineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rulesEngineArgs)(nil)).Elem()
}

type RulesEngineInput interface {
	pulumi.Input

	ToRulesEngineOutput() RulesEngineOutput
	ToRulesEngineOutputWithContext(ctx context.Context) RulesEngineOutput
}

func (*RulesEngine) ElementType() reflect.Type {
	return reflect.TypeOf((**RulesEngine)(nil)).Elem()
}

func (i *RulesEngine) ToRulesEngineOutput() RulesEngineOutput {
	return i.ToRulesEngineOutputWithContext(context.Background())
}

func (i *RulesEngine) ToRulesEngineOutputWithContext(ctx context.Context) RulesEngineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineOutput)
}

type RulesEngineOutput struct{ *pulumi.OutputState }

func (RulesEngineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RulesEngine)(nil)).Elem()
}

func (o RulesEngineOutput) ToRulesEngineOutput() RulesEngineOutput {
	return o
}

func (o RulesEngineOutput) ToRulesEngineOutputWithContext(ctx context.Context) RulesEngineOutput {
	return o
}

func (o RulesEngineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RulesEngine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RulesEngineOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *RulesEngine) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o RulesEngineOutput) Rules() RulesEngineRuleResponseArrayOutput {
	return o.ApplyT(func(v *RulesEngine) RulesEngineRuleResponseArrayOutput { return v.Rules }).(RulesEngineRuleResponseArrayOutput)
}

func (o RulesEngineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RulesEngine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RulesEngineOutput{})
}
