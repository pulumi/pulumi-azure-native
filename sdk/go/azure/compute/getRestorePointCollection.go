


package compute

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRestorePointCollection(ctx *pulumi.Context, args *LookupRestorePointCollectionArgs, opts ...pulumi.InvokeOption) (*LookupRestorePointCollectionResult, error) {
	var rv LookupRestorePointCollectionResult
	err := ctx.Invoke("azure-native:compute:getRestorePointCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRestorePointCollectionArgs struct {
	Expand                     *string `pulumi:"expand"`
	ResourceGroupName          string  `pulumi:"resourceGroupName"`
	RestorePointCollectionName string  `pulumi:"restorePointCollectionName"`
}


type LookupRestorePointCollectionResult struct {
	Id                       string                                          `pulumi:"id"`
	Location                 string                                          `pulumi:"location"`
	Name                     string                                          `pulumi:"name"`
	ProvisioningState        string                                          `pulumi:"provisioningState"`
	RestorePointCollectionId string                                          `pulumi:"restorePointCollectionId"`
	RestorePoints            []RestorePointResponse                          `pulumi:"restorePoints"`
	Source                   *RestorePointCollectionSourcePropertiesResponse `pulumi:"source"`
	Tags                     map[string]string                               `pulumi:"tags"`
	Type                     string                                          `pulumi:"type"`
}
