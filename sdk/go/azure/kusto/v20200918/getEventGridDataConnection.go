


package v20200918

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventGridDataConnection(ctx *pulumi.Context, args *LookupEventGridDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEventGridDataConnectionResult, error) {
	var rv LookupEventGridDataConnectionResult
	err := ctx.Invoke("azure-native:kusto/v20200918:getEventGridDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventGridDataConnectionArgs struct {
	ClusterName        string `pulumi:"clusterName"`
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupEventGridDataConnectionResult struct {
	BlobStorageEventType     *string `pulumi:"blobStorageEventType"`
	ConsumerGroup            string  `pulumi:"consumerGroup"`
	DataFormat               *string `pulumi:"dataFormat"`
	EventHubResourceId       string  `pulumi:"eventHubResourceId"`
	Id                       string  `pulumi:"id"`
	IgnoreFirstRecord        *bool   `pulumi:"ignoreFirstRecord"`
	Kind                     string  `pulumi:"kind"`
	Location                 *string `pulumi:"location"`
	MappingRuleName          *string `pulumi:"mappingRuleName"`
	Name                     string  `pulumi:"name"`
	ProvisioningState        string  `pulumi:"provisioningState"`
	StorageAccountResourceId string  `pulumi:"storageAccountResourceId"`
	TableName                *string `pulumi:"tableName"`
	Type                     string  `pulumi:"type"`
}

func LookupEventGridDataConnectionOutput(ctx *pulumi.Context, args LookupEventGridDataConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupEventGridDataConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventGridDataConnectionResult, error) {
			args := v.(LookupEventGridDataConnectionArgs)
			r, err := LookupEventGridDataConnection(ctx, &args, opts...)
			var s LookupEventGridDataConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventGridDataConnectionResultOutput)
}

type LookupEventGridDataConnectionOutputArgs struct {
	ClusterName        pulumi.StringInput `pulumi:"clusterName"`
	DataConnectionName pulumi.StringInput `pulumi:"dataConnectionName"`
	DatabaseName       pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventGridDataConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventGridDataConnectionArgs)(nil)).Elem()
}


type LookupEventGridDataConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupEventGridDataConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventGridDataConnectionResult)(nil)).Elem()
}

func (o LookupEventGridDataConnectionResultOutput) ToLookupEventGridDataConnectionResultOutput() LookupEventGridDataConnectionResultOutput {
	return o
}

func (o LookupEventGridDataConnectionResultOutput) ToLookupEventGridDataConnectionResultOutputWithContext(ctx context.Context) LookupEventGridDataConnectionResultOutput {
	return o
}

func (o LookupEventGridDataConnectionResultOutput) BlobStorageEventType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *string { return v.BlobStorageEventType }).(pulumi.StringPtrOutput)
}

func (o LookupEventGridDataConnectionResultOutput) ConsumerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.ConsumerGroup }).(pulumi.StringOutput)
}

func (o LookupEventGridDataConnectionResultOutput) DataFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *string { return v.DataFormat }).(pulumi.StringPtrOutput)
}

func (o LookupEventGridDataConnectionResultOutput) EventHubResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.EventHubResourceId }).(pulumi.StringOutput)
}

func (o LookupEventGridDataConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEventGridDataConnectionResultOutput) IgnoreFirstRecord() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *bool { return v.IgnoreFirstRecord }).(pulumi.BoolPtrOutput)
}

func (o LookupEventGridDataConnectionResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupEventGridDataConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupEventGridDataConnectionResultOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *string { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

func (o LookupEventGridDataConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEventGridDataConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupEventGridDataConnectionResultOutput) StorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.StorageAccountResourceId }).(pulumi.StringOutput)
}

func (o LookupEventGridDataConnectionResultOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *string { return v.TableName }).(pulumi.StringPtrOutput)
}

func (o LookupEventGridDataConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventGridDataConnectionResultOutput{})
}
