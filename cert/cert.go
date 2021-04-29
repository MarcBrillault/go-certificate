package cert

import (
	"fmt"
	"strings"
	"time"
)

var MaxLenCourse = 20
var MaxLenName = 30

type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

type Saver interface {
	Save(c Cert) error
}

func New(course, name, date string) (*Cert, error) {
	c, err := validateCourse(course)
	if err != nil {
		return nil, err
	}

	n, err := validateName(name)
	if err != nil {
		return nil, err
	}

	d, err := parseDate(date)

	cert := &Cert{
		Course:             c,
		Name:               n,
		Date:               d,
		LabelTitle:         fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion:    "Certificate of completion",
		LabelPresented:     "This certificate is presented to",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2006")),
	}

	return cert, nil
}

func validateCourse(course string) (string, error) {
	course, err := validateStr(course, MaxLenCourse)
	if err != nil {
		return "", err
	}

	if !strings.HasSuffix(strings.ToLower(course), " course") {
		course += " course"
	}

	return strings.ToUpper(course), nil
}

func validateName(n string) (string, error) {
	n, err := validateStr(n, MaxLenName)
	if err != nil {
		return "", err
	}

	n = strings.Title(n)

	return n, nil
}

func validateStr(str string, maxLen int) (string, error) {
	str = strings.TrimSpace(str)

	if len(str) == 0 {
		return str, fmt.Errorf("Invalid string. Got '%s', len=%d", str, len(str))
	} else if maxLen > 0 && len(str) > maxLen {
		return str, fmt.Errorf("Invalid string length. Got '%s', len=%d. Max length is %d", str, len(str), maxLen)
	}

	return str, nil
}

func parseDate(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}

	return t, nil
}
