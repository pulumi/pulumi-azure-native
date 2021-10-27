


package v20140401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseThreatDetectionPolicy(ctx *pulumi.Context, args *LookupDatabaseThreatDetectionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseThreatDetectionPolicyResult, error) {
	var rv LookupDatabaseThreatDetectionPolicyResult
	err := ctx.Invoke("azure-native:sql/v20140401:getDatabaseThreatDetectionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseThreatDetectionPolicyArgs struct {
	DatabaseName            string `pulumi:"databaseName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SecurityAlertPolicyName string `pulumi:"securityAlertPolicyName"`
	ServerName              string `pulumi:"serverName"`
}


type LookupDatabaseThreatDetectionPolicyResult struct {
	DisabledAlerts     *string `pulumi:"disabledAlerts"`
	EmailAccountAdmins *string `pulumi:"emailAccountAdmins"`
	EmailAddresses     *string `pulumi:"emailAddresses"`
	Id                 string  `pulumi:"id"`
	Kind               string  `pulumi:"kind"`
	Location           *string `pulumi:"location"`
	Name               string  `pulumi:"name"`
	RetentionDays      *int    `pulumi:"retentionDays"`
	State              string  `pulumi:"state"`
	StorageEndpoint    *string `pulumi:"storageEndpoint"`
	Type               string  `pulumi:"type"`
	UseServerDefault   *string `pulumi:"useServerDefault"`
}
