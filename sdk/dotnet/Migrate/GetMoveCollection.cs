// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate
{
    public static class GetMoveCollection
    {
        /// <summary>
        /// Gets the move collection.
        /// 
        /// Uses Azure REST API version 2023-08-01.
        /// 
        /// Other available API versions: 2019-10-01-preview, 2021-01-01, 2021-08-01, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetMoveCollectionResult> InvokeAsync(GetMoveCollectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMoveCollectionResult>("azure-native:migrate:getMoveCollection", args ?? new GetMoveCollectionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the move collection.
        /// 
        /// Uses Azure REST API version 2023-08-01.
        /// 
        /// Other available API versions: 2019-10-01-preview, 2021-01-01, 2021-08-01, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMoveCollectionResult> Invoke(GetMoveCollectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMoveCollectionResult>("azure-native:migrate:getMoveCollection", args ?? new GetMoveCollectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the move collection.
        /// 
        /// Uses Azure REST API version 2023-08-01.
        /// 
        /// Other available API versions: 2019-10-01-preview, 2021-01-01, 2021-08-01, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMoveCollectionResult> Invoke(GetMoveCollectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMoveCollectionResult>("azure-native:migrate:getMoveCollection", args ?? new GetMoveCollectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMoveCollectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Move Collection Name.
        /// </summary>
        [Input("moveCollectionName", required: true)]
        public string MoveCollectionName { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMoveCollectionArgs()
        {
        }
        public static new GetMoveCollectionArgs Empty => new GetMoveCollectionArgs();
    }

    public sealed class GetMoveCollectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Move Collection Name.
        /// </summary>
        [Input("moveCollectionName", required: true)]
        public Input<string> MoveCollectionName { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetMoveCollectionInvokeArgs()
        {
        }
        public static new GetMoveCollectionInvokeArgs Empty => new GetMoveCollectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetMoveCollectionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The etag of the resource.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Fully qualified resource Id for the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Defines the MSI properties of the Move Collection.
        /// </summary>
        public readonly Outputs.IdentityResponse? Identity;
        /// <summary>
        /// The geo-location where the resource lives.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Defines the move collection properties.
        /// </summary>
        public readonly Outputs.MoveCollectionPropertiesResponse Properties;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMoveCollectionResult(
            string azureApiVersion,

            string etag,

            string id,

            Outputs.IdentityResponse? identity,

            string? location,

            string name,

            Outputs.MoveCollectionPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Etag = etag;
            Id = id;
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
