// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20250101
{
    public static class GetManagedCertificate
    {
        /// <summary>
        /// Managed certificates used for Custom Domain bindings of Container Apps in a Managed Environment
        /// </summary>
        public static Task<GetManagedCertificateResult> InvokeAsync(GetManagedCertificateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetManagedCertificateResult>("azure-native:app/v20250101:getManagedCertificate", args ?? new GetManagedCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Managed certificates used for Custom Domain bindings of Container Apps in a Managed Environment
        /// </summary>
        public static Output<GetManagedCertificateResult> Invoke(GetManagedCertificateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagedCertificateResult>("azure-native:app/v20250101:getManagedCertificate", args ?? new GetManagedCertificateInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Managed certificates used for Custom Domain bindings of Container Apps in a Managed Environment
        /// </summary>
        public static Output<GetManagedCertificateResult> Invoke(GetManagedCertificateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagedCertificateResult>("azure-native:app/v20250101:getManagedCertificate", args ?? new GetManagedCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetManagedCertificateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Managed Environment.
        /// </summary>
        [Input("environmentName", required: true)]
        public string EnvironmentName { get; set; } = null!;

        /// <summary>
        /// Name of the Managed Certificate.
        /// </summary>
        [Input("managedCertificateName", required: true)]
        public string ManagedCertificateName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetManagedCertificateArgs()
        {
        }
        public static new GetManagedCertificateArgs Empty => new GetManagedCertificateArgs();
    }

    public sealed class GetManagedCertificateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Managed Environment.
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        /// <summary>
        /// Name of the Managed Certificate.
        /// </summary>
        [Input("managedCertificateName", required: true)]
        public Input<string> ManagedCertificateName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetManagedCertificateInvokeArgs()
        {
        }
        public static new GetManagedCertificateInvokeArgs Empty => new GetManagedCertificateInvokeArgs();
    }


    [OutputType]
    public sealed class GetManagedCertificateResult
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Certificate resource specific properties
        /// </summary>
        public readonly Outputs.ManagedCertificateResponseProperties Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetManagedCertificateResult(
            string id,

            string location,

            string name,

            Outputs.ManagedCertificateResponseProperties properties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
