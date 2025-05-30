// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Resources
{
    public static class GetDeploymentAtSubscriptionScope
    {
        /// <summary>
        /// Gets a deployment.
        /// 
        /// Uses Azure REST API version 2024-03-01.
        /// 
        /// Other available API versions: 2020-10-01, 2021-01-01, 2021-04-01, 2022-09-01, 2023-07-01, 2024-07-01, 2024-11-01, 2025-03-01, 2025-04-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetDeploymentAtSubscriptionScopeResult> InvokeAsync(GetDeploymentAtSubscriptionScopeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeploymentAtSubscriptionScopeResult>("azure-native:resources:getDeploymentAtSubscriptionScope", args ?? new GetDeploymentAtSubscriptionScopeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a deployment.
        /// 
        /// Uses Azure REST API version 2024-03-01.
        /// 
        /// Other available API versions: 2020-10-01, 2021-01-01, 2021-04-01, 2022-09-01, 2023-07-01, 2024-07-01, 2024-11-01, 2025-03-01, 2025-04-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDeploymentAtSubscriptionScopeResult> Invoke(GetDeploymentAtSubscriptionScopeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeploymentAtSubscriptionScopeResult>("azure-native:resources:getDeploymentAtSubscriptionScope", args ?? new GetDeploymentAtSubscriptionScopeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a deployment.
        /// 
        /// Uses Azure REST API version 2024-03-01.
        /// 
        /// Other available API versions: 2020-10-01, 2021-01-01, 2021-04-01, 2022-09-01, 2023-07-01, 2024-07-01, 2024-11-01, 2025-03-01, 2025-04-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDeploymentAtSubscriptionScopeResult> Invoke(GetDeploymentAtSubscriptionScopeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeploymentAtSubscriptionScopeResult>("azure-native:resources:getDeploymentAtSubscriptionScope", args ?? new GetDeploymentAtSubscriptionScopeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeploymentAtSubscriptionScopeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the deployment.
        /// </summary>
        [Input("deploymentName", required: true)]
        public string DeploymentName { get; set; } = null!;

        public GetDeploymentAtSubscriptionScopeArgs()
        {
        }
        public static new GetDeploymentAtSubscriptionScopeArgs Empty => new GetDeploymentAtSubscriptionScopeArgs();
    }

    public sealed class GetDeploymentAtSubscriptionScopeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the deployment.
        /// </summary>
        [Input("deploymentName", required: true)]
        public Input<string> DeploymentName { get; set; } = null!;

        public GetDeploymentAtSubscriptionScopeInvokeArgs()
        {
        }
        public static new GetDeploymentAtSubscriptionScopeInvokeArgs Empty => new GetDeploymentAtSubscriptionScopeInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeploymentAtSubscriptionScopeResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The ID of the deployment.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// the location of the deployment.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the deployment.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Deployment properties.
        /// </summary>
        public readonly Outputs.DeploymentPropertiesExtendedResponse Properties;
        /// <summary>
        /// Deployment tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the deployment.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDeploymentAtSubscriptionScopeResult(
            string azureApiVersion,

            string id,

            string? location,

            string name,

            Outputs.DeploymentPropertiesExtendedResponse properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}
