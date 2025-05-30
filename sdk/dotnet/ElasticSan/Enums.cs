// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.ElasticSan
{
    /// <summary>
    /// The action of virtual network rule.
    /// </summary>
    [EnumType]
    public readonly struct Action : IEquatable<Action>
    {
        private readonly string _value;

        private Action(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Action Allow { get; } = new Action("Allow");

        public static bool operator ==(Action left, Action right) => left.Equals(right);
        public static bool operator !=(Action left, Action right) => !left.Equals(right);

        public static explicit operator string(Action value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Action other && Equals(other);
        public bool Equals(Action other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of encryption
    /// </summary>
    [EnumType]
    public readonly struct EncryptionType : IEquatable<EncryptionType>
    {
        private readonly string _value;

        private EncryptionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Volume is encrypted at rest with Platform managed key. It is the default encryption type.
        /// </summary>
        public static EncryptionType EncryptionAtRestWithPlatformKey { get; } = new EncryptionType("EncryptionAtRestWithPlatformKey");
        /// <summary>
        /// Volume is encrypted at rest with Customer managed key that can be changed and revoked by a customer.
        /// </summary>
        public static EncryptionType EncryptionAtRestWithCustomerManagedKey { get; } = new EncryptionType("EncryptionAtRestWithCustomerManagedKey");

        public static bool operator ==(EncryptionType left, EncryptionType right) => left.Equals(right);
        public static bool operator !=(EncryptionType left, EncryptionType right) => !left.Equals(right);

        public static explicit operator string(EncryptionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EncryptionType other && Equals(other);
        public bool Equals(EncryptionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The identity type.
    /// </summary>
    [EnumType]
    public readonly struct IdentityType : IEquatable<IdentityType>
    {
        private readonly string _value;

        private IdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IdentityType None { get; } = new IdentityType("None");
        public static IdentityType SystemAssigned { get; } = new IdentityType("SystemAssigned");
        public static IdentityType UserAssigned { get; } = new IdentityType("UserAssigned");

        public static bool operator ==(IdentityType left, IdentityType right) => left.Equals(right);
        public static bool operator !=(IdentityType left, IdentityType right) => !left.Equals(right);

        public static explicit operator string(IdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IdentityType other && Equals(other);
        public bool Equals(IdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
    /// </summary>
    [EnumType]
    public readonly struct PrivateEndpointServiceConnectionStatus : IEquatable<PrivateEndpointServiceConnectionStatus>
    {
        private readonly string _value;

        private PrivateEndpointServiceConnectionStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PrivateEndpointServiceConnectionStatus Pending { get; } = new PrivateEndpointServiceConnectionStatus("Pending");
        public static PrivateEndpointServiceConnectionStatus Approved { get; } = new PrivateEndpointServiceConnectionStatus("Approved");
        public static PrivateEndpointServiceConnectionStatus Failed { get; } = new PrivateEndpointServiceConnectionStatus("Failed");
        public static PrivateEndpointServiceConnectionStatus Rejected { get; } = new PrivateEndpointServiceConnectionStatus("Rejected");

        public static bool operator ==(PrivateEndpointServiceConnectionStatus left, PrivateEndpointServiceConnectionStatus right) => left.Equals(right);
        public static bool operator !=(PrivateEndpointServiceConnectionStatus left, PrivateEndpointServiceConnectionStatus right) => !left.Equals(right);

        public static explicit operator string(PrivateEndpointServiceConnectionStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PrivateEndpointServiceConnectionStatus other && Equals(other);
        public bool Equals(PrivateEndpointServiceConnectionStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Allow or disallow public network access to ElasticSan. Value is optional but if passed in, must be 'Enabled' or 'Disabled'.
    /// </summary>
    [EnumType]
    public readonly struct PublicNetworkAccess : IEquatable<PublicNetworkAccess>
    {
        private readonly string _value;

        private PublicNetworkAccess(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PublicNetworkAccess Enabled { get; } = new PublicNetworkAccess("Enabled");
        public static PublicNetworkAccess Disabled { get; } = new PublicNetworkAccess("Disabled");

        public static bool operator ==(PublicNetworkAccess left, PublicNetworkAccess right) => left.Equals(right);
        public static bool operator !=(PublicNetworkAccess left, PublicNetworkAccess right) => !left.Equals(right);

        public static explicit operator string(PublicNetworkAccess value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PublicNetworkAccess other && Equals(other);
        public bool Equals(PublicNetworkAccess other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The sku name.
    /// </summary>
    [EnumType]
    public readonly struct SkuName : IEquatable<SkuName>
    {
        private readonly string _value;

        private SkuName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Premium locally redundant storage
        /// </summary>
        public static SkuName Premium_LRS { get; } = new SkuName("Premium_LRS");
        /// <summary>
        /// Premium zone redundant storage
        /// </summary>
        public static SkuName Premium_ZRS { get; } = new SkuName("Premium_ZRS");

        public static bool operator ==(SkuName left, SkuName right) => left.Equals(right);
        public static bool operator !=(SkuName left, SkuName right) => !left.Equals(right);

        public static explicit operator string(SkuName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SkuName other && Equals(other);
        public bool Equals(SkuName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The sku tier.
    /// </summary>
    [EnumType]
    public readonly struct SkuTier : IEquatable<SkuTier>
    {
        private readonly string _value;

        private SkuTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Premium Tier
        /// </summary>
        public static SkuTier Premium { get; } = new SkuTier("Premium");

        public static bool operator ==(SkuTier left, SkuTier right) => left.Equals(right);
        public static bool operator !=(SkuTier left, SkuTier right) => !left.Equals(right);

        public static explicit operator string(SkuTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SkuTier other && Equals(other);
        public bool Equals(SkuTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of storage target
    /// </summary>
    [EnumType]
    public readonly struct StorageTargetType : IEquatable<StorageTargetType>
    {
        private readonly string _value;

        private StorageTargetType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static StorageTargetType Iscsi { get; } = new StorageTargetType("Iscsi");
        public static StorageTargetType None { get; } = new StorageTargetType("None");

        public static bool operator ==(StorageTargetType left, StorageTargetType right) => left.Equals(right);
        public static bool operator !=(StorageTargetType left, StorageTargetType right) => !left.Equals(right);

        public static explicit operator string(StorageTargetType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is StorageTargetType other && Equals(other);
        public bool Equals(StorageTargetType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// This enumerates the possible sources of a volume creation.
    /// </summary>
    [EnumType]
    public readonly struct VolumeCreateOption : IEquatable<VolumeCreateOption>
    {
        private readonly string _value;

        private VolumeCreateOption(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static VolumeCreateOption None { get; } = new VolumeCreateOption("None");
        public static VolumeCreateOption VolumeSnapshot { get; } = new VolumeCreateOption("VolumeSnapshot");
        public static VolumeCreateOption DiskSnapshot { get; } = new VolumeCreateOption("DiskSnapshot");
        public static VolumeCreateOption Disk { get; } = new VolumeCreateOption("Disk");
        public static VolumeCreateOption DiskRestorePoint { get; } = new VolumeCreateOption("DiskRestorePoint");

        public static bool operator ==(VolumeCreateOption left, VolumeCreateOption right) => left.Equals(right);
        public static bool operator !=(VolumeCreateOption left, VolumeCreateOption right) => !left.Equals(right);

        public static explicit operator string(VolumeCreateOption value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VolumeCreateOption other && Equals(other);
        public bool Equals(VolumeCreateOption other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
