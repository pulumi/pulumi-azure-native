// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Chaos.V20241101Preview
{
    public static class GetExperimentExecutionDetails
    {
        /// <summary>
        /// Execution details of an experiment resource.
        /// </summary>
        public static Task<GetExperimentExecutionDetailsResult> InvokeAsync(GetExperimentExecutionDetailsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExperimentExecutionDetailsResult>("azure-native:chaos/v20241101preview:getExperimentExecutionDetails", args ?? new GetExperimentExecutionDetailsArgs(), options.WithDefaults());

        /// <summary>
        /// Execution details of an experiment resource.
        /// </summary>
        public static Output<GetExperimentExecutionDetailsResult> Invoke(GetExperimentExecutionDetailsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExperimentExecutionDetailsResult>("azure-native:chaos/v20241101preview:getExperimentExecutionDetails", args ?? new GetExperimentExecutionDetailsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Execution details of an experiment resource.
        /// </summary>
        public static Output<GetExperimentExecutionDetailsResult> Invoke(GetExperimentExecutionDetailsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetExperimentExecutionDetailsResult>("azure-native:chaos/v20241101preview:getExperimentExecutionDetails", args ?? new GetExperimentExecutionDetailsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExperimentExecutionDetailsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// GUID that represents a Experiment execution detail.
        /// </summary>
        [Input("executionId", required: true)]
        public string ExecutionId { get; set; } = null!;

        /// <summary>
        /// String that represents a Experiment resource name.
        /// </summary>
        [Input("experimentName", required: true)]
        public string ExperimentName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetExperimentExecutionDetailsArgs()
        {
        }
        public static new GetExperimentExecutionDetailsArgs Empty => new GetExperimentExecutionDetailsArgs();
    }

    public sealed class GetExperimentExecutionDetailsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// GUID that represents a Experiment execution detail.
        /// </summary>
        [Input("executionId", required: true)]
        public Input<string> ExecutionId { get; set; } = null!;

        /// <summary>
        /// String that represents a Experiment resource name.
        /// </summary>
        [Input("experimentName", required: true)]
        public Input<string> ExperimentName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetExperimentExecutionDetailsInvokeArgs()
        {
        }
        public static new GetExperimentExecutionDetailsInvokeArgs Empty => new GetExperimentExecutionDetailsInvokeArgs();
    }


    [OutputType]
    public sealed class GetExperimentExecutionDetailsResult
    {
        /// <summary>
        /// String of the fully qualified resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// String of the resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties of the experiment execution details.
        /// </summary>
        public readonly Outputs.ExperimentExecutionDetailsPropertiesResponse Properties;
        /// <summary>
        /// String of the resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetExperimentExecutionDetailsResult(
            string id,

            string name,

            Outputs.ExperimentExecutionDetailsPropertiesResponse properties,

            string type)
        {
            Id = id;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
