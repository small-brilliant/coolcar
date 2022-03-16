package profile

import (
	"context"
	"coolcar/shared/id"
)

type Manager struct {
}

func (p *Manager) Verify(context.Context, id.AccountID) (id.IdentiTyID, error) {
	return id.IdentiTyID("account1"), nil
}
