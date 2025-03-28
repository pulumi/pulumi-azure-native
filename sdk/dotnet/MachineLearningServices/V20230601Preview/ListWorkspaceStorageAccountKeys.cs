// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20230601Preview
{
    public static class ListWorkspaceStorageAccountKeys
    {
        public static Task<ListWorkspaceStorageAccountKeysResult> InvokeAsync(ListWorkspaceStorageAccountKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListWorkspaceStorageAccountKeysResult>("azure-native:machinelearningservices/v20230601preview:listWorkspaceStorageAccountKeys", args ?? new ListWorkspaceStorageAccountKeysArgs(), options.WithDefaults());

        public static Output<ListWorkspaceStorageAccountKeysResult> Invoke(ListWorkspaceStorageAccountKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListWorkspaceStorageAccountKeysResult>("azure-native:machinelearningservices/v20230601preview:listWorkspaceStorageAccountKeys", args ?? new ListWorkspaceStorageAccountKeysInvokeArgs(), options.WithDefaults());

        public static Output<ListWorkspaceStorageAccountKeysResult> Invoke(ListWorkspaceStorageAccountKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListWorkspaceStorageAccountKeysResult>("azure-native:machinelearningservices/v20230601preview:listWorkspaceStorageAccountKeys", args ?? new ListWorkspaceStorageAccountKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListWorkspaceStorageAccountKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public ListWorkspaceStorageAccountKeysArgs()
        {
        }
        public static new ListWorkspaceStorageAccountKeysArgs Empty => new ListWorkspaceStorageAccountKeysArgs();
    }

    public sealed class ListWorkspaceStorageAccountKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public ListWorkspaceStorageAccountKeysInvokeArgs()
        {
        }
        public static new ListWorkspaceStorageAccountKeysInvokeArgs Empty => new ListWorkspaceStorageAccountKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListWorkspaceStorageAccountKeysResult
    {
        /// <summary>
        /// The access key of the storage
        /// </summary>
        public readonly string UserStorageKey;

        [OutputConstructor]
        private ListWorkspaceStorageAccountKeysResult(string userStorageKey)
        {
            UserStorageKey = userStorageKey;
        }
    }
}
