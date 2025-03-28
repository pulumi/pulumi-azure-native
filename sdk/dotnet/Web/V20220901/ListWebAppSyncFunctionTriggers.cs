// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20220901
{
    public static class ListWebAppSyncFunctionTriggers
    {
        /// <summary>
        /// Description for This is to allow calling via powershell and ARM template.
        /// </summary>
        public static Task<ListWebAppSyncFunctionTriggersResult> InvokeAsync(ListWebAppSyncFunctionTriggersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListWebAppSyncFunctionTriggersResult>("azure-native:web/v20220901:listWebAppSyncFunctionTriggers", args ?? new ListWebAppSyncFunctionTriggersArgs(), options.WithDefaults());

        /// <summary>
        /// Description for This is to allow calling via powershell and ARM template.
        /// </summary>
        public static Output<ListWebAppSyncFunctionTriggersResult> Invoke(ListWebAppSyncFunctionTriggersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebAppSyncFunctionTriggersResult>("azure-native:web/v20220901:listWebAppSyncFunctionTriggers", args ?? new ListWebAppSyncFunctionTriggersInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Description for This is to allow calling via powershell and ARM template.
        /// </summary>
        public static Output<ListWebAppSyncFunctionTriggersResult> Invoke(ListWebAppSyncFunctionTriggersInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebAppSyncFunctionTriggersResult>("azure-native:web/v20220901:listWebAppSyncFunctionTriggers", args ?? new ListWebAppSyncFunctionTriggersInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListWebAppSyncFunctionTriggersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListWebAppSyncFunctionTriggersArgs()
        {
        }
        public static new ListWebAppSyncFunctionTriggersArgs Empty => new ListWebAppSyncFunctionTriggersArgs();
    }

    public sealed class ListWebAppSyncFunctionTriggersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListWebAppSyncFunctionTriggersInvokeArgs()
        {
        }
        public static new ListWebAppSyncFunctionTriggersInvokeArgs Empty => new ListWebAppSyncFunctionTriggersInvokeArgs();
    }


    [OutputType]
    public sealed class ListWebAppSyncFunctionTriggersResult
    {
        /// <summary>
        /// Secret key.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// Trigger URL.
        /// </summary>
        public readonly string? TriggerUrl;

        [OutputConstructor]
        private ListWebAppSyncFunctionTriggersResult(
            string? key,

            string? triggerUrl)
        {
            Key = key;
            TriggerUrl = triggerUrl;
        }
    }
}
