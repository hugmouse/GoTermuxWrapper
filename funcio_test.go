package gotermuxwrapper

import (
	"testing"
)

var (
	Title = "test"
	Hint = "Press \"Yes\" then \"No\""
	TestConfirm = TDialogConfirm{
		Hint,
		TDialog{Title},
	}
	TestCheckbox = TDialogCheckbox{
		[]string{"Value"},
		TDialog{Title},
	}
)

func TestTermuxDialog(t *testing.T) {

	resultYES := TermuxDialog(Title)
	if resultYES.Code != -1 || len(resultYES.Text) != 0 {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"-1, \"\" \".", resultYES.Code, resultYES.Text)
	}

	resultNO := TermuxDialog(Title)
	if resultNO.Code != -2 || len(resultNO.Text) != 0 {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"-2, \"\" \".", resultNO.Code, resultNO.Text)
	}

}

func TestTermuxDialogConfirm(t *testing.T) {

	resultYES := TermuxDialogConfirm(TestConfirm)
	if resultYES.Code != 0 || resultYES.Text != "yes" {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"0, \"yes\" \".", resultYES.Code, resultYES.Text)
	}

	resultNO := TermuxDialogConfirm(TestConfirm)
	if resultNO.Code != 0 || resultNO.Text != "no" {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"0, \"no\" \".", resultNO.Code, resultNO.Text)
	}

}

func TestTermuxDialogCheckbox(t *testing.T) {
	resultYES := TermuxDialogCheckbox(TestCheckbox)
	if resultYES.Code != -1 || resultYES.Text != "[value]" {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"-1, \"[value]\" \".", resultYES.Code, resultYES.Text)
	}
	if resultYES.Values[0].Index != 0 || resultYES.Values[0].Text != "value" {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"0, \"value\" \".", resultYES.Values[0].Index, resultYES.Values[0].Text)
	}

	resultNO := TermuxDialogCheckbox(TestCheckbox)
	if resultNO.Code != -2 || len(resultNO.Text) != 0 {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"-2, \"\" \".", resultNO.Code, resultNO.Text)
	}
}
