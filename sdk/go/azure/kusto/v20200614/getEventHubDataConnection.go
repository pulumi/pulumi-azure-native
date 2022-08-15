


package v20200614

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventHubDataConnection(ctx *pulumi.Context, args *LookupEventHubDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEventHubDataConnectionResult, error) {
	var rv LookupEventHubDataConnectionResult
	err := ctx.Invoke("azure-native:kusto/v20200614:getEventHubDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubDataConnectionArgs struct {
	ClusterName        string `pulumi:"clusterName"`
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupEventHubDataConnectionResult struct {
	Compression           *string  `pulumi:"compression"`
	ConsumerGroup         string   `pulumi:"consumerGroup"`
	DataFormat            *string  `pulumi:"dataFormat"`
	EventHubResourceId    string   `pulumi:"eventHubResourceId"`
	EventSystemProperties []string `pulumi:"eventSystemProperties"`
	Id                    string   `pulumi:"id"`
	Kind                  string   `pulumi:"kind"`
	Location              *string  `pulumi:"location"`
	MappingRuleName       *string  `pulumi:"mappingRuleName"`
	Name                  string   `pulumi:"name"`
	TableName             *string  `pulumi:"tableName"`
	Type                  string   `pulumi:"type"`
}

func LookupEventHubDataConnectionOutput(ctx *pulumi.Context, args LookupEventHubDataConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupEventHubDataConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventHubDataConnectionResult, error) {
			args := v.(LookupEventHubDataConnectionArgs)
			r, err := LookupEventHubDataConnection(ctx, &args, opts...)
			var s LookupEventHubDataConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventHubDataConnectionResultOutput)
}

type LookupEventHubDataConnectionOutputArgs struct {
	ClusterName        pulumi.StringInput `pulumi:"clusterName"`
	DataConnectionName pulumi.StringInput `pulumi:"dataConnectionName"`
	DatabaseName       pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventHubDataConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubDataConnectionArgs)(nil)).Elem()
}


type LookupEventHubDataConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupEventHubDataConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubDataConnectionResult)(nil)).Elem()
}

func (o LookupEventHubDataConnectionResultOutput) ToLookupEventHubDataConnectionResultOutput() LookupEventHubDataConnectionResultOutput {
	return o
}

func (o LookupEventHubDataConnectionResultOutput) ToLookupEventHubDataConnectionResultOutputWithContext(ctx context.Context) LookupEventHubDataConnectionResultOutput {
	return o
}

func (o LookupEventHubDataConnectionResultOutput) Compression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) *string { return v.Compression }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubDataConnectionResultOutput) ConsumerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.ConsumerGroup }).(pulumi.StringOutput)
}

func (o LookupEventHubDataConnectionResultOutput) DataFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) *string { return v.DataFormat }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubDataConnectionResultOutput) EventHubResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.EventHubResourceId }).(pulumi.StringOutput)
}

func (o LookupEventHubDataConnectionResultOutput) EventSystemProperties() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) []string { return v.EventSystemProperties }).(pulumi.StringArrayOutput)
}

func (o LookupEventHubDataConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEventHubDataConnectionResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupEventHubDataConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubDataConnectionResultOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) *string { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubDataConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEventHubDataConnectionResultOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) *string { return v.TableName }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubDataConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventHubDataConnectionResultOutput{})
}
