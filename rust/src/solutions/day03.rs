use std::collections::HashSet;

use crate::problem::Problem;

pub struct Day3 {
    lines: Vec<String>,
}

impl Day3 {
    pub fn new() -> Self {
        Day3 { lines: Vec::new() }
    }
}

fn to_set(s: &str) -> HashSet<char> {
    let mut set = HashSet::new();
    for b in s.chars() {
        set.insert(b);
    }
    set
}

fn convert_to_priority(item_type: char) -> i32 {
    let sub = (item_type as i32) - ('a' as i32);
    if sub >= 0 {
        return sub + 1;
    }

    (item_type as i32) - ('A' as i32) + 27
}

fn find_item_type(first_compartment: &str, second_compartment: &str) -> i32 {
    for i in 0..first_compartment.len() {
        for j in 0..first_compartment.len() {
            if first_compartment.chars().nth(i).unwrap()
                == second_compartment.chars().nth(j).unwrap()
            {
                return convert_to_priority(first_compartment.chars().nth(i).unwrap());
            }
        }
    }
    return 0;
}

impl Problem for Day3 {
    fn parse_input(&mut self, input: &str) {
        self.lines.append(
            &mut input
                .split("\r\n")
                .map(str::to_string)
                .collect::<Vec<String>>(),
        );
    }

    fn part1(&self) {
        let sum: i32 = self
            .lines
            .iter()
            .map(|a| {
                let a = a.trim();
                let pivot = a.len() / 2;
                let first_compartment: &str = &a[0..pivot];
                let second_compartment: &str = &a[pivot..a.len()];
                find_item_type(first_compartment, second_compartment)
            })
            .sum();

        println!("part 1) {sum}");
    }

    fn part2(&self) {
        let sum: i32 = self
            .lines
            .chunks(3)
            .map(|rucksacks| {
                let r1 = to_set(rucksacks[0].as_str());
                let r2 = to_set(rucksacks[1].as_str());
                let r3 = to_set(rucksacks[2].as_str());
                convert_to_priority(
                    *r1.iter()
                        .filter(|c| r2.contains(c) && r3.contains(c))
                        .next()
                        .unwrap(),
                )
            })
            .sum();

        println!("part 2) {sum}");
    }
}
