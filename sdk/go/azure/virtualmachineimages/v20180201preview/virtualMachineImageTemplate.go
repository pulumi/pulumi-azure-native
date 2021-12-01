


package v20180201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineImageTemplate struct {
	pulumi.CustomResourceState

	Customize         ImageTemplateShellCustomizerResponseArrayOutput `pulumi:"customize"`
	Distribute        pulumi.ArrayOutput                              `pulumi:"distribute"`
	LastRunStatus     ImageTemplateLastRunStatusResponseOutput        `pulumi:"lastRunStatus"`
	Location          pulumi.StringOutput                             `pulumi:"location"`
	Name              pulumi.StringOutput                             `pulumi:"name"`
	ProvisioningError ProvisioningErrorResponseOutput                 `pulumi:"provisioningError"`
	ProvisioningState pulumi.StringOutput                             `pulumi:"provisioningState"`
	Source            pulumi.AnyOutput                                `pulumi:"source"`
	Tags              pulumi.StringMapOutput                          `pulumi:"tags"`
	Type              pulumi.StringOutput                             `pulumi:"type"`
}


func NewVirtualMachineImageTemplate(ctx *pulumi.Context,
	name string, args *VirtualMachineImageTemplateArgs, opts ...pulumi.ResourceOption) (*VirtualMachineImageTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Distribute == nil {
		return nil, errors.New("invalid value for required argument 'Distribute'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:virtualmachineimages:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20190201preview:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20190501preview:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20200214:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20211001:VirtualMachineImageTemplate"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineImageTemplate
	err := ctx.RegisterResource("azure-native:virtualmachineimages/v20180201preview:VirtualMachineImageTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineImageTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineImageTemplateState, opts ...pulumi.ResourceOption) (*VirtualMachineImageTemplate, error) {
	var resource VirtualMachineImageTemplate
	err := ctx.ReadResource("azure-native:virtualmachineimages/v20180201preview:VirtualMachineImageTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineImageTemplateState struct {
}

type VirtualMachineImageTemplateState struct {
}

func (VirtualMachineImageTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineImageTemplateState)(nil)).Elem()
}

type virtualMachineImageTemplateArgs struct {
	Customize         []ImageTemplateShellCustomizer `pulumi:"customize"`
	Distribute        []interface{}                  `pulumi:"distribute"`
	ImageTemplateName *string                        `pulumi:"imageTemplateName"`
	Location          *string                        `pulumi:"location"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	Source            interface{}                    `pulumi:"source"`
	Tags              map[string]string              `pulumi:"tags"`
}


type VirtualMachineImageTemplateArgs struct {
	Customize         ImageTemplateShellCustomizerArrayInput
	Distribute        pulumi.ArrayInput
	ImageTemplateName pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Source            pulumi.Input
	Tags              pulumi.StringMapInput
}

func (VirtualMachineImageTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineImageTemplateArgs)(nil)).Elem()
}

type VirtualMachineImageTemplateInput interface {
	pulumi.Input

	ToVirtualMachineImageTemplateOutput() VirtualMachineImageTemplateOutput
	ToVirtualMachineImageTemplateOutputWithContext(ctx context.Context) VirtualMachineImageTemplateOutput
}

func (*VirtualMachineImageTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineImageTemplate)(nil))
}

func (i *VirtualMachineImageTemplate) ToVirtualMachineImageTemplateOutput() VirtualMachineImageTemplateOutput {
	return i.ToVirtualMachineImageTemplateOutputWithContext(context.Background())
}

func (i *VirtualMachineImageTemplate) ToVirtualMachineImageTemplateOutputWithContext(ctx context.Context) VirtualMachineImageTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImageTemplateOutput)
}

type VirtualMachineImageTemplateOutput struct{ *pulumi.OutputState }

func (VirtualMachineImageTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineImageTemplate)(nil))
}

func (o VirtualMachineImageTemplateOutput) ToVirtualMachineImageTemplateOutput() VirtualMachineImageTemplateOutput {
	return o
}

func (o VirtualMachineImageTemplateOutput) ToVirtualMachineImageTemplateOutputWithContext(ctx context.Context) VirtualMachineImageTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineImageTemplateOutput{})
}
