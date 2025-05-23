// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement
{
    public static class GetAuthorizationLoginLinkPost
    {
        /// <summary>
        /// Gets authorization login links.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetAuthorizationLoginLinkPostResult> InvokeAsync(GetAuthorizationLoginLinkPostArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAuthorizationLoginLinkPostResult>("azure-native:apimanagement:getAuthorizationLoginLinkPost", args ?? new GetAuthorizationLoginLinkPostArgs(), options.WithDefaults());

        /// <summary>
        /// Gets authorization login links.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAuthorizationLoginLinkPostResult> Invoke(GetAuthorizationLoginLinkPostInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthorizationLoginLinkPostResult>("azure-native:apimanagement:getAuthorizationLoginLinkPost", args ?? new GetAuthorizationLoginLinkPostInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets authorization login links.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAuthorizationLoginLinkPostResult> Invoke(GetAuthorizationLoginLinkPostInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthorizationLoginLinkPostResult>("azure-native:apimanagement:getAuthorizationLoginLinkPost", args ?? new GetAuthorizationLoginLinkPostInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthorizationLoginLinkPostArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier of the authorization.
        /// </summary>
        [Input("authorizationId", required: true)]
        public string AuthorizationId { get; set; } = null!;

        /// <summary>
        /// Identifier of the authorization provider.
        /// </summary>
        [Input("authorizationProviderId", required: true)]
        public string AuthorizationProviderId { get; set; } = null!;

        /// <summary>
        /// The redirect URL after login has completed.
        /// </summary>
        [Input("postLoginRedirectUrl")]
        public string? PostLoginRedirectUrl { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetAuthorizationLoginLinkPostArgs()
        {
        }
        public static new GetAuthorizationLoginLinkPostArgs Empty => new GetAuthorizationLoginLinkPostArgs();
    }

    public sealed class GetAuthorizationLoginLinkPostInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier of the authorization.
        /// </summary>
        [Input("authorizationId", required: true)]
        public Input<string> AuthorizationId { get; set; } = null!;

        /// <summary>
        /// Identifier of the authorization provider.
        /// </summary>
        [Input("authorizationProviderId", required: true)]
        public Input<string> AuthorizationProviderId { get; set; } = null!;

        /// <summary>
        /// The redirect URL after login has completed.
        /// </summary>
        [Input("postLoginRedirectUrl")]
        public Input<string>? PostLoginRedirectUrl { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetAuthorizationLoginLinkPostInvokeArgs()
        {
        }
        public static new GetAuthorizationLoginLinkPostInvokeArgs Empty => new GetAuthorizationLoginLinkPostInvokeArgs();
    }


    [OutputType]
    public sealed class GetAuthorizationLoginLinkPostResult
    {
        /// <summary>
        /// The login link
        /// </summary>
        public readonly string? LoginLink;

        [OutputConstructor]
        private GetAuthorizationLoginLinkPostResult(string? loginLink)
        {
            LoginLink = loginLink;
        }
    }
}
