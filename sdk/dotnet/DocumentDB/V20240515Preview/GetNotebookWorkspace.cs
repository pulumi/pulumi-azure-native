// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DocumentDB.V20240515Preview
{
    public static class GetNotebookWorkspace
    {
        /// <summary>
        /// Gets the notebook workspace for a Cosmos DB account.
        /// </summary>
        public static Task<GetNotebookWorkspaceResult> InvokeAsync(GetNotebookWorkspaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNotebookWorkspaceResult>("azure-native:documentdb/v20240515preview:getNotebookWorkspace", args ?? new GetNotebookWorkspaceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the notebook workspace for a Cosmos DB account.
        /// </summary>
        public static Output<GetNotebookWorkspaceResult> Invoke(GetNotebookWorkspaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNotebookWorkspaceResult>("azure-native:documentdb/v20240515preview:getNotebookWorkspace", args ?? new GetNotebookWorkspaceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the notebook workspace for a Cosmos DB account.
        /// </summary>
        public static Output<GetNotebookWorkspaceResult> Invoke(GetNotebookWorkspaceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNotebookWorkspaceResult>("azure-native:documentdb/v20240515preview:getNotebookWorkspace", args ?? new GetNotebookWorkspaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNotebookWorkspaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the notebook workspace resource.
        /// </summary>
        [Input("notebookWorkspaceName", required: true)]
        public string NotebookWorkspaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNotebookWorkspaceArgs()
        {
        }
        public static new GetNotebookWorkspaceArgs Empty => new GetNotebookWorkspaceArgs();
    }

    public sealed class GetNotebookWorkspaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the notebook workspace resource.
        /// </summary>
        [Input("notebookWorkspaceName", required: true)]
        public Input<string> NotebookWorkspaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetNotebookWorkspaceInvokeArgs()
        {
        }
        public static new GetNotebookWorkspaceInvokeArgs Empty => new GetNotebookWorkspaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetNotebookWorkspaceResult
    {
        /// <summary>
        /// The unique resource identifier of the database account.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the database account.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the endpoint of Notebook server.
        /// </summary>
        public readonly string NotebookServerEndpoint;
        /// <summary>
        /// Status of the notebook workspace. Possible values are: Creating, Online, Deleting, Failed, Updating.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetNotebookWorkspaceResult(
            string id,

            string name,

            string notebookServerEndpoint,

            string status,

            string type)
        {
            Id = id;
            Name = name;
            NotebookServerEndpoint = notebookServerEndpoint;
            Status = status;
            Type = type;
        }
    }
}
