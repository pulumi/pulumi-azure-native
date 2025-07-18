// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Communication
{
    public static class GetCommunicationService
    {
        /// <summary>
        /// Get the CommunicationService and its properties.
        /// 
        /// Uses Azure REST API version 2023-06-01-preview.
        /// 
        /// Other available API versions: 2023-03-31, 2023-04-01, 2023-04-01-preview, 2024-09-01-preview, 2025-05-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native communication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetCommunicationServiceResult> InvokeAsync(GetCommunicationServiceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCommunicationServiceResult>("azure-native:communication:getCommunicationService", args ?? new GetCommunicationServiceArgs(), options.WithDefaults());

        /// <summary>
        /// Get the CommunicationService and its properties.
        /// 
        /// Uses Azure REST API version 2023-06-01-preview.
        /// 
        /// Other available API versions: 2023-03-31, 2023-04-01, 2023-04-01-preview, 2024-09-01-preview, 2025-05-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native communication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetCommunicationServiceResult> Invoke(GetCommunicationServiceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCommunicationServiceResult>("azure-native:communication:getCommunicationService", args ?? new GetCommunicationServiceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the CommunicationService and its properties.
        /// 
        /// Uses Azure REST API version 2023-06-01-preview.
        /// 
        /// Other available API versions: 2023-03-31, 2023-04-01, 2023-04-01-preview, 2024-09-01-preview, 2025-05-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native communication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetCommunicationServiceResult> Invoke(GetCommunicationServiceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCommunicationServiceResult>("azure-native:communication:getCommunicationService", args ?? new GetCommunicationServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCommunicationServiceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the CommunicationService resource.
        /// </summary>
        [Input("communicationServiceName", required: true)]
        public string CommunicationServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetCommunicationServiceArgs()
        {
        }
        public static new GetCommunicationServiceArgs Empty => new GetCommunicationServiceArgs();
    }

    public sealed class GetCommunicationServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the CommunicationService resource.
        /// </summary>
        [Input("communicationServiceName", required: true)]
        public Input<string> CommunicationServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetCommunicationServiceInvokeArgs()
        {
        }
        public static new GetCommunicationServiceInvokeArgs Empty => new GetCommunicationServiceInvokeArgs();
    }


    [OutputType]
    public sealed class GetCommunicationServiceResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The location where the communication service stores its data at rest.
        /// </summary>
        public readonly string DataLocation;
        /// <summary>
        /// FQDN of the CommunicationService instance.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Managed service identity (system assigned and/or user assigned identities)
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// The immutable resource Id of the communication service.
        /// </summary>
        public readonly string ImmutableResourceId;
        /// <summary>
        /// List of email Domain resource Ids.
        /// </summary>
        public readonly ImmutableArray<string> LinkedDomains;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource ID of an Azure Notification Hub linked to this resource.
        /// </summary>
        public readonly string NotificationHubId;
        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
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
        /// <summary>
        /// Version of the CommunicationService resource. Probably you need the same or higher version of client SDKs.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetCommunicationServiceResult(
            string azureApiVersion,

            string dataLocation,

            string hostName,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            string immutableResourceId,

            ImmutableArray<string> linkedDomains,

            string location,

            string name,

            string notificationHubId,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            string version)
        {
            AzureApiVersion = azureApiVersion;
            DataLocation = dataLocation;
            HostName = hostName;
            Id = id;
            Identity = identity;
            ImmutableResourceId = immutableResourceId;
            LinkedDomains = linkedDomains;
            Location = location;
            Name = name;
            NotificationHubId = notificationHubId;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            Version = version;
        }
    }
}
