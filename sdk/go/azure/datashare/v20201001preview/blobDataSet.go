


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlobDataSet struct {
	pulumi.CustomResourceState

	ContainerName      pulumi.StringOutput      `pulumi:"containerName"`
	DataSetId          pulumi.StringOutput      `pulumi:"dataSetId"`
	FilePath           pulumi.StringOutput      `pulumi:"filePath"`
	Kind               pulumi.StringOutput      `pulumi:"kind"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	ResourceGroup      pulumi.StringOutput      `pulumi:"resourceGroup"`
	StorageAccountName pulumi.StringOutput      `pulumi:"storageAccountName"`
	SubscriptionId     pulumi.StringOutput      `pulumi:"subscriptionId"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewBlobDataSet(ctx *pulumi.Context,
	name string, args *BlobDataSetArgs, opts ...pulumi.ResourceOption) (*BlobDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
	}
	if args.FilePath == nil {
		return nil, errors.New("invalid value for required argument 'FilePath'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
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
	args.Kind = pulumi.String("Blob")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datashare/v20201001preview:BlobDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare:BlobDataSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare:BlobDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:BlobDataSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20181101preview:BlobDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:BlobDataSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20191101:BlobDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:BlobDataSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20200901:BlobDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:BlobDataSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20210801:BlobDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:BlobDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlobDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobDataSetState, opts ...pulumi.ResourceOption) (*BlobDataSet, error) {
	var resource BlobDataSet
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:BlobDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blobDataSetState struct {
}

type BlobDataSetState struct {
}

func (BlobDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobDataSetState)(nil)).Elem()
}

type blobDataSetArgs struct {
	AccountName        string  `pulumi:"accountName"`
	ContainerName      string  `pulumi:"containerName"`
	DataSetName        *string `pulumi:"dataSetName"`
	FilePath           string  `pulumi:"filePath"`
	Kind               string  `pulumi:"kind"`
	ResourceGroup      string  `pulumi:"resourceGroup"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	ShareName          string  `pulumi:"shareName"`
	StorageAccountName string  `pulumi:"storageAccountName"`
	SubscriptionId     string  `pulumi:"subscriptionId"`
}


type BlobDataSetArgs struct {
	AccountName        pulumi.StringInput
	ContainerName      pulumi.StringInput
	DataSetName        pulumi.StringPtrInput
	FilePath           pulumi.StringInput
	Kind               pulumi.StringInput
	ResourceGroup      pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	ShareName          pulumi.StringInput
	StorageAccountName pulumi.StringInput
	SubscriptionId     pulumi.StringInput
}

func (BlobDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobDataSetArgs)(nil)).Elem()
}

type BlobDataSetInput interface {
	pulumi.Input

	ToBlobDataSetOutput() BlobDataSetOutput
	ToBlobDataSetOutputWithContext(ctx context.Context) BlobDataSetOutput
}

func (*BlobDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobDataSet)(nil))
}

func (i *BlobDataSet) ToBlobDataSetOutput() BlobDataSetOutput {
	return i.ToBlobDataSetOutputWithContext(context.Background())
}

func (i *BlobDataSet) ToBlobDataSetOutputWithContext(ctx context.Context) BlobDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobDataSetOutput)
}

type BlobDataSetOutput struct{ *pulumi.OutputState }

func (BlobDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobDataSet)(nil))
}

func (o BlobDataSetOutput) ToBlobDataSetOutput() BlobDataSetOutput {
	return o
}

func (o BlobDataSetOutput) ToBlobDataSetOutputWithContext(ctx context.Context) BlobDataSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BlobDataSetOutput{})
}
