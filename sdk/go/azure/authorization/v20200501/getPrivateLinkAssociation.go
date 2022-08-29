


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateLinkAssociation(ctx *pulumi.Context, args *LookupPrivateLinkAssociationArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkAssociationResult, error) {
	var rv LookupPrivateLinkAssociationResult
	err := ctx.Invoke("azure-native:authorization/v20200501:getPrivateLinkAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkAssociationArgs struct {
	GroupId string `pulumi:"groupId"`
	PlaId   string `pulumi:"plaId"`
}

type LookupPrivateLinkAssociationResult struct {
	Id         string                                           `pulumi:"id"`
	Name       string                                           `pulumi:"name"`
	Properties PrivateLinkAssociationPropertiesExpandedResponse `pulumi:"properties"`
	Type       string                                           `pulumi:"type"`
}

func LookupPrivateLinkAssociationOutput(ctx *pulumi.Context, args LookupPrivateLinkAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkAssociationResult, error) {
			args := v.(LookupPrivateLinkAssociationArgs)
			r, err := LookupPrivateLinkAssociation(ctx, &args, opts...)
			var s LookupPrivateLinkAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateLinkAssociationResultOutput)
}

type LookupPrivateLinkAssociationOutputArgs struct {
	GroupId pulumi.StringInput `pulumi:"groupId"`
	PlaId   pulumi.StringInput `pulumi:"plaId"`
}

func (LookupPrivateLinkAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkAssociationArgs)(nil)).Elem()
}

type LookupPrivateLinkAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkAssociationResult)(nil)).Elem()
}

func (o LookupPrivateLinkAssociationResultOutput) ToLookupPrivateLinkAssociationResultOutput() LookupPrivateLinkAssociationResultOutput {
	return o
}

func (o LookupPrivateLinkAssociationResultOutput) ToLookupPrivateLinkAssociationResultOutputWithContext(ctx context.Context) LookupPrivateLinkAssociationResultOutput {
	return o
}

func (o LookupPrivateLinkAssociationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkAssociationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkAssociationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkAssociationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkAssociationResultOutput) Properties() PrivateLinkAssociationPropertiesExpandedResponseOutput {
	return o.ApplyT(func(v LookupPrivateLinkAssociationResult) PrivateLinkAssociationPropertiesExpandedResponse {
		return v.Properties
	}).(PrivateLinkAssociationPropertiesExpandedResponseOutput)
}

func (o LookupPrivateLinkAssociationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkAssociationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkAssociationResultOutput{})
}
