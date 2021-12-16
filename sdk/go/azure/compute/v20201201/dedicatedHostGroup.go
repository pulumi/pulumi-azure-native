


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DedicatedHostGroup struct {
	pulumi.CustomResourceState

	Hosts                     SubResourceReadOnlyResponseArrayOutput       `pulumi:"hosts"`
	InstanceView              DedicatedHostGroupInstanceViewResponseOutput `pulumi:"instanceView"`
	Location                  pulumi.StringOutput                          `pulumi:"location"`
	Name                      pulumi.StringOutput                          `pulumi:"name"`
	PlatformFaultDomainCount  pulumi.IntOutput                             `pulumi:"platformFaultDomainCount"`
	SupportAutomaticPlacement pulumi.BoolPtrOutput                         `pulumi:"supportAutomaticPlacement"`
	Tags                      pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                      pulumi.StringOutput                          `pulumi:"type"`
	Zones                     pulumi.StringArrayOutput                     `pulumi:"zones"`
}


func NewDedicatedHostGroup(ctx *pulumi.Context,
	name string, args *DedicatedHostGroupArgs, opts ...pulumi.ResourceOption) (*DedicatedHostGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PlatformFaultDomainCount == nil {
		return nil, errors.New("invalid value for required argument 'PlatformFaultDomainCount'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:DedicatedHostGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource DedicatedHostGroup
	err := ctx.RegisterResource("azure-native:compute/v20201201:DedicatedHostGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDedicatedHostGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedHostGroupState, opts ...pulumi.ResourceOption) (*DedicatedHostGroup, error) {
	var resource DedicatedHostGroup
	err := ctx.ReadResource("azure-native:compute/v20201201:DedicatedHostGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dedicatedHostGroupState struct {
}

type DedicatedHostGroupState struct {
}

func (DedicatedHostGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostGroupState)(nil)).Elem()
}

type dedicatedHostGroupArgs struct {
	HostGroupName             *string           `pulumi:"hostGroupName"`
	Location                  *string           `pulumi:"location"`
	PlatformFaultDomainCount  int               `pulumi:"platformFaultDomainCount"`
	ResourceGroupName         string            `pulumi:"resourceGroupName"`
	SupportAutomaticPlacement *bool             `pulumi:"supportAutomaticPlacement"`
	Tags                      map[string]string `pulumi:"tags"`
	Zones                     []string          `pulumi:"zones"`
}


type DedicatedHostGroupArgs struct {
	HostGroupName             pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	PlatformFaultDomainCount  pulumi.IntInput
	ResourceGroupName         pulumi.StringInput
	SupportAutomaticPlacement pulumi.BoolPtrInput
	Tags                      pulumi.StringMapInput
	Zones                     pulumi.StringArrayInput
}

func (DedicatedHostGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostGroupArgs)(nil)).Elem()
}

type DedicatedHostGroupInput interface {
	pulumi.Input

	ToDedicatedHostGroupOutput() DedicatedHostGroupOutput
	ToDedicatedHostGroupOutputWithContext(ctx context.Context) DedicatedHostGroupOutput
}

func (*DedicatedHostGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHostGroup)(nil)).Elem()
}

func (i *DedicatedHostGroup) ToDedicatedHostGroupOutput() DedicatedHostGroupOutput {
	return i.ToDedicatedHostGroupOutputWithContext(context.Background())
}

func (i *DedicatedHostGroup) ToDedicatedHostGroupOutputWithContext(ctx context.Context) DedicatedHostGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostGroupOutput)
}

type DedicatedHostGroupOutput struct{ *pulumi.OutputState }

func (DedicatedHostGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHostGroup)(nil)).Elem()
}

func (o DedicatedHostGroupOutput) ToDedicatedHostGroupOutput() DedicatedHostGroupOutput {
	return o
}

func (o DedicatedHostGroupOutput) ToDedicatedHostGroupOutputWithContext(ctx context.Context) DedicatedHostGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DedicatedHostGroupOutput{})
}
