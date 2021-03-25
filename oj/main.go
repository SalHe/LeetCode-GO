package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {

	httpClient := &http.Client{}

	request, _ := http.NewRequest("GET", "https://leetcode-cn.com/api/problems/algorithms/", nil)
	request.Header.Set("Cookies", leetcodeCookies)
	response, err := httpClient.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode != 200 {
		fmt.Println("请求失败！")
		return
	}

	var leetcodeProblems leetcodeProblems
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	if err = json.Unmarshal(bodyBytes, &leetcodeProblems); err != nil {
		fmt.Println(err)
		return
	}

	displayUserInfo(&leetcodeProblems)
	processProblems(&leetcodeProblems)

	generateDirectories(&leetcodeProblems)

}

func generateDirectories(problems *leetcodeProblems) {
	fmt.Println("Generate directories for these questions? (Y\\N)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	c := scanner.Text()
	if c != "Y" && c != "y" {
		return
	}

	fmt.Println("Please input a path to save directories: ")
	scanner.Scan()
	dir := scanner.Text()
	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("Directory not exists, creating...")
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			_ = fmt.Errorf("Mkdir failed!\n")
			return
		} else {
			fmt.Println("Mkdir success.")
		}
	}

	generateProblemDirectories(problems, dir)
	fmt.Println("Over!")
}

func generateProblemDirectories(problems *leetcodeProblems, dir string) {
	for _, problem := range problems.StatStatusPairs {
		title := strings.ReplaceAll(problem.Stat.QuestionTitle, " ", "-")
		subdir := fmt.Sprintf("%04d-%s", problem.Stat.QuestionID, title)
		problemDir := path.Join(dir, subdir)
		err := os.Mkdir(problemDir, os.ModePerm)
		if err != nil {
			_ = fmt.Errorf("Mkdir for problem %4d\"%s\" failed: %s\n", problem.Stat.QuestionID, title, problemDir)
		} else {
			fmt.Printf("Mkdir for problem %4d\"%s\" success: %s\n", problem.Stat.QuestionID, title, problemDir)
		}
	}
}

func displayUserInfo(problems *leetcodeProblems) {
	fmt.Println("----------------------------------------")
	fmt.Printf("Username:\t%s\n", problems.UserName)
	fmt.Printf("Solved/Total:\t%d/%d\n", problems.NumSolved, problems.NumTotal)
	fmt.Println("----------------------------------------")
}

func processProblems(problems *leetcodeProblems) {
	for id := range problems.StatStatusPairs {
		problem := problems.StatStatusPairs[len(problems.StatStatusPairs)-id-1]
		fmt.Printf("Question ID: %d\n", problem.Stat.QuestionID)
		fmt.Printf("Question Title: %s\n", problem.Stat.QuestionTitle)
		fmt.Println("")
	}
}
