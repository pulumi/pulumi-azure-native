


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CapacityDetails struct {
	pulumi.CustomResourceState

	Administration    DedicatedCapacityAdministratorsResponsePtrOutput `pulumi:"administration"`
	Location          pulumi.StringOutput                              `pulumi:"location"`
	Mode              pulumi.StringPtrOutput                           `pulumi:"mode"`
	Name              pulumi.StringOutput                              `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                              `pulumi:"provisioningState"`
	Sku               CapacitySkuResponseOutput                        `pulumi:"sku"`
	State             pulumi.StringOutput                              `pulumi:"state"`
	SystemData        SystemDataResponsePtrOutput                      `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                           `pulumi:"tags"`
	Type              pulumi.StringOutput                              `pulumi:"type"`
}


func NewCapacityDetails(ctx *pulumi.Context,
	name string, args *CapacityDetailsArgs, opts ...pulumi.ResourceOption) (*CapacityDetails, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:powerbidedicated/v20210101:CapacityDetails"),
		},
		{
			Type: pulumi.String("azure-native:powerbidedicated:CapacityDetails"),
		},
		{
			Type: pulumi.String("azure-nextgen:powerbidedicated:CapacityDetails"),
		},
		{
			Type: pulumi.String("azure-native:powerbidedicated/v20171001:CapacityDetails"),
		},
		{
			Type: pulumi.String("azure-nextgen:powerbidedicated/v20171001:CapacityDetails"),
		},
	})
	opts = append(opts, aliases)
	var resource CapacityDetails
	err := ctx.RegisterResource("azure-native:powerbidedicated/v20210101:CapacityDetails", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCapacityDetails(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityDetailsState, opts ...pulumi.ResourceOption) (*CapacityDetails, error) {
	var resource CapacityDetails
	err := ctx.ReadResource("azure-native:powerbidedicated/v20210101:CapacityDetails", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type capacityDetailsState struct {
}

type CapacityDetailsState struct {
}

func (CapacityDetailsState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityDetailsState)(nil)).Elem()
}

type capacityDetailsArgs struct {
	Administration        *DedicatedCapacityAdministrators `pulumi:"administration"`
	DedicatedCapacityName *string                          `pulumi:"dedicatedCapacityName"`
	Location              *string                          `pulumi:"location"`
	Mode                  *string                          `pulumi:"mode"`
	ResourceGroupName     string                           `pulumi:"resourceGroupName"`
	Sku                   CapacitySku                      `pulumi:"sku"`
	SystemData            *SystemData                      `pulumi:"systemData"`
	Tags                  map[string]string                `pulumi:"tags"`
}


type CapacityDetailsArgs struct {
	Administration        DedicatedCapacityAdministratorsPtrInput
	DedicatedCapacityName pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	Mode                  pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Sku                   CapacitySkuInput
	SystemData            SystemDataPtrInput
	Tags                  pulumi.StringMapInput
}

func (CapacityDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityDetailsArgs)(nil)).Elem()
}

type CapacityDetailsInput interface {
	pulumi.Input

	ToCapacityDetailsOutput() CapacityDetailsOutput
	ToCapacityDetailsOutputWithContext(ctx context.Context) CapacityDetailsOutput
}

func (*CapacityDetails) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityDetails)(nil))
}

func (i *CapacityDetails) ToCapacityDetailsOutput() CapacityDetailsOutput {
	return i.ToCapacityDetailsOutputWithContext(context.Background())
}

func (i *CapacityDetails) ToCapacityDetailsOutputWithContext(ctx context.Context) CapacityDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityDetailsOutput)
}

type CapacityDetailsOutput struct{ *pulumi.OutputState }

func (CapacityDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityDetails)(nil))
}

func (o CapacityDetailsOutput) ToCapacityDetailsOutput() CapacityDetailsOutput {
	return o
}

func (o CapacityDetailsOutput) ToCapacityDetailsOutputWithContext(ctx context.Context) CapacityDetailsOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CapacityDetailsInput)(nil)).Elem(), &CapacityDetails{})
	pulumi.RegisterOutputType(CapacityDetailsOutput{})
}
