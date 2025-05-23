// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetFirewallPolicyRuleCollectionGroup
    {
        /// <summary>
        /// Gets the specified FirewallPolicyRuleCollectionGroup.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetFirewallPolicyRuleCollectionGroupResult> InvokeAsync(GetFirewallPolicyRuleCollectionGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallPolicyRuleCollectionGroupResult>("azure-native:network:getFirewallPolicyRuleCollectionGroup", args ?? new GetFirewallPolicyRuleCollectionGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified FirewallPolicyRuleCollectionGroup.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetFirewallPolicyRuleCollectionGroupResult> Invoke(GetFirewallPolicyRuleCollectionGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallPolicyRuleCollectionGroupResult>("azure-native:network:getFirewallPolicyRuleCollectionGroup", args ?? new GetFirewallPolicyRuleCollectionGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified FirewallPolicyRuleCollectionGroup.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetFirewallPolicyRuleCollectionGroupResult> Invoke(GetFirewallPolicyRuleCollectionGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallPolicyRuleCollectionGroupResult>("azure-native:network:getFirewallPolicyRuleCollectionGroup", args ?? new GetFirewallPolicyRuleCollectionGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallPolicyRuleCollectionGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Firewall Policy.
        /// </summary>
        [Input("firewallPolicyName", required: true)]
        public string FirewallPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the FirewallPolicyRuleCollectionGroup.
        /// </summary>
        [Input("ruleCollectionGroupName", required: true)]
        public string RuleCollectionGroupName { get; set; } = null!;

        public GetFirewallPolicyRuleCollectionGroupArgs()
        {
        }
        public static new GetFirewallPolicyRuleCollectionGroupArgs Empty => new GetFirewallPolicyRuleCollectionGroupArgs();
    }

    public sealed class GetFirewallPolicyRuleCollectionGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Firewall Policy.
        /// </summary>
        [Input("firewallPolicyName", required: true)]
        public Input<string> FirewallPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the FirewallPolicyRuleCollectionGroup.
        /// </summary>
        [Input("ruleCollectionGroupName", required: true)]
        public Input<string> RuleCollectionGroupName { get; set; } = null!;

        public GetFirewallPolicyRuleCollectionGroupInvokeArgs()
        {
        }
        public static new GetFirewallPolicyRuleCollectionGroupInvokeArgs Empty => new GetFirewallPolicyRuleCollectionGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallPolicyRuleCollectionGroupResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Priority of the Firewall Policy Rule Collection Group resource.
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// The provisioning state of the firewall policy rule collection group resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Group of Firewall Policy rule collections.
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.FirewallPolicyFilterRuleCollectionResponse, Outputs.FirewallPolicyNatRuleCollectionResponse>> RuleCollections;
        /// <summary>
        /// A read-only string that represents the size of the FirewallPolicyRuleCollectionGroupProperties in MB. (ex 1.2MB)
        /// </summary>
        public readonly string Size;
        /// <summary>
        /// Rule Group type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetFirewallPolicyRuleCollectionGroupResult(
            string azureApiVersion,

            string etag,

            string? id,

            string? name,

            int? priority,

            string provisioningState,

            ImmutableArray<Union<Outputs.FirewallPolicyFilterRuleCollectionResponse, Outputs.FirewallPolicyNatRuleCollectionResponse>> ruleCollections,

            string size,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Etag = etag;
            Id = id;
            Name = name;
            Priority = priority;
            ProvisioningState = provisioningState;
            RuleCollections = ruleCollections;
            Size = size;
            Type = type;
        }
    }
}
