


package v20151001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddDataLakeStoreWithAccountParameters struct {
	Name   string  `pulumi:"name"`
	Suffix *string `pulumi:"suffix"`
}





type AddDataLakeStoreWithAccountParametersInput interface {
	pulumi.Input

	ToAddDataLakeStoreWithAccountParametersOutput() AddDataLakeStoreWithAccountParametersOutput
	ToAddDataLakeStoreWithAccountParametersOutputWithContext(context.Context) AddDataLakeStoreWithAccountParametersOutput
}

type AddDataLakeStoreWithAccountParametersArgs struct {
	Name   pulumi.StringInput    `pulumi:"name"`
	Suffix pulumi.StringPtrInput `pulumi:"suffix"`
}

func (AddDataLakeStoreWithAccountParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddDataLakeStoreWithAccountParameters)(nil)).Elem()
}

func (i AddDataLakeStoreWithAccountParametersArgs) ToAddDataLakeStoreWithAccountParametersOutput() AddDataLakeStoreWithAccountParametersOutput {
	return i.ToAddDataLakeStoreWithAccountParametersOutputWithContext(context.Background())
}

func (i AddDataLakeStoreWithAccountParametersArgs) ToAddDataLakeStoreWithAccountParametersOutputWithContext(ctx context.Context) AddDataLakeStoreWithAccountParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddDataLakeStoreWithAccountParametersOutput)
}





type AddDataLakeStoreWithAccountParametersArrayInput interface {
	pulumi.Input

	ToAddDataLakeStoreWithAccountParametersArrayOutput() AddDataLakeStoreWithAccountParametersArrayOutput
	ToAddDataLakeStoreWithAccountParametersArrayOutputWithContext(context.Context) AddDataLakeStoreWithAccountParametersArrayOutput
}

type AddDataLakeStoreWithAccountParametersArray []AddDataLakeStoreWithAccountParametersInput

func (AddDataLakeStoreWithAccountParametersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AddDataLakeStoreWithAccountParameters)(nil)).Elem()
}

func (i AddDataLakeStoreWithAccountParametersArray) ToAddDataLakeStoreWithAccountParametersArrayOutput() AddDataLakeStoreWithAccountParametersArrayOutput {
	return i.ToAddDataLakeStoreWithAccountParametersArrayOutputWithContext(context.Background())
}

func (i AddDataLakeStoreWithAccountParametersArray) ToAddDataLakeStoreWithAccountParametersArrayOutputWithContext(ctx context.Context) AddDataLakeStoreWithAccountParametersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddDataLakeStoreWithAccountParametersArrayOutput)
}

type AddDataLakeStoreWithAccountParametersOutput struct{ *pulumi.OutputState }

func (AddDataLakeStoreWithAccountParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddDataLakeStoreWithAccountParameters)(nil)).Elem()
}

func (o AddDataLakeStoreWithAccountParametersOutput) ToAddDataLakeStoreWithAccountParametersOutput() AddDataLakeStoreWithAccountParametersOutput {
	return o
}

func (o AddDataLakeStoreWithAccountParametersOutput) ToAddDataLakeStoreWithAccountParametersOutputWithContext(ctx context.Context) AddDataLakeStoreWithAccountParametersOutput {
	return o
}

func (o AddDataLakeStoreWithAccountParametersOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AddDataLakeStoreWithAccountParameters) string { return v.Name }).(pulumi.StringOutput)
}

func (o AddDataLakeStoreWithAccountParametersOutput) Suffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddDataLakeStoreWithAccountParameters) *string { return v.Suffix }).(pulumi.StringPtrOutput)
}

type AddDataLakeStoreWithAccountParametersArrayOutput struct{ *pulumi.OutputState }

func (AddDataLakeStoreWithAccountParametersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AddDataLakeStoreWithAccountParameters)(nil)).Elem()
}

func (o AddDataLakeStoreWithAccountParametersArrayOutput) ToAddDataLakeStoreWithAccountParametersArrayOutput() AddDataLakeStoreWithAccountParametersArrayOutput {
	return o
}

func (o AddDataLakeStoreWithAccountParametersArrayOutput) ToAddDataLakeStoreWithAccountParametersArrayOutputWithContext(ctx context.Context) AddDataLakeStoreWithAccountParametersArrayOutput {
	return o
}

func (o AddDataLakeStoreWithAccountParametersArrayOutput) Index(i pulumi.IntInput) AddDataLakeStoreWithAccountParametersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AddDataLakeStoreWithAccountParameters {
		return vs[0].([]AddDataLakeStoreWithAccountParameters)[vs[1].(int)]
	}).(AddDataLakeStoreWithAccountParametersOutput)
}

type AddStorageAccountWithAccountParameters struct {
	AccessKey string  `pulumi:"accessKey"`
	Name      string  `pulumi:"name"`
	Suffix    *string `pulumi:"suffix"`
}





type AddStorageAccountWithAccountParametersInput interface {
	pulumi.Input

	ToAddStorageAccountWithAccountParametersOutput() AddStorageAccountWithAccountParametersOutput
	ToAddStorageAccountWithAccountParametersOutputWithContext(context.Context) AddStorageAccountWithAccountParametersOutput
}

type AddStorageAccountWithAccountParametersArgs struct {
	AccessKey pulumi.StringInput    `pulumi:"accessKey"`
	Name      pulumi.StringInput    `pulumi:"name"`
	Suffix    pulumi.StringPtrInput `pulumi:"suffix"`
}

func (AddStorageAccountWithAccountParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddStorageAccountWithAccountParameters)(nil)).Elem()
}

func (i AddStorageAccountWithAccountParametersArgs) ToAddStorageAccountWithAccountParametersOutput() AddStorageAccountWithAccountParametersOutput {
	return i.ToAddStorageAccountWithAccountParametersOutputWithContext(context.Background())
}

func (i AddStorageAccountWithAccountParametersArgs) ToAddStorageAccountWithAccountParametersOutputWithContext(ctx context.Context) AddStorageAccountWithAccountParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddStorageAccountWithAccountParametersOutput)
}





type AddStorageAccountWithAccountParametersArrayInput interface {
	pulumi.Input

	ToAddStorageAccountWithAccountParametersArrayOutput() AddStorageAccountWithAccountParametersArrayOutput
	ToAddStorageAccountWithAccountParametersArrayOutputWithContext(context.Context) AddStorageAccountWithAccountParametersArrayOutput
}

type AddStorageAccountWithAccountParametersArray []AddStorageAccountWithAccountParametersInput

func (AddStorageAccountWithAccountParametersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AddStorageAccountWithAccountParameters)(nil)).Elem()
}

func (i AddStorageAccountWithAccountParametersArray) ToAddStorageAccountWithAccountParametersArrayOutput() AddStorageAccountWithAccountParametersArrayOutput {
	return i.ToAddStorageAccountWithAccountParametersArrayOutputWithContext(context.Background())
}

func (i AddStorageAccountWithAccountParametersArray) ToAddStorageAccountWithAccountParametersArrayOutputWithContext(ctx context.Context) AddStorageAccountWithAccountParametersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddStorageAccountWithAccountParametersArrayOutput)
}

type AddStorageAccountWithAccountParametersOutput struct{ *pulumi.OutputState }

func (AddStorageAccountWithAccountParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddStorageAccountWithAccountParameters)(nil)).Elem()
}

func (o AddStorageAccountWithAccountParametersOutput) ToAddStorageAccountWithAccountParametersOutput() AddStorageAccountWithAccountParametersOutput {
	return o
}

func (o AddStorageAccountWithAccountParametersOutput) ToAddStorageAccountWithAccountParametersOutputWithContext(ctx context.Context) AddStorageAccountWithAccountParametersOutput {
	return o
}

func (o AddStorageAccountWithAccountParametersOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v AddStorageAccountWithAccountParameters) string { return v.AccessKey }).(pulumi.StringOutput)
}

func (o AddStorageAccountWithAccountParametersOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AddStorageAccountWithAccountParameters) string { return v.Name }).(pulumi.StringOutput)
}

func (o AddStorageAccountWithAccountParametersOutput) Suffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddStorageAccountWithAccountParameters) *string { return v.Suffix }).(pulumi.StringPtrOutput)
}

type AddStorageAccountWithAccountParametersArrayOutput struct{ *pulumi.OutputState }

func (AddStorageAccountWithAccountParametersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AddStorageAccountWithAccountParameters)(nil)).Elem()
}

func (o AddStorageAccountWithAccountParametersArrayOutput) ToAddStorageAccountWithAccountParametersArrayOutput() AddStorageAccountWithAccountParametersArrayOutput {
	return o
}

func (o AddStorageAccountWithAccountParametersArrayOutput) ToAddStorageAccountWithAccountParametersArrayOutputWithContext(ctx context.Context) AddStorageAccountWithAccountParametersArrayOutput {
	return o
}

func (o AddStorageAccountWithAccountParametersArrayOutput) Index(i pulumi.IntInput) AddStorageAccountWithAccountParametersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AddStorageAccountWithAccountParameters {
		return vs[0].([]AddStorageAccountWithAccountParameters)[vs[1].(int)]
	}).(AddStorageAccountWithAccountParametersOutput)
}

type ComputePolicyResponse struct {
	Id                           string `pulumi:"id"`
	MaxDegreeOfParallelismPerJob int    `pulumi:"maxDegreeOfParallelismPerJob"`
	MinPriorityPerJob            int    `pulumi:"minPriorityPerJob"`
	Name                         string `pulumi:"name"`
	ObjectId                     string `pulumi:"objectId"`
	ObjectType                   string `pulumi:"objectType"`
	Type                         string `pulumi:"type"`
}

type ComputePolicyResponseOutput struct{ *pulumi.OutputState }

func (ComputePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputePolicyResponse)(nil)).Elem()
}

func (o ComputePolicyResponseOutput) ToComputePolicyResponseOutput() ComputePolicyResponseOutput {
	return o
}

func (o ComputePolicyResponseOutput) ToComputePolicyResponseOutputWithContext(ctx context.Context) ComputePolicyResponseOutput {
	return o
}

func (o ComputePolicyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ComputePolicyResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ComputePolicyResponseOutput) MaxDegreeOfParallelismPerJob() pulumi.IntOutput {
	return o.ApplyT(func(v ComputePolicyResponse) int { return v.MaxDegreeOfParallelismPerJob }).(pulumi.IntOutput)
}

func (o ComputePolicyResponseOutput) MinPriorityPerJob() pulumi.IntOutput {
	return o.ApplyT(func(v ComputePolicyResponse) int { return v.MinPriorityPerJob }).(pulumi.IntOutput)
}

func (o ComputePolicyResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ComputePolicyResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ComputePolicyResponseOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v ComputePolicyResponse) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o ComputePolicyResponseOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v ComputePolicyResponse) string { return v.ObjectType }).(pulumi.StringOutput)
}

func (o ComputePolicyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ComputePolicyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ComputePolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (ComputePolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ComputePolicyResponse)(nil)).Elem()
}

func (o ComputePolicyResponseArrayOutput) ToComputePolicyResponseArrayOutput() ComputePolicyResponseArrayOutput {
	return o
}

func (o ComputePolicyResponseArrayOutput) ToComputePolicyResponseArrayOutputWithContext(ctx context.Context) ComputePolicyResponseArrayOutput {
	return o
}

func (o ComputePolicyResponseArrayOutput) Index(i pulumi.IntInput) ComputePolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ComputePolicyResponse {
		return vs[0].([]ComputePolicyResponse)[vs[1].(int)]
	}).(ComputePolicyResponseOutput)
}

type CreateComputePolicyWithAccountParameters struct {
	MaxDegreeOfParallelismPerJob *int   `pulumi:"maxDegreeOfParallelismPerJob"`
	MinPriorityPerJob            *int   `pulumi:"minPriorityPerJob"`
	Name                         string `pulumi:"name"`
	ObjectId                     string `pulumi:"objectId"`
	ObjectType                   string `pulumi:"objectType"`
}





type CreateComputePolicyWithAccountParametersInput interface {
	pulumi.Input

	ToCreateComputePolicyWithAccountParametersOutput() CreateComputePolicyWithAccountParametersOutput
	ToCreateComputePolicyWithAccountParametersOutputWithContext(context.Context) CreateComputePolicyWithAccountParametersOutput
}

type CreateComputePolicyWithAccountParametersArgs struct {
	MaxDegreeOfParallelismPerJob pulumi.IntPtrInput `pulumi:"maxDegreeOfParallelismPerJob"`
	MinPriorityPerJob            pulumi.IntPtrInput `pulumi:"minPriorityPerJob"`
	Name                         pulumi.StringInput `pulumi:"name"`
	ObjectId                     pulumi.StringInput `pulumi:"objectId"`
	ObjectType                   pulumi.StringInput `pulumi:"objectType"`
}

func (CreateComputePolicyWithAccountParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateComputePolicyWithAccountParameters)(nil)).Elem()
}

func (i CreateComputePolicyWithAccountParametersArgs) ToCreateComputePolicyWithAccountParametersOutput() CreateComputePolicyWithAccountParametersOutput {
	return i.ToCreateComputePolicyWithAccountParametersOutputWithContext(context.Background())
}

func (i CreateComputePolicyWithAccountParametersArgs) ToCreateComputePolicyWithAccountParametersOutputWithContext(ctx context.Context) CreateComputePolicyWithAccountParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateComputePolicyWithAccountParametersOutput)
}





type CreateComputePolicyWithAccountParametersArrayInput interface {
	pulumi.Input

	ToCreateComputePolicyWithAccountParametersArrayOutput() CreateComputePolicyWithAccountParametersArrayOutput
	ToCreateComputePolicyWithAccountParametersArrayOutputWithContext(context.Context) CreateComputePolicyWithAccountParametersArrayOutput
}

type CreateComputePolicyWithAccountParametersArray []CreateComputePolicyWithAccountParametersInput

func (CreateComputePolicyWithAccountParametersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CreateComputePolicyWithAccountParameters)(nil)).Elem()
}

func (i CreateComputePolicyWithAccountParametersArray) ToCreateComputePolicyWithAccountParametersArrayOutput() CreateComputePolicyWithAccountParametersArrayOutput {
	return i.ToCreateComputePolicyWithAccountParametersArrayOutputWithContext(context.Background())
}

func (i CreateComputePolicyWithAccountParametersArray) ToCreateComputePolicyWithAccountParametersArrayOutputWithContext(ctx context.Context) CreateComputePolicyWithAccountParametersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateComputePolicyWithAccountParametersArrayOutput)
}

type CreateComputePolicyWithAccountParametersOutput struct{ *pulumi.OutputState }

func (CreateComputePolicyWithAccountParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateComputePolicyWithAccountParameters)(nil)).Elem()
}

func (o CreateComputePolicyWithAccountParametersOutput) ToCreateComputePolicyWithAccountParametersOutput() CreateComputePolicyWithAccountParametersOutput {
	return o
}

func (o CreateComputePolicyWithAccountParametersOutput) ToCreateComputePolicyWithAccountParametersOutputWithContext(ctx context.Context) CreateComputePolicyWithAccountParametersOutput {
	return o
}

func (o CreateComputePolicyWithAccountParametersOutput) MaxDegreeOfParallelismPerJob() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CreateComputePolicyWithAccountParameters) *int { return v.MaxDegreeOfParallelismPerJob }).(pulumi.IntPtrOutput)
}

func (o CreateComputePolicyWithAccountParametersOutput) MinPriorityPerJob() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CreateComputePolicyWithAccountParameters) *int { return v.MinPriorityPerJob }).(pulumi.IntPtrOutput)
}

func (o CreateComputePolicyWithAccountParametersOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CreateComputePolicyWithAccountParameters) string { return v.Name }).(pulumi.StringOutput)
}

func (o CreateComputePolicyWithAccountParametersOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v CreateComputePolicyWithAccountParameters) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o CreateComputePolicyWithAccountParametersOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v CreateComputePolicyWithAccountParameters) string { return v.ObjectType }).(pulumi.StringOutput)
}

type CreateComputePolicyWithAccountParametersArrayOutput struct{ *pulumi.OutputState }

func (CreateComputePolicyWithAccountParametersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CreateComputePolicyWithAccountParameters)(nil)).Elem()
}

func (o CreateComputePolicyWithAccountParametersArrayOutput) ToCreateComputePolicyWithAccountParametersArrayOutput() CreateComputePolicyWithAccountParametersArrayOutput {
	return o
}

func (o CreateComputePolicyWithAccountParametersArrayOutput) ToCreateComputePolicyWithAccountParametersArrayOutputWithContext(ctx context.Context) CreateComputePolicyWithAccountParametersArrayOutput {
	return o
}

func (o CreateComputePolicyWithAccountParametersArrayOutput) Index(i pulumi.IntInput) CreateComputePolicyWithAccountParametersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CreateComputePolicyWithAccountParameters {
		return vs[0].([]CreateComputePolicyWithAccountParameters)[vs[1].(int)]
	}).(CreateComputePolicyWithAccountParametersOutput)
}

type CreateFirewallRuleWithAccountParameters struct {
	EndIpAddress   string `pulumi:"endIpAddress"`
	Name           string `pulumi:"name"`
	StartIpAddress string `pulumi:"startIpAddress"`
}





type CreateFirewallRuleWithAccountParametersInput interface {
	pulumi.Input

	ToCreateFirewallRuleWithAccountParametersOutput() CreateFirewallRuleWithAccountParametersOutput
	ToCreateFirewallRuleWithAccountParametersOutputWithContext(context.Context) CreateFirewallRuleWithAccountParametersOutput
}

type CreateFirewallRuleWithAccountParametersArgs struct {
	EndIpAddress   pulumi.StringInput `pulumi:"endIpAddress"`
	Name           pulumi.StringInput `pulumi:"name"`
	StartIpAddress pulumi.StringInput `pulumi:"startIpAddress"`
}

func (CreateFirewallRuleWithAccountParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateFirewallRuleWithAccountParameters)(nil)).Elem()
}

func (i CreateFirewallRuleWithAccountParametersArgs) ToCreateFirewallRuleWithAccountParametersOutput() CreateFirewallRuleWithAccountParametersOutput {
	return i.ToCreateFirewallRuleWithAccountParametersOutputWithContext(context.Background())
}

func (i CreateFirewallRuleWithAccountParametersArgs) ToCreateFirewallRuleWithAccountParametersOutputWithContext(ctx context.Context) CreateFirewallRuleWithAccountParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateFirewallRuleWithAccountParametersOutput)
}





type CreateFirewallRuleWithAccountParametersArrayInput interface {
	pulumi.Input

	ToCreateFirewallRuleWithAccountParametersArrayOutput() CreateFirewallRuleWithAccountParametersArrayOutput
	ToCreateFirewallRuleWithAccountParametersArrayOutputWithContext(context.Context) CreateFirewallRuleWithAccountParametersArrayOutput
}

type CreateFirewallRuleWithAccountParametersArray []CreateFirewallRuleWithAccountParametersInput

func (CreateFirewallRuleWithAccountParametersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CreateFirewallRuleWithAccountParameters)(nil)).Elem()
}

func (i CreateFirewallRuleWithAccountParametersArray) ToCreateFirewallRuleWithAccountParametersArrayOutput() CreateFirewallRuleWithAccountParametersArrayOutput {
	return i.ToCreateFirewallRuleWithAccountParametersArrayOutputWithContext(context.Background())
}

func (i CreateFirewallRuleWithAccountParametersArray) ToCreateFirewallRuleWithAccountParametersArrayOutputWithContext(ctx context.Context) CreateFirewallRuleWithAccountParametersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateFirewallRuleWithAccountParametersArrayOutput)
}

type CreateFirewallRuleWithAccountParametersOutput struct{ *pulumi.OutputState }

func (CreateFirewallRuleWithAccountParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateFirewallRuleWithAccountParameters)(nil)).Elem()
}

func (o CreateFirewallRuleWithAccountParametersOutput) ToCreateFirewallRuleWithAccountParametersOutput() CreateFirewallRuleWithAccountParametersOutput {
	return o
}

func (o CreateFirewallRuleWithAccountParametersOutput) ToCreateFirewallRuleWithAccountParametersOutputWithContext(ctx context.Context) CreateFirewallRuleWithAccountParametersOutput {
	return o
}

func (o CreateFirewallRuleWithAccountParametersOutput) EndIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v CreateFirewallRuleWithAccountParameters) string { return v.EndIpAddress }).(pulumi.StringOutput)
}

func (o CreateFirewallRuleWithAccountParametersOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CreateFirewallRuleWithAccountParameters) string { return v.Name }).(pulumi.StringOutput)
}

func (o CreateFirewallRuleWithAccountParametersOutput) StartIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v CreateFirewallRuleWithAccountParameters) string { return v.StartIpAddress }).(pulumi.StringOutput)
}

type CreateFirewallRuleWithAccountParametersArrayOutput struct{ *pulumi.OutputState }

func (CreateFirewallRuleWithAccountParametersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CreateFirewallRuleWithAccountParameters)(nil)).Elem()
}

func (o CreateFirewallRuleWithAccountParametersArrayOutput) ToCreateFirewallRuleWithAccountParametersArrayOutput() CreateFirewallRuleWithAccountParametersArrayOutput {
	return o
}

func (o CreateFirewallRuleWithAccountParametersArrayOutput) ToCreateFirewallRuleWithAccountParametersArrayOutputWithContext(ctx context.Context) CreateFirewallRuleWithAccountParametersArrayOutput {
	return o
}

func (o CreateFirewallRuleWithAccountParametersArrayOutput) Index(i pulumi.IntInput) CreateFirewallRuleWithAccountParametersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CreateFirewallRuleWithAccountParameters {
		return vs[0].([]CreateFirewallRuleWithAccountParameters)[vs[1].(int)]
	}).(CreateFirewallRuleWithAccountParametersOutput)
}

type DataLakeStoreAccountInformationResponse struct {
	Id     string `pulumi:"id"`
	Name   string `pulumi:"name"`
	Suffix string `pulumi:"suffix"`
	Type   string `pulumi:"type"`
}

type DataLakeStoreAccountInformationResponseOutput struct{ *pulumi.OutputState }

func (DataLakeStoreAccountInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeStoreAccountInformationResponse)(nil)).Elem()
}

func (o DataLakeStoreAccountInformationResponseOutput) ToDataLakeStoreAccountInformationResponseOutput() DataLakeStoreAccountInformationResponseOutput {
	return o
}

func (o DataLakeStoreAccountInformationResponseOutput) ToDataLakeStoreAccountInformationResponseOutputWithContext(ctx context.Context) DataLakeStoreAccountInformationResponseOutput {
	return o
}

func (o DataLakeStoreAccountInformationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeStoreAccountInformationResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o DataLakeStoreAccountInformationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeStoreAccountInformationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DataLakeStoreAccountInformationResponseOutput) Suffix() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeStoreAccountInformationResponse) string { return v.Suffix }).(pulumi.StringOutput)
}

func (o DataLakeStoreAccountInformationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeStoreAccountInformationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type DataLakeStoreAccountInformationResponseArrayOutput struct{ *pulumi.OutputState }

func (DataLakeStoreAccountInformationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataLakeStoreAccountInformationResponse)(nil)).Elem()
}

func (o DataLakeStoreAccountInformationResponseArrayOutput) ToDataLakeStoreAccountInformationResponseArrayOutput() DataLakeStoreAccountInformationResponseArrayOutput {
	return o
}

func (o DataLakeStoreAccountInformationResponseArrayOutput) ToDataLakeStoreAccountInformationResponseArrayOutputWithContext(ctx context.Context) DataLakeStoreAccountInformationResponseArrayOutput {
	return o
}

func (o DataLakeStoreAccountInformationResponseArrayOutput) Index(i pulumi.IntInput) DataLakeStoreAccountInformationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataLakeStoreAccountInformationResponse {
		return vs[0].([]DataLakeStoreAccountInformationResponse)[vs[1].(int)]
	}).(DataLakeStoreAccountInformationResponseOutput)
}

type FirewallRuleResponse struct {
	EndIpAddress   string `pulumi:"endIpAddress"`
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	StartIpAddress string `pulumi:"startIpAddress"`
	Type           string `pulumi:"type"`
}

type FirewallRuleResponseOutput struct{ *pulumi.OutputState }

func (FirewallRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRuleResponse)(nil)).Elem()
}

func (o FirewallRuleResponseOutput) ToFirewallRuleResponseOutput() FirewallRuleResponseOutput {
	return o
}

func (o FirewallRuleResponseOutput) ToFirewallRuleResponseOutputWithContext(ctx context.Context) FirewallRuleResponseOutput {
	return o
}

func (o FirewallRuleResponseOutput) EndIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v FirewallRuleResponse) string { return v.EndIpAddress }).(pulumi.StringOutput)
}

func (o FirewallRuleResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FirewallRuleResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o FirewallRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v FirewallRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallRuleResponseOutput) StartIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v FirewallRuleResponse) string { return v.StartIpAddress }).(pulumi.StringOutput)
}

func (o FirewallRuleResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FirewallRuleResponse) string { return v.Type }).(pulumi.StringOutput)
}

type FirewallRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (FirewallRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallRuleResponse)(nil)).Elem()
}

func (o FirewallRuleResponseArrayOutput) ToFirewallRuleResponseArrayOutput() FirewallRuleResponseArrayOutput {
	return o
}

func (o FirewallRuleResponseArrayOutput) ToFirewallRuleResponseArrayOutputWithContext(ctx context.Context) FirewallRuleResponseArrayOutput {
	return o
}

func (o FirewallRuleResponseArrayOutput) Index(i pulumi.IntInput) FirewallRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallRuleResponse {
		return vs[0].([]FirewallRuleResponse)[vs[1].(int)]
	}).(FirewallRuleResponseOutput)
}

type HiveMetastoreResponse struct {
	DatabaseName                    string `pulumi:"databaseName"`
	Id                              string `pulumi:"id"`
	Name                            string `pulumi:"name"`
	NestedResourceProvisioningState string `pulumi:"nestedResourceProvisioningState"`
	Password                        string `pulumi:"password"`
	RuntimeVersion                  string `pulumi:"runtimeVersion"`
	ServerUri                       string `pulumi:"serverUri"`
	Type                            string `pulumi:"type"`
	UserName                        string `pulumi:"userName"`
}

type HiveMetastoreResponseOutput struct{ *pulumi.OutputState }

func (HiveMetastoreResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HiveMetastoreResponse)(nil)).Elem()
}

func (o HiveMetastoreResponseOutput) ToHiveMetastoreResponseOutput() HiveMetastoreResponseOutput {
	return o
}

func (o HiveMetastoreResponseOutput) ToHiveMetastoreResponseOutputWithContext(ctx context.Context) HiveMetastoreResponseOutput {
	return o
}

func (o HiveMetastoreResponseOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v HiveMetastoreResponse) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o HiveMetastoreResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v HiveMetastoreResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o HiveMetastoreResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HiveMetastoreResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o HiveMetastoreResponseOutput) NestedResourceProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v HiveMetastoreResponse) string { return v.NestedResourceProvisioningState }).(pulumi.StringOutput)
}

func (o HiveMetastoreResponseOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v HiveMetastoreResponse) string { return v.Password }).(pulumi.StringOutput)
}

func (o HiveMetastoreResponseOutput) RuntimeVersion() pulumi.StringOutput {
	return o.ApplyT(func(v HiveMetastoreResponse) string { return v.RuntimeVersion }).(pulumi.StringOutput)
}

func (o HiveMetastoreResponseOutput) ServerUri() pulumi.StringOutput {
	return o.ApplyT(func(v HiveMetastoreResponse) string { return v.ServerUri }).(pulumi.StringOutput)
}

func (o HiveMetastoreResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v HiveMetastoreResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o HiveMetastoreResponseOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v HiveMetastoreResponse) string { return v.UserName }).(pulumi.StringOutput)
}

type HiveMetastoreResponseArrayOutput struct{ *pulumi.OutputState }

func (HiveMetastoreResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HiveMetastoreResponse)(nil)).Elem()
}

func (o HiveMetastoreResponseArrayOutput) ToHiveMetastoreResponseArrayOutput() HiveMetastoreResponseArrayOutput {
	return o
}

func (o HiveMetastoreResponseArrayOutput) ToHiveMetastoreResponseArrayOutputWithContext(ctx context.Context) HiveMetastoreResponseArrayOutput {
	return o
}

func (o HiveMetastoreResponseArrayOutput) Index(i pulumi.IntInput) HiveMetastoreResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HiveMetastoreResponse {
		return vs[0].([]HiveMetastoreResponse)[vs[1].(int)]
	}).(HiveMetastoreResponseOutput)
}

type SasTokenInformationResponse struct {
	AccessToken string `pulumi:"accessToken"`
}

type StorageAccountInformationResponse struct {
	Id     string `pulumi:"id"`
	Name   string `pulumi:"name"`
	Suffix string `pulumi:"suffix"`
	Type   string `pulumi:"type"`
}

type StorageAccountInformationResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountInformationResponse)(nil)).Elem()
}

func (o StorageAccountInformationResponseOutput) ToStorageAccountInformationResponseOutput() StorageAccountInformationResponseOutput {
	return o
}

func (o StorageAccountInformationResponseOutput) ToStorageAccountInformationResponseOutputWithContext(ctx context.Context) StorageAccountInformationResponseOutput {
	return o
}

func (o StorageAccountInformationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountInformationResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o StorageAccountInformationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountInformationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o StorageAccountInformationResponseOutput) Suffix() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountInformationResponse) string { return v.Suffix }).(pulumi.StringOutput)
}

func (o StorageAccountInformationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountInformationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type StorageAccountInformationResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountInformationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountInformationResponse)(nil)).Elem()
}

func (o StorageAccountInformationResponseArrayOutput) ToStorageAccountInformationResponseArrayOutput() StorageAccountInformationResponseArrayOutput {
	return o
}

func (o StorageAccountInformationResponseArrayOutput) ToStorageAccountInformationResponseArrayOutputWithContext(ctx context.Context) StorageAccountInformationResponseArrayOutput {
	return o
}

func (o StorageAccountInformationResponseArrayOutput) Index(i pulumi.IntInput) StorageAccountInformationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccountInformationResponse {
		return vs[0].([]StorageAccountInformationResponse)[vs[1].(int)]
	}).(StorageAccountInformationResponseOutput)
}

type VirtualNetworkRuleResponse struct {
	Id                      string `pulumi:"id"`
	Name                    string `pulumi:"name"`
	SubnetId                string `pulumi:"subnetId"`
	Type                    string `pulumi:"type"`
	VirtualNetworkRuleState string `pulumi:"virtualNetworkRuleState"`
}

type VirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleResponseOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleResponseOutput) VirtualNetworkRuleState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.VirtualNetworkRuleState }).(pulumi.StringOutput)
}

type VirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRuleResponse {
		return vs[0].([]VirtualNetworkRuleResponse)[vs[1].(int)]
	}).(VirtualNetworkRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AddDataLakeStoreWithAccountParametersOutput{})
	pulumi.RegisterOutputType(AddDataLakeStoreWithAccountParametersArrayOutput{})
	pulumi.RegisterOutputType(AddStorageAccountWithAccountParametersOutput{})
	pulumi.RegisterOutputType(AddStorageAccountWithAccountParametersArrayOutput{})
	pulumi.RegisterOutputType(ComputePolicyResponseOutput{})
	pulumi.RegisterOutputType(ComputePolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(CreateComputePolicyWithAccountParametersOutput{})
	pulumi.RegisterOutputType(CreateComputePolicyWithAccountParametersArrayOutput{})
	pulumi.RegisterOutputType(CreateFirewallRuleWithAccountParametersOutput{})
	pulumi.RegisterOutputType(CreateFirewallRuleWithAccountParametersArrayOutput{})
	pulumi.RegisterOutputType(DataLakeStoreAccountInformationResponseOutput{})
	pulumi.RegisterOutputType(DataLakeStoreAccountInformationResponseArrayOutput{})
	pulumi.RegisterOutputType(FirewallRuleResponseOutput{})
	pulumi.RegisterOutputType(FirewallRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(HiveMetastoreResponseOutput{})
	pulumi.RegisterOutputType(HiveMetastoreResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountInformationResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountInformationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
