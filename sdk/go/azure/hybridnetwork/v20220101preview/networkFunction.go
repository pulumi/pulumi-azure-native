


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkFunction struct {
	pulumi.CustomResourceState

	Device                                 SubResourceResponsePtrOutput                        `pulumi:"device"`
	Etag                                   pulumi.StringPtrOutput                              `pulumi:"etag"`
	Location                               pulumi.StringOutput                                 `pulumi:"location"`
	ManagedApplication                     SubResourceResponseOutput                           `pulumi:"managedApplication"`
	ManagedApplicationParameters           pulumi.AnyOutput                                    `pulumi:"managedApplicationParameters"`
	Name                                   pulumi.StringOutput                                 `pulumi:"name"`
	NetworkFunctionContainerConfigurations pulumi.AnyOutput                                    `pulumi:"networkFunctionContainerConfigurations"`
	NetworkFunctionUserConfigurations      NetworkFunctionUserConfigurationResponseArrayOutput `pulumi:"networkFunctionUserConfigurations"`
	ProvisioningState                      pulumi.StringOutput                                 `pulumi:"provisioningState"`
	ServiceKey                             pulumi.StringOutput                                 `pulumi:"serviceKey"`
	SkuName                                pulumi.StringPtrOutput                              `pulumi:"skuName"`
	SkuType                                pulumi.StringOutput                                 `pulumi:"skuType"`
	SystemData                             SystemDataResponseOutput                            `pulumi:"systemData"`
	Tags                                   pulumi.StringMapOutput                              `pulumi:"tags"`
	Type                                   pulumi.StringOutput                                 `pulumi:"type"`
	VendorName                             pulumi.StringPtrOutput                              `pulumi:"vendorName"`
	VendorProvisioningState                pulumi.StringOutput                                 `pulumi:"vendorProvisioningState"`
}


func NewNetworkFunction(ctx *pulumi.Context,
	name string, args *NetworkFunctionArgs, opts ...pulumi.ResourceOption) (*NetworkFunction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridnetwork:NetworkFunction"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20200101preview:NetworkFunction"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20210501:NetworkFunction"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkFunction
	err := ctx.RegisterResource("azure-native:hybridnetwork/v20220101preview:NetworkFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkFunctionState, opts ...pulumi.ResourceOption) (*NetworkFunction, error) {
	var resource NetworkFunction
	err := ctx.ReadResource("azure-native:hybridnetwork/v20220101preview:NetworkFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkFunctionState struct {
}

type NetworkFunctionState struct {
}

func (NetworkFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkFunctionState)(nil)).Elem()
}

type networkFunctionArgs struct {
	Device                                 *SubResource                       `pulumi:"device"`
	Location                               *string                            `pulumi:"location"`
	ManagedApplicationParameters           interface{}                        `pulumi:"managedApplicationParameters"`
	NetworkFunctionContainerConfigurations interface{}                        `pulumi:"networkFunctionContainerConfigurations"`
	NetworkFunctionName                    *string                            `pulumi:"networkFunctionName"`
	NetworkFunctionUserConfigurations      []NetworkFunctionUserConfiguration `pulumi:"networkFunctionUserConfigurations"`
	ResourceGroupName                      string                             `pulumi:"resourceGroupName"`
	SkuName                                *string                            `pulumi:"skuName"`
	Tags                                   map[string]string                  `pulumi:"tags"`
	VendorName                             *string                            `pulumi:"vendorName"`
}


type NetworkFunctionArgs struct {
	Device                                 SubResourcePtrInput
	Location                               pulumi.StringPtrInput
	ManagedApplicationParameters           pulumi.Input
	NetworkFunctionContainerConfigurations pulumi.Input
	NetworkFunctionName                    pulumi.StringPtrInput
	NetworkFunctionUserConfigurations      NetworkFunctionUserConfigurationArrayInput
	ResourceGroupName                      pulumi.StringInput
	SkuName                                pulumi.StringPtrInput
	Tags                                   pulumi.StringMapInput
	VendorName                             pulumi.StringPtrInput
}

func (NetworkFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkFunctionArgs)(nil)).Elem()
}

type NetworkFunctionInput interface {
	pulumi.Input

	ToNetworkFunctionOutput() NetworkFunctionOutput
	ToNetworkFunctionOutputWithContext(ctx context.Context) NetworkFunctionOutput
}

func (*NetworkFunction) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunction)(nil)).Elem()
}

func (i *NetworkFunction) ToNetworkFunctionOutput() NetworkFunctionOutput {
	return i.ToNetworkFunctionOutputWithContext(context.Background())
}

func (i *NetworkFunction) ToNetworkFunctionOutputWithContext(ctx context.Context) NetworkFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionOutput)
}

type NetworkFunctionOutput struct{ *pulumi.OutputState }

func (NetworkFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunction)(nil)).Elem()
}

func (o NetworkFunctionOutput) ToNetworkFunctionOutput() NetworkFunctionOutput {
	return o
}

func (o NetworkFunctionOutput) ToNetworkFunctionOutputWithContext(ctx context.Context) NetworkFunctionOutput {
	return o
}

func (o NetworkFunctionOutput) Device() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *NetworkFunction) SubResourceResponsePtrOutput { return v.Device }).(SubResourceResponsePtrOutput)
}

func (o NetworkFunctionOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkFunction) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkFunctionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFunction) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o NetworkFunctionOutput) ManagedApplication() SubResourceResponseOutput {
	return o.ApplyT(func(v *NetworkFunction) SubResourceResponseOutput { return v.ManagedApplication }).(SubResourceResponseOutput)
}

func (o NetworkFunctionOutput) ManagedApplicationParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *NetworkFunction) pulumi.AnyOutput { return v.ManagedApplicationParameters }).(pulumi.AnyOutput)
}

func (o NetworkFunctionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFunction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkFunctionOutput) NetworkFunctionContainerConfigurations() pulumi.AnyOutput {
	return o.ApplyT(func(v *NetworkFunction) pulumi.AnyOutput { return v.NetworkFunctionContainerConfigurations }).(pulumi.AnyOutput)
}

func (o NetworkFunctionOutput) NetworkFunctionUserConfigurations() NetworkFunctionUserConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkFunction) NetworkFunctionUserConfigurationResponseArrayOutput {
		return v.NetworkFunctionUserConfigurations
	}).(NetworkFunctionUserConfigurationResponseArrayOutput)
}

func (o NetworkFunctionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFunction) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NetworkFunctionOutput) ServiceKey() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFunction) pulumi.StringOutput { return v.ServiceKey }).(pulumi.StringOutput)
}

func (o NetworkFunctionOutput) SkuName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkFunction) pulumi.StringPtrOutput { return v.SkuName }).(pulumi.StringPtrOutput)
}

func (o NetworkFunctionOutput) SkuType() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFunction) pulumi.StringOutput { return v.SkuType }).(pulumi.StringOutput)
}

func (o NetworkFunctionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkFunction) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o NetworkFunctionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkFunction) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkFunctionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFunction) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NetworkFunctionOutput) VendorName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkFunction) pulumi.StringPtrOutput { return v.VendorName }).(pulumi.StringPtrOutput)
}

func (o NetworkFunctionOutput) VendorProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFunction) pulumi.StringOutput { return v.VendorProvisioningState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkFunctionOutput{})
}
