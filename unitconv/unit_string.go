// Code generated by "stringer -type=Unit -linecomment"; DO NOT EDIT.

package unitconv

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Wei-0]
	_ = x[KWei-1]
	_ = x[MWei-2]
	_ = x[GWei-3]
	_ = x[Szabo-4]
	_ = x[Finney-5]
	_ = x[Ether-6]
	_ = x[KEther-7]
	_ = x[MEther-8]
	_ = x[GEther-9]
	_ = x[TEther-10]
}

const _Unit_name = "WeiKWeiMWeiGWeiSzaboFinneyEtherKEtherMEtherGEtherTEther"

var _Unit_index = [...]uint8{0, 3, 7, 11, 15, 20, 26, 31, 37, 43, 49, 55}

func (i Unit) String() string {
	if i >= Unit(len(_Unit_index)-1) {
		return "Unit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Unit_name[_Unit_index[i]:_Unit_index[i+1]]
}
