


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupKustoPoolDatabase(ctx *pulumi.Context, args *LookupKustoPoolDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupKustoPoolDatabaseResult, error) {
	var rv LookupKustoPoolDatabaseResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getKustoPoolDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoPoolDatabaseArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	KustoPoolName     string `pulumi:"kustoPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupKustoPoolDatabaseResult struct {
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Location   *string            `pulumi:"location"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupKustoPoolDatabaseOutput(ctx *pulumi.Context, args LookupKustoPoolDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupKustoPoolDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoPoolDatabaseResult, error) {
			args := v.(LookupKustoPoolDatabaseArgs)
			r, err := LookupKustoPoolDatabase(ctx, &args, opts...)
			var s LookupKustoPoolDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoPoolDatabaseResultOutput)
}

type LookupKustoPoolDatabaseOutputArgs struct {
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	KustoPoolName     pulumi.StringInput `pulumi:"kustoPoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupKustoPoolDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolDatabaseArgs)(nil)).Elem()
}


type LookupKustoPoolDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupKustoPoolDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolDatabaseResult)(nil)).Elem()
}

func (o LookupKustoPoolDatabaseResultOutput) ToLookupKustoPoolDatabaseResultOutput() LookupKustoPoolDatabaseResultOutput {
	return o
}

func (o LookupKustoPoolDatabaseResultOutput) ToLookupKustoPoolDatabaseResultOutputWithContext(ctx context.Context) LookupKustoPoolDatabaseResultOutput {
	return o
}

func (o LookupKustoPoolDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDatabaseResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabaseResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupKustoPoolDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDatabaseResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabaseResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupKustoPoolDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoPoolDatabaseResultOutput{})
}
