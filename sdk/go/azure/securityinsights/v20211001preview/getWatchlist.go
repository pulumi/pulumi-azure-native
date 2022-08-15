


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWatchlist(ctx *pulumi.Context, args *LookupWatchlistArgs, opts ...pulumi.InvokeOption) (*LookupWatchlistResult, error) {
	var rv LookupWatchlistResult
	err := ctx.Invoke("azure-native:securityinsights/v20211001preview:getWatchlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWatchlistArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WatchlistAlias    string `pulumi:"watchlistAlias"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWatchlistResult struct {
	ContentType         *string                    `pulumi:"contentType"`
	Created             *string                    `pulumi:"created"`
	CreatedBy           *WatchlistUserInfoResponse `pulumi:"createdBy"`
	DefaultDuration     *string                    `pulumi:"defaultDuration"`
	Description         *string                    `pulumi:"description"`
	DisplayName         string                     `pulumi:"displayName"`
	Etag                *string                    `pulumi:"etag"`
	Id                  string                     `pulumi:"id"`
	IsDeleted           *bool                      `pulumi:"isDeleted"`
	ItemsSearchKey      string                     `pulumi:"itemsSearchKey"`
	Labels              []string                   `pulumi:"labels"`
	Name                string                     `pulumi:"name"`
	NumberOfLinesToSkip *int                       `pulumi:"numberOfLinesToSkip"`
	Provider            string                     `pulumi:"provider"`
	RawContent          *string                    `pulumi:"rawContent"`
	Source              string                     `pulumi:"source"`
	SystemData          SystemDataResponse         `pulumi:"systemData"`
	TenantId            *string                    `pulumi:"tenantId"`
	Type                string                     `pulumi:"type"`
	Updated             *string                    `pulumi:"updated"`
	UpdatedBy           *WatchlistUserInfoResponse `pulumi:"updatedBy"`
	UploadStatus        *string                    `pulumi:"uploadStatus"`
	WatchlistAlias      *string                    `pulumi:"watchlistAlias"`
	WatchlistId         *string                    `pulumi:"watchlistId"`
	WatchlistItemsCount *int                       `pulumi:"watchlistItemsCount"`
	WatchlistType       *string                    `pulumi:"watchlistType"`
}

func LookupWatchlistOutput(ctx *pulumi.Context, args LookupWatchlistOutputArgs, opts ...pulumi.InvokeOption) LookupWatchlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWatchlistResult, error) {
			args := v.(LookupWatchlistArgs)
			r, err := LookupWatchlist(ctx, &args, opts...)
			var s LookupWatchlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWatchlistResultOutput)
}

type LookupWatchlistOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WatchlistAlias    pulumi.StringInput `pulumi:"watchlistAlias"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWatchlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWatchlistArgs)(nil)).Elem()
}


type LookupWatchlistResultOutput struct{ *pulumi.OutputState }

func (LookupWatchlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWatchlistResult)(nil)).Elem()
}

func (o LookupWatchlistResultOutput) ToLookupWatchlistResultOutput() LookupWatchlistResultOutput {
	return o
}

func (o LookupWatchlistResultOutput) ToLookupWatchlistResultOutputWithContext(ctx context.Context) LookupWatchlistResultOutput {
	return o
}

func (o LookupWatchlistResultOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistResultOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.Created }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistResultOutput) CreatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *WatchlistUserInfoResponse { return v.CreatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

func (o LookupWatchlistResultOutput) DefaultDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.DefaultDuration }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupWatchlistResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWatchlistResultOutput) IsDeleted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *bool { return v.IsDeleted }).(pulumi.BoolPtrOutput)
}

func (o LookupWatchlistResultOutput) ItemsSearchKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistResult) string { return v.ItemsSearchKey }).(pulumi.StringOutput)
}

func (o LookupWatchlistResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWatchlistResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o LookupWatchlistResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWatchlistResultOutput) NumberOfLinesToSkip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *int { return v.NumberOfLinesToSkip }).(pulumi.IntPtrOutput)
}

func (o LookupWatchlistResultOutput) Provider() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistResult) string { return v.Provider }).(pulumi.StringOutput)
}

func (o LookupWatchlistResultOutput) RawContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.RawContent }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistResult) string { return v.Source }).(pulumi.StringOutput)
}

func (o LookupWatchlistResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWatchlistResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWatchlistResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWatchlistResultOutput) Updated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.Updated }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistResultOutput) UpdatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *WatchlistUserInfoResponse { return v.UpdatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

func (o LookupWatchlistResultOutput) UploadStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.UploadStatus }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistResultOutput) WatchlistAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.WatchlistAlias }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistResultOutput) WatchlistId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.WatchlistId }).(pulumi.StringPtrOutput)
}

func (o LookupWatchlistResultOutput) WatchlistItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *int { return v.WatchlistItemsCount }).(pulumi.IntPtrOutput)
}

func (o LookupWatchlistResultOutput) WatchlistType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.WatchlistType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWatchlistResultOutput{})
}
