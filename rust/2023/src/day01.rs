#[aoc(day1, part1)]
pub fn part1(input: &str) -> u32 {
    let lines = input.lines();

    let mut sum: u32 = 0;
    for line in lines {
        for byte in line.bytes() {
            if byte >= '0' as u8 && byte <= '9' as u8{
                sum += ((byte - '0' as u8)*10) as u32;
                break;
            }
        }

        for byte in line.bytes().rev() {
            if byte >= '0' as u8 && byte <= '9' as u8{
                sum += (byte - '0' as u8) as u32;
                break;
            }
        }
    }

    sum
}

#[aoc(day1, part2)]
pub fn part2(input: &str) -> u32 {
    let lines = input.lines().map(|l| l.as_bytes());
    let digits: [&str; 9] = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"];

    let mut sum: u32 = 0;
    for line in lines {

        'first: for (i, byte) in line.into_iter().enumerate() {
            if *byte >= '0' as u8 && *byte <= '9' as u8{
                sum += ((byte - '0' as u8)*10) as u32;
                break;
            }

            for (j, digit) in digits.iter().enumerate(){
                for (k, digit_byte) in digit.as_bytes().iter().enumerate(){
                    if *digit_byte != line[i+k] {
                        break;
                    } 
                    if digit.len()-1 == k {
                        sum += ((j+1)*10) as u32;
                        break 'first;
                    }
                }
            }
        }

        'second: for (i, byte) in line.into_iter().rev().enumerate() {
            if *byte >= '0' as u8 && *byte <= '9' as u8{
                sum += (byte - '0' as u8) as u32;
                break;
            }

            for (j, digit) in digits.iter().enumerate(){
                for (k, digit_byte) in digit.as_bytes().iter().rev().enumerate(){
                    if i+k > line.len()-1 || *digit_byte != line[line.len()-1-i-k] {
                        break;
                    } 
                    if digit.len()-1 == k {
                        sum += (j+1) as u32;
                        break 'second;
                    }
                }
            }
        }
    }

    sum
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let data = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";

        assert_eq!(part1(data), 142)
    }

    #[test]
    fn test_part2() {
        let data = "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen";

        assert_eq!(part2(data), 281)
    }
}