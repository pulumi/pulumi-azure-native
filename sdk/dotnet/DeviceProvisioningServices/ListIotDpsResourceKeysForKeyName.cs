// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DeviceProvisioningServices
{
    public static class ListIotDpsResourceKeysForKeyName
    {
        /// <summary>
        /// List primary and secondary keys for a specific key name
        /// 
        /// Uses Azure REST API version 2023-03-01-preview.
        /// 
        /// Other available API versions: 2017-08-21-preview, 2017-11-15, 2018-01-22, 2020-01-01, 2020-03-01, 2020-09-01-preview, 2021-10-15, 2022-02-05, 2022-12-12, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native deviceprovisioningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListIotDpsResourceKeysForKeyNameResult> InvokeAsync(ListIotDpsResourceKeysForKeyNameArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListIotDpsResourceKeysForKeyNameResult>("azure-native:deviceprovisioningservices:listIotDpsResourceKeysForKeyName", args ?? new ListIotDpsResourceKeysForKeyNameArgs(), options.WithDefaults());

        /// <summary>
        /// List primary and secondary keys for a specific key name
        /// 
        /// Uses Azure REST API version 2023-03-01-preview.
        /// 
        /// Other available API versions: 2017-08-21-preview, 2017-11-15, 2018-01-22, 2020-01-01, 2020-03-01, 2020-09-01-preview, 2021-10-15, 2022-02-05, 2022-12-12, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native deviceprovisioningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListIotDpsResourceKeysForKeyNameResult> Invoke(ListIotDpsResourceKeysForKeyNameInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListIotDpsResourceKeysForKeyNameResult>("azure-native:deviceprovisioningservices:listIotDpsResourceKeysForKeyName", args ?? new ListIotDpsResourceKeysForKeyNameInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List primary and secondary keys for a specific key name
        /// 
        /// Uses Azure REST API version 2023-03-01-preview.
        /// 
        /// Other available API versions: 2017-08-21-preview, 2017-11-15, 2018-01-22, 2020-01-01, 2020-03-01, 2020-09-01-preview, 2021-10-15, 2022-02-05, 2022-12-12, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native deviceprovisioningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListIotDpsResourceKeysForKeyNameResult> Invoke(ListIotDpsResourceKeysForKeyNameInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListIotDpsResourceKeysForKeyNameResult>("azure-native:deviceprovisioningservices:listIotDpsResourceKeysForKeyName", args ?? new ListIotDpsResourceKeysForKeyNameInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListIotDpsResourceKeysForKeyNameArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Logical key name to get key-values for.
        /// </summary>
        [Input("keyName", required: true)]
        public string KeyName { get; set; } = null!;

        /// <summary>
        /// Name of the provisioning service.
        /// </summary>
        [Input("provisioningServiceName", required: true)]
        public string ProvisioningServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the provisioning service.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListIotDpsResourceKeysForKeyNameArgs()
        {
        }
        public static new ListIotDpsResourceKeysForKeyNameArgs Empty => new ListIotDpsResourceKeysForKeyNameArgs();
    }

    public sealed class ListIotDpsResourceKeysForKeyNameInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Logical key name to get key-values for.
        /// </summary>
        [Input("keyName", required: true)]
        public Input<string> KeyName { get; set; } = null!;

        /// <summary>
        /// Name of the provisioning service.
        /// </summary>
        [Input("provisioningServiceName", required: true)]
        public Input<string> ProvisioningServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the provisioning service.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListIotDpsResourceKeysForKeyNameInvokeArgs()
        {
        }
        public static new ListIotDpsResourceKeysForKeyNameInvokeArgs Empty => new ListIotDpsResourceKeysForKeyNameInvokeArgs();
    }


    [OutputType]
    public sealed class ListIotDpsResourceKeysForKeyNameResult
    {
        /// <summary>
        /// Name of the key.
        /// </summary>
        public readonly string KeyName;
        /// <summary>
        /// Primary SAS key value.
        /// </summary>
        public readonly string? PrimaryKey;
        /// <summary>
        /// Rights that this key has.
        /// </summary>
        public readonly string Rights;
        /// <summary>
        /// Secondary SAS key value.
        /// </summary>
        public readonly string? SecondaryKey;

        [OutputConstructor]
        private ListIotDpsResourceKeysForKeyNameResult(
            string keyName,

            string? primaryKey,

            string rights,

            string? secondaryKey)
        {
            KeyName = keyName;
            PrimaryKey = primaryKey;
            Rights = rights;
            SecondaryKey = secondaryKey;
        }
    }
}
