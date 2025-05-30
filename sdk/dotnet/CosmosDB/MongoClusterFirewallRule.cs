// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB
{
    /// <summary>
    /// Represents a mongo cluster firewall rule.
    /// 
    /// Uses Azure REST API version 2024-02-15-preview.
    /// 
    /// Other available API versions: 2023-03-01-preview, 2023-03-15-preview, 2023-09-15-preview, 2023-11-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:cosmosdb:MongoClusterFirewallRule")]
    public partial class MongoClusterFirewallRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The end IP address of the mongo cluster firewall rule. Must be IPv4 format.
        /// </summary>
        [Output("endIpAddress")]
        public Output<string> EndIpAddress { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the firewall rule.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The start IP address of the mongo cluster firewall rule. Must be IPv4 format.
        /// </summary>
        [Output("startIpAddress")]
        public Output<string> StartIpAddress { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a MongoClusterFirewallRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MongoClusterFirewallRule(string name, MongoClusterFirewallRuleArgs args, CustomResourceOptions? options = null)
            : base("azure-native:cosmosdb:MongoClusterFirewallRule", name, args ?? new MongoClusterFirewallRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MongoClusterFirewallRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:cosmosdb:MongoClusterFirewallRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230301preview:MongoClusterFirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230315preview:MongoClusterFirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230915preview:MongoClusterFirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20231115preview:MongoClusterFirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20240215preview:MongoClusterFirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230315preview:MongoClusterFirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230915preview:MongoClusterFirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20231115preview:MongoClusterFirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240215preview:MongoClusterFirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240301preview:FirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240601preview:FirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240701:FirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20241001preview:FirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb:FirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb:MongoClusterFirewallRule" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MongoClusterFirewallRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MongoClusterFirewallRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MongoClusterFirewallRule(name, id, options);
        }
    }

    public sealed class MongoClusterFirewallRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The end IP address of the mongo cluster firewall rule. Must be IPv4 format.
        /// </summary>
        [Input("endIpAddress", required: true)]
        public Input<string> EndIpAddress { get; set; } = null!;

        /// <summary>
        /// The name of the mongo cluster firewall rule.
        /// </summary>
        [Input("firewallRuleName")]
        public Input<string>? FirewallRuleName { get; set; }

        /// <summary>
        /// The name of the mongo cluster.
        /// </summary>
        [Input("mongoClusterName", required: true)]
        public Input<string> MongoClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The start IP address of the mongo cluster firewall rule. Must be IPv4 format.
        /// </summary>
        [Input("startIpAddress", required: true)]
        public Input<string> StartIpAddress { get; set; } = null!;

        public MongoClusterFirewallRuleArgs()
        {
        }
        public static new MongoClusterFirewallRuleArgs Empty => new MongoClusterFirewallRuleArgs();
    }
}
