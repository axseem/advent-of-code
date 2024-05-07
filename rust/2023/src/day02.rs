use std::cmp::max;

use aoc_runner_derive::{aoc, aoc_generator};

#[aoc_generator(day2)]
pub fn parse(input: &str) -> Vec<Vec<[u8; 3]>> {
    input
        .lines()
        .map(|line| {
            line.split_once(": ")
                .unwrap()
                .1
                .split("; ")
                .map(|set| {
                    let mut array: [u8; 3] = [0; 3];

                    for cubes in set.split(", ") {
                        let (amount, color) = cubes.split_once(' ').unwrap();

                        match color {
                            "red" => array[0] = amount.parse().unwrap(),
                            "green" => array[1] = amount.parse().unwrap(),
                            "blue" => array[2] = amount.parse().unwrap(),
                            _ => {}
                        }
                    }

                    array
                })
                .collect::<Vec<[u8; 3]>>()
        })
        .collect::<Vec<Vec<[u8; 3]>>>()
}

#[aoc(day2, part1)]
pub fn part1(input: &Vec<Vec<[u8; 3]>>) -> usize {
    const MAX: [u8; 3] = [12, 13, 14];

    input
        .iter()
        .enumerate()
        .map(|(i, game)| {
            if game
                .iter()
                .any(|set| set[0] > MAX[0] || set[1] > MAX[1] || set[2] > MAX[2])
            {
                return 0;
            }
            i + 1
        })
        .sum()
}

#[aoc(day2, part2)]
pub fn part2(input: &Vec<Vec<[u8; 3]>>) -> usize {
    input
        .iter()
        .map(|game| {
            let mut amount: [usize; 3] = [0; 3];
            for set in game {
                amount[0] = max(amount[0], set[0] as usize);
                amount[1] = max(amount[1], set[1] as usize);
                amount[2] = max(amount[2], set[2] as usize);
            }
            amount[0] * amount[1] * amount[2]
        })
        .sum()
}

#[cfg(test)]
mod test {
    use super::*;

    const DATA: &str = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green";

    #[test]
    fn test_part1() {
        assert_eq!(part1(&parse(DATA)), 8)
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&parse(DATA)), 2286)
    }
}
