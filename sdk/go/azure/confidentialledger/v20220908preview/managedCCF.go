


package v20220908preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedCCF struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                `pulumi:"location"`
	Name       pulumi.StringOutput                `pulumi:"name"`
	Properties ManagedCCFPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput           `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput             `pulumi:"tags"`
	Type       pulumi.StringOutput                `pulumi:"type"`
}


func NewManagedCCF(ctx *pulumi.Context,
	name string, args *ManagedCCFArgs, opts ...pulumi.ResourceOption) (*ManagedCCF, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ManagedCCF
	err := ctx.RegisterResource("azure-native:confidentialledger/v20220908preview:ManagedCCF", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedCCF(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedCCFState, opts ...pulumi.ResourceOption) (*ManagedCCF, error) {
	var resource ManagedCCF
	err := ctx.ReadResource("azure-native:confidentialledger/v20220908preview:ManagedCCF", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedCCFState struct {
}

type ManagedCCFState struct {
}

func (ManagedCCFState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedCCFState)(nil)).Elem()
}

type managedCCFArgs struct {
	AppName           *string               `pulumi:"appName"`
	Location          *string               `pulumi:"location"`
	Properties        *ManagedCCFProperties `pulumi:"properties"`
	ResourceGroupName string                `pulumi:"resourceGroupName"`
	Tags              map[string]string     `pulumi:"tags"`
}


type ManagedCCFArgs struct {
	AppName           pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        ManagedCCFPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (ManagedCCFArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedCCFArgs)(nil)).Elem()
}

type ManagedCCFInput interface {
	pulumi.Input

	ToManagedCCFOutput() ManagedCCFOutput
	ToManagedCCFOutputWithContext(ctx context.Context) ManagedCCFOutput
}

func (*ManagedCCF) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCCF)(nil)).Elem()
}

func (i *ManagedCCF) ToManagedCCFOutput() ManagedCCFOutput {
	return i.ToManagedCCFOutputWithContext(context.Background())
}

func (i *ManagedCCF) ToManagedCCFOutputWithContext(ctx context.Context) ManagedCCFOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedCCFOutput)
}

type ManagedCCFOutput struct{ *pulumi.OutputState }

func (ManagedCCFOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCCF)(nil)).Elem()
}

func (o ManagedCCFOutput) ToManagedCCFOutput() ManagedCCFOutput {
	return o
}

func (o ManagedCCFOutput) ToManagedCCFOutputWithContext(ctx context.Context) ManagedCCFOutput {
	return o
}

func (o ManagedCCFOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCCF) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ManagedCCFOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCCF) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedCCFOutput) Properties() ManagedCCFPropertiesResponseOutput {
	return o.ApplyT(func(v *ManagedCCF) ManagedCCFPropertiesResponseOutput { return v.Properties }).(ManagedCCFPropertiesResponseOutput)
}

func (o ManagedCCFOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedCCF) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ManagedCCFOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedCCF) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedCCFOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCCF) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedCCFOutput{})
}
