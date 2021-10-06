


package v20210501

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
			Type: pulumi.String("azure-nextgen:hybridnetwork/v20210501:NetworkFunction"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork:NetworkFunction"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridnetwork:NetworkFunction"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20200101preview:NetworkFunction"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridnetwork/v20200101preview:NetworkFunction"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkFunction
	err := ctx.RegisterResource("azure-native:hybridnetwork/v20210501:NetworkFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkFunctionState, opts ...pulumi.ResourceOption) (*NetworkFunction, error) {
	var resource NetworkFunction
	err := ctx.ReadResource("azure-native:hybridnetwork/v20210501:NetworkFunction", name, id, state, &resource, opts...)
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
	Etag                                   *string                            `pulumi:"etag"`
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
	Etag                                   pulumi.StringPtrInput
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
	return reflect.TypeOf((*NetworkFunction)(nil))
}

func (i *NetworkFunction) ToNetworkFunctionOutput() NetworkFunctionOutput {
	return i.ToNetworkFunctionOutputWithContext(context.Background())
}

func (i *NetworkFunction) ToNetworkFunctionOutputWithContext(ctx context.Context) NetworkFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionOutput)
}

type NetworkFunctionOutput struct{ *pulumi.OutputState }

func (NetworkFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunction)(nil))
}

func (o NetworkFunctionOutput) ToNetworkFunctionOutput() NetworkFunctionOutput {
	return o
}

func (o NetworkFunctionOutput) ToNetworkFunctionOutputWithContext(ctx context.Context) NetworkFunctionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkFunctionOutput{})
}
