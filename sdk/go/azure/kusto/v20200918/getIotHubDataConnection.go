


package v20200918

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotHubDataConnection(ctx *pulumi.Context, args *LookupIotHubDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupIotHubDataConnectionResult, error) {
	var rv LookupIotHubDataConnectionResult
	err := ctx.Invoke("azure-native:kusto/v20200918:getIotHubDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotHubDataConnectionArgs struct {
	ClusterName        string `pulumi:"clusterName"`
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupIotHubDataConnectionResult struct {
	ConsumerGroup          string   `pulumi:"consumerGroup"`
	DataFormat             *string  `pulumi:"dataFormat"`
	EventSystemProperties  []string `pulumi:"eventSystemProperties"`
	Id                     string   `pulumi:"id"`
	IotHubResourceId       string   `pulumi:"iotHubResourceId"`
	Kind                   string   `pulumi:"kind"`
	Location               *string  `pulumi:"location"`
	MappingRuleName        *string  `pulumi:"mappingRuleName"`
	Name                   string   `pulumi:"name"`
	ProvisioningState      string   `pulumi:"provisioningState"`
	SharedAccessPolicyName string   `pulumi:"sharedAccessPolicyName"`
	TableName              *string  `pulumi:"tableName"`
	Type                   string   `pulumi:"type"`
}

func LookupIotHubDataConnectionOutput(ctx *pulumi.Context, args LookupIotHubDataConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupIotHubDataConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotHubDataConnectionResult, error) {
			args := v.(LookupIotHubDataConnectionArgs)
			r, err := LookupIotHubDataConnection(ctx, &args, opts...)
			var s LookupIotHubDataConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotHubDataConnectionResultOutput)
}

type LookupIotHubDataConnectionOutputArgs struct {
	ClusterName        pulumi.StringInput `pulumi:"clusterName"`
	DataConnectionName pulumi.StringInput `pulumi:"dataConnectionName"`
	DatabaseName       pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIotHubDataConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubDataConnectionArgs)(nil)).Elem()
}


type LookupIotHubDataConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupIotHubDataConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubDataConnectionResult)(nil)).Elem()
}

func (o LookupIotHubDataConnectionResultOutput) ToLookupIotHubDataConnectionResultOutput() LookupIotHubDataConnectionResultOutput {
	return o
}

func (o LookupIotHubDataConnectionResultOutput) ToLookupIotHubDataConnectionResultOutputWithContext(ctx context.Context) LookupIotHubDataConnectionResultOutput {
	return o
}

func (o LookupIotHubDataConnectionResultOutput) ConsumerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.ConsumerGroup }).(pulumi.StringOutput)
}

func (o LookupIotHubDataConnectionResultOutput) DataFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) *string { return v.DataFormat }).(pulumi.StringPtrOutput)
}

func (o LookupIotHubDataConnectionResultOutput) EventSystemProperties() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) []string { return v.EventSystemProperties }).(pulumi.StringArrayOutput)
}

func (o LookupIotHubDataConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIotHubDataConnectionResultOutput) IotHubResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.IotHubResourceId }).(pulumi.StringOutput)
}

func (o LookupIotHubDataConnectionResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupIotHubDataConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIotHubDataConnectionResultOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) *string { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

func (o LookupIotHubDataConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIotHubDataConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupIotHubDataConnectionResultOutput) SharedAccessPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.SharedAccessPolicyName }).(pulumi.StringOutput)
}

func (o LookupIotHubDataConnectionResultOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) *string { return v.TableName }).(pulumi.StringPtrOutput)
}

func (o LookupIotHubDataConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotHubDataConnectionResultOutput{})
}
