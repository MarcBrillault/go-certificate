package cert

import (
	"testing"
)

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2021-04-21")
	if err != nil {
		t.Errorf("Cert data should be valid. Err: %v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference, got nil. Error: %v", err)
	}
	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name should be 'GOLANG COURSE', but got '%v'", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2021-04-21")
	if err == nil {
		t.Error("Error should be returned on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	_, err := New("azertyuiopqsdfghjklmwxcvbn", "Bob", "2021-04-21")
	if err == nil {
		t.Error("Error should be returned on a too long course")
	}
}

func TestNameTitleCase(t *testing.T) {
	c, _ := New("Golang", "bob", "2021-04-21")

	if c.Name != "Bob" {
		t.Errorf("Invalid name. Expected 'Bob', got '%s'", c.Name)
	}
}
func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2021-04-21")

	if err == nil {
		t.Error("Error should be returned on an empty name")
	}
}

func TestNameTooLong(t *testing.T) {
	c, err := New("golang", "Daenerys Targaryen, Queen of Meereen, Khaleesi of the Great Grass Sea, Mother of Dragons, The Unburnt, Breaker of Chains", "2021-04-21")

	if err == nil {
		t.Errorf("Names longer than 30 characters should not be allowed. Got '%s' with a length of %d characters", c.Name, len(c.Name))
	}
}
