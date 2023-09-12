package models

type InvoiceStorage interface {
	AddInvoice()
	GetInvoice()
	DeleteInvoice()
	UpdateInvoice()
}
