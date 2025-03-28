// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.V20240701
{
    /// <summary>
    /// Specifies information about the SSH public key.
    /// </summary>
    [AzureNativeResourceType("azure-native:compute/v20240701:SshPublicKey")]
    public partial class SshPublicKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// SSH public key used to authenticate to a virtual machine through ssh. If this property is not initially provided when the resource is created, the publicKey property will be populated when generateKeyPair is called. If the public key is provided upon resource creation, the provided public key needs to be at least 2048-bit and in ssh-rsa format.
        /// </summary>
        [Output("publicKey")]
        public Output<string?> PublicKey { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SshPublicKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SshPublicKey(string name, SshPublicKeyArgs args, CustomResourceOptions? options = null)
            : base("azure-native:compute/v20240701:SshPublicKey", name, args ?? new SshPublicKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SshPublicKey(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:compute/v20240701:SshPublicKey", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20191201:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20200601:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20201201:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20210301:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20210401:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20210701:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20211101:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20220301:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20220801:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20221101:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20230301:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20230701:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20230901:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20240301:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20241101:SshPublicKey" },
                    new global::Pulumi.Alias { Type = "azure-native:compute:SshPublicKey" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SshPublicKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SshPublicKey Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SshPublicKey(name, id, options);
        }
    }

    public sealed class SshPublicKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// SSH public key used to authenticate to a virtual machine through ssh. If this property is not initially provided when the resource is created, the publicKey property will be populated when generateKeyPair is called. If the public key is provided upon resource creation, the provided public key needs to be at least 2048-bit and in ssh-rsa format.
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the SSH public key.
        /// </summary>
        [Input("sshPublicKeyName")]
        public Input<string>? SshPublicKeyName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public SshPublicKeyArgs()
        {
        }
        public static new SshPublicKeyArgs Empty => new SshPublicKeyArgs();
    }
}
