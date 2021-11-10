


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlobFolderDataSetMapping struct {
	pulumi.CustomResourceState

	ContainerName        pulumi.StringOutput      `pulumi:"containerName"`
	DataSetId            pulumi.StringOutput      `pulumi:"dataSetId"`
	DataSetMappingStatus pulumi.StringOutput      `pulumi:"dataSetMappingStatus"`
	Kind                 pulumi.StringOutput      `pulumi:"kind"`
	Name                 pulumi.StringOutput      `pulumi:"name"`
	Prefix               pulumi.StringOutput      `pulumi:"prefix"`
	ProvisioningState    pulumi.StringOutput      `pulumi:"provisioningState"`
	ResourceGroup        pulumi.StringOutput      `pulumi:"resourceGroup"`
	StorageAccountName   pulumi.StringOutput      `pulumi:"storageAccountName"`
	SubscriptionId       pulumi.StringOutput      `pulumi:"subscriptionId"`
	SystemData           SystemDataResponseOutput `pulumi:"systemData"`
	Type                 pulumi.StringOutput      `pulumi:"type"`
}


func NewBlobFolderDataSetMapping(ctx *pulumi.Context,
	name string, args *BlobFolderDataSetMappingArgs, opts ...pulumi.ResourceOption) (*BlobFolderDataSetMapping, error) {
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
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
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
			Type: pulumi.String("azure-native:datashare:BlobFolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:BlobFolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:BlobFolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:BlobFolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:BlobFolderDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobFolderDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:BlobFolderDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlobFolderDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobFolderDataSetMappingState, opts ...pulumi.ResourceOption) (*BlobFolderDataSetMapping, error) {
	var resource BlobFolderDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:BlobFolderDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blobFolderDataSetMappingState struct {
}

type BlobFolderDataSetMappingState struct {
}

func (BlobFolderDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobFolderDataSetMappingState)(nil)).Elem()
}

type blobFolderDataSetMappingArgs struct {
	AccountName           string  `pulumi:"accountName"`
	ContainerName         string  `pulumi:"containerName"`
	DataSetId             string  `pulumi:"dataSetId"`
	DataSetMappingName    *string `pulumi:"dataSetMappingName"`
	Kind                  string  `pulumi:"kind"`
	Prefix                string  `pulumi:"prefix"`
	ResourceGroup         string  `pulumi:"resourceGroup"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	StorageAccountName    string  `pulumi:"storageAccountName"`
	SubscriptionId        string  `pulumi:"subscriptionId"`
}


type BlobFolderDataSetMappingArgs struct {
	AccountName           pulumi.StringInput
	ContainerName         pulumi.StringInput
	DataSetId             pulumi.StringInput
	DataSetMappingName    pulumi.StringPtrInput
	Kind                  pulumi.StringInput
	Prefix                pulumi.StringInput
	ResourceGroup         pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	ShareSubscriptionName pulumi.StringInput
	StorageAccountName    pulumi.StringInput
	SubscriptionId        pulumi.StringInput
}

func (BlobFolderDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobFolderDataSetMappingArgs)(nil)).Elem()
}

type BlobFolderDataSetMappingInput interface {
	pulumi.Input

	ToBlobFolderDataSetMappingOutput() BlobFolderDataSetMappingOutput
	ToBlobFolderDataSetMappingOutputWithContext(ctx context.Context) BlobFolderDataSetMappingOutput
}

func (*BlobFolderDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobFolderDataSetMapping)(nil))
}

func (i *BlobFolderDataSetMapping) ToBlobFolderDataSetMappingOutput() BlobFolderDataSetMappingOutput {
	return i.ToBlobFolderDataSetMappingOutputWithContext(context.Background())
}

func (i *BlobFolderDataSetMapping) ToBlobFolderDataSetMappingOutputWithContext(ctx context.Context) BlobFolderDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobFolderDataSetMappingOutput)
}

type BlobFolderDataSetMappingOutput struct{ *pulumi.OutputState }

func (BlobFolderDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobFolderDataSetMapping)(nil))
}

func (o BlobFolderDataSetMappingOutput) ToBlobFolderDataSetMappingOutput() BlobFolderDataSetMappingOutput {
	return o
}

func (o BlobFolderDataSetMappingOutput) ToBlobFolderDataSetMappingOutputWithContext(ctx context.Context) BlobFolderDataSetMappingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BlobFolderDataSetMappingOutput{})
}
