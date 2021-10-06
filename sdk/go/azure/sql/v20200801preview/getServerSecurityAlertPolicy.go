


package v20200801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerSecurityAlertPolicy(ctx *pulumi.Context, args *LookupServerSecurityAlertPolicyArgs, opts ...pulumi.InvokeOption) (*LookupServerSecurityAlertPolicyResult, error) {
	var rv LookupServerSecurityAlertPolicyResult
	err := ctx.Invoke("azure-native:sql/v20200801preview:getServerSecurityAlertPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerSecurityAlertPolicyArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SecurityAlertPolicyName string `pulumi:"securityAlertPolicyName"`
	ServerName              string `pulumi:"serverName"`
}


type LookupServerSecurityAlertPolicyResult struct {
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
