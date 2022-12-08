


package v20221101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SimPolicy struct {
	pulumi.CustomResourceState

	DefaultSlice          SliceResourceIdResponseOutput         `pulumi:"defaultSlice"`
	Location              pulumi.StringOutput                   `pulumi:"location"`
	Name                  pulumi.StringOutput                   `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput                   `pulumi:"provisioningState"`
	RegistrationTimer     pulumi.IntPtrOutput                   `pulumi:"registrationTimer"`
	RfspIndex             pulumi.IntPtrOutput                   `pulumi:"rfspIndex"`
	SiteProvisioningState pulumi.StringMapOutput                `pulumi:"siteProvisioningState"`
	SliceConfigurations   SliceConfigurationResponseArrayOutput `pulumi:"sliceConfigurations"`
	SystemData            SystemDataResponseOutput              `pulumi:"systemData"`
	Tags                  pulumi.StringMapOutput                `pulumi:"tags"`
	Type                  pulumi.StringOutput                   `pulumi:"type"`
	UeAmbr                AmbrResponseOutput                    `pulumi:"ueAmbr"`
}


func NewSimPolicy(ctx *pulumi.Context,
	name string, args *SimPolicyArgs, opts ...pulumi.ResourceOption) (*SimPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultSlice == nil {
		return nil, errors.New("invalid value for required argument 'DefaultSlice'")
	}
	if args.MobileNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'MobileNetworkName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SliceConfigurations == nil {
		return nil, errors.New("invalid value for required argument 'SliceConfigurations'")
	}
	if args.UeAmbr == nil {
		return nil, errors.New("invalid value for required argument 'UeAmbr'")
	}
	if isZero(args.RegistrationTimer) {
		args.RegistrationTimer = pulumi.IntPtr(3240)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:SimPolicy"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:SimPolicy"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:SimPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource SimPolicy
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20221101:SimPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSimPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SimPolicyState, opts ...pulumi.ResourceOption) (*SimPolicy, error) {
	var resource SimPolicy
	err := ctx.ReadResource("azure-native:mobilenetwork/v20221101:SimPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type simPolicyState struct {
}

type SimPolicyState struct {
}

func (SimPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*simPolicyState)(nil)).Elem()
}

type simPolicyArgs struct {
	DefaultSlice        SliceResourceId      `pulumi:"defaultSlice"`
	Location            *string              `pulumi:"location"`
	MobileNetworkName   string               `pulumi:"mobileNetworkName"`
	RegistrationTimer   *int                 `pulumi:"registrationTimer"`
	ResourceGroupName   string               `pulumi:"resourceGroupName"`
	RfspIndex           *int                 `pulumi:"rfspIndex"`
	SimPolicyName       *string              `pulumi:"simPolicyName"`
	SliceConfigurations []SliceConfiguration `pulumi:"sliceConfigurations"`
	Tags                map[string]string    `pulumi:"tags"`
	UeAmbr              Ambr                 `pulumi:"ueAmbr"`
}


type SimPolicyArgs struct {
	DefaultSlice        SliceResourceIdInput
	Location            pulumi.StringPtrInput
	MobileNetworkName   pulumi.StringInput
	RegistrationTimer   pulumi.IntPtrInput
	ResourceGroupName   pulumi.StringInput
	RfspIndex           pulumi.IntPtrInput
	SimPolicyName       pulumi.StringPtrInput
	SliceConfigurations SliceConfigurationArrayInput
	Tags                pulumi.StringMapInput
	UeAmbr              AmbrInput
}

func (SimPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*simPolicyArgs)(nil)).Elem()
}

type SimPolicyInput interface {
	pulumi.Input

	ToSimPolicyOutput() SimPolicyOutput
	ToSimPolicyOutputWithContext(ctx context.Context) SimPolicyOutput
}

func (*SimPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SimPolicy)(nil)).Elem()
}

func (i *SimPolicy) ToSimPolicyOutput() SimPolicyOutput {
	return i.ToSimPolicyOutputWithContext(context.Background())
}

func (i *SimPolicy) ToSimPolicyOutputWithContext(ctx context.Context) SimPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimPolicyOutput)
}

type SimPolicyOutput struct{ *pulumi.OutputState }

func (SimPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SimPolicy)(nil)).Elem()
}

func (o SimPolicyOutput) ToSimPolicyOutput() SimPolicyOutput {
	return o
}

func (o SimPolicyOutput) ToSimPolicyOutputWithContext(ctx context.Context) SimPolicyOutput {
	return o
}

func (o SimPolicyOutput) DefaultSlice() SliceResourceIdResponseOutput {
	return o.ApplyT(func(v *SimPolicy) SliceResourceIdResponseOutput { return v.DefaultSlice }).(SliceResourceIdResponseOutput)
}

func (o SimPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SimPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SimPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SimPolicyOutput) RegistrationTimer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.IntPtrOutput { return v.RegistrationTimer }).(pulumi.IntPtrOutput)
}

func (o SimPolicyOutput) RfspIndex() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.IntPtrOutput { return v.RfspIndex }).(pulumi.IntPtrOutput)
}

func (o SimPolicyOutput) SiteProvisioningState() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringMapOutput { return v.SiteProvisioningState }).(pulumi.StringMapOutput)
}

func (o SimPolicyOutput) SliceConfigurations() SliceConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *SimPolicy) SliceConfigurationResponseArrayOutput { return v.SliceConfigurations }).(SliceConfigurationResponseArrayOutput)
}

func (o SimPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SimPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SimPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SimPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SimPolicyOutput) UeAmbr() AmbrResponseOutput {
	return o.ApplyT(func(v *SimPolicy) AmbrResponseOutput { return v.UeAmbr }).(AmbrResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(SimPolicyOutput{})
}
