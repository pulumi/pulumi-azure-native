


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlobFolderDataSet struct {
	pulumi.CustomResourceState

	ContainerName      pulumi.StringOutput      `pulumi:"containerName"`
	DataSetId          pulumi.StringOutput      `pulumi:"dataSetId"`
	Kind               pulumi.StringOutput      `pulumi:"kind"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	Prefix             pulumi.StringOutput      `pulumi:"prefix"`
	ResourceGroup      pulumi.StringOutput      `pulumi:"resourceGroup"`
	StorageAccountName pulumi.StringOutput      `pulumi:"storageAccountName"`
	SubscriptionId     pulumi.StringOutput      `pulumi:"subscriptionId"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewBlobFolderDataSet(ctx *pulumi.Context,
	name string, args *BlobFolderDataSetArgs, opts ...pulumi.ResourceOption) (*BlobFolderDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Prefix == nil {
		return nil, errors.New("invalid value for required argument 'Prefix'")
	}
	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	args.Kind = pulumi.String("BlobFolder")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:BlobFolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:BlobFolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:BlobFolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:BlobFolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:BlobFolderDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobFolderDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20210801:BlobFolderDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlobFolderDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobFolderDataSetState, opts ...pulumi.ResourceOption) (*BlobFolderDataSet, error) {
	var resource BlobFolderDataSet
	err := ctx.ReadResource("azure-native:datashare/v20210801:BlobFolderDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blobFolderDataSetState struct {
}

type BlobFolderDataSetState struct {
}

func (BlobFolderDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobFolderDataSetState)(nil)).Elem()
}

type blobFolderDataSetArgs struct {
	AccountName        string  `pulumi:"accountName"`
	ContainerName      string  `pulumi:"containerName"`
	DataSetName        *string `pulumi:"dataSetName"`
	Kind               string  `pulumi:"kind"`
	Prefix             string  `pulumi:"prefix"`
	ResourceGroup      string  `pulumi:"resourceGroup"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	ShareName          string  `pulumi:"shareName"`
	StorageAccountName string  `pulumi:"storageAccountName"`
	SubscriptionId     string  `pulumi:"subscriptionId"`
}


type BlobFolderDataSetArgs struct {
	AccountName        pulumi.StringInput
	ContainerName      pulumi.StringInput
	DataSetName        pulumi.StringPtrInput
	Kind               pulumi.StringInput
	Prefix             pulumi.StringInput
	ResourceGroup      pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	ShareName          pulumi.StringInput
	StorageAccountName pulumi.StringInput
	SubscriptionId     pulumi.StringInput
}

func (BlobFolderDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobFolderDataSetArgs)(nil)).Elem()
}

type BlobFolderDataSetInput interface {
	pulumi.Input

	ToBlobFolderDataSetOutput() BlobFolderDataSetOutput
	ToBlobFolderDataSetOutputWithContext(ctx context.Context) BlobFolderDataSetOutput
}

func (*BlobFolderDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobFolderDataSet)(nil)).Elem()
}

func (i *BlobFolderDataSet) ToBlobFolderDataSetOutput() BlobFolderDataSetOutput {
	return i.ToBlobFolderDataSetOutputWithContext(context.Background())
}

func (i *BlobFolderDataSet) ToBlobFolderDataSetOutputWithContext(ctx context.Context) BlobFolderDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobFolderDataSetOutput)
}

type BlobFolderDataSetOutput struct{ *pulumi.OutputState }

func (BlobFolderDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobFolderDataSet)(nil)).Elem()
}

func (o BlobFolderDataSetOutput) ToBlobFolderDataSetOutput() BlobFolderDataSetOutput {
	return o
}

func (o BlobFolderDataSetOutput) ToBlobFolderDataSetOutputWithContext(ctx context.Context) BlobFolderDataSetOutput {
	return o
}

func (o BlobFolderDataSetOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobFolderDataSet) pulumi.StringOutput { return v.ContainerName }).(pulumi.StringOutput)
}

func (o BlobFolderDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobFolderDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o BlobFolderDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobFolderDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o BlobFolderDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobFolderDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BlobFolderDataSetOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobFolderDataSet) pulumi.StringOutput { return v.Prefix }).(pulumi.StringOutput)
}

func (o BlobFolderDataSetOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobFolderDataSet) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o BlobFolderDataSetOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobFolderDataSet) pulumi.StringOutput { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o BlobFolderDataSetOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobFolderDataSet) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o BlobFolderDataSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BlobFolderDataSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BlobFolderDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobFolderDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BlobFolderDataSetOutput{})
}
