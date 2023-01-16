


package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOuContainer(ctx *pulumi.Context, args *LookupOuContainerArgs, opts ...pulumi.InvokeOption) (*LookupOuContainerResult, error) {
	var rv LookupOuContainerResult
	err := ctx.Invoke("azure-native:aad/v20220901:getOuContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOuContainerArgs struct {
	DomainServiceName string `pulumi:"domainServiceName"`
	OuContainerName   string `pulumi:"ouContainerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOuContainerResult struct {
	Accounts          []ContainerAccountResponse `pulumi:"accounts"`
	ContainerId       string                     `pulumi:"containerId"`
	DeploymentId      string                     `pulumi:"deploymentId"`
	DistinguishedName string                     `pulumi:"distinguishedName"`
	DomainName        string                     `pulumi:"domainName"`
	Etag              *string                    `pulumi:"etag"`
	Id                string                     `pulumi:"id"`
	Location          *string                    `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	ServiceStatus     string                     `pulumi:"serviceStatus"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	Tags              map[string]string          `pulumi:"tags"`
	TenantId          string                     `pulumi:"tenantId"`
	Type              string                     `pulumi:"type"`
}

func LookupOuContainerOutput(ctx *pulumi.Context, args LookupOuContainerOutputArgs, opts ...pulumi.InvokeOption) LookupOuContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOuContainerResult, error) {
			args := v.(LookupOuContainerArgs)
			r, err := LookupOuContainer(ctx, &args, opts...)
			var s LookupOuContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOuContainerResultOutput)
}

type LookupOuContainerOutputArgs struct {
	DomainServiceName pulumi.StringInput `pulumi:"domainServiceName"`
	OuContainerName   pulumi.StringInput `pulumi:"ouContainerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOuContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOuContainerArgs)(nil)).Elem()
}


type LookupOuContainerResultOutput struct{ *pulumi.OutputState }

func (LookupOuContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOuContainerResult)(nil)).Elem()
}

func (o LookupOuContainerResultOutput) ToLookupOuContainerResultOutput() LookupOuContainerResultOutput {
	return o
}

func (o LookupOuContainerResultOutput) ToLookupOuContainerResultOutputWithContext(ctx context.Context) LookupOuContainerResultOutput {
	return o
}

func (o LookupOuContainerResultOutput) Accounts() ContainerAccountResponseArrayOutput {
	return o.ApplyT(func(v LookupOuContainerResult) []ContainerAccountResponse { return v.Accounts }).(ContainerAccountResponseArrayOutput)
}

func (o LookupOuContainerResultOutput) ContainerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.ContainerId }).(pulumi.StringOutput)
}

func (o LookupOuContainerResultOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.DeploymentId }).(pulumi.StringOutput)
}

func (o LookupOuContainerResultOutput) DistinguishedName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.DistinguishedName }).(pulumi.StringOutput)
}

func (o LookupOuContainerResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o LookupOuContainerResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOuContainerResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupOuContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOuContainerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOuContainerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupOuContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOuContainerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupOuContainerResultOutput) ServiceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.ServiceStatus }).(pulumi.StringOutput)
}

func (o LookupOuContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOuContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOuContainerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOuContainerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupOuContainerResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupOuContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOuContainerResultOutput{})
}
