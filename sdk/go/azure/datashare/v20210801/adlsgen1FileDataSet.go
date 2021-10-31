


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ADLSGen1FileDataSet struct {
	pulumi.CustomResourceState

	AccountName    pulumi.StringOutput      `pulumi:"accountName"`
	DataSetId      pulumi.StringOutput      `pulumi:"dataSetId"`
	FileName       pulumi.StringOutput      `pulumi:"fileName"`
	FolderPath     pulumi.StringOutput      `pulumi:"folderPath"`
	Kind           pulumi.StringOutput      `pulumi:"kind"`
	Name           pulumi.StringOutput      `pulumi:"name"`
	ResourceGroup  pulumi.StringOutput      `pulumi:"resourceGroup"`
	SubscriptionId pulumi.StringOutput      `pulumi:"subscriptionId"`
	SystemData     SystemDataResponseOutput `pulumi:"systemData"`
	Type           pulumi.StringOutput      `pulumi:"type"`
}


func NewADLSGen1FileDataSet(ctx *pulumi.Context,
	name string, args *ADLSGen1FileDataSetArgs, opts ...pulumi.ResourceOption) (*ADLSGen1FileDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.FileName == nil {
		return nil, errors.New("invalid value for required argument 'FileName'")
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
	args.Kind = pulumi.String("AdlsGen1File")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datashare/v20210801:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20181101preview:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20191101:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20200901:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20201001preview:ADLSGen1FileDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource ADLSGen1FileDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20210801:ADLSGen1FileDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetADLSGen1FileDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen1FileDataSetState, opts ...pulumi.ResourceOption) (*ADLSGen1FileDataSet, error) {
	var resource ADLSGen1FileDataSet
	err := ctx.ReadResource("azure-native:datashare/v20210801:ADLSGen1FileDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adlsgen1FileDataSetState struct {
}

type ADLSGen1FileDataSetState struct {
}

func (ADLSGen1FileDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen1FileDataSetState)(nil)).Elem()
}

type adlsgen1FileDataSetArgs struct {
	AccountName       string  `pulumi:"accountName"`
	DataSetName       *string `pulumi:"dataSetName"`
	FileName          string  `pulumi:"fileName"`
	FolderPath        string  `pulumi:"folderPath"`
	Kind              string  `pulumi:"kind"`
	ResourceGroup     string  `pulumi:"resourceGroup"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ShareName         string  `pulumi:"shareName"`
	SubscriptionId    string  `pulumi:"subscriptionId"`
}


type ADLSGen1FileDataSetArgs struct {
	AccountName       pulumi.StringInput
	DataSetName       pulumi.StringPtrInput
	FileName          pulumi.StringInput
	FolderPath        pulumi.StringInput
	Kind              pulumi.StringInput
	ResourceGroup     pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ShareName         pulumi.StringInput
	SubscriptionId    pulumi.StringInput
}

func (ADLSGen1FileDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen1FileDataSetArgs)(nil)).Elem()
}

type ADLSGen1FileDataSetInput interface {
	pulumi.Input

	ToADLSGen1FileDataSetOutput() ADLSGen1FileDataSetOutput
	ToADLSGen1FileDataSetOutputWithContext(ctx context.Context) ADLSGen1FileDataSetOutput
}

func (*ADLSGen1FileDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ADLSGen1FileDataSet)(nil))
}

func (i *ADLSGen1FileDataSet) ToADLSGen1FileDataSetOutput() ADLSGen1FileDataSetOutput {
	return i.ToADLSGen1FileDataSetOutputWithContext(context.Background())
}

func (i *ADLSGen1FileDataSet) ToADLSGen1FileDataSetOutputWithContext(ctx context.Context) ADLSGen1FileDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen1FileDataSetOutput)
}

type ADLSGen1FileDataSetOutput struct{ *pulumi.OutputState }

func (ADLSGen1FileDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ADLSGen1FileDataSet)(nil))
}

func (o ADLSGen1FileDataSetOutput) ToADLSGen1FileDataSetOutput() ADLSGen1FileDataSetOutput {
	return o
}

func (o ADLSGen1FileDataSetOutput) ToADLSGen1FileDataSetOutputWithContext(ctx context.Context) ADLSGen1FileDataSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ADLSGen1FileDataSetOutput{})
}
