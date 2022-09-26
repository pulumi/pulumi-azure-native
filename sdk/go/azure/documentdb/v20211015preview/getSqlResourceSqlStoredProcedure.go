


package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlResourceSqlStoredProcedure(ctx *pulumi.Context, args *LookupSqlResourceSqlStoredProcedureArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlStoredProcedureResult, error) {
	var rv LookupSqlResourceSqlStoredProcedureResult
	err := ctx.Invoke("azure-native:documentdb/v20211015preview:getSqlResourceSqlStoredProcedure", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlStoredProcedureArgs struct {
	AccountName         string `pulumi:"accountName"`
	ContainerName       string `pulumi:"containerName"`
	DatabaseName        string `pulumi:"databaseName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	StoredProcedureName string `pulumi:"storedProcedureName"`
}


type LookupSqlResourceSqlStoredProcedureResult struct {
	Id       string                                           `pulumi:"id"`
	Identity *ManagedServiceIdentityResponse                  `pulumi:"identity"`
	Location *string                                          `pulumi:"location"`
	Name     string                                           `pulumi:"name"`
	Resource *SqlStoredProcedureGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                                `pulumi:"tags"`
	Type     string                                           `pulumi:"type"`
}

func LookupSqlResourceSqlStoredProcedureOutput(ctx *pulumi.Context, args LookupSqlResourceSqlStoredProcedureOutputArgs, opts ...pulumi.InvokeOption) LookupSqlResourceSqlStoredProcedureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlResourceSqlStoredProcedureResult, error) {
			args := v.(LookupSqlResourceSqlStoredProcedureArgs)
			r, err := LookupSqlResourceSqlStoredProcedure(ctx, &args, opts...)
			var s LookupSqlResourceSqlStoredProcedureResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlResourceSqlStoredProcedureResultOutput)
}

type LookupSqlResourceSqlStoredProcedureOutputArgs struct {
	AccountName         pulumi.StringInput `pulumi:"accountName"`
	ContainerName       pulumi.StringInput `pulumi:"containerName"`
	DatabaseName        pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	StoredProcedureName pulumi.StringInput `pulumi:"storedProcedureName"`
}

func (LookupSqlResourceSqlStoredProcedureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlStoredProcedureArgs)(nil)).Elem()
}


type LookupSqlResourceSqlStoredProcedureResultOutput struct{ *pulumi.OutputState }

func (LookupSqlResourceSqlStoredProcedureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlStoredProcedureResult)(nil)).Elem()
}

func (o LookupSqlResourceSqlStoredProcedureResultOutput) ToLookupSqlResourceSqlStoredProcedureResultOutput() LookupSqlResourceSqlStoredProcedureResultOutput {
	return o
}

func (o LookupSqlResourceSqlStoredProcedureResultOutput) ToLookupSqlResourceSqlStoredProcedureResultOutputWithContext(ctx context.Context) LookupSqlResourceSqlStoredProcedureResultOutput {
	return o
}

func (o LookupSqlResourceSqlStoredProcedureResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlStoredProcedureResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlStoredProcedureResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlStoredProcedureResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupSqlResourceSqlStoredProcedureResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlStoredProcedureResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSqlResourceSqlStoredProcedureResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlStoredProcedureResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlStoredProcedureResultOutput) Resource() SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlStoredProcedureResult) *SqlStoredProcedureGetPropertiesResponseResource {
		return v.Resource
	}).(SqlStoredProcedureGetPropertiesResponseResourcePtrOutput)
}

func (o LookupSqlResourceSqlStoredProcedureResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlStoredProcedureResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlResourceSqlStoredProcedureResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlStoredProcedureResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlResourceSqlStoredProcedureResultOutput{})
}
