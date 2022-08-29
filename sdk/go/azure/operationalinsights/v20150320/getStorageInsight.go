


package v20150320

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupStorageInsight(ctx *pulumi.Context, args *LookupStorageInsightArgs, opts ...pulumi.InvokeOption) (*LookupStorageInsightResult, error) {
	var rv LookupStorageInsightResult
	err := ctx.Invoke("azure-native:operationalinsights/v20150320:getStorageInsight", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageInsightArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	StorageInsightName string `pulumi:"storageInsightName"`
	WorkspaceName      string `pulumi:"workspaceName"`
}


type LookupStorageInsightResult struct {
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

func LookupStorageInsightOutput(ctx *pulumi.Context, args LookupStorageInsightOutputArgs, opts ...pulumi.InvokeOption) LookupStorageInsightResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageInsightResult, error) {
			args := v.(LookupStorageInsightArgs)
			r, err := LookupStorageInsight(ctx, &args, opts...)
			var s LookupStorageInsightResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageInsightResultOutput)
}

type LookupStorageInsightOutputArgs struct {
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageInsightName pulumi.StringInput `pulumi:"storageInsightName"`
	WorkspaceName      pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupStorageInsightOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageInsightArgs)(nil)).Elem()
}


type LookupStorageInsightResultOutput struct{ *pulumi.OutputState }

func (LookupStorageInsightResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageInsightResult)(nil)).Elem()
}

func (o LookupStorageInsightResultOutput) ToLookupStorageInsightResultOutput() LookupStorageInsightResultOutput {
	return o
}

func (o LookupStorageInsightResultOutput) ToLookupStorageInsightResultOutputWithContext(ctx context.Context) LookupStorageInsightResultOutput {
	return o
}

func (o LookupStorageInsightResultOutput) Containers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStorageInsightResult) []string { return v.Containers }).(pulumi.StringArrayOutput)
}

func (o LookupStorageInsightResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageInsightResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupStorageInsightResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageInsightResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageInsightResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageInsightResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageInsightResultOutput) Status() StorageInsightStatusResponseOutput {
	return o.ApplyT(func(v LookupStorageInsightResult) StorageInsightStatusResponse { return v.Status }).(StorageInsightStatusResponseOutput)
}

func (o LookupStorageInsightResultOutput) StorageAccount() StorageAccountResponseOutput {
	return o.ApplyT(func(v LookupStorageInsightResult) StorageAccountResponse { return v.StorageAccount }).(StorageAccountResponseOutput)
}

func (o LookupStorageInsightResultOutput) Tables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStorageInsightResult) []string { return v.Tables }).(pulumi.StringArrayOutput)
}

func (o LookupStorageInsightResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStorageInsightResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStorageInsightResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageInsightResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageInsightResultOutput{})
}
