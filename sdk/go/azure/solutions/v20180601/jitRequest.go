


package v20180601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type JitRequest struct {
	pulumi.CustomResourceState

	ApplicationResourceId    pulumi.StringOutput                         `pulumi:"applicationResourceId"`
	CreatedBy                ApplicationClientDetailsResponseOutput      `pulumi:"createdBy"`
	JitAuthorizationPolicies JitAuthorizationPoliciesResponseArrayOutput `pulumi:"jitAuthorizationPolicies"`
	JitRequestState          pulumi.StringOutput                         `pulumi:"jitRequestState"`
	JitSchedulingPolicy      JitSchedulingPolicyResponseOutput           `pulumi:"jitSchedulingPolicy"`
	Location                 pulumi.StringPtrOutput                      `pulumi:"location"`
	Name                     pulumi.StringOutput                         `pulumi:"name"`
	ProvisioningState        pulumi.StringOutput                         `pulumi:"provisioningState"`
	PublisherTenantId        pulumi.StringOutput                         `pulumi:"publisherTenantId"`
	Tags                     pulumi.StringMapOutput                      `pulumi:"tags"`
	Type                     pulumi.StringOutput                         `pulumi:"type"`
	UpdatedBy                ApplicationClientDetailsResponseOutput      `pulumi:"updatedBy"`
}


func NewJitRequest(ctx *pulumi.Context,
	name string, args *JitRequestArgs, opts ...pulumi.ResourceOption) (*JitRequest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationResourceId'")
	}
	if args.JitAuthorizationPolicies == nil {
		return nil, errors.New("invalid value for required argument 'JitAuthorizationPolicies'")
	}
	if args.JitSchedulingPolicy == nil {
		return nil, errors.New("invalid value for required argument 'JitSchedulingPolicy'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:solutions:JitRequest"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180301:JitRequest"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180901preview:JitRequest"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20190701:JitRequest"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20200821preview:JitRequest"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20210201preview:JitRequest"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20210701:JitRequest"),
		},
	})
	opts = append(opts, aliases)
	var resource JitRequest
	err := ctx.RegisterResource("azure-native:solutions/v20180601:JitRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJitRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JitRequestState, opts ...pulumi.ResourceOption) (*JitRequest, error) {
	var resource JitRequest
	err := ctx.ReadResource("azure-native:solutions/v20180601:JitRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type jitRequestState struct {
}

type JitRequestState struct {
}

func (JitRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*jitRequestState)(nil)).Elem()
}

type jitRequestArgs struct {
	ApplicationResourceId    string                     `pulumi:"applicationResourceId"`
	JitAuthorizationPolicies []JitAuthorizationPolicies `pulumi:"jitAuthorizationPolicies"`
	JitRequestName           *string                    `pulumi:"jitRequestName"`
	JitSchedulingPolicy      JitSchedulingPolicy        `pulumi:"jitSchedulingPolicy"`
	Location                 *string                    `pulumi:"location"`
	ResourceGroupName        string                     `pulumi:"resourceGroupName"`
	Tags                     map[string]string          `pulumi:"tags"`
}


type JitRequestArgs struct {
	ApplicationResourceId    pulumi.StringInput
	JitAuthorizationPolicies JitAuthorizationPoliciesArrayInput
	JitRequestName           pulumi.StringPtrInput
	JitSchedulingPolicy      JitSchedulingPolicyInput
	Location                 pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	Tags                     pulumi.StringMapInput
}

func (JitRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jitRequestArgs)(nil)).Elem()
}

type JitRequestInput interface {
	pulumi.Input

	ToJitRequestOutput() JitRequestOutput
	ToJitRequestOutputWithContext(ctx context.Context) JitRequestOutput
}

func (*JitRequest) ElementType() reflect.Type {
	return reflect.TypeOf((**JitRequest)(nil)).Elem()
}

func (i *JitRequest) ToJitRequestOutput() JitRequestOutput {
	return i.ToJitRequestOutputWithContext(context.Background())
}

func (i *JitRequest) ToJitRequestOutputWithContext(ctx context.Context) JitRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitRequestOutput)
}

type JitRequestOutput struct{ *pulumi.OutputState }

func (JitRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JitRequest)(nil)).Elem()
}

func (o JitRequestOutput) ToJitRequestOutput() JitRequestOutput {
	return o
}

func (o JitRequestOutput) ToJitRequestOutputWithContext(ctx context.Context) JitRequestOutput {
	return o
}

func (o JitRequestOutput) ApplicationResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *JitRequest) pulumi.StringOutput { return v.ApplicationResourceId }).(pulumi.StringOutput)
}

func (o JitRequestOutput) CreatedBy() ApplicationClientDetailsResponseOutput {
	return o.ApplyT(func(v *JitRequest) ApplicationClientDetailsResponseOutput { return v.CreatedBy }).(ApplicationClientDetailsResponseOutput)
}

func (o JitRequestOutput) JitAuthorizationPolicies() JitAuthorizationPoliciesResponseArrayOutput {
	return o.ApplyT(func(v *JitRequest) JitAuthorizationPoliciesResponseArrayOutput { return v.JitAuthorizationPolicies }).(JitAuthorizationPoliciesResponseArrayOutput)
}

func (o JitRequestOutput) JitRequestState() pulumi.StringOutput {
	return o.ApplyT(func(v *JitRequest) pulumi.StringOutput { return v.JitRequestState }).(pulumi.StringOutput)
}

func (o JitRequestOutput) JitSchedulingPolicy() JitSchedulingPolicyResponseOutput {
	return o.ApplyT(func(v *JitRequest) JitSchedulingPolicyResponseOutput { return v.JitSchedulingPolicy }).(JitSchedulingPolicyResponseOutput)
}

func (o JitRequestOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JitRequest) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o JitRequestOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JitRequest) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o JitRequestOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *JitRequest) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o JitRequestOutput) PublisherTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *JitRequest) pulumi.StringOutput { return v.PublisherTenantId }).(pulumi.StringOutput)
}

func (o JitRequestOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *JitRequest) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o JitRequestOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JitRequest) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o JitRequestOutput) UpdatedBy() ApplicationClientDetailsResponseOutput {
	return o.ApplyT(func(v *JitRequest) ApplicationClientDetailsResponseOutput { return v.UpdatedBy }).(ApplicationClientDetailsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(JitRequestOutput{})
}
