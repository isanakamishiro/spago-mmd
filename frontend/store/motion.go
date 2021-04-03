package store

import "app/lib/threejs/animation"

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

var (
	_motions = []string{
		"",
		"./assets/models/mmd/vmds/wavefile_v2.vmd",
		"./assets/models/mmd/vmds/みんなみっくみくにしてあげる(Lat式).vmd",
		"./assets/models/mmd/vmds/ダブルラリアット.vmd",
	}
)

// CurrentMotion is ..
var CurrentMotion Motion = Dance1
var MotionDictionay map[Motion]animation.Clip = make(map[Motion]animation.Clip)

// Path gets MMD motion file path.
func (c Motion) Path() string {

	return _motions[c]
}
