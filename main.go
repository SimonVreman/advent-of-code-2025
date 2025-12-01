package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"simonvreman/advent-of-code-2025/src/days"
	"simonvreman/advent-of-code-2025/src/util"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	command := getCommandArg()
	day := getDayArg()
	dayPath := filepath.Join("src", "days", strconv.Itoa(day))

	if command == "create" {
		writeDayModule(dayPath)
		writeDayFile(day, dayPath)
		createInputFiles(dayPath)
		buildImportMap()
		fmt.Println("created files for day", day)
	} else {
		runDay(day, dayPath)
	}
}

func getCommandArg() string {
	if len(os.Args) < 2 {
		log.Fatal(errors.New("too few arguments, specify command (run/create)"))
	}

	command := os.Args[1]
	if command != "run" && command != "create" {
		log.Fatal("expected `run` or `create` as the first argument")
	}

	return command
}

func getDayArg() int {
	if len(os.Args) < 3 {
		log.Fatal(errors.New("too few arguments, specify day number"))
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(fmt.Errorf("expected a day number: %v", err.Error()))
	}

	return day
}

func writeDayModule(dayPath string) string {
	err := os.Mkdir(dayPath, 0755)
	check(err)

	return dayPath
}

func getTemplate(name string) []byte {
	template, err := os.ReadFile(filepath.Join("src", "empty", fmt.Sprintf("%v.template", name)))
	check(err)

	return template
}

func writeDayFile(day int, dayPath string) {
	file, err := os.Create(filepath.Join(dayPath, fmt.Sprintf("day_%v.go", day)))
	check(err)

	template := getTemplate("main")

	_, err = file.Write(bytes.ReplaceAll(template, []byte("_daynumber_"), []byte(strconv.Itoa(day))))
	check(err)
}

func createInputFiles(dayPath string) {
	files := []string{"example_1.txt", "example_2.txt", "input.txt"}

	for _, file := range files {
		_, err := os.Create(filepath.Join(dayPath, file))
		check(err)
	}
}

func buildImportMap() {
	daysPath := filepath.Join("src", "days")
	allFiles, err := os.ReadDir(daysPath)
	check(err)

	subDirs := util.Filter(allFiles, func(file os.DirEntry) bool { return file.IsDir() })
	dayNumbers := util.Map(subDirs, func(dir os.DirEntry) int {
		number, err := strconv.Atoi(dir.Name())
		check(err)
		return number
	})

	template := getTemplate("index")
	file, err := os.Create(filepath.Join("src", "days", "index.go"))
	check(err)

	imports := []byte{}
	entries := []byte{}
	for i, day := range dayNumbers {
		imports = fmt.Appendf(imports, "	days_%v \"simonvreman/advent-of-code-2025/src/days/%v\"", day, day)
		entries = fmt.Appendf(entries, "	%v: {{days_%v.First, days_%v.FirstExpected}, {days_%v.Second, days_%v.SecondExpected}},", day, day, day, day, day)

		if i < (len(dayNumbers) - 1) {
			imports = append(imports, '\n')
			entries = append(entries, '\n')
		}
	}

	_, err = file.Write(
		bytes.ReplaceAll(
			bytes.ReplaceAll(template, []byte("_dayimports_"), imports),
			[]byte("_dayentries_"),
			entries,
		),
	)
	check(err)
}

func runDay(day int, dayPath string) {
	input, err := os.ReadFile(filepath.Join("src", "days", strconv.Itoa(day), "input.txt"))
	check(err)

	for i, part := range days.Solutions[day] {
		fmt.Printf("Running part %v for day %v\n", i+1, day)
		first := part.Fn(input)
		fmt.Printf("Answer to part %v: %v\n", i+1, first)
	}
}
