


package synapse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventHubDataConnection(ctx *pulumi.Context, args *LookupEventHubDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEventHubDataConnectionResult, error) {
	var rv LookupEventHubDataConnectionResult
	err := ctx.Invoke("azure-native:synapse:getEventHubDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubDataConnectionArgs struct {
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	KustoPoolName      string `pulumi:"kustoPoolName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	WorkspaceName      string `pulumi:"workspaceName"`
}


type LookupEventHubDataConnectionResult struct {
	Compression           *string            `pulumi:"compression"`
	ConsumerGroup         string             `pulumi:"consumerGroup"`
	DataFormat            *string            `pulumi:"dataFormat"`
	EventHubResourceId    string             `pulumi:"eventHubResourceId"`
	EventSystemProperties []string           `pulumi:"eventSystemProperties"`
	Id                    string             `pulumi:"id"`
	Kind                  string             `pulumi:"kind"`
	Location              *string            `pulumi:"location"`
	MappingRuleName       *string            `pulumi:"mappingRuleName"`
	Name                  string             `pulumi:"name"`
	ProvisioningState     string             `pulumi:"provisioningState"`
	SystemData            SystemDataResponse `pulumi:"systemData"`
	TableName             *string            `pulumi:"tableName"`
	Type                  string             `pulumi:"type"`
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
	DataConnectionName pulumi.StringInput `pulumi:"dataConnectionName"`
	DatabaseName       pulumi.StringInput `pulumi:"databaseName"`
	KustoPoolName      pulumi.StringInput `pulumi:"kustoPoolName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName      pulumi.StringInput `pulumi:"workspaceName"`
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

func (o LookupEventHubDataConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupEventHubDataConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
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
