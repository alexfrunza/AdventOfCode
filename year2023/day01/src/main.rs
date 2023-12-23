use std::env;
use std::fs;

mod second_part;

use second_part::solution;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() != 2 {
        panic!("The program must have 1 argument which is the input file!");
    }

    let input =
        fs::read_to_string(args[1].clone()).expect("Should have been able to read the file");

    solution(&input);
}
