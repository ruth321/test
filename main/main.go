package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//type group struct {
//	Name     string
//	GroupID  int
//	ParentID int
//}
//
//type groups []group
//
//func sortByName(grs []group, s int, e int) []group {
//	for i := s + 1; i < e; i++ {
//		if grs[i].Name < grs[i-1].Name {
//			gr := grs[i]
//			g := i
//			for g > 0 && gr.Name < grs[g-1].Name {
//				grs[g] = grs[g-1]
//				g--
//			}
//			grs[g] = gr
//		}
//	}
//	return grs
//}
//
//func sortByParentsFirst(grs []group) []group {
//	parentID := 0
//	c := 0
//	for i := 0; i < len(grs); i++ {
//		s := c
//		for g := c; g < len(grs); g++ {
//			if parentID == grs[g].ParentID {
//				grs[c], grs[g] = grs[g], grs[c]
//				c++
//			}
//		}
//		grs = sortByName(grs, s, c)
//		parentID = grs[i].GroupID
//	}
//	return grs
//}
//
//func sortByParentWithChildren(grs []group) []group {
//	var gr group
//	var aGrs []group
//	for k := 0; k < len(grs); k++ {
//		for i := 0; i < len(grs); i++ {
//			for g := 0; g < len(grs); g++ {
//				if grs[i].ParentID == grs[g].GroupID {
//					if contains(aGrs, grs[i]) {
//						continue
//					}
//					gr = grs[i]
//					if i < g {
//						for j := i; j < g; j++ {
//							grs[j] = grs[j+1]
//						}
//						grs[g] = gr
//					} else {
//						for j := i; j > g+1; j-- {
//							grs[j] = grs[j-1]
//						}
//						grs[g+1] = gr
//					}
//					aGrs = append(aGrs, gr)
//					i--
//					break
//				}
//			}
//		}
//	}
//	return grs
//}
//
//func anotherSort(grs []group) []group {
//	var gr group
//	for i := 0; i < len(grs); i++ {
//		for g := 0; g < len(grs); g++ {
//			if grs[i].ParentID == grs[g].GroupID {
//				gr = grs[i]
//				if i < g {
//					for j := i; j < g; j++ {
//						grs[j] = grs[j+1]
//					}
//					grs[g] = gr
//				} else {
//					for j := i; j > g+1; j-- {
//						grs[j] = grs[j-1]
//					}
//					grs[g+1] = gr
//				}
//				i--
//				break
//			}
//		}
//	}
//	return grs
//}
//
//func contains(grs []group, gr group) bool {
//	for i := 0; i < len(grs); i++ {
//		if gr == grs[i] {
//			return true
//		}
//	}
//	return false
//}

func main() {
	//taskGroups := groups{
	//	//	{Name: "д", GroupID: 21, ParentID: 0}, //0
	//	//	{Name: "па", GroupID: 22, ParentID: 21}, //1
	//	//	{Name: "пг", GroupID: 32, ParentID: 21}, //2
	//	//	{Name: "л", GroupID: 17, ParentID: 0}, //3
	//	//}
	//
	//	{Name: "ggg", GroupID: 18, ParentID: 13},
	//	{Name: "hhh", GroupID: 19, ParentID: 18},
	//	{Name: "iii", GroupID: 20, ParentID: 0},
	//	{Name: "ccc", GroupID: 14, ParentID: 12},
	//	{Name: "aaa", GroupID: 12, ParentID: 0},
	//	{Name: "ddd", GroupID: 15, ParentID: 12},
	//	{Name: "bbb", GroupID: 13, ParentID: 0},
	//	{Name: "kkk", GroupID: 22, ParentID: 21},
	//	{Name: "eee", GroupID: 16, ParentID: 14},
	//	{Name: "jjj", GroupID: 21, ParentID: 16},
	//	{Name: "fff", GroupID: 17, ParentID: 15},
	//	{Name: "lll", GroupID: 23, ParentID: 16},
	//}
	//taskGroups = sortByParentWithChildren(taskGroups)
	//for i := 0; i < len(taskGroups); i++ {
	//	fmt.Println(taskGroups[i])
	//}
	r := mux.NewRouter()
	r.HandleFunc("/aaa/{qqq}", aaa)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8070", nil))
}

func aaa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars)
}
