file = "test.txt"


def cal(input, result):
    base_score = ord(input) - ord('A') + 1
    bonus = 0
    if result == "X":  # lost
        base_score -= 1
    elif result == "Y":  # draw
        bonus = 3
    elif result == "Z":  # win
        bonus = 6
        base_score += 1

    if base_score == 0:
        base_score = 3
    if base_score == 4:
        base_score = 1
    # print(">> ", base_score, bonus)
    return base_score + bonus


with open(file, 'r') as f:
    sum = 0
    for line in f:
        list = line.split()
        if len(list) >= 2:
            sum += cal(list[0], list[1])

print(sum)
