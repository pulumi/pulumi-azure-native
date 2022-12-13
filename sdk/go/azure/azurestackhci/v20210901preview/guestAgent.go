


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GuestAgent struct {
	pulumi.CustomResourceState

	Credentials        GuestCredentialResponsePtrOutput        `pulumi:"credentials"`
	HttpProxyConfig    HttpProxyConfigurationResponsePtrOutput `pulumi:"httpProxyConfig"`
	Name               pulumi.StringOutput                     `pulumi:"name"`
	ProvisioningAction pulumi.StringPtrOutput                  `pulumi:"provisioningAction"`
	ProvisioningState  pulumi.StringOutput                     `pulumi:"provisioningState"`
	Status             pulumi.StringOutput                     `pulumi:"status"`
	SystemData         SystemDataResponseOutput                `pulumi:"systemData"`
	Type               pulumi.StringOutput                     `pulumi:"type"`
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
	var resource GuestAgent
	err := ctx.RegisterResource("azure-native:azurestackhci/v20210901preview:GuestAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGuestAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestAgentState, opts ...pulumi.ResourceOption) (*GuestAgent, error) {
	var resource GuestAgent
	err := ctx.ReadResource("azure-native:azurestackhci/v20210901preview:GuestAgent", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**GuestAgent)(nil)).Elem()
}

func (i *GuestAgent) ToGuestAgentOutput() GuestAgentOutput {
	return i.ToGuestAgentOutputWithContext(context.Background())
}

func (i *GuestAgent) ToGuestAgentOutputWithContext(ctx context.Context) GuestAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestAgentOutput)
}

type GuestAgentOutput struct{ *pulumi.OutputState }

func (GuestAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestAgent)(nil)).Elem()
}

func (o GuestAgentOutput) ToGuestAgentOutput() GuestAgentOutput {
	return o
}

func (o GuestAgentOutput) ToGuestAgentOutputWithContext(ctx context.Context) GuestAgentOutput {
	return o
}

func (o GuestAgentOutput) Credentials() GuestCredentialResponsePtrOutput {
	return o.ApplyT(func(v *GuestAgent) GuestCredentialResponsePtrOutput { return v.Credentials }).(GuestCredentialResponsePtrOutput)
}

func (o GuestAgentOutput) HttpProxyConfig() HttpProxyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *GuestAgent) HttpProxyConfigurationResponsePtrOutput { return v.HttpProxyConfig }).(HttpProxyConfigurationResponsePtrOutput)
}

func (o GuestAgentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GuestAgentOutput) ProvisioningAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringPtrOutput { return v.ProvisioningAction }).(pulumi.StringPtrOutput)
}

func (o GuestAgentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GuestAgentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o GuestAgentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GuestAgent) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GuestAgentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GuestAgentOutput{})
}
