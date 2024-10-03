# Simple_Math_Quiz
Here we have Ten questions which appear and we have to answer them and it shows the score at the end. The total number of questions is 10
# Table of Contents
-Features 

-Installation

-Usage

-CSV File Format

-Contributing


# Features
-Command-line tool for taking a math quiz.

-Automatically reads a list of questions and answers from a CSV file.

-Keeps track of the score and provides a result after the quiz.

# Installation
1.Make sure you have Go installed on your machine.

2.Clone this repository to your local machine:

     git clone https://github.com/your-username/math-quiz-game.git
     cd math-quiz-game
     
3.Build the executable:

     go build -o quiz main.go
     
# Usage 
-To run the Math Quiz Game, use the following command:

     ./quiz -csv=problems.csv
     
This command will read the quiz questions from problems.csv and start the game.     

# CSV File Format
The CSV file should contain questions and answers in two columns, separated by a comma. The first column is the question, and the second is the correct answer.

Example problems.csv:

5+5,10

7-3,4

12/3,4

9*9,81

# Contributing
Contributions are welcome! If you find bugs or have suggestions for improvement, feel free to open an issue or submit a pull request.
