


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EncryptionScope struct {
	pulumi.CustomResourceState

	CreationTime                    pulumi.StringOutput                                `pulumi:"creationTime"`
	KeyVaultProperties              EncryptionScopeKeyVaultPropertiesResponsePtrOutput `pulumi:"keyVaultProperties"`
	LastModifiedTime                pulumi.StringOutput                                `pulumi:"lastModifiedTime"`
	Name                            pulumi.StringOutput                                `pulumi:"name"`
	RequireInfrastructureEncryption pulumi.BoolPtrOutput                               `pulumi:"requireInfrastructureEncryption"`
	Source                          pulumi.StringPtrOutput                             `pulumi:"source"`
	State                           pulumi.StringPtrOutput                             `pulumi:"state"`
	Type                            pulumi.StringOutput                                `pulumi:"type"`
}


func NewEncryptionScope(ctx *pulumi.Context,
	name string, args *EncryptionScopeArgs, opts ...pulumi.ResourceOption) (*EncryptionScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:EncryptionScope"),
		},
	})
	opts = append(opts, aliases)
	var resource EncryptionScope
	err := ctx.RegisterResource("azure-native:storage/v20210801:EncryptionScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEncryptionScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EncryptionScopeState, opts ...pulumi.ResourceOption) (*EncryptionScope, error) {
	var resource EncryptionScope
	err := ctx.ReadResource("azure-native:storage/v20210801:EncryptionScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type encryptionScopeState struct {
}

type EncryptionScopeState struct {
}

func (EncryptionScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptionScopeState)(nil)).Elem()
}

type encryptionScopeArgs struct {
	AccountName                     string                             `pulumi:"accountName"`
	EncryptionScopeName             *string                            `pulumi:"encryptionScopeName"`
	KeyVaultProperties              *EncryptionScopeKeyVaultProperties `pulumi:"keyVaultProperties"`
	RequireInfrastructureEncryption *bool                              `pulumi:"requireInfrastructureEncryption"`
	ResourceGroupName               string                             `pulumi:"resourceGroupName"`
	Source                          *string                            `pulumi:"source"`
	State                           *string                            `pulumi:"state"`
}


type EncryptionScopeArgs struct {
	AccountName                     pulumi.StringInput
	EncryptionScopeName             pulumi.StringPtrInput
	KeyVaultProperties              EncryptionScopeKeyVaultPropertiesPtrInput
	RequireInfrastructureEncryption pulumi.BoolPtrInput
	ResourceGroupName               pulumi.StringInput
	Source                          pulumi.StringPtrInput
	State                           pulumi.StringPtrInput
}

func (EncryptionScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptionScopeArgs)(nil)).Elem()
}

type EncryptionScopeInput interface {
	pulumi.Input

	ToEncryptionScopeOutput() EncryptionScopeOutput
	ToEncryptionScopeOutputWithContext(ctx context.Context) EncryptionScopeOutput
}

func (*EncryptionScope) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionScope)(nil)).Elem()
}

func (i *EncryptionScope) ToEncryptionScopeOutput() EncryptionScopeOutput {
	return i.ToEncryptionScopeOutputWithContext(context.Background())
}

func (i *EncryptionScope) ToEncryptionScopeOutputWithContext(ctx context.Context) EncryptionScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionScopeOutput)
}

type EncryptionScopeOutput struct{ *pulumi.OutputState }

func (EncryptionScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionScope)(nil)).Elem()
}

func (o EncryptionScopeOutput) ToEncryptionScopeOutput() EncryptionScopeOutput {
	return o
}

func (o EncryptionScopeOutput) ToEncryptionScopeOutputWithContext(ctx context.Context) EncryptionScopeOutput {
	return o
}

func (o EncryptionScopeOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o EncryptionScopeOutput) KeyVaultProperties() EncryptionScopeKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionScope) EncryptionScopeKeyVaultPropertiesResponsePtrOutput {
		return v.KeyVaultProperties
	}).(EncryptionScopeKeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionScopeOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o EncryptionScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EncryptionScopeOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.BoolPtrOutput { return v.RequireInfrastructureEncryption }).(pulumi.BoolPtrOutput)
}

func (o EncryptionScopeOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

func (o EncryptionScopeOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o EncryptionScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EncryptionScopeOutput{})
}
