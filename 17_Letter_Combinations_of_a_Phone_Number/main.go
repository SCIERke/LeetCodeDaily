package main

var s map[string]bool;

var numberToset = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
};

func recurr(curr_word []byte ,curr_idx int ,digits *string) {
	if curr_idx == len(*digits) {
    s[string(curr_word)] = true
    return
	}

	currentDigit := (*digits)[curr_idx]
	letters := numberToset[rune(currentDigit)]

	for _, c := range []byte(letters) {
			curr_word = append(curr_word, byte(c))
			recurr(curr_word, curr_idx+1, digits)
			curr_word = curr_word[:len(curr_word)-1]
	}
}

func letterCombinations(digits string) []string {
	var init_word []byte;

	s = make(map[string]bool)
	recurr(init_word ,0 , &digits);

	rs := make([]string, 0);
	for key := range s {
		rs = append(rs, key);
		// fmt.Println(key);
	}

	return rs;
}

func main() {
	digits := "23";

	letterCombinations(digits);


}