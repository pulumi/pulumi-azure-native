


package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ADLSGen2FileSystemDataSet struct {
	pulumi.CustomResourceState

	DataSetId          pulumi.StringOutput `pulumi:"dataSetId"`
	FileSystem         pulumi.StringOutput `pulumi:"fileSystem"`
	Kind               pulumi.StringOutput `pulumi:"kind"`
	Name               pulumi.StringOutput `pulumi:"name"`
	ResourceGroup      pulumi.StringOutput `pulumi:"resourceGroup"`
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	SubscriptionId     pulumi.StringOutput `pulumi:"subscriptionId"`
	Type               pulumi.StringOutput `pulumi:"type"`
}


func NewADLSGen2FileSystemDataSet(ctx *pulumi.Context,
	name string, args *ADLSGen2FileSystemDataSetArgs, opts ...pulumi.ResourceOption) (*ADLSGen2FileSystemDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
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
	args.Kind = pulumi.String("AdlsGen2FileSystem")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:ADLSGen2FileSystemDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen2FileSystemDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen2FileSystemDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ADLSGen2FileSystemDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ADLSGen2FileSystemDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource ADLSGen2FileSystemDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20181101preview:ADLSGen2FileSystemDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetADLSGen2FileSystemDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen2FileSystemDataSetState, opts ...pulumi.ResourceOption) (*ADLSGen2FileSystemDataSet, error) {
	var resource ADLSGen2FileSystemDataSet
	err := ctx.ReadResource("azure-native:datashare/v20181101preview:ADLSGen2FileSystemDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adlsgen2FileSystemDataSetState struct {
}

type ADLSGen2FileSystemDataSetState struct {
}

func (ADLSGen2FileSystemDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FileSystemDataSetState)(nil)).Elem()
}

type adlsgen2FileSystemDataSetArgs struct {
	AccountName        string  `pulumi:"accountName"`
	DataSetName        *string `pulumi:"dataSetName"`
	FileSystem         string  `pulumi:"fileSystem"`
	Kind               string  `pulumi:"kind"`
	ResourceGroup      string  `pulumi:"resourceGroup"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	ShareName          string  `pulumi:"shareName"`
	StorageAccountName string  `pulumi:"storageAccountName"`
	SubscriptionId     string  `pulumi:"subscriptionId"`
}


type ADLSGen2FileSystemDataSetArgs struct {
	AccountName        pulumi.StringInput
	DataSetName        pulumi.StringPtrInput
	FileSystem         pulumi.StringInput
	Kind               pulumi.StringInput
	ResourceGroup      pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	ShareName          pulumi.StringInput
	StorageAccountName pulumi.StringInput
	SubscriptionId     pulumi.StringInput
}

func (ADLSGen2FileSystemDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FileSystemDataSetArgs)(nil)).Elem()
}

type ADLSGen2FileSystemDataSetInput interface {
	pulumi.Input

	ToADLSGen2FileSystemDataSetOutput() ADLSGen2FileSystemDataSetOutput
	ToADLSGen2FileSystemDataSetOutputWithContext(ctx context.Context) ADLSGen2FileSystemDataSetOutput
}

func (*ADLSGen2FileSystemDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FileSystemDataSet)(nil)).Elem()
}

func (i *ADLSGen2FileSystemDataSet) ToADLSGen2FileSystemDataSetOutput() ADLSGen2FileSystemDataSetOutput {
	return i.ToADLSGen2FileSystemDataSetOutputWithContext(context.Background())
}

func (i *ADLSGen2FileSystemDataSet) ToADLSGen2FileSystemDataSetOutputWithContext(ctx context.Context) ADLSGen2FileSystemDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2FileSystemDataSetOutput)
}

type ADLSGen2FileSystemDataSetOutput struct{ *pulumi.OutputState }

func (ADLSGen2FileSystemDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FileSystemDataSet)(nil)).Elem()
}

func (o ADLSGen2FileSystemDataSetOutput) ToADLSGen2FileSystemDataSetOutput() ADLSGen2FileSystemDataSetOutput {
	return o
}

func (o ADLSGen2FileSystemDataSetOutput) ToADLSGen2FileSystemDataSetOutputWithContext(ctx context.Context) ADLSGen2FileSystemDataSetOutput {
	return o
}

func (o ADLSGen2FileSystemDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o ADLSGen2FileSystemDataSetOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSet) pulumi.StringOutput { return v.FileSystem }).(pulumi.StringOutput)
}

func (o ADLSGen2FileSystemDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ADLSGen2FileSystemDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ADLSGen2FileSystemDataSetOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSet) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o ADLSGen2FileSystemDataSetOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSet) pulumi.StringOutput { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o ADLSGen2FileSystemDataSetOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSet) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o ADLSGen2FileSystemDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ADLSGen2FileSystemDataSetOutput{})
}
