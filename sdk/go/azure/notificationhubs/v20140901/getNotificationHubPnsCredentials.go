


package v20140901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNotificationHubPnsCredentials(ctx *pulumi.Context, args *GetNotificationHubPnsCredentialsArgs, opts ...pulumi.InvokeOption) (*GetNotificationHubPnsCredentialsResult, error) {
	var rv GetNotificationHubPnsCredentialsResult
	err := ctx.Invoke("azure-native:notificationhubs/v20140901:getNotificationHubPnsCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetNotificationHubPnsCredentialsArgs struct {
	NamespaceName       string `pulumi:"namespaceName"`
	NotificationHubName string `pulumi:"notificationHubName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type GetNotificationHubPnsCredentialsResult struct {
	Id         *string                           `pulumi:"id"`
	Location   *string                           `pulumi:"location"`
	Name       *string                           `pulumi:"name"`
	Properties NotificationHubPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       *string                           `pulumi:"type"`
}

func GetNotificationHubPnsCredentialsOutput(ctx *pulumi.Context, args GetNotificationHubPnsCredentialsOutputArgs, opts ...pulumi.InvokeOption) GetNotificationHubPnsCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNotificationHubPnsCredentialsResult, error) {
			args := v.(GetNotificationHubPnsCredentialsArgs)
			r, err := GetNotificationHubPnsCredentials(ctx, &args, opts...)
			var s GetNotificationHubPnsCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNotificationHubPnsCredentialsResultOutput)
}

type GetNotificationHubPnsCredentialsOutputArgs struct {
	NamespaceName       pulumi.StringInput `pulumi:"namespaceName"`
	NotificationHubName pulumi.StringInput `pulumi:"notificationHubName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetNotificationHubPnsCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNotificationHubPnsCredentialsArgs)(nil)).Elem()
}


type GetNotificationHubPnsCredentialsResultOutput struct{ *pulumi.OutputState }

func (GetNotificationHubPnsCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNotificationHubPnsCredentialsResult)(nil)).Elem()
}

func (o GetNotificationHubPnsCredentialsResultOutput) ToGetNotificationHubPnsCredentialsResultOutput() GetNotificationHubPnsCredentialsResultOutput {
	return o
}

func (o GetNotificationHubPnsCredentialsResultOutput) ToGetNotificationHubPnsCredentialsResultOutputWithContext(ctx context.Context) GetNotificationHubPnsCredentialsResultOutput {
	return o
}

func (o GetNotificationHubPnsCredentialsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) Properties() NotificationHubPropertiesResponseOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) NotificationHubPropertiesResponse { return v.Properties }).(NotificationHubPropertiesResponseOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNotificationHubPnsCredentialsResultOutput{})
}
