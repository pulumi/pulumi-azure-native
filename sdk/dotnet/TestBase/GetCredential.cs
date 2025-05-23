// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TestBase
{
    public static class GetCredential
    {
        /// <summary>
        /// Gets a test base credential Resource
        /// 
        /// Uses Azure REST API version 2023-11-01-preview.
        /// </summary>
        public static Task<GetCredentialResult> InvokeAsync(GetCredentialArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCredentialResult>("azure-native:testbase:getCredential", args ?? new GetCredentialArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a test base credential Resource
        /// 
        /// Uses Azure REST API version 2023-11-01-preview.
        /// </summary>
        public static Output<GetCredentialResult> Invoke(GetCredentialInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCredentialResult>("azure-native:testbase:getCredential", args ?? new GetCredentialInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a test base credential Resource
        /// 
        /// Uses Azure REST API version 2023-11-01-preview.
        /// </summary>
        public static Output<GetCredentialResult> Invoke(GetCredentialInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCredentialResult>("azure-native:testbase:getCredential", args ?? new GetCredentialInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCredentialArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The credential resource name.
        /// </summary>
        [Input("credentialName", required: true)]
        public string CredentialName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource name of the Test Base Account.
        /// </summary>
        [Input("testBaseAccountName", required: true)]
        public string TestBaseAccountName { get; set; } = null!;

        public GetCredentialArgs()
        {
        }
        public static new GetCredentialArgs Empty => new GetCredentialArgs();
    }

    public sealed class GetCredentialInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The credential resource name.
        /// </summary>
        [Input("credentialName", required: true)]
        public Input<string> CredentialName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource name of the Test Base Account.
        /// </summary>
        [Input("testBaseAccountName", required: true)]
        public Input<string> TestBaseAccountName { get; set; } = null!;

        public GetCredentialInvokeArgs()
        {
        }
        public static new GetCredentialInvokeArgs Empty => new GetCredentialInvokeArgs();
    }


    [OutputType]
    public sealed class GetCredentialResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Credential type.
        /// </summary>
        public readonly string CredentialType;
        /// <summary>
        /// Credential display name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCredentialResult(
            string azureApiVersion,

            string credentialType,

            string displayName,

            string id,

            string name,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            CredentialType = credentialType;
            DisplayName = displayName;
            Id = id;
            Name = name;
            SystemData = systemData;
            Type = type;
        }
    }
}
