// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataShare
{
    public static class GetShare
    {
        /// <summary>
        /// Get a share 
        /// 
        /// Uses Azure REST API version 2021-08-01.
        /// </summary>
        public static Task<GetShareResult> InvokeAsync(GetShareArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetShareResult>("azure-native:datashare:getShare", args ?? new GetShareArgs(), options.WithDefaults());

        /// <summary>
        /// Get a share 
        /// 
        /// Uses Azure REST API version 2021-08-01.
        /// </summary>
        public static Output<GetShareResult> Invoke(GetShareInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetShareResult>("azure-native:datashare:getShare", args ?? new GetShareInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a share 
        /// 
        /// Uses Azure REST API version 2021-08-01.
        /// </summary>
        public static Output<GetShareResult> Invoke(GetShareInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetShareResult>("azure-native:datashare:getShare", args ?? new GetShareInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetShareArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the share to retrieve.
        /// </summary>
        [Input("shareName", required: true)]
        public string ShareName { get; set; } = null!;

        public GetShareArgs()
        {
        }
        public static new GetShareArgs Empty => new GetShareArgs();
    }

    public sealed class GetShareInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the share to retrieve.
        /// </summary>
        [Input("shareName", required: true)]
        public Input<string> ShareName { get; set; } = null!;

        public GetShareInvokeArgs()
        {
        }
        public static new GetShareInvokeArgs Empty => new GetShareInvokeArgs();
    }


    [OutputType]
    public sealed class GetShareResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Time at which the share was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Share description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The resource id of the azure resource
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the azure resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets the provisioning state
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Share kind.
        /// </summary>
        public readonly string? ShareKind;
        /// <summary>
        /// System Data of the Azure resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Share terms.
        /// </summary>
        public readonly string? Terms;
        /// <summary>
        /// Type of the azure resource
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Email of the user who created the resource
        /// </summary>
        public readonly string UserEmail;
        /// <summary>
        /// Name of the user who created the resource
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private GetShareResult(
            string azureApiVersion,

            string createdAt,

            string? description,

            string id,

            string name,

            string provisioningState,

            string? shareKind,

            Outputs.SystemDataResponse systemData,

            string? terms,

            string type,

            string userEmail,

            string userName)
        {
            AzureApiVersion = azureApiVersion;
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            ShareKind = shareKind;
            SystemData = systemData;
            Terms = terms;
            Type = type;
            UserEmail = userEmail;
            UserName = userName;
        }
    }
}
