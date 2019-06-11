package pm

import (
	"math"
	"math/big"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestAuxData(t *testing.T) {
	round := int64(5)
	blkHash := ethcommon.BytesToHash(ethcommon.FromHex("7624778dedc75f8b322b9fa1632a610d40b85e106c7d9bf0e743a9ce291b9c6f"))

	assert := assert.New(t)

	// Test CreationRound = 0

	ticket := &Ticket{
		CreationRound:          0,
		CreationRoundBlockHash: blkHash,
	}

	assert.Equal(
		ethcommon.FromHex("00000000000000000000000000000000000000000000000000000000000000007624778dedc75f8b322b9fa1632a610d40b85e106c7d9bf0e743a9ce291b9c6f"),
		ticket.AuxData(),
	)

	// Test empty block hash

	emptyBlkHash := ethcommon.BytesToHash(ethcommon.LeftPadBytes([]byte{}, 32))

	ticket = &Ticket{
		CreationRound:          round,
		CreationRoundBlockHash: emptyBlkHash,
	}

	assert.Equal(
		ethcommon.FromHex("00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000"),
		ticket.AuxData(),
	)

	// Test nil block hash

	ticket = &Ticket{
		CreationRound: round,
	}

	assert.Equal(
		ethcommon.FromHex("00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000"),
		ticket.AuxData(),
	)

	// Test round = 0 and empty block hash

	ticket = &Ticket{
		CreationRound:          0,
		CreationRoundBlockHash: emptyBlkHash,
	}

	assert.Equal(
		[]byte{},
		ticket.AuxData(),
	)

	// Test normal case

	ticket = &Ticket{
		CreationRound:          round,
		CreationRoundBlockHash: blkHash,
	}

	assert.Equal(
		ethcommon.FromHex("00000000000000000000000000000000000000000000000000000000000000057624778dedc75f8b322b9fa1632a610d40b85e106c7d9bf0e743a9ce291b9c6f"),
		ticket.AuxData(),
	)
}

func TestHash(t *testing.T) {
	exp := ethcommon.HexToHash("e1393fc7f6de093780674022f96cb8e3872235167d037c04d554e58c0e63d280")
	ticket := &Ticket{
		Recipient:         ethcommon.Address{},
		Sender:            ethcommon.Address{},
		FaceValue:         big.NewInt(0),
		WinProb:           big.NewInt(0),
		SenderNonce:       math.MaxUint32,
		RecipientRandHash: ethcommon.Hash{},
	}
	h := ticket.Hash()

	if h != exp {
		t.Errorf("Expected %v got %v", exp, h)
	}

	exp = ethcommon.HexToHash("ce0918ba94518293e9712effbe5fca4f1f431089833a5b8c257cb1e024595f68")
	ticket = &Ticket{
		Recipient:         ethcommon.Address{},
		Sender:            ethcommon.Address{},
		FaceValue:         big.NewInt(0),
		WinProb:           big.NewInt(0),
		SenderNonce:       1,
		RecipientRandHash: ethcommon.Hash{},
	}
	h = ticket.Hash()

	if h != exp {
		t.Errorf("Expected %v got %v", exp, h)
	}

	exp = ethcommon.HexToHash("87ef13f2d37e4d5352a01d3c77b8179d80e0887f1953bfabfd0bfd7b0f689ddd")
	ticket = &Ticket{
		Recipient:         ethcommon.Address{},
		Sender:            ethcommon.Address{},
		FaceValue:         big.NewInt(1),
		WinProb:           big.NewInt(500),
		SenderNonce:       0,
		RecipientRandHash: ethcommon.Hash{},
	}
	h = ticket.Hash()

	if h != exp {
		t.Errorf("Expected %v got %v", exp, h)
	}

	exp = ethcommon.HexToHash("aa4b35071043587992ac8b9dd1b2cf1d8311130e6458cf0b2342484e21af5f5b")
	ticket = &Ticket{
		Recipient:         ethcommon.Address{},
		Sender:            ethcommon.Address{},
		FaceValue:         big.NewInt(500),
		WinProb:           big.NewInt(1),
		SenderNonce:       0,
		RecipientRandHash: ethcommon.Hash{},
	}
	h = ticket.Hash()

	if h != exp {
		t.Errorf("Expected %v got %v", exp, h)
	}

	exp = ethcommon.HexToHash("81e353ae39046e2b77b9f996abbaeed527fda559b61ef90f299c4499dd508ccc")
	ticket = &Ticket{
		Recipient:         ethcommon.HexToAddress("73AEd7b5dEb30222fa896f399d46cC99c7BEe57F"),
		Sender:            ethcommon.Address{},
		FaceValue:         big.NewInt(0),
		WinProb:           big.NewInt(0),
		SenderNonce:       0,
		RecipientRandHash: ethcommon.Hash{},
	}
	h = ticket.Hash()

	if h != exp {
		t.Errorf("Expected %v got %v", exp, h)
	}

	var recipientRandHash [32]byte
	copy(recipientRandHash[:], ethcommon.FromHex("41b1a0649752af1b28b3dc29a1556eee781e4a4c3a1f7f53f90fa834de098c4d")[:32])

	exp = ethcommon.HexToHash("be5eaf9a49a39540b9bb02f1a21904f61324a29819e2c032824bfaa10c100b17")
	ticket = &Ticket{
		Recipient:         ethcommon.Address{},
		Sender:            ethcommon.Address{},
		FaceValue:         big.NewInt(0),
		WinProb:           big.NewInt(0),
		SenderNonce:       0,
		RecipientRandHash: recipientRandHash,
	}
	h = ticket.Hash()

	if h != exp {
		t.Errorf("Expected %v got %v", exp, h)
	}

	exp = ethcommon.HexToHash("0xe502907c16036ab3d11b78c5a2e93c20f3d6415c67f91de4ed01348182b3b2e2")
	ticket = &Ticket{
		Recipient:              ethcommon.Address{},
		Sender:                 ethcommon.Address{},
		FaceValue:              big.NewInt(0),
		WinProb:                big.NewInt(0),
		SenderNonce:            0,
		RecipientRandHash:      ethcommon.Hash{},
		CreationRound:          10,
		CreationRoundBlockHash: ethcommon.HexToHash("0x41b1a0649752af1b28b3dc29a1556eee781e4a4c3a1f7f53f90fa834de098c4d"),
	}
	h = ticket.Hash()

	if h != exp {
		t.Errorf("Expected %x got %x", exp, h)
	}
}
