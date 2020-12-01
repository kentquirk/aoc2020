#! /usr/bin/env python3
import itertools

target = 2020

# part A - find a pair of numbers that add to 2020 and emit their product
def day1a(data):
    # create a list of differences between 2020 and the items in the list
    diffs = [target - d for d in data]
    # find the set intersection between that and the data -- these are the only possibilities
    candidates = list(set(data).intersection(diffs))
    print(candidates, candidates[0] * candidates[1])


# part B - find a triplet of numbers that add to 2020 and emit that product
def day1b(data):
    # create the product of all values that are not themselves
    data2 = [x[0] + x[1] for x in itertools.product(data, data) if x[0] != x[1]]
    # then follow the plan above -- except now we have 3 values instead of 2
    diffs = [target - d for d in data2]
    candidates = list(set(data).intersection(diffs))
    print(candidates, candidates[0] * candidates[1] * candidates[2])


if __name__ == "__main__":
    f = open("./input.txt")
    lines = f.readlines()
    data = [int(l) for l in lines]
    day1a(data)
    day1b(data)
