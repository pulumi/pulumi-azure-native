


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticMember(ctx *pulumi.Context, args *LookupStaticMemberArgs, opts ...pulumi.InvokeOption) (*LookupStaticMemberResult, error) {
	var rv LookupStaticMemberResult
	err := ctx.Invoke("azure-native:network/v20220201preview:getStaticMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticMemberArgs struct {
	NetworkGroupName   string `pulumi:"networkGroupName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	StaticMemberName   string `pulumi:"staticMemberName"`
}


type LookupStaticMemberResult struct {
	Etag       string             `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Name       string             `pulumi:"name"`
	ResourceId *string            `pulumi:"resourceId"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupStaticMemberOutput(ctx *pulumi.Context, args LookupStaticMemberOutputArgs, opts ...pulumi.InvokeOption) LookupStaticMemberResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticMemberResult, error) {
			args := v.(LookupStaticMemberArgs)
			r, err := LookupStaticMember(ctx, &args, opts...)
			var s LookupStaticMemberResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticMemberResultOutput)
}

type LookupStaticMemberOutputArgs struct {
	NetworkGroupName   pulumi.StringInput `pulumi:"networkGroupName"`
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	StaticMemberName   pulumi.StringInput `pulumi:"staticMemberName"`
}

func (LookupStaticMemberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticMemberArgs)(nil)).Elem()
}


type LookupStaticMemberResultOutput struct{ *pulumi.OutputState }

func (LookupStaticMemberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticMemberResult)(nil)).Elem()
}

func (o LookupStaticMemberResultOutput) ToLookupStaticMemberResultOutput() LookupStaticMemberResultOutput {
	return o
}

func (o LookupStaticMemberResultOutput) ToLookupStaticMemberResultOutputWithContext(ctx context.Context) LookupStaticMemberResultOutput {
	return o
}

func (o LookupStaticMemberResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupStaticMemberResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStaticMemberResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStaticMemberResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupStaticMemberResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupStaticMemberResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticMemberResultOutput{})
}
