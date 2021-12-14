


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlobStorageAccountDataSetMapping struct {
	pulumi.CustomResourceState

	ContainerName            pulumi.StringOutput      `pulumi:"containerName"`
	DataSetId                pulumi.StringOutput      `pulumi:"dataSetId"`
	DataSetMappingStatus     pulumi.StringOutput      `pulumi:"dataSetMappingStatus"`
	Folder                   pulumi.StringOutput      `pulumi:"folder"`
	Kind                     pulumi.StringOutput      `pulumi:"kind"`
	Location                 pulumi.StringOutput      `pulumi:"location"`
	MountPath                pulumi.StringPtrOutput   `pulumi:"mountPath"`
	Name                     pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState        pulumi.StringOutput      `pulumi:"provisioningState"`
	StorageAccountResourceId pulumi.StringOutput      `pulumi:"storageAccountResourceId"`
	SystemData               SystemDataResponseOutput `pulumi:"systemData"`
	Type                     pulumi.StringOutput      `pulumi:"type"`
}


func NewBlobStorageAccountDataSetMapping(ctx *pulumi.Context,
	name string, args *BlobStorageAccountDataSetMappingArgs, opts ...pulumi.ResourceOption) (*BlobStorageAccountDataSetMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
	}
	if args.DataSetId == nil {
		return nil, errors.New("invalid value for required argument 'DataSetId'")
	}
	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
	}
	if args.StorageAccountResourceId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountResourceId'")
	}
	args.Kind = pulumi.String("BlobStorageAccount")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:BlobStorageAccountDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:BlobStorageAccountDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:BlobStorageAccountDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:BlobStorageAccountDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:BlobStorageAccountDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobStorageAccountDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:BlobStorageAccountDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlobStorageAccountDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobStorageAccountDataSetMappingState, opts ...pulumi.ResourceOption) (*BlobStorageAccountDataSetMapping, error) {
	var resource BlobStorageAccountDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:BlobStorageAccountDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blobStorageAccountDataSetMappingState struct {
}

type BlobStorageAccountDataSetMappingState struct {
}

func (BlobStorageAccountDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobStorageAccountDataSetMappingState)(nil)).Elem()
}

type blobStorageAccountDataSetMappingArgs struct {
	AccountName              string  `pulumi:"accountName"`
	ContainerName            string  `pulumi:"containerName"`
	DataSetId                string  `pulumi:"dataSetId"`
	DataSetMappingName       *string `pulumi:"dataSetMappingName"`
	Folder                   string  `pulumi:"folder"`
	Kind                     string  `pulumi:"kind"`
	MountPath                *string `pulumi:"mountPath"`
	ResourceGroupName        string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName    string  `pulumi:"shareSubscriptionName"`
	StorageAccountResourceId string  `pulumi:"storageAccountResourceId"`
}


type BlobStorageAccountDataSetMappingArgs struct {
	AccountName              pulumi.StringInput
	ContainerName            pulumi.StringInput
	DataSetId                pulumi.StringInput
	DataSetMappingName       pulumi.StringPtrInput
	Folder                   pulumi.StringInput
	Kind                     pulumi.StringInput
	MountPath                pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	ShareSubscriptionName    pulumi.StringInput
	StorageAccountResourceId pulumi.StringInput
}

func (BlobStorageAccountDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobStorageAccountDataSetMappingArgs)(nil)).Elem()
}

type BlobStorageAccountDataSetMappingInput interface {
	pulumi.Input

	ToBlobStorageAccountDataSetMappingOutput() BlobStorageAccountDataSetMappingOutput
	ToBlobStorageAccountDataSetMappingOutputWithContext(ctx context.Context) BlobStorageAccountDataSetMappingOutput
}

func (*BlobStorageAccountDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobStorageAccountDataSetMapping)(nil)).Elem()
}

func (i *BlobStorageAccountDataSetMapping) ToBlobStorageAccountDataSetMappingOutput() BlobStorageAccountDataSetMappingOutput {
	return i.ToBlobStorageAccountDataSetMappingOutputWithContext(context.Background())
}

func (i *BlobStorageAccountDataSetMapping) ToBlobStorageAccountDataSetMappingOutputWithContext(ctx context.Context) BlobStorageAccountDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobStorageAccountDataSetMappingOutput)
}

type BlobStorageAccountDataSetMappingOutput struct{ *pulumi.OutputState }

func (BlobStorageAccountDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobStorageAccountDataSetMapping)(nil)).Elem()
}

func (o BlobStorageAccountDataSetMappingOutput) ToBlobStorageAccountDataSetMappingOutput() BlobStorageAccountDataSetMappingOutput {
	return o
}

func (o BlobStorageAccountDataSetMappingOutput) ToBlobStorageAccountDataSetMappingOutputWithContext(ctx context.Context) BlobStorageAccountDataSetMappingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BlobStorageAccountDataSetMappingOutput{})
}
