// *** WARNING: this file was generated by pulumi. ***
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
        /// Uses Azure REST API version 2020-01-01-preview.
        /// 
        /// Other available API versions: 2017-08-01-preview, 2023-12-01-preview.
        /// </summary>
        public static Task<GetSecurityContactResult> InvokeAsync(GetSecurityContactArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityContactResult>("azure-native:security:getSecurityContact", args ?? new GetSecurityContactArgs(), options.WithDefaults());

        /// <summary>
        /// Get Default Security contact configurations for the subscription
        /// 
        /// Uses Azure REST API version 2020-01-01-preview.
        /// 
        /// Other available API versions: 2017-08-01-preview, 2023-12-01-preview.
        /// </summary>
        public static Output<GetSecurityContactResult> Invoke(GetSecurityContactInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityContactResult>("azure-native:security:getSecurityContact", args ?? new GetSecurityContactInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get Default Security contact configurations for the subscription
        /// 
        /// Uses Azure REST API version 2020-01-01-preview.
        /// 
        /// Other available API versions: 2017-08-01-preview, 2023-12-01-preview.
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
        /// Defines whether to send email notifications about new security alerts
        /// </summary>
        public readonly Outputs.SecurityContactPropertiesResponseAlertNotifications? AlertNotifications;
        /// <summary>
        /// List of email addresses which will get notifications from Microsoft Defender for Cloud by the configurations defined in this security contact.
        /// </summary>
        public readonly string? Emails;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Defines whether to send email notifications from Microsoft Defender for Cloud to persons with specific RBAC roles on the subscription.
        /// </summary>
        public readonly Outputs.SecurityContactPropertiesResponseNotificationsByRole? NotificationsByRole;
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
            Outputs.SecurityContactPropertiesResponseAlertNotifications? alertNotifications,

            string? emails,

            string id,

            string name,

            Outputs.SecurityContactPropertiesResponseNotificationsByRole? notificationsByRole,

            string? phone,

            string type)
        {
            AlertNotifications = alertNotifications;
            Emails = emails;
            Id = id;
            Name = name;
            NotificationsByRole = notificationsByRole;
            Phone = phone;
            Type = type;
        }
    }
}
