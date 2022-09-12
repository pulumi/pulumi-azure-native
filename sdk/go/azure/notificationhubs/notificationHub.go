


package notificationhubs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NotificationHub struct {
	pulumi.CustomResourceState

	AdmCredential      AdmCredentialResponsePtrOutput                             `pulumi:"admCredential"`
	ApnsCredential     ApnsCredentialResponsePtrOutput                            `pulumi:"apnsCredential"`
	AuthorizationRules SharedAccessAuthorizationRulePropertiesResponseArrayOutput `pulumi:"authorizationRules"`
	BaiduCredential    BaiduCredentialResponsePtrOutput                           `pulumi:"baiduCredential"`
	GcmCredential      GcmCredentialResponsePtrOutput                             `pulumi:"gcmCredential"`
	Location           pulumi.StringPtrOutput                                     `pulumi:"location"`
	MpnsCredential     MpnsCredentialResponsePtrOutput                            `pulumi:"mpnsCredential"`
	Name               pulumi.StringOutput                                        `pulumi:"name"`
	RegistrationTtl    pulumi.StringPtrOutput                                     `pulumi:"registrationTtl"`
	Sku                SkuResponsePtrOutput                                       `pulumi:"sku"`
	Tags               pulumi.StringMapOutput                                     `pulumi:"tags"`
	Type               pulumi.StringOutput                                        `pulumi:"type"`
	WnsCredential      WnsCredentialResponsePtrOutput                             `pulumi:"wnsCredential"`
}


func NewNotificationHub(ctx *pulumi.Context,
	name string, args *NotificationHubArgs, opts ...pulumi.ResourceOption) (*NotificationHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:notificationhubs/v20140901:NotificationHub"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20160301:NotificationHub"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20170401:NotificationHub"),
		},
	})
	opts = append(opts, aliases)
	var resource NotificationHub
	err := ctx.RegisterResource("azure-native:notificationhubs:NotificationHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNotificationHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationHubState, opts ...pulumi.ResourceOption) (*NotificationHub, error) {
	var resource NotificationHub
	err := ctx.ReadResource("azure-native:notificationhubs:NotificationHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type notificationHubState struct {
}

type NotificationHubState struct {
}

func (NotificationHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubState)(nil)).Elem()
}

type notificationHubArgs struct {
	AdmCredential       *AdmCredential                            `pulumi:"admCredential"`
	ApnsCredential      *ApnsCredential                           `pulumi:"apnsCredential"`
	AuthorizationRules  []SharedAccessAuthorizationRuleProperties `pulumi:"authorizationRules"`
	BaiduCredential     *BaiduCredential                          `pulumi:"baiduCredential"`
	GcmCredential       *GcmCredential                            `pulumi:"gcmCredential"`
	Location            *string                                   `pulumi:"location"`
	MpnsCredential      *MpnsCredential                           `pulumi:"mpnsCredential"`
	Name                *string                                   `pulumi:"name"`
	NamespaceName       string                                    `pulumi:"namespaceName"`
	NotificationHubName *string                                   `pulumi:"notificationHubName"`
	RegistrationTtl     *string                                   `pulumi:"registrationTtl"`
	ResourceGroupName   string                                    `pulumi:"resourceGroupName"`
	Sku                 *Sku                                      `pulumi:"sku"`
	Tags                map[string]string                         `pulumi:"tags"`
	WnsCredential       *WnsCredential                            `pulumi:"wnsCredential"`
}


type NotificationHubArgs struct {
	AdmCredential       AdmCredentialPtrInput
	ApnsCredential      ApnsCredentialPtrInput
	AuthorizationRules  SharedAccessAuthorizationRulePropertiesArrayInput
	BaiduCredential     BaiduCredentialPtrInput
	GcmCredential       GcmCredentialPtrInput
	Location            pulumi.StringPtrInput
	MpnsCredential      MpnsCredentialPtrInput
	Name                pulumi.StringPtrInput
	NamespaceName       pulumi.StringInput
	NotificationHubName pulumi.StringPtrInput
	RegistrationTtl     pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Sku                 SkuPtrInput
	Tags                pulumi.StringMapInput
	WnsCredential       WnsCredentialPtrInput
}

func (NotificationHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubArgs)(nil)).Elem()
}

type NotificationHubInput interface {
	pulumi.Input

	ToNotificationHubOutput() NotificationHubOutput
	ToNotificationHubOutputWithContext(ctx context.Context) NotificationHubOutput
}

func (*NotificationHub) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationHub)(nil)).Elem()
}

func (i *NotificationHub) ToNotificationHubOutput() NotificationHubOutput {
	return i.ToNotificationHubOutputWithContext(context.Background())
}

func (i *NotificationHub) ToNotificationHubOutputWithContext(ctx context.Context) NotificationHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationHubOutput)
}

type NotificationHubOutput struct{ *pulumi.OutputState }

func (NotificationHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationHub)(nil)).Elem()
}

func (o NotificationHubOutput) ToNotificationHubOutput() NotificationHubOutput {
	return o
}

func (o NotificationHubOutput) ToNotificationHubOutputWithContext(ctx context.Context) NotificationHubOutput {
	return o
}

func (o NotificationHubOutput) AdmCredential() AdmCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) AdmCredentialResponsePtrOutput { return v.AdmCredential }).(AdmCredentialResponsePtrOutput)
}

func (o NotificationHubOutput) ApnsCredential() ApnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) ApnsCredentialResponsePtrOutput { return v.ApnsCredential }).(ApnsCredentialResponsePtrOutput)
}

func (o NotificationHubOutput) AuthorizationRules() SharedAccessAuthorizationRulePropertiesResponseArrayOutput {
	return o.ApplyT(func(v *NotificationHub) SharedAccessAuthorizationRulePropertiesResponseArrayOutput {
		return v.AuthorizationRules
	}).(SharedAccessAuthorizationRulePropertiesResponseArrayOutput)
}

func (o NotificationHubOutput) BaiduCredential() BaiduCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) BaiduCredentialResponsePtrOutput { return v.BaiduCredential }).(BaiduCredentialResponsePtrOutput)
}

func (o NotificationHubOutput) GcmCredential() GcmCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) GcmCredentialResponsePtrOutput { return v.GcmCredential }).(GcmCredentialResponsePtrOutput)
}

func (o NotificationHubOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationHub) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NotificationHubOutput) MpnsCredential() MpnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) MpnsCredentialResponsePtrOutput { return v.MpnsCredential }).(MpnsCredentialResponsePtrOutput)
}

func (o NotificationHubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NotificationHubOutput) RegistrationTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationHub) pulumi.StringPtrOutput { return v.RegistrationTtl }).(pulumi.StringPtrOutput)
}

func (o NotificationHubOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o NotificationHubOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NotificationHub) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NotificationHubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NotificationHubOutput) WnsCredential() WnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHub) WnsCredentialResponsePtrOutput { return v.WnsCredential }).(WnsCredentialResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NotificationHubOutput{})
}
