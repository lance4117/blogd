package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMsgUpdateBlog{}

func NewMsgMsgUpdateBlog(creator string, id int64, title string, content string) *MsgMsgUpdateBlog {
	return &MsgMsgUpdateBlog{
		Creator: creator,
		Id:      id,
		Title:   title,
		Content: content,
	}
}

func (msg *MsgMsgUpdateBlog) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
