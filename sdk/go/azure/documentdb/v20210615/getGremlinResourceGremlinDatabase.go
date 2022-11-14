


package v20210615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGremlinResourceGremlinDatabase(ctx *pulumi.Context, args *LookupGremlinResourceGremlinDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupGremlinResourceGremlinDatabaseResult, error) {
	var rv LookupGremlinResourceGremlinDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20210615:getGremlinResourceGremlinDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGremlinResourceGremlinDatabaseArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupGremlinResourceGremlinDatabaseResult struct {
	Id       string                                        `pulumi:"id"`
	Location *string                                       `pulumi:"location"`
	Name     string                                        `pulumi:"name"`
	Options  *GremlinDatabaseGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *GremlinDatabaseGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                             `pulumi:"tags"`
	Type     string                                        `pulumi:"type"`
}

func LookupGremlinResourceGremlinDatabaseOutput(ctx *pulumi.Context, args LookupGremlinResourceGremlinDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupGremlinResourceGremlinDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGremlinResourceGremlinDatabaseResult, error) {
			args := v.(LookupGremlinResourceGremlinDatabaseArgs)
			r, err := LookupGremlinResourceGremlinDatabase(ctx, &args, opts...)
			var s LookupGremlinResourceGremlinDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGremlinResourceGremlinDatabaseResultOutput)
}

type LookupGremlinResourceGremlinDatabaseOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGremlinResourceGremlinDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGremlinResourceGremlinDatabaseArgs)(nil)).Elem()
}


type LookupGremlinResourceGremlinDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupGremlinResourceGremlinDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGremlinResourceGremlinDatabaseResult)(nil)).Elem()
}

func (o LookupGremlinResourceGremlinDatabaseResultOutput) ToLookupGremlinResourceGremlinDatabaseResultOutput() LookupGremlinResourceGremlinDatabaseResultOutput {
	return o
}

func (o LookupGremlinResourceGremlinDatabaseResultOutput) ToLookupGremlinResourceGremlinDatabaseResultOutputWithContext(ctx context.Context) LookupGremlinResourceGremlinDatabaseResultOutput {
	return o
}

func (o LookupGremlinResourceGremlinDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGremlinResourceGremlinDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupGremlinResourceGremlinDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGremlinResourceGremlinDatabaseResultOutput) Options() GremlinDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) *GremlinDatabaseGetPropertiesResponseOptions {
		return v.Options
	}).(GremlinDatabaseGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupGremlinResourceGremlinDatabaseResultOutput) Resource() GremlinDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) *GremlinDatabaseGetPropertiesResponseResource {
		return v.Resource
	}).(GremlinDatabaseGetPropertiesResponseResourcePtrOutput)
}

func (o LookupGremlinResourceGremlinDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGremlinResourceGremlinDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGremlinResourceGremlinDatabaseResultOutput{})
}
