fn main() {
    let input = include_str!("./input1.txt");

    // A, X -> Rock
    // B, Y -> Paper
    // C, Z -> Scissors

    let part1_score = |split: Vec<&str>| match split[..] {
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

    let part2_score = |split: Vec<&str>| match split[..] {
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

    let score: i32 = input
        .split("\r\n")
        .map(|a| part1_score(a.split(" ").collect::<Vec<&str>>()))
        .sum();

    println!("sum: {}", score);

    let score: i32 = input
        .split("\r\n")
        .map(|a| part2_score(a.split(" ").collect::<Vec<&str>>()))
        .sum();

    println!("sum 2: {}", score);
}
