


package v20170418

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CognitiveServicesAccountApiProperties struct {
	AadClientId                    *string `pulumi:"aadClientId"`
	AadTenantId                    *string `pulumi:"aadTenantId"`
	EventHubConnectionString       *string `pulumi:"eventHubConnectionString"`
	QnaAzureSearchEndpointId       *string `pulumi:"qnaAzureSearchEndpointId"`
	QnaAzureSearchEndpointKey      *string `pulumi:"qnaAzureSearchEndpointKey"`
	QnaRuntimeEndpoint             *string `pulumi:"qnaRuntimeEndpoint"`
	StatisticsEnabled              *bool   `pulumi:"statisticsEnabled"`
	StorageAccountConnectionString *string `pulumi:"storageAccountConnectionString"`
	SuperUser                      *string `pulumi:"superUser"`
	WebsiteName                    *string `pulumi:"websiteName"`
}





type CognitiveServicesAccountApiPropertiesInput interface {
	pulumi.Input

	ToCognitiveServicesAccountApiPropertiesOutput() CognitiveServicesAccountApiPropertiesOutput
	ToCognitiveServicesAccountApiPropertiesOutputWithContext(context.Context) CognitiveServicesAccountApiPropertiesOutput
}

type CognitiveServicesAccountApiPropertiesArgs struct {
	AadClientId                    pulumi.StringPtrInput `pulumi:"aadClientId"`
	AadTenantId                    pulumi.StringPtrInput `pulumi:"aadTenantId"`
	EventHubConnectionString       pulumi.StringPtrInput `pulumi:"eventHubConnectionString"`
	QnaAzureSearchEndpointId       pulumi.StringPtrInput `pulumi:"qnaAzureSearchEndpointId"`
	QnaAzureSearchEndpointKey      pulumi.StringPtrInput `pulumi:"qnaAzureSearchEndpointKey"`
	QnaRuntimeEndpoint             pulumi.StringPtrInput `pulumi:"qnaRuntimeEndpoint"`
	StatisticsEnabled              pulumi.BoolPtrInput   `pulumi:"statisticsEnabled"`
	StorageAccountConnectionString pulumi.StringPtrInput `pulumi:"storageAccountConnectionString"`
	SuperUser                      pulumi.StringPtrInput `pulumi:"superUser"`
	WebsiteName                    pulumi.StringPtrInput `pulumi:"websiteName"`
}

func (CognitiveServicesAccountApiPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CognitiveServicesAccountApiProperties)(nil)).Elem()
}

func (i CognitiveServicesAccountApiPropertiesArgs) ToCognitiveServicesAccountApiPropertiesOutput() CognitiveServicesAccountApiPropertiesOutput {
	return i.ToCognitiveServicesAccountApiPropertiesOutputWithContext(context.Background())
}

func (i CognitiveServicesAccountApiPropertiesArgs) ToCognitiveServicesAccountApiPropertiesOutputWithContext(ctx context.Context) CognitiveServicesAccountApiPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountApiPropertiesOutput)
}

func (i CognitiveServicesAccountApiPropertiesArgs) ToCognitiveServicesAccountApiPropertiesPtrOutput() CognitiveServicesAccountApiPropertiesPtrOutput {
	return i.ToCognitiveServicesAccountApiPropertiesPtrOutputWithContext(context.Background())
}

func (i CognitiveServicesAccountApiPropertiesArgs) ToCognitiveServicesAccountApiPropertiesPtrOutputWithContext(ctx context.Context) CognitiveServicesAccountApiPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountApiPropertiesOutput).ToCognitiveServicesAccountApiPropertiesPtrOutputWithContext(ctx)
}









type CognitiveServicesAccountApiPropertiesPtrInput interface {
	pulumi.Input

	ToCognitiveServicesAccountApiPropertiesPtrOutput() CognitiveServicesAccountApiPropertiesPtrOutput
	ToCognitiveServicesAccountApiPropertiesPtrOutputWithContext(context.Context) CognitiveServicesAccountApiPropertiesPtrOutput
}

type cognitiveServicesAccountApiPropertiesPtrType CognitiveServicesAccountApiPropertiesArgs

func CognitiveServicesAccountApiPropertiesPtr(v *CognitiveServicesAccountApiPropertiesArgs) CognitiveServicesAccountApiPropertiesPtrInput {
	return (*cognitiveServicesAccountApiPropertiesPtrType)(v)
}

func (*cognitiveServicesAccountApiPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccountApiProperties)(nil)).Elem()
}

func (i *cognitiveServicesAccountApiPropertiesPtrType) ToCognitiveServicesAccountApiPropertiesPtrOutput() CognitiveServicesAccountApiPropertiesPtrOutput {
	return i.ToCognitiveServicesAccountApiPropertiesPtrOutputWithContext(context.Background())
}

func (i *cognitiveServicesAccountApiPropertiesPtrType) ToCognitiveServicesAccountApiPropertiesPtrOutputWithContext(ctx context.Context) CognitiveServicesAccountApiPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountApiPropertiesPtrOutput)
}

type CognitiveServicesAccountApiPropertiesOutput struct{ *pulumi.OutputState }

func (CognitiveServicesAccountApiPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CognitiveServicesAccountApiProperties)(nil)).Elem()
}

func (o CognitiveServicesAccountApiPropertiesOutput) ToCognitiveServicesAccountApiPropertiesOutput() CognitiveServicesAccountApiPropertiesOutput {
	return o
}

func (o CognitiveServicesAccountApiPropertiesOutput) ToCognitiveServicesAccountApiPropertiesOutputWithContext(ctx context.Context) CognitiveServicesAccountApiPropertiesOutput {
	return o
}

func (o CognitiveServicesAccountApiPropertiesOutput) ToCognitiveServicesAccountApiPropertiesPtrOutput() CognitiveServicesAccountApiPropertiesPtrOutput {
	return o.ToCognitiveServicesAccountApiPropertiesPtrOutputWithContext(context.Background())
}

func (o CognitiveServicesAccountApiPropertiesOutput) ToCognitiveServicesAccountApiPropertiesPtrOutputWithContext(ctx context.Context) CognitiveServicesAccountApiPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CognitiveServicesAccountApiProperties) *CognitiveServicesAccountApiProperties {
		return &v
	}).(CognitiveServicesAccountApiPropertiesPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesOutput) AadClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiProperties) *string { return v.AadClientId }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiProperties) *string { return v.AadTenantId }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesOutput) EventHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiProperties) *string { return v.EventHubConnectionString }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesOutput) QnaAzureSearchEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiProperties) *string { return v.QnaAzureSearchEndpointId }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesOutput) QnaAzureSearchEndpointKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiProperties) *string { return v.QnaAzureSearchEndpointKey }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesOutput) QnaRuntimeEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiProperties) *string { return v.QnaRuntimeEndpoint }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesOutput) StatisticsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiProperties) *bool { return v.StatisticsEnabled }).(pulumi.BoolPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesOutput) StorageAccountConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiProperties) *string { return v.StorageAccountConnectionString }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesOutput) SuperUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiProperties) *string { return v.SuperUser }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesOutput) WebsiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiProperties) *string { return v.WebsiteName }).(pulumi.StringPtrOutput)
}

type CognitiveServicesAccountApiPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CognitiveServicesAccountApiPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccountApiProperties)(nil)).Elem()
}

func (o CognitiveServicesAccountApiPropertiesPtrOutput) ToCognitiveServicesAccountApiPropertiesPtrOutput() CognitiveServicesAccountApiPropertiesPtrOutput {
	return o
}

func (o CognitiveServicesAccountApiPropertiesPtrOutput) ToCognitiveServicesAccountApiPropertiesPtrOutputWithContext(ctx context.Context) CognitiveServicesAccountApiPropertiesPtrOutput {
	return o
}

func (o CognitiveServicesAccountApiPropertiesPtrOutput) Elem() CognitiveServicesAccountApiPropertiesOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiProperties) CognitiveServicesAccountApiProperties {
		if v != nil {
			return *v
		}
		var ret CognitiveServicesAccountApiProperties
		return ret
	}).(CognitiveServicesAccountApiPropertiesOutput)
}

func (o CognitiveServicesAccountApiPropertiesPtrOutput) AadClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.AadClientId
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesPtrOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.AadTenantId
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesPtrOutput) EventHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.EventHubConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesPtrOutput) QnaAzureSearchEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.QnaAzureSearchEndpointId
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesPtrOutput) QnaAzureSearchEndpointKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.QnaAzureSearchEndpointKey
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesPtrOutput) QnaRuntimeEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.QnaRuntimeEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesPtrOutput) StatisticsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiProperties) *bool {
		if v == nil {
			return nil
		}
		return v.StatisticsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesPtrOutput) StorageAccountConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesPtrOutput) SuperUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.SuperUser
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesPtrOutput) WebsiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.WebsiteName
	}).(pulumi.StringPtrOutput)
}

type CognitiveServicesAccountApiPropertiesResponse struct {
	AadClientId                    *string `pulumi:"aadClientId"`
	AadTenantId                    *string `pulumi:"aadTenantId"`
	EventHubConnectionString       *string `pulumi:"eventHubConnectionString"`
	QnaAzureSearchEndpointId       *string `pulumi:"qnaAzureSearchEndpointId"`
	QnaAzureSearchEndpointKey      *string `pulumi:"qnaAzureSearchEndpointKey"`
	QnaRuntimeEndpoint             *string `pulumi:"qnaRuntimeEndpoint"`
	StatisticsEnabled              *bool   `pulumi:"statisticsEnabled"`
	StorageAccountConnectionString *string `pulumi:"storageAccountConnectionString"`
	SuperUser                      *string `pulumi:"superUser"`
	WebsiteName                    *string `pulumi:"websiteName"`
}





type CognitiveServicesAccountApiPropertiesResponseInput interface {
	pulumi.Input

	ToCognitiveServicesAccountApiPropertiesResponseOutput() CognitiveServicesAccountApiPropertiesResponseOutput
	ToCognitiveServicesAccountApiPropertiesResponseOutputWithContext(context.Context) CognitiveServicesAccountApiPropertiesResponseOutput
}

type CognitiveServicesAccountApiPropertiesResponseArgs struct {
	AadClientId                    pulumi.StringPtrInput `pulumi:"aadClientId"`
	AadTenantId                    pulumi.StringPtrInput `pulumi:"aadTenantId"`
	EventHubConnectionString       pulumi.StringPtrInput `pulumi:"eventHubConnectionString"`
	QnaAzureSearchEndpointId       pulumi.StringPtrInput `pulumi:"qnaAzureSearchEndpointId"`
	QnaAzureSearchEndpointKey      pulumi.StringPtrInput `pulumi:"qnaAzureSearchEndpointKey"`
	QnaRuntimeEndpoint             pulumi.StringPtrInput `pulumi:"qnaRuntimeEndpoint"`
	StatisticsEnabled              pulumi.BoolPtrInput   `pulumi:"statisticsEnabled"`
	StorageAccountConnectionString pulumi.StringPtrInput `pulumi:"storageAccountConnectionString"`
	SuperUser                      pulumi.StringPtrInput `pulumi:"superUser"`
	WebsiteName                    pulumi.StringPtrInput `pulumi:"websiteName"`
}

func (CognitiveServicesAccountApiPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CognitiveServicesAccountApiPropertiesResponse)(nil)).Elem()
}

func (i CognitiveServicesAccountApiPropertiesResponseArgs) ToCognitiveServicesAccountApiPropertiesResponseOutput() CognitiveServicesAccountApiPropertiesResponseOutput {
	return i.ToCognitiveServicesAccountApiPropertiesResponseOutputWithContext(context.Background())
}

func (i CognitiveServicesAccountApiPropertiesResponseArgs) ToCognitiveServicesAccountApiPropertiesResponseOutputWithContext(ctx context.Context) CognitiveServicesAccountApiPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountApiPropertiesResponseOutput)
}

func (i CognitiveServicesAccountApiPropertiesResponseArgs) ToCognitiveServicesAccountApiPropertiesResponsePtrOutput() CognitiveServicesAccountApiPropertiesResponsePtrOutput {
	return i.ToCognitiveServicesAccountApiPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CognitiveServicesAccountApiPropertiesResponseArgs) ToCognitiveServicesAccountApiPropertiesResponsePtrOutputWithContext(ctx context.Context) CognitiveServicesAccountApiPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountApiPropertiesResponseOutput).ToCognitiveServicesAccountApiPropertiesResponsePtrOutputWithContext(ctx)
}









type CognitiveServicesAccountApiPropertiesResponsePtrInput interface {
	pulumi.Input

	ToCognitiveServicesAccountApiPropertiesResponsePtrOutput() CognitiveServicesAccountApiPropertiesResponsePtrOutput
	ToCognitiveServicesAccountApiPropertiesResponsePtrOutputWithContext(context.Context) CognitiveServicesAccountApiPropertiesResponsePtrOutput
}

type cognitiveServicesAccountApiPropertiesResponsePtrType CognitiveServicesAccountApiPropertiesResponseArgs

func CognitiveServicesAccountApiPropertiesResponsePtr(v *CognitiveServicesAccountApiPropertiesResponseArgs) CognitiveServicesAccountApiPropertiesResponsePtrInput {
	return (*cognitiveServicesAccountApiPropertiesResponsePtrType)(v)
}

func (*cognitiveServicesAccountApiPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccountApiPropertiesResponse)(nil)).Elem()
}

func (i *cognitiveServicesAccountApiPropertiesResponsePtrType) ToCognitiveServicesAccountApiPropertiesResponsePtrOutput() CognitiveServicesAccountApiPropertiesResponsePtrOutput {
	return i.ToCognitiveServicesAccountApiPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *cognitiveServicesAccountApiPropertiesResponsePtrType) ToCognitiveServicesAccountApiPropertiesResponsePtrOutputWithContext(ctx context.Context) CognitiveServicesAccountApiPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountApiPropertiesResponsePtrOutput)
}

type CognitiveServicesAccountApiPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CognitiveServicesAccountApiPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CognitiveServicesAccountApiPropertiesResponse)(nil)).Elem()
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) ToCognitiveServicesAccountApiPropertiesResponseOutput() CognitiveServicesAccountApiPropertiesResponseOutput {
	return o
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) ToCognitiveServicesAccountApiPropertiesResponseOutputWithContext(ctx context.Context) CognitiveServicesAccountApiPropertiesResponseOutput {
	return o
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) ToCognitiveServicesAccountApiPropertiesResponsePtrOutput() CognitiveServicesAccountApiPropertiesResponsePtrOutput {
	return o.ToCognitiveServicesAccountApiPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) ToCognitiveServicesAccountApiPropertiesResponsePtrOutputWithContext(ctx context.Context) CognitiveServicesAccountApiPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CognitiveServicesAccountApiPropertiesResponse) *CognitiveServicesAccountApiPropertiesResponse {
		return &v
	}).(CognitiveServicesAccountApiPropertiesResponsePtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) AadClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiPropertiesResponse) *string { return v.AadClientId }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiPropertiesResponse) *string { return v.AadTenantId }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) EventHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiPropertiesResponse) *string { return v.EventHubConnectionString }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) QnaAzureSearchEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiPropertiesResponse) *string { return v.QnaAzureSearchEndpointId }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) QnaAzureSearchEndpointKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiPropertiesResponse) *string { return v.QnaAzureSearchEndpointKey }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) QnaRuntimeEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiPropertiesResponse) *string { return v.QnaRuntimeEndpoint }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) StatisticsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiPropertiesResponse) *bool { return v.StatisticsEnabled }).(pulumi.BoolPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) StorageAccountConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiPropertiesResponse) *string { return v.StorageAccountConnectionString }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) SuperUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiPropertiesResponse) *string { return v.SuperUser }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponseOutput) WebsiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountApiPropertiesResponse) *string { return v.WebsiteName }).(pulumi.StringPtrOutput)
}

type CognitiveServicesAccountApiPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CognitiveServicesAccountApiPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccountApiPropertiesResponse)(nil)).Elem()
}

func (o CognitiveServicesAccountApiPropertiesResponsePtrOutput) ToCognitiveServicesAccountApiPropertiesResponsePtrOutput() CognitiveServicesAccountApiPropertiesResponsePtrOutput {
	return o
}

func (o CognitiveServicesAccountApiPropertiesResponsePtrOutput) ToCognitiveServicesAccountApiPropertiesResponsePtrOutputWithContext(ctx context.Context) CognitiveServicesAccountApiPropertiesResponsePtrOutput {
	return o
}

func (o CognitiveServicesAccountApiPropertiesResponsePtrOutput) Elem() CognitiveServicesAccountApiPropertiesResponseOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiPropertiesResponse) CognitiveServicesAccountApiPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CognitiveServicesAccountApiPropertiesResponse
		return ret
	}).(CognitiveServicesAccountApiPropertiesResponseOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponsePtrOutput) AadClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadClientId
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponsePtrOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadTenantId
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponsePtrOutput) EventHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.EventHubConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponsePtrOutput) QnaAzureSearchEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.QnaAzureSearchEndpointId
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponsePtrOutput) QnaAzureSearchEndpointKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.QnaAzureSearchEndpointKey
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponsePtrOutput) QnaRuntimeEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.QnaRuntimeEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponsePtrOutput) StatisticsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.StatisticsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponsePtrOutput) StorageAccountConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponsePtrOutput) SuperUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SuperUser
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountApiPropertiesResponsePtrOutput) WebsiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.WebsiteName
	}).(pulumi.StringPtrOutput)
}

type CognitiveServicesAccountProperties struct {
	ApiProperties              *CognitiveServicesAccountApiProperties `pulumi:"apiProperties"`
	CustomSubDomainName        *string                                `pulumi:"customSubDomainName"`
	Encryption                 *Encryption                            `pulumi:"encryption"`
	NetworkAcls                *NetworkRuleSet                        `pulumi:"networkAcls"`
	PrivateEndpointConnections []PrivateEndpointConnectionType        `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess        *string                                `pulumi:"publicNetworkAccess"`
	UserOwnedStorage           []UserOwnedStorage                     `pulumi:"userOwnedStorage"`
}


func (val *CognitiveServicesAccountProperties) Defaults() *CognitiveServicesAccountProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Encryption = tmp.Encryption.Defaults()

	return &tmp
}





type CognitiveServicesAccountPropertiesInput interface {
	pulumi.Input

	ToCognitiveServicesAccountPropertiesOutput() CognitiveServicesAccountPropertiesOutput
	ToCognitiveServicesAccountPropertiesOutputWithContext(context.Context) CognitiveServicesAccountPropertiesOutput
}

type CognitiveServicesAccountPropertiesArgs struct {
	ApiProperties              CognitiveServicesAccountApiPropertiesPtrInput `pulumi:"apiProperties"`
	CustomSubDomainName        pulumi.StringPtrInput                         `pulumi:"customSubDomainName"`
	Encryption                 EncryptionPtrInput                            `pulumi:"encryption"`
	NetworkAcls                NetworkRuleSetPtrInput                        `pulumi:"networkAcls"`
	PrivateEndpointConnections PrivateEndpointConnectionTypeArrayInput       `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess        pulumi.StringPtrInput                         `pulumi:"publicNetworkAccess"`
	UserOwnedStorage           UserOwnedStorageArrayInput                    `pulumi:"userOwnedStorage"`
}

func (CognitiveServicesAccountPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CognitiveServicesAccountProperties)(nil)).Elem()
}

func (i CognitiveServicesAccountPropertiesArgs) ToCognitiveServicesAccountPropertiesOutput() CognitiveServicesAccountPropertiesOutput {
	return i.ToCognitiveServicesAccountPropertiesOutputWithContext(context.Background())
}

func (i CognitiveServicesAccountPropertiesArgs) ToCognitiveServicesAccountPropertiesOutputWithContext(ctx context.Context) CognitiveServicesAccountPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountPropertiesOutput)
}

func (i CognitiveServicesAccountPropertiesArgs) ToCognitiveServicesAccountPropertiesPtrOutput() CognitiveServicesAccountPropertiesPtrOutput {
	return i.ToCognitiveServicesAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i CognitiveServicesAccountPropertiesArgs) ToCognitiveServicesAccountPropertiesPtrOutputWithContext(ctx context.Context) CognitiveServicesAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountPropertiesOutput).ToCognitiveServicesAccountPropertiesPtrOutputWithContext(ctx)
}









type CognitiveServicesAccountPropertiesPtrInput interface {
	pulumi.Input

	ToCognitiveServicesAccountPropertiesPtrOutput() CognitiveServicesAccountPropertiesPtrOutput
	ToCognitiveServicesAccountPropertiesPtrOutputWithContext(context.Context) CognitiveServicesAccountPropertiesPtrOutput
}

type cognitiveServicesAccountPropertiesPtrType CognitiveServicesAccountPropertiesArgs

func CognitiveServicesAccountPropertiesPtr(v *CognitiveServicesAccountPropertiesArgs) CognitiveServicesAccountPropertiesPtrInput {
	return (*cognitiveServicesAccountPropertiesPtrType)(v)
}

func (*cognitiveServicesAccountPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccountProperties)(nil)).Elem()
}

func (i *cognitiveServicesAccountPropertiesPtrType) ToCognitiveServicesAccountPropertiesPtrOutput() CognitiveServicesAccountPropertiesPtrOutput {
	return i.ToCognitiveServicesAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i *cognitiveServicesAccountPropertiesPtrType) ToCognitiveServicesAccountPropertiesPtrOutputWithContext(ctx context.Context) CognitiveServicesAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountPropertiesPtrOutput)
}

type CognitiveServicesAccountPropertiesOutput struct{ *pulumi.OutputState }

func (CognitiveServicesAccountPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CognitiveServicesAccountProperties)(nil)).Elem()
}

func (o CognitiveServicesAccountPropertiesOutput) ToCognitiveServicesAccountPropertiesOutput() CognitiveServicesAccountPropertiesOutput {
	return o
}

func (o CognitiveServicesAccountPropertiesOutput) ToCognitiveServicesAccountPropertiesOutputWithContext(ctx context.Context) CognitiveServicesAccountPropertiesOutput {
	return o
}

func (o CognitiveServicesAccountPropertiesOutput) ToCognitiveServicesAccountPropertiesPtrOutput() CognitiveServicesAccountPropertiesPtrOutput {
	return o.ToCognitiveServicesAccountPropertiesPtrOutputWithContext(context.Background())
}

func (o CognitiveServicesAccountPropertiesOutput) ToCognitiveServicesAccountPropertiesPtrOutputWithContext(ctx context.Context) CognitiveServicesAccountPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CognitiveServicesAccountProperties) *CognitiveServicesAccountProperties {
		return &v
	}).(CognitiveServicesAccountPropertiesPtrOutput)
}

func (o CognitiveServicesAccountPropertiesOutput) ApiProperties() CognitiveServicesAccountApiPropertiesPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountProperties) *CognitiveServicesAccountApiProperties {
		return v.ApiProperties
	}).(CognitiveServicesAccountApiPropertiesPtrOutput)
}

func (o CognitiveServicesAccountPropertiesOutput) CustomSubDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountProperties) *string { return v.CustomSubDomainName }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountPropertiesOutput) Encryption() EncryptionPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountProperties) *Encryption { return v.Encryption }).(EncryptionPtrOutput)
}

func (o CognitiveServicesAccountPropertiesOutput) NetworkAcls() NetworkRuleSetPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountProperties) *NetworkRuleSet { return v.NetworkAcls }).(NetworkRuleSetPtrOutput)
}

func (o CognitiveServicesAccountPropertiesOutput) PrivateEndpointConnections() PrivateEndpointConnectionTypeArrayOutput {
	return o.ApplyT(func(v CognitiveServicesAccountProperties) []PrivateEndpointConnectionType {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionTypeArrayOutput)
}

func (o CognitiveServicesAccountPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountPropertiesOutput) UserOwnedStorage() UserOwnedStorageArrayOutput {
	return o.ApplyT(func(v CognitiveServicesAccountProperties) []UserOwnedStorage { return v.UserOwnedStorage }).(UserOwnedStorageArrayOutput)
}

type CognitiveServicesAccountPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CognitiveServicesAccountPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccountProperties)(nil)).Elem()
}

func (o CognitiveServicesAccountPropertiesPtrOutput) ToCognitiveServicesAccountPropertiesPtrOutput() CognitiveServicesAccountPropertiesPtrOutput {
	return o
}

func (o CognitiveServicesAccountPropertiesPtrOutput) ToCognitiveServicesAccountPropertiesPtrOutputWithContext(ctx context.Context) CognitiveServicesAccountPropertiesPtrOutput {
	return o
}

func (o CognitiveServicesAccountPropertiesPtrOutput) Elem() CognitiveServicesAccountPropertiesOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountProperties) CognitiveServicesAccountProperties {
		if v != nil {
			return *v
		}
		var ret CognitiveServicesAccountProperties
		return ret
	}).(CognitiveServicesAccountPropertiesOutput)
}

func (o CognitiveServicesAccountPropertiesPtrOutput) ApiProperties() CognitiveServicesAccountApiPropertiesPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountProperties) *CognitiveServicesAccountApiProperties {
		if v == nil {
			return nil
		}
		return v.ApiProperties
	}).(CognitiveServicesAccountApiPropertiesPtrOutput)
}

func (o CognitiveServicesAccountPropertiesPtrOutput) CustomSubDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountProperties) *string {
		if v == nil {
			return nil
		}
		return v.CustomSubDomainName
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountPropertiesPtrOutput) Encryption() EncryptionPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountProperties) *Encryption {
		if v == nil {
			return nil
		}
		return v.Encryption
	}).(EncryptionPtrOutput)
}

func (o CognitiveServicesAccountPropertiesPtrOutput) NetworkAcls() NetworkRuleSetPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountProperties) *NetworkRuleSet {
		if v == nil {
			return nil
		}
		return v.NetworkAcls
	}).(NetworkRuleSetPtrOutput)
}

func (o CognitiveServicesAccountPropertiesPtrOutput) PrivateEndpointConnections() PrivateEndpointConnectionTypeArrayOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountProperties) []PrivateEndpointConnectionType {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionTypeArrayOutput)
}

func (o CognitiveServicesAccountPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountPropertiesPtrOutput) UserOwnedStorage() UserOwnedStorageArrayOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountProperties) []UserOwnedStorage {
		if v == nil {
			return nil
		}
		return v.UserOwnedStorage
	}).(UserOwnedStorageArrayOutput)
}

type CognitiveServicesAccountPropertiesResponse struct {
	ApiProperties              *CognitiveServicesAccountApiPropertiesResponse `pulumi:"apiProperties"`
	Capabilities               []SkuCapabilityResponse                        `pulumi:"capabilities"`
	CustomSubDomainName        *string                                        `pulumi:"customSubDomainName"`
	DateCreated                string                                         `pulumi:"dateCreated"`
	Encryption                 *EncryptionResponse                            `pulumi:"encryption"`
	Endpoint                   string                                         `pulumi:"endpoint"`
	InternalId                 string                                         `pulumi:"internalId"`
	IsMigrated                 bool                                           `pulumi:"isMigrated"`
	NetworkAcls                *NetworkRuleSetResponse                        `pulumi:"networkAcls"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse            `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                                         `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                        `pulumi:"publicNetworkAccess"`
	SkuChangeInfo              CognitiveServicesAccountSkuChangeInfoResponse  `pulumi:"skuChangeInfo"`
	UserOwnedStorage           []UserOwnedStorageResponse                     `pulumi:"userOwnedStorage"`
}


func (val *CognitiveServicesAccountPropertiesResponse) Defaults() *CognitiveServicesAccountPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Encryption = tmp.Encryption.Defaults()

	return &tmp
}





type CognitiveServicesAccountPropertiesResponseInput interface {
	pulumi.Input

	ToCognitiveServicesAccountPropertiesResponseOutput() CognitiveServicesAccountPropertiesResponseOutput
	ToCognitiveServicesAccountPropertiesResponseOutputWithContext(context.Context) CognitiveServicesAccountPropertiesResponseOutput
}

type CognitiveServicesAccountPropertiesResponseArgs struct {
	ApiProperties              CognitiveServicesAccountApiPropertiesResponsePtrInput `pulumi:"apiProperties"`
	Capabilities               SkuCapabilityResponseArrayInput                       `pulumi:"capabilities"`
	CustomSubDomainName        pulumi.StringPtrInput                                 `pulumi:"customSubDomainName"`
	DateCreated                pulumi.StringInput                                    `pulumi:"dateCreated"`
	Encryption                 EncryptionResponsePtrInput                            `pulumi:"encryption"`
	Endpoint                   pulumi.StringInput                                    `pulumi:"endpoint"`
	InternalId                 pulumi.StringInput                                    `pulumi:"internalId"`
	IsMigrated                 pulumi.BoolInput                                      `pulumi:"isMigrated"`
	NetworkAcls                NetworkRuleSetResponsePtrInput                        `pulumi:"networkAcls"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayInput           `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringInput                                    `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrInput                                 `pulumi:"publicNetworkAccess"`
	SkuChangeInfo              CognitiveServicesAccountSkuChangeInfoResponseInput    `pulumi:"skuChangeInfo"`
	UserOwnedStorage           UserOwnedStorageResponseArrayInput                    `pulumi:"userOwnedStorage"`
}

func (CognitiveServicesAccountPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CognitiveServicesAccountPropertiesResponse)(nil)).Elem()
}

func (i CognitiveServicesAccountPropertiesResponseArgs) ToCognitiveServicesAccountPropertiesResponseOutput() CognitiveServicesAccountPropertiesResponseOutput {
	return i.ToCognitiveServicesAccountPropertiesResponseOutputWithContext(context.Background())
}

func (i CognitiveServicesAccountPropertiesResponseArgs) ToCognitiveServicesAccountPropertiesResponseOutputWithContext(ctx context.Context) CognitiveServicesAccountPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountPropertiesResponseOutput)
}

func (i CognitiveServicesAccountPropertiesResponseArgs) ToCognitiveServicesAccountPropertiesResponsePtrOutput() CognitiveServicesAccountPropertiesResponsePtrOutput {
	return i.ToCognitiveServicesAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CognitiveServicesAccountPropertiesResponseArgs) ToCognitiveServicesAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) CognitiveServicesAccountPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountPropertiesResponseOutput).ToCognitiveServicesAccountPropertiesResponsePtrOutputWithContext(ctx)
}









type CognitiveServicesAccountPropertiesResponsePtrInput interface {
	pulumi.Input

	ToCognitiveServicesAccountPropertiesResponsePtrOutput() CognitiveServicesAccountPropertiesResponsePtrOutput
	ToCognitiveServicesAccountPropertiesResponsePtrOutputWithContext(context.Context) CognitiveServicesAccountPropertiesResponsePtrOutput
}

type cognitiveServicesAccountPropertiesResponsePtrType CognitiveServicesAccountPropertiesResponseArgs

func CognitiveServicesAccountPropertiesResponsePtr(v *CognitiveServicesAccountPropertiesResponseArgs) CognitiveServicesAccountPropertiesResponsePtrInput {
	return (*cognitiveServicesAccountPropertiesResponsePtrType)(v)
}

func (*cognitiveServicesAccountPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccountPropertiesResponse)(nil)).Elem()
}

func (i *cognitiveServicesAccountPropertiesResponsePtrType) ToCognitiveServicesAccountPropertiesResponsePtrOutput() CognitiveServicesAccountPropertiesResponsePtrOutput {
	return i.ToCognitiveServicesAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *cognitiveServicesAccountPropertiesResponsePtrType) ToCognitiveServicesAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) CognitiveServicesAccountPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountPropertiesResponsePtrOutput)
}

type CognitiveServicesAccountPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CognitiveServicesAccountPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CognitiveServicesAccountPropertiesResponse)(nil)).Elem()
}

func (o CognitiveServicesAccountPropertiesResponseOutput) ToCognitiveServicesAccountPropertiesResponseOutput() CognitiveServicesAccountPropertiesResponseOutput {
	return o
}

func (o CognitiveServicesAccountPropertiesResponseOutput) ToCognitiveServicesAccountPropertiesResponseOutputWithContext(ctx context.Context) CognitiveServicesAccountPropertiesResponseOutput {
	return o
}

func (o CognitiveServicesAccountPropertiesResponseOutput) ToCognitiveServicesAccountPropertiesResponsePtrOutput() CognitiveServicesAccountPropertiesResponsePtrOutput {
	return o.ToCognitiveServicesAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CognitiveServicesAccountPropertiesResponseOutput) ToCognitiveServicesAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) CognitiveServicesAccountPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CognitiveServicesAccountPropertiesResponse) *CognitiveServicesAccountPropertiesResponse {
		return &v
	}).(CognitiveServicesAccountPropertiesResponsePtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) ApiProperties() CognitiveServicesAccountApiPropertiesResponsePtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) *CognitiveServicesAccountApiPropertiesResponse {
		return v.ApiProperties
	}).(CognitiveServicesAccountApiPropertiesResponsePtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) Capabilities() SkuCapabilityResponseArrayOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) []SkuCapabilityResponse { return v.Capabilities }).(SkuCapabilityResponseArrayOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) CustomSubDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) *string { return v.CustomSubDomainName }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) string { return v.DateCreated }).(pulumi.StringOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) *EncryptionResponse { return v.Encryption }).(EncryptionResponsePtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) string { return v.InternalId }).(pulumi.StringOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) IsMigrated() pulumi.BoolOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) bool { return v.IsMigrated }).(pulumi.BoolOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) NetworkAcls() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) *NetworkRuleSetResponse { return v.NetworkAcls }).(NetworkRuleSetResponsePtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) SkuChangeInfo() CognitiveServicesAccountSkuChangeInfoResponseOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) CognitiveServicesAccountSkuChangeInfoResponse {
		return v.SkuChangeInfo
	}).(CognitiveServicesAccountSkuChangeInfoResponseOutput)
}

func (o CognitiveServicesAccountPropertiesResponseOutput) UserOwnedStorage() UserOwnedStorageResponseArrayOutput {
	return o.ApplyT(func(v CognitiveServicesAccountPropertiesResponse) []UserOwnedStorageResponse {
		return v.UserOwnedStorage
	}).(UserOwnedStorageResponseArrayOutput)
}

type CognitiveServicesAccountPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CognitiveServicesAccountPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccountPropertiesResponse)(nil)).Elem()
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) ToCognitiveServicesAccountPropertiesResponsePtrOutput() CognitiveServicesAccountPropertiesResponsePtrOutput {
	return o
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) ToCognitiveServicesAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) CognitiveServicesAccountPropertiesResponsePtrOutput {
	return o
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) Elem() CognitiveServicesAccountPropertiesResponseOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) CognitiveServicesAccountPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CognitiveServicesAccountPropertiesResponse
		return ret
	}).(CognitiveServicesAccountPropertiesResponseOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) ApiProperties() CognitiveServicesAccountApiPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) *CognitiveServicesAccountApiPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.ApiProperties
	}).(CognitiveServicesAccountApiPropertiesResponsePtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) Capabilities() SkuCapabilityResponseArrayOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) []SkuCapabilityResponse {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(SkuCapabilityResponseArrayOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) CustomSubDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomSubDomainName
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) DateCreated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DateCreated
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) *EncryptionResponse {
		if v == nil {
			return nil
		}
		return v.Encryption
	}).(EncryptionResponsePtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Endpoint
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) InternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.InternalId
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) IsMigrated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsMigrated
	}).(pulumi.BoolPtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) NetworkAcls() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) *NetworkRuleSetResponse {
		if v == nil {
			return nil
		}
		return v.NetworkAcls
	}).(NetworkRuleSetResponsePtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) []PrivateEndpointConnectionResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) SkuChangeInfo() CognitiveServicesAccountSkuChangeInfoResponsePtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) *CognitiveServicesAccountSkuChangeInfoResponse {
		if v == nil {
			return nil
		}
		return &v.SkuChangeInfo
	}).(CognitiveServicesAccountSkuChangeInfoResponsePtrOutput)
}

func (o CognitiveServicesAccountPropertiesResponsePtrOutput) UserOwnedStorage() UserOwnedStorageResponseArrayOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountPropertiesResponse) []UserOwnedStorageResponse {
		if v == nil {
			return nil
		}
		return v.UserOwnedStorage
	}).(UserOwnedStorageResponseArrayOutput)
}

type CognitiveServicesAccountSkuChangeInfoResponse struct {
	CountOfDowngrades              float64 `pulumi:"countOfDowngrades"`
	CountOfUpgradesAfterDowngrades float64 `pulumi:"countOfUpgradesAfterDowngrades"`
	LastChangeDate                 string  `pulumi:"lastChangeDate"`
}





type CognitiveServicesAccountSkuChangeInfoResponseInput interface {
	pulumi.Input

	ToCognitiveServicesAccountSkuChangeInfoResponseOutput() CognitiveServicesAccountSkuChangeInfoResponseOutput
	ToCognitiveServicesAccountSkuChangeInfoResponseOutputWithContext(context.Context) CognitiveServicesAccountSkuChangeInfoResponseOutput
}

type CognitiveServicesAccountSkuChangeInfoResponseArgs struct {
	CountOfDowngrades              pulumi.Float64Input `pulumi:"countOfDowngrades"`
	CountOfUpgradesAfterDowngrades pulumi.Float64Input `pulumi:"countOfUpgradesAfterDowngrades"`
	LastChangeDate                 pulumi.StringInput  `pulumi:"lastChangeDate"`
}

func (CognitiveServicesAccountSkuChangeInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CognitiveServicesAccountSkuChangeInfoResponse)(nil)).Elem()
}

func (i CognitiveServicesAccountSkuChangeInfoResponseArgs) ToCognitiveServicesAccountSkuChangeInfoResponseOutput() CognitiveServicesAccountSkuChangeInfoResponseOutput {
	return i.ToCognitiveServicesAccountSkuChangeInfoResponseOutputWithContext(context.Background())
}

func (i CognitiveServicesAccountSkuChangeInfoResponseArgs) ToCognitiveServicesAccountSkuChangeInfoResponseOutputWithContext(ctx context.Context) CognitiveServicesAccountSkuChangeInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountSkuChangeInfoResponseOutput)
}

func (i CognitiveServicesAccountSkuChangeInfoResponseArgs) ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutput() CognitiveServicesAccountSkuChangeInfoResponsePtrOutput {
	return i.ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutputWithContext(context.Background())
}

func (i CognitiveServicesAccountSkuChangeInfoResponseArgs) ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutputWithContext(ctx context.Context) CognitiveServicesAccountSkuChangeInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountSkuChangeInfoResponseOutput).ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutputWithContext(ctx)
}









type CognitiveServicesAccountSkuChangeInfoResponsePtrInput interface {
	pulumi.Input

	ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutput() CognitiveServicesAccountSkuChangeInfoResponsePtrOutput
	ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutputWithContext(context.Context) CognitiveServicesAccountSkuChangeInfoResponsePtrOutput
}

type cognitiveServicesAccountSkuChangeInfoResponsePtrType CognitiveServicesAccountSkuChangeInfoResponseArgs

func CognitiveServicesAccountSkuChangeInfoResponsePtr(v *CognitiveServicesAccountSkuChangeInfoResponseArgs) CognitiveServicesAccountSkuChangeInfoResponsePtrInput {
	return (*cognitiveServicesAccountSkuChangeInfoResponsePtrType)(v)
}

func (*cognitiveServicesAccountSkuChangeInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccountSkuChangeInfoResponse)(nil)).Elem()
}

func (i *cognitiveServicesAccountSkuChangeInfoResponsePtrType) ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutput() CognitiveServicesAccountSkuChangeInfoResponsePtrOutput {
	return i.ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutputWithContext(context.Background())
}

func (i *cognitiveServicesAccountSkuChangeInfoResponsePtrType) ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutputWithContext(ctx context.Context) CognitiveServicesAccountSkuChangeInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountSkuChangeInfoResponsePtrOutput)
}

type CognitiveServicesAccountSkuChangeInfoResponseOutput struct{ *pulumi.OutputState }

func (CognitiveServicesAccountSkuChangeInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CognitiveServicesAccountSkuChangeInfoResponse)(nil)).Elem()
}

func (o CognitiveServicesAccountSkuChangeInfoResponseOutput) ToCognitiveServicesAccountSkuChangeInfoResponseOutput() CognitiveServicesAccountSkuChangeInfoResponseOutput {
	return o
}

func (o CognitiveServicesAccountSkuChangeInfoResponseOutput) ToCognitiveServicesAccountSkuChangeInfoResponseOutputWithContext(ctx context.Context) CognitiveServicesAccountSkuChangeInfoResponseOutput {
	return o
}

func (o CognitiveServicesAccountSkuChangeInfoResponseOutput) ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutput() CognitiveServicesAccountSkuChangeInfoResponsePtrOutput {
	return o.ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutputWithContext(context.Background())
}

func (o CognitiveServicesAccountSkuChangeInfoResponseOutput) ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutputWithContext(ctx context.Context) CognitiveServicesAccountSkuChangeInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CognitiveServicesAccountSkuChangeInfoResponse) *CognitiveServicesAccountSkuChangeInfoResponse {
		return &v
	}).(CognitiveServicesAccountSkuChangeInfoResponsePtrOutput)
}

func (o CognitiveServicesAccountSkuChangeInfoResponseOutput) CountOfDowngrades() pulumi.Float64Output {
	return o.ApplyT(func(v CognitiveServicesAccountSkuChangeInfoResponse) float64 { return v.CountOfDowngrades }).(pulumi.Float64Output)
}

func (o CognitiveServicesAccountSkuChangeInfoResponseOutput) CountOfUpgradesAfterDowngrades() pulumi.Float64Output {
	return o.ApplyT(func(v CognitiveServicesAccountSkuChangeInfoResponse) float64 { return v.CountOfUpgradesAfterDowngrades }).(pulumi.Float64Output)
}

func (o CognitiveServicesAccountSkuChangeInfoResponseOutput) LastChangeDate() pulumi.StringOutput {
	return o.ApplyT(func(v CognitiveServicesAccountSkuChangeInfoResponse) string { return v.LastChangeDate }).(pulumi.StringOutput)
}

type CognitiveServicesAccountSkuChangeInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (CognitiveServicesAccountSkuChangeInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccountSkuChangeInfoResponse)(nil)).Elem()
}

func (o CognitiveServicesAccountSkuChangeInfoResponsePtrOutput) ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutput() CognitiveServicesAccountSkuChangeInfoResponsePtrOutput {
	return o
}

func (o CognitiveServicesAccountSkuChangeInfoResponsePtrOutput) ToCognitiveServicesAccountSkuChangeInfoResponsePtrOutputWithContext(ctx context.Context) CognitiveServicesAccountSkuChangeInfoResponsePtrOutput {
	return o
}

func (o CognitiveServicesAccountSkuChangeInfoResponsePtrOutput) Elem() CognitiveServicesAccountSkuChangeInfoResponseOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountSkuChangeInfoResponse) CognitiveServicesAccountSkuChangeInfoResponse {
		if v != nil {
			return *v
		}
		var ret CognitiveServicesAccountSkuChangeInfoResponse
		return ret
	}).(CognitiveServicesAccountSkuChangeInfoResponseOutput)
}

func (o CognitiveServicesAccountSkuChangeInfoResponsePtrOutput) CountOfDowngrades() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountSkuChangeInfoResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.CountOfDowngrades
	}).(pulumi.Float64PtrOutput)
}

func (o CognitiveServicesAccountSkuChangeInfoResponsePtrOutput) CountOfUpgradesAfterDowngrades() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountSkuChangeInfoResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.CountOfUpgradesAfterDowngrades
	}).(pulumi.Float64PtrOutput)
}

func (o CognitiveServicesAccountSkuChangeInfoResponsePtrOutput) LastChangeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccountSkuChangeInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastChangeDate
	}).(pulumi.StringPtrOutput)
}

type Encryption struct {
	KeySource          *string             `pulumi:"keySource"`
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
}


func (val *Encryption) Defaults() *Encryption {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		keySource_ := "Microsoft.KeyVault"
		tmp.KeySource = &keySource_
	}
	return &tmp
}





type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

type EncryptionArgs struct {
	KeySource          pulumi.StringPtrInput      `pulumi:"keySource"`
	KeyVaultProperties KeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
}

func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}









type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

type EncryptionOutput struct{ *pulumi.OutputState }

func (EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (o EncryptionOutput) ToEncryptionOutput() EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o.ToEncryptionPtrOutputWithContext(context.Background())
}

func (o EncryptionOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Encryption) *Encryption {
		return &v
	}).(EncryptionPtrOutput)
}

func (o EncryptionOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Encryption) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v Encryption) *KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesPtrOutput)
}

type EncryptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) Elem() EncryptionOutput {
	return o.ApplyT(func(v *Encryption) Encryption {
		if v != nil {
			return *v
		}
		var ret Encryption
		return ret
	}).(EncryptionOutput)
}

func (o EncryptionPtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *Encryption) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

type EncryptionResponse struct {
	KeySource          *string                     `pulumi:"keySource"`
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
}


func (val *EncryptionResponse) Defaults() *EncryptionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		keySource_ := "Microsoft.KeyVault"
		tmp.KeySource = &keySource_
	}
	return &tmp
}





type EncryptionResponseInput interface {
	pulumi.Input

	ToEncryptionResponseOutput() EncryptionResponseOutput
	ToEncryptionResponseOutputWithContext(context.Context) EncryptionResponseOutput
}

type EncryptionResponseArgs struct {
	KeySource          pulumi.StringPtrInput              `pulumi:"keySource"`
	KeyVaultProperties KeyVaultPropertiesResponsePtrInput `pulumi:"keyVaultProperties"`
}

func (EncryptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return i.ToEncryptionResponseOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput)
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput).ToEncryptionResponsePtrOutputWithContext(ctx)
}









type EncryptionResponsePtrInput interface {
	pulumi.Input

	ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput
	ToEncryptionResponsePtrOutputWithContext(context.Context) EncryptionResponsePtrOutput
}

type encryptionResponsePtrType EncryptionResponseArgs

func EncryptionResponsePtr(v *EncryptionResponseArgs) EncryptionResponsePtrInput {
	return (*encryptionResponsePtrType)(v)
}

func (*encryptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponsePtrOutput)
}

type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionResponse) *EncryptionResponse {
		return &v
	}).(EncryptionResponsePtrOutput)
}

func (o EncryptionResponseOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

type EncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) Elem() EncryptionResponseOutput {
	return o.ApplyT(func(v *EncryptionResponse) EncryptionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionResponse
		return ret
	}).(EncryptionResponseOutput)
}

func (o EncryptionResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

type Identity struct {
	Type                   *IdentityType                   `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentity `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   IdentityTypePtrInput         `pulumi:"type"`
	UserAssignedIdentities UserAssignedIdentityMapInput `pulumi:"userAssignedIdentities"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() IdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *IdentityType { return v.Type }).(IdentityTypePtrOutput)
}

func (o IdentityOutput) UserAssignedIdentities() UserAssignedIdentityMapOutput {
	return o.ApplyT(func(v Identity) map[string]UserAssignedIdentity { return v.UserAssignedIdentities }).(UserAssignedIdentityMapOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() IdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *IdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(IdentityTypePtrOutput)
}

func (o IdentityPtrOutput) UserAssignedIdentities() UserAssignedIdentityMapOutput {
	return o.ApplyT(func(v *Identity) map[string]UserAssignedIdentity {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityMapOutput)
}

type IdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                   `pulumi:"principalId"`
	TenantId               pulumi.StringInput                   `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                `pulumi:"type"`
	UserAssignedIdentities UserAssignedIdentityResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]UserAssignedIdentityResponse { return v.UserAssignedIdentities }).(UserAssignedIdentityResponseMapOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type IpRule struct {
	Value string `pulumi:"value"`
}





type IpRuleInput interface {
	pulumi.Input

	ToIpRuleOutput() IpRuleOutput
	ToIpRuleOutputWithContext(context.Context) IpRuleOutput
}

type IpRuleArgs struct {
	Value pulumi.StringInput `pulumi:"value"`
}

func (IpRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRule)(nil)).Elem()
}

func (i IpRuleArgs) ToIpRuleOutput() IpRuleOutput {
	return i.ToIpRuleOutputWithContext(context.Background())
}

func (i IpRuleArgs) ToIpRuleOutputWithContext(ctx context.Context) IpRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleOutput)
}





type IpRuleArrayInput interface {
	pulumi.Input

	ToIpRuleArrayOutput() IpRuleArrayOutput
	ToIpRuleArrayOutputWithContext(context.Context) IpRuleArrayOutput
}

type IpRuleArray []IpRuleInput

func (IpRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRule)(nil)).Elem()
}

func (i IpRuleArray) ToIpRuleArrayOutput() IpRuleArrayOutput {
	return i.ToIpRuleArrayOutputWithContext(context.Background())
}

func (i IpRuleArray) ToIpRuleArrayOutputWithContext(ctx context.Context) IpRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleArrayOutput)
}

type IpRuleOutput struct{ *pulumi.OutputState }

func (IpRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRule)(nil)).Elem()
}

func (o IpRuleOutput) ToIpRuleOutput() IpRuleOutput {
	return o
}

func (o IpRuleOutput) ToIpRuleOutputWithContext(ctx context.Context) IpRuleOutput {
	return o
}

func (o IpRuleOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v IpRule) string { return v.Value }).(pulumi.StringOutput)
}

type IpRuleArrayOutput struct{ *pulumi.OutputState }

func (IpRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRule)(nil)).Elem()
}

func (o IpRuleArrayOutput) ToIpRuleArrayOutput() IpRuleArrayOutput {
	return o
}

func (o IpRuleArrayOutput) ToIpRuleArrayOutputWithContext(ctx context.Context) IpRuleArrayOutput {
	return o
}

func (o IpRuleArrayOutput) Index(i pulumi.IntInput) IpRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpRule {
		return vs[0].([]IpRule)[vs[1].(int)]
	}).(IpRuleOutput)
}

type IpRuleResponse struct {
	Value string `pulumi:"value"`
}





type IpRuleResponseInput interface {
	pulumi.Input

	ToIpRuleResponseOutput() IpRuleResponseOutput
	ToIpRuleResponseOutputWithContext(context.Context) IpRuleResponseOutput
}

type IpRuleResponseArgs struct {
	Value pulumi.StringInput `pulumi:"value"`
}

func (IpRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRuleResponse)(nil)).Elem()
}

func (i IpRuleResponseArgs) ToIpRuleResponseOutput() IpRuleResponseOutput {
	return i.ToIpRuleResponseOutputWithContext(context.Background())
}

func (i IpRuleResponseArgs) ToIpRuleResponseOutputWithContext(ctx context.Context) IpRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleResponseOutput)
}





type IpRuleResponseArrayInput interface {
	pulumi.Input

	ToIpRuleResponseArrayOutput() IpRuleResponseArrayOutput
	ToIpRuleResponseArrayOutputWithContext(context.Context) IpRuleResponseArrayOutput
}

type IpRuleResponseArray []IpRuleResponseInput

func (IpRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRuleResponse)(nil)).Elem()
}

func (i IpRuleResponseArray) ToIpRuleResponseArrayOutput() IpRuleResponseArrayOutput {
	return i.ToIpRuleResponseArrayOutputWithContext(context.Background())
}

func (i IpRuleResponseArray) ToIpRuleResponseArrayOutputWithContext(ctx context.Context) IpRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleResponseArrayOutput)
}

type IpRuleResponseOutput struct{ *pulumi.OutputState }

func (IpRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRuleResponse)(nil)).Elem()
}

func (o IpRuleResponseOutput) ToIpRuleResponseOutput() IpRuleResponseOutput {
	return o
}

func (o IpRuleResponseOutput) ToIpRuleResponseOutputWithContext(ctx context.Context) IpRuleResponseOutput {
	return o
}

func (o IpRuleResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v IpRuleResponse) string { return v.Value }).(pulumi.StringOutput)
}

type IpRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IpRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRuleResponse)(nil)).Elem()
}

func (o IpRuleResponseArrayOutput) ToIpRuleResponseArrayOutput() IpRuleResponseArrayOutput {
	return o
}

func (o IpRuleResponseArrayOutput) ToIpRuleResponseArrayOutputWithContext(ctx context.Context) IpRuleResponseArrayOutput {
	return o
}

func (o IpRuleResponseArrayOutput) Index(i pulumi.IntInput) IpRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpRuleResponse {
		return vs[0].([]IpRuleResponse)[vs[1].(int)]
	}).(IpRuleResponseOutput)
}

type KeyVaultProperties struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	KeyVersion  *string `pulumi:"keyVersion"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput).ToKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type KeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput
	ToKeyVaultPropertiesPtrOutputWithContext(context.Context) KeyVaultPropertiesPtrOutput
}

type keyVaultPropertiesPtrType KeyVaultPropertiesArgs

func KeyVaultPropertiesPtr(v *KeyVaultPropertiesArgs) KeyVaultPropertiesPtrInput {
	return (*keyVaultPropertiesPtrType)(v)
}

func (*keyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesPtrOutput)
}

type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultProperties) *KeyVaultProperties {
		return &v
	}).(KeyVaultPropertiesPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) Elem() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v *KeyVaultProperties) KeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret KeyVaultProperties
		return ret
	}).(KeyVaultPropertiesOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	KeyVersion  *string `pulumi:"keyVersion"`
}





type KeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput
	ToKeyVaultPropertiesResponseOutputWithContext(context.Context) KeyVaultPropertiesResponseOutput
}

type KeyVaultPropertiesResponseArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
}

func (KeyVaultPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return i.ToKeyVaultPropertiesResponseOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput)
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput).ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx)
}









type KeyVaultPropertiesResponsePtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput
	ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Context) KeyVaultPropertiesResponsePtrOutput
}

type keyVaultPropertiesResponsePtrType KeyVaultPropertiesResponseArgs

func KeyVaultPropertiesResponsePtr(v *KeyVaultPropertiesResponseArgs) KeyVaultPropertiesResponsePtrInput {
	return (*keyVaultPropertiesResponsePtrType)(v)
}

func (*keyVaultPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponsePtrOutput)
}

type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultPropertiesResponse) *KeyVaultPropertiesResponse {
		return &v
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) Elem() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) KeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultPropertiesResponse
		return ret
	}).(KeyVaultPropertiesResponseOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type NetworkRuleSet struct {
	DefaultAction       *string              `pulumi:"defaultAction"`
	IpRules             []IpRule             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRule `pulumi:"virtualNetworkRules"`
}





type NetworkRuleSetInput interface {
	pulumi.Input

	ToNetworkRuleSetOutput() NetworkRuleSetOutput
	ToNetworkRuleSetOutputWithContext(context.Context) NetworkRuleSetOutput
}

type NetworkRuleSetArgs struct {
	DefaultAction       pulumi.StringPtrInput        `pulumi:"defaultAction"`
	IpRules             IpRuleArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules VirtualNetworkRuleArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return i.ToNetworkRuleSetOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput)
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput).ToNetworkRuleSetPtrOutputWithContext(ctx)
}









type NetworkRuleSetPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput
	ToNetworkRuleSetPtrOutputWithContext(context.Context) NetworkRuleSetPtrOutput
}

type networkRuleSetPtrType NetworkRuleSetArgs

func NetworkRuleSetPtr(v *NetworkRuleSetArgs) NetworkRuleSetPtrInput {
	return (*networkRuleSetPtrType)(v)
}

func (*networkRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPtrOutput)
}

type NetworkRuleSetOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSet) *NetworkRuleSet {
		return &v
	}).(NetworkRuleSetPtrOutput)
}

func (o NetworkRuleSetOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSet) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetOutput) IpRules() IpRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []IpRule { return v.IpRules }).(IpRuleArrayOutput)
}

func (o NetworkRuleSetOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []VirtualNetworkRule { return v.VirtualNetworkRules }).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) Elem() NetworkRuleSetOutput {
	return o.ApplyT(func(v *NetworkRuleSet) NetworkRuleSet {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSet
		return ret
	}).(NetworkRuleSetOutput)
}

func (o NetworkRuleSetPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPtrOutput) IpRules() IpRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []IpRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IpRuleArrayOutput)
}

func (o NetworkRuleSetPtrOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []VirtualNetworkRule {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetResponse struct {
	DefaultAction       *string                      `pulumi:"defaultAction"`
	IpRules             []IpRuleResponse             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}





type NetworkRuleSetResponseInput interface {
	pulumi.Input

	ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput
	ToNetworkRuleSetResponseOutputWithContext(context.Context) NetworkRuleSetResponseOutput
}

type NetworkRuleSetResponseArgs struct {
	DefaultAction       pulumi.StringPtrInput                `pulumi:"defaultAction"`
	IpRules             IpRuleResponseArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules VirtualNetworkRuleResponseArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return i.ToNetworkRuleSetResponseOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput)
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput).ToNetworkRuleSetResponsePtrOutputWithContext(ctx)
}









type NetworkRuleSetResponsePtrInput interface {
	pulumi.Input

	ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput
	ToNetworkRuleSetResponsePtrOutputWithContext(context.Context) NetworkRuleSetResponsePtrOutput
}

type networkRuleSetResponsePtrType NetworkRuleSetResponseArgs

func NetworkRuleSetResponsePtr(v *NetworkRuleSetResponseArgs) NetworkRuleSetResponsePtrInput {
	return (*networkRuleSetResponsePtrType)(v)
}

func (*networkRuleSetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponsePtrOutput)
}

type NetworkRuleSetResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSetResponse) *NetworkRuleSetResponse {
		return &v
	}).(NetworkRuleSetResponsePtrOutput)
}

func (o NetworkRuleSetResponseOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponseOutput) IpRules() IpRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []IpRuleResponse { return v.IpRules }).(IpRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponseOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

type NetworkRuleSetResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) Elem() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) NetworkRuleSetResponse {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetResponse
		return ret
	}).(NetworkRuleSetResponseOutput)
}

func (o NetworkRuleSetResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponsePtrOutput) IpRules() IpRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []IpRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IpRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponsePtrOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []VirtualNetworkRuleResponse {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleResponseArrayOutput)
}

type PrivateEndpointConnectionType struct {
	Location   *string                              `pulumi:"location"`
	Properties *PrivateEndpointConnectionProperties `pulumi:"properties"`
}





type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	Location   pulumi.StringPtrInput                       `pulumi:"location"`
	Properties PrivateEndpointConnectionPropertiesPtrInput `pulumi:"properties"`
}

func (PrivateEndpointConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return i.ToPrivateEndpointConnectionTypeOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput)
}





type PrivateEndpointConnectionTypeArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput
	ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Context) PrivateEndpointConnectionTypeArrayOutput
}

type PrivateEndpointConnectionTypeArray []PrivateEndpointConnectionTypeInput

func (PrivateEndpointConnectionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return i.ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeArrayOutput)
}

type PrivateEndpointConnectionTypeOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionTypeOutput) Properties() PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *PrivateEndpointConnectionProperties { return v.Properties }).(PrivateEndpointConnectionPropertiesPtrOutput)
}

type PrivateEndpointConnectionTypeArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionType {
		return vs[0].([]PrivateEndpointConnectionType)[vs[1].(int)]
	}).(PrivateEndpointConnectionTypeOutput)
}

type PrivateEndpointConnectionProperties struct {
	GroupIds                          []string                          `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput
	ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesOutput
}

type PrivateEndpointConnectionPropertiesArgs struct {
	GroupIds                          pulumi.StringArrayInput                `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return i.ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput)
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput).ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput
	ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPtrOutput
}

type privateEndpointConnectionPropertiesPtrType PrivateEndpointConnectionPropertiesArgs

func PrivateEndpointConnectionPropertiesPtr(v *PrivateEndpointConnectionPropertiesArgs) PrivateEndpointConnectionPropertiesPtrInput {
	return (*privateEndpointConnectionPropertiesPtrType)(v)
}

func (*privateEndpointConnectionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPtrOutput)
}

type PrivateEndpointConnectionPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionProperties) *PrivateEndpointConnectionProperties {
		return &v
	}).(PrivateEndpointConnectionPropertiesPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateEndpointConnectionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) Elem() PrivateEndpointConnectionPropertiesOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) PrivateEndpointConnectionProperties {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionProperties
		return ret
	}).(PrivateEndpointConnectionPropertiesOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) []string {
		if v == nil {
			return nil
		}
		return v.GroupIds
	}).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return &v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	GroupIds                          []string                                  `pulumi:"groupIds"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput
	ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponseOutput
}

type PrivateEndpointConnectionPropertiesResponseArgs struct {
	GroupIds                          pulumi.StringArrayInput                        `pulumi:"groupIds"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput).ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput
	ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput
}

type privateEndpointConnectionPropertiesResponsePtrType PrivateEndpointConnectionPropertiesResponseArgs

func PrivateEndpointConnectionPropertiesResponsePtr(v *PrivateEndpointConnectionPropertiesResponseArgs) PrivateEndpointConnectionPropertiesResponsePtrInput {
	return (*privateEndpointConnectionPropertiesResponsePtrType)(v)
}

func (*privateEndpointConnectionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponse {
		return &v
	}).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

type PrivateEndpointConnectionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) Elem() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) PrivateEndpointConnectionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesResponse
		return ret
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.GroupIds
	}).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateLinkServiceConnectionStateResponse {
		if v == nil {
			return nil
		}
		return &v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	Etag       string                                       `pulumi:"etag"`
	Id         string                                       `pulumi:"id"`
	Location   *string                                      `pulumi:"location"`
	Name       string                                       `pulumi:"name"`
	Properties *PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	Type       string                                       `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Etag       pulumi.StringInput                                  `pulumi:"etag"`
	Id         pulumi.StringInput                                  `pulumi:"id"`
	Location   pulumi.StringPtrInput                               `pulumi:"location"`
	Name       pulumi.StringInput                                  `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponsePtrInput `pulumi:"properties"`
	Type       pulumi.StringInput                                  `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuCapabilityResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type SkuCapabilityResponseInput interface {
	pulumi.Input

	ToSkuCapabilityResponseOutput() SkuCapabilityResponseOutput
	ToSkuCapabilityResponseOutputWithContext(context.Context) SkuCapabilityResponseOutput
}

type SkuCapabilityResponseArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (SkuCapabilityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapabilityResponse)(nil)).Elem()
}

func (i SkuCapabilityResponseArgs) ToSkuCapabilityResponseOutput() SkuCapabilityResponseOutput {
	return i.ToSkuCapabilityResponseOutputWithContext(context.Background())
}

func (i SkuCapabilityResponseArgs) ToSkuCapabilityResponseOutputWithContext(ctx context.Context) SkuCapabilityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapabilityResponseOutput)
}





type SkuCapabilityResponseArrayInput interface {
	pulumi.Input

	ToSkuCapabilityResponseArrayOutput() SkuCapabilityResponseArrayOutput
	ToSkuCapabilityResponseArrayOutputWithContext(context.Context) SkuCapabilityResponseArrayOutput
}

type SkuCapabilityResponseArray []SkuCapabilityResponseInput

func (SkuCapabilityResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SkuCapabilityResponse)(nil)).Elem()
}

func (i SkuCapabilityResponseArray) ToSkuCapabilityResponseArrayOutput() SkuCapabilityResponseArrayOutput {
	return i.ToSkuCapabilityResponseArrayOutputWithContext(context.Background())
}

func (i SkuCapabilityResponseArray) ToSkuCapabilityResponseArrayOutputWithContext(ctx context.Context) SkuCapabilityResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapabilityResponseArrayOutput)
}

type SkuCapabilityResponseOutput struct{ *pulumi.OutputState }

func (SkuCapabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapabilityResponse)(nil)).Elem()
}

func (o SkuCapabilityResponseOutput) ToSkuCapabilityResponseOutput() SkuCapabilityResponseOutput {
	return o
}

func (o SkuCapabilityResponseOutput) ToSkuCapabilityResponseOutputWithContext(ctx context.Context) SkuCapabilityResponseOutput {
	return o
}

func (o SkuCapabilityResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuCapabilityResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuCapabilityResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuCapabilityResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SkuCapabilityResponseArrayOutput struct{ *pulumi.OutputState }

func (SkuCapabilityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SkuCapabilityResponse)(nil)).Elem()
}

func (o SkuCapabilityResponseArrayOutput) ToSkuCapabilityResponseArrayOutput() SkuCapabilityResponseArrayOutput {
	return o
}

func (o SkuCapabilityResponseArrayOutput) ToSkuCapabilityResponseArrayOutputWithContext(ctx context.Context) SkuCapabilityResponseArrayOutput {
	return o
}

func (o SkuCapabilityResponseArrayOutput) Index(i pulumi.IntInput) SkuCapabilityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SkuCapabilityResponse {
		return vs[0].([]SkuCapabilityResponse)[vs[1].(int)]
	}).(SkuCapabilityResponseOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type UserAssignedIdentity struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}





type UserAssignedIdentityInput interface {
	pulumi.Input

	ToUserAssignedIdentityOutput() UserAssignedIdentityOutput
	ToUserAssignedIdentityOutputWithContext(context.Context) UserAssignedIdentityOutput
}

type UserAssignedIdentityArgs struct {
	ClientId    pulumi.StringPtrInput `pulumi:"clientId"`
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
}

func (UserAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentity)(nil)).Elem()
}

func (i UserAssignedIdentityArgs) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return i.ToUserAssignedIdentityOutputWithContext(context.Background())
}

func (i UserAssignedIdentityArgs) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityOutput)
}





type UserAssignedIdentityMapInput interface {
	pulumi.Input

	ToUserAssignedIdentityMapOutput() UserAssignedIdentityMapOutput
	ToUserAssignedIdentityMapOutputWithContext(context.Context) UserAssignedIdentityMapOutput
}

type UserAssignedIdentityMap map[string]UserAssignedIdentityInput

func (UserAssignedIdentityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentity)(nil)).Elem()
}

func (i UserAssignedIdentityMap) ToUserAssignedIdentityMapOutput() UserAssignedIdentityMapOutput {
	return i.ToUserAssignedIdentityMapOutputWithContext(context.Background())
}

func (i UserAssignedIdentityMap) ToUserAssignedIdentityMapOutputWithContext(ctx context.Context) UserAssignedIdentityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityMapOutput)
}

type UserAssignedIdentityOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentity)(nil)).Elem()
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentity) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserAssignedIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentity)(nil)).Elem()
}

func (o UserAssignedIdentityMapOutput) ToUserAssignedIdentityMapOutput() UserAssignedIdentityMapOutput {
	return o
}

func (o UserAssignedIdentityMapOutput) ToUserAssignedIdentityMapOutputWithContext(ctx context.Context) UserAssignedIdentityMapOutput {
	return o
}

func (o UserAssignedIdentityMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentity {
		return vs[0].(map[string]UserAssignedIdentity)[vs[1].(string)]
	}).(UserAssignedIdentityOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}





type UserAssignedIdentityResponseInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput
	ToUserAssignedIdentityResponseOutputWithContext(context.Context) UserAssignedIdentityResponseOutput
}

type UserAssignedIdentityResponseArgs struct {
	ClientId    pulumi.StringPtrInput `pulumi:"clientId"`
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
}

func (UserAssignedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return i.ToUserAssignedIdentityResponseOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseOutput)
}





type UserAssignedIdentityResponseMapInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput
	ToUserAssignedIdentityResponseMapOutputWithContext(context.Context) UserAssignedIdentityResponseMapOutput
}

type UserAssignedIdentityResponseMap map[string]UserAssignedIdentityResponseInput

func (UserAssignedIdentityResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return i.ToUserAssignedIdentityResponseMapOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseMapOutput)
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

type UserOwnedStorage struct {
	ResourceId *string `pulumi:"resourceId"`
}





type UserOwnedStorageInput interface {
	pulumi.Input

	ToUserOwnedStorageOutput() UserOwnedStorageOutput
	ToUserOwnedStorageOutputWithContext(context.Context) UserOwnedStorageOutput
}

type UserOwnedStorageArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (UserOwnedStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserOwnedStorage)(nil)).Elem()
}

func (i UserOwnedStorageArgs) ToUserOwnedStorageOutput() UserOwnedStorageOutput {
	return i.ToUserOwnedStorageOutputWithContext(context.Background())
}

func (i UserOwnedStorageArgs) ToUserOwnedStorageOutputWithContext(ctx context.Context) UserOwnedStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOwnedStorageOutput)
}





type UserOwnedStorageArrayInput interface {
	pulumi.Input

	ToUserOwnedStorageArrayOutput() UserOwnedStorageArrayOutput
	ToUserOwnedStorageArrayOutputWithContext(context.Context) UserOwnedStorageArrayOutput
}

type UserOwnedStorageArray []UserOwnedStorageInput

func (UserOwnedStorageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserOwnedStorage)(nil)).Elem()
}

func (i UserOwnedStorageArray) ToUserOwnedStorageArrayOutput() UserOwnedStorageArrayOutput {
	return i.ToUserOwnedStorageArrayOutputWithContext(context.Background())
}

func (i UserOwnedStorageArray) ToUserOwnedStorageArrayOutputWithContext(ctx context.Context) UserOwnedStorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOwnedStorageArrayOutput)
}

type UserOwnedStorageOutput struct{ *pulumi.OutputState }

func (UserOwnedStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserOwnedStorage)(nil)).Elem()
}

func (o UserOwnedStorageOutput) ToUserOwnedStorageOutput() UserOwnedStorageOutput {
	return o
}

func (o UserOwnedStorageOutput) ToUserOwnedStorageOutputWithContext(ctx context.Context) UserOwnedStorageOutput {
	return o
}

func (o UserOwnedStorageOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserOwnedStorage) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type UserOwnedStorageArrayOutput struct{ *pulumi.OutputState }

func (UserOwnedStorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserOwnedStorage)(nil)).Elem()
}

func (o UserOwnedStorageArrayOutput) ToUserOwnedStorageArrayOutput() UserOwnedStorageArrayOutput {
	return o
}

func (o UserOwnedStorageArrayOutput) ToUserOwnedStorageArrayOutputWithContext(ctx context.Context) UserOwnedStorageArrayOutput {
	return o
}

func (o UserOwnedStorageArrayOutput) Index(i pulumi.IntInput) UserOwnedStorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserOwnedStorage {
		return vs[0].([]UserOwnedStorage)[vs[1].(int)]
	}).(UserOwnedStorageOutput)
}

type UserOwnedStorageResponse struct {
	ResourceId *string `pulumi:"resourceId"`
}





type UserOwnedStorageResponseInput interface {
	pulumi.Input

	ToUserOwnedStorageResponseOutput() UserOwnedStorageResponseOutput
	ToUserOwnedStorageResponseOutputWithContext(context.Context) UserOwnedStorageResponseOutput
}

type UserOwnedStorageResponseArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (UserOwnedStorageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserOwnedStorageResponse)(nil)).Elem()
}

func (i UserOwnedStorageResponseArgs) ToUserOwnedStorageResponseOutput() UserOwnedStorageResponseOutput {
	return i.ToUserOwnedStorageResponseOutputWithContext(context.Background())
}

func (i UserOwnedStorageResponseArgs) ToUserOwnedStorageResponseOutputWithContext(ctx context.Context) UserOwnedStorageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOwnedStorageResponseOutput)
}





type UserOwnedStorageResponseArrayInput interface {
	pulumi.Input

	ToUserOwnedStorageResponseArrayOutput() UserOwnedStorageResponseArrayOutput
	ToUserOwnedStorageResponseArrayOutputWithContext(context.Context) UserOwnedStorageResponseArrayOutput
}

type UserOwnedStorageResponseArray []UserOwnedStorageResponseInput

func (UserOwnedStorageResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserOwnedStorageResponse)(nil)).Elem()
}

func (i UserOwnedStorageResponseArray) ToUserOwnedStorageResponseArrayOutput() UserOwnedStorageResponseArrayOutput {
	return i.ToUserOwnedStorageResponseArrayOutputWithContext(context.Background())
}

func (i UserOwnedStorageResponseArray) ToUserOwnedStorageResponseArrayOutputWithContext(ctx context.Context) UserOwnedStorageResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOwnedStorageResponseArrayOutput)
}

type UserOwnedStorageResponseOutput struct{ *pulumi.OutputState }

func (UserOwnedStorageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserOwnedStorageResponse)(nil)).Elem()
}

func (o UserOwnedStorageResponseOutput) ToUserOwnedStorageResponseOutput() UserOwnedStorageResponseOutput {
	return o
}

func (o UserOwnedStorageResponseOutput) ToUserOwnedStorageResponseOutputWithContext(ctx context.Context) UserOwnedStorageResponseOutput {
	return o
}

func (o UserOwnedStorageResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserOwnedStorageResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type UserOwnedStorageResponseArrayOutput struct{ *pulumi.OutputState }

func (UserOwnedStorageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserOwnedStorageResponse)(nil)).Elem()
}

func (o UserOwnedStorageResponseArrayOutput) ToUserOwnedStorageResponseArrayOutput() UserOwnedStorageResponseArrayOutput {
	return o
}

func (o UserOwnedStorageResponseArrayOutput) ToUserOwnedStorageResponseArrayOutputWithContext(ctx context.Context) UserOwnedStorageResponseArrayOutput {
	return o
}

func (o UserOwnedStorageResponseArrayOutput) Index(i pulumi.IntInput) UserOwnedStorageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserOwnedStorageResponse {
		return vs[0].([]UserOwnedStorageResponse)[vs[1].(int)]
	}).(UserOwnedStorageResponseOutput)
}

type VirtualNetworkRule struct {
	Id                               string  `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint *bool   `pulumi:"ignoreMissingVnetServiceEndpoint"`
	State                            *string `pulumi:"state"`
}





type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(context.Context) VirtualNetworkRuleOutput
}

type VirtualNetworkRuleArgs struct {
	Id                               pulumi.StringInput    `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput   `pulumi:"ignoreMissingVnetServiceEndpoint"`
	State                            pulumi.StringPtrInput `pulumi:"state"`
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}





type VirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput
	ToVirtualNetworkRuleArrayOutputWithContext(context.Context) VirtualNetworkRuleArrayOutput
}

type VirtualNetworkRuleArray []VirtualNetworkRuleInput

func (VirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return i.ToVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleArrayOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRule) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkRuleOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *string { return v.State }).(pulumi.StringPtrOutput)
}

type VirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].([]VirtualNetworkRule)[vs[1].(int)]
	}).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleResponse struct {
	Id                               string  `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint *bool   `pulumi:"ignoreMissingVnetServiceEndpoint"`
	State                            *string `pulumi:"state"`
}





type VirtualNetworkRuleResponseInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput
	ToVirtualNetworkRuleResponseOutputWithContext(context.Context) VirtualNetworkRuleResponseOutput
}

type VirtualNetworkRuleResponseArgs struct {
	Id                               pulumi.StringInput    `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput   `pulumi:"ignoreMissingVnetServiceEndpoint"`
	State                            pulumi.StringPtrInput `pulumi:"state"`
}

func (VirtualNetworkRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return i.ToVirtualNetworkRuleResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseOutput)
}





type VirtualNetworkRuleResponseArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput
	ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Context) VirtualNetworkRuleResponseArrayOutput
}

type VirtualNetworkRuleResponseArray []VirtualNetworkRuleResponseInput

func (VirtualNetworkRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return i.ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseArrayOutput)
}

type VirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleResponseOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type VirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRuleResponse {
		return vs[0].([]VirtualNetworkRuleResponse)[vs[1].(int)]
	}).(VirtualNetworkRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(CognitiveServicesAccountApiPropertiesOutput{})
	pulumi.RegisterOutputType(CognitiveServicesAccountApiPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CognitiveServicesAccountApiPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CognitiveServicesAccountApiPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CognitiveServicesAccountPropertiesOutput{})
	pulumi.RegisterOutputType(CognitiveServicesAccountPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CognitiveServicesAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CognitiveServicesAccountPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CognitiveServicesAccountSkuChangeInfoResponseOutput{})
	pulumi.RegisterOutputType(CognitiveServicesAccountSkuChangeInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(IpRuleOutput{})
	pulumi.RegisterOutputType(IpRuleArrayOutput{})
	pulumi.RegisterOutputType(IpRuleResponseOutput{})
	pulumi.RegisterOutputType(IpRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuCapabilityResponseOutput{})
	pulumi.RegisterOutputType(SkuCapabilityResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityMapOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(UserOwnedStorageOutput{})
	pulumi.RegisterOutputType(UserOwnedStorageArrayOutput{})
	pulumi.RegisterOutputType(UserOwnedStorageResponseOutput{})
	pulumi.RegisterOutputType(UserOwnedStorageResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
