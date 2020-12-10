#! /usr/bin/env python3
import itertools
from collections import defaultdict


def day10a(data):
    diffs = [b - a for a, b in zip(data[:-1], data[1:])]
    results = defaultdict(int)
    for d in diffs:
        results[d] += 1
    print(results)
    return results[1] * results[3]


if __name__ == "__main__":
    f = open("./input.txt")
    lines = f.readlines()
    data = sorted([int(l) for l in lines])
    data.append(data[-1] + 3)
    data.append(0)
    data.sort()
    print(f"product is {day10a(data)}")