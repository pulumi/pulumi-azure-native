// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Automation
{
    public static class GetWatcher
    {
        /// <summary>
        /// Retrieve the watcher identified by watcher name.
        /// 
        /// Uses Azure REST API version 2023-05-15-preview.
        /// 
        /// Other available API versions: 2015-10-31, 2019-06-01, 2020-01-13-preview, 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetWatcherResult> InvokeAsync(GetWatcherArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWatcherResult>("azure-native:automation:getWatcher", args ?? new GetWatcherArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve the watcher identified by watcher name.
        /// 
        /// Uses Azure REST API version 2023-05-15-preview.
        /// 
        /// Other available API versions: 2015-10-31, 2019-06-01, 2020-01-13-preview, 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWatcherResult> Invoke(GetWatcherInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWatcherResult>("azure-native:automation:getWatcher", args ?? new GetWatcherInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve the watcher identified by watcher name.
        /// 
        /// Uses Azure REST API version 2023-05-15-preview.
        /// 
        /// Other available API versions: 2015-10-31, 2019-06-01, 2020-01-13-preview, 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWatcherResult> Invoke(GetWatcherInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWatcherResult>("azure-native:automation:getWatcher", args ?? new GetWatcherInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWatcherArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public string AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The watcher name.
        /// </summary>
        [Input("watcherName", required: true)]
        public string WatcherName { get; set; } = null!;

        public GetWatcherArgs()
        {
        }
        public static new GetWatcherArgs Empty => new GetWatcherArgs();
    }

    public sealed class GetWatcherInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The watcher name.
        /// </summary>
        [Input("watcherName", required: true)]
        public Input<string> WatcherName { get; set; } = null!;

        public GetWatcherInvokeArgs()
        {
        }
        public static new GetWatcherInvokeArgs Empty => new GetWatcherInvokeArgs();
    }


    [OutputType]
    public sealed class GetWatcherResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Gets or sets the creation time.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// Gets or sets the description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Gets or sets the etag of the resource.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Gets or sets the frequency at which the watcher is invoked.
        /// </summary>
        public readonly double? ExecutionFrequencyInSeconds;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Details of the user who last modified the watcher.
        /// </summary>
        public readonly string LastModifiedBy;
        /// <summary>
        /// Gets or sets the last modified time.
        /// </summary>
        public readonly string LastModifiedTime;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets the name of the script the watcher is attached to, i.e. the name of an existing runbook.
        /// </summary>
        public readonly string? ScriptName;
        /// <summary>
        /// Gets or sets the parameters of the script.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ScriptParameters;
        /// <summary>
        /// Gets or sets the name of the hybrid worker group the watcher will run on.
        /// </summary>
        public readonly string? ScriptRunOn;
        /// <summary>
        /// Gets the current status of the watcher.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetWatcherResult(
            string azureApiVersion,

            string creationTime,

            string? description,

            string? etag,

            double? executionFrequencyInSeconds,

            string id,

            string lastModifiedBy,

            string lastModifiedTime,

            string? location,

            string name,

            string? scriptName,

            ImmutableDictionary<string, string>? scriptParameters,

            string? scriptRunOn,

            string status,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            CreationTime = creationTime;
            Description = description;
            Etag = etag;
            ExecutionFrequencyInSeconds = executionFrequencyInSeconds;
            Id = id;
            LastModifiedBy = lastModifiedBy;
            LastModifiedTime = lastModifiedTime;
            Location = location;
            Name = name;
            ScriptName = scriptName;
            ScriptParameters = scriptParameters;
            ScriptRunOn = scriptRunOn;
            Status = status;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
