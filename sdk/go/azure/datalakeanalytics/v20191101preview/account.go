


package v20191101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Account struct {
	pulumi.CustomResourceState

	AccountId                       pulumi.StringOutput                                `pulumi:"accountId"`
	ComputePolicies                 ComputePolicyResponseArrayOutput                   `pulumi:"computePolicies"`
	CreationTime                    pulumi.StringOutput                                `pulumi:"creationTime"`
	CurrentTier                     pulumi.StringOutput                                `pulumi:"currentTier"`
	DataLakeStoreAccounts           DataLakeStoreAccountInformationResponseArrayOutput `pulumi:"dataLakeStoreAccounts"`
	DebugDataAccessLevel            pulumi.StringOutput                                `pulumi:"debugDataAccessLevel"`
	DefaultDataLakeStoreAccount     pulumi.StringOutput                                `pulumi:"defaultDataLakeStoreAccount"`
	DefaultDataLakeStoreAccountType pulumi.StringOutput                                `pulumi:"defaultDataLakeStoreAccountType"`
	Endpoint                        pulumi.StringOutput                                `pulumi:"endpoint"`
	FirewallAllowAzureIps           pulumi.StringPtrOutput                             `pulumi:"firewallAllowAzureIps"`
	FirewallRules                   FirewallRuleResponseArrayOutput                    `pulumi:"firewallRules"`
	FirewallState                   pulumi.StringPtrOutput                             `pulumi:"firewallState"`
	HiveMetastores                  HiveMetastoreResponseArrayOutput                   `pulumi:"hiveMetastores"`
	LastModifiedTime                pulumi.StringOutput                                `pulumi:"lastModifiedTime"`
	Location                        pulumi.StringOutput                                `pulumi:"location"`
	MaxActiveJobCountPerUser        pulumi.IntOutput                                   `pulumi:"maxActiveJobCountPerUser"`
	MaxDegreeOfParallelism          pulumi.IntPtrOutput                                `pulumi:"maxDegreeOfParallelism"`
	MaxDegreeOfParallelismPerJob    pulumi.IntPtrOutput                                `pulumi:"maxDegreeOfParallelismPerJob"`
	MaxJobCount                     pulumi.IntPtrOutput                                `pulumi:"maxJobCount"`
	MaxJobRunningTimeInMin          pulumi.IntOutput                                   `pulumi:"maxJobRunningTimeInMin"`
	MaxQueuedJobCountPerUser        pulumi.IntOutput                                   `pulumi:"maxQueuedJobCountPerUser"`
	MinPriorityPerJob               pulumi.IntOutput                                   `pulumi:"minPriorityPerJob"`
	Name                            pulumi.StringOutput                                `pulumi:"name"`
	NewTier                         pulumi.StringPtrOutput                             `pulumi:"newTier"`
	ProvisioningState               pulumi.StringOutput                                `pulumi:"provisioningState"`
	PublicDataLakeStoreAccounts     DataLakeStoreAccountInformationResponseArrayOutput `pulumi:"publicDataLakeStoreAccounts"`
	QueryStoreRetention             pulumi.IntPtrOutput                                `pulumi:"queryStoreRetention"`
	State                           pulumi.StringOutput                                `pulumi:"state"`
	StorageAccounts                 StorageAccountInformationResponseArrayOutput       `pulumi:"storageAccounts"`
	SystemMaxDegreeOfParallelism    pulumi.IntOutput                                   `pulumi:"systemMaxDegreeOfParallelism"`
	SystemMaxJobCount               pulumi.IntOutput                                   `pulumi:"systemMaxJobCount"`
	Tags                            pulumi.StringMapOutput                             `pulumi:"tags"`
	Type                            pulumi.StringOutput                                `pulumi:"type"`
	VirtualNetworkRules             VirtualNetworkRuleResponseArrayOutput              `pulumi:"virtualNetworkRules"`
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
	if args.FirewallAllowAzureIps == nil {
		args.FirewallAllowAzureIps = FirewallAllowAzureIpsState("Disabled")
	}
	if args.FirewallState == nil {
		args.FirewallState = FirewallState("Disabled")
	}
	if args.MaxDegreeOfParallelism == nil {
		args.MaxDegreeOfParallelism = pulumi.IntPtr(30)
	}
	if args.MaxDegreeOfParallelismPerJob == nil {
		args.MaxDegreeOfParallelismPerJob = pulumi.IntPtr(32)
	}
	if args.MaxJobCount == nil {
		args.MaxJobCount = pulumi.IntPtr(3)
	}
	if args.NewTier == nil {
		args.NewTier = TierType("Consumption")
	}
	if args.QueryStoreRetention == nil {
		args.QueryStoreRetention = pulumi.IntPtr(30)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datalakeanalytics/v20191101preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:datalakeanalytics:Account"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20151001preview:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:datalakeanalytics/v20151001preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20161101:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:datalakeanalytics/v20161101:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-native:datalakeanalytics/v20191101preview:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:datalakeanalytics/v20191101preview:Account", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*Account)(nil))
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
