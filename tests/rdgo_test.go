package rdgo

import (
	"log"
	"testing"

	"github.com/ofio/rdgo"
)

func TestImportStruct(t *testing.T) {
	sample := &rdgo.ApprovalRequestApprover{Approver: rdgo.Approver{Email: "test@raindrop.com", Name: "test"}}
	log.Println(sample)
}
