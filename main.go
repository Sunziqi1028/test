/**
 * @Author: Sun
 * @Company: 苏州天和汇工业
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/1/29 14:26
 */

package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("git log")

	outInfo := bytes.Buffer{}

	cmd.Stdout = &outInfo
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(outInfo.String())
}
