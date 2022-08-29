


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataController struct {
	pulumi.CustomResourceState

	ExtendedLocation ExtendedLocationResponsePtrOutput      `pulumi:"extendedLocation"`
	Location         pulumi.StringOutput                    `pulumi:"location"`
	Name             pulumi.StringOutput                    `pulumi:"name"`
	Properties       DataControllerPropertiesResponseOutput `pulumi:"properties"`
	SystemData       SystemDataResponseOutput               `pulumi:"systemData"`
	Tags             pulumi.StringMapOutput                 `pulumi:"tags"`
	Type             pulumi.StringOutput                    `pulumi:"type"`
}


func NewDataController(ctx *pulumi.Context,
	name string, args *DataControllerArgs, opts ...pulumi.ResourceOption) (*DataController, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Properties = args.Properties.ToDataControllerPropertiesOutput().ApplyT(func(v DataControllerProperties) DataControllerProperties { return *v.Defaults() }).(DataControllerPropertiesOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurearcdata:DataController"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210601preview:DataController"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210701preview:DataController"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210801:DataController"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20220301preview:DataController"),
		},
	})
	opts = append(opts, aliases)
	var resource DataController
	err := ctx.RegisterResource("azure-native:azurearcdata/v20211101:DataController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataControllerState, opts ...pulumi.ResourceOption) (*DataController, error) {
	var resource DataController
	err := ctx.ReadResource("azure-native:azurearcdata/v20211101:DataController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataControllerState struct {
}

type DataControllerState struct {
}

func (DataControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataControllerState)(nil)).Elem()
}

type dataControllerArgs struct {
	DataControllerName *string                  `pulumi:"dataControllerName"`
	ExtendedLocation   *ExtendedLocation        `pulumi:"extendedLocation"`
	Location           *string                  `pulumi:"location"`
	Properties         DataControllerProperties `pulumi:"properties"`
	ResourceGroupName  string                   `pulumi:"resourceGroupName"`
	Tags               map[string]string        `pulumi:"tags"`
}


type DataControllerArgs struct {
	DataControllerName pulumi.StringPtrInput
	ExtendedLocation   ExtendedLocationPtrInput
	Location           pulumi.StringPtrInput
	Properties         DataControllerPropertiesInput
	ResourceGroupName  pulumi.StringInput
	Tags               pulumi.StringMapInput
}

func (DataControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataControllerArgs)(nil)).Elem()
}

type DataControllerInput interface {
	pulumi.Input

	ToDataControllerOutput() DataControllerOutput
	ToDataControllerOutputWithContext(ctx context.Context) DataControllerOutput
}

func (*DataController) ElementType() reflect.Type {
	return reflect.TypeOf((**DataController)(nil)).Elem()
}

func (i *DataController) ToDataControllerOutput() DataControllerOutput {
	return i.ToDataControllerOutputWithContext(context.Background())
}

func (i *DataController) ToDataControllerOutputWithContext(ctx context.Context) DataControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataControllerOutput)
}

type DataControllerOutput struct{ *pulumi.OutputState }

func (DataControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataController)(nil)).Elem()
}

func (o DataControllerOutput) ToDataControllerOutput() DataControllerOutput {
	return o
}

func (o DataControllerOutput) ToDataControllerOutputWithContext(ctx context.Context) DataControllerOutput {
	return o
}

func (o DataControllerOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *DataController) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o DataControllerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataController) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataController) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataControllerOutput) Properties() DataControllerPropertiesResponseOutput {
	return o.ApplyT(func(v *DataController) DataControllerPropertiesResponseOutput { return v.Properties }).(DataControllerPropertiesResponseOutput)
}

func (o DataControllerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DataController) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DataControllerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataController) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DataControllerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataController) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataControllerOutput{})
}
