package quickPerm_test

import (
	"testing"
	"fmt"

	"github.com/Ramshackle-Jamathon/quickPerm"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    for permutation := range quickPerm.GeneratePermutationsString([]string{"a","b","c","d","e","f","g","h"}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsInt([]int{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsInt8([]int8{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsInt16([]int16{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsInt32([]int32{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsInt64([]int64{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsUint([]uint{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsUint8([]uint8{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsUint16([]uint16{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsUint32([]uint32{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsUint64([]uint64{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsFloat32([]float32{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsFloat64([]float64{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsComplex64([]complex64{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
    for permutation := range quickPerm.GeneratePermutationsComplex128([]complex128{1,2,3,4,6,7,8}) {
        fmt.Println(permutation)
        assert.NotNil(t, permutation)
    }
}