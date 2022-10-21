


package v20210901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Extension struct {
	pulumi.CustomResourceState

	AggregateState          pulumi.StringOutput                      `pulumi:"aggregateState"`
	AutoUpgradeMinorVersion pulumi.BoolPtrOutput                     `pulumi:"autoUpgradeMinorVersion"`
	CreatedAt               pulumi.StringPtrOutput                   `pulumi:"createdAt"`
	CreatedBy               pulumi.StringPtrOutput                   `pulumi:"createdBy"`
	CreatedByType           pulumi.StringPtrOutput                   `pulumi:"createdByType"`
	ForceUpdateTag          pulumi.StringPtrOutput                   `pulumi:"forceUpdateTag"`
	LastModifiedAt          pulumi.StringPtrOutput                   `pulumi:"lastModifiedAt"`
	LastModifiedBy          pulumi.StringPtrOutput                   `pulumi:"lastModifiedBy"`
	LastModifiedByType      pulumi.StringPtrOutput                   `pulumi:"lastModifiedByType"`
	Name                    pulumi.StringOutput                      `pulumi:"name"`
	PerNodeExtensionDetails PerNodeExtensionStateResponseArrayOutput `pulumi:"perNodeExtensionDetails"`
	ProtectedSettings       pulumi.AnyOutput                         `pulumi:"protectedSettings"`
	ProvisioningState       pulumi.StringOutput                      `pulumi:"provisioningState"`
	Publisher               pulumi.StringPtrOutput                   `pulumi:"publisher"`
	Settings                pulumi.AnyOutput                         `pulumi:"settings"`
	Type                    pulumi.StringOutput                      `pulumi:"type"`
	TypeHandlerVersion      pulumi.StringPtrOutput                   `pulumi:"typeHandlerVersion"`
}


func NewExtension(ctx *pulumi.Context,
	name string, args *ExtensionArgs, opts ...pulumi.ResourceOption) (*Extension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ArcSettingName == nil {
		return nil, errors.New("invalid value for required argument 'ArcSettingName'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210101preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220101:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220301:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220501:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220901:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221001:Extension"),
		},
	})
	opts = append(opts, aliases)
	var resource Extension
	err := ctx.RegisterResource("azure-native:azurestackhci/v20210901:Extension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionState, opts ...pulumi.ResourceOption) (*Extension, error) {
	var resource Extension
	err := ctx.ReadResource("azure-native:azurestackhci/v20210901:Extension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type extensionState struct {
}

type ExtensionState struct {
}

func (ExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionState)(nil)).Elem()
}

type extensionArgs struct {
	ArcSettingName          string      `pulumi:"arcSettingName"`
	AutoUpgradeMinorVersion *bool       `pulumi:"autoUpgradeMinorVersion"`
	ClusterName             string      `pulumi:"clusterName"`
	CreatedAt               *string     `pulumi:"createdAt"`
	CreatedBy               *string     `pulumi:"createdBy"`
	CreatedByType           *string     `pulumi:"createdByType"`
	ExtensionName           *string     `pulumi:"extensionName"`
	ForceUpdateTag          *string     `pulumi:"forceUpdateTag"`
	LastModifiedAt          *string     `pulumi:"lastModifiedAt"`
	LastModifiedBy          *string     `pulumi:"lastModifiedBy"`
	LastModifiedByType      *string     `pulumi:"lastModifiedByType"`
	ProtectedSettings       interface{} `pulumi:"protectedSettings"`
	Publisher               *string     `pulumi:"publisher"`
	ResourceGroupName       string      `pulumi:"resourceGroupName"`
	Settings                interface{} `pulumi:"settings"`
	Type                    *string     `pulumi:"type"`
	TypeHandlerVersion      *string     `pulumi:"typeHandlerVersion"`
}


type ExtensionArgs struct {
	ArcSettingName          pulumi.StringInput
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	ClusterName             pulumi.StringInput
	CreatedAt               pulumi.StringPtrInput
	CreatedBy               pulumi.StringPtrInput
	CreatedByType           pulumi.StringPtrInput
	ExtensionName           pulumi.StringPtrInput
	ForceUpdateTag          pulumi.StringPtrInput
	LastModifiedAt          pulumi.StringPtrInput
	LastModifiedBy          pulumi.StringPtrInput
	LastModifiedByType      pulumi.StringPtrInput
	ProtectedSettings       pulumi.Input
	Publisher               pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	Settings                pulumi.Input
	Type                    pulumi.StringPtrInput
	TypeHandlerVersion      pulumi.StringPtrInput
}

func (ExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionArgs)(nil)).Elem()
}

type ExtensionInput interface {
	pulumi.Input

	ToExtensionOutput() ExtensionOutput
	ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput
}

func (*Extension) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (i *Extension) ToExtensionOutput() ExtensionOutput {
	return i.ToExtensionOutputWithContext(context.Background())
}

func (i *Extension) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionOutput)
}

type ExtensionOutput struct{ *pulumi.OutputState }

func (ExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (o ExtensionOutput) ToExtensionOutput() ExtensionOutput {
	return o
}

func (o ExtensionOutput) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return o
}

func (o ExtensionOutput) AggregateState() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.AggregateState }).(pulumi.StringOutput)
}

func (o ExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.BoolPtrOutput { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o ExtensionOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o ExtensionOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o ExtensionOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o ExtensionOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o ExtensionOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o ExtensionOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o ExtensionOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o ExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExtensionOutput) PerNodeExtensionDetails() PerNodeExtensionStateResponseArrayOutput {
	return o.ApplyT(func(v *Extension) PerNodeExtensionStateResponseArrayOutput { return v.PerNodeExtensionDetails }).(PerNodeExtensionStateResponseArrayOutput)
}

func (o ExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *Extension) pulumi.AnyOutput { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o ExtensionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ExtensionOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *Extension) pulumi.AnyOutput { return v.Settings }).(pulumi.AnyOutput)
}

func (o ExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ExtensionOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ExtensionOutput{})
}
