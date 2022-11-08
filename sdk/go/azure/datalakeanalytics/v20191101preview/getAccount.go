


package v20191101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:datalakeanalytics/v20191101preview:getAccount", args, &rv, opts...)
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
	AccountId                       string                                    `pulumi:"accountId"`
	ComputePolicies                 []ComputePolicyResponse                   `pulumi:"computePolicies"`
	CreationTime                    string                                    `pulumi:"creationTime"`
	CurrentTier                     string                                    `pulumi:"currentTier"`
	DataLakeStoreAccounts           []DataLakeStoreAccountInformationResponse `pulumi:"dataLakeStoreAccounts"`
	DebugDataAccessLevel            string                                    `pulumi:"debugDataAccessLevel"`
	DefaultDataLakeStoreAccount     string                                    `pulumi:"defaultDataLakeStoreAccount"`
	DefaultDataLakeStoreAccountType string                                    `pulumi:"defaultDataLakeStoreAccountType"`
	Endpoint                        string                                    `pulumi:"endpoint"`
	FirewallAllowAzureIps           *string                                   `pulumi:"firewallAllowAzureIps"`
	FirewallRules                   []FirewallRuleResponse                    `pulumi:"firewallRules"`
	FirewallState                   *string                                   `pulumi:"firewallState"`
	HiveMetastores                  []HiveMetastoreResponse                   `pulumi:"hiveMetastores"`
	Id                              string                                    `pulumi:"id"`
	LastModifiedTime                string                                    `pulumi:"lastModifiedTime"`
	Location                        string                                    `pulumi:"location"`
	MaxActiveJobCountPerUser        int                                       `pulumi:"maxActiveJobCountPerUser"`
	MaxDegreeOfParallelism          *int                                      `pulumi:"maxDegreeOfParallelism"`
	MaxDegreeOfParallelismPerJob    *int                                      `pulumi:"maxDegreeOfParallelismPerJob"`
	MaxJobCount                     *int                                      `pulumi:"maxJobCount"`
	MaxJobRunningTimeInMin          int                                       `pulumi:"maxJobRunningTimeInMin"`
	MaxQueuedJobCountPerUser        int                                       `pulumi:"maxQueuedJobCountPerUser"`
	MinPriorityPerJob               int                                       `pulumi:"minPriorityPerJob"`
	Name                            string                                    `pulumi:"name"`
	NewTier                         *string                                   `pulumi:"newTier"`
	ProvisioningState               string                                    `pulumi:"provisioningState"`
	PublicDataLakeStoreAccounts     []DataLakeStoreAccountInformationResponse `pulumi:"publicDataLakeStoreAccounts"`
	QueryStoreRetention             *int                                      `pulumi:"queryStoreRetention"`
	State                           string                                    `pulumi:"state"`
	StorageAccounts                 []StorageAccountInformationResponse       `pulumi:"storageAccounts"`
	SystemMaxDegreeOfParallelism    int                                       `pulumi:"systemMaxDegreeOfParallelism"`
	SystemMaxJobCount               int                                       `pulumi:"systemMaxJobCount"`
	Tags                            map[string]string                         `pulumi:"tags"`
	Type                            string                                    `pulumi:"type"`
	VirtualNetworkRules             []VirtualNetworkRuleResponse              `pulumi:"virtualNetworkRules"`
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

func LookupAccountOutput(ctx *pulumi.Context, args LookupAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountResult, error) {
			args := v.(LookupAccountArgs)
			r, err := LookupAccount(ctx, &args, opts...)
			var s LookupAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountResultOutput)
}

type LookupAccountOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}


type LookupAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountResult)(nil)).Elem()
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutput() LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutputWithContext(ctx context.Context) LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) ComputePolicies() ComputePolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []ComputePolicyResponse { return v.ComputePolicies }).(ComputePolicyResponseArrayOutput)
}

func (o LookupAccountResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) CurrentTier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.CurrentTier }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) DataLakeStoreAccounts() DataLakeStoreAccountInformationResponseArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []DataLakeStoreAccountInformationResponse { return v.DataLakeStoreAccounts }).(DataLakeStoreAccountInformationResponseArrayOutput)
}

func (o LookupAccountResultOutput) DebugDataAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.DebugDataAccessLevel }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) DefaultDataLakeStoreAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.DefaultDataLakeStoreAccount }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) DefaultDataLakeStoreAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.DefaultDataLakeStoreAccountType }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) FirewallAllowAzureIps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.FirewallAllowAzureIps }).(pulumi.StringPtrOutput)
}

func (o LookupAccountResultOutput) FirewallRules() FirewallRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []FirewallRuleResponse { return v.FirewallRules }).(FirewallRuleResponseArrayOutput)
}

func (o LookupAccountResultOutput) FirewallState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.FirewallState }).(pulumi.StringPtrOutput)
}

func (o LookupAccountResultOutput) HiveMetastores() HiveMetastoreResponseArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []HiveMetastoreResponse { return v.HiveMetastores }).(HiveMetastoreResponseArrayOutput)
}

func (o LookupAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) MaxActiveJobCountPerUser() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAccountResult) int { return v.MaxActiveJobCountPerUser }).(pulumi.IntOutput)
}

func (o LookupAccountResultOutput) MaxDegreeOfParallelism() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *int { return v.MaxDegreeOfParallelism }).(pulumi.IntPtrOutput)
}

func (o LookupAccountResultOutput) MaxDegreeOfParallelismPerJob() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *int { return v.MaxDegreeOfParallelismPerJob }).(pulumi.IntPtrOutput)
}

func (o LookupAccountResultOutput) MaxJobCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *int { return v.MaxJobCount }).(pulumi.IntPtrOutput)
}

func (o LookupAccountResultOutput) MaxJobRunningTimeInMin() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAccountResult) int { return v.MaxJobRunningTimeInMin }).(pulumi.IntOutput)
}

func (o LookupAccountResultOutput) MaxQueuedJobCountPerUser() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAccountResult) int { return v.MaxQueuedJobCountPerUser }).(pulumi.IntOutput)
}

func (o LookupAccountResultOutput) MinPriorityPerJob() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAccountResult) int { return v.MinPriorityPerJob }).(pulumi.IntOutput)
}

func (o LookupAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) NewTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.NewTier }).(pulumi.StringPtrOutput)
}

func (o LookupAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) PublicDataLakeStoreAccounts() DataLakeStoreAccountInformationResponseArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []DataLakeStoreAccountInformationResponse {
		return v.PublicDataLakeStoreAccounts
	}).(DataLakeStoreAccountInformationResponseArrayOutput)
}

func (o LookupAccountResultOutput) QueryStoreRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *int { return v.QueryStoreRetention }).(pulumi.IntPtrOutput)
}

func (o LookupAccountResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) StorageAccounts() StorageAccountInformationResponseArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []StorageAccountInformationResponse { return v.StorageAccounts }).(StorageAccountInformationResponseArrayOutput)
}

func (o LookupAccountResultOutput) SystemMaxDegreeOfParallelism() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAccountResult) int { return v.SystemMaxDegreeOfParallelism }).(pulumi.IntOutput)
}

func (o LookupAccountResultOutput) SystemMaxJobCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAccountResult) int { return v.SystemMaxJobCount }).(pulumi.IntOutput)
}

func (o LookupAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountResultOutput{})
}
