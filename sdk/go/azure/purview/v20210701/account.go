


package v20210701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Account struct {
	pulumi.CustomResourceState

	CloudConnectors            CloudConnectorsResponsePtrOutput                `pulumi:"cloudConnectors"`
	CreatedAt                  pulumi.StringOutput                             `pulumi:"createdAt"`
	CreatedBy                  pulumi.StringOutput                             `pulumi:"createdBy"`
	CreatedByObjectId          pulumi.StringOutput                             `pulumi:"createdByObjectId"`
	Endpoints                  AccountPropertiesResponseEndpointsOutput        `pulumi:"endpoints"`
	FriendlyName               pulumi.StringOutput                             `pulumi:"friendlyName"`
	Identity                   IdentityResponsePtrOutput                       `pulumi:"identity"`
	Location                   pulumi.StringPtrOutput                          `pulumi:"location"`
	ManagedResourceGroupName   pulumi.StringPtrOutput                          `pulumi:"managedResourceGroupName"`
	ManagedResources           AccountPropertiesResponseManagedResourcesOutput `pulumi:"managedResources"`
	Name                       pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput    `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                             `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrOutput                          `pulumi:"publicNetworkAccess"`
	Sku                        AccountResponseSkuOutput                        `pulumi:"sku"`
	SystemData                 TrackedResourceResponseSystemDataOutput         `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                          `pulumi:"tags"`
	Type                       pulumi.StringOutput                             `pulumi:"type"`
}


func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.PublicNetworkAccess) {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:purview:Account"),
		},
		{
			Type: pulumi.String("azure-native:purview/v20201201preview:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-native:purview/v20210701:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:purview/v20210701:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type accountState struct {
}

type AccountState struct {
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	AccountName              *string           `pulumi:"accountName"`
	Identity                 *Identity         `pulumi:"identity"`
	Location                 *string           `pulumi:"location"`
	ManagedResourceGroupName *string           `pulumi:"managedResourceGroupName"`
	PublicNetworkAccess      *string           `pulumi:"publicNetworkAccess"`
	ResourceGroupName        string            `pulumi:"resourceGroupName"`
	Tags                     map[string]string `pulumi:"tags"`
}


type AccountArgs struct {
	AccountName              pulumi.StringPtrInput
	Identity                 IdentityPtrInput
	Location                 pulumi.StringPtrInput
	ManagedResourceGroupName pulumi.StringPtrInput
	PublicNetworkAccess      pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	Tags                     pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func (o AccountOutput) CloudConnectors() CloudConnectorsResponsePtrOutput {
	return o.ApplyT(func(v *Account) CloudConnectorsResponsePtrOutput { return v.CloudConnectors }).(CloudConnectorsResponsePtrOutput)
}

func (o AccountOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o AccountOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o AccountOutput) CreatedByObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CreatedByObjectId }).(pulumi.StringOutput)
}

func (o AccountOutput) Endpoints() AccountPropertiesResponseEndpointsOutput {
	return o.ApplyT(func(v *Account) AccountPropertiesResponseEndpointsOutput { return v.Endpoints }).(AccountPropertiesResponseEndpointsOutput)
}

func (o AccountOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.FriendlyName }).(pulumi.StringOutput)
}

func (o AccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Account) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o AccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AccountOutput) ManagedResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.ManagedResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o AccountOutput) ManagedResources() AccountPropertiesResponseManagedResourcesOutput {
	return o.ApplyT(func(v *Account) AccountPropertiesResponseManagedResourcesOutput { return v.ManagedResources }).(AccountPropertiesResponseManagedResourcesOutput)
}

func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AccountOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Account) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o AccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AccountOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o AccountOutput) Sku() AccountResponseSkuOutput {
	return o.ApplyT(func(v *Account) AccountResponseSkuOutput { return v.Sku }).(AccountResponseSkuOutput)
}

func (o AccountOutput) SystemData() TrackedResourceResponseSystemDataOutput {
	return o.ApplyT(func(v *Account) TrackedResourceResponseSystemDataOutput { return v.SystemData }).(TrackedResourceResponseSystemDataOutput)
}

func (o AccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Account) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
