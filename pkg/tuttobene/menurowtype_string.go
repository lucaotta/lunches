// Code generated by "stringer -type=MenuRowType"; DO NOT EDIT.

package tuttobene

import "strconv"

const _MenuRowType_name = "UnknonwnPrimoSecondoContornoVegetarianoFruttaPanino"

var _MenuRowType_index = [...]uint8{0, 8, 13, 20, 28, 39, 45, 51}

func (i MenuRowType) String() string {
	if i < 0 || i >= MenuRowType(len(_MenuRowType_index)-1) {
		return "MenuRowType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MenuRowType_name[_MenuRowType_index[i]:_MenuRowType_index[i+1]]
}
