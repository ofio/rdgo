package rdgo

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/leekchan/accounting"
	gopdf "github.com/ofio/gopdf"
	uuid "github.com/satori/go.uuid"
)

func createAddressHeader(pdf *gopdf.Fpdf, rows [][]string, cols []float64, lineHeight float64, yLocRight float64) float64 {

	for r, row := range rows {
		curx, y := pdf.GetXY()
		x := curx

		height := 0.

		for i, txt := range row {
			lines := pdf.SplitLines([]byte(txt), cols[i])
			h := float64(len(lines)) * lineHeight
			if h > height {
				height = h
			}
		}

		for j, txt := range row {
			width := cols[j]
			pdf.SetFont("GothamHTF", "Book", 10)
			if r == 0 {
				pdf.SetFillColor(240, 240, 240)
				pdf.SetFont("GothamHTF", "Medium", 10)
				pdf.CellFormat(width, lineHeight, txt, "", 0, "C", true, 0, "")
			} else {
				pdf.SetFont("GothamHTF", "Book", 10)
				pdf.CellFormat(width, lineHeight, txt, "", 0, "L", false, 0, "")
			}

			// pdf.MultiCell(width, lineHeight, txt, "", "", false)
			x += width
			pdf.SetXY(x, y)
		}

		pdf.SetXY(curx, y+height)
	}
	return yLocRight
}

func createTotalItems(pdf *gopdf.Fpdf, rows [][]string, cols []float64, lineHeight float64) {
	for r, row := range rows {
		curx, y := pdf.GetXY()
		x := curx

		height := 0.

		for i, txt := range row {
			lines := pdf.SplitLines([]byte(txt), cols[i])
			h := float64(len(lines)) * lineHeight
			if h > height {
				height = h
			}
		}

		for j, txt := range row {
			width := cols[j]

			fill := false
			if r == len(rows)-1 {
				fill = true
			}
			if j == 0 {
				pdf.SetFillColor(240, 240, 240)
				pdf.SetFont("GothamHTF", "Medium", 10)

				pdf.CellFormat(width, lineHeight, txt, "", 0, "L", fill, 0, "")
			} else {
				pdf.SetFont("GothamHTF", "Book", 10)
				pdf.CellFormat(width, lineHeight, txt, "", 0, "R", fill, 0, "")
			}

			// pdf.MultiCell(width, lineHeight, txt, "", "", false)
			x += width
			pdf.SetXY(x, y)
		}
		pdf.SetXY(curx, y+height)
	}
}

func createLineItem(pdf *gopdf.Fpdf, rows [][]string, cols []float64, lineHeight float64, mtop float64, mleft float64) {
	fill := true
	for r, row := range rows {
		curx, y := pdf.GetXY()

		x := curx

		description := row[1]
		splitLinesDescription := pdf.SplitLines([]byte(description), cols[1])
		numDescriptionRows := len(splitLinesDescription)

		commodity := row[2]
		splitLinesCommodity := pdf.SplitLines([]byte(commodity), cols[2])
		numCommodityRows := len(splitLinesCommodity)

		numRows := 1
		if numCommodityRows > 1 || numDescriptionRows > 1 {
			if numCommodityRows > numDescriptionRows {
				numRows = numCommodityRows
			} else {
				numRows = numDescriptionRows
			}
		}

		// _, ht, _ := pdf.PageSize(1)
		if y > 268 {
			pdf.AddPage()
			y = mtop
			fill := true

			//insert row header
			for j, txt := range rows[0] {
				width := cols[j]
				pdf.SetFont("GothamHTF", "Medium", 10)
				if j == 0 {
					pdf.CellFormat(width, lineHeight, txt, "", 0, "L", fill, 0, "")
				} else if j == 1 {
					pdf.CellFormat(width, lineHeight, txt, "", 0, "L", fill, 0, "")
				} else if j == 2 {
					pdf.CellFormat(width, lineHeight, txt, "", 0, "L", fill, 0, "")
				} else if j == 3 {
					pdf.CellFormat(width, lineHeight, txt, "", 0, "R", fill, 0, "")
				} else if j == 4 {
					pdf.CellFormat(width, lineHeight, txt, "", 0, "R", fill, 0, "")
				} else if j == 5 {
					pdf.CellFormat(width, lineHeight, txt, "", 0, "R", fill, 0, "")
				}

				x += width
				pdf.SetXY(x, y)
			}
			x = mleft
			y = y + lineHeight

			pdf.SetXY(x, y)
		}

		rowHeight := float64(numRows) * lineHeight
		if rowHeight <= lineHeight {
			rowHeight = lineHeight
		}

		for idx := 0; idx < numRows; idx++ {
			if idx != 0 {
				x = mleft
				y += lineHeight
				if y > 268 {
					fill = true
					pdf.AddPage()
					y = mtop
					pdf.SetXY(x, y)
				} else {
					pdf.SetXY(x, y)
				}
			}
			for j, txt := range row {
				width := cols[j]

				pdf.SetFillColor(240, 240, 240)

				if r == 0 {
					fill = true
					pdf.SetFont("GothamHTF", "Medium", 10)
				}
				if fill {
					pdf.SetFillColor(244, 244, 244)
					pdf.SetFont("GothamHTF", "Book", 10)
					pdf.SetDrawColor(255, 255, 255)
					pdf.SetLineWidth(0)
					pdf.Rect(x, y, width, rowHeight, "")
				} else {
					pdf.SetFont("GothamHTF", "Book", 10)
					pdf.SetDrawColor(255, 255, 255)
					pdf.SetLineWidth(0)
					pdf.Rect(x, y, width, rowHeight, "")
				}

				align := "L"
				if j > 2 {
					align = "R"
				}
				text := ""
				if idx == 0 {
					text = txt
				}

				if j == 1 {
					if idx < len(splitLinesDescription) {
						text = string(splitLinesDescription[idx])
					}
				} else if j == 2 {
					if idx < len(splitLinesCommodity) {
						text = string(splitLinesCommodity[idx])
					}
				}

				pdf.CellFormat(width, lineHeight, text, "", 0, align, fill, 0, "")
				x += width

				pdf.SetXY(x, y)
			}
		}

		if fill {
			fill = false
		} else {
			fill = true
		}

		pdf.SetXY(curx, y+lineHeight)
	}
}

func createTextBox(pdf *gopdf.Fpdf, x float64, y float64, w float64, h float64, businessName string, align string, fill bool, fontSize float64, fontStyle string) {
	// Arial bold 15
	pdf.SetFont("GothamHTF", fontStyle, fontSize)
	// Calculate width of title and position
	pdf.SetXY(x, y)
	// Title
	pdf.CellFormat(w, h, businessName, "", 1, align, fill, 0, "")
}

func queryPurchaseOrder(instance int, poNumber string, revision int, token string, xHasuraAdminSecret string, hasuraEndpoint string) (PoHeader, error) {
	queryPO := `query purchaseOrder($rev: Int, $instance_id: Int, $po_number: String) {
		po_header(where: {_and: {}, rev: {_eq: $rev}, instance_id: {_eq: $instance_id}, po_number: {_eq: $po_number}}) {
			id
			uuid
			currency_code
			created_by
			po_number
			payment_terms
			status
			invoicing_instructions
			terms_and_conditions
			notes
			department_id
			sold_to_entity
			instance {
				business {
					name
				}
				id
				instance_settings {
					branding_logo_uuid
				}
			}
			department_id
			rev
			buyer_jsonb
			requester_jsonb
			supplier_contact {
				name
				email
			}
			updated_at
			businessBillTo {
				name
				address
				city
				state_province
				postal_code
				country
			}
			businessShipTo {
				name
				shipping_address
				shipping_city
				shipping_state_province
				shipping_postal_code
				shipping_country
			}
			businessSupplier {
				name
				address
				city
				state_province
				postal_code
				country
			}
			department {
				id
				name
			}
			po_lines(order_by: {line_number: asc}) {
				id
				commodity {
					id
					name
				}
				line_number
				item_description
				quantity
				net_price_per_unit
				commodity_id
			}
			contract {
				name
				effective_date
				end_date
				renewal_type
				payment_schedule
				signed_date
			}
		}
	}	
	`

	purchaseOrder := PoHeader{}
	queryVar := map[string]interface{}{"rev": revision, "instance_id": instance, "po_number": poNumber}
	smartResponseData := Responsedata{}
	err := NewError()
	if len(xHasuraAdminSecret) > 0 {
		smartResponseData, err = SmartQuery(queryPO, queryVar, hasuraEndpoint, xHasuraAdminSecret, "")
		if err != nil {
			log.Println("query error", err)
			return purchaseOrder, err
		}
	} else {
		smartResponseData, err = SmartQuery(queryPO, queryVar, hasuraEndpoint, "", token)
		if err != nil {
			log.Println("query error", err)
			return purchaseOrder, err
		}

	}

	for _, poh := range smartResponseData.Data.PoHeader {
		purchaseOrder = poh
	}
	return purchaseOrder, nil
}

func queryInvoice(invoiceID int, revision int, token string, xHasuraAdminSecret string, HasuraEndpoint string) (Invoice, error) {
	queryPO := `query invoice($id: Int!) {
		invoice(where: {id: {_eq: $id}}) {
			id
			created_by
			currency_code
			business_id
			po_number
			type
			po_number
			amount
			approved_at
			updated_at
			bank_name
			checks_payable_to
			business_phone
			routing_number
			account_number
			import_status
			import_data
			status
			invoice_number
			terms_and_conditions
			instance {
				business {
					name
				}
				id
				instance_settings {
					branding_logo_uuid
				}
			}
			updated_at
			business {
				id
				address
				name
				city
				country
				state_province
				postal_code
				phone
			}
			invoice_lines(order_by: {line_number: asc}) {
				id
				line_number
				item_description
				quantity
				uom_code
				item_code
			}
		}
	}	
	`

	invoice := Invoice{}
	queryVar := map[string]interface{}{"id": invoiceID}
	smartResponseData := Responsedata{}
	err := NewError()
	if len(xHasuraAdminSecret) > 0 {
		smartResponseData, err = SmartQuery(queryPO, queryVar, HasuraEndpoint, xHasuraAdminSecret, "")
		if err != nil {
			log.Println("query error", err)
			return invoice, err
		}
	} else {
		smartResponseData, err = SmartQuery(queryPO, queryVar, HasuraEndpoint, "", token)
		if err != nil {
			log.Println("query error", err)
			return invoice, err
		}

	}

	for _, poh := range smartResponseData.Data.Invoice {
		invoice = poh
	}
	return invoice, nil
}

func InvoicePurchaseOrderHandler(pdf *gopdf.Fpdf, instance int, poNumber string, invoiceID int, revision int, token string, adminSecret string, hasuraEndpoint string, isInvoice bool, bucket string, publicBucket string, saveAttachment bool) ([]byte, string, error) {
	poHeader := PoHeader{}
	invoice := Invoice{}

	fileName := ""
	brandingLogoUUID := ""
	err := NewError()
	objectID := -1
	createdBy := ""
	if isInvoice {
		invoice, err = queryInvoice(invoiceID, revision, token, adminSecret, hasuraEndpoint)
		if err != nil {
			fmt.Println(err)
			return nil, "", err
		}
		for _, settings := range invoice.Instance.InstanceSettings {
			brandingLogoUUID = settings.BrandingLogoUUID
		}
		fileName = "Invoice " + invoice.InvoiceNumber + " " + invoice.BusinessName + ".pdf"
		objectID = invoice.ID
		createdBy = invoice.CreatedBy
	} else {
		poHeader, err = queryPurchaseOrder(instance, poNumber, revision, token, adminSecret, hasuraEndpoint)
		if err != nil {
			fmt.Println(err)
			return nil, "", err
		}
		for _, settings := range poHeader.Instance.InstanceSettings {
			brandingLogoUUID = settings.BrandingLogoUUID
		}
		fileName = "PO " + poNumber + " Rev " + strconv.Itoa(revision) + " " + poHeader.BusinessSupplier.Name + ".pdf"
		objectID = poHeader.ID
		createdBy = poHeader.CreatedBy
	}

	var logob []byte
	if len(brandingLogoUUID) > 0 {
		_, logob, err = ReadObj(brandingLogoUUID, strconv.Itoa(instance), publicBucket)
		if err != nil {
			fmt.Println(err)
		}
	}

	var pdfb []byte
	pdfb, err = CreatePurchaseOrderInvoice(pdf, poHeader, invoice, isInvoice, &logob)
	if err != nil {
		fmt.Println(err)
		return nil, "", err
	}

	if saveAttachment {
		err := savePDFAttachment(pdfb, objectID, createdBy, fileName, instance, isInvoice, hasuraEndpoint, adminSecret, token, bucket)
		if err != nil {
			fmt.Println(err)
			return nil, "", err
		}
	}

	return pdfb, fileName, nil
}

func savePDFAttachment(pdfb []byte, objectID int, createdBy string, fileName string, instance int, isInvoice bool, hasuraEndpoint string, adminSecret string, token string, bucket string) error {
	uuid := uuid.NewV4().String()
	if isInvoice {
		queryInvoiceAttachment := `query invoice_attachment($invoice_id: Int!) {
			invoice_attachment(where: { invoice_id: { _eq: $invoice_id }, is_deleted: { _eq: false } }) {
				uuid
				name
				generation
			}
		}
		`
		queryVar := map[string]interface{}{"invoice_id": objectID}
		smartResponseData, err := SmartQuery(queryInvoiceAttachment, queryVar, hasuraEndpoint, adminSecret, "")
		if err != nil {
			log.Println("query error", err)
			return err
		}

		existingAttachment := Attachment{}
		for _, attachment := range smartResponseData.Data.InvoiceAttachment {
			existingAttachment = attachment
		}

		if len(existingAttachment.UUID) > 0 {
			uuid = existingAttachment.UUID
		}

		id, uuid, gen, err := FileUpsert(bufio.NewReader(bytes.NewReader(pdfb)), instance, fileName, "application/pdf", createdBy, uuid, objectID, bucket, "invoice", hasuraEndpoint, adminSecret, "", instance, instance, "invoice_attachment_pkey")
		if err != nil {
			log.Println("upsert error", err)
			return err
		}
		log.Println("file upserted: ", id, uuid, gen)

	} else {
		queryInvoiceAttachment := `query po_header_attachment($po_header_id: Int!) {
			po_header_attachment(where: {po_header_id: {_eq: $po_header_id}, is_purchase_order: {_eq: true}, is_deleted: {_eq: false}}) {
				uuid
				name
				generation
			}
		}
		`
		queryVar := map[string]interface{}{"po_header_id": objectID}
		smartResponseData, err := SmartQuery(queryInvoiceAttachment, queryVar, hasuraEndpoint, adminSecret, "")
		if err != nil {
			log.Println("query error", err)
			return err
		}

		existingAttachment := Attachment{}
		for _, attachment := range smartResponseData.Data.InvoiceAttachment {
			existingAttachment = attachment
		}

		if len(existingAttachment.UUID) > 0 {
			uuid = existingAttachment.UUID
		}

		id, uuid, gen, err := FileUpsert(bufio.NewReader(bytes.NewReader(pdfb)), instance, fileName, "application/pdf", createdBy, uuid, objectID, bucket, "po_header", hasuraEndpoint, adminSecret, "", instance, instance, "po_header_attachment_pkey")
		if err != nil {
			log.Println("upsert error", err)
			return err
		}
		log.Println("file upserted: ", id, uuid, gen)
	}
	return nil
}
func createLogo(pdf *gopdf.Fpdf, bc []byte, mleft float64, mtop float64, imageName string) {
	opt := gopdf.ImageOptions{
		ImageType:             "image/png",
		ReadDpi:               false,
		AllowNegativePosition: false,
	}

	createBusinessLogo := func() {
		pdf.ImageOptions(bc, "image/png", imageName, mleft, mtop, 60, 0, false, opt, 0, "")
	}
	createBusinessLogo()
}

func CreatePurchaseOrderInvoice(pdf *gopdf.Fpdf, po PoHeader, invoice Invoice, isInvoice bool, logob *[]byte) ([]byte, error) {
	var err error

	lc := accounting.LocaleInfo[po.CurrencyCode]

	ac := accounting.Accounting{Symbol: po.CurrencyCode, Precision: 2, Thousand: lc.ThouSep, Decimal: lc.DecSep}

	pagew, pageh := pdf.GetPageSize()
	mleft, mtop, mright, _ := pdf.GetMargins()
	_ = mtop
	sumWidth := pagew - mleft - mright
	itemWidth := 15.
	quantityWidth := 20.
	unitPriceWidth := 30.
	totalPriceWidth := 30.
	commodityWidth := 40.
	descriptionWidth := sumWidth - commodityWidth - itemWidth - quantityWidth - unitPriceWidth - totalPriceWidth
	cols := []float64{itemWidth, descriptionWidth, commodityWidth, quantityWidth, unitPriceWidth, totalPriceWidth}

	lineHeight := 6.

	lineItems := [][]string{}

	if isInvoice {
		lineItems = append(lineItems, []string{"Item", "Description", "UOM", "Quantity", "Unit Price", "Total"})
		if len(invoice.InvoiceLines) == 0 {
			strArr := []string{"", "", "", "", "", "-"}
			lineItems = append(lineItems, strArr, strArr, strArr, strArr)
		} else {
			for _, v := range invoice.InvoiceLines {
				quantity := ""
				if _, frac := math.Modf(v.Quantity); frac < 1e-6 || frac > 1-1e-6 {
					quantity = fmt.Sprint(int(v.Quantity))
				} else {
					quantity = strconv.FormatFloat(v.Quantity, 'f', 3, 64)
				}
				strArr := []string{strconv.Itoa(v.LineNumber), v.ItemDescription, v.UomCode, quantity, ac.FormatMoney(v.UnitPrice), ac.FormatMoney(v.Quantity * v.UnitPrice)}
				lineItems = append(lineItems, strArr)
			}
		}
	} else {
		lineItems = append(lineItems, []string{"Item", "Description", "Commodity", "Quantity", "Unit Price", "Total"})
		if len(po.PoLines) == 0 {
			strArr := []string{"", "", "", "", "", "-"}
			lineItems = append(lineItems, strArr, strArr, strArr, strArr)
		} else {
			for _, v := range po.PoLines {
				quantity := ""
				if _, frac := math.Modf(v.Quantity); frac < 1e-6 || frac > 1-1e-6 {
					quantity = fmt.Sprint(int(v.Quantity))
				} else {
					quantity = strconv.FormatFloat(v.Quantity, 'f', 3, 64)
				}
				strArr := []string{strconv.Itoa(v.LineNumber), v.ItemDescription, v.Commodity.Name, quantity, ac.FormatMoney(v.NetPricePerUnit), ac.FormatMoney(v.Quantity * v.NetPricePerUnit)}
				lineItems = append(lineItems, strArr)
			}
		}
	}

	pdf.AddPage()

	bc := make([]byte, len(*logob))
	copy(bc, *logob)
	CreateNewOrderPage(0, bc, pdf, bc, mleft, mtop, lineHeight, unitPriceWidth, totalPriceWidth, quantityWidth, sumWidth, cols, lineItems, po, invoice, isInvoice)

	// headerX := pdf.GetX()

	pdf.SetXY(mleft, mtop+(lineHeight*6))

	for i := 1; i <= pdf.PageCount(); i++ {
		pdf.SetPage(i)
		pageStr := fmt.Sprintf("Page %d of "+strconv.Itoa(pdf.PageCount()), pdf.PageNo())
		wd := pdf.GetStringWidth(pageStr) + 6
		pageStrX := ((pagew - wd) / 2)
		// Arial italic 8
		pdf.SetFont("GothamHTF", "Medium", 14)
		// Text color in gray
		pdf.SetTextColor(128, 128, 128)
		// Page number
		pdf.Text(pageStrX, pageh-10, pageStr)

	}
	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		fmt.Println(err)
	}
	pdfb := buf.Bytes()
	return pdfb, err
}

func createPOHeaderItems(pdf *gopdf.Fpdf, rows [][]string, cols []float64, lineHeight float64, mleft float64, mtop float64, sumWidth float64, secondColumnXLoc float64) {
	_, y := pdf.GetXY()

	for _, row := range rows {
		if y > 268 {
			pdf.AddPage()
			pdf.SetXY(mleft, mtop)
		}

		for j, txt := range row {
			width := cols[j]

			if j == 0 {
				pdf.SetFont("GothamHTF", "Medium", 10)
				pdf.CellFormat(width, lineHeight, txt, "", 0, "L", false, 0, "")
			} else {
				pdf.SetFont("GothamHTF", "Book", 10)
				lines := pdf.SplitLines([]byte(txt), width)
				for idx, line := range lines {
					if y > 268 {
						pdf.AddPage()
						pdf.SetXY(mleft, mtop)
					}

					pdf.CellFormat(sumWidth, lineHeight, string(line), "", 0, "L", false, 0, "")
					if len(lines) > 1 && idx != len(lines)-1 {
						xPos := secondColumnXLoc + width
						y += (lineHeight)
						pdf.SetXY(xPos, y)
					}

				}
			}

		}

		y += (lineHeight)
		pdf.SetXY(secondColumnXLoc, y)

	}
}

func addFullWidthText(title string, text string, pdf *gopdf.Fpdf, lineHeight float64, yLoc float64, sumWidth float64, titleWidth float64, marginLeft float64, marginTop float64) {
	if yLoc > 268 {
		pdf.AddPage()
		yLoc = marginTop
		pdf.SetXY(marginLeft, yLoc)
	} else {
		pdf.SetXY(marginLeft, yLoc)
	}

	pdf.SetFillColor(240, 240, 240)
	pdf.SetFont("GothamHTF", "Medium", 10)

	pdf.CellFormat(titleWidth, lineHeight, title, "", 0, "L", false, 0, "")
	pdf.SetFont("GothamHTF", "Book", 10)
	//keep on same line
	pdf.SetXY(marginLeft+titleWidth, yLoc+lineHeight)
	lines := pdf.SplitLines([]byte(text), sumWidth)
	for _, line := range lines {
		y := pdf.GetY()
		if y > 268 {
			pdf.AddPage()
			yLoc = marginTop
			pdf.SetXY(marginLeft, yLoc)
		}
		yLoc = yLoc + (lineHeight)
		pdf.SetXY(marginLeft, yLoc)
		pdf.CellFormat(sumWidth, lineHeight, string(line), "", 1, "L", false, 0, "")
	}
	yLoc = yLoc + (lineHeight * 2)
	pdf.SetXY(marginLeft, yLoc)
}

func createContacts(rows [][]string, cols []float64, lineHeight float64, pdf *gopdf.Fpdf, firstColumn float64) {
	for r, row := range rows {
		curx, y := pdf.GetXY()
		x := curx

		for j, txt := range row {
			width := cols[j]

			fill := false
			if r == 0 {
				pdf.SetFillColor(240, 240, 240)
				pdf.SetFont("GothamHTF", "Medium", 10)
				pdf.CellFormat(firstColumn, lineHeight, txt, "", 0, "C", true, 0, "")
			} else {
				if j == 0 {
					pdf.SetFont("GothamHTF", "Medium", 10)
					pdf.CellFormat(width, lineHeight, txt, "", 0, "L", fill, 0, "")
				} else {
					pdf.SetFont("GothamHTF", "Book", 10)
					pdf.CellFormat(width, lineHeight, txt, "", 0, "L", fill, 0, "")
				}

			}

			// pdf.MultiCell(width, lineHeight, txt, "", "", false)
			x += width
			pdf.SetXY(x, y)
		}
		pdf.SetXY(curx, y+lineHeight)
	}
}
