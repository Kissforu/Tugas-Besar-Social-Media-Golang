package main

import (
	"bufio"
	"fmt"
	"os"
)

const NMAX int = 100

type comment struct {
	user    string
	content string
}
type posts struct {
	id              int
	content, userby string
	commentz        arr3
}

type users struct {
	username             string
	password             string
	bio                  string
	followers, following arr4
	blocks, blockedby    arr
	is_you               bool
	pozts                arr2
}
type yourdetails struct {
	you       string
	yourindex int
}
type follow2 struct {
	followid int
	followby string
}
type arr4 = [NMAX]follow2
type allusers = [NMAX]users
type arr = [NMAX]string
type allposts = [NMAX]posts
type arr2 = [NMAX]posts
type arr3 = [NMAX]comment

var userposts allposts

var templist allposts

var bb allusers

// break
func main() {
	menu()
}

func postlist(i *int, logged bool, detailsofu yourdetails, templistisempty *bool) {
	j := *i
	var templist2 allposts
	count := 0
	blockcount := 0
	block := false
	var followed string
	var empty bool = true
	for *i < j+5 {
		if userposts[*i].content != "" {
			block = false

			if logged {
				for g := 0; g < NMAX; g++ {
					if userposts[*i].userby == bb[detailsofu.yourindex].blocks[g] {
						block = true
						fmt.Println("check1")
					}
					empty = false
					for s := 0; s < NMAX; s++ {
						if userposts[*i].userby == bb[s].username {
							for d := 0; d < NMAX; d++ {
								if bb[s].blocks[d] == detailsofu.you {
									block = true
									fmt.Println("check2")

								}
							}
						}
					}
				}
				if block {
					blockcount++
				}
			}

			if blockcount == 5 {
				q := *i
				postlist(&q, true, detailsofu, templistisempty)
				return
			}
			if !block {
				empty = false
				templist2[count] = userposts[*i]
				count++

				// Determine follow status
				followed = "Follow"
				if userposts[*i].userby == detailsofu.you {
					followed = ""
				} else {
					for k := 0; k < NMAX; k++ {
						if bb[detailsofu.yourindex].following[k].followby == userposts[*i].userby {
							followed = "Followed"

						}
					}
				}

				fmt.Printf("%d. User: %s %s\n", count, userposts[*i].userby, followed)
				fmt.Printf("    Post:  %s\n", userposts[*i].content)
				fmt.Println("Comments:")
				for j := 0; j < NMAX; j++ {
					if userposts[*i].commentz[j].content != "" {
						fmt.Printf(" - %s: %s\n", userposts[*i].commentz[j].user, userposts[*i].commentz[j].content)
					}
				}
				fmt.Println()
			}
		}
		*i++

	}
	templist = templist2
	*i = j
	if empty {
		fmt.Println("No More Posts")
		*templistisempty = true
	} else {
		*templistisempty = false
	}
}

func seeposts(logged bool, detailsofu yourdetails) {
	var s int
	i := 1
	alrsaw := false
	var templistisempty bool = true
	for {
		if alrsaw == false {
			postlist(&i, logged, detailsofu, &templistisempty)
		}

		if logged {
			fmt.Println("[◉ °]--------------------")
			fmt.Println("|_1 Comment _____|")
			fmt.Println("|2 Make Post_|")
			fmt.Println("|3 Log Out_|")
			fmt.Println("|4 Search Account|")
			fmt.Println("|5 Next Page_|")
			fmt.Println("|6 Prev Page_|")
			fmt.Println("--------------------------")
			fmt.Scan(&s)

			bufio.NewReader(os.Stdin).ReadString('\n')

			if s == 1 {
				makecomment(detailsofu)
				if alrsaw == true {
					postlist(&i, true, detailsofu, &templistisempty)
				}
			} else if s == 2 {
				makepost(detailsofu)
				if alrsaw == true {
					postlist(&i, true, detailsofu, &templistisempty)
				}
			} else if s == 3 {
				logged = false
				detailsofu.you = ""
				bb[detailsofu.yourindex].is_you = false
				detailsofu.yourindex = 0
				return
			} else if s == 4 {
				search(detailsofu)
				alrsaw = false
			} else if s == 5 {
				if i+5 > NMAX {
					fmt.Println("No More Posts")
					postlist(&i, true, detailsofu, &templistisempty)
					alrsaw = true
				} else {
					if templistisempty {

					} else {
						i += 5
					}
					postlist(&i, true, detailsofu, &templistisempty)
					alrsaw = true
				}

			} else if s == 6 {
				if i-5 < 0 {
					fmt.Println("No More Previous Posts")
					postlist(&i, true, detailsofu, &templistisempty)
					alrsaw = true
				} else {
					i -= 5
					postlist(&i, true, detailsofu, &templistisempty)
					alrsaw = true
				}

			} else {
				fmt.Println("invalid input")
			}
		} else {
			fmt.Println("|")
			fmt.Println("1 Next Page_|")
			fmt.Println("2 Prev Page_|")
			fmt.Println("3 Back|")
			fmt.Println("|")
			fmt.Scan(&s)
			bufio.NewReader(os.Stdin).ReadString('\n')
			if s == 1 {
				if i+5 > NMAX {
					fmt.Println("No More Posts")
					postlist(&i, false, detailsofu, &templistisempty)
					alrsaw = true
				} else {
					if templistisempty {

					} else {
						i += 5
					}
					postlist(&i, false, detailsofu, &templistisempty)
					alrsaw = true
				}
			} else if s == 2 {
				if i-5 < 0 {
					fmt.Println("No More Previous Posts")
					postlist(&i, false, detailsofu, &templistisempty)
					alrsaw = true
				} else {
					i -= 5
					postlist(&i, false, detailsofu, &templistisempty)
					alrsaw = true
				}
			} else {
				menu()
			}
		}
	}

}
func makecomment(detailsofu yourdetails) {
	var n int
	for w := 0; w < NMAX; w++ {
		if userposts[w].content == "" {
			userposts[w].id = 101
		}
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("insert which posts to comment:")
	fmt.Scan(&n)
	if n > 5 || n < 1 {
		fmt.Println("invalid input: No posts in index")
		return
	}
	left, right := 1, NMAX-1
	for left <= right {
		mid := (left + right) / 2
		if templist[n-1].id == userposts[mid].id {
			if userposts[mid].content != "" {
				bufio.NewReader(os.Stdin).ReadString('\n')
				fmt.Println("write comment:")
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading input:", err)
					return
				}

				input = input[:len(input)-1]

				for j := 0; j < NMAX; j++ {
					if userposts[mid].commentz[j].content == "" {
						userposts[mid].commentz[j].content = input
						userposts[mid].commentz[j].user = detailsofu.you
						return
					}
				}
			} else {
				fmt.Println("Invalid Input: No posts in index")
				return
			}

		} else if templist[n-1].id > userposts[mid].id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("asiduh")
	return

}
func makepost(detailsofu yourdetails) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write caption")

	// Read the input until a newline character
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Remove the newline character from the end of the input
	input = input[:len(input)-1]

	// Store the post in the userposts array
	for i := 1; i < NMAX; i++ {
		if userposts[i].content == "" {
			userposts[i].id = i
			userposts[i].content = input
			userposts[i].userby = detailsofu.you
			for j := 0; j < NMAX; j++ {
				if bb[detailsofu.yourindex].pozts[i].content == "" {
					bb[detailsofu.yourindex].pozts[0] = userposts[i]
					return
				}
			}
			return

		}
	}
}

func menu() {
	for true {
		var s int
		var youracc yourdetails
		fmt.Println("[◉ °]_")
		fmt.Println("|      1 See posts      |")
		fmt.Println("|      2 Register       |")
		fmt.Println("|      3 Log In         |")
		fmt.Println("-------------------------")
		fmt.Scan(&s)

		// Clear the input buffer after reading integer input
		bufio.NewReader(os.Stdin).ReadString('\n')

		if s == 1 {
			seeposts(false, youracc)
		} else if s == 2 {
			register()
		} else if s == 3 {
			youracc = login()
			seeposts(true, youracc)
		} else {
			fmt.Println("invalid input")
		}
	}
}

func login() yourdetails {
	success := false
	var usern string
	var pass string
	var detailsofu yourdetails
	for success == false {
		fmt.Println("    Insert username")
		fmt.Scan(&usern)
		fmt.Println("")
		fmt.Println("    Insert Password")
		fmt.Scan(&pass)
		fmt.Println("")
		i := 0
		for i < NMAX {
			if bb[i].username == usern {
				if bb[i].password == pass {
					fmt.Println("Logged In to", usern)
					fmt.Println("")
					bb[i].is_you = true
					detailsofu.you = bb[i].username
					detailsofu.yourindex = i
					success = true
					return detailsofu
				} else {

				}
			}
			i++
			if i == NMAX {
				fmt.Println("Username or Password does not match")
			}

		}

	}
	return detailsofu

}
func register() {
	success := false
	reader := bufio.NewReader(os.Stdin)
	for success == false {
		var usern, userp string
		fmt.Println("    Insert username (spacebars are prohibited)")
		fmt.Scan(&usern)
		fmt.Print("")
		fmt.Println("    Insert Password (spacebars are prohibited)")
		fmt.Scan(&userp)
		fmt.Println("    Insert Bio:")
		bufio.NewReader(os.Stdin).ReadString('\n')
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		input = input[:len(input)-1]
		fmt.Println("")
		for i := 0; i < NMAX; i++ {
			if bb[i].username == usern {
				fmt.Println("Username taken")
				return
			} else if bb[i].username == "" {
				bb[i].username = usern
				bb[i].password = userp
				bb[i].bio = input
				fmt.Println("Your account has been registered!")
				fmt.Println("")
				success = true
				return
			}
		}
	}

}

func search(detailsofu yourdetails) {
	var target string
	var targetidx int
	var s int
	success := false
	fmt.Println("Insert username:")
	fmt.Scan(&target)
	if target == detailsofu.you {
		targetidx = detailsofu.yourindex
		success = true
	} else {
		for i := 0; i < NMAX; i++ {
			if bb[i].username == target {
				for k := 0; k < NMAX; k++ {
					if bb[i].blocks[k] == detailsofu.you {
						success = false
						fmt.Println("target not found")
						return
					}
					targetidx = i
					success = true
				}

			}
		}
	}
	if success {
		fmt.Println("----------------------------")
		fmt.Println("User:")
		fmt.Printf("%v \n", bb[targetidx].username)
		fmt.Println("Bio:")
		fmt.Printf("%s\n", bb[targetidx].bio)
		fmt.Println("----------------------------")
		fmt.Println("----------------------------")
		if target != detailsofu.you {
			fmt.Println("1. Follow")
			fmt.Println("2. Unfollow")
			fmt.Println("3. Block")
			fmt.Println("4. Unblock")
			fmt.Println("5. See Followers")
			fmt.Println("6. See Following")
			fmt.Println("7. Back")
			fmt.Println("----------------------------")
			fmt.Scan(&s)
			if s == 3 {
				blockuser(targetidx, detailsofu)
			} else if s == 1 {
				follow(targetidx, detailsofu)
			} else if s == 2 {
				unfollow(targetidx, detailsofu)
			} else if s == 4 {
				unblock(bb[targetidx], detailsofu)
			} else if s == 6 {
				seefollowing(true, targetidx, true)
			} else if s == 5 {
				seefollowers(true, targetidx, true)
			} else {

			}
		} else {
			fmt.Println("1. Change Bio")
			fmt.Println("2. See Following") //sort
			fmt.Println("3. See Follower")
			fmt.Println("4. Back")
			fmt.Println("----------------------------")
			fmt.Scan(&s)
			if s == 1 {
				changebio(detailsofu)
			} else if s == 2 {
				seefollowing(true, detailsofu.yourindex, true)
			} else if s == 3 {
				seefollowers(true, detailsofu.yourindex, true)
			}
		}
		fmt.Println("----------------------------")
	} else {
		fmt.Println("user not found")
	}

}
func blockuser(target int, detailsofu yourdetails) {
	for i := 0; i < NMAX; i++ {
		if bb[detailsofu.yourindex].blocks[i] == "" {
			bb[detailsofu.yourindex].blocks[i] = bb[target].username
			for j := 0; j < NMAX; j++ {
				if bb[target].blockedby[j] == "" {
					bb[target].blockedby[j] = detailsofu.you
					unfollow(target, detailsofu)
					return
				}
			}
		} else if bb[detailsofu.yourindex].blocks[i] == bb[target].username {
			fmt.Println("you have already blocked this user")
			return
		}
	}

}

func follow(target int, detailsofu yourdetails) {
	for i := 0; i < NMAX; i++ {
		if bb[detailsofu.yourindex].blocks[i] == bb[target].username {
			fmt.Println("you are blocking this user")
			return
		}
	}
	for j := 1; j < NMAX; j++ {
		if bb[target].followers[j].followby == "" {
			bb[target].followers[j].followby = detailsofu.you
			bb[target].followers[j].followid = j
			for k := 1; k < NMAX; k++ {
				if bb[detailsofu.yourindex].following[k].followby == "" {
					bb[detailsofu.yourindex].following[k].followby = bb[target].username
					bb[detailsofu.yourindex].following[k].followid = k
					return
				}
			}
		} else if bb[target].followers[j].followby == detailsofu.you {
			fmt.Println("you have already followed this user")
			return
		}
	}

}

func unfollow(target int, detailsofu yourdetails) {
	for j := 1; j < NMAX; j++ {
		if bb[target].followers[j].followby == detailsofu.you {
			bb[target].followers[j].followby = ""
			bb[target].followers[j].followid = 0
		}
	}
	for k := 1; k < NMAX; k++ {
		if bb[detailsofu.yourindex].following[k].followby == bb[target].username {
			bb[detailsofu.yourindex].following[k].followby = ""
			bb[detailsofu.yourindex].following[k].followid = 0
		}
	}
}

func unblock(target users, detailsofu yourdetails) {
	for i := 0; i < NMAX; i++ {
		if bb[detailsofu.yourindex].blocks[i] == target.username {
			bb[detailsofu.yourindex].blocks[i] = ""
			for j := 0; j < NMAX; j++ {
				if target.blockedby[j] == detailsofu.you {
					target.blockedby[j] = ""
					return
				}
			}
		}
	}
}

func changebio(detailsofu yourdetails) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("    Insert Bio:")
	bufio.NewReader(os.Stdin).ReadString('\n')
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = input[:len(input)-1]
	bb[detailsofu.yourindex].bio = input

}

func seefollowing(oldest bool, youridx int, seeresults bool) {
	fmt.Println("---------------")
	var sorted users
	var s int
	var count int = 0
	for i := 1; i < NMAX; i++ {
		if bb[youridx].following[i].followid == 0 {

		} else {
			sorted.following[count] = bb[youridx].following[i]
			count++
		}
	}
	if oldest {
		sorted = sortselect(sorted, true)
	} else {
		sorted = sortinsert(sorted, true)
	}
	if seeresults {
		for i := 0; i < NMAX; i++ {
			if sorted.following[i].followid != 0 {
				fmt.Printf("%d. %s  \n", sorted.following[i].followid, sorted.following[i].followby)
			}
		}
		fmt.Println("---------------")
		fmt.Println("-sort by:")
		fmt.Println("1.Oldest")
		fmt.Println("2.Most Recent")
		fmt.Println("4.Back")
		fmt.Println("---------------")
		fmt.Scan(&s) //ask to wether you want to sort it asc or desc

		if s == 1 {
			seefollowing(true, youridx, true)
		} else if s == 2 {
			seefollowing(false, youridx, true)
		} else {
			seefollowing(true, youridx, false)
			return
		}
	}
}

func seefollowers(oldest bool, youridx int, seeresults bool) {
	fmt.Println("---------------")
	var sorted users
	var s int
	var count int = 0
	for i := 1; i < NMAX; i++ {
		if bb[youridx].followers[i].followid == 0 {

		} else {
			sorted.followers[count] = bb[youridx].followers[i]
			count++
		}
	}
	if oldest {
		sorted = sortselect(sorted, false)
	} else {
		sorted = sortinsert(sorted, false)
	}
	if seeresults {
		for i := 0; i < NMAX; i++ {
			if sorted.followers[i].followid != 0 {
				fmt.Printf("%d. %s  \n", sorted.followers[i].followid, sorted.followers[i].followby)
			}
		}
		fmt.Println("---------------")
		fmt.Println("-sort by:")
		fmt.Println("1.Oldest")
		fmt.Println("2.Most Recent")
		fmt.Println("4.Back")
		fmt.Println("---------------")
		fmt.Scan(&s) //ask to wether you want to sort it asc or desc

		if s == 1 {
			seefollowers(true, youridx, true)
		} else if s == 2 {
			seefollowers(false, youridx, true)
		} else {
			seefollowers(true, youridx, false)
			return
		}
	}
}

func sortselect(array users, following bool) users {
	if following {
		for i := 0; i < NMAX; i++ {
			minidx := i
			for j := i + 1; j < NMAX; j++ {
				if array.following[j].followid < array.following[minidx].followid {
					minidx = j
				}
			}
			array.following[i], array.following[minidx] = array.following[minidx], array.following[i]
		}
	} else {
		for i := 0; i < NMAX; i++ {
			minidx := i
			for j := i + 1; j < NMAX; j++ {
				if array.followers[j].followid < array.followers[minidx].followid {
					minidx = j
				}
			}
			array.followers[i], array.followers[minidx] = array.followers[minidx], array.followers[i]
		}
	}
	return array
}
func sortinsert(array users, following bool) users {
	if following {
		for i := 1; i < NMAX; i++ {
			key := array.following[i]
			j := i - 1

			for j >= 0 && array.following[j].followid < key.followid {
				array.following[j+1] = array.following[j]
				j = j - 1
			}
			array.following[j+1] = key
		}

	} else {
		for i := 1; i < NMAX; i++ {
			key := array.followers[i]
			j := i - 1

			for j >= 0 && array.followers[j].followid < key.followid {
				array.followers[j+1] = array.followers[j]
				j = j - 1
			}
			array.followers[j+1] = key
		}
	}
	return array
}