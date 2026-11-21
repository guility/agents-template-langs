//! DDD Rust Template
//!
//! Domain-Driven Design с чистой архитектурой

pub mod application;
pub mod domain;
pub mod infrastructure;
pub mod interface;

// Re-exports
pub use application::*;
pub use domain::*;
pub use infrastructure::*;
pub use interface::*;
