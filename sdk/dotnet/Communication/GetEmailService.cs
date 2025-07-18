// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Communication
{
    public static class GetEmailService
    {
        /// <summary>
        /// Get the EmailService and its properties.
        /// 
        /// Uses Azure REST API version 2023-06-01-preview.
        /// 
        /// Other available API versions: 2023-03-31, 2023-04-01, 2023-04-01-preview, 2024-09-01-preview, 2025-05-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native communication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetEmailServiceResult> InvokeAsync(GetEmailServiceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEmailServiceResult>("azure-native:communication:getEmailService", args ?? new GetEmailServiceArgs(), options.WithDefaults());

        /// <summary>
        /// Get the EmailService and its properties.
        /// 
        /// Uses Azure REST API version 2023-06-01-preview.
        /// 
        /// Other available API versions: 2023-03-31, 2023-04-01, 2023-04-01-preview, 2024-09-01-preview, 2025-05-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native communication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetEmailServiceResult> Invoke(GetEmailServiceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEmailServiceResult>("azure-native:communication:getEmailService", args ?? new GetEmailServiceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the EmailService and its properties.
        /// 
        /// Uses Azure REST API version 2023-06-01-preview.
        /// 
        /// Other available API versions: 2023-03-31, 2023-04-01, 2023-04-01-preview, 2024-09-01-preview, 2025-05-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native communication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetEmailServiceResult> Invoke(GetEmailServiceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEmailServiceResult>("azure-native:communication:getEmailService", args ?? new GetEmailServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEmailServiceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the EmailService resource.
        /// </summary>
        [Input("emailServiceName", required: true)]
        public string EmailServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetEmailServiceArgs()
        {
        }
        public static new GetEmailServiceArgs Empty => new GetEmailServiceArgs();
    }

    public sealed class GetEmailServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the EmailService resource.
        /// </summary>
        [Input("emailServiceName", required: true)]
        public Input<string> EmailServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetEmailServiceInvokeArgs()
        {
        }
        public static new GetEmailServiceInvokeArgs Empty => new GetEmailServiceInvokeArgs();
    }


    [OutputType]
    public sealed class GetEmailServiceResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The location where the email service stores its data at rest.
        /// </summary>
        public readonly string DataLocation;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetEmailServiceResult(
            string azureApiVersion,

            string dataLocation,

            string id,

            string location,

            string name,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            DataLocation = dataLocation;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
