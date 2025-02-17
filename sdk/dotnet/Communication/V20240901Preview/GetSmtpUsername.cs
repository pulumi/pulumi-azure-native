// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Communication.V20240901Preview
{
    public static class GetSmtpUsername
    {
        /// <summary>
        /// Get a SmtpUsernameResource.
        /// </summary>
        public static Task<GetSmtpUsernameResult> InvokeAsync(GetSmtpUsernameArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSmtpUsernameResult>("azure-native:communication/v20240901preview:getSmtpUsername", args ?? new GetSmtpUsernameArgs(), options.WithDefaults());

        /// <summary>
        /// Get a SmtpUsernameResource.
        /// </summary>
        public static Output<GetSmtpUsernameResult> Invoke(GetSmtpUsernameInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSmtpUsernameResult>("azure-native:communication/v20240901preview:getSmtpUsername", args ?? new GetSmtpUsernameInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a SmtpUsernameResource.
        /// </summary>
        public static Output<GetSmtpUsernameResult> Invoke(GetSmtpUsernameInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSmtpUsernameResult>("azure-native:communication/v20240901preview:getSmtpUsername", args ?? new GetSmtpUsernameInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSmtpUsernameArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the CommunicationService resource.
        /// </summary>
        [Input("communicationServiceName", required: true)]
        public string CommunicationServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the SmtpUsernameResource.
        /// </summary>
        [Input("smtpUsername", required: true)]
        public string SmtpUsername { get; set; } = null!;

        public GetSmtpUsernameArgs()
        {
        }
        public static new GetSmtpUsernameArgs Empty => new GetSmtpUsernameArgs();
    }

    public sealed class GetSmtpUsernameInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the CommunicationService resource.
        /// </summary>
        [Input("communicationServiceName", required: true)]
        public Input<string> CommunicationServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the SmtpUsernameResource.
        /// </summary>
        [Input("smtpUsername", required: true)]
        public Input<string> SmtpUsername { get; set; } = null!;

        public GetSmtpUsernameInvokeArgs()
        {
        }
        public static new GetSmtpUsernameInvokeArgs Empty => new GetSmtpUsernameInvokeArgs();
    }


    [OutputType]
    public sealed class GetSmtpUsernameResult
    {
        /// <summary>
        /// The application Id for the linked Entra Application.
        /// </summary>
        public readonly string EntraApplicationId;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The tenant of the linked Entra Application.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The SMTP username. Could be free form or in the email address format.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetSmtpUsernameResult(
            string entraApplicationId,

            string id,

            string name,

            Outputs.SystemDataResponse systemData,

            string tenantId,

            string type,

            string username)
        {
            EntraApplicationId = entraApplicationId;
            Id = id;
            Name = name;
            SystemData = systemData;
            TenantId = tenantId;
            Type = type;
            Username = username;
        }
    }
}
