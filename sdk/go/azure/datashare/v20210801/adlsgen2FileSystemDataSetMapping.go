


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ADLSGen2FileSystemDataSetMapping struct {
	pulumi.CustomResourceState

	DataSetId            pulumi.StringOutput      `pulumi:"dataSetId"`
	DataSetMappingStatus pulumi.StringOutput      `pulumi:"dataSetMappingStatus"`
	FileSystem           pulumi.StringOutput      `pulumi:"fileSystem"`
	Kind                 pulumi.StringOutput      `pulumi:"kind"`
	Name                 pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput      `pulumi:"provisioningState"`
	ResourceGroup        pulumi.StringOutput      `pulumi:"resourceGroup"`
	StorageAccountName   pulumi.StringOutput      `pulumi:"storageAccountName"`
	SubscriptionId       pulumi.StringOutput      `pulumi:"subscriptionId"`
	SystemData           SystemDataResponseOutput `pulumi:"systemData"`
	Type                 pulumi.StringOutput      `pulumi:"type"`
}


func NewADLSGen2FileSystemDataSetMapping(ctx *pulumi.Context,
	name string, args *ADLSGen2FileSystemDataSetMappingArgs, opts ...pulumi.ResourceOption) (*ADLSGen2FileSystemDataSetMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DataSetId == nil {
		return nil, errors.New("invalid value for required argument 'DataSetId'")
	}
	if args.FileSystem == nil {
		return nil, errors.New("invalid value for required argument 'FileSystem'")
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
	args.Kind = pulumi.String("AdlsGen2FileSystem")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datashare/v20210801:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20181101preview:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20191101:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20200901:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20201001preview:ADLSGen2FileSystemDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource ADLSGen2FileSystemDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20210801:ADLSGen2FileSystemDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetADLSGen2FileSystemDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen2FileSystemDataSetMappingState, opts ...pulumi.ResourceOption) (*ADLSGen2FileSystemDataSetMapping, error) {
	var resource ADLSGen2FileSystemDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20210801:ADLSGen2FileSystemDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adlsgen2FileSystemDataSetMappingState struct {
}

type ADLSGen2FileSystemDataSetMappingState struct {
}

func (ADLSGen2FileSystemDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FileSystemDataSetMappingState)(nil)).Elem()
}

type adlsgen2FileSystemDataSetMappingArgs struct {
	AccountName           string  `pulumi:"accountName"`
	DataSetId             string  `pulumi:"dataSetId"`
	DataSetMappingName    *string `pulumi:"dataSetMappingName"`
	FileSystem            string  `pulumi:"fileSystem"`
	Kind                  string  `pulumi:"kind"`
	ResourceGroup         string  `pulumi:"resourceGroup"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	StorageAccountName    string  `pulumi:"storageAccountName"`
	SubscriptionId        string  `pulumi:"subscriptionId"`
}


type ADLSGen2FileSystemDataSetMappingArgs struct {
	AccountName           pulumi.StringInput
	DataSetId             pulumi.StringInput
	DataSetMappingName    pulumi.StringPtrInput
	FileSystem            pulumi.StringInput
	Kind                  pulumi.StringInput
	ResourceGroup         pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	ShareSubscriptionName pulumi.StringInput
	StorageAccountName    pulumi.StringInput
	SubscriptionId        pulumi.StringInput
}

func (ADLSGen2FileSystemDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FileSystemDataSetMappingArgs)(nil)).Elem()
}

type ADLSGen2FileSystemDataSetMappingInput interface {
	pulumi.Input

	ToADLSGen2FileSystemDataSetMappingOutput() ADLSGen2FileSystemDataSetMappingOutput
	ToADLSGen2FileSystemDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FileSystemDataSetMappingOutput
}

func (*ADLSGen2FileSystemDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*ADLSGen2FileSystemDataSetMapping)(nil))
}

func (i *ADLSGen2FileSystemDataSetMapping) ToADLSGen2FileSystemDataSetMappingOutput() ADLSGen2FileSystemDataSetMappingOutput {
	return i.ToADLSGen2FileSystemDataSetMappingOutputWithContext(context.Background())
}

func (i *ADLSGen2FileSystemDataSetMapping) ToADLSGen2FileSystemDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FileSystemDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2FileSystemDataSetMappingOutput)
}

type ADLSGen2FileSystemDataSetMappingOutput struct{ *pulumi.OutputState }

func (ADLSGen2FileSystemDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ADLSGen2FileSystemDataSetMapping)(nil))
}

func (o ADLSGen2FileSystemDataSetMappingOutput) ToADLSGen2FileSystemDataSetMappingOutput() ADLSGen2FileSystemDataSetMappingOutput {
	return o
}

func (o ADLSGen2FileSystemDataSetMappingOutput) ToADLSGen2FileSystemDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FileSystemDataSetMappingOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ADLSGen2FileSystemDataSetMappingInput)(nil)).Elem(), &ADLSGen2FileSystemDataSetMapping{})
	pulumi.RegisterOutputType(ADLSGen2FileSystemDataSetMappingOutput{})
}
