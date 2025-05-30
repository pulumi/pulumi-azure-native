// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App
{
    public static class GetDotNetComponent
    {
        /// <summary>
        /// .NET Component.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetDotNetComponentResult> InvokeAsync(GetDotNetComponentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDotNetComponentResult>("azure-native:app:getDotNetComponent", args ?? new GetDotNetComponentArgs(), options.WithDefaults());

        /// <summary>
        /// .NET Component.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDotNetComponentResult> Invoke(GetDotNetComponentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDotNetComponentResult>("azure-native:app:getDotNetComponent", args ?? new GetDotNetComponentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// .NET Component.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDotNetComponentResult> Invoke(GetDotNetComponentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDotNetComponentResult>("azure-native:app:getDotNetComponent", args ?? new GetDotNetComponentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDotNetComponentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Managed Environment.
        /// </summary>
        [Input("environmentName", required: true)]
        public string EnvironmentName { get; set; } = null!;

        /// <summary>
        /// Name of the .NET Component.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDotNetComponentArgs()
        {
        }
        public static new GetDotNetComponentArgs Empty => new GetDotNetComponentArgs();
    }

    public sealed class GetDotNetComponentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Managed Environment.
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        /// <summary>
        /// Name of the .NET Component.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDotNetComponentInvokeArgs()
        {
        }
        public static new GetDotNetComponentInvokeArgs Empty => new GetDotNetComponentInvokeArgs();
    }


    [OutputType]
    public sealed class GetDotNetComponentResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Type of the .NET Component.
        /// </summary>
        public readonly string? ComponentType;
        /// <summary>
        /// List of .NET Components configuration properties
        /// </summary>
        public readonly ImmutableArray<Outputs.DotNetComponentConfigurationPropertyResponse> Configurations;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state of the .NET Component.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// List of .NET Components that are bound to the .NET component
        /// </summary>
        public readonly ImmutableArray<Outputs.DotNetComponentServiceBindResponse> ServiceBinds;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDotNetComponentResult(
            string azureApiVersion,

            string? componentType,

            ImmutableArray<Outputs.DotNetComponentConfigurationPropertyResponse> configurations,

            string id,

            string name,

            string provisioningState,

            ImmutableArray<Outputs.DotNetComponentServiceBindResponse> serviceBinds,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            ComponentType = componentType;
            Configurations = configurations;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            ServiceBinds = serviceBinds;
            SystemData = systemData;
            Type = type;
        }
    }
}
