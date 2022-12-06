package rdgo

import (
	"strconv"

	"github.com/leekchan/accounting"
	gopdf "github.com/ofio/gopdf"
)

func CreateNewOrderPage(pageNum int, image []byte, pdf *gopdf.Fpdf, logob []byte, mleft, mtop, lineHeight, unitPriceWidth, totalPriceWidth, quantityWidth, sumWidth float64, cols []float64, lineItems [][]string, po PoHeader, invoice Invoice, isInvoice bool) {
	yLocLeft := mtop

	businessName := ""
	if isInvoice {
		businessName = invoice.Business.Name
		createTextBox(pdf, mleft, mtop, pdf.GetStringWidth(businessName)+6, lineHeight*2, businessName, "L", false, 16, "Medium")
		yLocLeft += (lineHeight * 3)
	} else {
		if len(logob) <= 0 {
			businessName = po.Instance.Business.Name
			createTextBox(pdf, mleft, mtop, pdf.GetStringWidth(businessName)+6, lineHeight*2, businessName, "L", false, 16, "Medium")
			yLocLeft += (lineHeight * 3)
		} else {
			createLogo(pdf, image, mleft, mtop, "logo_pg"+strconv.Itoa(pageNum)+".png")
			yLocLeft += (lineHeight * 4)
		}
	}

	//move below business name
	pdf.SetXY(mleft, yLocLeft)

	var firstColumnWidth float64 = sumWidth/2 - 3
	var secondColumnWidth float64 = sumWidth/2 - 3
	var secondColumnXLoc float64 = mleft + firstColumnWidth + 6

	contactColumnWidths := []float64{firstColumnWidth/2 - 24, firstColumnWidth/2 + 24}

	requesterExists := len(po.BuyerJsonb.Email) > 0
	if requesterExists {
		requesterItems := [][]string{
			{"Requester"},
			{"Name", po.RequesterJsonb.Name},
			{"Email", po.RequesterJsonb.Email},
		}
		yLocLeft += (lineHeight * 3)

		if len(po.Department.Name) > 0 {
			requesterItems = append(requesterItems, []string{"Dept", po.Department.Name})
			yLocLeft += (lineHeight)
		}
		if len(po.SoldToEntity) > 0 {
			requesterItems = append(requesterItems, []string{"Entity", po.SoldToEntity})
			yLocLeft += (lineHeight)
		}
		createContacts(requesterItems, contactColumnWidths, lineHeight, pdf, firstColumnWidth)
	}

	buyerExists := len(po.BuyerJsonb.Email) > 0
	if buyerExists {
		buyerItems := [][]string{
			{"Buyer"},
			{"Name", po.BuyerJsonb.Name},
			{"Email", po.BuyerJsonb.Email},
		}
		yLocLeft += (lineHeight * 3)
		createContacts(buyerItems, contactColumnWidths, lineHeight, pdf, firstColumnWidth)
	}
	vendorItems := [][]string{}
	billToItems := [][]string{}
	if isInvoice {
		vendorItems = [][]string{
			{"Sold By"},
			{"Name", invoice.Business.Name},
			{"Address", invoice.Business.Address},
			{"", invoice.Business.City + ", " + invoice.Business.StateProvince + " " + invoice.Business.PostalCode},
			{"", invoice.Business.Country},
		}
		yLocLeft += (lineHeight * 5)
		if len(invoice.Business.Phone) > 0 {
			vendorItems = append(vendorItems, []string{"Phone", invoice.Business.Phone})
			yLocLeft += (lineHeight * 1)
		}
		if len(invoice.UserEmail) > 0 {
			vendorItems = append(vendorItems, []string{"Email", invoice.UserEmail})
			yLocLeft += (lineHeight * 1)
		}

		billToItems = [][]string{
			{"Bill To"},
			{"Name", invoice.Instance.Business.Name},
			{"Address", invoice.Instance.Business.Address},
			{"", invoice.Instance.Business.City + ", " + invoice.Instance.Business.StateProvince + " " + invoice.Instance.Business.PostalCode},
			{"", invoice.Instance.Business.Country},
		}
		yLocLeft += (lineHeight * 5)
		if len(invoice.Instance.Business.Phone) > 0 {
			billToItems = append(billToItems, []string{"Phone", invoice.Instance.Business.Phone})
			yLocLeft += (lineHeight * 1)
		}
	} else {
		vendorItems = [][]string{
			{"Supplier"},
			{"Name", po.BusinessSupplier.Name},
			{"Address", po.BusinessSupplier.Address},
			{"", po.BusinessSupplier.City + ", " + po.BusinessSupplier.StateProvince + " " + po.BusinessSupplier.PostalCode},
			{"", po.BusinessSupplier.Country},
		}
		yLocLeft += (lineHeight * 5)
		if len(po.SupplierContact.Name) > 0 {
			vendorItems = append(vendorItems, []string{"Contact", po.SupplierContact.Name})
			yLocLeft += (lineHeight * 1)
		}
		if len(po.SupplierContact.Email) > 0 {
			vendorItems = append(vendorItems, []string{"Email", po.SupplierContact.Email})
			yLocLeft += (lineHeight * 1)
		}
		if len(po.SupplierContact.Phone) > 0 {
			vendorItems = append(vendorItems, []string{"Phone", po.SupplierContact.Phone})
			yLocLeft += (lineHeight)
		}
	}

	createContacts(vendorItems, contactColumnWidths, lineHeight, pdf, firstColumnWidth)
	if isInvoice {
		createContacts(billToItems, contactColumnWidths, lineHeight, pdf, firstColumnWidth)
	}

	//ORDER HEADER TEXT
	pdf.SetXY(secondColumnXLoc, mtop)
	//"order" header empty box
	createTextBox(pdf, secondColumnXLoc, mtop, secondColumnWidth, lineHeight*2, "", "C", true, 18, "Book")
	//"order" header text
	if isInvoice {
		createTextBox(pdf, secondColumnXLoc, mtop, secondColumnWidth, (lineHeight*2)+1, "INVOICE", "C", false, 18, "Book")
	} else {
		createTextBox(pdf, secondColumnXLoc, mtop, secondColumnWidth, (lineHeight*2)+1, "PURCHASE ORDER", "C", false, 18, "Book")
	}

	createOrderRevision := func(rows [][]string, cols []float64, lineHeight float64) {
		for _, row := range rows {
			curx, y := pdf.GetXY()
			x := curx

			for j, txt := range row {
				width := cols[j]

				if j%2 == 0 {
					pdf.SetFont("GothamHTF", "Medium", 10)
					pdf.CellFormat(width, lineHeight, txt, "", 0, "L", false, 0, "")
				} else {
					pdf.SetFont("GothamHTF", "Book", 10)
					pdf.CellFormat(width, lineHeight, txt, "", 0, "R", false, 0, "")
				}

				// pdf.MultiCell(width, lineHeight, txt, "", "", false)
				x += width
				pdf.SetXY(x, y)
			}
			pdf.SetXY(curx, y+lineHeight)
		}
	}
	revWidth := (secondColumnWidth) / 4
	orderRevCols := []float64{revWidth, revWidth, revWidth, revWidth}
	orderRevisionDetails := [][]string{}
	if isInvoice {
		//orderRevisionDetails = append(orderRevisionDetails, []string{"Order", invoice.InvoiceNumber, "", ""})
	} else {
		orderRevisionDetails = append(orderRevisionDetails, []string{"Order", po.PoNumber, "Revision", strconv.Itoa(po.Rev)})
	}

	pdf.SetXY(secondColumnXLoc, mtop+(lineHeight*2))
	createOrderRevision(orderRevisionDetails, orderRevCols, lineHeight)

	//		{"Order", strconv.Itoa(po.PONumber),"Revision", strconv.Itoa(po.Rev)},
	poCols := []float64{revWidth * 2, revWidth * 2}

	poHeaderItems := [][]string{}
	if isInvoice {
		poHeaderItems = [][]string{
			{"Invoice Number", invoice.InvoiceNumber},
			{"Purchase Order", invoice.PoNumber},
			{"Date", invoice.UpdatedAt.Format("January 2, 2006")},
		}
	} else {
		poHeaderItems = [][]string{
			{"Date", po.UpdatedAt.Format("January 2, 2006")},
			{"Payment Terms", po.PaymentTerms},
		}
	}

	createPOHeaderItems(pdf, poHeaderItems, poCols, lineHeight, mleft, mtop, sumWidth, secondColumnXLoc)

	if len(po.BusinessShipTo.Name) > 0 {
		yLocRight := pdf.GetY()
		shipTo := []float64{secondColumnWidth}
		shipToItems := [][]string{
			{"Ship To"},
			{po.BusinessShipTo.Name},
			{po.BusinessShipTo.ShippingAddress},
			{po.BusinessShipTo.ShippingCity + ", " + po.BusinessShipTo.ShippingStateProvince + " " + po.BusinessShipTo.ShippingPostalCode},
			{po.BusinessShipTo.ShippingCountry},
		}

		pdf.SetXY(secondColumnXLoc, yLocRight)
		_ = createAddressHeader(pdf, shipToItems, shipTo, lineHeight, yLocRight)
		yLocRight = yLocRight + (lineHeight * 5)
	}

	if len(po.BusinessBillTo.Name) > 0 {
		yLocRight := pdf.GetY()
		billTo := []float64{secondColumnWidth}
		billToItems := [][]string{
			{"Invoice To"},
			{po.BusinessBillTo.Name},
			{po.BusinessBillTo.Address},
			{po.BusinessBillTo.City + ", " + po.BusinessBillTo.StateProvince + " " + po.BusinessBillTo.PostalCode},
			{po.BusinessBillTo.Country},
		}

		pdf.SetXY(secondColumnXLoc, yLocRight)
		_ = createAddressHeader(pdf, billToItems, billTo, lineHeight, yLocRight)
		yLocRight = yLocRight + (lineHeight * 5)

	}

	if len(po.InvoicingInstructions) > 0 {
		yLocRight := pdf.GetY()
		pdf.SetXY(secondColumnXLoc, yLocRight)
		pdf.SetFillColor(240, 240, 240)
		pdf.SetFont("GothamHTF", "Medium", 10)

		pdf.CellFormat(secondColumnWidth, lineHeight, "Invoicing Instructions", "", 0, "C", true, 0, "")
		yLocRight += lineHeight
		pdf.SetFont("GothamHTF", "Book", 10)
		pdf.SetXY(secondColumnXLoc, yLocRight)
		lines := pdf.SplitLines([]byte(po.InvoicingInstructions), secondColumnWidth)
		for lineNum, line := range lines {
			y := pdf.GetY()
			if y > 268 {
				pdf.AddPage()
				yLocLeft = mtop
				yLocRight = mtop
			}

			if lineNum != 0 {
				yLocRight += (lineHeight)
				pdf.SetXY(secondColumnXLoc, yLocRight)
			}
			pdf.CellFormat(sumWidth, lineHeight, string(line), "", 1, "L", false, 0, "")
		}
	}

	if len(po.Contract.Name) > 0 {
		yLocRight := pdf.GetY()
		pdf.SetXY(secondColumnXLoc, yLocRight)
		pdf.SetFillColor(240, 240, 240)
		pdf.SetFont("GothamHTF", "Medium", 10)
		pdf.CellFormat((revWidth * 4), lineHeight, "Contract Information", "", 0, "C", true, 0, "")
		yLocRight += lineHeight
		poCols := []float64{(revWidth * 2), (revWidth * 2)}
		//	endDate := "none"
		signedDate := "none"

		// if !po.Contract.EndDate.IsZero() {
		// 	endDate = po.Contract.EndDate.Format("January 2, 2006")
		// }

		if !po.Contract.SignedDate.IsZero() {
			signedDate = po.Contract.SignedDate.Format("January 2, 2006")
		}

		lines := pdf.SplitLines([]byte(po.Contract.Name), sumWidth)
		contractName := ""
		if len(lines) > 0 {
			contractName = string(lines[0])
		}
		poHeaderItems := [][]string{
			{"Name", contractName},
			{"Date Executed", signedDate},
			//{"End Date", endDate},
			// {"Renewal Type", po.Contract.RenewalType},
			{"Payment Schedule", po.Contract.PaymentSchedule},
		}

		//		{"Order", strconv.Itoa(po.PONumber),"Revision", strconv.Itoa(po.Rev)},
		pdf.SetXY(secondColumnXLoc, yLocRight)
		createPOHeaderItems(pdf, poHeaderItems, poCols, lineHeight, mleft, mtop, sumWidth, secondColumnXLoc)

	}
	pdf.SetX(mleft)

	//FULLWIDTH SECTION
	pageCount := pdf.PageCount()
	notesYLoc := pdf.GetY()
	if pageCount > 1 {
		notesYLoc = pdf.GetY()
	} else {
		if yLocLeft > pdf.GetY() {
			notesYLoc = yLocLeft
		}
	}
	notesYLoc += lineHeight

	if len(po.Notes) > 0 {
		addFullWidthText("Notes", po.Notes, pdf, lineHeight, notesYLoc, sumWidth, 18, mleft, mtop)
	} else {
		pdf.SetY(notesYLoc)
	}
	createLineItem(pdf, lineItems, cols, lineHeight, mtop, mleft)

	y0 := pdf.GetY()
	total := 0.0
	if isInvoice {
		for _, v := range invoice.InvoiceLines {
			total += v.LineAmount
		}
	} else {
		for _, v := range po.PoLines {
			total += v.LineAmount
		}
	}

	lc := accounting.LocaleInfo[po.CurrencyCode]
	ac := accounting.Accounting{Symbol: po.CurrencyCode, Precision: 2, Thousand: lc.ThouSep, Decimal: lc.DecSep}

	totalColWidths := []float64{unitPriceWidth, totalPriceWidth}
	totalItems := [][]string{
		{"Subtotal", ac.FormatMoney(total)},
		//{"Tax", "540.00"},
		//{"Shipping"},
		//{"Other", ""},
		{"Total", ac.FormatMoney(total)},
	}

	pdf.SetXY(mleft+sumWidth-unitPriceWidth-totalPriceWidth, y0+10)
	createTotalItems(pdf, totalItems, totalColWidths, lineHeight)

	if len(po.TermsAndConditions) > 0 {
		addFullWidthText("Terms and Conditions", po.TermsAndConditions, pdf, lineHeight, pdf.GetY()+lineHeight, sumWidth, 47, mleft, mtop)
	}
}
