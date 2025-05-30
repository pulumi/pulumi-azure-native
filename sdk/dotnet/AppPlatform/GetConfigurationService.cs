// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform
{
    public static class GetConfigurationService
    {
        /// <summary>
        /// Get the Application Configuration Service and its properties.
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetConfigurationServiceResult> InvokeAsync(GetConfigurationServiceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationServiceResult>("azure-native:appplatform:getConfigurationService", args ?? new GetConfigurationServiceArgs(), options.WithDefaults());

        /// <summary>
        /// Get the Application Configuration Service and its properties.
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetConfigurationServiceResult> Invoke(GetConfigurationServiceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigurationServiceResult>("azure-native:appplatform:getConfigurationService", args ?? new GetConfigurationServiceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the Application Configuration Service and its properties.
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetConfigurationServiceResult> Invoke(GetConfigurationServiceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigurationServiceResult>("azure-native:appplatform:getConfigurationService", args ?? new GetConfigurationServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigurationServiceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of Application Configuration Service.
        /// </summary>
        [Input("configurationServiceName", required: true)]
        public string ConfigurationServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Service resource.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetConfigurationServiceArgs()
        {
        }
        public static new GetConfigurationServiceArgs Empty => new GetConfigurationServiceArgs();
    }

    public sealed class GetConfigurationServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of Application Configuration Service.
        /// </summary>
        [Input("configurationServiceName", required: true)]
        public Input<string> ConfigurationServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Service resource.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetConfigurationServiceInvokeArgs()
        {
        }
        public static new GetConfigurationServiceInvokeArgs Empty => new GetConfigurationServiceInvokeArgs();
    }


    [OutputType]
    public sealed class GetConfigurationServiceResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource Id for the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Application Configuration Service properties payload
        /// </summary>
        public readonly Outputs.ConfigurationServicePropertiesResponse Properties;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetConfigurationServiceResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.ConfigurationServicePropertiesResponse properties,

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
