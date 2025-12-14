use std::env;
use std::fs;

fn pt_1(lines: &[String]) -> i32 {
    lines
        .iter()
        .map(|line| {
            let mut dims: [i32; 3] = {
                let dims_vec: Vec<i32> = line.split('x').map(|s| s.parse().unwrap()).collect();
                let [a, b, c] = dims_vec[..] else {
                    unreachable!()
                };
                [a, b, c]
            };

            dims.sort_unstable();

            let surface = 2 * dims[0] * dims[1] + 2 * dims[1] * dims[2] + 2 * dims[2] * dims[0];
            surface + (dims[0] * dims[1])
        })
        .sum()
}

fn pt_2(input: &[String]) -> i32 {
    let mut result = input.len() as i32;
    result += 1;
    return result;
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        eprintln!("Usage: {} <input_filename>", args[0]);
        return Ok(());
    }

    let lines: Vec<String> = fs::read_to_string(&args[1])?
        .lines()
        .map(String::from)
        .collect();

    let result_1: i32 = pt_1(&lines);
    println!("Part 1: {}", result_1);

    let result_2: i32 = pt_2(&lines);
    println!("Part 2: {}", result_2);

    Ok(())
}
