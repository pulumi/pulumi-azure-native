


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWatchlistItem(ctx *pulumi.Context, args *LookupWatchlistItemArgs, opts ...pulumi.InvokeOption) (*LookupWatchlistItemResult, error) {
	var rv LookupWatchlistItemResult
	err := ctx.Invoke("azure-native:securityinsights/v20220501preview:getWatchlistItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWatchlistItemArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WatchlistAlias    string `pulumi:"watchlistAlias"`
	WatchlistItemId   string `pulumi:"watchlistItemId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWatchlistItemResult struct {
	Created           *string                    `pulumi:"created"`
	CreatedBy         *WatchlistUserInfoResponse `pulumi:"createdBy"`
	EntityMapping     interface{}                `pulumi:"entityMapping"`
	Etag              *string                    `pulumi:"etag"`
	Id                string                     `pulumi:"id"`
	IsDeleted         *bool                      `pulumi:"isDeleted"`
	ItemsKeyValue     interface{}                `pulumi:"itemsKeyValue"`
	Name              string                     `pulumi:"name"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	TenantId          *string                    `pulumi:"tenantId"`
	Type              string                     `pulumi:"type"`
	Updated           *string                    `pulumi:"updated"`
	UpdatedBy         *WatchlistUserInfoResponse `pulumi:"updatedBy"`
	WatchlistItemId   *string                    `pulumi:"watchlistItemId"`
	WatchlistItemType *string                    `pulumi:"watchlistItemType"`
}

func LookupWatchlistItemOutput(ctx *pulumi.Context, args LookupWatchlistItemOutputArgs, opts ...pulumi.InvokeOption) LookupWatchlistItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWatchlistItemResult, error) {
			args := v.(LookupWatchlistItemArgs)
			r, err := LookupWatchlistItem(ctx, &args, opts...)
			var s LookupWatchlistItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWatchlistItemResultOutput)
}

type LookupWatchlistItemOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WatchlistAlias    pulumi.StringInput `pulumi:"watchlistAlias"`
	WatchlistItemId   pulumi.StringInput `pulumi:"watchlistItemId"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWatchlistItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWatchlistItemArgs)(nil)).Elem()
}


type LookupWatchlistItemResultOutput struct{ *pulumi.OutputState }

func (LookupWatchlistItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWatchlistItemResult)(nil)).Elem()
}

func (o LookupWatchlistItemResultOutput) ToLookupWatchlistItemResultOutput() LookupWatchlistItemResultOutput {
	return o
}

func (o LookupWatchlistItemResultOutput) ToLookupWatchlistItemResultOutputWithContext(ctx context.Context) LookupWatchlistItemResultOutput {
	return o
}

func (o LookupWatchlistItemResultOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *string { return v.Created }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistItemResultOutput) CreatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *WatchlistUserInfoResponse { return v.CreatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

func (o LookupWatchlistItemResultOutput) EntityMapping() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) interface{} { return v.EntityMapping }).(pulumi.AnyOutput)
}

func (o LookupWatchlistItemResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistItemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWatchlistItemResultOutput) IsDeleted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *bool { return v.IsDeleted }).(pulumi.BoolPtrOutput)
}

func (o LookupWatchlistItemResultOutput) ItemsKeyValue() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) interface{} { return v.ItemsKeyValue }).(pulumi.AnyOutput)
}

func (o LookupWatchlistItemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWatchlistItemResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWatchlistItemResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistItemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWatchlistItemResultOutput) Updated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *string { return v.Updated }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistItemResultOutput) UpdatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *WatchlistUserInfoResponse { return v.UpdatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

func (o LookupWatchlistItemResultOutput) WatchlistItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *string { return v.WatchlistItemId }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistItemResultOutput) WatchlistItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *string { return v.WatchlistItemType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWatchlistItemResultOutput{})
}
