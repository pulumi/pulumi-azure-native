


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ADLSGen2FolderDataSetMapping struct {
	pulumi.CustomResourceState

	DataSetId            pulumi.StringOutput      `pulumi:"dataSetId"`
	DataSetMappingStatus pulumi.StringOutput      `pulumi:"dataSetMappingStatus"`
	FileSystem           pulumi.StringOutput      `pulumi:"fileSystem"`
	FolderPath           pulumi.StringOutput      `pulumi:"folderPath"`
	Kind                 pulumi.StringOutput      `pulumi:"kind"`
	Name                 pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput      `pulumi:"provisioningState"`
	ResourceGroup        pulumi.StringOutput      `pulumi:"resourceGroup"`
	StorageAccountName   pulumi.StringOutput      `pulumi:"storageAccountName"`
	SubscriptionId       pulumi.StringOutput      `pulumi:"subscriptionId"`
	SystemData           SystemDataResponseOutput `pulumi:"systemData"`
	Type                 pulumi.StringOutput      `pulumi:"type"`
}


func NewADLSGen2FolderDataSetMapping(ctx *pulumi.Context,
	name string, args *ADLSGen2FolderDataSetMappingArgs, opts ...pulumi.ResourceOption) (*ADLSGen2FolderDataSetMapping, error) {
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
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
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
			Type: pulumi.String("azure-native:datashare:ADLSGen2FolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ADLSGen2FolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen2FolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen2FolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ADLSGen2FolderDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource ADLSGen2FolderDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20210801:ADLSGen2FolderDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetADLSGen2FolderDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen2FolderDataSetMappingState, opts ...pulumi.ResourceOption) (*ADLSGen2FolderDataSetMapping, error) {
	var resource ADLSGen2FolderDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20210801:ADLSGen2FolderDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adlsgen2FolderDataSetMappingState struct {
}

type ADLSGen2FolderDataSetMappingState struct {
}

func (ADLSGen2FolderDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FolderDataSetMappingState)(nil)).Elem()
}

type adlsgen2FolderDataSetMappingArgs struct {
	AccountName           string  `pulumi:"accountName"`
	DataSetId             string  `pulumi:"dataSetId"`
	DataSetMappingName    *string `pulumi:"dataSetMappingName"`
	FileSystem            string  `pulumi:"fileSystem"`
	FolderPath            string  `pulumi:"folderPath"`
	Kind                  string  `pulumi:"kind"`
	ResourceGroup         string  `pulumi:"resourceGroup"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	StorageAccountName    string  `pulumi:"storageAccountName"`
	SubscriptionId        string  `pulumi:"subscriptionId"`
}


type ADLSGen2FolderDataSetMappingArgs struct {
	AccountName           pulumi.StringInput
	DataSetId             pulumi.StringInput
	DataSetMappingName    pulumi.StringPtrInput
	FileSystem            pulumi.StringInput
	FolderPath            pulumi.StringInput
	Kind                  pulumi.StringInput
	ResourceGroup         pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	ShareSubscriptionName pulumi.StringInput
	StorageAccountName    pulumi.StringInput
	SubscriptionId        pulumi.StringInput
}

func (ADLSGen2FolderDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FolderDataSetMappingArgs)(nil)).Elem()
}

type ADLSGen2FolderDataSetMappingInput interface {
	pulumi.Input

	ToADLSGen2FolderDataSetMappingOutput() ADLSGen2FolderDataSetMappingOutput
	ToADLSGen2FolderDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FolderDataSetMappingOutput
}

func (*ADLSGen2FolderDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FolderDataSetMapping)(nil)).Elem()
}

func (i *ADLSGen2FolderDataSetMapping) ToADLSGen2FolderDataSetMappingOutput() ADLSGen2FolderDataSetMappingOutput {
	return i.ToADLSGen2FolderDataSetMappingOutputWithContext(context.Background())
}

func (i *ADLSGen2FolderDataSetMapping) ToADLSGen2FolderDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FolderDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2FolderDataSetMappingOutput)
}

type ADLSGen2FolderDataSetMappingOutput struct{ *pulumi.OutputState }

func (ADLSGen2FolderDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FolderDataSetMapping)(nil)).Elem()
}

func (o ADLSGen2FolderDataSetMappingOutput) ToADLSGen2FolderDataSetMappingOutput() ADLSGen2FolderDataSetMappingOutput {
	return o
}

func (o ADLSGen2FolderDataSetMappingOutput) ToADLSGen2FolderDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FolderDataSetMappingOutput {
	return o
}

func (o ADLSGen2FolderDataSetMappingOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetMappingOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetMappingOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.FileSystem }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetMappingOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.FolderPath }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetMappingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetMappingOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetMappingOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetMappingOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o ADLSGen2FolderDataSetMappingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ADLSGen2FolderDataSetMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ADLSGen2FolderDataSetMappingOutput{})
}
