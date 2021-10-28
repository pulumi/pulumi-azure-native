


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServiceFabric struct {
	pulumi.CustomResourceState

	ApplicableSchedule      ApplicableScheduleResponseOutput `pulumi:"applicableSchedule"`
	EnvironmentId           pulumi.StringPtrOutput           `pulumi:"environmentId"`
	ExternalServiceFabricId pulumi.StringPtrOutput           `pulumi:"externalServiceFabricId"`
	Location                pulumi.StringPtrOutput           `pulumi:"location"`
	Name                    pulumi.StringOutput              `pulumi:"name"`
	ProvisioningState       pulumi.StringOutput              `pulumi:"provisioningState"`
	Tags                    pulumi.StringMapOutput           `pulumi:"tags"`
	Type                    pulumi.StringOutput              `pulumi:"type"`
	UniqueIdentifier        pulumi.StringOutput              `pulumi:"uniqueIdentifier"`
}


func NewServiceFabric(ctx *pulumi.Context,
	name string, args *ServiceFabricArgs, opts ...pulumi.ResourceOption) (*ServiceFabric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devtestlab:ServiceFabric"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:ServiceFabric"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20180915:ServiceFabric"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceFabric
	err := ctx.RegisterResource("azure-native:devtestlab:ServiceFabric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServiceFabric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceFabricState, opts ...pulumi.ResourceOption) (*ServiceFabric, error) {
	var resource ServiceFabric
	err := ctx.ReadResource("azure-native:devtestlab:ServiceFabric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceFabricState struct {
}

type ServiceFabricState struct {
}

func (ServiceFabricState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceFabricState)(nil)).Elem()
}

type serviceFabricArgs struct {
	EnvironmentId           *string           `pulumi:"environmentId"`
	ExternalServiceFabricId *string           `pulumi:"externalServiceFabricId"`
	LabName                 string            `pulumi:"labName"`
	Location                *string           `pulumi:"location"`
	Name                    *string           `pulumi:"name"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
	Tags                    map[string]string `pulumi:"tags"`
	UserName                string            `pulumi:"userName"`
}


type ServiceFabricArgs struct {
	EnvironmentId           pulumi.StringPtrInput
	ExternalServiceFabricId pulumi.StringPtrInput
	LabName                 pulumi.StringInput
	Location                pulumi.StringPtrInput
	Name                    pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	Tags                    pulumi.StringMapInput
	UserName                pulumi.StringInput
}

func (ServiceFabricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceFabricArgs)(nil)).Elem()
}

type ServiceFabricInput interface {
	pulumi.Input

	ToServiceFabricOutput() ServiceFabricOutput
	ToServiceFabricOutputWithContext(ctx context.Context) ServiceFabricOutput
}

func (*ServiceFabric) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceFabric)(nil))
}

func (i *ServiceFabric) ToServiceFabricOutput() ServiceFabricOutput {
	return i.ToServiceFabricOutputWithContext(context.Background())
}

func (i *ServiceFabric) ToServiceFabricOutputWithContext(ctx context.Context) ServiceFabricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceFabricOutput)
}

type ServiceFabricOutput struct{ *pulumi.OutputState }

func (ServiceFabricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceFabric)(nil))
}

func (o ServiceFabricOutput) ToServiceFabricOutput() ServiceFabricOutput {
	return o
}

func (o ServiceFabricOutput) ToServiceFabricOutputWithContext(ctx context.Context) ServiceFabricOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceFabricOutput{})
}
