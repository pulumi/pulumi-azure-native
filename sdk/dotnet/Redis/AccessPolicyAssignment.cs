// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Redis
{
    /// <summary>
    /// Response to an operation on access policy assignment
    /// 
    /// Uses Azure REST API version 2024-11-01.
    /// 
    /// Other available API versions: 2023-05-01-preview, 2023-08-01, 2024-03-01, 2024-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redis [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:redis:AccessPolicyAssignment")]
    public partial class AccessPolicyAssignment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the access policy that is being assigned
        /// </summary>
        [Output("accessPolicyName")]
        public Output<string> AccessPolicyName { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Object Id to assign access policy to
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;

        /// <summary>
        /// User friendly name for object id. Also represents username for token based authentication
        /// </summary>
        [Output("objectIdAlias")]
        public Output<string> ObjectIdAlias { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of an access policy assignment set
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AccessPolicyAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessPolicyAssignment(string name, AccessPolicyAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azure-native:redis:AccessPolicyAssignment", name, args ?? new AccessPolicyAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessPolicyAssignment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:redis:AccessPolicyAssignment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:cache/v20230501preview:AccessPolicyAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:cache/v20230801:AccessPolicyAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:cache/v20240301:AccessPolicyAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:cache/v20240401preview:AccessPolicyAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:cache/v20241101:AccessPolicyAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:cache:AccessPolicyAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:redis/v20230501preview:AccessPolicyAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:redis/v20230801:AccessPolicyAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:redis/v20240301:AccessPolicyAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:redis/v20240401preview:AccessPolicyAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:redis/v20241101:AccessPolicyAssignment" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AccessPolicyAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessPolicyAssignment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccessPolicyAssignment(name, id, options);
        }
    }

    public sealed class AccessPolicyAssignmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the access policy assignment.
        /// </summary>
        [Input("accessPolicyAssignmentName")]
        public Input<string>? AccessPolicyAssignmentName { get; set; }

        /// <summary>
        /// The name of the access policy that is being assigned
        /// </summary>
        [Input("accessPolicyName", required: true)]
        public Input<string> AccessPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the Redis cache.
        /// </summary>
        [Input("cacheName", required: true)]
        public Input<string> CacheName { get; set; } = null!;

        /// <summary>
        /// Object Id to assign access policy to
        /// </summary>
        [Input("objectId", required: true)]
        public Input<string> ObjectId { get; set; } = null!;

        /// <summary>
        /// User friendly name for object id. Also represents username for token based authentication
        /// </summary>
        [Input("objectIdAlias", required: true)]
        public Input<string> ObjectIdAlias { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public AccessPolicyAssignmentArgs()
        {
        }
        public static new AccessPolicyAssignmentArgs Empty => new AccessPolicyAssignmentArgs();
    }
}
