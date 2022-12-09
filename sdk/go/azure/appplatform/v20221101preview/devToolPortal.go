


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DevToolPortal struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                   `pulumi:"name"`
	Properties DevToolPortalPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput              `pulumi:"systemData"`
	Type       pulumi.StringOutput                   `pulumi:"type"`
}


func NewDevToolPortal(ctx *pulumi.Context,
	name string, args *DevToolPortalArgs, opts ...pulumi.ResourceOption) (*DevToolPortal, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToDevToolPortalPropertiesPtrOutput().ApplyT(func(v *DevToolPortalProperties) *DevToolPortalProperties { return v.Defaults() }).(DevToolPortalPropertiesPtrOutput)
	}
	var resource DevToolPortal
	err := ctx.RegisterResource("azure-native:appplatform/v20221101preview:DevToolPortal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDevToolPortal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DevToolPortalState, opts ...pulumi.ResourceOption) (*DevToolPortal, error) {
	var resource DevToolPortal
	err := ctx.ReadResource("azure-native:appplatform/v20221101preview:DevToolPortal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type devToolPortalState struct {
}

type DevToolPortalState struct {
}

func (DevToolPortalState) ElementType() reflect.Type {
	return reflect.TypeOf((*devToolPortalState)(nil)).Elem()
}

type devToolPortalArgs struct {
	DevToolPortalName *string                  `pulumi:"devToolPortalName"`
	Properties        *DevToolPortalProperties `pulumi:"properties"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	ServiceName       string                   `pulumi:"serviceName"`
}


type DevToolPortalArgs struct {
	DevToolPortalName pulumi.StringPtrInput
	Properties        DevToolPortalPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (DevToolPortalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*devToolPortalArgs)(nil)).Elem()
}

type DevToolPortalInput interface {
	pulumi.Input

	ToDevToolPortalOutput() DevToolPortalOutput
	ToDevToolPortalOutputWithContext(ctx context.Context) DevToolPortalOutput
}

func (*DevToolPortal) ElementType() reflect.Type {
	return reflect.TypeOf((**DevToolPortal)(nil)).Elem()
}

func (i *DevToolPortal) ToDevToolPortalOutput() DevToolPortalOutput {
	return i.ToDevToolPortalOutputWithContext(context.Background())
}

func (i *DevToolPortal) ToDevToolPortalOutputWithContext(ctx context.Context) DevToolPortalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevToolPortalOutput)
}

type DevToolPortalOutput struct{ *pulumi.OutputState }

func (DevToolPortalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DevToolPortal)(nil)).Elem()
}

func (o DevToolPortalOutput) ToDevToolPortalOutput() DevToolPortalOutput {
	return o
}

func (o DevToolPortalOutput) ToDevToolPortalOutputWithContext(ctx context.Context) DevToolPortalOutput {
	return o
}

func (o DevToolPortalOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DevToolPortal) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DevToolPortalOutput) Properties() DevToolPortalPropertiesResponseOutput {
	return o.ApplyT(func(v *DevToolPortal) DevToolPortalPropertiesResponseOutput { return v.Properties }).(DevToolPortalPropertiesResponseOutput)
}

func (o DevToolPortalOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DevToolPortal) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DevToolPortalOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DevToolPortal) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DevToolPortalOutput{})
}
