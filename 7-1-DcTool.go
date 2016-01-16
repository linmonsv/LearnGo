// Copyright (C) 2013 Andras Belicza. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// A Gowut "Showcase of Features" application.

package main

import (
	"fmt"
	"github.com/icza/gowut/gwu"
	"os"
	"strconv"
	"time"
)

// plural returns an empty string if i is equal to 1,
// "s" otherwise.
func plural(i int) string {
	if i == 1 {
		return ""
	}
	return "s"
}

func buildHomeDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	//p.Add(gwu.NewLabel("This app is written in and showcases Gowut version " + gwu.GOWUT_VERSION + "."))
	p.Add(gwu.NewLabel("The old version of DC_PBOC_CLIENT.exe is written by VC6/MFC."))
	p.AddVSpace(20)
	p.Add(gwu.NewLabel("It's usually crashed in win8.1/10 64bit, so it's time to rewrite now!"))
	p.AddVSpace(20)
	p.Add(gwu.NewLabel("I think Web-Ui is a good choice, and do it in my free time!"))
	p.AddVSpace(20)
	p.Add(gwu.NewLabel("If you have any comments, just email qin2@qq.com or QQ 57235742."))

	return p
}

func buildExpanderDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	l := gwu.NewLabel("Click on the Expander's header.")
	l.Style().SetColor(gwu.CLR_GREEN)
	p.Add(l)
	p.AddVSpace(5)
	e := gwu.NewExpander()
	e.SetHeader(gwu.NewLabel("I'm an Expander."))
	e.SetContent(gwu.NewLabel("I'm the content of the Expander."))
	p.Add(e)
	e.AddEHandlerFunc(func(ev gwu.Event) {
		if e.Expanded() {
			l.SetText("You expanded it.")
		} else {
			l.SetText("You collapsed it.")
		}
		ev.MarkDirty(l)
	}, gwu.ETYPE_STATE_CHANGE)

	p.AddVSpace(20)
	var ee gwu.Expander
	for i := 4; i >= 0; i-- {
		e2 := gwu.NewExpander()
		e2.SetHeader(gwu.NewLabel("I hide embedded expanders. #" + strconv.Itoa(i)))
		if i == 4 {
			e2.SetContent(gwu.NewLabel("No more."))
		} else {
			e2.SetContent(ee)
		}
		ee = e2
	}
	p.Add(ee)

	return p
}

func buildLinkContainerDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	link := gwu.NewLink("An obvious link, to Google Home", "https://google.com/")
	inside := gwu.NewPanel()
	inside.Style().SetBorder2(1, gwu.BRD_STYLE_SOLID, gwu.CLR_GRAY)
	inside.Add(gwu.NewLabel("Everything inside this box also links to Google!"))
	inside.Add(gwu.NewButton("Me too!"))
	link.SetComp(inside)
	p.Add(link)

	return p
}

func buildPanelDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Panel with horizontal layout:"))
	h := gwu.NewHorizontalPanel()
	for i := 1; i <= 5; i++ {
		h.Add(gwu.NewButton("Button " + strconv.Itoa(i)))
	}
	p.Add(h)

	p.AddVSpace(20)
	p.Add(gwu.NewLabel("Panel with vertical layout:"))
	v := gwu.NewVerticalPanel()
	for i := 1; i <= 5; i++ {
		v.Add(gwu.NewButton("Button " + strconv.Itoa(i)))
	}
	p.Add(v)

	p.AddVSpace(20)
	p.Add(gwu.NewLabel("Panel with natural layout:"))
	n := gwu.NewNaturalPanel()
	for i := 1; i <= 20; i++ {
		n.Add(gwu.NewButton("LONG BUTTON " + strconv.Itoa(i)))
	}
	p.Add(n)

	return p
}

func buildTableDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	l := gwu.NewLabel("Tip: Switch to the 'debug' theme (top right) to see cell borders.")
	l.Style().SetColor(gwu.CLR_RED).SetFontStyle(gwu.FONT_STYLE_ITALIC)
	p.Add(l)

	p.AddVSpace(20)
	p.Add(gwu.NewLabel("A simple form aligned with a table:"))
	p.AddVSpace(10)
	t := gwu.NewTable()
	t.SetCellPadding(2)
	t.EnsureSize(2, 2)
	var c gwu.Comp
	t.Add(gwu.NewLabel("User name:"), 0, 0)
	c = gwu.NewTextBox("")
	c.Style().SetWidthPx(160)
	t.Add(c, 0, 1)
	t.Add(gwu.NewLabel("Password:"), 1, 0)
	c = gwu.NewPasswBox("")
	c.Style().SetWidthPx(160)
	t.Add(c, 1, 1)
	t.Add(gwu.NewLabel("Go to:"), 2, 0)
	c = gwu.NewListBox([]string{"Inbox", "User preferences", "Last visited page"})
	c.Style().SetWidthPx(160)
	t.Add(c, 2, 1)
	p.Add(t)

	p.AddVSpace(30)
	p.Add(gwu.NewLabel("Advanced table structure with modified alignment, row and col spans:"))
	p.AddVSpace(10)
	t = gwu.NewTable()
	t.Style().SetBorder2(1, gwu.BRD_STYLE_SOLID, gwu.CLR_GREY)
	t.SetAlign(gwu.HA_RIGHT, gwu.VA_TOP)
	t.EnsureSize(5, 5)
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			t.Add(gwu.NewButton("Button "+strconv.Itoa(row)+strconv.Itoa(col)), row, col)
		}
	}
	t.SetColSpan(2, 1, 2)
	t.SetRowSpan(3, 1, 2)
	t.CellFmt(2, 2).Style().SetSizePx(150, 80)
	t.CellFmt(2, 2).SetAlign(gwu.HA_RIGHT, gwu.VA_BOTTOM)
	t.RowFmt(2).SetAlign(gwu.HA_DEFAULT, gwu.VA_MIDDLE)
	t.CompAt(2, 1).Style().SetFullSize()
	t.CompAt(4, 2).Style().SetFullWidth()
	t.RowFmt(0).Style().SetBackground(gwu.CLR_RED)
	t.RowFmt(1).Style().SetBackground(gwu.CLR_GREEN)
	t.RowFmt(2).Style().SetBackground(gwu.CLR_BLUE)
	t.RowFmt(3).Style().SetBackground(gwu.CLR_GREY)
	t.RowFmt(4).Style().SetBackground(gwu.CLR_TEAL)
	p.Add(t)

	return p
}

func buildTabPanelDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	t := gwu.NewTabPanel()
	t.Style().SetSizePx(500, 300)

	table := gwu.NewTable()
	table.SetCellPadding(2)
	table.EnsureSize(3, 2)
	table.Add(gwu.NewLabel("Change tab bar placement:"), 0, 0)
	table.Add(gwu.NewLabel("Tab bar horizontal align:"), 1, 0)
	table.Add(gwu.NewLabel("Tab bar vertical align:"), 2, 0)

	placemslb := gwu.NewListBox([]string{"Top", "Right", "Bottom", "Left"})
	placems := []gwu.TabBarPlacement{gwu.TB_PLACEMENT_TOP, gwu.TB_PLACEMENT_RIGHT, gwu.TB_PLACEMENT_BOTTOM, gwu.TB_PLACEMENT_LEFT}
	halignslb := gwu.NewListBox([]string{"Left", "Center", "Right"})
	haligns := []gwu.HAlign{gwu.HA_LEFT, gwu.HA_CENTER, gwu.HA_RIGHT}
	valignslb := gwu.NewListBox([]string{"Top", "Middle", "Bottom"})
	valigns := []gwu.VAlign{gwu.VA_TOP, gwu.VA_MIDDLE, gwu.VA_BOTTOM}
	placemslb.Style().SetFullWidth()
	halignslb.Style().SetFullWidth()
	valignslb.Style().SetFullWidth()
	table.Add(placemslb, 0, 1)
	table.Add(halignslb, 1, 1)
	table.Add(valignslb, 2, 1)

	placemslb.AddEHandlerFunc(func(e gwu.Event) {
		t.SetTabBarPlacement(placems[placemslb.SelectedIdx()])
		e.MarkDirty(t)
	}, gwu.ETYPE_CHANGE)
	halignslb.AddEHandlerFunc(func(e gwu.Event) {
		t.TabBarFmt().SetHAlign(haligns[halignslb.SelectedIdx()])
		e.MarkDirty(t)
	}, gwu.ETYPE_CHANGE)
	valignslb.AddEHandlerFunc(func(e gwu.Event) {
		t.TabBarFmt().SetVAlign(valigns[valignslb.SelectedIdx()])
		e.MarkDirty(t)
	}, gwu.ETYPE_CHANGE)

	p.Add(table)

	fix := gwu.NewCheckBox("Fixed size")
	fix.SetState(true)
	fix.AddEHandlerFunc(func(e gwu.Event) {
		if fix.State() {
			t.Style().SetSizePx(500, 300)
		} else {
			t.Style().SetSize("", "")
		}
		e.MarkDirty(t)
	}, gwu.ETYPE_CLICK)
	p.Add(fix)

	p.AddVSpace(10)
	l := gwu.NewLabel("Click on tabs...")
	l.Style().SetColor(gwu.CLR_GREEN)
	p.Add(l)
	t.AddEHandlerFunc(func(e gwu.Event) {
		l.SetText("Clicked on tab: " + strconv.Itoa(t.Selected()))
		e.MarkDirty(l)
	}, gwu.ETYPE_STATE_CHANGE)
	p.AddVSpace(10)
	c := gwu.NewPanel()
	c.Add(gwu.NewLabel("This is a TabPanel."))
	c.Add(gwu.NewLabel("Click on other tabs to see their content."))
	c.AddVSpace(15)
	c.Add(gwu.NewLabel("Or click here to see what's in the Hollow:"))
	b := gwu.NewButton("Take me to the Hollow!")
	b.AddEHandlerFunc(func(e gwu.Event) {
		t.SetSelected(3)
		e.MarkDirty(t)
	}, gwu.ETYPE_CLICK)
	c.Add(b)
	t.AddString("Home", c)
	c = gwu.NewPanel()
	c.Add(gwu.NewLabel("You have no new messages."))
	t.AddString("Inbox", c)
	c = gwu.NewPanel()
	c.Add(gwu.NewLabel("You have no sent messages."))
	t.AddString("Sent", c)
	c = gwu.NewPanel()
	c.Add(gwu.NewLabel("There is nothing in the hollow."))
	t.AddString("Hollow", c)
	c = gwu.NewPanel()
	tb := gwu.NewTextBox("Click to edit this comment.")
	tb.SetRows(10)
	tb.SetCols(40)
	c.Add(tb)
	t.AddString("Comment", c)
	p.Add(t)

	return p
}

func buildWindowDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("The Window represents the whole window, the page inside the browser."))
	p.AddVSpace(5)
	p.Add(gwu.NewLabel("The Window is the top of the component hierarchy. It is an extension of the Panel."))

	return p
}

func buildCheckBoxDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	suml := gwu.NewLabel("")

	p.Add(gwu.NewLabel("Check the days you want to work on:"))

	cbs := []gwu.CheckBox{gwu.NewCheckBox("Monday"), gwu.NewCheckBox("Tuesday"), gwu.NewCheckBox("Wednesday"),
		gwu.NewCheckBox("Thursday"), gwu.NewCheckBox("Friday"), gwu.NewCheckBox("Saturday"), gwu.NewCheckBox("Sunday")}
	cbs[5].SetEnabled(false)
	cbs[6].SetEnabled(false)

	for _, cb := range cbs {
		p.Add(cb)
		cb.AddEHandlerFunc(func(e gwu.Event) {
			sum := 0
			for _, cb2 := range cbs {
				if cb2.State() {
					sum++
				}
			}
			suml.SetText(fmt.Sprintf("%d day%s is a total of %d hours a week.", sum, plural(sum), sum*8))
			e.MarkDirty(suml)
		}, gwu.ETYPE_CLICK)
	}

	p.Add(suml)

	return p
}

func buildListBoxDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	row := gwu.NewHorizontalPanel()
	l := gwu.NewLabel("Select a background color:")
	row.Add(l)
	lb := gwu.NewListBox([]string{"", "Black", "Red", "Green", "Blue", "White"})
	lb.AddEHandlerFunc(func(e gwu.Event) {
		l.Style().SetBackground(lb.SelectedValue())
		e.MarkDirty(l)
	}, gwu.ETYPE_CHANGE)
	row.Add(lb)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Select numbers that add up to 89:"))
	sumLabel := gwu.NewLabel("")
	lb2 := gwu.NewListBox([]string{"1", "2", "4", "8", "16", "32", "64", "128"})
	lb2.SetMulti(true)
	lb2.SetRows(10)
	lb2.AddEHandlerFunc(func(e gwu.Event) {
		sum := 0
		for _, idx := range lb2.SelectedIndices() {
			sum += 1 << uint(idx)
		}
		if sum == 89 {
			sumLabel.SetText("Hooray! You did it!")
		} else {
			sumLabel.SetText(fmt.Sprintf("Now quite there... (sum = %d)", sum))
		}
		e.MarkDirty(sumLabel)
	}, gwu.ETYPE_CHANGE)
	p.Add(lb2)
	p.Add(sumLabel)

	return p
}

func buildTextBoxDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Input Key(max 32 characters):"))
	row := gwu.NewHorizontalPanel()
	txBox_des_key := gwu.NewTextBox("")
	txBox_des_key.SetCols(32)
	txBox_des_key.SetMaxLength(32)
	txBox_des_key.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	length := gwu.NewLabel("")
	length.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_des_key.AddEHandlerFunc(func(e gwu.Event) {
		rem := 32 - len(txBox_des_key.Text())
		length.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_des_key)
	row.Add(length)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Input Data(max 32 characters):"))
	row = gwu.NewHorizontalPanel()
	txBox_des_data := gwu.NewTextBox("")
	txBox_des_data.SetCols(32)
	txBox_des_data.SetMaxLength(32)
	txBox_des_data.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	length_2 := gwu.NewLabel("")
	length_2.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_des_data.AddEHandlerFunc(func(e gwu.Event) {
		rem := 32 - len(txBox_des_data.Text())
		length_2.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length_2)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_des_data)
	row.Add(length_2)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Encrypt Result:"))
	txBox_des_result_encrypt := gwu.NewTextBox("")
	txBox_des_result_encrypt.SetCols(32)
	p.Add(txBox_des_result_encrypt)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Decrypt Result:"))
	txBox_des_result_decrypt := gwu.NewTextBox("")
	txBox_des_result_decrypt.SetCols(32)
	p.Add(txBox_des_result_decrypt)

	/*p.AddVSpace(10
	p.Add(gwu.NewLabel("Short biography:"))
	bio := gwu.NewTextBox("")
	bio.SetRows(5)
	bio.SetCols(40)
	p.Add(bio)

	p.AddVSpace(10)
	rtb := gwu.NewTextBox("This is just a read-only text box...")
	rtb.SetReadOnly(true)
	p.Add(rtb)

	p.AddVSpace(10)
	dtb := gwu.NewTextBox("...and a disabled one.")
	dtb.SetEnabled(false)
	p.Add(dtb)*/

	return p
}

func buildConvertView(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Input Ascii(max 1024 characters):"))
	row := gwu.NewHorizontalPanel()
	txBox_ascii := gwu.NewTextBox("")
	txBox_ascii.SetRows(8)
	txBox_ascii.SetCols(128)
	txBox_ascii.SetMaxLength(1024)
	txBox_ascii.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	length := gwu.NewLabel("")
	length.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_ascii.AddEHandlerFunc(func(e gwu.Event) {
		rem := 1024 - len(txBox_ascii.Text())
		length.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_ascii)
	row.Add(length)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Input Hex(max 1024 characters):"))
	row = gwu.NewHorizontalPanel()
	txBox_hex := gwu.NewTextBox("")
	txBox_hex.SetRows(8)
	txBox_hex.SetCols(128)
	txBox_hex.SetMaxLength(1024)
	txBox_hex.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	length_3 := gwu.NewLabel("")
	length_3.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_hex.AddEHandlerFunc(func(e gwu.Event) {
		rem := 1024 - len(txBox_hex.Text())
		length_3.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length_3)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_hex)
	row.Add(length_3)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Input XorSum Data(max 1024 characters):"))
	row = gwu.NewHorizontalPanel()
	txBox_xor_data := gwu.NewTextBox("")
	txBox_xor_data.SetRows(8)
	txBox_xor_data.SetCols(128)
	txBox_xor_data.SetMaxLength(1024)
	txBox_xor_data.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	length_4 := gwu.NewLabel("")
	length_4.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_xor_data.AddEHandlerFunc(func(e gwu.Event) {
		rem := 1024 - len(txBox_xor_data.Text())
		length_4.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length_4)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_xor_data)
	row.Add(length_4)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("XorSum Result:"))
	txBox_xor_result := gwu.NewTextBox("")
	txBox_xor_result.SetCols(128)
	p.Add(txBox_xor_result)

	return p
}

func buildDesView(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Input Key(max 32 characters):"))
	row := gwu.NewHorizontalPanel()
	txBox_des_key := gwu.NewTextBox("")
	txBox_des_key.SetCols(32)
	txBox_des_key.SetMaxLength(32)
	txBox_des_key.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	length := gwu.NewLabel("")
	length.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_des_key.AddEHandlerFunc(func(e gwu.Event) {
		rem := 32 - len(txBox_des_key.Text())
		length.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_des_key)
	row.Add(length)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Input Data(max 32 characters):"))
	row = gwu.NewHorizontalPanel()
	txBox_des_data := gwu.NewTextBox("")
	txBox_des_data.SetCols(32)
	txBox_des_data.SetMaxLength(32)
	txBox_des_data.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	length_2 := gwu.NewLabel("")
	length_2.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_des_data.AddEHandlerFunc(func(e gwu.Event) {
		rem := 32 - len(txBox_des_data.Text())
		length_2.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length_2)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_des_data)
	row.Add(length_2)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Encrypt Result:"))
	txBox_des_result_encrypt := gwu.NewTextBox("")
	txBox_des_result_encrypt.SetCols(32)
	p.Add(txBox_des_result_encrypt)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Decrypt Result:"))
	txBox_des_result_decrypt := gwu.NewTextBox("")
	txBox_des_result_decrypt.SetCols(32)
	p.Add(txBox_des_result_decrypt)

	return p
}

func buildMacView(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Input Key(max 32 characters):"))
	row := gwu.NewHorizontalPanel()
	txBox_des_key := gwu.NewTextBox("")
	txBox_des_key.SetCols(32)
	txBox_des_key.SetMaxLength(32)
	txBox_des_key.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	length := gwu.NewLabel("")
	length.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_des_key.AddEHandlerFunc(func(e gwu.Event) {
		rem := 32 - len(txBox_des_key.Text())
		length.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_des_key)
	row.Add(length)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Input Data(max 2048 characters):"))
	row = gwu.NewHorizontalPanel()
	txBox_des_data := gwu.NewTextBox("")
	txBox_des_data.SetRows(16)
	txBox_des_data.SetCols(128)
	txBox_des_data.SetMaxLength(2048)
	txBox_des_data.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	length_2 := gwu.NewLabel("")
	length_2.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_des_data.AddEHandlerFunc(func(e gwu.Event) {
		rem := 2048 - len(txBox_des_data.Text())
		length_2.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length_2)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_des_data)
	row.Add(length_2)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Input Init Data(max 16 characters):"))
	row = gwu.NewHorizontalPanel()
	txBox_init_data := gwu.NewTextBox("")
	txBox_init_data.SetCols(16)
	txBox_init_data.SetMaxLength(16)
	txBox_init_data.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	txBox_init_data.SetText("0000000000000000")
	length_3 := gwu.NewLabel("")
	length_3.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_init_data.AddEHandlerFunc(func(e gwu.Event) {
		rem := 16 - len(txBox_init_data.Text())
		length_3.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length_3)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_init_data)
	row.Add(length_3)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("DES/3DES PBOC Result:"))
	txBox_pboc_result_encrypt := gwu.NewTextBox("")
	txBox_pboc_result_encrypt.SetCols(32)
	p.Add(txBox_pboc_result_encrypt)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("DES/3DES POS ECB Result:"))
	txBox_pos_result_decrypt := gwu.NewTextBox("")
	txBox_pos_result_decrypt.SetCols(32)
	p.Add(txBox_pos_result_decrypt)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("DES/3DES ANSI X9.9 Result:"))
	txBox_x99_result_decrypt := gwu.NewTextBox("")
	txBox_x99_result_decrypt.SetCols(32)
	p.Add(txBox_x99_result_decrypt)

	return p
}

func buildRsaView(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Input Public Key(max 1024 characters):"))
	row := gwu.NewHorizontalPanel()
	txBox_public_key := gwu.NewTextBox("")
	txBox_public_key.SetRows(8)
	txBox_public_key.SetCols(128)
	txBox_public_key.SetMaxLength(1024)
	txBox_public_key.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	txBox_public_key.SetText("B417E5AB8961649679F037EE0799BE6BF5478203DE483832A4F0C906DC2064A451DF3913AEFE700D8D6F2729EB24600E4DF59CF54863C7C39016883CB6DD800A239FDADA5E227E130F633B7171FBF3CA8A27722CFD0DDC1A5B7FA02E161FB4C16FC1306F3DE9CE2F1D5E376E484BD77CAB8377F379805219EA9855672DB7C6B12D5E273D8BD6A2C3B3FD67E6BCD4A9795510E2E59D83ECA46AB1DEA0C31ED8AD8BB052F84DBD5FE5BA2A58E53C0C65ED82797CF8EA04F6E382EB64EBD096E87B34C59BF94CA1E8329F10DB3D05E124B6810F92F58AF900079E66AF566C8E985AD729FDA28637FADC80B11C3AD550217214A612B8E7B1454D")
	length := gwu.NewLabel("")
	length.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_public_key.AddEHandlerFunc(func(e gwu.Event) {
		rem := 1024 - len(txBox_public_key.Text())
		length.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_public_key)
	row.Add(length)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Input Expoent(max 6 characters):"))
	row = gwu.NewHorizontalPanel()
	txBox_expoent_data := gwu.NewTextBox("")
	txBox_expoent_data.SetCols(6)
	txBox_expoent_data.SetMaxLength(6)
	txBox_expoent_data.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	txBox_expoent_data.SetText("010001")
	length_2 := gwu.NewLabel("")
	length_2.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_expoent_data.AddEHandlerFunc(func(e gwu.Event) {
		rem := 6 - len(txBox_expoent_data.Text())
		length_2.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length_2)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_expoent_data)
	row.Add(length_2)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Input Private Key(max 1024 characters):"))
	row = gwu.NewHorizontalPanel()
	txBox_private_key := gwu.NewTextBox("")
	txBox_private_key.SetRows(8)
	txBox_private_key.SetCols(128)
	txBox_private_key.SetMaxLength(1024)
	txBox_private_key.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	txBox_private_key.SetText("29EA29ED0B935C4ADD64784FF0FAEA63CBF9719C1EDEE61D7B60031992A0036C6F15A9FE17CB661EF15CFE5B763BC8136C378EE95388CD948973850FC46BB702251A6E60466A76B74EEC848839C42F1063A91C9CB5E4C4E6698741AB984096DC327B5F1F74E476B5667B682145B84915DE014228AB8D9F0E2318833F951C4E74B0691D9062BC668EA333141049FDCC6EBC1049DD001E02C5CE5C45CA65DEE1468BFC203E6AD50B756E9867465470D039BD4A4253592E0BEC45A51B7F48861117CA369A10A00007D4DA8E3A69FCFFD0EE8E350AC97D3EB5D5B56BDAD2919AA05927FB6AEA9326279D74A8B144A742CE3727B3545D56DABAA1")
	length_3 := gwu.NewLabel("")
	length_3.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_private_key.AddEHandlerFunc(func(e gwu.Event) {
		rem := 1024 - len(txBox_private_key.Text())
		length_3.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length_3)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_private_key)
	row.Add(length_3)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Input Data(max 1024 characters):"))
	row = gwu.NewHorizontalPanel()
	txBox_rsa_data := gwu.NewTextBox("")
	txBox_rsa_data.SetRows(8)
	txBox_rsa_data.SetCols(128)
	txBox_rsa_data.SetMaxLength(1024)
	txBox_rsa_data.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	txBox_rsa_data.SetText("35E262DD2FC4CF5BF284D3EE5B53DFF287A80C814DCB48FAA565AEB56EF7D246929D3AF8CA60DCB14A24F6E94C44DFE0732B4FEF4CDA86F5D7096D3080BBA7F28612373046AAC0A48F1F7E81A2A4B403D8CAE4C8F028D9958BE249D4C72E846C53844D23D95389912D14D420292F9ACEF35D121ED0BAE34CEEFB9378EB2A9A5E96F848257BCD2456972142A54E27B8175C74BD3ED885652120CFBBCF1B7F2B33DDE928FA9F3E2F6FB9ED58F84CC026CA2D2CE030641D86596D0741521C04D9F5F3A8E49FB751F9EF3CF1641A6DCA7AECC6C1469E2E9B90424DFF28F2740F36319999D3A3F4306E133F4C5B8065DB70BDCB96205F7D13DE00")
	length_4 := gwu.NewLabel("")
	length_4.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_rsa_data.AddEHandlerFunc(func(e gwu.Event) {
		rem := 1024 - len(txBox_rsa_data.Text())
		length_4.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length_4)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_rsa_data)
	row.Add(length_4)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Encrypt Result:"))
	txBox_rsa_result_encrypt := gwu.NewTextBox("")
	txBox_rsa_result_encrypt.SetRows(8)
	txBox_rsa_result_encrypt.SetCols(128)
	p.Add(txBox_rsa_result_encrypt)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Decrypt Result:"))
	txBox_rsa_result_decrypt := gwu.NewTextBox("")
	txBox_rsa_result_decrypt.SetRows(8)
	txBox_rsa_result_decrypt.SetCols(128)
	p.Add(txBox_rsa_result_decrypt)

	return p
}

func buildPasswBoxDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Enter your password:"))
	p.Add(gwu.NewPasswBox(""))

	return p
}

func buildRadioButtonDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Select your favorite programming language:"))

	group := gwu.NewRadioGroup("lang")
	rbs := []gwu.RadioButton{gwu.NewRadioButton("Go", group), gwu.NewRadioButton("Java", group), gwu.NewRadioButton("C / C++", group),
		gwu.NewRadioButton("Python", group), gwu.NewRadioButton("QBasic (nah this can't be your favorite)", group)}
	rbs[4].SetEnabled(false)

	for _, rb := range rbs {
		p.Add(rb)
	}

	p.AddVSpace(20)
	p.Add(gwu.NewLabel("Select your favorite computer game:"))

	group = gwu.NewRadioGroup("game")
	rbs = []gwu.RadioButton{gwu.NewRadioButton("StarCraft II", group), gwu.NewRadioButton("Minecraft", group),
		gwu.NewRadioButton("Other", group)}

	for _, rb := range rbs {
		p.Add(rb)
	}

	return p
}

func buildSwitchButtonDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()
	p.SetCellPadding(1)

	row := gwu.NewHorizontalPanel()
	row.Add(gwu.NewLabel("Here's an ON/OFF switch which enables/disables the other one:"))
	sw := gwu.NewSwitchButton()
	sw.SetOnOff("ENB", "DISB")
	sw.SetState(true)
	row.Add(sw)
	p.Add(row)

	p.AddVSpace(10)
	row = gwu.NewHorizontalPanel()
	row.Add(gwu.NewLabel("And the other one:"))
	sw2 := gwu.NewSwitchButton()
	sw2.SetEnabled(true)
	sw2.Style().SetWidthPx(100)
	row.Add(sw2)
	sw.AddEHandlerFunc(func(e gwu.Event) {
		sw2.SetEnabled(sw.State())
		e.MarkDirty(sw2)
	}, gwu.ETYPE_CLICK)
	p.Add(row)

	return p
}

func buildButtonDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	l := gwu.NewLabel("")

	btnp := gwu.NewHorizontalPanel()
	b := gwu.NewButton("Normal Button")
	b.AddEHandlerFunc(func(e gwu.Event) {
		switch e.Type() {
		case gwu.ETYPE_MOUSE_OVER:
			l.SetText("Mouse is over...")
		case gwu.ETYPE_MOUSE_OUT:
			l.SetText("Mouse is out.")
		case gwu.ETYPE_CLICK:
			x, y := e.Mouse()
			l.SetText(fmt.Sprintf("Clicked at x=%d, y=%d", x, y))
		}
		e.MarkDirty(l)
	}, gwu.ETYPE_CLICK, gwu.ETYPE_MOUSE_OVER, gwu.ETYPE_MOUSE_OUT)
	btnp.Add(b)

	b = gwu.NewButton("Disabled Button")
	b.SetEnabled(false)
	btnp.Add(b)

	p.Add(btnp)

	p.Add(l)

	return p
}

func buildHtmlDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	html := `<span onclick="alert('Hi from Html!');">Hi! I'm inserted as HTML. Click on me!</span>`

	p.Add(gwu.NewLabel("The following HTML code is inserted after the text box as an Html component:"))
	ta := gwu.NewTextBox(html)
	ta.SetReadOnly(true)
	ta.Style().SetWidthPx(500)
	ta.SetRows(4)
	p.Add(ta)

	p.AddVSpace(20)
	h := gwu.NewHtml(html)
	p.Add(h)

	return p
}

func buildImageDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Google's logo:"))
	img := gwu.NewImage("Google's logo", "https://www.google.com/images/srpr/logo3w.png")
	img.Style().SetSizePx(275, 95)
	p.Add(img)

	p.AddVSpace(20)
	p.Add(gwu.NewLabel("Go's Gopher:"))
	img = gwu.NewImage("Go's Gopher", "http://golang.org/doc/gopher/frontpage.png")
	img.Style().SetSizePx(250, 340)
	p.Add(img)

	return p
}

func buildLabelDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("This is a Label."))
	p.Add(gwu.NewLabel("世界 And another one. ㅈㅈ"))
	p.Add(gwu.NewLabel("Nothing special about them, but they may be the mostly used components."))

	p.AddVSpace(20)
	p.Add(gwu.NewLabel("You can change their text:"))
	b := gwu.NewButton("Change!")
	b.AddEHandlerFunc(func(e gwu.Event) {
		for i := 0; i < p.CompsCount(); i++ {
			if l, ok := p.CompAt(i).(gwu.Label); ok && l != b {
				reversed := []rune(l.Text())
				for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
					reversed[i], reversed[j] = reversed[j], reversed[i]
				}
				l.SetText(string(reversed))
			}
		}
		e.MarkDirty(p)
	}, gwu.ETYPE_CLICK)
	p.Add(b)

	return p
}

func buildLinkDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()
	p.SetCellPadding(3)

	p.Add(gwu.NewLink("Visit Gowut Home page", "https://sites.google.com/site/gowebuitoolkit/"))
	p.Add(gwu.NewLink("Visit Gowut Project page", "https://github.com/icza/gowut"))

	row := gwu.NewHorizontalPanel()
	row.SetCellPadding(3)
	row.Add(gwu.NewLabel("Discussion forum:"))
	row.Add(gwu.NewLink("https://groups.google.com/d/forum/gowebuitoolkit", "https://groups.google.com/d/forum/gowebuitoolkit"))
	p.Add(row)

	row = gwu.NewHorizontalPanel()
	row.SetCellPadding(3)
	row.Add(gwu.NewLabel("Send e-mail to the Gowut author:"))
	email := "iczaaa" + "@" + "gmail.com"
	row.Add(gwu.NewLink("András Belicza <"+email+">", "mailto:"+email))
	p.Add(row)

	return p
}

func buildTimerDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()
	p.SetCellPadding(3)

	// Add timers to a panel which is always attached instead of our panel
	// because the user can switch to another component demo causing this panel to be removed
	// and that way timer events would address components that are not part of the window (returning error).
	hiddenPan := event.Session().Attr("hiddenPan").(gwu.Panel)

	p.Add(gwu.NewLabel("A Timer is used to detonate a bomb after 3 seconds."))
	p.AddVSpace(10)
	defText := "You can defuse the bomb with the button below. Tick... Tack..."
	l := gwu.NewLabel(defText)
	p.Add(l)
	t := gwu.NewTimer(3 * time.Second)
	b := gwu.NewButton("Defuse!")
	t.AddEHandlerFunc(func(e gwu.Event) {
		l.SetText("BOOOOM! You were too slow!")
		l.Style().SetColor(gwu.CLR_RED)
		b.SetEnabled(false)
		e.MarkDirty(l, b)
	}, gwu.ETYPE_STATE_CHANGE)
	hiddenPan.Add(t)
	row := gwu.NewHorizontalPanel()
	b.AddEHandlerFunc(func(e gwu.Event) {
		t.SetActive(false)
		l.SetText("Bomb defused! Phew! Good Job!")
		l.Style().SetColor(gwu.CLR_GREEN)
		b.SetEnabled(false)
		e.MarkDirty(t, l, b)
	}, gwu.ETYPE_CLICK)
	row.Add(b)
	b2 := gwu.NewButton("Plant a new Bomb!")
	b2.AddEHandlerFunc(func(e gwu.Event) {
		t.SetActive(true)
		t.Reset()
		l.SetText(defText)
		l.Style().SetColor("")
		b.SetEnabled(true)
		e.MarkDirty(t, l, b)
	}, gwu.ETYPE_CLICK)
	row.Add(b2)
	p.Add(row)

	p.AddVSpace(20)
	p.Add(gwu.NewLabel("A Timer is used to refresh the time below repeatedly in every second for half a minute."))
	tl := gwu.NewLabel("")
	p.Add(tl)
	t2 := gwu.NewTimer(time.Second)
	t2.SetRepeat(true)
	counter := 30
	t2.AddEHandlerFunc(func(e gwu.Event) {
		counter--
		tl.SetText(fmt.Sprintf("%s (%d remaining)", time.Now().Format("2006-01-02 15:04:05"), counter))
		e.MarkDirty(tl)
		if counter <= 0 {
			t2.SetActive(false)
			e.MarkDirty(t2)
		}
	}, gwu.ETYPE_STATE_CHANGE)
	hiddenPan.Add(t2)
	b3 := gwu.NewButton("Restart")
	b3.AddEHandlerFunc(func(e gwu.Event) {
		counter = 30
		t2.SetActive(true)
		e.MarkDirty(t2)
	}, gwu.ETYPE_CLICK)
	p.Add(b3)

	event.MarkDirty(hiddenPan)

	return p
}

type demo struct {
	link      gwu.Label
	buildFunc func(gwu.Event) gwu.Comp
	comp      gwu.Comp // Lazily initialized demo comp
}
type pdemo *demo

func buildShowcaseWin(sess gwu.Session) {
	win := gwu.NewWindow("web_view", "DC_PBOC_CLIENT Go Go Go")
	win.Style().SetFullSize()
	win.AddEHandlerFunc(func(e gwu.Event) {
		switch e.Type() {
		case gwu.ETYPE_WIN_LOAD:
			fmt.Println("LOADING window:", e.Src().Id())
		case gwu.ETYPE_WIN_UNLOAD:
			fmt.Println("UNLOADING window:", e.Src().Id())
		}
	}, gwu.ETYPE_WIN_LOAD, gwu.ETYPE_WIN_UNLOAD)

	hiddenPan := gwu.NewNaturalPanel()
	sess.SetAttr("hiddenPan", hiddenPan)

	header := gwu.NewHorizontalPanel()
	header.Style().SetFullWidth().SetBorderBottom2(2, gwu.BRD_STYLE_SOLID, "#777777")
	l := gwu.NewLabel("Welcome to use this version !")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetFontSize("120%")
	header.Add(l)
	/*
		header.AddHConsumer()
		header.Add(gwu.NewLabel("Theme:"))
		themes := gwu.NewListBox([]string{"default", "debug"})
		themes.AddEHandlerFunc(func(e gwu.Event) {
			win.SetTheme(themes.SelectedValue())
			e.ReloadWin("show")
		}, gwu.ETYPE_CHANGE)
		header.Add(themes)
		header.AddHSpace(10)
		reset := gwu.NewLink("Reset", "#")
		reset.SetTarget("")
		reset.AddEHandlerFunc(func(e gwu.Event) {
			e.RemoveSess()
			e.ReloadWin("show")
		}, gwu.ETYPE_CLICK)
		header.Add(reset)
	*/
	setNoWrap(header)
	win.Add(header)

	content := gwu.NewHorizontalPanel()
	content.SetCellPadding(1)
	content.SetVAlign(gwu.VA_TOP)
	content.Style().SetFullSize()

	demoWrapper := gwu.NewPanel()
	demoWrapper.Style().SetPaddingLeftPx(5)
	demoWrapper.AddVSpace(10)
	demoTitle := gwu.NewLabel("")
	demoTitle.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetFontSize("110%")
	demoWrapper.Add(demoTitle)
	demoWrapper.AddVSpace(10)

	links := gwu.NewPanel()
	links.SetCellPadding(1)
	links.Style().SetPaddingRightPx(5)

	demos := make(map[string]pdemo)
	var selDemo pdemo

	selectDemo := func(d pdemo, e gwu.Event) {
		if selDemo != nil {
			selDemo.link.Style().SetBackground("")
			if e != nil {
				e.MarkDirty(selDemo.link)
			}
			demoWrapper.Remove(selDemo.comp)
		}
		selDemo = d
		d.link.Style().SetBackground("#88ff88")
		demoTitle.SetText(d.link.Text())
		if d.comp == nil {
			d.comp = d.buildFunc(e)
		}
		demoWrapper.Add(d.comp)
		if e != nil {
			e.MarkDirty(d.link, demoWrapper)
		}
	}

	createDemo := func(name string, buildFunc func(gwu.Event) gwu.Comp) pdemo {
		link := gwu.NewLabel(name)
		link.Style().SetFullWidth().SetCursor(gwu.CURSOR_POINTER).SetDisplay(gwu.DISPLAY_BLOCK).SetColor(gwu.CLR_BLUE)
		demo := &demo{link: link, buildFunc: buildFunc}
		link.AddEHandlerFunc(func(e gwu.Event) {
			selectDemo(demo, e)
		}, gwu.ETYPE_CLICK)
		links.Add(link)
		demos[name] = demo
		return demo
	}

	links.Style().SetFullHeight().SetBorderRight2(2, gwu.BRD_STYLE_SOLID, "#777777")
	links.AddVSpace(5)
	homeDemo := createDemo("Home", buildHomeDemo)
	selectDemo(homeDemo, nil)

	/*
		links.AddVSpace(5)
		l = gwu.NewLabel("Module Demo")
		l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetFontSize("110%")
		links.Add(l)
	*/

	links.AddVSpace(5)
	l = gwu.NewLabel("IC Card")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD)
	links.Add(l)
	//createDemo("Expander", buildExpanderDemo)
	//createDemo("Link (as Container)", buildLinkContainerDemo)
	//createDemo("Panel", buildPanelDemo)
	//createDemo("Table", buildTableDemo)
	createDemo("PBOC3.0", buildTabPanelDemo)
	//createDemo("Window", buildWindowDemo)

	links.AddVSpace(5)
	l = gwu.NewLabel("8583")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetDisplay(gwu.DISPLAY_BLOCK)
	links.Add(l)
	//createDemo("CheckBox", buildCheckBoxDemo)
	//createDemo("ListBox", buildListBoxDemo)
	createDemo("Analysis", buildTextBoxDemo)
	//createDemo("PasswBox", buildPasswBoxDemo)
	//createDemo("RadioButton", buildRadioButtonDemo)
	//createDemo("SwitchButton", buildSwitchButtonDemo)

	links.AddVSpace(5)
	l = gwu.NewLabel("TLV")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD)
	links.Add(l)
	createDemo("Analysis", buildListBoxDemo)
	//createDemo("Button", buildButtonDemo)
	//createDemo("Html", buildHtmlDemo)
	//createDemo("Image", buildImageDemo)
	//createDemo("Label", buildLabelDemo)
	//createDemo("Link", buildLinkDemo)
	//createDemo("Timer", buildTimerDemo)

	links.AddVSpace(5)
	l = gwu.NewLabel("STRING")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD)
	links.Add(l)
	createDemo("Convert", buildConvertView)

	links.AddVSpace(5)
	l = gwu.NewLabel("CALCULATE")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD)
	links.Add(l)
	createDemo("Des", buildDesView)
	createDemo("Mac", buildMacView)
	createDemo("Rsa", buildRsaView)

	links.AddVConsumer()
	setNoWrap(links)
	content.Add(links)
	content.Add(demoWrapper)
	content.CellFmt(demoWrapper).Style().SetFullWidth()

	win.Add(content)
	win.CellFmt(content).Style().SetFullSize()

	footer := gwu.NewHorizontalPanel()
	footer.Style().SetFullWidth().SetBorderTop2(2, gwu.BRD_STYLE_SOLID, "#777777")
	footer.Add(hiddenPan)
	footer.AddHConsumer()
	l = gwu.NewLabel("Copyright © 2016-2016 Qfx. All rights reserved.")
	l.Style().SetFontStyle(gwu.FONT_STYLE_ITALIC).SetFontSize("95%")
	footer.Add(l)
	footer.AddHSpace(10)
	link := gwu.NewLink("Visit Qinuu Home page", "http://www.qinuu.com/")
	link.Style().SetFontStyle(gwu.FONT_STYLE_ITALIC).SetFontSize("95%")
	footer.Add(link)
	setNoWrap(footer)
	win.Add(footer)

	sess.AddWin(win)
}

// setNoWrap sets WHITE_SPACE_NOWRAP to all children of the specified panel.
func setNoWrap(panel gwu.Panel) {
	count := panel.CompsCount()
	for i := count - 1; i >= 0; i-- {
		panel.CompAt(i).Style().SetWhiteSpace(gwu.WHITE_SPACE_NOWRAP)
	}
}

// SessHandler is our session handler to build the showcases window.
type SessHandler struct{}

func (h SessHandler) Created(s gwu.Session) {
	buildShowcaseWin(s)
}

func (h SessHandler) Removed(s gwu.Session) {}

func main() {
	// Allow app control from command line (in co-operation with the starter script):
	fmt.Println("Type 'r' to restart, 'e' to exit.")
	go func() {
		var cmd string
		for {
			fmt.Scanf("%s", &cmd)
			switch cmd {
			case "r": // restart
				os.Exit(1)
			case "e": // exit
				os.Exit(0)
			}
		}
	}()

	// Create GUI server
	server := gwu.NewServer("dc_pboc_client", "")
	server.SetText("")

	server.AddSessCreatorName("web_view", "")
	server.AddSHandler(SessHandler{})

	// Start GUI server
	if err := server.Start("web_view"); err != nil {
		fmt.Println("Error: Cound not start GUI server:", err)
		return
	}
}
