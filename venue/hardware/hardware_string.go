// Code generated by "stringer -type=Hardware"; DO NOT EDIT

package hardware

import "fmt"

const _Hardware_name = "StageBoxLocalProTools"

var _Hardware_index = [...]uint8{0, 8, 13, 21}

func (i Hardware) String() string {
	if i < 0 || i >= Hardware(len(_Hardware_index)-1) {
		return fmt.Sprintf("Hardware(%d)", i)
	}
	return _Hardware_name[_Hardware_index[i]:_Hardware_index[i+1]]
}
