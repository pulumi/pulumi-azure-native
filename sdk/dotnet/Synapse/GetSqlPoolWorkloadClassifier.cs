// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse
{
    public static class GetSqlPoolWorkloadClassifier
    {
        /// <summary>
        /// Get a workload classifier of Sql pool's workload group.
        /// 
        /// Uses Azure REST API version 2021-06-01.
        /// 
        /// Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetSqlPoolWorkloadClassifierResult> InvokeAsync(GetSqlPoolWorkloadClassifierArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSqlPoolWorkloadClassifierResult>("azure-native:synapse:getSqlPoolWorkloadClassifier", args ?? new GetSqlPoolWorkloadClassifierArgs(), options.WithDefaults());

        /// <summary>
        /// Get a workload classifier of Sql pool's workload group.
        /// 
        /// Uses Azure REST API version 2021-06-01.
        /// 
        /// Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSqlPoolWorkloadClassifierResult> Invoke(GetSqlPoolWorkloadClassifierInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSqlPoolWorkloadClassifierResult>("azure-native:synapse:getSqlPoolWorkloadClassifier", args ?? new GetSqlPoolWorkloadClassifierInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a workload classifier of Sql pool's workload group.
        /// 
        /// Uses Azure REST API version 2021-06-01.
        /// 
        /// Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSqlPoolWorkloadClassifierResult> Invoke(GetSqlPoolWorkloadClassifierInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSqlPoolWorkloadClassifierResult>("azure-native:synapse:getSqlPoolWorkloadClassifier", args ?? new GetSqlPoolWorkloadClassifierInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSqlPoolWorkloadClassifierArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// SQL pool name
        /// </summary>
        [Input("sqlPoolName", required: true)]
        public string SqlPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the workload classifier.
        /// </summary>
        [Input("workloadClassifierName", required: true)]
        public string WorkloadClassifierName { get; set; } = null!;

        /// <summary>
        /// The name of the workload group.
        /// </summary>
        [Input("workloadGroupName", required: true)]
        public string WorkloadGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetSqlPoolWorkloadClassifierArgs()
        {
        }
        public static new GetSqlPoolWorkloadClassifierArgs Empty => new GetSqlPoolWorkloadClassifierArgs();
    }

    public sealed class GetSqlPoolWorkloadClassifierInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// SQL pool name
        /// </summary>
        [Input("sqlPoolName", required: true)]
        public Input<string> SqlPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the workload classifier.
        /// </summary>
        [Input("workloadClassifierName", required: true)]
        public Input<string> WorkloadClassifierName { get; set; } = null!;

        /// <summary>
        /// The name of the workload group.
        /// </summary>
        [Input("workloadGroupName", required: true)]
        public Input<string> WorkloadGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetSqlPoolWorkloadClassifierInvokeArgs()
        {
        }
        public static new GetSqlPoolWorkloadClassifierInvokeArgs Empty => new GetSqlPoolWorkloadClassifierInvokeArgs();
    }


    [OutputType]
    public sealed class GetSqlPoolWorkloadClassifierResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The workload classifier context.
        /// </summary>
        public readonly string? Context;
        /// <summary>
        /// The workload classifier end time for classification.
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The workload classifier importance.
        /// </summary>
        public readonly string? Importance;
        /// <summary>
        /// The workload classifier label.
        /// </summary>
        public readonly string? Label;
        /// <summary>
        /// The workload classifier member name.
        /// </summary>
        public readonly string MemberName;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The workload classifier start time for classification.
        /// </summary>
        public readonly string? StartTime;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSqlPoolWorkloadClassifierResult(
            string azureApiVersion,

            string? context,

            string? endTime,

            string id,

            string? importance,

            string? label,

            string memberName,

            string name,

            string? startTime,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Context = context;
            EndTime = endTime;
            Id = id;
            Importance = importance;
            Label = label;
            MemberName = memberName;
            Name = name;
            StartTime = startTime;
            Type = type;
        }
    }
}
