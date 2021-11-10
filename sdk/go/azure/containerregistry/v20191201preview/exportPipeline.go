


package v20191201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExportPipeline struct {
	pulumi.CustomResourceState

	Identity          IdentityPropertiesResponsePtrOutput          `pulumi:"identity"`
	Location          pulumi.StringPtrOutput                       `pulumi:"location"`
	Name              pulumi.StringOutput                          `pulumi:"name"`
	Options           pulumi.StringArrayOutput                     `pulumi:"options"`
	ProvisioningState pulumi.StringOutput                          `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput                     `pulumi:"systemData"`
	Target            ExportPipelineTargetPropertiesResponseOutput `pulumi:"target"`
	Type              pulumi.StringOutput                          `pulumi:"type"`
}


func NewExportPipeline(ctx *pulumi.Context,
	name string, args *ExportPipelineArgs, opts ...pulumi.ResourceOption) (*ExportPipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:ExportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:ExportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:ExportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:ExportPipeline"),
		},
	})
	opts = append(opts, aliases)
	var resource ExportPipeline
	err := ctx.RegisterResource("azure-native:containerregistry/v20191201preview:ExportPipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExportPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExportPipelineState, opts ...pulumi.ResourceOption) (*ExportPipeline, error) {
	var resource ExportPipeline
	err := ctx.ReadResource("azure-native:containerregistry/v20191201preview:ExportPipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type exportPipelineState struct {
}

type ExportPipelineState struct {
}

func (ExportPipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*exportPipelineState)(nil)).Elem()
}

type exportPipelineArgs struct {
	ExportPipelineName *string                        `pulumi:"exportPipelineName"`
	Identity           *IdentityProperties            `pulumi:"identity"`
	Location           *string                        `pulumi:"location"`
	Options            []string                       `pulumi:"options"`
	RegistryName       string                         `pulumi:"registryName"`
	ResourceGroupName  string                         `pulumi:"resourceGroupName"`
	Target             ExportPipelineTargetProperties `pulumi:"target"`
}


type ExportPipelineArgs struct {
	ExportPipelineName pulumi.StringPtrInput
	Identity           IdentityPropertiesPtrInput
	Location           pulumi.StringPtrInput
	Options            pulumi.StringArrayInput
	RegistryName       pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	Target             ExportPipelineTargetPropertiesInput
}

func (ExportPipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*exportPipelineArgs)(nil)).Elem()
}

type ExportPipelineInput interface {
	pulumi.Input

	ToExportPipelineOutput() ExportPipelineOutput
	ToExportPipelineOutputWithContext(ctx context.Context) ExportPipelineOutput
}

func (*ExportPipeline) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPipeline)(nil))
}

func (i *ExportPipeline) ToExportPipelineOutput() ExportPipelineOutput {
	return i.ToExportPipelineOutputWithContext(context.Background())
}

func (i *ExportPipeline) ToExportPipelineOutputWithContext(ctx context.Context) ExportPipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPipelineOutput)
}

type ExportPipelineOutput struct{ *pulumi.OutputState }

func (ExportPipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPipeline)(nil))
}

func (o ExportPipelineOutput) ToExportPipelineOutput() ExportPipelineOutput {
	return o
}

func (o ExportPipelineOutput) ToExportPipelineOutputWithContext(ctx context.Context) ExportPipelineOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExportPipelineOutput{})
}
