


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Sim struct {
	pulumi.CustomResourceState

	CreatedAt                             pulumi.StringPtrOutput                   `pulumi:"createdAt"`
	CreatedBy                             pulumi.StringPtrOutput                   `pulumi:"createdBy"`
	CreatedByType                         pulumi.StringPtrOutput                   `pulumi:"createdByType"`
	DeviceType                            pulumi.StringPtrOutput                   `pulumi:"deviceType"`
	IntegratedCircuitCardIdentifier       pulumi.StringPtrOutput                   `pulumi:"integratedCircuitCardIdentifier"`
	InternationalMobileSubscriberIdentity pulumi.StringOutput                      `pulumi:"internationalMobileSubscriberIdentity"`
	LastModifiedAt                        pulumi.StringPtrOutput                   `pulumi:"lastModifiedAt"`
	LastModifiedBy                        pulumi.StringPtrOutput                   `pulumi:"lastModifiedBy"`
	LastModifiedByType                    pulumi.StringPtrOutput                   `pulumi:"lastModifiedByType"`
	Name                                  pulumi.StringOutput                      `pulumi:"name"`
	ProvisioningState                     pulumi.StringOutput                      `pulumi:"provisioningState"`
	SimPolicy                             SimPolicyResourceIdResponsePtrOutput     `pulumi:"simPolicy"`
	SimState                              pulumi.StringOutput                      `pulumi:"simState"`
	StaticIpConfiguration                 SimStaticIpPropertiesResponseArrayOutput `pulumi:"staticIpConfiguration"`
	SystemData                            SystemDataResponseOutput                 `pulumi:"systemData"`
	Type                                  pulumi.StringOutput                      `pulumi:"type"`
}


func NewSim(ctx *pulumi.Context,
	name string, args *SimArgs, opts ...pulumi.ResourceOption) (*Sim, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InternationalMobileSubscriberIdentity == nil {
		return nil, errors.New("invalid value for required argument 'InternationalMobileSubscriberIdentity'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SimGroupName == nil {
		return nil, errors.New("invalid value for required argument 'SimGroupName'")
	}
	var resource Sim
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20220401preview:Sim", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSim(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SimState, opts ...pulumi.ResourceOption) (*Sim, error) {
	var resource Sim
	err := ctx.ReadResource("azure-native:mobilenetwork/v20220401preview:Sim", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type simState struct {
}

type SimState struct {
}

func (SimState) ElementType() reflect.Type {
	return reflect.TypeOf((*simState)(nil)).Elem()
}

type simArgs struct {
	AuthenticationKey                     *string                 `pulumi:"authenticationKey"`
	CreatedAt                             *string                 `pulumi:"createdAt"`
	CreatedBy                             *string                 `pulumi:"createdBy"`
	CreatedByType                         *string                 `pulumi:"createdByType"`
	DeviceType                            *string                 `pulumi:"deviceType"`
	IntegratedCircuitCardIdentifier       *string                 `pulumi:"integratedCircuitCardIdentifier"`
	InternationalMobileSubscriberIdentity string                  `pulumi:"internationalMobileSubscriberIdentity"`
	LastModifiedAt                        *string                 `pulumi:"lastModifiedAt"`
	LastModifiedBy                        *string                 `pulumi:"lastModifiedBy"`
	LastModifiedByType                    *string                 `pulumi:"lastModifiedByType"`
	OperatorKeyCode                       *string                 `pulumi:"operatorKeyCode"`
	ResourceGroupName                     string                  `pulumi:"resourceGroupName"`
	SimGroupName                          string                  `pulumi:"simGroupName"`
	SimName                               *string                 `pulumi:"simName"`
	SimPolicy                             *SimPolicyResourceId    `pulumi:"simPolicy"`
	StaticIpConfiguration                 []SimStaticIpProperties `pulumi:"staticIpConfiguration"`
}


type SimArgs struct {
	AuthenticationKey                     pulumi.StringPtrInput
	CreatedAt                             pulumi.StringPtrInput
	CreatedBy                             pulumi.StringPtrInput
	CreatedByType                         pulumi.StringPtrInput
	DeviceType                            pulumi.StringPtrInput
	IntegratedCircuitCardIdentifier       pulumi.StringPtrInput
	InternationalMobileSubscriberIdentity pulumi.StringInput
	LastModifiedAt                        pulumi.StringPtrInput
	LastModifiedBy                        pulumi.StringPtrInput
	LastModifiedByType                    pulumi.StringPtrInput
	OperatorKeyCode                       pulumi.StringPtrInput
	ResourceGroupName                     pulumi.StringInput
	SimGroupName                          pulumi.StringInput
	SimName                               pulumi.StringPtrInput
	SimPolicy                             SimPolicyResourceIdPtrInput
	StaticIpConfiguration                 SimStaticIpPropertiesArrayInput
}

func (SimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*simArgs)(nil)).Elem()
}

type SimInput interface {
	pulumi.Input

	ToSimOutput() SimOutput
	ToSimOutputWithContext(ctx context.Context) SimOutput
}

func (*Sim) ElementType() reflect.Type {
	return reflect.TypeOf((**Sim)(nil)).Elem()
}

func (i *Sim) ToSimOutput() SimOutput {
	return i.ToSimOutputWithContext(context.Background())
}

func (i *Sim) ToSimOutputWithContext(ctx context.Context) SimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimOutput)
}

type SimOutput struct{ *pulumi.OutputState }

func (SimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sim)(nil)).Elem()
}

func (o SimOutput) ToSimOutput() SimOutput {
	return o
}

func (o SimOutput) ToSimOutputWithContext(ctx context.Context) SimOutput {
	return o
}

func (o SimOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sim) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SimOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sim) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SimOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sim) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SimOutput) DeviceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sim) pulumi.StringPtrOutput { return v.DeviceType }).(pulumi.StringPtrOutput)
}

func (o SimOutput) IntegratedCircuitCardIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sim) pulumi.StringPtrOutput { return v.IntegratedCircuitCardIdentifier }).(pulumi.StringPtrOutput)
}

func (o SimOutput) InternationalMobileSubscriberIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v *Sim) pulumi.StringOutput { return v.InternationalMobileSubscriberIdentity }).(pulumi.StringOutput)
}

func (o SimOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sim) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SimOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sim) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SimOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sim) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o SimOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Sim) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SimOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Sim) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SimOutput) SimPolicy() SimPolicyResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *Sim) SimPolicyResourceIdResponsePtrOutput { return v.SimPolicy }).(SimPolicyResourceIdResponsePtrOutput)
}

func (o SimOutput) SimState() pulumi.StringOutput {
	return o.ApplyT(func(v *Sim) pulumi.StringOutput { return v.SimState }).(pulumi.StringOutput)
}

func (o SimOutput) StaticIpConfiguration() SimStaticIpPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *Sim) SimStaticIpPropertiesResponseArrayOutput { return v.StaticIpConfiguration }).(SimStaticIpPropertiesResponseArrayOutput)
}

func (o SimOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Sim) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SimOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Sim) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SimOutput{})
}
