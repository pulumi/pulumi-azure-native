


package v20191101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ADLSGen2FileDataSet struct {
	pulumi.CustomResourceState

	DataSetId          pulumi.StringOutput `pulumi:"dataSetId"`
	FilePath           pulumi.StringOutput `pulumi:"filePath"`
	FileSystem         pulumi.StringOutput `pulumi:"fileSystem"`
	Kind               pulumi.StringOutput `pulumi:"kind"`
	Name               pulumi.StringOutput `pulumi:"name"`
	ResourceGroup      pulumi.StringOutput `pulumi:"resourceGroup"`
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	SubscriptionId     pulumi.StringOutput `pulumi:"subscriptionId"`
	Type               pulumi.StringOutput `pulumi:"type"`
}


func NewADLSGen2FileDataSet(ctx *pulumi.Context,
	name string, args *ADLSGen2FileDataSetArgs, opts ...pulumi.ResourceOption) (*ADLSGen2FileDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.FilePath == nil {
		return nil, errors.New("invalid value for required argument 'FilePath'")
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
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	args.Kind = pulumi.String("AdlsGen2File")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:ADLSGen2FileDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ADLSGen2FileDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen2FileDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ADLSGen2FileDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ADLSGen2FileDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource ADLSGen2FileDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20191101:ADLSGen2FileDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetADLSGen2FileDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen2FileDataSetState, opts ...pulumi.ResourceOption) (*ADLSGen2FileDataSet, error) {
	var resource ADLSGen2FileDataSet
	err := ctx.ReadResource("azure-native:datashare/v20191101:ADLSGen2FileDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adlsgen2FileDataSetState struct {
}

type ADLSGen2FileDataSetState struct {
}

func (ADLSGen2FileDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FileDataSetState)(nil)).Elem()
}

type adlsgen2FileDataSetArgs struct {
	AccountName        string  `pulumi:"accountName"`
	DataSetName        *string `pulumi:"dataSetName"`
	FilePath           string  `pulumi:"filePath"`
	FileSystem         string  `pulumi:"fileSystem"`
	Kind               string  `pulumi:"kind"`
	ResourceGroup      string  `pulumi:"resourceGroup"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	ShareName          string  `pulumi:"shareName"`
	StorageAccountName string  `pulumi:"storageAccountName"`
	SubscriptionId     string  `pulumi:"subscriptionId"`
}


type ADLSGen2FileDataSetArgs struct {
	AccountName        pulumi.StringInput
	DataSetName        pulumi.StringPtrInput
	FilePath           pulumi.StringInput
	FileSystem         pulumi.StringInput
	Kind               pulumi.StringInput
	ResourceGroup      pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	ShareName          pulumi.StringInput
	StorageAccountName pulumi.StringInput
	SubscriptionId     pulumi.StringInput
}

func (ADLSGen2FileDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FileDataSetArgs)(nil)).Elem()
}

type ADLSGen2FileDataSetInput interface {
	pulumi.Input

	ToADLSGen2FileDataSetOutput() ADLSGen2FileDataSetOutput
	ToADLSGen2FileDataSetOutputWithContext(ctx context.Context) ADLSGen2FileDataSetOutput
}

func (*ADLSGen2FileDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FileDataSet)(nil)).Elem()
}

func (i *ADLSGen2FileDataSet) ToADLSGen2FileDataSetOutput() ADLSGen2FileDataSetOutput {
	return i.ToADLSGen2FileDataSetOutputWithContext(context.Background())
}

func (i *ADLSGen2FileDataSet) ToADLSGen2FileDataSetOutputWithContext(ctx context.Context) ADLSGen2FileDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2FileDataSetOutput)
}

type ADLSGen2FileDataSetOutput struct{ *pulumi.OutputState }

func (ADLSGen2FileDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FileDataSet)(nil)).Elem()
}

func (o ADLSGen2FileDataSetOutput) ToADLSGen2FileDataSetOutput() ADLSGen2FileDataSetOutput {
	return o
}

func (o ADLSGen2FileDataSetOutput) ToADLSGen2FileDataSetOutputWithContext(ctx context.Context) ADLSGen2FileDataSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ADLSGen2FileDataSetOutput{})
}
