


package datamigration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureActiveDirectoryApp struct {
	AppKey        string `pulumi:"appKey"`
	ApplicationId string `pulumi:"applicationId"`
	TenantId      string `pulumi:"tenantId"`
}

type AzureActiveDirectoryAppResponse struct {
	AppKey        string `pulumi:"appKey"`
	ApplicationId string `pulumi:"applicationId"`
	TenantId      string `pulumi:"tenantId"`
}

type BackupFileInfoResponse struct {
	FamilySequenceNumber *int    `pulumi:"familySequenceNumber"`
	FileLocation         *string `pulumi:"fileLocation"`
	Status               *string `pulumi:"status"`
}

type BackupSetInfoResponse struct {
	BackupFinishedDate *string                  `pulumi:"backupFinishedDate"`
	BackupSetId        *string                  `pulumi:"backupSetId"`
	BackupStartDate    *string                  `pulumi:"backupStartDate"`
	BackupType         *string                  `pulumi:"backupType"`
	DatabaseName       *string                  `pulumi:"databaseName"`
	FirstLsn           *string                  `pulumi:"firstLsn"`
	IsBackupRestored   *bool                    `pulumi:"isBackupRestored"`
	LastLsn            *string                  `pulumi:"lastLsn"`
	LastModifiedTime   *string                  `pulumi:"lastModifiedTime"`
	ListOfBackupFiles  []BackupFileInfoResponse `pulumi:"listOfBackupFiles"`
}

type BlobShare struct {
	SasUri string `pulumi:"sasUri"`
}

type BlobShareResponse struct {
	SasUri string `pulumi:"sasUri"`
}

type ConnectToSourcePostgreSqlSyncTaskInput struct {
	SourceConnectionInfo PostgreSqlConnectionInfo `pulumi:"sourceConnectionInfo"`
}

type ConnectToSourcePostgreSqlSyncTaskInputResponse struct {
	SourceConnectionInfo PostgreSqlConnectionInfoResponse `pulumi:"sourceConnectionInfo"`
}

type ConnectToSourcePostgreSqlSyncTaskOutputResponse struct {
	Databases                []string                      `pulumi:"databases"`
	Id                       string                        `pulumi:"id"`
	SourceServerBrandVersion string                        `pulumi:"sourceServerBrandVersion"`
	SourceServerVersion      string                        `pulumi:"sourceServerVersion"`
	ValidationErrors         []ReportableExceptionResponse `pulumi:"validationErrors"`
}

type ConnectToSourcePostgreSqlSyncTaskProperties struct {
	Input    *ConnectToSourcePostgreSqlSyncTaskInput `pulumi:"input"`
	TaskType string                                  `pulumi:"taskType"`
}

type ConnectToSourcePostgreSqlSyncTaskPropertiesResponse struct {
	Commands []interface{}                                     `pulumi:"commands"`
	Errors   []ODataErrorResponse                              `pulumi:"errors"`
	Input    *ConnectToSourcePostgreSqlSyncTaskInputResponse   `pulumi:"input"`
	Output   []ConnectToSourcePostgreSqlSyncTaskOutputResponse `pulumi:"output"`
	State    string                                            `pulumi:"state"`
	TaskType string                                            `pulumi:"taskType"`
}

type ConnectToSourceSqlServerSyncTaskProperties struct {
	Input    *ConnectToSourceSqlServerTaskInput `pulumi:"input"`
	TaskType string                             `pulumi:"taskType"`
}


func (val *ConnectToSourceSqlServerSyncTaskProperties) Defaults() *ConnectToSourceSqlServerSyncTaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ConnectToSourceSqlServerSyncTaskPropertiesResponse struct {
	Commands []interface{}                              `pulumi:"commands"`
	Errors   []ODataErrorResponse                       `pulumi:"errors"`
	Input    *ConnectToSourceSqlServerTaskInputResponse `pulumi:"input"`
	Output   []interface{}                              `pulumi:"output"`
	State    string                                     `pulumi:"state"`
	TaskType string                                     `pulumi:"taskType"`
}


func (val *ConnectToSourceSqlServerSyncTaskPropertiesResponse) Defaults() *ConnectToSourceSqlServerSyncTaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ConnectToSourceSqlServerTaskInput struct {
	CheckPermissionsGroup *string           `pulumi:"checkPermissionsGroup"`
	CollectAgentJobs      *bool             `pulumi:"collectAgentJobs"`
	CollectLogins         *bool             `pulumi:"collectLogins"`
	SourceConnectionInfo  SqlConnectionInfo `pulumi:"sourceConnectionInfo"`
}


func (val *ConnectToSourceSqlServerTaskInput) Defaults() *ConnectToSourceSqlServerTaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CollectAgentJobs) {
		collectAgentJobs_ := false
		tmp.CollectAgentJobs = &collectAgentJobs_
	}
	if isZero(tmp.CollectLogins) {
		collectLogins_ := false
		tmp.CollectLogins = &collectLogins_
	}
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	return &tmp
}

type ConnectToSourceSqlServerTaskInputResponse struct {
	CheckPermissionsGroup *string                   `pulumi:"checkPermissionsGroup"`
	CollectAgentJobs      *bool                     `pulumi:"collectAgentJobs"`
	CollectLogins         *bool                     `pulumi:"collectLogins"`
	SourceConnectionInfo  SqlConnectionInfoResponse `pulumi:"sourceConnectionInfo"`
}


func (val *ConnectToSourceSqlServerTaskInputResponse) Defaults() *ConnectToSourceSqlServerTaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CollectAgentJobs) {
		collectAgentJobs_ := false
		tmp.CollectAgentJobs = &collectAgentJobs_
	}
	if isZero(tmp.CollectLogins) {
		collectLogins_ := false
		tmp.CollectLogins = &collectLogins_
	}
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	return &tmp
}

type ConnectToSourceSqlServerTaskOutputAgentJobLevelResponse struct {
	Id                   string                           `pulumi:"id"`
	IsEnabled            bool                             `pulumi:"isEnabled"`
	JobCategory          string                           `pulumi:"jobCategory"`
	JobOwner             string                           `pulumi:"jobOwner"`
	LastExecutedOn       string                           `pulumi:"lastExecutedOn"`
	MigrationEligibility MigrationEligibilityInfoResponse `pulumi:"migrationEligibility"`
	Name                 string                           `pulumi:"name"`
	ResultType           string                           `pulumi:"resultType"`
}

type ConnectToSourceSqlServerTaskOutputDatabaseLevelResponse struct {
	CompatibilityLevel string                     `pulumi:"compatibilityLevel"`
	DatabaseFiles      []DatabaseFileInfoResponse `pulumi:"databaseFiles"`
	DatabaseState      string                     `pulumi:"databaseState"`
	Id                 string                     `pulumi:"id"`
	Name               string                     `pulumi:"name"`
	ResultType         string                     `pulumi:"resultType"`
	SizeMB             float64                    `pulumi:"sizeMB"`
}

type ConnectToSourceSqlServerTaskOutputLoginLevelResponse struct {
	DefaultDatabase      string                           `pulumi:"defaultDatabase"`
	Id                   string                           `pulumi:"id"`
	IsEnabled            bool                             `pulumi:"isEnabled"`
	LoginType            string                           `pulumi:"loginType"`
	MigrationEligibility MigrationEligibilityInfoResponse `pulumi:"migrationEligibility"`
	Name                 string                           `pulumi:"name"`
	ResultType           string                           `pulumi:"resultType"`
}

type ConnectToSourceSqlServerTaskOutputTaskLevelResponse struct {
	AgentJobs                map[string]string             `pulumi:"agentJobs"`
	Databases                map[string]string             `pulumi:"databases"`
	Id                       string                        `pulumi:"id"`
	Logins                   map[string]string             `pulumi:"logins"`
	ResultType               string                        `pulumi:"resultType"`
	SourceServerBrandVersion string                        `pulumi:"sourceServerBrandVersion"`
	SourceServerVersion      string                        `pulumi:"sourceServerVersion"`
	ValidationErrors         []ReportableExceptionResponse `pulumi:"validationErrors"`
}

type ConnectToSourceSqlServerTaskProperties struct {
	Input    *ConnectToSourceSqlServerTaskInput `pulumi:"input"`
	TaskType string                             `pulumi:"taskType"`
}


func (val *ConnectToSourceSqlServerTaskProperties) Defaults() *ConnectToSourceSqlServerTaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ConnectToSourceSqlServerTaskPropertiesResponse struct {
	Commands []interface{}                              `pulumi:"commands"`
	Errors   []ODataErrorResponse                       `pulumi:"errors"`
	Input    *ConnectToSourceSqlServerTaskInputResponse `pulumi:"input"`
	Output   []interface{}                              `pulumi:"output"`
	State    string                                     `pulumi:"state"`
	TaskType string                                     `pulumi:"taskType"`
}


func (val *ConnectToSourceSqlServerTaskPropertiesResponse) Defaults() *ConnectToSourceSqlServerTaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ConnectToTargetAzureDbForMySqlTaskInput struct {
	SourceConnectionInfo MySqlConnectionInfo `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo MySqlConnectionInfo `pulumi:"targetConnectionInfo"`
}

type ConnectToTargetAzureDbForMySqlTaskInputResponse struct {
	SourceConnectionInfo MySqlConnectionInfoResponse `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo MySqlConnectionInfoResponse `pulumi:"targetConnectionInfo"`
}

type ConnectToTargetAzureDbForMySqlTaskOutputResponse struct {
	Databases                []string                      `pulumi:"databases"`
	Id                       string                        `pulumi:"id"`
	ServerVersion            string                        `pulumi:"serverVersion"`
	TargetServerBrandVersion string                        `pulumi:"targetServerBrandVersion"`
	ValidationErrors         []ReportableExceptionResponse `pulumi:"validationErrors"`
}

type ConnectToTargetAzureDbForMySqlTaskProperties struct {
	Input    *ConnectToTargetAzureDbForMySqlTaskInput `pulumi:"input"`
	TaskType string                                   `pulumi:"taskType"`
}

type ConnectToTargetAzureDbForMySqlTaskPropertiesResponse struct {
	Commands []interface{}                                      `pulumi:"commands"`
	Errors   []ODataErrorResponse                               `pulumi:"errors"`
	Input    *ConnectToTargetAzureDbForMySqlTaskInputResponse   `pulumi:"input"`
	Output   []ConnectToTargetAzureDbForMySqlTaskOutputResponse `pulumi:"output"`
	State    string                                             `pulumi:"state"`
	TaskType string                                             `pulumi:"taskType"`
}

type ConnectToTargetAzureDbForPostgreSqlSyncTaskInput struct {
	SourceConnectionInfo PostgreSqlConnectionInfo `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo PostgreSqlConnectionInfo `pulumi:"targetConnectionInfo"`
}

type ConnectToTargetAzureDbForPostgreSqlSyncTaskInputResponse struct {
	SourceConnectionInfo PostgreSqlConnectionInfoResponse `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo PostgreSqlConnectionInfoResponse `pulumi:"targetConnectionInfo"`
}

type ConnectToTargetAzureDbForPostgreSqlSyncTaskOutputResponse struct {
	Databases                []string                      `pulumi:"databases"`
	Id                       string                        `pulumi:"id"`
	TargetServerBrandVersion string                        `pulumi:"targetServerBrandVersion"`
	TargetServerVersion      string                        `pulumi:"targetServerVersion"`
	ValidationErrors         []ReportableExceptionResponse `pulumi:"validationErrors"`
}

type ConnectToTargetAzureDbForPostgreSqlSyncTaskProperties struct {
	Input    *ConnectToTargetAzureDbForPostgreSqlSyncTaskInput `pulumi:"input"`
	TaskType string                                            `pulumi:"taskType"`
}

type ConnectToTargetAzureDbForPostgreSqlSyncTaskPropertiesResponse struct {
	Commands []interface{}                                               `pulumi:"commands"`
	Errors   []ODataErrorResponse                                        `pulumi:"errors"`
	Input    *ConnectToTargetAzureDbForPostgreSqlSyncTaskInputResponse   `pulumi:"input"`
	Output   []ConnectToTargetAzureDbForPostgreSqlSyncTaskOutputResponse `pulumi:"output"`
	State    string                                                      `pulumi:"state"`
	TaskType string                                                      `pulumi:"taskType"`
}

type ConnectToTargetSqlDbTaskInput struct {
	TargetConnectionInfo SqlConnectionInfo `pulumi:"targetConnectionInfo"`
}


func (val *ConnectToTargetSqlDbTaskInput) Defaults() *ConnectToTargetSqlDbTaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type ConnectToTargetSqlDbTaskInputResponse struct {
	TargetConnectionInfo SqlConnectionInfoResponse `pulumi:"targetConnectionInfo"`
}


func (val *ConnectToTargetSqlDbTaskInputResponse) Defaults() *ConnectToTargetSqlDbTaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type ConnectToTargetSqlDbTaskOutputResponse struct {
	Databases                map[string]string `pulumi:"databases"`
	Id                       string            `pulumi:"id"`
	TargetServerBrandVersion string            `pulumi:"targetServerBrandVersion"`
	TargetServerVersion      string            `pulumi:"targetServerVersion"`
}

type ConnectToTargetSqlDbTaskProperties struct {
	Input    *ConnectToTargetSqlDbTaskInput `pulumi:"input"`
	TaskType string                         `pulumi:"taskType"`
}


func (val *ConnectToTargetSqlDbTaskProperties) Defaults() *ConnectToTargetSqlDbTaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ConnectToTargetSqlDbTaskPropertiesResponse struct {
	Commands []interface{}                            `pulumi:"commands"`
	Errors   []ODataErrorResponse                     `pulumi:"errors"`
	Input    *ConnectToTargetSqlDbTaskInputResponse   `pulumi:"input"`
	Output   []ConnectToTargetSqlDbTaskOutputResponse `pulumi:"output"`
	State    string                                   `pulumi:"state"`
	TaskType string                                   `pulumi:"taskType"`
}


func (val *ConnectToTargetSqlDbTaskPropertiesResponse) Defaults() *ConnectToTargetSqlDbTaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ConnectToTargetSqlMISyncTaskInput struct {
	AzureApp             AzureActiveDirectoryApp `pulumi:"azureApp"`
	TargetConnectionInfo MiSqlConnectionInfo     `pulumi:"targetConnectionInfo"`
}

type ConnectToTargetSqlMISyncTaskInputResponse struct {
	AzureApp             AzureActiveDirectoryAppResponse `pulumi:"azureApp"`
	TargetConnectionInfo MiSqlConnectionInfoResponse     `pulumi:"targetConnectionInfo"`
}

type ConnectToTargetSqlMISyncTaskOutputResponse struct {
	TargetServerBrandVersion string                        `pulumi:"targetServerBrandVersion"`
	TargetServerVersion      string                        `pulumi:"targetServerVersion"`
	ValidationErrors         []ReportableExceptionResponse `pulumi:"validationErrors"`
}

type ConnectToTargetSqlMISyncTaskProperties struct {
	Input    *ConnectToTargetSqlMISyncTaskInput `pulumi:"input"`
	TaskType string                             `pulumi:"taskType"`
}

type ConnectToTargetSqlMISyncTaskPropertiesResponse struct {
	Commands []interface{}                                `pulumi:"commands"`
	Errors   []ODataErrorResponse                         `pulumi:"errors"`
	Input    *ConnectToTargetSqlMISyncTaskInputResponse   `pulumi:"input"`
	Output   []ConnectToTargetSqlMISyncTaskOutputResponse `pulumi:"output"`
	State    string                                       `pulumi:"state"`
	TaskType string                                       `pulumi:"taskType"`
}

type ConnectToTargetSqlMITaskInput struct {
	TargetConnectionInfo SqlConnectionInfo `pulumi:"targetConnectionInfo"`
}


func (val *ConnectToTargetSqlMITaskInput) Defaults() *ConnectToTargetSqlMITaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type ConnectToTargetSqlMITaskInputResponse struct {
	TargetConnectionInfo SqlConnectionInfoResponse `pulumi:"targetConnectionInfo"`
}


func (val *ConnectToTargetSqlMITaskInputResponse) Defaults() *ConnectToTargetSqlMITaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type ConnectToTargetSqlMITaskOutputResponse struct {
	AgentJobs                []string                      `pulumi:"agentJobs"`
	Id                       string                        `pulumi:"id"`
	Logins                   []string                      `pulumi:"logins"`
	TargetServerBrandVersion string                        `pulumi:"targetServerBrandVersion"`
	TargetServerVersion      string                        `pulumi:"targetServerVersion"`
	ValidationErrors         []ReportableExceptionResponse `pulumi:"validationErrors"`
}

type ConnectToTargetSqlMITaskProperties struct {
	Input    *ConnectToTargetSqlMITaskInput `pulumi:"input"`
	TaskType string                         `pulumi:"taskType"`
}


func (val *ConnectToTargetSqlMITaskProperties) Defaults() *ConnectToTargetSqlMITaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ConnectToTargetSqlMITaskPropertiesResponse struct {
	Commands []interface{}                            `pulumi:"commands"`
	Errors   []ODataErrorResponse                     `pulumi:"errors"`
	Input    *ConnectToTargetSqlMITaskInputResponse   `pulumi:"input"`
	Output   []ConnectToTargetSqlMITaskOutputResponse `pulumi:"output"`
	State    string                                   `pulumi:"state"`
	TaskType string                                   `pulumi:"taskType"`
}


func (val *ConnectToTargetSqlMITaskPropertiesResponse) Defaults() *ConnectToTargetSqlMITaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ConnectToTargetSqlSqlDbSyncTaskInput struct {
	SourceConnectionInfo SqlConnectionInfo `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo SqlConnectionInfo `pulumi:"targetConnectionInfo"`
}


func (val *ConnectToTargetSqlSqlDbSyncTaskInput) Defaults() *ConnectToTargetSqlSqlDbSyncTaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type ConnectToTargetSqlSqlDbSyncTaskInputResponse struct {
	SourceConnectionInfo SqlConnectionInfoResponse `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo SqlConnectionInfoResponse `pulumi:"targetConnectionInfo"`
}


func (val *ConnectToTargetSqlSqlDbSyncTaskInputResponse) Defaults() *ConnectToTargetSqlSqlDbSyncTaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type ConnectToTargetSqlSqlDbSyncTaskProperties struct {
	Input    *ConnectToTargetSqlSqlDbSyncTaskInput `pulumi:"input"`
	TaskType string                                `pulumi:"taskType"`
}


func (val *ConnectToTargetSqlSqlDbSyncTaskProperties) Defaults() *ConnectToTargetSqlSqlDbSyncTaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ConnectToTargetSqlSqlDbSyncTaskPropertiesResponse struct {
	Commands []interface{}                                 `pulumi:"commands"`
	Errors   []ODataErrorResponse                          `pulumi:"errors"`
	Input    *ConnectToTargetSqlSqlDbSyncTaskInputResponse `pulumi:"input"`
	Output   []ConnectToTargetSqlDbTaskOutputResponse      `pulumi:"output"`
	State    string                                        `pulumi:"state"`
	TaskType string                                        `pulumi:"taskType"`
}


func (val *ConnectToTargetSqlSqlDbSyncTaskPropertiesResponse) Defaults() *ConnectToTargetSqlSqlDbSyncTaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type CopyProgressDetailsResponse struct {
	CopyDuration       int     `pulumi:"copyDuration"`
	CopyStart          string  `pulumi:"copyStart"`
	CopyThroughput     float64 `pulumi:"copyThroughput"`
	DataRead           float64 `pulumi:"dataRead"`
	DataWritten        float64 `pulumi:"dataWritten"`
	ParallelCopyType   string  `pulumi:"parallelCopyType"`
	RowsCopied         float64 `pulumi:"rowsCopied"`
	RowsRead           float64 `pulumi:"rowsRead"`
	Status             string  `pulumi:"status"`
	TableName          string  `pulumi:"tableName"`
	UsedParallelCopies int     `pulumi:"usedParallelCopies"`
}

type CopyProgressDetailsResponseOutput struct{ *pulumi.OutputState }

func (CopyProgressDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyProgressDetailsResponse)(nil)).Elem()
}

func (o CopyProgressDetailsResponseOutput) ToCopyProgressDetailsResponseOutput() CopyProgressDetailsResponseOutput {
	return o
}

func (o CopyProgressDetailsResponseOutput) ToCopyProgressDetailsResponseOutputWithContext(ctx context.Context) CopyProgressDetailsResponseOutput {
	return o
}

func (o CopyProgressDetailsResponseOutput) CopyDuration() pulumi.IntOutput {
	return o.ApplyT(func(v CopyProgressDetailsResponse) int { return v.CopyDuration }).(pulumi.IntOutput)
}

func (o CopyProgressDetailsResponseOutput) CopyStart() pulumi.StringOutput {
	return o.ApplyT(func(v CopyProgressDetailsResponse) string { return v.CopyStart }).(pulumi.StringOutput)
}

func (o CopyProgressDetailsResponseOutput) CopyThroughput() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressDetailsResponse) float64 { return v.CopyThroughput }).(pulumi.Float64Output)
}

func (o CopyProgressDetailsResponseOutput) DataRead() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressDetailsResponse) float64 { return v.DataRead }).(pulumi.Float64Output)
}

func (o CopyProgressDetailsResponseOutput) DataWritten() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressDetailsResponse) float64 { return v.DataWritten }).(pulumi.Float64Output)
}

func (o CopyProgressDetailsResponseOutput) ParallelCopyType() pulumi.StringOutput {
	return o.ApplyT(func(v CopyProgressDetailsResponse) string { return v.ParallelCopyType }).(pulumi.StringOutput)
}

func (o CopyProgressDetailsResponseOutput) RowsCopied() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressDetailsResponse) float64 { return v.RowsCopied }).(pulumi.Float64Output)
}

func (o CopyProgressDetailsResponseOutput) RowsRead() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressDetailsResponse) float64 { return v.RowsRead }).(pulumi.Float64Output)
}

func (o CopyProgressDetailsResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v CopyProgressDetailsResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o CopyProgressDetailsResponseOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v CopyProgressDetailsResponse) string { return v.TableName }).(pulumi.StringOutput)
}

func (o CopyProgressDetailsResponseOutput) UsedParallelCopies() pulumi.IntOutput {
	return o.ApplyT(func(v CopyProgressDetailsResponse) int { return v.UsedParallelCopies }).(pulumi.IntOutput)
}

type CopyProgressDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (CopyProgressDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CopyProgressDetailsResponse)(nil)).Elem()
}

func (o CopyProgressDetailsResponseArrayOutput) ToCopyProgressDetailsResponseArrayOutput() CopyProgressDetailsResponseArrayOutput {
	return o
}

func (o CopyProgressDetailsResponseArrayOutput) ToCopyProgressDetailsResponseArrayOutputWithContext(ctx context.Context) CopyProgressDetailsResponseArrayOutput {
	return o
}

func (o CopyProgressDetailsResponseArrayOutput) Index(i pulumi.IntInput) CopyProgressDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CopyProgressDetailsResponse {
		return vs[0].([]CopyProgressDetailsResponse)[vs[1].(int)]
	}).(CopyProgressDetailsResponseOutput)
}

type DataIntegrityValidationResultResponse struct {
	FailedObjects    map[string]string       `pulumi:"failedObjects"`
	ValidationErrors ValidationErrorResponse `pulumi:"validationErrors"`
}

type DataItemMigrationSummaryResultResponse struct {
	EndedOn             string  `pulumi:"endedOn"`
	ErrorPrefix         string  `pulumi:"errorPrefix"`
	ItemsCompletedCount float64 `pulumi:"itemsCompletedCount"`
	ItemsCount          float64 `pulumi:"itemsCount"`
	Name                string  `pulumi:"name"`
	ResultPrefix        string  `pulumi:"resultPrefix"`
	StartedOn           string  `pulumi:"startedOn"`
	State               string  `pulumi:"state"`
	StatusMessage       string  `pulumi:"statusMessage"`
}

type DatabaseBackupInfoResponse struct {
	BackupFiles      []string `pulumi:"backupFiles"`
	BackupFinishDate string   `pulumi:"backupFinishDate"`
	BackupType       string   `pulumi:"backupType"`
	DatabaseName     string   `pulumi:"databaseName"`
	FamilyCount      int      `pulumi:"familyCount"`
	IsCompressed     bool     `pulumi:"isCompressed"`
	IsDamaged        bool     `pulumi:"isDamaged"`
	Position         int      `pulumi:"position"`
}

type DatabaseFileInfoResponse struct {
	DatabaseName     *string  `pulumi:"databaseName"`
	FileType         *string  `pulumi:"fileType"`
	Id               *string  `pulumi:"id"`
	LogicalName      *string  `pulumi:"logicalName"`
	PhysicalFullName *string  `pulumi:"physicalFullName"`
	RestoreFullName  *string  `pulumi:"restoreFullName"`
	SizeMB           *float64 `pulumi:"sizeMB"`
}

type DatabaseInfo struct {
	SourceDatabaseName string `pulumi:"sourceDatabaseName"`
}





type DatabaseInfoInput interface {
	pulumi.Input

	ToDatabaseInfoOutput() DatabaseInfoOutput
	ToDatabaseInfoOutputWithContext(context.Context) DatabaseInfoOutput
}

type DatabaseInfoArgs struct {
	SourceDatabaseName pulumi.StringInput `pulumi:"sourceDatabaseName"`
}

func (DatabaseInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseInfo)(nil)).Elem()
}

func (i DatabaseInfoArgs) ToDatabaseInfoOutput() DatabaseInfoOutput {
	return i.ToDatabaseInfoOutputWithContext(context.Background())
}

func (i DatabaseInfoArgs) ToDatabaseInfoOutputWithContext(ctx context.Context) DatabaseInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInfoOutput)
}





type DatabaseInfoArrayInput interface {
	pulumi.Input

	ToDatabaseInfoArrayOutput() DatabaseInfoArrayOutput
	ToDatabaseInfoArrayOutputWithContext(context.Context) DatabaseInfoArrayOutput
}

type DatabaseInfoArray []DatabaseInfoInput

func (DatabaseInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseInfo)(nil)).Elem()
}

func (i DatabaseInfoArray) ToDatabaseInfoArrayOutput() DatabaseInfoArrayOutput {
	return i.ToDatabaseInfoArrayOutputWithContext(context.Background())
}

func (i DatabaseInfoArray) ToDatabaseInfoArrayOutputWithContext(ctx context.Context) DatabaseInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInfoArrayOutput)
}

type DatabaseInfoOutput struct{ *pulumi.OutputState }

func (DatabaseInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseInfo)(nil)).Elem()
}

func (o DatabaseInfoOutput) ToDatabaseInfoOutput() DatabaseInfoOutput {
	return o
}

func (o DatabaseInfoOutput) ToDatabaseInfoOutputWithContext(ctx context.Context) DatabaseInfoOutput {
	return o
}

func (o DatabaseInfoOutput) SourceDatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseInfo) string { return v.SourceDatabaseName }).(pulumi.StringOutput)
}

type DatabaseInfoArrayOutput struct{ *pulumi.OutputState }

func (DatabaseInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseInfo)(nil)).Elem()
}

func (o DatabaseInfoArrayOutput) ToDatabaseInfoArrayOutput() DatabaseInfoArrayOutput {
	return o
}

func (o DatabaseInfoArrayOutput) ToDatabaseInfoArrayOutputWithContext(ctx context.Context) DatabaseInfoArrayOutput {
	return o
}

func (o DatabaseInfoArrayOutput) Index(i pulumi.IntInput) DatabaseInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseInfo {
		return vs[0].([]DatabaseInfo)[vs[1].(int)]
	}).(DatabaseInfoOutput)
}

type DatabaseInfoResponse struct {
	SourceDatabaseName string `pulumi:"sourceDatabaseName"`
}

type DatabaseInfoResponseOutput struct{ *pulumi.OutputState }

func (DatabaseInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseInfoResponse)(nil)).Elem()
}

func (o DatabaseInfoResponseOutput) ToDatabaseInfoResponseOutput() DatabaseInfoResponseOutput {
	return o
}

func (o DatabaseInfoResponseOutput) ToDatabaseInfoResponseOutputWithContext(ctx context.Context) DatabaseInfoResponseOutput {
	return o
}

func (o DatabaseInfoResponseOutput) SourceDatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseInfoResponse) string { return v.SourceDatabaseName }).(pulumi.StringOutput)
}

type DatabaseInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseInfoResponse)(nil)).Elem()
}

func (o DatabaseInfoResponseArrayOutput) ToDatabaseInfoResponseArrayOutput() DatabaseInfoResponseArrayOutput {
	return o
}

func (o DatabaseInfoResponseArrayOutput) ToDatabaseInfoResponseArrayOutputWithContext(ctx context.Context) DatabaseInfoResponseArrayOutput {
	return o
}

func (o DatabaseInfoResponseArrayOutput) Index(i pulumi.IntInput) DatabaseInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseInfoResponse {
		return vs[0].([]DatabaseInfoResponse)[vs[1].(int)]
	}).(DatabaseInfoResponseOutput)
}

type DatabaseMigrationPropertiesSqlDb struct {
	Kind                    string                    `pulumi:"kind"`
	MigrationOperationId    *string                   `pulumi:"migrationOperationId"`
	MigrationService        *string                   `pulumi:"migrationService"`
	ProvisioningError       *string                   `pulumi:"provisioningError"`
	Scope                   *string                   `pulumi:"scope"`
	SourceDatabaseName      *string                   `pulumi:"sourceDatabaseName"`
	SourceSqlConnection     *SqlConnectionInformation `pulumi:"sourceSqlConnection"`
	TableList               []string                  `pulumi:"tableList"`
	TargetDatabaseCollation *string                   `pulumi:"targetDatabaseCollation"`
	TargetSqlConnection     *SqlConnectionInformation `pulumi:"targetSqlConnection"`
}





type DatabaseMigrationPropertiesSqlDbInput interface {
	pulumi.Input

	ToDatabaseMigrationPropertiesSqlDbOutput() DatabaseMigrationPropertiesSqlDbOutput
	ToDatabaseMigrationPropertiesSqlDbOutputWithContext(context.Context) DatabaseMigrationPropertiesSqlDbOutput
}

type DatabaseMigrationPropertiesSqlDbArgs struct {
	Kind                    pulumi.StringInput               `pulumi:"kind"`
	MigrationOperationId    pulumi.StringPtrInput            `pulumi:"migrationOperationId"`
	MigrationService        pulumi.StringPtrInput            `pulumi:"migrationService"`
	ProvisioningError       pulumi.StringPtrInput            `pulumi:"provisioningError"`
	Scope                   pulumi.StringPtrInput            `pulumi:"scope"`
	SourceDatabaseName      pulumi.StringPtrInput            `pulumi:"sourceDatabaseName"`
	SourceSqlConnection     SqlConnectionInformationPtrInput `pulumi:"sourceSqlConnection"`
	TableList               pulumi.StringArrayInput          `pulumi:"tableList"`
	TargetDatabaseCollation pulumi.StringPtrInput            `pulumi:"targetDatabaseCollation"`
	TargetSqlConnection     SqlConnectionInformationPtrInput `pulumi:"targetSqlConnection"`
}

func (DatabaseMigrationPropertiesSqlDbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseMigrationPropertiesSqlDb)(nil)).Elem()
}

func (i DatabaseMigrationPropertiesSqlDbArgs) ToDatabaseMigrationPropertiesSqlDbOutput() DatabaseMigrationPropertiesSqlDbOutput {
	return i.ToDatabaseMigrationPropertiesSqlDbOutputWithContext(context.Background())
}

func (i DatabaseMigrationPropertiesSqlDbArgs) ToDatabaseMigrationPropertiesSqlDbOutputWithContext(ctx context.Context) DatabaseMigrationPropertiesSqlDbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMigrationPropertiesSqlDbOutput)
}

func (i DatabaseMigrationPropertiesSqlDbArgs) ToDatabaseMigrationPropertiesSqlDbPtrOutput() DatabaseMigrationPropertiesSqlDbPtrOutput {
	return i.ToDatabaseMigrationPropertiesSqlDbPtrOutputWithContext(context.Background())
}

func (i DatabaseMigrationPropertiesSqlDbArgs) ToDatabaseMigrationPropertiesSqlDbPtrOutputWithContext(ctx context.Context) DatabaseMigrationPropertiesSqlDbPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMigrationPropertiesSqlDbOutput).ToDatabaseMigrationPropertiesSqlDbPtrOutputWithContext(ctx)
}









type DatabaseMigrationPropertiesSqlDbPtrInput interface {
	pulumi.Input

	ToDatabaseMigrationPropertiesSqlDbPtrOutput() DatabaseMigrationPropertiesSqlDbPtrOutput
	ToDatabaseMigrationPropertiesSqlDbPtrOutputWithContext(context.Context) DatabaseMigrationPropertiesSqlDbPtrOutput
}

type databaseMigrationPropertiesSqlDbPtrType DatabaseMigrationPropertiesSqlDbArgs

func DatabaseMigrationPropertiesSqlDbPtr(v *DatabaseMigrationPropertiesSqlDbArgs) DatabaseMigrationPropertiesSqlDbPtrInput {
	return (*databaseMigrationPropertiesSqlDbPtrType)(v)
}

func (*databaseMigrationPropertiesSqlDbPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseMigrationPropertiesSqlDb)(nil)).Elem()
}

func (i *databaseMigrationPropertiesSqlDbPtrType) ToDatabaseMigrationPropertiesSqlDbPtrOutput() DatabaseMigrationPropertiesSqlDbPtrOutput {
	return i.ToDatabaseMigrationPropertiesSqlDbPtrOutputWithContext(context.Background())
}

func (i *databaseMigrationPropertiesSqlDbPtrType) ToDatabaseMigrationPropertiesSqlDbPtrOutputWithContext(ctx context.Context) DatabaseMigrationPropertiesSqlDbPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMigrationPropertiesSqlDbPtrOutput)
}

type DatabaseMigrationPropertiesSqlDbOutput struct{ *pulumi.OutputState }

func (DatabaseMigrationPropertiesSqlDbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseMigrationPropertiesSqlDb)(nil)).Elem()
}

func (o DatabaseMigrationPropertiesSqlDbOutput) ToDatabaseMigrationPropertiesSqlDbOutput() DatabaseMigrationPropertiesSqlDbOutput {
	return o
}

func (o DatabaseMigrationPropertiesSqlDbOutput) ToDatabaseMigrationPropertiesSqlDbOutputWithContext(ctx context.Context) DatabaseMigrationPropertiesSqlDbOutput {
	return o
}

func (o DatabaseMigrationPropertiesSqlDbOutput) ToDatabaseMigrationPropertiesSqlDbPtrOutput() DatabaseMigrationPropertiesSqlDbPtrOutput {
	return o.ToDatabaseMigrationPropertiesSqlDbPtrOutputWithContext(context.Background())
}

func (o DatabaseMigrationPropertiesSqlDbOutput) ToDatabaseMigrationPropertiesSqlDbPtrOutputWithContext(ctx context.Context) DatabaseMigrationPropertiesSqlDbPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseMigrationPropertiesSqlDb) *DatabaseMigrationPropertiesSqlDb {
		return &v
	}).(DatabaseMigrationPropertiesSqlDbPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDb) string { return v.Kind }).(pulumi.StringOutput)
}

func (o DatabaseMigrationPropertiesSqlDbOutput) MigrationOperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDb) *string { return v.MigrationOperationId }).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbOutput) MigrationService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDb) *string { return v.MigrationService }).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbOutput) ProvisioningError() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDb) *string { return v.ProvisioningError }).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDb) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbOutput) SourceDatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDb) *string { return v.SourceDatabaseName }).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbOutput) SourceSqlConnection() SqlConnectionInformationPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDb) *SqlConnectionInformation { return v.SourceSqlConnection }).(SqlConnectionInformationPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbOutput) TableList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDb) []string { return v.TableList }).(pulumi.StringArrayOutput)
}

func (o DatabaseMigrationPropertiesSqlDbOutput) TargetDatabaseCollation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDb) *string { return v.TargetDatabaseCollation }).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbOutput) TargetSqlConnection() SqlConnectionInformationPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDb) *SqlConnectionInformation { return v.TargetSqlConnection }).(SqlConnectionInformationPtrOutput)
}

type DatabaseMigrationPropertiesSqlDbPtrOutput struct{ *pulumi.OutputState }

func (DatabaseMigrationPropertiesSqlDbPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseMigrationPropertiesSqlDb)(nil)).Elem()
}

func (o DatabaseMigrationPropertiesSqlDbPtrOutput) ToDatabaseMigrationPropertiesSqlDbPtrOutput() DatabaseMigrationPropertiesSqlDbPtrOutput {
	return o
}

func (o DatabaseMigrationPropertiesSqlDbPtrOutput) ToDatabaseMigrationPropertiesSqlDbPtrOutputWithContext(ctx context.Context) DatabaseMigrationPropertiesSqlDbPtrOutput {
	return o
}

func (o DatabaseMigrationPropertiesSqlDbPtrOutput) Elem() DatabaseMigrationPropertiesSqlDbOutput {
	return o.ApplyT(func(v *DatabaseMigrationPropertiesSqlDb) DatabaseMigrationPropertiesSqlDb {
		if v != nil {
			return *v
		}
		var ret DatabaseMigrationPropertiesSqlDb
		return ret
	}).(DatabaseMigrationPropertiesSqlDbOutput)
}

func (o DatabaseMigrationPropertiesSqlDbPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationPropertiesSqlDb) *string {
		if v == nil {
			return nil
		}
		return &v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbPtrOutput) MigrationOperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationPropertiesSqlDb) *string {
		if v == nil {
			return nil
		}
		return v.MigrationOperationId
	}).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbPtrOutput) MigrationService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationPropertiesSqlDb) *string {
		if v == nil {
			return nil
		}
		return v.MigrationService
	}).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbPtrOutput) ProvisioningError() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationPropertiesSqlDb) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningError
	}).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbPtrOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationPropertiesSqlDb) *string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbPtrOutput) SourceDatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationPropertiesSqlDb) *string {
		if v == nil {
			return nil
		}
		return v.SourceDatabaseName
	}).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbPtrOutput) SourceSqlConnection() SqlConnectionInformationPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationPropertiesSqlDb) *SqlConnectionInformation {
		if v == nil {
			return nil
		}
		return v.SourceSqlConnection
	}).(SqlConnectionInformationPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbPtrOutput) TableList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseMigrationPropertiesSqlDb) []string {
		if v == nil {
			return nil
		}
		return v.TableList
	}).(pulumi.StringArrayOutput)
}

func (o DatabaseMigrationPropertiesSqlDbPtrOutput) TargetDatabaseCollation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationPropertiesSqlDb) *string {
		if v == nil {
			return nil
		}
		return v.TargetDatabaseCollation
	}).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbPtrOutput) TargetSqlConnection() SqlConnectionInformationPtrOutput {
	return o.ApplyT(func(v *DatabaseMigrationPropertiesSqlDb) *SqlConnectionInformation {
		if v == nil {
			return nil
		}
		return v.TargetSqlConnection
	}).(SqlConnectionInformationPtrOutput)
}

type DatabaseMigrationPropertiesSqlDbResponse struct {
	EndedOn                 string                              `pulumi:"endedOn"`
	Kind                    string                              `pulumi:"kind"`
	MigrationFailureError   ErrorInfoResponse                   `pulumi:"migrationFailureError"`
	MigrationOperationId    *string                             `pulumi:"migrationOperationId"`
	MigrationService        *string                             `pulumi:"migrationService"`
	MigrationStatus         string                              `pulumi:"migrationStatus"`
	MigrationStatusDetails  SqlDbMigrationStatusDetailsResponse `pulumi:"migrationStatusDetails"`
	OfflineConfiguration    SqlDbOfflineConfigurationResponse   `pulumi:"offlineConfiguration"`
	ProvisioningError       *string                             `pulumi:"provisioningError"`
	ProvisioningState       string                              `pulumi:"provisioningState"`
	Scope                   *string                             `pulumi:"scope"`
	SourceDatabaseName      *string                             `pulumi:"sourceDatabaseName"`
	SourceServerName        string                              `pulumi:"sourceServerName"`
	SourceSqlConnection     *SqlConnectionInformationResponse   `pulumi:"sourceSqlConnection"`
	StartedOn               string                              `pulumi:"startedOn"`
	TableList               []string                            `pulumi:"tableList"`
	TargetDatabaseCollation *string                             `pulumi:"targetDatabaseCollation"`
	TargetSqlConnection     *SqlConnectionInformationResponse   `pulumi:"targetSqlConnection"`
}

type DatabaseMigrationPropertiesSqlDbResponseOutput struct{ *pulumi.OutputState }

func (DatabaseMigrationPropertiesSqlDbResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseMigrationPropertiesSqlDbResponse)(nil)).Elem()
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) ToDatabaseMigrationPropertiesSqlDbResponseOutput() DatabaseMigrationPropertiesSqlDbResponseOutput {
	return o
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) ToDatabaseMigrationPropertiesSqlDbResponseOutputWithContext(ctx context.Context) DatabaseMigrationPropertiesSqlDbResponseOutput {
	return o
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) EndedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) string { return v.EndedOn }).(pulumi.StringOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) MigrationFailureError() ErrorInfoResponseOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) ErrorInfoResponse { return v.MigrationFailureError }).(ErrorInfoResponseOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) MigrationOperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) *string { return v.MigrationOperationId }).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) MigrationService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) *string { return v.MigrationService }).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) MigrationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) string { return v.MigrationStatus }).(pulumi.StringOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) MigrationStatusDetails() SqlDbMigrationStatusDetailsResponseOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) SqlDbMigrationStatusDetailsResponse {
		return v.MigrationStatusDetails
	}).(SqlDbMigrationStatusDetailsResponseOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) OfflineConfiguration() SqlDbOfflineConfigurationResponseOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) SqlDbOfflineConfigurationResponse {
		return v.OfflineConfiguration
	}).(SqlDbOfflineConfigurationResponseOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) ProvisioningError() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) *string { return v.ProvisioningError }).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) SourceDatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) *string { return v.SourceDatabaseName }).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) SourceServerName() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) string { return v.SourceServerName }).(pulumi.StringOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) SourceSqlConnection() SqlConnectionInformationResponsePtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) *SqlConnectionInformationResponse {
		return v.SourceSqlConnection
	}).(SqlConnectionInformationResponsePtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) StartedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) string { return v.StartedOn }).(pulumi.StringOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) TableList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) []string { return v.TableList }).(pulumi.StringArrayOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) TargetDatabaseCollation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) *string { return v.TargetDatabaseCollation }).(pulumi.StringPtrOutput)
}

func (o DatabaseMigrationPropertiesSqlDbResponseOutput) TargetSqlConnection() SqlConnectionInformationResponsePtrOutput {
	return o.ApplyT(func(v DatabaseMigrationPropertiesSqlDbResponse) *SqlConnectionInformationResponse {
		return v.TargetSqlConnection
	}).(SqlConnectionInformationResponsePtrOutput)
}

type DatabaseSummaryResultResponse struct {
	EndedOn             string  `pulumi:"endedOn"`
	ErrorPrefix         string  `pulumi:"errorPrefix"`
	ItemsCompletedCount float64 `pulumi:"itemsCompletedCount"`
	ItemsCount          float64 `pulumi:"itemsCount"`
	Name                string  `pulumi:"name"`
	ResultPrefix        string  `pulumi:"resultPrefix"`
	SizeMB              float64 `pulumi:"sizeMB"`
	StartedOn           string  `pulumi:"startedOn"`
	State               string  `pulumi:"state"`
	StatusMessage       string  `pulumi:"statusMessage"`
}

type DatabaseTableResponse struct {
	HasRows bool   `pulumi:"hasRows"`
	Name    string `pulumi:"name"`
}

type ErrorInfoResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}

type ErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (ErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorInfoResponse)(nil)).Elem()
}

func (o ErrorInfoResponseOutput) ToErrorInfoResponseOutput() ErrorInfoResponseOutput {
	return o
}

func (o ErrorInfoResponseOutput) ToErrorInfoResponseOutputWithContext(ctx context.Context) ErrorInfoResponseOutput {
	return o
}

func (o ErrorInfoResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorInfoResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorInfoResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorInfoResponse) string { return v.Message }).(pulumi.StringOutput)
}

type ExecutionStatisticsResponse struct {
	CpuTimeMs      float64                           `pulumi:"cpuTimeMs"`
	ElapsedTimeMs  float64                           `pulumi:"elapsedTimeMs"`
	ExecutionCount float64                           `pulumi:"executionCount"`
	HasErrors      bool                              `pulumi:"hasErrors"`
	SqlErrors      []string                          `pulumi:"sqlErrors"`
	WaitStats      map[string]WaitStatisticsResponse `pulumi:"waitStats"`
}

type FileShare struct {
	Password *string `pulumi:"password"`
	Path     string  `pulumi:"path"`
	UserName *string `pulumi:"userName"`
}

type FileShareResponse struct {
	Password *string `pulumi:"password"`
	Path     string  `pulumi:"path"`
	UserName *string `pulumi:"userName"`
}

type GetTdeCertificatesSqlTaskInput struct {
	BackupFileShare      FileShare                  `pulumi:"backupFileShare"`
	ConnectionInfo       SqlConnectionInfo          `pulumi:"connectionInfo"`
	SelectedCertificates []SelectedCertificateInput `pulumi:"selectedCertificates"`
}


func (val *GetTdeCertificatesSqlTaskInput) Defaults() *GetTdeCertificatesSqlTaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConnectionInfo = *tmp.ConnectionInfo.Defaults()

	return &tmp
}

type GetTdeCertificatesSqlTaskInputResponse struct {
	BackupFileShare      FileShareResponse                  `pulumi:"backupFileShare"`
	ConnectionInfo       SqlConnectionInfoResponse          `pulumi:"connectionInfo"`
	SelectedCertificates []SelectedCertificateInputResponse `pulumi:"selectedCertificates"`
}


func (val *GetTdeCertificatesSqlTaskInputResponse) Defaults() *GetTdeCertificatesSqlTaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConnectionInfo = *tmp.ConnectionInfo.Defaults()

	return &tmp
}

type GetTdeCertificatesSqlTaskOutputResponse struct {
	Base64EncodedCertificates map[string][]string           `pulumi:"base64EncodedCertificates"`
	ValidationErrors          []ReportableExceptionResponse `pulumi:"validationErrors"`
}

type GetTdeCertificatesSqlTaskProperties struct {
	Input    *GetTdeCertificatesSqlTaskInput `pulumi:"input"`
	TaskType string                          `pulumi:"taskType"`
}


func (val *GetTdeCertificatesSqlTaskProperties) Defaults() *GetTdeCertificatesSqlTaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type GetTdeCertificatesSqlTaskPropertiesResponse struct {
	Commands []interface{}                             `pulumi:"commands"`
	Errors   []ODataErrorResponse                      `pulumi:"errors"`
	Input    *GetTdeCertificatesSqlTaskInputResponse   `pulumi:"input"`
	Output   []GetTdeCertificatesSqlTaskOutputResponse `pulumi:"output"`
	State    string                                    `pulumi:"state"`
	TaskType string                                    `pulumi:"taskType"`
}


func (val *GetTdeCertificatesSqlTaskPropertiesResponse) Defaults() *GetTdeCertificatesSqlTaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type GetUserTablesSqlSyncTaskInput struct {
	SelectedSourceDatabases []string          `pulumi:"selectedSourceDatabases"`
	SelectedTargetDatabases []string          `pulumi:"selectedTargetDatabases"`
	SourceConnectionInfo    SqlConnectionInfo `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo    SqlConnectionInfo `pulumi:"targetConnectionInfo"`
}


func (val *GetUserTablesSqlSyncTaskInput) Defaults() *GetUserTablesSqlSyncTaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type GetUserTablesSqlSyncTaskInputResponse struct {
	SelectedSourceDatabases []string                  `pulumi:"selectedSourceDatabases"`
	SelectedTargetDatabases []string                  `pulumi:"selectedTargetDatabases"`
	SourceConnectionInfo    SqlConnectionInfoResponse `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo    SqlConnectionInfoResponse `pulumi:"targetConnectionInfo"`
}


func (val *GetUserTablesSqlSyncTaskInputResponse) Defaults() *GetUserTablesSqlSyncTaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type GetUserTablesSqlSyncTaskOutputResponse struct {
	DatabasesToSourceTables map[string][]DatabaseTableResponse `pulumi:"databasesToSourceTables"`
	DatabasesToTargetTables map[string][]DatabaseTableResponse `pulumi:"databasesToTargetTables"`
	TableValidationErrors   map[string][]string                `pulumi:"tableValidationErrors"`
	ValidationErrors        []ReportableExceptionResponse      `pulumi:"validationErrors"`
}

type GetUserTablesSqlSyncTaskProperties struct {
	Input    *GetUserTablesSqlSyncTaskInput `pulumi:"input"`
	TaskType string                         `pulumi:"taskType"`
}


func (val *GetUserTablesSqlSyncTaskProperties) Defaults() *GetUserTablesSqlSyncTaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type GetUserTablesSqlSyncTaskPropertiesResponse struct {
	Commands []interface{}                            `pulumi:"commands"`
	Errors   []ODataErrorResponse                     `pulumi:"errors"`
	Input    *GetUserTablesSqlSyncTaskInputResponse   `pulumi:"input"`
	Output   []GetUserTablesSqlSyncTaskOutputResponse `pulumi:"output"`
	State    string                                   `pulumi:"state"`
	TaskType string                                   `pulumi:"taskType"`
}


func (val *GetUserTablesSqlSyncTaskPropertiesResponse) Defaults() *GetUserTablesSqlSyncTaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type GetUserTablesSqlTaskInput struct {
	ConnectionInfo    SqlConnectionInfo `pulumi:"connectionInfo"`
	SelectedDatabases []string          `pulumi:"selectedDatabases"`
}


func (val *GetUserTablesSqlTaskInput) Defaults() *GetUserTablesSqlTaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConnectionInfo = *tmp.ConnectionInfo.Defaults()

	return &tmp
}

type GetUserTablesSqlTaskInputResponse struct {
	ConnectionInfo    SqlConnectionInfoResponse `pulumi:"connectionInfo"`
	SelectedDatabases []string                  `pulumi:"selectedDatabases"`
}


func (val *GetUserTablesSqlTaskInputResponse) Defaults() *GetUserTablesSqlTaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConnectionInfo = *tmp.ConnectionInfo.Defaults()

	return &tmp
}

type GetUserTablesSqlTaskOutputResponse struct {
	DatabasesToTables map[string][]DatabaseTableResponse `pulumi:"databasesToTables"`
	Id                string                             `pulumi:"id"`
	ValidationErrors  []ReportableExceptionResponse      `pulumi:"validationErrors"`
}

type GetUserTablesSqlTaskProperties struct {
	Input    *GetUserTablesSqlTaskInput `pulumi:"input"`
	TaskType string                     `pulumi:"taskType"`
}


func (val *GetUserTablesSqlTaskProperties) Defaults() *GetUserTablesSqlTaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type GetUserTablesSqlTaskPropertiesResponse struct {
	Commands []interface{}                        `pulumi:"commands"`
	Errors   []ODataErrorResponse                 `pulumi:"errors"`
	Input    *GetUserTablesSqlTaskInputResponse   `pulumi:"input"`
	Output   []GetUserTablesSqlTaskOutputResponse `pulumi:"output"`
	State    string                               `pulumi:"state"`
	TaskType string                               `pulumi:"taskType"`
}


func (val *GetUserTablesSqlTaskPropertiesResponse) Defaults() *GetUserTablesSqlTaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type MiSqlConnectionInfo struct {
	ManagedInstanceResourceId string  `pulumi:"managedInstanceResourceId"`
	Password                  *string `pulumi:"password"`
	Type                      string  `pulumi:"type"`
	UserName                  *string `pulumi:"userName"`
}

type MiSqlConnectionInfoResponse struct {
	ManagedInstanceResourceId string  `pulumi:"managedInstanceResourceId"`
	Password                  *string `pulumi:"password"`
	Type                      string  `pulumi:"type"`
	UserName                  *string `pulumi:"userName"`
}

type MigrateMISyncCompleteCommandInputResponse struct {
	SourceDatabaseName string `pulumi:"sourceDatabaseName"`
}

type MigrateMISyncCompleteCommandOutputResponse struct {
	Errors []ReportableExceptionResponse `pulumi:"errors"`
}

type MigrateMISyncCompleteCommandPropertiesResponse struct {
	CommandType string                                     `pulumi:"commandType"`
	Errors      []ODataErrorResponse                       `pulumi:"errors"`
	Input       *MigrateMISyncCompleteCommandInputResponse `pulumi:"input"`
	Output      MigrateMISyncCompleteCommandOutputResponse `pulumi:"output"`
	State       string                                     `pulumi:"state"`
}

type MigrateMySqlAzureDbForMySqlSyncDatabaseInput struct {
	MigrationSetting   map[string]string `pulumi:"migrationSetting"`
	Name               *string           `pulumi:"name"`
	SourceSetting      map[string]string `pulumi:"sourceSetting"`
	TargetDatabaseName *string           `pulumi:"targetDatabaseName"`
	TargetSetting      map[string]string `pulumi:"targetSetting"`
}

type MigrateMySqlAzureDbForMySqlSyncDatabaseInputResponse struct {
	MigrationSetting   map[string]string `pulumi:"migrationSetting"`
	Name               *string           `pulumi:"name"`
	SourceSetting      map[string]string `pulumi:"sourceSetting"`
	TargetDatabaseName *string           `pulumi:"targetDatabaseName"`
	TargetSetting      map[string]string `pulumi:"targetSetting"`
}

type MigrateMySqlAzureDbForMySqlSyncTaskInput struct {
	SelectedDatabases    []MigrateMySqlAzureDbForMySqlSyncDatabaseInput `pulumi:"selectedDatabases"`
	SourceConnectionInfo MySqlConnectionInfo                            `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo MySqlConnectionInfo                            `pulumi:"targetConnectionInfo"`
}

type MigrateMySqlAzureDbForMySqlSyncTaskInputResponse struct {
	SelectedDatabases    []MigrateMySqlAzureDbForMySqlSyncDatabaseInputResponse `pulumi:"selectedDatabases"`
	SourceConnectionInfo MySqlConnectionInfoResponse                            `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo MySqlConnectionInfoResponse                            `pulumi:"targetConnectionInfo"`
}

type MigrateMySqlAzureDbForMySqlSyncTaskOutputDatabaseErrorResponse struct {
	ErrorMessage *string                                   `pulumi:"errorMessage"`
	Events       []SyncMigrationDatabaseErrorEventResponse `pulumi:"events"`
	Id           string                                    `pulumi:"id"`
	ResultType   string                                    `pulumi:"resultType"`
}

type MigrateMySqlAzureDbForMySqlSyncTaskOutputDatabaseLevelResponse struct {
	AppliedChanges          float64 `pulumi:"appliedChanges"`
	CdcDeleteCounter        float64 `pulumi:"cdcDeleteCounter"`
	CdcInsertCounter        float64 `pulumi:"cdcInsertCounter"`
	CdcUpdateCounter        float64 `pulumi:"cdcUpdateCounter"`
	DatabaseName            string  `pulumi:"databaseName"`
	EndedOn                 string  `pulumi:"endedOn"`
	FullLoadCompletedTables float64 `pulumi:"fullLoadCompletedTables"`
	FullLoadErroredTables   float64 `pulumi:"fullLoadErroredTables"`
	FullLoadLoadingTables   float64 `pulumi:"fullLoadLoadingTables"`
	FullLoadQueuedTables    float64 `pulumi:"fullLoadQueuedTables"`
	Id                      string  `pulumi:"id"`
	IncomingChanges         float64 `pulumi:"incomingChanges"`
	InitializationCompleted bool    `pulumi:"initializationCompleted"`
	Latency                 float64 `pulumi:"latency"`
	MigrationState          string  `pulumi:"migrationState"`
	ResultType              string  `pulumi:"resultType"`
	StartedOn               string  `pulumi:"startedOn"`
}

type MigrateMySqlAzureDbForMySqlSyncTaskOutputErrorResponse struct {
	Error      ReportableExceptionResponse `pulumi:"error"`
	Id         string                      `pulumi:"id"`
	ResultType string                      `pulumi:"resultType"`
}

type MigrateMySqlAzureDbForMySqlSyncTaskOutputMigrationLevelResponse struct {
	EndedOn             string `pulumi:"endedOn"`
	Id                  string `pulumi:"id"`
	ResultType          string `pulumi:"resultType"`
	SourceServer        string `pulumi:"sourceServer"`
	SourceServerVersion string `pulumi:"sourceServerVersion"`
	StartedOn           string `pulumi:"startedOn"`
	TargetServer        string `pulumi:"targetServer"`
	TargetServerVersion string `pulumi:"targetServerVersion"`
}

type MigrateMySqlAzureDbForMySqlSyncTaskOutputTableLevelResponse struct {
	CdcDeleteCounter      string  `pulumi:"cdcDeleteCounter"`
	CdcInsertCounter      string  `pulumi:"cdcInsertCounter"`
	CdcUpdateCounter      string  `pulumi:"cdcUpdateCounter"`
	DataErrorsCounter     float64 `pulumi:"dataErrorsCounter"`
	DatabaseName          string  `pulumi:"databaseName"`
	FullLoadEndedOn       string  `pulumi:"fullLoadEndedOn"`
	FullLoadEstFinishTime string  `pulumi:"fullLoadEstFinishTime"`
	FullLoadStartedOn     string  `pulumi:"fullLoadStartedOn"`
	FullLoadTotalRows     float64 `pulumi:"fullLoadTotalRows"`
	Id                    string  `pulumi:"id"`
	LastModifiedTime      string  `pulumi:"lastModifiedTime"`
	ResultType            string  `pulumi:"resultType"`
	State                 string  `pulumi:"state"`
	TableName             string  `pulumi:"tableName"`
	TotalChangesApplied   float64 `pulumi:"totalChangesApplied"`
}

type MigrateMySqlAzureDbForMySqlSyncTaskProperties struct {
	Input    *MigrateMySqlAzureDbForMySqlSyncTaskInput `pulumi:"input"`
	TaskType string                                    `pulumi:"taskType"`
}

type MigrateMySqlAzureDbForMySqlSyncTaskPropertiesResponse struct {
	Commands []interface{}                                     `pulumi:"commands"`
	Errors   []ODataErrorResponse                              `pulumi:"errors"`
	Input    *MigrateMySqlAzureDbForMySqlSyncTaskInputResponse `pulumi:"input"`
	Output   []interface{}                                     `pulumi:"output"`
	State    string                                            `pulumi:"state"`
	TaskType string                                            `pulumi:"taskType"`
}

type MigratePostgreSqlAzureDbForPostgreSqlSyncDatabaseInput struct {
	MigrationSetting   map[string]string `pulumi:"migrationSetting"`
	Name               *string           `pulumi:"name"`
	SourceSetting      map[string]string `pulumi:"sourceSetting"`
	TargetDatabaseName *string           `pulumi:"targetDatabaseName"`
	TargetSetting      map[string]string `pulumi:"targetSetting"`
}

type MigratePostgreSqlAzureDbForPostgreSqlSyncDatabaseInputResponse struct {
	MigrationSetting   map[string]string `pulumi:"migrationSetting"`
	Name               *string           `pulumi:"name"`
	SourceSetting      map[string]string `pulumi:"sourceSetting"`
	TargetDatabaseName *string           `pulumi:"targetDatabaseName"`
	TargetSetting      map[string]string `pulumi:"targetSetting"`
}

type MigratePostgreSqlAzureDbForPostgreSqlSyncTaskInput struct {
	SelectedDatabases    []MigratePostgreSqlAzureDbForPostgreSqlSyncDatabaseInput `pulumi:"selectedDatabases"`
	SourceConnectionInfo PostgreSqlConnectionInfo                                 `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo PostgreSqlConnectionInfo                                 `pulumi:"targetConnectionInfo"`
}

type MigratePostgreSqlAzureDbForPostgreSqlSyncTaskInputResponse struct {
	SelectedDatabases    []MigratePostgreSqlAzureDbForPostgreSqlSyncDatabaseInputResponse `pulumi:"selectedDatabases"`
	SourceConnectionInfo PostgreSqlConnectionInfoResponse                                 `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo PostgreSqlConnectionInfoResponse                                 `pulumi:"targetConnectionInfo"`
}

type MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputDatabaseErrorResponse struct {
	ErrorMessage *string                                   `pulumi:"errorMessage"`
	Events       []SyncMigrationDatabaseErrorEventResponse `pulumi:"events"`
	Id           string                                    `pulumi:"id"`
	ResultType   string                                    `pulumi:"resultType"`
}

type MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputDatabaseLevelResponse struct {
	AppliedChanges          float64 `pulumi:"appliedChanges"`
	CdcDeleteCounter        float64 `pulumi:"cdcDeleteCounter"`
	CdcInsertCounter        float64 `pulumi:"cdcInsertCounter"`
	CdcUpdateCounter        float64 `pulumi:"cdcUpdateCounter"`
	DatabaseName            string  `pulumi:"databaseName"`
	EndedOn                 string  `pulumi:"endedOn"`
	FullLoadCompletedTables float64 `pulumi:"fullLoadCompletedTables"`
	FullLoadErroredTables   float64 `pulumi:"fullLoadErroredTables"`
	FullLoadLoadingTables   float64 `pulumi:"fullLoadLoadingTables"`
	FullLoadQueuedTables    float64 `pulumi:"fullLoadQueuedTables"`
	Id                      string  `pulumi:"id"`
	IncomingChanges         float64 `pulumi:"incomingChanges"`
	InitializationCompleted bool    `pulumi:"initializationCompleted"`
	Latency                 float64 `pulumi:"latency"`
	MigrationState          string  `pulumi:"migrationState"`
	ResultType              string  `pulumi:"resultType"`
	StartedOn               string  `pulumi:"startedOn"`
}

type MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputErrorResponse struct {
	Error      ReportableExceptionResponse `pulumi:"error"`
	Id         string                      `pulumi:"id"`
	ResultType string                      `pulumi:"resultType"`
}

type MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputMigrationLevelResponse struct {
	EndedOn             string `pulumi:"endedOn"`
	Id                  string `pulumi:"id"`
	ResultType          string `pulumi:"resultType"`
	SourceServer        string `pulumi:"sourceServer"`
	SourceServerVersion string `pulumi:"sourceServerVersion"`
	StartedOn           string `pulumi:"startedOn"`
	TargetServer        string `pulumi:"targetServer"`
	TargetServerVersion string `pulumi:"targetServerVersion"`
}

type MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputTableLevelResponse struct {
	CdcDeleteCounter      float64 `pulumi:"cdcDeleteCounter"`
	CdcInsertCounter      float64 `pulumi:"cdcInsertCounter"`
	CdcUpdateCounter      float64 `pulumi:"cdcUpdateCounter"`
	DataErrorsCounter     float64 `pulumi:"dataErrorsCounter"`
	DatabaseName          string  `pulumi:"databaseName"`
	FullLoadEndedOn       string  `pulumi:"fullLoadEndedOn"`
	FullLoadEstFinishTime string  `pulumi:"fullLoadEstFinishTime"`
	FullLoadStartedOn     string  `pulumi:"fullLoadStartedOn"`
	FullLoadTotalRows     float64 `pulumi:"fullLoadTotalRows"`
	Id                    string  `pulumi:"id"`
	LastModifiedTime      string  `pulumi:"lastModifiedTime"`
	ResultType            string  `pulumi:"resultType"`
	State                 string  `pulumi:"state"`
	TableName             string  `pulumi:"tableName"`
	TotalChangesApplied   float64 `pulumi:"totalChangesApplied"`
}

type MigratePostgreSqlAzureDbForPostgreSqlSyncTaskProperties struct {
	Input    *MigratePostgreSqlAzureDbForPostgreSqlSyncTaskInput `pulumi:"input"`
	TaskType string                                              `pulumi:"taskType"`
}

type MigratePostgreSqlAzureDbForPostgreSqlSyncTaskPropertiesResponse struct {
	Commands []interface{}                                               `pulumi:"commands"`
	Errors   []ODataErrorResponse                                        `pulumi:"errors"`
	Input    *MigratePostgreSqlAzureDbForPostgreSqlSyncTaskInputResponse `pulumi:"input"`
	Output   []interface{}                                               `pulumi:"output"`
	State    string                                                      `pulumi:"state"`
	TaskType string                                                      `pulumi:"taskType"`
}

type MigrateSqlServerSqlDbDatabaseInput struct {
	MakeSourceDbReadOnly *bool             `pulumi:"makeSourceDbReadOnly"`
	Name                 *string           `pulumi:"name"`
	TableMap             map[string]string `pulumi:"tableMap"`
	TargetDatabaseName   *string           `pulumi:"targetDatabaseName"`
}

type MigrateSqlServerSqlDbDatabaseInputResponse struct {
	MakeSourceDbReadOnly *bool             `pulumi:"makeSourceDbReadOnly"`
	Name                 *string           `pulumi:"name"`
	TableMap             map[string]string `pulumi:"tableMap"`
	TargetDatabaseName   *string           `pulumi:"targetDatabaseName"`
}

type MigrateSqlServerSqlDbSyncDatabaseInput struct {
	Id                 *string           `pulumi:"id"`
	MigrationSetting   map[string]string `pulumi:"migrationSetting"`
	Name               *string           `pulumi:"name"`
	SchemaName         *string           `pulumi:"schemaName"`
	SourceSetting      map[string]string `pulumi:"sourceSetting"`
	TableMap           map[string]string `pulumi:"tableMap"`
	TargetDatabaseName *string           `pulumi:"targetDatabaseName"`
	TargetSetting      map[string]string `pulumi:"targetSetting"`
}

type MigrateSqlServerSqlDbSyncDatabaseInputResponse struct {
	Id                 *string           `pulumi:"id"`
	MigrationSetting   map[string]string `pulumi:"migrationSetting"`
	Name               *string           `pulumi:"name"`
	SchemaName         *string           `pulumi:"schemaName"`
	SourceSetting      map[string]string `pulumi:"sourceSetting"`
	TableMap           map[string]string `pulumi:"tableMap"`
	TargetDatabaseName *string           `pulumi:"targetDatabaseName"`
	TargetSetting      map[string]string `pulumi:"targetSetting"`
}

type MigrateSqlServerSqlDbSyncTaskInput struct {
	SelectedDatabases    []MigrateSqlServerSqlDbSyncDatabaseInput `pulumi:"selectedDatabases"`
	SourceConnectionInfo SqlConnectionInfo                        `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo SqlConnectionInfo                        `pulumi:"targetConnectionInfo"`
	ValidationOptions    *MigrationValidationOptions              `pulumi:"validationOptions"`
}


func (val *MigrateSqlServerSqlDbSyncTaskInput) Defaults() *MigrateSqlServerSqlDbSyncTaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type MigrateSqlServerSqlDbSyncTaskInputResponse struct {
	SelectedDatabases    []MigrateSqlServerSqlDbSyncDatabaseInputResponse `pulumi:"selectedDatabases"`
	SourceConnectionInfo SqlConnectionInfoResponse                        `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo SqlConnectionInfoResponse                        `pulumi:"targetConnectionInfo"`
	ValidationOptions    *MigrationValidationOptionsResponse              `pulumi:"validationOptions"`
}


func (val *MigrateSqlServerSqlDbSyncTaskInputResponse) Defaults() *MigrateSqlServerSqlDbSyncTaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type MigrateSqlServerSqlDbSyncTaskOutputDatabaseErrorResponse struct {
	ErrorMessage *string                                   `pulumi:"errorMessage"`
	Events       []SyncMigrationDatabaseErrorEventResponse `pulumi:"events"`
	Id           string                                    `pulumi:"id"`
	ResultType   string                                    `pulumi:"resultType"`
}

type MigrateSqlServerSqlDbSyncTaskOutputDatabaseLevelResponse struct {
	AppliedChanges          float64 `pulumi:"appliedChanges"`
	CdcDeleteCounter        float64 `pulumi:"cdcDeleteCounter"`
	CdcInsertCounter        float64 `pulumi:"cdcInsertCounter"`
	CdcUpdateCounter        float64 `pulumi:"cdcUpdateCounter"`
	DatabaseName            string  `pulumi:"databaseName"`
	EndedOn                 string  `pulumi:"endedOn"`
	FullLoadCompletedTables float64 `pulumi:"fullLoadCompletedTables"`
	FullLoadErroredTables   float64 `pulumi:"fullLoadErroredTables"`
	FullLoadLoadingTables   float64 `pulumi:"fullLoadLoadingTables"`
	FullLoadQueuedTables    float64 `pulumi:"fullLoadQueuedTables"`
	Id                      string  `pulumi:"id"`
	IncomingChanges         float64 `pulumi:"incomingChanges"`
	InitializationCompleted bool    `pulumi:"initializationCompleted"`
	Latency                 float64 `pulumi:"latency"`
	MigrationState          string  `pulumi:"migrationState"`
	ResultType              string  `pulumi:"resultType"`
	StartedOn               string  `pulumi:"startedOn"`
}

type MigrateSqlServerSqlDbSyncTaskOutputErrorResponse struct {
	Error      ReportableExceptionResponse `pulumi:"error"`
	Id         string                      `pulumi:"id"`
	ResultType string                      `pulumi:"resultType"`
}

type MigrateSqlServerSqlDbSyncTaskOutputMigrationLevelResponse struct {
	DatabaseCount       int    `pulumi:"databaseCount"`
	EndedOn             string `pulumi:"endedOn"`
	Id                  string `pulumi:"id"`
	ResultType          string `pulumi:"resultType"`
	SourceServer        string `pulumi:"sourceServer"`
	SourceServerVersion string `pulumi:"sourceServerVersion"`
	StartedOn           string `pulumi:"startedOn"`
	TargetServer        string `pulumi:"targetServer"`
	TargetServerVersion string `pulumi:"targetServerVersion"`
}

type MigrateSqlServerSqlDbSyncTaskOutputTableLevelResponse struct {
	CdcDeleteCounter      float64 `pulumi:"cdcDeleteCounter"`
	CdcInsertCounter      float64 `pulumi:"cdcInsertCounter"`
	CdcUpdateCounter      float64 `pulumi:"cdcUpdateCounter"`
	DataErrorsCounter     float64 `pulumi:"dataErrorsCounter"`
	DatabaseName          string  `pulumi:"databaseName"`
	FullLoadEndedOn       string  `pulumi:"fullLoadEndedOn"`
	FullLoadEstFinishTime string  `pulumi:"fullLoadEstFinishTime"`
	FullLoadStartedOn     string  `pulumi:"fullLoadStartedOn"`
	FullLoadTotalRows     float64 `pulumi:"fullLoadTotalRows"`
	Id                    string  `pulumi:"id"`
	LastModifiedTime      string  `pulumi:"lastModifiedTime"`
	ResultType            string  `pulumi:"resultType"`
	State                 string  `pulumi:"state"`
	TableName             string  `pulumi:"tableName"`
	TotalChangesApplied   float64 `pulumi:"totalChangesApplied"`
}

type MigrateSqlServerSqlDbSyncTaskProperties struct {
	Input    *MigrateSqlServerSqlDbSyncTaskInput `pulumi:"input"`
	TaskType string                              `pulumi:"taskType"`
}


func (val *MigrateSqlServerSqlDbSyncTaskProperties) Defaults() *MigrateSqlServerSqlDbSyncTaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type MigrateSqlServerSqlDbSyncTaskPropertiesResponse struct {
	Commands []interface{}                               `pulumi:"commands"`
	Errors   []ODataErrorResponse                        `pulumi:"errors"`
	Input    *MigrateSqlServerSqlDbSyncTaskInputResponse `pulumi:"input"`
	Output   []interface{}                               `pulumi:"output"`
	State    string                                      `pulumi:"state"`
	TaskType string                                      `pulumi:"taskType"`
}


func (val *MigrateSqlServerSqlDbSyncTaskPropertiesResponse) Defaults() *MigrateSqlServerSqlDbSyncTaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type MigrateSqlServerSqlDbTaskInput struct {
	SelectedDatabases    []MigrateSqlServerSqlDbDatabaseInput `pulumi:"selectedDatabases"`
	SourceConnectionInfo SqlConnectionInfo                    `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo SqlConnectionInfo                    `pulumi:"targetConnectionInfo"`
	ValidationOptions    *MigrationValidationOptions          `pulumi:"validationOptions"`
}


func (val *MigrateSqlServerSqlDbTaskInput) Defaults() *MigrateSqlServerSqlDbTaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type MigrateSqlServerSqlDbTaskInputResponse struct {
	SelectedDatabases    []MigrateSqlServerSqlDbDatabaseInputResponse `pulumi:"selectedDatabases"`
	SourceConnectionInfo SqlConnectionInfoResponse                    `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo SqlConnectionInfoResponse                    `pulumi:"targetConnectionInfo"`
	ValidationOptions    *MigrationValidationOptionsResponse          `pulumi:"validationOptions"`
}


func (val *MigrateSqlServerSqlDbTaskInputResponse) Defaults() *MigrateSqlServerSqlDbTaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type MigrateSqlServerSqlDbTaskOutputDatabaseLevelResponse struct {
	DatabaseName             string                                            `pulumi:"databaseName"`
	EndedOn                  string                                            `pulumi:"endedOn"`
	ErrorCount               float64                                           `pulumi:"errorCount"`
	ErrorPrefix              string                                            `pulumi:"errorPrefix"`
	ExceptionsAndWarnings    []ReportableExceptionResponse                     `pulumi:"exceptionsAndWarnings"`
	Id                       string                                            `pulumi:"id"`
	Message                  string                                            `pulumi:"message"`
	NumberOfObjects          float64                                           `pulumi:"numberOfObjects"`
	NumberOfObjectsCompleted float64                                           `pulumi:"numberOfObjectsCompleted"`
	ObjectSummary            map[string]DataItemMigrationSummaryResultResponse `pulumi:"objectSummary"`
	ResultPrefix             string                                            `pulumi:"resultPrefix"`
	ResultType               string                                            `pulumi:"resultType"`
	Stage                    string                                            `pulumi:"stage"`
	StartedOn                string                                            `pulumi:"startedOn"`
	State                    string                                            `pulumi:"state"`
	StatusMessage            string                                            `pulumi:"statusMessage"`
}

type MigrateSqlServerSqlDbTaskOutputDatabaseLevelValidationResultResponse struct {
	DataIntegrityValidationResult DataIntegrityValidationResultResponse    `pulumi:"dataIntegrityValidationResult"`
	EndedOn                       string                                   `pulumi:"endedOn"`
	Id                            string                                   `pulumi:"id"`
	MigrationId                   string                                   `pulumi:"migrationId"`
	QueryAnalysisValidationResult QueryAnalysisValidationResultResponse    `pulumi:"queryAnalysisValidationResult"`
	ResultType                    string                                   `pulumi:"resultType"`
	SchemaValidationResult        SchemaComparisonValidationResultResponse `pulumi:"schemaValidationResult"`
	SourceDatabaseName            string                                   `pulumi:"sourceDatabaseName"`
	StartedOn                     string                                   `pulumi:"startedOn"`
	Status                        string                                   `pulumi:"status"`
	TargetDatabaseName            string                                   `pulumi:"targetDatabaseName"`
}

type MigrateSqlServerSqlDbTaskOutputErrorResponse struct {
	Error      ReportableExceptionResponse `pulumi:"error"`
	Id         string                      `pulumi:"id"`
	ResultType string                      `pulumi:"resultType"`
}

type MigrateSqlServerSqlDbTaskOutputMigrationLevelResponse struct {
	DatabaseSummary          map[string]DatabaseSummaryResultResponse `pulumi:"databaseSummary"`
	Databases                map[string]string                        `pulumi:"databases"`
	DurationInSeconds        float64                                  `pulumi:"durationInSeconds"`
	EndedOn                  string                                   `pulumi:"endedOn"`
	ExceptionsAndWarnings    []ReportableExceptionResponse            `pulumi:"exceptionsAndWarnings"`
	Id                       string                                   `pulumi:"id"`
	Message                  string                                   `pulumi:"message"`
	MigrationReport          MigrationReportResultResponse            `pulumi:"migrationReport"`
	ResultType               string                                   `pulumi:"resultType"`
	SourceServerBrandVersion string                                   `pulumi:"sourceServerBrandVersion"`
	SourceServerVersion      string                                   `pulumi:"sourceServerVersion"`
	StartedOn                string                                   `pulumi:"startedOn"`
	Status                   string                                   `pulumi:"status"`
	StatusMessage            string                                   `pulumi:"statusMessage"`
	TargetServerBrandVersion string                                   `pulumi:"targetServerBrandVersion"`
	TargetServerVersion      string                                   `pulumi:"targetServerVersion"`
}

type MigrateSqlServerSqlDbTaskOutputTableLevelResponse struct {
	EndedOn             string  `pulumi:"endedOn"`
	ErrorPrefix         string  `pulumi:"errorPrefix"`
	Id                  string  `pulumi:"id"`
	ItemsCompletedCount float64 `pulumi:"itemsCompletedCount"`
	ItemsCount          float64 `pulumi:"itemsCount"`
	ObjectName          string  `pulumi:"objectName"`
	ResultPrefix        string  `pulumi:"resultPrefix"`
	ResultType          string  `pulumi:"resultType"`
	StartedOn           string  `pulumi:"startedOn"`
	State               string  `pulumi:"state"`
	StatusMessage       string  `pulumi:"statusMessage"`
}

type MigrateSqlServerSqlDbTaskOutputValidationResultResponse struct {
	Id             string                                                      `pulumi:"id"`
	MigrationId    string                                                      `pulumi:"migrationId"`
	ResultType     string                                                      `pulumi:"resultType"`
	Status         string                                                      `pulumi:"status"`
	SummaryResults map[string]MigrationValidationDatabaseSummaryResultResponse `pulumi:"summaryResults"`
}

type MigrateSqlServerSqlDbTaskProperties struct {
	Input    *MigrateSqlServerSqlDbTaskInput `pulumi:"input"`
	TaskType string                          `pulumi:"taskType"`
}


func (val *MigrateSqlServerSqlDbTaskProperties) Defaults() *MigrateSqlServerSqlDbTaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type MigrateSqlServerSqlDbTaskPropertiesResponse struct {
	Commands []interface{}                           `pulumi:"commands"`
	Errors   []ODataErrorResponse                    `pulumi:"errors"`
	Input    *MigrateSqlServerSqlDbTaskInputResponse `pulumi:"input"`
	Output   []interface{}                           `pulumi:"output"`
	State    string                                  `pulumi:"state"`
	TaskType string                                  `pulumi:"taskType"`
}


func (val *MigrateSqlServerSqlDbTaskPropertiesResponse) Defaults() *MigrateSqlServerSqlDbTaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type MigrateSqlServerSqlMIDatabaseInput struct {
	BackupFilePaths     []string   `pulumi:"backupFilePaths"`
	BackupFileShare     *FileShare `pulumi:"backupFileShare"`
	Name                string     `pulumi:"name"`
	RestoreDatabaseName string     `pulumi:"restoreDatabaseName"`
}

type MigrateSqlServerSqlMIDatabaseInputResponse struct {
	BackupFilePaths     []string           `pulumi:"backupFilePaths"`
	BackupFileShare     *FileShareResponse `pulumi:"backupFileShare"`
	Name                string             `pulumi:"name"`
	RestoreDatabaseName string             `pulumi:"restoreDatabaseName"`
}

type MigrateSqlServerSqlMISyncTaskInput struct {
	AzureApp             AzureActiveDirectoryApp              `pulumi:"azureApp"`
	BackupFileShare      *FileShare                           `pulumi:"backupFileShare"`
	SelectedDatabases    []MigrateSqlServerSqlMIDatabaseInput `pulumi:"selectedDatabases"`
	SourceConnectionInfo SqlConnectionInfo                    `pulumi:"sourceConnectionInfo"`
	StorageResourceId    string                               `pulumi:"storageResourceId"`
	TargetConnectionInfo MiSqlConnectionInfo                  `pulumi:"targetConnectionInfo"`
}


func (val *MigrateSqlServerSqlMISyncTaskInput) Defaults() *MigrateSqlServerSqlMISyncTaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	return &tmp
}

type MigrateSqlServerSqlMISyncTaskInputResponse struct {
	AzureApp             AzureActiveDirectoryAppResponse              `pulumi:"azureApp"`
	BackupFileShare      *FileShareResponse                           `pulumi:"backupFileShare"`
	SelectedDatabases    []MigrateSqlServerSqlMIDatabaseInputResponse `pulumi:"selectedDatabases"`
	SourceConnectionInfo SqlConnectionInfoResponse                    `pulumi:"sourceConnectionInfo"`
	StorageResourceId    string                                       `pulumi:"storageResourceId"`
	TargetConnectionInfo MiSqlConnectionInfoResponse                  `pulumi:"targetConnectionInfo"`
}


func (val *MigrateSqlServerSqlMISyncTaskInputResponse) Defaults() *MigrateSqlServerSqlMISyncTaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	return &tmp
}

type MigrateSqlServerSqlMISyncTaskOutputDatabaseLevelResponse struct {
	ActiveBackupSets          []BackupSetInfoResponse       `pulumi:"activeBackupSets"`
	ContainerName             string                        `pulumi:"containerName"`
	EndedOn                   string                        `pulumi:"endedOn"`
	ErrorPrefix               string                        `pulumi:"errorPrefix"`
	ExceptionsAndWarnings     []ReportableExceptionResponse `pulumi:"exceptionsAndWarnings"`
	FullBackupSetInfo         BackupSetInfoResponse         `pulumi:"fullBackupSetInfo"`
	Id                        string                        `pulumi:"id"`
	IsFullBackupRestored      bool                          `pulumi:"isFullBackupRestored"`
	LastRestoredBackupSetInfo BackupSetInfoResponse         `pulumi:"lastRestoredBackupSetInfo"`
	MigrationState            string                        `pulumi:"migrationState"`
	ResultType                string                        `pulumi:"resultType"`
	SourceDatabaseName        string                        `pulumi:"sourceDatabaseName"`
	StartedOn                 string                        `pulumi:"startedOn"`
}

type MigrateSqlServerSqlMISyncTaskOutputErrorResponse struct {
	Error      ReportableExceptionResponse `pulumi:"error"`
	Id         string                      `pulumi:"id"`
	ResultType string                      `pulumi:"resultType"`
}

type MigrateSqlServerSqlMISyncTaskOutputMigrationLevelResponse struct {
	DatabaseCount            int    `pulumi:"databaseCount"`
	DatabaseErrorCount       int    `pulumi:"databaseErrorCount"`
	EndedOn                  string `pulumi:"endedOn"`
	Id                       string `pulumi:"id"`
	ResultType               string `pulumi:"resultType"`
	SourceServerBrandVersion string `pulumi:"sourceServerBrandVersion"`
	SourceServerName         string `pulumi:"sourceServerName"`
	SourceServerVersion      string `pulumi:"sourceServerVersion"`
	StartedOn                string `pulumi:"startedOn"`
	State                    string `pulumi:"state"`
	TargetServerBrandVersion string `pulumi:"targetServerBrandVersion"`
	TargetServerName         string `pulumi:"targetServerName"`
	TargetServerVersion      string `pulumi:"targetServerVersion"`
}

type MigrateSqlServerSqlMISyncTaskProperties struct {
	Input    *MigrateSqlServerSqlMISyncTaskInput `pulumi:"input"`
	TaskType string                              `pulumi:"taskType"`
}


func (val *MigrateSqlServerSqlMISyncTaskProperties) Defaults() *MigrateSqlServerSqlMISyncTaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type MigrateSqlServerSqlMISyncTaskPropertiesResponse struct {
	Commands []interface{}                               `pulumi:"commands"`
	Errors   []ODataErrorResponse                        `pulumi:"errors"`
	Input    *MigrateSqlServerSqlMISyncTaskInputResponse `pulumi:"input"`
	Output   []interface{}                               `pulumi:"output"`
	State    string                                      `pulumi:"state"`
	TaskType string                                      `pulumi:"taskType"`
}


func (val *MigrateSqlServerSqlMISyncTaskPropertiesResponse) Defaults() *MigrateSqlServerSqlMISyncTaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type MigrateSqlServerSqlMITaskInput struct {
	BackupBlobShare      BlobShare                            `pulumi:"backupBlobShare"`
	BackupFileShare      *FileShare                           `pulumi:"backupFileShare"`
	BackupMode           *string                              `pulumi:"backupMode"`
	SelectedAgentJobs    []string                             `pulumi:"selectedAgentJobs"`
	SelectedDatabases    []MigrateSqlServerSqlMIDatabaseInput `pulumi:"selectedDatabases"`
	SelectedLogins       []string                             `pulumi:"selectedLogins"`
	SourceConnectionInfo SqlConnectionInfo                    `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo SqlConnectionInfo                    `pulumi:"targetConnectionInfo"`
}


func (val *MigrateSqlServerSqlMITaskInput) Defaults() *MigrateSqlServerSqlMITaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type MigrateSqlServerSqlMITaskInputResponse struct {
	BackupBlobShare      BlobShareResponse                            `pulumi:"backupBlobShare"`
	BackupFileShare      *FileShareResponse                           `pulumi:"backupFileShare"`
	BackupMode           *string                                      `pulumi:"backupMode"`
	SelectedAgentJobs    []string                                     `pulumi:"selectedAgentJobs"`
	SelectedDatabases    []MigrateSqlServerSqlMIDatabaseInputResponse `pulumi:"selectedDatabases"`
	SelectedLogins       []string                                     `pulumi:"selectedLogins"`
	SourceConnectionInfo SqlConnectionInfoResponse                    `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo SqlConnectionInfoResponse                    `pulumi:"targetConnectionInfo"`
}


func (val *MigrateSqlServerSqlMITaskInputResponse) Defaults() *MigrateSqlServerSqlMITaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type MigrateSqlServerSqlMITaskOutputAgentJobLevelResponse struct {
	EndedOn               string                        `pulumi:"endedOn"`
	ExceptionsAndWarnings []ReportableExceptionResponse `pulumi:"exceptionsAndWarnings"`
	Id                    string                        `pulumi:"id"`
	IsEnabled             bool                          `pulumi:"isEnabled"`
	Message               string                        `pulumi:"message"`
	Name                  string                        `pulumi:"name"`
	ResultType            string                        `pulumi:"resultType"`
	StartedOn             string                        `pulumi:"startedOn"`
	State                 string                        `pulumi:"state"`
}

type MigrateSqlServerSqlMITaskOutputDatabaseLevelResponse struct {
	DatabaseName          string                        `pulumi:"databaseName"`
	EndedOn               string                        `pulumi:"endedOn"`
	ExceptionsAndWarnings []ReportableExceptionResponse `pulumi:"exceptionsAndWarnings"`
	Id                    string                        `pulumi:"id"`
	Message               string                        `pulumi:"message"`
	ResultType            string                        `pulumi:"resultType"`
	SizeMB                float64                       `pulumi:"sizeMB"`
	Stage                 string                        `pulumi:"stage"`
	StartedOn             string                        `pulumi:"startedOn"`
	State                 string                        `pulumi:"state"`
}

type MigrateSqlServerSqlMITaskOutputErrorResponse struct {
	Error      ReportableExceptionResponse `pulumi:"error"`
	Id         string                      `pulumi:"id"`
	ResultType string                      `pulumi:"resultType"`
}

type MigrateSqlServerSqlMITaskOutputLoginLevelResponse struct {
	EndedOn               string                        `pulumi:"endedOn"`
	ExceptionsAndWarnings []ReportableExceptionResponse `pulumi:"exceptionsAndWarnings"`
	Id                    string                        `pulumi:"id"`
	LoginName             string                        `pulumi:"loginName"`
	Message               string                        `pulumi:"message"`
	ResultType            string                        `pulumi:"resultType"`
	Stage                 string                        `pulumi:"stage"`
	StartedOn             string                        `pulumi:"startedOn"`
	State                 string                        `pulumi:"state"`
}

type MigrateSqlServerSqlMITaskOutputMigrationLevelResponse struct {
	AgentJobs                map[string]string                                         `pulumi:"agentJobs"`
	Databases                map[string]string                                         `pulumi:"databases"`
	EndedOn                  string                                                    `pulumi:"endedOn"`
	ExceptionsAndWarnings    []ReportableExceptionResponse                             `pulumi:"exceptionsAndWarnings"`
	Id                       string                                                    `pulumi:"id"`
	Logins                   map[string]string                                         `pulumi:"logins"`
	Message                  string                                                    `pulumi:"message"`
	OrphanedUsersInfo        []OrphanedUserInfoResponse                                `pulumi:"orphanedUsersInfo"`
	ResultType               string                                                    `pulumi:"resultType"`
	ServerRoleResults        map[string]StartMigrationScenarioServerRoleResultResponse `pulumi:"serverRoleResults"`
	SourceServerBrandVersion string                                                    `pulumi:"sourceServerBrandVersion"`
	SourceServerVersion      string                                                    `pulumi:"sourceServerVersion"`
	StartedOn                string                                                    `pulumi:"startedOn"`
	State                    string                                                    `pulumi:"state"`
	Status                   string                                                    `pulumi:"status"`
	TargetServerBrandVersion string                                                    `pulumi:"targetServerBrandVersion"`
	TargetServerVersion      string                                                    `pulumi:"targetServerVersion"`
}

type MigrateSqlServerSqlMITaskProperties struct {
	Input    *MigrateSqlServerSqlMITaskInput `pulumi:"input"`
	TaskType string                          `pulumi:"taskType"`
}


func (val *MigrateSqlServerSqlMITaskProperties) Defaults() *MigrateSqlServerSqlMITaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type MigrateSqlServerSqlMITaskPropertiesResponse struct {
	Commands []interface{}                           `pulumi:"commands"`
	Errors   []ODataErrorResponse                    `pulumi:"errors"`
	Input    *MigrateSqlServerSqlMITaskInputResponse `pulumi:"input"`
	Output   []interface{}                           `pulumi:"output"`
	State    string                                  `pulumi:"state"`
	TaskType string                                  `pulumi:"taskType"`
}


func (val *MigrateSqlServerSqlMITaskPropertiesResponse) Defaults() *MigrateSqlServerSqlMITaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type MigrateSyncCompleteCommandInputResponse struct {
	CommitTimeStamp *string `pulumi:"commitTimeStamp"`
	DatabaseName    string  `pulumi:"databaseName"`
}

type MigrateSyncCompleteCommandOutputResponse struct {
	Errors []ReportableExceptionResponse `pulumi:"errors"`
	Id     string                        `pulumi:"id"`
}

type MigrateSyncCompleteCommandPropertiesResponse struct {
	CommandType string                                   `pulumi:"commandType"`
	Errors      []ODataErrorResponse                     `pulumi:"errors"`
	Input       *MigrateSyncCompleteCommandInputResponse `pulumi:"input"`
	Output      MigrateSyncCompleteCommandOutputResponse `pulumi:"output"`
	State       string                                   `pulumi:"state"`
}

type MigrationEligibilityInfoResponse struct {
	IsEligibleForMigration bool     `pulumi:"isEligibleForMigration"`
	ValidationMessages     []string `pulumi:"validationMessages"`
}

type MigrationReportResultResponse struct {
	Id        string `pulumi:"id"`
	ReportUrl string `pulumi:"reportUrl"`
}

type MigrationValidationDatabaseSummaryResultResponse struct {
	EndedOn            string `pulumi:"endedOn"`
	Id                 string `pulumi:"id"`
	MigrationId        string `pulumi:"migrationId"`
	SourceDatabaseName string `pulumi:"sourceDatabaseName"`
	StartedOn          string `pulumi:"startedOn"`
	Status             string `pulumi:"status"`
	TargetDatabaseName string `pulumi:"targetDatabaseName"`
}

type MigrationValidationOptions struct {
	EnableDataIntegrityValidation *bool `pulumi:"enableDataIntegrityValidation"`
	EnableQueryAnalysisValidation *bool `pulumi:"enableQueryAnalysisValidation"`
	EnableSchemaValidation        *bool `pulumi:"enableSchemaValidation"`
}

type MigrationValidationOptionsResponse struct {
	EnableDataIntegrityValidation *bool `pulumi:"enableDataIntegrityValidation"`
	EnableQueryAnalysisValidation *bool `pulumi:"enableQueryAnalysisValidation"`
	EnableSchemaValidation        *bool `pulumi:"enableSchemaValidation"`
}

type MySqlConnectionInfo struct {
	Password   *string `pulumi:"password"`
	Port       int     `pulumi:"port"`
	ServerName string  `pulumi:"serverName"`
	Type       string  `pulumi:"type"`
	UserName   *string `pulumi:"userName"`
}

type MySqlConnectionInfoResponse struct {
	Password   *string `pulumi:"password"`
	Port       int     `pulumi:"port"`
	ServerName string  `pulumi:"serverName"`
	Type       string  `pulumi:"type"`
	UserName   *string `pulumi:"userName"`
}

type NodeMonitoringDataResponse struct {
	AdditionalProperties  map[string]interface{} `pulumi:"additionalProperties"`
	AvailableMemoryInMB   int                    `pulumi:"availableMemoryInMB"`
	ConcurrentJobsLimit   int                    `pulumi:"concurrentJobsLimit"`
	ConcurrentJobsRunning int                    `pulumi:"concurrentJobsRunning"`
	CpuUtilization        int                    `pulumi:"cpuUtilization"`
	MaxConcurrentJobs     int                    `pulumi:"maxConcurrentJobs"`
	NodeName              string                 `pulumi:"nodeName"`
	ReceivedBytes         float64                `pulumi:"receivedBytes"`
	SentBytes             float64                `pulumi:"sentBytes"`
}

type NodeMonitoringDataResponseOutput struct{ *pulumi.OutputState }

func (NodeMonitoringDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeMonitoringDataResponse)(nil)).Elem()
}

func (o NodeMonitoringDataResponseOutput) ToNodeMonitoringDataResponseOutput() NodeMonitoringDataResponseOutput {
	return o
}

func (o NodeMonitoringDataResponseOutput) ToNodeMonitoringDataResponseOutputWithContext(ctx context.Context) NodeMonitoringDataResponseOutput {
	return o
}

func (o NodeMonitoringDataResponseOutput) AdditionalProperties() pulumi.MapOutput {
	return o.ApplyT(func(v NodeMonitoringDataResponse) map[string]interface{} { return v.AdditionalProperties }).(pulumi.MapOutput)
}

func (o NodeMonitoringDataResponseOutput) AvailableMemoryInMB() pulumi.IntOutput {
	return o.ApplyT(func(v NodeMonitoringDataResponse) int { return v.AvailableMemoryInMB }).(pulumi.IntOutput)
}

func (o NodeMonitoringDataResponseOutput) ConcurrentJobsLimit() pulumi.IntOutput {
	return o.ApplyT(func(v NodeMonitoringDataResponse) int { return v.ConcurrentJobsLimit }).(pulumi.IntOutput)
}

func (o NodeMonitoringDataResponseOutput) ConcurrentJobsRunning() pulumi.IntOutput {
	return o.ApplyT(func(v NodeMonitoringDataResponse) int { return v.ConcurrentJobsRunning }).(pulumi.IntOutput)
}

func (o NodeMonitoringDataResponseOutput) CpuUtilization() pulumi.IntOutput {
	return o.ApplyT(func(v NodeMonitoringDataResponse) int { return v.CpuUtilization }).(pulumi.IntOutput)
}

func (o NodeMonitoringDataResponseOutput) MaxConcurrentJobs() pulumi.IntOutput {
	return o.ApplyT(func(v NodeMonitoringDataResponse) int { return v.MaxConcurrentJobs }).(pulumi.IntOutput)
}

func (o NodeMonitoringDataResponseOutput) NodeName() pulumi.StringOutput {
	return o.ApplyT(func(v NodeMonitoringDataResponse) string { return v.NodeName }).(pulumi.StringOutput)
}

func (o NodeMonitoringDataResponseOutput) ReceivedBytes() pulumi.Float64Output {
	return o.ApplyT(func(v NodeMonitoringDataResponse) float64 { return v.ReceivedBytes }).(pulumi.Float64Output)
}

func (o NodeMonitoringDataResponseOutput) SentBytes() pulumi.Float64Output {
	return o.ApplyT(func(v NodeMonitoringDataResponse) float64 { return v.SentBytes }).(pulumi.Float64Output)
}

type NodeMonitoringDataResponseArrayOutput struct{ *pulumi.OutputState }

func (NodeMonitoringDataResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeMonitoringDataResponse)(nil)).Elem()
}

func (o NodeMonitoringDataResponseArrayOutput) ToNodeMonitoringDataResponseArrayOutput() NodeMonitoringDataResponseArrayOutput {
	return o
}

func (o NodeMonitoringDataResponseArrayOutput) ToNodeMonitoringDataResponseArrayOutputWithContext(ctx context.Context) NodeMonitoringDataResponseArrayOutput {
	return o
}

func (o NodeMonitoringDataResponseArrayOutput) Index(i pulumi.IntInput) NodeMonitoringDataResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeMonitoringDataResponse {
		return vs[0].([]NodeMonitoringDataResponse)[vs[1].(int)]
	}).(NodeMonitoringDataResponseOutput)
}

type ODataErrorResponse struct {
	Code    string               `pulumi:"code"`
	Details []ODataErrorResponse `pulumi:"details"`
	Message string               `pulumi:"message"`
}

type OrphanedUserInfoResponse struct {
	DatabaseName *string `pulumi:"databaseName"`
	Name         *string `pulumi:"name"`
}

type PostgreSqlConnectionInfo struct {
	DatabaseName *string `pulumi:"databaseName"`
	Password     *string `pulumi:"password"`
	Port         int     `pulumi:"port"`
	ServerName   string  `pulumi:"serverName"`
	Type         string  `pulumi:"type"`
	UserName     *string `pulumi:"userName"`
}

type PostgreSqlConnectionInfoResponse struct {
	DatabaseName *string `pulumi:"databaseName"`
	Password     *string `pulumi:"password"`
	Port         int     `pulumi:"port"`
	ServerName   string  `pulumi:"serverName"`
	Type         string  `pulumi:"type"`
	UserName     *string `pulumi:"userName"`
}

type ProjectFileProperties struct {
	Extension *string `pulumi:"extension"`
	FilePath  *string `pulumi:"filePath"`
	MediaType *string `pulumi:"mediaType"`
}





type ProjectFilePropertiesInput interface {
	pulumi.Input

	ToProjectFilePropertiesOutput() ProjectFilePropertiesOutput
	ToProjectFilePropertiesOutputWithContext(context.Context) ProjectFilePropertiesOutput
}

type ProjectFilePropertiesArgs struct {
	Extension pulumi.StringPtrInput `pulumi:"extension"`
	FilePath  pulumi.StringPtrInput `pulumi:"filePath"`
	MediaType pulumi.StringPtrInput `pulumi:"mediaType"`
}

func (ProjectFilePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectFileProperties)(nil)).Elem()
}

func (i ProjectFilePropertiesArgs) ToProjectFilePropertiesOutput() ProjectFilePropertiesOutput {
	return i.ToProjectFilePropertiesOutputWithContext(context.Background())
}

func (i ProjectFilePropertiesArgs) ToProjectFilePropertiesOutputWithContext(ctx context.Context) ProjectFilePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFilePropertiesOutput)
}

func (i ProjectFilePropertiesArgs) ToProjectFilePropertiesPtrOutput() ProjectFilePropertiesPtrOutput {
	return i.ToProjectFilePropertiesPtrOutputWithContext(context.Background())
}

func (i ProjectFilePropertiesArgs) ToProjectFilePropertiesPtrOutputWithContext(ctx context.Context) ProjectFilePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFilePropertiesOutput).ToProjectFilePropertiesPtrOutputWithContext(ctx)
}









type ProjectFilePropertiesPtrInput interface {
	pulumi.Input

	ToProjectFilePropertiesPtrOutput() ProjectFilePropertiesPtrOutput
	ToProjectFilePropertiesPtrOutputWithContext(context.Context) ProjectFilePropertiesPtrOutput
}

type projectFilePropertiesPtrType ProjectFilePropertiesArgs

func ProjectFilePropertiesPtr(v *ProjectFilePropertiesArgs) ProjectFilePropertiesPtrInput {
	return (*projectFilePropertiesPtrType)(v)
}

func (*projectFilePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectFileProperties)(nil)).Elem()
}

func (i *projectFilePropertiesPtrType) ToProjectFilePropertiesPtrOutput() ProjectFilePropertiesPtrOutput {
	return i.ToProjectFilePropertiesPtrOutputWithContext(context.Background())
}

func (i *projectFilePropertiesPtrType) ToProjectFilePropertiesPtrOutputWithContext(ctx context.Context) ProjectFilePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFilePropertiesPtrOutput)
}

type ProjectFilePropertiesOutput struct{ *pulumi.OutputState }

func (ProjectFilePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectFileProperties)(nil)).Elem()
}

func (o ProjectFilePropertiesOutput) ToProjectFilePropertiesOutput() ProjectFilePropertiesOutput {
	return o
}

func (o ProjectFilePropertiesOutput) ToProjectFilePropertiesOutputWithContext(ctx context.Context) ProjectFilePropertiesOutput {
	return o
}

func (o ProjectFilePropertiesOutput) ToProjectFilePropertiesPtrOutput() ProjectFilePropertiesPtrOutput {
	return o.ToProjectFilePropertiesPtrOutputWithContext(context.Background())
}

func (o ProjectFilePropertiesOutput) ToProjectFilePropertiesPtrOutputWithContext(ctx context.Context) ProjectFilePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectFileProperties) *ProjectFileProperties {
		return &v
	}).(ProjectFilePropertiesPtrOutput)
}

func (o ProjectFilePropertiesOutput) Extension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectFileProperties) *string { return v.Extension }).(pulumi.StringPtrOutput)
}

func (o ProjectFilePropertiesOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectFileProperties) *string { return v.FilePath }).(pulumi.StringPtrOutput)
}

func (o ProjectFilePropertiesOutput) MediaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectFileProperties) *string { return v.MediaType }).(pulumi.StringPtrOutput)
}

type ProjectFilePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ProjectFilePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectFileProperties)(nil)).Elem()
}

func (o ProjectFilePropertiesPtrOutput) ToProjectFilePropertiesPtrOutput() ProjectFilePropertiesPtrOutput {
	return o
}

func (o ProjectFilePropertiesPtrOutput) ToProjectFilePropertiesPtrOutputWithContext(ctx context.Context) ProjectFilePropertiesPtrOutput {
	return o
}

func (o ProjectFilePropertiesPtrOutput) Elem() ProjectFilePropertiesOutput {
	return o.ApplyT(func(v *ProjectFileProperties) ProjectFileProperties {
		if v != nil {
			return *v
		}
		var ret ProjectFileProperties
		return ret
	}).(ProjectFilePropertiesOutput)
}

func (o ProjectFilePropertiesPtrOutput) Extension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectFileProperties) *string {
		if v == nil {
			return nil
		}
		return v.Extension
	}).(pulumi.StringPtrOutput)
}

func (o ProjectFilePropertiesPtrOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectFileProperties) *string {
		if v == nil {
			return nil
		}
		return v.FilePath
	}).(pulumi.StringPtrOutput)
}

func (o ProjectFilePropertiesPtrOutput) MediaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectFileProperties) *string {
		if v == nil {
			return nil
		}
		return v.MediaType
	}).(pulumi.StringPtrOutput)
}

type ProjectFilePropertiesResponse struct {
	Extension    *string `pulumi:"extension"`
	FilePath     *string `pulumi:"filePath"`
	LastModified string  `pulumi:"lastModified"`
	MediaType    *string `pulumi:"mediaType"`
	Size         float64 `pulumi:"size"`
}

type ProjectFilePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ProjectFilePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectFilePropertiesResponse)(nil)).Elem()
}

func (o ProjectFilePropertiesResponseOutput) ToProjectFilePropertiesResponseOutput() ProjectFilePropertiesResponseOutput {
	return o
}

func (o ProjectFilePropertiesResponseOutput) ToProjectFilePropertiesResponseOutputWithContext(ctx context.Context) ProjectFilePropertiesResponseOutput {
	return o
}

func (o ProjectFilePropertiesResponseOutput) Extension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectFilePropertiesResponse) *string { return v.Extension }).(pulumi.StringPtrOutput)
}

func (o ProjectFilePropertiesResponseOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectFilePropertiesResponse) *string { return v.FilePath }).(pulumi.StringPtrOutput)
}

func (o ProjectFilePropertiesResponseOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectFilePropertiesResponse) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o ProjectFilePropertiesResponseOutput) MediaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectFilePropertiesResponse) *string { return v.MediaType }).(pulumi.StringPtrOutput)
}

func (o ProjectFilePropertiesResponseOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v ProjectFilePropertiesResponse) float64 { return v.Size }).(pulumi.Float64Output)
}

type QueryAnalysisValidationResultResponse struct {
	QueryResults     QueryExecutionResultResponse `pulumi:"queryResults"`
	ValidationErrors ValidationErrorResponse      `pulumi:"validationErrors"`
}

type QueryExecutionResultResponse struct {
	QueryText         string                      `pulumi:"queryText"`
	SourceResult      ExecutionStatisticsResponse `pulumi:"sourceResult"`
	StatementsInBatch float64                     `pulumi:"statementsInBatch"`
	TargetResult      ExecutionStatisticsResponse `pulumi:"targetResult"`
}

type ReportableExceptionResponse struct {
	ActionableMessage *string `pulumi:"actionableMessage"`
	FilePath          string  `pulumi:"filePath"`
	HResult           int     `pulumi:"hResult"`
	LineNumber        string  `pulumi:"lineNumber"`
	Message           string  `pulumi:"message"`
	StackTrace        string  `pulumi:"stackTrace"`
}

type SchemaComparisonValidationResultResponse struct {
	SchemaDifferences         SchemaComparisonValidationResultTypeResponse `pulumi:"schemaDifferences"`
	SourceDatabaseObjectCount map[string]float64                           `pulumi:"sourceDatabaseObjectCount"`
	TargetDatabaseObjectCount map[string]float64                           `pulumi:"targetDatabaseObjectCount"`
	ValidationErrors          ValidationErrorResponse                      `pulumi:"validationErrors"`
}

type SchemaComparisonValidationResultTypeResponse struct {
	ObjectName   string `pulumi:"objectName"`
	ObjectType   string `pulumi:"objectType"`
	UpdateAction string `pulumi:"updateAction"`
}

type SelectedCertificateInput struct {
	CertificateName string `pulumi:"certificateName"`
	Password        string `pulumi:"password"`
}

type SelectedCertificateInputResponse struct {
	CertificateName string `pulumi:"certificateName"`
	Password        string `pulumi:"password"`
}

type ServiceSku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     *string `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type ServiceSkuInput interface {
	pulumi.Input

	ToServiceSkuOutput() ServiceSkuOutput
	ToServiceSkuOutputWithContext(context.Context) ServiceSkuOutput
}

type ServiceSkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ServiceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceSku)(nil)).Elem()
}

func (i ServiceSkuArgs) ToServiceSkuOutput() ServiceSkuOutput {
	return i.ToServiceSkuOutputWithContext(context.Background())
}

func (i ServiceSkuArgs) ToServiceSkuOutputWithContext(ctx context.Context) ServiceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceSkuOutput)
}

func (i ServiceSkuArgs) ToServiceSkuPtrOutput() ServiceSkuPtrOutput {
	return i.ToServiceSkuPtrOutputWithContext(context.Background())
}

func (i ServiceSkuArgs) ToServiceSkuPtrOutputWithContext(ctx context.Context) ServiceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceSkuOutput).ToServiceSkuPtrOutputWithContext(ctx)
}









type ServiceSkuPtrInput interface {
	pulumi.Input

	ToServiceSkuPtrOutput() ServiceSkuPtrOutput
	ToServiceSkuPtrOutputWithContext(context.Context) ServiceSkuPtrOutput
}

type serviceSkuPtrType ServiceSkuArgs

func ServiceSkuPtr(v *ServiceSkuArgs) ServiceSkuPtrInput {
	return (*serviceSkuPtrType)(v)
}

func (*serviceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceSku)(nil)).Elem()
}

func (i *serviceSkuPtrType) ToServiceSkuPtrOutput() ServiceSkuPtrOutput {
	return i.ToServiceSkuPtrOutputWithContext(context.Background())
}

func (i *serviceSkuPtrType) ToServiceSkuPtrOutputWithContext(ctx context.Context) ServiceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceSkuPtrOutput)
}

type ServiceSkuOutput struct{ *pulumi.OutputState }

func (ServiceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceSku)(nil)).Elem()
}

func (o ServiceSkuOutput) ToServiceSkuOutput() ServiceSkuOutput {
	return o
}

func (o ServiceSkuOutput) ToServiceSkuOutputWithContext(ctx context.Context) ServiceSkuOutput {
	return o
}

func (o ServiceSkuOutput) ToServiceSkuPtrOutput() ServiceSkuPtrOutput {
	return o.ToServiceSkuPtrOutputWithContext(context.Background())
}

func (o ServiceSkuOutput) ToServiceSkuPtrOutputWithContext(ctx context.Context) ServiceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceSku) *ServiceSku {
		return &v
	}).(ServiceSkuPtrOutput)
}

func (o ServiceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ServiceSkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceSku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o ServiceSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceSkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceSku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o ServiceSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ServiceSkuPtrOutput struct{ *pulumi.OutputState }

func (ServiceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceSku)(nil)).Elem()
}

func (o ServiceSkuPtrOutput) ToServiceSkuPtrOutput() ServiceSkuPtrOutput {
	return o
}

func (o ServiceSkuPtrOutput) ToServiceSkuPtrOutputWithContext(ctx context.Context) ServiceSkuPtrOutput {
	return o
}

func (o ServiceSkuPtrOutput) Elem() ServiceSkuOutput {
	return o.ApplyT(func(v *ServiceSku) ServiceSku {
		if v != nil {
			return *v
		}
		var ret ServiceSku
		return ret
	}).(ServiceSkuOutput)
}

func (o ServiceSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ServiceSkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ServiceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ServiceSkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o ServiceSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ServiceSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     *string `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}

type ServiceSkuResponseOutput struct{ *pulumi.OutputState }

func (ServiceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceSkuResponse)(nil)).Elem()
}

func (o ServiceSkuResponseOutput) ToServiceSkuResponseOutput() ServiceSkuResponseOutput {
	return o
}

func (o ServiceSkuResponseOutput) ToServiceSkuResponseOutputWithContext(ctx context.Context) ServiceSkuResponseOutput {
	return o
}

func (o ServiceSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ServiceSkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceSkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o ServiceSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceSkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceSkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o ServiceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ServiceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceSkuResponse)(nil)).Elem()
}

func (o ServiceSkuResponsePtrOutput) ToServiceSkuResponsePtrOutput() ServiceSkuResponsePtrOutput {
	return o
}

func (o ServiceSkuResponsePtrOutput) ToServiceSkuResponsePtrOutputWithContext(ctx context.Context) ServiceSkuResponsePtrOutput {
	return o
}

func (o ServiceSkuResponsePtrOutput) Elem() ServiceSkuResponseOutput {
	return o.ApplyT(func(v *ServiceSkuResponse) ServiceSkuResponse {
		if v != nil {
			return *v
		}
		var ret ServiceSkuResponse
		return ret
	}).(ServiceSkuResponseOutput)
}

func (o ServiceSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ServiceSkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ServiceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ServiceSkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o ServiceSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SqlConnectionInfo struct {
	AdditionalSettings     *string `pulumi:"additionalSettings"`
	Authentication         *string `pulumi:"authentication"`
	DataSource             string  `pulumi:"dataSource"`
	EncryptConnection      *bool   `pulumi:"encryptConnection"`
	Password               *string `pulumi:"password"`
	Platform               *string `pulumi:"platform"`
	TrustServerCertificate *bool   `pulumi:"trustServerCertificate"`
	Type                   string  `pulumi:"type"`
	UserName               *string `pulumi:"userName"`
}


func (val *SqlConnectionInfo) Defaults() *SqlConnectionInfo {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EncryptConnection) {
		encryptConnection_ := true
		tmp.EncryptConnection = &encryptConnection_
	}
	if isZero(tmp.TrustServerCertificate) {
		trustServerCertificate_ := false
		tmp.TrustServerCertificate = &trustServerCertificate_
	}
	return &tmp
}

type SqlConnectionInfoResponse struct {
	AdditionalSettings     *string `pulumi:"additionalSettings"`
	Authentication         *string `pulumi:"authentication"`
	DataSource             string  `pulumi:"dataSource"`
	EncryptConnection      *bool   `pulumi:"encryptConnection"`
	Password               *string `pulumi:"password"`
	Platform               *string `pulumi:"platform"`
	TrustServerCertificate *bool   `pulumi:"trustServerCertificate"`
	Type                   string  `pulumi:"type"`
	UserName               *string `pulumi:"userName"`
}


func (val *SqlConnectionInfoResponse) Defaults() *SqlConnectionInfoResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EncryptConnection) {
		encryptConnection_ := true
		tmp.EncryptConnection = &encryptConnection_
	}
	if isZero(tmp.TrustServerCertificate) {
		trustServerCertificate_ := false
		tmp.TrustServerCertificate = &trustServerCertificate_
	}
	return &tmp
}

type SqlConnectionInformation struct {
	Authentication         *string `pulumi:"authentication"`
	DataSource             *string `pulumi:"dataSource"`
	EncryptConnection      *bool   `pulumi:"encryptConnection"`
	Password               *string `pulumi:"password"`
	TrustServerCertificate *bool   `pulumi:"trustServerCertificate"`
	UserName               *string `pulumi:"userName"`
}





type SqlConnectionInformationInput interface {
	pulumi.Input

	ToSqlConnectionInformationOutput() SqlConnectionInformationOutput
	ToSqlConnectionInformationOutputWithContext(context.Context) SqlConnectionInformationOutput
}

type SqlConnectionInformationArgs struct {
	Authentication         pulumi.StringPtrInput `pulumi:"authentication"`
	DataSource             pulumi.StringPtrInput `pulumi:"dataSource"`
	EncryptConnection      pulumi.BoolPtrInput   `pulumi:"encryptConnection"`
	Password               pulumi.StringPtrInput `pulumi:"password"`
	TrustServerCertificate pulumi.BoolPtrInput   `pulumi:"trustServerCertificate"`
	UserName               pulumi.StringPtrInput `pulumi:"userName"`
}

func (SqlConnectionInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlConnectionInformation)(nil)).Elem()
}

func (i SqlConnectionInformationArgs) ToSqlConnectionInformationOutput() SqlConnectionInformationOutput {
	return i.ToSqlConnectionInformationOutputWithContext(context.Background())
}

func (i SqlConnectionInformationArgs) ToSqlConnectionInformationOutputWithContext(ctx context.Context) SqlConnectionInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlConnectionInformationOutput)
}

func (i SqlConnectionInformationArgs) ToSqlConnectionInformationPtrOutput() SqlConnectionInformationPtrOutput {
	return i.ToSqlConnectionInformationPtrOutputWithContext(context.Background())
}

func (i SqlConnectionInformationArgs) ToSqlConnectionInformationPtrOutputWithContext(ctx context.Context) SqlConnectionInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlConnectionInformationOutput).ToSqlConnectionInformationPtrOutputWithContext(ctx)
}









type SqlConnectionInformationPtrInput interface {
	pulumi.Input

	ToSqlConnectionInformationPtrOutput() SqlConnectionInformationPtrOutput
	ToSqlConnectionInformationPtrOutputWithContext(context.Context) SqlConnectionInformationPtrOutput
}

type sqlConnectionInformationPtrType SqlConnectionInformationArgs

func SqlConnectionInformationPtr(v *SqlConnectionInformationArgs) SqlConnectionInformationPtrInput {
	return (*sqlConnectionInformationPtrType)(v)
}

func (*sqlConnectionInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlConnectionInformation)(nil)).Elem()
}

func (i *sqlConnectionInformationPtrType) ToSqlConnectionInformationPtrOutput() SqlConnectionInformationPtrOutput {
	return i.ToSqlConnectionInformationPtrOutputWithContext(context.Background())
}

func (i *sqlConnectionInformationPtrType) ToSqlConnectionInformationPtrOutputWithContext(ctx context.Context) SqlConnectionInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlConnectionInformationPtrOutput)
}

type SqlConnectionInformationOutput struct{ *pulumi.OutputState }

func (SqlConnectionInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlConnectionInformation)(nil)).Elem()
}

func (o SqlConnectionInformationOutput) ToSqlConnectionInformationOutput() SqlConnectionInformationOutput {
	return o
}

func (o SqlConnectionInformationOutput) ToSqlConnectionInformationOutputWithContext(ctx context.Context) SqlConnectionInformationOutput {
	return o
}

func (o SqlConnectionInformationOutput) ToSqlConnectionInformationPtrOutput() SqlConnectionInformationPtrOutput {
	return o.ToSqlConnectionInformationPtrOutputWithContext(context.Background())
}

func (o SqlConnectionInformationOutput) ToSqlConnectionInformationPtrOutputWithContext(ctx context.Context) SqlConnectionInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlConnectionInformation) *SqlConnectionInformation {
		return &v
	}).(SqlConnectionInformationPtrOutput)
}

func (o SqlConnectionInformationOutput) Authentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInformation) *string { return v.Authentication }).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInformationOutput) DataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInformation) *string { return v.DataSource }).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInformationOutput) EncryptConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlConnectionInformation) *bool { return v.EncryptConnection }).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInformationOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInformation) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInformationOutput) TrustServerCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlConnectionInformation) *bool { return v.TrustServerCertificate }).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInformationOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInformation) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

type SqlConnectionInformationPtrOutput struct{ *pulumi.OutputState }

func (SqlConnectionInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlConnectionInformation)(nil)).Elem()
}

func (o SqlConnectionInformationPtrOutput) ToSqlConnectionInformationPtrOutput() SqlConnectionInformationPtrOutput {
	return o
}

func (o SqlConnectionInformationPtrOutput) ToSqlConnectionInformationPtrOutputWithContext(ctx context.Context) SqlConnectionInformationPtrOutput {
	return o
}

func (o SqlConnectionInformationPtrOutput) Elem() SqlConnectionInformationOutput {
	return o.ApplyT(func(v *SqlConnectionInformation) SqlConnectionInformation {
		if v != nil {
			return *v
		}
		var ret SqlConnectionInformation
		return ret
	}).(SqlConnectionInformationOutput)
}

func (o SqlConnectionInformationPtrOutput) Authentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInformation) *string {
		if v == nil {
			return nil
		}
		return v.Authentication
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInformationPtrOutput) DataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInformation) *string {
		if v == nil {
			return nil
		}
		return v.DataSource
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInformationPtrOutput) EncryptConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInformation) *bool {
		if v == nil {
			return nil
		}
		return v.EncryptConnection
	}).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInformationPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInformation) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInformationPtrOutput) TrustServerCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInformation) *bool {
		if v == nil {
			return nil
		}
		return v.TrustServerCertificate
	}).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInformationPtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInformation) *string {
		if v == nil {
			return nil
		}
		return v.UserName
	}).(pulumi.StringPtrOutput)
}

type SqlConnectionInformationResponse struct {
	Authentication         *string `pulumi:"authentication"`
	DataSource             *string `pulumi:"dataSource"`
	EncryptConnection      *bool   `pulumi:"encryptConnection"`
	Password               *string `pulumi:"password"`
	TrustServerCertificate *bool   `pulumi:"trustServerCertificate"`
	UserName               *string `pulumi:"userName"`
}

type SqlConnectionInformationResponseOutput struct{ *pulumi.OutputState }

func (SqlConnectionInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlConnectionInformationResponse)(nil)).Elem()
}

func (o SqlConnectionInformationResponseOutput) ToSqlConnectionInformationResponseOutput() SqlConnectionInformationResponseOutput {
	return o
}

func (o SqlConnectionInformationResponseOutput) ToSqlConnectionInformationResponseOutputWithContext(ctx context.Context) SqlConnectionInformationResponseOutput {
	return o
}

func (o SqlConnectionInformationResponseOutput) Authentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInformationResponse) *string { return v.Authentication }).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInformationResponseOutput) DataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInformationResponse) *string { return v.DataSource }).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInformationResponseOutput) EncryptConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlConnectionInformationResponse) *bool { return v.EncryptConnection }).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInformationResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInformationResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInformationResponseOutput) TrustServerCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlConnectionInformationResponse) *bool { return v.TrustServerCertificate }).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInformationResponseOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInformationResponse) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

type SqlConnectionInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlConnectionInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlConnectionInformationResponse)(nil)).Elem()
}

func (o SqlConnectionInformationResponsePtrOutput) ToSqlConnectionInformationResponsePtrOutput() SqlConnectionInformationResponsePtrOutput {
	return o
}

func (o SqlConnectionInformationResponsePtrOutput) ToSqlConnectionInformationResponsePtrOutputWithContext(ctx context.Context) SqlConnectionInformationResponsePtrOutput {
	return o
}

func (o SqlConnectionInformationResponsePtrOutput) Elem() SqlConnectionInformationResponseOutput {
	return o.ApplyT(func(v *SqlConnectionInformationResponse) SqlConnectionInformationResponse {
		if v != nil {
			return *v
		}
		var ret SqlConnectionInformationResponse
		return ret
	}).(SqlConnectionInformationResponseOutput)
}

func (o SqlConnectionInformationResponsePtrOutput) Authentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Authentication
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInformationResponsePtrOutput) DataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DataSource
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInformationResponsePtrOutput) EncryptConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInformationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EncryptConnection
	}).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInformationResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInformationResponsePtrOutput) TrustServerCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInformationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.TrustServerCertificate
	}).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInformationResponsePtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserName
	}).(pulumi.StringPtrOutput)
}

type SqlDbMigrationStatusDetailsResponse struct {
	ListOfCopyProgressDetails []CopyProgressDetailsResponse `pulumi:"listOfCopyProgressDetails"`
	MigrationState            string                        `pulumi:"migrationState"`
	SqlDataCopyErrors         []string                      `pulumi:"sqlDataCopyErrors"`
}

type SqlDbMigrationStatusDetailsResponseOutput struct{ *pulumi.OutputState }

func (SqlDbMigrationStatusDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDbMigrationStatusDetailsResponse)(nil)).Elem()
}

func (o SqlDbMigrationStatusDetailsResponseOutput) ToSqlDbMigrationStatusDetailsResponseOutput() SqlDbMigrationStatusDetailsResponseOutput {
	return o
}

func (o SqlDbMigrationStatusDetailsResponseOutput) ToSqlDbMigrationStatusDetailsResponseOutputWithContext(ctx context.Context) SqlDbMigrationStatusDetailsResponseOutput {
	return o
}

func (o SqlDbMigrationStatusDetailsResponseOutput) ListOfCopyProgressDetails() CopyProgressDetailsResponseArrayOutput {
	return o.ApplyT(func(v SqlDbMigrationStatusDetailsResponse) []CopyProgressDetailsResponse {
		return v.ListOfCopyProgressDetails
	}).(CopyProgressDetailsResponseArrayOutput)
}

func (o SqlDbMigrationStatusDetailsResponseOutput) MigrationState() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDbMigrationStatusDetailsResponse) string { return v.MigrationState }).(pulumi.StringOutput)
}

func (o SqlDbMigrationStatusDetailsResponseOutput) SqlDataCopyErrors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SqlDbMigrationStatusDetailsResponse) []string { return v.SqlDataCopyErrors }).(pulumi.StringArrayOutput)
}

type SqlDbOfflineConfigurationResponse struct {
	Offline bool `pulumi:"offline"`
}

type SqlDbOfflineConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SqlDbOfflineConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDbOfflineConfigurationResponse)(nil)).Elem()
}

func (o SqlDbOfflineConfigurationResponseOutput) ToSqlDbOfflineConfigurationResponseOutput() SqlDbOfflineConfigurationResponseOutput {
	return o
}

func (o SqlDbOfflineConfigurationResponseOutput) ToSqlDbOfflineConfigurationResponseOutputWithContext(ctx context.Context) SqlDbOfflineConfigurationResponseOutput {
	return o
}

func (o SqlDbOfflineConfigurationResponseOutput) Offline() pulumi.BoolOutput {
	return o.ApplyT(func(v SqlDbOfflineConfigurationResponse) bool { return v.Offline }).(pulumi.BoolOutput)
}

type StartMigrationScenarioServerRoleResultResponse struct {
	ExceptionsAndWarnings []ReportableExceptionResponse `pulumi:"exceptionsAndWarnings"`
	Name                  string                        `pulumi:"name"`
	State                 string                        `pulumi:"state"`
}

type SyncMigrationDatabaseErrorEventResponse struct {
	EventText       string `pulumi:"eventText"`
	EventTypeString string `pulumi:"eventTypeString"`
	TimestampString string `pulumi:"timestampString"`
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type ValidateMigrationInputSqlServerSqlDbSyncTaskProperties struct {
	Input    *ValidateSyncMigrationInputSqlServerTaskInput `pulumi:"input"`
	TaskType string                                        `pulumi:"taskType"`
}


func (val *ValidateMigrationInputSqlServerSqlDbSyncTaskProperties) Defaults() *ValidateMigrationInputSqlServerSqlDbSyncTaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ValidateMigrationInputSqlServerSqlDbSyncTaskPropertiesResponse struct {
	Commands []interface{}                                           `pulumi:"commands"`
	Errors   []ODataErrorResponse                                    `pulumi:"errors"`
	Input    *ValidateSyncMigrationInputSqlServerTaskInputResponse   `pulumi:"input"`
	Output   []ValidateSyncMigrationInputSqlServerTaskOutputResponse `pulumi:"output"`
	State    string                                                  `pulumi:"state"`
	TaskType string                                                  `pulumi:"taskType"`
}


func (val *ValidateMigrationInputSqlServerSqlDbSyncTaskPropertiesResponse) Defaults() *ValidateMigrationInputSqlServerSqlDbSyncTaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ValidateMigrationInputSqlServerSqlMISyncTaskInput struct {
	AzureApp             AzureActiveDirectoryApp              `pulumi:"azureApp"`
	BackupFileShare      *FileShare                           `pulumi:"backupFileShare"`
	SelectedDatabases    []MigrateSqlServerSqlMIDatabaseInput `pulumi:"selectedDatabases"`
	SourceConnectionInfo SqlConnectionInfo                    `pulumi:"sourceConnectionInfo"`
	StorageResourceId    string                               `pulumi:"storageResourceId"`
	TargetConnectionInfo MiSqlConnectionInfo                  `pulumi:"targetConnectionInfo"`
}


func (val *ValidateMigrationInputSqlServerSqlMISyncTaskInput) Defaults() *ValidateMigrationInputSqlServerSqlMISyncTaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	return &tmp
}

type ValidateMigrationInputSqlServerSqlMISyncTaskInputResponse struct {
	AzureApp             AzureActiveDirectoryAppResponse              `pulumi:"azureApp"`
	BackupFileShare      *FileShareResponse                           `pulumi:"backupFileShare"`
	SelectedDatabases    []MigrateSqlServerSqlMIDatabaseInputResponse `pulumi:"selectedDatabases"`
	SourceConnectionInfo SqlConnectionInfoResponse                    `pulumi:"sourceConnectionInfo"`
	StorageResourceId    string                                       `pulumi:"storageResourceId"`
	TargetConnectionInfo MiSqlConnectionInfoResponse                  `pulumi:"targetConnectionInfo"`
}


func (val *ValidateMigrationInputSqlServerSqlMISyncTaskInputResponse) Defaults() *ValidateMigrationInputSqlServerSqlMISyncTaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	return &tmp
}

type ValidateMigrationInputSqlServerSqlMISyncTaskOutputResponse struct {
	Id               string                        `pulumi:"id"`
	Name             string                        `pulumi:"name"`
	ValidationErrors []ReportableExceptionResponse `pulumi:"validationErrors"`
}

type ValidateMigrationInputSqlServerSqlMISyncTaskProperties struct {
	Input    *ValidateMigrationInputSqlServerSqlMISyncTaskInput `pulumi:"input"`
	TaskType string                                             `pulumi:"taskType"`
}


func (val *ValidateMigrationInputSqlServerSqlMISyncTaskProperties) Defaults() *ValidateMigrationInputSqlServerSqlMISyncTaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ValidateMigrationInputSqlServerSqlMISyncTaskPropertiesResponse struct {
	Commands []interface{}                                                `pulumi:"commands"`
	Errors   []ODataErrorResponse                                         `pulumi:"errors"`
	Input    *ValidateMigrationInputSqlServerSqlMISyncTaskInputResponse   `pulumi:"input"`
	Output   []ValidateMigrationInputSqlServerSqlMISyncTaskOutputResponse `pulumi:"output"`
	State    string                                                       `pulumi:"state"`
	TaskType string                                                       `pulumi:"taskType"`
}


func (val *ValidateMigrationInputSqlServerSqlMISyncTaskPropertiesResponse) Defaults() *ValidateMigrationInputSqlServerSqlMISyncTaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ValidateMigrationInputSqlServerSqlMITaskInput struct {
	BackupBlobShare      BlobShare                            `pulumi:"backupBlobShare"`
	BackupFileShare      *FileShare                           `pulumi:"backupFileShare"`
	BackupMode           *string                              `pulumi:"backupMode"`
	SelectedDatabases    []MigrateSqlServerSqlMIDatabaseInput `pulumi:"selectedDatabases"`
	SelectedLogins       []string                             `pulumi:"selectedLogins"`
	SourceConnectionInfo SqlConnectionInfo                    `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo SqlConnectionInfo                    `pulumi:"targetConnectionInfo"`
}


func (val *ValidateMigrationInputSqlServerSqlMITaskInput) Defaults() *ValidateMigrationInputSqlServerSqlMITaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type ValidateMigrationInputSqlServerSqlMITaskInputResponse struct {
	BackupBlobShare      BlobShareResponse                            `pulumi:"backupBlobShare"`
	BackupFileShare      *FileShareResponse                           `pulumi:"backupFileShare"`
	BackupMode           *string                                      `pulumi:"backupMode"`
	SelectedDatabases    []MigrateSqlServerSqlMIDatabaseInputResponse `pulumi:"selectedDatabases"`
	SelectedLogins       []string                                     `pulumi:"selectedLogins"`
	SourceConnectionInfo SqlConnectionInfoResponse                    `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo SqlConnectionInfoResponse                    `pulumi:"targetConnectionInfo"`
}


func (val *ValidateMigrationInputSqlServerSqlMITaskInputResponse) Defaults() *ValidateMigrationInputSqlServerSqlMITaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type ValidateMigrationInputSqlServerSqlMITaskOutputResponse struct {
	BackupFolderErrors           []ReportableExceptionResponse `pulumi:"backupFolderErrors"`
	BackupShareCredentialsErrors []ReportableExceptionResponse `pulumi:"backupShareCredentialsErrors"`
	BackupStorageAccountErrors   []ReportableExceptionResponse `pulumi:"backupStorageAccountErrors"`
	DatabaseBackupInfo           *DatabaseBackupInfoResponse   `pulumi:"databaseBackupInfo"`
	ExistingBackupErrors         []ReportableExceptionResponse `pulumi:"existingBackupErrors"`
	Id                           string                        `pulumi:"id"`
	Name                         string                        `pulumi:"name"`
	RestoreDatabaseNameErrors    []ReportableExceptionResponse `pulumi:"restoreDatabaseNameErrors"`
}

type ValidateMigrationInputSqlServerSqlMITaskProperties struct {
	Input    *ValidateMigrationInputSqlServerSqlMITaskInput `pulumi:"input"`
	TaskType string                                         `pulumi:"taskType"`
}


func (val *ValidateMigrationInputSqlServerSqlMITaskProperties) Defaults() *ValidateMigrationInputSqlServerSqlMITaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ValidateMigrationInputSqlServerSqlMITaskPropertiesResponse struct {
	Commands []interface{}                                            `pulumi:"commands"`
	Errors   []ODataErrorResponse                                     `pulumi:"errors"`
	Input    *ValidateMigrationInputSqlServerSqlMITaskInputResponse   `pulumi:"input"`
	Output   []ValidateMigrationInputSqlServerSqlMITaskOutputResponse `pulumi:"output"`
	State    string                                                   `pulumi:"state"`
	TaskType string                                                   `pulumi:"taskType"`
}


func (val *ValidateMigrationInputSqlServerSqlMITaskPropertiesResponse) Defaults() *ValidateMigrationInputSqlServerSqlMITaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Input = tmp.Input.Defaults()

	return &tmp
}

type ValidateSyncMigrationInputSqlServerTaskInput struct {
	SelectedDatabases    []MigrateSqlServerSqlDbSyncDatabaseInput `pulumi:"selectedDatabases"`
	SourceConnectionInfo SqlConnectionInfo                        `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo SqlConnectionInfo                        `pulumi:"targetConnectionInfo"`
}


func (val *ValidateSyncMigrationInputSqlServerTaskInput) Defaults() *ValidateSyncMigrationInputSqlServerTaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type ValidateSyncMigrationInputSqlServerTaskInputResponse struct {
	SelectedDatabases    []MigrateSqlServerSqlDbSyncDatabaseInputResponse `pulumi:"selectedDatabases"`
	SourceConnectionInfo SqlConnectionInfoResponse                        `pulumi:"sourceConnectionInfo"`
	TargetConnectionInfo SqlConnectionInfoResponse                        `pulumi:"targetConnectionInfo"`
}


func (val *ValidateSyncMigrationInputSqlServerTaskInputResponse) Defaults() *ValidateSyncMigrationInputSqlServerTaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = *tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type ValidateSyncMigrationInputSqlServerTaskOutputResponse struct {
	Id               string                        `pulumi:"id"`
	Name             string                        `pulumi:"name"`
	ValidationErrors []ReportableExceptionResponse `pulumi:"validationErrors"`
}

type ValidationErrorResponse struct {
	Severity string `pulumi:"severity"`
	Text     string `pulumi:"text"`
}

type WaitStatisticsResponse struct {
	WaitCount  float64 `pulumi:"waitCount"`
	WaitTimeMs float64 `pulumi:"waitTimeMs"`
	WaitType   string  `pulumi:"waitType"`
}


func (val *WaitStatisticsResponse) Defaults() *WaitStatisticsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.WaitTimeMs) {
		tmp.WaitTimeMs = 0.0
	}
	return &tmp
}
func init() {
	pulumi.RegisterOutputType(CopyProgressDetailsResponseOutput{})
	pulumi.RegisterOutputType(CopyProgressDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseInfoOutput{})
	pulumi.RegisterOutputType(DatabaseInfoArrayOutput{})
	pulumi.RegisterOutputType(DatabaseInfoResponseOutput{})
	pulumi.RegisterOutputType(DatabaseInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseMigrationPropertiesSqlDbOutput{})
	pulumi.RegisterOutputType(DatabaseMigrationPropertiesSqlDbPtrOutput{})
	pulumi.RegisterOutputType(DatabaseMigrationPropertiesSqlDbResponseOutput{})
	pulumi.RegisterOutputType(ErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(NodeMonitoringDataResponseOutput{})
	pulumi.RegisterOutputType(NodeMonitoringDataResponseArrayOutput{})
	pulumi.RegisterOutputType(ProjectFilePropertiesOutput{})
	pulumi.RegisterOutputType(ProjectFilePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ProjectFilePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServiceSkuOutput{})
	pulumi.RegisterOutputType(ServiceSkuPtrOutput{})
	pulumi.RegisterOutputType(ServiceSkuResponseOutput{})
	pulumi.RegisterOutputType(ServiceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlConnectionInformationOutput{})
	pulumi.RegisterOutputType(SqlConnectionInformationPtrOutput{})
	pulumi.RegisterOutputType(SqlConnectionInformationResponseOutput{})
	pulumi.RegisterOutputType(SqlConnectionInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlDbMigrationStatusDetailsResponseOutput{})
	pulumi.RegisterOutputType(SqlDbOfflineConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
