


package blueprint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PublishedBlueprint struct {
	pulumi.CustomResourceState

	BlueprintName  pulumi.StringPtrOutput                   `pulumi:"blueprintName"`
	ChangeNotes    pulumi.StringPtrOutput                   `pulumi:"changeNotes"`
	Description    pulumi.StringPtrOutput                   `pulumi:"description"`
	DisplayName    pulumi.StringPtrOutput                   `pulumi:"displayName"`
	Name           pulumi.StringOutput                      `pulumi:"name"`
	Parameters     ParameterDefinitionResponseMapOutput     `pulumi:"parameters"`
	ResourceGroups ResourceGroupDefinitionResponseMapOutput `pulumi:"resourceGroups"`
	Status         BlueprintStatusResponseOutput            `pulumi:"status"`
	TargetScope    pulumi.StringPtrOutput                   `pulumi:"targetScope"`
	Type           pulumi.StringOutput                      `pulumi:"type"`
}


func NewPublishedBlueprint(ctx *pulumi.Context,
	name string, args *PublishedBlueprintArgs, opts ...pulumi.ResourceOption) (*PublishedBlueprint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlueprintName == nil {
		return nil, errors.New("invalid value for required argument 'BlueprintName'")
	}
	if args.ResourceScope == nil {
		return nil, errors.New("invalid value for required argument 'ResourceScope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:blueprint:PublishedBlueprint"),
		},
		{
			Type: pulumi.String("azure-native:blueprint/v20181101preview:PublishedBlueprint"),
		},
		{
			Type: pulumi.String("azure-nextgen:blueprint/v20181101preview:PublishedBlueprint"),
		},
	})
	opts = append(opts, aliases)
	var resource PublishedBlueprint
	err := ctx.RegisterResource("azure-native:blueprint:PublishedBlueprint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPublishedBlueprint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublishedBlueprintState, opts ...pulumi.ResourceOption) (*PublishedBlueprint, error) {
	var resource PublishedBlueprint
	err := ctx.ReadResource("azure-native:blueprint:PublishedBlueprint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type publishedBlueprintState struct {
}

type PublishedBlueprintState struct {
}

func (PublishedBlueprintState) ElementType() reflect.Type {
	return reflect.TypeOf((*publishedBlueprintState)(nil)).Elem()
}

type publishedBlueprintArgs struct {
	BlueprintName  string                             `pulumi:"blueprintName"`
	ChangeNotes    *string                            `pulumi:"changeNotes"`
	Description    *string                            `pulumi:"description"`
	DisplayName    *string                            `pulumi:"displayName"`
	Parameters     map[string]ParameterDefinition     `pulumi:"parameters"`
	ResourceGroups map[string]ResourceGroupDefinition `pulumi:"resourceGroups"`
	ResourceScope  string                             `pulumi:"resourceScope"`
	TargetScope    *string                            `pulumi:"targetScope"`
	VersionId      *string                            `pulumi:"versionId"`
}


type PublishedBlueprintArgs struct {
	BlueprintName  pulumi.StringInput
	ChangeNotes    pulumi.StringPtrInput
	Description    pulumi.StringPtrInput
	DisplayName    pulumi.StringPtrInput
	Parameters     ParameterDefinitionMapInput
	ResourceGroups ResourceGroupDefinitionMapInput
	ResourceScope  pulumi.StringInput
	TargetScope    pulumi.StringPtrInput
	VersionId      pulumi.StringPtrInput
}

func (PublishedBlueprintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publishedBlueprintArgs)(nil)).Elem()
}

type PublishedBlueprintInput interface {
	pulumi.Input

	ToPublishedBlueprintOutput() PublishedBlueprintOutput
	ToPublishedBlueprintOutputWithContext(ctx context.Context) PublishedBlueprintOutput
}

func (*PublishedBlueprint) ElementType() reflect.Type {
	return reflect.TypeOf((*PublishedBlueprint)(nil))
}

func (i *PublishedBlueprint) ToPublishedBlueprintOutput() PublishedBlueprintOutput {
	return i.ToPublishedBlueprintOutputWithContext(context.Background())
}

func (i *PublishedBlueprint) ToPublishedBlueprintOutputWithContext(ctx context.Context) PublishedBlueprintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublishedBlueprintOutput)
}

type PublishedBlueprintOutput struct{ *pulumi.OutputState }

func (PublishedBlueprintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublishedBlueprint)(nil))
}

func (o PublishedBlueprintOutput) ToPublishedBlueprintOutput() PublishedBlueprintOutput {
	return o
}

func (o PublishedBlueprintOutput) ToPublishedBlueprintOutputWithContext(ctx context.Context) PublishedBlueprintOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublishedBlueprintInput)(nil)).Elem(), &PublishedBlueprint{})
	pulumi.RegisterOutputType(PublishedBlueprintOutput{})
}
