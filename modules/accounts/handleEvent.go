package accounts

import (
	"github.com/eapache/go-resiliency/breaker"
	"github.com/forbole/flowJuno/types"
)
func HandleEvent(event types.Event, m.messagesParser, m.db, tx.Height){
	switch event.Type{
	case "flow.AccountKeyAdded":
		
		break
	case "flow.AccountCreated":
		break
	case "flow.AccountKeyRemoved":
		break
	case "flow.AccountCodeUpdated":
		break
	}
}