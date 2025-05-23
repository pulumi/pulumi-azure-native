// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.SecretSyncController
{
    /// <summary>
    /// The type of the extended location.
    /// </summary>
    [EnumType]
    public readonly struct ExtendedLocationType : IEquatable<ExtendedLocationType>
    {
        private readonly string _value;

        private ExtendedLocationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Azure Edge Zones location type
        /// </summary>
        public static ExtendedLocationType EdgeZone { get; } = new ExtendedLocationType("EdgeZone");
        /// <summary>
        /// Azure Custom Locations type
        /// </summary>
        public static ExtendedLocationType CustomLocation { get; } = new ExtendedLocationType("CustomLocation");

        public static bool operator ==(ExtendedLocationType left, ExtendedLocationType right) => left.Equals(right);
        public static bool operator !=(ExtendedLocationType left, ExtendedLocationType right) => !left.Equals(right);

        public static explicit operator string(ExtendedLocationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ExtendedLocationType other && Equals(other);
        public bool Equals(ExtendedLocationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type specifies the type of the Kubernetes secret object, e.g. "Opaque" or"kubernetes.io/tls". The controller must have permission to create secrets of the specified type.
    /// </summary>
    [EnumType]
    public readonly struct KubernetesSecretType : IEquatable<KubernetesSecretType>
    {
        private readonly string _value;

        private KubernetesSecretType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Opaque is the default secret type.
        /// </summary>
        public static KubernetesSecretType Opaque { get; } = new KubernetesSecretType("Opaque");
        /// <summary>
        /// The kubernetes.io/tls secret type is for storing a certificate and its associated key that are typically used for TLS.
        /// </summary>
        public static KubernetesSecretType Tls { get; } = new KubernetesSecretType("kubernetes.io/tls");

        public static bool operator ==(KubernetesSecretType left, KubernetesSecretType right) => left.Equals(right);
        public static bool operator !=(KubernetesSecretType left, KubernetesSecretType right) => !left.Equals(right);

        public static explicit operator string(KubernetesSecretType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KubernetesSecretType other && Equals(other);
        public bool Equals(KubernetesSecretType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
