// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Billing
{
    public static class GetInvoiceSection
    {
        /// <summary>
        /// Gets an invoice section by its ID. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// </summary>
        public static Task<GetInvoiceSectionResult> InvokeAsync(GetInvoiceSectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInvoiceSectionResult>("azure-native:billing:getInvoiceSection", args ?? new GetInvoiceSectionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an invoice section by its ID. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// </summary>
        public static Output<GetInvoiceSectionResult> Invoke(GetInvoiceSectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInvoiceSectionResult>("azure-native:billing:getInvoiceSection", args ?? new GetInvoiceSectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an invoice section by its ID. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// </summary>
        public static Output<GetInvoiceSectionResult> Invoke(GetInvoiceSectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetInvoiceSectionResult>("azure-native:billing:getInvoiceSection", args ?? new GetInvoiceSectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInvoiceSectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID that uniquely identifies a billing account.
        /// </summary>
        [Input("billingAccountName", required: true)]
        public string BillingAccountName { get; set; } = null!;

        /// <summary>
        /// The ID that uniquely identifies a billing profile.
        /// </summary>
        [Input("billingProfileName", required: true)]
        public string BillingProfileName { get; set; } = null!;

        /// <summary>
        /// The ID that uniquely identifies an invoice section.
        /// </summary>
        [Input("invoiceSectionName", required: true)]
        public string InvoiceSectionName { get; set; } = null!;

        public GetInvoiceSectionArgs()
        {
        }
        public static new GetInvoiceSectionArgs Empty => new GetInvoiceSectionArgs();
    }

    public sealed class GetInvoiceSectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID that uniquely identifies a billing account.
        /// </summary>
        [Input("billingAccountName", required: true)]
        public Input<string> BillingAccountName { get; set; } = null!;

        /// <summary>
        /// The ID that uniquely identifies a billing profile.
        /// </summary>
        [Input("billingProfileName", required: true)]
        public Input<string> BillingProfileName { get; set; } = null!;

        /// <summary>
        /// The ID that uniquely identifies an invoice section.
        /// </summary>
        [Input("invoiceSectionName", required: true)]
        public Input<string> InvoiceSectionName { get; set; } = null!;

        public GetInvoiceSectionInvokeArgs()
        {
        }
        public static new GetInvoiceSectionInvokeArgs Empty => new GetInvoiceSectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetInvoiceSectionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An invoice section.
        /// </summary>
        public readonly Outputs.InvoiceSectionPropertiesResponse Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Dictionary of metadata associated with the resource. It may not be populated for all resource types. Maximum key/value length supported of 256 characters. Keys/value should not empty value nor null. Keys can not contain &lt; &gt; % &amp; \ ? /
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetInvoiceSectionResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.InvoiceSectionPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
