// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Portal
{
    public static class GetTenantConfiguration
    {
        /// <summary>
        /// Gets the tenant configuration.
        /// 
        /// Uses Azure REST API version 2022-12-01-preview.
        /// 
        /// Other available API versions: 2019-01-01-preview, 2020-09-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native portal [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetTenantConfigurationResult> InvokeAsync(GetTenantConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTenantConfigurationResult>("azure-native:portal:getTenantConfiguration", args ?? new GetTenantConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the tenant configuration.
        /// 
        /// Uses Azure REST API version 2022-12-01-preview.
        /// 
        /// Other available API versions: 2019-01-01-preview, 2020-09-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native portal [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTenantConfigurationResult> Invoke(GetTenantConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTenantConfigurationResult>("azure-native:portal:getTenantConfiguration", args ?? new GetTenantConfigurationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the tenant configuration.
        /// 
        /// Uses Azure REST API version 2022-12-01-preview.
        /// 
        /// Other available API versions: 2019-01-01-preview, 2020-09-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native portal [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTenantConfigurationResult> Invoke(GetTenantConfigurationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTenantConfigurationResult>("azure-native:portal:getTenantConfiguration", args ?? new GetTenantConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTenantConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Configuration
        /// </summary>
        [Input("configurationName", required: true)]
        public string ConfigurationName { get; set; } = null!;

        public GetTenantConfigurationArgs()
        {
        }
        public static new GetTenantConfigurationArgs Empty => new GetTenantConfigurationArgs();
    }

    public sealed class GetTenantConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Configuration
        /// </summary>
        [Input("configurationName", required: true)]
        public Input<string> ConfigurationName { get; set; } = null!;

        public GetTenantConfigurationInvokeArgs()
        {
        }
        public static new GetTenantConfigurationInvokeArgs Empty => new GetTenantConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetTenantConfigurationResult
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
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource-specific properties for this resource.
        /// </summary>
        public readonly Outputs.ConfigurationPropertiesResponse Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTenantConfigurationResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.ConfigurationPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Type = type;
        }
    }
}
