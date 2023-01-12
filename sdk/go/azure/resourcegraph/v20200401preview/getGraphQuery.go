


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGraphQuery(ctx *pulumi.Context, args *LookupGraphQueryArgs, opts ...pulumi.InvokeOption) (*LookupGraphQueryResult, error) {
	var rv LookupGraphQueryResult
	err := ctx.Invoke("azure-native:resourcegraph/v20200401preview:getGraphQuery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGraphQueryArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupGraphQueryResult struct {
	Description  *string            `pulumi:"description"`
	Etag         *string            `pulumi:"etag"`
	Id           string             `pulumi:"id"`
	Location     string             `pulumi:"location"`
	Name         string             `pulumi:"name"`
	Query        string             `pulumi:"query"`
	ResultKind   string             `pulumi:"resultKind"`
	SystemData   SystemDataResponse `pulumi:"systemData"`
	Tags         map[string]string  `pulumi:"tags"`
	TimeModified string             `pulumi:"timeModified"`
	Type         string             `pulumi:"type"`
}

func LookupGraphQueryOutput(ctx *pulumi.Context, args LookupGraphQueryOutputArgs, opts ...pulumi.InvokeOption) LookupGraphQueryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGraphQueryResult, error) {
			args := v.(LookupGraphQueryArgs)
			r, err := LookupGraphQuery(ctx, &args, opts...)
			var s LookupGraphQueryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGraphQueryResultOutput)
}

type LookupGraphQueryOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupGraphQueryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphQueryArgs)(nil)).Elem()
}


type LookupGraphQueryResultOutput struct{ *pulumi.OutputState }

func (LookupGraphQueryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphQueryResult)(nil)).Elem()
}

func (o LookupGraphQueryResultOutput) ToLookupGraphQueryResultOutput() LookupGraphQueryResultOutput {
	return o
}

func (o LookupGraphQueryResultOutput) ToLookupGraphQueryResultOutputWithContext(ctx context.Context) LookupGraphQueryResultOutput {
	return o
}

func (o LookupGraphQueryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQueryResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQueryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGraphQueryResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGraphQueryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGraphQueryResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.Query }).(pulumi.StringOutput)
}

func (o LookupGraphQueryResultOutput) ResultKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.ResultKind }).(pulumi.StringOutput)
}

func (o LookupGraphQueryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupGraphQueryResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGraphQueryResultOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.TimeModified }).(pulumi.StringOutput)
}

func (o LookupGraphQueryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGraphQueryResultOutput{})
}
