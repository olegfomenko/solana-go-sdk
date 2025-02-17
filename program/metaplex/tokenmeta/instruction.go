package tokenmeta

import (
	"github.com/near/borsh-go"
	"github.com/olegfomenko/solana-go-sdk/common"
	"github.com/olegfomenko/solana-go-sdk/types"
	"github.com/pkg/errors"
)

type Instruction uint8

const (
	InstructionCreateMetadataAccount Instruction = iota
	InstructionUpdateMetadataAccount
	InstructionDeprecatedCreateMasterEdition
	InstructionDeprecatedMintNewEditionFromMasterEditionViaPrintingToken
	InstructionUpdatePrimarySaleHappenedViaToken
	InstructionDeprecatedSetReservationList
	InstructionDeprecatedCreateReservationList
	InstructionSignMetadata
	InstructionDeprecatedMintPrintingTokensViaToken
	InstructionDeprecatedMintPrintingTokens
	InstructionCreateMasterEdition
	InstructionMintNewEditionFromMasterEditionViaToken
	InstructionConvertMasterEditionV1ToV2
	InstructionMintNewEditionFromMasterEditionViaVaultProxy
	InstructionPuffMetadata
)

func CreateMetadataAccount(metadata, mint, mintAuthority, payer, updateAuthority common.PublicKey, updateAuthorityIsSigner, isMutable bool, mintData Data) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Data        Data
		IsMutable   bool
	}{
		Instruction: InstructionCreateMetadataAccount,
		Data:        mintData,
		IsMutable:   isMutable,
	})

	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     mintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     updateAuthority,
				IsSigner:   updateAuthorityIsSigner,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}, nil
}

func UpdatePrimarySaleHappenedViaToken(metadata, owner, account common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionUpdatePrimarySaleHappenedViaToken,
	})

	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     owner,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     account,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}, nil
}

func UpdateMetadataAccount(metadata, owner common.PublicKey, updateData *Data, updateAuthority *common.PublicKey, updatePrimarySaleHappened *bool) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction         Instruction
		Data                *Data
		UpdateAuthority     *common.PublicKey
		PrimarySaleHappened *bool
	}{
		Instruction:         InstructionUpdateMetadataAccount,
		Data:                updateData,
		UpdateAuthority:     updateAuthority,
		PrimarySaleHappened: updatePrimarySaleHappened,
	})

	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     owner,
				IsSigner:   true,
				IsWritable: false,
			},
		},
		Data: data,
	}, nil
}

func SignMetadata(metadata, creator common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionSignMetadata,
	})

	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     creator,
				IsSigner:   true,
				IsWritable: false,
			},
		},
		Data: data,
	}, nil
}
