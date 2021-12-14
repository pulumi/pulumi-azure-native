


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlobContainerDataSet struct {
	pulumi.CustomResourceState

	ContainerName      pulumi.StringOutput      `pulumi:"containerName"`
	DataSetId          pulumi.StringOutput      `pulumi:"dataSetId"`
	Kind               pulumi.StringOutput      `pulumi:"kind"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	ResourceGroup      pulumi.StringOutput      `pulumi:"resourceGroup"`
	StorageAccountName pulumi.StringOutput      `pulumi:"storageAccountName"`
	SubscriptionId     pulumi.StringOutput      `pulumi:"subscriptionId"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewBlobContainerDataSet(ctx *pulumi.Context,
	name string, args *BlobContainerDataSetArgs, opts ...pulumi.ResourceOption) (*BlobContainerDataSet, error) {
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
	args.Kind = pulumi.String("Container")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:BlobContainerDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:BlobContainerDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:BlobContainerDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:BlobContainerDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:BlobContainerDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobContainerDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:BlobContainerDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlobContainerDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobContainerDataSetState, opts ...pulumi.ResourceOption) (*BlobContainerDataSet, error) {
	var resource BlobContainerDataSet
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:BlobContainerDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blobContainerDataSetState struct {
}

type BlobContainerDataSetState struct {
}

func (BlobContainerDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerDataSetState)(nil)).Elem()
}

type blobContainerDataSetArgs struct {
	AccountName        string  `pulumi:"accountName"`
	ContainerName      string  `pulumi:"containerName"`
	DataSetName        *string `pulumi:"dataSetName"`
	Kind               string  `pulumi:"kind"`
	ResourceGroup      string  `pulumi:"resourceGroup"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	ShareName          string  `pulumi:"shareName"`
	StorageAccountName string  `pulumi:"storageAccountName"`
	SubscriptionId     string  `pulumi:"subscriptionId"`
}


type BlobContainerDataSetArgs struct {
	AccountName        pulumi.StringInput
	ContainerName      pulumi.StringInput
	DataSetName        pulumi.StringPtrInput
	Kind               pulumi.StringInput
	ResourceGroup      pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	ShareName          pulumi.StringInput
	StorageAccountName pulumi.StringInput
	SubscriptionId     pulumi.StringInput
}

func (BlobContainerDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerDataSetArgs)(nil)).Elem()
}

type BlobContainerDataSetInput interface {
	pulumi.Input

	ToBlobContainerDataSetOutput() BlobContainerDataSetOutput
	ToBlobContainerDataSetOutputWithContext(ctx context.Context) BlobContainerDataSetOutput
}

func (*BlobContainerDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobContainerDataSet)(nil)).Elem()
}

func (i *BlobContainerDataSet) ToBlobContainerDataSetOutput() BlobContainerDataSetOutput {
	return i.ToBlobContainerDataSetOutputWithContext(context.Background())
}

func (i *BlobContainerDataSet) ToBlobContainerDataSetOutputWithContext(ctx context.Context) BlobContainerDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobContainerDataSetOutput)
}

type BlobContainerDataSetOutput struct{ *pulumi.OutputState }

func (BlobContainerDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobContainerDataSet)(nil)).Elem()
}

func (o BlobContainerDataSetOutput) ToBlobContainerDataSetOutput() BlobContainerDataSetOutput {
	return o
}

func (o BlobContainerDataSetOutput) ToBlobContainerDataSetOutputWithContext(ctx context.Context) BlobContainerDataSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BlobContainerDataSetOutput{})
}
