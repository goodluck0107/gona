package utils

import "strings"
/*下划线字符串转驼峰字符串
table_test_only ==> TableTestOnly
*/
func NormalizeString(srcStr string)(string){
	srcStrArr :=strings.Split(srcStr,"_")
	destStr := ""
	for _,str :=range srcStrArr{
		for charIndex,tChar :=range str{
			if charIndex == 0{
				destStr=destStr+strings.ToUpper(string(tChar))
			}else{
				destStr=destStr+string(tChar)
			}
		}
	}
	return destStr
}