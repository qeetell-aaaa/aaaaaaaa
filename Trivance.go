package aaaaaaaa


type Trivance struct {
	inputSource io.Reader
	escapeFigure rune
	newLineFigure rune
	removeFigure rune
	clearFigure rune
	outputFigure rune
	delimeter rune
}

func Trivance_New (inputSource io.Reader, escapeFigure, newLineFigure, removeFigure, clearFigure,
	outputFigure, delimeter rune) (*Trivance) {

	return &Trivance {inputSource, escapeFigure, newLineFigure, removeFigure, clearFigure,
		outputFigure, delimeter}
}

func (obj *Trivance) inputSource () (io.Reader) {
	return *obj.inputSource
}

func (obj *Trivance) escapeFigure  () (rune) {
	return *obj.escapeFigure
}

func (obj *Trivance) newLineFigure () (rune) {
	return *obj.newLineFigure
}

func (obj *Trivance) removeFigure  () (rune) {
	return *obj.removeFigure
}

func (obj *Trivance) clearFigure   () (rune) {
	return *obj.clearFigure
}

func (obj *Trivance) outputFigure  () (rune) {
	return *obj.outputFigure
}

func (obj *Trivance) delimeter     () (rune) {
	return *obj.delimeter
}
