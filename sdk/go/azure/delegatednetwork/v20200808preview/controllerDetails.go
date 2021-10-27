


package v20200808preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ControllerDetails struct {
	pulumi.CustomResourceState

	DncAppId          pulumi.StringOutput    `pulumi:"dncAppId"`
	DncEndpoint       pulumi.StringOutput    `pulumi:"dncEndpoint"`
	DncTenantId       pulumi.StringOutput    `pulumi:"dncTenantId"`
	Location          pulumi.StringPtrOutput `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	ResourceGuid      pulumi.StringOutput    `pulumi:"resourceGuid"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewControllerDetails(ctx *pulumi.Context,
	name string, args *ControllerDetailsArgs, opts ...pulumi.ResourceOption) (*ControllerDetails, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:delegatednetwork/v20200808preview:ControllerDetails"),
		},
		{
			Type: pulumi.String("azure-native:delegatednetwork:ControllerDetails"),
		},
		{
			Type: pulumi.String("azure-nextgen:delegatednetwork:ControllerDetails"),
		},
		{
			Type: pulumi.String("azure-native:delegatednetwork/v20210315:ControllerDetails"),
		},
		{
			Type: pulumi.String("azure-nextgen:delegatednetwork/v20210315:ControllerDetails"),
		},
	})
	opts = append(opts, aliases)
	var resource ControllerDetails
	err := ctx.RegisterResource("azure-native:delegatednetwork/v20200808preview:ControllerDetails", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetControllerDetails(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ControllerDetailsState, opts ...pulumi.ResourceOption) (*ControllerDetails, error) {
	var resource ControllerDetails
	err := ctx.ReadResource("azure-native:delegatednetwork/v20200808preview:ControllerDetails", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type controllerDetailsState struct {
}

type ControllerDetailsState struct {
}

func (ControllerDetailsState) ElementType() reflect.Type {
	return reflect.TypeOf((*controllerDetailsState)(nil)).Elem()
}

type controllerDetailsArgs struct {
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	Tags              map[string]string `pulumi:"tags"`
}


type ControllerDetailsArgs struct {
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (ControllerDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*controllerDetailsArgs)(nil)).Elem()
}

type ControllerDetailsInput interface {
	pulumi.Input

	ToControllerDetailsOutput() ControllerDetailsOutput
	ToControllerDetailsOutputWithContext(ctx context.Context) ControllerDetailsOutput
}

func (*ControllerDetails) ElementType() reflect.Type {
	return reflect.TypeOf((*ControllerDetails)(nil))
}

func (i *ControllerDetails) ToControllerDetailsOutput() ControllerDetailsOutput {
	return i.ToControllerDetailsOutputWithContext(context.Background())
}

func (i *ControllerDetails) ToControllerDetailsOutputWithContext(ctx context.Context) ControllerDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControllerDetailsOutput)
}

type ControllerDetailsOutput struct{ *pulumi.OutputState }

func (ControllerDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ControllerDetails)(nil))
}

func (o ControllerDetailsOutput) ToControllerDetailsOutput() ControllerDetailsOutput {
	return o
}

func (o ControllerDetailsOutput) ToControllerDetailsOutputWithContext(ctx context.Context) ControllerDetailsOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ControllerDetailsInput)(nil)).Elem(), &ControllerDetails{})
	pulumi.RegisterOutputType(ControllerDetailsOutput{})
}
