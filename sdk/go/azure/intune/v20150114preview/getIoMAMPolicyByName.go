


package v20150114preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIoMAMPolicyByName(ctx *pulumi.Context, args *LookupIoMAMPolicyByNameArgs, opts ...pulumi.InvokeOption) (*LookupIoMAMPolicyByNameResult, error) {
	var rv LookupIoMAMPolicyByNameResult
	err := ctx.Invoke("azure-native:intune/v20150114preview:getIoMAMPolicyByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
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
