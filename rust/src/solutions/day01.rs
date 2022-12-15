use crate::problem::Problem;
use std::vec::Vec;

pub struct Day1 {
    calories: Vec<i32>,
}

impl Day1 {
    pub fn new() -> Self {
        Day1 {
            calories: Vec::new(),
        }
    }
}

impl Problem for Day1 {
    fn parse_input(&mut self, input: &str) {
        let mut calories: Vec<i32> = input
            .split("\n\r\n")
            .map(|a| {
                a.split("\r\n")
                    .filter(|s| *s != "")
                    .map(|b| b.trim().parse::<i32>().unwrap())
                    .sum()
            })
            .collect();
        calories.sort();
        self.calories = calories;
    }

    fn part1(&self) {
        println!("part 1) {}", self.calories[self.calories.len() - 1]);
    }

    fn part2(&self) {
        let len = self.calories.len();
        println!(
            "part 2) {}",
            self.calories[len - 1] + self.calories[len - 2] + self.calories[len - 3]
        );
    }
}
