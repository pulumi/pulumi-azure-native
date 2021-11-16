


package v20190401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBatchAccount(ctx *pulumi.Context, args *LookupBatchAccountArgs, opts ...pulumi.InvokeOption) (*LookupBatchAccountResult, error) {
	var rv LookupBatchAccountResult
	err := ctx.Invoke("azure-native:batch/v20190401:getBatchAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBatchAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBatchAccountResult struct {
	AccountEndpoint                       string                                  `pulumi:"accountEndpoint"`
	ActiveJobAndJobScheduleQuota          int                                     `pulumi:"activeJobAndJobScheduleQuota"`
	AutoStorage                           AutoStoragePropertiesResponse           `pulumi:"autoStorage"`
	DedicatedCoreQuota                    int                                     `pulumi:"dedicatedCoreQuota"`
	DedicatedCoreQuotaPerVMFamily         []VirtualMachineFamilyCoreQuotaResponse `pulumi:"dedicatedCoreQuotaPerVMFamily"`
	DedicatedCoreQuotaPerVMFamilyEnforced bool                                    `pulumi:"dedicatedCoreQuotaPerVMFamilyEnforced"`
	Id                                    string                                  `pulumi:"id"`
	KeyVaultReference                     KeyVaultReferenceResponse               `pulumi:"keyVaultReference"`
	Location                              string                                  `pulumi:"location"`
	LowPriorityCoreQuota                  int                                     `pulumi:"lowPriorityCoreQuota"`
	Name                                  string                                  `pulumi:"name"`
	PoolAllocationMode                    string                                  `pulumi:"poolAllocationMode"`
	PoolQuota                             int                                     `pulumi:"poolQuota"`
	ProvisioningState                     string                                  `pulumi:"provisioningState"`
	Tags                                  map[string]string                       `pulumi:"tags"`
	Type                                  string                                  `pulumi:"type"`
}
