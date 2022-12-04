from collections import defaultdict
from pathlib import Path
import time

SCRIPT_DIR = Path(__file__).parent
INPUT_FILE = Path(SCRIPT_DIR,"input.txt")

GRP_SZ = 3

def main():
    with open(INPUT_FILE,mode="rt") as f:
        data = f.read().splitlines()

    item_to_priority = {}
    for i, ordinal in enumerate(range(ord('a'),ord('z')+1), start=1):
        item_to_priority[chr(ordinal)]=i
    for i, ordinal in enumerate(range(ord('A'),ord('Z')+1), start=27):
        item_to_priority[chr(ordinal)]=i

    priorities = []
    for rucksack in data:
        compartment_1 = set(rucksack[0:len(rucksack)//2])
        compartment_2 = set(rucksack[len(rucksack)//2:])
        common = compartment_1 & compartment_2 #intersaction
        for item in common:
            priorities.append(item_to_priority[item])

    print(f"Part 1 : Sum of priorities = {sum(priorities)}")



if __name__ == "__main__":
    t1 = time.perf_counter()
    main()
