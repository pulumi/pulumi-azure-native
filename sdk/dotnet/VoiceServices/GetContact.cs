// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VoiceServices
{
    public static class GetContact
    {
        /// <summary>
        /// Get a Contact
        /// 
        /// Uses Azure REST API version 2022-12-01-preview.
        /// </summary>
        public static Task<GetContactResult> InvokeAsync(GetContactArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContactResult>("azure-native:voiceservices:getContact", args ?? new GetContactArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Contact
        /// 
        /// Uses Azure REST API version 2022-12-01-preview.
        /// </summary>
        public static Output<GetContactResult> Invoke(GetContactInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContactResult>("azure-native:voiceservices:getContact", args ?? new GetContactInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Contact
        /// 
        /// Uses Azure REST API version 2022-12-01-preview.
        /// </summary>
        public static Output<GetContactResult> Invoke(GetContactInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetContactResult>("azure-native:voiceservices:getContact", args ?? new GetContactInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContactArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for this deployment
        /// </summary>
        [Input("communicationsGatewayName", required: true)]
        public string CommunicationsGatewayName { get; set; } = null!;

        /// <summary>
        /// Unique identifier for this contact
        /// </summary>
        [Input("contactName", required: true)]
        public string ContactName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetContactArgs()
        {
        }
        public static new GetContactArgs Empty => new GetContactArgs();
    }

    public sealed class GetContactInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for this deployment
        /// </summary>
        [Input("communicationsGatewayName", required: true)]
        public Input<string> CommunicationsGatewayName { get; set; } = null!;

        /// <summary>
        /// Unique identifier for this contact
        /// </summary>
        [Input("contactName", required: true)]
        public Input<string> ContactName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetContactInvokeArgs()
        {
        }
        public static new GetContactInvokeArgs Empty => new GetContactInvokeArgs();
    }


    [OutputType]
    public sealed class GetContactResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Full name of contact
        /// </summary>
        public readonly string ContactName;
        /// <summary>
        /// Email address of contact
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
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
        /// Telephone number of contact
        /// </summary>
        public readonly string PhoneNumber;
        /// <summary>
        /// Resource provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Job title of contact
        /// </summary>
        public readonly string Role;
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
        private GetContactResult(
            string azureApiVersion,

            string contactName,

            string email,

            string id,

            string location,

            string name,

            string phoneNumber,

            string provisioningState,

            string role,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            ContactName = contactName;
            Email = email;
            Id = id;
            Location = location;
            Name = name;
            PhoneNumber = phoneNumber;
            ProvisioningState = provisioningState;
            Role = role;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
