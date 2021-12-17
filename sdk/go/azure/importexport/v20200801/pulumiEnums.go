


package v20200801

type DriveState string

const (
	DriveStateSpecified         = DriveState("Specified")
	DriveStateReceived          = DriveState("Received")
	DriveStateNeverReceived     = DriveState("NeverReceived")
	DriveStateTransferring      = DriveState("Transferring")
	DriveStateCompleted         = DriveState("Completed")
	DriveStateCompletedMoreInfo = DriveState("CompletedMoreInfo")
	DriveStateShippedBack       = DriveState("ShippedBack")
)

type EncryptionKekType string

const (
	EncryptionKekTypeMicrosoftManaged = EncryptionKekType("MicrosoftManaged")
	EncryptionKekTypeCustomerManaged  = EncryptionKekType("CustomerManaged")
)

func init() {
}
