// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Relationships
{
    public static class GetDependencyOfRelationship
    {
        /// <summary>
        /// Get a DependencyOfRelationship
        /// 
        /// Uses Azure REST API version 2023-09-01-preview.
        /// </summary>
        public static Task<GetDependencyOfRelationshipResult> InvokeAsync(GetDependencyOfRelationshipArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDependencyOfRelationshipResult>("azure-native:relationships:getDependencyOfRelationship", args ?? new GetDependencyOfRelationshipArgs(), options.WithDefaults());

        /// <summary>
        /// Get a DependencyOfRelationship
        /// 
        /// Uses Azure REST API version 2023-09-01-preview.
        /// </summary>
        public static Output<GetDependencyOfRelationshipResult> Invoke(GetDependencyOfRelationshipInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDependencyOfRelationshipResult>("azure-native:relationships:getDependencyOfRelationship", args ?? new GetDependencyOfRelationshipInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a DependencyOfRelationship
        /// 
        /// Uses Azure REST API version 2023-09-01-preview.
        /// </summary>
        public static Output<GetDependencyOfRelationshipResult> Invoke(GetDependencyOfRelationshipInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDependencyOfRelationshipResult>("azure-native:relationships:getDependencyOfRelationship", args ?? new GetDependencyOfRelationshipInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDependencyOfRelationshipArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of dependencyOf relationship.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public string ResourceUri { get; set; } = null!;

        public GetDependencyOfRelationshipArgs()
        {
        }
        public static new GetDependencyOfRelationshipArgs Empty => new GetDependencyOfRelationshipArgs();
    }

    public sealed class GetDependencyOfRelationshipInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of dependencyOf relationship.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public Input<string> ResourceUri { get; set; } = null!;

        public GetDependencyOfRelationshipInvokeArgs()
        {
        }
        public static new GetDependencyOfRelationshipInvokeArgs Empty => new GetDependencyOfRelationshipInvokeArgs();
    }


    [OutputType]
    public sealed class GetDependencyOfRelationshipResult
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
        /// The resource-specific properties for this resource.
        /// </summary>
        public readonly Outputs.DependencyOfRelationshipPropertiesResponse Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDependencyOfRelationshipResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.DependencyOfRelationshipPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Type = type;
        }
    }
}
