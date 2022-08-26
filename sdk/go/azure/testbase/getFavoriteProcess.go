


package testbase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFavoriteProcess(ctx *pulumi.Context, args *LookupFavoriteProcessArgs, opts ...pulumi.InvokeOption) (*LookupFavoriteProcessResult, error) {
	var rv LookupFavoriteProcessResult
	err := ctx.Invoke("azure-native:testbase:getFavoriteProcess", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFavoriteProcessArgs struct {
	FavoriteProcessResourceName string `pulumi:"favoriteProcessResourceName"`
	PackageName                 string `pulumi:"packageName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	TestBaseAccountName         string `pulumi:"testBaseAccountName"`
}


type LookupFavoriteProcessResult struct {
	ActualProcessName string             `pulumi:"actualProcessName"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupFavoriteProcessOutput(ctx *pulumi.Context, args LookupFavoriteProcessOutputArgs, opts ...pulumi.InvokeOption) LookupFavoriteProcessResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFavoriteProcessResult, error) {
			args := v.(LookupFavoriteProcessArgs)
			r, err := LookupFavoriteProcess(ctx, &args, opts...)
			var s LookupFavoriteProcessResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFavoriteProcessResultOutput)
}

type LookupFavoriteProcessOutputArgs struct {
	FavoriteProcessResourceName pulumi.StringInput `pulumi:"favoriteProcessResourceName"`
	PackageName                 pulumi.StringInput `pulumi:"packageName"`
	ResourceGroupName           pulumi.StringInput `pulumi:"resourceGroupName"`
	TestBaseAccountName         pulumi.StringInput `pulumi:"testBaseAccountName"`
}

func (LookupFavoriteProcessOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFavoriteProcessArgs)(nil)).Elem()
}


type LookupFavoriteProcessResultOutput struct{ *pulumi.OutputState }

func (LookupFavoriteProcessResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFavoriteProcessResult)(nil)).Elem()
}

func (o LookupFavoriteProcessResultOutput) ToLookupFavoriteProcessResultOutput() LookupFavoriteProcessResultOutput {
	return o
}

func (o LookupFavoriteProcessResultOutput) ToLookupFavoriteProcessResultOutputWithContext(ctx context.Context) LookupFavoriteProcessResultOutput {
	return o
}

func (o LookupFavoriteProcessResultOutput) ActualProcessName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFavoriteProcessResult) string { return v.ActualProcessName }).(pulumi.StringOutput)
}

func (o LookupFavoriteProcessResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFavoriteProcessResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFavoriteProcessResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFavoriteProcessResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFavoriteProcessResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFavoriteProcessResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupFavoriteProcessResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFavoriteProcessResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFavoriteProcessResultOutput{})
}
