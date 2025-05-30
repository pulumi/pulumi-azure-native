// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData
{
    public static class GetActiveDirectoryConnector
    {
        /// <summary>
        /// Retrieves an Active Directory connector resource
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2023-01-15-preview, 2024-05-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurearcdata [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetActiveDirectoryConnectorResult> InvokeAsync(GetActiveDirectoryConnectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetActiveDirectoryConnectorResult>("azure-native:azurearcdata:getActiveDirectoryConnector", args ?? new GetActiveDirectoryConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves an Active Directory connector resource
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2023-01-15-preview, 2024-05-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurearcdata [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetActiveDirectoryConnectorResult> Invoke(GetActiveDirectoryConnectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetActiveDirectoryConnectorResult>("azure-native:azurearcdata:getActiveDirectoryConnector", args ?? new GetActiveDirectoryConnectorInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves an Active Directory connector resource
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2023-01-15-preview, 2024-05-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurearcdata [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetActiveDirectoryConnectorResult> Invoke(GetActiveDirectoryConnectorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetActiveDirectoryConnectorResult>("azure-native:azurearcdata:getActiveDirectoryConnector", args ?? new GetActiveDirectoryConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetActiveDirectoryConnectorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Active Directory connector instance
        /// </summary>
        [Input("activeDirectoryConnectorName", required: true)]
        public string ActiveDirectoryConnectorName { get; set; } = null!;

        /// <summary>
        /// The name of the data controller
        /// </summary>
        [Input("dataControllerName", required: true)]
        public string DataControllerName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetActiveDirectoryConnectorArgs()
        {
        }
        public static new GetActiveDirectoryConnectorArgs Empty => new GetActiveDirectoryConnectorArgs();
    }

    public sealed class GetActiveDirectoryConnectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Active Directory connector instance
        /// </summary>
        [Input("activeDirectoryConnectorName", required: true)]
        public Input<string> ActiveDirectoryConnectorName { get; set; } = null!;

        /// <summary>
        /// The name of the data controller
        /// </summary>
        [Input("dataControllerName", required: true)]
        public Input<string> DataControllerName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetActiveDirectoryConnectorInvokeArgs()
        {
        }
        public static new GetActiveDirectoryConnectorInvokeArgs Empty => new GetActiveDirectoryConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetActiveDirectoryConnectorResult
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
        /// null
        /// </summary>
        public readonly Outputs.ActiveDirectoryConnectorPropertiesResponse Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetActiveDirectoryConnectorResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.ActiveDirectoryConnectorPropertiesResponse properties,

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
