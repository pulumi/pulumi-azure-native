


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationLiveView struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                         `pulumi:"name"`
	Properties ApplicationLiveViewPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                    `pulumi:"systemData"`
	Type       pulumi.StringOutput                         `pulumi:"type"`
}


func NewApplicationLiveView(ctx *pulumi.Context,
	name string, args *ApplicationLiveViewArgs, opts ...pulumi.ResourceOption) (*ApplicationLiveView, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource ApplicationLiveView
	err := ctx.RegisterResource("azure-native:appplatform/v20221101preview:ApplicationLiveView", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationLiveView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationLiveViewState, opts ...pulumi.ResourceOption) (*ApplicationLiveView, error) {
	var resource ApplicationLiveView
	err := ctx.ReadResource("azure-native:appplatform/v20221101preview:ApplicationLiveView", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationLiveViewState struct {
}

type ApplicationLiveViewState struct {
}

func (ApplicationLiveViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationLiveViewState)(nil)).Elem()
}

type applicationLiveViewArgs struct {
	ApplicationLiveViewName *string `pulumi:"applicationLiveViewName"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	ServiceName             string  `pulumi:"serviceName"`
}


type ApplicationLiveViewArgs struct {
	ApplicationLiveViewName pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	ServiceName             pulumi.StringInput
}

func (ApplicationLiveViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationLiveViewArgs)(nil)).Elem()
}

type ApplicationLiveViewInput interface {
	pulumi.Input

	ToApplicationLiveViewOutput() ApplicationLiveViewOutput
	ToApplicationLiveViewOutputWithContext(ctx context.Context) ApplicationLiveViewOutput
}

func (*ApplicationLiveView) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLiveView)(nil)).Elem()
}

func (i *ApplicationLiveView) ToApplicationLiveViewOutput() ApplicationLiveViewOutput {
	return i.ToApplicationLiveViewOutputWithContext(context.Background())
}

func (i *ApplicationLiveView) ToApplicationLiveViewOutputWithContext(ctx context.Context) ApplicationLiveViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLiveViewOutput)
}

type ApplicationLiveViewOutput struct{ *pulumi.OutputState }

func (ApplicationLiveViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLiveView)(nil)).Elem()
}

func (o ApplicationLiveViewOutput) ToApplicationLiveViewOutput() ApplicationLiveViewOutput {
	return o
}

func (o ApplicationLiveViewOutput) ToApplicationLiveViewOutputWithContext(ctx context.Context) ApplicationLiveViewOutput {
	return o
}

func (o ApplicationLiveViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationLiveView) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationLiveViewOutput) Properties() ApplicationLiveViewPropertiesResponseOutput {
	return o.ApplyT(func(v *ApplicationLiveView) ApplicationLiveViewPropertiesResponseOutput { return v.Properties }).(ApplicationLiveViewPropertiesResponseOutput)
}

func (o ApplicationLiveViewOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApplicationLiveView) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ApplicationLiveViewOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationLiveView) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationLiveViewOutput{})
}
