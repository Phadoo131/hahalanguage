package model

import (
	"testing"

	"github.com/Phadoo131/hahalanguage/util"
	"github.com/stretchr/testify/require"
)

func TestFunOrFalse(t *testing.T) {
	require := require.New(t)
	
	 // Test with int
	result := test.FunOrFalse(5)
	require.True(result)
	  
	// Test with rune
	result = test.FunOrFalse(rune('5'))
	require.True(result)
	
	// Test with byte
	result = test.FunOrFalse(byte('5'))
	require.True(result)
	
	// Test with string
	result = test.FunOrFalse("Hello, World!")
	require.False(result)
	
	// Test with []string
	result = test.FunOrFalse([]string{"hello", "5"})
	require.True(result)
	
	// Test with unsupported type
	result = test.FunOrFalse(true)
	require.False(result)
	
	// Test with random byte
	randByte := util.RandomByte()
	result = test.FunOrFalse(randByte)
	require.NotEmpty(t, result)
	
	    // Test with random int
	randInt := util.RandomInt(12345, 123456)
	result = test.FunOrFalse(randInt)
	require.NotEmpty(t, result)
	
	// Test with random string
	randString := util.RandomString(7) + "!"
	result = test.FunOrFalse(randString)
	require.NotEmpty(t, result)
}

func TestWhereIsFun(t *testing.T) {
	h := Hahalanguage{}

	require.Equal(t, "The fun is at the index: 0, 1", h.WhereIsFun("ha"))
	require.NotEqual(t, "The fun is at the index: 1, 2", h.WhereIsFun("lol"))

	require.Equal(t, "The fun is at the index: No, there's no fun in this", h.WhereIsFun("hello"))
	require.Equal(t, "The fun is at the index: No, there's no fun in this", h.WhereIsFun("world"))
}
