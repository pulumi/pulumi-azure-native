


package v20220515preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCassandraResourceCassandraTable(ctx *pulumi.Context, args *LookupCassandraResourceCassandraTableArgs, opts ...pulumi.InvokeOption) (*LookupCassandraResourceCassandraTableResult, error) {
	var rv LookupCassandraResourceCassandraTableResult
	err := ctx.Invoke("azure-native:documentdb/v20220515preview:getCassandraResourceCassandraTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCassandraResourceCassandraTableArgs struct {
	AccountName       string `pulumi:"accountName"`
	KeyspaceName      string `pulumi:"keyspaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TableName         string `pulumi:"tableName"`
}


type LookupCassandraResourceCassandraTableResult struct {
	Id       string                                       `pulumi:"id"`
	Identity *ManagedServiceIdentityResponse              `pulumi:"identity"`
	Location *string                                      `pulumi:"location"`
	Name     string                                       `pulumi:"name"`
	Options  *CassandraTableGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *CassandraTableGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                            `pulumi:"tags"`
	Type     string                                       `pulumi:"type"`
}

func LookupCassandraResourceCassandraTableOutput(ctx *pulumi.Context, args LookupCassandraResourceCassandraTableOutputArgs, opts ...pulumi.InvokeOption) LookupCassandraResourceCassandraTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCassandraResourceCassandraTableResult, error) {
			args := v.(LookupCassandraResourceCassandraTableArgs)
			r, err := LookupCassandraResourceCassandraTable(ctx, &args, opts...)
			var s LookupCassandraResourceCassandraTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCassandraResourceCassandraTableResultOutput)
}

type LookupCassandraResourceCassandraTableOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	KeyspaceName      pulumi.StringInput `pulumi:"keyspaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TableName         pulumi.StringInput `pulumi:"tableName"`
}

func (LookupCassandraResourceCassandraTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraResourceCassandraTableArgs)(nil)).Elem()
}


type LookupCassandraResourceCassandraTableResultOutput struct{ *pulumi.OutputState }

func (LookupCassandraResourceCassandraTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraResourceCassandraTableResult)(nil)).Elem()
}

func (o LookupCassandraResourceCassandraTableResultOutput) ToLookupCassandraResourceCassandraTableResultOutput() LookupCassandraResourceCassandraTableResultOutput {
	return o
}

func (o LookupCassandraResourceCassandraTableResultOutput) ToLookupCassandraResourceCassandraTableResultOutputWithContext(ctx context.Context) LookupCassandraResourceCassandraTableResultOutput {
	return o
}

func (o LookupCassandraResourceCassandraTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCassandraResourceCassandraTableResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupCassandraResourceCassandraTableResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupCassandraResourceCassandraTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCassandraResourceCassandraTableResultOutput) Options() CassandraTableGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) *CassandraTableGetPropertiesResponseOptions {
		return v.Options
	}).(CassandraTableGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupCassandraResourceCassandraTableResultOutput) Resource() CassandraTableGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) *CassandraTableGetPropertiesResponseResource {
		return v.Resource
	}).(CassandraTableGetPropertiesResponseResourcePtrOutput)
}

func (o LookupCassandraResourceCassandraTableResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCassandraResourceCassandraTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCassandraResourceCassandraTableResultOutput{})
}
