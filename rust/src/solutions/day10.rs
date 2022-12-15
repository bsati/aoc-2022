use crate::problem::Problem;

pub struct Day10 {
    sum_part1: i32,
}

impl Day10 {
    pub fn new() -> Self {
        Day10 { sum_part1: 0 }
    }
}

fn relevant_cycle(cycle: i32) -> bool {
    match cycle {
        20 | 60 | 100 | 140 | 180 | 220 => true,
        _ => false,
    }
}

fn add_cycle(cycle: i32, x: i32) -> i32 {
    if !relevant_cycle(cycle) {
        return 0;
    }

    return cycle * x;
}

fn print_cycle(cycle: i32, x: i32) {
    if ((cycle % 40) - x).abs() <= 1 {
        print!("#");
    } else {
        print!(".");
    }
    if cycle % 40 == 0 {
        print!("\n");
    }
}

impl Problem for Day10 {
    fn parse_input(&mut self, input: &str) {
        let mut x = 1;
        let mut cycles = 0;
        let mut sum = 0;
        for line in input.split("\r\n") {
            print_cycle(cycles, x);
            let mut split = line.split(" ");
            let op = split.next().unwrap();
            if op == "noop" {
                cycles += 1;
                sum += add_cycle(cycles, x);
            } else {
                let adder = split.next().unwrap().parse::<i32>().unwrap();
                cycles += 1;
                sum += add_cycle(cycles, x);
                print_cycle(cycles, x);
                cycles += 1;
                x += adder;
                sum += add_cycle(cycles, x);
            }
        }

        self.sum_part1 = sum;
    }

    fn part1(&self) {
        println!("\npart 1) {}", self.sum_part1);
    }

    fn part2(&self) {}
}
