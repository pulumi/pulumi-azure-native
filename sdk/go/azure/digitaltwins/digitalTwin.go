


package digitaltwins

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DigitalTwin struct {
	pulumi.CustomResourceState

	CreatedTime                pulumi.StringOutput                          `pulumi:"createdTime"`
	HostName                   pulumi.StringOutput                          `pulumi:"hostName"`
	Identity                   DigitalTwinsIdentityResponsePtrOutput        `pulumi:"identity"`
	LastUpdatedTime            pulumi.StringOutput                          `pulumi:"lastUpdatedTime"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrOutput                       `pulumi:"publicNetworkAccess"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
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
			Type: pulumi.String("azure-nextgen:digitaltwins:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20200301preview:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-nextgen:digitaltwins/v20200301preview:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20201031:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-nextgen:digitaltwins/v20201031:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20201201:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-nextgen:digitaltwins/v20201201:DigitalTwin"),
		},
	})
	opts = append(opts, aliases)
	var resource DigitalTwin
	err := ctx.RegisterResource("azure-native:digitaltwins:DigitalTwin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDigitalTwin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DigitalTwinState, opts ...pulumi.ResourceOption) (*DigitalTwin, error) {
	var resource DigitalTwin
	err := ctx.ReadResource("azure-native:digitaltwins:DigitalTwin", name, id, state, &resource, opts...)
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
	Identity                   *DigitalTwinsIdentity           `pulumi:"identity"`
	Location                   *string                         `pulumi:"location"`
	PrivateEndpointConnections []PrivateEndpointConnectionType `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess        *string                         `pulumi:"publicNetworkAccess"`
	ResourceGroupName          string                          `pulumi:"resourceGroupName"`
	ResourceName               *string                         `pulumi:"resourceName"`
	Tags                       map[string]string               `pulumi:"tags"`
}


type DigitalTwinArgs struct {
	Identity                   DigitalTwinsIdentityPtrInput
	Location                   pulumi.StringPtrInput
	PrivateEndpointConnections PrivateEndpointConnectionTypeArrayInput
	PublicNetworkAccess        pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	ResourceName               pulumi.StringPtrInput
	Tags                       pulumi.StringMapInput
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
	return reflect.TypeOf((*DigitalTwin)(nil))
}

func (i *DigitalTwin) ToDigitalTwinOutput() DigitalTwinOutput {
	return i.ToDigitalTwinOutputWithContext(context.Background())
}

func (i *DigitalTwin) ToDigitalTwinOutputWithContext(ctx context.Context) DigitalTwinOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinOutput)
}

type DigitalTwinOutput struct{ *pulumi.OutputState }

func (DigitalTwinOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwin)(nil))
}

func (o DigitalTwinOutput) ToDigitalTwinOutput() DigitalTwinOutput {
	return o
}

func (o DigitalTwinOutput) ToDigitalTwinOutputWithContext(ctx context.Context) DigitalTwinOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DigitalTwinOutput{})
}
