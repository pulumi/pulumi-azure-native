


package v20190701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Pipeline struct {
	pulumi.CustomResourceState

	BootstrapConfiguration BootstrapConfigurationResponseOutput `pulumi:"bootstrapConfiguration"`
	Location               pulumi.StringPtrOutput               `pulumi:"location"`
	Name                   pulumi.StringOutput                  `pulumi:"name"`
	Organization           OrganizationReferenceResponseOutput  `pulumi:"organization"`
	PipelineId             pulumi.IntOutput                     `pulumi:"pipelineId"`
	Project                ProjectReferenceResponseOutput       `pulumi:"project"`
	Tags                   pulumi.StringMapOutput               `pulumi:"tags"`
	Type                   pulumi.StringOutput                  `pulumi:"type"`
}


func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BootstrapConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'BootstrapConfiguration'")
	}
	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devops/v20190701preview:Pipeline"),
		},
		{
			Type: pulumi.String("azure-native:devops:Pipeline"),
		},
		{
			Type: pulumi.String("azure-nextgen:devops:Pipeline"),
		},
		{
			Type: pulumi.String("azure-native:devops/v20200713preview:Pipeline"),
		},
		{
			Type: pulumi.String("azure-nextgen:devops/v20200713preview:Pipeline"),
		},
	})
	opts = append(opts, aliases)
	var resource Pipeline
	err := ctx.RegisterResource("azure-native:devops/v20190701preview:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("azure-native:devops/v20190701preview:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type pipelineState struct {
}

type PipelineState struct {
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	BootstrapConfiguration BootstrapConfiguration `pulumi:"bootstrapConfiguration"`
	Location               *string                `pulumi:"location"`
	Organization           OrganizationReference  `pulumi:"organization"`
	PipelineName           *string                `pulumi:"pipelineName"`
	Project                ProjectReference       `pulumi:"project"`
	ResourceGroupName      string                 `pulumi:"resourceGroupName"`
	Tags                   map[string]string      `pulumi:"tags"`
}


type PipelineArgs struct {
	BootstrapConfiguration BootstrapConfigurationInput
	Location               pulumi.StringPtrInput
	Organization           OrganizationReferenceInput
	PipelineName           pulumi.StringPtrInput
	Project                ProjectReferenceInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}

type PipelineInput interface {
	pulumi.Input

	ToPipelineOutput() PipelineOutput
	ToPipelineOutputWithContext(ctx context.Context) PipelineOutput
}

func (*Pipeline) ElementType() reflect.Type {
	return reflect.TypeOf((*Pipeline)(nil))
}

func (i *Pipeline) ToPipelineOutput() PipelineOutput {
	return i.ToPipelineOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOutput)
}

type PipelineOutput struct{ *pulumi.OutputState }

func (PipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pipeline)(nil))
}

func (o PipelineOutput) ToPipelineOutput() PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineInput)(nil)).Elem(), &Pipeline{})
	pulumi.RegisterOutputType(PipelineOutput{})
}
