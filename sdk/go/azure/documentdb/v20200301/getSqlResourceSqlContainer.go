


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSqlResourceSqlContainer(ctx *pulumi.Context, args *LookupSqlResourceSqlContainerArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlContainerResult, error) {
	var rv LookupSqlResourceSqlContainerResult
	err := ctx.Invoke("azure-native:documentdb/v20200301:getSqlResourceSqlContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSqlResourceSqlContainerArgs struct {
	AccountName       string `pulumi:"accountName"`
	ContainerName     string `pulumi:"containerName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSqlResourceSqlContainerResult struct {
	Id       string                                     `pulumi:"id"`
	Location *string                                    `pulumi:"location"`
	Name     string                                     `pulumi:"name"`
	Options  *SqlContainerGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *SqlContainerGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                          `pulumi:"tags"`
	Type     string                                     `pulumi:"type"`
}


func (val *LookupSqlResourceSqlContainerResult) Defaults() *LookupSqlResourceSqlContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Resource = tmp.Resource.Defaults()

	return &tmp
}

func LookupSqlResourceSqlContainerOutput(ctx *pulumi.Context, args LookupSqlResourceSqlContainerOutputArgs, opts ...pulumi.InvokeOption) LookupSqlResourceSqlContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlResourceSqlContainerResult, error) {
			args := v.(LookupSqlResourceSqlContainerArgs)
			r, err := LookupSqlResourceSqlContainer(ctx, &args, opts...)
			var s LookupSqlResourceSqlContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlResourceSqlContainerResultOutput)
}

type LookupSqlResourceSqlContainerOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ContainerName     pulumi.StringInput `pulumi:"containerName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSqlResourceSqlContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlContainerArgs)(nil)).Elem()
}


type LookupSqlResourceSqlContainerResultOutput struct{ *pulumi.OutputState }

func (LookupSqlResourceSqlContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlContainerResult)(nil)).Elem()
}

func (o LookupSqlResourceSqlContainerResultOutput) ToLookupSqlResourceSqlContainerResultOutput() LookupSqlResourceSqlContainerResultOutput {
	return o
}

func (o LookupSqlResourceSqlContainerResultOutput) ToLookupSqlResourceSqlContainerResultOutputWithContext(ctx context.Context) LookupSqlResourceSqlContainerResultOutput {
	return o
}

func (o LookupSqlResourceSqlContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlContainerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSqlResourceSqlContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlContainerResultOutput) Options() SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) *SqlContainerGetPropertiesResponseOptions {
		return v.Options
	}).(SqlContainerGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupSqlResourceSqlContainerResultOutput) Resource() SqlContainerGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) *SqlContainerGetPropertiesResponseResource {
		return v.Resource
	}).(SqlContainerGetPropertiesResponseResourcePtrOutput)
}

func (o LookupSqlResourceSqlContainerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlResourceSqlContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlResourceSqlContainerResultOutput{})
}
