package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
	IDs []int
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// sort by ID
	sort.Slice(records, func(a, b int) bool {
		return records[a].ID < records[b].ID
	})

	if records[0].ID != 0 {
		return nil, errors.New("no root node")
	}

	if records[0].Parent > 0 {
		return nil, errors.New("root cannot have parent")
	}

	var n = &Node{ID: 0}
	for _, r := range records {
		if err := n.addRecord(r); err != nil {
			return nil, err
		}
	}

	return n, nil
}

func (n *Node) addRecord(r Record) error {
	// check if ID has already been added
	if err := n.hasId(r.ID); err != nil {
		return err
	}

	// root element, only append ID
	if r.ID == 0 {
		return n.addId(r.ID)
	}

	// target is root, append to current node, and return
	if r.ID != 0 && r.Parent == 0 {
		n.Children = append(n.Children, &Node{ID: r.ID})
		return n.addId(r.ID)
	}

	if r.ID == r.Parent {
		return errors.New("a node cannot be a parent to itself")
	}

	// find appropriate child node to append to
	for _, nc := range n.Children {
		if nc.ID == r.Parent {
			nc.Children = append(nc.Children, &Node{ID: r.ID})
			if err := n.addId(r.ID); err != nil {
				return err
			}
			break
		}
		if len(nc.Children) > 0 {
			if err := nc.addRecord(r); err != nil {
				return err
			}
		}
	}

	return nil
}

func (n *Node) hasId(id int) error {
	for _, i := range n.IDs {
		if i == id {
			return errors.New("duplicate record")
		}
	}
	return nil
}

func (n *Node) addId(id int) error {
	if len(n.IDs) > 0 {
		lastId := n.IDs[len(n.IDs)-1]
		if lastId != (id - 1) {
			return errors.New("ids must be sequential")
		}
	}

	n.IDs = append(n.IDs, id)
	return nil
}
