


package v20160331

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDatabaseAccountSqlContainer(ctx *pulumi.Context, args *LookupDatabaseAccountSqlContainerArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountSqlContainerResult, error) {
	var rv LookupDatabaseAccountSqlContainerResult
	err := ctx.Invoke("azure-native:documentdb/v20160331:getDatabaseAccountSqlContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDatabaseAccountSqlContainerArgs struct {
	AccountName       string `pulumi:"accountName"`
	ContainerName     string `pulumi:"containerName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountSqlContainerResult struct {
	ConflictResolutionPolicy *ConflictResolutionPolicyResponse `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               *int                              `pulumi:"defaultTtl"`
	Etag                     *string                           `pulumi:"etag"`
	Id                       string                            `pulumi:"id"`
	IndexingPolicy           *IndexingPolicyResponse           `pulumi:"indexingPolicy"`
	Location                 *string                           `pulumi:"location"`
	Name                     string                            `pulumi:"name"`
	PartitionKey             *ContainerPartitionKeyResponse    `pulumi:"partitionKey"`
	Rid                      *string                           `pulumi:"rid"`
	Tags                     map[string]string                 `pulumi:"tags"`
	Ts                       interface{}                       `pulumi:"ts"`
	Type                     string                            `pulumi:"type"`
	UniqueKeyPolicy          *UniqueKeyPolicyResponse          `pulumi:"uniqueKeyPolicy"`
}


func (val *LookupDatabaseAccountSqlContainerResult) Defaults() *LookupDatabaseAccountSqlContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConflictResolutionPolicy = tmp.ConflictResolutionPolicy.Defaults()

	tmp.IndexingPolicy = tmp.IndexingPolicy.Defaults()

	tmp.PartitionKey = tmp.PartitionKey.Defaults()

	return &tmp
}

func LookupDatabaseAccountSqlContainerOutput(ctx *pulumi.Context, args LookupDatabaseAccountSqlContainerOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAccountSqlContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAccountSqlContainerResult, error) {
			args := v.(LookupDatabaseAccountSqlContainerArgs)
			r, err := LookupDatabaseAccountSqlContainer(ctx, &args, opts...)
			var s LookupDatabaseAccountSqlContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAccountSqlContainerResultOutput)
}

type LookupDatabaseAccountSqlContainerOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ContainerName     pulumi.StringInput `pulumi:"containerName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabaseAccountSqlContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountSqlContainerArgs)(nil)).Elem()
}


type LookupDatabaseAccountSqlContainerResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAccountSqlContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountSqlContainerResult)(nil)).Elem()
}

func (o LookupDatabaseAccountSqlContainerResultOutput) ToLookupDatabaseAccountSqlContainerResultOutput() LookupDatabaseAccountSqlContainerResultOutput {
	return o
}

func (o LookupDatabaseAccountSqlContainerResultOutput) ToLookupDatabaseAccountSqlContainerResultOutputWithContext(ctx context.Context) LookupDatabaseAccountSqlContainerResultOutput {
	return o
}

func (o LookupDatabaseAccountSqlContainerResultOutput) ConflictResolutionPolicy() ConflictResolutionPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlContainerResult) *ConflictResolutionPolicyResponse {
		return v.ConflictResolutionPolicy
	}).(ConflictResolutionPolicyResponsePtrOutput)
}

func (o LookupDatabaseAccountSqlContainerResultOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlContainerResult) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o LookupDatabaseAccountSqlContainerResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlContainerResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountSqlContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountSqlContainerResultOutput) IndexingPolicy() IndexingPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlContainerResult) *IndexingPolicyResponse { return v.IndexingPolicy }).(IndexingPolicyResponsePtrOutput)
}

func (o LookupDatabaseAccountSqlContainerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlContainerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountSqlContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountSqlContainerResultOutput) PartitionKey() ContainerPartitionKeyResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlContainerResult) *ContainerPartitionKeyResponse { return v.PartitionKey }).(ContainerPartitionKeyResponsePtrOutput)
}

func (o LookupDatabaseAccountSqlContainerResultOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlContainerResult) *string { return v.Rid }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountSqlContainerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlContainerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseAccountSqlContainerResultOutput) Ts() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlContainerResult) interface{} { return v.Ts }).(pulumi.AnyOutput)
}

func (o LookupDatabaseAccountSqlContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountSqlContainerResultOutput) UniqueKeyPolicy() UniqueKeyPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlContainerResult) *UniqueKeyPolicyResponse { return v.UniqueKeyPolicy }).(UniqueKeyPolicyResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAccountSqlContainerResultOutput{})
}
