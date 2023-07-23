fn main() {
    println!("Hello, world!");
}

#[derive(Debug, PartialEq)]
pub enum Token {
    LET,
    IDENTIFIER(String),
    ASSIGN,
    INTEGER(i32),
    PLUS,
    MINUS,
    SEMICOLON,
    COMMA,
    ILLEGAL,
    IDENT(String),
    LPAREN,
    RPAREN,
    LBRACE,
    RBRACE,
    FUNCTION,
    EOF,
}

pub struct Lexer {
    input: Vec<u8>,
    curr: u8,
    pos: usize,
    read_pos: usize,
}

impl Lexer {
    pub fn new(input: String) -> Lexer {
        let mut lexer = Lexer {
            input: input.into_bytes(),
            curr: 0,
            pos: 0,
            read_pos: 0,
        };
        lexer.read_char();
        return lexer;
    }

    pub fn next_token(&mut self) -> Token {
        let tok = match self.curr {
            b',' => Token::COMMA,
            b';' => Token::SEMICOLON,
            b'(' => Token::LPAREN,
            b')' => Token::RPAREN,
            b'{' => Token::LBRACE,
            b'}' => Token::RBRACE,
            b'+' => Token::PLUS,
            b'-' => Token::MINUS,
            b'=' => Token::ASSIGN,
            0 => Token::EOF,
            _ => Token::ILLEGAL,
        };
        self.read_char();
        return tok;
    }

    fn read_char(&mut self) {
        if self.read_pos >= self.input.len() {
            self.curr = 0;
        } else {
            self.curr = self.input[self.read_pos];
        }
        self.pos = self.read_pos;
        self.read_pos += 1;
    }
}

#[cfg(test)]
mod test {
    use super::{Lexer, Token};

    #[test]
    fn get_next_token() {
        let input = String::from("{}()+-,;");
        let mut lexer = Lexer::new(input);

        let test_cases = vec![
            Token::LBRACE,
            Token::RBRACE,
            Token::LPAREN,
            Token::RPAREN,
            Token::PLUS,
            Token::MINUS,
            Token::COMMA,
            Token::SEMICOLON,
            Token::EOF,
        ];

        for test in test_cases {
            let next_tok = lexer.next_token();
            assert_eq!(next_tok, test);
        }
    }
}
