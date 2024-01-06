package data

import "errors"

type BranchInfo struct {
	Version      int           `json:"version"`
	Branches     []Branch      `json:"branches"`
	BranchGroups []BranchGroup `json:"branch_groups"`
}

type Branch struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BranchGroup struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Owner       string   `json:"owner"`
	Branches    []string `json:"branches"`
}

// 检查分支，移除不存在的分支
func (branchInfo *BranchInfo) Verify(newBranches *[]string) {
	branchInfo.VerifyBranch(newBranches)
	branchInfo.VerifyGroupBranch()
}

// 检查分支，移除不存在的分支
func (branchInfo *BranchInfo) VerifyBranch(newBranches *[]string) {
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
		if ndb, ok := recordBranchMap[b]; ok {
			notDeletedBranches = append(notDeletedBranches, ndb)
		} else {
			notDeletedBranches = append(notDeletedBranches, Branch{
				Name:        b,
				Description: "",
			})
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
		if _, ok := recordBranchMap[bg.Owner]; ok {
			branchGroupHasOwner = append(branchGroupHasOwner, bg)
		}
	}

	branchInfo.BranchGroups = branchGroupHasOwner

	// 移除分组中已经被移除的branch
	for i := range branchGroupHasOwner {
		bg := &branchInfo.BranchGroups[i]
		bArray := []string{}

		for j := range bg.Branches {
			b := bg.Branches[j]
			if _, ok := recordBranchMap[b]; ok && b != bg.Owner {
				bArray = append(bArray, b)
			}
		}
		bg.Branches = bArray
	}
}

// 设置分支的描述
func (branchInfo *BranchInfo) SetBranchDesc(name, desc string) {
	for i := range branchInfo.Branches {
		b := &branchInfo.Branches[i]
		if b.Name == name {
			b.Description = desc
		}
	}
}

// 向group添加 branch, 会掉重复的
func (group *BranchGroup) addBranch(branchName ...string) {
	group.Branches = append(group.Branches, branchName...)
	group.Branches = RemoveDuplicates(group.Branches)
}

// 设置group 描述
func (branchInfo *BranchInfo) SetGroupDesc(name, desc string) {
	groupMap := branchInfo.getGroupMap()
	if groupP, ok := groupMap[name]; ok {
		groupP.Description = desc
	}
}

// 创建group
func (branchInfo *BranchInfo) CreateGroup(name, owner string, branches []string) error {
	branches = RemoveDuplicates(branches)
	branchNameMap := *branchInfo.getBranchNameMap()

	if _, ok := branchNameMap[owner]; !ok {
		return errors.New("owner branch '" + owner + "' cannot find")
	}

	groupMap := branchInfo.getGroupMap()

	validBranch := []string{}

	for _, b := range branches {
		if _, ok := branchNameMap[b]; ok && b != owner {
			validBranch = append(validBranch, b)
		}
	}

	if groupP, ok := groupMap[name]; ok {
		groupP.addBranch(validBranch...)
	} else {
		branchInfo.BranchGroups = append(branchInfo.BranchGroups, BranchGroup{
			Name:        name,
			Owner:       owner,
			Branches:    validBranch,
			Description: "",
		})
	}

	return nil

}

func (branchInfo *BranchInfo) RemoveBranchFromGroup(groupName string, branches []string) {
	branches = RemoveDuplicates(branches)

	willDeleteBranchMap := make(map[string]bool)
	for _, b := range branches {
		willDeleteBranchMap[b] = true
	}

	groupMap := branchInfo.getGroupMap()

	groupP, ok := groupMap[groupName]

	if !ok {
		return
	}

	resultBranch := []string{}

	for _, b := range groupP.Branches {
		if _, ok := willDeleteBranchMap[b]; !ok {
			resultBranch = append(resultBranch, b)
		}
	}

	groupP.Branches = resultBranch

}

func (branchInfo *BranchInfo) AddBranchToGroup(groupName string, branches []string) {
	branches = RemoveDuplicates(branches)

	groupMap := branchInfo.getGroupMap()
	groupP, ok := groupMap[groupName]
	if !ok {
		return
	}

	branchNameMap := *branchInfo.getBranchNameMap()

	validBranch := []string{}
	for _, b := range branches {
		if _, ok := branchNameMap[b]; ok && groupP.Name != b {
			validBranch = append(validBranch, b)
		}
	}

	groupP.addBranch(validBranch...)

}

func (branchInfo *BranchInfo) getBranchNameMap() *map[string]bool {
	branchNameMap := make(map[string]bool)
	for _, branch := range branchInfo.Branches {
		branchNameMap[branch.Name] = true
	}

	return &branchNameMap
}

func (branchInfo *BranchInfo) getGroupMap() map[string]*BranchGroup {
	branchNameMap := make(map[string]*BranchGroup)
	for i := range branchInfo.BranchGroups {
		bg := &branchInfo.BranchGroups[i]
		branchNameMap[bg.Name] = bg
	}

	return branchNameMap
}

func RemoveDuplicates(strings []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, v := range strings {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}
