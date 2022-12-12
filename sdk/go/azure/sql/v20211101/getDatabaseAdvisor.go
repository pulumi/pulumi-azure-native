


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseAdvisor(ctx *pulumi.Context, args *LookupDatabaseAdvisorArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAdvisorResult, error) {
	var rv LookupDatabaseAdvisorResult
	err := ctx.Invoke("azure-native:sql/v20211101:getDatabaseAdvisor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAdvisorArgs struct {
	AdvisorName       string `pulumi:"advisorName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupDatabaseAdvisorResult struct {
	AdvisorStatus                  string                      `pulumi:"advisorStatus"`
	AutoExecuteStatus              string                      `pulumi:"autoExecuteStatus"`
	AutoExecuteStatusInheritedFrom string                      `pulumi:"autoExecuteStatusInheritedFrom"`
	Id                             string                      `pulumi:"id"`
	Kind                           string                      `pulumi:"kind"`
	LastChecked                    string                      `pulumi:"lastChecked"`
	Location                       string                      `pulumi:"location"`
	Name                           string                      `pulumi:"name"`
	RecommendationsStatus          string                      `pulumi:"recommendationsStatus"`
	RecommendedActions             []RecommendedActionResponse `pulumi:"recommendedActions"`
	Type                           string                      `pulumi:"type"`
}

func LookupDatabaseAdvisorOutput(ctx *pulumi.Context, args LookupDatabaseAdvisorOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAdvisorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAdvisorResult, error) {
			args := v.(LookupDatabaseAdvisorArgs)
			r, err := LookupDatabaseAdvisor(ctx, &args, opts...)
			var s LookupDatabaseAdvisorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAdvisorResultOutput)
}

type LookupDatabaseAdvisorOutputArgs struct {
	AdvisorName       pulumi.StringInput `pulumi:"advisorName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupDatabaseAdvisorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAdvisorArgs)(nil)).Elem()
}


type LookupDatabaseAdvisorResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAdvisorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAdvisorResult)(nil)).Elem()
}

func (o LookupDatabaseAdvisorResultOutput) ToLookupDatabaseAdvisorResultOutput() LookupDatabaseAdvisorResultOutput {
	return o
}

func (o LookupDatabaseAdvisorResultOutput) ToLookupDatabaseAdvisorResultOutputWithContext(ctx context.Context) LookupDatabaseAdvisorResultOutput {
	return o
}

func (o LookupDatabaseAdvisorResultOutput) AdvisorStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.AdvisorStatus }).(pulumi.StringOutput)
}

func (o LookupDatabaseAdvisorResultOutput) AutoExecuteStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.AutoExecuteStatus }).(pulumi.StringOutput)
}

func (o LookupDatabaseAdvisorResultOutput) AutoExecuteStatusInheritedFrom() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.AutoExecuteStatusInheritedFrom }).(pulumi.StringOutput)
}

func (o LookupDatabaseAdvisorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseAdvisorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDatabaseAdvisorResultOutput) LastChecked() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.LastChecked }).(pulumi.StringOutput)
}

func (o LookupDatabaseAdvisorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDatabaseAdvisorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseAdvisorResultOutput) RecommendationsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.RecommendationsStatus }).(pulumi.StringOutput)
}

func (o LookupDatabaseAdvisorResultOutput) RecommendedActions() RecommendedActionResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) []RecommendedActionResponse { return v.RecommendedActions }).(RecommendedActionResponseArrayOutput)
}

func (o LookupDatabaseAdvisorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAdvisorResultOutput{})
}
