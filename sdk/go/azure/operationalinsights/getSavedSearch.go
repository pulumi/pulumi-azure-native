


package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSavedSearch(ctx *pulumi.Context, args *LookupSavedSearchArgs, opts ...pulumi.InvokeOption) (*LookupSavedSearchResult, error) {
	var rv LookupSavedSearchResult
	err := ctx.Invoke("azure-native:operationalinsights:getSavedSearch", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSavedSearchArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SavedSearchId     string `pulumi:"savedSearchId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupSavedSearchResult struct {
	Category           string        `pulumi:"category"`
	DisplayName        string        `pulumi:"displayName"`
	Etag               *string       `pulumi:"etag"`
	FunctionAlias      *string       `pulumi:"functionAlias"`
	FunctionParameters *string       `pulumi:"functionParameters"`
	Id                 string        `pulumi:"id"`
	Name               string        `pulumi:"name"`
	Query              string        `pulumi:"query"`
	Tags               []TagResponse `pulumi:"tags"`
	Type               string        `pulumi:"type"`
	Version            *float64      `pulumi:"version"`
}

func LookupSavedSearchOutput(ctx *pulumi.Context, args LookupSavedSearchOutputArgs, opts ...pulumi.InvokeOption) LookupSavedSearchResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSavedSearchResult, error) {
			args := v.(LookupSavedSearchArgs)
			r, err := LookupSavedSearch(ctx, &args, opts...)
			var s LookupSavedSearchResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSavedSearchResultOutput)
}

type LookupSavedSearchOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SavedSearchId     pulumi.StringInput `pulumi:"savedSearchId"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSavedSearchOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSavedSearchArgs)(nil)).Elem()
}


type LookupSavedSearchResultOutput struct{ *pulumi.OutputState }

func (LookupSavedSearchResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSavedSearchResult)(nil)).Elem()
}

func (o LookupSavedSearchResultOutput) ToLookupSavedSearchResultOutput() LookupSavedSearchResultOutput {
	return o
}

func (o LookupSavedSearchResultOutput) ToLookupSavedSearchResultOutputWithContext(ctx context.Context) LookupSavedSearchResultOutput {
	return o
}

func (o LookupSavedSearchResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSavedSearchResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o LookupSavedSearchResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSavedSearchResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupSavedSearchResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSavedSearchResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupSavedSearchResultOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSavedSearchResult) *string { return v.FunctionAlias }).(pulumi.StringPtrOutput)
}

func (o LookupSavedSearchResultOutput) FunctionParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSavedSearchResult) *string { return v.FunctionParameters }).(pulumi.StringPtrOutput)
}

func (o LookupSavedSearchResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSavedSearchResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSavedSearchResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSavedSearchResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSavedSearchResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSavedSearchResult) string { return v.Query }).(pulumi.StringOutput)
}

func (o LookupSavedSearchResultOutput) Tags() TagResponseArrayOutput {
	return o.ApplyT(func(v LookupSavedSearchResult) []TagResponse { return v.Tags }).(TagResponseArrayOutput)
}

func (o LookupSavedSearchResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSavedSearchResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSavedSearchResultOutput) Version() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupSavedSearchResult) *float64 { return v.Version }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSavedSearchResultOutput{})
}
