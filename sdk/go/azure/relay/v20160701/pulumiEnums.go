


package v20160701

type AccessRights string

const (
	AccessRightsManage = AccessRights("Manage")
	AccessRightsSend   = AccessRights("Send")
	AccessRightsListen = AccessRights("Listen")
)

type Relaytype string

const (
	RelaytypeNetTcp = Relaytype("NetTcp")
	RelaytypeHttp   = Relaytype("Http")
)

type SkuName string

const (
	SkuNameStandard = SkuName("Standard")
)

type SkuTier string

const (
	SkuTierStandard = SkuTier("Standard")
)

func init() {
}
