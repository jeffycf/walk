// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/lxn/walk"
)

type TreeView struct {
	AssignTo             **walk.TreeView
	Name                 string
	Enabled              Property
	Visible              Property
	Font                 Font
	ToolTipText          Property
	MinSize              Size
	MaxSize              Size
	StretchFactor        int
	Row                  int
	RowSpan              int
	Column               int
	ColumnSpan           int
	ContextMenuItems     []MenuItem
	OnKeyDown            walk.KeyEventHandler
	OnKeyPress           walk.KeyEventHandler
	OnKeyUp              walk.KeyEventHandler
	OnMouseDown          walk.MouseEventHandler
	OnMouseMove          walk.MouseEventHandler
	OnMouseUp            walk.MouseEventHandler
	OnSizeChanged        walk.EventHandler
	Model                walk.TreeModel
	OnCurrentItemChanged walk.EventHandler
	OnItemCollapsed      walk.TreeItemEventHandler
	OnItemExpanded       walk.TreeItemEventHandler
}

func (tv TreeView) Create(builder *Builder) error {
	w, err := walk.NewTreeView(builder.Parent())
	if err != nil {
		return err
	}

	return builder.InitWidget(tv, w, func() error {
		if err := w.SetModel(tv.Model); err != nil {
			return err
		}

		if tv.OnCurrentItemChanged != nil {
			w.CurrentItemChanged().Attach(tv.OnCurrentItemChanged)
		}

		if tv.OnItemCollapsed != nil {
			w.ItemCollapsed().Attach(tv.OnItemCollapsed)
		}

		if tv.OnItemExpanded != nil {
			w.ItemExpanded().Attach(tv.OnItemExpanded)
		}

		if tv.AssignTo != nil {
			*tv.AssignTo = w
		}

		return nil
	})
}

func (w TreeView) WidgetInfo() (name string, disabled, hidden bool, font *Font, toolTipText string, minSize, maxSize Size, stretchFactor, row, rowSpan, column, columnSpan int, contextMenuItems []MenuItem, OnKeyDown walk.KeyEventHandler, OnKeyPress walk.KeyEventHandler, OnKeyUp walk.KeyEventHandler, OnMouseDown walk.MouseEventHandler, OnMouseMove walk.MouseEventHandler, OnMouseUp walk.MouseEventHandler, OnSizeChanged walk.EventHandler) {
	return w.Name, false, false, &w.Font, "", w.MinSize, w.MaxSize, w.StretchFactor, w.Row, w.RowSpan, w.Column, w.ColumnSpan, w.ContextMenuItems, w.OnKeyDown, w.OnKeyPress, w.OnKeyUp, w.OnMouseDown, w.OnMouseMove, w.OnMouseUp, w.OnSizeChanged
}
