package models

type PasswordResetRequest struct {
	RequestedForUsername  string `json:"requestedForUsername" firestore:"requestedForUsername"`
	RequestedForEmail     string `json:"requestedForEmail" firestore:"requestedForEmail"`
	RequestedTimestamp    int64  `json:"requestedTimestamp" firestore:"requestedTimestamp"`
	ExpirationTimestamp   int64  `json:"expirationTimestamp" firestore:"expirationTimestamp"`
	RequestId			  string `json:"requestId" firestore:"requestId"`
	ResetCompleted        bool   `json:"resetCompleted" firestore:"resetCompleted"`

	FirestoreID string `json:"-"`
}

func (u *PasswordResetRequest) SetFirestoreID(id string) {
    u.FirestoreID = id
}

func (u *PasswordResetRequest) GetFirestoreID() string {
	return u.FirestoreID
}