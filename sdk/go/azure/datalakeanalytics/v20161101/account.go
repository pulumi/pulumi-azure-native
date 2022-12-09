


package v20161101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Account struct {
	pulumi.CustomResourceState

	AccountId                    pulumi.StringOutput                                `pulumi:"accountId"`
	ComputePolicies              ComputePolicyResponseArrayOutput                   `pulumi:"computePolicies"`
	CreationTime                 pulumi.StringOutput                                `pulumi:"creationTime"`
	CurrentTier                  pulumi.StringOutput                                `pulumi:"currentTier"`
	DataLakeStoreAccounts        DataLakeStoreAccountInformationResponseArrayOutput `pulumi:"dataLakeStoreAccounts"`
	DebugDataAccessLevel         pulumi.StringOutput                                `pulumi:"debugDataAccessLevel"`
	DefaultDataLakeStoreAccount  pulumi.StringOutput                                `pulumi:"defaultDataLakeStoreAccount"`
	Endpoint                     pulumi.StringOutput                                `pulumi:"endpoint"`
	FirewallAllowAzureIps        pulumi.StringPtrOutput                             `pulumi:"firewallAllowAzureIps"`
	FirewallRules                FirewallRuleResponseArrayOutput                    `pulumi:"firewallRules"`
	FirewallState                pulumi.StringPtrOutput                             `pulumi:"firewallState"`
	HiveMetastores               HiveMetastoreResponseArrayOutput                   `pulumi:"hiveMetastores"`
	LastModifiedTime             pulumi.StringOutput                                `pulumi:"lastModifiedTime"`
	Location                     pulumi.StringOutput                                `pulumi:"location"`
	MaxActiveJobCountPerUser     pulumi.IntOutput                                   `pulumi:"maxActiveJobCountPerUser"`
	MaxDegreeOfParallelism       pulumi.IntPtrOutput                                `pulumi:"maxDegreeOfParallelism"`
	MaxDegreeOfParallelismPerJob pulumi.IntPtrOutput                                `pulumi:"maxDegreeOfParallelismPerJob"`
	MaxJobCount                  pulumi.IntPtrOutput                                `pulumi:"maxJobCount"`
	MaxJobRunningTimeInMin       pulumi.IntOutput                                   `pulumi:"maxJobRunningTimeInMin"`
	MaxQueuedJobCountPerUser     pulumi.IntOutput                                   `pulumi:"maxQueuedJobCountPerUser"`
	MinPriorityPerJob            pulumi.IntOutput                                   `pulumi:"minPriorityPerJob"`
	Name                         pulumi.StringOutput                                `pulumi:"name"`
	NewTier                      pulumi.StringPtrOutput                             `pulumi:"newTier"`
	ProvisioningState            pulumi.StringOutput                                `pulumi:"provisioningState"`
	PublicDataLakeStoreAccounts  DataLakeStoreAccountInformationResponseArrayOutput `pulumi:"publicDataLakeStoreAccounts"`
	QueryStoreRetention          pulumi.IntPtrOutput                                `pulumi:"queryStoreRetention"`
	State                        pulumi.StringOutput                                `pulumi:"state"`
	StorageAccounts              StorageAccountInformationResponseArrayOutput       `pulumi:"storageAccounts"`
	SystemMaxDegreeOfParallelism pulumi.IntOutput                                   `pulumi:"systemMaxDegreeOfParallelism"`
	SystemMaxJobCount            pulumi.IntOutput                                   `pulumi:"systemMaxJobCount"`
	Tags                         pulumi.StringMapOutput                             `pulumi:"tags"`
	Type                         pulumi.StringOutput                                `pulumi:"type"`
	VirtualNetworkRules          VirtualNetworkRuleResponseArrayOutput              `pulumi:"virtualNetworkRules"`
}


func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataLakeStoreAccounts == nil {
		return nil, errors.New("invalid value for required argument 'DataLakeStoreAccounts'")
	}
	if args.DefaultDataLakeStoreAccount == nil {
		return nil, errors.New("invalid value for required argument 'DefaultDataLakeStoreAccount'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.FirewallAllowAzureIps) {
		args.FirewallAllowAzureIps = FirewallAllowAzureIpsState("Disabled")
	}
	if isZero(args.FirewallState) {
		args.FirewallState = FirewallState("Disabled")
	}
	if isZero(args.MaxDegreeOfParallelism) {
		args.MaxDegreeOfParallelism = pulumi.IntPtr(30)
	}
	if isZero(args.MaxDegreeOfParallelismPerJob) {
		args.MaxDegreeOfParallelismPerJob = pulumi.IntPtr(32)
	}
	if isZero(args.MaxJobCount) {
		args.MaxJobCount = pulumi.IntPtr(3)
	}
	if isZero(args.NewTier) {
		args.NewTier = TierType("Consumption")
	}
	if isZero(args.QueryStoreRetention) {
		args.QueryStoreRetention = pulumi.IntPtr(30)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datalakeanalytics:Account"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20151001preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20191101preview:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-native:datalakeanalytics/v20161101:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:datalakeanalytics/v20161101:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type accountState struct {
}

type AccountState struct {
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	AccountName                  *string                                    `pulumi:"accountName"`
	ComputePolicies              []CreateComputePolicyWithAccountParameters `pulumi:"computePolicies"`
	DataLakeStoreAccounts        []AddDataLakeStoreWithAccountParameters    `pulumi:"dataLakeStoreAccounts"`
	DefaultDataLakeStoreAccount  string                                     `pulumi:"defaultDataLakeStoreAccount"`
	FirewallAllowAzureIps        *FirewallAllowAzureIpsState                `pulumi:"firewallAllowAzureIps"`
	FirewallRules                []CreateFirewallRuleWithAccountParameters  `pulumi:"firewallRules"`
	FirewallState                *FirewallState                             `pulumi:"firewallState"`
	Location                     *string                                    `pulumi:"location"`
	MaxDegreeOfParallelism       *int                                       `pulumi:"maxDegreeOfParallelism"`
	MaxDegreeOfParallelismPerJob *int                                       `pulumi:"maxDegreeOfParallelismPerJob"`
	MaxJobCount                  *int                                       `pulumi:"maxJobCount"`
	MinPriorityPerJob            *int                                       `pulumi:"minPriorityPerJob"`
	NewTier                      *TierType                                  `pulumi:"newTier"`
	QueryStoreRetention          *int                                       `pulumi:"queryStoreRetention"`
	ResourceGroupName            string                                     `pulumi:"resourceGroupName"`
	StorageAccounts              []AddStorageAccountWithAccountParameters   `pulumi:"storageAccounts"`
	Tags                         map[string]string                          `pulumi:"tags"`
}


type AccountArgs struct {
	AccountName                  pulumi.StringPtrInput
	ComputePolicies              CreateComputePolicyWithAccountParametersArrayInput
	DataLakeStoreAccounts        AddDataLakeStoreWithAccountParametersArrayInput
	DefaultDataLakeStoreAccount  pulumi.StringInput
	FirewallAllowAzureIps        FirewallAllowAzureIpsStatePtrInput
	FirewallRules                CreateFirewallRuleWithAccountParametersArrayInput
	FirewallState                FirewallStatePtrInput
	Location                     pulumi.StringPtrInput
	MaxDegreeOfParallelism       pulumi.IntPtrInput
	MaxDegreeOfParallelismPerJob pulumi.IntPtrInput
	MaxJobCount                  pulumi.IntPtrInput
	MinPriorityPerJob            pulumi.IntPtrInput
	NewTier                      TierTypePtrInput
	QueryStoreRetention          pulumi.IntPtrInput
	ResourceGroupName            pulumi.StringInput
	StorageAccounts              AddStorageAccountWithAccountParametersArrayInput
	Tags                         pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func (o AccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o AccountOutput) ComputePolicies() ComputePolicyResponseArrayOutput {
	return o.ApplyT(func(v *Account) ComputePolicyResponseArrayOutput { return v.ComputePolicies }).(ComputePolicyResponseArrayOutput)
}

func (o AccountOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o AccountOutput) CurrentTier() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CurrentTier }).(pulumi.StringOutput)
}

func (o AccountOutput) DataLakeStoreAccounts() DataLakeStoreAccountInformationResponseArrayOutput {
	return o.ApplyT(func(v *Account) DataLakeStoreAccountInformationResponseArrayOutput { return v.DataLakeStoreAccounts }).(DataLakeStoreAccountInformationResponseArrayOutput)
}

func (o AccountOutput) DebugDataAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.DebugDataAccessLevel }).(pulumi.StringOutput)
}

func (o AccountOutput) DefaultDataLakeStoreAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.DefaultDataLakeStoreAccount }).(pulumi.StringOutput)
}

func (o AccountOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

func (o AccountOutput) FirewallAllowAzureIps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.FirewallAllowAzureIps }).(pulumi.StringPtrOutput)
}

func (o AccountOutput) FirewallRules() FirewallRuleResponseArrayOutput {
	return o.ApplyT(func(v *Account) FirewallRuleResponseArrayOutput { return v.FirewallRules }).(FirewallRuleResponseArrayOutput)
}

func (o AccountOutput) FirewallState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.FirewallState }).(pulumi.StringPtrOutput)
}

func (o AccountOutput) HiveMetastores() HiveMetastoreResponseArrayOutput {
	return o.ApplyT(func(v *Account) HiveMetastoreResponseArrayOutput { return v.HiveMetastores }).(HiveMetastoreResponseArrayOutput)
}

func (o AccountOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o AccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AccountOutput) MaxActiveJobCountPerUser() pulumi.IntOutput {
	return o.ApplyT(func(v *Account) pulumi.IntOutput { return v.MaxActiveJobCountPerUser }).(pulumi.IntOutput)
}

func (o AccountOutput) MaxDegreeOfParallelism() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.IntPtrOutput { return v.MaxDegreeOfParallelism }).(pulumi.IntPtrOutput)
}

func (o AccountOutput) MaxDegreeOfParallelismPerJob() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.IntPtrOutput { return v.MaxDegreeOfParallelismPerJob }).(pulumi.IntPtrOutput)
}

func (o AccountOutput) MaxJobCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.IntPtrOutput { return v.MaxJobCount }).(pulumi.IntPtrOutput)
}

func (o AccountOutput) MaxJobRunningTimeInMin() pulumi.IntOutput {
	return o.ApplyT(func(v *Account) pulumi.IntOutput { return v.MaxJobRunningTimeInMin }).(pulumi.IntOutput)
}

func (o AccountOutput) MaxQueuedJobCountPerUser() pulumi.IntOutput {
	return o.ApplyT(func(v *Account) pulumi.IntOutput { return v.MaxQueuedJobCountPerUser }).(pulumi.IntOutput)
}

func (o AccountOutput) MinPriorityPerJob() pulumi.IntOutput {
	return o.ApplyT(func(v *Account) pulumi.IntOutput { return v.MinPriorityPerJob }).(pulumi.IntOutput)
}

func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AccountOutput) NewTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.NewTier }).(pulumi.StringPtrOutput)
}

func (o AccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AccountOutput) PublicDataLakeStoreAccounts() DataLakeStoreAccountInformationResponseArrayOutput {
	return o.ApplyT(func(v *Account) DataLakeStoreAccountInformationResponseArrayOutput {
		return v.PublicDataLakeStoreAccounts
	}).(DataLakeStoreAccountInformationResponseArrayOutput)
}

func (o AccountOutput) QueryStoreRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.IntPtrOutput { return v.QueryStoreRetention }).(pulumi.IntPtrOutput)
}

func (o AccountOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o AccountOutput) StorageAccounts() StorageAccountInformationResponseArrayOutput {
	return o.ApplyT(func(v *Account) StorageAccountInformationResponseArrayOutput { return v.StorageAccounts }).(StorageAccountInformationResponseArrayOutput)
}

func (o AccountOutput) SystemMaxDegreeOfParallelism() pulumi.IntOutput {
	return o.ApplyT(func(v *Account) pulumi.IntOutput { return v.SystemMaxDegreeOfParallelism }).(pulumi.IntOutput)
}

func (o AccountOutput) SystemMaxJobCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Account) pulumi.IntOutput { return v.SystemMaxJobCount }).(pulumi.IntOutput)
}

func (o AccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Account) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AccountOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *Account) VirtualNetworkRuleResponseArrayOutput { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
