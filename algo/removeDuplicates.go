func removeDuplicates(nums []int) int {
	lens = len(nums)
    for i := lens - 1; i > 0; i-- {
    	if nums[i] == nums[i-1] {
    		for j := i; j < lens; j++ {
    			nums[j-1] = nums[j]
    		}
    		lens--
    		nums = nums[0:lens]
    	}
    }
    return lens
}