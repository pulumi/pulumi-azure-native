


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDomainService(ctx *pulumi.Context, args *LookupDomainServiceArgs, opts ...pulumi.InvokeOption) (*LookupDomainServiceResult, error) {
	var rv LookupDomainServiceResult
	err := ctx.Invoke("azure-native:aad/v20210501:getDomainService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDomainServiceArgs struct {
	DomainServiceName string `pulumi:"domainServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDomainServiceResult struct {
	ConfigDiagnostics       *ConfigDiagnosticsResponse      `pulumi:"configDiagnostics"`
	DeploymentId            string                          `pulumi:"deploymentId"`
	DomainConfigurationType *string                         `pulumi:"domainConfigurationType"`
	DomainName              *string                         `pulumi:"domainName"`
	DomainSecuritySettings  *DomainSecuritySettingsResponse `pulumi:"domainSecuritySettings"`
	Etag                    *string                         `pulumi:"etag"`
	FilteredSync            *string                         `pulumi:"filteredSync"`
	Id                      string                          `pulumi:"id"`
	LdapsSettings           *LdapsSettingsResponse          `pulumi:"ldapsSettings"`
	Location                *string                         `pulumi:"location"`
	MigrationProperties     MigrationPropertiesResponse     `pulumi:"migrationProperties"`
	Name                    string                          `pulumi:"name"`
	NotificationSettings    *NotificationSettingsResponse   `pulumi:"notificationSettings"`
	ProvisioningState       string                          `pulumi:"provisioningState"`
	ReplicaSets             []ReplicaSetResponse            `pulumi:"replicaSets"`
	ResourceForestSettings  *ResourceForestSettingsResponse `pulumi:"resourceForestSettings"`
	Sku                     *string                         `pulumi:"sku"`
	SyncOwner               string                          `pulumi:"syncOwner"`
	SystemData              SystemDataResponse              `pulumi:"systemData"`
	Tags                    map[string]string               `pulumi:"tags"`
	TenantId                string                          `pulumi:"tenantId"`
	Type                    string                          `pulumi:"type"`
	Version                 int                             `pulumi:"version"`
}


func (val *LookupDomainServiceResult) Defaults() *LookupDomainServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DomainSecuritySettings = tmp.DomainSecuritySettings.Defaults()

	tmp.LdapsSettings = tmp.LdapsSettings.Defaults()

	return &tmp
}

func LookupDomainServiceOutput(ctx *pulumi.Context, args LookupDomainServiceOutputArgs, opts ...pulumi.InvokeOption) LookupDomainServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainServiceResult, error) {
			args := v.(LookupDomainServiceArgs)
			r, err := LookupDomainService(ctx, &args, opts...)
			var s LookupDomainServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainServiceResultOutput)
}

type LookupDomainServiceOutputArgs struct {
	DomainServiceName pulumi.StringInput `pulumi:"domainServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainServiceArgs)(nil)).Elem()
}


type LookupDomainServiceResultOutput struct{ *pulumi.OutputState }

func (LookupDomainServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainServiceResult)(nil)).Elem()
}

func (o LookupDomainServiceResultOutput) ToLookupDomainServiceResultOutput() LookupDomainServiceResultOutput {
	return o
}

func (o LookupDomainServiceResultOutput) ToLookupDomainServiceResultOutputWithContext(ctx context.Context) LookupDomainServiceResultOutput {
	return o
}

func (o LookupDomainServiceResultOutput) ConfigDiagnostics() ConfigDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *ConfigDiagnosticsResponse { return v.ConfigDiagnostics }).(ConfigDiagnosticsResponsePtrOutput)
}

func (o LookupDomainServiceResultOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.DeploymentId }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) DomainConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.DomainConfigurationType }).(pulumi.StringPtrOutput)
}

func (o LookupDomainServiceResultOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

func (o LookupDomainServiceResultOutput) DomainSecuritySettings() DomainSecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *DomainSecuritySettingsResponse { return v.DomainSecuritySettings }).(DomainSecuritySettingsResponsePtrOutput)
}

func (o LookupDomainServiceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupDomainServiceResultOutput) FilteredSync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.FilteredSync }).(pulumi.StringPtrOutput)
}

func (o LookupDomainServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) LdapsSettings() LdapsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *LdapsSettingsResponse { return v.LdapsSettings }).(LdapsSettingsResponsePtrOutput)
}

func (o LookupDomainServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDomainServiceResultOutput) MigrationProperties() MigrationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) MigrationPropertiesResponse { return v.MigrationProperties }).(MigrationPropertiesResponseOutput)
}

func (o LookupDomainServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *NotificationSettingsResponse { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

func (o LookupDomainServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) ReplicaSets() ReplicaSetResponseArrayOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) []ReplicaSetResponse { return v.ReplicaSets }).(ReplicaSetResponseArrayOutput)
}

func (o LookupDomainServiceResultOutput) ResourceForestSettings() ResourceForestSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *ResourceForestSettingsResponse { return v.ResourceForestSettings }).(ResourceForestSettingsResponsePtrOutput)
}

func (o LookupDomainServiceResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o LookupDomainServiceResultOutput) SyncOwner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.SyncOwner }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDomainServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDomainServiceResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainServiceResultOutput{})
}
