// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NotificationHubs
{
    public static class GetNotificationHubPnsCredentials
    {
        /// <summary>
        /// Lists the PNS Credentials associated with a notification hub.
        /// 
        /// Uses Azure REST API version 2023-10-01-preview.
        /// 
        /// Other available API versions: 2023-01-01-preview, 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native notificationhubs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetNotificationHubPnsCredentialsResult> InvokeAsync(GetNotificationHubPnsCredentialsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNotificationHubPnsCredentialsResult>("azure-native:notificationhubs:getNotificationHubPnsCredentials", args ?? new GetNotificationHubPnsCredentialsArgs(), options.WithDefaults());

        /// <summary>
        /// Lists the PNS Credentials associated with a notification hub.
        /// 
        /// Uses Azure REST API version 2023-10-01-preview.
        /// 
        /// Other available API versions: 2023-01-01-preview, 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native notificationhubs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetNotificationHubPnsCredentialsResult> Invoke(GetNotificationHubPnsCredentialsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNotificationHubPnsCredentialsResult>("azure-native:notificationhubs:getNotificationHubPnsCredentials", args ?? new GetNotificationHubPnsCredentialsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Lists the PNS Credentials associated with a notification hub.
        /// 
        /// Uses Azure REST API version 2023-10-01-preview.
        /// 
        /// Other available API versions: 2023-01-01-preview, 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native notificationhubs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetNotificationHubPnsCredentialsResult> Invoke(GetNotificationHubPnsCredentialsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNotificationHubPnsCredentialsResult>("azure-native:notificationhubs:getNotificationHubPnsCredentials", args ?? new GetNotificationHubPnsCredentialsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNotificationHubPnsCredentialsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// Notification Hub name
        /// </summary>
        [Input("notificationHubName", required: true)]
        public string NotificationHubName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNotificationHubPnsCredentialsArgs()
        {
        }
        public static new GetNotificationHubPnsCredentialsArgs Empty => new GetNotificationHubPnsCredentialsArgs();
    }

    public sealed class GetNotificationHubPnsCredentialsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Notification Hub name
        /// </summary>
        [Input("notificationHubName", required: true)]
        public Input<string> NotificationHubName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetNotificationHubPnsCredentialsInvokeArgs()
        {
        }
        public static new GetNotificationHubPnsCredentialsInvokeArgs Empty => new GetNotificationHubPnsCredentialsInvokeArgs();
    }


    [OutputType]
    public sealed class GetNotificationHubPnsCredentialsResult
    {
        /// <summary>
        /// Description of a NotificationHub AdmCredential.
        /// </summary>
        public readonly Outputs.AdmCredentialResponse? AdmCredential;
        /// <summary>
        /// Description of a NotificationHub ApnsCredential.
        /// </summary>
        public readonly Outputs.ApnsCredentialResponse? ApnsCredential;
        /// <summary>
        /// Description of a NotificationHub BaiduCredential.
        /// </summary>
        public readonly Outputs.BaiduCredentialResponse? BaiduCredential;
        /// <summary>
        /// Description of a NotificationHub BrowserCredential.
        /// </summary>
        public readonly Outputs.BrowserCredentialResponse? BrowserCredential;
        /// <summary>
        /// Description of a NotificationHub FcmV1Credential.
        /// </summary>
        public readonly Outputs.FcmV1CredentialResponse? FcmV1Credential;
        /// <summary>
        /// Description of a NotificationHub GcmCredential.
        /// </summary>
        public readonly Outputs.GcmCredentialResponse? GcmCredential;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Deprecated - only for compatibility.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Description of a NotificationHub MpnsCredential.
        /// </summary>
        public readonly Outputs.MpnsCredentialResponse? MpnsCredential;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Deprecated - only for compatibility.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Description of a NotificationHub WnsCredential.
        /// </summary>
        public readonly Outputs.WnsCredentialResponse? WnsCredential;
        /// <summary>
        /// Description of a NotificationHub XiaomiCredential.
        /// </summary>
        public readonly Outputs.XiaomiCredentialResponse? XiaomiCredential;

        [OutputConstructor]
        private GetNotificationHubPnsCredentialsResult(
            Outputs.AdmCredentialResponse? admCredential,

            Outputs.ApnsCredentialResponse? apnsCredential,

            Outputs.BaiduCredentialResponse? baiduCredential,

            Outputs.BrowserCredentialResponse? browserCredential,

            Outputs.FcmV1CredentialResponse? fcmV1Credential,

            Outputs.GcmCredentialResponse? gcmCredential,

            string id,

            string? location,

            Outputs.MpnsCredentialResponse? mpnsCredential,

            string name,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.WnsCredentialResponse? wnsCredential,

            Outputs.XiaomiCredentialResponse? xiaomiCredential)
        {
            AdmCredential = admCredential;
            ApnsCredential = apnsCredential;
            BaiduCredential = baiduCredential;
            BrowserCredential = browserCredential;
            FcmV1Credential = fcmV1Credential;
            GcmCredential = gcmCredential;
            Id = id;
            Location = location;
            MpnsCredential = mpnsCredential;
            Name = name;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            WnsCredential = wnsCredential;
            XiaomiCredential = xiaomiCredential;
        }
    }
}
