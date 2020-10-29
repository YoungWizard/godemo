package ident33

func search(nums []int, target int) int {
	n:=len(nums)
	if n==0{
		return -1
	}
	left,right:=0,n
	for left<right{
		mid:=left+(right-left)/2
		if nums[mid]==target{
			return mid
		}
		if nums[0]<=nums[mid]{
			if target<nums[mid]&&target>=nums[0]{
				right=mid
			}else{
				left=mid+1
			}
		}else{
			if target>nums[mid]&&target<=nums[n-1]{
				left=mid+1
			}else{
				right=mid
			}
		}
	}
	return -1
}