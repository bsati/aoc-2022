use crate::problem::Problem;

pub struct Day2 {
    lines: Vec<String>,
}

impl Day2 {
    pub fn new() -> Self {
        Day2 { lines: Vec::new() }
    }
}

impl Problem for Day2 {
    fn parse_input<'a>(&mut self, input: &'a str) {
        self.lines = input
            .split("\r\n")
            .map(str::to_string)
            .collect::<Vec<String>>();
    }

    fn part1(&self) {
        let score_fn = |split: Vec<&str>| match split[..] {
            ["A", "X"] => 1 + 3,
            ["A", "Y"] => 2 + 6,
            ["A", "Z"] => 3 + 0,
            ["B", "X"] => 1 + 0,
            ["B", "Y"] => 2 + 3,
            ["B", "Z"] => 3 + 6,
            ["C", "X"] => 1 + 6,
            ["C", "Y"] => 2 + 0,
            ["C", "Z"] => 3 + 3,
            _ => -1,
        };

        let score: i32 = self
            .lines
            .iter()
            .map(|a| score_fn(a.split(" ").collect::<Vec<&str>>()))
            .sum();

        println!("sum: {}", score);
    }

    fn part2(&self) {
        let score_fn = |split: Vec<&str>| match split[..] {
            ["A", "X"] => 3 + 0,
            ["A", "Y"] => 1 + 3,
            ["A", "Z"] => 2 + 6,
            ["B", "X"] => 1 + 0,
            ["B", "Y"] => 2 + 3,
            ["B", "Z"] => 3 + 6,
            ["C", "X"] => 2 + 0,
            ["C", "Y"] => 3 + 3,
            ["C", "Z"] => 1 + 6,
            _ => -1,
        };

        let score: i32 = self
            .lines
            .iter()
            .map(|a| score_fn(a.split(" ").collect::<Vec<&str>>()))
            .sum();

        println!("sum: {}", score);
    }
}
