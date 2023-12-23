pub fn solution(input: &String) {
    let mut sum: u32 = 0;

    for line in input.lines() {
        if line.is_empty() {
            break;
        }

        let digits: Vec<_> = line.chars().filter(|c| c.is_digit(10)).collect();

        let first_digit = digits[0] as u32 - '0' as u32;
        let last_digit = digits[digits.len() - 1] as u32 - '0' as u32;

        let n = first_digit * 10 + last_digit;

        sum += n;
    }

    println!("The solution is: {}", sum);
}
