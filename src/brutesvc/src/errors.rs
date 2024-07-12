use std::fmt;
use std::io;

#[derive(Debug)]
pub enum SourceFileError {
    OpenError(io::Error),
    ReadError(io::Error),
}

impl fmt::Display for SourceFileError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            SourceFileError::OpenError(e) => write!(f, "Failed to open the file: {}", e),
            SourceFileError::ReadError(e ) => write!(f, "Failed to read the file: {}", e),
        }
    }
}

impl From<io::Error> for SourceFileError {
    fn from(value: io::Error) -> Self {
        SourceFileError::OpenError(value)
    } 
}
