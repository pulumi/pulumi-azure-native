


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationStore struct {
	pulumi.CustomResourceState

	CreationDate               pulumi.StringOutput                                   `pulumi:"creationDate"`
	DisableLocalAuth           pulumi.BoolPtrOutput                                  `pulumi:"disableLocalAuth"`
	EnablePurgeProtection      pulumi.BoolPtrOutput                                  `pulumi:"enablePurgeProtection"`
	Encryption                 EncryptionPropertiesResponsePtrOutput                 `pulumi:"encryption"`
	Endpoint                   pulumi.StringOutput                                   `pulumi:"endpoint"`
	Identity                   ResourceIdentityResponsePtrOutput                     `pulumi:"identity"`
	Location                   pulumi.StringOutput                                   `pulumi:"location"`
	Name                       pulumi.StringOutput                                   `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionReferenceResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                                   `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrOutput                                `pulumi:"publicNetworkAccess"`
	Sku                        SkuResponseOutput                                     `pulumi:"sku"`
	SoftDeleteRetentionInDays  pulumi.IntPtrOutput                                   `pulumi:"softDeleteRetentionInDays"`
	SystemData                 SystemDataResponseOutput                              `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                                `pulumi:"tags"`
	Type                       pulumi.StringOutput                                   `pulumi:"type"`
}


func NewConfigurationStore(ctx *pulumi.Context,
	name string, args *ConfigurationStoreArgs, opts ...pulumi.ResourceOption) (*ConfigurationStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if isZero(args.DisableLocalAuth) {
		args.DisableLocalAuth = pulumi.BoolPtr(false)
	}
	if isZero(args.EnablePurgeProtection) {
		args.EnablePurgeProtection = pulumi.BoolPtr(false)
	}
	if isZero(args.SoftDeleteRetentionInDays) {
		args.SoftDeleteRetentionInDays = pulumi.IntPtr(7)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appconfiguration:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20190201preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20191001:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20191101preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20200601:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20200701preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20210301preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20220301preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20220501:ConfigurationStore"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationStore
	err := ctx.RegisterResource("azure-native:appconfiguration/v20211001preview:ConfigurationStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationStoreState, opts ...pulumi.ResourceOption) (*ConfigurationStore, error) {
	var resource ConfigurationStore
	err := ctx.ReadResource("azure-native:appconfiguration/v20211001preview:ConfigurationStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationStoreState struct {
}

type ConfigurationStoreState struct {
}

func (ConfigurationStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationStoreState)(nil)).Elem()
}

type configurationStoreArgs struct {
	ConfigStoreName           *string               `pulumi:"configStoreName"`
	CreateMode                *CreateMode           `pulumi:"createMode"`
	DisableLocalAuth          *bool                 `pulumi:"disableLocalAuth"`
	EnablePurgeProtection     *bool                 `pulumi:"enablePurgeProtection"`
	Encryption                *EncryptionProperties `pulumi:"encryption"`
	Identity                  *ResourceIdentity     `pulumi:"identity"`
	Location                  *string               `pulumi:"location"`
	PublicNetworkAccess       *string               `pulumi:"publicNetworkAccess"`
	ResourceGroupName         string                `pulumi:"resourceGroupName"`
	Sku                       Sku                   `pulumi:"sku"`
	SoftDeleteRetentionInDays *int                  `pulumi:"softDeleteRetentionInDays"`
	Tags                      map[string]string     `pulumi:"tags"`
}


type ConfigurationStoreArgs struct {
	ConfigStoreName           pulumi.StringPtrInput
	CreateMode                CreateModePtrInput
	DisableLocalAuth          pulumi.BoolPtrInput
	EnablePurgeProtection     pulumi.BoolPtrInput
	Encryption                EncryptionPropertiesPtrInput
	Identity                  ResourceIdentityPtrInput
	Location                  pulumi.StringPtrInput
	PublicNetworkAccess       pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Sku                       SkuInput
	SoftDeleteRetentionInDays pulumi.IntPtrInput
	Tags                      pulumi.StringMapInput
}

func (ConfigurationStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationStoreArgs)(nil)).Elem()
}

type ConfigurationStoreInput interface {
	pulumi.Input

	ToConfigurationStoreOutput() ConfigurationStoreOutput
	ToConfigurationStoreOutputWithContext(ctx context.Context) ConfigurationStoreOutput
}

func (*ConfigurationStore) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationStore)(nil)).Elem()
}

func (i *ConfigurationStore) ToConfigurationStoreOutput() ConfigurationStoreOutput {
	return i.ToConfigurationStoreOutputWithContext(context.Background())
}

func (i *ConfigurationStore) ToConfigurationStoreOutputWithContext(ctx context.Context) ConfigurationStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationStoreOutput)
}

type ConfigurationStoreOutput struct{ *pulumi.OutputState }

func (ConfigurationStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationStore)(nil)).Elem()
}

func (o ConfigurationStoreOutput) ToConfigurationStoreOutput() ConfigurationStoreOutput {
	return o
}

func (o ConfigurationStoreOutput) ToConfigurationStoreOutputWithContext(ctx context.Context) ConfigurationStoreOutput {
	return o
}

func (o ConfigurationStoreOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o ConfigurationStoreOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o ConfigurationStoreOutput) EnablePurgeProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.BoolPtrOutput { return v.EnablePurgeProtection }).(pulumi.BoolPtrOutput)
}

func (o ConfigurationStoreOutput) Encryption() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationStore) EncryptionPropertiesResponsePtrOutput { return v.Encryption }).(EncryptionPropertiesResponsePtrOutput)
}

func (o ConfigurationStoreOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

func (o ConfigurationStoreOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationStore) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o ConfigurationStoreOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ConfigurationStoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationStoreOutput) PrivateEndpointConnections() PrivateEndpointConnectionReferenceResponseArrayOutput {
	return o.ApplyT(func(v *ConfigurationStore) PrivateEndpointConnectionReferenceResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionReferenceResponseArrayOutput)
}

func (o ConfigurationStoreOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConfigurationStoreOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o ConfigurationStoreOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *ConfigurationStore) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o ConfigurationStoreOutput) SoftDeleteRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.IntPtrOutput { return v.SoftDeleteRetentionInDays }).(pulumi.IntPtrOutput)
}

func (o ConfigurationStoreOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationStore) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConfigurationStoreOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ConfigurationStoreOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationStoreOutput{})
}
