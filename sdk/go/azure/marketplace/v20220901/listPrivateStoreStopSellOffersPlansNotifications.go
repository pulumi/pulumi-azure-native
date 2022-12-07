


package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListPrivateStoreStopSellOffersPlansNotifications(ctx *pulumi.Context, args *ListPrivateStoreStopSellOffersPlansNotificationsArgs, opts ...pulumi.InvokeOption) (*ListPrivateStoreStopSellOffersPlansNotificationsResult, error) {
	var rv ListPrivateStoreStopSellOffersPlansNotificationsResult
	err := ctx.Invoke("azure-native:marketplace/v20220901:listPrivateStoreStopSellOffersPlansNotifications", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPrivateStoreStopSellOffersPlansNotificationsArgs struct {
	PrivateStoreId string   `pulumi:"privateStoreId"`
	Subscriptions  []string `pulumi:"subscriptions"`
}


type ListPrivateStoreStopSellOffersPlansNotificationsResult struct {
	StopSellNotifications []StopSellOffersPlansNotificationsListPropertiesResponse `pulumi:"stopSellNotifications"`
}

func ListPrivateStoreStopSellOffersPlansNotificationsOutput(ctx *pulumi.Context, args ListPrivateStoreStopSellOffersPlansNotificationsOutputArgs, opts ...pulumi.InvokeOption) ListPrivateStoreStopSellOffersPlansNotificationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPrivateStoreStopSellOffersPlansNotificationsResult, error) {
			args := v.(ListPrivateStoreStopSellOffersPlansNotificationsArgs)
			r, err := ListPrivateStoreStopSellOffersPlansNotifications(ctx, &args, opts...)
			var s ListPrivateStoreStopSellOffersPlansNotificationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPrivateStoreStopSellOffersPlansNotificationsResultOutput)
}

type ListPrivateStoreStopSellOffersPlansNotificationsOutputArgs struct {
	PrivateStoreId pulumi.StringInput      `pulumi:"privateStoreId"`
	Subscriptions  pulumi.StringArrayInput `pulumi:"subscriptions"`
}

func (ListPrivateStoreStopSellOffersPlansNotificationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPrivateStoreStopSellOffersPlansNotificationsArgs)(nil)).Elem()
}


type ListPrivateStoreStopSellOffersPlansNotificationsResultOutput struct{ *pulumi.OutputState }

func (ListPrivateStoreStopSellOffersPlansNotificationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPrivateStoreStopSellOffersPlansNotificationsResult)(nil)).Elem()
}

func (o ListPrivateStoreStopSellOffersPlansNotificationsResultOutput) ToListPrivateStoreStopSellOffersPlansNotificationsResultOutput() ListPrivateStoreStopSellOffersPlansNotificationsResultOutput {
	return o
}

func (o ListPrivateStoreStopSellOffersPlansNotificationsResultOutput) ToListPrivateStoreStopSellOffersPlansNotificationsResultOutputWithContext(ctx context.Context) ListPrivateStoreStopSellOffersPlansNotificationsResultOutput {
	return o
}

func (o ListPrivateStoreStopSellOffersPlansNotificationsResultOutput) StopSellNotifications() StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput {
	return o.ApplyT(func(v ListPrivateStoreStopSellOffersPlansNotificationsResult) []StopSellOffersPlansNotificationsListPropertiesResponse {
		return v.StopSellNotifications
	}).(StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPrivateStoreStopSellOffersPlansNotificationsResultOutput{})
}
