fn main() {
    let input = include_str!("./input1.txt");
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
    println!("Max calories: {}", calories[calories.len() - 1]);
    println!(
        "Max calories sum: {}",
        calories[calories.len() - 1] + calories[calories.len() - 2] + calories[calories.len() - 3]
    );
}
