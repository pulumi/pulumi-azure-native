// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SignalRService.V20240401Preview
{
    public static class ListSignalRKeys
    {
        /// <summary>
        /// Get the access keys of the resource.
        /// </summary>
        public static Task<ListSignalRKeysResult> InvokeAsync(ListSignalRKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListSignalRKeysResult>("azure-native:signalrservice/v20240401preview:listSignalRKeys", args ?? new ListSignalRKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Get the access keys of the resource.
        /// </summary>
        public static Output<ListSignalRKeysResult> Invoke(ListSignalRKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListSignalRKeysResult>("azure-native:signalrservice/v20240401preview:listSignalRKeys", args ?? new ListSignalRKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the access keys of the resource.
        /// </summary>
        public static Output<ListSignalRKeysResult> Invoke(ListSignalRKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListSignalRKeysResult>("azure-native:signalrservice/v20240401preview:listSignalRKeys", args ?? new ListSignalRKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListSignalRKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public ListSignalRKeysArgs()
        {
        }
        public static new ListSignalRKeysArgs Empty => new ListSignalRKeysArgs();
    }

    public sealed class ListSignalRKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public ListSignalRKeysInvokeArgs()
        {
        }
        public static new ListSignalRKeysInvokeArgs Empty => new ListSignalRKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListSignalRKeysResult
    {
        /// <summary>
        /// Connection string constructed via the primaryKey
        /// </summary>
        public readonly string? PrimaryConnectionString;
        /// <summary>
        /// The primary access key.
        /// </summary>
        public readonly string? PrimaryKey;
        /// <summary>
        /// Connection string constructed via the secondaryKey
        /// </summary>
        public readonly string? SecondaryConnectionString;
        /// <summary>
        /// The secondary access key.
        /// </summary>
        public readonly string? SecondaryKey;

        [OutputConstructor]
        private ListSignalRKeysResult(
            string? primaryConnectionString,

            string? primaryKey,

            string? secondaryConnectionString,

            string? secondaryKey)
        {
            PrimaryConnectionString = primaryConnectionString;
            PrimaryKey = primaryKey;
            SecondaryConnectionString = secondaryConnectionString;
            SecondaryKey = secondaryKey;
        }
    }
}
