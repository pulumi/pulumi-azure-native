// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Resources
{
    public static class GetTagAtScope
    {
        /// <summary>
        /// Wrapper resource for tags API requests and responses.
        /// 
        /// Uses Azure REST API version 2024-03-01.
        /// 
        /// Other available API versions: 2020-10-01, 2021-01-01, 2021-04-01, 2022-09-01, 2023-07-01, 2024-07-01, 2024-11-01, 2025-03-01, 2025-04-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetTagAtScopeResult> InvokeAsync(GetTagAtScopeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTagAtScopeResult>("azure-native:resources:getTagAtScope", args ?? new GetTagAtScopeArgs(), options.WithDefaults());

        /// <summary>
        /// Wrapper resource for tags API requests and responses.
        /// 
        /// Uses Azure REST API version 2024-03-01.
        /// 
        /// Other available API versions: 2020-10-01, 2021-01-01, 2021-04-01, 2022-09-01, 2023-07-01, 2024-07-01, 2024-11-01, 2025-03-01, 2025-04-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTagAtScopeResult> Invoke(GetTagAtScopeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTagAtScopeResult>("azure-native:resources:getTagAtScope", args ?? new GetTagAtScopeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Wrapper resource for tags API requests and responses.
        /// 
        /// Uses Azure REST API version 2024-03-01.
        /// 
        /// Other available API versions: 2020-10-01, 2021-01-01, 2021-04-01, 2022-09-01, 2023-07-01, 2024-07-01, 2024-11-01, 2025-03-01, 2025-04-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTagAtScopeResult> Invoke(GetTagAtScopeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTagAtScopeResult>("azure-native:resources:getTagAtScope", args ?? new GetTagAtScopeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTagAtScopeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource scope.
        /// </summary>
        [Input("scope", required: true)]
        public string Scope { get; set; } = null!;

        public GetTagAtScopeArgs()
        {
        }
        public static new GetTagAtScopeArgs Empty => new GetTagAtScopeArgs();
    }

    public sealed class GetTagAtScopeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource scope.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public GetTagAtScopeInvokeArgs()
        {
        }
        public static new GetTagAtScopeInvokeArgs Empty => new GetTagAtScopeInvokeArgs();
    }


    [OutputType]
    public sealed class GetTagAtScopeResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The ID of the tags wrapper resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the tags wrapper resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The set of tags.
        /// </summary>
        public readonly Outputs.TagsResponse Properties;
        /// <summary>
        /// The type of the tags wrapper resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTagAtScopeResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.TagsResponse properties,

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
