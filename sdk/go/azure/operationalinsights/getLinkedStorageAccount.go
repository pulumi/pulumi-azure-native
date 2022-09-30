


package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedStorageAccount(ctx *pulumi.Context, args *LookupLinkedStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupLinkedStorageAccountResult, error) {
	var rv LookupLinkedStorageAccountResult
	err := ctx.Invoke("azure-native:operationalinsights:getLinkedStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedStorageAccountArgs struct {
	DataSourceType    string `pulumi:"dataSourceType"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupLinkedStorageAccountResult struct {
	DataSourceType    string   `pulumi:"dataSourceType"`
	Id                string   `pulumi:"id"`
	Name              string   `pulumi:"name"`
	StorageAccountIds []string `pulumi:"storageAccountIds"`
	Type              string   `pulumi:"type"`
}

func LookupLinkedStorageAccountOutput(ctx *pulumi.Context, args LookupLinkedStorageAccountOutputArgs, opts ...pulumi.InvokeOption) LookupLinkedStorageAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkedStorageAccountResult, error) {
			args := v.(LookupLinkedStorageAccountArgs)
			r, err := LookupLinkedStorageAccount(ctx, &args, opts...)
			var s LookupLinkedStorageAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkedStorageAccountResultOutput)
}

type LookupLinkedStorageAccountOutputArgs struct {
	DataSourceType    pulumi.StringInput `pulumi:"dataSourceType"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupLinkedStorageAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedStorageAccountArgs)(nil)).Elem()
}


type LookupLinkedStorageAccountResultOutput struct{ *pulumi.OutputState }

func (LookupLinkedStorageAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedStorageAccountResult)(nil)).Elem()
}

func (o LookupLinkedStorageAccountResultOutput) ToLookupLinkedStorageAccountResultOutput() LookupLinkedStorageAccountResultOutput {
	return o
}

func (o LookupLinkedStorageAccountResultOutput) ToLookupLinkedStorageAccountResultOutputWithContext(ctx context.Context) LookupLinkedStorageAccountResultOutput {
	return o
}

func (o LookupLinkedStorageAccountResultOutput) DataSourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedStorageAccountResult) string { return v.DataSourceType }).(pulumi.StringOutput)
}

func (o LookupLinkedStorageAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedStorageAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLinkedStorageAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedStorageAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLinkedStorageAccountResultOutput) StorageAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLinkedStorageAccountResult) []string { return v.StorageAccountIds }).(pulumi.StringArrayOutput)
}

func (o LookupLinkedStorageAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedStorageAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkedStorageAccountResultOutput{})
}
