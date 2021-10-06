


package v20201019preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MSIXPackage struct {
	pulumi.CustomResourceState

	DisplayName           pulumi.StringPtrOutput                     `pulumi:"displayName"`
	ImagePath             pulumi.StringPtrOutput                     `pulumi:"imagePath"`
	IsActive              pulumi.BoolPtrOutput                       `pulumi:"isActive"`
	IsRegularRegistration pulumi.BoolPtrOutput                       `pulumi:"isRegularRegistration"`
	LastUpdated           pulumi.StringPtrOutput                     `pulumi:"lastUpdated"`
	Name                  pulumi.StringOutput                        `pulumi:"name"`
	PackageApplications   MsixPackageApplicationsResponseArrayOutput `pulumi:"packageApplications"`
	PackageDependencies   MsixPackageDependenciesResponseArrayOutput `pulumi:"packageDependencies"`
	PackageFamilyName     pulumi.StringPtrOutput                     `pulumi:"packageFamilyName"`
	PackageName           pulumi.StringPtrOutput                     `pulumi:"packageName"`
	PackageRelativePath   pulumi.StringPtrOutput                     `pulumi:"packageRelativePath"`
	Type                  pulumi.StringOutput                        `pulumi:"type"`
	Version               pulumi.StringPtrOutput                     `pulumi:"version"`
}


func NewMSIXPackage(ctx *pulumi.Context,
	name string, args *MSIXPackageArgs, opts ...pulumi.ResourceOption) (*MSIXPackage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostPoolName == nil {
		return nil, errors.New("invalid value for required argument 'HostPoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20201019preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20200921preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20200921preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201102preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20201102preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201110preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20201110preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210114preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210114preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210201preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210201preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210309preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210401preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210712:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210712:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:MSIXPackage"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210903preview:MSIXPackage"),
		},
	})
	opts = append(opts, aliases)
	var resource MSIXPackage
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20201019preview:MSIXPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMSIXPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MSIXPackageState, opts ...pulumi.ResourceOption) (*MSIXPackage, error) {
	var resource MSIXPackage
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20201019preview:MSIXPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type msixpackageState struct {
}

type MSIXPackageState struct {
}

func (MSIXPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*msixpackageState)(nil)).Elem()
}

type msixpackageArgs struct {
	DisplayName           *string                   `pulumi:"displayName"`
	HostPoolName          string                    `pulumi:"hostPoolName"`
	ImagePath             *string                   `pulumi:"imagePath"`
	IsActive              *bool                     `pulumi:"isActive"`
	IsRegularRegistration *bool                     `pulumi:"isRegularRegistration"`
	LastUpdated           *string                   `pulumi:"lastUpdated"`
	MsixPackageFullName   *string                   `pulumi:"msixPackageFullName"`
	PackageApplications   []MsixPackageApplications `pulumi:"packageApplications"`
	PackageDependencies   []MsixPackageDependencies `pulumi:"packageDependencies"`
	PackageFamilyName     *string                   `pulumi:"packageFamilyName"`
	PackageName           *string                   `pulumi:"packageName"`
	PackageRelativePath   *string                   `pulumi:"packageRelativePath"`
	ResourceGroupName     string                    `pulumi:"resourceGroupName"`
	Version               *string                   `pulumi:"version"`
}


type MSIXPackageArgs struct {
	DisplayName           pulumi.StringPtrInput
	HostPoolName          pulumi.StringInput
	ImagePath             pulumi.StringPtrInput
	IsActive              pulumi.BoolPtrInput
	IsRegularRegistration pulumi.BoolPtrInput
	LastUpdated           pulumi.StringPtrInput
	MsixPackageFullName   pulumi.StringPtrInput
	PackageApplications   MsixPackageApplicationsArrayInput
	PackageDependencies   MsixPackageDependenciesArrayInput
	PackageFamilyName     pulumi.StringPtrInput
	PackageName           pulumi.StringPtrInput
	PackageRelativePath   pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Version               pulumi.StringPtrInput
}

func (MSIXPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*msixpackageArgs)(nil)).Elem()
}

type MSIXPackageInput interface {
	pulumi.Input

	ToMSIXPackageOutput() MSIXPackageOutput
	ToMSIXPackageOutputWithContext(ctx context.Context) MSIXPackageOutput
}

func (*MSIXPackage) ElementType() reflect.Type {
	return reflect.TypeOf((*MSIXPackage)(nil))
}

func (i *MSIXPackage) ToMSIXPackageOutput() MSIXPackageOutput {
	return i.ToMSIXPackageOutputWithContext(context.Background())
}

func (i *MSIXPackage) ToMSIXPackageOutputWithContext(ctx context.Context) MSIXPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSIXPackageOutput)
}

type MSIXPackageOutput struct{ *pulumi.OutputState }

func (MSIXPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MSIXPackage)(nil))
}

func (o MSIXPackageOutput) ToMSIXPackageOutput() MSIXPackageOutput {
	return o
}

func (o MSIXPackageOutput) ToMSIXPackageOutputWithContext(ctx context.Context) MSIXPackageOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MSIXPackageOutput{})
}
