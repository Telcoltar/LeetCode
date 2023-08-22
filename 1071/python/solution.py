def solution(str1: str, str2: str) -> str:
    len_str1 = len(str1)
    len_str2 = len(str2)
    min_len: int = min(len_str1, len_str2)
    for i in range(min_len, 0, -1):
        sub_str1 = str1[:i]
        sub_str2 = str2[:i]
        if sub_str1 == sub_str2:
            if len_str1 % i == 0 and len_str2 % i == 0:
                if str1 == sub_str1 * (len_str1 // i) and str2 == sub_str2 * (len_str2 // i):
                    return sub_str1
    return ""
