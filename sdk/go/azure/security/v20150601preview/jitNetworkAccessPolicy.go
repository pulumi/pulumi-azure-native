


package v20150601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type JitNetworkAccessPolicy struct {
	pulumi.CustomResourceState

	Kind              pulumi.StringPtrOutput                                  `pulumi:"kind"`
	Location          pulumi.StringOutput                                     `pulumi:"location"`
	Name              pulumi.StringOutput                                     `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                                     `pulumi:"provisioningState"`
	Requests          JitNetworkAccessRequestResponseArrayOutput              `pulumi:"requests"`
	Type              pulumi.StringOutput                                     `pulumi:"type"`
	VirtualMachines   JitNetworkAccessPolicyVirtualMachineResponseArrayOutput `pulumi:"virtualMachines"`
}


func NewJitNetworkAccessPolicy(ctx *pulumi.Context,
	name string, args *JitNetworkAccessPolicyArgs, opts ...pulumi.ResourceOption) (*JitNetworkAccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AscLocation == nil {
		return nil, errors.New("invalid value for required argument 'AscLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualMachines == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachines'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:JitNetworkAccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:security/v20200101:JitNetworkAccessPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource JitNetworkAccessPolicy
	err := ctx.RegisterResource("azure-native:security/v20150601preview:JitNetworkAccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJitNetworkAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JitNetworkAccessPolicyState, opts ...pulumi.ResourceOption) (*JitNetworkAccessPolicy, error) {
	var resource JitNetworkAccessPolicy
	err := ctx.ReadResource("azure-native:security/v20150601preview:JitNetworkAccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type jitNetworkAccessPolicyState struct {
}

type JitNetworkAccessPolicyState struct {
}

func (JitNetworkAccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*jitNetworkAccessPolicyState)(nil)).Elem()
}

type jitNetworkAccessPolicyArgs struct {
	AscLocation                string                                 `pulumi:"ascLocation"`
	JitNetworkAccessPolicyName *string                                `pulumi:"jitNetworkAccessPolicyName"`
	Kind                       *string                                `pulumi:"kind"`
	Requests                   []JitNetworkAccessRequest              `pulumi:"requests"`
	ResourceGroupName          string                                 `pulumi:"resourceGroupName"`
	VirtualMachines            []JitNetworkAccessPolicyVirtualMachine `pulumi:"virtualMachines"`
}


type JitNetworkAccessPolicyArgs struct {
	AscLocation                pulumi.StringInput
	JitNetworkAccessPolicyName pulumi.StringPtrInput
	Kind                       pulumi.StringPtrInput
	Requests                   JitNetworkAccessRequestArrayInput
	ResourceGroupName          pulumi.StringInput
	VirtualMachines            JitNetworkAccessPolicyVirtualMachineArrayInput
}

func (JitNetworkAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jitNetworkAccessPolicyArgs)(nil)).Elem()
}

type JitNetworkAccessPolicyInput interface {
	pulumi.Input

	ToJitNetworkAccessPolicyOutput() JitNetworkAccessPolicyOutput
	ToJitNetworkAccessPolicyOutputWithContext(ctx context.Context) JitNetworkAccessPolicyOutput
}

func (*JitNetworkAccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**JitNetworkAccessPolicy)(nil)).Elem()
}

func (i *JitNetworkAccessPolicy) ToJitNetworkAccessPolicyOutput() JitNetworkAccessPolicyOutput {
	return i.ToJitNetworkAccessPolicyOutputWithContext(context.Background())
}

func (i *JitNetworkAccessPolicy) ToJitNetworkAccessPolicyOutputWithContext(ctx context.Context) JitNetworkAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPolicyOutput)
}

type JitNetworkAccessPolicyOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JitNetworkAccessPolicy)(nil)).Elem()
}

func (o JitNetworkAccessPolicyOutput) ToJitNetworkAccessPolicyOutput() JitNetworkAccessPolicyOutput {
	return o
}

func (o JitNetworkAccessPolicyOutput) ToJitNetworkAccessPolicyOutputWithContext(ctx context.Context) JitNetworkAccessPolicyOutput {
	return o
}

func (o JitNetworkAccessPolicyOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o JitNetworkAccessPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o JitNetworkAccessPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o JitNetworkAccessPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o JitNetworkAccessPolicyOutput) Requests() JitNetworkAccessRequestResponseArrayOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) JitNetworkAccessRequestResponseArrayOutput { return v.Requests }).(JitNetworkAccessRequestResponseArrayOutput)
}

func (o JitNetworkAccessPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o JitNetworkAccessPolicyOutput) VirtualMachines() JitNetworkAccessPolicyVirtualMachineResponseArrayOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) JitNetworkAccessPolicyVirtualMachineResponseArrayOutput {
		return v.VirtualMachines
	}).(JitNetworkAccessPolicyVirtualMachineResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(JitNetworkAccessPolicyOutput{})
}
