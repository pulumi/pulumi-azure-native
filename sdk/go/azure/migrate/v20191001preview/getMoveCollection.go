


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMoveCollection(ctx *pulumi.Context, args *LookupMoveCollectionArgs, opts ...pulumi.InvokeOption) (*LookupMoveCollectionResult, error) {
	var rv LookupMoveCollectionResult
	err := ctx.Invoke("azure-native:migrate/v20191001preview:getMoveCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMoveCollectionArgs struct {
	MoveCollectionName string `pulumi:"moveCollectionName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupMoveCollectionResult struct {
	Etag       string                           `pulumi:"etag"`
	Id         string                           `pulumi:"id"`
	Identity   *IdentityResponse                `pulumi:"identity"`
	Location   *string                          `pulumi:"location"`
	Name       string                           `pulumi:"name"`
	Properties MoveCollectionPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                `pulumi:"tags"`
	Type       string                           `pulumi:"type"`
}

func LookupMoveCollectionOutput(ctx *pulumi.Context, args LookupMoveCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupMoveCollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMoveCollectionResult, error) {
			args := v.(LookupMoveCollectionArgs)
			r, err := LookupMoveCollection(ctx, &args, opts...)
			var s LookupMoveCollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMoveCollectionResultOutput)
}

type LookupMoveCollectionOutputArgs struct {
	MoveCollectionName pulumi.StringInput `pulumi:"moveCollectionName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMoveCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMoveCollectionArgs)(nil)).Elem()
}


type LookupMoveCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupMoveCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMoveCollectionResult)(nil)).Elem()
}

func (o LookupMoveCollectionResultOutput) ToLookupMoveCollectionResultOutput() LookupMoveCollectionResultOutput {
	return o
}

func (o LookupMoveCollectionResultOutput) ToLookupMoveCollectionResultOutputWithContext(ctx context.Context) LookupMoveCollectionResultOutput {
	return o
}

func (o LookupMoveCollectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMoveCollectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupMoveCollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMoveCollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMoveCollectionResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupMoveCollectionResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupMoveCollectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMoveCollectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMoveCollectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMoveCollectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMoveCollectionResultOutput) Properties() MoveCollectionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMoveCollectionResult) MoveCollectionPropertiesResponse { return v.Properties }).(MoveCollectionPropertiesResponseOutput)
}

func (o LookupMoveCollectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMoveCollectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMoveCollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMoveCollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMoveCollectionResultOutput{})
}
