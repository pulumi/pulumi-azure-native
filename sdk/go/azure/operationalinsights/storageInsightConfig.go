


package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageInsightConfig struct {
	pulumi.CustomResourceState

	Containers     pulumi.StringArrayOutput           `pulumi:"containers"`
	ETag           pulumi.StringPtrOutput             `pulumi:"eTag"`
	Name           pulumi.StringOutput                `pulumi:"name"`
	Status         StorageInsightStatusResponseOutput `pulumi:"status"`
	StorageAccount StorageAccountResponseOutput       `pulumi:"storageAccount"`
	Tables         pulumi.StringArrayOutput           `pulumi:"tables"`
	Tags           pulumi.StringMapOutput             `pulumi:"tags"`
	Type           pulumi.StringOutput                `pulumi:"type"`
}


func NewStorageInsightConfig(ctx *pulumi.Context,
	name string, args *StorageInsightConfigArgs, opts ...pulumi.ResourceOption) (*StorageInsightConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccount == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccount'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights/v20150320:StorageInsightConfig"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:StorageInsightConfig"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:StorageInsightConfig"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageInsightConfig
	err := ctx.RegisterResource("azure-native:operationalinsights:StorageInsightConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageInsightConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageInsightConfigState, opts ...pulumi.ResourceOption) (*StorageInsightConfig, error) {
	var resource StorageInsightConfig
	err := ctx.ReadResource("azure-native:operationalinsights:StorageInsightConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storageInsightConfigState struct {
}

type StorageInsightConfigState struct {
}

func (StorageInsightConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageInsightConfigState)(nil)).Elem()
}

type storageInsightConfigArgs struct {
	Containers         []string          `pulumi:"containers"`
	ETag               *string           `pulumi:"eTag"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	StorageAccount     StorageAccount    `pulumi:"storageAccount"`
	StorageInsightName *string           `pulumi:"storageInsightName"`
	Tables             []string          `pulumi:"tables"`
	Tags               map[string]string `pulumi:"tags"`
	WorkspaceName      string            `pulumi:"workspaceName"`
}


type StorageInsightConfigArgs struct {
	Containers         pulumi.StringArrayInput
	ETag               pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	StorageAccount     StorageAccountInput
	StorageInsightName pulumi.StringPtrInput
	Tables             pulumi.StringArrayInput
	Tags               pulumi.StringMapInput
	WorkspaceName      pulumi.StringInput
}

func (StorageInsightConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageInsightConfigArgs)(nil)).Elem()
}

type StorageInsightConfigInput interface {
	pulumi.Input

	ToStorageInsightConfigOutput() StorageInsightConfigOutput
	ToStorageInsightConfigOutputWithContext(ctx context.Context) StorageInsightConfigOutput
}

func (*StorageInsightConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageInsightConfig)(nil)).Elem()
}

func (i *StorageInsightConfig) ToStorageInsightConfigOutput() StorageInsightConfigOutput {
	return i.ToStorageInsightConfigOutputWithContext(context.Background())
}

func (i *StorageInsightConfig) ToStorageInsightConfigOutputWithContext(ctx context.Context) StorageInsightConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageInsightConfigOutput)
}

type StorageInsightConfigOutput struct{ *pulumi.OutputState }

func (StorageInsightConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageInsightConfig)(nil)).Elem()
}

func (o StorageInsightConfigOutput) ToStorageInsightConfigOutput() StorageInsightConfigOutput {
	return o
}

func (o StorageInsightConfigOutput) ToStorageInsightConfigOutputWithContext(ctx context.Context) StorageInsightConfigOutput {
	return o
}

func (o StorageInsightConfigOutput) Containers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StorageInsightConfig) pulumi.StringArrayOutput { return v.Containers }).(pulumi.StringArrayOutput)
}

func (o StorageInsightConfigOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageInsightConfig) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o StorageInsightConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageInsightConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StorageInsightConfigOutput) Status() StorageInsightStatusResponseOutput {
	return o.ApplyT(func(v *StorageInsightConfig) StorageInsightStatusResponseOutput { return v.Status }).(StorageInsightStatusResponseOutput)
}

func (o StorageInsightConfigOutput) StorageAccount() StorageAccountResponseOutput {
	return o.ApplyT(func(v *StorageInsightConfig) StorageAccountResponseOutput { return v.StorageAccount }).(StorageAccountResponseOutput)
}

func (o StorageInsightConfigOutput) Tables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StorageInsightConfig) pulumi.StringArrayOutput { return v.Tables }).(pulumi.StringArrayOutput)
}

func (o StorageInsightConfigOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageInsightConfig) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o StorageInsightConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageInsightConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageInsightConfigOutput{})
}
