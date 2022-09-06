


package v20210801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSyncMember(ctx *pulumi.Context, args *LookupSyncMemberArgs, opts ...pulumi.InvokeOption) (*LookupSyncMemberResult, error) {
	var rv LookupSyncMemberResult
	err := ctx.Invoke("azure-native:sql/v20210801preview:getSyncMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncMemberArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	SyncGroupName     string `pulumi:"syncGroupName"`
	SyncMemberName    string `pulumi:"syncMemberName"`
}


type LookupSyncMemberResult struct {
	DatabaseName                      *string `pulumi:"databaseName"`
	DatabaseType                      *string `pulumi:"databaseType"`
	Id                                string  `pulumi:"id"`
	Name                              string  `pulumi:"name"`
	PrivateEndpointName               string  `pulumi:"privateEndpointName"`
	ServerName                        *string `pulumi:"serverName"`
	SqlServerDatabaseId               *string `pulumi:"sqlServerDatabaseId"`
	SyncAgentId                       *string `pulumi:"syncAgentId"`
	SyncDirection                     *string `pulumi:"syncDirection"`
	SyncMemberAzureDatabaseResourceId *string `pulumi:"syncMemberAzureDatabaseResourceId"`
	SyncState                         string  `pulumi:"syncState"`
	Type                              string  `pulumi:"type"`
	UsePrivateLinkConnection          *bool   `pulumi:"usePrivateLinkConnection"`
	UserName                          *string `pulumi:"userName"`
}

func LookupSyncMemberOutput(ctx *pulumi.Context, args LookupSyncMemberOutputArgs, opts ...pulumi.InvokeOption) LookupSyncMemberResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSyncMemberResult, error) {
			args := v.(LookupSyncMemberArgs)
			r, err := LookupSyncMember(ctx, &args, opts...)
			var s LookupSyncMemberResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSyncMemberResultOutput)
}

type LookupSyncMemberOutputArgs struct {
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
	SyncGroupName     pulumi.StringInput `pulumi:"syncGroupName"`
	SyncMemberName    pulumi.StringInput `pulumi:"syncMemberName"`
}

func (LookupSyncMemberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncMemberArgs)(nil)).Elem()
}


type LookupSyncMemberResultOutput struct{ *pulumi.OutputState }

func (LookupSyncMemberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncMemberResult)(nil)).Elem()
}

func (o LookupSyncMemberResultOutput) ToLookupSyncMemberResultOutput() LookupSyncMemberResultOutput {
	return o
}

func (o LookupSyncMemberResultOutput) ToLookupSyncMemberResultOutputWithContext(ctx context.Context) LookupSyncMemberResultOutput {
	return o
}

func (o LookupSyncMemberResultOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o LookupSyncMemberResultOutput) DatabaseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) *string { return v.DatabaseType }).(pulumi.StringPtrOutput)
}

func (o LookupSyncMemberResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSyncMemberResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSyncMemberResultOutput) PrivateEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) string { return v.PrivateEndpointName }).(pulumi.StringOutput)
}

func (o LookupSyncMemberResultOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o LookupSyncMemberResultOutput) SqlServerDatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) *string { return v.SqlServerDatabaseId }).(pulumi.StringPtrOutput)
}

func (o LookupSyncMemberResultOutput) SyncAgentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) *string { return v.SyncAgentId }).(pulumi.StringPtrOutput)
}

func (o LookupSyncMemberResultOutput) SyncDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) *string { return v.SyncDirection }).(pulumi.StringPtrOutput)
}

func (o LookupSyncMemberResultOutput) SyncMemberAzureDatabaseResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) *string { return v.SyncMemberAzureDatabaseResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupSyncMemberResultOutput) SyncState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) string { return v.SyncState }).(pulumi.StringOutput)
}

func (o LookupSyncMemberResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSyncMemberResultOutput) UsePrivateLinkConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) *bool { return v.UsePrivateLinkConnection }).(pulumi.BoolPtrOutput)
}

func (o LookupSyncMemberResultOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncMemberResult) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSyncMemberResultOutput{})
}
