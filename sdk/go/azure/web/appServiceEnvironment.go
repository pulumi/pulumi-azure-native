


package web

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppServiceEnvironment struct {
	pulumi.CustomResourceState

	ClusterSettings           NameValuePairResponseArrayOutput    `pulumi:"clusterSettings"`
	DedicatedHostCount        pulumi.IntOutput                    `pulumi:"dedicatedHostCount"`
	DnsSuffix                 pulumi.StringPtrOutput              `pulumi:"dnsSuffix"`
	FrontEndScaleFactor       pulumi.IntPtrOutput                 `pulumi:"frontEndScaleFactor"`
	HasLinuxWorkers           pulumi.BoolOutput                   `pulumi:"hasLinuxWorkers"`
	InternalLoadBalancingMode pulumi.StringPtrOutput              `pulumi:"internalLoadBalancingMode"`
	IpsslAddressCount         pulumi.IntPtrOutput                 `pulumi:"ipsslAddressCount"`
	Kind                      pulumi.StringPtrOutput              `pulumi:"kind"`
	Location                  pulumi.StringOutput                 `pulumi:"location"`
	MaximumNumberOfMachines   pulumi.IntOutput                    `pulumi:"maximumNumberOfMachines"`
	MultiRoleCount            pulumi.IntOutput                    `pulumi:"multiRoleCount"`
	MultiSize                 pulumi.StringPtrOutput              `pulumi:"multiSize"`
	Name                      pulumi.StringOutput                 `pulumi:"name"`
	ProvisioningState         pulumi.StringOutput                 `pulumi:"provisioningState"`
	Status                    pulumi.StringOutput                 `pulumi:"status"`
	Suspended                 pulumi.BoolOutput                   `pulumi:"suspended"`
	Tags                      pulumi.StringMapOutput              `pulumi:"tags"`
	Type                      pulumi.StringOutput                 `pulumi:"type"`
	UserWhitelistedIpRanges   pulumi.StringArrayOutput            `pulumi:"userWhitelistedIpRanges"`
	VirtualNetwork            VirtualNetworkProfileResponseOutput `pulumi:"virtualNetwork"`
}


func NewAppServiceEnvironment(ctx *pulumi.Context,
	name string, args *AppServiceEnvironmentArgs, opts ...pulumi.ResourceOption) (*AppServiceEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualNetwork == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetwork'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20150801:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160901:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:AppServiceEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource AppServiceEnvironment
	err := ctx.RegisterResource("azure-native:web:AppServiceEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAppServiceEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceEnvironmentState, opts ...pulumi.ResourceOption) (*AppServiceEnvironment, error) {
	var resource AppServiceEnvironment
	err := ctx.ReadResource("azure-native:web:AppServiceEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type appServiceEnvironmentState struct {
}

type AppServiceEnvironmentState struct {
}

func (AppServiceEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentState)(nil)).Elem()
}

type appServiceEnvironmentArgs struct {
	ClusterSettings           []NameValuePair       `pulumi:"clusterSettings"`
	DnsSuffix                 *string               `pulumi:"dnsSuffix"`
	FrontEndScaleFactor       *int                  `pulumi:"frontEndScaleFactor"`
	InternalLoadBalancingMode *string               `pulumi:"internalLoadBalancingMode"`
	IpsslAddressCount         *int                  `pulumi:"ipsslAddressCount"`
	Kind                      *string               `pulumi:"kind"`
	Location                  *string               `pulumi:"location"`
	MultiSize                 *string               `pulumi:"multiSize"`
	Name                      *string               `pulumi:"name"`
	ResourceGroupName         string                `pulumi:"resourceGroupName"`
	Tags                      map[string]string     `pulumi:"tags"`
	UserWhitelistedIpRanges   []string              `pulumi:"userWhitelistedIpRanges"`
	VirtualNetwork            VirtualNetworkProfile `pulumi:"virtualNetwork"`
}


type AppServiceEnvironmentArgs struct {
	ClusterSettings           NameValuePairArrayInput
	DnsSuffix                 pulumi.StringPtrInput
	FrontEndScaleFactor       pulumi.IntPtrInput
	InternalLoadBalancingMode pulumi.StringPtrInput
	IpsslAddressCount         pulumi.IntPtrInput
	Kind                      pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	MultiSize                 pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Tags                      pulumi.StringMapInput
	UserWhitelistedIpRanges   pulumi.StringArrayInput
	VirtualNetwork            VirtualNetworkProfileInput
}

func (AppServiceEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentArgs)(nil)).Elem()
}

type AppServiceEnvironmentInput interface {
	pulumi.Input

	ToAppServiceEnvironmentOutput() AppServiceEnvironmentOutput
	ToAppServiceEnvironmentOutputWithContext(ctx context.Context) AppServiceEnvironmentOutput
}

func (*AppServiceEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceEnvironment)(nil)).Elem()
}

func (i *AppServiceEnvironment) ToAppServiceEnvironmentOutput() AppServiceEnvironmentOutput {
	return i.ToAppServiceEnvironmentOutputWithContext(context.Background())
}

func (i *AppServiceEnvironment) ToAppServiceEnvironmentOutputWithContext(ctx context.Context) AppServiceEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceEnvironmentOutput)
}

type AppServiceEnvironmentOutput struct{ *pulumi.OutputState }

func (AppServiceEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceEnvironment)(nil)).Elem()
}

func (o AppServiceEnvironmentOutput) ToAppServiceEnvironmentOutput() AppServiceEnvironmentOutput {
	return o
}

func (o AppServiceEnvironmentOutput) ToAppServiceEnvironmentOutputWithContext(ctx context.Context) AppServiceEnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AppServiceEnvironmentOutput{})
}
