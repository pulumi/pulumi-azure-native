


package v20210801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseSecurityAlertPolicy(ctx *pulumi.Context, args *LookupDatabaseSecurityAlertPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseSecurityAlertPolicyResult, error) {
	var rv LookupDatabaseSecurityAlertPolicyResult
	err := ctx.Invoke("azure-native:sql/v20210801preview:getDatabaseSecurityAlertPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseSecurityAlertPolicyArgs struct {
	DatabaseName            string `pulumi:"databaseName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SecurityAlertPolicyName string `pulumi:"securityAlertPolicyName"`
	ServerName              string `pulumi:"serverName"`
}


type LookupDatabaseSecurityAlertPolicyResult struct {
	CreationTime            string             `pulumi:"creationTime"`
	DisabledAlerts          []string           `pulumi:"disabledAlerts"`
	EmailAccountAdmins      *bool              `pulumi:"emailAccountAdmins"`
	EmailAddresses          []string           `pulumi:"emailAddresses"`
	Id                      string             `pulumi:"id"`
	Name                    string             `pulumi:"name"`
	RetentionDays           *int               `pulumi:"retentionDays"`
	State                   string             `pulumi:"state"`
	StorageAccountAccessKey *string            `pulumi:"storageAccountAccessKey"`
	StorageEndpoint         *string            `pulumi:"storageEndpoint"`
	SystemData              SystemDataResponse `pulumi:"systemData"`
	Type                    string             `pulumi:"type"`
}
