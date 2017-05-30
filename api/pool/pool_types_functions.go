package pool


func NewMemberNodes(host string,priority int, state string, weight int ) MemberNodes{
     memberNodes := MemberNodes{
        Node: host,
	Priority: priority,
	State: state,
	Weight: weight,
     }
	return memberNodes
}

