


package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ADLSGen1FolderDataSet struct {
	pulumi.CustomResourceState

	AccountName    pulumi.StringOutput `pulumi:"accountName"`
	DataSetId      pulumi.StringOutput `pulumi:"dataSetId"`
	FolderPath     pulumi.StringOutput `pulumi:"folderPath"`
	Kind           pulumi.StringOutput `pulumi:"kind"`
	Name           pulumi.StringOutput `pulumi:"name"`
	ResourceGroup  pulumi.StringOutput `pulumi:"resourceGroup"`
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	Type           pulumi.StringOutput `pulumi:"type"`
}


func NewADLSGen1FolderDataSet(ctx *pulumi.Context,
	name string, args *ADLSGen1FolderDataSetArgs, opts ...pulumi.ResourceOption) (*ADLSGen1FolderDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.FolderPath == nil {
		return nil, errors.New("invalid value for required argument 'FolderPath'")
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
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	args.Kind = pulumi.String("AdlsGen1Folder")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:ADLSGen1FolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen1FolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen1FolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ADLSGen1FolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ADLSGen1FolderDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource ADLSGen1FolderDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20181101preview:ADLSGen1FolderDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetADLSGen1FolderDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen1FolderDataSetState, opts ...pulumi.ResourceOption) (*ADLSGen1FolderDataSet, error) {
	var resource ADLSGen1FolderDataSet
	err := ctx.ReadResource("azure-native:datashare/v20181101preview:ADLSGen1FolderDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adlsgen1FolderDataSetState struct {
}

type ADLSGen1FolderDataSetState struct {
}

func (ADLSGen1FolderDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen1FolderDataSetState)(nil)).Elem()
}

type adlsgen1FolderDataSetArgs struct {
	AccountName       string  `pulumi:"accountName"`
	DataSetName       *string `pulumi:"dataSetName"`
	FolderPath        string  `pulumi:"folderPath"`
	Kind              string  `pulumi:"kind"`
	ResourceGroup     string  `pulumi:"resourceGroup"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ShareName         string  `pulumi:"shareName"`
	SubscriptionId    string  `pulumi:"subscriptionId"`
}


type ADLSGen1FolderDataSetArgs struct {
	AccountName       pulumi.StringInput
	DataSetName       pulumi.StringPtrInput
	FolderPath        pulumi.StringInput
	Kind              pulumi.StringInput
	ResourceGroup     pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ShareName         pulumi.StringInput
	SubscriptionId    pulumi.StringInput
}

func (ADLSGen1FolderDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen1FolderDataSetArgs)(nil)).Elem()
}

type ADLSGen1FolderDataSetInput interface {
	pulumi.Input

	ToADLSGen1FolderDataSetOutput() ADLSGen1FolderDataSetOutput
	ToADLSGen1FolderDataSetOutputWithContext(ctx context.Context) ADLSGen1FolderDataSetOutput
}

func (*ADLSGen1FolderDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen1FolderDataSet)(nil)).Elem()
}

func (i *ADLSGen1FolderDataSet) ToADLSGen1FolderDataSetOutput() ADLSGen1FolderDataSetOutput {
	return i.ToADLSGen1FolderDataSetOutputWithContext(context.Background())
}

func (i *ADLSGen1FolderDataSet) ToADLSGen1FolderDataSetOutputWithContext(ctx context.Context) ADLSGen1FolderDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen1FolderDataSetOutput)
}

type ADLSGen1FolderDataSetOutput struct{ *pulumi.OutputState }

func (ADLSGen1FolderDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen1FolderDataSet)(nil)).Elem()
}

func (o ADLSGen1FolderDataSetOutput) ToADLSGen1FolderDataSetOutput() ADLSGen1FolderDataSetOutput {
	return o
}

func (o ADLSGen1FolderDataSetOutput) ToADLSGen1FolderDataSetOutputWithContext(ctx context.Context) ADLSGen1FolderDataSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ADLSGen1FolderDataSetOutput{})
}
