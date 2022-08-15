


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ADLSGen2FolderDataSet struct {
	pulumi.CustomResourceState

	DataSetId          pulumi.StringOutput      `pulumi:"dataSetId"`
	FileSystem         pulumi.StringOutput      `pulumi:"fileSystem"`
	FolderPath         pulumi.StringOutput      `pulumi:"folderPath"`
	Kind               pulumi.StringOutput      `pulumi:"kind"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	ResourceGroup      pulumi.StringOutput      `pulumi:"resourceGroup"`
	StorageAccountName pulumi.StringOutput      `pulumi:"storageAccountName"`
	SubscriptionId     pulumi.StringOutput      `pulumi:"subscriptionId"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewADLSGen2FolderDataSet(ctx *pulumi.Context,
	name string, args *ADLSGen2FolderDataSetArgs, opts ...pulumi.ResourceOption) (*ADLSGen2FolderDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.FileSystem == nil {
		return nil, errors.New("invalid value for required argument 'FileSystem'")
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
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	args.Kind = pulumi.String("AdlsGen2Folder")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:ADLSGen2FolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ADLSGen2FolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen2FolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen2FolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ADLSGen2FolderDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource ADLSGen2FolderDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:ADLSGen2FolderDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetADLSGen2FolderDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen2FolderDataSetState, opts ...pulumi.ResourceOption) (*ADLSGen2FolderDataSet, error) {
	var resource ADLSGen2FolderDataSet
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:ADLSGen2FolderDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adlsgen2FolderDataSetState struct {
}

type ADLSGen2FolderDataSetState struct {
}

func (ADLSGen2FolderDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FolderDataSetState)(nil)).Elem()
}

type adlsgen2FolderDataSetArgs struct {
	AccountName        string  `pulumi:"accountName"`
	DataSetName        *string `pulumi:"dataSetName"`
	FileSystem         string  `pulumi:"fileSystem"`
	FolderPath         string  `pulumi:"folderPath"`
	Kind               string  `pulumi:"kind"`
	ResourceGroup      string  `pulumi:"resourceGroup"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	ShareName          string  `pulumi:"shareName"`
	StorageAccountName string  `pulumi:"storageAccountName"`
	SubscriptionId     string  `pulumi:"subscriptionId"`
}


type ADLSGen2FolderDataSetArgs struct {
	AccountName        pulumi.StringInput
	DataSetName        pulumi.StringPtrInput
	FileSystem         pulumi.StringInput
	FolderPath         pulumi.StringInput
	Kind               pulumi.StringInput
	ResourceGroup      pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	ShareName          pulumi.StringInput
	StorageAccountName pulumi.StringInput
	SubscriptionId     pulumi.StringInput
}

func (ADLSGen2FolderDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FolderDataSetArgs)(nil)).Elem()
}

type ADLSGen2FolderDataSetInput interface {
	pulumi.Input

	ToADLSGen2FolderDataSetOutput() ADLSGen2FolderDataSetOutput
	ToADLSGen2FolderDataSetOutputWithContext(ctx context.Context) ADLSGen2FolderDataSetOutput
}

func (*ADLSGen2FolderDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FolderDataSet)(nil)).Elem()
}

func (i *ADLSGen2FolderDataSet) ToADLSGen2FolderDataSetOutput() ADLSGen2FolderDataSetOutput {
	return i.ToADLSGen2FolderDataSetOutputWithContext(context.Background())
}

func (i *ADLSGen2FolderDataSet) ToADLSGen2FolderDataSetOutputWithContext(ctx context.Context) ADLSGen2FolderDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2FolderDataSetOutput)
}

type ADLSGen2FolderDataSetOutput struct{ *pulumi.OutputState }

func (ADLSGen2FolderDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FolderDataSet)(nil)).Elem()
}

func (o ADLSGen2FolderDataSetOutput) ToADLSGen2FolderDataSetOutput() ADLSGen2FolderDataSetOutput {
	return o
}

func (o ADLSGen2FolderDataSetOutput) ToADLSGen2FolderDataSetOutputWithContext(ctx context.Context) ADLSGen2FolderDataSetOutput {
	return o
}

func (o ADLSGen2FolderDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.FileSystem }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.FolderPath }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ADLSGen2FolderDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ADLSGen2FolderDataSetOutput{})
}
