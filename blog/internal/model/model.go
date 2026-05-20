package model

import "time"

type BlogPost struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
	Is_Published bool `json:"is_published"`
	Created_At time.Time `json:"created_at"`
}
