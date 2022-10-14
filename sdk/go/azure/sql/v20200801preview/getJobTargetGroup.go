


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobTargetGroup(ctx *pulumi.Context, args *LookupJobTargetGroupArgs, opts ...pulumi.InvokeOption) (*LookupJobTargetGroupResult, error) {
	var rv LookupJobTargetGroupResult
	err := ctx.Invoke("azure-native:sql/v20200801preview:getJobTargetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobTargetGroupArgs struct {
	JobAgentName      string `pulumi:"jobAgentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	TargetGroupName   string `pulumi:"targetGroupName"`
}


type LookupJobTargetGroupResult struct {
	Id      string              `pulumi:"id"`
	Members []JobTargetResponse `pulumi:"members"`
	Name    string              `pulumi:"name"`
	Type    string              `pulumi:"type"`
}

func LookupJobTargetGroupOutput(ctx *pulumi.Context, args LookupJobTargetGroupOutputArgs, opts ...pulumi.InvokeOption) LookupJobTargetGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobTargetGroupResult, error) {
			args := v.(LookupJobTargetGroupArgs)
			r, err := LookupJobTargetGroup(ctx, &args, opts...)
			var s LookupJobTargetGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobTargetGroupResultOutput)
}

type LookupJobTargetGroupOutputArgs struct {
	JobAgentName      pulumi.StringInput `pulumi:"jobAgentName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
	TargetGroupName   pulumi.StringInput `pulumi:"targetGroupName"`
}

func (LookupJobTargetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobTargetGroupArgs)(nil)).Elem()
}


type LookupJobTargetGroupResultOutput struct{ *pulumi.OutputState }

func (LookupJobTargetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobTargetGroupResult)(nil)).Elem()
}

func (o LookupJobTargetGroupResultOutput) ToLookupJobTargetGroupResultOutput() LookupJobTargetGroupResultOutput {
	return o
}

func (o LookupJobTargetGroupResultOutput) ToLookupJobTargetGroupResultOutputWithContext(ctx context.Context) LookupJobTargetGroupResultOutput {
	return o
}

func (o LookupJobTargetGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobTargetGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJobTargetGroupResultOutput) Members() JobTargetResponseArrayOutput {
	return o.ApplyT(func(v LookupJobTargetGroupResult) []JobTargetResponse { return v.Members }).(JobTargetResponseArrayOutput)
}

func (o LookupJobTargetGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobTargetGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupJobTargetGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobTargetGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobTargetGroupResultOutput{})
}
