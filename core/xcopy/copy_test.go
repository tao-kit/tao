package xcopy

import (
	"testing"
	"time"
)

type Person struct {
	Name     string
	Birthday *time.Time
	Age      int32
}

type User struct {
	Name     string
	Birthday *time.Time
	Age      int32
	Roles    []string
}

func checkPerson(person Person, user User, t *testing.T, testCase string) {
	if person.Name != user.Name {
		t.Errorf("%v: Name haven't been copied correctly.", testCase)
	}

	if person.Birthday == nil && user.Birthday != nil {
		t.Errorf("%v: Birthday haven't been copied correctly.", testCase)
	}

	if person.Birthday != nil && user.Birthday == nil {
		t.Errorf("%v: Birthday haven't been copied correctly.", testCase)
	}

	if person.Birthday != nil && user.Birthday != nil &&
		!person.Birthday.Equal(*user.Birthday) {
		t.Errorf("%v: Birthday haven't been copied correctly.", testCase)
	}

}

func TestCopyStruct(t *testing.T) {
	user := &User{
		Name:  "sllt",
		Age:   18,
		Roles: []string{"hello", "world", "3"},
	}
	person := &Person{}

	Copy(person, user)
	checkPerson(*person, *user, t, "check copy struct")
}
