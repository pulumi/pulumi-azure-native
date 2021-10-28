


package v20191101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServiceTopology struct {
	pulumi.CustomResourceState

	ArtifactSourceId pulumi.StringPtrOutput `pulumi:"artifactSourceId"`
	Location         pulumi.StringOutput    `pulumi:"location"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	Tags             pulumi.StringMapOutput `pulumi:"tags"`
	Type             pulumi.StringOutput    `pulumi:"type"`
}


func NewServiceTopology(ctx *pulumi.Context,
	name string, args *ServiceTopologyArgs, opts ...pulumi.ResourceOption) (*ServiceTopology, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:deploymentmanager/v20191101preview:ServiceTopology"),
		},
		{
			Type: pulumi.String("azure-native:deploymentmanager:ServiceTopology"),
		},
		{
			Type: pulumi.String("azure-nextgen:deploymentmanager:ServiceTopology"),
		},
		{
			Type: pulumi.String("azure-native:deploymentmanager/v20180901preview:ServiceTopology"),
		},
		{
			Type: pulumi.String("azure-nextgen:deploymentmanager/v20180901preview:ServiceTopology"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceTopology
	err := ctx.RegisterResource("azure-native:deploymentmanager/v20191101preview:ServiceTopology", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServiceTopology(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceTopologyState, opts ...pulumi.ResourceOption) (*ServiceTopology, error) {
	var resource ServiceTopology
	err := ctx.ReadResource("azure-native:deploymentmanager/v20191101preview:ServiceTopology", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceTopologyState struct {
}

type ServiceTopologyState struct {
}

func (ServiceTopologyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTopologyState)(nil)).Elem()
}

type serviceTopologyArgs struct {
	ArtifactSourceId    *string           `pulumi:"artifactSourceId"`
	Location            *string           `pulumi:"location"`
	ResourceGroupName   string            `pulumi:"resourceGroupName"`
	ServiceTopologyName *string           `pulumi:"serviceTopologyName"`
	Tags                map[string]string `pulumi:"tags"`
}


type ServiceTopologyArgs struct {
	ArtifactSourceId    pulumi.StringPtrInput
	Location            pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	ServiceTopologyName pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
}

func (ServiceTopologyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTopologyArgs)(nil)).Elem()
}

type ServiceTopologyInput interface {
	pulumi.Input

	ToServiceTopologyOutput() ServiceTopologyOutput
	ToServiceTopologyOutputWithContext(ctx context.Context) ServiceTopologyOutput
}

func (*ServiceTopology) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTopology)(nil))
}

func (i *ServiceTopology) ToServiceTopologyOutput() ServiceTopologyOutput {
	return i.ToServiceTopologyOutputWithContext(context.Background())
}

func (i *ServiceTopology) ToServiceTopologyOutputWithContext(ctx context.Context) ServiceTopologyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTopologyOutput)
}

type ServiceTopologyOutput struct{ *pulumi.OutputState }

func (ServiceTopologyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTopology)(nil))
}

func (o ServiceTopologyOutput) ToServiceTopologyOutput() ServiceTopologyOutput {
	return o
}

func (o ServiceTopologyOutput) ToServiceTopologyOutputWithContext(ctx context.Context) ServiceTopologyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTopologyInput)(nil)).Elem(), &ServiceTopology{})
	pulumi.RegisterOutputType(ServiceTopologyOutput{})
}
