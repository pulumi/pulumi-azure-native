// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse
{
    public static class ListKustoPoolFollowerDatabases
    {
        /// <summary>
        /// Returns a list of databases that are owned by this Kusto Pool and were followed by another Kusto Pool.
        /// 
        /// Uses Azure REST API version 2021-06-01-preview.
        /// </summary>
        public static Task<ListKustoPoolFollowerDatabasesResult> InvokeAsync(ListKustoPoolFollowerDatabasesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListKustoPoolFollowerDatabasesResult>("azure-native:synapse:listKustoPoolFollowerDatabases", args ?? new ListKustoPoolFollowerDatabasesArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a list of databases that are owned by this Kusto Pool and were followed by another Kusto Pool.
        /// 
        /// Uses Azure REST API version 2021-06-01-preview.
        /// </summary>
        public static Output<ListKustoPoolFollowerDatabasesResult> Invoke(ListKustoPoolFollowerDatabasesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListKustoPoolFollowerDatabasesResult>("azure-native:synapse:listKustoPoolFollowerDatabases", args ?? new ListKustoPoolFollowerDatabasesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a list of databases that are owned by this Kusto Pool and were followed by another Kusto Pool.
        /// 
        /// Uses Azure REST API version 2021-06-01-preview.
        /// </summary>
        public static Output<ListKustoPoolFollowerDatabasesResult> Invoke(ListKustoPoolFollowerDatabasesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListKustoPoolFollowerDatabasesResult>("azure-native:synapse:listKustoPoolFollowerDatabases", args ?? new ListKustoPoolFollowerDatabasesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListKustoPoolFollowerDatabasesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kusto pool.
        /// </summary>
        [Input("kustoPoolName", required: true)]
        public string KustoPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public ListKustoPoolFollowerDatabasesArgs()
        {
        }
        public static new ListKustoPoolFollowerDatabasesArgs Empty => new ListKustoPoolFollowerDatabasesArgs();
    }

    public sealed class ListKustoPoolFollowerDatabasesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kusto pool.
        /// </summary>
        [Input("kustoPoolName", required: true)]
        public Input<string> KustoPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public ListKustoPoolFollowerDatabasesInvokeArgs()
        {
        }
        public static new ListKustoPoolFollowerDatabasesInvokeArgs Empty => new ListKustoPoolFollowerDatabasesInvokeArgs();
    }


    [OutputType]
    public sealed class ListKustoPoolFollowerDatabasesResult
    {
        /// <summary>
        /// The list of follower database result.
        /// </summary>
        public readonly ImmutableArray<Outputs.FollowerDatabaseDefinitionResponse> Value;

        [OutputConstructor]
        private ListKustoPoolFollowerDatabasesResult(ImmutableArray<Outputs.FollowerDatabaseDefinitionResponse> value)
        {
            Value = value;
        }
    }
}
