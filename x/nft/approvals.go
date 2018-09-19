package nft

import "github.com/iov-one/weave/errors"

const (
	ActionKindUpdateDetails = "baseUpdateDetails"
)

func (i *ActionApprovals) Validate() error {
	if i == nil {
		return errors.ErrInternal("must not be nil")
	}
	// todo: do proper validation
	return nil
}

func (i *ActionApprovals) Clone() *ActionApprovals {
	if i == nil {
		return nil
	}
	approvals := make([]*Approval, len(i.Approvals))
	for i, a := range i.Approvals {
		approvals[i] = a.Clone()
	}
	return &ActionApprovals{
		Action:    i.Action,
		Approvals: approvals,
	}
}

func (i *Approval) Clone() *Approval {
	if i == nil {
		return nil
	}
	return &Approval{
		Address: i.Address,
		Options: i.Options.Clone(),
	}
}

func (i *ApprovalOptions) Clone() *ApprovalOptions {
	if i == nil {
		return nil
	}
	return &ApprovalOptions{
		UntilBlockHeight: i.UntilBlockHeight,
		Count:            i.Count,
		Immutable:        i.Immutable,
	}
}
