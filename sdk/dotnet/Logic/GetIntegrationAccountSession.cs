// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Logic
{
    public static class GetIntegrationAccountSession
    {
        /// <summary>
        /// Gets an integration account session.
        /// 
        /// Uses Azure REST API version 2019-05-01.
        /// 
        /// Other available API versions: 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetIntegrationAccountSessionResult> InvokeAsync(GetIntegrationAccountSessionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationAccountSessionResult>("azure-native:logic:getIntegrationAccountSession", args ?? new GetIntegrationAccountSessionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an integration account session.
        /// 
        /// Uses Azure REST API version 2019-05-01.
        /// 
        /// Other available API versions: 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetIntegrationAccountSessionResult> Invoke(GetIntegrationAccountSessionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIntegrationAccountSessionResult>("azure-native:logic:getIntegrationAccountSession", args ?? new GetIntegrationAccountSessionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an integration account session.
        /// 
        /// Uses Azure REST API version 2019-05-01.
        /// 
        /// Other available API versions: 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetIntegrationAccountSessionResult> Invoke(GetIntegrationAccountSessionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIntegrationAccountSessionResult>("azure-native:logic:getIntegrationAccountSession", args ?? new GetIntegrationAccountSessionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIntegrationAccountSessionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The integration account name.
        /// </summary>
        [Input("integrationAccountName", required: true)]
        public string IntegrationAccountName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The integration account session name.
        /// </summary>
        [Input("sessionName", required: true)]
        public string SessionName { get; set; } = null!;

        public GetIntegrationAccountSessionArgs()
        {
        }
        public static new GetIntegrationAccountSessionArgs Empty => new GetIntegrationAccountSessionArgs();
    }

    public sealed class GetIntegrationAccountSessionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The integration account name.
        /// </summary>
        [Input("integrationAccountName", required: true)]
        public Input<string> IntegrationAccountName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The integration account session name.
        /// </summary>
        [Input("sessionName", required: true)]
        public Input<string> SessionName { get; set; } = null!;

        public GetIntegrationAccountSessionInvokeArgs()
        {
        }
        public static new GetIntegrationAccountSessionInvokeArgs Empty => new GetIntegrationAccountSessionInvokeArgs();
    }


    [OutputType]
    public sealed class GetIntegrationAccountSessionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The changed time.
        /// </summary>
        public readonly string ChangedTime;
        /// <summary>
        /// The session content.
        /// </summary>
        public readonly object? Content;
        /// <summary>
        /// The created time.
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// The resource id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Gets the resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Gets the resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetIntegrationAccountSessionResult(
            string azureApiVersion,

            string changedTime,

            object? content,

            string createdTime,

            string id,

            string? location,

            string name,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            ChangedTime = changedTime;
            Content = content;
            CreatedTime = createdTime;
            Id = id;
            Location = location;
            Name = name;
            Tags = tags;
            Type = type;
        }
    }
}
