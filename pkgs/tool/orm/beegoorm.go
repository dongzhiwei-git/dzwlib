package beegoorm

import (
	"fmt"
	"os"
	"strings"
)

// Format 将BeegoORM格式化成Mysql能直接执行的语句
func Format(str string) (string, error) {
	// 读取文件
	//content, err := ioutil.ReadFile("./filename.txt")
	//if err != nil {
	//	return "", err
	//}

	// 将字符串分割成字符串数组
	strArr := strings.Split(str, " - ")
	par := strings.Replace(strArr[1], "`", "'", -1)

	newStrArr := strings.Split(par, ", ")

	count := 0
	retStr := ""
	for i, _ := range strArr[0] {

		if strArr[0][i] == '?' {
			retStr += newStrArr[count]
			count++
		} else {
			retStr += string(strArr[0][i])
		}
	}
	retStr = strings.TrimPrefix(retStr, "[")
	retStr = strings.TrimSuffix(retStr, "]")
	fmt.Println(retStr)
	f, err := os.Create("./result.txt")
	defer f.Close()
	if err != nil {
		return "", err
	}
	_, err = f.WriteString(retStr)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return "", err
	}

	return retStr, nil
}
