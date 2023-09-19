## Feedback

- +Tasks creates watchfulness
- -Not obvious, that there is few answers and answer has not only one right code line

### Changes
1. [x] Show task instructions
   - Show "one answer" or "few|0/2 answers" tag
   - Show "with any line" or "with related lines"
2. [x] Add tutorial challenge (with just syntax errors)
3. [x] ~~Add copy to clipboard button~~ Add auto copy on challenge start
4. Add learning materials links button [go install golang.org/x/website/tour@latest | go doc sync.RWMutex | go doc sync.WaitGroup | https://www.educative.io/answers/what-are-pointers-in-golang | https://neilalexander.dev/2021/08/29/go-pass-by-value]

## Feedback 2

- Low terminal height -> "I don't know that to do"
- "Do you understand that code has copied?" -> No
- "I can't select question", I don't understand "Select an answer and a problem code line using arrows and '⮐ '" -> different step in tutorial

### Changes

1. [x] Low height error
2. [x] Code copy question in tutorial
3. [x] Fix code selection bug. Add tutorial step

### Future idea

- Optimization quizz

```
func somethink(str string) string {
	ss := strings.Split(str, "\n")

	for i, s := range ss {
		newStr := "    " + s

		ss[i] = newStr
	}

	return strings.Join(ss, "\n")
}

func somethink2(str string) string {
	sb := strings.Builder{}
	var i int

	sb.Grow(len(str) + 4)
	sb.WriteString("    ")

	for {
		i = strings.IndexRune(str, '\n')
		if i < 0 {
			sb.WriteString(str)
			break
		}

		sb.WriteString(str[:i+1])
		str = str[i+1:]
		sb.WriteString("    ")
	}

	return sb.String()
}
```
