package markdown

import (
	"amireshoon/gow/gow"
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// AddTitle add or change existing title
func AddTitle(title string) {

}

// AddTodo adds new todo to TODO.md file
func AddTodo(desc string, path string) {
	c, err := gow.GetTodo(path)

	if err != nil {
		fmt.Println("Could not read TODO.md file")
	}

	c += `
- [ ] ` + desc + ``
	err = gow.FillTodo(c, path)
	if err != nil {
		fmt.Println("Could not write to file")
	}
}

// CheckTodo checks existing todo
func CheckTodo(index int, path string) {
	c, err := gow.GetTodo(path)

	if err != nil {
		fmt.Println("Could not read TODO.md file")
	}

	// fmt.Println(c)
	reGeneratedTodo := ``
	scanner := bufio.NewScanner(strings.NewReader(c))
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, `- [ ] `) {
			if counter == index {
				reGeneratedTodo += strings.Replace(line, `- [ ] `, `- [x] `, -1) + "\n"
			} else {
				reGeneratedTodo += line + "\n"
			}
			counter++
		} else {
			reGeneratedTodo += line + "\n"
		}
	}
	err = gow.FillTodo(reGeneratedTodo, path)
	if err != nil {
		fmt.Println("Could not write to file")
	}
}

// HasTodo returns true if it's already initilized or TODO.md already exists
func HasTodo(path string) bool {
	_, err := gow.GetTodo(path)
	if err != nil {
		return false
	}
	return true
}

func loadReadme(path string) (string, error) {

openingFile:
	content, err := ioutil.ReadFile(path + "/README.md")
	if err != nil {
		err := ioutil.WriteFile(path+"README.md", []byte(""), 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
		goto openingFile
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text, err
}

// AddToReadme will add TODO.md to bottom of README.md
func AddToReadme(path string) {

	readme, err := loadReadme(path)

	if err != nil {
		fmt.Println("Could not read or write into README.md")
	}

	c, err := gow.GetTodo(path)

	if err != nil {
		fmt.Println("Could not find TODO.md")
	}

	readme += "\n" + c + "\n"
	gow.FillReadme(readme, path)
}

func getTodoVersion(path string) (string, error) {
	c, _ := gow.GetTodo(path)

	scanner := bufio.NewScanner(strings.NewReader(c))

	for scanner.Scan() {
		r, _ := regexp.Compile(`\(+([a-z A-Z 0-9 :])+\)`)
		if r.MatchString(scanner.Text()) {
			s := strings.Split(r.FindString(scanner.Text()), ":")
			return strings.Replace(s[1], ")", "", -1), nil
		}
		break
	}

	return "", nil
}
