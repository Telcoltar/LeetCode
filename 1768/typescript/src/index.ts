function solution(word1: string, word2: string) {
    let len_word1 = word1.length;
    let len_word2 = word2.length;

    let longer_word: string;
    let shorter_len: number;
    let output: string = "";

    if (len_word1 < len_word2) {
        longer_word = word2;
        shorter_len = len_word1;
    } else {
        longer_word = word1;
        shorter_len = len_word2;
    }

    for (let i = 0; i < shorter_len; i++) {
        output += word1[i] + word2[i];
    }

    output += longer_word.slice(shorter_len)

    return output
}