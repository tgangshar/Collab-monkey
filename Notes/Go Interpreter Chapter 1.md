# Lexical Analysis
	Change our source code two times
Source Code -> Tokens -> Abstract Syntax Tree

- **Lexical Analysis:** From Source code to Tokens
	- Done by *Lexer or tokenize/scanner* with each having subtle differences
	- **Tokens** are small, easily categorizable data structures 
		- Fed to the parser 
- Parser
	- token to Abstract Syntax tree
"let x = 5 + 5" -> Lexer -> 
[
Let, IDENTIFIER("x"), EQUAL_SIGN, INTEGER(5), PLUS_SIGN, INTEGER(5), SEMICOLON
]
