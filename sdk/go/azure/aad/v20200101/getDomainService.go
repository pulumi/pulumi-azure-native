


package v20200101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDomainService(ctx *pulumi.Context, args *LookupDomainServiceArgs, opts ...pulumi.InvokeOption) (*LookupDomainServiceResult, error) {
	var rv LookupDomainServiceResult
	err := ctx.Invoke("azure-native:aad/v20200101:getDomainService", args, &rv, opts...)
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
