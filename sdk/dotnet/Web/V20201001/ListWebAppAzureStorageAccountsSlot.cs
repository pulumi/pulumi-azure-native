// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20201001
{
    public static class ListWebAppAzureStorageAccountsSlot
    {
        /// <summary>
        /// Gets the Azure storage account configurations of an app.
        /// </summary>
        public static Task<ListWebAppAzureStorageAccountsSlotResult> InvokeAsync(ListWebAppAzureStorageAccountsSlotArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListWebAppAzureStorageAccountsSlotResult>("azure-native:web/v20201001:listWebAppAzureStorageAccountsSlot", args ?? new ListWebAppAzureStorageAccountsSlotArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the Azure storage account configurations of an app.
        /// </summary>
        public static Output<ListWebAppAzureStorageAccountsSlotResult> Invoke(ListWebAppAzureStorageAccountsSlotInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebAppAzureStorageAccountsSlotResult>("azure-native:web/v20201001:listWebAppAzureStorageAccountsSlot", args ?? new ListWebAppAzureStorageAccountsSlotInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the Azure storage account configurations of an app.
        /// </summary>
        public static Output<ListWebAppAzureStorageAccountsSlotResult> Invoke(ListWebAppAzureStorageAccountsSlotInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebAppAzureStorageAccountsSlotResult>("azure-native:web/v20201001:listWebAppAzureStorageAccountsSlot", args ?? new ListWebAppAzureStorageAccountsSlotInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListWebAppAzureStorageAccountsSlotArgs : global::Pulumi.InvokeArgs
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

        /// <summary>
        /// Name of the deployment slot. If a slot is not specified, the API will update the Azure storage account configurations for the production slot.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        public ListWebAppAzureStorageAccountsSlotArgs()
        {
        }
        public static new ListWebAppAzureStorageAccountsSlotArgs Empty => new ListWebAppAzureStorageAccountsSlotArgs();
    }

    public sealed class ListWebAppAzureStorageAccountsSlotInvokeArgs : global::Pulumi.InvokeArgs
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

        /// <summary>
        /// Name of the deployment slot. If a slot is not specified, the API will update the Azure storage account configurations for the production slot.
        /// </summary>
        [Input("slot", required: true)]
        public Input<string> Slot { get; set; } = null!;

        public ListWebAppAzureStorageAccountsSlotInvokeArgs()
        {
        }
        public static new ListWebAppAzureStorageAccountsSlotInvokeArgs Empty => new ListWebAppAzureStorageAccountsSlotInvokeArgs();
    }


    [OutputType]
    public sealed class ListWebAppAzureStorageAccountsSlotResult
    {
        /// <summary>
        /// Resource Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure storage accounts.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.AzureStorageInfoValueResponse> Properties;
        /// <summary>
        /// The system metadata relating to this resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListWebAppAzureStorageAccountsSlotResult(
            string id,

            string? kind,

            string name,

            ImmutableDictionary<string, Outputs.AzureStorageInfoValueResponse> properties,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            Id = id;
            Kind = kind;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Type = type;
        }
    }
}
