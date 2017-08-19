package column_test

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/kshvakov/clickhouse/lib/binary"
	columns "github.com/kshvakov/clickhouse/lib/column"
	"github.com/stretchr/testify/assert"
)

func Test_Column_Int8(t *testing.T) {
	var (
		buf     bytes.Buffer
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "Int8", time.Local); assert.NoError(t, err) {
		for i := -128; i <= 127; i++ {
			if err := column.Write(encoder, int8(i)); assert.NoError(t, err) {
				if v, err := column.Read(decoder); assert.NoError(t, err) {
					assert.Equal(t, int8(i), v)
				}
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "Int8", column.CHType()) {
			assert.Equal(t, reflect.Int8, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int16(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int16(0), e.T)
			}
		}
	}
}

func Test_Column_Int16(t *testing.T) {
	var (
		buf     bytes.Buffer
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "Int16", time.Local); assert.NoError(t, err) {
		for i := -32768; i <= 32767; i++ {
			if err := column.Write(encoder, int16(i)); assert.NoError(t, err) {
				if v, err := column.Read(decoder); assert.NoError(t, err) {
					assert.Equal(t, int16(i), v)
				}
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "Int16", column.CHType()) {
			assert.Equal(t, reflect.Int16, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int8(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int8(0), e.T)
			}
		}
	}
}

func Test_Column_Int32(t *testing.T) {
	var (
		buf     bytes.Buffer
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "Int32", time.Local); assert.NoError(t, err) {
		for i := -2147483648; i <= 2147483648; i += 100000 {
			if err := column.Write(encoder, int32(i)); assert.NoError(t, err) {
				if v, err := column.Read(decoder); assert.NoError(t, err) {
					assert.Equal(t, int32(i), v)
				}
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "Int32", column.CHType()) {
			assert.Equal(t, reflect.Int32, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int8(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int8(0), e.T)
			}
		}
	}
}

func Test_Column_Int64(t *testing.T) {
	var (
		buf     bytes.Buffer
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "Int64", time.Local); assert.NoError(t, err) {
		for i := -2147483648; i <= 2147483648*2; i += 100000 {
			if err := column.Write(encoder, int64(i)); assert.NoError(t, err) {
				if v, err := column.Read(decoder); assert.NoError(t, err) {
					assert.Equal(t, int64(i), v)
				}
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "Int64", column.CHType()) {
			assert.Equal(t, reflect.Int64, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int8(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int8(0), e.T)
			}
		}
	}
}

func Test_Column_UInt8(t *testing.T) {
	var (
		buf     bytes.Buffer
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "UInt8", time.Local); assert.NoError(t, err) {
		for i := 0; i <= 255; i++ {
			if err := column.Write(encoder, uint8(i)); assert.NoError(t, err) {
				if v, err := column.Read(decoder); assert.NoError(t, err) {
					assert.Equal(t, uint8(i), v)
				}
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "UInt8", column.CHType()) {
			assert.Equal(t, reflect.Uint8, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int16(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int16(0), e.T)
			}
		}
	}
}

func Test_Column_UInt16(t *testing.T) {
	var (
		buf     bytes.Buffer
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "UInt16", time.Local); assert.NoError(t, err) {
		for i := 0; i <= 65535; i++ {
			if err := column.Write(encoder, uint16(i)); assert.NoError(t, err) {
				if v, err := column.Read(decoder); assert.NoError(t, err) {
					assert.Equal(t, uint16(i), v)
				}
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "UInt16", column.CHType()) {
			assert.Equal(t, reflect.Uint16, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int8(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int8(0), e.T)
			}
		}
	}
}

func Test_Column_UInt32(t *testing.T) {
	var (
		buf     bytes.Buffer
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "UInt32", time.Local); assert.NoError(t, err) {
		for i := 0; i <= 4294967295; i += 100000 {
			if err := column.Write(encoder, uint32(i)); assert.NoError(t, err) {
				if v, err := column.Read(decoder); assert.NoError(t, err) {
					assert.Equal(t, uint32(i), v)
				}
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "UInt32", column.CHType()) {
			assert.Equal(t, reflect.Uint32, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int8(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int8(0), e.T)
			}
		}
	}
}

func Test_Column_UInt64(t *testing.T) {
	var (
		buf     bytes.Buffer
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "UInt64", time.Local); assert.NoError(t, err) {
		for i := 0; i <= 4294967295*2; i += 100000 {
			if err := column.Write(encoder, uint64(i)); assert.NoError(t, err) {
				if v, err := column.Read(decoder); assert.NoError(t, err) {
					assert.Equal(t, uint64(i), v)
				}
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "UInt64", column.CHType()) {
			assert.Equal(t, reflect.Uint64, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int8(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int8(0), e.T)
			}
		}
	}
}

func Test_Column_Float32(t *testing.T) {
	var (
		buf     bytes.Buffer
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "Float32", time.Local); assert.NoError(t, err) {
		for i := -2147483648; i <= 2147483648; i += 100000 {
			if err := column.Write(encoder, float32(i)); assert.NoError(t, err) {
				if v, err := column.Read(decoder); assert.NoError(t, err) {
					assert.Equal(t, float32(i), v)
				}
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "Float32", column.CHType()) {
			assert.Equal(t, reflect.Float32, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int8(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int8(0), e.T)
			}
		}
	}
}

func Test_Column_Float64(t *testing.T) {
	var (
		buf     bytes.Buffer
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "Float64", time.Local); assert.NoError(t, err) {
		for i := -2147483648; i <= 2147483648*2; i += 100000 {
			if err := column.Write(encoder, float64(i)); assert.NoError(t, err) {
				if v, err := column.Read(decoder); assert.NoError(t, err) {
					assert.Equal(t, float64(i), v)
				}
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "Float64", column.CHType()) {
			assert.Equal(t, reflect.Float64, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int8(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int8(0), e.T)
			}
		}
	}
}

func Test_Column_String(t *testing.T) {
	var (
		buf     bytes.Buffer
		str     = fmt.Sprintf("str_%d", time.Now().Unix())
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "String", time.Local); assert.NoError(t, err) {
		if err := column.Write(encoder, str); assert.NoError(t, err) {
			if v, err := column.Read(decoder); assert.NoError(t, err) {
				assert.Equal(t, str, v)
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "String", column.CHType()) {
			assert.Equal(t, reflect.String, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int8(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int8(0), e.T)
			}
		}
	}
}

func Test_Column_FixedString(t *testing.T) {
	var (
		buf     bytes.Buffer
		str     = fmt.Sprintf("str_%d", time.Now().Unix())
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "FixedString(14)", time.Local); assert.NoError(t, err) {
		if err := column.Write(encoder, str); assert.NoError(t, err) {
			if v, err := column.Read(decoder); assert.NoError(t, err) {
				assert.Equal(t, str, v)
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "FixedString(14)", column.CHType()) {
			assert.Equal(t, reflect.String, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int8(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int8(0), e.T)
			}
		}
	}
}

func Test_Column_Enum8(t *testing.T) {
	var (
		buf     bytes.Buffer
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "Enum8('A'=1,'B'=2,'C'=3)", time.Local); assert.NoError(t, err) {
		if err := column.Write(encoder, "B"); assert.NoError(t, err) {
			if v, err := column.Read(decoder); assert.NoError(t, err) {
				assert.Equal(t, "B", v)
			}
		}
		if err := column.Write(encoder, int16(3)); assert.Error(t, err) {
			if err := column.Write(encoder, int8(3)); assert.NoError(t, err) {
				if v, err := column.Read(decoder); assert.NoError(t, err) {
					assert.Equal(t, "C", v)
				}
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "Enum8('A'=1,'B'=2,'C'=3)", column.CHType()) {
			assert.Equal(t, reflect.String, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int(0), e.T)
			}
		}
	}
}

func Test_Column_Enum16(t *testing.T) {
	var (
		buf     bytes.Buffer
		encoder = binary.NewEncoder(&buf)
		decoder = binary.NewDecoder(&buf)
	)
	if column, err := columns.Factory("column_name", "Enum16('A'=1,'B'=2,'C'=3)", time.Local); assert.NoError(t, err) {
		if err := column.Write(encoder, "B"); assert.NoError(t, err) {
			if v, err := column.Read(decoder); assert.NoError(t, err) {
				assert.Equal(t, "B", v)
			}
		}
		if err := column.Write(encoder, int8(3)); assert.Error(t, err) {
			if err := column.Write(encoder, int16(3)); assert.NoError(t, err) {
				if v, err := column.Read(decoder); assert.NoError(t, err) {
					assert.Equal(t, "C", v)
				}
			}
		}
		if assert.Equal(t, "column_name", column.Name()) && assert.Equal(t, "Enum16('A'=1,'B'=2,'C'=3)", column.CHType()) {
			assert.Equal(t, reflect.String, column.ScanType().Kind())
		}
		if err := column.Write(encoder, int(0)); assert.Error(t, err) {
			if e, ok := err.(*columns.ErrUnexpectedType); assert.True(t, ok) {
				assert.Equal(t, int(0), e.T)
			}
		}
	}
}
