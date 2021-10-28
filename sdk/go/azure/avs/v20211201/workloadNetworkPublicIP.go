


package v20211201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadNetworkPublicIP struct {
	pulumi.CustomResourceState

	DisplayName       pulumi.StringPtrOutput  `pulumi:"displayName"`
	Name              pulumi.StringOutput     `pulumi:"name"`
	NumberOfPublicIPs pulumi.Float64PtrOutput `pulumi:"numberOfPublicIPs"`
	ProvisioningState pulumi.StringOutput     `pulumi:"provisioningState"`
	PublicIPBlock     pulumi.StringOutput     `pulumi:"publicIPBlock"`
	Type              pulumi.StringOutput     `pulumi:"type"`
}


func NewWorkloadNetworkPublicIP(ctx *pulumi.Context,
	name string, args *WorkloadNetworkPublicIPArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkPublicIP, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:avs/v20211201:WorkloadNetworkPublicIP"),
		},
		{
			Type: pulumi.String("azure-native:avs:WorkloadNetworkPublicIP"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs:WorkloadNetworkPublicIP"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkPublicIP"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20210601:WorkloadNetworkPublicIP"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkPublicIP
	err := ctx.RegisterResource("azure-native:avs/v20211201:WorkloadNetworkPublicIP", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkloadNetworkPublicIP(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkPublicIPState, opts ...pulumi.ResourceOption) (*WorkloadNetworkPublicIP, error) {
	var resource WorkloadNetworkPublicIP
	err := ctx.ReadResource("azure-native:avs/v20211201:WorkloadNetworkPublicIP", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workloadNetworkPublicIPState struct {
}

type WorkloadNetworkPublicIPState struct {
}

func (WorkloadNetworkPublicIPState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkPublicIPState)(nil)).Elem()
}

type workloadNetworkPublicIPArgs struct {
	DisplayName       *string  `pulumi:"displayName"`
	NumberOfPublicIPs *float64 `pulumi:"numberOfPublicIPs"`
	PrivateCloudName  string   `pulumi:"privateCloudName"`
	PublicIPId        *string  `pulumi:"publicIPId"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
}


type WorkloadNetworkPublicIPArgs struct {
	DisplayName       pulumi.StringPtrInput
	NumberOfPublicIPs pulumi.Float64PtrInput
	PrivateCloudName  pulumi.StringInput
	PublicIPId        pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (WorkloadNetworkPublicIPArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkPublicIPArgs)(nil)).Elem()
}

type WorkloadNetworkPublicIPInput interface {
	pulumi.Input

	ToWorkloadNetworkPublicIPOutput() WorkloadNetworkPublicIPOutput
	ToWorkloadNetworkPublicIPOutputWithContext(ctx context.Context) WorkloadNetworkPublicIPOutput
}

func (*WorkloadNetworkPublicIP) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkPublicIP)(nil))
}

func (i *WorkloadNetworkPublicIP) ToWorkloadNetworkPublicIPOutput() WorkloadNetworkPublicIPOutput {
	return i.ToWorkloadNetworkPublicIPOutputWithContext(context.Background())
}

func (i *WorkloadNetworkPublicIP) ToWorkloadNetworkPublicIPOutputWithContext(ctx context.Context) WorkloadNetworkPublicIPOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkPublicIPOutput)
}

type WorkloadNetworkPublicIPOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkPublicIPOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkPublicIP)(nil))
}

func (o WorkloadNetworkPublicIPOutput) ToWorkloadNetworkPublicIPOutput() WorkloadNetworkPublicIPOutput {
	return o
}

func (o WorkloadNetworkPublicIPOutput) ToWorkloadNetworkPublicIPOutputWithContext(ctx context.Context) WorkloadNetworkPublicIPOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkPublicIPOutput{})
}
