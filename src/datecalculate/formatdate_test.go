package datecalculate

func TestFormatDateInput412018ShouldBeThurday4January2018(t *testing.T)
{
	expected := "Thurday, 4 Januaray 2018"
	startDate := "04012018"
	actualDate := FormatDate(startDate)
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actualDate)
	}
}
