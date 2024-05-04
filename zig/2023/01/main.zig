const std = @import("std");

pub fn part1(input: []const u8) usize {
    var lines = std.mem.split(u8, input, "\n");

    var sum: usize = 0;
    while (lines.next()) |line| {
        var i: usize = 0;
        while (i < line.len) : (i += 1) {
            const c: u8 = line[i];
            if (c >= '1' and c <= '9') {
                sum += (c - 48) * 10;
                break;
            }
        }

        i = line.len - 1;
        while (i >= 0) : (i -= 1) {
            const c: u8 = line[i];
            if (c >= '1' and c <= '9') {
                sum += (c - 48);
                break;
            }
        }
    }

    return sum;
}

pub fn part2(input: []const u8) usize {
    var lines = std.mem.split(u8, input, "\n");
    const digits = [_][]const u8{ "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" };

    var sum: usize = 0;
    while (lines.next()) |line| {
        var i: usize = 0;
        first: while (i < line.len) : (i += 1) {
            const c: u8 = line[i];
            if (c >= '1' and c <= '9') {
                sum += (c - 48) * 10;
                break;
            }

            for (digits, 0..) |digit, j| {
                var k: usize = 0;
                while (line[i + k] == digit[k]) : (k += 1) {
                    if (digit.len - 1 == k) {
                        sum += (j + 1) * 10;
                        break :first;
                    }
                }
            }
        }

        i = line.len - 1;
        last: while (i >= 0) : (i -= 1) {
            const c: u8 = line[i];
            if (c >= '1' and c <= '9') {
                sum += (c - 48);
                break;
            }

            for (digits, 0..) |digit, j| {
                var k: usize = 0;
                while (line[i - k] == digit[digit.len - 1 - k]) : (k += 1) {
                    if (digit.len - 1 == k) {
                        sum += j + 1;
                        break :last;
                    }
                }
            }
        }
    }

    return sum;
}

pub fn main() void {
    const input = @embedFile("./input.txt");

    std.debug.print("--- 2023 day 01 answer ---\n", .{});
    std.debug.print("part 1:\t{d}\n", .{part1(input)});
    std.debug.print("part 2:\t{d}\n", .{part2(input)});
}
