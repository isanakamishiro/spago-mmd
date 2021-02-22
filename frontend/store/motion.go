package store

// Motion is MMD Model animation number.
type Motion int

const (
	// Dance1 is ...
	Dance1 Motion = iota + 1
	// Dance2 is ...
	Dance2
	// Dance3 is ...
	Dance3
)

// CurrentMotion is ..
var CurrentMotion Motion = Dance1

// Path gets MMD motion file path.
func (c Motion) Path() string {

	switch c {
	case Dance1:
		return "./assets/models/mmd/vmds/wavefile_v2.vmd"
	case Dance2:
		return "./assets/models/mmd/vmds/ハッピーシンセサイザモーション.vmd"
	case Dance3:
		return "./assets/models/mmd/vmds/wavefile_v2.vmd"
	default:
		return ""
	}

}
