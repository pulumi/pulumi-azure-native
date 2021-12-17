


package v20190401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Controller struct {
	pulumi.CustomResourceState

	DataPlaneFqdn                    pulumi.StringOutput    `pulumi:"dataPlaneFqdn"`
	HostSuffix                       pulumi.StringOutput    `pulumi:"hostSuffix"`
	Location                         pulumi.StringOutput    `pulumi:"location"`
	Name                             pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState                pulumi.StringOutput    `pulumi:"provisioningState"`
	Sku                              SkuResponseOutput      `pulumi:"sku"`
	Tags                             pulumi.StringMapOutput `pulumi:"tags"`
	TargetContainerHostApiServerFqdn pulumi.StringOutput    `pulumi:"targetContainerHostApiServerFqdn"`
	TargetContainerHostResourceId    pulumi.StringOutput    `pulumi:"targetContainerHostResourceId"`
	Type                             pulumi.StringOutput    `pulumi:"type"`
}


func NewController(ctx *pulumi.Context,
	name string, args *ControllerArgs, opts ...pulumi.ResourceOption) (*Controller, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.TargetContainerHostCredentialsBase64 == nil {
		return nil, errors.New("invalid value for required argument 'TargetContainerHostCredentialsBase64'")
	}
	if args.TargetContainerHostResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetContainerHostResourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devspaces:Controller"),
		},
	})
	opts = append(opts, aliases)
	var resource Controller
	err := ctx.RegisterResource("azure-native:devspaces/v20190401:Controller", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ControllerState, opts ...pulumi.ResourceOption) (*Controller, error) {
	var resource Controller
	err := ctx.ReadResource("azure-native:devspaces/v20190401:Controller", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type controllerState struct {
}

type ControllerState struct {
}

func (ControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*controllerState)(nil)).Elem()
}

type controllerArgs struct {
	Location                             *string           `pulumi:"location"`
	Name                                 *string           `pulumi:"name"`
	ResourceGroupName                    string            `pulumi:"resourceGroupName"`
	Sku                                  Sku               `pulumi:"sku"`
	Tags                                 map[string]string `pulumi:"tags"`
	TargetContainerHostCredentialsBase64 string            `pulumi:"targetContainerHostCredentialsBase64"`
	TargetContainerHostResourceId        string            `pulumi:"targetContainerHostResourceId"`
}


type ControllerArgs struct {
	Location                             pulumi.StringPtrInput
	Name                                 pulumi.StringPtrInput
	ResourceGroupName                    pulumi.StringInput
	Sku                                  SkuInput
	Tags                                 pulumi.StringMapInput
	TargetContainerHostCredentialsBase64 pulumi.StringInput
	TargetContainerHostResourceId        pulumi.StringInput
}

func (ControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*controllerArgs)(nil)).Elem()
}

type ControllerInput interface {
	pulumi.Input

	ToControllerOutput() ControllerOutput
	ToControllerOutputWithContext(ctx context.Context) ControllerOutput
}

func (*Controller) ElementType() reflect.Type {
	return reflect.TypeOf((**Controller)(nil)).Elem()
}

func (i *Controller) ToControllerOutput() ControllerOutput {
	return i.ToControllerOutputWithContext(context.Background())
}

func (i *Controller) ToControllerOutputWithContext(ctx context.Context) ControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControllerOutput)
}

type ControllerOutput struct{ *pulumi.OutputState }

func (ControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Controller)(nil)).Elem()
}

func (o ControllerOutput) ToControllerOutput() ControllerOutput {
	return o
}

func (o ControllerOutput) ToControllerOutputWithContext(ctx context.Context) ControllerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ControllerOutput{})
}
