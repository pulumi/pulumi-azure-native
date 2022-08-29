


package v20220201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSignalRSharedPrivateLinkResource(ctx *pulumi.Context, args *LookupSignalRSharedPrivateLinkResourceArgs, opts ...pulumi.InvokeOption) (*LookupSignalRSharedPrivateLinkResourceResult, error) {
	var rv LookupSignalRSharedPrivateLinkResourceResult
	err := ctx.Invoke("azure-native:signalrservice/v20220201:getSignalRSharedPrivateLinkResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalRSharedPrivateLinkResourceArgs struct {
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
	SharedPrivateLinkResourceName string `pulumi:"sharedPrivateLinkResourceName"`
}


type LookupSignalRSharedPrivateLinkResourceResult struct {
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

func LookupSignalRSharedPrivateLinkResourceOutput(ctx *pulumi.Context, args LookupSignalRSharedPrivateLinkResourceOutputArgs, opts ...pulumi.InvokeOption) LookupSignalRSharedPrivateLinkResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSignalRSharedPrivateLinkResourceResult, error) {
			args := v.(LookupSignalRSharedPrivateLinkResourceArgs)
			r, err := LookupSignalRSharedPrivateLinkResource(ctx, &args, opts...)
			var s LookupSignalRSharedPrivateLinkResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSignalRSharedPrivateLinkResourceResultOutput)
}

type LookupSignalRSharedPrivateLinkResourceOutputArgs struct {
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                  pulumi.StringInput `pulumi:"resourceName"`
	SharedPrivateLinkResourceName pulumi.StringInput `pulumi:"sharedPrivateLinkResourceName"`
}

func (LookupSignalRSharedPrivateLinkResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRSharedPrivateLinkResourceArgs)(nil)).Elem()
}


type LookupSignalRSharedPrivateLinkResourceResultOutput struct{ *pulumi.OutputState }

func (LookupSignalRSharedPrivateLinkResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRSharedPrivateLinkResourceResult)(nil)).Elem()
}

func (o LookupSignalRSharedPrivateLinkResourceResultOutput) ToLookupSignalRSharedPrivateLinkResourceResultOutput() LookupSignalRSharedPrivateLinkResourceResultOutput {
	return o
}

func (o LookupSignalRSharedPrivateLinkResourceResultOutput) ToLookupSignalRSharedPrivateLinkResourceResultOutputWithContext(ctx context.Context) LookupSignalRSharedPrivateLinkResourceResultOutput {
	return o
}

func (o LookupSignalRSharedPrivateLinkResourceResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.GroupId }).(pulumi.StringOutput)
}

func (o LookupSignalRSharedPrivateLinkResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSignalRSharedPrivateLinkResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSignalRSharedPrivateLinkResourceResultOutput) PrivateLinkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.PrivateLinkResourceId }).(pulumi.StringOutput)
}

func (o LookupSignalRSharedPrivateLinkResourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSignalRSharedPrivateLinkResourceResultOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

func (o LookupSignalRSharedPrivateLinkResourceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSignalRSharedPrivateLinkResourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSignalRSharedPrivateLinkResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalRSharedPrivateLinkResourceResultOutput{})
}
