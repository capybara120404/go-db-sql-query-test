package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := 1

	// напиши тест здесь
	client, err := selectClient(db, clientID)
	require.NoError(t, err)
	require.Equal(t, client.ID, clientID)
	require.NotEmpty(t, client.FIO)
	require.NotEmpty(t, client.Login)
	require.NotEmpty(t, client.Birthday)
	require.NotEmpty(t, client.Email)
	
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()
	
	clientID := -1

	// напиши тест здесь
	client, err := selectClient(db, clientID)
	require.Equal(t, err, sql.ErrNoRows)
	require.Empty(t, client.ID)
	require.Empty(t, client.FIO)
	require.Empty(t, client.Login)
	require.Empty(t, client.Birthday)
	require.Empty(t, client.Email)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	cl.ID, err = insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)
	client, err := selectClient(db, cl.ID)
	require.NoError(t, err)
	require.Equal(t, cl.ID, client.ID)
	require.Equal(t, cl.FIO, client.FIO)
	require.Equal(t, cl.Login, client.Login)
	require.Equal(t, cl.Birthday, client.Birthday)
	require.Equal(t, cl.Email, client.Email)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	cl.ID, err = insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)
	client, err := selectClient(db, cl.ID)
	require.NoError(t, err)
	err = deleteClient(db, client.ID)
	require.NoError(t,err)
	newClient, err := selectClient(db, client.ID)
	require.Empty(t,newClient)
	require.Equal(t, err, sql.ErrNoRows)
}
