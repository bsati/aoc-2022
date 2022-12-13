use std::cmp::Ordering;
use std::string::String;
use std::vec::Vec;

#[derive(Debug, Clone, PartialEq, Eq)]
enum Packet {
    Value(i32),
    List(Vec<Packet>),
}

impl PartialOrd for Packet {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        use Packet::*;

        match (self, other) {
            (Value(val1), Value(val2)) => val1.partial_cmp(val2),
            (List(l1), List(l2)) => l1
                .iter()
                .zip(l2)
                .fold(None, |ord, (p1, p2)| {
                    ord.or(match p1.cmp(p2) {
                        Ordering::Less => Some(Ordering::Less),
                        Ordering::Equal => None,
                        Ordering::Greater => Some(Ordering::Greater),
                    })
                })
                .or_else(|| l1.len().partial_cmp(&l2.len())),
            (List(_), Value(_)) => self.partial_cmp(&Packet::List(vec![other.clone()])),
            (Value(_), List(_)) => Packet::List(vec![self.clone()]).partial_cmp(other),
        }
    }
}

impl Ord for Packet {
    fn cmp(&self, other: &Self) -> Ordering {
        self.partial_cmp(other).unwrap()
    }
}

fn find_next_correct_closing_bracket(s: &str) -> usize {
    let left = s.find(']').unwrap();
    if left == s.rfind(']').unwrap() {
        return left;
    }
    let mut stack = Vec::new();
    stack.push('[');
    for (i, c) in s[1..].chars().enumerate() {
        match c {
            '[' => stack.push('['),
            ']' => {
                stack.pop();
                if stack.len() == 0 {
                    return i + 1;
                }
            }
            _ => {}
        }
    }
    return 0;
}

fn parse_packet(input: &str) -> Packet {
    let mut children = parse_data(input);
    if children.len() == 1 {
        return children.remove(0);
    }
    Packet::List(children)
}

fn parse_data(s: &str) -> Vec<Packet> {
    let stripped = s
        .strip_prefix('[')
        .and_then(|s| s.strip_suffix(']'))
        .unwrap_or(s);

    let mut collected = Vec::new();

    let mut current = String::new();

    let chars = stripped.chars().collect::<Vec<char>>();
    let mut i = 0;
    while i < chars.len() {
        match chars[i] {
            '0'..='9' => current.push(chars[i]),
            ',' => {
                if let Ok(parsed) = current.parse::<i32>() {
                    collected.push(Packet::Value(parsed));
                }
                current.clear();
            }
            '[' => {
                let right_index = find_next_correct_closing_bracket(&stripped[i..]);
                let children = parse_data(&stripped[i..=i + right_index]);
                collected.push(Packet::List(children));
                i += right_index;
            }
            _ => {
                println!("Non-matched char: {}", chars[i]);
            }
        }
        i += 1;
    }

    if current.len() > 0 {
        collected.push(Packet::Value(current.parse::<i32>().unwrap()));
    }

    collected
}

fn main() {
    let input = include_str!("./input.txt");

    let packets = input
        .split("\r\n\r\n")
        .enumerate()
        .map(|(i, pair)| {
            let mut split = pair.split("\r\n");
            let first_packet = split.next().unwrap();
            let second_packet = split.next().unwrap();

            let packet1 = parse_packet(first_packet);
            let packet2 = parse_packet(second_packet);

            (packet1, packet2, i + 1)
        })
        .collect::<Vec<(Packet, Packet, usize)>>();

    let sum: usize = packets
        .iter()
        .filter(|(packet1, packet2, _index)| packet1 < packet2)
        .map(|(_packet1, _packet2, index)| index)
        .sum();

    println!("part 1) {}", sum);

    let divider1 = Packet::List(vec![Packet::List(vec![Packet::Value(2)])]);
    let divider2 = Packet::List(vec![Packet::List(vec![Packet::Value(6)])]);

    let mut simplified = packets
        .into_iter()
        .flat_map(|(packet1, packet2, _index)| [packet1, packet2])
        .chain([divider1.clone(), divider2.clone()])
        .collect::<Vec<Packet>>();

    simplified.sort();

    let product: usize = simplified
        .iter()
        .enumerate()
        .map(|(i, packet)| {
            if packet == &divider1 || packet == &divider2 {
                return i + 1;
            }
            1
        })
        .product();

    println!("part 2) {}", product);
}
