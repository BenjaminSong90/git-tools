package data

type BranchInfo struct {
	Version      int           `json:"version"`
	Branches     []Branch      `json:"branches"`
	BranchGroups []BranchGroup `json:"branch_groups"`
}

type Branch struct {
	Name     string `json:"name"`
	Describe string `json:"describe"`
}

type BranchGroup struct {
	Name     string   `json:"name"`
	Describe string   `json:"describe"`
	Owner    Branch   `json:"owner"`
	Branches []Branch `json:"branches"`
}
