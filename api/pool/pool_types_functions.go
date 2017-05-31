package pool

<<<<<<< HEAD

=======
>>>>>>> 8b461f7e26a3bc1f8cd0128d196fa7e2f0107ef6
// NewMemberNode - Returns the Membernodes object
func NewMemberNode(host string, priority int, state string, weight int) MemberNode {
	memberNode := MemberNode{
		Node:     host,
		Priority: priority,
		State:    state,
		Weight:   weight,
	}
	return memberNode
<<<<<<< HEAD

=======
>>>>>>> 8b461f7e26a3bc1f8cd0128d196fa7e2f0107ef6
}
