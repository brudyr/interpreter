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
        return lexer;
    }

    pub fn next_token(&mut self) -> Token {
        return Token::RBRACE;
    }
}

#[cfg(test)]
mod test {
    use super::{Lexer, Token};

    #[test]
    fn get_next_token() {
        let input = String::from("{}()+-,;");
        let mut lexer = Lexer::new(input);

        let tok = lexer.next_token();

        assert_eq!(tok, Token::LBRACE);
    }
}
