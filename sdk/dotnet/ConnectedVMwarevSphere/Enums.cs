// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.ConnectedVMwarevSphere
{
    /// <summary>
    /// Gets or sets the disk mode.
    /// </summary>
    [EnumType]
    public readonly struct DiskMode : IEquatable<DiskMode>
    {
        private readonly string _value;

        private DiskMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DiskMode Persistent { get; } = new DiskMode("persistent");
        public static DiskMode Independent_persistent { get; } = new DiskMode("independent_persistent");
        public static DiskMode Independent_nonpersistent { get; } = new DiskMode("independent_nonpersistent");

        public static bool operator ==(DiskMode left, DiskMode right) => left.Equals(right);
        public static bool operator !=(DiskMode left, DiskMode right) => !left.Equals(right);

        public static explicit operator string(DiskMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DiskMode other && Equals(other);
        public bool Equals(DiskMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the disk backing type.
    /// </summary>
    [EnumType]
    public readonly struct DiskType : IEquatable<DiskType>
    {
        private readonly string _value;

        private DiskType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DiskType Flat { get; } = new DiskType("flat");
        public static DiskType Pmem { get; } = new DiskType("pmem");
        public static DiskType Rawphysical { get; } = new DiskType("rawphysical");
        public static DiskType Rawvirtual { get; } = new DiskType("rawvirtual");
        public static DiskType Sparse { get; } = new DiskType("sparse");
        public static DiskType Sesparse { get; } = new DiskType("sesparse");
        public static DiskType Unknown { get; } = new DiskType("unknown");

        public static bool operator ==(DiskType left, DiskType right) => left.Equals(right);
        public static bool operator !=(DiskType left, DiskType right) => !left.Equals(right);

        public static explicit operator string(DiskType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DiskType other && Equals(other);
        public bool Equals(DiskType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Firmware type
    /// </summary>
    [EnumType]
    public readonly struct FirmwareType : IEquatable<FirmwareType>
    {
        private readonly string _value;

        private FirmwareType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FirmwareType Bios { get; } = new FirmwareType("bios");
        public static FirmwareType Efi { get; } = new FirmwareType("efi");

        public static bool operator ==(FirmwareType left, FirmwareType right) => left.Equals(right);
        public static bool operator !=(FirmwareType left, FirmwareType right) => !left.Equals(right);

        public static explicit operator string(FirmwareType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FirmwareType other && Equals(other);
        public bool Equals(FirmwareType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the nic allocation method.
    /// </summary>
    [EnumType]
    public readonly struct IPAddressAllocationMethod : IEquatable<IPAddressAllocationMethod>
    {
        private readonly string _value;

        private IPAddressAllocationMethod(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IPAddressAllocationMethod Unset { get; } = new IPAddressAllocationMethod("unset");
        public static IPAddressAllocationMethod @Dynamic { get; } = new IPAddressAllocationMethod("dynamic");
        public static IPAddressAllocationMethod @Static { get; } = new IPAddressAllocationMethod("static");
        public static IPAddressAllocationMethod Linklayer { get; } = new IPAddressAllocationMethod("linklayer");
        public static IPAddressAllocationMethod Random { get; } = new IPAddressAllocationMethod("random");
        public static IPAddressAllocationMethod Other { get; } = new IPAddressAllocationMethod("other");

        public static bool operator ==(IPAddressAllocationMethod left, IPAddressAllocationMethod right) => left.Equals(right);
        public static bool operator !=(IPAddressAllocationMethod left, IPAddressAllocationMethod right) => !left.Equals(right);

        public static explicit operator string(IPAddressAllocationMethod value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IPAddressAllocationMethod other && Equals(other);
        public bool Equals(IPAddressAllocationMethod other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of managed service identity.
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
    /// They inventory type.
    /// </summary>
    [EnumType]
    public readonly struct InventoryType : IEquatable<InventoryType>
    {
        private readonly string _value;

        private InventoryType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InventoryType ResourcePool { get; } = new InventoryType("ResourcePool");
        public static InventoryType VirtualMachine { get; } = new InventoryType("VirtualMachine");
        public static InventoryType VirtualMachineTemplate { get; } = new InventoryType("VirtualMachineTemplate");
        public static InventoryType VirtualNetwork { get; } = new InventoryType("VirtualNetwork");
        public static InventoryType Cluster { get; } = new InventoryType("Cluster");
        public static InventoryType Datastore { get; } = new InventoryType("Datastore");
        public static InventoryType Host { get; } = new InventoryType("Host");

        public static bool operator ==(InventoryType left, InventoryType right) => left.Equals(right);
        public static bool operator !=(InventoryType left, InventoryType right) => !left.Equals(right);

        public static explicit operator string(InventoryType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InventoryType other && Equals(other);
        public bool Equals(InventoryType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// NIC type
    /// </summary>
    [EnumType]
    public readonly struct NICType : IEquatable<NICType>
    {
        private readonly string _value;

        private NICType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static NICType Vmxnet3 { get; } = new NICType("vmxnet3");
        public static NICType Vmxnet2 { get; } = new NICType("vmxnet2");
        public static NICType Vmxnet { get; } = new NICType("vmxnet");
        public static NICType E1000 { get; } = new NICType("e1000");
        public static NICType E1000e { get; } = new NICType("e1000e");
        public static NICType Pcnet32 { get; } = new NICType("pcnet32");

        public static bool operator ==(NICType left, NICType right) => left.Equals(right);
        public static bool operator !=(NICType left, NICType right) => !left.Equals(right);

        public static explicit operator string(NICType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NICType other && Equals(other);
        public bool Equals(NICType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the type of the os.
    /// </summary>
    [EnumType]
    public readonly struct OsType : IEquatable<OsType>
    {
        private readonly string _value;

        private OsType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OsType Windows { get; } = new OsType("Windows");
        public static OsType Linux { get; } = new OsType("Linux");
        public static OsType Other { get; } = new OsType("Other");

        public static bool operator ==(OsType left, OsType right) => left.Equals(right);
        public static bool operator !=(OsType left, OsType right) => !left.Equals(right);

        public static explicit operator string(OsType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OsType other && Equals(other);
        public bool Equals(OsType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the power on boot.
    /// </summary>
    [EnumType]
    public readonly struct PowerOnBootOption : IEquatable<PowerOnBootOption>
    {
        private readonly string _value;

        private PowerOnBootOption(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PowerOnBootOption Enabled { get; } = new PowerOnBootOption("enabled");
        public static PowerOnBootOption Disabled { get; } = new PowerOnBootOption("disabled");

        public static bool operator ==(PowerOnBootOption left, PowerOnBootOption right) => left.Equals(right);
        public static bool operator !=(PowerOnBootOption left, PowerOnBootOption right) => !left.Equals(right);

        public static explicit operator string(PowerOnBootOption value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PowerOnBootOption other && Equals(other);
        public bool Equals(PowerOnBootOption other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the guest agent provisioning action.
    /// </summary>
    [EnumType]
    public readonly struct ProvisioningAction : IEquatable<ProvisioningAction>
    {
        private readonly string _value;

        private ProvisioningAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProvisioningAction Install { get; } = new ProvisioningAction("install");
        public static ProvisioningAction Uninstall { get; } = new ProvisioningAction("uninstall");
        public static ProvisioningAction Repair { get; } = new ProvisioningAction("repair");

        public static bool operator ==(ProvisioningAction left, ProvisioningAction right) => left.Equals(right);
        public static bool operator !=(ProvisioningAction left, ProvisioningAction right) => !left.Equals(right);

        public static explicit operator string(ProvisioningAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProvisioningAction other && Equals(other);
        public bool Equals(ProvisioningAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
