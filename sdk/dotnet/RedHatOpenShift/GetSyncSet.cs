// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RedHatOpenShift
{
    public static class GetSyncSet
    {
        /// <summary>
        /// The operation returns properties of a SyncSet.
        /// 
        /// Uses Azure REST API version 2023-11-22.
        /// 
        /// Other available API versions: 2022-09-04, 2023-04-01, 2023-07-01-preview, 2023-09-04. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redhatopenshift [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetSyncSetResult> InvokeAsync(GetSyncSetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSyncSetResult>("azure-native:redhatopenshift:getSyncSet", args ?? new GetSyncSetArgs(), options.WithDefaults());

        /// <summary>
        /// The operation returns properties of a SyncSet.
        /// 
        /// Uses Azure REST API version 2023-11-22.
        /// 
        /// Other available API versions: 2022-09-04, 2023-04-01, 2023-07-01-preview, 2023-09-04. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redhatopenshift [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSyncSetResult> Invoke(GetSyncSetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSyncSetResult>("azure-native:redhatopenshift:getSyncSet", args ?? new GetSyncSetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The operation returns properties of a SyncSet.
        /// 
        /// Uses Azure REST API version 2023-11-22.
        /// 
        /// Other available API versions: 2022-09-04, 2023-04-01, 2023-07-01-preview, 2023-09-04. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redhatopenshift [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSyncSetResult> Invoke(GetSyncSetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSyncSetResult>("azure-native:redhatopenshift:getSyncSet", args ?? new GetSyncSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSyncSetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the SyncSet resource.
        /// </summary>
        [Input("childResourceName", required: true)]
        public string ChildResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the OpenShift cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetSyncSetArgs()
        {
        }
        public static new GetSyncSetArgs Empty => new GetSyncSetArgs();
    }

    public sealed class GetSyncSetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the SyncSet resource.
        /// </summary>
        [Input("childResourceName", required: true)]
        public Input<string> ChildResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the OpenShift cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public GetSyncSetInvokeArgs()
        {
        }
        public static new GetSyncSetInvokeArgs Empty => new GetSyncSetInvokeArgs();
    }


    [OutputType]
    public sealed class GetSyncSetResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resources represents the SyncSets configuration.
        /// </summary>
        public readonly string? Resources;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSyncSetResult(
            string azureApiVersion,

            string id,

            string name,

            string? resources,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            Resources = resources;
            SystemData = systemData;
            Type = type;
        }
    }
}
