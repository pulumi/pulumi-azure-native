


package v20211015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMongoDBResourceMongoDBDatabase(ctx *pulumi.Context, args *LookupMongoDBResourceMongoDBDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupMongoDBResourceMongoDBDatabaseResult, error) {
	var rv LookupMongoDBResourceMongoDBDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20211015:getMongoDBResourceMongoDBDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoDBResourceMongoDBDatabaseArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMongoDBResourceMongoDBDatabaseResult struct {
	Id       string                                        `pulumi:"id"`
	Location *string                                       `pulumi:"location"`
	Name     string                                        `pulumi:"name"`
	Options  *MongoDBDatabaseGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *MongoDBDatabaseGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                             `pulumi:"tags"`
	Type     string                                        `pulumi:"type"`
}

func LookupMongoDBResourceMongoDBDatabaseOutput(ctx *pulumi.Context, args LookupMongoDBResourceMongoDBDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupMongoDBResourceMongoDBDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMongoDBResourceMongoDBDatabaseResult, error) {
			args := v.(LookupMongoDBResourceMongoDBDatabaseArgs)
			r, err := LookupMongoDBResourceMongoDBDatabase(ctx, &args, opts...)
			var s LookupMongoDBResourceMongoDBDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMongoDBResourceMongoDBDatabaseResultOutput)
}

type LookupMongoDBResourceMongoDBDatabaseOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMongoDBResourceMongoDBDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoDBDatabaseArgs)(nil)).Elem()
}


type LookupMongoDBResourceMongoDBDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupMongoDBResourceMongoDBDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoDBDatabaseResult)(nil)).Elem()
}

func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) ToLookupMongoDBResourceMongoDBDatabaseResultOutput() LookupMongoDBResourceMongoDBDatabaseResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) ToLookupMongoDBResourceMongoDBDatabaseResultOutputWithContext(ctx context.Context) LookupMongoDBResourceMongoDBDatabaseResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Options() MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) *MongoDBDatabaseGetPropertiesResponseOptions {
		return v.Options
	}).(MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Resource() MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) *MongoDBDatabaseGetPropertiesResponseResource {
		return v.Resource
	}).(MongoDBDatabaseGetPropertiesResponseResourcePtrOutput)
}

func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMongoDBResourceMongoDBDatabaseResultOutput{})
}
