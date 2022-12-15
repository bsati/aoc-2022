pub trait Problem {
    fn parse_input(&mut self, input: &str);

    fn part1(&self);

    fn part2(&self);
}
