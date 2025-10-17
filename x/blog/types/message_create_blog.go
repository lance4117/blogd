package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateBlog{}

func NewMsgCreateBlog(creator string, title string, content string) *MsgCreateBlog {
	return &MsgCreateBlog{
		Creator: creator,
		Title:   title,
		Content: content,
	}
}

func (msg *MsgCreateBlog) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
