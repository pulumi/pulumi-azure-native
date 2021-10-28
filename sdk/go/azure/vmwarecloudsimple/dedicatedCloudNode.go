


package vmwarecloudsimple

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DedicatedCloudNode struct {
	pulumi.CustomResourceState

	AvailabilityZoneId   pulumi.StringOutput    `pulumi:"availabilityZoneId"`
	AvailabilityZoneName pulumi.StringOutput    `pulumi:"availabilityZoneName"`
	CloudRackName        pulumi.StringOutput    `pulumi:"cloudRackName"`
	Created              pulumi.StringOutput    `pulumi:"created"`
	Location             pulumi.StringOutput    `pulumi:"location"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	NodesCount           pulumi.IntOutput       `pulumi:"nodesCount"`
	PlacementGroupId     pulumi.StringOutput    `pulumi:"placementGroupId"`
	PlacementGroupName   pulumi.StringOutput    `pulumi:"placementGroupName"`
	PrivateCloudId       pulumi.StringOutput    `pulumi:"privateCloudId"`
	PrivateCloudName     pulumi.StringOutput    `pulumi:"privateCloudName"`
	ProvisioningState    pulumi.StringOutput    `pulumi:"provisioningState"`
	PurchaseId           pulumi.StringOutput    `pulumi:"purchaseId"`
	Sku                  SkuResponsePtrOutput   `pulumi:"sku"`
	Status               pulumi.StringOutput    `pulumi:"status"`
	Tags                 pulumi.StringMapOutput `pulumi:"tags"`
	Type                 pulumi.StringOutput    `pulumi:"type"`
	VmwareClusterName    pulumi.StringOutput    `pulumi:"vmwareClusterName"`
}


func NewDedicatedCloudNode(ctx *pulumi.Context,
	name string, args *DedicatedCloudNodeArgs, opts ...pulumi.ResourceOption) (*DedicatedCloudNode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZoneId == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZoneId'")
	}
	if args.Id == nil {
		return nil, errors.New("invalid value for required argument 'Id'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.NodesCount == nil {
		return nil, errors.New("invalid value for required argument 'NodesCount'")
	}
	if args.PlacementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'PlacementGroupId'")
	}
	if args.PurchaseId == nil {
		return nil, errors.New("invalid value for required argument 'PurchaseId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:vmwarecloudsimple:DedicatedCloudNode"),
		},
		{
			Type: pulumi.String("azure-native:vmwarecloudsimple/v20190401:DedicatedCloudNode"),
		},
		{
			Type: pulumi.String("azure-nextgen:vmwarecloudsimple/v20190401:DedicatedCloudNode"),
		},
	})
	opts = append(opts, aliases)
	var resource DedicatedCloudNode
	err := ctx.RegisterResource("azure-native:vmwarecloudsimple:DedicatedCloudNode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDedicatedCloudNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedCloudNodeState, opts ...pulumi.ResourceOption) (*DedicatedCloudNode, error) {
	var resource DedicatedCloudNode
	err := ctx.ReadResource("azure-native:vmwarecloudsimple:DedicatedCloudNode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dedicatedCloudNodeState struct {
}

type DedicatedCloudNodeState struct {
}

func (DedicatedCloudNodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedCloudNodeState)(nil)).Elem()
}

type dedicatedCloudNodeArgs struct {
	AvailabilityZoneId     string            `pulumi:"availabilityZoneId"`
	DedicatedCloudNodeName *string           `pulumi:"dedicatedCloudNodeName"`
	Id                     string            `pulumi:"id"`
	Location               *string           `pulumi:"location"`
	Name                   string            `pulumi:"name"`
	NodesCount             int               `pulumi:"nodesCount"`
	PlacementGroupId       string            `pulumi:"placementGroupId"`
	PurchaseId             string            `pulumi:"purchaseId"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Sku                    *Sku              `pulumi:"sku"`
	Tags                   map[string]string `pulumi:"tags"`
}


type DedicatedCloudNodeArgs struct {
	AvailabilityZoneId     pulumi.StringInput
	DedicatedCloudNodeName pulumi.StringPtrInput
	Id                     pulumi.StringInput
	Location               pulumi.StringPtrInput
	Name                   pulumi.StringInput
	NodesCount             pulumi.IntInput
	PlacementGroupId       pulumi.StringInput
	PurchaseId             pulumi.StringInput
	ResourceGroupName      pulumi.StringInput
	Sku                    SkuPtrInput
	Tags                   pulumi.StringMapInput
}

func (DedicatedCloudNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedCloudNodeArgs)(nil)).Elem()
}

type DedicatedCloudNodeInput interface {
	pulumi.Input

	ToDedicatedCloudNodeOutput() DedicatedCloudNodeOutput
	ToDedicatedCloudNodeOutputWithContext(ctx context.Context) DedicatedCloudNodeOutput
}

func (*DedicatedCloudNode) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedCloudNode)(nil))
}

func (i *DedicatedCloudNode) ToDedicatedCloudNodeOutput() DedicatedCloudNodeOutput {
	return i.ToDedicatedCloudNodeOutputWithContext(context.Background())
}

func (i *DedicatedCloudNode) ToDedicatedCloudNodeOutputWithContext(ctx context.Context) DedicatedCloudNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCloudNodeOutput)
}

type DedicatedCloudNodeOutput struct{ *pulumi.OutputState }

func (DedicatedCloudNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedCloudNode)(nil))
}

func (o DedicatedCloudNodeOutput) ToDedicatedCloudNodeOutput() DedicatedCloudNodeOutput {
	return o
}

func (o DedicatedCloudNodeOutput) ToDedicatedCloudNodeOutputWithContext(ctx context.Context) DedicatedCloudNodeOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCloudNodeInput)(nil)).Elem(), &DedicatedCloudNode{})
	pulumi.RegisterOutputType(DedicatedCloudNodeOutput{})
}
