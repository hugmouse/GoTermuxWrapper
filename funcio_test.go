package gotermuxwrapper

import (
	"testing"
)

var (
	Title = "test"
	Hint  = "Press \"Yes\" then \"No\""
)

func TestTermuxDialog(t *testing.T) {

	resultYes := TermuxDialog(Title)
	if resultYes.Code != -1 || len(resultYes.Text) != 0 {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"-1, \"\" \".", resultYes.Code, resultYes.Text)
	}

	resultNo := TermuxDialog(Title)
	if resultNo.Code != -2 || len(resultNo.Text) != 0 {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"-2, \"\" \".", resultNo.Code, resultNo.Text)
	}

}

func TestTermuxDialogConfirm(t *testing.T) {

	resultYes := TermuxDialogConfirm(TDialogConfirm{
		Hint,
		TDialog{Title},
	})
	if resultYes.Code != 0 || resultYes.Text != "yes" {
		t.Errorf("TermuxDialogConfirm() was incorrect, got: \"%d, %s\" want: \"0, \"yes\" \".", resultYes.Code, resultYes.Text)
	}

	resultNo := TermuxDialogConfirm(TDialogConfirm{
		Hint,
		TDialog{Title},
	})
	if resultNo.Code != 0 || resultNo.Text != "no" {
		t.Errorf("TermuxDialogConfirm() was incorrect, got: \"%d, %s\" want: \"0, \"no\" \".", resultNo.Code, resultNo.Text)
	}

}

func TestTermuxDialogCheckbox(t *testing.T) {

	resultYes := TermuxDialogCheckbox(TDialogCheckbox{
		[]string{"value"},
		TDialog{Title},
	})
	if resultYes.Values[0].Index != 0 || resultYes.Values[0].Text != "value" {
		t.Errorf("TermuxDialogCheckbox() was incorrect, got: \"%d, %s\" want: \"0, \"value\" \".", resultYes.Values[0].Index, resultYes.Values[0].Text)
	}

}

func TestTermuxDialogCounter(t *testing.T) {
	tests := []struct {
		Min          int
		Max          int
		Start        int
		WantedCode   int8
		WantedString string
	}{
		{0, 2, 1, -1, "1"},
		{0, 2, 2, -1, "2"},
		{0, 2, 0, -1, "0"},
		{0, 2, 0, -2, ""},
	}

	for _, test := range tests {
		result := TermuxDialogCounter(TDialogCounter{
			test.Min,
			test.Max,
			test.Start,
			TDialog{"Just press \"ok\". On the 4th test press \"Cancel\""},
		})
		if result.Code != test.WantedCode || result.Text != test.WantedString {
			t.Errorf("TermuxDialogCounter() was incorrect, got: \"%d, %s\". Want: \"%d, \"%s\" \".", result.Code, result.Text, test.WantedCode, test.WantedString)
		}
	}

}

// TODO: Rewrite this function due bad code of this function in API
func TestTermuxDialogDate(t *testing.T) {
	tests := []struct {
		Day        uint
		Month      uint
		Year       uint
		Khour      uint
		Minutes    uint
		Seconds    uint
		WantedText string
		WantedCode int8
	}{
		// Here. "WantedText".
		{01, 01, 2000, 12, 00, 00, "1-1-2000 12:0:0", -1},
		{1, 1, 2000, 12, 02, 02, "1-1-2000 12:2:2", -1},
	}

	for _, test := range tests {
		result := TermuxDialogDate(TDialogDate{
			TDialogDatePattern{
				test.Day,
				test.Month,
				test.Year,
				test.Khour,
				test.Minutes,
				test.Seconds,
			},
			TDialog{"Just press \"OK\""},
		})
		if result.Code != test.WantedCode || result.Text != test.WantedText {
			t.Errorf("TermuxDialogDate() was incorrect, got: \"%d, %s\". Want: \"%d, \"%s\" \"", result.Code, result.Text, test.WantedCode, test.WantedText)
		}
	}
}

// TODO: do time.Now() and make format like in Termux output.
//func TestTermuxDialogDateWithoutDate(t *testing.T) {
//	tests := []struct{
//		WantedText string
//		WantedCode int8
//	} {
//
//	}
//
//	for _, test := range tests {
//		result := TermuxDialogWithoutDate(TDialog{"Just press \"OK\""})
//		if result.Code != test.WantedCode || result.Text != test.WantedText {
//			t.Errorf("TermuxDialogDate() was incorrect, got: \"%d, %s\". Want: \"%d, \"%s\" \"", result.Code, result.Text, test.WantedCode, test.WantedText)
//		}
//	}
//}

// Because they're the same
func TestTermuxDialogRadioSheetSpinner(t *testing.T) {
	tests := []struct {
		Value       []string
		WantedText  string
		WantedCode  int8
		WantedIndex uint
	}{
		{[]string{"Check me!"}, "Check me!", -1, 0},
		{[]string{"Do NOT check me!"}, "Do NOT check me!", -1, 0},
	}
	for _, test := range tests {
		resultRadio := TermuxDialogRadio(TDialogRadio{TDialogCheckbox{
			test.Value,
			TDialog{"Read carefully"},
		}})

		resultSheet := TermuxDialogSheet(TDialogSheet{TDialogCheckbox{
			test.Value,
			TDialog{"Read carefully"},
		}})

		resultSpinner := TermuxDialogSpinner(TDialogSpinner{TDialogCheckbox{
			test.Value,
			TDialog{"Read carefully"},
		}})

		if resultRadio.Code != test.WantedCode || resultRadio.Text != test.WantedText {
			t.Errorf("TermuxDialogRadio() was incorrect, got: \"%d, %s\". Want: \"%d, \"%s\" \".", resultRadio.Code, resultRadio.Text, test.WantedCode, test.WantedText)
		}

		// How funny is that? Thanks for absolutely different result code, Termux API!
		if resultSheet.Code != 0 || resultSheet.Text != test.WantedText {
			t.Errorf("TermuxDialogSheet() was incorrect, got: \"%d, %s\". Want: \"%d, \"%s\" \".", resultSheet.Code, resultSheet.Text, 0, test.WantedText)
		}

		if resultSpinner.Code != test.WantedCode || resultSpinner.Text != test.WantedText {
			t.Errorf("TermuxDialogSpinner() was incorrect, got: \"%d, %s\". Want: \"%d, \"%s\" \".", resultSpinner.Code, resultSpinner.Text, test.WantedCode, test.WantedText)
		}
	}
}
