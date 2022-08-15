


package v20201005privatepreview

type CitusVersion string

const (
	CitusVersion_8_3 = CitusVersion("8.3")
	CitusVersion_9_0 = CitusVersion("9.0")
	CitusVersion_9_1 = CitusVersion("9.1")
	CitusVersion_9_2 = CitusVersion("9.2")
	CitusVersion_9_3 = CitusVersion("9.3")
	CitusVersion_9_4 = CitusVersion("9.4")
	CitusVersion_9_5 = CitusVersion("9.5")
)

type CreateMode string

const (
	CreateModeDefault            = CreateMode("Default")
	CreateModePointInTimeRestore = CreateMode("PointInTimeRestore")
	CreateModeReadReplica        = CreateMode("ReadReplica")
)

type PostgreSQLVersion string

const (
	PostgreSQLVersion_11 = PostgreSQLVersion("11")
	PostgreSQLVersion_12 = PostgreSQLVersion("12")
)

type ServerEdition string

const (
	ServerEditionGeneralPurpose  = ServerEdition("GeneralPurpose")
	ServerEditionMemoryOptimized = ServerEdition("MemoryOptimized")
)

type ServerRole string

const (
	ServerRoleCoordinator = ServerRole("Coordinator")
	ServerRoleWorker      = ServerRole("Worker")
)

func init() {
}
