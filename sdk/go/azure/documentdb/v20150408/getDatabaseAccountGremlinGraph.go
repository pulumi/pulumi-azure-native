


package v20150408

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDatabaseAccountGremlinGraph(ctx *pulumi.Context, args *LookupDatabaseAccountGremlinGraphArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountGremlinGraphResult, error) {
	var rv LookupDatabaseAccountGremlinGraphResult
	err := ctx.Invoke("azure-native:documentdb/v20150408:getDatabaseAccountGremlinGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDatabaseAccountGremlinGraphArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	GraphName         string `pulumi:"graphName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountGremlinGraphResult struct {
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


func (val *LookupDatabaseAccountGremlinGraphResult) Defaults() *LookupDatabaseAccountGremlinGraphResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConflictResolutionPolicy = tmp.ConflictResolutionPolicy.Defaults()

	tmp.IndexingPolicy = tmp.IndexingPolicy.Defaults()

	tmp.PartitionKey = tmp.PartitionKey.Defaults()

	return &tmp
}

func LookupDatabaseAccountGremlinGraphOutput(ctx *pulumi.Context, args LookupDatabaseAccountGremlinGraphOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAccountGremlinGraphResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAccountGremlinGraphResult, error) {
			args := v.(LookupDatabaseAccountGremlinGraphArgs)
			r, err := LookupDatabaseAccountGremlinGraph(ctx, &args, opts...)
			var s LookupDatabaseAccountGremlinGraphResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAccountGremlinGraphResultOutput)
}

type LookupDatabaseAccountGremlinGraphOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	GraphName         pulumi.StringInput `pulumi:"graphName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabaseAccountGremlinGraphOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountGremlinGraphArgs)(nil)).Elem()
}


type LookupDatabaseAccountGremlinGraphResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAccountGremlinGraphResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountGremlinGraphResult)(nil)).Elem()
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) ToLookupDatabaseAccountGremlinGraphResultOutput() LookupDatabaseAccountGremlinGraphResultOutput {
	return o
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) ToLookupDatabaseAccountGremlinGraphResultOutputWithContext(ctx context.Context) LookupDatabaseAccountGremlinGraphResultOutput {
	return o
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) ConflictResolutionPolicy() ConflictResolutionPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinGraphResult) *ConflictResolutionPolicyResponse {
		return v.ConflictResolutionPolicy
	}).(ConflictResolutionPolicyResponsePtrOutput)
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinGraphResult) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinGraphResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinGraphResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) IndexingPolicy() IndexingPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinGraphResult) *IndexingPolicyResponse { return v.IndexingPolicy }).(IndexingPolicyResponsePtrOutput)
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinGraphResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinGraphResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) PartitionKey() ContainerPartitionKeyResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinGraphResult) *ContainerPartitionKeyResponse { return v.PartitionKey }).(ContainerPartitionKeyResponsePtrOutput)
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinGraphResult) *string { return v.Rid }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinGraphResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) Ts() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinGraphResult) interface{} { return v.Ts }).(pulumi.AnyOutput)
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinGraphResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountGremlinGraphResultOutput) UniqueKeyPolicy() UniqueKeyPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinGraphResult) *UniqueKeyPolicyResponse { return v.UniqueKeyPolicy }).(UniqueKeyPolicyResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAccountGremlinGraphResultOutput{})
}
