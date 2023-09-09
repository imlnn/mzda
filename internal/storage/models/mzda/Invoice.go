package mzda

type InvoiceStorage interface {
	AddInvoice()
	GetInvoice()
	DeleteInvoice()
	UpdateInvoice()
}
