// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI
{
    /// <summary>
    /// Details of an Update run
    /// 
    /// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.
    /// 
    /// Other available API versions: 2022-12-15-preview, 2023-02-01, 2023-03-01, 2023-06-01, 2023-08-01, 2023-08-01-preview, 2023-11-01-preview, 2024-01-01, 2024-02-15-preview, 2024-09-01-preview, 2024-12-01-preview, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestackhci [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:azurestackhci:UpdateRun")]
    public partial class UpdateRun : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// More detailed description of the step.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Duration of the update run.
        /// </summary>
        [Output("duration")]
        public Output<string?> Duration { get; private set; } = null!;

        /// <summary>
        /// When the step reached a terminal state.
        /// </summary>
        [Output("endTimeUtc")]
        public Output<string?> EndTimeUtc { get; private set; } = null!;

        /// <summary>
        /// Error message, specified if the step is in a failed state.
        /// </summary>
        [Output("errorMessage")]
        public Output<string?> ErrorMessage { get; private set; } = null!;

        /// <summary>
        /// Expected execution time of a given step. This is optionally authored in the update action plan and can be empty.
        /// </summary>
        [Output("expectedExecutionTime")]
        public Output<string?> ExpectedExecutionTime { get; private set; } = null!;

        /// <summary>
        /// Timestamp of the most recently completed step in the update run.
        /// </summary>
        [Output("lastUpdatedTime")]
        public Output<string?> LastUpdatedTime { get; private set; } = null!;

        /// <summary>
        /// Completion time of this step or the last completed sub-step.
        /// </summary>
        [Output("lastUpdatedTimeUtc")]
        public Output<string?> LastUpdatedTimeUtc { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the UpdateRuns proxy resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// When the step started, or empty if it has not started executing.
        /// </summary>
        [Output("startTimeUtc")]
        public Output<string?> StartTimeUtc { get; private set; } = null!;

        /// <summary>
        /// State of the update run.
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// Status of the step, bubbled up from the ECE action plan for installation attempts. Values are: 'Success', 'Error', 'InProgress', and 'Unknown status'.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Recursive model for child steps of this step.
        /// </summary>
        [Output("steps")]
        public Output<ImmutableArray<Outputs.StepResponse>> Steps { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Timestamp of the update run was started.
        /// </summary>
        [Output("timeStarted")]
        public Output<string?> TimeStarted { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a UpdateRun resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UpdateRun(string name, UpdateRunArgs args, CustomResourceOptions? options = null)
            : base("azure-native:azurestackhci:UpdateRun", name, args ?? new UpdateRunArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UpdateRun(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:azurestackhci:UpdateRun", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20221201:UpdateRun" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20221215preview:UpdateRun" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230201:UpdateRun" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230301:UpdateRun" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230601:UpdateRun" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230801:UpdateRun" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230801preview:UpdateRun" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20231101preview:UpdateRun" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240101:UpdateRun" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240215preview:UpdateRun" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240401:UpdateRun" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240901preview:UpdateRun" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20241201preview:UpdateRun" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20250201preview:UpdateRun" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing UpdateRun resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UpdateRun Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new UpdateRun(name, id, options);
        }
    }

    public sealed class UpdateRunArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// More detailed description of the step.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Duration of the update run.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// When the step reached a terminal state.
        /// </summary>
        [Input("endTimeUtc")]
        public Input<string>? EndTimeUtc { get; set; }

        /// <summary>
        /// Error message, specified if the step is in a failed state.
        /// </summary>
        [Input("errorMessage")]
        public Input<string>? ErrorMessage { get; set; }

        /// <summary>
        /// Expected execution time of a given step. This is optionally authored in the update action plan and can be empty.
        /// </summary>
        [Input("expectedExecutionTime")]
        public Input<string>? ExpectedExecutionTime { get; set; }

        /// <summary>
        /// Timestamp of the most recently completed step in the update run.
        /// </summary>
        [Input("lastUpdatedTime")]
        public Input<string>? LastUpdatedTime { get; set; }

        /// <summary>
        /// Completion time of this step or the last completed sub-step.
        /// </summary>
        [Input("lastUpdatedTimeUtc")]
        public Input<string>? LastUpdatedTimeUtc { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the step.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// When the step started, or empty if it has not started executing.
        /// </summary>
        [Input("startTimeUtc")]
        public Input<string>? StartTimeUtc { get; set; }

        /// <summary>
        /// State of the update run.
        /// </summary>
        [Input("state")]
        public InputUnion<string, Pulumi.AzureNative.AzureStackHCI.UpdateRunPropertiesState>? State { get; set; }

        /// <summary>
        /// Status of the step, bubbled up from the ECE action plan for installation attempts. Values are: 'Success', 'Error', 'InProgress', and 'Unknown status'.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("steps")]
        private InputList<Inputs.StepArgs>? _steps;

        /// <summary>
        /// Recursive model for child steps of this step.
        /// </summary>
        public InputList<Inputs.StepArgs> Steps
        {
            get => _steps ?? (_steps = new InputList<Inputs.StepArgs>());
            set => _steps = value;
        }

        /// <summary>
        /// Timestamp of the update run was started.
        /// </summary>
        [Input("timeStarted")]
        public Input<string>? TimeStarted { get; set; }

        /// <summary>
        /// The name of the Update
        /// </summary>
        [Input("updateName", required: true)]
        public Input<string> UpdateName { get; set; } = null!;

        /// <summary>
        /// The name of the Update Run
        /// </summary>
        [Input("updateRunName")]
        public Input<string>? UpdateRunName { get; set; }

        public UpdateRunArgs()
        {
        }
        public static new UpdateRunArgs Empty => new UpdateRunArgs();
    }
}
