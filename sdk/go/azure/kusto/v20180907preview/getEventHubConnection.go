


package v20180907preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventHubConnection(ctx *pulumi.Context, args *LookupEventHubConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEventHubConnectionResult, error) {
	var rv LookupEventHubConnectionResult
	err := ctx.Invoke("azure-native:kusto/v20180907preview:getEventHubConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubConnectionArgs struct {
	ClusterName            string `pulumi:"clusterName"`
	DatabaseName           string `pulumi:"databaseName"`
	EventHubConnectionName string `pulumi:"eventHubConnectionName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupEventHubConnectionResult struct {
	ConsumerGroup      string  `pulumi:"consumerGroup"`
	DataFormat         *string `pulumi:"dataFormat"`
	EventHubResourceId string  `pulumi:"eventHubResourceId"`
	Id                 string  `pulumi:"id"`
	Location           *string `pulumi:"location"`
	MappingRuleName    *string `pulumi:"mappingRuleName"`
	Name               string  `pulumi:"name"`
	TableName          *string `pulumi:"tableName"`
	Type               string  `pulumi:"type"`
}

func LookupEventHubConnectionOutput(ctx *pulumi.Context, args LookupEventHubConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupEventHubConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventHubConnectionResult, error) {
			args := v.(LookupEventHubConnectionArgs)
			r, err := LookupEventHubConnection(ctx, &args, opts...)
			var s LookupEventHubConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventHubConnectionResultOutput)
}

type LookupEventHubConnectionOutputArgs struct {
	ClusterName            pulumi.StringInput `pulumi:"clusterName"`
	DatabaseName           pulumi.StringInput `pulumi:"databaseName"`
	EventHubConnectionName pulumi.StringInput `pulumi:"eventHubConnectionName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventHubConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubConnectionArgs)(nil)).Elem()
}


type LookupEventHubConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupEventHubConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubConnectionResult)(nil)).Elem()
}

func (o LookupEventHubConnectionResultOutput) ToLookupEventHubConnectionResultOutput() LookupEventHubConnectionResultOutput {
	return o
}

func (o LookupEventHubConnectionResultOutput) ToLookupEventHubConnectionResultOutputWithContext(ctx context.Context) LookupEventHubConnectionResultOutput {
	return o
}

func (o LookupEventHubConnectionResultOutput) ConsumerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubConnectionResult) string { return v.ConsumerGroup }).(pulumi.StringOutput)
}

func (o LookupEventHubConnectionResultOutput) DataFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubConnectionResult) *string { return v.DataFormat }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubConnectionResultOutput) EventHubResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubConnectionResult) string { return v.EventHubResourceId }).(pulumi.StringOutput)
}

func (o LookupEventHubConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEventHubConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubConnectionResultOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubConnectionResult) *string { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEventHubConnectionResultOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubConnectionResult) *string { return v.TableName }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventHubConnectionResultOutput{})
}
