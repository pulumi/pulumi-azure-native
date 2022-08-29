


package v20170228preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ReferenceDataSet struct {
	pulumi.CustomResourceState

	CreationTime      pulumi.StringOutput                            `pulumi:"creationTime"`
	KeyProperties     ReferenceDataSetKeyPropertyResponseArrayOutput `pulumi:"keyProperties"`
	Location          pulumi.StringOutput                            `pulumi:"location"`
	Name              pulumi.StringOutput                            `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                            `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput                         `pulumi:"tags"`
	Type              pulumi.StringOutput                            `pulumi:"type"`
}


func NewReferenceDataSet(ctx *pulumi.Context,
	name string, args *ReferenceDataSetArgs, opts ...pulumi.ResourceOption) (*ReferenceDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.KeyProperties == nil {
		return nil, errors.New("invalid value for required argument 'KeyProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:timeseriesinsights:ReferenceDataSet"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:ReferenceDataSet"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20180815preview:ReferenceDataSet"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20200515:ReferenceDataSet"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:ReferenceDataSet"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:ReferenceDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource ReferenceDataSet
	err := ctx.RegisterResource("azure-native:timeseriesinsights/v20170228preview:ReferenceDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReferenceDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReferenceDataSetState, opts ...pulumi.ResourceOption) (*ReferenceDataSet, error) {
	var resource ReferenceDataSet
	err := ctx.ReadResource("azure-native:timeseriesinsights/v20170228preview:ReferenceDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type referenceDataSetState struct {
}

type ReferenceDataSetState struct {
}

func (ReferenceDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceDataSetState)(nil)).Elem()
}

type referenceDataSetArgs struct {
	EnvironmentName      string                        `pulumi:"environmentName"`
	KeyProperties        []ReferenceDataSetKeyProperty `pulumi:"keyProperties"`
	Location             *string                       `pulumi:"location"`
	ReferenceDataSetName *string                       `pulumi:"referenceDataSetName"`
	ResourceGroupName    string                        `pulumi:"resourceGroupName"`
	Tags                 map[string]string             `pulumi:"tags"`
}


type ReferenceDataSetArgs struct {
	EnvironmentName      pulumi.StringInput
	KeyProperties        ReferenceDataSetKeyPropertyArrayInput
	Location             pulumi.StringPtrInput
	ReferenceDataSetName pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
}

func (ReferenceDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceDataSetArgs)(nil)).Elem()
}

type ReferenceDataSetInput interface {
	pulumi.Input

	ToReferenceDataSetOutput() ReferenceDataSetOutput
	ToReferenceDataSetOutputWithContext(ctx context.Context) ReferenceDataSetOutput
}

func (*ReferenceDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceDataSet)(nil)).Elem()
}

func (i *ReferenceDataSet) ToReferenceDataSetOutput() ReferenceDataSetOutput {
	return i.ToReferenceDataSetOutputWithContext(context.Background())
}

func (i *ReferenceDataSet) ToReferenceDataSetOutputWithContext(ctx context.Context) ReferenceDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceDataSetOutput)
}

type ReferenceDataSetOutput struct{ *pulumi.OutputState }

func (ReferenceDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceDataSet)(nil)).Elem()
}

func (o ReferenceDataSetOutput) ToReferenceDataSetOutput() ReferenceDataSetOutput {
	return o
}

func (o ReferenceDataSetOutput) ToReferenceDataSetOutputWithContext(ctx context.Context) ReferenceDataSetOutput {
	return o
}

func (o ReferenceDataSetOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ReferenceDataSet) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o ReferenceDataSetOutput) KeyProperties() ReferenceDataSetKeyPropertyResponseArrayOutput {
	return o.ApplyT(func(v *ReferenceDataSet) ReferenceDataSetKeyPropertyResponseArrayOutput { return v.KeyProperties }).(ReferenceDataSetKeyPropertyResponseArrayOutput)
}

func (o ReferenceDataSetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ReferenceDataSet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ReferenceDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReferenceDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReferenceDataSetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ReferenceDataSet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ReferenceDataSetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReferenceDataSet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ReferenceDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReferenceDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReferenceDataSetOutput{})
}
