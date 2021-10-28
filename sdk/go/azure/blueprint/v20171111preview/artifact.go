


package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Artifact struct {
	pulumi.CustomResourceState

	Kind pulumi.StringOutput `pulumi:"kind"`
	Name pulumi.StringOutput `pulumi:"name"`
	Type pulumi.StringOutput `pulumi:"type"`
}


func NewArtifact(ctx *pulumi.Context,
	name string, args *ArtifactArgs, opts ...pulumi.ResourceOption) (*Artifact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlueprintName == nil {
		return nil, errors.New("invalid value for required argument 'BlueprintName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ManagementGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:blueprint/v20171111preview:Artifact"),
		},
	})
	opts = append(opts, aliases)
	var resource Artifact
	err := ctx.RegisterResource("azure-native:blueprint/v20171111preview:Artifact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetArtifact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactState, opts ...pulumi.ResourceOption) (*Artifact, error) {
	var resource Artifact
	err := ctx.ReadResource("azure-native:blueprint/v20171111preview:Artifact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type artifactState struct {
}

type ArtifactState struct {
}

func (ArtifactState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactState)(nil)).Elem()
}

type artifactArgs struct {
	ArtifactName        *string `pulumi:"artifactName"`
	BlueprintName       string  `pulumi:"blueprintName"`
	Kind                string  `pulumi:"kind"`
	ManagementGroupName string  `pulumi:"managementGroupName"`
}


type ArtifactArgs struct {
	ArtifactName        pulumi.StringPtrInput
	BlueprintName       pulumi.StringInput
	Kind                pulumi.StringInput
	ManagementGroupName pulumi.StringInput
}

func (ArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactArgs)(nil)).Elem()
}

type ArtifactInput interface {
	pulumi.Input

	ToArtifactOutput() ArtifactOutput
	ToArtifactOutputWithContext(ctx context.Context) ArtifactOutput
}

func (*Artifact) ElementType() reflect.Type {
	return reflect.TypeOf((*Artifact)(nil))
}

func (i *Artifact) ToArtifactOutput() ArtifactOutput {
	return i.ToArtifactOutputWithContext(context.Background())
}

func (i *Artifact) ToArtifactOutputWithContext(ctx context.Context) ArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactOutput)
}

type ArtifactOutput struct{ *pulumi.OutputState }

func (ArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Artifact)(nil))
}

func (o ArtifactOutput) ToArtifactOutput() ArtifactOutput {
	return o
}

func (o ArtifactOutput) ToArtifactOutputWithContext(ctx context.Context) ArtifactOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactInput)(nil)).Elem(), &Artifact{})
	pulumi.RegisterOutputType(ArtifactOutput{})
}
