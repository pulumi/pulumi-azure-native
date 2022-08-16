


package v20160331

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDatabaseAccountMongoDBCollection(ctx *pulumi.Context, args *LookupDatabaseAccountMongoDBCollectionArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountMongoDBCollectionResult, error) {
	var rv LookupDatabaseAccountMongoDBCollectionResult
	err := ctx.Invoke("azure-native:documentdb/v20160331:getDatabaseAccountMongoDBCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountMongoDBCollectionArgs struct {
	AccountName       string `pulumi:"accountName"`
	CollectionName    string `pulumi:"collectionName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountMongoDBCollectionResult struct {
	Id       string               `pulumi:"id"`
	Indexes  []MongoIndexResponse `pulumi:"indexes"`
	Location *string              `pulumi:"location"`
	Name     string               `pulumi:"name"`
	ShardKey map[string]string    `pulumi:"shardKey"`
	Tags     map[string]string    `pulumi:"tags"`
	Type     string               `pulumi:"type"`
}

func LookupDatabaseAccountMongoDBCollectionOutput(ctx *pulumi.Context, args LookupDatabaseAccountMongoDBCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAccountMongoDBCollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAccountMongoDBCollectionResult, error) {
			args := v.(LookupDatabaseAccountMongoDBCollectionArgs)
			r, err := LookupDatabaseAccountMongoDBCollection(ctx, &args, opts...)
			var s LookupDatabaseAccountMongoDBCollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAccountMongoDBCollectionResultOutput)
}

type LookupDatabaseAccountMongoDBCollectionOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	CollectionName    pulumi.StringInput `pulumi:"collectionName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabaseAccountMongoDBCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountMongoDBCollectionArgs)(nil)).Elem()
}


type LookupDatabaseAccountMongoDBCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAccountMongoDBCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountMongoDBCollectionResult)(nil)).Elem()
}

func (o LookupDatabaseAccountMongoDBCollectionResultOutput) ToLookupDatabaseAccountMongoDBCollectionResultOutput() LookupDatabaseAccountMongoDBCollectionResultOutput {
	return o
}

func (o LookupDatabaseAccountMongoDBCollectionResultOutput) ToLookupDatabaseAccountMongoDBCollectionResultOutputWithContext(ctx context.Context) LookupDatabaseAccountMongoDBCollectionResultOutput {
	return o
}

func (o LookupDatabaseAccountMongoDBCollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountMongoDBCollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountMongoDBCollectionResultOutput) Indexes() MongoIndexResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountMongoDBCollectionResult) []MongoIndexResponse { return v.Indexes }).(MongoIndexResponseArrayOutput)
}

func (o LookupDatabaseAccountMongoDBCollectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountMongoDBCollectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountMongoDBCollectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountMongoDBCollectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountMongoDBCollectionResultOutput) ShardKey() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountMongoDBCollectionResult) map[string]string { return v.ShardKey }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseAccountMongoDBCollectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountMongoDBCollectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseAccountMongoDBCollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountMongoDBCollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAccountMongoDBCollectionResultOutput{})
}
