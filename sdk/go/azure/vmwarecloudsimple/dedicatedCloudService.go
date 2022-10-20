


package vmwarecloudsimple

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DedicatedCloudService struct {
	pulumi.CustomResourceState

	GatewaySubnet      pulumi.StringOutput    `pulumi:"gatewaySubnet"`
	IsAccountOnboarded pulumi.StringOutput    `pulumi:"isAccountOnboarded"`
	Location           pulumi.StringOutput    `pulumi:"location"`
	Name               pulumi.StringOutput    `pulumi:"name"`
	Nodes              pulumi.IntOutput       `pulumi:"nodes"`
	ServiceURL         pulumi.StringOutput    `pulumi:"serviceURL"`
	Tags               pulumi.StringMapOutput `pulumi:"tags"`
	Type               pulumi.StringOutput    `pulumi:"type"`
}


func NewDedicatedCloudService(ctx *pulumi.Context,
	name string, args *DedicatedCloudServiceArgs, opts ...pulumi.ResourceOption) (*DedicatedCloudService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewaySubnet == nil {
		return nil, errors.New("invalid value for required argument 'GatewaySubnet'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:vmwarecloudsimple/v20190401:DedicatedCloudService"),
		},
	})
	opts = append(opts, aliases)
	var resource DedicatedCloudService
	err := ctx.RegisterResource("azure-native:vmwarecloudsimple:DedicatedCloudService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDedicatedCloudService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedCloudServiceState, opts ...pulumi.ResourceOption) (*DedicatedCloudService, error) {
	var resource DedicatedCloudService
	err := ctx.ReadResource("azure-native:vmwarecloudsimple:DedicatedCloudService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dedicatedCloudServiceState struct {
}

type DedicatedCloudServiceState struct {
}

func (DedicatedCloudServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedCloudServiceState)(nil)).Elem()
}

type dedicatedCloudServiceArgs struct {
	DedicatedCloudServiceName *string           `pulumi:"dedicatedCloudServiceName"`
	GatewaySubnet             string            `pulumi:"gatewaySubnet"`
	Location                  *string           `pulumi:"location"`
	ResourceGroupName         string            `pulumi:"resourceGroupName"`
	Tags                      map[string]string `pulumi:"tags"`
}


type DedicatedCloudServiceArgs struct {
	DedicatedCloudServiceName pulumi.StringPtrInput
	GatewaySubnet             pulumi.StringInput
	Location                  pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Tags                      pulumi.StringMapInput
}

func (DedicatedCloudServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedCloudServiceArgs)(nil)).Elem()
}

type DedicatedCloudServiceInput interface {
	pulumi.Input

	ToDedicatedCloudServiceOutput() DedicatedCloudServiceOutput
	ToDedicatedCloudServiceOutputWithContext(ctx context.Context) DedicatedCloudServiceOutput
}

func (*DedicatedCloudService) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCloudService)(nil)).Elem()
}

func (i *DedicatedCloudService) ToDedicatedCloudServiceOutput() DedicatedCloudServiceOutput {
	return i.ToDedicatedCloudServiceOutputWithContext(context.Background())
}

func (i *DedicatedCloudService) ToDedicatedCloudServiceOutputWithContext(ctx context.Context) DedicatedCloudServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCloudServiceOutput)
}

type DedicatedCloudServiceOutput struct{ *pulumi.OutputState }

func (DedicatedCloudServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCloudService)(nil)).Elem()
}

func (o DedicatedCloudServiceOutput) ToDedicatedCloudServiceOutput() DedicatedCloudServiceOutput {
	return o
}

func (o DedicatedCloudServiceOutput) ToDedicatedCloudServiceOutputWithContext(ctx context.Context) DedicatedCloudServiceOutput {
	return o
}

func (o DedicatedCloudServiceOutput) GatewaySubnet() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringOutput { return v.GatewaySubnet }).(pulumi.StringOutput)
}

func (o DedicatedCloudServiceOutput) IsAccountOnboarded() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringOutput { return v.IsAccountOnboarded }).(pulumi.StringOutput)
}

func (o DedicatedCloudServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DedicatedCloudServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DedicatedCloudServiceOutput) Nodes() pulumi.IntOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.IntOutput { return v.Nodes }).(pulumi.IntOutput)
}

func (o DedicatedCloudServiceOutput) ServiceURL() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringOutput { return v.ServiceURL }).(pulumi.StringOutput)
}

func (o DedicatedCloudServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DedicatedCloudServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DedicatedCloudServiceOutput{})
}
