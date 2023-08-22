from itertools import zip_longest

def solution(word1: str, word2: str):
    output: str = ""
    for (char1, char2) in zip_longest(word1, word2, fillvalue=""):
        output += char1
        output += char2
    return output
    