package db

import (
	"log"
	"os"
	"path"
	"yagc/models"
	"yagc/util"
)

func WriteIndex(content []byte) {
	root, ok := util.GetRepoRoot()
	if !ok {
		log.Fatalln("Failed to get repo root")
	}

	indexPath := path.Join(root, ".yagc", "index")
	file, err := os.Create(indexPath)

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatalf("Failed to close file: %s", err)
		}
	}(file)

	if err != nil {
		log.Fatalf("Failed to create file %s: %s\n", indexPath, err)
	}

	if _, err := file.Write(util.Compress(content)); err != nil {
		log.Fatalf("Failed to write to file %s: %s\n", indexPath, err)
	}
}

func ReadIndex() []byte {
	root, ok := util.GetRepoRoot()
	if !ok {
		log.Fatalln("Failed to get repo root")
	}

	indexPath := path.Join(root, ".yagc", "index")
	content, err := os.ReadFile(indexPath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %s\n", indexPath, err)
	}

	return util.Decompress(content)
}

func CreateIndex(files []string) {
	root, ok := util.GetRepoRoot()
	if !ok {
		log.Fatalln("Failed to get repo root")
	}

	indexPath := path.Join(root, ".yagc", "index")
	tree := models.TreeObject{}

	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		if _, err := os.Create(indexPath); err != nil {
			log.Fatalf("Failed to create index file %s: %s\n", indexPath, err)
		}
	} else {
		index := ReadIndex()
		tree.Parse(index)
	}

	for _, file := range files {
		addTree(&tree, file)
	}

	WriteIndex(tree.GetContent())
}

func addTree(currTree *models.TreeObject, name string) {
	info, err := os.Stat(name)
	if err != nil {
		log.Printf("Failed to stat file %s: %s\n", name, err)
		return
	}

	if info.IsDir() {
		entries, err := os.ReadDir(name)
		if err != nil || len(entries) == 0 {
			log.Printf("Skipping empty directory %s", name)
			return
		}

		newTree := models.TreeObject{
			Entries: make([]models.Entry, 0),
		}
		for _, entry := range entries {
			newPath := path.Join(name, entry.Name())
			addTree(&newTree, newPath)
		}
		sha1, content := newTree.GetSha1()
		WriteObject(sha1, content)
		currTree.Entries = append(currTree.Entries, models.Entry{
			File: name,
			Mode: "000644",
			Id:   sha1,
			Type: "tree",
		})
	} else {
		blob := &models.BlobObject{
			File: name,
		}

		sha1, content := blob.GetSha1()
		if currTree.Exists(sha1) {
			log.Printf("File %s already exists in index\n", name)
			return
		}

		WriteObject(sha1, content)
		currTree.AddEntry(models.Entry{
			Id:   sha1,
			Type: "blob",
			Mode: "100644",
			File: name,
		})
	}
}
