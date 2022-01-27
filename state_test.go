package state

import (
	"testing"
)

func TestViewState(t *testing.T) {
	wanter := "[ Kylling Rev Korn ---V \\ \\_HS_/ _____________/ Ø---]"
	state := ViewState()
	if state != wanted {
		t.errorf("Feil, fikk %q, ønsket %q, state, wanted")
	}
}