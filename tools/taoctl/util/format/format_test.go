package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	list, err := split("A")
	assert.Nil(t, err)
	assert.Equal(t, []string{"A"}, list)

	list, err = split("goTao")
	assert.Nil(t, err)
	assert.Equal(t, []string{"go", "Tao"}, list)

	list, err = split("Gotao")
	assert.Nil(t, err)
	assert.Equal(t, []string{"Gotao"}, list)

	list, err = split("go_tao")
	assert.Nil(t, err)
	assert.Equal(t, []string{"go", "tao"}, list)

	list, err = split("talGo_tao")
	assert.Nil(t, err)
	assert.Equal(t, []string{"tal", "Go", "tao"}, list)

	list, err = split("GOTAO")
	assert.Nil(t, err)
	assert.Equal(t, []string{"G", "O", "T", "A", "O"}, list)

	list, err = split("gotao")
	assert.Nil(t, err)
	assert.Equal(t, []string{"gotao"}, list)

	list, err = split("")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(list))

	list, err = split("a_b_CD_EF")
	assert.Nil(t, err)
	assert.Equal(t, []string{"a", "b", "C", "D", "E", "F"}, list)

	list, err = split("_")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(list))

	list, err = split("__")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(list))

	list, err = split("_A")
	assert.Nil(t, err)
	assert.Equal(t, []string{"A"}, list)

	list, err = split("_A_")
	assert.Nil(t, err)
	assert.Equal(t, []string{"A"}, list)

	list, err = split("A_")
	assert.Nil(t, err)
	assert.Equal(t, []string{"A"}, list)

	list, err = split("welcome_to_go_tao")
	assert.Nil(t, err)
	assert.Equal(t, []string{"welcome", "to", "go", "tao"}, list)
}

func TestFileNamingFormat(t *testing.T) {
	testFileNamingFormat(t, "gotao", "welcome_to_go_tao", "welcometogotao")
	testFileNamingFormat(t, "_go#tao_", "welcome_to_go_tao", "_welcome#to#go#tao_")
	testFileNamingFormat(t, "Go#tao", "welcome_to_go_tao", "Welcome#to#go#tao")
	testFileNamingFormat(t, "Go#Tao", "welcome_to_go_tao", "Welcome#To#Go#Tao")
	testFileNamingFormat(t, "Go_Tao", "welcome_to_go_tao", "Welcome_To_Go_Tao")
	testFileNamingFormat(t, "go_Tao", "welcome_to_go_tao", "welcome_To_Go_Tao")
	testFileNamingFormat(t, "goTao", "welcome_to_go_tao", "welcomeToGoTao")
	testFileNamingFormat(t, "GoTao", "welcome_to_go_tao", "WelcomeToGoTao")
	testFileNamingFormat(t, "GOTao", "welcome_to_go_tao", "WELCOMEToGoTao")
	testFileNamingFormat(t, "GoTAO", "welcome_to_go_tao", "WelcomeTOGOTAO")
	testFileNamingFormat(t, "GOTAO", "welcome_to_go_tao", "WELCOMETOGOTAO")
	testFileNamingFormat(t, "GO*TAO", "welcome_to_go_tao", "WELCOME*TO*GO*TAO")
	testFileNamingFormat(t, "[GO#TAO]", "welcome_to_go_tao", "[WELCOME#TO#GO#TAO]")
	testFileNamingFormat(t, "{go###tao}", "welcome_to_go_tao", "{welcome###to###go###tao}")
	testFileNamingFormat(t, "{go###taogo_tao}", "welcome_to_go_tao", "{welcome###to###go###taogo_tao}")
	testFileNamingFormat(t, "GogoTaotao", "welcome_to_go_tao", "WelcomegoTogoGogoTaotao")
	testFileNamingFormat(t, "前缀GoTao后缀", "welcome_to_go_tao", "前缀WelcomeToGoTao后缀")
	testFileNamingFormat(t, "GoTao", "welcometogotao", "Welcometogotao")
	testFileNamingFormat(t, "GoTao", "WelcomeToGoTao", "WelcomeToGoTao")
	testFileNamingFormat(t, "gotao", "WelcomeToGoTao", "welcometogotao")
	testFileNamingFormat(t, "go_tao", "WelcomeToGoTao", "welcome_to_go_tao")
	testFileNamingFormat(t, "Go_Tao", "WelcomeToGoTao", "Welcome_To_Go_Tao")
	testFileNamingFormat(t, "Go_Tao", "", "")
	testFileNamingFormatErr(t, "go", "")
	testFileNamingFormatErr(t, "gOTao", "")
	testFileNamingFormatErr(t, "tao", "")
	testFileNamingFormatErr(t, "goZEro", "welcome_to_go_tao")
	testFileNamingFormatErr(t, "goZERo", "welcome_to_go_tao")
	testFileNamingFormatErr(t, "taogo", "welcome_to_go_tao")
}

func testFileNamingFormat(t *testing.T, format, in, expected string) {
	format, err := FileNamingFormat(format, in)
	assert.Nil(t, err)
	assert.Equal(t, expected, format)
}

func testFileNamingFormatErr(t *testing.T, format, in string) {
	_, err := FileNamingFormat(format, in)
	assert.Error(t, err)
}
