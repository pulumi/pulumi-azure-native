


package v20210615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMongoDBResourceMongoDBCollection(ctx *pulumi.Context, args *LookupMongoDBResourceMongoDBCollectionArgs, opts ...pulumi.InvokeOption) (*LookupMongoDBResourceMongoDBCollectionResult, error) {
	var rv LookupMongoDBResourceMongoDBCollectionResult
	err := ctx.Invoke("azure-native:documentdb/v20210615:getMongoDBResourceMongoDBCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoDBResourceMongoDBCollectionArgs struct {
	AccountName       string `pulumi:"accountName"`
	CollectionName    string `pulumi:"collectionName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMongoDBResourceMongoDBCollectionResult struct {
	Id       string                                          `pulumi:"id"`
	Location *string                                         `pulumi:"location"`
	Name     string                                          `pulumi:"name"`
	Options  *MongoDBCollectionGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *MongoDBCollectionGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                               `pulumi:"tags"`
	Type     string                                          `pulumi:"type"`
}

func LookupMongoDBResourceMongoDBCollectionOutput(ctx *pulumi.Context, args LookupMongoDBResourceMongoDBCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupMongoDBResourceMongoDBCollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMongoDBResourceMongoDBCollectionResult, error) {
			args := v.(LookupMongoDBResourceMongoDBCollectionArgs)
			r, err := LookupMongoDBResourceMongoDBCollection(ctx, &args, opts...)
			var s LookupMongoDBResourceMongoDBCollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMongoDBResourceMongoDBCollectionResultOutput)
}

type LookupMongoDBResourceMongoDBCollectionOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	CollectionName    pulumi.StringInput `pulumi:"collectionName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMongoDBResourceMongoDBCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoDBCollectionArgs)(nil)).Elem()
}


type LookupMongoDBResourceMongoDBCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupMongoDBResourceMongoDBCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoDBCollectionResult)(nil)).Elem()
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) ToLookupMongoDBResourceMongoDBCollectionResultOutput() LookupMongoDBResourceMongoDBCollectionResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) ToLookupMongoDBResourceMongoDBCollectionResultOutputWithContext(ctx context.Context) LookupMongoDBResourceMongoDBCollectionResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Options() MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) *MongoDBCollectionGetPropertiesResponseOptions {
		return v.Options
	}).(MongoDBCollectionGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Resource() MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) *MongoDBCollectionGetPropertiesResponseResource {
		return v.Resource
	}).(MongoDBCollectionGetPropertiesResponseResourcePtrOutput)
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMongoDBResourceMongoDBCollectionResultOutput{})
}
