package store

// Model is MMD Model Number.
type Model int

const (
	// Diluc is ...
	Diluc Model = iota + 1
	// Lisa is ...
	Lisa
	// Miku is ...
	Miku
)

// CurrentModel is ...
var CurrentModel Model = Diluc

// Path gets MMD model file path.
func (c Model) Path() string {

	switch c {
	case Diluc:
		return "./assets/models/mmd/diluc/diluc.pmx"
	case Lisa:
		return "./assets/models/mmd/lisa/lisa.pmx"
	case Miku:
		return "./assets/models/mmd/miku/miku_v2.pmd"
	default:
		return ""
	}

}
