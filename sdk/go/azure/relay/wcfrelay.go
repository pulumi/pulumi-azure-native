


package relay

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WCFRelay struct {
	pulumi.CustomResourceState

	CreatedAt                   pulumi.StringOutput    `pulumi:"createdAt"`
	IsDynamic                   pulumi.BoolOutput      `pulumi:"isDynamic"`
	ListenerCount               pulumi.IntOutput       `pulumi:"listenerCount"`
	Name                        pulumi.StringOutput    `pulumi:"name"`
	RelayType                   pulumi.StringPtrOutput `pulumi:"relayType"`
	RequiresClientAuthorization pulumi.BoolPtrOutput   `pulumi:"requiresClientAuthorization"`
	RequiresTransportSecurity   pulumi.BoolPtrOutput   `pulumi:"requiresTransportSecurity"`
	Type                        pulumi.StringOutput    `pulumi:"type"`
	UpdatedAt                   pulumi.StringOutput    `pulumi:"updatedAt"`
	UserMetadata                pulumi.StringPtrOutput `pulumi:"userMetadata"`
}


func NewWCFRelay(ctx *pulumi.Context,
	name string, args *WCFRelayArgs, opts ...pulumi.ResourceOption) (*WCFRelay, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:relay/v20160701:WCFRelay"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20170401:WCFRelay"),
		},
	})
	opts = append(opts, aliases)
	var resource WCFRelay
	err := ctx.RegisterResource("azure-native:relay:WCFRelay", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWCFRelay(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WCFRelayState, opts ...pulumi.ResourceOption) (*WCFRelay, error) {
	var resource WCFRelay
	err := ctx.ReadResource("azure-native:relay:WCFRelay", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type wcfrelayState struct {
}

type WCFRelayState struct {
}

func (WCFRelayState) ElementType() reflect.Type {
	return reflect.TypeOf((*wcfrelayState)(nil)).Elem()
}

type wcfrelayArgs struct {
	NamespaceName               string     `pulumi:"namespaceName"`
	RelayName                   *string    `pulumi:"relayName"`
	RelayType                   *Relaytype `pulumi:"relayType"`
	RequiresClientAuthorization *bool      `pulumi:"requiresClientAuthorization"`
	RequiresTransportSecurity   *bool      `pulumi:"requiresTransportSecurity"`
	ResourceGroupName           string     `pulumi:"resourceGroupName"`
	UserMetadata                *string    `pulumi:"userMetadata"`
}


type WCFRelayArgs struct {
	NamespaceName               pulumi.StringInput
	RelayName                   pulumi.StringPtrInput
	RelayType                   RelaytypePtrInput
	RequiresClientAuthorization pulumi.BoolPtrInput
	RequiresTransportSecurity   pulumi.BoolPtrInput
	ResourceGroupName           pulumi.StringInput
	UserMetadata                pulumi.StringPtrInput
}

func (WCFRelayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wcfrelayArgs)(nil)).Elem()
}

type WCFRelayInput interface {
	pulumi.Input

	ToWCFRelayOutput() WCFRelayOutput
	ToWCFRelayOutputWithContext(ctx context.Context) WCFRelayOutput
}

func (*WCFRelay) ElementType() reflect.Type {
	return reflect.TypeOf((**WCFRelay)(nil)).Elem()
}

func (i *WCFRelay) ToWCFRelayOutput() WCFRelayOutput {
	return i.ToWCFRelayOutputWithContext(context.Background())
}

func (i *WCFRelay) ToWCFRelayOutputWithContext(ctx context.Context) WCFRelayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WCFRelayOutput)
}

type WCFRelayOutput struct{ *pulumi.OutputState }

func (WCFRelayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WCFRelay)(nil)).Elem()
}

func (o WCFRelayOutput) ToWCFRelayOutput() WCFRelayOutput {
	return o
}

func (o WCFRelayOutput) ToWCFRelayOutputWithContext(ctx context.Context) WCFRelayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WCFRelayOutput{})
}
