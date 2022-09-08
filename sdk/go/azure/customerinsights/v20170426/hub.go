


package v20170426

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Hub struct {
	pulumi.CustomResourceState

	ApiEndpoint       pulumi.StringOutput                   `pulumi:"apiEndpoint"`
	HubBillingInfo    HubBillingInfoFormatResponsePtrOutput `pulumi:"hubBillingInfo"`
	Location          pulumi.StringPtrOutput                `pulumi:"location"`
	Name              pulumi.StringOutput                   `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                   `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput                `pulumi:"tags"`
	TenantFeatures    pulumi.IntPtrOutput                   `pulumi:"tenantFeatures"`
	Type              pulumi.StringOutput                   `pulumi:"type"`
	WebEndpoint       pulumi.StringOutput                   `pulumi:"webEndpoint"`
}


func NewHub(ctx *pulumi.Context,
	name string, args *HubArgs, opts ...pulumi.ResourceOption) (*Hub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights:Hub"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170101:Hub"),
		},
	})
	opts = append(opts, aliases)
	var resource Hub
	err := ctx.RegisterResource("azure-native:customerinsights/v20170426:Hub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HubState, opts ...pulumi.ResourceOption) (*Hub, error) {
	var resource Hub
	err := ctx.ReadResource("azure-native:customerinsights/v20170426:Hub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hubState struct {
}

type HubState struct {
}

func (HubState) ElementType() reflect.Type {
	return reflect.TypeOf((*hubState)(nil)).Elem()
}

type hubArgs struct {
	HubBillingInfo    *HubBillingInfoFormat `pulumi:"hubBillingInfo"`
	HubName           *string               `pulumi:"hubName"`
	Location          *string               `pulumi:"location"`
	ResourceGroupName string                `pulumi:"resourceGroupName"`
	Tags              map[string]string     `pulumi:"tags"`
	TenantFeatures    *int                  `pulumi:"tenantFeatures"`
}


type HubArgs struct {
	HubBillingInfo    HubBillingInfoFormatPtrInput
	HubName           pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	TenantFeatures    pulumi.IntPtrInput
}

func (HubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hubArgs)(nil)).Elem()
}

type HubInput interface {
	pulumi.Input

	ToHubOutput() HubOutput
	ToHubOutputWithContext(ctx context.Context) HubOutput
}

func (*Hub) ElementType() reflect.Type {
	return reflect.TypeOf((**Hub)(nil)).Elem()
}

func (i *Hub) ToHubOutput() HubOutput {
	return i.ToHubOutputWithContext(context.Background())
}

func (i *Hub) ToHubOutputWithContext(ctx context.Context) HubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubOutput)
}

type HubOutput struct{ *pulumi.OutputState }

func (HubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Hub)(nil)).Elem()
}

func (o HubOutput) ToHubOutput() HubOutput {
	return o
}

func (o HubOutput) ToHubOutputWithContext(ctx context.Context) HubOutput {
	return o
}

func (o HubOutput) ApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringOutput { return v.ApiEndpoint }).(pulumi.StringOutput)
}

func (o HubOutput) HubBillingInfo() HubBillingInfoFormatResponsePtrOutput {
	return o.ApplyT(func(v *Hub) HubBillingInfoFormatResponsePtrOutput { return v.HubBillingInfo }).(HubBillingInfoFormatResponsePtrOutput)
}

func (o HubOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o HubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o HubOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o HubOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o HubOutput) TenantFeatures() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Hub) pulumi.IntPtrOutput { return v.TenantFeatures }).(pulumi.IntPtrOutput)
}

func (o HubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o HubOutput) WebEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringOutput { return v.WebEndpoint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HubOutput{})
}
