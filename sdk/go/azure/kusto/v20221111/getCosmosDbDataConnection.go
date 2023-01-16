


package v20221111

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCosmosDbDataConnection(ctx *pulumi.Context, args *LookupCosmosDbDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupCosmosDbDataConnectionResult, error) {
	var rv LookupCosmosDbDataConnectionResult
	err := ctx.Invoke("azure-native:kusto/v20221111:getCosmosDbDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCosmosDbDataConnectionArgs struct {
	ClusterName        string `pulumi:"clusterName"`
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupCosmosDbDataConnectionResult struct {
	CosmosDbAccountResourceId string  `pulumi:"cosmosDbAccountResourceId"`
	CosmosDbContainer         string  `pulumi:"cosmosDbContainer"`
	CosmosDbDatabase          string  `pulumi:"cosmosDbDatabase"`
	Id                        string  `pulumi:"id"`
	Kind                      string  `pulumi:"kind"`
	Location                  *string `pulumi:"location"`
	ManagedIdentityObjectId   string  `pulumi:"managedIdentityObjectId"`
	ManagedIdentityResourceId string  `pulumi:"managedIdentityResourceId"`
	MappingRuleName           *string `pulumi:"mappingRuleName"`
	Name                      string  `pulumi:"name"`
	ProvisioningState         string  `pulumi:"provisioningState"`
	RetrievalStartDate        *string `pulumi:"retrievalStartDate"`
	TableName                 string  `pulumi:"tableName"`
	Type                      string  `pulumi:"type"`
}

func LookupCosmosDbDataConnectionOutput(ctx *pulumi.Context, args LookupCosmosDbDataConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupCosmosDbDataConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCosmosDbDataConnectionResult, error) {
			args := v.(LookupCosmosDbDataConnectionArgs)
			r, err := LookupCosmosDbDataConnection(ctx, &args, opts...)
			var s LookupCosmosDbDataConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCosmosDbDataConnectionResultOutput)
}

type LookupCosmosDbDataConnectionOutputArgs struct {
	ClusterName        pulumi.StringInput `pulumi:"clusterName"`
	DataConnectionName pulumi.StringInput `pulumi:"dataConnectionName"`
	DatabaseName       pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCosmosDbDataConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCosmosDbDataConnectionArgs)(nil)).Elem()
}


type LookupCosmosDbDataConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupCosmosDbDataConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCosmosDbDataConnectionResult)(nil)).Elem()
}

func (o LookupCosmosDbDataConnectionResultOutput) ToLookupCosmosDbDataConnectionResultOutput() LookupCosmosDbDataConnectionResultOutput {
	return o
}

func (o LookupCosmosDbDataConnectionResultOutput) ToLookupCosmosDbDataConnectionResultOutputWithContext(ctx context.Context) LookupCosmosDbDataConnectionResultOutput {
	return o
}

func (o LookupCosmosDbDataConnectionResultOutput) CosmosDbAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) string { return v.CosmosDbAccountResourceId }).(pulumi.StringOutput)
}

func (o LookupCosmosDbDataConnectionResultOutput) CosmosDbContainer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) string { return v.CosmosDbContainer }).(pulumi.StringOutput)
}

func (o LookupCosmosDbDataConnectionResultOutput) CosmosDbDatabase() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) string { return v.CosmosDbDatabase }).(pulumi.StringOutput)
}

func (o LookupCosmosDbDataConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCosmosDbDataConnectionResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupCosmosDbDataConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupCosmosDbDataConnectionResultOutput) ManagedIdentityObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) string { return v.ManagedIdentityObjectId }).(pulumi.StringOutput)
}

func (o LookupCosmosDbDataConnectionResultOutput) ManagedIdentityResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) string { return v.ManagedIdentityResourceId }).(pulumi.StringOutput)
}

func (o LookupCosmosDbDataConnectionResultOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) *string { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

func (o LookupCosmosDbDataConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCosmosDbDataConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCosmosDbDataConnectionResultOutput) RetrievalStartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) *string { return v.RetrievalStartDate }).(pulumi.StringPtrOutput)
}

func (o LookupCosmosDbDataConnectionResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) string { return v.TableName }).(pulumi.StringOutput)
}

func (o LookupCosmosDbDataConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCosmosDbDataConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCosmosDbDataConnectionResultOutput{})
}
