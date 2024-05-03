const std = @import("std");

pub fn part1(input: []const u8) usize {
    var lines = std.mem.split(u8, input, "\n");

    var sum: usize = 0;
    while (lines.next()) |line| {
        var i: usize = 0;
        while (i < line.len) : (i += 1) {
            if (line[i] >= '1' and line[i] <= '9') {
                sum += (line[i] - 48) * 10;
                break;
            }
        }

        i = line.len - 1;
        while (i >= 0) : (i -= 1) {
            if (line[i] >= '1' and line[i] <= '9') {
                sum += (line[i] - 48);
                break;
            }
        }
    }

    return sum;
}

pub fn main() void {
    const input = @embedFile("./input.txt");

    std.debug.print("{d}", .{part1(input)});
}
