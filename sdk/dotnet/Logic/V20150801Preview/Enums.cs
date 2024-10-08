// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Logic.V20150801Preview
{
    /// <summary>
    /// The agreement type.
    /// </summary>
    [EnumType]
    public readonly struct AgreementType : IEquatable<AgreementType>
    {
        private readonly string _value;

        private AgreementType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AgreementType NotSpecified { get; } = new AgreementType("NotSpecified");
        public static AgreementType AS2 { get; } = new AgreementType("AS2");
        public static AgreementType X12 { get; } = new AgreementType("X12");
        public static AgreementType Edifact { get; } = new AgreementType("Edifact");

        public static bool operator ==(AgreementType left, AgreementType right) => left.Equals(right);
        public static bool operator !=(AgreementType left, AgreementType right) => !left.Equals(right);

        public static explicit operator string(AgreementType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AgreementType other && Equals(other);
        public bool Equals(AgreementType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The EDIFACT frame setting characterSet.
    /// </summary>
    [EnumType]
    public readonly struct EdifactCharacterSet : IEquatable<EdifactCharacterSet>
    {
        private readonly string _value;

        private EdifactCharacterSet(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EdifactCharacterSet NotSpecified { get; } = new EdifactCharacterSet("NotSpecified");
        public static EdifactCharacterSet UNOB { get; } = new EdifactCharacterSet("UNOB");
        public static EdifactCharacterSet UNOA { get; } = new EdifactCharacterSet("UNOA");
        public static EdifactCharacterSet UNOC { get; } = new EdifactCharacterSet("UNOC");
        public static EdifactCharacterSet UNOD { get; } = new EdifactCharacterSet("UNOD");
        public static EdifactCharacterSet UNOE { get; } = new EdifactCharacterSet("UNOE");
        public static EdifactCharacterSet UNOF { get; } = new EdifactCharacterSet("UNOF");
        public static EdifactCharacterSet UNOG { get; } = new EdifactCharacterSet("UNOG");
        public static EdifactCharacterSet UNOH { get; } = new EdifactCharacterSet("UNOH");
        public static EdifactCharacterSet UNOI { get; } = new EdifactCharacterSet("UNOI");
        public static EdifactCharacterSet UNOJ { get; } = new EdifactCharacterSet("UNOJ");
        public static EdifactCharacterSet UNOK { get; } = new EdifactCharacterSet("UNOK");
        public static EdifactCharacterSet UNOX { get; } = new EdifactCharacterSet("UNOX");
        public static EdifactCharacterSet UNOY { get; } = new EdifactCharacterSet("UNOY");
        public static EdifactCharacterSet KECA { get; } = new EdifactCharacterSet("KECA");

        public static bool operator ==(EdifactCharacterSet left, EdifactCharacterSet right) => left.Equals(right);
        public static bool operator !=(EdifactCharacterSet left, EdifactCharacterSet right) => !left.Equals(right);

        public static explicit operator string(EdifactCharacterSet value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EdifactCharacterSet other && Equals(other);
        public bool Equals(EdifactCharacterSet other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The EDIFACT frame setting decimal indicator.
    /// </summary>
    [EnumType]
    public readonly struct EdifactDecimalIndicator : IEquatable<EdifactDecimalIndicator>
    {
        private readonly string _value;

        private EdifactDecimalIndicator(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EdifactDecimalIndicator NotSpecified { get; } = new EdifactDecimalIndicator("NotSpecified");
        public static EdifactDecimalIndicator Comma { get; } = new EdifactDecimalIndicator("Comma");
        public static EdifactDecimalIndicator Decimal { get; } = new EdifactDecimalIndicator("Decimal");

        public static bool operator ==(EdifactDecimalIndicator left, EdifactDecimalIndicator right) => left.Equals(right);
        public static bool operator !=(EdifactDecimalIndicator left, EdifactDecimalIndicator right) => !left.Equals(right);

        public static explicit operator string(EdifactDecimalIndicator value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EdifactDecimalIndicator other && Equals(other);
        public bool Equals(EdifactDecimalIndicator other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The encryption algorithm.
    /// </summary>
    [EnumType]
    public readonly struct EncryptionAlgorithm : IEquatable<EncryptionAlgorithm>
    {
        private readonly string _value;

        private EncryptionAlgorithm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EncryptionAlgorithm NotSpecified { get; } = new EncryptionAlgorithm("NotSpecified");
        public static EncryptionAlgorithm None { get; } = new EncryptionAlgorithm("None");
        public static EncryptionAlgorithm DES3 { get; } = new EncryptionAlgorithm("DES3");
        public static EncryptionAlgorithm RC2 { get; } = new EncryptionAlgorithm("RC2");
        public static EncryptionAlgorithm AES128 { get; } = new EncryptionAlgorithm("AES128");
        public static EncryptionAlgorithm AES192 { get; } = new EncryptionAlgorithm("AES192");
        public static EncryptionAlgorithm AES256 { get; } = new EncryptionAlgorithm("AES256");

        public static bool operator ==(EncryptionAlgorithm left, EncryptionAlgorithm right) => left.Equals(right);
        public static bool operator !=(EncryptionAlgorithm left, EncryptionAlgorithm right) => !left.Equals(right);

        public static explicit operator string(EncryptionAlgorithm value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EncryptionAlgorithm other && Equals(other);
        public bool Equals(EncryptionAlgorithm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The signing or hashing algorithm.
    /// </summary>
    [EnumType]
    public readonly struct HashingAlgorithm : IEquatable<HashingAlgorithm>
    {
        private readonly string _value;

        private HashingAlgorithm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static HashingAlgorithm NotSpecified { get; } = new HashingAlgorithm("NotSpecified");
        public static HashingAlgorithm None { get; } = new HashingAlgorithm("None");
        public static HashingAlgorithm SHA2256 { get; } = new HashingAlgorithm("SHA2256");
        public static HashingAlgorithm SHA2384 { get; } = new HashingAlgorithm("SHA2384");
        public static HashingAlgorithm SHA2512 { get; } = new HashingAlgorithm("SHA2512");

        public static bool operator ==(HashingAlgorithm left, HashingAlgorithm right) => left.Equals(right);
        public static bool operator !=(HashingAlgorithm left, HashingAlgorithm right) => !left.Equals(right);

        public static explicit operator string(HashingAlgorithm value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HashingAlgorithm other && Equals(other);
        public bool Equals(HashingAlgorithm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The map type.
    /// </summary>
    [EnumType]
    public readonly struct MapType : IEquatable<MapType>
    {
        private readonly string _value;

        private MapType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MapType NotSpecified { get; } = new MapType("NotSpecified");
        public static MapType Xslt { get; } = new MapType("Xslt");

        public static bool operator ==(MapType left, MapType right) => left.Equals(right);
        public static bool operator !=(MapType left, MapType right) => !left.Equals(right);

        public static explicit operator string(MapType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MapType other && Equals(other);
        public bool Equals(MapType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The message filter type.
    /// </summary>
    [EnumType]
    public readonly struct MessageFilterType : IEquatable<MessageFilterType>
    {
        private readonly string _value;

        private MessageFilterType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MessageFilterType NotSpecified { get; } = new MessageFilterType("NotSpecified");
        public static MessageFilterType Include { get; } = new MessageFilterType("Include");
        public static MessageFilterType Exclude { get; } = new MessageFilterType("Exclude");

        public static bool operator ==(MessageFilterType left, MessageFilterType right) => left.Equals(right);
        public static bool operator !=(MessageFilterType left, MessageFilterType right) => !left.Equals(right);

        public static explicit operator string(MessageFilterType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MessageFilterType other && Equals(other);
        public bool Equals(MessageFilterType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The partner type.
    /// </summary>
    [EnumType]
    public readonly struct PartnerType : IEquatable<PartnerType>
    {
        private readonly string _value;

        private PartnerType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PartnerType NotSpecified { get; } = new PartnerType("NotSpecified");
        public static PartnerType B2B { get; } = new PartnerType("B2B");

        public static bool operator ==(PartnerType left, PartnerType right) => left.Equals(right);
        public static bool operator !=(PartnerType left, PartnerType right) => !left.Equals(right);

        public static explicit operator string(PartnerType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PartnerType other && Equals(other);
        public bool Equals(PartnerType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The schema type.
    /// </summary>
    [EnumType]
    public readonly struct SchemaType : IEquatable<SchemaType>
    {
        private readonly string _value;

        private SchemaType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SchemaType NotSpecified { get; } = new SchemaType("NotSpecified");
        public static SchemaType Xml { get; } = new SchemaType("Xml");

        public static bool operator ==(SchemaType left, SchemaType right) => left.Equals(right);
        public static bool operator !=(SchemaType left, SchemaType right) => !left.Equals(right);

        public static explicit operator string(SchemaType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SchemaType other && Equals(other);
        public bool Equals(SchemaType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The segment terminator suffix.
    /// </summary>
    [EnumType]
    public readonly struct SegmentTerminatorSuffix : IEquatable<SegmentTerminatorSuffix>
    {
        private readonly string _value;

        private SegmentTerminatorSuffix(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SegmentTerminatorSuffix NotSpecified { get; } = new SegmentTerminatorSuffix("NotSpecified");
        public static SegmentTerminatorSuffix None { get; } = new SegmentTerminatorSuffix("None");
        public static SegmentTerminatorSuffix CR { get; } = new SegmentTerminatorSuffix("CR");
        public static SegmentTerminatorSuffix LF { get; } = new SegmentTerminatorSuffix("LF");
        public static SegmentTerminatorSuffix CRLF { get; } = new SegmentTerminatorSuffix("CRLF");

        public static bool operator ==(SegmentTerminatorSuffix left, SegmentTerminatorSuffix right) => left.Equals(right);
        public static bool operator !=(SegmentTerminatorSuffix left, SegmentTerminatorSuffix right) => !left.Equals(right);

        public static explicit operator string(SegmentTerminatorSuffix value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SegmentTerminatorSuffix other && Equals(other);
        public bool Equals(SegmentTerminatorSuffix other) => string.Equals(_value, other._value, StringComparison.Ordinal);

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

        public static SkuName NotSpecified { get; } = new SkuName("NotSpecified");
        public static SkuName Free { get; } = new SkuName("Free");
        public static SkuName Shared { get; } = new SkuName("Shared");
        public static SkuName Basic { get; } = new SkuName("Basic");
        public static SkuName Standard { get; } = new SkuName("Standard");
        public static SkuName Premium { get; } = new SkuName("Premium");

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
    /// The trailing separator policy.
    /// </summary>
    [EnumType]
    public readonly struct TrailingSeparatorPolicy : IEquatable<TrailingSeparatorPolicy>
    {
        private readonly string _value;

        private TrailingSeparatorPolicy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TrailingSeparatorPolicy NotSpecified { get; } = new TrailingSeparatorPolicy("NotSpecified");
        public static TrailingSeparatorPolicy NotAllowed { get; } = new TrailingSeparatorPolicy("NotAllowed");
        public static TrailingSeparatorPolicy Optional { get; } = new TrailingSeparatorPolicy("Optional");
        public static TrailingSeparatorPolicy Mandatory { get; } = new TrailingSeparatorPolicy("Mandatory");

        public static bool operator ==(TrailingSeparatorPolicy left, TrailingSeparatorPolicy right) => left.Equals(right);
        public static bool operator !=(TrailingSeparatorPolicy left, TrailingSeparatorPolicy right) => !left.Equals(right);

        public static explicit operator string(TrailingSeparatorPolicy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TrailingSeparatorPolicy other && Equals(other);
        public bool Equals(TrailingSeparatorPolicy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The usage indicator.
    /// </summary>
    [EnumType]
    public readonly struct UsageIndicator : IEquatable<UsageIndicator>
    {
        private readonly string _value;

        private UsageIndicator(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static UsageIndicator NotSpecified { get; } = new UsageIndicator("NotSpecified");
        public static UsageIndicator Test { get; } = new UsageIndicator("Test");
        public static UsageIndicator Information { get; } = new UsageIndicator("Information");
        public static UsageIndicator Production { get; } = new UsageIndicator("Production");

        public static bool operator ==(UsageIndicator left, UsageIndicator right) => left.Equals(right);
        public static bool operator !=(UsageIndicator left, UsageIndicator right) => !left.Equals(right);

        public static explicit operator string(UsageIndicator value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is UsageIndicator other && Equals(other);
        public bool Equals(UsageIndicator other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The X12 character set.
    /// </summary>
    [EnumType]
    public readonly struct X12CharacterSet : IEquatable<X12CharacterSet>
    {
        private readonly string _value;

        private X12CharacterSet(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static X12CharacterSet NotSpecified { get; } = new X12CharacterSet("NotSpecified");
        public static X12CharacterSet Basic { get; } = new X12CharacterSet("Basic");
        public static X12CharacterSet Extended { get; } = new X12CharacterSet("Extended");
        public static X12CharacterSet UTF8 { get; } = new X12CharacterSet("UTF8");

        public static bool operator ==(X12CharacterSet left, X12CharacterSet right) => left.Equals(right);
        public static bool operator !=(X12CharacterSet left, X12CharacterSet right) => !left.Equals(right);

        public static explicit operator string(X12CharacterSet value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is X12CharacterSet other && Equals(other);
        public bool Equals(X12CharacterSet other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The group header date format.
    /// </summary>
    [EnumType]
    public readonly struct X12DateFormat : IEquatable<X12DateFormat>
    {
        private readonly string _value;

        private X12DateFormat(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static X12DateFormat NotSpecified { get; } = new X12DateFormat("NotSpecified");
        public static X12DateFormat CCYYMMDD { get; } = new X12DateFormat("CCYYMMDD");
        public static X12DateFormat YYMMDD { get; } = new X12DateFormat("YYMMDD");

        public static bool operator ==(X12DateFormat left, X12DateFormat right) => left.Equals(right);
        public static bool operator !=(X12DateFormat left, X12DateFormat right) => !left.Equals(right);

        public static explicit operator string(X12DateFormat value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is X12DateFormat other && Equals(other);
        public bool Equals(X12DateFormat other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The group header time format.
    /// </summary>
    [EnumType]
    public readonly struct X12TimeFormat : IEquatable<X12TimeFormat>
    {
        private readonly string _value;

        private X12TimeFormat(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static X12TimeFormat NotSpecified { get; } = new X12TimeFormat("NotSpecified");
        public static X12TimeFormat HHMM { get; } = new X12TimeFormat("HHMM");
        public static X12TimeFormat HHMMSS { get; } = new X12TimeFormat("HHMMSS");
        public static X12TimeFormat HHMMSSdd { get; } = new X12TimeFormat("HHMMSSdd");
        public static X12TimeFormat HHMMSSd { get; } = new X12TimeFormat("HHMMSSd");

        public static bool operator ==(X12TimeFormat left, X12TimeFormat right) => left.Equals(right);
        public static bool operator !=(X12TimeFormat left, X12TimeFormat right) => !left.Equals(right);

        public static explicit operator string(X12TimeFormat value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is X12TimeFormat other && Equals(other);
        public bool Equals(X12TimeFormat other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
