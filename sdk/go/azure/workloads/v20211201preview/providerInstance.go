


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProviderInstance struct {
	pulumi.CustomResourceState

	Errors            ProviderInstancePropertiesResponseErrorsOutput `pulumi:"errors"`
	Identity          UserAssignedServiceIdentityResponsePtrOutput   `pulumi:"identity"`
	Name              pulumi.StringOutput                            `pulumi:"name"`
	ProviderSettings  pulumi.AnyOutput                               `pulumi:"providerSettings"`
	ProvisioningState pulumi.StringOutput                            `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput                       `pulumi:"systemData"`
	Type              pulumi.StringOutput                            `pulumi:"type"`
}


func NewProviderInstance(ctx *pulumi.Context,
	name string, args *ProviderInstanceArgs, opts ...pulumi.ResourceOption) (*ProviderInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MonitorName == nil {
		return nil, errors.New("invalid value for required argument 'MonitorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:workloads:ProviderInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource ProviderInstance
	err := ctx.RegisterResource("azure-native:workloads/v20211201preview:ProviderInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProviderInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderInstanceState, opts ...pulumi.ResourceOption) (*ProviderInstance, error) {
	var resource ProviderInstance
	err := ctx.ReadResource("azure-native:workloads/v20211201preview:ProviderInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type providerInstanceState struct {
}

type ProviderInstanceState struct {
}

func (ProviderInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerInstanceState)(nil)).Elem()
}

type providerInstanceArgs struct {
	Identity             *UserAssignedServiceIdentity `pulumi:"identity"`
	MonitorName          string                       `pulumi:"monitorName"`
	ProviderInstanceName *string                      `pulumi:"providerInstanceName"`
	ProviderSettings     interface{}                  `pulumi:"providerSettings"`
	ResourceGroupName    string                       `pulumi:"resourceGroupName"`
}


type ProviderInstanceArgs struct {
	Identity             UserAssignedServiceIdentityPtrInput
	MonitorName          pulumi.StringInput
	ProviderInstanceName pulumi.StringPtrInput
	ProviderSettings     pulumi.Input
	ResourceGroupName    pulumi.StringInput
}

func (ProviderInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerInstanceArgs)(nil)).Elem()
}

type ProviderInstanceInput interface {
	pulumi.Input

	ToProviderInstanceOutput() ProviderInstanceOutput
	ToProviderInstanceOutputWithContext(ctx context.Context) ProviderInstanceOutput
}

func (*ProviderInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderInstance)(nil)).Elem()
}

func (i *ProviderInstance) ToProviderInstanceOutput() ProviderInstanceOutput {
	return i.ToProviderInstanceOutputWithContext(context.Background())
}

func (i *ProviderInstance) ToProviderInstanceOutputWithContext(ctx context.Context) ProviderInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderInstanceOutput)
}

type ProviderInstanceOutput struct{ *pulumi.OutputState }

func (ProviderInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderInstance)(nil)).Elem()
}

func (o ProviderInstanceOutput) ToProviderInstanceOutput() ProviderInstanceOutput {
	return o
}

func (o ProviderInstanceOutput) ToProviderInstanceOutputWithContext(ctx context.Context) ProviderInstanceOutput {
	return o
}

func (o ProviderInstanceOutput) Errors() ProviderInstancePropertiesResponseErrorsOutput {
	return o.ApplyT(func(v *ProviderInstance) ProviderInstancePropertiesResponseErrorsOutput { return v.Errors }).(ProviderInstancePropertiesResponseErrorsOutput)
}

func (o ProviderInstanceOutput) Identity() UserAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ProviderInstance) UserAssignedServiceIdentityResponsePtrOutput { return v.Identity }).(UserAssignedServiceIdentityResponsePtrOutput)
}

func (o ProviderInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProviderInstanceOutput) ProviderSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *ProviderInstance) pulumi.AnyOutput { return v.ProviderSettings }).(pulumi.AnyOutput)
}

func (o ProviderInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ProviderInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ProviderInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ProviderInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProviderInstanceOutput{})
}
