


package v20201031

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DigitalTwin struct {
	pulumi.CustomResourceState

	CreatedTime       pulumi.StringOutput    `pulumi:"createdTime"`
	HostName          pulumi.StringOutput    `pulumi:"hostName"`
	LastUpdatedTime   pulumi.StringOutput    `pulumi:"lastUpdatedTime"`
	Location          pulumi.StringOutput    `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewDigitalTwin(ctx *pulumi.Context,
	name string, args *DigitalTwinArgs, opts ...pulumi.ResourceOption) (*DigitalTwin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:digitaltwins:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20200301preview:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20201201:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20210630preview:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20220531:DigitalTwin"),
		},
	})
	opts = append(opts, aliases)
	var resource DigitalTwin
	err := ctx.RegisterResource("azure-native:digitaltwins/v20201031:DigitalTwin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDigitalTwin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DigitalTwinState, opts ...pulumi.ResourceOption) (*DigitalTwin, error) {
	var resource DigitalTwin
	err := ctx.ReadResource("azure-native:digitaltwins/v20201031:DigitalTwin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type digitalTwinState struct {
}

type DigitalTwinState struct {
}

func (DigitalTwinState) ElementType() reflect.Type {
	return reflect.TypeOf((*digitalTwinState)(nil)).Elem()
}

type digitalTwinArgs struct {
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	Tags              map[string]string `pulumi:"tags"`
}


type DigitalTwinArgs struct {
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (DigitalTwinArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*digitalTwinArgs)(nil)).Elem()
}

type DigitalTwinInput interface {
	pulumi.Input

	ToDigitalTwinOutput() DigitalTwinOutput
	ToDigitalTwinOutputWithContext(ctx context.Context) DigitalTwinOutput
}

func (*DigitalTwin) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwin)(nil)).Elem()
}

func (i *DigitalTwin) ToDigitalTwinOutput() DigitalTwinOutput {
	return i.ToDigitalTwinOutputWithContext(context.Background())
}

func (i *DigitalTwin) ToDigitalTwinOutputWithContext(ctx context.Context) DigitalTwinOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinOutput)
}

type DigitalTwinOutput struct{ *pulumi.OutputState }

func (DigitalTwinOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwin)(nil)).Elem()
}

func (o DigitalTwinOutput) ToDigitalTwinOutput() DigitalTwinOutput {
	return o
}

func (o DigitalTwinOutput) ToDigitalTwinOutputWithContext(ctx context.Context) DigitalTwinOutput {
	return o
}

func (o DigitalTwinOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o DigitalTwinOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

func (o DigitalTwinOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o DigitalTwinOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DigitalTwinOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DigitalTwinOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DigitalTwinOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DigitalTwinOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DigitalTwinOutput{})
}
