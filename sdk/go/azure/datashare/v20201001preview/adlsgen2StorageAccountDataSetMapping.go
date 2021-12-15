


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ADLSGen2StorageAccountDataSetMapping struct {
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


func NewADLSGen2StorageAccountDataSetMapping(ctx *pulumi.Context,
	name string, args *ADLSGen2StorageAccountDataSetMappingArgs, opts ...pulumi.ResourceOption) (*ADLSGen2StorageAccountDataSetMapping, error) {
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
	args.Kind = pulumi.String("AdlsGen2StorageAccount")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:ADLSGen2StorageAccountDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ADLSGen2StorageAccountDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen2StorageAccountDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen2StorageAccountDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ADLSGen2StorageAccountDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource ADLSGen2StorageAccountDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:ADLSGen2StorageAccountDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetADLSGen2StorageAccountDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen2StorageAccountDataSetMappingState, opts ...pulumi.ResourceOption) (*ADLSGen2StorageAccountDataSetMapping, error) {
	var resource ADLSGen2StorageAccountDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:ADLSGen2StorageAccountDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adlsgen2StorageAccountDataSetMappingState struct {
}

type ADLSGen2StorageAccountDataSetMappingState struct {
}

func (ADLSGen2StorageAccountDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2StorageAccountDataSetMappingState)(nil)).Elem()
}

type adlsgen2StorageAccountDataSetMappingArgs struct {
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


type ADLSGen2StorageAccountDataSetMappingArgs struct {
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

func (ADLSGen2StorageAccountDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2StorageAccountDataSetMappingArgs)(nil)).Elem()
}

type ADLSGen2StorageAccountDataSetMappingInput interface {
	pulumi.Input

	ToADLSGen2StorageAccountDataSetMappingOutput() ADLSGen2StorageAccountDataSetMappingOutput
	ToADLSGen2StorageAccountDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2StorageAccountDataSetMappingOutput
}

func (*ADLSGen2StorageAccountDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2StorageAccountDataSetMapping)(nil)).Elem()
}

func (i *ADLSGen2StorageAccountDataSetMapping) ToADLSGen2StorageAccountDataSetMappingOutput() ADLSGen2StorageAccountDataSetMappingOutput {
	return i.ToADLSGen2StorageAccountDataSetMappingOutputWithContext(context.Background())
}

func (i *ADLSGen2StorageAccountDataSetMapping) ToADLSGen2StorageAccountDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2StorageAccountDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2StorageAccountDataSetMappingOutput)
}

type ADLSGen2StorageAccountDataSetMappingOutput struct{ *pulumi.OutputState }

func (ADLSGen2StorageAccountDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2StorageAccountDataSetMapping)(nil)).Elem()
}

func (o ADLSGen2StorageAccountDataSetMappingOutput) ToADLSGen2StorageAccountDataSetMappingOutput() ADLSGen2StorageAccountDataSetMappingOutput {
	return o
}

func (o ADLSGen2StorageAccountDataSetMappingOutput) ToADLSGen2StorageAccountDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2StorageAccountDataSetMappingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ADLSGen2StorageAccountDataSetMappingOutput{})
}
