


package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualApplianceSite struct {
	pulumi.CustomResourceState

	AddressPrefix     pulumi.StringPtrOutput                     `pulumi:"addressPrefix"`
	Etag              pulumi.StringOutput                        `pulumi:"etag"`
	Name              pulumi.StringPtrOutput                     `pulumi:"name"`
	O365Policy        Office365PolicyPropertiesResponsePtrOutput `pulumi:"o365Policy"`
	ProvisioningState pulumi.StringOutput                        `pulumi:"provisioningState"`
	Type              pulumi.StringOutput                        `pulumi:"type"`
}


func NewVirtualApplianceSite(ctx *pulumi.Context,
	name string, args *VirtualApplianceSiteArgs, opts ...pulumi.ResourceOption) (*VirtualApplianceSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkVirtualApplianceName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkVirtualApplianceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualApplianceSite"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualApplianceSite
	err := ctx.RegisterResource("azure-native:network:VirtualApplianceSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualApplianceSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualApplianceSiteState, opts ...pulumi.ResourceOption) (*VirtualApplianceSite, error) {
	var resource VirtualApplianceSite
	err := ctx.ReadResource("azure-native:network:VirtualApplianceSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualApplianceSiteState struct {
}

type VirtualApplianceSiteState struct {
}

func (VirtualApplianceSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualApplianceSiteState)(nil)).Elem()
}

type virtualApplianceSiteArgs struct {
	AddressPrefix               *string                    `pulumi:"addressPrefix"`
	Id                          *string                    `pulumi:"id"`
	Name                        *string                    `pulumi:"name"`
	NetworkVirtualApplianceName string                     `pulumi:"networkVirtualApplianceName"`
	O365Policy                  *Office365PolicyProperties `pulumi:"o365Policy"`
	ResourceGroupName           string                     `pulumi:"resourceGroupName"`
	SiteName                    *string                    `pulumi:"siteName"`
}


type VirtualApplianceSiteArgs struct {
	AddressPrefix               pulumi.StringPtrInput
	Id                          pulumi.StringPtrInput
	Name                        pulumi.StringPtrInput
	NetworkVirtualApplianceName pulumi.StringInput
	O365Policy                  Office365PolicyPropertiesPtrInput
	ResourceGroupName           pulumi.StringInput
	SiteName                    pulumi.StringPtrInput
}

func (VirtualApplianceSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualApplianceSiteArgs)(nil)).Elem()
}

type VirtualApplianceSiteInput interface {
	pulumi.Input

	ToVirtualApplianceSiteOutput() VirtualApplianceSiteOutput
	ToVirtualApplianceSiteOutputWithContext(ctx context.Context) VirtualApplianceSiteOutput
}

func (*VirtualApplianceSite) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualApplianceSite)(nil)).Elem()
}

func (i *VirtualApplianceSite) ToVirtualApplianceSiteOutput() VirtualApplianceSiteOutput {
	return i.ToVirtualApplianceSiteOutputWithContext(context.Background())
}

func (i *VirtualApplianceSite) ToVirtualApplianceSiteOutputWithContext(ctx context.Context) VirtualApplianceSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualApplianceSiteOutput)
}

type VirtualApplianceSiteOutput struct{ *pulumi.OutputState }

func (VirtualApplianceSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualApplianceSite)(nil)).Elem()
}

func (o VirtualApplianceSiteOutput) ToVirtualApplianceSiteOutput() VirtualApplianceSiteOutput {
	return o
}

func (o VirtualApplianceSiteOutput) ToVirtualApplianceSiteOutputWithContext(ctx context.Context) VirtualApplianceSiteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualApplianceSiteOutput{})
}
