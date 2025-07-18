// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NotificationHubs
{
    public static class ListNamespaceKeys
    {
        /// <summary>
        /// Gets the Primary and Secondary ConnectionStrings to the namespace.
        /// 
        /// Uses Azure REST API version 2023-10-01-preview.
        /// 
        /// Other available API versions: 2023-01-01-preview, 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native notificationhubs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListNamespaceKeysResult> InvokeAsync(ListNamespaceKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListNamespaceKeysResult>("azure-native:notificationhubs:listNamespaceKeys", args ?? new ListNamespaceKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the Primary and Secondary ConnectionStrings to the namespace.
        /// 
        /// Uses Azure REST API version 2023-10-01-preview.
        /// 
        /// Other available API versions: 2023-01-01-preview, 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native notificationhubs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListNamespaceKeysResult> Invoke(ListNamespaceKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListNamespaceKeysResult>("azure-native:notificationhubs:listNamespaceKeys", args ?? new ListNamespaceKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the Primary and Secondary ConnectionStrings to the namespace.
        /// 
        /// Uses Azure REST API version 2023-10-01-preview.
        /// 
        /// Other available API versions: 2023-01-01-preview, 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native notificationhubs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListNamespaceKeysResult> Invoke(ListNamespaceKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListNamespaceKeysResult>("azure-native:notificationhubs:listNamespaceKeys", args ?? new ListNamespaceKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListNamespaceKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Authorization Rule Name
        /// </summary>
        [Input("authorizationRuleName", required: true)]
        public string AuthorizationRuleName { get; set; } = null!;

        /// <summary>
        /// Namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListNamespaceKeysArgs()
        {
        }
        public static new ListNamespaceKeysArgs Empty => new ListNamespaceKeysArgs();
    }

    public sealed class ListNamespaceKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Authorization Rule Name
        /// </summary>
        [Input("authorizationRuleName", required: true)]
        public Input<string> AuthorizationRuleName { get; set; } = null!;

        /// <summary>
        /// Namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListNamespaceKeysInvokeArgs()
        {
        }
        public static new ListNamespaceKeysInvokeArgs Empty => new ListNamespaceKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListNamespaceKeysResult
    {
        /// <summary>
        /// Gets or sets keyName of the created AuthorizationRule
        /// </summary>
        public readonly string KeyName;
        /// <summary>
        /// Gets or sets primaryConnectionString of the AuthorizationRule.
        /// </summary>
        public readonly string PrimaryConnectionString;
        /// <summary>
        /// Gets or sets primaryKey of the created AuthorizationRule.
        /// </summary>
        public readonly string PrimaryKey;
        /// <summary>
        /// Gets or sets secondaryConnectionString of the created
        /// AuthorizationRule
        /// </summary>
        public readonly string SecondaryConnectionString;
        /// <summary>
        /// Gets or sets secondaryKey of the created AuthorizationRule
        /// </summary>
        public readonly string SecondaryKey;

        [OutputConstructor]
        private ListNamespaceKeysResult(
            string keyName,

            string primaryConnectionString,

            string primaryKey,

            string secondaryConnectionString,

            string secondaryKey)
        {
            KeyName = keyName;
            PrimaryConnectionString = primaryConnectionString;
            PrimaryKey = primaryKey;
            SecondaryConnectionString = secondaryConnectionString;
            SecondaryKey = secondaryKey;
        }
    }
}
