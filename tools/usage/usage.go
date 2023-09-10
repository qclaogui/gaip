// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package usage

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/grafana/dskit/flagext"
	"github.com/qclaogui/golang-api-server/tools/fieldcategory"
)

const (
	customStr      = "string"
	customDuration = "duration"
	customInt      = "int"
	customUint     = "uint"
	customFloat    = "float"
)

// Usage prints command-line usage.
// printAll controls whether only basic flags or all flags are included.
// configs are expected to be pointers to structs.
func Usage(printAll bool, configs ...interface{}) error {
	fields := map[uintptr]reflect.StructField{}
	for _, c := range configs {
		if err := parseStructure(c, fields); err != nil {
			return err
		}
	}

	fs := flag.CommandLine
	_, _ = fmt.Fprintf(fs.Output(), "Usage of %s:\n", os.Args[0])
	fs.VisitAll(func(fl *flag.Flag) {
		v := reflect.ValueOf(fl.Value)
		fieldCat := fieldcategory.Basic
		var field reflect.StructField
		var hasField bool

		// Do not print usage for deprecated flags.
		if fl.Value.String() == "deprecated" {
			return
		}
		if v.Kind() == reflect.Ptr {
			ptr := v.Pointer()
			field, hasField = fields[ptr]
			if hasField && isFieldHidden(field) {
				// Don't print help for this flag since it's hidden
				return
			}
		}

		if override, ok := fieldcategory.GetOverride(fl.Name); ok {
			fieldCat = override
		} else if hasField {
			catStr := field.Tag.Get("category")
			switch catStr {
			case "advanced":
				fieldCat = fieldcategory.Advanced
			case "experimental":
				fieldCat = fieldcategory.Experimental
			case "deprecated":
				fieldCat = fieldcategory.Deprecated
			}
		} else {
			// The field is neither an override nor has been parsed, so we'll skip it.
			return
		}

		if fieldCat != fieldcategory.Basic && !printAll {
			// Don't print help for this flag since we're supposed to print only basic flags
			return
		}

		var b strings.Builder
		// Two spaces before -; see next two comments.
		_, _ = fmt.Fprintf(&b, "  -%s", fl.Name)
		name := getFlagName(fl)
		if len(name) > 0 {
			b.WriteString(" ")
			b.WriteString(strings.ReplaceAll(name, " ", "-"))
		}
		// Four spaces before the tab triggers good alignment
		// for both 4- and 8-space tab stops.
		b.WriteString("\n    \t")
		switch fieldCat {
		case fieldcategory.Experimental:
			b.WriteString("[experimental] ")
		case fieldcategory.Deprecated:
			b.WriteString("[deprecated] ")
		}

		b.WriteString(strings.ReplaceAll(fl.Usage, "\n", "\n    \t"))

		if defValue := getFlagDefault(fl, field); !isZeroValue(fl, defValue) {
			v := reflect.ValueOf(fl.Value)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.String {
				// put quotes on the value
				_, _ = fmt.Fprintf(&b, " (default %q)", defValue)
			} else {
				_, _ = fmt.Fprintf(&b, " (default %v)", defValue)
			}
		}
		_, _ = fmt.Fprint(fs.Output(), b.String(), "\n")
	})

	if !printAll {
		_, _ = fmt.Fprintf(fs.Output(), "\nTo see all flags, use -help-all\n")
	}

	return nil
}

// isZeroValue determines whether the string represents the zero
// value for a flag.
func isZeroValue(fl *flag.Flag, value string) bool {
	// Build a zero value of the flag's Value type, and see if the
	// result of calling its String method equals the value passed in.
	// This works unless the Value type is itself an interface type.
	typ := reflect.TypeOf(fl.Value)
	var z reflect.Value
	if typ.Kind() == reflect.Ptr {
		z = reflect.New(typ.Elem())
	} else {
		z = reflect.Zero(typ)
	}
	return value == z.Interface().(flag.Value).String()
}

// parseStructure parses a struct and populates fields.
func parseStructure(structure interface{}, fields map[uintptr]reflect.StructField) error {
	// structure is expected to be a pointer to a struct
	if reflect.TypeOf(structure).Kind() != reflect.Ptr {
		t := reflect.TypeOf(structure)
		return fmt.Errorf("%s is a %s while a %s is expected", t, t.Kind(), reflect.Ptr)
	}
	v := reflect.ValueOf(structure).Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("%s is a %s while a %s is expected", v, v.Kind(), reflect.Struct)
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Func {
			continue
		}

		fieldValue := v.FieldByIndex(field.Index)

		// Take address of field value and map it to field
		fields[fieldValue.Addr().Pointer()] = field

		// Recurse if a struct
		if field.Type.Kind() != reflect.Struct || isFieldHidden(field) || ignoreStructType(field.Type) || !field.IsExported() {
			continue
		}

		if err := parseStructure(fieldValue.Addr().Interface(), fields); err != nil {
			return err
		}
	}

	return nil
}

// Descending into some structs breaks check for "advanced" category for some fields (eg. flagext.Secret),
// because field itself is at the same memory address as the internal field in the struct, and advanced-category-check
// then gets confused.
var ignoredStructTypes = []reflect.Type{
	reflect.TypeOf(flagext.Secret{}),
}

func ignoreStructType(fieldType reflect.Type) bool {
	for _, t := range ignoredStructTypes {
		if fieldType == t {
			return true
		}
	}
	return false
}

func getFlagName(fl *flag.Flag) string {
	if getter, ok := fl.Value.(flag.Getter); ok {
		if v := reflect.ValueOf(getter.Get()); v.IsValid() {
			t := v.Type()
			switch t.Name() {
			case "bool":
				return ""
			case "Duration":
				return customDuration
			case "float64":
				return customFloat
			case "int", "int64":
				return customInt
			case "string":
				return customStr
			case "uint", "uint64":
				return customUint
			case "Secret":
				return customStr
			default:
				return "value"
			}
		}
	}

	// Check custom types.
	if v := reflect.ValueOf(fl.Value); v.IsValid() {
		switch v.Type().String() {
		case "*flagext.Secret":
			return customStr
		case "*flagext.StringSlice":
			return customStr
		case "*flagext.StringSliceCSV":
			return "comma-separated list of strings"
		case "*flagext.CIDRSliceCSV":
			return "comma-separated list of strings"
		case "*flagext.URLValue":
			return customStr
		case "*url.URL":
			return customStr
		case "*model.Duration":
			return customDuration
		case "*tsdb.DurationList":
			return "comma-separated list of durations"
		}
	}
	return "value"
}

func isFieldHidden(f reflect.StructField) bool {
	return getDocTagFlag(f, "hidden")
}

func getDocTagFlag(f reflect.StructField, name string) bool {
	cfg := parseDocTag(f)
	_, ok := cfg[name]
	return ok
}

func getFlagDefault(fl *flag.Flag, field reflect.StructField) string {
	if docDefault := parseDocTag(field)["default"]; docDefault != "" {
		return docDefault
	}
	return fl.DefValue
}

func parseDocTag(f reflect.StructField) map[string]string {
	cfg := map[string]string{}
	tag := f.Tag.Get("doc")

	if tag == "" {
		return cfg
	}

	for _, entry := range strings.Split(tag, "|") {
		parts := strings.SplitN(entry, "=", 2)

		switch len(parts) {
		case 1:
			cfg[parts[0]] = ""
		case 2:
			cfg[parts[0]] = parts[1]
		}
	}

	return cfg
}
