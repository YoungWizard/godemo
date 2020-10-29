package ident1246

func MinimumMoves(arr []int) int {
	n:=len(arr)
	dp:= make([][]int,n)
	for i:=range dp{
		dp[i]=make([]int,n)
	}
	for i:=0;i<n;i++{
		dp[i][i]=1
	}
	for i:=0;i<n-1;i++{
		if(arr[i]==arr[i+1]){
			dp[i][i+1]=1
		}else{
			dp[i][i+1]=2
		}
	}
	for l:=2;l<n;l++{
		for i:=0;i<n-l;i++{
			dp[i][i+l]=l+1
			for k:=0;k<l;k++{
				if(dp[i][i + l] > dp[i][i + k] + dp[i + k + 1][i + l]){
					dp[i][i + l] = dp[i][i + k] + dp[i + k + 1][i + l]
				}
			}
			if(arr[i] == arr[i + l] && dp[i][i + l] > dp[i + 1][i + l - 1]){
				dp[i][i + l] = dp[i + 1][i + l - 1]
			}
		}
	}
	
	return dp[0][n-1]
}
