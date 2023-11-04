package merkletree

import (
	"crypto/sha256"
	"errors"
	"log"
	"strconv"

	"github.com/cbergoon/merkletree"
)

// TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type TestContent struct {
	x string
}

// CalculateHash hashes the values of a TestContent
func (t TestContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.x)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (t TestContent) Equals(other merkletree.Content) (bool, error) {
	otherTC, ok := other.(TestContent)
	if !ok {
		return false, errors.New("value is not of type TestContent")
	}
	return t.x == otherTC.x, nil

}

func TestMerkleTree() {
	//Build list of Content to build tree
	var list []merkletree.Content
	for i := 1; i <= 50000; i++ {
		content := TestContent{strconv.Itoa(i)}
		list = append(list, content)
	}
	list = append(list, TestContent{x: "Hello"})
	list = append(list, TestContent{x: "Hi"})
	list = append(list, TestContent{x: "Hey"})
	list = append(list, TestContent{x: "Hola"})

	//Create a new Merkle Tree from the list of Content
	t, err := merkletree.NewTree(list)
	if err != nil {
		log.Fatal(err)
	}

	//Get the Merkle Root of the tree
	mr := t.MerkleRoot()
	log.Println(mr)

	//Verify the entire tree (hashes for each node) is valid
	vt, err := t.VerifyTree()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Verify Tree: ", vt)

	//Verify a specific content in in the tree
	testCase := TestContent{
		x: "1",
	}
	vc, err := t.VerifyContent(testCase)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Verify Content: ", vc)

	//String representation
	//log.Println(t)
}

type Tree struct {
	originId string
}
type Article struct {
	Id      string
	Contend string
	Author  string
	Images  []string
	Users   []string
}

func (t *Tree) getProof(interface{}) string {
	return ""
}
func (t *Tree) verifyLeaf(p1, p2 interface{}) bool {
	return false
}

var ResponseArticle []*Article

// 用协程池去跑这些计算
func rangeTree(param interface{}) {
	var forest []*Tree
	forest = GetForest()
	for i, tree := range forest {
		proof := tree.getProof(param)
		isLeaf := tree.verifyLeaf(proof, param)
		if isLeaf {
			articles := GetArticles(tree.originId)
			ResponseArticle = append(ResponseArticle, articles...)
		}
	}
}
func GetArticles(originId string) []*Article {
	return []*Article{
		{
			Id:      "",
			Contend: "",
			Author:  "",
			Images:  nil,
		},
	}
}
func GetForest() []*Tree {
	return []*Tree{
		{},
	}
}
