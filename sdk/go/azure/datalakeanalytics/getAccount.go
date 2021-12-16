


package datalakeanalytics

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:datalakeanalytics:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountResult struct {
	AccountId                    string                                    `pulumi:"accountId"`
	ComputePolicies              []ComputePolicyResponse                   `pulumi:"computePolicies"`
	CreationTime                 string                                    `pulumi:"creationTime"`
	CurrentTier                  string                                    `pulumi:"currentTier"`
	DataLakeStoreAccounts        []DataLakeStoreAccountInformationResponse `pulumi:"dataLakeStoreAccounts"`
	DebugDataAccessLevel         string                                    `pulumi:"debugDataAccessLevel"`
	DefaultDataLakeStoreAccount  string                                    `pulumi:"defaultDataLakeStoreAccount"`
	Endpoint                     string                                    `pulumi:"endpoint"`
	FirewallAllowAzureIps        *string                                   `pulumi:"firewallAllowAzureIps"`
	FirewallRules                []FirewallRuleResponse                    `pulumi:"firewallRules"`
	FirewallState                *string                                   `pulumi:"firewallState"`
	HiveMetastores               []HiveMetastoreResponse                   `pulumi:"hiveMetastores"`
	Id                           string                                    `pulumi:"id"`
	LastModifiedTime             string                                    `pulumi:"lastModifiedTime"`
	Location                     string                                    `pulumi:"location"`
	MaxActiveJobCountPerUser     int                                       `pulumi:"maxActiveJobCountPerUser"`
	MaxDegreeOfParallelism       *int                                      `pulumi:"maxDegreeOfParallelism"`
	MaxDegreeOfParallelismPerJob *int                                      `pulumi:"maxDegreeOfParallelismPerJob"`
	MaxJobCount                  *int                                      `pulumi:"maxJobCount"`
	MaxJobRunningTimeInMin       int                                       `pulumi:"maxJobRunningTimeInMin"`
	MaxQueuedJobCountPerUser     int                                       `pulumi:"maxQueuedJobCountPerUser"`
	MinPriorityPerJob            int                                       `pulumi:"minPriorityPerJob"`
	Name                         string                                    `pulumi:"name"`
	NewTier                      *string                                   `pulumi:"newTier"`
	ProvisioningState            string                                    `pulumi:"provisioningState"`
	PublicDataLakeStoreAccounts  []DataLakeStoreAccountInformationResponse `pulumi:"publicDataLakeStoreAccounts"`
	QueryStoreRetention          *int                                      `pulumi:"queryStoreRetention"`
	State                        string                                    `pulumi:"state"`
	StorageAccounts              []StorageAccountInformationResponse       `pulumi:"storageAccounts"`
	SystemMaxDegreeOfParallelism int                                       `pulumi:"systemMaxDegreeOfParallelism"`
	SystemMaxJobCount            int                                       `pulumi:"systemMaxJobCount"`
	Tags                         map[string]string                         `pulumi:"tags"`
	Type                         string                                    `pulumi:"type"`
	VirtualNetworkRules          []VirtualNetworkRuleResponse              `pulumi:"virtualNetworkRules"`
}


func (val *LookupAccountResult) Defaults() *LookupAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxDegreeOfParallelism) {
		maxDegreeOfParallelism_ := 30
		tmp.MaxDegreeOfParallelism = &maxDegreeOfParallelism_
	}
	if isZero(tmp.MaxJobCount) {
		maxJobCount_ := 3
		tmp.MaxJobCount = &maxJobCount_
	}
	if isZero(tmp.QueryStoreRetention) {
		queryStoreRetention_ := 30
		tmp.QueryStoreRetention = &queryStoreRetention_
	}
	return &tmp
}
