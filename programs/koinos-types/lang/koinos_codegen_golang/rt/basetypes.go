package koinos

import (
    "bytes"
    "math/big"
    "encoding/binary"
    "encoding/json"
    "errors"
    "sync"
    "unicode/utf8"

    "github.com/btcsuite/btcutil/base58"
)

const bigIntNumericLiteralMin int64 = -9007199254740991 // -1 << 53
const bigIntNumericLiteralMax int64 =  9007199254740991 // 1 << 53 - 1


type Serializeable interface {
    Serialize(vb *VariableBlob) *VariableBlob
}

// --------------------------------
//  String
// --------------------------------

type String string

func NewString() *String {
    o := String("")
    return &o
}

func (n *String) Serialize(vb *VariableBlob) *VariableBlob {
    nvb := VariableBlob(make([]byte, len(*n)))
    copy(nvb, *n)

    return nvb.Serialize(vb)
}

func DeserializeString(vb *VariableBlob) (uint64,*String,error) {
    bytes, vb_ptr, err := DeserializeVariableBlob(vb)
    s := String("")
    if err != nil {
        return 0, &s, err
    }

    if !utf8.Valid(*vb_ptr) {
        return 0, &s, errors.New("String is not UTF-8 encoded")
    }
    s = String(*vb_ptr)

    return bytes, &s, nil
}

// --------------------------------
//  Boolean
// --------------------------------

type Boolean bool

func NewBoolean() *Boolean {
    o := Boolean(false)
    return &o
}

func (n *Boolean) Serialize(vb *VariableBlob) *VariableBlob {
    var b byte
    if *n {
        b = 1
    }
    x := append(*vb, b)
    return &x
}

func DeserializeBoolean(vb *VariableBlob) (uint64,*Boolean,error) {
    var b Boolean

    if len(*vb) < 1 {
        return 0, &b, errors.New("Unexpected EOF")
    }

    if (*vb)[0] == 1 {
        b = true
    } else if (*vb)[0] != 0 {
        return 0, &b, errors.New("Boolean must be 0 or 1")
    }

    return 1, &b, nil
}

// --------------------------------
//  Int8
// --------------------------------

type Int8 int8

func NewInt8() *Int8 {
    o := Int8(0)
    return &o
}

func (n *Int8) Serialize(vb *VariableBlob) *VariableBlob {
    ov := append(*vb, byte(*n))
    return &ov
}

func DeserializeInt8(vb *VariableBlob) (uint64,*Int8,error) {
    var i Int8

    if len(*vb) < 1 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = Int8((*vb)[0])

    return 1, &i, nil
}

// --------------------------------
//  UInt8
// --------------------------------

type UInt8 uint8

func NewUInt8() *UInt8 {
    o := UInt8(0)
    return &o
}

func (n *UInt8) Serialize(vb *VariableBlob) *VariableBlob {
    ov := append(*vb, byte(*n))
    return &ov
}

func DeserializeUInt8(vb *VariableBlob) (uint64,*UInt8,error) {
    var i UInt8

    if len(*vb) < 1 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = UInt8((*vb)[0])

    return 1, &i, nil
}

// --------------------------------
//  Int16
// --------------------------------

type Int16 int16

func NewInt16() *Int16 {
    o := Int16(0)
    return &o
}

func (n *Int16) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 2)
    binary.BigEndian.PutUint16(b, uint16(*n))
    ov := append(*vb, b...)
    return &ov
}

func DeserializeInt16(vb *VariableBlob) (uint64,*Int16,error) {
    var i Int16

    if len(*vb) < 2 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = Int16(binary.BigEndian.Uint16(*vb))

    return 2, &i, nil
}

// --------------------------------
//  UInt16
// --------------------------------

type UInt16 uint16

func NewUInt16() *UInt16 {
    o := UInt16(0)
    return &o
}

func (n *UInt16) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 2)
    binary.BigEndian.PutUint16(b, uint16(*n))
    ov := append(*vb, b...)
    return &ov
}

func DeserializeUInt16(vb *VariableBlob) (uint64,*UInt16,error) {
    var i UInt16

    if len(*vb) < 2 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = UInt16(binary.BigEndian.Uint16(*vb))

    return 2, &i, nil
}


// --------------------------------
//  Int32
// --------------------------------

type Int32 int32

func NewInt32() *Int32 {
    o := Int32(0)
    return &o
}

func (n *Int32) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 4)
    binary.BigEndian.PutUint32(b, uint32(*n))
    ov := append(*vb, b...)
    return &ov
}

func DeserializeInt32(vb *VariableBlob) (uint64,*Int32,error) {
    var i Int32

    if len(*vb) < 4 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = Int32(binary.BigEndian.Uint32(*vb))

    return 4, &i, nil
}

// --------------------------------
//  UInt32
// --------------------------------

type UInt32 uint32

func NewUInt32() *UInt32 {
    o := UInt32(0)
    return &o
}

func (n *UInt32) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 4)
    binary.BigEndian.PutUint32(b, uint32(*n))
    ov := append(*vb, b...)
    return &ov
}

func DeserializeUInt32(vb *VariableBlob) (uint64,*UInt32,error) {
    var i UInt32

    if len(*vb) < 4 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = UInt32(binary.BigEndian.Uint32(*vb))

    return 4, &i, nil
}

// --------------------------------
//  Int64
// --------------------------------

type Int64 int64

func NewInt64() *Int64 {
    o := Int64(0)
    return &o
}

func (n *Int64) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(*n))
    ov := append(*vb, b...)
    return &ov
}

func DeserializeInt64(vb *VariableBlob) (uint64,*Int64,error) {
    var i Int64

    if len(*vb) < 8 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = Int64(binary.BigEndian.Uint64(*vb))

    return 8, &i, nil
}

// --------------------------------
//  UInt64
// --------------------------------

type UInt64 uint64

func NewUInt64() *UInt64 {
    o := UInt64(0)
    return &o
}

func (n *UInt64) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(*n))
    ov := append(*vb, b...)
    return &ov
}

func DeserializeUInt64(vb *VariableBlob) (uint64,*UInt64,error) {
    var i UInt64

    if len(*vb) < 8 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = UInt64(binary.BigEndian.Uint64(*vb))

    return 8, &i, nil
}

// ----------------------------------------
//  Int128
// ----------------------------------------

type Int128 struct {
    Value big.Int
}

var (
    int128Max *Int128
    int128Min *Int128
    int128Once sync.Once
)

func initInt128Bounds() {
    int128Once.Do(func() {
        int128Max = NewInt128()
        nv, _ := int128Max.Value.SetString("170141183460469231731687303715884105727", 10)
        int128Max.Value = *nv
        int128Min = NewInt128()
        nv, _ = int128Min.Value.SetString("-170141183460469231731687303715884105728", 10)
        int128Min.Value = *nv
    })
}

func Int128Max() Int128 {
    initInt128Bounds()
    return *int128Max
}

func Int128Min() Int128 {
    initInt128Bounds()
    return *int128Min
}

func NewInt128() *Int128 {
    result := Int128{}
    result.Value = *big.NewInt(0)
    return &result
}

func NewInt128FromString(value string) (*Int128,error) {
    var result Int128 = Int128{}
    if nv, ok := result.Value.SetString(value, 10); ok == false {
        return nil, errors.New("Could not parse Int128")
    } else {
        max := Int128Max()
        min := Int128Min()
        if nv.Cmp(&max.Value) == 1 || nv.Cmp(&min.Value) == -1 {
            return nil, errors.New("Int128 is out of bounds")
        }
        result.Value = *nv
    }
    return &result, nil
}

func (n *Int128) Serialize(vb *VariableBlob) *VariableBlob {
    s := SerializeBigInt(&n.Value, 16, true)
    ov := append(*vb, *s...)
    return &ov
}

func DeserializeInt128(vb *VariableBlob) (uint64,*Int128,error) {
    bi_ptr, err := DeserializeBigInt(vb, 16, true)
    i := Int128{}

    if err != nil {
        return 0, &i, err
    }

    i.Value = *bi_ptr

    return 16, &i, nil
}

func (n *Int128) MarshalJSON() ([]byte, error) {
    if i := n.Value.Int64(); n.Value.IsInt64() && i <= bigIntNumericLiteralMax && i >= bigIntNumericLiteralMin {
        return json.Marshal(i)
    }

    s := n.Value.String()
    return json.Marshal(s)
}

func (n *Int128) UnmarshalJSON(b []byte) error {
    var s string
    err := json.Unmarshal(b, &s)

    if err != nil {
        var i int64
        if err = json.Unmarshal(b, &i); err != nil {
            return err
        }
        n.Value = *big.NewInt(i)
    } else {
        if nv, err := NewInt128FromString(s); err != nil {
            return err
        } else {
            *n = *nv
        }
    }

    return nil
}

// ----------------------------------------
//  UInt128
// ----------------------------------------

type UInt128 struct {
    Value big.Int
}

var (
    uint128Max *UInt128
    uint128Min *UInt128
    uint128Once sync.Once
)

func initUInt128Bounds() {
    uint128Once.Do(func() {
        uint128Max = NewUInt128()
        nv, _ := uint128Max.Value.SetString("340282366920938463463374607431768211455", 10)
        uint128Max.Value = *nv
        uint128Min = NewUInt128()
    })
}

func UInt128Max() UInt128 {
    initUInt128Bounds()
    return *uint128Max
}

func UInt128Min() UInt128 {
    initUInt128Bounds()
    return *uint128Min
}

func NewUInt128() *UInt128 {
    result := UInt128{}
    result.Value = *big.NewInt(0)
    return &result
}

func NewUInt128FromString(value string) (*UInt128,error) {
    var result UInt128 = UInt128{}
    if nv, ok := result.Value.SetString(value, 10); ok == false {
        return nil, errors.New("Could not parse UInt128")
    } else {
        max := UInt128Max()
        min := UInt128Min()
        if nv.Cmp(&max.Value) == 1 || nv.Cmp(&min.Value) == -1 {
            return nil, errors.New("UInt128 is out of bounds")
        }
        result.Value = *nv
    }
    return &result, nil
}

func (n *UInt128) Serialize(vb *VariableBlob) *VariableBlob {
    x := SerializeBigInt(&n.Value, 16, false)
    ov := append(*vb, *x...)
    return &ov
}

func DeserializeUInt128(vb *VariableBlob) (uint64,*UInt128,error) {
    bi_ptr, err := DeserializeBigInt(vb, 16, false)
    i := UInt128{}

    if err != nil {
        return 0, &i, err
    }

    i.Value = *bi_ptr

    return 16, &i, nil
}

func (n *UInt128) MarshalJSON() ([]byte, error) {
    if i := n.Value.Int64(); n.Value.IsInt64() && i <= bigIntNumericLiteralMax && i >= bigIntNumericLiteralMin {
        return json.Marshal(i)
    }

    s := n.Value.String()
    return json.Marshal(s)
}

func (n *UInt128) UnmarshalJSON(b []byte) error {
    var s string
    err := json.Unmarshal(b, &s)

    if err != nil {
        var i int64
        if err = json.Unmarshal(b, &i); err != nil {
            return err
        }
        if i < 0 {
            return errors.New("UInt128 is out of bounds")
        }
        n.Value = *big.NewInt(i)
    } else {
        if nv, err := NewUInt128FromString(s); err != nil {
            return err
        } else {
            *n = *nv
        }
    }

    return nil
}

// ----------------------------------------
//  Int160
// ----------------------------------------

type Int160 struct {
    Value big.Int
}

var (
    int160Max *Int160
    int160Min *Int160
    int160Once sync.Once
)

func initInt160Bounds() {
    int160Once.Do(func() {
        int160Max = NewInt160()
        nv, _ := int160Max.Value.SetString("730750818665451459101842416358141509827966271487", 10)
        int160Max.Value = *nv
        int160Min = NewInt160()
        nv, _ = int160Min.Value.SetString("-730750818665451459101842416358141509827966271488", 10)
        int160Min.Value = *nv
    })
}

func Int160Max() Int160 {
    initInt160Bounds()
    return *int160Max
}

func Int160Min() Int160 {
    initInt160Bounds()
    return *int160Min
}

func NewInt160() *Int160 {
    result := Int160{}
    result.Value = *big.NewInt(0)
    return &result
}

func NewInt160FromString(value string) (*Int160,error) {
    var result Int160 = Int160{}
    if nv, ok := result.Value.SetString(value, 10); ok == false {
        return nil, errors.New("Could not parse Int160")
    } else {
        max := Int160Max()
        min := Int160Min()
        if nv.Cmp(&max.Value) == 1 || nv.Cmp(&min.Value) == -1 {
            return nil, errors.New("Int160 is out of bounds")
        }
        result.Value = *nv
    }
    return &result, nil
}

func (n *Int160) Serialize(vb *VariableBlob) *VariableBlob {
    x := SerializeBigInt(&n.Value, 20, true)
    ov := append(*vb, *x...)
    return &ov
}

func DeserializeInt160(vb *VariableBlob) (uint64,*Int160,error) {
    bi_ptr, err := DeserializeBigInt(vb, 20, true)
    i := Int160{}

    if err != nil {
        return 0, &i, err
    }

    i.Value = *bi_ptr

    return 16, &i, nil
}

func (n *Int160) MarshalJSON() ([]byte, error) {
    if i := n.Value.Int64(); n.Value.IsInt64() && i <= bigIntNumericLiteralMax && i >= bigIntNumericLiteralMin {
        return json.Marshal(i)
    }

    s := n.Value.String()
    return json.Marshal(s)
}

func (n *Int160) UnmarshalJSON(b []byte) error {
    var s string
    err := json.Unmarshal(b, &s)

    if err != nil {
        var i int64
        if err = json.Unmarshal(b, &i); err != nil {
            return err
        }
        n.Value = *big.NewInt(i)
    } else {
        if nv, err := NewInt160FromString(s); err != nil {
            return err
        } else {
            *n = *nv
        }
    }

    return nil
}

// ----------------------------------------
//  UInt160
// ----------------------------------------

type UInt160 struct {
    Value big.Int
}

var (
    uint160Max *UInt160
    uint160Min *UInt160
    uint160Once sync.Once
)

func initUInt160Bounds() {
    uint160Once.Do(func() {
        uint160Max = NewUInt160()
        nv, _ := uint160Max.Value.SetString("1461501637330902918203684832716283019655932542975", 10)
        uint160Max.Value = *nv
        uint160Min = NewUInt160()
    })
}

func UInt160Max() UInt160 {
    initUInt160Bounds()
    return *uint160Max
}

func UInt160Min() UInt160 {
    initUInt160Bounds()
    return *uint160Min
}

func NewUInt160() *UInt160 {
    result := UInt160{}
    result.Value = *big.NewInt(0)
    return &result
}

func NewUInt160FromString(value string) (*UInt160,error) {
    var result UInt160 = UInt160{}
    if nv, ok := result.Value.SetString(value, 10); ok == false {
        return nil, errors.New("Could not parse UInt160")
    } else {
        max := UInt160Max()
        min := UInt160Min()
        if nv.Cmp(&max.Value) == 1 || nv.Cmp(&min.Value) == -1 {
            return nil, errors.New("UInt160 is out of bounds")
        }
        result.Value = *nv
    }
    return &result, nil
}

func (n *UInt160) Serialize(vb *VariableBlob) *VariableBlob {
    x := SerializeBigInt(&n.Value, 20, false)
    ov := append(*vb, *x...)
    return &ov
}

func DeserializeUInt160(vb *VariableBlob) (uint64,*UInt160,error) {
    bi_ptr, err := DeserializeBigInt(vb, 20, false)
    i := UInt160{}

    if err != nil {
        return 0, &i, err
    }

    i.Value = *bi_ptr

    return 16, &i, nil
}

func (n *UInt160) MarshalJSON() ([]byte, error) {
    if i := n.Value.Int64(); n.Value.IsInt64() && i <= bigIntNumericLiteralMax && i >= bigIntNumericLiteralMin {
        return json.Marshal(i)
    }

    s := n.Value.String()
    return json.Marshal(s)
}

func (n *UInt160) UnmarshalJSON(b []byte) error {
    var s string
    err := json.Unmarshal(b, &s)

    if err != nil {
        var i int64
        if err = json.Unmarshal(b, &i); err != nil {
            return err
        }
        if i < 0 {
            return errors.New("UInt160 is out of bounds")
        }
        n.Value = *big.NewInt(i)
    } else {
        if nv, err := NewUInt160FromString(s); err != nil {
            return err
        } else {
            *n = *nv
        }
    }

    return nil
}

// ----------------------------------------
//  Int256
// ----------------------------------------

type Int256 struct {
    Value big.Int
}

var (
    int256Max *Int256
    int256Min *Int256
    int256Once sync.Once
)

func initInt256Bounds() {
    int256Once.Do(func() {
        int256Max = NewInt256()
        nv, _ := int256Max.Value.SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
        int256Max.Value = *nv
        int256Min = NewInt256()
        nv, _ = int256Min.Value.SetString("-57896044618658097711785492504343953926634992332820282019728792003956564819968", 10)
        int256Min.Value = *nv
    })
}

func Int256Max() Int256 {
    initInt256Bounds()
    return *int256Max
}

func Int256Min() Int256 {
    initInt256Bounds()
    return *int256Min
}

func NewInt256() *Int256 {
    result := Int256{}
    result.Value = *big.NewInt(0)
    return &result
}

func NewInt256FromString(value string) (*Int256,error) {
    var result Int256 = Int256{}
    if nv, ok := result.Value.SetString(value, 10); ok == false {
        return nil, errors.New("Could not parse Int256")
    } else {
        max := Int256Max()
        min := Int256Min()
        if nv.Cmp(&max.Value) == 1 || nv.Cmp(&min.Value) == -1 {
            return nil, errors.New("Int256 is out of bounds")
        }
        result.Value = *nv
    }
    return &result, nil
}

func (n *Int256) Serialize(vb *VariableBlob) *VariableBlob {
    x := SerializeBigInt(&n.Value, 32, true)
    ov := append(*vb, *x...)
    return &ov
}

func DeserializeInt256(vb *VariableBlob) (uint64,*Int256,error) {
    bi_ptr, err := DeserializeBigInt(vb, 32, true)
    i := Int256{}

    if err != nil {
        return 0, &i, err
    }

    i.Value = *bi_ptr

    return 16, &i, nil
}

func (n *Int256) MarshalJSON() ([]byte, error) {
    if i := n.Value.Int64(); n.Value.IsInt64() && i <= bigIntNumericLiteralMax && i >= bigIntNumericLiteralMin {
        return json.Marshal(i)
    }

    s := n.Value.String()
    return json.Marshal(s)
}

func (n *Int256) UnmarshalJSON(b []byte) error {
    var s string
    err := json.Unmarshal(b, &s)

    if err != nil {
        var i int64
        if err = json.Unmarshal(b, &i); err != nil {
            return err
        }
        n.Value = *big.NewInt(i)
    } else {
        if nv, err := NewInt256FromString(s); err != nil {
            return err
        } else {
            *n = *nv
        }
    }

    return nil
}

// ----------------------------------------
//  UInt256
// ----------------------------------------

type UInt256 struct {
    Value big.Int
}

var (
    uint256Max *UInt256
    uint256Min *UInt256
    uint256Once sync.Once
)

func initUInt256Bounds() {
    uint256Once.Do(func() {
        uint256Max = NewUInt256()
        nv, _ := uint256Max.Value.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
        uint256Max.Value = *nv
        uint256Min = NewUInt256()
    })
}

func UInt256Max() UInt256 {
    initUInt256Bounds()
    return *uint256Max
}

func UInt256Min() UInt256 {
    initUInt256Bounds()
    return *uint256Min
}

func NewUInt256() *UInt256 {
    result := UInt256{}
    result.Value = *big.NewInt(0)
    return &result
}

func NewUInt256FromString(value string) (*UInt256,error) {
    var result UInt256 = UInt256{}
    if nv, ok := result.Value.SetString(value, 10); ok == false {
        return nil, errors.New("Could not parse UInt256")
    } else {
        max := UInt256Max()
        min := UInt256Min()
        if nv.Cmp(&max.Value) == 1 || nv.Cmp(&min.Value) == -1 {
            return nil, errors.New("UInt256 is out of bounds")
        }
        result.Value = *nv
    }
    return &result, nil
}

func (n *UInt256) Serialize(vb *VariableBlob) *VariableBlob {
    x := SerializeBigInt(&n.Value, 32, false)
    ov := append(*vb, *x...)
    return &ov
}

func DeserializeUInt256(vb *VariableBlob) (uint64,*UInt256,error) {
    bi_ptr, err := DeserializeBigInt(vb, 32, false)
    i := UInt256{}

    if err != nil {
        return 0, &i, err
    }

    i.Value = *bi_ptr

    return 16, &i, nil
}

func (n *UInt256) MarshalJSON() ([]byte, error) {
    if i := n.Value.Int64(); n.Value.IsInt64() && i <= bigIntNumericLiteralMax && i >= bigIntNumericLiteralMin {
        return json.Marshal(i)
    }

    s := n.Value.String()
    return json.Marshal(s)
}

func (n *UInt256) UnmarshalJSON(b []byte) error {
    var s string
    err := json.Unmarshal(b, &s)

    if err != nil {
        var i int64
        if err = json.Unmarshal(b, &i); err != nil {
            return err
        }
        if i < 0 {
            return errors.New("UInt256 is out of bounds")
        }
        n.Value = *big.NewInt(i)
    } else {
        if nv, err := NewUInt256FromString(s); err != nil {
            return err
        } else {
            *n = *nv
        }
    }

    return nil
}

// --------------------------------
//  VariableBlob
// --------------------------------

type VariableBlob []byte

// TODO: Make this variadic for size and size_hint
func NewVariableBlob() *VariableBlob {
    vb := VariableBlob(make([]byte, 0))
    return &vb
}

func (n *VariableBlob) Serialize(vb *VariableBlob) *VariableBlob {
    header := make([]byte, binary.MaxVarintLen64)
    bytes := binary.PutUvarint(header, uint64(len(*n)))
    ovb := append(*vb, header[:bytes]...)
    ovb = append(ovb, *n...)
    return &ovb
}

func DeserializeVariableBlob(vb *VariableBlob) (uint64,*VariableBlob,error) {
    size,bytes := binary.Uvarint(*vb)
    var result VariableBlob = VariableBlob(make([]byte, 0, size))
    if bytes <= 0 {
        return 0, &result, errors.New("Could not deserialize variable blob size")
    }

    if len(*vb) < bytes + int(size) {
        return 0, &result, errors.New("Unexpected EOF")
    }

    ovb := append(result, (*vb)[bytes:uint64(bytes)+size]...)
    return uint64(uint64(bytes)+size), &ovb, nil
}

func (n *VariableBlob) MarshalJSON() ([]byte, error) {
    s := EncodeBytes(*n)
    return json.Marshal(s)
}

func (n *VariableBlob) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }

    db,err := DecodeBytes(s)
    if err != nil {
        return err
    }
    if len(db) == 0 && len(s[1:]) > 0 {
        return errors.New("Unable to decode base58")
    }

    *n = db
    return nil
}

// --------------------------------
//  TimestampType
// --------------------------------

type TimestampType UInt64

func NewTimestampType() *TimestampType {
    o := TimestampType(0)
    return &o
}

func (n *TimestampType) Serialize(vb *VariableBlob) *VariableBlob {
    un := UInt64(*n)
    return un.Serialize(vb)
}

func DeserializeTimestampType(vb *VariableBlob) (uint64,*TimestampType,error) {
    i,x,err := DeserializeUInt64(vb)
    ox := TimestampType(*x)
    return i,&ox,err
}

// --------------------------------
//  BlockHeightType
// --------------------------------

type BlockHeightType UInt64

func NewBlockHeightType() *BlockHeightType {
    o := BlockHeightType(0)
    return &o
}

func (n *BlockHeightType) Serialize(vb *VariableBlob) *VariableBlob {
    un := UInt64(*n)
    return un.Serialize(vb)
}

func DeserializeBlockHeightType(vb *VariableBlob) (uint64,*BlockHeightType,error) {
    i,x,err := DeserializeUInt64(vb)
    ox := BlockHeightType(*x)
    return i,&ox,err
}

// --------------------------------
//  Multihash
// --------------------------------

type Multihash struct {
    Id UInt64 `json:"hash"`
    Digest VariableBlob `json:"digest"`
}

func NewMultihash() *Multihash {
    o := Multihash{}
    o.Id = UInt64(0)
    o.Digest = *NewVariableBlob()
    return &o
}

func (m0 *Multihash) Equals(m1 *Multihash) bool {
    return (m0.Id == m1.Id) && bytes.Equal(m0.Digest, m1.Digest)
}

func (m0 *Multihash) LessThan(m1 *Multihash) bool {
    r := m0.Id - m1.Id
    if (r < 0) {
        return true
    }
    if (r > 0) {
        return false
    }
    return (len(m0.Digest) - len(m1.Digest)) < 0
}

func (m0 *Multihash) GreaterThan(m1 *Multihash) bool {
    return !m0.Equals(m1) && !m0.LessThan(m1)
}

func (n *Multihash) Serialize(vb *VariableBlob) *VariableBlob {
    vb = EncodeVarint(vb, uint64(n.Id))
    return n.Digest.Serialize(vb)
}

func DeserializeMultihash(vb *VariableBlob) (uint64,*Multihash,error) {
    omh := Multihash{}
    id,isize := binary.Uvarint(*vb)
    if isize <= 0 {
        return 0, &omh, errors.New("Could not deserialize multihash id")
    }
    rvb := (*vb)[isize:]
    dsize,d,err := DeserializeVariableBlob(&rvb)
    if err != nil {
        return 0, &omh, err
    }
    omh.Id = UInt64(id)
    omh.Digest = *d
    return uint64(isize)+dsize, &omh, nil
}

// --------------------------------
//  MultihashVector
// --------------------------------

type MultihashVector struct {
    Id UInt64
    Digests []VariableBlob
}

func NewMultihashVector() *MultihashVector {
    o := MultihashVector{}
    o.Id = UInt64(0)
    o.Digests = make([]VariableBlob, 0)
    return &o
}

func (n *MultihashVector) Serialize(vb *VariableBlob) *VariableBlob {
    vb = EncodeVarint(vb, uint64(n.Id))
    size := uint64(0)
    if len(n.Digests) > 0 {
        size = uint64(len(n.Digests[0]))
    }
    vb = EncodeVarint(vb, size)
    vb = EncodeVarint(vb, uint64(len(n.Digests)))

    for _, item := range n.Digests {
        if uint64(len(item)) != size {
            panic("Multihash vector size mismatch")
        }
        *vb = append(*vb, item...)
    }

    return vb
}

func DeserializeMultihashVector(vb *VariableBlob) (uint64,*MultihashVector,error) {
    omv := MultihashVector{}
    id,i := binary.Uvarint(*vb)
    if i <= 0 {
        return 0, &omv, errors.New("Could not deserialize multihash vector id")
    }
    omv.Id = UInt64(id)
    size,j := binary.Uvarint((*vb)[i:])
    if j <= 0 {
        return 0, &omv, errors.New("Could not deserialize multihash vector hash size")
    }
    i += j
    entries,j := binary.Uvarint((*vb)[i:])
    if j <= 0 {
        return 0, &omv, errors.New("Could not deserialize multihash vector size")
    }
    i += j

    if len(*vb) < i + int(entries) * int(size) {
        return 0, &omv, errors.New("Unexpected EOF")
    }

    for num := uint64(0); num < entries; num++ {
        if uint64(len((*vb)[i:i+int(size)])) != size {
            return 0, &omv, errors.New("Multihash vector size mismatch")
        }
        omv.Digests = append(omv.Digests, (*vb)[i:i+int(size)])
        i += int(size)
    }

    return uint64(i), &omv, nil
}

func (n *MultihashVector) MarshalJSON() ([]byte, error) {
    mhv := struct {
        Id uint64 `json:"hash"`
        Digests []string `json:"digests"`
    }{Id: uint64(n.Id)}
    size := uint64(0)
    if len(n.Digests) > 0 {
        size = uint64(len(n.Digests[0]))
    }
    for _,item := range n.Digests {
        if uint64(len(item)) != size {
            panic("Multihash vector size mismatch")
        }
        mhv.Digests = append(mhv.Digests, EncodeBytes(item))
    }

    return json.Marshal(&mhv)
}

func (n *MultihashVector) UnmarshalJSON(b []byte) error {
    mhv := struct {
        Id uint64 `json:"hash"`
        Digests []string `json:"digests"`
    }{}

    err := json.Unmarshal(b, &mhv)
    if err != nil {
        return err
    }

    n.Id = UInt64(mhv.Id)
    size := uint64(0)
    for i, item := range mhv.Digests {
        db,err := DecodeBytes(item)
        if err != nil {
            return err
        }
        if i == 0 {
            size = uint64(len(db))
        } else {
            if uint64(len(db)) != size {
                return errors.New("Multihash vector size mismatch")
            }
        }
        n.Digests = append(n.Digests, VariableBlob(db))
    }

    return nil
}

// --------------------------------
//  Utility Functions
// --------------------------------

func EncodeBytes(b []byte) string {
    return "z" + base58.Encode(b)
}

func DecodeBytes(s string) ([]byte,error) {
    if len(s) <= 1 {
        return make([]byte, 0),nil
    }

    switch s[0] {
    case 'z':
        return base58.Decode(s[1:]),nil
    default:
        return nil,errors.New("Unknown encoding: " + string(s[0]))
    }
}

func SerializeBigInt(num *big.Int, byte_size int, signed bool) *VariableBlob {
    v := VariableBlob(make([]byte, byte_size))

    if signed && num.Sign() == -1 {
        x := big.NewInt(1)
        x = x.Add(x, num)
        v = x.FillBytes(v)
        for i := 0; i < byte_size; i++ {
            v[i] = ^v[i]
        }
        return &v
    }

    v = num.FillBytes(v)
    return &v
}

func DeserializeBigInt(vb *VariableBlob, byte_size int, signed bool) (*big.Int,error) {
    num := new(big.Int)

    if len(*vb) < byte_size {
        return num, errors.New("Unexpected EOF")
    }

    if signed && (0x80 & (*vb)[0]) == 0x80 {
        v := VariableBlob(make([]byte, byte_size))
        for i := 0; i < byte_size; i++ {
            v[i] = ^((*vb)[i])
        }
        neg := big.NewInt(-1)
        return num.SetBytes(v).Mul(neg, num).Add(neg, num), nil
    }

    return num.SetBytes((*vb)[:byte_size]), nil
}

func EncodeVarint(vb* VariableBlob, value uint64) *VariableBlob {
    header := make([]byte, binary.MaxVarintLen64)
    bytes := binary.PutUvarint(header, value)
    *vb = append(*vb, header[:bytes]...)
    return vb
}
