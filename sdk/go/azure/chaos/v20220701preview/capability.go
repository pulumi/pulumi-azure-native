


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Capability struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                `pulumi:"name"`
	Properties CapabilityPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput           `pulumi:"systemData"`
	Type       pulumi.StringOutput                `pulumi:"type"`
}


func NewCapability(ctx *pulumi.Context,
	name string, args *CapabilityArgs, opts ...pulumi.ResourceOption) (*Capability, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ParentProviderNamespace'")
	}
	if args.ParentResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourceName'")
	}
	if args.ParentResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourceType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TargetName == nil {
		return nil, errors.New("invalid value for required argument 'TargetName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:chaos:Capability"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20210915preview:Capability"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20221001preview:Capability"),
		},
	})
	opts = append(opts, aliases)
	var resource Capability
	err := ctx.RegisterResource("azure-native:chaos/v20220701preview:Capability", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCapability(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapabilityState, opts ...pulumi.ResourceOption) (*Capability, error) {
	var resource Capability
	err := ctx.ReadResource("azure-native:chaos/v20220701preview:Capability", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type capabilityState struct {
}

type CapabilityState struct {
}

func (CapabilityState) ElementType() reflect.Type {
	return reflect.TypeOf((*capabilityState)(nil)).Elem()
}

type capabilityArgs struct {
	CapabilityName          *string `pulumi:"capabilityName"`
	ParentProviderNamespace string  `pulumi:"parentProviderNamespace"`
	ParentResourceName      string  `pulumi:"parentResourceName"`
	ParentResourceType      string  `pulumi:"parentResourceType"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	TargetName              string  `pulumi:"targetName"`
}


type CapabilityArgs struct {
	CapabilityName          pulumi.StringPtrInput
	ParentProviderNamespace pulumi.StringInput
	ParentResourceName      pulumi.StringInput
	ParentResourceType      pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
	TargetName              pulumi.StringInput
}

func (CapabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capabilityArgs)(nil)).Elem()
}

type CapabilityInput interface {
	pulumi.Input

	ToCapabilityOutput() CapabilityOutput
	ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput
}

func (*Capability) ElementType() reflect.Type {
	return reflect.TypeOf((**Capability)(nil)).Elem()
}

func (i *Capability) ToCapabilityOutput() CapabilityOutput {
	return i.ToCapabilityOutputWithContext(context.Background())
}

func (i *Capability) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityOutput)
}

type CapabilityOutput struct{ *pulumi.OutputState }

func (CapabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Capability)(nil)).Elem()
}

func (o CapabilityOutput) ToCapabilityOutput() CapabilityOutput {
	return o
}

func (o CapabilityOutput) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return o
}

func (o CapabilityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Capability) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CapabilityOutput) Properties() CapabilityPropertiesResponseOutput {
	return o.ApplyT(func(v *Capability) CapabilityPropertiesResponseOutput { return v.Properties }).(CapabilityPropertiesResponseOutput)
}

func (o CapabilityOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Capability) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o CapabilityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Capability) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CapabilityOutput{})
}
