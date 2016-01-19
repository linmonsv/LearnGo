package main

import (
	"errors"
	"fmt"
	"github.com/icza/gowut/gwu"
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

func buildPboc3View(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("This demo need the related hardware support!"))

	p.Add(gwu.NewLabel("Select your device:"))

	group := gwu.NewRadioGroup("device")
	rbs := []gwu.RadioButton{
		gwu.NewRadioButton("T10", group),
		gwu.NewRadioButton("D8", group),
		gwu.NewRadioButton("D6", group),
		gwu.NewRadioButton("Z9", group)}

	for _, rb := range rbs {
		p.Add(rb)
	}

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Select Usb/Serial:"))

	group = gwu.NewRadioGroup("commtype")
	rbs = []gwu.RadioButton{
		gwu.NewRadioButton("USB", group),
		gwu.NewRadioButton("COM1", group),
		gwu.NewRadioButton("COM2", group),
		gwu.NewRadioButton("COM3", group),
		gwu.NewRadioButton("COM4", group)}

	for _, rb := range rbs {
		p.Add(rb)
	}

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Select Usb/Serial:"))

	group = gwu.NewRadioGroup("transtype")
	rbs = []gwu.RadioButton{
		gwu.NewRadioButton("ReadCardInfo", group),
		gwu.NewRadioButton("ARQC/ARPC", group),
		gwu.NewRadioButton("ReadCardLog", group)}

	for _, rb := range rbs {
		p.Add(rb)
	}

	txBox_result := gwu.NewTextBox("")

	p.AddVSpace(10)
	b := gwu.NewButton("Start")
	b.AddEHandlerFunc(func(e gwu.Event) {
		switch e.Type() {
		case gwu.ETYPE_MOUSE_OVER:
			txBox_result.SetText("Mouse is over...")
		case gwu.ETYPE_MOUSE_OUT:
			txBox_result.SetText("Mouse is out.")
		case gwu.ETYPE_CLICK:
			x, y := e.Mouse()
			txBox_result.SetText(fmt.Sprintf("Clicked at x=%d, y=%d", x, y))
		}
		e.MarkDirty(txBox_result)
	}, gwu.ETYPE_CLICK, gwu.ETYPE_MOUSE_OVER, gwu.ETYPE_MOUSE_OUT)
	p.Add(b)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Result:"))

	txBox_result.SetRows(8)
	txBox_result.SetCols(128)
	p.Add(txBox_result)

	return p
}

func build8583AnalysisView(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Input 8583 Data(max 2048 characters):"))
	row := gwu.NewHorizontalPanel()
	txBox_ascii := gwu.NewTextBox("")
	txBox_ascii.SetRows(8)
	txBox_ascii.SetCols(128)
	txBox_ascii.SetMaxLength(2048)
	txBox_ascii.AddSyncOnETypes(gwu.ETYPE_KEY_UP)

	txBox_result := gwu.NewTextBox("")
	txBox_result.SetRows(32)
	txBox_result.SetCols(128)
	txBox_result.AddSyncOnETypes(gwu.ETYPE_KEY_UP)

	length := gwu.NewLabel("")
	length.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_ascii.AddEHandlerFunc(func(e gwu.Event) {
		txBox_result.SetText(txBox_ascii.Text())
		e.MarkDirty(txBox_result)
		rem := 2048 - len(txBox_ascii.Text())
		length.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_ascii)
	row.Add(length)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("The Result:"))
	row = gwu.NewHorizontalPanel()

	row.Add(txBox_result)
	p.Add(row)

	return p
}

func buildTlvAnalysisView(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Input Tlv Data(max 1024 characters):"))
	row := gwu.NewHorizontalPanel()
	txBox_ascii := gwu.NewTextBox("")
	txBox_ascii.SetRows(8)
	txBox_ascii.SetCols(128)
	txBox_ascii.SetMaxLength(1024)
	txBox_ascii.AddSyncOnETypes(gwu.ETYPE_KEY_UP)

	txBox_result := gwu.NewTextBox("")
	txBox_result.SetRows(16)
	txBox_result.SetCols(128)
	txBox_result.AddSyncOnETypes(gwu.ETYPE_KEY_UP)

	length := gwu.NewLabel("")
	length.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_ascii.AddEHandlerFunc(func(e gwu.Event) {
		txBox_result.SetText(txBox_ascii.Text())
		e.MarkDirty(txBox_result)
		rem := 1024 - len(txBox_ascii.Text())
		length.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_ascii)
	row.Add(length)
	p.Add(row)

	p.AddVSpace(10)

	lb2 := gwu.NewListBox([]string{"1", "2", "4", "8", "16", "32", "64", "128"})
	lb2.SetMulti(true)
	lb2.SetRows(10)
	lb2.AddEHandlerFunc(func(e gwu.Event) {
		sum := 0
		for _, idx := range lb2.SelectedIndices() {
			sum += 1 << uint(idx)
		}
		if sum == 89 {
			txBox_result.SetText("Hooray! You did it!")
		} else {
			txBox_result.SetText(fmt.Sprintf("Now quite there... (sum = %d)", sum))
		}
		e.MarkDirty(txBox_result)
	}, gwu.ETYPE_CHANGE)
	p.Add(lb2)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("The Result:"))
	row = gwu.NewHorizontalPanel()

	row.Add(txBox_result)
	p.Add(row)

	return p
}

func HexToBytes(str string) ([]byte, error) {
	bytes := make([]byte, len(str)/2)
	var highBits byte
	var lowBits byte
	for i := 0; i < len(str); i += 2 {
		highBits = 0x00
		lowBits = 0x00
		switch {
		case str[i] >= '0' && str[i] <= '9':
			highBits = str[i] - '0'
		case str[i] >= 'a' && str[i] <= 'z':
			highBits = str[i] - 'a' + 10
		case str[i] >= 'A' && str[i] <= 'Z':
			highBits = str[i] - 'A' + 10
		default:
			return nil, errors.New(fmt.Sprintf("invalid hex character: %c", str[i]))
		}
		switch {
		case str[i+1] >= '0' && str[i] <= '9':
			lowBits = str[i+1] - '0'
		case str[i+1] >= 'a' && str[i] <= 'z':
			lowBits = str[i+1] - 'a' + 10
		case str[i+1] >= 'A' && str[i] <= 'Z':
			lowBits = str[i+1] - 'A' + 10
		default:
			return nil, errors.New(fmt.Sprintf("invalid hex character: %c", str[i]))
		}
		bytes[i/2] = highBits<<4 | lowBits

	}
	return bytes, nil
}

func XorSum(mybytes []byte) string {
	var result byte
	result = 0x00
	for _, one := range mybytes {
		result ^= one
	}
	return string(result)
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

	txBox_hex := gwu.NewTextBox("")
	txBox_hex.SetRows(8)
	txBox_hex.SetCols(128)
	txBox_hex.SetMaxLength(1024)
	txBox_hex.AddSyncOnETypes(gwu.ETYPE_KEY_UP)

	length := gwu.NewLabel("")
	length.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_ascii.AddEHandlerFunc(func(e gwu.Event) {
		rem := 1024 - len(txBox_ascii.Text())
		length.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
		if rem%2 == 0 {
			strH, _ := HexToBytes(txBox_ascii.Text())
			txBox_hex.SetText(string(strH))
			e.MarkDirty(txBox_hex)
			length.Style().SetBackground(gwu.CLR_GREEN)
		} else {
			length.Style().SetBackground(gwu.CLR_RED)
		}
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_ascii)
	row.Add(length)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Input Hex(max 1024 characters):"))
	row = gwu.NewHorizontalPanel()

	length_3 := gwu.NewLabel("")
	length_3.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_hex.AddEHandlerFunc(func(e gwu.Event) {
		rem := 1024 - len(txBox_hex.Text())
		length_3.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length_3)
		txBox_ascii.SetText(fmt.Sprintf("%X", txBox_hex.Text()))
		e.MarkDirty(txBox_ascii)
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

	txBox_xor_result := gwu.NewTextBox("")
	txBox_xor_result.SetCols(128)

	length_4 := gwu.NewLabel("")
	length_4.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_xor_data.AddEHandlerFunc(func(e gwu.Event) {
		rem := 1024 - len(txBox_xor_data.Text())
		length_4.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length_4)
		if rem%2 == 0 {
			strH, _ := HexToBytes(txBox_xor_data.Text())
			txBox_xor_result.SetText(fmt.Sprintf("%X\n", XorSum(strH)))
			e.MarkDirty(txBox_xor_result)
			length_4.Style().SetBackground(gwu.CLR_GREEN)
		} else {
			length_4.Style().SetBackground(gwu.CLR_RED)
		}
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(txBox_xor_data)
	row.Add(length_4)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("XorSum Result:"))

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
	length.Style().SetColor(gwu.CLR_WHITE)
	length.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_des_key.AddEHandlerFunc(func(e gwu.Event) {
		rem := 32 - len(txBox_des_key.Text())
		if rem == 0 || rem == 16 {
			length.Style().SetBackground(gwu.CLR_GREEN)
		} else {
			length.Style().SetBackground(gwu.CLR_RED)
		}
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
	length_2.Style().SetColor(gwu.CLR_WHITE)
	length_2.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	txBox_des_data.AddEHandlerFunc(func(e gwu.Event) {
		rem := 32 - len(txBox_des_data.Text())
		if rem%16 == 0 {
			length_2.Style().SetBackground(gwu.CLR_GREEN)
		} else {
			length_2.Style().SetBackground(gwu.CLR_RED)
		}
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

	links.AddVSpace(5)
	l = gwu.NewLabel("IC Card")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD)
	links.Add(l)
	createDemo("PBOC3", buildPboc3View)

	links.AddVSpace(5)
	l = gwu.NewLabel("8583")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetDisplay(gwu.DISPLAY_BLOCK)
	links.Add(l)
	createDemo("Analysis", build8583AnalysisView)

	links.AddVSpace(5)
	l = gwu.NewLabel("TLV")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD)
	links.Add(l)
	createDemo("Analysis", buildTlvAnalysisView)

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
	// Create GUI server
	server := gwu.NewServer("dc_pboc_client", "127.0.0.1:5761")
	server.SetText("")

	server.AddSessCreatorName("web_view", "")
	server.AddSHandler(SessHandler{})

	// Start GUI server
	if err := server.Start("web_view"); err != nil {
		fmt.Println("Error: Cound not start GUI server:", err)
		return
	}
}
