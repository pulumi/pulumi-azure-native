


package v20180701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Asset struct {
	pulumi.CustomResourceState

	AlternateId             pulumi.StringPtrOutput `pulumi:"alternateId"`
	AssetId                 pulumi.StringOutput    `pulumi:"assetId"`
	Container               pulumi.StringPtrOutput `pulumi:"container"`
	Created                 pulumi.StringOutput    `pulumi:"created"`
	Description             pulumi.StringPtrOutput `pulumi:"description"`
	LastModified            pulumi.StringOutput    `pulumi:"lastModified"`
	Name                    pulumi.StringOutput    `pulumi:"name"`
	StorageAccountName      pulumi.StringPtrOutput `pulumi:"storageAccountName"`
	StorageEncryptionFormat pulumi.StringOutput    `pulumi:"storageEncryptionFormat"`
	Type                    pulumi.StringOutput    `pulumi:"type"`
}


func NewAsset(ctx *pulumi.Context,
	name string, args *AssetArgs, opts ...pulumi.ResourceOption) (*Asset, error) {
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
			Type: pulumi.String("azure-native:media:Asset"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:Asset"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:Asset"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:Asset"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:Asset"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:Asset"),
		},
	})
	opts = append(opts, aliases)
	var resource Asset
	err := ctx.RegisterResource("azure-native:media/v20180701:Asset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAsset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssetState, opts ...pulumi.ResourceOption) (*Asset, error) {
	var resource Asset
	err := ctx.ReadResource("azure-native:media/v20180701:Asset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type assetState struct {
}

type AssetState struct {
}

func (AssetState) ElementType() reflect.Type {
	return reflect.TypeOf((*assetState)(nil)).Elem()
}

type assetArgs struct {
	AccountName        string  `pulumi:"accountName"`
	AlternateId        *string `pulumi:"alternateId"`
	AssetName          *string `pulumi:"assetName"`
	Container          *string `pulumi:"container"`
	Description        *string `pulumi:"description"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	StorageAccountName *string `pulumi:"storageAccountName"`
}


type AssetArgs struct {
	AccountName        pulumi.StringInput
	AlternateId        pulumi.StringPtrInput
	AssetName          pulumi.StringPtrInput
	Container          pulumi.StringPtrInput
	Description        pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	StorageAccountName pulumi.StringPtrInput
}

func (AssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assetArgs)(nil)).Elem()
}

type AssetInput interface {
	pulumi.Input

	ToAssetOutput() AssetOutput
	ToAssetOutputWithContext(ctx context.Context) AssetOutput
}

func (*Asset) ElementType() reflect.Type {
	return reflect.TypeOf((**Asset)(nil)).Elem()
}

func (i *Asset) ToAssetOutput() AssetOutput {
	return i.ToAssetOutputWithContext(context.Background())
}

func (i *Asset) ToAssetOutputWithContext(ctx context.Context) AssetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetOutput)
}

type AssetOutput struct{ *pulumi.OutputState }

func (AssetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Asset)(nil)).Elem()
}

func (o AssetOutput) ToAssetOutput() AssetOutput {
	return o
}

func (o AssetOutput) ToAssetOutputWithContext(ctx context.Context) AssetOutput {
	return o
}

func (o AssetOutput) AlternateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.AlternateId }).(pulumi.StringPtrOutput)
}

func (o AssetOutput) AssetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.AssetId }).(pulumi.StringOutput)
}

func (o AssetOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.Container }).(pulumi.StringPtrOutput)
}

func (o AssetOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

func (o AssetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AssetOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

func (o AssetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AssetOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

func (o AssetOutput) StorageEncryptionFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.StorageEncryptionFormat }).(pulumi.StringOutput)
}

func (o AssetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssetOutput{})
}
