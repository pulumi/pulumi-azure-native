// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevHub
{
    public static class GetIacProfile
    {
        /// <summary>
        /// Resource representation of a IacProfile.
        /// 
        /// Uses Azure REST API version 2024-05-01-preview.
        /// 
        /// Other available API versions: 2024-08-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devhub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetIacProfileResult> InvokeAsync(GetIacProfileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIacProfileResult>("azure-native:devhub:getIacProfile", args ?? new GetIacProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Resource representation of a IacProfile.
        /// 
        /// Uses Azure REST API version 2024-05-01-preview.
        /// 
        /// Other available API versions: 2024-08-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devhub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetIacProfileResult> Invoke(GetIacProfileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIacProfileResult>("azure-native:devhub:getIacProfile", args ?? new GetIacProfileInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource representation of a IacProfile.
        /// 
        /// Uses Azure REST API version 2024-05-01-preview.
        /// 
        /// Other available API versions: 2024-08-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devhub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetIacProfileResult> Invoke(GetIacProfileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIacProfileResult>("azure-native:devhub:getIacProfile", args ?? new GetIacProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIacProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the IacProfile.
        /// </summary>
        [Input("iacProfileName", required: true)]
        public string IacProfileName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetIacProfileArgs()
        {
        }
        public static new GetIacProfileArgs Empty => new GetIacProfileArgs();
    }

    public sealed class GetIacProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the IacProfile.
        /// </summary>
        [Input("iacProfileName", required: true)]
        public Input<string> IacProfileName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetIacProfileInvokeArgs()
        {
        }
        public static new GetIacProfileInvokeArgs Empty => new GetIacProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetIacProfileResult
    {
        /// <summary>
        /// Determines the authorization status of requests.
        /// </summary>
        public readonly string AuthStatus;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Repository Branch Name
        /// </summary>
        public readonly string? BranchName;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
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
        /// The status of the Pull Request submitted against the users repository.
        /// </summary>
        public readonly string PrStatus;
        /// <summary>
        /// The number associated with the submitted pull request.
        /// </summary>
        public readonly int PullNumber;
        /// <summary>
        /// Repository Main Branch
        /// </summary>
        public readonly string? RepositoryMainBranch;
        /// <summary>
        /// Repository Name
        /// </summary>
        public readonly string? RepositoryName;
        /// <summary>
        /// Repository Owner
        /// </summary>
        public readonly string? RepositoryOwner;
        public readonly ImmutableArray<Outputs.StagePropertiesResponse> Stages;
        /// <summary>
        /// Terraform Storage Account Name
        /// </summary>
        public readonly string? StorageAccountName;
        /// <summary>
        /// Terraform Storage Account Resource Group
        /// </summary>
        public readonly string? StorageAccountResourceGroup;
        /// <summary>
        /// Terraform Storage Account Subscription
        /// </summary>
        public readonly string? StorageAccountSubscription;
        /// <summary>
        /// Terraform Container Name
        /// </summary>
        public readonly string? StorageContainerName;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        public readonly ImmutableArray<Outputs.IacTemplatePropertiesResponse> Templates;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetIacProfileResult(
            string authStatus,

            string azureApiVersion,

            string? branchName,

            string etag,

            string id,

            string location,

            string name,

            string prStatus,

            int pullNumber,

            string? repositoryMainBranch,

            string? repositoryName,

            string? repositoryOwner,

            ImmutableArray<Outputs.StagePropertiesResponse> stages,

            string? storageAccountName,

            string? storageAccountResourceGroup,

            string? storageAccountSubscription,

            string? storageContainerName,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            ImmutableArray<Outputs.IacTemplatePropertiesResponse> templates,

            string type)
        {
            AuthStatus = authStatus;
            AzureApiVersion = azureApiVersion;
            BranchName = branchName;
            Etag = etag;
            Id = id;
            Location = location;
            Name = name;
            PrStatus = prStatus;
            PullNumber = pullNumber;
            RepositoryMainBranch = repositoryMainBranch;
            RepositoryName = repositoryName;
            RepositoryOwner = repositoryOwner;
            Stages = stages;
            StorageAccountName = storageAccountName;
            StorageAccountResourceGroup = storageAccountResourceGroup;
            StorageAccountSubscription = storageAccountSubscription;
            StorageContainerName = storageContainerName;
            SystemData = systemData;
            Tags = tags;
            Templates = templates;
            Type = type;
        }
    }
}
