


package v20220210preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHostPool(ctx *pulumi.Context, args *LookupHostPoolArgs, opts ...pulumi.InvokeOption) (*LookupHostPoolResult, error) {
	var rv LookupHostPoolResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20220210preview:getHostPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHostPoolArgs struct {
	HostPoolName      string `pulumi:"hostPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupHostPoolResult struct {
	AgentUpdate                   *AgentUpdatePropertiesResponse                       `pulumi:"agentUpdate"`
	ApplicationGroupReferences    []string                                             `pulumi:"applicationGroupReferences"`
	CloudPcResource               bool                                                 `pulumi:"cloudPcResource"`
	CustomRdpProperty             *string                                              `pulumi:"customRdpProperty"`
	Description                   *string                                              `pulumi:"description"`
	Etag                          string                                               `pulumi:"etag"`
	FriendlyName                  *string                                              `pulumi:"friendlyName"`
	HostPoolType                  string                                               `pulumi:"hostPoolType"`
	Id                            string                                               `pulumi:"id"`
	Identity                      *ResourceModelWithAllowedPropertySetResponseIdentity `pulumi:"identity"`
	Kind                          *string                                              `pulumi:"kind"`
	LoadBalancerType              string                                               `pulumi:"loadBalancerType"`
	Location                      *string                                              `pulumi:"location"`
	ManagedBy                     *string                                              `pulumi:"managedBy"`
	MaxSessionLimit               *int                                                 `pulumi:"maxSessionLimit"`
	MigrationRequest              *MigrationRequestPropertiesResponse                  `pulumi:"migrationRequest"`
	Name                          string                                               `pulumi:"name"`
	ObjectId                      string                                               `pulumi:"objectId"`
	PersonalDesktopAssignmentType *string                                              `pulumi:"personalDesktopAssignmentType"`
	Plan                          *ResourceModelWithAllowedPropertySetResponsePlan     `pulumi:"plan"`
	PreferredAppGroupType         string                                               `pulumi:"preferredAppGroupType"`
	PrivateEndpointConnections    []PrivateEndpointConnectionResponse                  `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess           *string                                              `pulumi:"publicNetworkAccess"`
	RegistrationInfo              *RegistrationInfoResponse                            `pulumi:"registrationInfo"`
	Ring                          *int                                                 `pulumi:"ring"`
	Sku                           *ResourceModelWithAllowedPropertySetResponseSku      `pulumi:"sku"`
	SsoClientId                   *string                                              `pulumi:"ssoClientId"`
	SsoClientSecretKeyVaultPath   *string                                              `pulumi:"ssoClientSecretKeyVaultPath"`
	SsoSecretType                 *string                                              `pulumi:"ssoSecretType"`
	SsoadfsAuthority              *string                                              `pulumi:"ssoadfsAuthority"`
	StartVMOnConnect              *bool                                                `pulumi:"startVMOnConnect"`
	SystemData                    SystemDataResponse                                   `pulumi:"systemData"`
	Tags                          map[string]string                                    `pulumi:"tags"`
	Type                          string                                               `pulumi:"type"`
	ValidationEnvironment         *bool                                                `pulumi:"validationEnvironment"`
	VmTemplate                    *string                                              `pulumi:"vmTemplate"`
}

func LookupHostPoolOutput(ctx *pulumi.Context, args LookupHostPoolOutputArgs, opts ...pulumi.InvokeOption) LookupHostPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHostPoolResult, error) {
			args := v.(LookupHostPoolArgs)
			r, err := LookupHostPool(ctx, &args, opts...)
			var s LookupHostPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHostPoolResultOutput)
}

type LookupHostPoolOutputArgs struct {
	HostPoolName      pulumi.StringInput `pulumi:"hostPoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHostPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostPoolArgs)(nil)).Elem()
}


type LookupHostPoolResultOutput struct{ *pulumi.OutputState }

func (LookupHostPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostPoolResult)(nil)).Elem()
}

func (o LookupHostPoolResultOutput) ToLookupHostPoolResultOutput() LookupHostPoolResultOutput {
	return o
}

func (o LookupHostPoolResultOutput) ToLookupHostPoolResultOutputWithContext(ctx context.Context) LookupHostPoolResultOutput {
	return o
}

func (o LookupHostPoolResultOutput) AgentUpdate() AgentUpdatePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *AgentUpdatePropertiesResponse { return v.AgentUpdate }).(AgentUpdatePropertiesResponsePtrOutput)
}

func (o LookupHostPoolResultOutput) ApplicationGroupReferences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostPoolResult) []string { return v.ApplicationGroupReferences }).(pulumi.StringArrayOutput)
}

func (o LookupHostPoolResultOutput) CloudPcResource() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostPoolResult) bool { return v.CloudPcResource }).(pulumi.BoolOutput)
}

func (o LookupHostPoolResultOutput) CustomRdpProperty() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.CustomRdpProperty }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) HostPoolType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.HostPoolType }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *ResourceModelWithAllowedPropertySetResponseIdentity { return v.Identity }).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

func (o LookupHostPoolResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) LoadBalancerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.LoadBalancerType }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) MaxSessionLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *int { return v.MaxSessionLimit }).(pulumi.IntPtrOutput)
}

func (o LookupHostPoolResultOutput) MigrationRequest() MigrationRequestPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *MigrationRequestPropertiesResponse { return v.MigrationRequest }).(MigrationRequestPropertiesResponsePtrOutput)
}

func (o LookupHostPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) PersonalDesktopAssignmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.PersonalDesktopAssignmentType }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *ResourceModelWithAllowedPropertySetResponsePlan { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

func (o LookupHostPoolResultOutput) PreferredAppGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.PreferredAppGroupType }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupHostPoolResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupHostPoolResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) RegistrationInfo() RegistrationInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *RegistrationInfoResponse { return v.RegistrationInfo }).(RegistrationInfoResponsePtrOutput)
}

func (o LookupHostPoolResultOutput) Ring() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *int { return v.Ring }).(pulumi.IntPtrOutput)
}

func (o LookupHostPoolResultOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *ResourceModelWithAllowedPropertySetResponseSku { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

func (o LookupHostPoolResultOutput) SsoClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.SsoClientId }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) SsoClientSecretKeyVaultPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.SsoClientSecretKeyVaultPath }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) SsoSecretType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.SsoSecretType }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) SsoadfsAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.SsoadfsAuthority }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) StartVMOnConnect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *bool { return v.StartVMOnConnect }).(pulumi.BoolPtrOutput)
}

func (o LookupHostPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHostPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupHostPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupHostPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupHostPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) ValidationEnvironment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *bool { return v.ValidationEnvironment }).(pulumi.BoolPtrOutput)
}

func (o LookupHostPoolResultOutput) VmTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.VmTemplate }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHostPoolResultOutput{})
}
