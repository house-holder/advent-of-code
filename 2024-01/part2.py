import sys

def buildLists(lines: list[str]) -> tuple[list[int], list[int]]:
    listL: list[int] = []
    listR: list[int] = []

    for line in lines:
        left, right = line.strip().split("   ")
        listL.append(int(left))
        listR.append(int(right)) 

    return sorted(listL), sorted(listR)

def main() -> None:
    filename: str = "input.txt"
    if "test" in sys.argv:
        filename = "example.txt"

    with open(filename, "r") as file:
        lines = file.readlines()
    listL, listR = buildLists(lines)

    score: int = 0
    for i in range(len(listL)):
        target: int = listL[i]
        count: int = 0

        for j in range(len(listR)):
            current: int = listR[j]
            if current == target:
                count += 1 
        score += target * count

    print(score)

if __name__ == "__main__":
    main()
