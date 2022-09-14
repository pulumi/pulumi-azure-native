


package agfoodplatform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FarmBeatsModel struct {
	pulumi.CustomResourceState

	InstanceUri       pulumi.StringOutput      `pulumi:"instanceUri"`
	Location          pulumi.StringOutput      `pulumi:"location"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewFarmBeatsModel(ctx *pulumi.Context,
	name string, args *FarmBeatsModelArgs, opts ...pulumi.ResourceOption) (*FarmBeatsModel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:agfoodplatform/v20200512preview:FarmBeatsModel"),
		},
		{
			Type: pulumi.String("azure-native:agfoodplatform/v20210901preview:FarmBeatsModel"),
		},
	})
	opts = append(opts, aliases)
	var resource FarmBeatsModel
	err := ctx.RegisterResource("azure-native:agfoodplatform:FarmBeatsModel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFarmBeatsModel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FarmBeatsModelState, opts ...pulumi.ResourceOption) (*FarmBeatsModel, error) {
	var resource FarmBeatsModel
	err := ctx.ReadResource("azure-native:agfoodplatform:FarmBeatsModel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type farmBeatsModelState struct {
}

type FarmBeatsModelState struct {
}

func (FarmBeatsModelState) ElementType() reflect.Type {
	return reflect.TypeOf((*farmBeatsModelState)(nil)).Elem()
}

type farmBeatsModelArgs struct {
	FarmBeatsResourceName *string           `pulumi:"farmBeatsResourceName"`
	Location              *string           `pulumi:"location"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	Tags                  map[string]string `pulumi:"tags"`
}


type FarmBeatsModelArgs struct {
	FarmBeatsResourceName pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Tags                  pulumi.StringMapInput
}

func (FarmBeatsModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*farmBeatsModelArgs)(nil)).Elem()
}

type FarmBeatsModelInput interface {
	pulumi.Input

	ToFarmBeatsModelOutput() FarmBeatsModelOutput
	ToFarmBeatsModelOutputWithContext(ctx context.Context) FarmBeatsModelOutput
}

func (*FarmBeatsModel) ElementType() reflect.Type {
	return reflect.TypeOf((**FarmBeatsModel)(nil)).Elem()
}

func (i *FarmBeatsModel) ToFarmBeatsModelOutput() FarmBeatsModelOutput {
	return i.ToFarmBeatsModelOutputWithContext(context.Background())
}

func (i *FarmBeatsModel) ToFarmBeatsModelOutputWithContext(ctx context.Context) FarmBeatsModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FarmBeatsModelOutput)
}

type FarmBeatsModelOutput struct{ *pulumi.OutputState }

func (FarmBeatsModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FarmBeatsModel)(nil)).Elem()
}

func (o FarmBeatsModelOutput) ToFarmBeatsModelOutput() FarmBeatsModelOutput {
	return o
}

func (o FarmBeatsModelOutput) ToFarmBeatsModelOutputWithContext(ctx context.Context) FarmBeatsModelOutput {
	return o
}

func (o FarmBeatsModelOutput) InstanceUri() pulumi.StringOutput {
	return o.ApplyT(func(v *FarmBeatsModel) pulumi.StringOutput { return v.InstanceUri }).(pulumi.StringOutput)
}

func (o FarmBeatsModelOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FarmBeatsModel) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o FarmBeatsModelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FarmBeatsModel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FarmBeatsModelOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FarmBeatsModel) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o FarmBeatsModelOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FarmBeatsModel) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o FarmBeatsModelOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FarmBeatsModel) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o FarmBeatsModelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FarmBeatsModel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FarmBeatsModelOutput{})
}
