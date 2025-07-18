// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate
{
    public static class GetAvsAssessmentsOperation
    {
        /// <summary>
        /// Get a AvsAssessment
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-03-15, 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-01-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetAvsAssessmentsOperationResult> InvokeAsync(GetAvsAssessmentsOperationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAvsAssessmentsOperationResult>("azure-native:migrate:getAvsAssessmentsOperation", args ?? new GetAvsAssessmentsOperationArgs(), options.WithDefaults());

        /// <summary>
        /// Get a AvsAssessment
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-03-15, 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-01-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAvsAssessmentsOperationResult> Invoke(GetAvsAssessmentsOperationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAvsAssessmentsOperationResult>("azure-native:migrate:getAvsAssessmentsOperation", args ?? new GetAvsAssessmentsOperationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a AvsAssessment
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-03-15, 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-01-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAvsAssessmentsOperationResult> Invoke(GetAvsAssessmentsOperationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAvsAssessmentsOperationResult>("azure-native:migrate:getAvsAssessmentsOperation", args ?? new GetAvsAssessmentsOperationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAvsAssessmentsOperationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// AVS Assessment ARM name
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

        public GetAvsAssessmentsOperationArgs()
        {
        }
        public static new GetAvsAssessmentsOperationArgs Empty => new GetAvsAssessmentsOperationArgs();
    }

    public sealed class GetAvsAssessmentsOperationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// AVS Assessment ARM name
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

        public GetAvsAssessmentsOperationInvokeArgs()
        {
        }
        public static new GetAvsAssessmentsOperationInvokeArgs Empty => new GetAvsAssessmentsOperationInvokeArgs();
    }


    [OutputType]
    public sealed class GetAvsAssessmentsOperationResult
    {
        /// <summary>
        /// Gets the assessment error summary.
        ///             This is the number of machines
        /// affected by each type of error in this assessment.
        /// </summary>
        public readonly ImmutableDictionary<string, int> AssessmentErrorSummary;
        /// <summary>
        /// Assessment type of the assessment.
        /// </summary>
        public readonly string AssessmentType;
        /// <summary>
        /// AVS Assessment Scenario.
        /// </summary>
        public readonly string? AvsAssessmentScenario;
        /// <summary>
        /// Estimated External Storage for Assessment.
        /// </summary>
        public readonly ImmutableArray<Outputs.AvsEstimatedExternalStorageResponse> AvsEstimatedExternalStorages;
        /// <summary>
        /// Estimated External Storage for Assessment.
        /// </summary>
        public readonly ImmutableArray<Outputs.AvsEstimatedNetworkResponse> AvsEstimatedNetworks;
        /// <summary>
        /// Estimated AVS SKU for Assessment.
        /// </summary>
        public readonly ImmutableArray<Outputs.AvsEstimatedNodeResponse> AvsEstimatedNodes;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Azure Location or Azure region where to which the machines will be migrated.
        /// </summary>
        public readonly string? AzureLocation;
        /// <summary>
        /// Azure Offer code according to which cost estimation is done.
        /// </summary>
        public readonly string? AzureOfferCode;
        /// <summary>
        /// Confidence Rating in Percentage.
        /// </summary>
        public readonly double ConfidenceRatingInPercentage;
        /// <summary>
        /// collection of cost components.
        /// </summary>
        public readonly ImmutableArray<Outputs.CostComponentResponse> CostComponents;
        /// <summary>
        /// Percentage of CPU capacity reserved for processing additional workloads.
        /// </summary>
        public readonly double? CpuHeadroom;
        /// <summary>
        /// Predicted CPU utilization.
        /// </summary>
        public readonly double CpuUtilization;
        /// <summary>
        /// Date and Time when assessment was created.
        /// </summary>
        public readonly string CreatedTimestamp;
        /// <summary>
        /// Currency in which prices should be reported.
        /// </summary>
        public readonly string? Currency;
        /// <summary>
        /// De-duplication compression.
        /// </summary>
        public readonly double? DedupeCompression;
        /// <summary>
        /// Custom discount percentage.
        /// </summary>
        public readonly double? DiscountPercentage;
        /// <summary>
        /// List of AVS external storage types.
        /// </summary>
        public readonly ImmutableArray<string> ExternalStorageTypes;
        /// <summary>
        /// Failures to tolerate and RAID level in a common property.
        /// </summary>
        public readonly string? FailuresToTolerateAndRaidLevel;
        /// <summary>
        /// List of Failures to tolerate and RAID levels in a common property.
        /// </summary>
        public readonly ImmutableArray<string> FailuresToTolerateAndRaidLevelList;
        /// <summary>
        /// Gets the group type for the assessment.
        /// </summary>
        public readonly string GroupType;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Is Stretch Cluster Enabled.
        /// </summary>
        public readonly bool? IsStretchClusterEnabled;
        /// <summary>
        /// Is VCF license applied
        /// </summary>
        public readonly bool? IsVcfByolEnabled;
        /// <summary>
        /// Limiting factor.
        /// </summary>
        public readonly string LimitingFactor;
        /// <summary>
        /// Memory overcommit.
        /// </summary>
        public readonly double? MemOvercommit;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// AVS node type.
        /// </summary>
        public readonly string? NodeType;
        /// <summary>
        /// AVS node types.
        /// </summary>
        public readonly ImmutableArray<string> NodeTypes;
        /// <summary>
        /// Number of machines part of the assessment.
        /// </summary>
        public readonly int NumberOfMachines;
        /// <summary>
        /// Recommended number of nodes.
        /// </summary>
        public readonly int NumberOfNodes;
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
        /// Time when the Azure Prices were queried. Date-Time represented in ISO-8601
        /// format.
        /// </summary>
        public readonly string PricesTimestamp;
        /// <summary>
        /// The status of the last operation.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Predicted RAM utilization.
        /// </summary>
        public readonly double RamUtilization;
        /// <summary>
        /// Reserved instance.
        /// </summary>
        public readonly string? ReservedInstance;
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
        /// User configurable setting to display the Stage of Assessment.
        /// </summary>
        public readonly string Stage;
        /// <summary>
        /// Whether assessment is in valid state and all machines have been assessed.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Predicted storage utilization.
        /// </summary>
        public readonly double StorageUtilization;
        /// <summary>
        /// Gets or sets the Assessment cloud suitability.
        /// </summary>
        public readonly string Suitability;
        /// <summary>
        /// Gets or sets the Assessment suitability explanation.
        /// </summary>
        public readonly string SuitabilityExplanation;
        /// <summary>
        /// Cloud suitability summary for all the machines in the assessment.
        /// </summary>
        public readonly ImmutableDictionary<string, int> SuitabilitySummary;
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
        /// Predicted total CPU cores used.
        /// </summary>
        public readonly double TotalCpuCores;
        /// <summary>
        /// Total monthly cost.
        /// </summary>
        public readonly double TotalMonthlyCost;
        /// <summary>
        /// Predicted total RAM used in GB.
        /// </summary>
        public readonly double TotalRamInGB;
        /// <summary>
        /// Predicted total Storage used in GB.
        /// </summary>
        public readonly double TotalStorageInGB;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Date and Time when assessment was last updated.
        /// </summary>
        public readonly string UpdatedTimestamp;
        /// <summary>
        /// VCPU over subscription.
        /// </summary>
        public readonly double? VcpuOversubscription;

        [OutputConstructor]
        private GetAvsAssessmentsOperationResult(
            ImmutableDictionary<string, int> assessmentErrorSummary,

            string assessmentType,

            string? avsAssessmentScenario,

            ImmutableArray<Outputs.AvsEstimatedExternalStorageResponse> avsEstimatedExternalStorages,

            ImmutableArray<Outputs.AvsEstimatedNetworkResponse> avsEstimatedNetworks,

            ImmutableArray<Outputs.AvsEstimatedNodeResponse> avsEstimatedNodes,

            string azureApiVersion,

            string? azureLocation,

            string? azureOfferCode,

            double confidenceRatingInPercentage,

            ImmutableArray<Outputs.CostComponentResponse> costComponents,

            double? cpuHeadroom,

            double cpuUtilization,

            string createdTimestamp,

            string? currency,

            double? dedupeCompression,

            double? discountPercentage,

            ImmutableArray<string> externalStorageTypes,

            string? failuresToTolerateAndRaidLevel,

            ImmutableArray<string> failuresToTolerateAndRaidLevelList,

            string groupType,

            string id,

            bool? isStretchClusterEnabled,

            bool? isVcfByolEnabled,

            string limitingFactor,

            double? memOvercommit,

            string name,

            string? nodeType,

            ImmutableArray<string> nodeTypes,

            int numberOfMachines,

            int numberOfNodes,

            string? percentile,

            string? perfDataEndTime,

            string? perfDataStartTime,

            string pricesTimestamp,

            string? provisioningState,

            double ramUtilization,

            string? reservedInstance,

            double? scalingFactor,

            string schemaVersion,

            string? sizingCriterion,

            string stage,

            string status,

            double storageUtilization,

            string suitability,

            string suitabilityExplanation,

            ImmutableDictionary<string, int> suitabilitySummary,

            Outputs.SystemDataResponse systemData,

            string? timeRange,

            double totalCpuCores,

            double totalMonthlyCost,

            double totalRamInGB,

            double totalStorageInGB,

            string type,

            string updatedTimestamp,

            double? vcpuOversubscription)
        {
            AssessmentErrorSummary = assessmentErrorSummary;
            AssessmentType = assessmentType;
            AvsAssessmentScenario = avsAssessmentScenario;
            AvsEstimatedExternalStorages = avsEstimatedExternalStorages;
            AvsEstimatedNetworks = avsEstimatedNetworks;
            AvsEstimatedNodes = avsEstimatedNodes;
            AzureApiVersion = azureApiVersion;
            AzureLocation = azureLocation;
            AzureOfferCode = azureOfferCode;
            ConfidenceRatingInPercentage = confidenceRatingInPercentage;
            CostComponents = costComponents;
            CpuHeadroom = cpuHeadroom;
            CpuUtilization = cpuUtilization;
            CreatedTimestamp = createdTimestamp;
            Currency = currency;
            DedupeCompression = dedupeCompression;
            DiscountPercentage = discountPercentage;
            ExternalStorageTypes = externalStorageTypes;
            FailuresToTolerateAndRaidLevel = failuresToTolerateAndRaidLevel;
            FailuresToTolerateAndRaidLevelList = failuresToTolerateAndRaidLevelList;
            GroupType = groupType;
            Id = id;
            IsStretchClusterEnabled = isStretchClusterEnabled;
            IsVcfByolEnabled = isVcfByolEnabled;
            LimitingFactor = limitingFactor;
            MemOvercommit = memOvercommit;
            Name = name;
            NodeType = nodeType;
            NodeTypes = nodeTypes;
            NumberOfMachines = numberOfMachines;
            NumberOfNodes = numberOfNodes;
            Percentile = percentile;
            PerfDataEndTime = perfDataEndTime;
            PerfDataStartTime = perfDataStartTime;
            PricesTimestamp = pricesTimestamp;
            ProvisioningState = provisioningState;
            RamUtilization = ramUtilization;
            ReservedInstance = reservedInstance;
            ScalingFactor = scalingFactor;
            SchemaVersion = schemaVersion;
            SizingCriterion = sizingCriterion;
            Stage = stage;
            Status = status;
            StorageUtilization = storageUtilization;
            Suitability = suitability;
            SuitabilityExplanation = suitabilityExplanation;
            SuitabilitySummary = suitabilitySummary;
            SystemData = systemData;
            TimeRange = timeRange;
            TotalCpuCores = totalCpuCores;
            TotalMonthlyCost = totalMonthlyCost;
            TotalRamInGB = totalRamInGB;
            TotalStorageInGB = totalStorageInGB;
            Type = type;
            UpdatedTimestamp = updatedTimestamp;
            VcpuOversubscription = vcpuOversubscription;
        }
    }
}
