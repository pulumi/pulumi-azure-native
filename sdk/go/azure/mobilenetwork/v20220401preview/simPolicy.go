


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SimPolicy struct {
	pulumi.CustomResourceState

	CreatedAt           pulumi.StringPtrOutput                `pulumi:"createdAt"`
	CreatedBy           pulumi.StringPtrOutput                `pulumi:"createdBy"`
	CreatedByType       pulumi.StringPtrOutput                `pulumi:"createdByType"`
	DefaultSlice        SliceResourceIdResponseOutput         `pulumi:"defaultSlice"`
	LastModifiedAt      pulumi.StringPtrOutput                `pulumi:"lastModifiedAt"`
	LastModifiedBy      pulumi.StringPtrOutput                `pulumi:"lastModifiedBy"`
	LastModifiedByType  pulumi.StringPtrOutput                `pulumi:"lastModifiedByType"`
	Location            pulumi.StringOutput                   `pulumi:"location"`
	Name                pulumi.StringOutput                   `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput                   `pulumi:"provisioningState"`
	RegistrationTimer   pulumi.IntPtrOutput                   `pulumi:"registrationTimer"`
	RfspIndex           pulumi.IntPtrOutput                   `pulumi:"rfspIndex"`
	SliceConfigurations SliceConfigurationResponseArrayOutput `pulumi:"sliceConfigurations"`
	SystemData          SystemDataResponseOutput              `pulumi:"systemData"`
	Tags                pulumi.StringMapOutput                `pulumi:"tags"`
	Type                pulumi.StringOutput                   `pulumi:"type"`
	UeAmbr              AmbrResponseOutput                    `pulumi:"ueAmbr"`
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
	})
	opts = append(opts, aliases)
	var resource SimPolicy
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20220401preview:SimPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSimPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SimPolicyState, opts ...pulumi.ResourceOption) (*SimPolicy, error) {
	var resource SimPolicy
	err := ctx.ReadResource("azure-native:mobilenetwork/v20220401preview:SimPolicy", name, id, state, &resource, opts...)
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
	CreatedAt           *string              `pulumi:"createdAt"`
	CreatedBy           *string              `pulumi:"createdBy"`
	CreatedByType       *string              `pulumi:"createdByType"`
	DefaultSlice        SliceResourceId      `pulumi:"defaultSlice"`
	LastModifiedAt      *string              `pulumi:"lastModifiedAt"`
	LastModifiedBy      *string              `pulumi:"lastModifiedBy"`
	LastModifiedByType  *string              `pulumi:"lastModifiedByType"`
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
	CreatedAt           pulumi.StringPtrInput
	CreatedBy           pulumi.StringPtrInput
	CreatedByType       pulumi.StringPtrInput
	DefaultSlice        SliceResourceIdInput
	LastModifiedAt      pulumi.StringPtrInput
	LastModifiedBy      pulumi.StringPtrInput
	LastModifiedByType  pulumi.StringPtrInput
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

func (o SimPolicyOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SimPolicyOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SimPolicyOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SimPolicyOutput) DefaultSlice() SliceResourceIdResponseOutput {
	return o.ApplyT(func(v *SimPolicy) SliceResourceIdResponseOutput { return v.DefaultSlice }).(SliceResourceIdResponseOutput)
}

func (o SimPolicyOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SimPolicyOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SimPolicyOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
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
