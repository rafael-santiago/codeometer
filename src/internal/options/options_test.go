//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package options

import (
    "testing"
    "os"
    "reflect"
)

func TestGetCommand(t *testing.T) {
    oldArgs := os.Args
    defer func() {
            os.Args = oldArgs
    }()
    testVector := []struct {
        Expected string
        Args []string
    }{
        { "", []string{} },
        { "one", []string{"./cmd", "one", "--foo=bar", "--bar=foo"} },
        { "two", []string{"./cmd", "two", "--abc=def", "--ghijklmn=opqrstuvwxyz." } },
        { "three", []string{"./cmd", "three", "--abc=def", "--2nd=option" } },
        { "go", []string{"./cmd", "go", "--abc=def", "--2nd=option" } },
    }
    for _, test := range testVector {
        os.Args = test.Args
        command := GetCommand()
        if command != test.Expected {
            t.Errorf(`command != test.Expected: %v != %v`, command, test.Expected)
        }
    }
}

func TestGetOption(t *testing.T) {
    oldArgs := os.Args
    defer func() {
            os.Args = oldArgs
    }()
    testVector := []struct {
        Option string
        Default string
        Expected string
        Args []string
    }{
        { "foo", "", "bar", []string{"./cmd", "command", "--foo=bar", "--bar=foo"} },
        { "boo", "ah!", "ah!", []string{"./cmd", "test", "--abc=def", "--ghijklmn=opqrstuvwxyz." } },
        { "2nd", "way", "option", []string{"./cmd", "go", "--abc=def", "--2nd=option" } },
    }
    for _, test := range testVector {
        os.Args = test.Args
        option := GetOption(test.Option, test.Default)
        if option != test.Expected {
            t.Errorf(`option != test.Expected: %v != %v`, option, test.Expected)
        }
    }
}

func TestGetBoolOption(t *testing.T) {
    oldArgs := os.Args
    defer func() {
            os.Args = oldArgs
    }()
    testVector := []struct {
        Option string
        Default bool
        Expected bool
        Args []string
    }{
        { "foo", false, true, []string{"./cmd", "go", "--bar", "--foo"} },
        { "boo", true, true, []string{"./cmd", "gogo", "--abc", "--ghijklmnopqrstuvwxyz" } },
        { "2nd", false, true, []string{"./cmd", "gogogo", "--abc", "--2nd"}},
    }
    for _, test := range testVector {
        os.Args = test.Args
        option := GetBoolOption(test.Option, test.Default)
        if option != test.Expected {
            t.Errorf(`option != test.Expected: %v != %v`, test.Option, test.Expected)
        }
    }
}

func TestGetArrayOption(t *testing.T) {
    oldArgs := os.Args
    defer func() {
            os.Args = oldArgs
    }()
    testVector := []struct {
        Option string
        Default []string
        Expected []string
        Args []string
    }{
        { "foo", []string{"foo", "boo", "bar", "baz"}, []string{"boo", "foo", "baz", "bar"}, []string{"./cmd", "c1", "--foo=boo,foo,baz,bar", "--bar=foo,boo,boo"} },
        { "boo", []string{"ah!", "hahaha"}, []string{"rabbit!", "aie!", "ruuuunnn-awwaaaay!!!"}, []string{"./cmd", "c2", "--abc=def", `--boo=rabbit!,aie!,ruuuunnn-awwaaaay!!!`} },
        { "2nd", []string{ "w", "a", "y"}, []string{"o", "p", "t", "i", "o", "n"}, []string{"./cmd", "c3", "--abc=def", "--2nd=o,p,t,i,o,n" } },
    }
    for _, test := range testVector {
        os.Args = test.Args
        options := GetArrayOption(test.Option, test.Default...)
        if reflect.DeepEqual(options, test.Expected) == false {
            t.Errorf(`options != test.Expected: %v != %v`, options, test.Expected)
        }
    }
}
