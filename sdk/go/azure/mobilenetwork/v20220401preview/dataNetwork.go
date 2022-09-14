


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataNetwork struct {
	pulumi.CustomResourceState

	CreatedAt          pulumi.StringPtrOutput   `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrOutput   `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrOutput   `pulumi:"createdByType"`
	Description        pulumi.StringPtrOutput   `pulumi:"description"`
	LastModifiedAt     pulumi.StringPtrOutput   `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrOutput   `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrOutput   `pulumi:"lastModifiedByType"`
	Location           pulumi.StringOutput      `pulumi:"location"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput   `pulumi:"tags"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewDataNetwork(ctx *pulumi.Context,
	name string, args *DataNetworkArgs, opts ...pulumi.ResourceOption) (*DataNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MobileNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'MobileNetworkName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:DataNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:DataNetwork"),
		},
	})
	opts = append(opts, aliases)
	var resource DataNetwork
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20220401preview:DataNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataNetworkState, opts ...pulumi.ResourceOption) (*DataNetwork, error) {
	var resource DataNetwork
	err := ctx.ReadResource("azure-native:mobilenetwork/v20220401preview:DataNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataNetworkState struct {
}

type DataNetworkState struct {
}

func (DataNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataNetworkState)(nil)).Elem()
}

type dataNetworkArgs struct {
	CreatedAt          *string           `pulumi:"createdAt"`
	CreatedBy          *string           `pulumi:"createdBy"`
	CreatedByType      *string           `pulumi:"createdByType"`
	DataNetworkName    *string           `pulumi:"dataNetworkName"`
	Description        *string           `pulumi:"description"`
	LastModifiedAt     *string           `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string           `pulumi:"lastModifiedBy"`
	LastModifiedByType *string           `pulumi:"lastModifiedByType"`
	Location           *string           `pulumi:"location"`
	MobileNetworkName  string            `pulumi:"mobileNetworkName"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	Tags               map[string]string `pulumi:"tags"`
}


type DataNetworkArgs struct {
	CreatedAt          pulumi.StringPtrInput
	CreatedBy          pulumi.StringPtrInput
	CreatedByType      pulumi.StringPtrInput
	DataNetworkName    pulumi.StringPtrInput
	Description        pulumi.StringPtrInput
	LastModifiedAt     pulumi.StringPtrInput
	LastModifiedBy     pulumi.StringPtrInput
	LastModifiedByType pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	MobileNetworkName  pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	Tags               pulumi.StringMapInput
}

func (DataNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataNetworkArgs)(nil)).Elem()
}

type DataNetworkInput interface {
	pulumi.Input

	ToDataNetworkOutput() DataNetworkOutput
	ToDataNetworkOutputWithContext(ctx context.Context) DataNetworkOutput
}

func (*DataNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**DataNetwork)(nil)).Elem()
}

func (i *DataNetwork) ToDataNetworkOutput() DataNetworkOutput {
	return i.ToDataNetworkOutputWithContext(context.Background())
}

func (i *DataNetwork) ToDataNetworkOutputWithContext(ctx context.Context) DataNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataNetworkOutput)
}

type DataNetworkOutput struct{ *pulumi.OutputState }

func (DataNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataNetwork)(nil)).Elem()
}

func (o DataNetworkOutput) ToDataNetworkOutput() DataNetworkOutput {
	return o
}

func (o DataNetworkOutput) ToDataNetworkOutputWithContext(ctx context.Context) DataNetworkOutput {
	return o
}

func (o DataNetworkOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o DataNetworkOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o DataNetworkOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o DataNetworkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataNetworkOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o DataNetworkOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o DataNetworkOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o DataNetworkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DataNetworkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DataNetwork) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DataNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DataNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataNetworkOutput{})
}
