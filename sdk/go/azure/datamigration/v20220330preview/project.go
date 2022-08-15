


package v20220330preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Project struct {
	pulumi.CustomResourceState

	AzureAuthenticationInfo AzureActiveDirectoryAppResponsePtrOutput `pulumi:"azureAuthenticationInfo"`
	CreationTime            pulumi.StringOutput                      `pulumi:"creationTime"`
	DatabasesInfo           DatabaseInfoResponseArrayOutput          `pulumi:"databasesInfo"`
	Etag                    pulumi.StringPtrOutput                   `pulumi:"etag"`
	Location                pulumi.StringPtrOutput                   `pulumi:"location"`
	Name                    pulumi.StringOutput                      `pulumi:"name"`
	ProvisioningState       pulumi.StringOutput                      `pulumi:"provisioningState"`
	SourceConnectionInfo    pulumi.AnyOutput                         `pulumi:"sourceConnectionInfo"`
	SourcePlatform          pulumi.StringOutput                      `pulumi:"sourcePlatform"`
	SystemData              SystemDataResponseOutput                 `pulumi:"systemData"`
	Tags                    pulumi.StringMapOutput                   `pulumi:"tags"`
	TargetConnectionInfo    pulumi.AnyOutput                         `pulumi:"targetConnectionInfo"`
	TargetPlatform          pulumi.StringOutput                      `pulumi:"targetPlatform"`
	Type                    pulumi.StringOutput                      `pulumi:"type"`
}


func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.SourcePlatform == nil {
		return nil, errors.New("invalid value for required argument 'SourcePlatform'")
	}
	if args.TargetPlatform == nil {
		return nil, errors.New("invalid value for required argument 'TargetPlatform'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datamigration:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20171115preview:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180315preview:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180331preview:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180419:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180715preview:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20210630:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20211030preview:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220130preview:Project"),
		},
	})
	opts = append(opts, aliases)
	var resource Project
	err := ctx.RegisterResource("azure-native:datamigration/v20220330preview:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("azure-native:datamigration/v20220330preview:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type projectState struct {
}

type ProjectState struct {
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	AzureAuthenticationInfo *AzureActiveDirectoryApp `pulumi:"azureAuthenticationInfo"`
	DatabasesInfo           []DatabaseInfo           `pulumi:"databasesInfo"`
	GroupName               string                   `pulumi:"groupName"`
	Location                *string                  `pulumi:"location"`
	ProjectName             *string                  `pulumi:"projectName"`
	ServiceName             string                   `pulumi:"serviceName"`
	SourceConnectionInfo    interface{}              `pulumi:"sourceConnectionInfo"`
	SourcePlatform          string                   `pulumi:"sourcePlatform"`
	Tags                    map[string]string        `pulumi:"tags"`
	TargetConnectionInfo    interface{}              `pulumi:"targetConnectionInfo"`
	TargetPlatform          string                   `pulumi:"targetPlatform"`
}


type ProjectArgs struct {
	AzureAuthenticationInfo AzureActiveDirectoryAppPtrInput
	DatabasesInfo           DatabaseInfoArrayInput
	GroupName               pulumi.StringInput
	Location                pulumi.StringPtrInput
	ProjectName             pulumi.StringPtrInput
	ServiceName             pulumi.StringInput
	SourceConnectionInfo    pulumi.Input
	SourcePlatform          pulumi.StringInput
	Tags                    pulumi.StringMapInput
	TargetConnectionInfo    pulumi.Input
	TargetPlatform          pulumi.StringInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

func (o ProjectOutput) AzureAuthenticationInfo() AzureActiveDirectoryAppResponsePtrOutput {
	return o.ApplyT(func(v *Project) AzureActiveDirectoryAppResponsePtrOutput { return v.AzureAuthenticationInfo }).(AzureActiveDirectoryAppResponsePtrOutput)
}

func (o ProjectOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o ProjectOutput) DatabasesInfo() DatabaseInfoResponseArrayOutput {
	return o.ApplyT(func(v *Project) DatabaseInfoResponseArrayOutput { return v.DatabasesInfo }).(DatabaseInfoResponseArrayOutput)
}

func (o ProjectOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ProjectOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProjectOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ProjectOutput) SourceConnectionInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v *Project) pulumi.AnyOutput { return v.SourceConnectionInfo }).(pulumi.AnyOutput)
}

func (o ProjectOutput) SourcePlatform() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.SourcePlatform }).(pulumi.StringOutput)
}

func (o ProjectOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Project) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ProjectOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Project) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ProjectOutput) TargetConnectionInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v *Project) pulumi.AnyOutput { return v.TargetConnectionInfo }).(pulumi.AnyOutput)
}

func (o ProjectOutput) TargetPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.TargetPlatform }).(pulumi.StringOutput)
}

func (o ProjectOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProjectOutput{})
}
