


package v20210622

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutomationAccount struct {
	pulumi.CustomResourceState

	AutomationHybridServiceUrl pulumi.StringPtrOutput                       `pulumi:"automationHybridServiceUrl"`
	CreationTime               pulumi.StringOutput                          `pulumi:"creationTime"`
	Description                pulumi.StringPtrOutput                       `pulumi:"description"`
	DisableLocalAuth           pulumi.BoolPtrOutput                         `pulumi:"disableLocalAuth"`
	Encryption                 EncryptionPropertiesResponsePtrOutput        `pulumi:"encryption"`
	Etag                       pulumi.StringPtrOutput                       `pulumi:"etag"`
	Identity                   IdentityResponsePtrOutput                    `pulumi:"identity"`
	LastModifiedBy             pulumi.StringPtrOutput                       `pulumi:"lastModifiedBy"`
	LastModifiedTime           pulumi.StringOutput                          `pulumi:"lastModifiedTime"`
	Location                   pulumi.StringPtrOutput                       `pulumi:"location"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess        pulumi.BoolPtrOutput                         `pulumi:"publicNetworkAccess"`
	Sku                        SkuResponsePtrOutput                         `pulumi:"sku"`
	State                      pulumi.StringOutput                          `pulumi:"state"`
	SystemData                 SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
}


func NewAutomationAccount(ctx *pulumi.Context,
	name string, args *AutomationAccountArgs, opts ...pulumi.ResourceOption) (*AutomationAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:AutomationAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource AutomationAccount
	err := ctx.RegisterResource("azure-native:automation/v20210622:AutomationAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAutomationAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationAccountState, opts ...pulumi.ResourceOption) (*AutomationAccount, error) {
	var resource AutomationAccount
	err := ctx.ReadResource("azure-native:automation/v20210622:AutomationAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type automationAccountState struct {
}

type AutomationAccountState struct {
}

func (AutomationAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountState)(nil)).Elem()
}

type automationAccountArgs struct {
	AutomationAccountName *string               `pulumi:"automationAccountName"`
	DisableLocalAuth      *bool                 `pulumi:"disableLocalAuth"`
	Encryption            *EncryptionProperties `pulumi:"encryption"`
	Identity              *Identity             `pulumi:"identity"`
	Location              *string               `pulumi:"location"`
	Name                  *string               `pulumi:"name"`
	PublicNetworkAccess   *bool                 `pulumi:"publicNetworkAccess"`
	ResourceGroupName     string                `pulumi:"resourceGroupName"`
	Sku                   *Sku                  `pulumi:"sku"`
	Tags                  map[string]string     `pulumi:"tags"`
}


type AutomationAccountArgs struct {
	AutomationAccountName pulumi.StringPtrInput
	DisableLocalAuth      pulumi.BoolPtrInput
	Encryption            EncryptionPropertiesPtrInput
	Identity              IdentityPtrInput
	Location              pulumi.StringPtrInput
	Name                  pulumi.StringPtrInput
	PublicNetworkAccess   pulumi.BoolPtrInput
	ResourceGroupName     pulumi.StringInput
	Sku                   SkuPtrInput
	Tags                  pulumi.StringMapInput
}

func (AutomationAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountArgs)(nil)).Elem()
}

type AutomationAccountInput interface {
	pulumi.Input

	ToAutomationAccountOutput() AutomationAccountOutput
	ToAutomationAccountOutputWithContext(ctx context.Context) AutomationAccountOutput
}

func (*AutomationAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationAccount)(nil)).Elem()
}

func (i *AutomationAccount) ToAutomationAccountOutput() AutomationAccountOutput {
	return i.ToAutomationAccountOutputWithContext(context.Background())
}

func (i *AutomationAccount) ToAutomationAccountOutputWithContext(ctx context.Context) AutomationAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationAccountOutput)
}

type AutomationAccountOutput struct{ *pulumi.OutputState }

func (AutomationAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationAccount)(nil)).Elem()
}

func (o AutomationAccountOutput) ToAutomationAccountOutput() AutomationAccountOutput {
	return o
}

func (o AutomationAccountOutput) ToAutomationAccountOutputWithContext(ctx context.Context) AutomationAccountOutput {
	return o
}

func (o AutomationAccountOutput) AutomationHybridServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringPtrOutput { return v.AutomationHybridServiceUrl }).(pulumi.StringPtrOutput)
}

func (o AutomationAccountOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o AutomationAccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AutomationAccountOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o AutomationAccountOutput) Encryption() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AutomationAccount) EncryptionPropertiesResponsePtrOutput { return v.Encryption }).(EncryptionPropertiesResponsePtrOutput)
}

func (o AutomationAccountOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o AutomationAccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *AutomationAccount) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o AutomationAccountOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o AutomationAccountOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o AutomationAccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AutomationAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AutomationAccountOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *AutomationAccount) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o AutomationAccountOutput) PublicNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.BoolPtrOutput { return v.PublicNetworkAccess }).(pulumi.BoolPtrOutput)
}

func (o AutomationAccountOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *AutomationAccount) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o AutomationAccountOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o AutomationAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AutomationAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AutomationAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AutomationAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AutomationAccountOutput{})
}
