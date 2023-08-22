class Solution {

    public static void main(String[] args) {
        System.out.println(solution("abc", "def"));
        System.out.println(solution("abc", "defgh"));
        System.out.println(solution("abcd", "ef"));
        System.out.println(solution("abcde", "fgh"));
        System.out.println(solution("abcde", "fghij"));
    }

    static String solution(String word1, String word2) {
        int word1_length = word1.length();
        int word2_length = word2.length();

        int shorter_length;
        String longer_word;

        if (word1_length > word2_length) {
            shorter_length = word2_length;
            longer_word = word1;
        } else {
            shorter_length = word1_length;
            longer_word = word2;
        }

        StringBuilder sb = new StringBuilder();
        for(int i = 0; i < shorter_length; i++) {
            sb.append(word1.charAt(i));
            sb.append(word2.charAt(i));
        }

        sb.append(longer_word.substring(shorter_length));

        return sb.toString();
    }
}