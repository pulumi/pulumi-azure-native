// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub
{
    public static class GetProviderRegistration
    {
        /// <summary>
        /// Gets the provider registration details.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2021-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native providerhub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetProviderRegistrationResult> InvokeAsync(GetProviderRegistrationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProviderRegistrationResult>("azure-native:providerhub:getProviderRegistration", args ?? new GetProviderRegistrationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the provider registration details.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2021-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native providerhub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetProviderRegistrationResult> Invoke(GetProviderRegistrationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProviderRegistrationResult>("azure-native:providerhub:getProviderRegistration", args ?? new GetProviderRegistrationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the provider registration details.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2021-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native providerhub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetProviderRegistrationResult> Invoke(GetProviderRegistrationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProviderRegistrationResult>("azure-native:providerhub:getProviderRegistration", args ?? new GetProviderRegistrationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProviderRegistrationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource provider hosted within ProviderHub.
        /// </summary>
        [Input("providerNamespace", required: true)]
        public string ProviderNamespace { get; set; } = null!;

        public GetProviderRegistrationArgs()
        {
        }
        public static new GetProviderRegistrationArgs Empty => new GetProviderRegistrationArgs();
    }

    public sealed class GetProviderRegistrationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource provider hosted within ProviderHub.
        /// </summary>
        [Input("providerNamespace", required: true)]
        public Input<string> ProviderNamespace { get; set; } = null!;

        public GetProviderRegistrationInvokeArgs()
        {
        }
        public static new GetProviderRegistrationInvokeArgs Empty => new GetProviderRegistrationInvokeArgs();
    }


    [OutputType]
    public sealed class GetProviderRegistrationResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Provider registration kind. This Metadata is also used by portal/tooling/etc to render different UX experiences for resources of the same type.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        public readonly Outputs.ProviderRegistrationPropertiesResponse Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetProviderRegistrationResult(
            string azureApiVersion,

            string id,

            string? kind,

            string name,

            Outputs.ProviderRegistrationPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Kind = kind;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Type = type;
        }
    }
}
