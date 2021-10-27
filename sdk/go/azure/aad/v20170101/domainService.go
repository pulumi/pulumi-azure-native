


package v20170101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DomainService struct {
	pulumi.CustomResourceState

	DeploymentId              pulumi.StringOutput                     `pulumi:"deploymentId"`
	DomainControllerIpAddress pulumi.StringArrayOutput                `pulumi:"domainControllerIpAddress"`
	DomainName                pulumi.StringPtrOutput                  `pulumi:"domainName"`
	DomainSecuritySettings    DomainSecuritySettingsResponsePtrOutput `pulumi:"domainSecuritySettings"`
	Etag                      pulumi.StringPtrOutput                  `pulumi:"etag"`
	FilteredSync              pulumi.StringPtrOutput                  `pulumi:"filteredSync"`
	HealthAlerts              HealthAlertResponseArrayOutput          `pulumi:"healthAlerts"`
	HealthLastEvaluated       pulumi.StringOutput                     `pulumi:"healthLastEvaluated"`
	HealthMonitors            HealthMonitorResponseArrayOutput        `pulumi:"healthMonitors"`
	LdapsSettings             LdapsSettingsResponsePtrOutput          `pulumi:"ldapsSettings"`
	Location                  pulumi.StringPtrOutput                  `pulumi:"location"`
	Name                      pulumi.StringOutput                     `pulumi:"name"`
	NotificationSettings      NotificationSettingsResponsePtrOutput   `pulumi:"notificationSettings"`
	ProvisioningState         pulumi.StringOutput                     `pulumi:"provisioningState"`
	ServiceStatus             pulumi.StringOutput                     `pulumi:"serviceStatus"`
	SubnetId                  pulumi.StringPtrOutput                  `pulumi:"subnetId"`
	Tags                      pulumi.StringMapOutput                  `pulumi:"tags"`
	TenantId                  pulumi.StringOutput                     `pulumi:"tenantId"`
	Type                      pulumi.StringOutput                     `pulumi:"type"`
	VnetSiteId                pulumi.StringOutput                     `pulumi:"vnetSiteId"`
}


func NewDomainService(ctx *pulumi.Context,
	name string, args *DomainServiceArgs, opts ...pulumi.ResourceOption) (*DomainService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:aad/v20170101:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad:DomainService"),
		},
		{
			Type: pulumi.String("azure-nextgen:aad:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20170601:DomainService"),
		},
		{
			Type: pulumi.String("azure-nextgen:aad/v20170601:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20200101:DomainService"),
		},
		{
			Type: pulumi.String("azure-nextgen:aad/v20200101:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20210301:DomainService"),
		},
		{
			Type: pulumi.String("azure-nextgen:aad/v20210301:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20210501:DomainService"),
		},
		{
			Type: pulumi.String("azure-nextgen:aad/v20210501:DomainService"),
		},
	})
	opts = append(opts, aliases)
	var resource DomainService
	err := ctx.RegisterResource("azure-native:aad/v20170101:DomainService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDomainService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainServiceState, opts ...pulumi.ResourceOption) (*DomainService, error) {
	var resource DomainService
	err := ctx.ReadResource("azure-native:aad/v20170101:DomainService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type domainServiceState struct {
}

type DomainServiceState struct {
}

func (DomainServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainServiceState)(nil)).Elem()
}

type domainServiceArgs struct {
	DomainName             *string                 `pulumi:"domainName"`
	DomainSecuritySettings *DomainSecuritySettings `pulumi:"domainSecuritySettings"`
	DomainServiceName      *string                 `pulumi:"domainServiceName"`
	Etag                   *string                 `pulumi:"etag"`
	FilteredSync           *string                 `pulumi:"filteredSync"`
	LdapsSettings          *LdapsSettings          `pulumi:"ldapsSettings"`
	Location               *string                 `pulumi:"location"`
	NotificationSettings   *NotificationSettings   `pulumi:"notificationSettings"`
	ResourceGroupName      string                  `pulumi:"resourceGroupName"`
	SubnetId               *string                 `pulumi:"subnetId"`
	Tags                   map[string]string       `pulumi:"tags"`
}


type DomainServiceArgs struct {
	DomainName             pulumi.StringPtrInput
	DomainSecuritySettings DomainSecuritySettingsPtrInput
	DomainServiceName      pulumi.StringPtrInput
	Etag                   pulumi.StringPtrInput
	FilteredSync           pulumi.StringPtrInput
	LdapsSettings          LdapsSettingsPtrInput
	Location               pulumi.StringPtrInput
	NotificationSettings   NotificationSettingsPtrInput
	ResourceGroupName      pulumi.StringInput
	SubnetId               pulumi.StringPtrInput
	Tags                   pulumi.StringMapInput
}

func (DomainServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainServiceArgs)(nil)).Elem()
}

type DomainServiceInput interface {
	pulumi.Input

	ToDomainServiceOutput() DomainServiceOutput
	ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput
}

func (*DomainService) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainService)(nil))
}

func (i *DomainService) ToDomainServiceOutput() DomainServiceOutput {
	return i.ToDomainServiceOutputWithContext(context.Background())
}

func (i *DomainService) ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainServiceOutput)
}

type DomainServiceOutput struct{ *pulumi.OutputState }

func (DomainServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainService)(nil))
}

func (o DomainServiceOutput) ToDomainServiceOutput() DomainServiceOutput {
	return o
}

func (o DomainServiceOutput) ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainServiceInput)(nil)).Elem(), &DomainService{})
	pulumi.RegisterOutputType(DomainServiceOutput{})
}
