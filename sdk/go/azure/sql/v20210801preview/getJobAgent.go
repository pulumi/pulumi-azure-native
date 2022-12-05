


package v20210801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobAgent(ctx *pulumi.Context, args *LookupJobAgentArgs, opts ...pulumi.InvokeOption) (*LookupJobAgentResult, error) {
	var rv LookupJobAgentResult
	err := ctx.Invoke("azure-native:sql/v20210801preview:getJobAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobAgentArgs struct {
	JobAgentName      string `pulumi:"jobAgentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupJobAgentResult struct {
	DatabaseId string            `pulumi:"databaseId"`
	Id         string            `pulumi:"id"`
	Location   string            `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Sku        *SkuResponse      `pulumi:"sku"`
	State      string            `pulumi:"state"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func LookupJobAgentOutput(ctx *pulumi.Context, args LookupJobAgentOutputArgs, opts ...pulumi.InvokeOption) LookupJobAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobAgentResult, error) {
			args := v.(LookupJobAgentArgs)
			r, err := LookupJobAgent(ctx, &args, opts...)
			var s LookupJobAgentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobAgentResultOutput)
}

type LookupJobAgentOutputArgs struct {
	JobAgentName      pulumi.StringInput `pulumi:"jobAgentName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupJobAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobAgentArgs)(nil)).Elem()
}


type LookupJobAgentResultOutput struct{ *pulumi.OutputState }

func (LookupJobAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobAgentResult)(nil)).Elem()
}

func (o LookupJobAgentResultOutput) ToLookupJobAgentResultOutput() LookupJobAgentResultOutput {
	return o
}

func (o LookupJobAgentResultOutput) ToLookupJobAgentResultOutputWithContext(ctx context.Context) LookupJobAgentResultOutput {
	return o
}

func (o LookupJobAgentResultOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobAgentResult) string { return v.DatabaseId }).(pulumi.StringOutput)
}

func (o LookupJobAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJobAgentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobAgentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupJobAgentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobAgentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupJobAgentResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupJobAgentResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupJobAgentResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobAgentResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupJobAgentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobAgentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupJobAgentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobAgentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobAgentResultOutput{})
}
