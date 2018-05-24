package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
)

func (h *httpApiHandler) ListContacts(r *http.Request) (interface{}, error) {
	uid, err := getUserId(r)
	if err != nil {
		return nil, err
	}

	dbContacts := h.rm.Contact().ListAll(uid)

	apiContacts := make([]*ApiContact, len(dbContacts))
	for i, cnt := range dbContacts {
		apiContacts[i] = dtaContact(cnt)
	}

	return apiContacts, nil
}

func (h *httpApiHandler) GetContact(r *http.Request) (interface{}, error) {
	uid, err := getUserId(r)
	if err != nil {
		return nil, err
	}

	cid, err := getContactId(r)
	if err != nil {
		return nil, err
	}

	dbContactRow := h.rm.Contact().FindWithUUID(uid, cid)
	if !dbContactRow.IsValid() {
		return nil, badRequestError{errors.New("not found")}
	}

	apiContact := dtaContact(dbContactRow.Get())
	return apiContact, nil
}

func (h *httpApiHandler) CreateContact(r *http.Request) (interface{}, error) {
	uid, err := getUserId(r)
	if err != nil {
		return nil, err
	}

	apiContact, err := getContactBody(r)
	if err != nil {
		return nil, err
	}

	dbContact := atdContact(apiContact)
	dbContact.UserId = uid

	h.rm.Contact().WithEntity(dbContact).Save()

	reloadedRow := h.rm.Contact().Find(dbContact.ID)
	if reloadedRow == nil {
		return nil, internalErrorErr
	}

	displayContact := dtaContact(reloadedRow.Get())
	return displayContact, nil
}

func (h *httpApiHandler) UpdateContact(r *http.Request) (interface{}, error) {
	uid, err := getUserId(r)
	if err != nil {
		return nil, err
	}

	cid, err := getContactId(r)
	if err != nil {
		return nil, err
	}

	apiContact, err := getContactBody(r)
	if err != nil {
		return nil, err
	}

	dbContactRow := h.rm.Contact().FindWithUUID(uid, cid)
	if !dbContactRow.IsValid() {
		return nil, badRequestError{errors.New("not found")}
	}

	dbContactRow.CleanRelated()

	dbUpdatedContact := atdContact(apiContact)
	dbContactRow.Get().UpdateWith(dbUpdatedContact)
	dbContactRow.Save()

	reloadedRow := h.rm.Contact().Find(dbContactRow.Get().ID)
	if reloadedRow == nil {
		return nil, internalErrorErr
	}

	displayContact := dtaContact(reloadedRow.Get())
	return displayContact, nil
}

func (h *httpApiHandler) DeleteContact(r *http.Request) (interface{}, error) {
	uid, err := getUserId(r)
	if err != nil {
		return nil, err
	}

	cid, err := getContactId(r)
	if err != nil {
		return nil, err
	}

	dbContactRow := h.rm.Contact().FindWithUUID(uid, cid)
	if !dbContactRow.IsValid() {
		return nil, badRequestError{errors.New("not found")}
	}

	dbContactRow.Delete()
	return nil, nil
}

func getUserId(r *http.Request) (string, error) {
	params := mux.Vars(r)
	if uid, ok := params["user_id"]; ok {
		return uid, nil
	}

	return "", badRequestError{errors.New("invalid user_id")}
}

func getContactId(r *http.Request) (string, error) {
	params := mux.Vars(r)
	if uid, ok := params["contact_id"]; ok {
		return uid, nil
	}

	return "", badRequestError{errors.New("invalid contact_id")}
}

func getContactBody(r *http.Request) (*ApiContact, error) {
	allowedJsonMimeTypes := []string{"application/json"}

	contentTypeHeaderValue := r.Header.Get("Content-Type")
	mimeType, _, err := mime.ParseMediaType(contentTypeHeaderValue)
	if err != nil || !strSliceContains(allowedJsonMimeTypes, mimeType) {
		return nil, badRequestError{errors.New("invalid content-type")}
	}

	var apiContact ApiContact
	err = json.NewDecoder(r.Body).Decode(&apiContact)
	defer r.Body.Close()

	if err != nil {
		return nil, badRequestError{errors.New("invalid body")}
	}

	err = apiContact.Validate()
	if err != nil {
		return nil, badRequestError{fmt.Errorf("invalid body, error: %v", err)}
	}

	return &apiContact, nil
}
