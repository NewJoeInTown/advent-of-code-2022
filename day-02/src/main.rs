use std::{error::Error, fs, str::FromStr};

fn main() {
    println!("Part 1: {}", solve_part_1(&file("input")));
    println!("Part 2: {}", solve_part_2(&file("input")));
}

fn solve_part_1(input: &str) -> u32 {
    input
        .lines()
        .map(|line| {
            let (elf, me) = line.split_once(' ').unwrap();
            let elf_shape = Shape::from_str(elf).unwrap();
            let me_shape = Shape::from_str(me).unwrap();
            me_shape.score() + Outcome::for_me(&elf_shape, &me_shape).score()
        })
        .sum()
}

fn solve_part_2(input: &str) -> u32 {
    input
        .lines()
        .map(|line| {
            let (elf, outcome) = line.split_once(' ').unwrap();
            let elf_shape = Shape::from_str(elf).unwrap();
            let outcome_type = Outcome::from_str(outcome).unwrap();
            outcome_type.with_other(&elf_shape).score() + outcome_type.score()
        })
        .sum()
}

enum Shape {
    Rock,
    Paper,
    Scissors,
}

impl FromStr for Shape {
    type Err = Box<dyn Error>;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "A" | "X" => Ok(Shape::Rock),
            "B" | "Y" => Ok(Shape::Paper),
            "C" | "Z" => Ok(Shape::Scissors),
            _ => panic!("Invalid shape: {}", s),
        }
    }
}

impl Shape {
    fn score(&self) -> u32 {
        match self {
            Shape::Rock => 1,
            Shape::Paper => 2,
            Shape::Scissors => 3,
        }
    }
}

enum Outcome {
    Lose,
    Draw,
    Win,
}

impl FromStr for Outcome {
    type Err = Box<dyn Error>;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "X" => Ok(Outcome::Lose),
            "Y" => Ok(Outcome::Draw),
            "Z" => Ok(Outcome::Win),
            _ => panic!("{}", s),
        }
    }
}

impl Outcome {
    fn for_me(other: &Shape, me: &Shape) -> Self {
        match (other, me) {
            (Shape::Rock, Shape::Rock) => Outcome::Draw,
            (Shape::Rock, Shape::Paper) => Outcome::Win,
            (Shape::Rock, Shape::Scissors) => Outcome::Lose,

            (Shape::Paper, Shape::Rock) => Outcome::Lose,
            (Shape::Paper, Shape::Paper) => Outcome::Draw,
            (Shape::Paper, Shape::Scissors) => Outcome::Win,

            (Shape::Scissors, Shape::Rock) => Outcome::Win,
            (Shape::Scissors, Shape::Paper) => Outcome::Lose,
            (Shape::Scissors, Shape::Scissors) => Outcome::Draw,
        }
    }

    fn with_other(&self, other: &Shape) -> Shape {
        match self {
            Outcome::Lose => match other {
                Shape::Rock => Shape::Scissors,
                Shape::Paper => Shape::Rock,
                Shape::Scissors => Shape::Paper,
            },
            Outcome::Draw => match other {
                Shape::Rock => Shape::Rock,
               
