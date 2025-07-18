// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate
{
    public static class GetSqlAssessmentV2Operation
    {
        /// <summary>
        /// Get a SqlAssessmentV2
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-03-15, 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-01-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetSqlAssessmentV2OperationResult> InvokeAsync(GetSqlAssessmentV2OperationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSqlAssessmentV2OperationResult>("azure-native:migrate:getSqlAssessmentV2Operation", args ?? new GetSqlAssessmentV2OperationArgs(), options.WithDefaults());

        /// <summary>
        /// Get a SqlAssessmentV2
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-03-15, 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-01-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSqlAssessmentV2OperationResult> Invoke(GetSqlAssessmentV2OperationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSqlAssessmentV2OperationResult>("azure-native:migrate:getSqlAssessmentV2Operation", args ?? new GetSqlAssessmentV2OperationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a SqlAssessmentV2
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-03-15, 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-01-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSqlAssessmentV2OperationResult> Invoke(GetSqlAssessmentV2OperationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSqlAssessmentV2OperationResult>("azure-native:migrate:getSqlAssessmentV2Operation", args ?? new GetSqlAssessmentV2OperationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSqlAssessmentV2OperationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// SQL Assessment arm name.
        /// </summary>
        [Input("assessmentName", required: true)]
        public string AssessmentName { get; set; } = null!;

        /// <summary>
        /// Group ARM name
        /// </summary>
        [Input("groupName", required: true)]
        public string GroupName { get; set; } = null!;

        /// <summary>
        /// Assessment Project Name
        /// </summary>
        [Input("projectName", required: true)]
        public string ProjectName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSqlAssessmentV2OperationArgs()
        {
        }
        public static new GetSqlAssessmentV2OperationArgs Empty => new GetSqlAssessmentV2OperationArgs();
    }

    public sealed class GetSqlAssessmentV2OperationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// SQL Assessment arm name.
        /// </summary>
        [Input("assessmentName", required: true)]
        public Input<string> AssessmentName { get; set; } = null!;

        /// <summary>
        /// Group ARM name
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// Assessment Project Name
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetSqlAssessmentV2OperationInvokeArgs()
        {
        }
        public static new GetSqlAssessmentV2OperationInvokeArgs Empty => new GetSqlAssessmentV2OperationInvokeArgs();
    }


    [OutputType]
    public sealed class GetSqlAssessmentV2OperationResult
    {
        /// <summary>
        /// Assessment type of the assessment.
        /// </summary>
        public readonly string? AssessmentType;
        /// <summary>
        /// Gets or sets user preference indicating intent of async commit mode.
        /// </summary>
        public readonly string? AsyncCommitModeIntent;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Azure Location or Azure region where to which the machines will be migrated.
        /// </summary>
        public readonly string? AzureLocation;
        /// <summary>
        /// Azure Offer Code.
        /// </summary>
        public readonly string? AzureOfferCode;
        /// <summary>
        /// Gets or sets Azure Offer Code for VM.
        /// </summary>
        public readonly string? AzureOfferCodeForVm;
        /// <summary>
        /// Gets or sets a value indicating azure security offering type.
        /// </summary>
        public readonly string? AzureSecurityOfferingType;
        /// <summary>
        /// Gets or sets user configurable SQL database settings.
        /// </summary>
        public readonly Outputs.SqlDbSettingsResponse? AzureSqlDatabaseSettings;
        /// <summary>
        /// Gets or sets user configurable SQL managed instance settings.
        /// </summary>
        public readonly Outputs.SqlMiSettingsResponse? AzureSqlManagedInstanceSettings;
        /// <summary>
        /// Gets or sets user configurable SQL VM settings.
        /// </summary>
        public readonly Outputs.SqlVmSettingsResponse? AzureSqlVmSettings;
        /// <summary>
        /// Confidence Rating in Percentage.
        /// </summary>
        public readonly double? ConfidenceRatingInPercentage;
        /// <summary>
        /// Date and Time when assessment was created.
        /// </summary>
        public readonly string CreatedTimestamp;
        /// <summary>
        /// Currency in which prices should be reported.
        /// </summary>
        public readonly string? Currency;
        /// <summary>
        /// Gets or sets the Azure Location or Azure region where to which the machines
        /// will be migrated.
        /// </summary>
        public readonly string? DisasterRecoveryLocation;
        /// <summary>
        /// Custom discount percentage.
        /// </summary>
        public readonly double? DiscountPercentage;
        /// <summary>
        /// Gets or sets the Enterprise agreement subscription id.
        /// </summary>
        public readonly string? EaSubscriptionId;
        /// <summary>
        /// Gets or sets a value indicating whether HADR assessments needs to be created.
        /// </summary>
        public readonly bool? EnableHadrAssessment;
        /// <summary>
        /// Gets or sets the duration for which the entity (SQL, VMs) are up in the
        /// on-premises environment.
        /// </summary>
        public readonly Outputs.EntityUptimeResponse? EntityUptime;
        /// <summary>
        /// Gets or sets user configurable setting to display the environment type.
        /// </summary>
        public readonly string? EnvironmentType;
        /// <summary>
        /// Gets the group type for the assessment.
        /// </summary>
        public readonly string? GroupType;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Gets or sets a value indicating whether internet access is available.
        /// </summary>
        public readonly bool? IsInternetAccessAvailable;
        /// <summary>
        /// Gets or sets user preference indicating intent of multi-subnet configuration.
        /// </summary>
        public readonly string? MultiSubnetIntent;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets SQL optimization logic.
        /// </summary>
        public readonly string? OptimizationLogic;
        /// <summary>
        /// Gets or sets user configurable setting to display the azure hybrid use benefit.
        /// </summary>
        public readonly string? OsLicense;
        /// <summary>
        /// Percentile of the utilization data values to be considered while assessing
        /// machines.
        /// </summary>
        public readonly string? Percentile;
        /// <summary>
        /// Gets or sets the end time to consider performance data for assessment.
        /// </summary>
        public readonly string? PerfDataEndTime;
        /// <summary>
        /// Gets or sets the start time to consider performance data for assessment.
        /// </summary>
        public readonly string? PerfDataStartTime;
        /// <summary>
        /// Last time when rates were queried.
        /// </summary>
        public readonly string PricesTimestamp;
        /// <summary>
        /// The status of the last operation.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Reserved instance.
        /// </summary>
        public readonly string? ReservedInstance;
        /// <summary>
        /// Gets or sets azure reserved instance for VM.
        /// </summary>
        public readonly string? ReservedInstanceForVm;
        /// <summary>
        /// Percentage of buffer that user wants on performance metrics when recommending
        /// Azure sizes.
        /// </summary>
        public readonly double? ScalingFactor;
        /// <summary>
        /// Schema version.
        /// </summary>
        public readonly string SchemaVersion;
        /// <summary>
        /// Assessment sizing criterion.
        /// </summary>
        public readonly string? SizingCriterion;
        /// <summary>
        /// SQL server license.
        /// </summary>
        public readonly string? SqlServerLicense;
        /// <summary>
        /// User configurable setting to display the Stage of Assessment.
        /// </summary>
        public readonly string Stage;
        /// <summary>
        /// Whether assessment is in valid state and all machines have been assessed.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Time Range for which the historic utilization data should be considered for
        /// assessment.
        /// </summary>
        public readonly string? TimeRange;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Date and Time when assessment was last updated.
        /// </summary>
        public readonly string UpdatedTimestamp;

        [OutputConstructor]
        private GetSqlAssessmentV2OperationResult(
            string? assessmentType,

            string? asyncCommitModeIntent,

            string azureApiVersion,

            string? azureLocation,

            string? azureOfferCode,

            string? azureOfferCodeForVm,

            string? azureSecurityOfferingType,

            Outputs.SqlDbSettingsResponse? azureSqlDatabaseSettings,

            Outputs.SqlMiSettingsResponse? azureSqlManagedInstanceSettings,

            Outputs.SqlVmSettingsResponse? azureSqlVmSettings,

            double? confidenceRatingInPercentage,

            string createdTimestamp,

            string? currency,

            string? disasterRecoveryLocation,

            double? discountPercentage,

            string? eaSubscriptionId,

            bool? enableHadrAssessment,

            Outputs.EntityUptimeResponse? entityUptime,

            string? environmentType,

            string? groupType,

            string id,

            bool? isInternetAccessAvailable,

            string? multiSubnetIntent,

            string name,

            string? optimizationLogic,

            string? osLicense,

            string? percentile,

            string? perfDataEndTime,

            string? perfDataStartTime,

            string pricesTimestamp,

            string? provisioningState,

            string? reservedInstance,

            string? reservedInstanceForVm,

            double? scalingFactor,

            string schemaVersion,

            string? sizingCriterion,

            string? sqlServerLicense,

            string stage,

            string status,

            Outputs.SystemDataResponse systemData,

            string? timeRange,

            string type,

            string updatedTimestamp)
        {
            AssessmentType = assessmentType;
            AsyncCommitModeIntent = asyncCommitModeIntent;
            AzureApiVersion = azureApiVersion;
            AzureLocation = azureLocation;
            AzureOfferCode = azureOfferCode;
            AzureOfferCodeForVm = azureOfferCodeForVm;
            AzureSecurityOfferingType = azureSecurityOfferingType;
            AzureSqlDatabaseSettings = azureSqlDatabaseSettings;
            AzureSqlManagedInstanceSettings = azureSqlManagedInstanceSettings;
            AzureSqlVmSettings = azureSqlVmSettings;
            ConfidenceRatingInPercentage = confidenceRatingInPercentage;
            CreatedTimestamp = createdTimestamp;
            Currency = currency;
            DisasterRecoveryLocation = disasterRecoveryLocation;
            DiscountPercentage = discountPercentage;
            EaSubscriptionId = eaSubscriptionId;
            EnableHadrAssessment = enableHadrAssessment;
            EntityUptime = entityUptime;
            EnvironmentType = environmentType;
            GroupType = groupType;
            Id = id;
            IsInternetAccessAvailable = isInternetAccessAvailable;
            MultiSubnetIntent = multiSubnetIntent;
            Name = name;
            OptimizationLogic = optimizationLogic;
            OsLicense = osLicense;
            Percentile = percentile;
            PerfDataEndTime = perfDataEndTime;
            PerfDataStartTime = perfDataStartTime;
            PricesTimestamp = pricesTimestamp;
            ProvisioningState = provisioningState;
            ReservedInstance = reservedInstance;
            ReservedInstanceForVm = reservedInstanceForVm;
            ScalingFactor = scalingFactor;
            SchemaVersion = schemaVersion;
            SizingCriterion = sizingCriterion;
            SqlServerLicense = sqlServerLicense;
            Stage = stage;
            Status = status;
            SystemData = systemData;
            TimeRange = timeRange;
            Type = type;
            UpdatedTimestamp = updatedTimestamp;
        }
    }
}
