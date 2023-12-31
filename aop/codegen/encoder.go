package codegen

import (
	"encoding"
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"google.golang.org/protobuf/proto"
)

// encoderError is the type of error passed to panic by encoding code that encounters an error.
type encoderError struct {
	err error
}

func (e encoderError) Error() string {
	if e.err == nil {
		return "encoder:"
	}
	return "encoder: " + e.err.Error()
}

func (e encoderError) Unwrap() error {
	return e.err
}

// Encoder serializes data in a byte slice data.
type Encoder struct {
	data  []byte    // Contains the serialized arguments.
	space [100]byte // Prellocated buffer to avoid allocations for small size arguments.
}

func NewEncoder() *Encoder {
	var enc Encoder
	enc.data = enc.space[:0] // Arrange to use builtin buffer
	return &enc
}

// Reset resets the Encoder to use a buffer with a capacity of at least the
// provided size. All encoded data is lost.
func (e *Encoder) Reset(n int) {
	// TODO(mwhittaker): Have a NewEncoder method that takes in an initial
	// buffer? Or at least an initial capacity? And then pipe that through
	// NewCaller.
	if n <= cap(e.data) {
		e.data = e.data[:0]
	} else {
		e.data = make([]byte, 0, n)
	}
}

// makeEncodeError creates and returns an encoder error.
func makeEncodeError(format string, args ...interface{}) encoderError {
	return encoderError{fmt.Errorf(format, args...)}
}

// EncodeProto serializes value into a byte slice using proto serialization.
func (e *Encoder) EncodeProto(value proto.Message) {
	enc, err := proto.Marshal(value)
	if err != nil {
		panic(makeEncodeError("error encoding to proto %T: %w", value, err))
	}
	e.Bytes(enc)
}

// EncodeBinaryMarshaler serializes value into a byte slice using its
// MarshalBinary method.
func (e *Encoder) EncodeBinaryMarshaler(value encoding.BinaryMarshaler) {
	enc, err := value.MarshalBinary()
	if err != nil {
		panic(makeEncodeError("error encoding BinaryMarshaler %T: %w", value, err))
	}
	e.Bytes(enc)
}

// Data returns the byte slice that contains the serialized arguments.
func (e *Encoder) Data() []byte {
	return e.data
}

// Grow increases the size of the encoder's data if needed. Only appends a new
// slice if there is not enough capacity to satisfy bytesNeeded.
// Returns the slice fragment that contains bytesNeeded.
func (e *Encoder) Grow(bytesNeeded int) []byte {
	n := len(e.data)
	if cap(e.data)-n >= bytesNeeded {
		e.data = e.data[:n+bytesNeeded] // Grow in place (common case)
	} else {
		// Create a new larger slice.
		e.data = append(e.data, make([]byte, bytesNeeded)...)
	}
	return e.data[n:]
}

// Uint8 encodes an arg of type uint8.
func (e *Encoder) Uint8(arg uint8) {
	e.Grow(1)[0] = arg
}

// Byte encodes an arg of type byte.
func (e *Encoder) Byte(arg byte) {
	e.Uint8(arg)
}

// Int8 encodes an arg of type int8.
func (e *Encoder) Int8(arg int8) {
	e.Uint8(byte(arg))
}

// Uint16 encodes an arg of type uint16.
func (e *Encoder) Uint16(arg uint16) {
	binary.LittleEndian.PutUint16(e.Grow(2), arg)
}

// Int16 encodes an arg of type int16.
func (e *Encoder) Int16(arg int16) {
	e.Uint16(uint16(arg))
}

// Uint32 encodes an arg of type uint32.
func (e *Encoder) Uint32(arg uint32) {
	binary.LittleEndian.PutUint32(e.Grow(4), arg)
}

// Int32 encodes an arg of type int32.
func (e *Encoder) Int32(arg int32) {
	e.Uint32(uint32(arg))
}

// Rune encodes an arg of type rune.
func (e *Encoder) Rune(arg rune) {
	e.Int32(arg)
}

// Uint64 encodes an arg of type uint64.
func (e *Encoder) Uint64(arg uint64) {
	binary.LittleEndian.PutUint64(e.Grow(8), arg)
}

// Int64 encodes an arg of type int64.
func (e *Encoder) Int64(arg int64) {
	e.Uint64(uint64(arg))
}

// Uint encodes an arg of type uint.
// Uint can have 32 bits or 64 bits based on the machine type. To simplify our
// reasoning, we encode the highest possible value.
func (e *Encoder) Uint(arg uint) {
	e.Uint64(uint64(arg))
}

// Int encodes an arg of type int.
// Int can have 32 bits or 64 bits based on the machine type. To simplify our
// reasoning, we encode the highest possible value.
func (e *Encoder) Int(arg int) {
	e.Uint64(uint64(arg))
}

// Bool encodes an arg of type bool.
// Serialize boolean values as an uint8 that encodes either 0 or 1.
func (e *Encoder) Bool(arg bool) {
	if arg {
		e.Uint8(1)
	} else {
		e.Uint8(0)
	}
}

// Float32 encodes an arg of type float32.
func (e *Encoder) Float32(arg float32) {
	binary.LittleEndian.PutUint32(e.Grow(4), math.Float32bits(arg))
}

// Float64 encodes an arg of type float64.
func (e *Encoder) Float64(arg float64) {
	binary.LittleEndian.PutUint64(e.Grow(8), math.Float64bits(arg))
}

// Complex64 encodes an arg of type complex64.
// We encode the real and the imaginary parts one after the other.
func (e *Encoder) Complex64(arg complex64) {
	e.Float32(real(arg))
	e.Float32(imag(arg))
}

// Complex128 encodes an arg of type complex128.
func (e *Encoder) Complex128(arg complex128) {
	e.Float64(real(arg))
	e.Float64(imag(arg))
}

// String encodes an arg of type string.
// For a string, we encode its length, followed by the serialized content.
func (e *Encoder) String(arg string) {
	n := len(arg)
	if n > math.MaxUint32 {
		panic(makeEncodeError("unable to encode string; length doesn't fit in 4 bytes"))
	}
	data := e.Grow(4 + n)
	binary.LittleEndian.PutUint32(data, uint32(n))
	copy(data[4:], arg)
}

// Bytes encodes an arg of type []byte.
// For a byte slice, we encode its length, followed by the serialized content.
// If the slice is nil, we encode length as -1.
func (e *Encoder) Bytes(arg []byte) {
	if arg == nil {
		e.Int32(-1)
		return
	}
	n := len(arg)
	if n > math.MaxUint32 {
		panic(makeEncodeError("unable to encode bytes; length doesn't fit in 4 bytes"))
	}
	data := e.Grow(4 + n)
	binary.LittleEndian.PutUint32(data, uint32(n))
	copy(data[4:], arg)
}

// Len attempts to encode l as an int32.
//
// Panics if l is bigger than an int32 or a negative length (except -1).
//
// NOTE that this method should be called only in the generated code, to avoid
// generating repetitive code that encodes the length of a non-basic type (e.g., slice, map).
func (e *Encoder) Len(l int) {
	if l < -1 {
		panic(makeEncodeError("unable to encode a negative length %d", l))
	}
	if l > math.MaxUint32 {
		panic(makeEncodeError("length can't be represented in 4 bytes"))
	}
	e.Int32(int32(l))
}

// Error encodes an arg of type error. We save enough type information
// to allow errors.Unwrap() and errors.Is() to work correctly.
func (e *Encoder) Error(err error) {
	// Get the stack of wrapped errors.
	stack := make([]error, 0, 4)
	for err != nil {
		stack = append(stack, err)
		err = errors.Unwrap(err)
	}

	e.Int(len(stack))
	for _, err := range stack {
		e.String(err.Error())
		e.String(fmtError(err))

		// TODO(sanjay): If a wrapped errors can be serialized using Gob, consider
		// saving that serialization. This may allow us to implement the As() method
		// that reconstructs an error value with the original type at the receiver.
	}
}
