


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvironmentType struct {
	pulumi.CustomResourceState

	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewEnvironmentType(ctx *pulumi.Context,
	name string, args *EnvironmentTypeArgs, opts ...pulumi.ResourceOption) (*EnvironmentType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevCenterName == nil {
		return nil, errors.New("invalid value for required argument 'DevCenterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:EnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:EnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:EnvironmentType"),
		},
	})
	opts = append(opts, aliases)
	var resource EnvironmentType
	err := ctx.RegisterResource("azure-native:devcenter/v20220801preview:EnvironmentType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnvironmentType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentTypeState, opts ...pulumi.ResourceOption) (*EnvironmentType, error) {
	var resource EnvironmentType
	err := ctx.ReadResource("azure-native:devcenter/v20220801preview:EnvironmentType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type environmentTypeState struct {
}

type EnvironmentTypeState struct {
}

func (EnvironmentTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentTypeState)(nil)).Elem()
}

type environmentTypeArgs struct {
	DevCenterName       string            `pulumi:"devCenterName"`
	EnvironmentTypeName *string           `pulumi:"environmentTypeName"`
	ResourceGroupName   string            `pulumi:"resourceGroupName"`
	Tags                map[string]string `pulumi:"tags"`
}


type EnvironmentTypeArgs struct {
	DevCenterName       pulumi.StringInput
	EnvironmentTypeName pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
}

func (EnvironmentTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentTypeArgs)(nil)).Elem()
}

type EnvironmentTypeInput interface {
	pulumi.Input

	ToEnvironmentTypeOutput() EnvironmentTypeOutput
	ToEnvironmentTypeOutputWithContext(ctx context.Context) EnvironmentTypeOutput
}

func (*EnvironmentType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentType)(nil)).Elem()
}

func (i *EnvironmentType) ToEnvironmentTypeOutput() EnvironmentTypeOutput {
	return i.ToEnvironmentTypeOutputWithContext(context.Background())
}

func (i *EnvironmentType) ToEnvironmentTypeOutputWithContext(ctx context.Context) EnvironmentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentTypeOutput)
}

type EnvironmentTypeOutput struct{ *pulumi.OutputState }

func (EnvironmentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentType)(nil)).Elem()
}

func (o EnvironmentTypeOutput) ToEnvironmentTypeOutput() EnvironmentTypeOutput {
	return o
}

func (o EnvironmentTypeOutput) ToEnvironmentTypeOutputWithContext(ctx context.Context) EnvironmentTypeOutput {
	return o
}

func (o EnvironmentTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EnvironmentTypeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentType) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o EnvironmentTypeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EnvironmentType) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o EnvironmentTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EnvironmentType) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o EnvironmentTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentType) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentTypeOutput{})
}
