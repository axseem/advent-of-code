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

fn get_numbers_entries(input: &Vec<Vec<char>>, row: i32, col: i32) -> Option<Vec<(u32, u32)>> {
    let mut entries: Vec<(u32, u32)> = Vec::new();

    for shift in -1..=1 {
        if row + shift >= 0 && row + shift < input.len() as i32 {
            if input[(row + shift) as usize][col as usize].is_digit(10) {
                entries.push(((row + shift) as u32, col as u32))
            } else {
                if col - 1 >= 0 && input[(row + shift) as usize][(col - 1) as usize].is_digit(10) {
                    entries.push(((row + shift) as u32, (col - 1) as u32))
                }
                if col + 1 < input[0].len() as i32
                    && input[(row + shift) as usize][(col + 1) as usize].is_digit(10)
                {
                    entries.push(((row + shift) as u32, (col + 1) as u32))
                }
            }
        }
    }

    return match entries.len() {
        2 => Some(entries),
        _ => None,
    };
}

fn get_number(line: &Vec<char>, entry: u32) -> u32 {
    let mut start: usize = 0;
    let mut end = line.len() - 1;
    for i in (0..=entry).rev() {
        if !line[i as usize].is_digit(10) {
            break;
        }
        start = i as usize;
    }
    for i in entry..line.len() as u32 {
        if !line[i as usize].is_digit(10) {
            break;
        }
        end = i as usize;
    }
    line[start..=end]
        .iter()
        .collect::<String>()
        .parse::<u32>()
        .unwrap()
}

#[aoc(day3, part2)]
pub fn part2(input: &Vec<Vec<char>>) -> u32 {
    let mut sum = 0;
    for (row, line) in input.iter().enumerate() {
        for (col, char) in line.iter().enumerate() {
            if *char == '*' {
                if let Some(entries) = get_numbers_entries(input, row as i32, col as i32) {
                    let n1 = get_number(&input[entries[0].0 as usize], entries[0].1);
                    let n2 = get_number(&input[entries[1].0 as usize], entries[1].1);
                    sum += n1 * n2;
                }
            }
        }
    }
    sum
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
        assert_eq!(part2(&parse(DATA)), 467835)
    }
}
