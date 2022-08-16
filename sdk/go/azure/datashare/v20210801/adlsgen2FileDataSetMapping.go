


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ADLSGen2FileDataSetMapping struct {
	pulumi.CustomResourceState

	DataSetId            pulumi.StringOutput      `pulumi:"dataSetId"`
	DataSetMappingStatus pulumi.StringOutput      `pulumi:"dataSetMappingStatus"`
	FilePath             pulumi.StringOutput      `pulumi:"filePath"`
	FileSystem           pulumi.StringOutput      `pulumi:"fileSystem"`
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


func NewADLSGen2FileDataSetMapping(ctx *pulumi.Context,
	name string, args *ADLSGen2FileDataSetMappingArgs, opts ...pulumi.ResourceOption) (*ADLSGen2FileDataSetMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DataSetId == nil {
		return nil, errors.New("invalid value for required argument 'DataSetId'")
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
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
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
			Type: pulumi.String("azure-native:datashare:ADLSGen2FileDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ADLSGen2FileDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen2FileDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen2FileDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ADLSGen2FileDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource ADLSGen2FileDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20210801:ADLSGen2FileDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetADLSGen2FileDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen2FileDataSetMappingState, opts ...pulumi.ResourceOption) (*ADLSGen2FileDataSetMapping, error) {
	var resource ADLSGen2FileDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20210801:ADLSGen2FileDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adlsgen2FileDataSetMappingState struct {
}

type ADLSGen2FileDataSetMappingState struct {
}

func (ADLSGen2FileDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FileDataSetMappingState)(nil)).Elem()
}

type adlsgen2FileDataSetMappingArgs struct {
	AccountName           string  `pulumi:"accountName"`
	DataSetId             string  `pulumi:"dataSetId"`
	DataSetMappingName    *string `pulumi:"dataSetMappingName"`
	FilePath              string  `pulumi:"filePath"`
	FileSystem            string  `pulumi:"fileSystem"`
	Kind                  string  `pulumi:"kind"`
	OutputType            *string `pulumi:"outputType"`
	ResourceGroup         string  `pulumi:"resourceGroup"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	StorageAccountName    string  `pulumi:"storageAccountName"`
	SubscriptionId        string  `pulumi:"subscriptionId"`
}


type ADLSGen2FileDataSetMappingArgs struct {
	AccountName           pulumi.StringInput
	DataSetId             pulumi.StringInput
	DataSetMappingName    pulumi.StringPtrInput
	FilePath              pulumi.StringInput
	FileSystem            pulumi.StringInput
	Kind                  pulumi.StringInput
	OutputType            pulumi.StringPtrInput
	ResourceGroup         pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	ShareSubscriptionName pulumi.StringInput
	StorageAccountName    pulumi.StringInput
	SubscriptionId        pulumi.StringInput
}

func (ADLSGen2FileDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FileDataSetMappingArgs)(nil)).Elem()
}

type ADLSGen2FileDataSetMappingInput interface {
	pulumi.Input

	ToADLSGen2FileDataSetMappingOutput() ADLSGen2FileDataSetMappingOutput
	ToADLSGen2FileDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FileDataSetMappingOutput
}

func (*ADLSGen2FileDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FileDataSetMapping)(nil)).Elem()
}

func (i *ADLSGen2FileDataSetMapping) ToADLSGen2FileDataSetMappingOutput() ADLSGen2FileDataSetMappingOutput {
	return i.ToADLSGen2FileDataSetMappingOutputWithContext(context.Background())
}

func (i *ADLSGen2FileDataSetMapping) ToADLSGen2FileDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FileDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2FileDataSetMappingOutput)
}

type ADLSGen2FileDataSetMappingOutput struct{ *pulumi.OutputState }

func (ADLSGen2FileDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FileDataSetMapping)(nil)).Elem()
}

func (o ADLSGen2FileDataSetMappingOutput) ToADLSGen2FileDataSetMappingOutput() ADLSGen2FileDataSetMappingOutput {
	return o
}

func (o ADLSGen2FileDataSetMappingOutput) ToADLSGen2FileDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FileDataSetMappingOutput {
	return o
}

func (o ADLSGen2FileDataSetMappingOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileDataSetMapping) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o ADLSGen2FileDataSetMappingOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileDataSetMapping) pulumi.StringOutput { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o ADLSGen2FileDataSetMappingOutput) FilePath() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileDataSetMapping) pulumi.StringOutput { return v.FilePath }).(pulumi.StringOutput)
}

func (o ADLSGen2FileDataSetMappingOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileDataSetMapping) pulumi.StringOutput { return v.FileSystem }).(pulumi.StringOutput)
}

func (o ADLSGen2FileDataSetMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileDataSetMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ADLSGen2FileDataSetMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileDataSetMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ADLSGen2FileDataSetMappingOutput) OutputType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ADLSGen2FileDataSetMapping) pulumi.StringPtrOutput { return v.OutputType }).(pulumi.StringPtrOutput)
}

func (o ADLSGen2FileDataSetMappingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileDataSetMapping) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ADLSGen2FileDataSetMappingOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileDataSetMapping) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o ADLSGen2FileDataSetMappingOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileDataSetMapping) pulumi.StringOutput { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o ADLSGen2FileDataSetMappingOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileDataSetMapping) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o ADLSGen2FileDataSetMappingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ADLSGen2FileDataSetMapping) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ADLSGen2FileDataSetMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileDataSetMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ADLSGen2FileDataSetMappingOutput{})
}
