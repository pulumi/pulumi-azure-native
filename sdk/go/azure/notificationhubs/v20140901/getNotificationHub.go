


package v20140901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupNotificationHub(ctx *pulumi.Context, args *LookupNotificationHubArgs, opts ...pulumi.InvokeOption) (*LookupNotificationHubResult, error) {
	var rv LookupNotificationHubResult
	err := ctx.Invoke("azure-native:notificationhubs/v20140901:getNotificationHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationHubArgs struct {
	NamespaceName       string `pulumi:"namespaceName"`
	NotificationHubName string `pulumi:"notificationHubName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupNotificationHubResult struct {
	Id         *string                           `pulumi:"id"`
	Location   *string                           `pulumi:"location"`
	Name       *string                           `pulumi:"name"`
	Properties NotificationHubPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       *string                           `pulumi:"type"`
}

func LookupNotificationHubOutput(ctx *pulumi.Context, args LookupNotificationHubOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotificationHubResult, error) {
			args := v.(LookupNotificationHubArgs)
			r, err := LookupNotificationHub(ctx, &args, opts...)
			var s LookupNotificationHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotificationHubResultOutput)
}

type LookupNotificationHubOutputArgs struct {
	NamespaceName       pulumi.StringInput `pulumi:"namespaceName"`
	NotificationHubName pulumi.StringInput `pulumi:"notificationHubName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNotificationHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubArgs)(nil)).Elem()
}


type LookupNotificationHubResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubResult)(nil)).Elem()
}

func (o LookupNotificationHubResultOutput) ToLookupNotificationHubResultOutput() LookupNotificationHubResultOutput {
	return o
}

func (o LookupNotificationHubResultOutput) ToLookupNotificationHubResultOutputWithContext(ctx context.Context) LookupNotificationHubResultOutput {
	return o
}

func (o LookupNotificationHubResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationHubResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationHubResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationHubResultOutput) Properties() NotificationHubPropertiesResponseOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) NotificationHubPropertiesResponse { return v.Properties }).(NotificationHubPropertiesResponseOutput)
}

func (o LookupNotificationHubResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNotificationHubResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationHubResultOutput{})
}
