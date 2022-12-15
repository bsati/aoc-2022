use crate::problem::Problem;

pub struct Day6 {
    input_chars: Vec<char>,
}

impl Day6 {
    pub fn new() -> Self {
        Day6 {
            input_chars: Vec::new(),
        }
    }

    fn find_marker<const M: usize>(&self) -> i32 {
        let mut previous: Vec<char> = Vec::with_capacity(M);
        for (i, c) in self.input_chars.iter().enumerate() {
            if previous.len() == M {
                previous.remove(0);
            }
            previous.push(*c);

            let mut cloned = previous.clone();
            cloned.sort();
            cloned.dedup();
            if cloned.len() == M {
                return i as i32 + 1;
            }
        }

        -1
    }
}

impl Problem for Day6 {
    fn parse_input(&mut self, input: &str) {
        self.input_chars = input.chars().collect();
    }

    fn part1(&self) {
        println!("part 1) {}", self.find_marker::<4>());
    }

    fn part2(&self) {
        println!("part 2) {}", self.find_marker::<14>());
    }
}
