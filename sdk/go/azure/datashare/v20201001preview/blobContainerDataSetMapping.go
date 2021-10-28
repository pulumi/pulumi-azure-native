


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlobContainerDataSetMapping struct {
	pulumi.CustomResourceState

	ContainerName        pulumi.StringOutput      `pulumi:"containerName"`
	DataSetId            pulumi.StringOutput      `pulumi:"dataSetId"`
	DataSetMappingStatus pulumi.StringOutput      `pulumi:"dataSetMappingStatus"`
	Kind                 pulumi.StringOutput      `pulumi:"kind"`
	Name                 pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput      `pulumi:"provisioningState"`
	ResourceGroup        pulumi.StringOutput      `pulumi:"resourceGroup"`
	StorageAccountName   pulumi.StringOutput      `pulumi:"storageAccountName"`
	SubscriptionId       pulumi.StringOutput      `pulumi:"subscriptionId"`
	SystemData           SystemDataResponseOutput `pulumi:"systemData"`
	Type                 pulumi.StringOutput      `pulumi:"type"`
}


func NewBlobContainerDataSetMapping(ctx *pulumi.Context,
	name string, args *BlobContainerDataSetMappingArgs, opts ...pulumi.ResourceOption) (*BlobContainerDataSetMapping, error) {
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
	args.Kind = pulumi.String("Container")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datashare/v20201001preview:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20181101preview:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20191101:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20200901:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20210801:BlobContainerDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobContainerDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:BlobContainerDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlobContainerDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobContainerDataSetMappingState, opts ...pulumi.ResourceOption) (*BlobContainerDataSetMapping, error) {
	var resource BlobContainerDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:BlobContainerDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blobContainerDataSetMappingState struct {
}

type BlobContainerDataSetMappingState struct {
}

func (BlobContainerDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerDataSetMappingState)(nil)).Elem()
}

type blobContainerDataSetMappingArgs struct {
	AccountName           string  `pulumi:"accountName"`
	ContainerName         string  `pulumi:"containerName"`
	DataSetId             string  `pulumi:"dataSetId"`
	DataSetMappingName    *string `pulumi:"dataSetMappingName"`
	Kind                  string  `pulumi:"kind"`
	ResourceGroup         string  `pulumi:"resourceGroup"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	StorageAccountName    string  `pulumi:"storageAccountName"`
	SubscriptionId        string  `pulumi:"subscriptionId"`
}


type BlobContainerDataSetMappingArgs struct {
	AccountName           pulumi.StringInput
	ContainerName         pulumi.StringInput
	DataSetId             pulumi.StringInput
	DataSetMappingName    pulumi.StringPtrInput
	Kind                  pulumi.StringInput
	ResourceGroup         pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	ShareSubscriptionName pulumi.StringInput
	StorageAccountName    pulumi.StringInput
	SubscriptionId        pulumi.StringInput
}

func (BlobContainerDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerDataSetMappingArgs)(nil)).Elem()
}

type BlobContainerDataSetMappingInput interface {
	pulumi.Input

	ToBlobContainerDataSetMappingOutput() BlobContainerDataSetMappingOutput
	ToBlobContainerDataSetMappingOutputWithContext(ctx context.Context) BlobContainerDataSetMappingOutput
}

func (*BlobContainerDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobContainerDataSetMapping)(nil))
}

func (i *BlobContainerDataSetMapping) ToBlobContainerDataSetMappingOutput() BlobContainerDataSetMappingOutput {
	return i.ToBlobContainerDataSetMappingOutputWithContext(context.Background())
}

func (i *BlobContainerDataSetMapping) ToBlobContainerDataSetMappingOutputWithContext(ctx context.Context) BlobContainerDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobContainerDataSetMappingOutput)
}

type BlobContainerDataSetMappingOutput struct{ *pulumi.OutputState }

func (BlobContainerDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobContainerDataSetMapping)(nil))
}

func (o BlobContainerDataSetMappingOutput) ToBlobContainerDataSetMappingOutput() BlobContainerDataSetMappingOutput {
	return o
}

func (o BlobContainerDataSetMappingOutput) ToBlobContainerDataSetMappingOutputWithContext(ctx context.Context) BlobContainerDataSetMappingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BlobContainerDataSetMappingOutput{})
}
