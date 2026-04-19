package models

type FirestoreRecordBase interface {
	GetFirestoreID() string
	SetFirestoreID(string)
}