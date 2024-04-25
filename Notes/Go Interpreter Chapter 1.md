# 1. Lexical Analysis
	Change our source code two times
Source Code -> Tokens -> Abstract Syntax Tree

- **Lexical Analysis:** From Source code to Tokens
	- Done by *Lexer or tokenize/scanner* with each having subtle differences
	- **Tokens** are small, easily categorizable data structures 
		- Fed to the parser 
- Parser
	- token to Abstract Syntax tree
ex. 
'''
"let x = 5 + 5" -> Lexer:
[
	Let, 
	IDENTIFIER("x"), 
	EQUAL_SIGN, INTEGER(5), 
	PLUS_SIGN, 
	INTEGER(5), 
	SEMICOLON
]
'''
- Some languages white space are tokens but white space is not significant in monkey
	- or new line characters
- Production-ready lexers attach meta information about line numbers, column num and filenames to token
	- for error handling in parsing
# 2. Defining our Tokens
- Defining outputs
	- integers, special characters
- Defining data structure
	- Type attribute to distinguish *integers* and *right brackets*
	- token/token.go
- Possible Token types as **const{}**
	- Two Special Types **ILLEGAL and EOF**
# 3. The Lexer
- Input: Source code, Output: Tokens
	- No buffer or saving
	- One method: **NextToken()**
# notes
- Lexer supports ASCII charters only  
# Code
- Initialize Lexer -> call NextToken() -> iterates through source code
- *lexer/lexer_test.go*
	- **func TestNextToken(t *testing.T)**
- *lexer/lexer.go*
	- **New()** function to return *Lexer
	- **readChar()** to give us next character and advance the position
