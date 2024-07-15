package types

import (
	"reflect"
	"time"
)

type Task struct {
	ID      uint
	Name    string
	Project string
	Status  Status
	Created time.Time
	Updated time.Time
}

func (t Task) FilterValue() string {
	return t.Name
}

func (t Task) Title() string {
	return t.Name
}

func (t Task) Description() string {
	return t.Project
}

func (ot *Task) Merge(nt Task) {
	uVal := reflect.ValueOf(&nt).Elem()
	oVal := reflect.ValueOf(ot).Elem()
	defaultDate := time.Time{}

	for i := 0; i < uVal.NumField(); i++ {
		uField := uVal.Field(i).Interface()
		if oVal.CanSet() {
			if v, ok := uField.(int64); ok && uField != 0 {
				oVal.Field(i).SetInt(v)
			}

			if v, ok := uField.(time.Time); ok && !v.Equal(defaultDate) {
				oVal.Field(i).Set(reflect.ValueOf(v))
			} else if v, ok := uField.(string); ok && uField != "" {
				oVal.Field(i).SetString(v)
			}
		}
	}
}
