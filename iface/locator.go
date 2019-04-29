package iface

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gopkg.in/yaml.v2"
)

type Locator interface {
	isLocator()
	GetScheme() string
	Validate() error
}

type LocatorBase struct {
	Scheme string
}

func (LocatorBase) isLocator() {
}

func (loc LocatorBase) GetScheme() string {
	return loc.Scheme
}

func ParseLocator(input string) (loc Locator, e error) {
	var locw LocatorWrapper
	if e = yaml.Unmarshal([]byte(input), &locw); e != nil {
		return loc, e
	}
	loc = locw.Locator
	return loc, nil
}

func MustParseLocator(input string) (loc Locator) {
	loc, e := ParseLocator(input)
	if e != nil {
		panic(e)
	}
	return loc
}

var locatorTypes = make(map[string]reflect.Type)

func RegisterLocatorType(locator Locator, schemes ...string) {
	typ := reflect.TypeOf(locator)
	if typ.Kind() != reflect.Struct {
		panic("locator must be a struct")
	}
	for _, scheme := range schemes {
		locatorTypes[scheme] = typ
	}
}

type LocatorWrapper struct {
	Locator
}

func (locw *LocatorWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(locw.Locator)
}

func (locw *LocatorWrapper) UnmarshalJSON(data []byte) error {
	return locw.UnmarshalYAML(func(v interface{}) error {
		return json.Unmarshal(data, v)
	})
}

func (locw *LocatorWrapper) MarshalYAML() (interface{}, error) {
	if locw.Locator == nil {
		return nil, nil
	}

	scheme := locw.Locator.GetScheme()
	if typ, ok := locatorTypes[scheme]; !ok {
		return nil, fmt.Errorf("unknown scheme %s", scheme)
	} else if typ != reflect.TypeOf(locw.Locator) {
		return nil, fmt.Errorf("unexpected type %T", locw.Locator)
	}

	if e := locw.Locator.Validate(); e != nil {
		return nil, e
	}
	return locw.Locator, nil
}

func (locw *LocatorWrapper) UnmarshalYAML(unmarshal func(interface{}) error) (e error) {
	schemeObj := struct {
		Scheme string
	}{}
	if e = unmarshal(&schemeObj); e != nil {
		return e
	}

	typ, ok := locatorTypes[schemeObj.Scheme]
	if !ok {
		return fmt.Errorf("unknown scheme %s", schemeObj.Scheme)
	}

	ptr := reflect.New(typ)
	if e = unmarshal(ptr.Interface()); e != nil {
		return e
	}

	loc := ptr.Elem().Interface().(Locator)
	if e = loc.Validate(); e != nil {
		return e
	}

	locw.Locator = loc
	return nil
}