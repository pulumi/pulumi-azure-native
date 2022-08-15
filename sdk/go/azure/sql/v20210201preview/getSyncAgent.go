


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSyncAgent(ctx *pulumi.Context, args *LookupSyncAgentArgs, opts ...pulumi.InvokeOption) (*LookupSyncAgentResult, error) {
	var rv LookupSyncAgentResult
	err := ctx.Invoke("azure-native:sql/v20210201preview:getSyncAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncAgentArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	SyncAgentName     string `pulumi:"syncAgentName"`
}


type LookupSyncAgentResult struct {
	ExpiryTime     string  `pulumi:"expiryTime"`
	Id             string  `pulumi:"id"`
	IsUpToDate     bool    `pulumi:"isUpToDate"`
	LastAliveTime  string  `pulumi:"lastAliveTime"`
	Name           string  `pulumi:"name"`
	State          string  `pulumi:"state"`
	SyncDatabaseId *string `pulumi:"syncDatabaseId"`
	Type           string  `pulumi:"type"`
	Version        string  `pulumi:"version"`
}

func LookupSyncAgentOutput(ctx *pulumi.Context, args LookupSyncAgentOutputArgs, opts ...pulumi.InvokeOption) LookupSyncAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSyncAgentResult, error) {
			args := v.(LookupSyncAgentArgs)
			r, err := LookupSyncAgent(ctx, &args, opts...)
			var s LookupSyncAgentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSyncAgentResultOutput)
}

type LookupSyncAgentOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
	SyncAgentName     pulumi.StringInput `pulumi:"syncAgentName"`
}

func (LookupSyncAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncAgentArgs)(nil)).Elem()
}


type LookupSyncAgentResultOutput struct{ *pulumi.OutputState }

func (LookupSyncAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncAgentResult)(nil)).Elem()
}

func (o LookupSyncAgentResultOutput) ToLookupSyncAgentResultOutput() LookupSyncAgentResultOutput {
	return o
}

func (o LookupSyncAgentResultOutput) ToLookupSyncAgentResultOutputWithContext(ctx context.Context) LookupSyncAgentResultOutput {
	return o
}

func (o LookupSyncAgentResultOutput) ExpiryTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncAgentResult) string { return v.ExpiryTime }).(pulumi.StringOutput)
}

func (o LookupSyncAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSyncAgentResultOutput) IsUpToDate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSyncAgentResult) bool { return v.IsUpToDate }).(pulumi.BoolOutput)
}

func (o LookupSyncAgentResultOutput) LastAliveTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncAgentResult) string { return v.LastAliveTime }).(pulumi.StringOutput)
}

func (o LookupSyncAgentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncAgentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSyncAgentResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncAgentResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupSyncAgentResultOutput) SyncDatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncAgentResult) *string { return v.SyncDatabaseId }).(pulumi.StringPtrOutput)
}

func (o LookupSyncAgentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncAgentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSyncAgentResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncAgentResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSyncAgentResultOutput{})
}
