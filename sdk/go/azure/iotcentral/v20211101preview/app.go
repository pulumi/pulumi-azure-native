


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type App struct {
	pulumi.CustomResourceState

	ApplicationId              pulumi.StringOutput                            `pulumi:"applicationId"`
	DisplayName                pulumi.StringPtrOutput                         `pulumi:"displayName"`
	Identity                   SystemAssignedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Location                   pulumi.StringOutput                            `pulumi:"location"`
	Name                       pulumi.StringOutput                            `pulumi:"name"`
	NetworkRuleSets            NetworkRuleSetsResponsePtrOutput               `pulumi:"networkRuleSets"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput   `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                            `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrOutput                         `pulumi:"publicNetworkAccess"`
	Sku                        AppSkuInfoResponseOutput                       `pulumi:"sku"`
	State                      pulumi.StringOutput                            `pulumi:"state"`
	Subdomain                  pulumi.StringPtrOutput                         `pulumi:"subdomain"`
	SystemData                 SystemDataResponseOutput                       `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                         `pulumi:"tags"`
	Template                   pulumi.StringPtrOutput                         `pulumi:"template"`
	Type                       pulumi.StringOutput                            `pulumi:"type"`
}


func NewApp(ctx *pulumi.Context,
	name string, args *AppArgs, opts ...pulumi.ResourceOption) (*App, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.NetworkRuleSets != nil {
		args.NetworkRuleSets = args.NetworkRuleSets.ToNetworkRuleSetsPtrOutput().ApplyT(func(v *NetworkRuleSets) *NetworkRuleSets { return v.Defaults() }).(NetworkRuleSetsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:iotcentral:App"),
		},
		{
			Type: pulumi.String("azure-native:iotcentral/v20180901:App"),
		},
		{
			Type: pulumi.String("azure-native:iotcentral/v20210601:App"),
		},
	})
	opts = append(opts, aliases)
	var resource App
	err := ctx.RegisterResource("azure-native:iotcentral/v20211101preview:App", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppState, opts ...pulumi.ResourceOption) (*App, error) {
	var resource App
	err := ctx.ReadResource("azure-native:iotcentral/v20211101preview:App", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type appState struct {
}

type AppState struct {
}

func (AppState) ElementType() reflect.Type {
	return reflect.TypeOf((*appState)(nil)).Elem()
}

type appArgs struct {
	DisplayName         *string                        `pulumi:"displayName"`
	Identity            *SystemAssignedServiceIdentity `pulumi:"identity"`
	Location            *string                        `pulumi:"location"`
	NetworkRuleSets     *NetworkRuleSets               `pulumi:"networkRuleSets"`
	PublicNetworkAccess *string                        `pulumi:"publicNetworkAccess"`
	ResourceGroupName   string                         `pulumi:"resourceGroupName"`
	ResourceName        *string                        `pulumi:"resourceName"`
	Sku                 AppSkuInfo                     `pulumi:"sku"`
	Subdomain           *string                        `pulumi:"subdomain"`
	Tags                map[string]string              `pulumi:"tags"`
	Template            *string                        `pulumi:"template"`
}


type AppArgs struct {
	DisplayName         pulumi.StringPtrInput
	Identity            SystemAssignedServiceIdentityPtrInput
	Location            pulumi.StringPtrInput
	NetworkRuleSets     NetworkRuleSetsPtrInput
	PublicNetworkAccess pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	ResourceName        pulumi.StringPtrInput
	Sku                 AppSkuInfoInput
	Subdomain           pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
	Template            pulumi.StringPtrInput
}

func (AppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appArgs)(nil)).Elem()
}

type AppInput interface {
	pulumi.Input

	ToAppOutput() AppOutput
	ToAppOutputWithContext(ctx context.Context) AppOutput
}

func (*App) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (i *App) ToAppOutput() AppOutput {
	return i.ToAppOutputWithContext(context.Background())
}

func (i *App) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppOutput)
}

type AppOutput struct{ *pulumi.OutputState }

func (AppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (o AppOutput) ToAppOutput() AppOutput {
	return o
}

func (o AppOutput) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return o
}

func (o AppOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o AppOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o AppOutput) Identity() SystemAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *App) SystemAssignedServiceIdentityResponsePtrOutput { return v.Identity }).(SystemAssignedServiceIdentityResponsePtrOutput)
}

func (o AppOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppOutput) NetworkRuleSets() NetworkRuleSetsResponsePtrOutput {
	return o.ApplyT(func(v *App) NetworkRuleSetsResponsePtrOutput { return v.NetworkRuleSets }).(NetworkRuleSetsResponsePtrOutput)
}

func (o AppOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *App) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o AppOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AppOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o AppOutput) Sku() AppSkuInfoResponseOutput {
	return o.ApplyT(func(v *App) AppSkuInfoResponseOutput { return v.Sku }).(AppSkuInfoResponseOutput)
}

func (o AppOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o AppOutput) Subdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.Subdomain }).(pulumi.StringPtrOutput)
}

func (o AppOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *App) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AppOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *App) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AppOutput) Template() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.Template }).(pulumi.StringPtrOutput)
}

func (o AppOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AppOutput{})
}
