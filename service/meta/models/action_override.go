package models

/**
代表在标准或自定义对象上的操作重写。
使用它可以创建、更新、编辑或删除操作覆盖。
只能通过指定的CustomObject来调用ActionOverride
*/

type ActionOverride struct {
	ActionName           string             `json:"actionName"`
	Comment              string             `json:"comment"`
	Content              string             `json:"content"`
	FormFactor           FormFactor         `json:"formFactor"`
	SkipRecordTypeSelect bool               `json:"skipRecordTypeSelect"`
	Type                 ActionOverrideType `json:"type"`
}
