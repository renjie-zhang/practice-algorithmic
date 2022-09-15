/*
The new "Avengers" movie has just been released! There are a lot of people at the cinema box office standing in a huge line.
Each of them has a single 100, 50 or 25 dollar bill. An "Avengers" ticket costs 25 dollars.

Vasya is currently working as a clerk. He wants to sell a ticket to every single person in this line.

Can Vasya sell a ticket to every person and give change if he initially has no money and sells the tickets strictly in the order people queue?

Return YES, if Vasya can sell a ticket to every person and give change with the bills he has at hand at that moment. Otherwise return NO.

Examples:
Tickets([]int{25, 25, 50}) // => YES
Tickets([]int{25, 100}) // => NO. Vasya will not have enough money to give change to 100 dollars
Tickets([]int{25, 25, 50, 50, 100}) // => NO. Vasya will not have the right bills to give 75 dollars of change (you can't make two bills of 25 from one of 50)

链接：https://www.codewars.com/kata/vasya-clerk/train/go
/我的第一种做法是使用一个切片维护收到的钱，虽然很好解决在找零钱补50，但是到补零钱100就犯难了。
在看了别人的题解之后，尝试使用map去维护收到的零钱。
*/
package vasyaclerk

//Tickets ...
func Tickets(peopleInLine []int) string {
	var money = map[int]int{
		25:  0,
		50:  0,
		100: 0,
	}
	for _, v := range peopleInLine {
		switch true {
		case v == 25:
			money[25]++
		case v == 50:
			if money[25] == 0 {
				return "NO"
			} else {
				money[25]--
				money[50]++
			}
		case v == 100:
			if money[50] >= 1 && money[25] >= 1 {
				money[100]++
				money[25]--
				money[50]--
				continue
			} else if money[25] >= 3 {
				money[100]++
				money[25] -= 3
				continue
			} else {
				return "NO"
			}
		default:
			return "NO"
		}
	}
	return "YES"
}
