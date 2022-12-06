use std::vec::Vec;

fn main() {
    let buffer = include_str!("./input.txt");

    let part1 = find_marker::<4>(buffer);
    let part2 = find_marker::<14>(buffer);

    println!("part 1) {}", part1);
    println!("part 2) {}", part2);
}

fn find_marker<const M: usize>(input: &str) -> i32 {
    let mut previous: Vec<char> = Vec::with_capacity(M);
    for (i, c) in input.chars().enumerate() {
        if previous.len() == M {
            previous.remove(0);
        }
        previous.push(c);

        let mut cloned = previous.clone();
        cloned.sort();
        cloned.dedup();
        if cloned.len() == M {
            return i as i32 + 1;
        }
    }

    -1
}
