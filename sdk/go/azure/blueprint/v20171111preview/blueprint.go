


package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Blueprint struct {
	pulumi.CustomResourceState

	Description    pulumi.StringPtrOutput                   `pulumi:"description"`
	DisplayName    pulumi.StringPtrOutput                   `pulumi:"displayName"`
	Layout         pulumi.AnyOutput                         `pulumi:"layout"`
	Name           pulumi.StringOutput                      `pulumi:"name"`
	Parameters     ParameterDefinitionResponseMapOutput     `pulumi:"parameters"`
	ResourceGroups ResourceGroupDefinitionResponseMapOutput `pulumi:"resourceGroups"`
	Status         BlueprintStatusResponseOutput            `pulumi:"status"`
	TargetScope    pulumi.StringOutput                      `pulumi:"targetScope"`
	Type           pulumi.StringOutput                      `pulumi:"type"`
	Versions       pulumi.AnyOutput                         `pulumi:"versions"`
}


func NewBlueprint(ctx *pulumi.Context,
	name string, args *BlueprintArgs, opts ...pulumi.ResourceOption) (*Blueprint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupName'")
	}
	if args.TargetScope == nil {
		return nil, errors.New("invalid value for required argument 'TargetScope'")
	}
	var resource Blueprint
	err := ctx.RegisterResource("azure-native:blueprint/v20171111preview:Blueprint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlueprint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlueprintState, opts ...pulumi.ResourceOption) (*Blueprint, error) {
	var resource Blueprint
	err := ctx.ReadResource("azure-native:blueprint/v20171111preview:Blueprint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blueprintState struct {
}

type BlueprintState struct {
}

func (BlueprintState) ElementType() reflect.Type {
	return reflect.TypeOf((*blueprintState)(nil)).Elem()
}

type blueprintArgs struct {
	BlueprintName       *string                            `pulumi:"blueprintName"`
	Description         *string                            `pulumi:"description"`
	DisplayName         *string                            `pulumi:"displayName"`
	Layout              interface{}                        `pulumi:"layout"`
	ManagementGroupName string                             `pulumi:"managementGroupName"`
	Parameters          map[string]ParameterDefinition     `pulumi:"parameters"`
	ResourceGroups      map[string]ResourceGroupDefinition `pulumi:"resourceGroups"`
	TargetScope         string                             `pulumi:"targetScope"`
	Versions            interface{}                        `pulumi:"versions"`
}


type BlueprintArgs struct {
	BlueprintName       pulumi.StringPtrInput
	Description         pulumi.StringPtrInput
	DisplayName         pulumi.StringPtrInput
	Layout              pulumi.Input
	ManagementGroupName pulumi.StringInput
	Parameters          ParameterDefinitionMapInput
	ResourceGroups      ResourceGroupDefinitionMapInput
	TargetScope         pulumi.StringInput
	Versions            pulumi.Input
}

func (BlueprintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blueprintArgs)(nil)).Elem()
}

type BlueprintInput interface {
	pulumi.Input

	ToBlueprintOutput() BlueprintOutput
	ToBlueprintOutputWithContext(ctx context.Context) BlueprintOutput
}

func (*Blueprint) ElementType() reflect.Type {
	return reflect.TypeOf((*Blueprint)(nil))
}

func (i *Blueprint) ToBlueprintOutput() BlueprintOutput {
	return i.ToBlueprintOutputWithContext(context.Background())
}

func (i *Blueprint) ToBlueprintOutputWithContext(ctx context.Context) BlueprintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlueprintOutput)
}

type BlueprintOutput struct{ *pulumi.OutputState }

func (BlueprintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Blueprint)(nil))
}

func (o BlueprintOutput) ToBlueprintOutput() BlueprintOutput {
	return o
}

func (o BlueprintOutput) ToBlueprintOutputWithContext(ctx context.Context) BlueprintOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BlueprintOutput{})
}
