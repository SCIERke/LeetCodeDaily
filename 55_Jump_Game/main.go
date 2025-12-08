package main

import "fmt"

func jumpStep(curr_num int ,max_reach int ,nums []int ,state []int) bool {
	if curr_num == len(nums) - 1 {
		return true;
	}

	if state[curr_num] != -1 {
		switch state[curr_num] {
			case 0:
				return false
			case 1:
				return true
		}
	}


	for i := 1; i <= max_reach ;i++ {
		if (curr_num + i < len(nums)) {
			if (jumpStep(curr_num + i , nums[curr_num + i],nums ,state)) {
				state[curr_num] = 1;
				return true;
			}
		}
	}
	state[curr_num] = 0;

	return false;
}

func canJump(nums []int) bool {
	curr_idx := 0;
	max_reach := nums[curr_idx];

	state := make([]int ,len(nums));
	for i := 0; i < len(nums);i++ {
		state[i] = -1; // -1 == not visited
	}

	return jumpStep(curr_idx ,max_reach ,nums ,state);
}

func main() {
	nums := []int{2,3,1,1,4};

	fmt.Println(canJump(nums));
}
