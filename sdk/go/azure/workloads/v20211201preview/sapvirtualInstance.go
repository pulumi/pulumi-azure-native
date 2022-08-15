


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SAPVirtualInstance struct {
	pulumi.CustomResourceState

	Configuration                     pulumi.AnyOutput                             `pulumi:"configuration"`
	Environment                       pulumi.StringOutput                          `pulumi:"environment"`
	Errors                            SAPVirtualInstanceErrorResponseOutput        `pulumi:"errors"`
	Health                            pulumi.StringOutput                          `pulumi:"health"`
	Identity                          UserAssignedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Location                          pulumi.StringOutput                          `pulumi:"location"`
	ManagedResourceGroupConfiguration ManagedRGConfigurationResponsePtrOutput      `pulumi:"managedResourceGroupConfiguration"`
	Name                              pulumi.StringOutput                          `pulumi:"name"`
	ProvisioningState                 pulumi.StringOutput                          `pulumi:"provisioningState"`
	SapProduct                        pulumi.StringOutput                          `pulumi:"sapProduct"`
	State                             pulumi.StringOutput                          `pulumi:"state"`
	Status                            pulumi.StringOutput                          `pulumi:"status"`
	SystemData                        SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                              pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                              pulumi.StringOutput                          `pulumi:"type"`
}


func NewSAPVirtualInstance(ctx *pulumi.Context,
	name string, args *SAPVirtualInstanceArgs, opts ...pulumi.ResourceOption) (*SAPVirtualInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SapProduct == nil {
		return nil, errors.New("invalid value for required argument 'SapProduct'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:workloads:SAPVirtualInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource SAPVirtualInstance
	err := ctx.RegisterResource("azure-native:workloads/v20211201preview:SAPVirtualInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSAPVirtualInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SAPVirtualInstanceState, opts ...pulumi.ResourceOption) (*SAPVirtualInstance, error) {
	var resource SAPVirtualInstance
	err := ctx.ReadResource("azure-native:workloads/v20211201preview:SAPVirtualInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sapvirtualInstanceState struct {
}

type SAPVirtualInstanceState struct {
}

func (SAPVirtualInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sapvirtualInstanceState)(nil)).Elem()
}

type sapvirtualInstanceArgs struct {
	Configuration                     interface{}                  `pulumi:"configuration"`
	Environment                       string                       `pulumi:"environment"`
	Identity                          *UserAssignedServiceIdentity `pulumi:"identity"`
	Location                          *string                      `pulumi:"location"`
	ManagedResourceGroupConfiguration *ManagedRGConfiguration      `pulumi:"managedResourceGroupConfiguration"`
	ResourceGroupName                 string                       `pulumi:"resourceGroupName"`
	SapProduct                        string                       `pulumi:"sapProduct"`
	SapVirtualInstanceName            *string                      `pulumi:"sapVirtualInstanceName"`
	Tags                              map[string]string            `pulumi:"tags"`
}


type SAPVirtualInstanceArgs struct {
	Configuration                     pulumi.Input
	Environment                       pulumi.StringInput
	Identity                          UserAssignedServiceIdentityPtrInput
	Location                          pulumi.StringPtrInput
	ManagedResourceGroupConfiguration ManagedRGConfigurationPtrInput
	ResourceGroupName                 pulumi.StringInput
	SapProduct                        pulumi.StringInput
	SapVirtualInstanceName            pulumi.StringPtrInput
	Tags                              pulumi.StringMapInput
}

func (SAPVirtualInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sapvirtualInstanceArgs)(nil)).Elem()
}

type SAPVirtualInstanceInput interface {
	pulumi.Input

	ToSAPVirtualInstanceOutput() SAPVirtualInstanceOutput
	ToSAPVirtualInstanceOutputWithContext(ctx context.Context) SAPVirtualInstanceOutput
}

func (*SAPVirtualInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPVirtualInstance)(nil)).Elem()
}

func (i *SAPVirtualInstance) ToSAPVirtualInstanceOutput() SAPVirtualInstanceOutput {
	return i.ToSAPVirtualInstanceOutputWithContext(context.Background())
}

func (i *SAPVirtualInstance) ToSAPVirtualInstanceOutputWithContext(ctx context.Context) SAPVirtualInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SAPVirtualInstanceOutput)
}

type SAPVirtualInstanceOutput struct{ *pulumi.OutputState }

func (SAPVirtualInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPVirtualInstance)(nil)).Elem()
}

func (o SAPVirtualInstanceOutput) ToSAPVirtualInstanceOutput() SAPVirtualInstanceOutput {
	return o
}

func (o SAPVirtualInstanceOutput) ToSAPVirtualInstanceOutputWithContext(ctx context.Context) SAPVirtualInstanceOutput {
	return o
}

func (o SAPVirtualInstanceOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.AnyOutput { return v.Configuration }).(pulumi.AnyOutput)
}

func (o SAPVirtualInstanceOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

func (o SAPVirtualInstanceOutput) Errors() SAPVirtualInstanceErrorResponseOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) SAPVirtualInstanceErrorResponseOutput { return v.Errors }).(SAPVirtualInstanceErrorResponseOutput)
}

func (o SAPVirtualInstanceOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.Health }).(pulumi.StringOutput)
}

func (o SAPVirtualInstanceOutput) Identity() UserAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) UserAssignedServiceIdentityResponsePtrOutput { return v.Identity }).(UserAssignedServiceIdentityResponsePtrOutput)
}

func (o SAPVirtualInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SAPVirtualInstanceOutput) ManagedResourceGroupConfiguration() ManagedRGConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) ManagedRGConfigurationResponsePtrOutput {
		return v.ManagedResourceGroupConfiguration
	}).(ManagedRGConfigurationResponsePtrOutput)
}

func (o SAPVirtualInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SAPVirtualInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SAPVirtualInstanceOutput) SapProduct() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.SapProduct }).(pulumi.StringOutput)
}

func (o SAPVirtualInstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o SAPVirtualInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SAPVirtualInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SAPVirtualInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SAPVirtualInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SAPVirtualInstanceOutput{})
}
