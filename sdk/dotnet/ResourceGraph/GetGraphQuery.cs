// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ResourceGraph
{
    public static class GetGraphQuery
    {
        /// <summary>
        /// Get a single graph query by its resourceName.
        /// 
        /// Uses Azure REST API version 2020-04-01-preview.
        /// 
        /// Other available API versions: 2018-09-01-preview, 2019-04-01, 2021-03-01, 2022-10-01, 2024-04-01.
        /// </summary>
        public static Task<GetGraphQueryResult> InvokeAsync(GetGraphQueryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGraphQueryResult>("azure-native:resourcegraph:getGraphQuery", args ?? new GetGraphQueryArgs(), options.WithDefaults());

        /// <summary>
        /// Get a single graph query by its resourceName.
        /// 
        /// Uses Azure REST API version 2020-04-01-preview.
        /// 
        /// Other available API versions: 2018-09-01-preview, 2019-04-01, 2021-03-01, 2022-10-01, 2024-04-01.
        /// </summary>
        public static Output<GetGraphQueryResult> Invoke(GetGraphQueryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGraphQueryResult>("azure-native:resourcegraph:getGraphQuery", args ?? new GetGraphQueryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a single graph query by its resourceName.
        /// 
        /// Uses Azure REST API version 2020-04-01-preview.
        /// 
        /// Other available API versions: 2018-09-01-preview, 2019-04-01, 2021-03-01, 2022-10-01, 2024-04-01.
        /// </summary>
        public static Output<GetGraphQueryResult> Invoke(GetGraphQueryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGraphQueryResult>("azure-native:resourcegraph:getGraphQuery", args ?? new GetGraphQueryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGraphQueryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Graph Query resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetGraphQueryArgs()
        {
        }
        public static new GetGraphQueryArgs Empty => new GetGraphQueryArgs();
    }

    public sealed class GetGraphQueryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Graph Query resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public GetGraphQueryInvokeArgs()
        {
        }
        public static new GetGraphQueryInvokeArgs Empty => new GetGraphQueryInvokeArgs();
    }


    [OutputType]
    public sealed class GetGraphQueryResult
    {
        /// <summary>
        /// The description of a graph query.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// This will be used to handle Optimistic Concurrency.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Azure resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The location of the resource
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Azure resource name. This is GUID value. The display name should be assigned within properties field.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// KQL query that will be graph.
        /// </summary>
        public readonly string Query;
        /// <summary>
        /// Enum indicating a type of graph query.
        /// </summary>
        public readonly string ResultKind;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Date and time in UTC of the last modification that was made to this graph query definition.
        /// </summary>
        public readonly string TimeModified;
        /// <summary>
        /// Azure resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetGraphQueryResult(
            string? description,

            string? etag,

            string id,

            string location,

            string name,

            string query,

            string resultKind,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string timeModified,

            string type)
        {
            Description = description;
            Etag = etag;
            Id = id;
            Location = location;
            Name = name;
            Query = query;
            ResultKind = resultKind;
            SystemData = systemData;
            Tags = tags;
            TimeModified = timeModified;
            Type = type;
        }
    }
}
