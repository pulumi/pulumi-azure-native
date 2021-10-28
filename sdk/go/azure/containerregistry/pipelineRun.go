


package containerregistry

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PipelineRun struct {
	pulumi.CustomResourceState

	ForceUpdateTag    pulumi.StringPtrOutput              `pulumi:"forceUpdateTag"`
	Name              pulumi.StringOutput                 `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                 `pulumi:"provisioningState"`
	Request           PipelineRunRequestResponsePtrOutput `pulumi:"request"`
	Response          PipelineRunResponseResponseOutput   `pulumi:"response"`
	SystemData        SystemDataResponseOutput            `pulumi:"systemData"`
	Type              pulumi.StringOutput                 `pulumi:"type"`
}


func NewPipelineRun(ctx *pulumi.Context,
	name string, args *PipelineRunArgs, opts ...pulumi.ResourceOption) (*PipelineRun, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:containerregistry:PipelineRun"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20191201preview:PipelineRun"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20191201preview:PipelineRun"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:PipelineRun"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20201101preview:PipelineRun"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:PipelineRun"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20210601preview:PipelineRun"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:PipelineRun"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20210801preview:PipelineRun"),
		},
	})
	opts = append(opts, aliases)
	var resource PipelineRun
	err := ctx.RegisterResource("azure-native:containerregistry:PipelineRun", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPipelineRun(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineRunState, opts ...pulumi.ResourceOption) (*PipelineRun, error) {
	var resource PipelineRun
	err := ctx.ReadResource("azure-native:containerregistry:PipelineRun", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type pipelineRunState struct {
}

type PipelineRunState struct {
}

func (PipelineRunState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineRunState)(nil)).Elem()
}

type pipelineRunArgs struct {
	ForceUpdateTag    *string             `pulumi:"forceUpdateTag"`
	PipelineRunName   *string             `pulumi:"pipelineRunName"`
	RegistryName      string              `pulumi:"registryName"`
	Request           *PipelineRunRequest `pulumi:"request"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
}


type PipelineRunArgs struct {
	ForceUpdateTag    pulumi.StringPtrInput
	PipelineRunName   pulumi.StringPtrInput
	RegistryName      pulumi.StringInput
	Request           PipelineRunRequestPtrInput
	ResourceGroupName pulumi.StringInput
}

func (PipelineRunArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineRunArgs)(nil)).Elem()
}

type PipelineRunInput interface {
	pulumi.Input

	ToPipelineRunOutput() PipelineRunOutput
	ToPipelineRunOutputWithContext(ctx context.Context) PipelineRunOutput
}

func (*PipelineRun) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRun)(nil))
}

func (i *PipelineRun) ToPipelineRunOutput() PipelineRunOutput {
	return i.ToPipelineRunOutputWithContext(context.Background())
}

func (i *PipelineRun) ToPipelineRunOutputWithContext(ctx context.Context) PipelineRunOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunOutput)
}

type PipelineRunOutput struct{ *pulumi.OutputState }

func (PipelineRunOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRun)(nil))
}

func (o PipelineRunOutput) ToPipelineRunOutput() PipelineRunOutput {
	return o
}

func (o PipelineRunOutput) ToPipelineRunOutputWithContext(ctx context.Context) PipelineRunOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PipelineRunOutput{})
}
