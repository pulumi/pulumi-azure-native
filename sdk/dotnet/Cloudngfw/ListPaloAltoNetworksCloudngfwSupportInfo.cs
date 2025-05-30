// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw
{
    public static class ListPaloAltoNetworksCloudngfwSupportInfo
    {
        /// <summary>
        /// Support information for the service
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2024-01-19-preview, 2024-02-07-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListPaloAltoNetworksCloudngfwSupportInfoResult> InvokeAsync(ListPaloAltoNetworksCloudngfwSupportInfoArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListPaloAltoNetworksCloudngfwSupportInfoResult>("azure-native:cloudngfw:listPaloAltoNetworksCloudngfwSupportInfo", args ?? new ListPaloAltoNetworksCloudngfwSupportInfoArgs(), options.WithDefaults());

        /// <summary>
        /// Support information for the service
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2024-01-19-preview, 2024-02-07-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListPaloAltoNetworksCloudngfwSupportInfoResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListPaloAltoNetworksCloudngfwSupportInfoResult>("azure-native:cloudngfw:listPaloAltoNetworksCloudngfwSupportInfo", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Support information for the service
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2024-01-19-preview, 2024-02-07-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListPaloAltoNetworksCloudngfwSupportInfoResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListPaloAltoNetworksCloudngfwSupportInfoResult>("azure-native:cloudngfw:listPaloAltoNetworksCloudngfwSupportInfo", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListPaloAltoNetworksCloudngfwSupportInfoArgs : global::Pulumi.InvokeArgs
    {
        public ListPaloAltoNetworksCloudngfwSupportInfoArgs()
        {
        }
        public static new ListPaloAltoNetworksCloudngfwSupportInfoArgs Empty => new ListPaloAltoNetworksCloudngfwSupportInfoArgs();
    }


    [OutputType]
    public sealed class ListPaloAltoNetworksCloudngfwSupportInfoResult
    {
        /// <summary>
        /// Support account associated with given resource when association type is tenant
        /// </summary>
        public readonly string? AccountId;
        /// <summary>
        /// Support account associated with given resource when association type is billing
        /// </summary>
        public readonly string? AccountIdForBilling;
        /// <summary>
        /// account registered in Customer Support Portal
        /// </summary>
        public readonly string? AccountRegistrationStatus;
        /// <summary>
        /// Association Type
        /// </summary>
        public readonly string? AssociationType;
        /// <summary>
        /// credits purchased, unit per hour
        /// </summary>
        public readonly int? Credits;
        /// <summary>
        /// date in format yyyy-mm-dd
        /// </summary>
        public readonly string? EndDateForCredits;
        /// <summary>
        /// Product usage is in free trial period
        /// </summary>
        public readonly string? FreeTrial;
        /// <summary>
        /// Free trial credit remaining
        /// </summary>
        public readonly int? FreeTrialCreditLeft;
        /// <summary>
        /// Free trial days remaining
        /// </summary>
        public readonly int? FreeTrialDaysLeft;
        /// <summary>
        /// URL for paloaltonetworks live community
        /// </summary>
        public readonly string? HelpURL;
        /// <summary>
        /// URL for Strata Cloud Manager
        /// </summary>
        public readonly string? HubUrl;
        /// <summary>
        /// monthly credit is computed as credits * days in calendar month
        /// </summary>
        public readonly int? MonthlyCreditLeft;
        /// <summary>
        /// product Serial associated with given resource
        /// </summary>
        public readonly string? ProductSerial;
        /// <summary>
        /// product SKU associated with given resource
        /// </summary>
        public readonly string? ProductSku;
        /// <summary>
        /// URL for registering product in paloaltonetworks Customer Service Portal
        /// </summary>
        public readonly string? RegisterURL;
        /// <summary>
        /// date in format yyyy-mm-dd
        /// </summary>
        public readonly string? StartDateForCredits;
        /// <summary>
        /// URL for paloaltonetworks Customer Service Portal
        /// </summary>
        public readonly string? SupportURL;

        [OutputConstructor]
        private ListPaloAltoNetworksCloudngfwSupportInfoResult(
            string? accountId,

            string? accountIdForBilling,

            string? accountRegistrationStatus,

            string? associationType,

            int? credits,

            string? endDateForCredits,

            string? freeTrial,

            int? freeTrialCreditLeft,

            int? freeTrialDaysLeft,

            string? helpURL,

            string? hubUrl,

            int? monthlyCreditLeft,

            string? productSerial,

            string? productSku,

            string? registerURL,

            string? startDateForCredits,

            string? supportURL)
        {
            AccountId = accountId;
            AccountIdForBilling = accountIdForBilling;
            AccountRegistrationStatus = accountRegistrationStatus;
            AssociationType = associationType;
            Credits = credits;
            EndDateForCredits = endDateForCredits;
            FreeTrial = freeTrial;
            FreeTrialCreditLeft = freeTrialCreditLeft;
            FreeTrialDaysLeft = freeTrialDaysLeft;
            HelpURL = helpURL;
            HubUrl = hubUrl;
            MonthlyCreditLeft = monthlyCreditLeft;
            ProductSerial = productSerial;
            ProductSku = productSku;
            RegisterURL = registerURL;
            StartDateForCredits = startDateForCredits;
            SupportURL = supportURL;
        }
    }
}
