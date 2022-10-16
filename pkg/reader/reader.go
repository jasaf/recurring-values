package reader

import (
	"bufio"
	"os"

	m "recurring-values/pkg/model"
)

func ReadWords(words *[]string, values *m.ValuesWrapper) {
	for _, arg := range *words {
		values.AppendValue(&arg)
	}
}

func ReadFiles(files *[]string, nestedValues *m.NestedValuesWrapper) {
	for _, arg := range *files {
		file, err := os.Open(arg)

		nestedValues.NestedValues[arg] = &m.ValuesWrapper{Values: make(map[string]*m.RecurringValue)}

		if err != nil {
			panic(err)
		}

		appendOccurances(file, nestedValues)

		file.Close()
	}
}

func appendOccurances(f *os.File, values *m.NestedValuesWrapper) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		text := input.Text()
		fileName := f.Name()

		values.AppendNestedValue(&text, &fileName)
	}
}
