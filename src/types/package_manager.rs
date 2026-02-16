use std::fmt::{Display, Formatter, Result};

#[derive(Debug, Clone, Copy, PartialEq)]
pub enum PackageManager {
    Bun,
    Npm,
    Pnpm,
    Yarn,
}

impl Display for PackageManager {
    fn fmt(&self, f: &mut Formatter) -> Result {
        write!(f, "{}", self.as_str())
    }
}

impl PackageManager {
    pub const ALL: &'static [PackageManager] = &[
        PackageManager::Npm,
        PackageManager::Yarn,
        PackageManager::Pnpm,
        PackageManager::Bun,
    ];

    pub fn as_str(&self) -> &str {
        match self {
            PackageManager::Npm => "npm",
            PackageManager::Yarn => "yarn",
            PackageManager::Pnpm => "pnpm",
            PackageManager::Bun => "bun",
        }
    }

    pub fn from_str(s: &str) -> Option<Self> {
        match s {
            "npm" => Some(PackageManager::Npm),
            "yarn" => Some(PackageManager::Yarn),
            "pnpm" => Some(PackageManager::Pnpm),
            "bun" => Some(PackageManager::Bun),
            _ => None,
        }
    }
}
