


package solutions

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Application struct {
	pulumi.CustomResourceState

	ApplicationDefinitionId pulumi.StringPtrOutput                            `pulumi:"applicationDefinitionId"`
	Artifacts               ApplicationArtifactResponseArrayOutput            `pulumi:"artifacts"`
	Authorizations          ApplicationAuthorizationResponseArrayOutput       `pulumi:"authorizations"`
	BillingDetails          ApplicationBillingDetailsDefinitionResponseOutput `pulumi:"billingDetails"`
	CreatedBy               ApplicationClientDetailsResponseOutput            `pulumi:"createdBy"`
	CustomerSupport         ApplicationPackageContactResponseOutput           `pulumi:"customerSupport"`
	Identity                IdentityResponsePtrOutput                         `pulumi:"identity"`
	JitAccessPolicy         ApplicationJitAccessPolicyResponsePtrOutput       `pulumi:"jitAccessPolicy"`
	Kind                    pulumi.StringOutput                               `pulumi:"kind"`
	Location                pulumi.StringPtrOutput                            `pulumi:"location"`
	ManagedBy               pulumi.StringPtrOutput                            `pulumi:"managedBy"`
	ManagedResourceGroupId  pulumi.StringPtrOutput                            `pulumi:"managedResourceGroupId"`
	ManagementMode          pulumi.StringOutput                               `pulumi:"managementMode"`
	Name                    pulumi.StringOutput                               `pulumi:"name"`
	Outputs                 pulumi.AnyOutput                                  `pulumi:"outputs"`
	Parameters              pulumi.AnyOutput                                  `pulumi:"parameters"`
	Plan                    PlanResponsePtrOutput                             `pulumi:"plan"`
	ProvisioningState       pulumi.StringOutput                               `pulumi:"provisioningState"`
	PublisherTenantId       pulumi.StringOutput                               `pulumi:"publisherTenantId"`
	Sku                     SkuResponsePtrOutput                              `pulumi:"sku"`
	SupportUrls             ApplicationPackageSupportUrlsResponseOutput       `pulumi:"supportUrls"`
	Tags                    pulumi.StringMapOutput                            `pulumi:"tags"`
	Type                    pulumi.StringOutput                               `pulumi:"type"`
	UpdatedBy               ApplicationClientDetailsResponseOutput            `pulumi:"updatedBy"`
}


func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:solutions/v20160901preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20170901:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20171201:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180201:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180301:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180601:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180901preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20190701:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20200821preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20210201preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20210701:Application"),
		},
	})
	opts = append(opts, aliases)
	var resource Application
	err := ctx.RegisterResource("azure-native:solutions:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure-native:solutions:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationState struct {
}

type ApplicationState struct {
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	ApplicationDefinitionId *string                     `pulumi:"applicationDefinitionId"`
	ApplicationName         *string                     `pulumi:"applicationName"`
	Identity                *Identity                   `pulumi:"identity"`
	JitAccessPolicy         *ApplicationJitAccessPolicy `pulumi:"jitAccessPolicy"`
	Kind                    string                      `pulumi:"kind"`
	Location                *string                     `pulumi:"location"`
	ManagedBy               *string                     `pulumi:"managedBy"`
	ManagedResourceGroupId  *string                     `pulumi:"managedResourceGroupId"`
	Parameters              interface{}                 `pulumi:"parameters"`
	Plan                    *Plan                       `pulumi:"plan"`
	ResourceGroupName       string                      `pulumi:"resourceGroupName"`
	Sku                     *Sku                        `pulumi:"sku"`
	Tags                    map[string]string           `pulumi:"tags"`
}


type ApplicationArgs struct {
	ApplicationDefinitionId pulumi.StringPtrInput
	ApplicationName         pulumi.StringPtrInput
	Identity                IdentityPtrInput
	JitAccessPolicy         ApplicationJitAccessPolicyPtrInput
	Kind                    pulumi.StringInput
	Location                pulumi.StringPtrInput
	ManagedBy               pulumi.StringPtrInput
	ManagedResourceGroupId  pulumi.StringPtrInput
	Parameters              pulumi.Input
	Plan                    PlanPtrInput
	ResourceGroupName       pulumi.StringInput
	Sku                     SkuPtrInput
	Tags                    pulumi.StringMapInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func (o ApplicationOutput) ApplicationDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.ApplicationDefinitionId }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) Artifacts() ApplicationArtifactResponseArrayOutput {
	return o.ApplyT(func(v *Application) ApplicationArtifactResponseArrayOutput { return v.Artifacts }).(ApplicationArtifactResponseArrayOutput)
}

func (o ApplicationOutput) Authorizations() ApplicationAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v *Application) ApplicationAuthorizationResponseArrayOutput { return v.Authorizations }).(ApplicationAuthorizationResponseArrayOutput)
}

func (o ApplicationOutput) BillingDetails() ApplicationBillingDetailsDefinitionResponseOutput {
	return o.ApplyT(func(v *Application) ApplicationBillingDetailsDefinitionResponseOutput { return v.BillingDetails }).(ApplicationBillingDetailsDefinitionResponseOutput)
}

func (o ApplicationOutput) CreatedBy() ApplicationClientDetailsResponseOutput {
	return o.ApplyT(func(v *Application) ApplicationClientDetailsResponseOutput { return v.CreatedBy }).(ApplicationClientDetailsResponseOutput)
}

func (o ApplicationOutput) CustomerSupport() ApplicationPackageContactResponseOutput {
	return o.ApplyT(func(v *Application) ApplicationPackageContactResponseOutput { return v.CustomerSupport }).(ApplicationPackageContactResponseOutput)
}

func (o ApplicationOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Application) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o ApplicationOutput) JitAccessPolicy() ApplicationJitAccessPolicyResponsePtrOutput {
	return o.ApplyT(func(v *Application) ApplicationJitAccessPolicyResponsePtrOutput { return v.JitAccessPolicy }).(ApplicationJitAccessPolicyResponsePtrOutput)
}

func (o ApplicationOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ApplicationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) ManagedResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.ManagedResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) ManagementMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ManagementMode }).(pulumi.StringOutput)
}

func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationOutput) Outputs() pulumi.AnyOutput {
	return o.ApplyT(func(v *Application) pulumi.AnyOutput { return v.Outputs }).(pulumi.AnyOutput)
}

func (o ApplicationOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *Application) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

func (o ApplicationOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v *Application) PlanResponsePtrOutput { return v.Plan }).(PlanResponsePtrOutput)
}

func (o ApplicationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ApplicationOutput) PublisherTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.PublisherTenantId }).(pulumi.StringOutput)
}

func (o ApplicationOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Application) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ApplicationOutput) SupportUrls() ApplicationPackageSupportUrlsResponseOutput {
	return o.ApplyT(func(v *Application) ApplicationPackageSupportUrlsResponseOutput { return v.SupportUrls }).(ApplicationPackageSupportUrlsResponseOutput)
}

func (o ApplicationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Application) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ApplicationOutput) UpdatedBy() ApplicationClientDetailsResponseOutput {
	return o.ApplyT(func(v *Application) ApplicationClientDetailsResponseOutput { return v.UpdatedBy }).(ApplicationClientDetailsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
}
