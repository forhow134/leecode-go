package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//l1 := newList([]int{2, 4, 3})
	//l2 := newList([]int{5, 6, 4})
	//l1 := newList([]int{0})
	//l2 := newList([]int{0})
	l1 := newList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := newList([]int{9, 9, 9, 9})
	addTwoNumbers(l1, l2)
}

func newList(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0], Next: nil}
	cur := head
	for i := 1; i < n; i++ {
		newNode := &ListNode{Val: nums[i], Next: nil}
		cur.Next = newNode
		cur = newNode
	}
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var listNodes []int
	overTen := false
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val
		if overTen {
			sum++
		}
		overTen = false
		if sum >= 10 {
			overTen = true
		}
		listNodes = append(listNodes, sum%10)
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum := l1.Val
		if overTen {
			sum++
		}
		overTen = false
		if sum >= 10 {
			overTen = true
		}
		listNodes = append(listNodes, sum%10)
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val
		if overTen {
			sum++
		}
		overTen = false
		if sum >= 10 {
			overTen = true
		}
		listNodes = append(listNodes, sum%10)
		l2 = l2.Next
	}
	if l1 == nil && l2 == nil && overTen {
		listNodes = append(listNodes, 1)
	}
	//fmt.Println(listNodes)
	return newList(listNodes)
}
