package ident45



func canJump(nums []int) bool {
	n,farthest:=len(nums),0
	for i:=0;i<n-1;i++{
		farthest=Max(farthest,i+nums[i])
		if farthest<=i{
			return false
		}
	}
	if(farthest>=n-1){
		return true
	}
	return false
}
func Max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

