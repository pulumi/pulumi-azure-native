


package v20190601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type BlobContainer struct {
	pulumi.CustomResourceState

	DefaultEncryptionScope      pulumi.StringPtrOutput                     `pulumi:"defaultEncryptionScope"`
	Deleted                     pulumi.BoolOutput                          `pulumi:"deleted"`
	DeletedTime                 pulumi.StringOutput                        `pulumi:"deletedTime"`
	DenyEncryptionScopeOverride pulumi.BoolPtrOutput                       `pulumi:"denyEncryptionScopeOverride"`
	Etag                        pulumi.StringOutput                        `pulumi:"etag"`
	HasImmutabilityPolicy       pulumi.BoolOutput                          `pulumi:"hasImmutabilityPolicy"`
	HasLegalHold                pulumi.BoolOutput                          `pulumi:"hasLegalHold"`
	ImmutabilityPolicy          ImmutabilityPolicyPropertiesResponseOutput `pulumi:"immutabilityPolicy"`
	LastModifiedTime            pulumi.StringOutput                        `pulumi:"lastModifiedTime"`
	LeaseDuration               pulumi.StringOutput                        `pulumi:"leaseDuration"`
	LeaseState                  pulumi.StringOutput                        `pulumi:"leaseState"`
	LeaseStatus                 pulumi.StringOutput                        `pulumi:"leaseStatus"`
	LegalHold                   LegalHoldPropertiesResponseOutput          `pulumi:"legalHold"`
	Metadata                    pulumi.StringMapOutput                     `pulumi:"metadata"`
	Name                        pulumi.StringOutput                        `pulumi:"name"`
	PublicAccess                pulumi.StringPtrOutput                     `pulumi:"publicAccess"`
	RemainingRetentionDays      pulumi.IntOutput                           `pulumi:"remainingRetentionDays"`
	Type                        pulumi.StringOutput                        `pulumi:"type"`
	Version                     pulumi.StringOutput                        `pulumi:"version"`
}


func NewBlobContainer(ctx *pulumi.Context,
	name string, args *BlobContainerArgs, opts ...pulumi.ResourceOption) (*BlobContainer, error) {
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
			Type: pulumi.String("azure-native:storage:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180201:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180301preview:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180701:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20181101:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190401:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:BlobContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobContainer
	err := ctx.RegisterResource("azure-native:storage/v20190601:BlobContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlobContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobContainerState, opts ...pulumi.ResourceOption) (*BlobContainer, error) {
	var resource BlobContainer
	err := ctx.ReadResource("azure-native:storage/v20190601:BlobContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blobContainerState struct {
}

type BlobContainerState struct {
}

func (BlobContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerState)(nil)).Elem()
}

type blobContainerArgs struct {
	AccountName                 string            `pulumi:"accountName"`
	ContainerName               *string           `pulumi:"containerName"`
	DefaultEncryptionScope      *string           `pulumi:"defaultEncryptionScope"`
	DenyEncryptionScopeOverride *bool             `pulumi:"denyEncryptionScopeOverride"`
	Metadata                    map[string]string `pulumi:"metadata"`
	PublicAccess                *PublicAccess     `pulumi:"publicAccess"`
	ResourceGroupName           string            `pulumi:"resourceGroupName"`
}


type BlobContainerArgs struct {
	AccountName                 pulumi.StringInput
	ContainerName               pulumi.StringPtrInput
	DefaultEncryptionScope      pulumi.StringPtrInput
	DenyEncryptionScopeOverride pulumi.BoolPtrInput
	Metadata                    pulumi.StringMapInput
	PublicAccess                PublicAccessPtrInput
	ResourceGroupName           pulumi.StringInput
}

func (BlobContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerArgs)(nil)).Elem()
}

type BlobContainerInput interface {
	pulumi.Input

	ToBlobContainerOutput() BlobContainerOutput
	ToBlobContainerOutputWithContext(ctx context.Context) BlobContainerOutput
}

func (*BlobContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobContainer)(nil)).Elem()
}

func (i *BlobContainer) ToBlobContainerOutput() BlobContainerOutput {
	return i.ToBlobContainerOutputWithContext(context.Background())
}

func (i *BlobContainer) ToBlobContainerOutputWithContext(ctx context.Context) BlobContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobContainerOutput)
}

type BlobContainerOutput struct{ *pulumi.OutputState }

func (BlobContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobContainer)(nil)).Elem()
}

func (o BlobContainerOutput) ToBlobContainerOutput() BlobContainerOutput {
	return o
}

func (o BlobContainerOutput) ToBlobContainerOutputWithContext(ctx context.Context) BlobContainerOutput {
	return o
}

func (o BlobContainerOutput) DefaultEncryptionScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringPtrOutput { return v.DefaultEncryptionScope }).(pulumi.StringPtrOutput)
}

func (o BlobContainerOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.BoolOutput { return v.Deleted }).(pulumi.BoolOutput)
}

func (o BlobContainerOutput) DeletedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.DeletedTime }).(pulumi.StringOutput)
}

func (o BlobContainerOutput) DenyEncryptionScopeOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.BoolPtrOutput { return v.DenyEncryptionScopeOverride }).(pulumi.BoolPtrOutput)
}

func (o BlobContainerOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BlobContainerOutput) HasImmutabilityPolicy() pulumi.BoolOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.BoolOutput { return v.HasImmutabilityPolicy }).(pulumi.BoolOutput)
}

func (o BlobContainerOutput) HasLegalHold() pulumi.BoolOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.BoolOutput { return v.HasLegalHold }).(pulumi.BoolOutput)
}

func (o BlobContainerOutput) ImmutabilityPolicy() ImmutabilityPolicyPropertiesResponseOutput {
	return o.ApplyT(func(v *BlobContainer) ImmutabilityPolicyPropertiesResponseOutput { return v.ImmutabilityPolicy }).(ImmutabilityPolicyPropertiesResponseOutput)
}

func (o BlobContainerOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o BlobContainerOutput) LeaseDuration() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.LeaseDuration }).(pulumi.StringOutput)
}

func (o BlobContainerOutput) LeaseState() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.LeaseState }).(pulumi.StringOutput)
}

func (o BlobContainerOutput) LeaseStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.LeaseStatus }).(pulumi.StringOutput)
}

func (o BlobContainerOutput) LegalHold() LegalHoldPropertiesResponseOutput {
	return o.ApplyT(func(v *BlobContainer) LegalHoldPropertiesResponseOutput { return v.LegalHold }).(LegalHoldPropertiesResponseOutput)
}

func (o BlobContainerOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o BlobContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BlobContainerOutput) PublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringPtrOutput { return v.PublicAccess }).(pulumi.StringPtrOutput)
}

func (o BlobContainerOutput) RemainingRetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.IntOutput { return v.RemainingRetentionDays }).(pulumi.IntOutput)
}

func (o BlobContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o BlobContainerOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BlobContainerOutput{})
}
