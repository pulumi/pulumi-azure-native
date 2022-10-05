


package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Blob struct {
	pulumi.CustomResourceState

	AccessTier  BlobAccessTierOutput   `pulumi:"accessTier"`
	ContentMd5  pulumi.StringPtrOutput `pulumi:"contentMd5"`
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	Metadata    pulumi.StringMapOutput `pulumi:"metadata"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Type        BlobTypeOutput         `pulumi:"type"`
	Url         pulumi.StringOutput    `pulumi:"url"`
}


func NewBlob(ctx *pulumi.Context,
	name string, args *BlobArgs, opts ...pulumi.ResourceOption) (*Blob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.Type) {
		args.Type = BlobType("Block")
	}
	var resource Blob
	err := ctx.RegisterResource("azure-native:storage:Blob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobState, opts ...pulumi.ResourceOption) (*Blob, error) {
	var resource Blob
	err := ctx.ReadResource("azure-native:storage:Blob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blobState struct {
}

type BlobState struct {
}

func (BlobState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobState)(nil)).Elem()
}

type blobArgs struct {
	AccessTier        *BlobAccessTier       `pulumi:"accessTier"`
	AccountName       string                `pulumi:"accountName"`
	BlobName          *string               `pulumi:"blobName"`
	ContainerName     string                `pulumi:"containerName"`
	ContentMd5        *string               `pulumi:"contentMd5"`
	ContentType       *string               `pulumi:"contentType"`
	Metadata          map[string]string     `pulumi:"metadata"`
	ResourceGroupName string                `pulumi:"resourceGroupName"`
	Source            pulumi.AssetOrArchive `pulumi:"source"`
	Type              *BlobType             `pulumi:"type"`
}


type BlobArgs struct {
	AccessTier        BlobAccessTierPtrInput
	AccountName       pulumi.StringInput
	BlobName          pulumi.StringPtrInput
	ContainerName     pulumi.StringInput
	ContentMd5        pulumi.StringPtrInput
	ContentType       pulumi.StringPtrInput
	Metadata          pulumi.StringMapInput
	ResourceGroupName pulumi.StringInput
	Source            pulumi.AssetOrArchiveInput
	Type              BlobTypePtrInput
}

func (BlobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobArgs)(nil)).Elem()
}

type BlobInput interface {
	pulumi.Input

	ToBlobOutput() BlobOutput
	ToBlobOutputWithContext(ctx context.Context) BlobOutput
}

func (*Blob) ElementType() reflect.Type {
	return reflect.TypeOf((**Blob)(nil)).Elem()
}

func (i *Blob) ToBlobOutput() BlobOutput {
	return i.ToBlobOutputWithContext(context.Background())
}

func (i *Blob) ToBlobOutputWithContext(ctx context.Context) BlobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobOutput)
}

type BlobOutput struct{ *pulumi.OutputState }

func (BlobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Blob)(nil)).Elem()
}

func (o BlobOutput) ToBlobOutput() BlobOutput {
	return o
}

func (o BlobOutput) ToBlobOutputWithContext(ctx context.Context) BlobOutput {
	return o
}

func (o BlobOutput) AccessTier() BlobAccessTierOutput {
	return o.ApplyT(func(v *Blob) BlobAccessTierOutput { return v.AccessTier }).(BlobAccessTierOutput)
}

func (o BlobOutput) ContentMd5() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Blob) pulumi.StringPtrOutput { return v.ContentMd5 }).(pulumi.StringPtrOutput)
}

func (o BlobOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Blob) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o BlobOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Blob) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o BlobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Blob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BlobOutput) Type() BlobTypeOutput {
	return o.ApplyT(func(v *Blob) BlobTypeOutput { return v.Type }).(BlobTypeOutput)
}

func (o BlobOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Blob) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BlobOutput{})
}
