# Go Quiz

![alt](https://freepngimg.com/thumb/hat/107810-hat-sorting-potter-harry-png-free-photo.png)

- A simple CLI-based quiz game.
- It takes a JSON file with an array of questions and answers as an input. Use the `--file` flag.
- The quiz can be timed as well, use the `--limit` flag and provide a time limit (in seconds) for your quiz.
- It isn't limited to single-worded answers.
- A final score is display in the end.
- A sample quiz file can be found in the `sample` folder.

### Usage

```bash
$ go build .
$ ./quiz-game --file=./sample/quiz.json --limit=20
```
