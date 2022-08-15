


package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageInsightConfig(ctx *pulumi.Context, args *LookupStorageInsightConfigArgs, opts ...pulumi.InvokeOption) (*LookupStorageInsightConfigResult, error) {
	var rv LookupStorageInsightConfigResult
	err := ctx.Invoke("azure-native:operationalinsights/v20200301preview:getStorageInsightConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageInsightConfigArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	StorageInsightName string `pulumi:"storageInsightName"`
	WorkspaceName      string `pulumi:"workspaceName"`
}


type LookupStorageInsightConfigResult struct {
	Containers     []string                     `pulumi:"containers"`
	ETag           *string                      `pulumi:"eTag"`
	Id             string                       `pulumi:"id"`
	Name           string                       `pulumi:"name"`
	Status         StorageInsightStatusResponse `pulumi:"status"`
	StorageAccount StorageAccountResponse       `pulumi:"storageAccount"`
	Tables         []string                     `pulumi:"tables"`
	Tags           map[string]string            `pulumi:"tags"`
	Type           string                       `pulumi:"type"`
}

func LookupStorageInsightConfigOutput(ctx *pulumi.Context, args LookupStorageInsightConfigOutputArgs, opts ...pulumi.InvokeOption) LookupStorageInsightConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageInsightConfigResult, error) {
			args := v.(LookupStorageInsightConfigArgs)
			r, err := LookupStorageInsightConfig(ctx, &args, opts...)
			var s LookupStorageInsightConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageInsightConfigResultOutput)
}

type LookupStorageInsightConfigOutputArgs struct {
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageInsightName pulumi.StringInput `pulumi:"storageInsightName"`
	WorkspaceName      pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupStorageInsightConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageInsightConfigArgs)(nil)).Elem()
}


type LookupStorageInsightConfigResultOutput struct{ *pulumi.OutputState }

func (LookupStorageInsightConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageInsightConfigResult)(nil)).Elem()
}

func (o LookupStorageInsightConfigResultOutput) ToLookupStorageInsightConfigResultOutput() LookupStorageInsightConfigResultOutput {
	return o
}

func (o LookupStorageInsightConfigResultOutput) ToLookupStorageInsightConfigResultOutputWithContext(ctx context.Context) LookupStorageInsightConfigResultOutput {
	return o
}

func (o LookupStorageInsightConfigResultOutput) Containers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStorageInsightConfigResult) []string { return v.Containers }).(pulumi.StringArrayOutput)
}

func (o LookupStorageInsightConfigResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageInsightConfigResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupStorageInsightConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageInsightConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageInsightConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageInsightConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageInsightConfigResultOutput) Status() StorageInsightStatusResponseOutput {
	return o.ApplyT(func(v LookupStorageInsightConfigResult) StorageInsightStatusResponse { return v.Status }).(StorageInsightStatusResponseOutput)
}

func (o LookupStorageInsightConfigResultOutput) StorageAccount() StorageAccountResponseOutput {
	return o.ApplyT(func(v LookupStorageInsightConfigResult) StorageAccountResponse { return v.StorageAccount }).(StorageAccountResponseOutput)
}

func (o LookupStorageInsightConfigResultOutput) Tables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStorageInsightConfigResult) []string { return v.Tables }).(pulumi.StringArrayOutput)
}

func (o LookupStorageInsightConfigResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStorageInsightConfigResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStorageInsightConfigResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageInsightConfigResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageInsightConfigResultOutput{})
}
