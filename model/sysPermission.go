package model

type SysPermission struct {
	Id  int `json:"id" form:"id"`
	Pid int `json:"parentId"`
	/**
	 * 权限类型[menu/permission]
	 */
	Type string `json:"type"`

	Title string `json:"title"`

	/**
	 * 权限编码[只有type= permission才有  user:view]
	 */
	Percode string `json:"percode"`

	Icon string `json:"icon"`

	Href string `json:"href"`

	Target string `json:"target"`

	Open int `json:"open"`

	Ordernum int `json:"Ordernum"`

	/**
	 * 状态【0不可用1可用】
	 */
	Available int `json:"available"`
}

type TreeNode struct {
	Id       int    `json:"id"`
	Pid      int    `json:"pId"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Href     string `json:"href"`
	Spread   bool   `json:"spread"`
	Title    string `json:"title"`
	CheckArr string `json:"checkArr" gorm:"default:0"`
	Children []TreeNode
}

func NewTreeNode(id int, pid int, spread bool, title string) *TreeNode {
	return &TreeNode{Id: id, Pid: pid, Spread: spread, Title: title}
}
