// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetVerifierWorkspace
    {
        /// <summary>
        /// Instance of Verifier Workspace.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2024-01-01-preview, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetVerifierWorkspaceResult> InvokeAsync(GetVerifierWorkspaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVerifierWorkspaceResult>("azure-native:network:getVerifierWorkspace", args ?? new GetVerifierWorkspaceArgs(), options.WithDefaults());

        /// <summary>
        /// Instance of Verifier Workspace.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2024-01-01-preview, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVerifierWorkspaceResult> Invoke(GetVerifierWorkspaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVerifierWorkspaceResult>("azure-native:network:getVerifierWorkspace", args ?? new GetVerifierWorkspaceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Instance of Verifier Workspace.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2024-01-01-preview, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVerifierWorkspaceResult> Invoke(GetVerifierWorkspaceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVerifierWorkspaceResult>("azure-native:network:getVerifierWorkspace", args ?? new GetVerifierWorkspaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVerifierWorkspaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the network manager.
        /// </summary>
        [Input("networkManagerName", required: true)]
        public string NetworkManagerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Workspace name.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetVerifierWorkspaceArgs()
        {
        }
        public static new GetVerifierWorkspaceArgs Empty => new GetVerifierWorkspaceArgs();
    }

    public sealed class GetVerifierWorkspaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the network manager.
        /// </summary>
        [Input("networkManagerName", required: true)]
        public Input<string> NetworkManagerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Workspace name.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetVerifierWorkspaceInvokeArgs()
        {
        }
        public static new GetVerifierWorkspaceInvokeArgs Empty => new GetVerifierWorkspaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetVerifierWorkspaceResult
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
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of Verifier Workspace resource.
        /// </summary>
        public readonly Outputs.VerifierWorkspacePropertiesResponse Properties;
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

        [OutputConstructor]
        private GetVerifierWorkspaceResult(
            string azureApiVersion,

            string id,

            string location,

            string name,

            Outputs.VerifierWorkspacePropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
