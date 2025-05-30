// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices
{
    public static class GetEnvironmentSpecificationVersion
    {
        /// <summary>
        /// Azure Resource Manager resource envelope.
        /// 
        /// Uses Azure REST API version 2021-03-01-preview.
        /// </summary>
        public static Task<GetEnvironmentSpecificationVersionResult> InvokeAsync(GetEnvironmentSpecificationVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnvironmentSpecificationVersionResult>("azure-native:machinelearningservices:getEnvironmentSpecificationVersion", args ?? new GetEnvironmentSpecificationVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Azure Resource Manager resource envelope.
        /// 
        /// Uses Azure REST API version 2021-03-01-preview.
        /// </summary>
        public static Output<GetEnvironmentSpecificationVersionResult> Invoke(GetEnvironmentSpecificationVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnvironmentSpecificationVersionResult>("azure-native:machinelearningservices:getEnvironmentSpecificationVersion", args ?? new GetEnvironmentSpecificationVersionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Azure Resource Manager resource envelope.
        /// 
        /// Uses Azure REST API version 2021-03-01-preview.
        /// </summary>
        public static Output<GetEnvironmentSpecificationVersionResult> Invoke(GetEnvironmentSpecificationVersionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnvironmentSpecificationVersionResult>("azure-native:machinelearningservices:getEnvironmentSpecificationVersion", args ?? new GetEnvironmentSpecificationVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvironmentSpecificationVersionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Container name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Version identifier.
        /// </summary>
        [Input("version", required: true)]
        public string Version { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetEnvironmentSpecificationVersionArgs()
        {
        }
        public static new GetEnvironmentSpecificationVersionArgs Empty => new GetEnvironmentSpecificationVersionArgs();
    }

    public sealed class GetEnvironmentSpecificationVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Container name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Version identifier.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetEnvironmentSpecificationVersionInvokeArgs()
        {
        }
        public static new GetEnvironmentSpecificationVersionInvokeArgs Empty => new GetEnvironmentSpecificationVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnvironmentSpecificationVersionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// [Required] Additional attributes of the entity.
        /// </summary>
        public readonly Outputs.EnvironmentSpecificationVersionResponse Properties;
        /// <summary>
        /// System data associated with resource provider
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetEnvironmentSpecificationVersionResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.EnvironmentSpecificationVersionResponse properties,

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
