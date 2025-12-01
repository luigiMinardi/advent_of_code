const std = @import("std");

pub fn main() !void {
    const file_name = "example.txt";
    var file = try std.fs.cwd().openFile(file_name, .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();

    var list1: []i32 = undefined;
    var list2: []i32 = undefined;
    var buf: [1024]u8 = undefined;
    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        var first_line: bool = false;
        for (line) |el| {
            std.debug.print("{any}\n", .{el});
        }
        std.debug.print("{any}\n", .{line});
    }
    // const location_id_arr: [val]i32 = undefined;
}
