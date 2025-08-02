# Quiz Game (golang)
The first project following the [Gophercises](https://gophercises.com/) course
for practicing and mastering Go.

## How it works
The program reads a CSV file, where the first field is the question, and the
second field is the answer to the respective question.

User is prompted with a question, and the program awaits his answer. If the
answer is correct, the user's score is incremented. If the answer is wrong,
the program continues as normal, but without incrementing the score.

Upon reaching the last question, or the time running out, the user is shown
their score against the total number of questions, and the program stops.

## Arguments
```
-f string
        path to problems file (default "problems.csv")
-s bool
        shuffle the problems list (default false)
-t int
        time to play in seconds (default 30)
```