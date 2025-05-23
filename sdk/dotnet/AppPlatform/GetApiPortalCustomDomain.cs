// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform
{
    public static class GetApiPortalCustomDomain
    {
        /// <summary>
        /// Get the API portal custom domain.
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetApiPortalCustomDomainResult> InvokeAsync(GetApiPortalCustomDomainArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApiPortalCustomDomainResult>("azure-native:appplatform:getApiPortalCustomDomain", args ?? new GetApiPortalCustomDomainArgs(), options.WithDefaults());

        /// <summary>
        /// Get the API portal custom domain.
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetApiPortalCustomDomainResult> Invoke(GetApiPortalCustomDomainInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApiPortalCustomDomainResult>("azure-native:appplatform:getApiPortalCustomDomain", args ?? new GetApiPortalCustomDomainInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the API portal custom domain.
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetApiPortalCustomDomainResult> Invoke(GetApiPortalCustomDomainInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApiPortalCustomDomainResult>("azure-native:appplatform:getApiPortalCustomDomain", args ?? new GetApiPortalCustomDomainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApiPortalCustomDomainArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of API portal.
        /// </summary>
        [Input("apiPortalName", required: true)]
        public string ApiPortalName { get; set; } = null!;

        /// <summary>
        /// The name of the API portal custom domain.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Service resource.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetApiPortalCustomDomainArgs()
        {
        }
        public static new GetApiPortalCustomDomainArgs Empty => new GetApiPortalCustomDomainArgs();
    }

    public sealed class GetApiPortalCustomDomainInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of API portal.
        /// </summary>
        [Input("apiPortalName", required: true)]
        public Input<string> ApiPortalName { get; set; } = null!;

        /// <summary>
        /// The name of the API portal custom domain.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Service resource.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetApiPortalCustomDomainInvokeArgs()
        {
        }
        public static new GetApiPortalCustomDomainInvokeArgs Empty => new GetApiPortalCustomDomainInvokeArgs();
    }


    [OutputType]
    public sealed class GetApiPortalCustomDomainResult
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
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties of custom domain for API portal
        /// </summary>
        public readonly Outputs.ApiPortalCustomDomainPropertiesResponse Properties;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApiPortalCustomDomainResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.ApiPortalCustomDomainPropertiesResponse properties,

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
