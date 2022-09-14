


package v20220301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VolumeQuotaRule struct {
	pulumi.CustomResourceState

	Location          pulumi.StringOutput      `pulumi:"location"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	QuotaSizeInKiBs   pulumi.Float64PtrOutput  `pulumi:"quotaSizeInKiBs"`
	QuotaTarget       pulumi.StringPtrOutput   `pulumi:"quotaTarget"`
	QuotaType         pulumi.StringPtrOutput   `pulumi:"quotaType"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewVolumeQuotaRule(ctx *pulumi.Context,
	name string, args *VolumeQuotaRuleArgs, opts ...pulumi.ResourceOption) (*VolumeQuotaRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VolumeName == nil {
		return nil, errors.New("invalid value for required argument 'VolumeName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp:VolumeQuotaRule"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:VolumeQuotaRule"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220501:VolumeQuotaRule"),
		},
	})
	opts = append(opts, aliases)
	var resource VolumeQuotaRule
	err := ctx.RegisterResource("azure-native:netapp/v20220301:VolumeQuotaRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVolumeQuotaRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeQuotaRuleState, opts ...pulumi.ResourceOption) (*VolumeQuotaRule, error) {
	var resource VolumeQuotaRule
	err := ctx.ReadResource("azure-native:netapp/v20220301:VolumeQuotaRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type volumeQuotaRuleState struct {
}

type VolumeQuotaRuleState struct {
}

func (VolumeQuotaRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeQuotaRuleState)(nil)).Elem()
}

type volumeQuotaRuleArgs struct {
	AccountName         string            `pulumi:"accountName"`
	Location            *string           `pulumi:"location"`
	PoolName            string            `pulumi:"poolName"`
	QuotaSizeInKiBs     *float64          `pulumi:"quotaSizeInKiBs"`
	QuotaTarget         *string           `pulumi:"quotaTarget"`
	QuotaType           *string           `pulumi:"quotaType"`
	ResourceGroupName   string            `pulumi:"resourceGroupName"`
	Tags                map[string]string `pulumi:"tags"`
	VolumeName          string            `pulumi:"volumeName"`
	VolumeQuotaRuleName *string           `pulumi:"volumeQuotaRuleName"`
}


type VolumeQuotaRuleArgs struct {
	AccountName         pulumi.StringInput
	Location            pulumi.StringPtrInput
	PoolName            pulumi.StringInput
	QuotaSizeInKiBs     pulumi.Float64PtrInput
	QuotaTarget         pulumi.StringPtrInput
	QuotaType           pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
	VolumeName          pulumi.StringInput
	VolumeQuotaRuleName pulumi.StringPtrInput
}

func (VolumeQuotaRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeQuotaRuleArgs)(nil)).Elem()
}

type VolumeQuotaRuleInput interface {
	pulumi.Input

	ToVolumeQuotaRuleOutput() VolumeQuotaRuleOutput
	ToVolumeQuotaRuleOutputWithContext(ctx context.Context) VolumeQuotaRuleOutput
}

func (*VolumeQuotaRule) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeQuotaRule)(nil)).Elem()
}

func (i *VolumeQuotaRule) ToVolumeQuotaRuleOutput() VolumeQuotaRuleOutput {
	return i.ToVolumeQuotaRuleOutputWithContext(context.Background())
}

func (i *VolumeQuotaRule) ToVolumeQuotaRuleOutputWithContext(ctx context.Context) VolumeQuotaRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeQuotaRuleOutput)
}

type VolumeQuotaRuleOutput struct{ *pulumi.OutputState }

func (VolumeQuotaRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeQuotaRule)(nil)).Elem()
}

func (o VolumeQuotaRuleOutput) ToVolumeQuotaRuleOutput() VolumeQuotaRuleOutput {
	return o
}

func (o VolumeQuotaRuleOutput) ToVolumeQuotaRuleOutputWithContext(ctx context.Context) VolumeQuotaRuleOutput {
	return o
}

func (o VolumeQuotaRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VolumeQuotaRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeQuotaRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VolumeQuotaRuleOutput) QuotaSizeInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.Float64PtrOutput { return v.QuotaSizeInKiBs }).(pulumi.Float64PtrOutput)
}

func (o VolumeQuotaRuleOutput) QuotaTarget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringPtrOutput { return v.QuotaTarget }).(pulumi.StringPtrOutput)
}

func (o VolumeQuotaRuleOutput) QuotaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringPtrOutput { return v.QuotaType }).(pulumi.StringPtrOutput)
}

func (o VolumeQuotaRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VolumeQuotaRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VolumeQuotaRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeQuotaRuleOutput{})
}
