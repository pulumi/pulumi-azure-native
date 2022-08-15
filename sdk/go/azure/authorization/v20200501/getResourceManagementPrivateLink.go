


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupResourceManagementPrivateLink(ctx *pulumi.Context, args *LookupResourceManagementPrivateLinkArgs, opts ...pulumi.InvokeOption) (*LookupResourceManagementPrivateLinkResult, error) {
	var rv LookupResourceManagementPrivateLinkResult
	err := ctx.Invoke("azure-native:authorization/v20200501:getResourceManagementPrivateLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceManagementPrivateLinkArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RmplName          string `pulumi:"rmplName"`
}

type LookupResourceManagementPrivateLinkResult struct {
	Id         string                                                   `pulumi:"id"`
	Location   *string                                                  `pulumi:"location"`
	Name       string                                                   `pulumi:"name"`
	Properties ResourceManagementPrivateLinkEndpointConnectionsResponse `pulumi:"properties"`
	Type       string                                                   `pulumi:"type"`
}

func LookupResourceManagementPrivateLinkOutput(ctx *pulumi.Context, args LookupResourceManagementPrivateLinkOutputArgs, opts ...pulumi.InvokeOption) LookupResourceManagementPrivateLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceManagementPrivateLinkResult, error) {
			args := v.(LookupResourceManagementPrivateLinkArgs)
			r, err := LookupResourceManagementPrivateLink(ctx, &args, opts...)
			var s LookupResourceManagementPrivateLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceManagementPrivateLinkResultOutput)
}

type LookupResourceManagementPrivateLinkOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RmplName          pulumi.StringInput `pulumi:"rmplName"`
}

func (LookupResourceManagementPrivateLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceManagementPrivateLinkArgs)(nil)).Elem()
}

type LookupResourceManagementPrivateLinkResultOutput struct{ *pulumi.OutputState }

func (LookupResourceManagementPrivateLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceManagementPrivateLinkResult)(nil)).Elem()
}

func (o LookupResourceManagementPrivateLinkResultOutput) ToLookupResourceManagementPrivateLinkResultOutput() LookupResourceManagementPrivateLinkResultOutput {
	return o
}

func (o LookupResourceManagementPrivateLinkResultOutput) ToLookupResourceManagementPrivateLinkResultOutputWithContext(ctx context.Context) LookupResourceManagementPrivateLinkResultOutput {
	return o
}

func (o LookupResourceManagementPrivateLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceManagementPrivateLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResourceManagementPrivateLinkResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceManagementPrivateLinkResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupResourceManagementPrivateLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceManagementPrivateLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupResourceManagementPrivateLinkResultOutput) Properties() ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return o.ApplyT(func(v LookupResourceManagementPrivateLinkResult) ResourceManagementPrivateLinkEndpointConnectionsResponse {
		return v.Properties
	}).(ResourceManagementPrivateLinkEndpointConnectionsResponseOutput)
}

func (o LookupResourceManagementPrivateLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceManagementPrivateLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceManagementPrivateLinkResultOutput{})
}
