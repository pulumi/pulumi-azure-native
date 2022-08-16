


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupKustoPoolDataConnection(ctx *pulumi.Context, args *LookupKustoPoolDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupKustoPoolDataConnectionResult, error) {
	var rv LookupKustoPoolDataConnectionResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getKustoPoolDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoPoolDataConnectionArgs struct {
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	KustoPoolName      string `pulumi:"kustoPoolName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	WorkspaceName      string `pulumi:"workspaceName"`
}


type LookupKustoPoolDataConnectionResult struct {
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Location   *string            `pulumi:"location"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupKustoPoolDataConnectionOutput(ctx *pulumi.Context, args LookupKustoPoolDataConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupKustoPoolDataConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoPoolDataConnectionResult, error) {
			args := v.(LookupKustoPoolDataConnectionArgs)
			r, err := LookupKustoPoolDataConnection(ctx, &args, opts...)
			var s LookupKustoPoolDataConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoPoolDataConnectionResultOutput)
}

type LookupKustoPoolDataConnectionOutputArgs struct {
	DataConnectionName pulumi.StringInput `pulumi:"dataConnectionName"`
	DatabaseName       pulumi.StringInput `pulumi:"databaseName"`
	KustoPoolName      pulumi.StringInput `pulumi:"kustoPoolName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName      pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupKustoPoolDataConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolDataConnectionArgs)(nil)).Elem()
}


type LookupKustoPoolDataConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupKustoPoolDataConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolDataConnectionResult)(nil)).Elem()
}

func (o LookupKustoPoolDataConnectionResultOutput) ToLookupKustoPoolDataConnectionResultOutput() LookupKustoPoolDataConnectionResultOutput {
	return o
}

func (o LookupKustoPoolDataConnectionResultOutput) ToLookupKustoPoolDataConnectionResultOutputWithContext(ctx context.Context) LookupKustoPoolDataConnectionResultOutput {
	return o
}

func (o LookupKustoPoolDataConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDataConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDataConnectionResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDataConnectionResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDataConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKustoPoolDataConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupKustoPoolDataConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDataConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKustoPoolDataConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoPoolDataConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupKustoPoolDataConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolDataConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoPoolDataConnectionResultOutput{})
}
