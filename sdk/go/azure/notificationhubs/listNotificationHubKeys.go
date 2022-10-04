


package notificationhubs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNotificationHubKeys(ctx *pulumi.Context, args *ListNotificationHubKeysArgs, opts ...pulumi.InvokeOption) (*ListNotificationHubKeysResult, error) {
	var rv ListNotificationHubKeysResult
	err := ctx.Invoke("azure-native:notificationhubs:listNotificationHubKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNotificationHubKeysArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	NotificationHubName   string `pulumi:"notificationHubName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type ListNotificationHubKeysResult struct {
	KeyName                   *string `pulumi:"keyName"`
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	PrimaryKey                *string `pulumi:"primaryKey"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	SecondaryKey              *string `pulumi:"secondaryKey"`
}

func ListNotificationHubKeysOutput(ctx *pulumi.Context, args ListNotificationHubKeysOutputArgs, opts ...pulumi.InvokeOption) ListNotificationHubKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNotificationHubKeysResult, error) {
			args := v.(ListNotificationHubKeysArgs)
			r, err := ListNotificationHubKeys(ctx, &args, opts...)
			var s ListNotificationHubKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNotificationHubKeysResultOutput)
}

type ListNotificationHubKeysOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	NotificationHubName   pulumi.StringInput `pulumi:"notificationHubName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListNotificationHubKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNotificationHubKeysArgs)(nil)).Elem()
}


type ListNotificationHubKeysResultOutput struct{ *pulumi.OutputState }

func (ListNotificationHubKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNotificationHubKeysResult)(nil)).Elem()
}

func (o ListNotificationHubKeysResultOutput) ToListNotificationHubKeysResultOutput() ListNotificationHubKeysResultOutput {
	return o
}

func (o ListNotificationHubKeysResultOutput) ToListNotificationHubKeysResultOutputWithContext(ctx context.Context) ListNotificationHubKeysResultOutput {
	return o
}

func (o ListNotificationHubKeysResultOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNotificationHubKeysResult) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o ListNotificationHubKeysResultOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNotificationHubKeysResult) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListNotificationHubKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNotificationHubKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListNotificationHubKeysResultOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNotificationHubKeysResult) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListNotificationHubKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNotificationHubKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNotificationHubKeysResultOutput{})
}
