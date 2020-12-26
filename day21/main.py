#! /usr/bin/env python3
import re
from collections import defaultdict

if __name__ == "__main__":
    f = open("./input.txt")
    lines = f.readlines()
    by_food = defaultdict(dict)
    by_word = defaultdict(set)
    allwords = defaultdict(int)
    for i in range(len(lines)):
        line = lines[i]
        rawwords, rawfoods = line.split(" (contains ")
        words = rawwords.split(" ")
        foods = [f for f in re.split("[^a-z]+", rawfoods) if f != ""]
        for w in words:
            allwords[w] += 1
            by_word[w].update(foods)
        for f in foods:
            by_food[f][i] = set(words)

    # print(f"food: {by_food}")
    # print(f"allwords: {allwords}")
    # print(f"word: {by_word}")

    possibles = dict()
    for food in by_food:
        groups = by_food[food]
        possible = None
        impossible = None
        for ix in groups:
            words = groups[ix]
            if possible is None:
                possible = words
            else:
                possible = possible.intersection(words)
        possibles[food] = possible
    print(f"{possibles}")

    impossibles = {w for w in allwords}
    for p in possibles:
        impossibles = impossibles - possibles[p]

    print(f"impossibles: {impossibles}")

    print(sum([allwords[w] for w in impossibles]))

    allergens = dict()
    while len(possibles) > 0:
        deletes = []
        for p in possibles:
            if len(possibles[p]) == 1:
                allergens[p] = list(possibles[p])[0]
                deletes.append(p)
        for d in deletes:
            del possibles[d]
        for p in possibles:
            possibles[p] = possibles[p] - set(allergens.values())

    print(",".join([allergens[k] for k in sorted(allergens.keys())]))