// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media
{
    public static class GetMediaService
    {
        /// <summary>
        /// Get the details of a Media Services account
        /// 
        /// Uses Azure REST API version 2023-01-01.
        /// 
        /// Other available API versions: 2015-10-01, 2018-03-30-preview, 2018-06-01-preview, 2018-07-01, 2020-05-01, 2021-05-01, 2021-06-01, 2021-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native media [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetMediaServiceResult> InvokeAsync(GetMediaServiceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMediaServiceResult>("azure-native:media:getMediaService", args ?? new GetMediaServiceArgs(), options.WithDefaults());

        /// <summary>
        /// Get the details of a Media Services account
        /// 
        /// Uses Azure REST API version 2023-01-01.
        /// 
        /// Other available API versions: 2015-10-01, 2018-03-30-preview, 2018-06-01-preview, 2018-07-01, 2020-05-01, 2021-05-01, 2021-06-01, 2021-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native media [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMediaServiceResult> Invoke(GetMediaServiceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMediaServiceResult>("azure-native:media:getMediaService", args ?? new GetMediaServiceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the details of a Media Services account
        /// 
        /// Uses Azure REST API version 2023-01-01.
        /// 
        /// Other available API versions: 2015-10-01, 2018-03-30-preview, 2018-06-01-preview, 2018-07-01, 2020-05-01, 2021-05-01, 2021-06-01, 2021-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native media [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMediaServiceResult> Invoke(GetMediaServiceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMediaServiceResult>("azure-native:media:getMediaService", args ?? new GetMediaServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMediaServiceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMediaServiceArgs()
        {
        }
        public static new GetMediaServiceArgs Empty => new GetMediaServiceArgs();
    }

    public sealed class GetMediaServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetMediaServiceInvokeArgs()
        {
        }
        public static new GetMediaServiceInvokeArgs Empty => new GetMediaServiceInvokeArgs();
    }


    [OutputType]
    public sealed class GetMediaServiceResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The account encryption properties.
        /// </summary>
        public readonly Outputs.AccountEncryptionResponse? Encryption;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Managed Identity for the Media Services account.
        /// </summary>
        public readonly Outputs.MediaServiceIdentityResponse? Identity;
        /// <summary>
        /// The Key Delivery properties for Media Services account.
        /// </summary>
        public readonly Outputs.KeyDeliveryResponse? KeyDelivery;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The Media Services account ID.
        /// </summary>
        public readonly string MediaServiceId;
        /// <summary>
        /// The minimum TLS version allowed for this account's requests. This is an optional property. If unspecified, a secure default value will be used.
        /// </summary>
        public readonly string? MinimumTlsVersion;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Private Endpoint Connections created for the Media Service account.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponse> PrivateEndpointConnections;
        /// <summary>
        /// Provisioning state of the Media Services account.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Whether or not public network access is allowed for resources under the Media Services account.
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// The storage accounts for this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.StorageAccountResponse> StorageAccounts;
        public readonly string? StorageAuthentication;
        /// <summary>
        /// The system metadata relating to this resource.
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
        private GetMediaServiceResult(
            string azureApiVersion,

            Outputs.AccountEncryptionResponse? encryption,

            string id,

            Outputs.MediaServiceIdentityResponse? identity,

            Outputs.KeyDeliveryResponse? keyDelivery,

            string location,

            string mediaServiceId,

            string? minimumTlsVersion,

            string name,

            ImmutableArray<Outputs.PrivateEndpointConnectionResponse> privateEndpointConnections,

            string provisioningState,

            string? publicNetworkAccess,

            ImmutableArray<Outputs.StorageAccountResponse> storageAccounts,

            string? storageAuthentication,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Encryption = encryption;
            Id = id;
            Identity = identity;
            KeyDelivery = keyDelivery;
            Location = location;
            MediaServiceId = mediaServiceId;
            MinimumTlsVersion = minimumTlsVersion;
            Name = name;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            PublicNetworkAccess = publicNetworkAccess;
            StorageAccounts = storageAccounts;
            StorageAuthentication = storageAuthentication;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
