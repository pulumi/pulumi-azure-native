


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActiveDirectory struct {
	ActiveDirectoryId  *string `pulumi:"activeDirectoryId"`
	Dns                *string `pulumi:"dns"`
	Domain             *string `pulumi:"domain"`
	OrganizationalUnit *string `pulumi:"organizationalUnit"`
	Password           *string `pulumi:"password"`
	SmbServerName      *string `pulumi:"smbServerName"`
	Status             *string `pulumi:"status"`
	Username           *string `pulumi:"username"`
}





type ActiveDirectoryInput interface {
	pulumi.Input

	ToActiveDirectoryOutput() ActiveDirectoryOutput
	ToActiveDirectoryOutputWithContext(context.Context) ActiveDirectoryOutput
}

type ActiveDirectoryArgs struct {
	ActiveDirectoryId  pulumi.StringPtrInput `pulumi:"activeDirectoryId"`
	Dns                pulumi.StringPtrInput `pulumi:"dns"`
	Domain             pulumi.StringPtrInput `pulumi:"domain"`
	OrganizationalUnit pulumi.StringPtrInput `pulumi:"organizationalUnit"`
	Password           pulumi.StringPtrInput `pulumi:"password"`
	SmbServerName      pulumi.StringPtrInput `pulumi:"smbServerName"`
	Status             pulumi.StringPtrInput `pulumi:"status"`
	Username           pulumi.StringPtrInput `pulumi:"username"`
}

func (ActiveDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectory)(nil)).Elem()
}

func (i ActiveDirectoryArgs) ToActiveDirectoryOutput() ActiveDirectoryOutput {
	return i.ToActiveDirectoryOutputWithContext(context.Background())
}

func (i ActiveDirectoryArgs) ToActiveDirectoryOutputWithContext(ctx context.Context) ActiveDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryOutput)
}





type ActiveDirectoryArrayInput interface {
	pulumi.Input

	ToActiveDirectoryArrayOutput() ActiveDirectoryArrayOutput
	ToActiveDirectoryArrayOutputWithContext(context.Context) ActiveDirectoryArrayOutput
}

type ActiveDirectoryArray []ActiveDirectoryInput

func (ActiveDirectoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveDirectory)(nil)).Elem()
}

func (i ActiveDirectoryArray) ToActiveDirectoryArrayOutput() ActiveDirectoryArrayOutput {
	return i.ToActiveDirectoryArrayOutputWithContext(context.Background())
}

func (i ActiveDirectoryArray) ToActiveDirectoryArrayOutputWithContext(ctx context.Context) ActiveDirectoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryArrayOutput)
}

type ActiveDirectoryOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectory)(nil)).Elem()
}

func (o ActiveDirectoryOutput) ToActiveDirectoryOutput() ActiveDirectoryOutput {
	return o
}

func (o ActiveDirectoryOutput) ToActiveDirectoryOutputWithContext(ctx context.Context) ActiveDirectoryOutput {
	return o
}

func (o ActiveDirectoryOutput) ActiveDirectoryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.ActiveDirectoryId }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) Dns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.Dns }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) OrganizationalUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.OrganizationalUnit }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) SmbServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.SmbServerName }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ActiveDirectoryArrayOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveDirectory)(nil)).Elem()
}

func (o ActiveDirectoryArrayOutput) ToActiveDirectoryArrayOutput() ActiveDirectoryArrayOutput {
	return o
}

func (o ActiveDirectoryArrayOutput) ToActiveDirectoryArrayOutputWithContext(ctx context.Context) ActiveDirectoryArrayOutput {
	return o
}

func (o ActiveDirectoryArrayOutput) Index(i pulumi.IntInput) ActiveDirectoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActiveDirectory {
		return vs[0].([]ActiveDirectory)[vs[1].(int)]
	}).(ActiveDirectoryOutput)
}

type ActiveDirectoryResponse struct {
	ActiveDirectoryId  *string `pulumi:"activeDirectoryId"`
	Dns                *string `pulumi:"dns"`
	Domain             *string `pulumi:"domain"`
	OrganizationalUnit *string `pulumi:"organizationalUnit"`
	Password           *string `pulumi:"password"`
	SmbServerName      *string `pulumi:"smbServerName"`
	Status             *string `pulumi:"status"`
	Username           *string `pulumi:"username"`
}

type ActiveDirectoryResponseOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryResponse)(nil)).Elem()
}

func (o ActiveDirectoryResponseOutput) ToActiveDirectoryResponseOutput() ActiveDirectoryResponseOutput {
	return o
}

func (o ActiveDirectoryResponseOutput) ToActiveDirectoryResponseOutputWithContext(ctx context.Context) ActiveDirectoryResponseOutput {
	return o
}

func (o ActiveDirectoryResponseOutput) ActiveDirectoryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.ActiveDirectoryId }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) Dns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.Dns }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) OrganizationalUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.OrganizationalUnit }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) SmbServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.SmbServerName }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ActiveDirectoryResponseArrayOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveDirectoryResponse)(nil)).Elem()
}

func (o ActiveDirectoryResponseArrayOutput) ToActiveDirectoryResponseArrayOutput() ActiveDirectoryResponseArrayOutput {
	return o
}

func (o ActiveDirectoryResponseArrayOutput) ToActiveDirectoryResponseArrayOutputWithContext(ctx context.Context) ActiveDirectoryResponseArrayOutput {
	return o
}

func (o ActiveDirectoryResponseArrayOutput) Index(i pulumi.IntInput) ActiveDirectoryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActiveDirectoryResponse {
		return vs[0].([]ActiveDirectoryResponse)[vs[1].(int)]
	}).(ActiveDirectoryResponseOutput)
}

type ExportPolicyRule struct {
	AllowedClients *string `pulumi:"allowedClients"`
	Cifs           *bool   `pulumi:"cifs"`
	Nfsv3          *bool   `pulumi:"nfsv3"`
	Nfsv4          *bool   `pulumi:"nfsv4"`
	RuleIndex      *int    `pulumi:"ruleIndex"`
	UnixReadOnly   *bool   `pulumi:"unixReadOnly"`
	UnixReadWrite  *bool   `pulumi:"unixReadWrite"`
}





type ExportPolicyRuleInput interface {
	pulumi.Input

	ToExportPolicyRuleOutput() ExportPolicyRuleOutput
	ToExportPolicyRuleOutputWithContext(context.Context) ExportPolicyRuleOutput
}

type ExportPolicyRuleArgs struct {
	AllowedClients pulumi.StringPtrInput `pulumi:"allowedClients"`
	Cifs           pulumi.BoolPtrInput   `pulumi:"cifs"`
	Nfsv3          pulumi.BoolPtrInput   `pulumi:"nfsv3"`
	Nfsv4          pulumi.BoolPtrInput   `pulumi:"nfsv4"`
	RuleIndex      pulumi.IntPtrInput    `pulumi:"ruleIndex"`
	UnixReadOnly   pulumi.BoolPtrInput   `pulumi:"unixReadOnly"`
	UnixReadWrite  pulumi.BoolPtrInput   `pulumi:"unixReadWrite"`
}

func (ExportPolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPolicyRule)(nil)).Elem()
}

func (i ExportPolicyRuleArgs) ToExportPolicyRuleOutput() ExportPolicyRuleOutput {
	return i.ToExportPolicyRuleOutputWithContext(context.Background())
}

func (i ExportPolicyRuleArgs) ToExportPolicyRuleOutputWithContext(ctx context.Context) ExportPolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPolicyRuleOutput)
}





type ExportPolicyRuleArrayInput interface {
	pulumi.Input

	ToExportPolicyRuleArrayOutput() ExportPolicyRuleArrayOutput
	ToExportPolicyRuleArrayOutputWithContext(context.Context) ExportPolicyRuleArrayOutput
}

type ExportPolicyRuleArray []ExportPolicyRuleInput

func (ExportPolicyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExportPolicyRule)(nil)).Elem()
}

func (i ExportPolicyRuleArray) ToExportPolicyRuleArrayOutput() ExportPolicyRuleArrayOutput {
	return i.ToExportPolicyRuleArrayOutputWithContext(context.Background())
}

func (i ExportPolicyRuleArray) ToExportPolicyRuleArrayOutputWithContext(ctx context.Context) ExportPolicyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPolicyRuleArrayOutput)
}

type ExportPolicyRuleOutput struct{ *pulumi.OutputState }

func (ExportPolicyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPolicyRule)(nil)).Elem()
}

func (o ExportPolicyRuleOutput) ToExportPolicyRuleOutput() ExportPolicyRuleOutput {
	return o
}

func (o ExportPolicyRuleOutput) ToExportPolicyRuleOutputWithContext(ctx context.Context) ExportPolicyRuleOutput {
	return o
}

func (o ExportPolicyRuleOutput) AllowedClients() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *string { return v.AllowedClients }).(pulumi.StringPtrOutput)
}

func (o ExportPolicyRuleOutput) Cifs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.Cifs }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) Nfsv3() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.Nfsv3 }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) Nfsv4() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.Nfsv4 }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) RuleIndex() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *int { return v.RuleIndex }).(pulumi.IntPtrOutput)
}

func (o ExportPolicyRuleOutput) UnixReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.UnixReadOnly }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleOutput) UnixReadWrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.UnixReadWrite }).(pulumi.BoolPtrOutput)
}

type ExportPolicyRuleArrayOutput struct{ *pulumi.OutputState }

func (ExportPolicyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExportPolicyRule)(nil)).Elem()
}

func (o ExportPolicyRuleArrayOutput) ToExportPolicyRuleArrayOutput() ExportPolicyRuleArrayOutput {
	return o
}

func (o ExportPolicyRuleArrayOutput) ToExportPolicyRuleArrayOutputWithContext(ctx context.Context) ExportPolicyRuleArrayOutput {
	return o
}

func (o ExportPolicyRuleArrayOutput) Index(i pulumi.IntInput) ExportPolicyRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExportPolicyRule {
		return vs[0].([]ExportPolicyRule)[vs[1].(int)]
	}).(ExportPolicyRuleOutput)
}

type ExportPolicyRuleResponse struct {
	AllowedClients *string `pulumi:"allowedClients"`
	Cifs           *bool   `pulumi:"cifs"`
	Nfsv3          *bool   `pulumi:"nfsv3"`
	Nfsv4          *bool   `pulumi:"nfsv4"`
	RuleIndex      *int    `pulumi:"ruleIndex"`
	UnixReadOnly   *bool   `pulumi:"unixReadOnly"`
	UnixReadWrite  *bool   `pulumi:"unixReadWrite"`
}

type ExportPolicyRuleResponseOutput struct{ *pulumi.OutputState }

func (ExportPolicyRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPolicyRuleResponse)(nil)).Elem()
}

func (o ExportPolicyRuleResponseOutput) ToExportPolicyRuleResponseOutput() ExportPolicyRuleResponseOutput {
	return o
}

func (o ExportPolicyRuleResponseOutput) ToExportPolicyRuleResponseOutputWithContext(ctx context.Context) ExportPolicyRuleResponseOutput {
	return o
}

func (o ExportPolicyRuleResponseOutput) AllowedClients() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *string { return v.AllowedClients }).(pulumi.StringPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) Cifs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.Cifs }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) Nfsv3() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.Nfsv3 }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) Nfsv4() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.Nfsv4 }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) RuleIndex() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *int { return v.RuleIndex }).(pulumi.IntPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) UnixReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.UnixReadOnly }).(pulumi.BoolPtrOutput)
}

func (o ExportPolicyRuleResponseOutput) UnixReadWrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.UnixReadWrite }).(pulumi.BoolPtrOutput)
}

type ExportPolicyRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ExportPolicyRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExportPolicyRuleResponse)(nil)).Elem()
}

func (o ExportPolicyRuleResponseArrayOutput) ToExportPolicyRuleResponseArrayOutput() ExportPolicyRuleResponseArrayOutput {
	return o
}

func (o ExportPolicyRuleResponseArrayOutput) ToExportPolicyRuleResponseArrayOutputWithContext(ctx context.Context) ExportPolicyRuleResponseArrayOutput {
	return o
}

func (o ExportPolicyRuleResponseArrayOutput) Index(i pulumi.IntInput) ExportPolicyRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExportPolicyRuleResponse {
		return vs[0].([]ExportPolicyRuleResponse)[vs[1].(int)]
	}).(ExportPolicyRuleResponseOutput)
}

type MountTargetProperties struct {
	EndIp         *string `pulumi:"endIp"`
	FileSystemId  string  `pulumi:"fileSystemId"`
	Gateway       *string `pulumi:"gateway"`
	Netmask       *string `pulumi:"netmask"`
	SmbServerFqdn *string `pulumi:"smbServerFqdn"`
	StartIp       *string `pulumi:"startIp"`
	Subnet        *string `pulumi:"subnet"`
}





type MountTargetPropertiesInput interface {
	pulumi.Input

	ToMountTargetPropertiesOutput() MountTargetPropertiesOutput
	ToMountTargetPropertiesOutputWithContext(context.Context) MountTargetPropertiesOutput
}

type MountTargetPropertiesArgs struct {
	EndIp         pulumi.StringPtrInput `pulumi:"endIp"`
	FileSystemId  pulumi.StringInput    `pulumi:"fileSystemId"`
	Gateway       pulumi.StringPtrInput `pulumi:"gateway"`
	Netmask       pulumi.StringPtrInput `pulumi:"netmask"`
	SmbServerFqdn pulumi.StringPtrInput `pulumi:"smbServerFqdn"`
	StartIp       pulumi.StringPtrInput `pulumi:"startIp"`
	Subnet        pulumi.StringPtrInput `pulumi:"subnet"`
}

func (MountTargetPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MountTargetProperties)(nil)).Elem()
}

func (i MountTargetPropertiesArgs) ToMountTargetPropertiesOutput() MountTargetPropertiesOutput {
	return i.ToMountTargetPropertiesOutputWithContext(context.Background())
}

func (i MountTargetPropertiesArgs) ToMountTargetPropertiesOutputWithContext(ctx context.Context) MountTargetPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountTargetPropertiesOutput)
}





type MountTargetPropertiesArrayInput interface {
	pulumi.Input

	ToMountTargetPropertiesArrayOutput() MountTargetPropertiesArrayOutput
	ToMountTargetPropertiesArrayOutputWithContext(context.Context) MountTargetPropertiesArrayOutput
}

type MountTargetPropertiesArray []MountTargetPropertiesInput

func (MountTargetPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MountTargetProperties)(nil)).Elem()
}

func (i MountTargetPropertiesArray) ToMountTargetPropertiesArrayOutput() MountTargetPropertiesArrayOutput {
	return i.ToMountTargetPropertiesArrayOutputWithContext(context.Background())
}

func (i MountTargetPropertiesArray) ToMountTargetPropertiesArrayOutputWithContext(ctx context.Context) MountTargetPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountTargetPropertiesArrayOutput)
}

type MountTargetPropertiesOutput struct{ *pulumi.OutputState }

func (MountTargetPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MountTargetProperties)(nil)).Elem()
}

func (o MountTargetPropertiesOutput) ToMountTargetPropertiesOutput() MountTargetPropertiesOutput {
	return o
}

func (o MountTargetPropertiesOutput) ToMountTargetPropertiesOutputWithContext(ctx context.Context) MountTargetPropertiesOutput {
	return o
}

func (o MountTargetPropertiesOutput) EndIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MountTargetProperties) *string { return v.EndIp }).(pulumi.StringPtrOutput)
}

func (o MountTargetPropertiesOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v MountTargetProperties) string { return v.FileSystemId }).(pulumi.StringOutput)
}

func (o MountTargetPropertiesOutput) Gateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MountTargetProperties) *string { return v.Gateway }).(pulumi.StringPtrOutput)
}

func (o MountTargetPropertiesOutput) Netmask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MountTargetProperties) *string { return v.Netmask }).(pulumi.StringPtrOutput)
}

func (o MountTargetPropertiesOutput) SmbServerFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MountTargetProperties) *string { return v.SmbServerFqdn }).(pulumi.StringPtrOutput)
}

func (o MountTargetPropertiesOutput) StartIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MountTargetProperties) *string { return v.StartIp }).(pulumi.StringPtrOutput)
}

func (o MountTargetPropertiesOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MountTargetProperties) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

type MountTargetPropertiesArrayOutput struct{ *pulumi.OutputState }

func (MountTargetPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MountTargetProperties)(nil)).Elem()
}

func (o MountTargetPropertiesArrayOutput) ToMountTargetPropertiesArrayOutput() MountTargetPropertiesArrayOutput {
	return o
}

func (o MountTargetPropertiesArrayOutput) ToMountTargetPropertiesArrayOutputWithContext(ctx context.Context) MountTargetPropertiesArrayOutput {
	return o
}

func (o MountTargetPropertiesArrayOutput) Index(i pulumi.IntInput) MountTargetPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MountTargetProperties {
		return vs[0].([]MountTargetProperties)[vs[1].(int)]
	}).(MountTargetPropertiesOutput)
}

type MountTargetPropertiesResponse struct {
	EndIp             *string `pulumi:"endIp"`
	FileSystemId      string  `pulumi:"fileSystemId"`
	Gateway           *string `pulumi:"gateway"`
	IpAddress         string  `pulumi:"ipAddress"`
	MountTargetId     string  `pulumi:"mountTargetId"`
	Netmask           *string `pulumi:"netmask"`
	ProvisioningState string  `pulumi:"provisioningState"`
	SmbServerFqdn     *string `pulumi:"smbServerFqdn"`
	StartIp           *string `pulumi:"startIp"`
	Subnet            *string `pulumi:"subnet"`
}

type MountTargetPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MountTargetPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MountTargetPropertiesResponse)(nil)).Elem()
}

func (o MountTargetPropertiesResponseOutput) ToMountTargetPropertiesResponseOutput() MountTargetPropertiesResponseOutput {
	return o
}

func (o MountTargetPropertiesResponseOutput) ToMountTargetPropertiesResponseOutputWithContext(ctx context.Context) MountTargetPropertiesResponseOutput {
	return o
}

func (o MountTargetPropertiesResponseOutput) EndIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) *string { return v.EndIp }).(pulumi.StringPtrOutput)
}

func (o MountTargetPropertiesResponseOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) string { return v.FileSystemId }).(pulumi.StringOutput)
}

func (o MountTargetPropertiesResponseOutput) Gateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) *string { return v.Gateway }).(pulumi.StringPtrOutput)
}

func (o MountTargetPropertiesResponseOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o MountTargetPropertiesResponseOutput) MountTargetId() pulumi.StringOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) string { return v.MountTargetId }).(pulumi.StringOutput)
}

func (o MountTargetPropertiesResponseOutput) Netmask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) *string { return v.Netmask }).(pulumi.StringPtrOutput)
}

func (o MountTargetPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MountTargetPropertiesResponseOutput) SmbServerFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) *string { return v.SmbServerFqdn }).(pulumi.StringPtrOutput)
}

func (o MountTargetPropertiesResponseOutput) StartIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) *string { return v.StartIp }).(pulumi.StringPtrOutput)
}

func (o MountTargetPropertiesResponseOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MountTargetPropertiesResponse) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

type MountTargetPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (MountTargetPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MountTargetPropertiesResponse)(nil)).Elem()
}

func (o MountTargetPropertiesResponseArrayOutput) ToMountTargetPropertiesResponseArrayOutput() MountTargetPropertiesResponseArrayOutput {
	return o
}

func (o MountTargetPropertiesResponseArrayOutput) ToMountTargetPropertiesResponseArrayOutputWithContext(ctx context.Context) MountTargetPropertiesResponseArrayOutput {
	return o
}

func (o MountTargetPropertiesResponseArrayOutput) Index(i pulumi.IntInput) MountTargetPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MountTargetPropertiesResponse {
		return vs[0].([]MountTargetPropertiesResponse)[vs[1].(int)]
	}).(MountTargetPropertiesResponseOutput)
}

type VolumePropertiesExportPolicy struct {
	Rules []ExportPolicyRule `pulumi:"rules"`
}





type VolumePropertiesExportPolicyInput interface {
	pulumi.Input

	ToVolumePropertiesExportPolicyOutput() VolumePropertiesExportPolicyOutput
	ToVolumePropertiesExportPolicyOutputWithContext(context.Context) VolumePropertiesExportPolicyOutput
}

type VolumePropertiesExportPolicyArgs struct {
	Rules ExportPolicyRuleArrayInput `pulumi:"rules"`
}

func (VolumePropertiesExportPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesExportPolicy)(nil)).Elem()
}

func (i VolumePropertiesExportPolicyArgs) ToVolumePropertiesExportPolicyOutput() VolumePropertiesExportPolicyOutput {
	return i.ToVolumePropertiesExportPolicyOutputWithContext(context.Background())
}

func (i VolumePropertiesExportPolicyArgs) ToVolumePropertiesExportPolicyOutputWithContext(ctx context.Context) VolumePropertiesExportPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesExportPolicyOutput)
}

func (i VolumePropertiesExportPolicyArgs) ToVolumePropertiesExportPolicyPtrOutput() VolumePropertiesExportPolicyPtrOutput {
	return i.ToVolumePropertiesExportPolicyPtrOutputWithContext(context.Background())
}

func (i VolumePropertiesExportPolicyArgs) ToVolumePropertiesExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesExportPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesExportPolicyOutput).ToVolumePropertiesExportPolicyPtrOutputWithContext(ctx)
}









type VolumePropertiesExportPolicyPtrInput interface {
	pulumi.Input

	ToVolumePropertiesExportPolicyPtrOutput() VolumePropertiesExportPolicyPtrOutput
	ToVolumePropertiesExportPolicyPtrOutputWithContext(context.Context) VolumePropertiesExportPolicyPtrOutput
}

type volumePropertiesExportPolicyPtrType VolumePropertiesExportPolicyArgs

func VolumePropertiesExportPolicyPtr(v *VolumePropertiesExportPolicyArgs) VolumePropertiesExportPolicyPtrInput {
	return (*volumePropertiesExportPolicyPtrType)(v)
}

func (*volumePropertiesExportPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesExportPolicy)(nil)).Elem()
}

func (i *volumePropertiesExportPolicyPtrType) ToVolumePropertiesExportPolicyPtrOutput() VolumePropertiesExportPolicyPtrOutput {
	return i.ToVolumePropertiesExportPolicyPtrOutputWithContext(context.Background())
}

func (i *volumePropertiesExportPolicyPtrType) ToVolumePropertiesExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesExportPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesExportPolicyPtrOutput)
}

type VolumePropertiesExportPolicyOutput struct{ *pulumi.OutputState }

func (VolumePropertiesExportPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesExportPolicy)(nil)).Elem()
}

func (o VolumePropertiesExportPolicyOutput) ToVolumePropertiesExportPolicyOutput() VolumePropertiesExportPolicyOutput {
	return o
}

func (o VolumePropertiesExportPolicyOutput) ToVolumePropertiesExportPolicyOutputWithContext(ctx context.Context) VolumePropertiesExportPolicyOutput {
	return o
}

func (o VolumePropertiesExportPolicyOutput) ToVolumePropertiesExportPolicyPtrOutput() VolumePropertiesExportPolicyPtrOutput {
	return o.ToVolumePropertiesExportPolicyPtrOutputWithContext(context.Background())
}

func (o VolumePropertiesExportPolicyOutput) ToVolumePropertiesExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesExportPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumePropertiesExportPolicy) *VolumePropertiesExportPolicy {
		return &v
	}).(VolumePropertiesExportPolicyPtrOutput)
}

func (o VolumePropertiesExportPolicyOutput) Rules() ExportPolicyRuleArrayOutput {
	return o.ApplyT(func(v VolumePropertiesExportPolicy) []ExportPolicyRule { return v.Rules }).(ExportPolicyRuleArrayOutput)
}

type VolumePropertiesExportPolicyPtrOutput struct{ *pulumi.OutputState }

func (VolumePropertiesExportPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesExportPolicy)(nil)).Elem()
}

func (o VolumePropertiesExportPolicyPtrOutput) ToVolumePropertiesExportPolicyPtrOutput() VolumePropertiesExportPolicyPtrOutput {
	return o
}

func (o VolumePropertiesExportPolicyPtrOutput) ToVolumePropertiesExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesExportPolicyPtrOutput {
	return o
}

func (o VolumePropertiesExportPolicyPtrOutput) Elem() VolumePropertiesExportPolicyOutput {
	return o.ApplyT(func(v *VolumePropertiesExportPolicy) VolumePropertiesExportPolicy {
		if v != nil {
			return *v
		}
		var ret VolumePropertiesExportPolicy
		return ret
	}).(VolumePropertiesExportPolicyOutput)
}

func (o VolumePropertiesExportPolicyPtrOutput) Rules() ExportPolicyRuleArrayOutput {
	return o.ApplyT(func(v *VolumePropertiesExportPolicy) []ExportPolicyRule {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(ExportPolicyRuleArrayOutput)
}

type VolumePropertiesResponseExportPolicy struct {
	Rules []ExportPolicyRuleResponse `pulumi:"rules"`
}

type VolumePropertiesResponseExportPolicyOutput struct{ *pulumi.OutputState }

func (VolumePropertiesResponseExportPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesResponseExportPolicy)(nil)).Elem()
}

func (o VolumePropertiesResponseExportPolicyOutput) ToVolumePropertiesResponseExportPolicyOutput() VolumePropertiesResponseExportPolicyOutput {
	return o
}

func (o VolumePropertiesResponseExportPolicyOutput) ToVolumePropertiesResponseExportPolicyOutputWithContext(ctx context.Context) VolumePropertiesResponseExportPolicyOutput {
	return o
}

func (o VolumePropertiesResponseExportPolicyOutput) Rules() ExportPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v VolumePropertiesResponseExportPolicy) []ExportPolicyRuleResponse { return v.Rules }).(ExportPolicyRuleResponseArrayOutput)
}

type VolumePropertiesResponseExportPolicyPtrOutput struct{ *pulumi.OutputState }

func (VolumePropertiesResponseExportPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesResponseExportPolicy)(nil)).Elem()
}

func (o VolumePropertiesResponseExportPolicyPtrOutput) ToVolumePropertiesResponseExportPolicyPtrOutput() VolumePropertiesResponseExportPolicyPtrOutput {
	return o
}

func (o VolumePropertiesResponseExportPolicyPtrOutput) ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseExportPolicyPtrOutput {
	return o
}

func (o VolumePropertiesResponseExportPolicyPtrOutput) Elem() VolumePropertiesResponseExportPolicyOutput {
	return o.ApplyT(func(v *VolumePropertiesResponseExportPolicy) VolumePropertiesResponseExportPolicy {
		if v != nil {
			return *v
		}
		var ret VolumePropertiesResponseExportPolicy
		return ret
	}).(VolumePropertiesResponseExportPolicyOutput)
}

func (o VolumePropertiesResponseExportPolicyPtrOutput) Rules() ExportPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v *VolumePropertiesResponseExportPolicy) []ExportPolicyRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(ExportPolicyRuleResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ActiveDirectoryOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryArrayOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryResponseOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryResponseArrayOutput{})
	pulumi.RegisterOutputType(ExportPolicyRuleOutput{})
	pulumi.RegisterOutputType(ExportPolicyRuleArrayOutput{})
	pulumi.RegisterOutputType(ExportPolicyRuleResponseOutput{})
	pulumi.RegisterOutputType(ExportPolicyRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(MountTargetPropertiesOutput{})
	pulumi.RegisterOutputType(MountTargetPropertiesArrayOutput{})
	pulumi.RegisterOutputType(MountTargetPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MountTargetPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(VolumePropertiesExportPolicyOutput{})
	pulumi.RegisterOutputType(VolumePropertiesExportPolicyPtrOutput{})
	pulumi.RegisterOutputType(VolumePropertiesResponseExportPolicyOutput{})
	pulumi.RegisterOutputType(VolumePropertiesResponseExportPolicyPtrOutput{})
}
