package db

import (
	"context"
	"simpleBank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountId: account1.ID,
		ToAccountId:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountId, transfer.FromAccountId)
	require.Equal(t, arg.ToAccountId, transfer.ToAccountId)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	CreateRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	transfer1 := CreateRandomTransfer(t, account1, account2)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountId, transfer2.FromAccountId)
	require.Equal(t, transfer1.ToAccountId, transfer2.ToAccountId)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t, account1, account2)
	}
	arg := ListTransfersParams{
		Limit:         5,
		Offset:        5,
		FromAccountId: account1.ID,
		ToAccountId:   account2.ID,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, trans := range transfers {
		require.NotEmpty(t, trans)
		require.True(t, trans.FromAccountId == account1.ID || trans.ToAccountId == account1.ID)
	}

}
