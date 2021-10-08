


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActiveDirectory struct {
	ActiveDirectoryId  *string  `pulumi:"activeDirectoryId"`
	BackupOperators    []string `pulumi:"backupOperators"`
	Dns                *string  `pulumi:"dns"`
	Domain             *string  `pulumi:"domain"`
	OrganizationalUnit *string  `pulumi:"organizationalUnit"`
	Password           *string  `pulumi:"password"`
	Site               *string  `pulumi:"site"`
	SmbServerName      *string  `pulumi:"smbServerName"`
	Status             *string  `pulumi:"status"`
	Username           *string  `pulumi:"username"`
}





type ActiveDirectoryInput interface {
	pulumi.Input

	ToActiveDirectoryOutput() ActiveDirectoryOutput
	ToActiveDirectoryOutputWithContext(context.Context) ActiveDirectoryOutput
}

type ActiveDirectoryArgs struct {
	ActiveDirectoryId  pulumi.StringPtrInput   `pulumi:"activeDirectoryId"`
	BackupOperators    pulumi.StringArrayInput `pulumi:"backupOperators"`
	Dns                pulumi.StringPtrInput   `pulumi:"dns"`
	Domain             pulumi.StringPtrInput   `pulumi:"domain"`
	OrganizationalUnit pulumi.StringPtrInput   `pulumi:"organizationalUnit"`
	Password           pulumi.StringPtrInput   `pulumi:"password"`
	Site               pulumi.StringPtrInput   `pulumi:"site"`
	SmbServerName      pulumi.StringPtrInput   `pulumi:"smbServerName"`
	Status             pulumi.StringPtrInput   `pulumi:"status"`
	Username           pulumi.StringPtrInput   `pulumi:"username"`
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

func (o ActiveDirectoryOutput) BackupOperators() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDirectory) []string { return v.BackupOperators }).(pulumi.StringArrayOutput)
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

func (o ActiveDirectoryOutput) Site() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectory) *string { return v.Site }).(pulumi.StringPtrOutput)
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
	ActiveDirectoryId  *string  `pulumi:"activeDirectoryId"`
	BackupOperators    []string `pulumi:"backupOperators"`
	Dns                *string  `pulumi:"dns"`
	Domain             *string  `pulumi:"domain"`
	OrganizationalUnit *string  `pulumi:"organizationalUnit"`
	Password           *string  `pulumi:"password"`
	Site               *string  `pulumi:"site"`
	SmbServerName      *string  `pulumi:"smbServerName"`
	Status             *string  `pulumi:"status"`
	Username           *string  `pulumi:"username"`
}





type ActiveDirectoryResponseInput interface {
	pulumi.Input

	ToActiveDirectoryResponseOutput() ActiveDirectoryResponseOutput
	ToActiveDirectoryResponseOutputWithContext(context.Context) ActiveDirectoryResponseOutput
}

type ActiveDirectoryResponseArgs struct {
	ActiveDirectoryId  pulumi.StringPtrInput   `pulumi:"activeDirectoryId"`
	BackupOperators    pulumi.StringArrayInput `pulumi:"backupOperators"`
	Dns                pulumi.StringPtrInput   `pulumi:"dns"`
	Domain             pulumi.StringPtrInput   `pulumi:"domain"`
	OrganizationalUnit pulumi.StringPtrInput   `pulumi:"organizationalUnit"`
	Password           pulumi.StringPtrInput   `pulumi:"password"`
	Site               pulumi.StringPtrInput   `pulumi:"site"`
	SmbServerName      pulumi.StringPtrInput   `pulumi:"smbServerName"`
	Status             pulumi.StringPtrInput   `pulumi:"status"`
	Username           pulumi.StringPtrInput   `pulumi:"username"`
}

func (ActiveDirectoryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryResponse)(nil)).Elem()
}

func (i ActiveDirectoryResponseArgs) ToActiveDirectoryResponseOutput() ActiveDirectoryResponseOutput {
	return i.ToActiveDirectoryResponseOutputWithContext(context.Background())
}

func (i ActiveDirectoryResponseArgs) ToActiveDirectoryResponseOutputWithContext(ctx context.Context) ActiveDirectoryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryResponseOutput)
}





type ActiveDirectoryResponseArrayInput interface {
	pulumi.Input

	ToActiveDirectoryResponseArrayOutput() ActiveDirectoryResponseArrayOutput
	ToActiveDirectoryResponseArrayOutputWithContext(context.Context) ActiveDirectoryResponseArrayOutput
}

type ActiveDirectoryResponseArray []ActiveDirectoryResponseInput

func (ActiveDirectoryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveDirectoryResponse)(nil)).Elem()
}

func (i ActiveDirectoryResponseArray) ToActiveDirectoryResponseArrayOutput() ActiveDirectoryResponseArrayOutput {
	return i.ToActiveDirectoryResponseArrayOutputWithContext(context.Background())
}

func (i ActiveDirectoryResponseArray) ToActiveDirectoryResponseArrayOutputWithContext(ctx context.Context) ActiveDirectoryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryResponseArrayOutput)
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

func (o ActiveDirectoryResponseOutput) BackupOperators() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) []string { return v.BackupOperators }).(pulumi.StringArrayOutput)
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

func (o ActiveDirectoryResponseOutput) Site() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryResponse) *string { return v.Site }).(pulumi.StringPtrOutput)
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
	Nfsv41         *bool   `pulumi:"nfsv41"`
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
	Nfsv41         pulumi.BoolPtrInput   `pulumi:"nfsv41"`
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

func (o ExportPolicyRuleOutput) Nfsv41() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRule) *bool { return v.Nfsv41 }).(pulumi.BoolPtrOutput)
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
	Nfsv41         *bool   `pulumi:"nfsv41"`
	RuleIndex      *int    `pulumi:"ruleIndex"`
	UnixReadOnly   *bool   `pulumi:"unixReadOnly"`
	UnixReadWrite  *bool   `pulumi:"unixReadWrite"`
}





type ExportPolicyRuleResponseInput interface {
	pulumi.Input

	ToExportPolicyRuleResponseOutput() ExportPolicyRuleResponseOutput
	ToExportPolicyRuleResponseOutputWithContext(context.Context) ExportPolicyRuleResponseOutput
}

type ExportPolicyRuleResponseArgs struct {
	AllowedClients pulumi.StringPtrInput `pulumi:"allowedClients"`
	Cifs           pulumi.BoolPtrInput   `pulumi:"cifs"`
	Nfsv3          pulumi.BoolPtrInput   `pulumi:"nfsv3"`
	Nfsv41         pulumi.BoolPtrInput   `pulumi:"nfsv41"`
	RuleIndex      pulumi.IntPtrInput    `pulumi:"ruleIndex"`
	UnixReadOnly   pulumi.BoolPtrInput   `pulumi:"unixReadOnly"`
	UnixReadWrite  pulumi.BoolPtrInput   `pulumi:"unixReadWrite"`
}

func (ExportPolicyRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPolicyRuleResponse)(nil)).Elem()
}

func (i ExportPolicyRuleResponseArgs) ToExportPolicyRuleResponseOutput() ExportPolicyRuleResponseOutput {
	return i.ToExportPolicyRuleResponseOutputWithContext(context.Background())
}

func (i ExportPolicyRuleResponseArgs) ToExportPolicyRuleResponseOutputWithContext(ctx context.Context) ExportPolicyRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPolicyRuleResponseOutput)
}





type ExportPolicyRuleResponseArrayInput interface {
	pulumi.Input

	ToExportPolicyRuleResponseArrayOutput() ExportPolicyRuleResponseArrayOutput
	ToExportPolicyRuleResponseArrayOutputWithContext(context.Context) ExportPolicyRuleResponseArrayOutput
}

type ExportPolicyRuleResponseArray []ExportPolicyRuleResponseInput

func (ExportPolicyRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExportPolicyRuleResponse)(nil)).Elem()
}

func (i ExportPolicyRuleResponseArray) ToExportPolicyRuleResponseArrayOutput() ExportPolicyRuleResponseArrayOutput {
	return i.ToExportPolicyRuleResponseArrayOutputWithContext(context.Background())
}

func (i ExportPolicyRuleResponseArray) ToExportPolicyRuleResponseArrayOutputWithContext(ctx context.Context) ExportPolicyRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPolicyRuleResponseArrayOutput)
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

func (o ExportPolicyRuleResponseOutput) Nfsv41() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExportPolicyRuleResponse) *bool { return v.Nfsv41 }).(pulumi.BoolPtrOutput)
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

type MountTargetPropertiesResponse struct {
	EndIp         *string `pulumi:"endIp"`
	FileSystemId  string  `pulumi:"fileSystemId"`
	Gateway       *string `pulumi:"gateway"`
	IpAddress     string  `pulumi:"ipAddress"`
	MountTargetId string  `pulumi:"mountTargetId"`
	Netmask       *string `pulumi:"netmask"`
	SmbServerFqdn *string `pulumi:"smbServerFqdn"`
	StartIp       *string `pulumi:"startIp"`
	Subnet        *string `pulumi:"subnet"`
}





type MountTargetPropertiesResponseInput interface {
	pulumi.Input

	ToMountTargetPropertiesResponseOutput() MountTargetPropertiesResponseOutput
	ToMountTargetPropertiesResponseOutputWithContext(context.Context) MountTargetPropertiesResponseOutput
}

type MountTargetPropertiesResponseArgs struct {
	EndIp         pulumi.StringPtrInput `pulumi:"endIp"`
	FileSystemId  pulumi.StringInput    `pulumi:"fileSystemId"`
	Gateway       pulumi.StringPtrInput `pulumi:"gateway"`
	IpAddress     pulumi.StringInput    `pulumi:"ipAddress"`
	MountTargetId pulumi.StringInput    `pulumi:"mountTargetId"`
	Netmask       pulumi.StringPtrInput `pulumi:"netmask"`
	SmbServerFqdn pulumi.StringPtrInput `pulumi:"smbServerFqdn"`
	StartIp       pulumi.StringPtrInput `pulumi:"startIp"`
	Subnet        pulumi.StringPtrInput `pulumi:"subnet"`
}

func (MountTargetPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MountTargetPropertiesResponse)(nil)).Elem()
}

func (i MountTargetPropertiesResponseArgs) ToMountTargetPropertiesResponseOutput() MountTargetPropertiesResponseOutput {
	return i.ToMountTargetPropertiesResponseOutputWithContext(context.Background())
}

func (i MountTargetPropertiesResponseArgs) ToMountTargetPropertiesResponseOutputWithContext(ctx context.Context) MountTargetPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountTargetPropertiesResponseOutput)
}





type MountTargetPropertiesResponseArrayInput interface {
	pulumi.Input

	ToMountTargetPropertiesResponseArrayOutput() MountTargetPropertiesResponseArrayOutput
	ToMountTargetPropertiesResponseArrayOutputWithContext(context.Context) MountTargetPropertiesResponseArrayOutput
}

type MountTargetPropertiesResponseArray []MountTargetPropertiesResponseInput

func (MountTargetPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MountTargetPropertiesResponse)(nil)).Elem()
}

func (i MountTargetPropertiesResponseArray) ToMountTargetPropertiesResponseArrayOutput() MountTargetPropertiesResponseArrayOutput {
	return i.ToMountTargetPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i MountTargetPropertiesResponseArray) ToMountTargetPropertiesResponseArrayOutputWithContext(ctx context.Context) MountTargetPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountTargetPropertiesResponseArrayOutput)
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

type ReplicationObject struct {
	EndpointType           *string `pulumi:"endpointType"`
	RemoteVolumeRegion     *string `pulumi:"remoteVolumeRegion"`
	RemoteVolumeResourceId string  `pulumi:"remoteVolumeResourceId"`
	ReplicationId          *string `pulumi:"replicationId"`
	ReplicationSchedule    string  `pulumi:"replicationSchedule"`
}





type ReplicationObjectInput interface {
	pulumi.Input

	ToReplicationObjectOutput() ReplicationObjectOutput
	ToReplicationObjectOutputWithContext(context.Context) ReplicationObjectOutput
}

type ReplicationObjectArgs struct {
	EndpointType           pulumi.StringPtrInput `pulumi:"endpointType"`
	RemoteVolumeRegion     pulumi.StringPtrInput `pulumi:"remoteVolumeRegion"`
	RemoteVolumeResourceId pulumi.StringInput    `pulumi:"remoteVolumeResourceId"`
	ReplicationId          pulumi.StringPtrInput `pulumi:"replicationId"`
	ReplicationSchedule    pulumi.StringInput    `pulumi:"replicationSchedule"`
}

func (ReplicationObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationObject)(nil)).Elem()
}

func (i ReplicationObjectArgs) ToReplicationObjectOutput() ReplicationObjectOutput {
	return i.ToReplicationObjectOutputWithContext(context.Background())
}

func (i ReplicationObjectArgs) ToReplicationObjectOutputWithContext(ctx context.Context) ReplicationObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationObjectOutput)
}

func (i ReplicationObjectArgs) ToReplicationObjectPtrOutput() ReplicationObjectPtrOutput {
	return i.ToReplicationObjectPtrOutputWithContext(context.Background())
}

func (i ReplicationObjectArgs) ToReplicationObjectPtrOutputWithContext(ctx context.Context) ReplicationObjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationObjectOutput).ToReplicationObjectPtrOutputWithContext(ctx)
}









type ReplicationObjectPtrInput interface {
	pulumi.Input

	ToReplicationObjectPtrOutput() ReplicationObjectPtrOutput
	ToReplicationObjectPtrOutputWithContext(context.Context) ReplicationObjectPtrOutput
}

type replicationObjectPtrType ReplicationObjectArgs

func ReplicationObjectPtr(v *ReplicationObjectArgs) ReplicationObjectPtrInput {
	return (*replicationObjectPtrType)(v)
}

func (*replicationObjectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationObject)(nil)).Elem()
}

func (i *replicationObjectPtrType) ToReplicationObjectPtrOutput() ReplicationObjectPtrOutput {
	return i.ToReplicationObjectPtrOutputWithContext(context.Background())
}

func (i *replicationObjectPtrType) ToReplicationObjectPtrOutputWithContext(ctx context.Context) ReplicationObjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationObjectPtrOutput)
}

type ReplicationObjectOutput struct{ *pulumi.OutputState }

func (ReplicationObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationObject)(nil)).Elem()
}

func (o ReplicationObjectOutput) ToReplicationObjectOutput() ReplicationObjectOutput {
	return o
}

func (o ReplicationObjectOutput) ToReplicationObjectOutputWithContext(ctx context.Context) ReplicationObjectOutput {
	return o
}

func (o ReplicationObjectOutput) ToReplicationObjectPtrOutput() ReplicationObjectPtrOutput {
	return o.ToReplicationObjectPtrOutputWithContext(context.Background())
}

func (o ReplicationObjectOutput) ToReplicationObjectPtrOutputWithContext(ctx context.Context) ReplicationObjectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicationObject) *ReplicationObject {
		return &v
	}).(ReplicationObjectPtrOutput)
}

func (o ReplicationObjectOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObject) *string { return v.EndpointType }).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectOutput) RemoteVolumeRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObject) *string { return v.RemoteVolumeRegion }).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectOutput) RemoteVolumeResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicationObject) string { return v.RemoteVolumeResourceId }).(pulumi.StringOutput)
}

func (o ReplicationObjectOutput) ReplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObject) *string { return v.ReplicationId }).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectOutput) ReplicationSchedule() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicationObject) string { return v.ReplicationSchedule }).(pulumi.StringOutput)
}

type ReplicationObjectPtrOutput struct{ *pulumi.OutputState }

func (ReplicationObjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationObject)(nil)).Elem()
}

func (o ReplicationObjectPtrOutput) ToReplicationObjectPtrOutput() ReplicationObjectPtrOutput {
	return o
}

func (o ReplicationObjectPtrOutput) ToReplicationObjectPtrOutputWithContext(ctx context.Context) ReplicationObjectPtrOutput {
	return o
}

func (o ReplicationObjectPtrOutput) Elem() ReplicationObjectOutput {
	return o.ApplyT(func(v *ReplicationObject) ReplicationObject {
		if v != nil {
			return *v
		}
		var ret ReplicationObject
		return ret
	}).(ReplicationObjectOutput)
}

func (o ReplicationObjectPtrOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObject) *string {
		if v == nil {
			return nil
		}
		return v.EndpointType
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectPtrOutput) RemoteVolumeRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObject) *string {
		if v == nil {
			return nil
		}
		return v.RemoteVolumeRegion
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectPtrOutput) RemoteVolumeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObject) *string {
		if v == nil {
			return nil
		}
		return &v.RemoteVolumeResourceId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectPtrOutput) ReplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObject) *string {
		if v == nil {
			return nil
		}
		return v.ReplicationId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectPtrOutput) ReplicationSchedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObject) *string {
		if v == nil {
			return nil
		}
		return &v.ReplicationSchedule
	}).(pulumi.StringPtrOutput)
}

type ReplicationObjectResponse struct {
	EndpointType           *string `pulumi:"endpointType"`
	RemoteVolumeRegion     *string `pulumi:"remoteVolumeRegion"`
	RemoteVolumeResourceId string  `pulumi:"remoteVolumeResourceId"`
	ReplicationId          *string `pulumi:"replicationId"`
	ReplicationSchedule    string  `pulumi:"replicationSchedule"`
}





type ReplicationObjectResponseInput interface {
	pulumi.Input

	ToReplicationObjectResponseOutput() ReplicationObjectResponseOutput
	ToReplicationObjectResponseOutputWithContext(context.Context) ReplicationObjectResponseOutput
}

type ReplicationObjectResponseArgs struct {
	EndpointType           pulumi.StringPtrInput `pulumi:"endpointType"`
	RemoteVolumeRegion     pulumi.StringPtrInput `pulumi:"remoteVolumeRegion"`
	RemoteVolumeResourceId pulumi.StringInput    `pulumi:"remoteVolumeResourceId"`
	ReplicationId          pulumi.StringPtrInput `pulumi:"replicationId"`
	ReplicationSchedule    pulumi.StringInput    `pulumi:"replicationSchedule"`
}

func (ReplicationObjectResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationObjectResponse)(nil)).Elem()
}

func (i ReplicationObjectResponseArgs) ToReplicationObjectResponseOutput() ReplicationObjectResponseOutput {
	return i.ToReplicationObjectResponseOutputWithContext(context.Background())
}

func (i ReplicationObjectResponseArgs) ToReplicationObjectResponseOutputWithContext(ctx context.Context) ReplicationObjectResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationObjectResponseOutput)
}

func (i ReplicationObjectResponseArgs) ToReplicationObjectResponsePtrOutput() ReplicationObjectResponsePtrOutput {
	return i.ToReplicationObjectResponsePtrOutputWithContext(context.Background())
}

func (i ReplicationObjectResponseArgs) ToReplicationObjectResponsePtrOutputWithContext(ctx context.Context) ReplicationObjectResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationObjectResponseOutput).ToReplicationObjectResponsePtrOutputWithContext(ctx)
}









type ReplicationObjectResponsePtrInput interface {
	pulumi.Input

	ToReplicationObjectResponsePtrOutput() ReplicationObjectResponsePtrOutput
	ToReplicationObjectResponsePtrOutputWithContext(context.Context) ReplicationObjectResponsePtrOutput
}

type replicationObjectResponsePtrType ReplicationObjectResponseArgs

func ReplicationObjectResponsePtr(v *ReplicationObjectResponseArgs) ReplicationObjectResponsePtrInput {
	return (*replicationObjectResponsePtrType)(v)
}

func (*replicationObjectResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationObjectResponse)(nil)).Elem()
}

func (i *replicationObjectResponsePtrType) ToReplicationObjectResponsePtrOutput() ReplicationObjectResponsePtrOutput {
	return i.ToReplicationObjectResponsePtrOutputWithContext(context.Background())
}

func (i *replicationObjectResponsePtrType) ToReplicationObjectResponsePtrOutputWithContext(ctx context.Context) ReplicationObjectResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationObjectResponsePtrOutput)
}

type ReplicationObjectResponseOutput struct{ *pulumi.OutputState }

func (ReplicationObjectResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationObjectResponse)(nil)).Elem()
}

func (o ReplicationObjectResponseOutput) ToReplicationObjectResponseOutput() ReplicationObjectResponseOutput {
	return o
}

func (o ReplicationObjectResponseOutput) ToReplicationObjectResponseOutputWithContext(ctx context.Context) ReplicationObjectResponseOutput {
	return o
}

func (o ReplicationObjectResponseOutput) ToReplicationObjectResponsePtrOutput() ReplicationObjectResponsePtrOutput {
	return o.ToReplicationObjectResponsePtrOutputWithContext(context.Background())
}

func (o ReplicationObjectResponseOutput) ToReplicationObjectResponsePtrOutputWithContext(ctx context.Context) ReplicationObjectResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicationObjectResponse) *ReplicationObjectResponse {
		return &v
	}).(ReplicationObjectResponsePtrOutput)
}

func (o ReplicationObjectResponseOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObjectResponse) *string { return v.EndpointType }).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponseOutput) RemoteVolumeRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObjectResponse) *string { return v.RemoteVolumeRegion }).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponseOutput) RemoteVolumeResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicationObjectResponse) string { return v.RemoteVolumeResourceId }).(pulumi.StringOutput)
}

func (o ReplicationObjectResponseOutput) ReplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationObjectResponse) *string { return v.ReplicationId }).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponseOutput) ReplicationSchedule() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicationObjectResponse) string { return v.ReplicationSchedule }).(pulumi.StringOutput)
}

type ReplicationObjectResponsePtrOutput struct{ *pulumi.OutputState }

func (ReplicationObjectResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationObjectResponse)(nil)).Elem()
}

func (o ReplicationObjectResponsePtrOutput) ToReplicationObjectResponsePtrOutput() ReplicationObjectResponsePtrOutput {
	return o
}

func (o ReplicationObjectResponsePtrOutput) ToReplicationObjectResponsePtrOutputWithContext(ctx context.Context) ReplicationObjectResponsePtrOutput {
	return o
}

func (o ReplicationObjectResponsePtrOutput) Elem() ReplicationObjectResponseOutput {
	return o.ApplyT(func(v *ReplicationObjectResponse) ReplicationObjectResponse {
		if v != nil {
			return *v
		}
		var ret ReplicationObjectResponse
		return ret
	}).(ReplicationObjectResponseOutput)
}

func (o ReplicationObjectResponsePtrOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObjectResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndpointType
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponsePtrOutput) RemoteVolumeRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObjectResponse) *string {
		if v == nil {
			return nil
		}
		return v.RemoteVolumeRegion
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponsePtrOutput) RemoteVolumeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObjectResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RemoteVolumeResourceId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponsePtrOutput) ReplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObjectResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReplicationId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationObjectResponsePtrOutput) ReplicationSchedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationObjectResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ReplicationSchedule
	}).(pulumi.StringPtrOutput)
}

type VolumePropertiesDataProtection struct {
	Replication *ReplicationObject        `pulumi:"replication"`
	Snapshot    *VolumeSnapshotProperties `pulumi:"snapshot"`
}





type VolumePropertiesDataProtectionInput interface {
	pulumi.Input

	ToVolumePropertiesDataProtectionOutput() VolumePropertiesDataProtectionOutput
	ToVolumePropertiesDataProtectionOutputWithContext(context.Context) VolumePropertiesDataProtectionOutput
}

type VolumePropertiesDataProtectionArgs struct {
	Replication ReplicationObjectPtrInput        `pulumi:"replication"`
	Snapshot    VolumeSnapshotPropertiesPtrInput `pulumi:"snapshot"`
}

func (VolumePropertiesDataProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesDataProtection)(nil)).Elem()
}

func (i VolumePropertiesDataProtectionArgs) ToVolumePropertiesDataProtectionOutput() VolumePropertiesDataProtectionOutput {
	return i.ToVolumePropertiesDataProtectionOutputWithContext(context.Background())
}

func (i VolumePropertiesDataProtectionArgs) ToVolumePropertiesDataProtectionOutputWithContext(ctx context.Context) VolumePropertiesDataProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesDataProtectionOutput)
}

func (i VolumePropertiesDataProtectionArgs) ToVolumePropertiesDataProtectionPtrOutput() VolumePropertiesDataProtectionPtrOutput {
	return i.ToVolumePropertiesDataProtectionPtrOutputWithContext(context.Background())
}

func (i VolumePropertiesDataProtectionArgs) ToVolumePropertiesDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesDataProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesDataProtectionOutput).ToVolumePropertiesDataProtectionPtrOutputWithContext(ctx)
}









type VolumePropertiesDataProtectionPtrInput interface {
	pulumi.Input

	ToVolumePropertiesDataProtectionPtrOutput() VolumePropertiesDataProtectionPtrOutput
	ToVolumePropertiesDataProtectionPtrOutputWithContext(context.Context) VolumePropertiesDataProtectionPtrOutput
}

type volumePropertiesDataProtectionPtrType VolumePropertiesDataProtectionArgs

func VolumePropertiesDataProtectionPtr(v *VolumePropertiesDataProtectionArgs) VolumePropertiesDataProtectionPtrInput {
	return (*volumePropertiesDataProtectionPtrType)(v)
}

func (*volumePropertiesDataProtectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesDataProtection)(nil)).Elem()
}

func (i *volumePropertiesDataProtectionPtrType) ToVolumePropertiesDataProtectionPtrOutput() VolumePropertiesDataProtectionPtrOutput {
	return i.ToVolumePropertiesDataProtectionPtrOutputWithContext(context.Background())
}

func (i *volumePropertiesDataProtectionPtrType) ToVolumePropertiesDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesDataProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesDataProtectionPtrOutput)
}

type VolumePropertiesDataProtectionOutput struct{ *pulumi.OutputState }

func (VolumePropertiesDataProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesDataProtection)(nil)).Elem()
}

func (o VolumePropertiesDataProtectionOutput) ToVolumePropertiesDataProtectionOutput() VolumePropertiesDataProtectionOutput {
	return o
}

func (o VolumePropertiesDataProtectionOutput) ToVolumePropertiesDataProtectionOutputWithContext(ctx context.Context) VolumePropertiesDataProtectionOutput {
	return o
}

func (o VolumePropertiesDataProtectionOutput) ToVolumePropertiesDataProtectionPtrOutput() VolumePropertiesDataProtectionPtrOutput {
	return o.ToVolumePropertiesDataProtectionPtrOutputWithContext(context.Background())
}

func (o VolumePropertiesDataProtectionOutput) ToVolumePropertiesDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesDataProtectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumePropertiesDataProtection) *VolumePropertiesDataProtection {
		return &v
	}).(VolumePropertiesDataProtectionPtrOutput)
}

func (o VolumePropertiesDataProtectionOutput) Replication() ReplicationObjectPtrOutput {
	return o.ApplyT(func(v VolumePropertiesDataProtection) *ReplicationObject { return v.Replication }).(ReplicationObjectPtrOutput)
}

func (o VolumePropertiesDataProtectionOutput) Snapshot() VolumeSnapshotPropertiesPtrOutput {
	return o.ApplyT(func(v VolumePropertiesDataProtection) *VolumeSnapshotProperties { return v.Snapshot }).(VolumeSnapshotPropertiesPtrOutput)
}

type VolumePropertiesDataProtectionPtrOutput struct{ *pulumi.OutputState }

func (VolumePropertiesDataProtectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesDataProtection)(nil)).Elem()
}

func (o VolumePropertiesDataProtectionPtrOutput) ToVolumePropertiesDataProtectionPtrOutput() VolumePropertiesDataProtectionPtrOutput {
	return o
}

func (o VolumePropertiesDataProtectionPtrOutput) ToVolumePropertiesDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesDataProtectionPtrOutput {
	return o
}

func (o VolumePropertiesDataProtectionPtrOutput) Elem() VolumePropertiesDataProtectionOutput {
	return o.ApplyT(func(v *VolumePropertiesDataProtection) VolumePropertiesDataProtection {
		if v != nil {
			return *v
		}
		var ret VolumePropertiesDataProtection
		return ret
	}).(VolumePropertiesDataProtectionOutput)
}

func (o VolumePropertiesDataProtectionPtrOutput) Replication() ReplicationObjectPtrOutput {
	return o.ApplyT(func(v *VolumePropertiesDataProtection) *ReplicationObject {
		if v == nil {
			return nil
		}
		return v.Replication
	}).(ReplicationObjectPtrOutput)
}

func (o VolumePropertiesDataProtectionPtrOutput) Snapshot() VolumeSnapshotPropertiesPtrOutput {
	return o.ApplyT(func(v *VolumePropertiesDataProtection) *VolumeSnapshotProperties {
		if v == nil {
			return nil
		}
		return v.Snapshot
	}).(VolumeSnapshotPropertiesPtrOutput)
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

type VolumePropertiesResponseDataProtection struct {
	Replication *ReplicationObjectResponse        `pulumi:"replication"`
	Snapshot    *VolumeSnapshotPropertiesResponse `pulumi:"snapshot"`
}





type VolumePropertiesResponseDataProtectionInput interface {
	pulumi.Input

	ToVolumePropertiesResponseDataProtectionOutput() VolumePropertiesResponseDataProtectionOutput
	ToVolumePropertiesResponseDataProtectionOutputWithContext(context.Context) VolumePropertiesResponseDataProtectionOutput
}

type VolumePropertiesResponseDataProtectionArgs struct {
	Replication ReplicationObjectResponsePtrInput        `pulumi:"replication"`
	Snapshot    VolumeSnapshotPropertiesResponsePtrInput `pulumi:"snapshot"`
}

func (VolumePropertiesResponseDataProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesResponseDataProtection)(nil)).Elem()
}

func (i VolumePropertiesResponseDataProtectionArgs) ToVolumePropertiesResponseDataProtectionOutput() VolumePropertiesResponseDataProtectionOutput {
	return i.ToVolumePropertiesResponseDataProtectionOutputWithContext(context.Background())
}

func (i VolumePropertiesResponseDataProtectionArgs) ToVolumePropertiesResponseDataProtectionOutputWithContext(ctx context.Context) VolumePropertiesResponseDataProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesResponseDataProtectionOutput)
}

func (i VolumePropertiesResponseDataProtectionArgs) ToVolumePropertiesResponseDataProtectionPtrOutput() VolumePropertiesResponseDataProtectionPtrOutput {
	return i.ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(context.Background())
}

func (i VolumePropertiesResponseDataProtectionArgs) ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseDataProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesResponseDataProtectionOutput).ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(ctx)
}









type VolumePropertiesResponseDataProtectionPtrInput interface {
	pulumi.Input

	ToVolumePropertiesResponseDataProtectionPtrOutput() VolumePropertiesResponseDataProtectionPtrOutput
	ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(context.Context) VolumePropertiesResponseDataProtectionPtrOutput
}

type volumePropertiesResponseDataProtectionPtrType VolumePropertiesResponseDataProtectionArgs

func VolumePropertiesResponseDataProtectionPtr(v *VolumePropertiesResponseDataProtectionArgs) VolumePropertiesResponseDataProtectionPtrInput {
	return (*volumePropertiesResponseDataProtectionPtrType)(v)
}

func (*volumePropertiesResponseDataProtectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesResponseDataProtection)(nil)).Elem()
}

func (i *volumePropertiesResponseDataProtectionPtrType) ToVolumePropertiesResponseDataProtectionPtrOutput() VolumePropertiesResponseDataProtectionPtrOutput {
	return i.ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(context.Background())
}

func (i *volumePropertiesResponseDataProtectionPtrType) ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseDataProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesResponseDataProtectionPtrOutput)
}

type VolumePropertiesResponseDataProtectionOutput struct{ *pulumi.OutputState }

func (VolumePropertiesResponseDataProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesResponseDataProtection)(nil)).Elem()
}

func (o VolumePropertiesResponseDataProtectionOutput) ToVolumePropertiesResponseDataProtectionOutput() VolumePropertiesResponseDataProtectionOutput {
	return o
}

func (o VolumePropertiesResponseDataProtectionOutput) ToVolumePropertiesResponseDataProtectionOutputWithContext(ctx context.Context) VolumePropertiesResponseDataProtectionOutput {
	return o
}

func (o VolumePropertiesResponseDataProtectionOutput) ToVolumePropertiesResponseDataProtectionPtrOutput() VolumePropertiesResponseDataProtectionPtrOutput {
	return o.ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(context.Background())
}

func (o VolumePropertiesResponseDataProtectionOutput) ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseDataProtectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumePropertiesResponseDataProtection) *VolumePropertiesResponseDataProtection {
		return &v
	}).(VolumePropertiesResponseDataProtectionPtrOutput)
}

func (o VolumePropertiesResponseDataProtectionOutput) Replication() ReplicationObjectResponsePtrOutput {
	return o.ApplyT(func(v VolumePropertiesResponseDataProtection) *ReplicationObjectResponse { return v.Replication }).(ReplicationObjectResponsePtrOutput)
}

func (o VolumePropertiesResponseDataProtectionOutput) Snapshot() VolumeSnapshotPropertiesResponsePtrOutput {
	return o.ApplyT(func(v VolumePropertiesResponseDataProtection) *VolumeSnapshotPropertiesResponse { return v.Snapshot }).(VolumeSnapshotPropertiesResponsePtrOutput)
}

type VolumePropertiesResponseDataProtectionPtrOutput struct{ *pulumi.OutputState }

func (VolumePropertiesResponseDataProtectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesResponseDataProtection)(nil)).Elem()
}

func (o VolumePropertiesResponseDataProtectionPtrOutput) ToVolumePropertiesResponseDataProtectionPtrOutput() VolumePropertiesResponseDataProtectionPtrOutput {
	return o
}

func (o VolumePropertiesResponseDataProtectionPtrOutput) ToVolumePropertiesResponseDataProtectionPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseDataProtectionPtrOutput {
	return o
}

func (o VolumePropertiesResponseDataProtectionPtrOutput) Elem() VolumePropertiesResponseDataProtectionOutput {
	return o.ApplyT(func(v *VolumePropertiesResponseDataProtection) VolumePropertiesResponseDataProtection {
		if v != nil {
			return *v
		}
		var ret VolumePropertiesResponseDataProtection
		return ret
	}).(VolumePropertiesResponseDataProtectionOutput)
}

func (o VolumePropertiesResponseDataProtectionPtrOutput) Replication() ReplicationObjectResponsePtrOutput {
	return o.ApplyT(func(v *VolumePropertiesResponseDataProtection) *ReplicationObjectResponse {
		if v == nil {
			return nil
		}
		return v.Replication
	}).(ReplicationObjectResponsePtrOutput)
}

func (o VolumePropertiesResponseDataProtectionPtrOutput) Snapshot() VolumeSnapshotPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *VolumePropertiesResponseDataProtection) *VolumeSnapshotPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Snapshot
	}).(VolumeSnapshotPropertiesResponsePtrOutput)
}

type VolumePropertiesResponseExportPolicy struct {
	Rules []ExportPolicyRuleResponse `pulumi:"rules"`
}





type VolumePropertiesResponseExportPolicyInput interface {
	pulumi.Input

	ToVolumePropertiesResponseExportPolicyOutput() VolumePropertiesResponseExportPolicyOutput
	ToVolumePropertiesResponseExportPolicyOutputWithContext(context.Context) VolumePropertiesResponseExportPolicyOutput
}

type VolumePropertiesResponseExportPolicyArgs struct {
	Rules ExportPolicyRuleResponseArrayInput `pulumi:"rules"`
}

func (VolumePropertiesResponseExportPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumePropertiesResponseExportPolicy)(nil)).Elem()
}

func (i VolumePropertiesResponseExportPolicyArgs) ToVolumePropertiesResponseExportPolicyOutput() VolumePropertiesResponseExportPolicyOutput {
	return i.ToVolumePropertiesResponseExportPolicyOutputWithContext(context.Background())
}

func (i VolumePropertiesResponseExportPolicyArgs) ToVolumePropertiesResponseExportPolicyOutputWithContext(ctx context.Context) VolumePropertiesResponseExportPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesResponseExportPolicyOutput)
}

func (i VolumePropertiesResponseExportPolicyArgs) ToVolumePropertiesResponseExportPolicyPtrOutput() VolumePropertiesResponseExportPolicyPtrOutput {
	return i.ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(context.Background())
}

func (i VolumePropertiesResponseExportPolicyArgs) ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseExportPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesResponseExportPolicyOutput).ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(ctx)
}









type VolumePropertiesResponseExportPolicyPtrInput interface {
	pulumi.Input

	ToVolumePropertiesResponseExportPolicyPtrOutput() VolumePropertiesResponseExportPolicyPtrOutput
	ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(context.Context) VolumePropertiesResponseExportPolicyPtrOutput
}

type volumePropertiesResponseExportPolicyPtrType VolumePropertiesResponseExportPolicyArgs

func VolumePropertiesResponseExportPolicyPtr(v *VolumePropertiesResponseExportPolicyArgs) VolumePropertiesResponseExportPolicyPtrInput {
	return (*volumePropertiesResponseExportPolicyPtrType)(v)
}

func (*volumePropertiesResponseExportPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumePropertiesResponseExportPolicy)(nil)).Elem()
}

func (i *volumePropertiesResponseExportPolicyPtrType) ToVolumePropertiesResponseExportPolicyPtrOutput() VolumePropertiesResponseExportPolicyPtrOutput {
	return i.ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(context.Background())
}

func (i *volumePropertiesResponseExportPolicyPtrType) ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseExportPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumePropertiesResponseExportPolicyPtrOutput)
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

func (o VolumePropertiesResponseExportPolicyOutput) ToVolumePropertiesResponseExportPolicyPtrOutput() VolumePropertiesResponseExportPolicyPtrOutput {
	return o.ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(context.Background())
}

func (o VolumePropertiesResponseExportPolicyOutput) ToVolumePropertiesResponseExportPolicyPtrOutputWithContext(ctx context.Context) VolumePropertiesResponseExportPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumePropertiesResponseExportPolicy) *VolumePropertiesResponseExportPolicy {
		return &v
	}).(VolumePropertiesResponseExportPolicyPtrOutput)
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

type VolumeSnapshotProperties struct {
	SnapshotPolicyId *string `pulumi:"snapshotPolicyId"`
}





type VolumeSnapshotPropertiesInput interface {
	pulumi.Input

	ToVolumeSnapshotPropertiesOutput() VolumeSnapshotPropertiesOutput
	ToVolumeSnapshotPropertiesOutputWithContext(context.Context) VolumeSnapshotPropertiesOutput
}

type VolumeSnapshotPropertiesArgs struct {
	SnapshotPolicyId pulumi.StringPtrInput `pulumi:"snapshotPolicyId"`
}

func (VolumeSnapshotPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeSnapshotProperties)(nil)).Elem()
}

func (i VolumeSnapshotPropertiesArgs) ToVolumeSnapshotPropertiesOutput() VolumeSnapshotPropertiesOutput {
	return i.ToVolumeSnapshotPropertiesOutputWithContext(context.Background())
}

func (i VolumeSnapshotPropertiesArgs) ToVolumeSnapshotPropertiesOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPropertiesOutput)
}

func (i VolumeSnapshotPropertiesArgs) ToVolumeSnapshotPropertiesPtrOutput() VolumeSnapshotPropertiesPtrOutput {
	return i.ToVolumeSnapshotPropertiesPtrOutputWithContext(context.Background())
}

func (i VolumeSnapshotPropertiesArgs) ToVolumeSnapshotPropertiesPtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPropertiesOutput).ToVolumeSnapshotPropertiesPtrOutputWithContext(ctx)
}









type VolumeSnapshotPropertiesPtrInput interface {
	pulumi.Input

	ToVolumeSnapshotPropertiesPtrOutput() VolumeSnapshotPropertiesPtrOutput
	ToVolumeSnapshotPropertiesPtrOutputWithContext(context.Context) VolumeSnapshotPropertiesPtrOutput
}

type volumeSnapshotPropertiesPtrType VolumeSnapshotPropertiesArgs

func VolumeSnapshotPropertiesPtr(v *VolumeSnapshotPropertiesArgs) VolumeSnapshotPropertiesPtrInput {
	return (*volumeSnapshotPropertiesPtrType)(v)
}

func (*volumeSnapshotPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeSnapshotProperties)(nil)).Elem()
}

func (i *volumeSnapshotPropertiesPtrType) ToVolumeSnapshotPropertiesPtrOutput() VolumeSnapshotPropertiesPtrOutput {
	return i.ToVolumeSnapshotPropertiesPtrOutputWithContext(context.Background())
}

func (i *volumeSnapshotPropertiesPtrType) ToVolumeSnapshotPropertiesPtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPropertiesPtrOutput)
}

type VolumeSnapshotPropertiesOutput struct{ *pulumi.OutputState }

func (VolumeSnapshotPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeSnapshotProperties)(nil)).Elem()
}

func (o VolumeSnapshotPropertiesOutput) ToVolumeSnapshotPropertiesOutput() VolumeSnapshotPropertiesOutput {
	return o
}

func (o VolumeSnapshotPropertiesOutput) ToVolumeSnapshotPropertiesOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesOutput {
	return o
}

func (o VolumeSnapshotPropertiesOutput) ToVolumeSnapshotPropertiesPtrOutput() VolumeSnapshotPropertiesPtrOutput {
	return o.ToVolumeSnapshotPropertiesPtrOutputWithContext(context.Background())
}

func (o VolumeSnapshotPropertiesOutput) ToVolumeSnapshotPropertiesPtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeSnapshotProperties) *VolumeSnapshotProperties {
		return &v
	}).(VolumeSnapshotPropertiesPtrOutput)
}

func (o VolumeSnapshotPropertiesOutput) SnapshotPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeSnapshotProperties) *string { return v.SnapshotPolicyId }).(pulumi.StringPtrOutput)
}

type VolumeSnapshotPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VolumeSnapshotPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeSnapshotProperties)(nil)).Elem()
}

func (o VolumeSnapshotPropertiesPtrOutput) ToVolumeSnapshotPropertiesPtrOutput() VolumeSnapshotPropertiesPtrOutput {
	return o
}

func (o VolumeSnapshotPropertiesPtrOutput) ToVolumeSnapshotPropertiesPtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesPtrOutput {
	return o
}

func (o VolumeSnapshotPropertiesPtrOutput) Elem() VolumeSnapshotPropertiesOutput {
	return o.ApplyT(func(v *VolumeSnapshotProperties) VolumeSnapshotProperties {
		if v != nil {
			return *v
		}
		var ret VolumeSnapshotProperties
		return ret
	}).(VolumeSnapshotPropertiesOutput)
}

func (o VolumeSnapshotPropertiesPtrOutput) SnapshotPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeSnapshotProperties) *string {
		if v == nil {
			return nil
		}
		return v.SnapshotPolicyId
	}).(pulumi.StringPtrOutput)
}

type VolumeSnapshotPropertiesResponse struct {
	SnapshotPolicyId *string `pulumi:"snapshotPolicyId"`
}





type VolumeSnapshotPropertiesResponseInput interface {
	pulumi.Input

	ToVolumeSnapshotPropertiesResponseOutput() VolumeSnapshotPropertiesResponseOutput
	ToVolumeSnapshotPropertiesResponseOutputWithContext(context.Context) VolumeSnapshotPropertiesResponseOutput
}

type VolumeSnapshotPropertiesResponseArgs struct {
	SnapshotPolicyId pulumi.StringPtrInput `pulumi:"snapshotPolicyId"`
}

func (VolumeSnapshotPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeSnapshotPropertiesResponse)(nil)).Elem()
}

func (i VolumeSnapshotPropertiesResponseArgs) ToVolumeSnapshotPropertiesResponseOutput() VolumeSnapshotPropertiesResponseOutput {
	return i.ToVolumeSnapshotPropertiesResponseOutputWithContext(context.Background())
}

func (i VolumeSnapshotPropertiesResponseArgs) ToVolumeSnapshotPropertiesResponseOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPropertiesResponseOutput)
}

func (i VolumeSnapshotPropertiesResponseArgs) ToVolumeSnapshotPropertiesResponsePtrOutput() VolumeSnapshotPropertiesResponsePtrOutput {
	return i.ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i VolumeSnapshotPropertiesResponseArgs) ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPropertiesResponseOutput).ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(ctx)
}









type VolumeSnapshotPropertiesResponsePtrInput interface {
	pulumi.Input

	ToVolumeSnapshotPropertiesResponsePtrOutput() VolumeSnapshotPropertiesResponsePtrOutput
	ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(context.Context) VolumeSnapshotPropertiesResponsePtrOutput
}

type volumeSnapshotPropertiesResponsePtrType VolumeSnapshotPropertiesResponseArgs

func VolumeSnapshotPropertiesResponsePtr(v *VolumeSnapshotPropertiesResponseArgs) VolumeSnapshotPropertiesResponsePtrInput {
	return (*volumeSnapshotPropertiesResponsePtrType)(v)
}

func (*volumeSnapshotPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeSnapshotPropertiesResponse)(nil)).Elem()
}

func (i *volumeSnapshotPropertiesResponsePtrType) ToVolumeSnapshotPropertiesResponsePtrOutput() VolumeSnapshotPropertiesResponsePtrOutput {
	return i.ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *volumeSnapshotPropertiesResponsePtrType) ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPropertiesResponsePtrOutput)
}

type VolumeSnapshotPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VolumeSnapshotPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeSnapshotPropertiesResponse)(nil)).Elem()
}

func (o VolumeSnapshotPropertiesResponseOutput) ToVolumeSnapshotPropertiesResponseOutput() VolumeSnapshotPropertiesResponseOutput {
	return o
}

func (o VolumeSnapshotPropertiesResponseOutput) ToVolumeSnapshotPropertiesResponseOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesResponseOutput {
	return o
}

func (o VolumeSnapshotPropertiesResponseOutput) ToVolumeSnapshotPropertiesResponsePtrOutput() VolumeSnapshotPropertiesResponsePtrOutput {
	return o.ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o VolumeSnapshotPropertiesResponseOutput) ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeSnapshotPropertiesResponse) *VolumeSnapshotPropertiesResponse {
		return &v
	}).(VolumeSnapshotPropertiesResponsePtrOutput)
}

func (o VolumeSnapshotPropertiesResponseOutput) SnapshotPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeSnapshotPropertiesResponse) *string { return v.SnapshotPolicyId }).(pulumi.StringPtrOutput)
}

type VolumeSnapshotPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VolumeSnapshotPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeSnapshotPropertiesResponse)(nil)).Elem()
}

func (o VolumeSnapshotPropertiesResponsePtrOutput) ToVolumeSnapshotPropertiesResponsePtrOutput() VolumeSnapshotPropertiesResponsePtrOutput {
	return o
}

func (o VolumeSnapshotPropertiesResponsePtrOutput) ToVolumeSnapshotPropertiesResponsePtrOutputWithContext(ctx context.Context) VolumeSnapshotPropertiesResponsePtrOutput {
	return o
}

func (o VolumeSnapshotPropertiesResponsePtrOutput) Elem() VolumeSnapshotPropertiesResponseOutput {
	return o.ApplyT(func(v *VolumeSnapshotPropertiesResponse) VolumeSnapshotPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VolumeSnapshotPropertiesResponse
		return ret
	}).(VolumeSnapshotPropertiesResponseOutput)
}

func (o VolumeSnapshotPropertiesResponsePtrOutput) SnapshotPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeSnapshotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SnapshotPolicyId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActiveDirectoryInput)(nil)).Elem(), ActiveDirectoryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActiveDirectoryArrayInput)(nil)).Elem(), ActiveDirectoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActiveDirectoryResponseInput)(nil)).Elem(), ActiveDirectoryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActiveDirectoryResponseArrayInput)(nil)).Elem(), ActiveDirectoryResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExportPolicyRuleInput)(nil)).Elem(), ExportPolicyRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExportPolicyRuleArrayInput)(nil)).Elem(), ExportPolicyRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExportPolicyRuleResponseInput)(nil)).Elem(), ExportPolicyRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExportPolicyRuleResponseArrayInput)(nil)).Elem(), ExportPolicyRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MountTargetPropertiesResponseInput)(nil)).Elem(), MountTargetPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MountTargetPropertiesResponseArrayInput)(nil)).Elem(), MountTargetPropertiesResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationObjectInput)(nil)).Elem(), ReplicationObjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationObjectPtrInput)(nil)).Elem(), ReplicationObjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationObjectResponseInput)(nil)).Elem(), ReplicationObjectResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationObjectResponsePtrInput)(nil)).Elem(), ReplicationObjectResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumePropertiesDataProtectionInput)(nil)).Elem(), VolumePropertiesDataProtectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumePropertiesDataProtectionPtrInput)(nil)).Elem(), VolumePropertiesDataProtectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumePropertiesExportPolicyInput)(nil)).Elem(), VolumePropertiesExportPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumePropertiesExportPolicyPtrInput)(nil)).Elem(), VolumePropertiesExportPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumePropertiesResponseDataProtectionInput)(nil)).Elem(), VolumePropertiesResponseDataProtectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumePropertiesResponseDataProtectionPtrInput)(nil)).Elem(), VolumePropertiesResponseDataProtectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumePropertiesResponseExportPolicyInput)(nil)).Elem(), VolumePropertiesResponseExportPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumePropertiesResponseExportPolicyPtrInput)(nil)).Elem(), VolumePropertiesResponseExportPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeSnapshotPropertiesInput)(nil)).Elem(), VolumeSnapshotPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeSnapshotPropertiesPtrInput)(nil)).Elem(), VolumeSnapshotPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeSnapshotPropertiesResponseInput)(nil)).Elem(), VolumeSnapshotPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeSnapshotPropertiesResponsePtrInput)(nil)).Elem(), VolumeSnapshotPropertiesResponseArgs{})
	pulumi.RegisterOutputType(ActiveDirectoryOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryArrayOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryResponseOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryResponseArrayOutput{})
	pulumi.RegisterOutputType(ExportPolicyRuleOutput{})
	pulumi.RegisterOutputType(ExportPolicyRuleArrayOutput{})
	pulumi.RegisterOutputType(ExportPolicyRuleResponseOutput{})
	pulumi.RegisterOutputType(ExportPolicyRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(MountTargetPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MountTargetPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(ReplicationObjectOutput{})
	pulumi.RegisterOutputType(ReplicationObjectPtrOutput{})
	pulumi.RegisterOutputType(ReplicationObjectResponseOutput{})
	pulumi.RegisterOutputType(ReplicationObjectResponsePtrOutput{})
	pulumi.RegisterOutputType(VolumePropertiesDataProtectionOutput{})
	pulumi.RegisterOutputType(VolumePropertiesDataProtectionPtrOutput{})
	pulumi.RegisterOutputType(VolumePropertiesExportPolicyOutput{})
	pulumi.RegisterOutputType(VolumePropertiesExportPolicyPtrOutput{})
	pulumi.RegisterOutputType(VolumePropertiesResponseDataProtectionOutput{})
	pulumi.RegisterOutputType(VolumePropertiesResponseDataProtectionPtrOutput{})
	pulumi.RegisterOutputType(VolumePropertiesResponseExportPolicyOutput{})
	pulumi.RegisterOutputType(VolumePropertiesResponseExportPolicyPtrOutput{})
	pulumi.RegisterOutputType(VolumeSnapshotPropertiesOutput{})
	pulumi.RegisterOutputType(VolumeSnapshotPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VolumeSnapshotPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VolumeSnapshotPropertiesResponsePtrOutput{})
}
