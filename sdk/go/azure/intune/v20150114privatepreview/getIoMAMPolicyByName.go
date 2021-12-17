


package v20150114privatepreview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIoMAMPolicyByName(ctx *pulumi.Context, args *LookupIoMAMPolicyByNameArgs, opts ...pulumi.InvokeOption) (*LookupIoMAMPolicyByNameResult, error) {
	var rv LookupIoMAMPolicyByNameResult
	err := ctx.Invoke("azure-native:intune/v20150114privatepreview:getIoMAMPolicyByName", args, &rv, opts...)
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
