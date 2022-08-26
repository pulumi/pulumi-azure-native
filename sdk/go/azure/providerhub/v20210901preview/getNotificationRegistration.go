


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNotificationRegistration(ctx *pulumi.Context, args *LookupNotificationRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupNotificationRegistrationResult, error) {
	var rv LookupNotificationRegistrationResult
	err := ctx.Invoke("azure-native:providerhub/v20210901preview:getNotificationRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationRegistrationArgs struct {
	NotificationRegistrationName string `pulumi:"notificationRegistrationName"`
	ProviderNamespace            string `pulumi:"providerNamespace"`
}


type LookupNotificationRegistrationResult struct {
	Id         string                                     `pulumi:"id"`
	Name       string                                     `pulumi:"name"`
	Properties NotificationRegistrationResponseProperties `pulumi:"properties"`
	SystemData SystemDataResponse                         `pulumi:"systemData"`
	Type       string                                     `pulumi:"type"`
}

func LookupNotificationRegistrationOutput(ctx *pulumi.Context, args LookupNotificationRegistrationOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationRegistrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotificationRegistrationResult, error) {
			args := v.(LookupNotificationRegistrationArgs)
			r, err := LookupNotificationRegistration(ctx, &args, opts...)
			var s LookupNotificationRegistrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotificationRegistrationResultOutput)
}

type LookupNotificationRegistrationOutputArgs struct {
	NotificationRegistrationName pulumi.StringInput `pulumi:"notificationRegistrationName"`
	ProviderNamespace            pulumi.StringInput `pulumi:"providerNamespace"`
}

func (LookupNotificationRegistrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationRegistrationArgs)(nil)).Elem()
}


type LookupNotificationRegistrationResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationRegistrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationRegistrationResult)(nil)).Elem()
}

func (o LookupNotificationRegistrationResultOutput) ToLookupNotificationRegistrationResultOutput() LookupNotificationRegistrationResultOutput {
	return o
}

func (o LookupNotificationRegistrationResultOutput) ToLookupNotificationRegistrationResultOutputWithContext(ctx context.Context) LookupNotificationRegistrationResultOutput {
	return o
}

func (o LookupNotificationRegistrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationRegistrationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNotificationRegistrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationRegistrationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNotificationRegistrationResultOutput) Properties() NotificationRegistrationResponsePropertiesOutput {
	return o.ApplyT(func(v LookupNotificationRegistrationResult) NotificationRegistrationResponseProperties {
		return v.Properties
	}).(NotificationRegistrationResponsePropertiesOutput)
}

func (o LookupNotificationRegistrationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNotificationRegistrationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupNotificationRegistrationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationRegistrationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationRegistrationResultOutput{})
}
