package calendar

import (
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type calendar struct {
	component.C
	name string
}

func NewCalendar(name string) *calendar {
	return &calendar{name: name}
}

func (c *calendar) SetToday() {
	attr := c.GetAttributes()
	if attr == nil {
		attr = make(templ.Attributes)
	}
	attr["today"] = true
	c.SetAttributes(attr)
}
func (c *calendar) EnableMonthYearSelection() {
	attr := c.GetAttributes()
	if attr == nil {
		attr = make(templ.Attributes)
	}
	attr["jumpable"] = true
	c.SetAttributes(attr)
}
func (c *calendar) StartsOnSunday() {
	attr := c.GetAttributes()
	if attr == nil {
		attr = make(templ.Attributes)
	}
	attr["starts-with"] = "0"
	c.SetAttributes(attr)
}
func (c *calendar) StartsOnMonday() {
	attr := c.GetAttributes()
	if attr == nil {
		attr = make(templ.Attributes)
	}
	attr["starts-with"] = "1"
	c.SetAttributes(attr)
}
func (c *calendar) DisableDates(dates []time.Time) {
	attr := c.GetAttributes()
	if attr == nil {
		attr = make(templ.Attributes)
	}
	var dateStrings []string
	for _, date := range dates {
		dateStrings = append(dateStrings, date.Format("2006-01-02"))
	}
	attr["disabled-dates"] = strings.Join(dateStrings, ",")
	c.SetAttributes(attr)
}
func (c *calendar) MarkDates(dates []time.Time) {
	attr := c.GetAttributes()
	if attr == nil {
		attr = make(templ.Attributes)
	}
	var dateStrings []string
	for _, date := range dates {
		dateStrings = append(dateStrings, date.Format("2006-01-02"))
	}
	attr["marked-dates"] = strings.Join(dateStrings, ",")
	c.SetAttributes(attr)
}

func (c *calendar) ViewDate(date time.Time) {
	attr := c.GetAttributes()
	if attr == nil {
		attr = make(templ.Attributes)
	}
	attr["view-date"] = date.Format("2006-01-02")
	c.SetAttributes(attr)
}
func (c *calendar) SetMinimumDate(date time.Time) {
	attr := c.GetAttributes()
	if attr == nil {
		attr = make(templ.Attributes)
	}
	attr["min"] = date.Format("2006-01-02")
	c.SetAttributes(attr)
}
func (c *calendar) SetMaximumDate(date time.Time) {
	attr := c.GetAttributes()
	if attr == nil {
		attr = make(templ.Attributes)
	}
	attr["max"] = date.Format("2006-01-02")
	c.SetAttributes(attr)
}
func (c *calendar) SetCurrentDate(date time.Time) {
	attr := c.GetAttributes()
	if attr == nil {
		attr = make(templ.Attributes)
	}
	attr["value"] = date.Format("2006-01-02")
	c.SetAttributes(attr)
}
func (c *calendar) Set118n(weekdays []string, months []string) {
	attr := c.GetAttributes()
	if attr == nil {
		attr = make(templ.Attributes)
	}
	wd := strings.Join(weekdays, ",")
	m := strings.Join(months, ",")
	attr["i18n"] = "weekdays: " + wd + "; months: " + m
	c.SetAttributes(attr)
}
func (c *calendar) AddCustomClasses(classes ...string) {
	class := strings.Join(classes, " ")
	attr := c.GetAttributes()
	if attr == nil {
		attr = make(templ.Attributes)
	}
	attr["cls-custom"] = class
	c.SetAttributes(attr)
}