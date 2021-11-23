


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GuestAgent struct {
	pulumi.CustomResourceState

	Credentials        GuestCredentialResponsePtrOutput        `pulumi:"credentials"`
	CustomResourceName pulumi.StringOutput                     `pulumi:"customResourceName"`
	HttpProxyConfig    HttpProxyConfigurationResponsePtrOutput `pulumi:"httpProxyConfig"`
	Name               pulumi.StringOutput                     `pulumi:"name"`
	ProvisioningAction pulumi.StringPtrOutput                  `pulumi:"provisioningAction"`
	ProvisioningState  pulumi.StringOutput                     `pulumi:"provisioningState"`
	Status             pulumi.StringOutput                     `pulumi:"status"`
	Statuses           ResourceStatusResponseArrayOutput       `pulumi:"statuses"`
	SystemData         SystemDataResponseOutput                `pulumi:"systemData"`
	Type               pulumi.StringOutput                     `pulumi:"type"`
	Uuid               pulumi.StringOutput                     `pulumi:"uuid"`
}


func NewGuestAgent(ctx *pulumi.Context,
	name string, args *GuestAgentArgs, opts ...pulumi.ResourceOption) (*GuestAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualMachineName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere:GuestAgent"),
		},
	})
	opts = append(opts, aliases)
	var resource GuestAgent
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere/v20201001preview:GuestAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGuestAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestAgentState, opts ...pulumi.ResourceOption) (*GuestAgent, error) {
	var resource GuestAgent
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere/v20201001preview:GuestAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type guestAgentState struct {
}

type GuestAgentState struct {
}

func (GuestAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestAgentState)(nil)).Elem()
}

type guestAgentArgs struct {
	Credentials        *GuestCredential        `pulumi:"credentials"`
	HttpProxyConfig    *HttpProxyConfiguration `pulumi:"httpProxyConfig"`
	Name               *string                 `pulumi:"name"`
	ProvisioningAction *string                 `pulumi:"provisioningAction"`
	ResourceGroupName  string                  `pulumi:"resourceGroupName"`
	VirtualMachineName string                  `pulumi:"virtualMachineName"`
}


type GuestAgentArgs struct {
	Credentials        GuestCredentialPtrInput
	HttpProxyConfig    HttpProxyConfigurationPtrInput
	Name               pulumi.StringPtrInput
	ProvisioningAction pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	VirtualMachineName pulumi.StringInput
}

func (GuestAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestAgentArgs)(nil)).Elem()
}

type GuestAgentInput interface {
	pulumi.Input

	ToGuestAgentOutput() GuestAgentOutput
	ToGuestAgentOutputWithContext(ctx context.Context) GuestAgentOutput
}

func (*GuestAgent) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestAgent)(nil))
}

func (i *GuestAgent) ToGuestAgentOutput() GuestAgentOutput {
	return i.ToGuestAgentOutputWithContext(context.Background())
}

func (i *GuestAgent) ToGuestAgentOutputWithContext(ctx context.Context) GuestAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestAgentOutput)
}

type GuestAgentOutput struct{ *pulumi.OutputState }

func (GuestAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestAgent)(nil))
}

func (o GuestAgentOutput) ToGuestAgentOutput() GuestAgentOutput {
	return o
}

func (o GuestAgentOutput) ToGuestAgentOutputWithContext(ctx context.Context) GuestAgentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GuestAgentOutput{})
}
