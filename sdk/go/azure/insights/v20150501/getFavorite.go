


package v20150501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFavorite(ctx *pulumi.Context, args *LookupFavoriteArgs, opts ...pulumi.InvokeOption) (*LookupFavoriteResult, error) {
	var rv LookupFavoriteResult
	err := ctx.Invoke("azure-native:insights/v20150501:getFavorite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFavoriteArgs struct {
	FavoriteId        string `pulumi:"favoriteId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupFavoriteResult struct {
	Category                *string  `pulumi:"category"`
	Config                  *string  `pulumi:"config"`
	FavoriteId              string   `pulumi:"favoriteId"`
	FavoriteType            *string  `pulumi:"favoriteType"`
	IsGeneratedFromTemplate *bool    `pulumi:"isGeneratedFromTemplate"`
	Name                    *string  `pulumi:"name"`
	SourceType              *string  `pulumi:"sourceType"`
	Tags                    []string `pulumi:"tags"`
	TimeModified            string   `pulumi:"timeModified"`
	UserId                  string   `pulumi:"userId"`
	Version                 *string  `pulumi:"version"`
}

func LookupFavoriteOutput(ctx *pulumi.Context, args LookupFavoriteOutputArgs, opts ...pulumi.InvokeOption) LookupFavoriteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFavoriteResult, error) {
			args := v.(LookupFavoriteArgs)
			r, err := LookupFavorite(ctx, &args, opts...)
			var s LookupFavoriteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFavoriteResultOutput)
}

type LookupFavoriteOutputArgs struct {
	FavoriteId        pulumi.StringInput `pulumi:"favoriteId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupFavoriteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFavoriteArgs)(nil)).Elem()
}


type LookupFavoriteResultOutput struct{ *pulumi.OutputState }

func (LookupFavoriteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFavoriteResult)(nil)).Elem()
}

func (o LookupFavoriteResultOutput) ToLookupFavoriteResultOutput() LookupFavoriteResultOutput {
	return o
}

func (o LookupFavoriteResultOutput) ToLookupFavoriteResultOutputWithContext(ctx context.Context) LookupFavoriteResultOutput {
	return o
}

func (o LookupFavoriteResultOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o LookupFavoriteResultOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *string { return v.Config }).(pulumi.StringPtrOutput)
}

func (o LookupFavoriteResultOutput) FavoriteId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFavoriteResult) string { return v.FavoriteId }).(pulumi.StringOutput)
}

func (o LookupFavoriteResultOutput) FavoriteType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *string { return v.FavoriteType }).(pulumi.StringPtrOutput)
}

func (o LookupFavoriteResultOutput) IsGeneratedFromTemplate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *bool { return v.IsGeneratedFromTemplate }).(pulumi.BoolPtrOutput)
}

func (o LookupFavoriteResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupFavoriteResultOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *string { return v.SourceType }).(pulumi.StringPtrOutput)
}

func (o LookupFavoriteResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFavoriteResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupFavoriteResultOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFavoriteResult) string { return v.TimeModified }).(pulumi.StringOutput)
}

func (o LookupFavoriteResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFavoriteResult) string { return v.UserId }).(pulumi.StringOutput)
}

func (o LookupFavoriteResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFavoriteResultOutput{})
}
