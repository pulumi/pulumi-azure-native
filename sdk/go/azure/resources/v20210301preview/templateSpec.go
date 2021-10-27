


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TemplateSpec struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput                   `pulumi:"description"`
	DisplayName pulumi.StringPtrOutput                   `pulumi:"displayName"`
	Location    pulumi.StringOutput                      `pulumi:"location"`
	Metadata    pulumi.AnyOutput                         `pulumi:"metadata"`
	Name        pulumi.StringOutput                      `pulumi:"name"`
	SystemData  SystemDataResponseOutput                 `pulumi:"systemData"`
	Tags        pulumi.StringMapOutput                   `pulumi:"tags"`
	Type        pulumi.StringOutput                      `pulumi:"type"`
	Versions    TemplateSpecVersionInfoResponseMapOutput `pulumi:"versions"`
}


func NewTemplateSpec(ctx *pulumi.Context,
	name string, args *TemplateSpecArgs, opts ...pulumi.ResourceOption) (*TemplateSpec, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:resources/v20210301preview:TemplateSpec"),
		},
		{
			Type: pulumi.String("azure-native:resources:TemplateSpec"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources:TemplateSpec"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190601preview:TemplateSpec"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20190601preview:TemplateSpec"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210501:TemplateSpec"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20210501:TemplateSpec"),
		},
	})
	opts = append(opts, aliases)
	var resource TemplateSpec
	err := ctx.RegisterResource("azure-native:resources/v20210301preview:TemplateSpec", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTemplateSpec(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateSpecState, opts ...pulumi.ResourceOption) (*TemplateSpec, error) {
	var resource TemplateSpec
	err := ctx.ReadResource("azure-native:resources/v20210301preview:TemplateSpec", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type templateSpecState struct {
}

type TemplateSpecState struct {
}

func (TemplateSpecState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateSpecState)(nil)).Elem()
}

type templateSpecArgs struct {
	Description       *string           `pulumi:"description"`
	DisplayName       *string           `pulumi:"displayName"`
	Location          *string           `pulumi:"location"`
	Metadata          interface{}       `pulumi:"metadata"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	TemplateSpecName  *string           `pulumi:"templateSpecName"`
}


type TemplateSpecArgs struct {
	Description       pulumi.StringPtrInput
	DisplayName       pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Metadata          pulumi.Input
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	TemplateSpecName  pulumi.StringPtrInput
}

func (TemplateSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateSpecArgs)(nil)).Elem()
}

type TemplateSpecInput interface {
	pulumi.Input

	ToTemplateSpecOutput() TemplateSpecOutput
	ToTemplateSpecOutputWithContext(ctx context.Context) TemplateSpecOutput
}

func (*TemplateSpec) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpec)(nil))
}

func (i *TemplateSpec) ToTemplateSpecOutput() TemplateSpecOutput {
	return i.ToTemplateSpecOutputWithContext(context.Background())
}

func (i *TemplateSpec) ToTemplateSpecOutputWithContext(ctx context.Context) TemplateSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateSpecOutput)
}

type TemplateSpecOutput struct{ *pulumi.OutputState }

func (TemplateSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpec)(nil))
}

func (o TemplateSpecOutput) ToTemplateSpecOutput() TemplateSpecOutput {
	return o
}

func (o TemplateSpecOutput) ToTemplateSpecOutputWithContext(ctx context.Context) TemplateSpecOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateSpecInput)(nil)).Elem(), &TemplateSpec{})
	pulumi.RegisterOutputType(TemplateSpecOutput{})
}
