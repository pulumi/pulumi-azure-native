


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlobDataSetMapping struct {
	pulumi.CustomResourceState

	ContainerName        pulumi.StringOutput      `pulumi:"containerName"`
	DataSetId            pulumi.StringOutput      `pulumi:"dataSetId"`
	DataSetMappingStatus pulumi.StringOutput      `pulumi:"dataSetMappingStatus"`
	FilePath             pulumi.StringOutput      `pulumi:"filePath"`
	Kind                 pulumi.StringOutput      `pulumi:"kind"`
	Name                 pulumi.StringOutput      `pulumi:"name"`
	OutputType           pulumi.StringPtrOutput   `pulumi:"outputType"`
	ProvisioningState    pulumi.StringOutput      `pulumi:"provisioningState"`
	ResourceGroup        pulumi.StringOutput      `pulumi:"resourceGroup"`
	StorageAccountName   pulumi.StringOutput      `pulumi:"storageAccountName"`
	SubscriptionId       pulumi.StringOutput      `pulumi:"subscriptionId"`
	SystemData           SystemDataResponseOutput `pulumi:"systemData"`
	Type                 pulumi.StringOutput      `pulumi:"type"`
}


func NewBlobDataSetMapping(ctx *pulumi.Context,
	name string, args *BlobDataSetMappingArgs, opts ...pulumi.ResourceOption) (*BlobDataSetMapping, error) {
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
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
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
			Type: pulumi.String("azure-native:datashare:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:BlobDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:BlobDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlobDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobDataSetMappingState, opts ...pulumi.ResourceOption) (*BlobDataSetMapping, error) {
	var resource BlobDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:BlobDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blobDataSetMappingState struct {
}

type BlobDataSetMappingState struct {
}

func (BlobDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobDataSetMappingState)(nil)).Elem()
}

type blobDataSetMappingArgs struct {
	AccountName           string  `pulumi:"accountName"`
	ContainerName         string  `pulumi:"containerName"`
	DataSetId             string  `pulumi:"dataSetId"`
	DataSetMappingName    *string `pulumi:"dataSetMappingName"`
	FilePath              string  `pulumi:"filePath"`
	Kind                  string  `pulumi:"kind"`
	OutputType            *string `pulumi:"outputType"`
	ResourceGroup         string  `pulumi:"resourceGroup"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	StorageAccountName    string  `pulumi:"storageAccountName"`
	SubscriptionId        string  `pulumi:"subscriptionId"`
}


type BlobDataSetMappingArgs struct {
	AccountName           pulumi.StringInput
	ContainerName         pulumi.StringInput
	DataSetId             pulumi.StringInput
	DataSetMappingName    pulumi.StringPtrInput
	FilePath              pulumi.StringInput
	Kind                  pulumi.StringInput
	OutputType            pulumi.StringPtrInput
	ResourceGroup         pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	ShareSubscriptionName pulumi.StringInput
	StorageAccountName    pulumi.StringInput
	SubscriptionId        pulumi.StringInput
}

func (BlobDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobDataSetMappingArgs)(nil)).Elem()
}

type BlobDataSetMappingInput interface {
	pulumi.Input

	ToBlobDataSetMappingOutput() BlobDataSetMappingOutput
	ToBlobDataSetMappingOutputWithContext(ctx context.Context) BlobDataSetMappingOutput
}

func (*BlobDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobDataSetMapping)(nil)).Elem()
}

func (i *BlobDataSetMapping) ToBlobDataSetMappingOutput() BlobDataSetMappingOutput {
	return i.ToBlobDataSetMappingOutputWithContext(context.Background())
}

func (i *BlobDataSetMapping) ToBlobDataSetMappingOutputWithContext(ctx context.Context) BlobDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobDataSetMappingOutput)
}

type BlobDataSetMappingOutput struct{ *pulumi.OutputState }

func (BlobDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobDataSetMapping)(nil)).Elem()
}

func (o BlobDataSetMappingOutput) ToBlobDataSetMappingOutput() BlobDataSetMappingOutput {
	return o
}

func (o BlobDataSetMappingOutput) ToBlobDataSetMappingOutputWithContext(ctx context.Context) BlobDataSetMappingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BlobDataSetMappingOutput{})
}
