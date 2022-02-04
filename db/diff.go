package db

import (
	"yagc/models"
)

func DiffIndex() []models.Change {
	var ret []models.Change
	index := ReadIndex()
	commit := LastCommitTree()

	diffTree(&ret, index, commit)

	return ret
}

func diffAll(
	ret *[]models.Change,
	tree *models.TreeObject,
	change models.FileChange,
) {
	for _, entry := range tree.Entries {
		switch entry.Type {
		case models.Blob:
			*ret = append(*ret, models.Change{
				Path:         entry.File,
				Kind:         change,
				OriginalSha1: entry.Id,
			})
		case models.Tree:
			var newTree models.TreeObject
			newTree.Parse(FindObject(entry.Id))
			diffAll(ret, &newTree, change)
		}
	}
}

func diffTree(
	ret *[]models.Change,
	indexTree, commitTree *models.TreeObject,
) {
	entryByName := map[string]models.Entry{}
	for _, entry := range commitTree.Entries {
		entryByName[entry.File+entry.Type] = entry
	}

	for _, entry := range indexTree.Entries {
		if otherEntry, ok := entryByName[entry.File+entry.Type]; ok {
			if otherEntry.Id != entry.Id {
				switch entry.Type {
				case models.Blob:
					*ret = append(*ret, models.Change{
						Path:         entry.File,
						Kind:         models.FileModified,
						OriginalSha1: entry.Id,
						Sha1:         otherEntry.Id,
					})
				case models.Tree:
					var newTree1, newTree2 models.TreeObject
					newTree1.Parse(FindObject(entry.Id))
					newTree2.Parse(FindObject(otherEntry.Id))
					diffTree(ret, &newTree1, &newTree2)
				}
			}
			delete(entryByName, entry.File+entry.Type)
		} else {
			switch entry.Type {
			case models.Blob:
				*ret = append(*ret, models.Change{
					Path:         entry.File,
					Kind:         models.FileAdded,
					OriginalSha1: entry.Id,
				})
			case models.Tree:
				var newTree models.TreeObject
				newTree.Parse(FindObject(entry.Id))
				diffAll(ret, &newTree, models.FileAdded)
			}
		}
	}

	for _, entry := range entryByName {
		switch entry.Type {
		case models.Blob:
			*ret = append(*ret, models.Change{
				Path:         entry.File,
				Kind:         models.FileDeleted,
				OriginalSha1: entry.Id,
			})
		case models.Tree:
			var newTree models.TreeObject
			newTree.Parse(FindObject(entry.Id))
			diffAll(ret, &newTree, models.FileDeleted)
		}
	}
}
