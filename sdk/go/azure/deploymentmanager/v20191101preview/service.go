


package v20191101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Service struct {
	pulumi.CustomResourceState

	Location             pulumi.StringOutput    `pulumi:"location"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	Tags                 pulumi.StringMapOutput `pulumi:"tags"`
	TargetLocation       pulumi.StringOutput    `pulumi:"targetLocation"`
	TargetSubscriptionId pulumi.StringOutput    `pulumi:"targetSubscriptionId"`
	Type                 pulumi.StringOutput    `pulumi:"type"`
}


func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceTopologyName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceTopologyName'")
	}
	if args.TargetLocation == nil {
		return nil, errors.New("invalid value for required argument 'TargetLocation'")
	}
	if args.TargetSubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'TargetSubscriptionId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:deploymentmanager:Service"),
		},
		{
			Type: pulumi.String("azure-native:deploymentmanager/v20180901preview:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azure-native:deploymentmanager/v20191101preview:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:deploymentmanager/v20191101preview:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	Location             *string           `pulumi:"location"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	ServiceName          *string           `pulumi:"serviceName"`
	ServiceTopologyName  string            `pulumi:"serviceTopologyName"`
	Tags                 map[string]string `pulumi:"tags"`
	TargetLocation       string            `pulumi:"targetLocation"`
	TargetSubscriptionId string            `pulumi:"targetSubscriptionId"`
}


type ServiceArgs struct {
	Location             pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	ServiceName          pulumi.StringPtrInput
	ServiceTopologyName  pulumi.StringInput
	Tags                 pulumi.StringMapInput
	TargetLocation       pulumi.StringInput
	TargetSubscriptionId pulumi.StringInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
