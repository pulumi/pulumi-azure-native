


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlobStorageAccountDataSet struct {
	pulumi.CustomResourceState

	DataSetId                pulumi.StringOutput                       `pulumi:"dataSetId"`
	Kind                     pulumi.StringOutput                       `pulumi:"kind"`
	Location                 pulumi.StringOutput                       `pulumi:"location"`
	Name                     pulumi.StringOutput                       `pulumi:"name"`
	Paths                    BlobStorageAccountPathResponseArrayOutput `pulumi:"paths"`
	StorageAccountResourceId pulumi.StringOutput                       `pulumi:"storageAccountResourceId"`
	SystemData               SystemDataResponseOutput                  `pulumi:"systemData"`
	Type                     pulumi.StringOutput                       `pulumi:"type"`
}


func NewBlobStorageAccountDataSet(ctx *pulumi.Context,
	name string, args *BlobStorageAccountDataSetArgs, opts ...pulumi.ResourceOption) (*BlobStorageAccountDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Paths == nil {
		return nil, errors.New("invalid value for required argument 'Paths'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.StorageAccountResourceId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountResourceId'")
	}
	args.Kind = pulumi.String("BlobStorageAccount")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:BlobStorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:BlobStorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:BlobStorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:BlobStorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:BlobStorageAccountDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobStorageAccountDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:BlobStorageAccountDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlobStorageAccountDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobStorageAccountDataSetState, opts ...pulumi.ResourceOption) (*BlobStorageAccountDataSet, error) {
	var resource BlobStorageAccountDataSet
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:BlobStorageAccountDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blobStorageAccountDataSetState struct {
}

type BlobStorageAccountDataSetState struct {
}

func (BlobStorageAccountDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobStorageAccountDataSetState)(nil)).Elem()
}

type blobStorageAccountDataSetArgs struct {
	AccountName              string                   `pulumi:"accountName"`
	DataSetName              *string                  `pulumi:"dataSetName"`
	Kind                     string                   `pulumi:"kind"`
	Paths                    []BlobStorageAccountPath `pulumi:"paths"`
	ResourceGroupName        string                   `pulumi:"resourceGroupName"`
	ShareName                string                   `pulumi:"shareName"`
	StorageAccountResourceId string                   `pulumi:"storageAccountResourceId"`
}


type BlobStorageAccountDataSetArgs struct {
	AccountName              pulumi.StringInput
	DataSetName              pulumi.StringPtrInput
	Kind                     pulumi.StringInput
	Paths                    BlobStorageAccountPathArrayInput
	ResourceGroupName        pulumi.StringInput
	ShareName                pulumi.StringInput
	StorageAccountResourceId pulumi.StringInput
}

func (BlobStorageAccountDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobStorageAccountDataSetArgs)(nil)).Elem()
}

type BlobStorageAccountDataSetInput interface {
	pulumi.Input

	ToBlobStorageAccountDataSetOutput() BlobStorageAccountDataSetOutput
	ToBlobStorageAccountDataSetOutputWithContext(ctx context.Context) BlobStorageAccountDataSetOutput
}

func (*BlobStorageAccountDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobStorageAccountDataSet)(nil)).Elem()
}

func (i *BlobStorageAccountDataSet) ToBlobStorageAccountDataSetOutput() BlobStorageAccountDataSetOutput {
	return i.ToBlobStorageAccountDataSetOutputWithContext(context.Background())
}

func (i *BlobStorageAccountDataSet) ToBlobStorageAccountDataSetOutputWithContext(ctx context.Context) BlobStorageAccountDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobStorageAccountDataSetOutput)
}

type BlobStorageAccountDataSetOutput struct{ *pulumi.OutputState }

func (BlobStorageAccountDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobStorageAccountDataSet)(nil)).Elem()
}

func (o BlobStorageAccountDataSetOutput) ToBlobStorageAccountDataSetOutput() BlobStorageAccountDataSetOutput {
	return o
}

func (o BlobStorageAccountDataSetOutput) ToBlobStorageAccountDataSetOutputWithContext(ctx context.Context) BlobStorageAccountDataSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BlobStorageAccountDataSetOutput{})
}
