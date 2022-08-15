


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebPubSubSharedPrivateLinkResource(ctx *pulumi.Context, args *LookupWebPubSubSharedPrivateLinkResourceArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubSharedPrivateLinkResourceResult, error) {
	var rv LookupWebPubSubSharedPrivateLinkResourceResult
	err := ctx.Invoke("azure-native:webpubsub/v20210401preview:getWebPubSubSharedPrivateLinkResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebPubSubSharedPrivateLinkResourceArgs struct {
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
	SharedPrivateLinkResourceName string `pulumi:"sharedPrivateLinkResourceName"`
}


type LookupWebPubSubSharedPrivateLinkResourceResult struct {
	GroupId               string             `pulumi:"groupId"`
	Id                    string             `pulumi:"id"`
	Name                  string             `pulumi:"name"`
	PrivateLinkResourceId string             `pulumi:"privateLinkResourceId"`
	ProvisioningState     string             `pulumi:"provisioningState"`
	RequestMessage        *string            `pulumi:"requestMessage"`
	Status                string             `pulumi:"status"`
	SystemData            SystemDataResponse `pulumi:"systemData"`
	Type                  string             `pulumi:"type"`
}

func LookupWebPubSubSharedPrivateLinkResourceOutput(ctx *pulumi.Context, args LookupWebPubSubSharedPrivateLinkResourceOutputArgs, opts ...pulumi.InvokeOption) LookupWebPubSubSharedPrivateLinkResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebPubSubSharedPrivateLinkResourceResult, error) {
			args := v.(LookupWebPubSubSharedPrivateLinkResourceArgs)
			r, err := LookupWebPubSubSharedPrivateLinkResource(ctx, &args, opts...)
			var s LookupWebPubSubSharedPrivateLinkResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebPubSubSharedPrivateLinkResourceResultOutput)
}

type LookupWebPubSubSharedPrivateLinkResourceOutputArgs struct {
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                  pulumi.StringInput `pulumi:"resourceName"`
	SharedPrivateLinkResourceName pulumi.StringInput `pulumi:"sharedPrivateLinkResourceName"`
}

func (LookupWebPubSubSharedPrivateLinkResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubSharedPrivateLinkResourceArgs)(nil)).Elem()
}


type LookupWebPubSubSharedPrivateLinkResourceResultOutput struct{ *pulumi.OutputState }

func (LookupWebPubSubSharedPrivateLinkResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubSharedPrivateLinkResourceResult)(nil)).Elem()
}

func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) ToLookupWebPubSubSharedPrivateLinkResourceResultOutput() LookupWebPubSubSharedPrivateLinkResourceResultOutput {
	return o
}

func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) ToLookupWebPubSubSharedPrivateLinkResourceResultOutputWithContext(ctx context.Context) LookupWebPubSubSharedPrivateLinkResourceResultOutput {
	return o
}

func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.GroupId }).(pulumi.StringOutput)
}

func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) PrivateLinkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.PrivateLinkResourceId }).(pulumi.StringOutput)
}

func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebPubSubSharedPrivateLinkResourceResultOutput{})
}
