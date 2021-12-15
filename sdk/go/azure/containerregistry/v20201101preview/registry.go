


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Registry struct {
	pulumi.CustomResourceState

	AdminUserEnabled           pulumi.BoolPtrOutput                         `pulumi:"adminUserEnabled"`
	AnonymousPullEnabled       pulumi.BoolPtrOutput                         `pulumi:"anonymousPullEnabled"`
	CreationDate               pulumi.StringOutput                          `pulumi:"creationDate"`
	DataEndpointEnabled        pulumi.BoolPtrOutput                         `pulumi:"dataEndpointEnabled"`
	DataEndpointHostNames      pulumi.StringArrayOutput                     `pulumi:"dataEndpointHostNames"`
	Encryption                 EncryptionPropertyResponsePtrOutput          `pulumi:"encryption"`
	Identity                   IdentityPropertiesResponsePtrOutput          `pulumi:"identity"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	LoginServer                pulumi.StringOutput                          `pulumi:"loginServer"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	NetworkRuleBypassOptions   pulumi.StringPtrOutput                       `pulumi:"networkRuleBypassOptions"`
	NetworkRuleSet             NetworkRuleSetResponsePtrOutput              `pulumi:"networkRuleSet"`
	Policies                   PoliciesResponsePtrOutput                    `pulumi:"policies"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrOutput                       `pulumi:"publicNetworkAccess"`
	Sku                        SkuResponseOutput                            `pulumi:"sku"`
	Status                     StatusResponseOutput                         `pulumi:"status"`
	SystemData                 SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
	ZoneRedundancy             pulumi.StringPtrOutput                       `pulumi:"zoneRedundancy"`
}


func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOption) (*Registry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if isZero(args.AdminUserEnabled) {
		args.AdminUserEnabled = pulumi.BoolPtr(false)
	}
	if isZero(args.AnonymousPullEnabled) {
		args.AnonymousPullEnabled = pulumi.BoolPtr(false)
	}
	if isZero(args.NetworkRuleBypassOptions) {
		args.NetworkRuleBypassOptions = pulumi.StringPtr("AzureServices")
	}
	if args.NetworkRuleSet != nil {
		args.NetworkRuleSet = args.NetworkRuleSet.ToNetworkRuleSetPtrOutput().ApplyT(func(v *NetworkRuleSet) *NetworkRuleSet { return v.Defaults() }).(NetworkRuleSetPtrOutput)
	}
	if args.Policies != nil {
		args.Policies = args.Policies.ToPoliciesPtrOutput().ApplyT(func(v *Policies) *Policies { return v.Defaults() }).(PoliciesPtrOutput)
	}
	if isZero(args.PublicNetworkAccess) {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	if isZero(args.ZoneRedundancy) {
		args.ZoneRedundancy = pulumi.StringPtr("Disabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20160627preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20170301:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20170601preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20171001:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20190501:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20191201preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210901:Registry"),
		},
	})
	opts = append(opts, aliases)
	var resource Registry
	err := ctx.RegisterResource("azure-native:containerregistry/v20201101preview:Registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryState, opts ...pulumi.ResourceOption) (*Registry, error) {
	var resource Registry
	err := ctx.ReadResource("azure-native:containerregistry/v20201101preview:Registry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registryState struct {
}

type RegistryState struct {
}

func (RegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryState)(nil)).Elem()
}

type registryArgs struct {
	AdminUserEnabled         *bool               `pulumi:"adminUserEnabled"`
	AnonymousPullEnabled     *bool               `pulumi:"anonymousPullEnabled"`
	DataEndpointEnabled      *bool               `pulumi:"dataEndpointEnabled"`
	Encryption               *EncryptionProperty `pulumi:"encryption"`
	Identity                 *IdentityProperties `pulumi:"identity"`
	Location                 *string             `pulumi:"location"`
	NetworkRuleBypassOptions *string             `pulumi:"networkRuleBypassOptions"`
	NetworkRuleSet           *NetworkRuleSet     `pulumi:"networkRuleSet"`
	Policies                 *Policies           `pulumi:"policies"`
	PublicNetworkAccess      *string             `pulumi:"publicNetworkAccess"`
	RegistryName             *string             `pulumi:"registryName"`
	ResourceGroupName        string              `pulumi:"resourceGroupName"`
	Sku                      Sku                 `pulumi:"sku"`
	Tags                     map[string]string   `pulumi:"tags"`
	ZoneRedundancy           *string             `pulumi:"zoneRedundancy"`
}


type RegistryArgs struct {
	AdminUserEnabled         pulumi.BoolPtrInput
	AnonymousPullEnabled     pulumi.BoolPtrInput
	DataEndpointEnabled      pulumi.BoolPtrInput
	Encryption               EncryptionPropertyPtrInput
	Identity                 IdentityPropertiesPtrInput
	Location                 pulumi.StringPtrInput
	NetworkRuleBypassOptions pulumi.StringPtrInput
	NetworkRuleSet           NetworkRuleSetPtrInput
	Policies                 PoliciesPtrInput
	PublicNetworkAccess      pulumi.StringPtrInput
	RegistryName             pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	Sku                      SkuInput
	Tags                     pulumi.StringMapInput
	ZoneRedundancy           pulumi.StringPtrInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}

type RegistryInput interface {
	pulumi.Input

	ToRegistryOutput() RegistryOutput
	ToRegistryOutputWithContext(ctx context.Context) RegistryOutput
}

func (*Registry) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil)).Elem()
}

func (i *Registry) ToRegistryOutput() RegistryOutput {
	return i.ToRegistryOutputWithContext(context.Background())
}

func (i *Registry) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryOutput)
}

type RegistryOutput struct{ *pulumi.OutputState }

func (RegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil)).Elem()
}

func (o RegistryOutput) ToRegistryOutput() RegistryOutput {
	return o
}

func (o RegistryOutput) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegistryOutput{})
}
