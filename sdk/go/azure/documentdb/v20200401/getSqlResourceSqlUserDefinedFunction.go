


package v20200401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSqlResourceSqlUserDefinedFunction(ctx *pulumi.Context, args *LookupSqlResourceSqlUserDefinedFunctionArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlUserDefinedFunctionResult, error) {
	var rv LookupSqlResourceSqlUserDefinedFunctionResult
	err := ctx.Invoke("azure-native:documentdb/v20200401:getSqlResourceSqlUserDefinedFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlUserDefinedFunctionArgs struct {
	AccountName             string `pulumi:"accountName"`
	ContainerName           string `pulumi:"containerName"`
	DatabaseName            string `pulumi:"databaseName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	UserDefinedFunctionName string `pulumi:"userDefinedFunctionName"`
}


type LookupSqlResourceSqlUserDefinedFunctionResult struct {
	Id       string                                               `pulumi:"id"`
	Location *string                                              `pulumi:"location"`
	Name     string                                               `pulumi:"name"`
	Resource *SqlUserDefinedFunctionGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                                    `pulumi:"tags"`
	Type     string                                               `pulumi:"type"`
}

func LookupSqlResourceSqlUserDefinedFunctionOutput(ctx *pulumi.Context, args LookupSqlResourceSqlUserDefinedFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupSqlResourceSqlUserDefinedFunctionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlResourceSqlUserDefinedFunctionResult, error) {
			args := v.(LookupSqlResourceSqlUserDefinedFunctionArgs)
			r, err := LookupSqlResourceSqlUserDefinedFunction(ctx, &args, opts...)
			var s LookupSqlResourceSqlUserDefinedFunctionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlResourceSqlUserDefinedFunctionResultOutput)
}

type LookupSqlResourceSqlUserDefinedFunctionOutputArgs struct {
	AccountName             pulumi.StringInput `pulumi:"accountName"`
	ContainerName           pulumi.StringInput `pulumi:"containerName"`
	DatabaseName            pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	UserDefinedFunctionName pulumi.StringInput `pulumi:"userDefinedFunctionName"`
}

func (LookupSqlResourceSqlUserDefinedFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlUserDefinedFunctionArgs)(nil)).Elem()
}


type LookupSqlResourceSqlUserDefinedFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupSqlResourceSqlUserDefinedFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlUserDefinedFunctionResult)(nil)).Elem()
}

func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) ToLookupSqlResourceSqlUserDefinedFunctionResultOutput() LookupSqlResourceSqlUserDefinedFunctionResultOutput {
	return o
}

func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) ToLookupSqlResourceSqlUserDefinedFunctionResultOutputWithContext(ctx context.Context) LookupSqlResourceSqlUserDefinedFunctionResultOutput {
	return o
}

func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlUserDefinedFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlUserDefinedFunctionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlUserDefinedFunctionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) Resource() SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlUserDefinedFunctionResult) *SqlUserDefinedFunctionGetPropertiesResponseResource {
		return v.Resource
	}).(SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput)
}

func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlUserDefinedFunctionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlUserDefinedFunctionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlResourceSqlUserDefinedFunctionResultOutput{})
}
