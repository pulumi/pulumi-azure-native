


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupShare(ctx *pulumi.Context, args *LookupShareArgs, opts ...pulumi.InvokeOption) (*LookupShareResult, error) {
	var rv LookupShareResult
	err := ctx.Invoke("azure-native:databoxedge/v20200901preview:getShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupShareArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupShareResult struct {
	AccessProtocol     string                      `pulumi:"accessProtocol"`
	AzureContainerInfo *AzureContainerInfoResponse `pulumi:"azureContainerInfo"`
	ClientAccessRights []ClientAccessRightResponse `pulumi:"clientAccessRights"`
	DataPolicy         *string                     `pulumi:"dataPolicy"`
	Description        *string                     `pulumi:"description"`
	Id                 string                      `pulumi:"id"`
	MonitoringStatus   string                      `pulumi:"monitoringStatus"`
	Name               string                      `pulumi:"name"`
	RefreshDetails     *RefreshDetailsResponse     `pulumi:"refreshDetails"`
	ShareMappings      []MountPointMapResponse     `pulumi:"shareMappings"`
	ShareStatus        string                      `pulumi:"shareStatus"`
	SystemData         SystemDataResponse          `pulumi:"systemData"`
	Type               string                      `pulumi:"type"`
	UserAccessRights   []UserAccessRightResponse   `pulumi:"userAccessRights"`
}

func LookupShareOutput(ctx *pulumi.Context, args LookupShareOutputArgs, opts ...pulumi.InvokeOption) LookupShareResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupShareResult, error) {
			args := v.(LookupShareArgs)
			r, err := LookupShare(ctx, &args, opts...)
			var s LookupShareResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupShareResultOutput)
}

type LookupShareOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupShareOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareArgs)(nil)).Elem()
}


type LookupShareResultOutput struct{ *pulumi.OutputState }

func (LookupShareResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareResult)(nil)).Elem()
}

func (o LookupShareResultOutput) ToLookupShareResultOutput() LookupShareResultOutput {
	return o
}

func (o LookupShareResultOutput) ToLookupShareResultOutputWithContext(ctx context.Context) LookupShareResultOutput {
	return o
}

func (o LookupShareResultOutput) AccessProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.AccessProtocol }).(pulumi.StringOutput)
}

func (o LookupShareResultOutput) AzureContainerInfo() AzureContainerInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupShareResult) *AzureContainerInfoResponse { return v.AzureContainerInfo }).(AzureContainerInfoResponsePtrOutput)
}

func (o LookupShareResultOutput) ClientAccessRights() ClientAccessRightResponseArrayOutput {
	return o.ApplyT(func(v LookupShareResult) []ClientAccessRightResponse { return v.ClientAccessRights }).(ClientAccessRightResponseArrayOutput)
}

func (o LookupShareResultOutput) DataPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupShareResult) *string { return v.DataPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupShareResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupShareResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupShareResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupShareResultOutput) MonitoringStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.MonitoringStatus }).(pulumi.StringOutput)
}

func (o LookupShareResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupShareResultOutput) RefreshDetails() RefreshDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupShareResult) *RefreshDetailsResponse { return v.RefreshDetails }).(RefreshDetailsResponsePtrOutput)
}

func (o LookupShareResultOutput) ShareMappings() MountPointMapResponseArrayOutput {
	return o.ApplyT(func(v LookupShareResult) []MountPointMapResponse { return v.ShareMappings }).(MountPointMapResponseArrayOutput)
}

func (o LookupShareResultOutput) ShareStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.ShareStatus }).(pulumi.StringOutput)
}

func (o LookupShareResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupShareResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupShareResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupShareResultOutput) UserAccessRights() UserAccessRightResponseArrayOutput {
	return o.ApplyT(func(v LookupShareResult) []UserAccessRightResponse { return v.UserAccessRights }).(UserAccessRightResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupShareResultOutput{})
}
