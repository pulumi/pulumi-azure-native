// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Confluent
{
    public static class ListAccessUsers
    {
        /// <summary>
        /// Organization users details
        /// 
        /// Uses Azure REST API version 2024-07-01.
        /// 
        /// Other available API versions: 2023-08-22, 2024-02-13. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native confluent [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListAccessUsersResult> InvokeAsync(ListAccessUsersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListAccessUsersResult>("azure-native:confluent:listAccessUsers", args ?? new ListAccessUsersArgs(), options.WithDefaults());

        /// <summary>
        /// Organization users details
        /// 
        /// Uses Azure REST API version 2024-07-01.
        /// 
        /// Other available API versions: 2023-08-22, 2024-02-13. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native confluent [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListAccessUsersResult> Invoke(ListAccessUsersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListAccessUsersResult>("azure-native:confluent:listAccessUsers", args ?? new ListAccessUsersInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Organization users details
        /// 
        /// Uses Azure REST API version 2024-07-01.
        /// 
        /// Other available API versions: 2023-08-22, 2024-02-13. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native confluent [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListAccessUsersResult> Invoke(ListAccessUsersInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListAccessUsersResult>("azure-native:confluent:listAccessUsers", args ?? new ListAccessUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListAccessUsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Organization resource name
        /// </summary>
        [Input("organizationName", required: true)]
        public string OrganizationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("searchFilters")]
        private Dictionary<string, string>? _searchFilters;

        /// <summary>
        /// Search filters for the request
        /// </summary>
        public Dictionary<string, string> SearchFilters
        {
            get => _searchFilters ?? (_searchFilters = new Dictionary<string, string>());
            set => _searchFilters = value;
        }

        public ListAccessUsersArgs()
        {
        }
        public static new ListAccessUsersArgs Empty => new ListAccessUsersArgs();
    }

    public sealed class ListAccessUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Organization resource name
        /// </summary>
        [Input("organizationName", required: true)]
        public Input<string> OrganizationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("searchFilters")]
        private InputMap<string>? _searchFilters;

        /// <summary>
        /// Search filters for the request
        /// </summary>
        public InputMap<string> SearchFilters
        {
            get => _searchFilters ?? (_searchFilters = new InputMap<string>());
            set => _searchFilters = value;
        }

        public ListAccessUsersInvokeArgs()
        {
        }
        public static new ListAccessUsersInvokeArgs Empty => new ListAccessUsersInvokeArgs();
    }


    [OutputType]
    public sealed class ListAccessUsersResult
    {
        /// <summary>
        /// Data of the users list
        /// </summary>
        public readonly ImmutableArray<Outputs.UserRecordResponse> Data;
        /// <summary>
        /// Type of response
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Metadata of the list
        /// </summary>
        public readonly Outputs.ConfluentListMetadataResponse? Metadata;

        [OutputConstructor]
        private ListAccessUsersResult(
            ImmutableArray<Outputs.UserRecordResponse> data,

            string? kind,

            Outputs.ConfluentListMetadataResponse? metadata)
        {
            Data = data;
            Kind = kind;
            Metadata = metadata;
        }
    }
}
