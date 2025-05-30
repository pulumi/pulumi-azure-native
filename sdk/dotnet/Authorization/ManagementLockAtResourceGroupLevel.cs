// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization
{
    /// <summary>
    /// The lock information.
    /// 
    /// Uses Azure REST API version 2020-05-01. In version 2.x of the Azure Native provider, it used API version 2020-05-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:authorization:ManagementLockAtResourceGroupLevel")]
    public partial class ManagementLockAtResourceGroupLevel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
        /// </summary>
        [Output("level")]
        public Output<string> Level { get; private set; } = null!;

        /// <summary>
        /// The name of the lock.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Notes about the lock. Maximum of 512 characters.
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        /// <summary>
        /// The owners of the lock.
        /// </summary>
        [Output("owners")]
        public Output<ImmutableArray<Outputs.ManagementLockOwnerResponse>> Owners { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The resource type of the lock - Microsoft.Authorization/locks.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ManagementLockAtResourceGroupLevel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagementLockAtResourceGroupLevel(string name, ManagementLockAtResourceGroupLevelArgs args, CustomResourceOptions? options = null)
            : base("azure-native:authorization:ManagementLockAtResourceGroupLevel", name, args ?? new ManagementLockAtResourceGroupLevelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagementLockAtResourceGroupLevel(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:authorization:ManagementLockAtResourceGroupLevel", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:authorization/v20150101:ManagementLockAtResourceGroupLevel" },
                    new global::Pulumi.Alias { Type = "azure-native:authorization/v20160901:ManagementLockAtResourceGroupLevel" },
                    new global::Pulumi.Alias { Type = "azure-native:authorization/v20170401:ManagementLockAtResourceGroupLevel" },
                    new global::Pulumi.Alias { Type = "azure-native:authorization/v20200501:ManagementLockAtResourceGroupLevel" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagementLockAtResourceGroupLevel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagementLockAtResourceGroupLevel Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagementLockAtResourceGroupLevel(name, id, options);
        }
    }

    public sealed class ManagementLockAtResourceGroupLevelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
        /// </summary>
        [Input("level", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Authorization.LockLevel> Level { get; set; } = null!;

        /// <summary>
        /// The lock name. The lock name can be a maximum of 260 characters. It cannot contain &lt;, &gt; %, &amp;, :, \, ?, /, or any control characters.
        /// </summary>
        [Input("lockName")]
        public Input<string>? LockName { get; set; }

        /// <summary>
        /// Notes about the lock. Maximum of 512 characters.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("owners")]
        private InputList<Inputs.ManagementLockOwnerArgs>? _owners;

        /// <summary>
        /// The owners of the lock.
        /// </summary>
        public InputList<Inputs.ManagementLockOwnerArgs> Owners
        {
            get => _owners ?? (_owners = new InputList<Inputs.ManagementLockOwnerArgs>());
            set => _owners = value;
        }

        /// <summary>
        /// The name of the resource group to lock.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ManagementLockAtResourceGroupLevelArgs()
        {
        }
        public static new ManagementLockAtResourceGroupLevelArgs Empty => new ManagementLockAtResourceGroupLevelArgs();
    }
}
