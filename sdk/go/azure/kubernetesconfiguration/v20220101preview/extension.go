


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Extension struct {
	pulumi.CustomResourceState

	AksAssignedIdentity            ExtensionResponseAksAssignedIdentityPtrOutput `pulumi:"aksAssignedIdentity"`
	AutoUpgradeMinorVersion        pulumi.BoolPtrOutput                          `pulumi:"autoUpgradeMinorVersion"`
	ConfigurationProtectedSettings pulumi.StringMapOutput                        `pulumi:"configurationProtectedSettings"`
	ConfigurationSettings          pulumi.StringMapOutput                        `pulumi:"configurationSettings"`
	CustomLocationSettings         pulumi.StringMapOutput                        `pulumi:"customLocationSettings"`
	ErrorInfo                      ErrorDetailResponseOutput                     `pulumi:"errorInfo"`
	ExtensionType                  pulumi.StringPtrOutput                        `pulumi:"extensionType"`
	Identity                       IdentityResponsePtrOutput                     `pulumi:"identity"`
	Name                           pulumi.StringOutput                           `pulumi:"name"`
	PackageUri                     pulumi.StringOutput                           `pulumi:"packageUri"`
	ProvisioningState              pulumi.StringOutput                           `pulumi:"provisioningState"`
	ReleaseTrain                   pulumi.StringPtrOutput                        `pulumi:"releaseTrain"`
	Scope                          ScopeResponsePtrOutput                        `pulumi:"scope"`
	Statuses                       ExtensionStatusResponseArrayOutput            `pulumi:"statuses"`
	SystemData                     SystemDataResponseOutput                      `pulumi:"systemData"`
	Type                           pulumi.StringOutput                           `pulumi:"type"`
	Version                        pulumi.StringPtrOutput                        `pulumi:"version"`
}


func NewExtension(ctx *pulumi.Context,
	name string, args *ExtensionArgs, opts ...pulumi.ResourceOption) (*Extension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ClusterResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterResourceName'")
	}
	if args.ClusterRp == nil {
		return nil, errors.New("invalid value for required argument 'ClusterRp'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.AutoUpgradeMinorVersion) {
		args.AutoUpgradeMinorVersion = pulumi.BoolPtr(true)
	}
	if isZero(args.ReleaseTrain) {
		args.ReleaseTrain = pulumi.StringPtr("Stable")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20200701preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20210501preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20210901:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20211101preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220301:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220402preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220701:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20221101:Extension"),
		},
	})
	opts = append(opts, aliases)
	var resource Extension
	err := ctx.RegisterResource("azure-native:kubernetesconfiguration/v20220101preview:Extension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionState, opts ...pulumi.ResourceOption) (*Extension, error) {
	var resource Extension
	err := ctx.ReadResource("azure-native:kubernetesconfiguration/v20220101preview:Extension", name, id, state, &resource, opts...)
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
	AksAssignedIdentity            *ExtensionAksAssignedIdentity `pulumi:"aksAssignedIdentity"`
	AutoUpgradeMinorVersion        *bool                         `pulumi:"autoUpgradeMinorVersion"`
	ClusterName                    string                        `pulumi:"clusterName"`
	ClusterResourceName            string                        `pulumi:"clusterResourceName"`
	ClusterRp                      string                        `pulumi:"clusterRp"`
	ConfigurationProtectedSettings map[string]string             `pulumi:"configurationProtectedSettings"`
	ConfigurationSettings          map[string]string             `pulumi:"configurationSettings"`
	ExtensionName                  *string                       `pulumi:"extensionName"`
	ExtensionType                  *string                       `pulumi:"extensionType"`
	Identity                       *Identity                     `pulumi:"identity"`
	ReleaseTrain                   *string                       `pulumi:"releaseTrain"`
	ResourceGroupName              string                        `pulumi:"resourceGroupName"`
	Scope                          *Scope                        `pulumi:"scope"`
	Statuses                       []ExtensionStatus             `pulumi:"statuses"`
	Version                        *string                       `pulumi:"version"`
}


type ExtensionArgs struct {
	AksAssignedIdentity            ExtensionAksAssignedIdentityPtrInput
	AutoUpgradeMinorVersion        pulumi.BoolPtrInput
	ClusterName                    pulumi.StringInput
	ClusterResourceName            pulumi.StringInput
	ClusterRp                      pulumi.StringInput
	ConfigurationProtectedSettings pulumi.StringMapInput
	ConfigurationSettings          pulumi.StringMapInput
	ExtensionName                  pulumi.StringPtrInput
	ExtensionType                  pulumi.StringPtrInput
	Identity                       IdentityPtrInput
	ReleaseTrain                   pulumi.StringPtrInput
	ResourceGroupName              pulumi.StringInput
	Scope                          ScopePtrInput
	Statuses                       ExtensionStatusArrayInput
	Version                        pulumi.StringPtrInput
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

func (o ExtensionOutput) AksAssignedIdentity() ExtensionResponseAksAssignedIdentityPtrOutput {
	return o.ApplyT(func(v *Extension) ExtensionResponseAksAssignedIdentityPtrOutput { return v.AksAssignedIdentity }).(ExtensionResponseAksAssignedIdentityPtrOutput)
}

func (o ExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.BoolPtrOutput { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o ExtensionOutput) ConfigurationProtectedSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringMapOutput { return v.ConfigurationProtectedSettings }).(pulumi.StringMapOutput)
}

func (o ExtensionOutput) ConfigurationSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringMapOutput { return v.ConfigurationSettings }).(pulumi.StringMapOutput)
}

func (o ExtensionOutput) CustomLocationSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringMapOutput { return v.CustomLocationSettings }).(pulumi.StringMapOutput)
}

func (o ExtensionOutput) ErrorInfo() ErrorDetailResponseOutput {
	return o.ApplyT(func(v *Extension) ErrorDetailResponseOutput { return v.ErrorInfo }).(ErrorDetailResponseOutput)
}

func (o ExtensionOutput) ExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.ExtensionType }).(pulumi.StringPtrOutput)
}

func (o ExtensionOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Extension) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o ExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExtensionOutput) PackageUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.PackageUri }).(pulumi.StringOutput)
}

func (o ExtensionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ExtensionOutput) ReleaseTrain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.ReleaseTrain }).(pulumi.StringPtrOutput)
}

func (o ExtensionOutput) Scope() ScopeResponsePtrOutput {
	return o.ApplyT(func(v *Extension) ScopeResponsePtrOutput { return v.Scope }).(ScopeResponsePtrOutput)
}

func (o ExtensionOutput) Statuses() ExtensionStatusResponseArrayOutput {
	return o.ApplyT(func(v *Extension) ExtensionStatusResponseArrayOutput { return v.Statuses }).(ExtensionStatusResponseArrayOutput)
}

func (o ExtensionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Extension) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ExtensionOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ExtensionOutput{})
}
