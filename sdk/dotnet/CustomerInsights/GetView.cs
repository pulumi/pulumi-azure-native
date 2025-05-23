// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CustomerInsights
{
    public static class GetView
    {
        /// <summary>
        /// Gets a view in the hub.
        /// 
        /// Uses Azure REST API version 2017-04-26.
        /// </summary>
        public static Task<GetViewResult> InvokeAsync(GetViewArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetViewResult>("azure-native:customerinsights:getView", args ?? new GetViewArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a view in the hub.
        /// 
        /// Uses Azure REST API version 2017-04-26.
        /// </summary>
        public static Output<GetViewResult> Invoke(GetViewInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetViewResult>("azure-native:customerinsights:getView", args ?? new GetViewInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a view in the hub.
        /// 
        /// Uses Azure REST API version 2017-04-26.
        /// </summary>
        public static Output<GetViewResult> Invoke(GetViewInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetViewResult>("azure-native:customerinsights:getView", args ?? new GetViewInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetViewArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the hub.
        /// </summary>
        [Input("hubName", required: true)]
        public string HubName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The user ID. Use * to retrieve hub level view.
        /// </summary>
        [Input("userId", required: true)]
        public string UserId { get; set; } = null!;

        /// <summary>
        /// The name of the view.
        /// </summary>
        [Input("viewName", required: true)]
        public string ViewName { get; set; } = null!;

        public GetViewArgs()
        {
        }
        public static new GetViewArgs Empty => new GetViewArgs();
    }

    public sealed class GetViewInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the hub.
        /// </summary>
        [Input("hubName", required: true)]
        public Input<string> HubName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The user ID. Use * to retrieve hub level view.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        /// <summary>
        /// The name of the view.
        /// </summary>
        [Input("viewName", required: true)]
        public Input<string> ViewName { get; set; } = null!;

        public GetViewInvokeArgs()
        {
        }
        public static new GetViewInvokeArgs Empty => new GetViewInvokeArgs();
    }


    [OutputType]
    public sealed class GetViewResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Date time when view was last modified.
        /// </summary>
        public readonly string Changed;
        /// <summary>
        /// Date time when view was created.
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// View definition.
        /// </summary>
        public readonly string Definition;
        /// <summary>
        /// Localized display name for the view.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? DisplayName;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// the hub name.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// the user ID.
        /// </summary>
        public readonly string? UserId;
        /// <summary>
        /// Name of the view.
        /// </summary>
        public readonly string ViewName;

        [OutputConstructor]
        private GetViewResult(
            string azureApiVersion,

            string changed,

            string created,

            string definition,

            ImmutableDictionary<string, string>? displayName,

            string id,

            string name,

            string tenantId,

            string type,

            string? userId,

            string viewName)
        {
            AzureApiVersion = azureApiVersion;
            Changed = changed;
            Created = created;
            Definition = definition;
            DisplayName = displayName;
            Id = id;
            Name = name;
            TenantId = tenantId;
            Type = type;
            UserId = userId;
            ViewName = viewName;
        }
    }
}
