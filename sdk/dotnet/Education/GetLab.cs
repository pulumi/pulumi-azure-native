// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Education
{
    public static class GetLab
    {
        /// <summary>
        /// Get the details for a specific lab associated with the provided billing account name, billing profile name, and invoice section name.
        /// 
        /// Uses Azure REST API version 2021-12-01-preview.
        /// </summary>
        public static Task<GetLabResult> InvokeAsync(GetLabArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLabResult>("azure-native:education:getLab", args ?? new GetLabArgs(), options.WithDefaults());

        /// <summary>
        /// Get the details for a specific lab associated with the provided billing account name, billing profile name, and invoice section name.
        /// 
        /// Uses Azure REST API version 2021-12-01-preview.
        /// </summary>
        public static Output<GetLabResult> Invoke(GetLabInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLabResult>("azure-native:education:getLab", args ?? new GetLabInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the details for a specific lab associated with the provided billing account name, billing profile name, and invoice section name.
        /// 
        /// Uses Azure REST API version 2021-12-01-preview.
        /// </summary>
        public static Output<GetLabResult> Invoke(GetLabInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLabResult>("azure-native:education:getLab", args ?? new GetLabInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLabArgs : global::Pulumi.InvokeArgs
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
        /// May be used to include budget information.
        /// </summary>
        [Input("includeBudget")]
        public bool? IncludeBudget { get; set; }

        /// <summary>
        /// The ID that uniquely identifies an invoice section.
        /// </summary>
        [Input("invoiceSectionName", required: true)]
        public string InvoiceSectionName { get; set; } = null!;

        public GetLabArgs()
        {
        }
        public static new GetLabArgs Empty => new GetLabArgs();
    }

    public sealed class GetLabInvokeArgs : global::Pulumi.InvokeArgs
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
        /// May be used to include budget information.
        /// </summary>
        [Input("includeBudget")]
        public Input<bool>? IncludeBudget { get; set; }

        /// <summary>
        /// The ID that uniquely identifies an invoice section.
        /// </summary>
        [Input("invoiceSectionName", required: true)]
        public Input<string> InvoiceSectionName { get; set; } = null!;

        public GetLabInvokeArgs()
        {
        }
        public static new GetLabInvokeArgs Empty => new GetLabInvokeArgs();
    }


    [OutputType]
    public sealed class GetLabResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Default monetary cap for each student in this lab
        /// </summary>
        public readonly Outputs.AmountResponse BudgetPerStudent;
        /// <summary>
        /// The type of currency being used for the value.
        /// </summary>
        public readonly string? Currency;
        /// <summary>
        /// Detail description of this lab
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Lab Display Name
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Lab creation date
        /// </summary>
        public readonly string EffectiveDate;
        /// <summary>
        /// Default expiration date for each student in this lab
        /// </summary>
        public readonly string ExpirationDate;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// invitation code for redeemable lab
        /// </summary>
        public readonly string InvitationCode;
        /// <summary>
        /// the total number of students that can be accepted to the lab.
        /// </summary>
        public readonly double MaxStudentCount;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The status of this lab
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Total budget
        /// </summary>
        public readonly Outputs.AmountResponse TotalBudget;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Amount value.
        /// </summary>
        public readonly double? Value;

        [OutputConstructor]
        private GetLabResult(
            string azureApiVersion,

            Outputs.AmountResponse budgetPerStudent,

            string? currency,

            string description,

            string displayName,

            string effectiveDate,

            string expirationDate,

            string id,

            string invitationCode,

            double maxStudentCount,

            string name,

            string status,

            Outputs.SystemDataResponse systemData,

            Outputs.AmountResponse totalBudget,

            string type,

            double? value)
        {
            AzureApiVersion = azureApiVersion;
            BudgetPerStudent = budgetPerStudent;
            Currency = currency;
            Description = description;
            DisplayName = displayName;
            EffectiveDate = effectiveDate;
            ExpirationDate = expirationDate;
            Id = id;
            InvitationCode = invitationCode;
            MaxStudentCount = maxStudentCount;
            Name = name;
            Status = status;
            SystemData = systemData;
            TotalBudget = totalBudget;
            Type = type;
            Value = value;
        }
    }
}
