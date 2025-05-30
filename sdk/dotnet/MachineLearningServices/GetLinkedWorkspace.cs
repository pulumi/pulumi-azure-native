// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices
{
    public static class GetLinkedWorkspace
    {
        /// <summary>
        /// Get the detail of a linked workspace.
        /// 
        /// Uses Azure REST API version 2020-05-15-preview.
        /// 
        /// Other available API versions: 2020-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetLinkedWorkspaceResult> InvokeAsync(GetLinkedWorkspaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLinkedWorkspaceResult>("azure-native:machinelearningservices:getLinkedWorkspace", args ?? new GetLinkedWorkspaceArgs(), options.WithDefaults());

        /// <summary>
        /// Get the detail of a linked workspace.
        /// 
        /// Uses Azure REST API version 2020-05-15-preview.
        /// 
        /// Other available API versions: 2020-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetLinkedWorkspaceResult> Invoke(GetLinkedWorkspaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLinkedWorkspaceResult>("azure-native:machinelearningservices:getLinkedWorkspace", args ?? new GetLinkedWorkspaceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the detail of a linked workspace.
        /// 
        /// Uses Azure REST API version 2020-05-15-preview.
        /// 
        /// Other available API versions: 2020-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetLinkedWorkspaceResult> Invoke(GetLinkedWorkspaceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLinkedWorkspaceResult>("azure-native:machinelearningservices:getLinkedWorkspace", args ?? new GetLinkedWorkspaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLinkedWorkspaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Friendly name of the linked workspace
        /// </summary>
        [Input("linkName", required: true)]
        public string LinkName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group in which workspace is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetLinkedWorkspaceArgs()
        {
        }
        public static new GetLinkedWorkspaceArgs Empty => new GetLinkedWorkspaceArgs();
    }

    public sealed class GetLinkedWorkspaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Friendly name of the linked workspace
        /// </summary>
        [Input("linkName", required: true)]
        public Input<string> LinkName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group in which workspace is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetLinkedWorkspaceInvokeArgs()
        {
        }
        public static new GetLinkedWorkspaceInvokeArgs Empty => new GetLinkedWorkspaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetLinkedWorkspaceResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// ResourceId of the link of the linked workspace.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Friendly name of the linked workspace.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// LinkedWorkspace specific properties.
        /// </summary>
        public readonly Outputs.LinkedWorkspacePropsResponse Properties;
        /// <summary>
        /// Resource type of linked workspace.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetLinkedWorkspaceResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.LinkedWorkspacePropsResponse properties,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
