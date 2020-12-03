#! /usr/bin/env python3
import functools


def bang_factor(lines, right, down):
    rowwidth = len(lines[0])
    collisions = 0
    posx = 0
    posy = 0
    while posy < len(lines):
        index = posx % rowwidth
        if lines[posy][index] == "#":
            collisions += 1
        posx += right
        posy += down
    return collisions


def day03a(lines):
    return bang_factor(lines, 3, 1)


def day03b(lines):
    ary = [
        bang_factor(lines, 1, 1),
        bang_factor(lines, 3, 1),
        bang_factor(lines, 5, 1),
        bang_factor(lines, 7, 1),
        bang_factor(lines, 1, 2),
    ]
    return functools.reduce(lambda a, b: a * b, ary, 1)


if __name__ == "__main__":
    f = open("./input.txt")
    lines = [l.strip() for l in f.readlines()]
    print(f"there were {day03a(lines)} collisions for part A")
    print(f"there were {day03b(lines)} collisions for part B")
