


package v20180301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupManagementGroup(ctx *pulumi.Context, args *LookupManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupManagementGroupResult, error) {
	var rv LookupManagementGroupResult
	err := ctx.Invoke("azure-native:management/v20180301preview:getManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementGroupArgs struct {
	Expand  *string `pulumi:"expand"`
	Filter  *string `pulumi:"filter"`
	GroupId string  `pulumi:"groupId"`
	Recurse *bool   `pulumi:"recurse"`
}


type LookupManagementGroupResult struct {
	Children    []ManagementGroupChildInfoResponse `pulumi:"children"`
	Details     *ManagementGroupDetailsResponse    `pulumi:"details"`
	DisplayName *string                            `pulumi:"displayName"`
	Id          string                             `pulumi:"id"`
	Name        string                             `pulumi:"name"`
	Roles       []string                           `pulumi:"roles"`
	TenantId    *string                            `pulumi:"tenantId"`
	Type        string                             `pulumi:"type"`
}

func LookupManagementGroupOutput(ctx *pulumi.Context, args LookupManagementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementGroupResult, error) {
			args := v.(LookupManagementGroupArgs)
			r, err := LookupManagementGroup(ctx, &args, opts...)
			var s LookupManagementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementGroupResultOutput)
}

type LookupManagementGroupOutputArgs struct {
	Expand  pulumi.StringPtrInput `pulumi:"expand"`
	Filter  pulumi.StringPtrInput `pulumi:"filter"`
	GroupId pulumi.StringInput    `pulumi:"groupId"`
	Recurse pulumi.BoolPtrInput   `pulumi:"recurse"`
}

func (LookupManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupArgs)(nil)).Elem()
}


type LookupManagementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupResult)(nil)).Elem()
}

func (o LookupManagementGroupResultOutput) ToLookupManagementGroupResultOutput() LookupManagementGroupResultOutput {
	return o
}

func (o LookupManagementGroupResultOutput) ToLookupManagementGroupResultOutputWithContext(ctx context.Context) LookupManagementGroupResultOutput {
	return o
}

func (o LookupManagementGroupResultOutput) Children() ManagementGroupChildInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) []ManagementGroupChildInfoResponse { return v.Children }).(ManagementGroupChildInfoResponseArrayOutput)
}

func (o LookupManagementGroupResultOutput) Details() ManagementGroupDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) *ManagementGroupDetailsResponse { return v.Details }).(ManagementGroupDetailsResponsePtrOutput)
}

func (o LookupManagementGroupResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagementGroupResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o LookupManagementGroupResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementGroupResultOutput{})
}
