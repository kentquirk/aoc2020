#! /usr/bin/env python3
import re
import itertools


def day06a(groups):
    total = 0
    for g in groups:
        s = set()
        for c in g:
            if c in "abcdefghijklmnopqrstuvwxyz":
                s.add(c)
        total += len(s)
    return total


def day06b(groups):
    total = 0
    for g in groups:
        intersectionSet = None
        for person in g.splitlines(keepends=False):
            s = set()
            for c in person:
                s.add(c)
            if intersectionSet is None:
                intersectionSet = s
            else:
                intersectionSet = intersectionSet & s
        total += len(intersectionSet)
    return total


if __name__ == "__main__":
    f = open("./input.txt")
    data = f.read()
    groups = re.split("\n[ \t]*\n", data)
    print(f"{day06a(groups)} questions were answered in part A")
    print(
        f"{day06b(groups)} questions were answered by all members of a group in part B"
    )
