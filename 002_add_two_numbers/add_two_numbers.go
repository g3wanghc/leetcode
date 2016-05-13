/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) ListNode {
    var l1_node = &l1
    var l2_node = &l2
    
    var l3 ListNode
    var l3_node = &l3
    var l3_empty bool = true
   
    var carry_over bool = false
    
    for *l1_node != nil || *l2_node != nil || carry_over {
        var sum int = 0
        
        if *l1_node != nil{
            sum += (*l1_node).Val
        }

        if *l2_node != nil{
            sum += (*l2_node).Val
        }

        if carry_over {
            sum += 1
            carry_over = false
        }
        
        if sum >= 10 {
            sum = sum % 10
            carry_over = true
        } 

        if *l1_node != nil{
            l1_node = &(*l1_node).Next
        }
        
        if *l2_node != nil{
            l2_node = &(*l2_node).Next
        }
        
        if !l3_empty{
            var tmp ListNode = ListNode {sum, nil}
            l3_node.Next = &tmp
            l3_node = l3_node.Next
        } else {
            l3 = ListNode {sum, nil}
            l3_empty = false
        }
    }
    
    return l3
}