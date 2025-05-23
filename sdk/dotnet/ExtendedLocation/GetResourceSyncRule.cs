// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ExtendedLocation
{
    public static class GetResourceSyncRule
    {
        /// <summary>
        /// Gets the details of the resourceSyncRule with a specified resource group, subscription id Custom Location resource name and Resource Sync Rule name.
        /// 
        /// Uses Azure REST API version 2021-08-31-preview.
        /// </summary>
        public static Task<GetResourceSyncRuleResult> InvokeAsync(GetResourceSyncRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourceSyncRuleResult>("azure-native:extendedlocation:getResourceSyncRule", args ?? new GetResourceSyncRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the resourceSyncRule with a specified resource group, subscription id Custom Location resource name and Resource Sync Rule name.
        /// 
        /// Uses Azure REST API version 2021-08-31-preview.
        /// </summary>
        public static Output<GetResourceSyncRuleResult> Invoke(GetResourceSyncRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceSyncRuleResult>("azure-native:extendedlocation:getResourceSyncRule", args ?? new GetResourceSyncRuleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the resourceSyncRule with a specified resource group, subscription id Custom Location resource name and Resource Sync Rule name.
        /// 
        /// Uses Azure REST API version 2021-08-31-preview.
        /// </summary>
        public static Output<GetResourceSyncRuleResult> Invoke(GetResourceSyncRuleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceSyncRuleResult>("azure-native:extendedlocation:getResourceSyncRule", args ?? new GetResourceSyncRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceSyncRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Resource Sync Rule name.
        /// </summary>
        [Input("childResourceName", required: true)]
        public string ChildResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Custom Locations name.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetResourceSyncRuleArgs()
        {
        }
        public static new GetResourceSyncRuleArgs Empty => new GetResourceSyncRuleArgs();
    }

    public sealed class GetResourceSyncRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Resource Sync Rule name.
        /// </summary>
        [Input("childResourceName", required: true)]
        public Input<string> ChildResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Custom Locations name.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public GetResourceSyncRuleInvokeArgs()
        {
        }
        public static new GetResourceSyncRuleInvokeArgs Empty => new GetResourceSyncRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourceSyncRuleResult
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
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Priority represents a priority of the Resource Sync Rule
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// Provisioning State for the Resource Sync Rule.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The second part, matchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn. The values set must be empty in the case of Exists and DoesNotExist. All of the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
        /// </summary>
        public readonly Outputs.ResourceSyncRulePropertiesResponseSelector? Selector;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// For an unmapped custom resource, its labels will be used to find matching resource sync rules. If this resource sync rule is one of the matching rules with highest priority, then the unmapped custom resource will be projected to the target resource group associated with this resource sync rule. The user creating this resource sync rule should have write permissions on the target resource group and this write permission will be validated when creating the resource sync rule.
        /// </summary>
        public readonly string? TargetResourceGroup;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetResourceSyncRuleResult(
            string azureApiVersion,

            string id,

            string location,

            string name,

            int? priority,

            string provisioningState,

            Outputs.ResourceSyncRulePropertiesResponseSelector? selector,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string? targetResourceGroup,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Location = location;
            Name = name;
            Priority = priority;
            ProvisioningState = provisioningState;
            Selector = selector;
            SystemData = systemData;
            Tags = tags;
            TargetResourceGroup = targetResourceGroup;
            Type = type;
        }
    }
}
