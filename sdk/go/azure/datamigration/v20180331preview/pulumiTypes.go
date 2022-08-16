


package v20180331preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlobShare struct {
	SasUri string `pulumi:"sasUri"`
}

type BlobShareResponse struct {
	SasUri string `pulumi:"sasUri"`
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
	Id         string `pulumi:"id"`
	ResultType string `pulumi:"resultType"`
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
	MigrationReportResult    *MigrationReportResultResponse           `pulumi:"migrationReportResult"`
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
	Id         string `pulumi:"id"`
	ResultType string `pulumi:"resultType"`
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
	BackupFileShare     *FileShare `pulumi:"backupFileShare"`
	Name                string     `pulumi:"name"`
	RestoreDatabaseName string     `pulumi:"restoreDatabaseName"`
}

type MigrateSqlServerSqlMIDatabaseInputResponse struct {
	BackupFileShare     *FileShareResponse `pulumi:"backupFileShare"`
	Name                string             `pulumi:"name"`
	RestoreDatabaseName string             `pulumi:"restoreDatabaseName"`
}

type MigrateSqlServerSqlMITaskInput struct {
	BackupBlobShare      BlobShare                            `pulumi:"backupBlobShare"`
	BackupFileShare      *FileShare                           `pulumi:"backupFileShare"`
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
	OrphanedUsers            map[string]string                                         `pulumi:"orphanedUsers"`
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

type MigrationEligibilityInfoResponse struct {
	IsEligibleForMigration bool     `pulumi:"isEligibleForMigration"`
	ValidationMessages     []string `pulumi:"validationMessages"`
}

type MigrationReportResultResponse struct {
	Id        *string `pulumi:"id"`
	ReportUrl *string `pulumi:"reportUrl"`
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

type ODataErrorResponse struct {
	Code    *string              `pulumi:"code"`
	Details []ODataErrorResponse `pulumi:"details"`
	Message *string              `pulumi:"message"`
}

type ReportableExceptionResponse struct {
	FilePath   *string `pulumi:"filePath"`
	HResult    *int    `pulumi:"hResult"`
	LineNumber *string `pulumi:"lineNumber"`
	Message    *string `pulumi:"message"`
	StackTrace *string `pulumi:"stackTrace"`
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





type SqlConnectionInfoInput interface {
	pulumi.Input

	ToSqlConnectionInfoOutput() SqlConnectionInfoOutput
	ToSqlConnectionInfoOutputWithContext(context.Context) SqlConnectionInfoOutput
}

type SqlConnectionInfoArgs struct {
	AdditionalSettings     pulumi.StringPtrInput `pulumi:"additionalSettings"`
	Authentication         pulumi.StringPtrInput `pulumi:"authentication"`
	DataSource             pulumi.StringInput    `pulumi:"dataSource"`
	EncryptConnection      pulumi.BoolPtrInput   `pulumi:"encryptConnection"`
	Password               pulumi.StringPtrInput `pulumi:"password"`
	TrustServerCertificate pulumi.BoolPtrInput   `pulumi:"trustServerCertificate"`
	Type                   pulumi.StringInput    `pulumi:"type"`
	UserName               pulumi.StringPtrInput `pulumi:"userName"`
}


func (val *SqlConnectionInfoArgs) Defaults() *SqlConnectionInfoArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EncryptConnection) {
		tmp.EncryptConnection = pulumi.BoolPtr(true)
	}
	if isZero(tmp.TrustServerCertificate) {
		tmp.TrustServerCertificate = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (SqlConnectionInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlConnectionInfo)(nil)).Elem()
}

func (i SqlConnectionInfoArgs) ToSqlConnectionInfoOutput() SqlConnectionInfoOutput {
	return i.ToSqlConnectionInfoOutputWithContext(context.Background())
}

func (i SqlConnectionInfoArgs) ToSqlConnectionInfoOutputWithContext(ctx context.Context) SqlConnectionInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlConnectionInfoOutput)
}

func (i SqlConnectionInfoArgs) ToSqlConnectionInfoPtrOutput() SqlConnectionInfoPtrOutput {
	return i.ToSqlConnectionInfoPtrOutputWithContext(context.Background())
}

func (i SqlConnectionInfoArgs) ToSqlConnectionInfoPtrOutputWithContext(ctx context.Context) SqlConnectionInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlConnectionInfoOutput).ToSqlConnectionInfoPtrOutputWithContext(ctx)
}









type SqlConnectionInfoPtrInput interface {
	pulumi.Input

	ToSqlConnectionInfoPtrOutput() SqlConnectionInfoPtrOutput
	ToSqlConnectionInfoPtrOutputWithContext(context.Context) SqlConnectionInfoPtrOutput
}

type sqlConnectionInfoPtrType SqlConnectionInfoArgs

func SqlConnectionInfoPtr(v *SqlConnectionInfoArgs) SqlConnectionInfoPtrInput {
	return (*sqlConnectionInfoPtrType)(v)
}

func (*sqlConnectionInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlConnectionInfo)(nil)).Elem()
}

func (i *sqlConnectionInfoPtrType) ToSqlConnectionInfoPtrOutput() SqlConnectionInfoPtrOutput {
	return i.ToSqlConnectionInfoPtrOutputWithContext(context.Background())
}

func (i *sqlConnectionInfoPtrType) ToSqlConnectionInfoPtrOutputWithContext(ctx context.Context) SqlConnectionInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlConnectionInfoPtrOutput)
}

type SqlConnectionInfoOutput struct{ *pulumi.OutputState }

func (SqlConnectionInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlConnectionInfo)(nil)).Elem()
}

func (o SqlConnectionInfoOutput) ToSqlConnectionInfoOutput() SqlConnectionInfoOutput {
	return o
}

func (o SqlConnectionInfoOutput) ToSqlConnectionInfoOutputWithContext(ctx context.Context) SqlConnectionInfoOutput {
	return o
}

func (o SqlConnectionInfoOutput) ToSqlConnectionInfoPtrOutput() SqlConnectionInfoPtrOutput {
	return o.ToSqlConnectionInfoPtrOutputWithContext(context.Background())
}

func (o SqlConnectionInfoOutput) ToSqlConnectionInfoPtrOutputWithContext(ctx context.Context) SqlConnectionInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlConnectionInfo) *SqlConnectionInfo {
		return &v
	}).(SqlConnectionInfoPtrOutput)
}

func (o SqlConnectionInfoOutput) AdditionalSettings() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInfo) *string { return v.AdditionalSettings }).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoOutput) Authentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInfo) *string { return v.Authentication }).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoOutput) DataSource() pulumi.StringOutput {
	return o.ApplyT(func(v SqlConnectionInfo) string { return v.DataSource }).(pulumi.StringOutput)
}

func (o SqlConnectionInfoOutput) EncryptConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlConnectionInfo) *bool { return v.EncryptConnection }).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInfoOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInfo) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoOutput) TrustServerCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlConnectionInfo) *bool { return v.TrustServerCertificate }).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInfoOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SqlConnectionInfo) string { return v.Type }).(pulumi.StringOutput)
}

func (o SqlConnectionInfoOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInfo) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

type SqlConnectionInfoPtrOutput struct{ *pulumi.OutputState }

func (SqlConnectionInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlConnectionInfo)(nil)).Elem()
}

func (o SqlConnectionInfoPtrOutput) ToSqlConnectionInfoPtrOutput() SqlConnectionInfoPtrOutput {
	return o
}

func (o SqlConnectionInfoPtrOutput) ToSqlConnectionInfoPtrOutputWithContext(ctx context.Context) SqlConnectionInfoPtrOutput {
	return o
}

func (o SqlConnectionInfoPtrOutput) Elem() SqlConnectionInfoOutput {
	return o.ApplyT(func(v *SqlConnectionInfo) SqlConnectionInfo {
		if v != nil {
			return *v
		}
		var ret SqlConnectionInfo
		return ret
	}).(SqlConnectionInfoOutput)
}

func (o SqlConnectionInfoPtrOutput) AdditionalSettings() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfo) *string {
		if v == nil {
			return nil
		}
		return v.AdditionalSettings
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoPtrOutput) Authentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfo) *string {
		if v == nil {
			return nil
		}
		return v.Authentication
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoPtrOutput) DataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfo) *string {
		if v == nil {
			return nil
		}
		return &v.DataSource
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoPtrOutput) EncryptConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfo) *bool {
		if v == nil {
			return nil
		}
		return v.EncryptConnection
	}).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInfoPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfo) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoPtrOutput) TrustServerCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfo) *bool {
		if v == nil {
			return nil
		}
		return v.TrustServerCertificate
	}).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInfoPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfo) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoPtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfo) *string {
		if v == nil {
			return nil
		}
		return v.UserName
	}).(pulumi.StringPtrOutput)
}

type SqlConnectionInfoResponse struct {
	AdditionalSettings     *string `pulumi:"additionalSettings"`
	Authentication         *string `pulumi:"authentication"`
	DataSource             string  `pulumi:"dataSource"`
	EncryptConnection      *bool   `pulumi:"encryptConnection"`
	Password               *string `pulumi:"password"`
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

type SqlConnectionInfoResponseOutput struct{ *pulumi.OutputState }

func (SqlConnectionInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlConnectionInfoResponse)(nil)).Elem()
}

func (o SqlConnectionInfoResponseOutput) ToSqlConnectionInfoResponseOutput() SqlConnectionInfoResponseOutput {
	return o
}

func (o SqlConnectionInfoResponseOutput) ToSqlConnectionInfoResponseOutputWithContext(ctx context.Context) SqlConnectionInfoResponseOutput {
	return o
}

func (o SqlConnectionInfoResponseOutput) AdditionalSettings() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInfoResponse) *string { return v.AdditionalSettings }).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoResponseOutput) Authentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInfoResponse) *string { return v.Authentication }).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoResponseOutput) DataSource() pulumi.StringOutput {
	return o.ApplyT(func(v SqlConnectionInfoResponse) string { return v.DataSource }).(pulumi.StringOutput)
}

func (o SqlConnectionInfoResponseOutput) EncryptConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlConnectionInfoResponse) *bool { return v.EncryptConnection }).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInfoResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInfoResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoResponseOutput) TrustServerCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlConnectionInfoResponse) *bool { return v.TrustServerCertificate }).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SqlConnectionInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o SqlConnectionInfoResponseOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectionInfoResponse) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

type SqlConnectionInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlConnectionInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlConnectionInfoResponse)(nil)).Elem()
}

func (o SqlConnectionInfoResponsePtrOutput) ToSqlConnectionInfoResponsePtrOutput() SqlConnectionInfoResponsePtrOutput {
	return o
}

func (o SqlConnectionInfoResponsePtrOutput) ToSqlConnectionInfoResponsePtrOutputWithContext(ctx context.Context) SqlConnectionInfoResponsePtrOutput {
	return o
}

func (o SqlConnectionInfoResponsePtrOutput) Elem() SqlConnectionInfoResponseOutput {
	return o.ApplyT(func(v *SqlConnectionInfoResponse) SqlConnectionInfoResponse {
		if v != nil {
			return *v
		}
		var ret SqlConnectionInfoResponse
		return ret
	}).(SqlConnectionInfoResponseOutput)
}

func (o SqlConnectionInfoResponsePtrOutput) AdditionalSettings() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdditionalSettings
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoResponsePtrOutput) Authentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Authentication
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoResponsePtrOutput) DataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DataSource
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoResponsePtrOutput) EncryptConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EncryptConnection
	}).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInfoResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoResponsePtrOutput) TrustServerCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.TrustServerCertificate
	}).(pulumi.BoolPtrOutput)
}

func (o SqlConnectionInfoResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectionInfoResponsePtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectionInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserName
	}).(pulumi.StringPtrOutput)
}

type StartMigrationScenarioServerRoleResultResponse struct {
	ExceptionsAndWarnings []ReportableExceptionResponse `pulumi:"exceptionsAndWarnings"`
	Name                  string                        `pulumi:"name"`
	State                 string                        `pulumi:"state"`
}

type ValidateMigrationInputSqlServerSqlMITaskInput struct {
	BackupBlobShare      BlobShare                            `pulumi:"backupBlobShare"`
	BackupFileShare      *FileShare                           `pulumi:"backupFileShare"`
	SelectedDatabases    []MigrateSqlServerSqlMIDatabaseInput `pulumi:"selectedDatabases"`
	TargetConnectionInfo SqlConnectionInfo                    `pulumi:"targetConnectionInfo"`
}


func (val *ValidateMigrationInputSqlServerSqlMITaskInput) Defaults() *ValidateMigrationInputSqlServerSqlMITaskInput {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type ValidateMigrationInputSqlServerSqlMITaskInputResponse struct {
	BackupBlobShare      BlobShareResponse                            `pulumi:"backupBlobShare"`
	BackupFileShare      *FileShareResponse                           `pulumi:"backupFileShare"`
	SelectedDatabases    []MigrateSqlServerSqlMIDatabaseInputResponse `pulumi:"selectedDatabases"`
	TargetConnectionInfo SqlConnectionInfoResponse                    `pulumi:"targetConnectionInfo"`
}


func (val *ValidateMigrationInputSqlServerSqlMITaskInputResponse) Defaults() *ValidateMigrationInputSqlServerSqlMITaskInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.TargetConnectionInfo = *tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

type ValidateMigrationInputSqlServerSqlMITaskOutputResponse struct {
	BackupFolderErrors           []ReportableExceptionResponse `pulumi:"backupFolderErrors"`
	BackupShareCredentialsErrors []ReportableExceptionResponse `pulumi:"backupShareCredentialsErrors"`
	BackupStorageAccountErrors   []ReportableExceptionResponse `pulumi:"backupStorageAccountErrors"`
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
func init() {
	pulumi.RegisterOutputType(DatabaseInfoOutput{})
	pulumi.RegisterOutputType(DatabaseInfoArrayOutput{})
	pulumi.RegisterOutputType(DatabaseInfoResponseOutput{})
	pulumi.RegisterOutputType(DatabaseInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceSkuOutput{})
	pulumi.RegisterOutputType(ServiceSkuPtrOutput{})
	pulumi.RegisterOutputType(ServiceSkuResponseOutput{})
	pulumi.RegisterOutputType(ServiceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlConnectionInfoOutput{})
	pulumi.RegisterOutputType(SqlConnectionInfoPtrOutput{})
	pulumi.RegisterOutputType(SqlConnectionInfoResponseOutput{})
	pulumi.RegisterOutputType(SqlConnectionInfoResponsePtrOutput{})
}
