// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices
{
    public static class GetRegistryEnvironmentVersion
    {
        /// <summary>
        /// Azure Resource Manager resource envelope.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetRegistryEnvironmentVersionResult> InvokeAsync(GetRegistryEnvironmentVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegistryEnvironmentVersionResult>("azure-native:machinelearningservices:getRegistryEnvironmentVersion", args ?? new GetRegistryEnvironmentVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Azure Resource Manager resource envelope.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetRegistryEnvironmentVersionResult> Invoke(GetRegistryEnvironmentVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegistryEnvironmentVersionResult>("azure-native:machinelearningservices:getRegistryEnvironmentVersion", args ?? new GetRegistryEnvironmentVersionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Azure Resource Manager resource envelope.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetRegistryEnvironmentVersionResult> Invoke(GetRegistryEnvironmentVersionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegistryEnvironmentVersionResult>("azure-native:machinelearningservices:getRegistryEnvironmentVersion", args ?? new GetRegistryEnvironmentVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegistryEnvironmentVersionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Container name. This is case-sensitive.
        /// </summary>
        [Input("environmentName", required: true)]
        public string EnvironmentName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning registry. This is case-insensitive
        /// </summary>
        [Input("registryName", required: true)]
        public string RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Version identifier. This is case-sensitive.
        /// </summary>
        [Input("version", required: true)]
        public string Version { get; set; } = null!;

        public GetRegistryEnvironmentVersionArgs()
        {
        }
        public static new GetRegistryEnvironmentVersionArgs Empty => new GetRegistryEnvironmentVersionArgs();
    }

    public sealed class GetRegistryEnvironmentVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Container name. This is case-sensitive.
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning registry. This is case-insensitive
        /// </summary>
        [Input("registryName", required: true)]
        public Input<string> RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Version identifier. This is case-sensitive.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public GetRegistryEnvironmentVersionInvokeArgs()
        {
        }
        public static new GetRegistryEnvironmentVersionInvokeArgs Empty => new GetRegistryEnvironmentVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegistryEnvironmentVersionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// [Required] Additional attributes of the entity.
        /// </summary>
        public readonly Outputs.EnvironmentVersionResponse EnvironmentVersionProperties;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRegistryEnvironmentVersionResult(
            string azureApiVersion,

            Outputs.EnvironmentVersionResponse environmentVersionProperties,

            string id,

            string name,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            EnvironmentVersionProperties = environmentVersionProperties;
            Id = id;
            Name = name;
            SystemData = systemData;
            Type = type;
        }
    }
}
