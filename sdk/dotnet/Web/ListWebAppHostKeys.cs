// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    public static class ListWebAppHostKeys
    {
        /// <summary>
        /// Description for Get host secrets for a function app.
        /// 
        /// Uses Azure REST API version 2022-09-01.
        /// 
        /// Other available API versions: 2020-10-01, 2023-01-01, 2023-12-01, 2024-04-01.
        /// </summary>
        public static Task<ListWebAppHostKeysResult> InvokeAsync(ListWebAppHostKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListWebAppHostKeysResult>("azure-native:web:listWebAppHostKeys", args ?? new ListWebAppHostKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Get host secrets for a function app.
        /// 
        /// Uses Azure REST API version 2022-09-01.
        /// 
        /// Other available API versions: 2020-10-01, 2023-01-01, 2023-12-01, 2024-04-01.
        /// </summary>
        public static Output<ListWebAppHostKeysResult> Invoke(ListWebAppHostKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebAppHostKeysResult>("azure-native:web:listWebAppHostKeys", args ?? new ListWebAppHostKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Get host secrets for a function app.
        /// 
        /// Uses Azure REST API version 2022-09-01.
        /// 
        /// Other available API versions: 2020-10-01, 2023-01-01, 2023-12-01, 2024-04-01.
        /// </summary>
        public static Output<ListWebAppHostKeysResult> Invoke(ListWebAppHostKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebAppHostKeysResult>("azure-native:web:listWebAppHostKeys", args ?? new ListWebAppHostKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListWebAppHostKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Site name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListWebAppHostKeysArgs()
        {
        }
        public static new ListWebAppHostKeysArgs Empty => new ListWebAppHostKeysArgs();
    }

    public sealed class ListWebAppHostKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Site name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListWebAppHostKeysInvokeArgs()
        {
        }
        public static new ListWebAppHostKeysInvokeArgs Empty => new ListWebAppHostKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListWebAppHostKeysResult
    {
        /// <summary>
        /// Host level function keys.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? FunctionKeys;
        /// <summary>
        /// Secret key.
        /// </summary>
        public readonly string? MasterKey;
        /// <summary>
        /// System keys.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? SystemKeys;

        [OutputConstructor]
        private ListWebAppHostKeysResult(
            ImmutableDictionary<string, string>? functionKeys,

            string? masterKey,

            ImmutableDictionary<string, string>? systemKeys)
        {
            FunctionKeys = functionKeys;
            MasterKey = masterKey;
            SystemKeys = systemKeys;
        }
    }
}
