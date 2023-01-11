


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CommunicationsGateway struct {
	pulumi.CustomResourceState

	ApiBridge         pulumi.AnyOutput                           `pulumi:"apiBridge"`
	Codecs            pulumi.StringArrayOutput                   `pulumi:"codecs"`
	Connectivity      pulumi.StringOutput                        `pulumi:"connectivity"`
	E911Type          pulumi.StringOutput                        `pulumi:"e911Type"`
	Location          pulumi.StringOutput                        `pulumi:"location"`
	Name              pulumi.StringOutput                        `pulumi:"name"`
	Platforms         pulumi.StringArrayOutput                   `pulumi:"platforms"`
	ProvisioningState pulumi.StringOutput                        `pulumi:"provisioningState"`
	ServiceLocations  ServiceRegionPropertiesResponseArrayOutput `pulumi:"serviceLocations"`
	Status            pulumi.StringOutput                        `pulumi:"status"`
	SystemData        SystemDataResponseOutput                   `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                     `pulumi:"tags"`
	Type              pulumi.StringOutput                        `pulumi:"type"`
}


func NewCommunicationsGateway(ctx *pulumi.Context,
	name string, args *CommunicationsGatewayArgs, opts ...pulumi.ResourceOption) (*CommunicationsGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Codecs == nil {
		return nil, errors.New("invalid value for required argument 'Codecs'")
	}
	if args.Connectivity == nil {
		return nil, errors.New("invalid value for required argument 'Connectivity'")
	}
	if args.E911Type == nil {
		return nil, errors.New("invalid value for required argument 'E911Type'")
	}
	if args.Platforms == nil {
		return nil, errors.New("invalid value for required argument 'Platforms'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceLocations == nil {
		return nil, errors.New("invalid value for required argument 'ServiceLocations'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:voiceservices:CommunicationsGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource CommunicationsGateway
	err := ctx.RegisterResource("azure-native:voiceservices/v20221201preview:CommunicationsGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCommunicationsGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommunicationsGatewayState, opts ...pulumi.ResourceOption) (*CommunicationsGateway, error) {
	var resource CommunicationsGateway
	err := ctx.ReadResource("azure-native:voiceservices/v20221201preview:CommunicationsGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type communicationsGatewayState struct {
}

type CommunicationsGatewayState struct {
}

func (CommunicationsGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*communicationsGatewayState)(nil)).Elem()
}

type communicationsGatewayArgs struct {
	ApiBridge                 interface{}               `pulumi:"apiBridge"`
	Codecs                    []string                  `pulumi:"codecs"`
	CommunicationsGatewayName *string                   `pulumi:"communicationsGatewayName"`
	Connectivity              string                    `pulumi:"connectivity"`
	E911Type                  string                    `pulumi:"e911Type"`
	Location                  *string                   `pulumi:"location"`
	Platforms                 []string                  `pulumi:"platforms"`
	ResourceGroupName         string                    `pulumi:"resourceGroupName"`
	ServiceLocations          []ServiceRegionProperties `pulumi:"serviceLocations"`
	Tags                      map[string]string         `pulumi:"tags"`
}


type CommunicationsGatewayArgs struct {
	ApiBridge                 pulumi.Input
	Codecs                    pulumi.StringArrayInput
	CommunicationsGatewayName pulumi.StringPtrInput
	Connectivity              pulumi.StringInput
	E911Type                  pulumi.StringInput
	Location                  pulumi.StringPtrInput
	Platforms                 pulumi.StringArrayInput
	ResourceGroupName         pulumi.StringInput
	ServiceLocations          ServiceRegionPropertiesArrayInput
	Tags                      pulumi.StringMapInput
}

func (CommunicationsGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*communicationsGatewayArgs)(nil)).Elem()
}

type CommunicationsGatewayInput interface {
	pulumi.Input

	ToCommunicationsGatewayOutput() CommunicationsGatewayOutput
	ToCommunicationsGatewayOutputWithContext(ctx context.Context) CommunicationsGatewayOutput
}

func (*CommunicationsGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**CommunicationsGateway)(nil)).Elem()
}

func (i *CommunicationsGateway) ToCommunicationsGatewayOutput() CommunicationsGatewayOutput {
	return i.ToCommunicationsGatewayOutputWithContext(context.Background())
}

func (i *CommunicationsGateway) ToCommunicationsGatewayOutputWithContext(ctx context.Context) CommunicationsGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunicationsGatewayOutput)
}

type CommunicationsGatewayOutput struct{ *pulumi.OutputState }

func (CommunicationsGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommunicationsGateway)(nil)).Elem()
}

func (o CommunicationsGatewayOutput) ToCommunicationsGatewayOutput() CommunicationsGatewayOutput {
	return o
}

func (o CommunicationsGatewayOutput) ToCommunicationsGatewayOutputWithContext(ctx context.Context) CommunicationsGatewayOutput {
	return o
}

func (o CommunicationsGatewayOutput) ApiBridge() pulumi.AnyOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.AnyOutput { return v.ApiBridge }).(pulumi.AnyOutput)
}

func (o CommunicationsGatewayOutput) Codecs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringArrayOutput { return v.Codecs }).(pulumi.StringArrayOutput)
}

func (o CommunicationsGatewayOutput) Connectivity() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.Connectivity }).(pulumi.StringOutput)
}

func (o CommunicationsGatewayOutput) E911Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.E911Type }).(pulumi.StringOutput)
}

func (o CommunicationsGatewayOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CommunicationsGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CommunicationsGatewayOutput) Platforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringArrayOutput { return v.Platforms }).(pulumi.StringArrayOutput)
}

func (o CommunicationsGatewayOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CommunicationsGatewayOutput) ServiceLocations() ServiceRegionPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *CommunicationsGateway) ServiceRegionPropertiesResponseArrayOutput { return v.ServiceLocations }).(ServiceRegionPropertiesResponseArrayOutput)
}

func (o CommunicationsGatewayOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o CommunicationsGatewayOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CommunicationsGateway) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o CommunicationsGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CommunicationsGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CommunicationsGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CommunicationsGatewayOutput{})
}
