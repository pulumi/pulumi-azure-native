// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Cache.V20201001Preview
{
    /// <summary>
    /// The type of RedisEnterprise cluster to deploy. Possible values: (Enterprise_E10, EnterpriseFlash_F300 etc.)
    /// </summary>
    [EnumType]
    public readonly struct SkuName : IEquatable<SkuName>
    {
        private readonly string _value;

        private SkuName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SkuName Enterprise_E10 { get; } = new SkuName("Enterprise_E10");
        public static SkuName Enterprise_E20 { get; } = new SkuName("Enterprise_E20");
        public static SkuName Enterprise_E50 { get; } = new SkuName("Enterprise_E50");
        public static SkuName Enterprise_E100 { get; } = new SkuName("Enterprise_E100");
        public static SkuName EnterpriseFlash_F300 { get; } = new SkuName("EnterpriseFlash_F300");
        public static SkuName EnterpriseFlash_F700 { get; } = new SkuName("EnterpriseFlash_F700");
        public static SkuName EnterpriseFlash_F1500 { get; } = new SkuName("EnterpriseFlash_F1500");

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
}
