


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HostPool struct {
	pulumi.CustomResourceState

	AgentUpdate                   AgentUpdatePropertiesResponsePtrOutput                       `pulumi:"agentUpdate"`
	ApplicationGroupReferences    pulumi.StringArrayOutput                                     `pulumi:"applicationGroupReferences"`
	CloudPcResource               pulumi.BoolOutput                                            `pulumi:"cloudPcResource"`
	CustomRdpProperty             pulumi.StringPtrOutput                                       `pulumi:"customRdpProperty"`
	Description                   pulumi.StringPtrOutput                                       `pulumi:"description"`
	Etag                          pulumi.StringOutput                                          `pulumi:"etag"`
	FriendlyName                  pulumi.StringPtrOutput                                       `pulumi:"friendlyName"`
	HostPoolType                  pulumi.StringOutput                                          `pulumi:"hostPoolType"`
	Identity                      ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput `pulumi:"identity"`
	Kind                          pulumi.StringPtrOutput                                       `pulumi:"kind"`
	LoadBalancerType              pulumi.StringOutput                                          `pulumi:"loadBalancerType"`
	Location                      pulumi.StringPtrOutput                                       `pulumi:"location"`
	ManagedBy                     pulumi.StringPtrOutput                                       `pulumi:"managedBy"`
	MaxSessionLimit               pulumi.IntPtrOutput                                          `pulumi:"maxSessionLimit"`
	MigrationRequest              MigrationRequestPropertiesResponsePtrOutput                  `pulumi:"migrationRequest"`
	Name                          pulumi.StringOutput                                          `pulumi:"name"`
	ObjectId                      pulumi.StringOutput                                          `pulumi:"objectId"`
	PersonalDesktopAssignmentType pulumi.StringPtrOutput                                       `pulumi:"personalDesktopAssignmentType"`
	Plan                          ResourceModelWithAllowedPropertySetResponsePlanPtrOutput     `pulumi:"plan"`
	PreferredAppGroupType         pulumi.StringOutput                                          `pulumi:"preferredAppGroupType"`
	PrivateEndpointConnections    PrivateEndpointConnectionResponseArrayOutput                 `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess           pulumi.StringPtrOutput                                       `pulumi:"publicNetworkAccess"`
	RegistrationInfo              RegistrationInfoResponsePtrOutput                            `pulumi:"registrationInfo"`
	Ring                          pulumi.IntPtrOutput                                          `pulumi:"ring"`
	Sku                           ResourceModelWithAllowedPropertySetResponseSkuPtrOutput      `pulumi:"sku"`
	SsoClientId                   pulumi.StringPtrOutput                                       `pulumi:"ssoClientId"`
	SsoClientSecretKeyVaultPath   pulumi.StringPtrOutput                                       `pulumi:"ssoClientSecretKeyVaultPath"`
	SsoSecretType                 pulumi.StringPtrOutput                                       `pulumi:"ssoSecretType"`
	SsoadfsAuthority              pulumi.StringPtrOutput                                       `pulumi:"ssoadfsAuthority"`
	StartVMOnConnect              pulumi.BoolPtrOutput                                         `pulumi:"startVMOnConnect"`
	SystemData                    SystemDataResponseOutput                                     `pulumi:"systemData"`
	Tags                          pulumi.StringMapOutput                                       `pulumi:"tags"`
	Type                          pulumi.StringOutput                                          `pulumi:"type"`
	ValidationEnvironment         pulumi.BoolPtrOutput                                         `pulumi:"validationEnvironment"`
	VmTemplate                    pulumi.StringPtrOutput                                       `pulumi:"vmTemplate"`
}


func NewHostPool(ctx *pulumi.Context,
	name string, args *HostPoolArgs, opts ...pulumi.ResourceOption) (*HostPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostPoolType == nil {
		return nil, errors.New("invalid value for required argument 'HostPoolType'")
	}
	if args.LoadBalancerType == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerType'")
	}
	if args.PreferredAppGroupType == nil {
		return nil, errors.New("invalid value for required argument 'PreferredAppGroupType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190123preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190924preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20191210preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20200921preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201019preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201102preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201110preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210114preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210201preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210712:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220210preview:HostPool"),
		},
	})
	opts = append(opts, aliases)
	var resource HostPool
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20220401preview:HostPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHostPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostPoolState, opts ...pulumi.ResourceOption) (*HostPool, error) {
	var resource HostPool
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20220401preview:HostPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hostPoolState struct {
}

type HostPoolState struct {
}

func (HostPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostPoolState)(nil)).Elem()
}

type hostPoolArgs struct {
	AgentUpdate                   *AgentUpdateProperties                       `pulumi:"agentUpdate"`
	CustomRdpProperty             *string                                      `pulumi:"customRdpProperty"`
	Description                   *string                                      `pulumi:"description"`
	FriendlyName                  *string                                      `pulumi:"friendlyName"`
	HostPoolName                  *string                                      `pulumi:"hostPoolName"`
	HostPoolType                  string                                       `pulumi:"hostPoolType"`
	Identity                      *ResourceModelWithAllowedPropertySetIdentity `pulumi:"identity"`
	Kind                          *string                                      `pulumi:"kind"`
	LoadBalancerType              string                                       `pulumi:"loadBalancerType"`
	Location                      *string                                      `pulumi:"location"`
	ManagedBy                     *string                                      `pulumi:"managedBy"`
	MaxSessionLimit               *int                                         `pulumi:"maxSessionLimit"`
	MigrationRequest              *MigrationRequestProperties                  `pulumi:"migrationRequest"`
	PersonalDesktopAssignmentType *string                                      `pulumi:"personalDesktopAssignmentType"`
	Plan                          *ResourceModelWithAllowedPropertySetPlan     `pulumi:"plan"`
	PreferredAppGroupType         string                                       `pulumi:"preferredAppGroupType"`
	PublicNetworkAccess           *string                                      `pulumi:"publicNetworkAccess"`
	RegistrationInfo              *RegistrationInfo                            `pulumi:"registrationInfo"`
	ResourceGroupName             string                                       `pulumi:"resourceGroupName"`
	Ring                          *int                                         `pulumi:"ring"`
	Sku                           *ResourceModelWithAllowedPropertySetSku      `pulumi:"sku"`
	SsoClientId                   *string                                      `pulumi:"ssoClientId"`
	SsoClientSecretKeyVaultPath   *string                                      `pulumi:"ssoClientSecretKeyVaultPath"`
	SsoSecretType                 *string                                      `pulumi:"ssoSecretType"`
	SsoadfsAuthority              *string                                      `pulumi:"ssoadfsAuthority"`
	StartVMOnConnect              *bool                                        `pulumi:"startVMOnConnect"`
	Tags                          map[string]string                            `pulumi:"tags"`
	ValidationEnvironment         *bool                                        `pulumi:"validationEnvironment"`
	VmTemplate                    *string                                      `pulumi:"vmTemplate"`
}


type HostPoolArgs struct {
	AgentUpdate                   AgentUpdatePropertiesPtrInput
	CustomRdpProperty             pulumi.StringPtrInput
	Description                   pulumi.StringPtrInput
	FriendlyName                  pulumi.StringPtrInput
	HostPoolName                  pulumi.StringPtrInput
	HostPoolType                  pulumi.StringInput
	Identity                      ResourceModelWithAllowedPropertySetIdentityPtrInput
	Kind                          pulumi.StringPtrInput
	LoadBalancerType              pulumi.StringInput
	Location                      pulumi.StringPtrInput
	ManagedBy                     pulumi.StringPtrInput
	MaxSessionLimit               pulumi.IntPtrInput
	MigrationRequest              MigrationRequestPropertiesPtrInput
	PersonalDesktopAssignmentType pulumi.StringPtrInput
	Plan                          ResourceModelWithAllowedPropertySetPlanPtrInput
	PreferredAppGroupType         pulumi.StringInput
	PublicNetworkAccess           pulumi.StringPtrInput
	RegistrationInfo              RegistrationInfoPtrInput
	ResourceGroupName             pulumi.StringInput
	Ring                          pulumi.IntPtrInput
	Sku                           ResourceModelWithAllowedPropertySetSkuPtrInput
	SsoClientId                   pulumi.StringPtrInput
	SsoClientSecretKeyVaultPath   pulumi.StringPtrInput
	SsoSecretType                 pulumi.StringPtrInput
	SsoadfsAuthority              pulumi.StringPtrInput
	StartVMOnConnect              pulumi.BoolPtrInput
	Tags                          pulumi.StringMapInput
	ValidationEnvironment         pulumi.BoolPtrInput
	VmTemplate                    pulumi.StringPtrInput
}

func (HostPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostPoolArgs)(nil)).Elem()
}

type HostPoolInput interface {
	pulumi.Input

	ToHostPoolOutput() HostPoolOutput
	ToHostPoolOutputWithContext(ctx context.Context) HostPoolOutput
}

func (*HostPool) ElementType() reflect.Type {
	return reflect.TypeOf((**HostPool)(nil)).Elem()
}

func (i *HostPool) ToHostPoolOutput() HostPoolOutput {
	return i.ToHostPoolOutputWithContext(context.Background())
}

func (i *HostPool) ToHostPoolOutputWithContext(ctx context.Context) HostPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPoolOutput)
}

type HostPoolOutput struct{ *pulumi.OutputState }

func (HostPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostPool)(nil)).Elem()
}

func (o HostPoolOutput) ToHostPoolOutput() HostPoolOutput {
	return o
}

func (o HostPoolOutput) ToHostPoolOutputWithContext(ctx context.Context) HostPoolOutput {
	return o
}

func (o HostPoolOutput) AgentUpdate() AgentUpdatePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *HostPool) AgentUpdatePropertiesResponsePtrOutput { return v.AgentUpdate }).(AgentUpdatePropertiesResponsePtrOutput)
}

func (o HostPoolOutput) ApplicationGroupReferences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringArrayOutput { return v.ApplicationGroupReferences }).(pulumi.StringArrayOutput)
}

func (o HostPoolOutput) CloudPcResource() pulumi.BoolOutput {
	return o.ApplyT(func(v *HostPool) pulumi.BoolOutput { return v.CloudPcResource }).(pulumi.BoolOutput)
}

func (o HostPoolOutput) CustomRdpProperty() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.CustomRdpProperty }).(pulumi.StringPtrOutput)
}

func (o HostPoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o HostPoolOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o HostPoolOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o HostPoolOutput) HostPoolType() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.HostPoolType }).(pulumi.StringOutput)
}

func (o HostPoolOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v *HostPool) ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput { return v.Identity }).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

func (o HostPoolOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o HostPoolOutput) LoadBalancerType() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.LoadBalancerType }).(pulumi.StringOutput)
}

func (o HostPoolOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o HostPoolOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o HostPoolOutput) MaxSessionLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.IntPtrOutput { return v.MaxSessionLimit }).(pulumi.IntPtrOutput)
}

func (o HostPoolOutput) MigrationRequest() MigrationRequestPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *HostPool) MigrationRequestPropertiesResponsePtrOutput { return v.MigrationRequest }).(MigrationRequestPropertiesResponsePtrOutput)
}

func (o HostPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o HostPoolOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.ObjectId }).(pulumi.StringOutput)
}

func (o HostPoolOutput) PersonalDesktopAssignmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.PersonalDesktopAssignmentType }).(pulumi.StringPtrOutput)
}

func (o HostPoolOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v *HostPool) ResourceModelWithAllowedPropertySetResponsePlanPtrOutput { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

func (o HostPoolOutput) PreferredAppGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.PreferredAppGroupType }).(pulumi.StringOutput)
}

func (o HostPoolOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *HostPool) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o HostPoolOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o HostPoolOutput) RegistrationInfo() RegistrationInfoResponsePtrOutput {
	return o.ApplyT(func(v *HostPool) RegistrationInfoResponsePtrOutput { return v.RegistrationInfo }).(RegistrationInfoResponsePtrOutput)
}

func (o HostPoolOutput) Ring() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.IntPtrOutput { return v.Ring }).(pulumi.IntPtrOutput)
}

func (o HostPoolOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v *HostPool) ResourceModelWithAllowedPropertySetResponseSkuPtrOutput { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

func (o HostPoolOutput) SsoClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.SsoClientId }).(pulumi.StringPtrOutput)
}

func (o HostPoolOutput) SsoClientSecretKeyVaultPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.SsoClientSecretKeyVaultPath }).(pulumi.StringPtrOutput)
}

func (o HostPoolOutput) SsoSecretType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.SsoSecretType }).(pulumi.StringPtrOutput)
}

func (o HostPoolOutput) SsoadfsAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.SsoadfsAuthority }).(pulumi.StringPtrOutput)
}

func (o HostPoolOutput) StartVMOnConnect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.BoolPtrOutput { return v.StartVMOnConnect }).(pulumi.BoolPtrOutput)
}

func (o HostPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HostPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o HostPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o HostPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o HostPoolOutput) ValidationEnvironment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.BoolPtrOutput { return v.ValidationEnvironment }).(pulumi.BoolPtrOutput)
}

func (o HostPoolOutput) VmTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.VmTemplate }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HostPoolOutput{})
}
