


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNspLink(ctx *pulumi.Context, args *LookupNspLinkArgs, opts ...pulumi.InvokeOption) (*LookupNspLinkResult, error) {
	var rv LookupNspLinkResult
	err := ctx.Invoke("azure-native:network/v20210201preview:getNspLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNspLinkArgs struct {
	LinkName                     string `pulumi:"linkName"`
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type LookupNspLinkResult struct {
	AutoApprovedRemotePerimeterResourceId *string  `pulumi:"autoApprovedRemotePerimeterResourceId"`
	Description                           *string  `pulumi:"description"`
	Etag                                  string   `pulumi:"etag"`
	Id                                    string   `pulumi:"id"`
	LocalInboundProfiles                  []string `pulumi:"localInboundProfiles"`
	LocalOutboundProfiles                 []string `pulumi:"localOutboundProfiles"`
	Name                                  string   `pulumi:"name"`
	ProvisioningState                     string   `pulumi:"provisioningState"`
	RemoteInboundProfiles                 []string `pulumi:"remoteInboundProfiles"`
	RemoteOutboundProfiles                []string `pulumi:"remoteOutboundProfiles"`
	RemotePerimeterGuid                   string   `pulumi:"remotePerimeterGuid"`
	Status                                string   `pulumi:"status"`
	Type                                  string   `pulumi:"type"`
}

func LookupNspLinkOutput(ctx *pulumi.Context, args LookupNspLinkOutputArgs, opts ...pulumi.InvokeOption) LookupNspLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNspLinkResult, error) {
			args := v.(LookupNspLinkArgs)
			r, err := LookupNspLink(ctx, &args, opts...)
			var s LookupNspLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNspLinkResultOutput)
}

type LookupNspLinkOutputArgs struct {
	LinkName                     pulumi.StringInput `pulumi:"linkName"`
	NetworkSecurityPerimeterName pulumi.StringInput `pulumi:"networkSecurityPerimeterName"`
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNspLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspLinkArgs)(nil)).Elem()
}


type LookupNspLinkResultOutput struct{ *pulumi.OutputState }

func (LookupNspLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspLinkResult)(nil)).Elem()
}

func (o LookupNspLinkResultOutput) ToLookupNspLinkResultOutput() LookupNspLinkResultOutput {
	return o
}

func (o LookupNspLinkResultOutput) ToLookupNspLinkResultOutputWithContext(ctx context.Context) LookupNspLinkResultOutput {
	return o
}

func (o LookupNspLinkResultOutput) AutoApprovedRemotePerimeterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspLinkResult) *string { return v.AutoApprovedRemotePerimeterResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupNspLinkResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspLinkResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupNspLinkResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspLinkResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupNspLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNspLinkResultOutput) LocalInboundProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNspLinkResult) []string { return v.LocalInboundProfiles }).(pulumi.StringArrayOutput)
}

func (o LookupNspLinkResultOutput) LocalOutboundProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNspLinkResult) []string { return v.LocalOutboundProfiles }).(pulumi.StringArrayOutput)
}

func (o LookupNspLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNspLinkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspLinkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNspLinkResultOutput) RemoteInboundProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNspLinkResult) []string { return v.RemoteInboundProfiles }).(pulumi.StringArrayOutput)
}

func (o LookupNspLinkResultOutput) RemoteOutboundProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNspLinkResult) []string { return v.RemoteOutboundProfiles }).(pulumi.StringArrayOutput)
}

func (o LookupNspLinkResultOutput) RemotePerimeterGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspLinkResult) string { return v.RemotePerimeterGuid }).(pulumi.StringOutput)
}

func (o LookupNspLinkResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspLinkResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupNspLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNspLinkResultOutput{})
}
