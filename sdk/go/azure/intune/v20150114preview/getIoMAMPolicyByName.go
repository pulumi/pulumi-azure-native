


package v20150114preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIoMAMPolicyByName(ctx *pulumi.Context, args *LookupIoMAMPolicyByNameArgs, opts ...pulumi.InvokeOption) (*LookupIoMAMPolicyByNameResult, error) {
	var rv LookupIoMAMPolicyByNameResult
	err := ctx.Invoke("azure-native:intune/v20150114preview:getIoMAMPolicyByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupIoMAMPolicyByNameArgs struct {
	HostName   string  `pulumi:"hostName"`
	PolicyName string  `pulumi:"policyName"`
	Select     *string `pulumi:"select"`
}


type LookupIoMAMPolicyByNameResult struct {
	AccessRecheckOfflineTimeout *string           `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  *string           `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         *string           `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           *string           `pulumi:"appSharingToLevel"`
	Authentication              *string           `pulumi:"authentication"`
	ClipboardSharingLevel       *string           `pulumi:"clipboardSharingLevel"`
	DataBackup                  *string           `pulumi:"dataBackup"`
	Description                 *string           `pulumi:"description"`
	DeviceCompliance            *string           `pulumi:"deviceCompliance"`
	FileEncryptionLevel         *string           `pulumi:"fileEncryptionLevel"`
	FileSharingSaveAs           *string           `pulumi:"fileSharingSaveAs"`
	FriendlyName                string            `pulumi:"friendlyName"`
	GroupStatus                 string            `pulumi:"groupStatus"`
	Id                          string            `pulumi:"id"`
	LastModifiedTime            string            `pulumi:"lastModifiedTime"`
	Location                    *string           `pulumi:"location"`
	ManagedBrowser              *string           `pulumi:"managedBrowser"`
	Name                        string            `pulumi:"name"`
	NumOfApps                   int               `pulumi:"numOfApps"`
	OfflineWipeTimeout          *string           `pulumi:"offlineWipeTimeout"`
	Pin                         *string           `pulumi:"pin"`
	PinNumRetry                 *int              `pulumi:"pinNumRetry"`
	Tags                        map[string]string `pulumi:"tags"`
	TouchId                     *string           `pulumi:"touchId"`
	Type                        string            `pulumi:"type"`
}


func (val *LookupIoMAMPolicyByNameResult) Defaults() *LookupIoMAMPolicyByNameResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AppSharingFromLevel) {
		appSharingFromLevel_ := "none"
		tmp.AppSharingFromLevel = &appSharingFromLevel_
	}
	if isZero(tmp.AppSharingToLevel) {
		appSharingToLevel_ := "none"
		tmp.AppSharingToLevel = &appSharingToLevel_
	}
	if isZero(tmp.Authentication) {
		authentication_ := "required"
		tmp.Authentication = &authentication_
	}
	if isZero(tmp.ClipboardSharingLevel) {
		clipboardSharingLevel_ := "blocked"
		tmp.ClipboardSharingLevel = &clipboardSharingLevel_
	}
	if isZero(tmp.DataBackup) {
		dataBackup_ := "allow"
		tmp.DataBackup = &dataBackup_
	}
	if isZero(tmp.DeviceCompliance) {
		deviceCompliance_ := "enable"
		tmp.DeviceCompliance = &deviceCompliance_
	}
	if isZero(tmp.FileEncryptionLevel) {
		fileEncryptionLevel_ := "deviceLocked"
		tmp.FileEncryptionLevel = &fileEncryptionLevel_
	}
	if isZero(tmp.FileSharingSaveAs) {
		fileSharingSaveAs_ := "allow"
		tmp.FileSharingSaveAs = &fileSharingSaveAs_
	}
	if isZero(tmp.GroupStatus) {
		tmp.GroupStatus = "notTargeted"
	}
	if isZero(tmp.ManagedBrowser) {
		managedBrowser_ := "required"
		tmp.ManagedBrowser = &managedBrowser_
	}
	if isZero(tmp.Pin) {
		pin_ := "required"
		tmp.Pin = &pin_
	}
	if isZero(tmp.TouchId) {
		touchId_ := "enable"
		tmp.TouchId = &touchId_
	}
	return &tmp
}

func LookupIoMAMPolicyByNameOutput(ctx *pulumi.Context, args LookupIoMAMPolicyByNameOutputArgs, opts ...pulumi.InvokeOption) LookupIoMAMPolicyByNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIoMAMPolicyByNameResult, error) {
			args := v.(LookupIoMAMPolicyByNameArgs)
			r, err := LookupIoMAMPolicyByName(ctx, &args, opts...)
			var s LookupIoMAMPolicyByNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIoMAMPolicyByNameResultOutput)
}

type LookupIoMAMPolicyByNameOutputArgs struct {
	HostName   pulumi.StringInput    `pulumi:"hostName"`
	PolicyName pulumi.StringInput    `pulumi:"policyName"`
	Select     pulumi.StringPtrInput `pulumi:"select"`
}

func (LookupIoMAMPolicyByNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoMAMPolicyByNameArgs)(nil)).Elem()
}


type LookupIoMAMPolicyByNameResultOutput struct{ *pulumi.OutputState }

func (LookupIoMAMPolicyByNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoMAMPolicyByNameResult)(nil)).Elem()
}

func (o LookupIoMAMPolicyByNameResultOutput) ToLookupIoMAMPolicyByNameResultOutput() LookupIoMAMPolicyByNameResultOutput {
	return o
}

func (o LookupIoMAMPolicyByNameResultOutput) ToLookupIoMAMPolicyByNameResultOutputWithContext(ctx context.Context) LookupIoMAMPolicyByNameResultOutput {
	return o
}

func (o LookupIoMAMPolicyByNameResultOutput) AccessRecheckOfflineTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.AccessRecheckOfflineTimeout }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) AccessRecheckOnlineTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.AccessRecheckOnlineTimeout }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) AppSharingFromLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.AppSharingFromLevel }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) AppSharingToLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.AppSharingToLevel }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) Authentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.Authentication }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) ClipboardSharingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.ClipboardSharingLevel }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) DataBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.DataBackup }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) DeviceCompliance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.DeviceCompliance }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) FileEncryptionLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.FileEncryptionLevel }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) FileSharingSaveAs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.FileSharingSaveAs }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) string { return v.FriendlyName }).(pulumi.StringOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) GroupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) string { return v.GroupStatus }).(pulumi.StringOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) ManagedBrowser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.ManagedBrowser }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) NumOfApps() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) int { return v.NumOfApps }).(pulumi.IntOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) OfflineWipeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.OfflineWipeTimeout }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) Pin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.Pin }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) PinNumRetry() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *int { return v.PinNumRetry }).(pulumi.IntPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) TouchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) *string { return v.TouchId }).(pulumi.StringPtrOutput)
}

func (o LookupIoMAMPolicyByNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoMAMPolicyByNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIoMAMPolicyByNameResultOutput{})
}
