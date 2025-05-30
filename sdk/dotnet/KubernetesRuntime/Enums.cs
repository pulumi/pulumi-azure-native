// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.KubernetesRuntime
{
    /// <summary>
    /// Storage Class Access Mode
    /// </summary>
    [EnumType]
    public readonly struct AccessMode : IEquatable<AccessMode>
    {
        private readonly string _value;

        private AccessMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Read Write Once (RWO) access mode
        /// </summary>
        public static AccessMode ReadWriteOnce { get; } = new AccessMode("ReadWriteOnce");
        /// <summary>
        /// Read Write Many (RWX) access mode
        /// </summary>
        public static AccessMode ReadWriteMany { get; } = new AccessMode("ReadWriteMany");

        public static bool operator ==(AccessMode left, AccessMode right) => left.Equals(right);
        public static bool operator !=(AccessMode left, AccessMode right) => !left.Equals(right);

        public static explicit operator string(AccessMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AccessMode other && Equals(other);
        public bool Equals(AccessMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Advertise Mode
    /// </summary>
    [EnumType]
    public readonly struct AdvertiseMode : IEquatable<AdvertiseMode>
    {
        private readonly string _value;

        private AdvertiseMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// ARP advertise mode
        /// </summary>
        public static AdvertiseMode ARP { get; } = new AdvertiseMode("ARP");
        /// <summary>
        /// BGP advertise mode
        /// </summary>
        public static AdvertiseMode BGP { get; } = new AdvertiseMode("BGP");
        /// <summary>
        /// both ARP and BGP advertise mode
        /// </summary>
        public static AdvertiseMode Both { get; } = new AdvertiseMode("Both");

        public static bool operator ==(AdvertiseMode left, AdvertiseMode right) => left.Equals(right);
        public static bool operator !=(AdvertiseMode left, AdvertiseMode right) => !left.Equals(right);

        public static explicit operator string(AdvertiseMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AdvertiseMode other && Equals(other);
        public bool Equals(AdvertiseMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Allow single data node failure
    /// </summary>
    [EnumType]
    public readonly struct DataResilienceTier : IEquatable<DataResilienceTier>
    {
        private readonly string _value;

        private DataResilienceTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not data resilient
        /// </summary>
        public static DataResilienceTier NotDataResilient { get; } = new DataResilienceTier("NotDataResilient");
        /// <summary>
        /// Data resilient
        /// </summary>
        public static DataResilienceTier DataResilient { get; } = new DataResilienceTier("DataResilient");

        public static bool operator ==(DataResilienceTier left, DataResilienceTier right) => left.Equals(right);
        public static bool operator !=(DataResilienceTier left, DataResilienceTier right) => !left.Equals(right);

        public static explicit operator string(DataResilienceTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DataResilienceTier other && Equals(other);
        public bool Equals(DataResilienceTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Failover speed: NA, Slow, Fast
    /// </summary>
    [EnumType]
    public readonly struct FailoverTier : IEquatable<FailoverTier>
    {
        private readonly string _value;

        private FailoverTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not available Failover Tier
        /// </summary>
        public static FailoverTier NotAvailable { get; } = new FailoverTier("NotAvailable");
        /// <summary>
        /// Slow Failover Tier
        /// </summary>
        public static FailoverTier Slow { get; } = new FailoverTier("Slow");
        /// <summary>
        /// Fast Failover Tier
        /// </summary>
        public static FailoverTier Fast { get; } = new FailoverTier("Fast");
        /// <summary>
        /// Super Failover Tier
        /// </summary>
        public static FailoverTier Super { get; } = new FailoverTier("Super");

        public static bool operator ==(FailoverTier left, FailoverTier right) => left.Equals(right);
        public static bool operator !=(FailoverTier left, FailoverTier right) => !left.Equals(right);

        public static explicit operator string(FailoverTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FailoverTier other && Equals(other);
        public bool Equals(FailoverTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The action to take when a NFS volume is deleted. Default is Delete
    /// </summary>
    [EnumType]
    public readonly struct NfsDirectoryActionOnVolumeDeletion : IEquatable<NfsDirectoryActionOnVolumeDeletion>
    {
        private readonly string _value;

        private NfsDirectoryActionOnVolumeDeletion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// When the volume is deleted, delete the directory
        /// </summary>
        public static NfsDirectoryActionOnVolumeDeletion Delete { get; } = new NfsDirectoryActionOnVolumeDeletion("Delete");
        /// <summary>
        /// When the volume is deleted, retain the directory
        /// </summary>
        public static NfsDirectoryActionOnVolumeDeletion Retain { get; } = new NfsDirectoryActionOnVolumeDeletion("Retain");

        public static bool operator ==(NfsDirectoryActionOnVolumeDeletion left, NfsDirectoryActionOnVolumeDeletion right) => left.Equals(right);
        public static bool operator !=(NfsDirectoryActionOnVolumeDeletion left, NfsDirectoryActionOnVolumeDeletion right) => !left.Equals(right);

        public static explicit operator string(NfsDirectoryActionOnVolumeDeletion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NfsDirectoryActionOnVolumeDeletion other && Equals(other);
        public bool Equals(NfsDirectoryActionOnVolumeDeletion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Performance tier
    /// </summary>
    [EnumType]
    public readonly struct PerformanceTier : IEquatable<PerformanceTier>
    {
        private readonly string _value;

        private PerformanceTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Undefined Performance Tier
        /// </summary>
        public static PerformanceTier Undefined { get; } = new PerformanceTier("Undefined");
        /// <summary>
        /// Basic Performance Tier
        /// </summary>
        public static PerformanceTier Basic { get; } = new PerformanceTier("Basic");
        /// <summary>
        /// Standard Performance Tier
        /// </summary>
        public static PerformanceTier Standard { get; } = new PerformanceTier("Standard");
        /// <summary>
        /// Premium Performance Tier
        /// </summary>
        public static PerformanceTier Premium { get; } = new PerformanceTier("Premium");
        /// <summary>
        /// Ultra Performance Tier
        /// </summary>
        public static PerformanceTier Ultra { get; } = new PerformanceTier("Ultra");

        public static bool operator ==(PerformanceTier left, PerformanceTier right) => left.Equals(right);
        public static bool operator !=(PerformanceTier left, PerformanceTier right) => !left.Equals(right);

        public static explicit operator string(PerformanceTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PerformanceTier other && Equals(other);
        public bool Equals(PerformanceTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of the storage class.
    /// </summary>
    [EnumType]
    public readonly struct SCType : IEquatable<SCType>
    {
        private readonly string _value;

        private SCType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Native storage class
        /// </summary>
        public static SCType Native { get; } = new SCType("Native");
        /// <summary>
        /// RWX storage class
        /// </summary>
        public static SCType RWX { get; } = new SCType("RWX");
        /// <summary>
        /// Blob storage class
        /// </summary>
        public static SCType Blob { get; } = new SCType("Blob");
        /// <summary>
        /// NFS storage class
        /// </summary>
        public static SCType NFS { get; } = new SCType("NFS");
        /// <summary>
        /// SMB storage class
        /// </summary>
        public static SCType SMB { get; } = new SCType("SMB");

        public static bool operator ==(SCType left, SCType right) => left.Equals(right);
        public static bool operator !=(SCType left, SCType right) => !left.Equals(right);

        public static explicit operator string(SCType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SCType other && Equals(other);
        public bool Equals(SCType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Binding mode of volumes: Immediate, WaitForFirstConsumer
    /// </summary>
    [EnumType]
    public readonly struct VolumeBindingMode : IEquatable<VolumeBindingMode>
    {
        private readonly string _value;

        private VolumeBindingMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Immediate binding mode
        /// </summary>
        public static VolumeBindingMode Immediate { get; } = new VolumeBindingMode("Immediate");
        /// <summary>
        /// Wait for first consumer binding mode
        /// </summary>
        public static VolumeBindingMode WaitForFirstConsumer { get; } = new VolumeBindingMode("WaitForFirstConsumer");

        public static bool operator ==(VolumeBindingMode left, VolumeBindingMode right) => left.Equals(right);
        public static bool operator !=(VolumeBindingMode left, VolumeBindingMode right) => !left.Equals(right);

        public static explicit operator string(VolumeBindingMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VolumeBindingMode other && Equals(other);
        public bool Equals(VolumeBindingMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Volume can be expanded or not
    /// </summary>
    [EnumType]
    public readonly struct VolumeExpansion : IEquatable<VolumeExpansion>
    {
        private readonly string _value;

        private VolumeExpansion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Allow volume expansion
        /// </summary>
        public static VolumeExpansion Allow { get; } = new VolumeExpansion("Allow");
        /// <summary>
        /// Disallow volume expansion
        /// </summary>
        public static VolumeExpansion Disallow { get; } = new VolumeExpansion("Disallow");

        public static bool operator ==(VolumeExpansion left, VolumeExpansion right) => left.Equals(right);
        public static bool operator !=(VolumeExpansion left, VolumeExpansion right) => !left.Equals(right);

        public static explicit operator string(VolumeExpansion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VolumeExpansion other && Equals(other);
        public bool Equals(VolumeExpansion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
