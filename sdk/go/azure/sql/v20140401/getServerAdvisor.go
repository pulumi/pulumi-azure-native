


package v20140401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerAdvisor(ctx *pulumi.Context, args *LookupServerAdvisorArgs, opts ...pulumi.InvokeOption) (*LookupServerAdvisorResult, error) {
	var rv LookupServerAdvisorResult
	err := ctx.Invoke("azure-native:sql/v20140401:getServerAdvisor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerAdvisorArgs struct {
	AdvisorName       string `pulumi:"advisorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerAdvisorResult struct {
	AdvisorStatus         string `pulumi:"advisorStatus"`
	AutoExecuteValue      string `pulumi:"autoExecuteValue"`
	Id                    string `pulumi:"id"`
	Kind                  string `pulumi:"kind"`
	LastChecked           string `pulumi:"lastChecked"`
	Location              string `pulumi:"location"`
	Name                  string `pulumi:"name"`
	RecommendationsStatus string `pulumi:"recommendationsStatus"`
	Type                  string `pulumi:"type"`
}

func LookupServerAdvisorOutput(ctx *pulumi.Context, args LookupServerAdvisorOutputArgs, opts ...pulumi.InvokeOption) LookupServerAdvisorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerAdvisorResult, error) {
			args := v.(LookupServerAdvisorArgs)
			r, err := LookupServerAdvisor(ctx, &args, opts...)
			var s LookupServerAdvisorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerAdvisorResultOutput)
}

type LookupServerAdvisorOutputArgs struct {
	AdvisorName       pulumi.StringInput `pulumi:"advisorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerAdvisorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerAdvisorArgs)(nil)).Elem()
}


type LookupServerAdvisorResultOutput struct{ *pulumi.OutputState }

func (LookupServerAdvisorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerAdvisorResult)(nil)).Elem()
}

func (o LookupServerAdvisorResultOutput) ToLookupServerAdvisorResultOutput() LookupServerAdvisorResultOutput {
	return o
}

func (o LookupServerAdvisorResultOutput) ToLookupServerAdvisorResultOutputWithContext(ctx context.Context) LookupServerAdvisorResultOutput {
	return o
}

func (o LookupServerAdvisorResultOutput) AdvisorStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdvisorResult) string { return v.AdvisorStatus }).(pulumi.StringOutput)
}

func (o LookupServerAdvisorResultOutput) AutoExecuteValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdvisorResult) string { return v.AutoExecuteValue }).(pulumi.StringOutput)
}

func (o LookupServerAdvisorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdvisorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerAdvisorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdvisorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupServerAdvisorResultOutput) LastChecked() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdvisorResult) string { return v.LastChecked }).(pulumi.StringOutput)
}

func (o LookupServerAdvisorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdvisorResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServerAdvisorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdvisorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerAdvisorResultOutput) RecommendationsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdvisorResult) string { return v.RecommendationsStatus }).(pulumi.StringOutput)
}

func (o LookupServerAdvisorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdvisorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerAdvisorResultOutput{})
}
