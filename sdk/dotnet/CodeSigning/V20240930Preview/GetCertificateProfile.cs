// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CodeSigning.V20240930Preview
{
    public static class GetCertificateProfile
    {
        /// <summary>
        /// Get details of a certificate profile.
        /// </summary>
        public static Task<GetCertificateProfileResult> InvokeAsync(GetCertificateProfileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCertificateProfileResult>("azure-native:codesigning/v20240930preview:getCertificateProfile", args ?? new GetCertificateProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Get details of a certificate profile.
        /// </summary>
        public static Output<GetCertificateProfileResult> Invoke(GetCertificateProfileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCertificateProfileResult>("azure-native:codesigning/v20240930preview:getCertificateProfile", args ?? new GetCertificateProfileInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get details of a certificate profile.
        /// </summary>
        public static Output<GetCertificateProfileResult> Invoke(GetCertificateProfileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCertificateProfileResult>("azure-native:codesigning/v20240930preview:getCertificateProfile", args ?? new GetCertificateProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Trusted Signing account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// Certificate profile name.
        /// </summary>
        [Input("profileName", required: true)]
        public string ProfileName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetCertificateProfileArgs()
        {
        }
        public static new GetCertificateProfileArgs Empty => new GetCertificateProfileArgs();
    }

    public sealed class GetCertificateProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Trusted Signing account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Certificate profile name.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetCertificateProfileInvokeArgs()
        {
        }
        public static new GetCertificateProfileInvokeArgs Empty => new GetCertificateProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetCertificateProfileResult
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Identity validation id used for the certificate subject name.
        /// </summary>
        public readonly string IdentityValidationId;
        /// <summary>
        /// Whether to include L in the certificate subject name. Applicable only for private trust, private trust ci profile types
        /// </summary>
        public readonly bool? IncludeCity;
        /// <summary>
        /// Whether to include C in the certificate subject name. Applicable only for private trust, private trust ci profile types
        /// </summary>
        public readonly bool? IncludeCountry;
        /// <summary>
        /// Whether to include PC in the certificate subject name.
        /// </summary>
        public readonly bool? IncludePostalCode;
        /// <summary>
        /// Whether to include S in the certificate subject name. Applicable only for private trust, private trust ci profile types
        /// </summary>
        public readonly bool? IncludeState;
        /// <summary>
        /// Whether to include STREET in the certificate subject name.
        /// </summary>
        public readonly bool? IncludeStreetAddress;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Profile type of the certificate.
        /// </summary>
        public readonly string ProfileType;
        /// <summary>
        /// Status of the current operation on certificate profile.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Status of the certificate profile.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCertificateProfileResult(
            string id,

            string identityValidationId,

            bool? includeCity,

            bool? includeCountry,

            bool? includePostalCode,

            bool? includeState,

            bool? includeStreetAddress,

            string name,

            string profileType,

            string provisioningState,

            string status,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            Id = id;
            IdentityValidationId = identityValidationId;
            IncludeCity = includeCity;
            IncludeCountry = includeCountry;
            IncludePostalCode = includePostalCode;
            IncludeState = includeState;
            IncludeStreetAddress = includeStreetAddress;
            Name = name;
            ProfileType = profileType;
            ProvisioningState = provisioningState;
            Status = status;
            SystemData = systemData;
            Type = type;
        }
    }
}
