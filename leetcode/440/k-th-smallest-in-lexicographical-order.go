package ident440

func findKthNumber(n int, k int) int {
	cur:=1
	p:=1
	for p<k{
		count:=getCount(cur,n)
		if p+count<=k{
			cur++
			p+=count
		}else{
			cur*=10
			p++
		}
	}
	return cur

}
func getCount(prefix,n int)int{
	count:=0
	for left,right:=prefix,prefix+1;
	left<=n;
	left,right=left*10,right*10{
		count+=min(right,n+1)-left
	}
	return count
}
func min(a,b int)int{
	if(a>b){
		return b
	}
	return a
}