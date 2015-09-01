/*
This is a home-work like pracitce for been familiar with golang
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const CHAR_LEN = 3

func process_str(org string) string {
	num_map := map[string]float64{
		"一": 1,
		"二": 2,
		"兩": 2,
		"叄": 3,
		"三": 3,
		"四": 4,
		"五": 5,
		"六": 6,
		"七": 7,
		"八": 8,
		"九": 9,
		"零": 0,
	}
	multi_map := map[string]float64{
		"萬": 10000,
		"千": 1000,
		"百": 100,
		"十": 10,
		"毛": 0.1,
		"分": 0.01,
	}
	multi_step_map := map[string]float64{
		"兆": 10000 * 10000 * 10000,
		"億": 10000 * 10000,
	}

	beg := strings.IndexAny(org, "一二三四五六七八九零")
	var end int

	if beg == -1 {
		return org
	}

	var temp, step_sum, final_sum float64
	for itor := beg; ; itor += CHAR_LEN {
		c := org[itor : itor+CHAR_LEN]
		if multi, ok := multi_map[c]; ok {
			temp *= multi
		} else if multi, ok := multi_step_map[c]; ok {
			step_sum += temp
			temp = 0
			step_sum *= multi
			final_sum += step_sum
			step_sum = 0
		} else if val, ok := num_map[c]; ok {
			step_sum += temp
			temp = val
		} else {
			final_sum += step_sum + temp
			end = itor
			break
		}

	}
	new_str := org[:beg] + fmt.Sprintf("%.4f", final_sum) + org[end:]
	fmt.Println(new_str)

	return org
}

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		print_str := process_str(scanner.Text())
		fmt.Println(print_str)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
