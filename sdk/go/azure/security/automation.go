


package security

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Automation struct {
	pulumi.CustomResourceState

	Actions     pulumi.ArrayOutput                  `pulumi:"actions"`
	Description pulumi.StringPtrOutput              `pulumi:"description"`
	Etag        pulumi.StringPtrOutput              `pulumi:"etag"`
	IsEnabled   pulumi.BoolPtrOutput                `pulumi:"isEnabled"`
	Kind        pulumi.StringPtrOutput              `pulumi:"kind"`
	Location    pulumi.StringPtrOutput              `pulumi:"location"`
	Name        pulumi.StringOutput                 `pulumi:"name"`
	Scopes      AutomationScopeResponseArrayOutput  `pulumi:"scopes"`
	Sources     AutomationSourceResponseArrayOutput `pulumi:"sources"`
	Tags        pulumi.StringMapOutput              `pulumi:"tags"`
	Type        pulumi.StringOutput                 `pulumi:"type"`
}


func NewAutomation(ctx *pulumi.Context,
	name string, args *AutomationArgs, opts ...pulumi.ResourceOption) (*Automation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security/v20190101preview:Automation"),
		},
	})
	opts = append(opts, aliases)
	var resource Automation
	err := ctx.RegisterResource("azure-native:security:Automation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAutomation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationState, opts ...pulumi.ResourceOption) (*Automation, error) {
	var resource Automation
	err := ctx.ReadResource("azure-native:security:Automation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type automationState struct {
}

type AutomationState struct {
}

func (AutomationState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationState)(nil)).Elem()
}

type automationArgs struct {
	Actions           []interface{}      `pulumi:"actions"`
	AutomationName    *string            `pulumi:"automationName"`
	Description       *string            `pulumi:"description"`
	IsEnabled         *bool              `pulumi:"isEnabled"`
	Kind              *string            `pulumi:"kind"`
	Location          *string            `pulumi:"location"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	Scopes            []AutomationScope  `pulumi:"scopes"`
	Sources           []AutomationSource `pulumi:"sources"`
	Tags              map[string]string  `pulumi:"tags"`
}


type AutomationArgs struct {
	Actions           pulumi.ArrayInput
	AutomationName    pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	IsEnabled         pulumi.BoolPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Scopes            AutomationScopeArrayInput
	Sources           AutomationSourceArrayInput
	Tags              pulumi.StringMapInput
}

func (AutomationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationArgs)(nil)).Elem()
}

type AutomationInput interface {
	pulumi.Input

	ToAutomationOutput() AutomationOutput
	ToAutomationOutputWithContext(ctx context.Context) AutomationOutput
}

func (*Automation) ElementType() reflect.Type {
	return reflect.TypeOf((**Automation)(nil)).Elem()
}

func (i *Automation) ToAutomationOutput() AutomationOutput {
	return i.ToAutomationOutputWithContext(context.Background())
}

func (i *Automation) ToAutomationOutputWithContext(ctx context.Context) AutomationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationOutput)
}

type AutomationOutput struct{ *pulumi.OutputState }

func (AutomationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Automation)(nil)).Elem()
}

func (o AutomationOutput) ToAutomationOutput() AutomationOutput {
	return o
}

func (o AutomationOutput) ToAutomationOutputWithContext(ctx context.Context) AutomationOutput {
	return o
}

func (o AutomationOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Automation) pulumi.ArrayOutput { return v.Actions }).(pulumi.ArrayOutput)
}

func (o AutomationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AutomationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o AutomationOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Automation) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o AutomationOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o AutomationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AutomationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AutomationOutput) Scopes() AutomationScopeResponseArrayOutput {
	return o.ApplyT(func(v *Automation) AutomationScopeResponseArrayOutput { return v.Scopes }).(AutomationScopeResponseArrayOutput)
}

func (o AutomationOutput) Sources() AutomationSourceResponseArrayOutput {
	return o.ApplyT(func(v *Automation) AutomationSourceResponseArrayOutput { return v.Sources }).(AutomationSourceResponseArrayOutput)
}

func (o AutomationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AutomationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AutomationOutput{})
}
