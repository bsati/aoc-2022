use crate::problem::Problem;

#[derive(Clone)]
struct Monkey {
    items: Vec<i64>,
    test: i64,
    operation: Op,
    dest_true: usize,
    dest_false: usize,
    counter: i64,
}

impl Monkey {
    fn inspect_items<F>(&mut self, reducer: F) -> (Vec<i64>, Vec<i64>)
    where
        F: Fn(i64) -> i64,
    {
        let mut dest_true = Vec::new();
        let mut dest_false = Vec::new();

        for item in &self.items {
            self.counter += 1;
            let new_worry_level = reducer(self.operation.apply(*item));
            if new_worry_level % self.test == 0 {
                dest_true.push(new_worry_level);
            } else {
                dest_false.push(new_worry_level);
            }
        }

        self.items.clear();

        (dest_true, dest_false)
    }
}

#[derive(Clone)]
enum Op {
    Add(i64),
    Multiply(i64),
    Square,
}

impl Op {
    fn apply(&self, old: i64) -> i64 {
        match self {
            Op::Add(val) => old + val,
            Op::Multiply(val) => old * val,
            Op::Square => old * old,
        }
    }
}

fn extract_line_info(line: &str) -> &str {
    let (_, info) = line.split_once(':').unwrap();
    &info[1..]
}

fn parse_monkey(input: &str) -> Monkey {
    let lines = input.split("\r\n").collect::<Vec<&str>>();

    let mut items = Vec::new();
    extract_line_info(lines[1])
        .split(", ")
        .map(|a| a.parse::<i64>().unwrap())
        .for_each(|a| items.push(a));

    Monkey {
        items,
        test: lines[3].split(" ").last().unwrap().parse::<i64>().unwrap(),
        operation: parse_operation(extract_line_info(lines[2])),
        dest_true: lines[4]
            .split(" ")
            .last()
            .unwrap()
            .parse::<usize>()
            .unwrap(),
        dest_false: lines[5]
            .split(" ")
            .last()
            .unwrap()
            .parse::<usize>()
            .unwrap(),
        counter: 0,
    }
}

fn parse_operation(s: &str) -> Op {
    let split = s.split(" ").collect::<Vec<&str>>();

    if split[4] == "old" {
        return Op::Square;
    }

    let other = split[4].parse::<i64>().unwrap();
    match split[3] {
        "+" => Op::Add(other),
        _ => Op::Multiply(other),
    }
}

pub struct Day11 {
    monkeys: Vec<Monkey>,
}

impl Day11 {
    pub fn new() -> Self {
        Day11 {
            monkeys: Vec::new(),
        }
    }

    fn runner<F>(&self, reducer: F, rounds: i32) -> Vec<i64>
    where
        F: Fn(i64) -> i64,
    {
        let mut monkeys = self.monkeys.clone();
        for _round in 0..rounds {
            for k in 0..monkeys.len() {
                let monkey = &mut monkeys[k];
                let dest_true = monkey.dest_true;
                let dest_false = monkey.dest_false;
                let (mut pos_test, mut neg_test) = monkey.inspect_items(&reducer);
                monkeys[dest_true].items.append(&mut pos_test);
                monkeys[dest_false].items.append(&mut neg_test);
            }
        }

        let mut scores = monkeys.iter().map(|m| m.counter).collect::<Vec<i64>>();
        scores.sort();

        scores
    }
}

impl Problem for Day11 {
    fn parse_input(&mut self, input: &str) {
        self.monkeys = input
            .split("\r\n\r\n")
            .map(parse_monkey)
            .collect::<Vec<Monkey>>();
    }

    fn part1(&self) {
        let scores = self.runner(|w| w / 3, 20);
        let len = scores.len();
        println!("part 1) {}", scores[len - 1] * scores[len - 2]);
    }

    fn part2(&self) {
        let lcm = self.monkeys.iter().map(|m| m.test).product::<i64>();
        let scores = self.runner(|w| w % lcm, 10000);
        let len = scores.len();
        println!("part 2) {}", scores[len - 1] * scores[len - 2]);
    }
}
