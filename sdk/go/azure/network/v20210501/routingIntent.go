


package v20210501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoutingIntent struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput              `pulumi:"etag"`
	Name              pulumi.StringPtrOutput           `pulumi:"name"`
	ProvisioningState pulumi.StringOutput              `pulumi:"provisioningState"`
	RoutingPolicies   RoutingPolicyResponseArrayOutput `pulumi:"routingPolicies"`
	Type              pulumi.StringOutput              `pulumi:"type"`
}


func NewRoutingIntent(ctx *pulumi.Context,
	name string, args *RoutingIntentArgs, opts ...pulumi.ResourceOption) (*RoutingIntent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualHubName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:RoutingIntent"),
		},
	})
	opts = append(opts, aliases)
	var resource RoutingIntent
	err := ctx.RegisterResource("azure-native:network/v20210501:RoutingIntent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRoutingIntent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoutingIntentState, opts ...pulumi.ResourceOption) (*RoutingIntent, error) {
	var resource RoutingIntent
	err := ctx.ReadResource("azure-native:network/v20210501:RoutingIntent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type routingIntentState struct {
}

type RoutingIntentState struct {
}

func (RoutingIntentState) ElementType() reflect.Type {
	return reflect.TypeOf((*routingIntentState)(nil)).Elem()
}

type routingIntentArgs struct {
	Id                *string         `pulumi:"id"`
	Name              *string         `pulumi:"name"`
	ResourceGroupName string          `pulumi:"resourceGroupName"`
	RoutingIntentName *string         `pulumi:"routingIntentName"`
	RoutingPolicies   []RoutingPolicy `pulumi:"routingPolicies"`
	VirtualHubName    string          `pulumi:"virtualHubName"`
}


type RoutingIntentArgs struct {
	Id                pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RoutingIntentName pulumi.StringPtrInput
	RoutingPolicies   RoutingPolicyArrayInput
	VirtualHubName    pulumi.StringInput
}

func (RoutingIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routingIntentArgs)(nil)).Elem()
}

type RoutingIntentInput interface {
	pulumi.Input

	ToRoutingIntentOutput() RoutingIntentOutput
	ToRoutingIntentOutputWithContext(ctx context.Context) RoutingIntentOutput
}

func (*RoutingIntent) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingIntent)(nil))
}

func (i *RoutingIntent) ToRoutingIntentOutput() RoutingIntentOutput {
	return i.ToRoutingIntentOutputWithContext(context.Background())
}

func (i *RoutingIntent) ToRoutingIntentOutputWithContext(ctx context.Context) RoutingIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingIntentOutput)
}

type RoutingIntentOutput struct{ *pulumi.OutputState }

func (RoutingIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingIntent)(nil))
}

func (o RoutingIntentOutput) ToRoutingIntentOutput() RoutingIntentOutput {
	return o
}

func (o RoutingIntentOutput) ToRoutingIntentOutputWithContext(ctx context.Context) RoutingIntentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RoutingIntentOutput{})
}
