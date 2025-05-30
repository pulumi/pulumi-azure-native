// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security
{
    public static class GetSecurityContact
    {
        /// <summary>
        /// Get Default Security contact configurations for the subscription
        /// 
        /// Uses Azure REST API version 2023-12-01-preview.
        /// 
        /// Other available API versions: 2017-08-01-preview, 2020-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native security [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetSecurityContactResult> InvokeAsync(GetSecurityContactArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityContactResult>("azure-native:security:getSecurityContact", args ?? new GetSecurityContactArgs(), options.WithDefaults());

        /// <summary>
        /// Get Default Security contact configurations for the subscription
        /// 
        /// Uses Azure REST API version 2023-12-01-preview.
        /// 
        /// Other available API versions: 2017-08-01-preview, 2020-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native security [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSecurityContactResult> Invoke(GetSecurityContactInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityContactResult>("azure-native:security:getSecurityContact", args ?? new GetSecurityContactInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get Default Security contact configurations for the subscription
        /// 
        /// Uses Azure REST API version 2023-12-01-preview.
        /// 
        /// Other available API versions: 2017-08-01-preview, 2020-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native security [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSecurityContactResult> Invoke(GetSecurityContactInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityContactResult>("azure-native:security:getSecurityContact", args ?? new GetSecurityContactInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityContactArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the security contact object
        /// </summary>
        [Input("securityContactName", required: true)]
        public string SecurityContactName { get; set; } = null!;

        public GetSecurityContactArgs()
        {
        }
        public static new GetSecurityContactArgs Empty => new GetSecurityContactArgs();
    }

    public sealed class GetSecurityContactInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the security contact object
        /// </summary>
        [Input("securityContactName", required: true)]
        public Input<string> SecurityContactName { get; set; } = null!;

        public GetSecurityContactInvokeArgs()
        {
        }
        public static new GetSecurityContactInvokeArgs Empty => new GetSecurityContactInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityContactResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// List of email addresses which will get notifications from Microsoft Defender for Cloud by the configurations defined in this security contact.
        /// </summary>
        public readonly string? Emails;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Indicates whether the security contact is enabled.
        /// </summary>
        public readonly bool? IsEnabled;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Defines whether to send email notifications from Microsoft Defender for Cloud to persons with specific RBAC roles on the subscription.
        /// </summary>
        public readonly Outputs.SecurityContactPropertiesResponseNotificationsByRole? NotificationsByRole;
        /// <summary>
        /// A collection of sources types which evaluate the email notification.
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.NotificationsSourceAlertResponse, Outputs.NotificationsSourceAttackPathResponse>> NotificationsSources;
        /// <summary>
        /// The security contact's phone number
        /// </summary>
        public readonly string? Phone;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSecurityContactResult(
            string azureApiVersion,

            string? emails,

            string id,

            bool? isEnabled,

            string name,

            Outputs.SecurityContactPropertiesResponseNotificationsByRole? notificationsByRole,

            ImmutableArray<Union<Outputs.NotificationsSourceAlertResponse, Outputs.NotificationsSourceAttackPathResponse>> notificationsSources,

            string? phone,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Emails = emails;
            Id = id;
            IsEnabled = isEnabled;
            Name = name;
            NotificationsByRole = notificationsByRole;
            NotificationsSources = notificationsSources;
            Phone = phone;
            Type = type;
        }
    }
}
