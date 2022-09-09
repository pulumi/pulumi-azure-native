


package web

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StaticSiteLinkedBackend struct {
	pulumi.CustomResourceState

	BackendResourceId pulumi.StringPtrOutput `pulumi:"backendResourceId"`
	CreatedOn         pulumi.StringOutput    `pulumi:"createdOn"`
	Kind              pulumi.StringPtrOutput `pulumi:"kind"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Region            pulumi.StringPtrOutput `pulumi:"region"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewStaticSiteLinkedBackend(ctx *pulumi.Context,
	name string, args *StaticSiteLinkedBackendArgs, opts ...pulumi.ResourceOption) (*StaticSiteLinkedBackend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20220301:StaticSiteLinkedBackend"),
		},
	})
	opts = append(opts, aliases)
	var resource StaticSiteLinkedBackend
	err := ctx.RegisterResource("azure-native:web:StaticSiteLinkedBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStaticSiteLinkedBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSiteLinkedBackendState, opts ...pulumi.ResourceOption) (*StaticSiteLinkedBackend, error) {
	var resource StaticSiteLinkedBackend
	err := ctx.ReadResource("azure-native:web:StaticSiteLinkedBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type staticSiteLinkedBackendState struct {
}

type StaticSiteLinkedBackendState struct {
}

func (StaticSiteLinkedBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteLinkedBackendState)(nil)).Elem()
}

type staticSiteLinkedBackendArgs struct {
	BackendResourceId *string `pulumi:"backendResourceId"`
	Kind              *string `pulumi:"kind"`
	LinkedBackendName *string `pulumi:"linkedBackendName"`
	Name              string  `pulumi:"name"`
	Region            *string `pulumi:"region"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type StaticSiteLinkedBackendArgs struct {
	BackendResourceId pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	LinkedBackendName pulumi.StringPtrInput
	Name              pulumi.StringInput
	Region            pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (StaticSiteLinkedBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteLinkedBackendArgs)(nil)).Elem()
}

type StaticSiteLinkedBackendInput interface {
	pulumi.Input

	ToStaticSiteLinkedBackendOutput() StaticSiteLinkedBackendOutput
	ToStaticSiteLinkedBackendOutputWithContext(ctx context.Context) StaticSiteLinkedBackendOutput
}

func (*StaticSiteLinkedBackend) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteLinkedBackend)(nil)).Elem()
}

func (i *StaticSiteLinkedBackend) ToStaticSiteLinkedBackendOutput() StaticSiteLinkedBackendOutput {
	return i.ToStaticSiteLinkedBackendOutputWithContext(context.Background())
}

func (i *StaticSiteLinkedBackend) ToStaticSiteLinkedBackendOutputWithContext(ctx context.Context) StaticSiteLinkedBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteLinkedBackendOutput)
}

type StaticSiteLinkedBackendOutput struct{ *pulumi.OutputState }

func (StaticSiteLinkedBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteLinkedBackend)(nil)).Elem()
}

func (o StaticSiteLinkedBackendOutput) ToStaticSiteLinkedBackendOutput() StaticSiteLinkedBackendOutput {
	return o
}

func (o StaticSiteLinkedBackendOutput) ToStaticSiteLinkedBackendOutputWithContext(ctx context.Context) StaticSiteLinkedBackendOutput {
	return o
}

func (o StaticSiteLinkedBackendOutput) BackendResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackend) pulumi.StringPtrOutput { return v.BackendResourceId }).(pulumi.StringPtrOutput)
}

func (o StaticSiteLinkedBackendOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackend) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o StaticSiteLinkedBackendOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackend) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o StaticSiteLinkedBackendOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackend) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StaticSiteLinkedBackendOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackend) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o StaticSiteLinkedBackendOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackend) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

func (o StaticSiteLinkedBackendOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteLinkedBackend) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticSiteLinkedBackendOutput{})
}
