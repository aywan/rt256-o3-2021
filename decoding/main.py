import json


def to_excel_format(number: int) -> str:
    base26 = ""
    while number > 0:
        char = number % 26
        if 0 == char:
            char = 26
        number = (number - char) // 26
        base26 = chr(char + 64) + base26
    return base26
    pass


if __name__ == "__main__":
    input_str = input()
    number = int(input_str)

    answer = to_excel_format(number)

    print(json.dumps(answer))