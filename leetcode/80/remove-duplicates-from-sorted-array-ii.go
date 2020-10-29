package ident80

func removeDuplicates(nums []int) int {
	n:=len(nums)
	if n==0{
		return 0
	}
	i:=0
	count:=1
	for j:=1;j<n;j++{
		if nums[i]!=nums[j]{
			i++
			nums[i]=nums[j]
			count=1
		}else{
			if count==1{
				i++
				nums[i]=nums[j]
				count--
			}
		}
	}
	return i+1
}