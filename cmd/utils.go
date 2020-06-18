package cmd


//type StringArray []string
//
//
//func NewStringArray(maxSize int) *StringArray {
//	ret := StringArray(make([]string, maxSize))
//	return &ret
//}
//
//
//func (sa *StringArray) Append(input []string) *StringArray {
//	for _, v := range input {
//		*sa = append(*sa, v)
//	}
//	return sa
//}
//
//
//func (sa *StringArray) ToArray() *[]string {
//	ret := []string(*sa)
//	return &ret
//}
//
//
//func (sa *StringArray) SplitFirst(maxSize int) (string, []string) {
//	var retString string
//	var retArray []string
//
//	for range onlyOnce {
//		if len(*sa) == 0 {
//			break
//		}
//
//		retString = (*sa)[0]
//		if len(*sa) <= 1 {
//			break
//		}
//		for index := 1; index < maxSize; index++ {
//			if index < len(*sa) {
//				retArray = append(retArray, (*sa)[index])
//				continue
//			}
//			retArray = append(retArray, "")
//		}
//	}
//
//	return retString, retArray
//}
//
//
//func (sa *StringArray) Divide(indexSplit int, maxSize int) ([]string, []string) {
//	var retString string
//	var retArray []string
//
//	//for index := 0; index < maxSize; index++ {
//	//	if index < indexSplit
//	//}
//
//	return retString, retArray
//}
