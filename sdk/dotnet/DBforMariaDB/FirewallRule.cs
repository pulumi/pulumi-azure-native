// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforMariaDB
{
    /// <summary>
    /// Represents a server firewall rule.
    /// 
    /// Uses Azure REST API version 2018-06-01. In version 2.x of the Azure Native provider, it used API version 2018-06-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:dbformariadb:FirewallRule")]
    public partial class FirewallRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The end IP address of the server firewall rule. Must be IPv4 format.
        /// </summary>
        [Output("endIpAddress")]
        public Output<string> EndIpAddress { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The start IP address of the server firewall rule. Must be IPv4 format.
        /// </summary>
        [Output("startIpAddress")]
        public Output<string> StartIpAddress { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallRule(string name, FirewallRuleArgs args, CustomResourceOptions? options = null)
            : base("azure-native:dbformariadb:FirewallRule", name, args ?? new FirewallRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:dbformariadb:FirewallRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:dbformariadb/v20180601:FirewallRule" },
                    new global::Pulumi.Alias { Type = "azure-native:dbformariadb/v20180601preview:FirewallRule" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FirewallRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FirewallRule(name, id, options);
        }
    }

    public sealed class FirewallRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The end IP address of the server firewall rule. Must be IPv4 format.
        /// </summary>
        [Input("endIpAddress", required: true)]
        public Input<string> EndIpAddress { get; set; } = null!;

        /// <summary>
        /// The name of the server firewall rule.
        /// </summary>
        [Input("firewallRuleName")]
        public Input<string>? FirewallRuleName { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// The start IP address of the server firewall rule. Must be IPv4 format.
        /// </summary>
        [Input("startIpAddress", required: true)]
        public Input<string> StartIpAddress { get; set; } = null!;

        public FirewallRuleArgs()
        {
        }
        public static new FirewallRuleArgs Empty => new FirewallRuleArgs();
    }
}
