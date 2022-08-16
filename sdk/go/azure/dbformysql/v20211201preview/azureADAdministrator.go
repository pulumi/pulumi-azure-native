


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureADAdministrator struct {
	pulumi.CustomResourceState

	AdministratorType  pulumi.StringPtrOutput   `pulumi:"administratorType"`
	IdentityResourceId pulumi.StringPtrOutput   `pulumi:"identityResourceId"`
	Login              pulumi.StringPtrOutput   `pulumi:"login"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	Sid                pulumi.StringPtrOutput   `pulumi:"sid"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	TenantId           pulumi.StringPtrOutput   `pulumi:"tenantId"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewAzureADAdministrator(ctx *pulumi.Context,
	name string, args *AzureADAdministratorArgs, opts ...pulumi.ResourceOption) (*AzureADAdministrator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	var resource AzureADAdministrator
	err := ctx.RegisterResource("azure-native:dbformysql/v20211201preview:AzureADAdministrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAzureADAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureADAdministratorState, opts ...pulumi.ResourceOption) (*AzureADAdministrator, error) {
	var resource AzureADAdministrator
	err := ctx.ReadResource("azure-native:dbformysql/v20211201preview:AzureADAdministrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type azureADAdministratorState struct {
}

type AzureADAdministratorState struct {
}

func (AzureADAdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureADAdministratorState)(nil)).Elem()
}

type azureADAdministratorArgs struct {
	AdministratorName  *string `pulumi:"administratorName"`
	AdministratorType  *string `pulumi:"administratorType"`
	IdentityResourceId *string `pulumi:"identityResourceId"`
	Login              *string `pulumi:"login"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	ServerName         string  `pulumi:"serverName"`
	Sid                *string `pulumi:"sid"`
	TenantId           *string `pulumi:"tenantId"`
}


type AzureADAdministratorArgs struct {
	AdministratorName  pulumi.StringPtrInput
	AdministratorType  pulumi.StringPtrInput
	IdentityResourceId pulumi.StringPtrInput
	Login              pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	ServerName         pulumi.StringInput
	Sid                pulumi.StringPtrInput
	TenantId           pulumi.StringPtrInput
}

func (AzureADAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureADAdministratorArgs)(nil)).Elem()
}

type AzureADAdministratorInput interface {
	pulumi.Input

	ToAzureADAdministratorOutput() AzureADAdministratorOutput
	ToAzureADAdministratorOutputWithContext(ctx context.Context) AzureADAdministratorOutput
}

func (*AzureADAdministrator) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureADAdministrator)(nil)).Elem()
}

func (i *AzureADAdministrator) ToAzureADAdministratorOutput() AzureADAdministratorOutput {
	return i.ToAzureADAdministratorOutputWithContext(context.Background())
}

func (i *AzureADAdministrator) ToAzureADAdministratorOutputWithContext(ctx context.Context) AzureADAdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureADAdministratorOutput)
}

type AzureADAdministratorOutput struct{ *pulumi.OutputState }

func (AzureADAdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureADAdministrator)(nil)).Elem()
}

func (o AzureADAdministratorOutput) ToAzureADAdministratorOutput() AzureADAdministratorOutput {
	return o
}

func (o AzureADAdministratorOutput) ToAzureADAdministratorOutputWithContext(ctx context.Context) AzureADAdministratorOutput {
	return o
}

func (o AzureADAdministratorOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringPtrOutput { return v.AdministratorType }).(pulumi.StringPtrOutput)
}

func (o AzureADAdministratorOutput) IdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringPtrOutput { return v.IdentityResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureADAdministratorOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringPtrOutput { return v.Login }).(pulumi.StringPtrOutput)
}

func (o AzureADAdministratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AzureADAdministratorOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringPtrOutput { return v.Sid }).(pulumi.StringPtrOutput)
}

func (o AzureADAdministratorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AzureADAdministrator) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AzureADAdministratorOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o AzureADAdministratorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureADAdministrator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureADAdministratorOutput{})
}
