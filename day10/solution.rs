use std::io::{self, Write};

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

fn main() {
    let input = include_str!("./input.txt");
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

    io::stdout().flush().unwrap();
    println!("");
    println!("part 1) {}", sum);
}
