package model

import "fmt"

type RecurringValue struct {
	Count    int
	FileName *string
}

type ValuesWrapper struct {
	Values map[string]*RecurringValue
}

type NestedValuesWrapper struct {
	NestedValues map[string]*ValuesWrapper
}

func (v *ValuesWrapper) AppendValue(word *string) {
	if val, ok := v.Values[*word]; !ok {
		v.Values[*word] = &RecurringValue{Count: 1}
	} else {
		val.Count++
	}
}

func (v *NestedValuesWrapper) AppendNestedValue(word *string, fileName *string) {
	if val, ok := v.NestedValues[*fileName].Values[*word]; !ok {
		v.NestedValues[*fileName].Values[*word] = &RecurringValue{Count: 1, FileName: fileName}
	} else {
		val.Count++
	}
}

func (v *ValuesWrapper) PrintValues() {
	printValues(v.Values)
}

func (v *NestedValuesWrapper) PrintNestedValues() {
	for _, vals := range v.NestedValues {
		printValues(vals.Values)
	}
}

func printValues(vals map[string]*RecurringValue) {
	for line, val := range vals {
		if val.Count > 1 {
			printLine(line, val)
		}
	}
}

func printLine(line string, val *RecurringValue) {
	if val.FileName != nil {
		fmt.Printf("%d\t%s\t%s\n", val.Count, line, *val.FileName)
	} else {
		fmt.Printf("%d\t%s\n", val.Count, line)
	}
}
