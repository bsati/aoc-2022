fn main() {
    let input = include_str!("./input.txt");

    let (sum1, sum2): (i32, i32) = input
        .split("\r\n")
        .map(|l| {
            let mut split = l.trim().split(',');
            let part1 = split.next().unwrap();
            let part2 = split.next().unwrap();
            let (a_min, a_max) = to_range(part1);
            let (b_min, b_max) = to_range(part2);
            let res1 = if (a_min <= b_min && a_max >= b_max) || (a_min >= b_min && a_max <= b_max) {
                1
            } else {
                0
            };
            let res2 = if (a_min >= b_min && a_min <= b_max)
                || (a_max >= b_min && a_max <= b_max)
                || (b_min >= a_min && b_min <= a_max)
                || (b_max >= a_min && b_max <= a_max)
            {
                1
            } else {
                0
            };
            (res1, res2)
        })
        .fold((0, 0), |(a1, a2), (b1, b2)| (a1 + b1, a2 + b2));

    println!("part 1) {}", sum1);
    println!("part 2) {}", sum2);
}

fn to_range(input: &str) -> (i32, i32) {
    let res: Vec<i32> = input
        .split('-')
        .map(|a| a.parse::<i32>().unwrap())
        .collect();

    (res[0], res[1])
}
