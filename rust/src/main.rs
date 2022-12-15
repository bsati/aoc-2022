use problem::Problem;
use solutions::{Day1, Day10, Day11, Day13, Day2, Day3, Day4, Day6};
use std::{env, fs};

mod problem;
mod solutions;

fn get_solver(day: i32) -> Option<Box<dyn Problem>> {
    match day {
        1 => Some(Box::new(Day1::new())),
        2 => Some(Box::new(Day2::new())),
        3 => Some(Box::new(Day3::new())),
        4 => Some(Box::new(Day4::new())),
        6 => Some(Box::new(Day6::new())),
        10 => Some(Box::new(Day10::new())),
        11 => Some(Box::new(Day11::new())),
        13 => Some(Box::new(Day13::new())),
        _ => None,
    }
}

fn main() {
    let mut args_iter = env::args().into_iter();
    args_iter.next();
    let day = args_iter
        .next()
        .expect("Expected one argument (the day so run solver for)")
        .parse::<i32>()
        .expect("Invalid format for day supplied (integer values only)");

    let solver = get_solver(day);

    match solver {
        Some(mut solver) => {
            let filepath = format!("../inputs/day{day:02}.txt");
            let input = fs::read_to_string(filepath)
                .expect(format!("Input for day{day:02} not found").as_str());

            let solver = solver.as_mut();

            solver.parse_input(input.as_str());

            solver.part1();
            solver.part2();
        }
        None => unimplemented!("Solver for day{day:02} has not been implemented"),
    }
}
