// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate
{
    public static class GetMoveResource
    {
        /// <summary>
        /// Gets the Move Resource.
        /// 
        /// Uses Azure REST API version 2023-08-01.
        /// 
        /// Other available API versions: 2019-10-01-preview, 2021-01-01, 2021-08-01, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetMoveResourceResult> InvokeAsync(GetMoveResourceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMoveResourceResult>("azure-native:migrate:getMoveResource", args ?? new GetMoveResourceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the Move Resource.
        /// 
        /// Uses Azure REST API version 2023-08-01.
        /// 
        /// Other available API versions: 2019-10-01-preview, 2021-01-01, 2021-08-01, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMoveResourceResult> Invoke(GetMoveResourceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMoveResourceResult>("azure-native:migrate:getMoveResource", args ?? new GetMoveResourceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the Move Resource.
        /// 
        /// Uses Azure REST API version 2023-08-01.
        /// 
        /// Other available API versions: 2019-10-01-preview, 2021-01-01, 2021-08-01, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMoveResourceResult> Invoke(GetMoveResourceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMoveResourceResult>("azure-native:migrate:getMoveResource", args ?? new GetMoveResourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMoveResourceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Move Collection Name.
        /// </summary>
        [Input("moveCollectionName", required: true)]
        public string MoveCollectionName { get; set; } = null!;

        /// <summary>
        /// The Move Resource Name.
        /// </summary>
        [Input("moveResourceName", required: true)]
        public string MoveResourceName { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMoveResourceArgs()
        {
        }
        public static new GetMoveResourceArgs Empty => new GetMoveResourceArgs();
    }

    public sealed class GetMoveResourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Move Collection Name.
        /// </summary>
        [Input("moveCollectionName", required: true)]
        public Input<string> MoveCollectionName { get; set; } = null!;

        /// <summary>
        /// The Move Resource Name.
        /// </summary>
        [Input("moveResourceName", required: true)]
        public Input<string> MoveResourceName { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetMoveResourceInvokeArgs()
        {
        }
        public static new GetMoveResourceInvokeArgs Empty => new GetMoveResourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetMoveResourceResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource Id for the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Defines the move resource properties.
        /// </summary>
        public readonly Outputs.MoveResourcePropertiesResponse Properties;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMoveResourceResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.MoveResourcePropertiesResponse properties,

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
