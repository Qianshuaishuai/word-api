package models

const (
	TYPE_RESOURCE_FRAME_BOOK     = 1
	TYPE_RESOURCE_FRAME_QUESTION = 2
)

//删除截图框
func DeleteFrame(frameId int) (err error) {
	// err = GetDb().Table("t_frames").Where("id = ?", frameId).Delete(Frame{}).Error

	// if err != nil {
	// 	return errors.New("删除失败")
	// }

	return nil
}
