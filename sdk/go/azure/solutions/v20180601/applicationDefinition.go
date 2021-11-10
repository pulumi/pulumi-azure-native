


package v20180601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationDefinition struct {
	pulumi.CustomResourceState

	Artifacts          ApplicationArtifactResponseArrayOutput              `pulumi:"artifacts"`
	Authorizations     ApplicationProviderAuthorizationResponseArrayOutput `pulumi:"authorizations"`
	CreateUiDefinition pulumi.AnyOutput                                    `pulumi:"createUiDefinition"`
	Description        pulumi.StringPtrOutput                              `pulumi:"description"`
	DisplayName        pulumi.StringPtrOutput                              `pulumi:"displayName"`
	Identity           IdentityResponsePtrOutput                           `pulumi:"identity"`
	IsEnabled          pulumi.StringPtrOutput                              `pulumi:"isEnabled"`
	Location           pulumi.StringPtrOutput                              `pulumi:"location"`
	LockLevel          pulumi.StringOutput                                 `pulumi:"lockLevel"`
	MainTemplate       pulumi.AnyOutput                                    `pulumi:"mainTemplate"`
	ManagedBy          pulumi.StringPtrOutput                              `pulumi:"managedBy"`
	Name               pulumi.StringOutput                                 `pulumi:"name"`
	PackageFileUri     pulumi.StringPtrOutput                              `pulumi:"packageFileUri"`
	Sku                SkuResponsePtrOutput                                `pulumi:"sku"`
	Tags               pulumi.StringMapOutput                              `pulumi:"tags"`
	Type               pulumi.StringOutput                                 `pulumi:"type"`
}


func NewApplicationDefinition(ctx *pulumi.Context,
	name string, args *ApplicationDefinitionArgs, opts ...pulumi.ResourceOption) (*ApplicationDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authorizations == nil {
		return nil, errors.New("invalid value for required argument 'Authorizations'")
	}
	if args.LockLevel == nil {
		return nil, errors.New("invalid value for required argument 'LockLevel'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:solutions/v20180601:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:solutions:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20160901preview:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:solutions/v20160901preview:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20170901:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:solutions/v20170901:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20190701:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:solutions/v20190701:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20200821preview:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:solutions/v20200821preview:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20210701:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:solutions/v20210701:ApplicationDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationDefinition
	err := ctx.RegisterResource("azure-native:solutions/v20180601:ApplicationDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationDefinitionState, opts ...pulumi.ResourceOption) (*ApplicationDefinition, error) {
	var resource ApplicationDefinition
	err := ctx.ReadResource("azure-native:solutions/v20180601:ApplicationDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationDefinitionState struct {
}

type ApplicationDefinitionState struct {
}

func (ApplicationDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationDefinitionState)(nil)).Elem()
}

type applicationDefinitionArgs struct {
	ApplicationDefinitionName *string                            `pulumi:"applicationDefinitionName"`
	Artifacts                 []ApplicationArtifact              `pulumi:"artifacts"`
	Authorizations            []ApplicationProviderAuthorization `pulumi:"authorizations"`
	CreateUiDefinition        interface{}                        `pulumi:"createUiDefinition"`
	Description               *string                            `pulumi:"description"`
	DisplayName               *string                            `pulumi:"displayName"`
	Identity                  *Identity                          `pulumi:"identity"`
	IsEnabled                 *string                            `pulumi:"isEnabled"`
	Location                  *string                            `pulumi:"location"`
	LockLevel                 ApplicationLockLevel               `pulumi:"lockLevel"`
	MainTemplate              interface{}                        `pulumi:"mainTemplate"`
	ManagedBy                 *string                            `pulumi:"managedBy"`
	PackageFileUri            *string                            `pulumi:"packageFileUri"`
	ResourceGroupName         string                             `pulumi:"resourceGroupName"`
	Sku                       *Sku                               `pulumi:"sku"`
	Tags                      map[string]string                  `pulumi:"tags"`
}


type ApplicationDefinitionArgs struct {
	ApplicationDefinitionName pulumi.StringPtrInput
	Artifacts                 ApplicationArtifactArrayInput
	Authorizations            ApplicationProviderAuthorizationArrayInput
	CreateUiDefinition        pulumi.Input
	Description               pulumi.StringPtrInput
	DisplayName               pulumi.StringPtrInput
	Identity                  IdentityPtrInput
	IsEnabled                 pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	LockLevel                 ApplicationLockLevelInput
	MainTemplate              pulumi.Input
	ManagedBy                 pulumi.StringPtrInput
	PackageFileUri            pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Sku                       SkuPtrInput
	Tags                      pulumi.StringMapInput
}

func (ApplicationDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationDefinitionArgs)(nil)).Elem()
}

type ApplicationDefinitionInput interface {
	pulumi.Input

	ToApplicationDefinitionOutput() ApplicationDefinitionOutput
	ToApplicationDefinitionOutputWithContext(ctx context.Context) ApplicationDefinitionOutput
}

func (*ApplicationDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDefinition)(nil))
}

func (i *ApplicationDefinition) ToApplicationDefinitionOutput() ApplicationDefinitionOutput {
	return i.ToApplicationDefinitionOutputWithContext(context.Background())
}

func (i *ApplicationDefinition) ToApplicationDefinitionOutputWithContext(ctx context.Context) ApplicationDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDefinitionOutput)
}

type ApplicationDefinitionOutput struct{ *pulumi.OutputState }

func (ApplicationDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDefinition)(nil))
}

func (o ApplicationDefinitionOutput) ToApplicationDefinitionOutput() ApplicationDefinitionOutput {
	return o
}

func (o ApplicationDefinitionOutput) ToApplicationDefinitionOutputWithContext(ctx context.Context) ApplicationDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplicationDefinitionOutput{})
}
