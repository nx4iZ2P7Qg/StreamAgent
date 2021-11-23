package minio

import "testing"

//func TestGetClient(t *testing.T) {
//	getClient()
//}

func TestCreateBucket(t *testing.T) {
	err := CreateBucket("mybucket", "east")
	if err != nil {
		panic(err)
	}

}

func TestUpload(t *testing.T) {
	err := Upload("another", "2021年7月电影推荐", "client.go")
	if err != nil {
		panic(err)
	}
}
