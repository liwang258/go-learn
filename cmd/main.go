package main

import (
	"encoding/json"
	"fmt"
	"go-learn/task3"
	"strconv"
)

func main() {
	uid1, _ := task3.AddUser("张三")
	uid2, _ := task3.AddUser("李四")
	uid3, _ := task3.AddUser("王五")
	user1 := task3.FindUser(uid1)
	postId, _ := task3.AddPost(user1, "user "+strconv.Itoa(user1.Id)+" add Post1")

	user2 := task3.FindUser(uid2)

	user3 := task3.FindUser(uid3)
	postId2, _ := task3.AddPost(user3, "user "+strconv.Itoa(user3.Id)+" add Post1")
	task3.AddComment(user2, postId, "user 2 add commet")
	task3.AddComment(user1, postId2, "user 1 add commet")
	task3.AddComment(user2, postId2, "user 2 add commet")
	hot := task3.FindHotPosts()
	if b, err := json.Marshal(hot); err == nil {
		fmt.Printf("找到最热门的帖子:\n%s\n", string(b))
	} else {
		fmt.Println("未找到有效记录")
	}
	task3.DeleteCommet(user1, postId2)
}
