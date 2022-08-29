


package v20180601preview

type BlockchainProtocol string

const (
	BlockchainProtocolNotSpecified = BlockchainProtocol("NotSpecified")
	BlockchainProtocolParity       = BlockchainProtocol("Parity")
	BlockchainProtocolQuorum       = BlockchainProtocol("Quorum")
	BlockchainProtocolCorda        = BlockchainProtocol("Corda")
)

func init() {
}
