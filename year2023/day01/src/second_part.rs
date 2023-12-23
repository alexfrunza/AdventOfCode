fn get_first_digit(arr: &[i32]) -> i32 {
    let values = [1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9];

    let mut min_idx = 0;
    let mut min_value = i32::MAX;

    for (i, value) in arr.iter().enumerate() {
        if *value < min_value {
            min_value = *value;
            min_idx = i;
        }
    }

    return values[min_idx];
}

fn get_last_digit(arr: &[i32]) -> i32 {
    let values = [1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9];

    let mut max_idx = 0;
    let mut max_value = -1;

    for (i, value) in arr.iter().enumerate() {
        if *value > max_value {
            max_value = *value;
            max_idx = i;
        }
    }

    return values[max_idx];
}

pub fn solution(input: &String) {
    let mut sum: i32 = 0;
    let tokens = [
        "1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six",
        "seven", "eight", "nine",
    ];

    for line in input.lines() {
        if line.is_empty() {
            break;
        }

        let results_first = tokens.map(|s| -> i32 {
            if let Some(x) = line.find(s) {
                return x as i32;
            } else {
                return i32::MAX;
            }
        });

        let results_last = tokens.map(|s| -> i32 {
            if let Some(x) = line.rfind(s) {
                return x as i32;
            } else {
                return -1;
            }
        });

        sum += 10 * get_first_digit(&results_first) + get_last_digit(&results_last);
    }

    println!("The solution is: {}", sum);
}
