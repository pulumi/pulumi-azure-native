// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MixedReality.V20210101
{
    public static class ListRemoteRenderingAccountKeys
    {
        /// <summary>
        /// 
        /// 
        /// &gt; [!NOTE]
        /// &gt;
        /// &gt; **Mixed Reality retirement**
        /// &gt;
        /// &gt; The Mixed Reality service is now deprecated and will be retired. 
        /// 
        ///  List Both of the 2 Keys of a Remote Rendering Account
        /// </summary>
        public static Task<ListRemoteRenderingAccountKeysResult> InvokeAsync(ListRemoteRenderingAccountKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListRemoteRenderingAccountKeysResult>("azure-native:mixedreality/v20210101:listRemoteRenderingAccountKeys", args ?? new ListRemoteRenderingAccountKeysArgs(), options.WithDefaults());

        /// <summary>
        /// 
        /// 
        /// &gt; [!NOTE]
        /// &gt;
        /// &gt; **Mixed Reality retirement**
        /// &gt;
        /// &gt; The Mixed Reality service is now deprecated and will be retired. 
        /// 
        ///  List Both of the 2 Keys of a Remote Rendering Account
        /// </summary>
        public static Output<ListRemoteRenderingAccountKeysResult> Invoke(ListRemoteRenderingAccountKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListRemoteRenderingAccountKeysResult>("azure-native:mixedreality/v20210101:listRemoteRenderingAccountKeys", args ?? new ListRemoteRenderingAccountKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// 
        /// 
        /// &gt; [!NOTE]
        /// &gt;
        /// &gt; **Mixed Reality retirement**
        /// &gt;
        /// &gt; The Mixed Reality service is now deprecated and will be retired. 
        /// 
        ///  List Both of the 2 Keys of a Remote Rendering Account
        /// </summary>
        public static Output<ListRemoteRenderingAccountKeysResult> Invoke(ListRemoteRenderingAccountKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListRemoteRenderingAccountKeysResult>("azure-native:mixedreality/v20210101:listRemoteRenderingAccountKeys", args ?? new ListRemoteRenderingAccountKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListRemoteRenderingAccountKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of an Mixed Reality Account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListRemoteRenderingAccountKeysArgs()
        {
        }
        public static new ListRemoteRenderingAccountKeysArgs Empty => new ListRemoteRenderingAccountKeysArgs();
    }

    public sealed class ListRemoteRenderingAccountKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of an Mixed Reality Account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListRemoteRenderingAccountKeysInvokeArgs()
        {
        }
        public static new ListRemoteRenderingAccountKeysInvokeArgs Empty => new ListRemoteRenderingAccountKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListRemoteRenderingAccountKeysResult
    {
        /// <summary>
        /// value of primary key.
        /// </summary>
        public readonly string PrimaryKey;
        /// <summary>
        /// value of secondary key.
        /// </summary>
        public readonly string SecondaryKey;

        [OutputConstructor]
        private ListRemoteRenderingAccountKeysResult(
            string primaryKey,

            string secondaryKey)
        {
            PrimaryKey = primaryKey;
            SecondaryKey = secondaryKey;
        }
    }
}
