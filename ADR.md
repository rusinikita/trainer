
```mermaid
---
title: Data model
---
erDiagram
  task ||--o{ question : contains
  task {
      string name
      string category "beginner, easy, concurrency, wtf"
      string code_window_default_text
  }
  question ||--|{ answer : contains
  question {
      string text
      string type "select_answer, code_line_to_answer, write_code"
  }
  answer {
      string name
      string text
      int[][] code_line_ranges "[[1,2][8,12]] or [] as none or NULL as all"
  }
```

```mermaid
---
title: Components and state
---
classDiagram
    note for TaskSelectionScreen "beginer easy concurrency wtf
    
   task1
   task2
   task3
   
   <> - category | up/down - task"
    class TaskSelectionScreen {
        -Category[] categories
        -Task[] loadedTasks
        -Task[] filteredTasks
    }

    class TaskScreen {
        -string name
        -Question[] questions
        -currentQuestion int
        -nextQuestion()
    }

    class CodeView {
        -string code
        -bool isLineCursorMode
        -int currentLine
        +selectedLine()
    }

    %% Question creates codeView with settings and calls CurrentLine() getter
    class Question {
        +isAnswered()
    }

    class Answer {
        string name
        string text
        int[][] code_line_ranges
        bool isAnswered
        bool isError
        bool isSelected
    }

```

State model:

Features plan:
-[ ] Select task category
  -[ ] Show categories list
  -[ ] Select category
  -[ ] Show task category stats (all - N, completed - X, updated - Y)
-[ ] Select task
  -[ ] Show tasks list
  -[ ] Select task
  -[ ] Show task status (all - N, completed - X, updated - Y)
-[ ] Do task
  -[ ] Show code with code line numbers
  -[ ] Step 1: Select code question answer
    -[ ] Show error state, remove error status when move
    -[ ] Show good state
  -[ ] Next step
    -[ ] Show go to next step
    -[ ] Save current task state
    -[ ] Load saved task state
  -[ ] Step 2: Select code line and question answer variant
    -[ ] Select code line
    -[ ] Select question answer
    -[ ] Show error status 
    -[ ] Show answered status
  -[ ] Keybindings help
  -[ ] Setup layout
    -[ ] Vertical layout
    -[ ] Split screen layout
  -[ ] Step 3: Rewrite code (Ctrl+C, Ctrl+V works)
    -[ ] Show code
    -[ ] Show packages allowed to import
    -[ ] Crtr+C for copy code into buffer
    -[ ] Crtl+V for paste code into editor
    -[ ] Create tmp folder, create files, go run, handle answer
    -[ ] Run tests
    -[ ] Run linter
    -[ ] Show problems list
    -[ ] Code editor inside app
-[ ] Mentoring
  -[ ] Create code review request to mentor
  -[ ] Sync app state mentor <-> mentee