


package v20210827

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAttachedDatabaseConfiguration(ctx *pulumi.Context, args *LookupAttachedDatabaseConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupAttachedDatabaseConfigurationResult, error) {
	var rv LookupAttachedDatabaseConfigurationResult
	err := ctx.Invoke("azure-native:kusto/v20210827:getAttachedDatabaseConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttachedDatabaseConfigurationArgs struct {
	AttachedDatabaseConfigurationName string `pulumi:"attachedDatabaseConfigurationName"`
	ClusterName                       string `pulumi:"clusterName"`
	ResourceGroupName                 string `pulumi:"resourceGroupName"`
}


type LookupAttachedDatabaseConfigurationResult struct {
	AttachedDatabaseNames             []string                             `pulumi:"attachedDatabaseNames"`
	ClusterResourceId                 string                               `pulumi:"clusterResourceId"`
	DatabaseName                      string                               `pulumi:"databaseName"`
	DefaultPrincipalsModificationKind string                               `pulumi:"defaultPrincipalsModificationKind"`
	Id                                string                               `pulumi:"id"`
	Location                          *string                              `pulumi:"location"`
	Name                              string                               `pulumi:"name"`
	ProvisioningState                 string                               `pulumi:"provisioningState"`
	TableLevelSharingProperties       *TableLevelSharingPropertiesResponse `pulumi:"tableLevelSharingProperties"`
	Type                              string                               `pulumi:"type"`
}
