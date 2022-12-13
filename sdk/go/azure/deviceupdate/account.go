


package deviceupdate

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Account struct {
	pulumi.CustomResourceState

	HostName                   pulumi.StringOutput                          `pulumi:"hostName"`
	Identity                   ManagedServiceIdentityResponsePtrOutput      `pulumi:"identity"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrOutput                       `pulumi:"publicNetworkAccess"`
	SystemData                 SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
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
			Type: pulumi.String("azure-native:deviceupdate/v20200301preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20220401preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20221001:Account"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20221201preview:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-native:deviceupdate:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:deviceupdate:Account", name, id, state, &resource, opts...)
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
	AccountName                *string                         `pulumi:"accountName"`
	Identity                   *ManagedServiceIdentity         `pulumi:"identity"`
	Location                   *string                         `pulumi:"location"`
	PrivateEndpointConnections []PrivateEndpointConnectionType `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess        *string                         `pulumi:"publicNetworkAccess"`
	ResourceGroupName          string                          `pulumi:"resourceGroupName"`
	Tags                       map[string]string               `pulumi:"tags"`
}


type AccountArgs struct {
	AccountName                pulumi.StringPtrInput
	Identity                   ManagedServiceIdentityPtrInput
	Location                   pulumi.StringPtrInput
	PrivateEndpointConnections PrivateEndpointConnectionTypeArrayInput
	PublicNetworkAccess        pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	Tags                       pulumi.StringMapInput
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

func (o AccountOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

func (o AccountOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Account) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o AccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
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

func (o AccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Account) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
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
