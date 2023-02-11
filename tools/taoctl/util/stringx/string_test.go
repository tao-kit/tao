package stringx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_IsEmptyOrSpace(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{
			want: true,
		},
		{
			input: " ",
			want:  true,
		},
		{
			input: "\t",
			want:  true,
		},
		{
			input: "\n",
			want:  true,
		},
		{
			input: "\f",
			want:  true,
		},
		{
			input: "		",
			want:  true,
		},
	}
	for _, v := range cases {
		s := From(v.input)
		assert.Equal(t, v.want, s.IsEmptyOrSpace())
	}
}

func TestString_Snake2Camel(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{
			input: "__",
			want:  "",
		},
		{
			input: "go_tao",
			want:  "GoTao",
		},
		{
			input: "の_go_tao",
			want:  "のGoTao",
		},
		{
			input: "goTao",
			want:  "GoTao",
		},
		{
			input: "goTao",
			want:  "GoTao",
		},
		{
			input: "goTao_",
			want:  "GoTao",
		},
		{
			input: "go_Tao_",
			want:  "GoTao",
		},
		{
			input: "_go_Tao_",
			want:  "GoTao",
		},
	}
	for _, c := range cases {
		ret := From(c.input).ToCamel()
		assert.Equal(t, c.want, ret)
	}
}

func TestString_Camel2Snake(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{
			input: "goTao",
			want:  "go_tao",
		},
		{
			input: "Gotao",
			want:  "gotao",
		},
		{
			input: "GoTao",
			want:  "go_tao",
		},
		{
			input: "Go_Tao",
			want:  "go__tao",
		},
	}
	for _, c := range cases {
		ret := From(c.input).ToSnake()
		assert.Equal(t, c.want, ret)
	}
}

func TestTitle(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{
			input: "go tao",
			want:  "Go Tao",
		},
		{
			input: "goTao",
			want:  "GoTao",
		},
		{
			input: "GoTao",
			want:  "GoTao",
		},
		{
			input: "の go tao",
			want:  "の Go Tao",
		},
		{
			input: "Gotao",
			want:  "Gotao",
		},
		{
			input: "Go_tao",
			want:  "Go_tao",
		},
		{
			input: "go_tao",
			want:  "Go_tao",
		},
		{
			input: "Go_Tao",
			want:  "Go_Tao",
		},
	}
	for _, c := range cases {
		ret := From(c.input).Title()
		assert.Equal(t, c.want, ret)
	}
}

func TestUntitle(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{
			input: "go tao",
			want:  "go tao",
		},
		{
			input: "GoTao",
			want:  "goTao",
		},
		{
			input: "Gotao",
			want:  "gotao",
		},
		{
			input: "Go_tao",
			want:  "go_tao",
		},
		{
			input: "go_tao",
			want:  "go_tao",
		},
		{
			input: "Go_Tao",
			want:  "go_Tao",
		},
	}
	for _, c := range cases {
		ret := From(c.input).Untitle()
		assert.Equal(t, c.want, ret)
	}
}
