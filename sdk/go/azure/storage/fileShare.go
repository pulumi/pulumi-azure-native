


package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FileShare struct {
	pulumi.CustomResourceState

	AccessTier             pulumi.StringPtrOutput `pulumi:"accessTier"`
	AccessTierChangeTime   pulumi.StringOutput    `pulumi:"accessTierChangeTime"`
	AccessTierStatus       pulumi.StringOutput    `pulumi:"accessTierStatus"`
	Deleted                pulumi.BoolOutput      `pulumi:"deleted"`
	DeletedTime            pulumi.StringOutput    `pulumi:"deletedTime"`
	EnabledProtocols       pulumi.StringPtrOutput `pulumi:"enabledProtocols"`
	Etag                   pulumi.StringOutput    `pulumi:"etag"`
	LastModifiedTime       pulumi.StringOutput    `pulumi:"lastModifiedTime"`
	Metadata               pulumi.StringMapOutput `pulumi:"metadata"`
	Name                   pulumi.StringOutput    `pulumi:"name"`
	RemainingRetentionDays pulumi.IntOutput       `pulumi:"remainingRetentionDays"`
	RootSquash             pulumi.StringPtrOutput `pulumi:"rootSquash"`
	ShareQuota             pulumi.IntPtrOutput    `pulumi:"shareQuota"`
	ShareUsageBytes        pulumi.Float64Output   `pulumi:"shareUsageBytes"`
	SnapshotTime           pulumi.StringOutput    `pulumi:"snapshotTime"`
	Type                   pulumi.StringOutput    `pulumi:"type"`
	Version                pulumi.StringOutput    `pulumi:"version"`
}


func NewFileShare(ctx *pulumi.Context,
	name string, args *FileShareArgs, opts ...pulumi.ResourceOption) (*FileShare, error) {
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
			Type: pulumi.String("azure-native:storage/v20190401:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:FileShare"),
		},
	})
	opts = append(opts, aliases)
	var resource FileShare
	err := ctx.RegisterResource("azure-native:storage:FileShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFileShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileShareState, opts ...pulumi.ResourceOption) (*FileShare, error) {
	var resource FileShare
	err := ctx.ReadResource("azure-native:storage:FileShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fileShareState struct {
}

type FileShareState struct {
}

func (FileShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileShareState)(nil)).Elem()
}

type fileShareArgs struct {
	AccessTier        *string           `pulumi:"accessTier"`
	AccountName       string            `pulumi:"accountName"`
	EnabledProtocols  *string           `pulumi:"enabledProtocols"`
	Expand            *string           `pulumi:"expand"`
	Metadata          map[string]string `pulumi:"metadata"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	RootSquash        *string           `pulumi:"rootSquash"`
	ShareName         *string           `pulumi:"shareName"`
	ShareQuota        *int              `pulumi:"shareQuota"`
}


type FileShareArgs struct {
	AccessTier        pulumi.StringPtrInput
	AccountName       pulumi.StringInput
	EnabledProtocols  pulumi.StringPtrInput
	Expand            pulumi.StringPtrInput
	Metadata          pulumi.StringMapInput
	ResourceGroupName pulumi.StringInput
	RootSquash        pulumi.StringPtrInput
	ShareName         pulumi.StringPtrInput
	ShareQuota        pulumi.IntPtrInput
}

func (FileShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileShareArgs)(nil)).Elem()
}

type FileShareInput interface {
	pulumi.Input

	ToFileShareOutput() FileShareOutput
	ToFileShareOutputWithContext(ctx context.Context) FileShareOutput
}

func (*FileShare) ElementType() reflect.Type {
	return reflect.TypeOf((**FileShare)(nil)).Elem()
}

func (i *FileShare) ToFileShareOutput() FileShareOutput {
	return i.ToFileShareOutputWithContext(context.Background())
}

func (i *FileShare) ToFileShareOutputWithContext(ctx context.Context) FileShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileShareOutput)
}

type FileShareOutput struct{ *pulumi.OutputState }

func (FileShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileShare)(nil)).Elem()
}

func (o FileShareOutput) ToFileShareOutput() FileShareOutput {
	return o
}

func (o FileShareOutput) ToFileShareOutputWithContext(ctx context.Context) FileShareOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FileShareOutput{})
}
