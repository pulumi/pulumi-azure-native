


package v20150320

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type StorageInsight struct {
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


func NewStorageInsight(ctx *pulumi.Context,
	name string, args *StorageInsightArgs, opts ...pulumi.ResourceOption) (*StorageInsight, error) {
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
			Type: pulumi.String("azure-native:operationalinsights:StorageInsight"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:StorageInsight"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:StorageInsight"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageInsight
	err := ctx.RegisterResource("azure-native:operationalinsights/v20150320:StorageInsight", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageInsight(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageInsightState, opts ...pulumi.ResourceOption) (*StorageInsight, error) {
	var resource StorageInsight
	err := ctx.ReadResource("azure-native:operationalinsights/v20150320:StorageInsight", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storageInsightState struct {
}

type StorageInsightState struct {
}

func (StorageInsightState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageInsightState)(nil)).Elem()
}

type storageInsightArgs struct {
	Containers         []string          `pulumi:"containers"`
	ETag               *string           `pulumi:"eTag"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	StorageAccount     StorageAccount    `pulumi:"storageAccount"`
	StorageInsightName *string           `pulumi:"storageInsightName"`
	Tables             []string          `pulumi:"tables"`
	Tags               map[string]string `pulumi:"tags"`
	WorkspaceName      string            `pulumi:"workspaceName"`
}


type StorageInsightArgs struct {
	Containers         pulumi.StringArrayInput
	ETag               pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	StorageAccount     StorageAccountInput
	StorageInsightName pulumi.StringPtrInput
	Tables             pulumi.StringArrayInput
	Tags               pulumi.StringMapInput
	WorkspaceName      pulumi.StringInput
}

func (StorageInsightArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageInsightArgs)(nil)).Elem()
}

type StorageInsightInput interface {
	pulumi.Input

	ToStorageInsightOutput() StorageInsightOutput
	ToStorageInsightOutputWithContext(ctx context.Context) StorageInsightOutput
}

func (*StorageInsight) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageInsight)(nil)).Elem()
}

func (i *StorageInsight) ToStorageInsightOutput() StorageInsightOutput {
	return i.ToStorageInsightOutputWithContext(context.Background())
}

func (i *StorageInsight) ToStorageInsightOutputWithContext(ctx context.Context) StorageInsightOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageInsightOutput)
}

type StorageInsightOutput struct{ *pulumi.OutputState }

func (StorageInsightOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageInsight)(nil)).Elem()
}

func (o StorageInsightOutput) ToStorageInsightOutput() StorageInsightOutput {
	return o
}

func (o StorageInsightOutput) ToStorageInsightOutputWithContext(ctx context.Context) StorageInsightOutput {
	return o
}

func (o StorageInsightOutput) Containers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StorageInsight) pulumi.StringArrayOutput { return v.Containers }).(pulumi.StringArrayOutput)
}

func (o StorageInsightOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageInsight) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o StorageInsightOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageInsight) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StorageInsightOutput) Status() StorageInsightStatusResponseOutput {
	return o.ApplyT(func(v *StorageInsight) StorageInsightStatusResponseOutput { return v.Status }).(StorageInsightStatusResponseOutput)
}

func (o StorageInsightOutput) StorageAccount() StorageAccountResponseOutput {
	return o.ApplyT(func(v *StorageInsight) StorageAccountResponseOutput { return v.StorageAccount }).(StorageAccountResponseOutput)
}

func (o StorageInsightOutput) Tables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StorageInsight) pulumi.StringArrayOutput { return v.Tables }).(pulumi.StringArrayOutput)
}

func (o StorageInsightOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageInsight) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o StorageInsightOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageInsight) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageInsightOutput{})
}
