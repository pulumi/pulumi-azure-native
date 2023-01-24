


package v20220904

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSyncSet(ctx *pulumi.Context, args *LookupSyncSetArgs, opts ...pulumi.InvokeOption) (*LookupSyncSetResult, error) {
	var rv LookupSyncSetResult
	err := ctx.Invoke("azure-native:redhatopenshift/v20220904:getSyncSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncSetArgs struct {
	ChildResourceName string `pulumi:"childResourceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupSyncSetResult struct {
	Id         string             `pulumi:"id"`
	Name       string             `pulumi:"name"`
	Resources  *string            `pulumi:"resources"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupSyncSetOutput(ctx *pulumi.Context, args LookupSyncSetOutputArgs, opts ...pulumi.InvokeOption) LookupSyncSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSyncSetResult, error) {
			args := v.(LookupSyncSetArgs)
			r, err := LookupSyncSet(ctx, &args, opts...)
			var s LookupSyncSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSyncSetResultOutput)
}

type LookupSyncSetOutputArgs struct {
	ChildResourceName pulumi.StringInput `pulumi:"childResourceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSyncSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncSetArgs)(nil)).Elem()
}


type LookupSyncSetResultOutput struct{ *pulumi.OutputState }

func (LookupSyncSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncSetResult)(nil)).Elem()
}

func (o LookupSyncSetResultOutput) ToLookupSyncSetResultOutput() LookupSyncSetResultOutput {
	return o
}

func (o LookupSyncSetResultOutput) ToLookupSyncSetResultOutputWithContext(ctx context.Context) LookupSyncSetResultOutput {
	return o
}

func (o LookupSyncSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSyncSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSyncSetResultOutput) Resources() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncSetResult) *string { return v.Resources }).(pulumi.StringPtrOutput)
}

func (o LookupSyncSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSyncSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSyncSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSyncSetResultOutput{})
}
