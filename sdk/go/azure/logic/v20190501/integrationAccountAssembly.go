


package v20190501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccountAssembly struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput           `pulumi:"location"`
	Name       pulumi.StringOutput              `pulumi:"name"`
	Properties AssemblyPropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput           `pulumi:"tags"`
	Type       pulumi.StringOutput              `pulumi:"type"`
}


func NewIntegrationAccountAssembly(ctx *pulumi.Context,
	name string, args *IntegrationAccountAssemblyArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountAssembly, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic:IntegrationAccountAssembly"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountAssembly"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountAssembly"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountAssembly
	err := ctx.RegisterResource("azure-native:logic/v20190501:IntegrationAccountAssembly", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationAccountAssembly(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountAssemblyState, opts ...pulumi.ResourceOption) (*IntegrationAccountAssembly, error) {
	var resource IntegrationAccountAssembly
	err := ctx.ReadResource("azure-native:logic/v20190501:IntegrationAccountAssembly", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationAccountAssemblyState struct {
}

type IntegrationAccountAssemblyState struct {
}

func (IntegrationAccountAssemblyState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountAssemblyState)(nil)).Elem()
}

type integrationAccountAssemblyArgs struct {
	AssemblyArtifactName   *string            `pulumi:"assemblyArtifactName"`
	IntegrationAccountName string             `pulumi:"integrationAccountName"`
	Location               *string            `pulumi:"location"`
	Properties             AssemblyProperties `pulumi:"properties"`
	ResourceGroupName      string             `pulumi:"resourceGroupName"`
	Tags                   map[string]string  `pulumi:"tags"`
}


type IntegrationAccountAssemblyArgs struct {
	AssemblyArtifactName   pulumi.StringPtrInput
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	Properties             AssemblyPropertiesInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (IntegrationAccountAssemblyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountAssemblyArgs)(nil)).Elem()
}

type IntegrationAccountAssemblyInput interface {
	pulumi.Input

	ToIntegrationAccountAssemblyOutput() IntegrationAccountAssemblyOutput
	ToIntegrationAccountAssemblyOutputWithContext(ctx context.Context) IntegrationAccountAssemblyOutput
}

func (*IntegrationAccountAssembly) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountAssembly)(nil)).Elem()
}

func (i *IntegrationAccountAssembly) ToIntegrationAccountAssemblyOutput() IntegrationAccountAssemblyOutput {
	return i.ToIntegrationAccountAssemblyOutputWithContext(context.Background())
}

func (i *IntegrationAccountAssembly) ToIntegrationAccountAssemblyOutputWithContext(ctx context.Context) IntegrationAccountAssemblyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountAssemblyOutput)
}

type IntegrationAccountAssemblyOutput struct{ *pulumi.OutputState }

func (IntegrationAccountAssemblyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountAssembly)(nil)).Elem()
}

func (o IntegrationAccountAssemblyOutput) ToIntegrationAccountAssemblyOutput() IntegrationAccountAssemblyOutput {
	return o
}

func (o IntegrationAccountAssemblyOutput) ToIntegrationAccountAssemblyOutputWithContext(ctx context.Context) IntegrationAccountAssemblyOutput {
	return o
}

func (o IntegrationAccountAssemblyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountAssembly) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountAssemblyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountAssembly) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IntegrationAccountAssemblyOutput) Properties() AssemblyPropertiesResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountAssembly) AssemblyPropertiesResponseOutput { return v.Properties }).(AssemblyPropertiesResponseOutput)
}

func (o IntegrationAccountAssemblyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationAccountAssembly) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o IntegrationAccountAssemblyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountAssembly) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountAssemblyOutput{})
}
