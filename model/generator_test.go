package model

import (
	"testing"
	"strconv"
	"github.com/stretchr/testify/require"
	"github.com/Phadoo131/hahalanguage/util"
)

var test *Hahalanguage

func TestHahainTh(t *testing.T){
	result := test.HahahaInTH()

	require.NotEmpty(t, result)
	require.Equal(t, 555, result)
}

func TestHahaThis(t *testing.T){
	randByte := util.RandomByte()
	randInt := util.RandomInt(12345, 123456)
	randString := util.RandomString(7) + "!"

	randAryByte := []byte{util.RandomByte(), util.RandomByte(), util.RandomByte(), util.RandomByte()}
	randAryInt := []int{util.RandomInt(12345, 123456), util.RandomInt(12345, 123456), util.RandomInt(12345, 123456)}
	randAryStr := []string{util.RandomString(3), util.RandomString(3), util.RandomString(3), util.RandomString(3), "!"}

	result1 := test.HahaThis(randByte)
	require.NotEmpty(t, result1)
	require.Equal(t, "5", result1)

	result2 := test.HahaThis(randInt)
	require.NotEmpty(t, result2)
	require.Equal(t, len(strconv.Itoa(randInt)), len(result2))

	result3 := test.HahaThis(randString)
	require.NotEmpty(t, result3)
	require.LessOrEqual(t, len(randString), len(result3))

	result4 := test.HahaThis(randAryByte)
	require.NotEmpty(t, result4)
	require.LessOrEqual(t, len(randAryByte), len(result4))

	result5 := test.HahaThis(randAryInt)
	require.NotEmpty(t, result5)
	require.Equal(t, len(randAryInt), len(result5))

	result6 := test.HahaThis(randAryStr)
	require.NotEmpty(t, result6)
	require.LessOrEqual(t, len(randAryStr), len(result6))
}