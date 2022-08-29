


package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TemplateArtifact struct {
	pulumi.CustomResourceState

	DependsOn     pulumi.StringArrayOutput        `pulumi:"dependsOn"`
	Description   pulumi.StringPtrOutput          `pulumi:"description"`
	DisplayName   pulumi.StringPtrOutput          `pulumi:"displayName"`
	Kind          pulumi.StringOutput             `pulumi:"kind"`
	Name          pulumi.StringOutput             `pulumi:"name"`
	Parameters    ParameterValueResponseMapOutput `pulumi:"parameters"`
	ResourceGroup pulumi.StringPtrOutput          `pulumi:"resourceGroup"`
	Template      pulumi.AnyOutput                `pulumi:"template"`
	Type          pulumi.StringOutput             `pulumi:"type"`
}


func NewTemplateArtifact(ctx *pulumi.Context,
	name string, args *TemplateArtifactArgs, opts ...pulumi.ResourceOption) (*TemplateArtifact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlueprintName == nil {
		return nil, errors.New("invalid value for required argument 'BlueprintName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	if args.ResourceScope == nil {
		return nil, errors.New("invalid value for required argument 'ResourceScope'")
	}
	if args.Template == nil {
		return nil, errors.New("invalid value for required argument 'Template'")
	}
	args.Kind = pulumi.String("template")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:blueprint:TemplateArtifact"),
		},
	})
	opts = append(opts, aliases)
	var resource TemplateArtifact
	err := ctx.RegisterResource("azure-native:blueprint/v20181101preview:TemplateArtifact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTemplateArtifact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateArtifactState, opts ...pulumi.ResourceOption) (*TemplateArtifact, error) {
	var resource TemplateArtifact
	err := ctx.ReadResource("azure-native:blueprint/v20181101preview:TemplateArtifact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type templateArtifactState struct {
}

type TemplateArtifactState struct {
}

func (TemplateArtifactState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArtifactState)(nil)).Elem()
}

type templateArtifactArgs struct {
	ArtifactName  *string                   `pulumi:"artifactName"`
	BlueprintName string                    `pulumi:"blueprintName"`
	DependsOn     []string                  `pulumi:"dependsOn"`
	Description   *string                   `pulumi:"description"`
	DisplayName   *string                   `pulumi:"displayName"`
	Kind          string                    `pulumi:"kind"`
	Parameters    map[string]ParameterValue `pulumi:"parameters"`
	ResourceGroup *string                   `pulumi:"resourceGroup"`
	ResourceScope string                    `pulumi:"resourceScope"`
	Template      interface{}               `pulumi:"template"`
}


type TemplateArtifactArgs struct {
	ArtifactName  pulumi.StringPtrInput
	BlueprintName pulumi.StringInput
	DependsOn     pulumi.StringArrayInput
	Description   pulumi.StringPtrInput
	DisplayName   pulumi.StringPtrInput
	Kind          pulumi.StringInput
	Parameters    ParameterValueMapInput
	ResourceGroup pulumi.StringPtrInput
	ResourceScope pulumi.StringInput
	Template      pulumi.Input
}

func (TemplateArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArtifactArgs)(nil)).Elem()
}

type TemplateArtifactInput interface {
	pulumi.Input

	ToTemplateArtifactOutput() TemplateArtifactOutput
	ToTemplateArtifactOutputWithContext(ctx context.Context) TemplateArtifactOutput
}

func (*TemplateArtifact) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateArtifact)(nil)).Elem()
}

func (i *TemplateArtifact) ToTemplateArtifactOutput() TemplateArtifactOutput {
	return i.ToTemplateArtifactOutputWithContext(context.Background())
}

func (i *TemplateArtifact) ToTemplateArtifactOutputWithContext(ctx context.Context) TemplateArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateArtifactOutput)
}

type TemplateArtifactOutput struct{ *pulumi.OutputState }

func (TemplateArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateArtifact)(nil)).Elem()
}

func (o TemplateArtifactOutput) ToTemplateArtifactOutput() TemplateArtifactOutput {
	return o
}

func (o TemplateArtifactOutput) ToTemplateArtifactOutputWithContext(ctx context.Context) TemplateArtifactOutput {
	return o
}

func (o TemplateArtifactOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringArrayOutput { return v.DependsOn }).(pulumi.StringArrayOutput)
}

func (o TemplateArtifactOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TemplateArtifactOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o TemplateArtifactOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o TemplateArtifactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TemplateArtifactOutput) Parameters() ParameterValueResponseMapOutput {
	return o.ApplyT(func(v *TemplateArtifact) ParameterValueResponseMapOutput { return v.Parameters }).(ParameterValueResponseMapOutput)
}

func (o TemplateArtifactOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringPtrOutput { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o TemplateArtifactOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.AnyOutput { return v.Template }).(pulumi.AnyOutput)
}

func (o TemplateArtifactOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateArtifact) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TemplateArtifactOutput{})
}
