#! /usr/bin/env python3


def transform(subject, target):
    return (target * subject) % 20201227


def loop(subject, loopsize):
    target = 1
    for i in range(loopsize):
        target = transform(subject, target)
    return target


def find(subject, public):
    target = 1
    i = 1
    while True:
        target = transform(subject, target)
        if target == public:
            return i
        i += 1


def day25a(subject, cardpub, doorpub):
    cardloop = find(subject, cardpub)
    doorloop = find(subject, doorpub)
    print(cardloop, doorloop)
    print(loop(cardpub, doorloop))
    print(loop(doorpub, cardloop))


if __name__ == "__main__":
    day25a(7, 6929599, 2448427)  # real
    # day25a(7, 5764801, 17807724) # sample
