


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Service struct {
	pulumi.CustomResourceState

	CreatedAt          pulumi.StringPtrOutput                  `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrOutput                  `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrOutput                  `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrOutput                  `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrOutput                  `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrOutput                  `pulumi:"lastModifiedByType"`
	Location           pulumi.StringOutput                     `pulumi:"location"`
	Name               pulumi.StringOutput                     `pulumi:"name"`
	PccRules           PccRuleConfigurationResponseArrayOutput `pulumi:"pccRules"`
	ProvisioningState  pulumi.StringOutput                     `pulumi:"provisioningState"`
	ServicePrecedence  pulumi.IntOutput                        `pulumi:"servicePrecedence"`
	ServiceQosPolicy   QosPolicyResponsePtrOutput              `pulumi:"serviceQosPolicy"`
	SystemData         SystemDataResponseOutput                `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput                  `pulumi:"tags"`
	Type               pulumi.StringOutput                     `pulumi:"type"`
}


func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MobileNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'MobileNetworkName'")
	}
	if args.PccRules == nil {
		return nil, errors.New("invalid value for required argument 'PccRules'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServicePrecedence == nil {
		return nil, errors.New("invalid value for required argument 'ServicePrecedence'")
	}
	if args.ServiceQosPolicy != nil {
		args.ServiceQosPolicy = args.ServiceQosPolicy.ToQosPolicyPtrOutput().ApplyT(func(v *QosPolicy) *QosPolicy { return v.Defaults() }).(QosPolicyPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:Service"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20221101:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20220401preview:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:mobilenetwork/v20220401preview:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	CreatedAt          *string                `pulumi:"createdAt"`
	CreatedBy          *string                `pulumi:"createdBy"`
	CreatedByType      *string                `pulumi:"createdByType"`
	LastModifiedAt     *string                `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string                `pulumi:"lastModifiedBy"`
	LastModifiedByType *string                `pulumi:"lastModifiedByType"`
	Location           *string                `pulumi:"location"`
	MobileNetworkName  string                 `pulumi:"mobileNetworkName"`
	PccRules           []PccRuleConfiguration `pulumi:"pccRules"`
	ResourceGroupName  string                 `pulumi:"resourceGroupName"`
	ServiceName        *string                `pulumi:"serviceName"`
	ServicePrecedence  int                    `pulumi:"servicePrecedence"`
	ServiceQosPolicy   *QosPolicy             `pulumi:"serviceQosPolicy"`
	Tags               map[string]string      `pulumi:"tags"`
}


type ServiceArgs struct {
	CreatedAt          pulumi.StringPtrInput
	CreatedBy          pulumi.StringPtrInput
	CreatedByType      pulumi.StringPtrInput
	LastModifiedAt     pulumi.StringPtrInput
	LastModifiedBy     pulumi.StringPtrInput
	LastModifiedByType pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	MobileNetworkName  pulumi.StringInput
	PccRules           PccRuleConfigurationArrayInput
	ResourceGroupName  pulumi.StringInput
	ServiceName        pulumi.StringPtrInput
	ServicePrecedence  pulumi.IntInput
	ServiceQosPolicy   QosPolicyPtrInput
	Tags               pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func (o ServiceOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceOutput) PccRules() PccRuleConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *Service) PccRuleConfigurationResponseArrayOutput { return v.PccRules }).(PccRuleConfigurationResponseArrayOutput)
}

func (o ServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceOutput) ServicePrecedence() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.ServicePrecedence }).(pulumi.IntOutput)
}

func (o ServiceOutput) ServiceQosPolicy() QosPolicyResponsePtrOutput {
	return o.ApplyT(func(v *Service) QosPolicyResponsePtrOutput { return v.ServiceQosPolicy }).(QosPolicyResponsePtrOutput)
}

func (o ServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Service) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
