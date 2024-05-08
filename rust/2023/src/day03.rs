use aoc_runner_derive::{aoc, aoc_generator};

#[aoc_generator(day3)]
pub fn parse(input: &str) -> Vec<Vec<char>> {
    input.lines().map(|line| line.chars().collect()).collect()
}

const NEIGHBERHOOD: [[i8; 2]; 8] = [
    [-1, 1],
    [0, 1],
    [1, 1],
    [-1, 0],
    [1, 0],
    [-1, -1],
    [0, -1],
    [1, -1],
];

#[aoc(day3, part1)]
pub fn part1(input: &Vec<Vec<char>>) -> u32 {
    let mut sum = 0;
    for (row, line) in input.iter().enumerate() {
        let mut number = 0;
        let mut is_valid = false;
        for (col, char) in line.iter().enumerate() {
            if char.is_digit(10) {
                number = number * 10 + char.to_digit(10).unwrap();
                is_valid |= NEIGHBERHOOD
                    .iter()
                    .map(|shift| {
                        let x = col as i32 + shift[0] as i32;
                        let y = row as i32 + shift[1] as i32;
                        if x < 0 || x >= line.len() as i32 || y < 0 || y >= input.len() as i32 {
                            return '.';
                        }
                        input[y as usize][x as usize]
                    })
                    .any(|neighbor| neighbor != '.' && !neighbor.is_digit(10));
                if col != line.len() - 1 {
                    continue;
                }
            }
            if is_valid {
                sum += number
            }
            number = 0;
            is_valid = false
        }
    }
    sum
}

#[aoc(day3, part2)]
pub fn part2(input: &Vec<Vec<char>>) -> u32 {
    0
}

#[cfg(test)]
mod test {
    use super::*;

    const DATA: &str = "467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..";

    #[test]
    fn test_part1() {
        assert_eq!(part1(&parse(DATA)), 4361)
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&parse(DATA)), 2286)
    }
}
