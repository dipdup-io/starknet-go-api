package abi

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func extractTupleTypes(typ string) ([]tupleItem, error) {
	if len(typ) < 2 {
		return nil, errors.Wrap(ErrInvalidTupleType, typ)
	}
	typ = strings.ReplaceAll(typ, " ", "")
	typ = strings.TrimPrefix(typ, "(")
	typ = strings.TrimSuffix(typ, ")")

	items, _ := subTuples(typ)
	return items, nil
}

type tupleItem struct {
	Type

	Childs []tupleItem
}

func newTupleItem(typ, name string, childs ...tupleItem) tupleItem {
	return tupleItem{
		Type: Type{
			Type: typ,
			Name: strings.TrimPrefix(name, ",("),
		},
		Childs: childs,
	}
}

func newTupleItemFromString(str string) tupleItem {
	if strings.Contains(str, ":") && !strings.Contains(str, ",") {
		parts := strings.Split(str, ":")
		return newTupleItem(parts[1], parts[0])
	}

	return newTupleItem(str, "")
}

func subTuples(typ string) ([]tupleItem, int) {
	var (
		lastIndex int
		current   []tupleItem
		items     = make([]tupleItem, 0)
		count     int
	)
	for i := 0; i < len(typ); i++ {
		switch typ[i] {
		case '(':
			tupleItems, parsedCount := subTuples(typ[i+1:])
			if i+parsedCount == len(typ)-1 {
				if len(items) == 0 {
					return tupleItems, i + parsedCount
				}
				return append(items, newTupleItem(typ[i:], tupleName(count), tupleItems...)), i
			}
			i += parsedCount
			current = tupleItems
		case ')':
			item := newTupleItemFromString(typ[lastIndex:i])
			item.Childs = append(item.Childs, current...)
			if item.Name == "" {
				item.Name = tupleName(count)
			}
			items = append(items, item)
			return items, i + 1
		case ',':
			item := newTupleItemFromString(typ[lastIndex:i])
			item.Childs = append(item.Childs, current...)
			if item.Name == "" {
				item.Name = tupleName(count)
			}
			current = current[:0]
			items = append(items, item)
			lastIndex = i + 1
			count++
		}
	}

	item := newTupleItemFromString(typ[lastIndex:])
	item.Childs = append(item.Childs, current...)
	if item.Name == "" {
		item.Name = tupleName(count)
	}
	items = append(items, item)

	return items, 0
}

func tupleName(idx int) string {
	return fmt.Sprintf("tuple_item_%d", idx)
}
