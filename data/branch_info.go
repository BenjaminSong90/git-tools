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

// 检查分支，移除不存在的分支
func (branchInfo *BranchInfo) Verify(newBranches *[]Branch) {
	branchInfo.VerifyBranch(newBranches)
	branchInfo.VerifyGroupBranch()
}

// 检查分支，移除不存在的分支
func (branchInfo *BranchInfo) VerifyBranch(newBranches *[]Branch) {
	recordBranchMap := make(map[string]Branch)

	if newBranches == nil {
		branchInfo.Branches = []Branch{}
		return
	}

	for _, b := range branchInfo.Branches {
		recordBranchMap[b.Name] = b
	}

	notDeletedBranches := []Branch{}

	for _, b := range *newBranches {
		if ndb, ok := recordBranchMap[b.Name]; ok {
			notDeletedBranches = append(notDeletedBranches, ndb)
		} else {
			notDeletedBranches = append(notDeletedBranches, b)
		}
	}

	branchInfo.Branches = notDeletedBranches
}

// 根据branchInfo 中的Branches 对 Group中的branch进行检查，移除不存在的branch
func (branchInfo *BranchInfo) VerifyGroupBranch() {
	recordBranchMap := make(map[string]bool)

	for _, b := range branchInfo.Branches {
		recordBranchMap[b.Name] = true
	}

	branchGroupHasOwner := []BranchGroup{}

	// 过滤有owner的group并添加到数组里
	for _, bg := range branchInfo.BranchGroups {
		if _, ok := recordBranchMap[bg.Owner.Name]; ok {
			branchGroupHasOwner = append(branchGroupHasOwner, bg)
		}
	}

	branchInfo.BranchGroups = branchGroupHasOwner

	// 移除分组中已经被移除的branch
	for i := range branchGroupHasOwner {
		bg := &branchInfo.BranchGroups[i]
		bArray := []Branch{}

		for j := range bg.Branches {
			b := &bg.Branches[j]
			if _, ok := recordBranchMap[b.Name]; ok && b.Name != bg.Owner.Name {
				bArray = append(bArray, *b)
			}
		}
		bg.Branches = bArray
	}
}
