package query

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gchaincl/dotsql"
)

func init() {
	mustReadQueries(&User, "sql/queries/users.sql")
}

func mustReadQueries(dest interface{}, name string) {
	a := MustAsset(name)
	if err := readSQLFromString(dest, string(a)); err != nil {
		log.Fatalf("Failed to load queries from: %s. Error is: %s", name, err)
	}
}

func readSQLFromString(dest interface{}, s string) error {
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("must pass a pointer, not a value, to resources reader destination")
	}
	if v.IsNil() {
		return fmt.Errorf("nil pointer passed to resources reader destination")
	}

	dot, err := dotsql.LoadFromString(s)
	if err != nil {
		return err
	}
	elem := v.Elem()
	for i := 0; i < elem.NumField(); i++ {
		valueField := elem.Field(i)
		typeField := elem.Type().Field(i)
		if !valueField.CanSet() {
			return fmt.Errorf("resources reader tried to read resources for %v, but field %s is not settable",
				reflect.TypeOf(dest),
				typeField.Name)
		}
		if valueField.Kind() != reflect.String {
			return fmt.Errorf("resources reader tried to read resources for %v, but field %s is not string",
				reflect.TypeOf(dest),
				typeField.Name)
		}
		expr, err := dot.Raw(typeField.Name)
		if err != nil {
			return err
		}
		valueField.SetString(expr)
	}
	return nil
}
