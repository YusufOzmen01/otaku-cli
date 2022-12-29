package mpv

type Progress struct {
	Time    int
	Paused  bool
	Loading bool
	Length  int
}
