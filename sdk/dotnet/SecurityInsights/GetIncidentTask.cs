// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights
{
    public static class GetIncidentTask
    {
        /// <summary>
        /// Gets an incident task.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-03-01, 2024-04-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-03-01, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetIncidentTaskResult> InvokeAsync(GetIncidentTaskArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIncidentTaskResult>("azure-native:securityinsights:getIncidentTask", args ?? new GetIncidentTaskArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an incident task.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-03-01, 2024-04-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-03-01, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetIncidentTaskResult> Invoke(GetIncidentTaskInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIncidentTaskResult>("azure-native:securityinsights:getIncidentTask", args ?? new GetIncidentTaskInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an incident task.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-03-01, 2024-04-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-03-01, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetIncidentTaskResult> Invoke(GetIncidentTaskInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIncidentTaskResult>("azure-native:securityinsights:getIncidentTask", args ?? new GetIncidentTaskInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIncidentTaskArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Incident ID
        /// </summary>
        [Input("incidentId", required: true)]
        public string IncidentId { get; set; } = null!;

        /// <summary>
        /// Incident task ID
        /// </summary>
        [Input("incidentTaskId", required: true)]
        public string IncidentTaskId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetIncidentTaskArgs()
        {
        }
        public static new GetIncidentTaskArgs Empty => new GetIncidentTaskArgs();
    }

    public sealed class GetIncidentTaskInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Incident ID
        /// </summary>
        [Input("incidentId", required: true)]
        public Input<string> IncidentId { get; set; } = null!;

        /// <summary>
        /// Incident task ID
        /// </summary>
        [Input("incidentTaskId", required: true)]
        public Input<string> IncidentTaskId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetIncidentTaskInvokeArgs()
        {
        }
        public static new GetIncidentTaskInvokeArgs Empty => new GetIncidentTaskInvokeArgs();
    }


    [OutputType]
    public sealed class GetIncidentTaskResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Information on the client (user or application) that made some action
        /// </summary>
        public readonly Outputs.ClientInfoResponse? CreatedBy;
        /// <summary>
        /// The time the task was created
        /// </summary>
        public readonly string CreatedTimeUtc;
        /// <summary>
        /// The description of the task
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Information on the client (user or application) that made some action
        /// </summary>
        public readonly Outputs.ClientInfoResponse? LastModifiedBy;
        /// <summary>
        /// The last time the task was updated
        /// </summary>
        public readonly string LastModifiedTimeUtc;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The status of the task
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The title of the task
        /// </summary>
        public readonly string Title;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetIncidentTaskResult(
            string azureApiVersion,

            Outputs.ClientInfoResponse? createdBy,

            string createdTimeUtc,

            string? description,

            string? etag,

            string id,

            Outputs.ClientInfoResponse? lastModifiedBy,

            string lastModifiedTimeUtc,

            string name,

            string status,

            Outputs.SystemDataResponse systemData,

            string title,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            CreatedBy = createdBy;
            CreatedTimeUtc = createdTimeUtc;
            Description = description;
            Etag = etag;
            Id = id;
            LastModifiedBy = lastModifiedBy;
            LastModifiedTimeUtc = lastModifiedTimeUtc;
            Name = name;
            Status = status;
            SystemData = systemData;
            Title = title;
            Type = type;
        }
    }
}
