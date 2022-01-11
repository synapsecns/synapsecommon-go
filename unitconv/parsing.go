package unitconv

import (
	"math"
	"math/big"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	"github.com/synapsecns/synapsecommon-go/bigint"
)

const (
	EtherDecimalPrecision uint = 18
	roundingMode               = big.ToNearestEven
)

func ToEther(amt interface{}, fromUnit Unit) (*big.Float, error) {
	if !fromUnit.valid() {
		return nil, errors.New("unknown Unit passed as param fromUnit")
	}

	var res *big.Float

	fromFloat := func(amt *big.Float) *big.Float {
		if fromUnit == Ether {
			return amt
		}

		return newEther().
			Mul(amt, makePow10(etherPowMap[fromUnit]))
	}

	fromInt := func(amt *big.Int) *big.Float {
		return fromFloat(newEther().SetInt(amt))
	}

	if s, ok := amt.(string); ok {
		return fromFloat(etherFromString(s)), nil
	}

	bigObj, err := makeBigObjFromInterface(amt)
	if err != nil {
		return nil, err
	}

	switch a := bigObj.(type) {
	case *big.Float:
		res = fromFloat(a)
	case *big.Int:
		res = fromInt(a)
	}

	return res, nil
}

func ToWei(amt interface{}, fromUnit Unit) (*big.Int, error) {
	if !fromUnit.valid() {
		return nil, errors.New("unknown Unit passed as param fromUnit")
	}

	var res *big.Int

	fromFloat := func(amt *big.Float) *big.Int {
		return parseWei(amt.Text('f', -1), fromUnit)
	}

	fromInt := func(amt *big.Int) *big.Int {
		if fromUnit == Wei {
			return amt
		}
		return parseWei(amt.String(), fromUnit)
	}

	if s, ok := amt.(string); ok {
		return parseWei(s, fromUnit), nil
	}

	bigObj, err := makeBigObjFromInterface(amt)
	if err != nil {
		return nil, err
	}

	switch a := bigObj.(type) {
	case *big.Float:
		res = fromFloat(a)
	case *big.Int:
		res = fromInt(a)
	}

	return res, nil
}

func parseWei(amt string, fromUnit Unit) *big.Int {
	baseWeiAmt := new(big.Int)
	makePow10(fromUnit.weiExponent()).Int(baseWeiAmt)

	baseLen := len(baseWeiAmt.String()) - 1

	isNegative := strings.HasPrefix(amt, "-")
	if isNegative {
		amt = strings.TrimPrefix(amt, "-")
	}

	var fracNumStr string

	splitAmt := strings.SplitN(amt, ".", 2)
	switch len(splitAmt) {
	case 1:
		fracNumStr = "0"
	case 2:
		fracNumStr = splitAmt[1]
	}

	wholeNum, ok := new(big.Int).SetString(splitAmt[0], 10)
	if !ok {
		panic(errors.Errorf("error parsing %[1]s to big.Int", splitAmt[0]))
	}

	for len(fracNumStr) < baseLen {
		fracNumStr += "0"
	}

	fracNum, ok := new(big.Int).SetString(fracNumStr, 10)
	if !ok {
		panic(errors.Errorf("error parsing %[1]s to big.Int", fracNumStr))
	}

	res := bigint.Mul(wholeNum, baseWeiAmt)
	res = bigint.Add(res, fracNum)

	if isNegative {
		res = res.Neg(res)
	}

	return res
}

func makeBigObjFromInterface(amt interface{}) (interface{}, error) {
	switch a := amt.(type) {
	case string:
		return nil, nil // handle case-by-case
	case *big.Float:
		return a, nil
	case float32, float64:
		n := reflect.ValueOf(a).Float()
		return newEther().SetFloat64(n), nil
	case *big.Int:
		return a, nil
	case int, int8, int16, int32, int64:
		n := reflect.ValueOf(a).Int()
		return bigint.FromInt64(n), nil
	case uint, uint8, uint16, uint32, uint64:
		n := reflect.ValueOf(a).Uint()
		return bigint.FromUint64(n), nil
	default:
		return nil, errors.Errorf("unsupported type %[1]T passed as parameter amt", amt)
	}
}

func makePow10(exp int) *big.Float {
	return new(big.Float).SetFloat64(math.Pow10(exp))
}

func etherFromString(s string) *big.Float {
	if !strings.Contains(s, ".") {
		s += ".0"
	}

	ether, _, err := newEther().Parse(s, 10)
	if err != nil {
		panic(errors.WithStack(errors.Wrapf(err, "error parsing string '%[1]s' to big.Float", s)))
	}

	return ether
}

func newEther() *big.Float {
	return new(big.Float).SetMode(roundingMode).SetPrec(18)
}
