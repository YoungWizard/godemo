package ident1259

func numberOfWays(num_people int) int {
	mod:=1000000007
	n:=num_people
	dp:=make([]int,n+1)
	dp[0]=1 
	dp[2]=1
	for i:=2;i<n+1;i+=2{
		res:=0
		for j:=1;j<i;j+=2{
			res=(res+(dp[j-1]%mod)*(dp[i-j-1]%mod))%mod
		}
		dp[i]=res
	}
	return dp[n]
}