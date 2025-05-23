// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement
{
    public static class GetAuthorizationProvider
    {
        /// <summary>
        /// Gets the details of the authorization provider specified by its identifier.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetAuthorizationProviderResult> InvokeAsync(GetAuthorizationProviderArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAuthorizationProviderResult>("azure-native:apimanagement:getAuthorizationProvider", args ?? new GetAuthorizationProviderArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the authorization provider specified by its identifier.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAuthorizationProviderResult> Invoke(GetAuthorizationProviderInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthorizationProviderResult>("azure-native:apimanagement:getAuthorizationProvider", args ?? new GetAuthorizationProviderInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the authorization provider specified by its identifier.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAuthorizationProviderResult> Invoke(GetAuthorizationProviderInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthorizationProviderResult>("azure-native:apimanagement:getAuthorizationProvider", args ?? new GetAuthorizationProviderInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthorizationProviderArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier of the authorization provider.
        /// </summary>
        [Input("authorizationProviderId", required: true)]
        public string AuthorizationProviderId { get; set; } = null!;

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

        public GetAuthorizationProviderArgs()
        {
        }
        public static new GetAuthorizationProviderArgs Empty => new GetAuthorizationProviderArgs();
    }

    public sealed class GetAuthorizationProviderInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier of the authorization provider.
        /// </summary>
        [Input("authorizationProviderId", required: true)]
        public Input<string> AuthorizationProviderId { get; set; } = null!;

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

        public GetAuthorizationProviderInvokeArgs()
        {
        }
        public static new GetAuthorizationProviderInvokeArgs Empty => new GetAuthorizationProviderInvokeArgs();
    }


    [OutputType]
    public sealed class GetAuthorizationProviderResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Authorization Provider name. Must be 1 to 300 characters long.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Identity provider name. Must be 1 to 300 characters long.
        /// </summary>
        public readonly string? IdentityProvider;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// OAuth2 settings
        /// </summary>
        public readonly Outputs.AuthorizationProviderOAuth2SettingsResponse? Oauth2;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAuthorizationProviderResult(
            string azureApiVersion,

            string? displayName,

            string id,

            string? identityProvider,

            string name,

            Outputs.AuthorizationProviderOAuth2SettingsResponse? oauth2,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            DisplayName = displayName;
            Id = id;
            IdentityProvider = identityProvider;
            Name = name;
            Oauth2 = oauth2;
            Type = type;
        }
    }
}
