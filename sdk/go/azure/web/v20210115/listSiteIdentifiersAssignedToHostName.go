


package v20210115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteIdentifiersAssignedToHostName(ctx *pulumi.Context, args *ListSiteIdentifiersAssignedToHostNameArgs, opts ...pulumi.InvokeOption) (*ListSiteIdentifiersAssignedToHostNameResult, error) {
	var rv ListSiteIdentifiersAssignedToHostNameResult
	err := ctx.Invoke("azure-native:web/v20210115:listSiteIdentifiersAssignedToHostName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteIdentifiersAssignedToHostNameArgs struct {
	Name *string `pulumi:"name"`
}


type ListSiteIdentifiersAssignedToHostNameResult struct {
	NextLink string               `pulumi:"nextLink"`
	Value    []IdentifierResponse `pulumi:"value"`
}

func ListSiteIdentifiersAssignedToHostNameOutput(ctx *pulumi.Context, args ListSiteIdentifiersAssignedToHostNameOutputArgs, opts ...pulumi.InvokeOption) ListSiteIdentifiersAssignedToHostNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteIdentifiersAssignedToHostNameResult, error) {
			args := v.(ListSiteIdentifiersAssignedToHostNameArgs)
			r, err := ListSiteIdentifiersAssignedToHostName(ctx, &args, opts...)
			var s ListSiteIdentifiersAssignedToHostNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteIdentifiersAssignedToHostNameResultOutput)
}

type ListSiteIdentifiersAssignedToHostNameOutputArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ListSiteIdentifiersAssignedToHostNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteIdentifiersAssignedToHostNameArgs)(nil)).Elem()
}


type ListSiteIdentifiersAssignedToHostNameResultOutput struct{ *pulumi.OutputState }

func (ListSiteIdentifiersAssignedToHostNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteIdentifiersAssignedToHostNameResult)(nil)).Elem()
}

func (o ListSiteIdentifiersAssignedToHostNameResultOutput) ToListSiteIdentifiersAssignedToHostNameResultOutput() ListSiteIdentifiersAssignedToHostNameResultOutput {
	return o
}

func (o ListSiteIdentifiersAssignedToHostNameResultOutput) ToListSiteIdentifiersAssignedToHostNameResultOutputWithContext(ctx context.Context) ListSiteIdentifiersAssignedToHostNameResultOutput {
	return o
}

func (o ListSiteIdentifiersAssignedToHostNameResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteIdentifiersAssignedToHostNameResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListSiteIdentifiersAssignedToHostNameResultOutput) Value() IdentifierResponseArrayOutput {
	return o.ApplyT(func(v ListSiteIdentifiersAssignedToHostNameResult) []IdentifierResponse { return v.Value }).(IdentifierResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteIdentifiersAssignedToHostNameResultOutput{})
}
