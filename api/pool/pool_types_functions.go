package pool


func NewMemberNode(host string, priority int, state string, weight int) MemberNode {
	memberNode := MemberNode{
		Node:     host,
		Priority: priority,
		State:    state,
		Weight:   weight,
	}
	return memberNode

}
