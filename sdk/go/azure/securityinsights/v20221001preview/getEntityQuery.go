


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupEntityQuery(ctx *pulumi.Context, args *LookupEntityQueryArgs, opts ...pulumi.InvokeOption) (*LookupEntityQueryResult, error) {
	var rv LookupEntityQueryResult
	err := ctx.Invoke("azure-native:securityinsights/v20221001preview:getEntityQuery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEntityQueryArgs struct {
	EntityQueryId     string `pulumi:"entityQueryId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupEntityQueryResult struct {
	Etag       *string            `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupEntityQueryOutput(ctx *pulumi.Context, args LookupEntityQueryOutputArgs, opts ...pulumi.InvokeOption) LookupEntityQueryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEntityQueryResult, error) {
			args := v.(LookupEntityQueryArgs)
			r, err := LookupEntityQuery(ctx, &args, opts...)
			var s LookupEntityQueryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEntityQueryResultOutput)
}

type LookupEntityQueryOutputArgs struct {
	EntityQueryId     pulumi.StringInput `pulumi:"entityQueryId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEntityQueryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntityQueryArgs)(nil)).Elem()
}


type LookupEntityQueryResultOutput struct{ *pulumi.OutputState }

func (LookupEntityQueryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntityQueryResult)(nil)).Elem()
}

func (o LookupEntityQueryResultOutput) ToLookupEntityQueryResultOutput() LookupEntityQueryResultOutput {
	return o
}

func (o LookupEntityQueryResultOutput) ToLookupEntityQueryResultOutputWithContext(ctx context.Context) LookupEntityQueryResultOutput {
	return o
}

func (o LookupEntityQueryResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEntityQueryResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupEntityQueryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityQueryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEntityQueryResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityQueryResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupEntityQueryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityQueryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEntityQueryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEntityQueryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEntityQueryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityQueryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEntityQueryResultOutput{})
}
