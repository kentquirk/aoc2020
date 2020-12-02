#! /usr/bin/env python3
import re
import itertools


def validateA(d):
    n = d["password"].count(d["ch"])
    return n >= int(d["ix1"]) and n <= int(d["ix2"])


def validateB(d):
    ix1 = int(d["ix1"]) - 1
    ix2 = int(d["ix2"]) - 1
    at1 = ix1 < len(d["password"]) and d["password"][ix1 : ix1 + 1] == d["ch"]
    at2 = ix2 < len(d["password"]) and d["password"][ix2 : ix2 + 1] == d["ch"]
    return at1 != at2


if __name__ == "__main__":
    f = open("./input.txt")
    lines = f.readlines()
    pat = re.compile(
        "(?P<ix1>[0-9]+)-(?P<ix2>[0-9]+) (?P<ch>[a-z]): (?P<password>[a-z]+)"
    )
    data = [pat.match(l).groupdict() for l in lines]
    valids = [x for x in data if validateA(x)]
    print(len(valids))
    valids2 = [x for x in data if validateB(x)]
    print(len(valids2))