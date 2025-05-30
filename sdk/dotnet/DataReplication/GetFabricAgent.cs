// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataReplication
{
    public static class GetFabricAgent
    {
        /// <summary>
        /// Gets the details of the fabric agent.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// </summary>
        public static Task<GetFabricAgentResult> InvokeAsync(GetFabricAgentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFabricAgentResult>("azure-native:datareplication:getFabricAgent", args ?? new GetFabricAgentArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the fabric agent.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// </summary>
        public static Output<GetFabricAgentResult> Invoke(GetFabricAgentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFabricAgentResult>("azure-native:datareplication:getFabricAgent", args ?? new GetFabricAgentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the fabric agent.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// </summary>
        public static Output<GetFabricAgentResult> Invoke(GetFabricAgentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFabricAgentResult>("azure-native:datareplication:getFabricAgent", args ?? new GetFabricAgentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFabricAgentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The fabric agent name.
        /// </summary>
        [Input("fabricAgentName", required: true)]
        public string FabricAgentName { get; set; } = null!;

        /// <summary>
        /// The fabric name.
        /// </summary>
        [Input("fabricName", required: true)]
        public string FabricName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetFabricAgentArgs()
        {
        }
        public static new GetFabricAgentArgs Empty => new GetFabricAgentArgs();
    }

    public sealed class GetFabricAgentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The fabric agent name.
        /// </summary>
        [Input("fabricAgentName", required: true)]
        public Input<string> FabricAgentName { get; set; } = null!;

        /// <summary>
        /// The fabric name.
        /// </summary>
        [Input("fabricName", required: true)]
        public Input<string> FabricName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetFabricAgentInvokeArgs()
        {
        }
        public static new GetFabricAgentInvokeArgs Empty => new GetFabricAgentInvokeArgs();
    }


    [OutputType]
    public sealed class GetFabricAgentResult
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
        public readonly Outputs.FabricAgentModelPropertiesResponse Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetFabricAgentResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.FabricAgentModelPropertiesResponse properties,

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
