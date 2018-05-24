package main

// Contact
func atdContact(contact *ApiContact) *DbContact {
	return &DbContact{
		IdentifiableDbModel: IdentifiableDbModel{UUID: contact.UUID},

		TitleBefore: contact.TitleBefore,
		FirstName:   contact.FirstName,
		MiddleName:  contact.MiddleName,
		LastName:    contact.LastName,
		TitleAfter:  contact.TitleAfter,

		Emails:    atdEmails(contact.Emails),
		Addresses: atdAddresses(contact.Addresses),
		Phones:    atdPhones(contact.Phones),

		Birthday:     contact.Birthday,
		Note:         contact.Note,
		Organization: contact.Organization,
	}
}

func dtaContact(contact *DbContact) *ApiContact {
	return &ApiContact{
		UUID: contact.UUID,

		TitleBefore: contact.TitleBefore,
		FirstName:   contact.FirstName,
		MiddleName:  contact.MiddleName,
		LastName:    contact.LastName,
		TitleAfter:  contact.TitleAfter,

		Emails:    dtaEmails(contact.Emails),
		Addresses: dtaAddresses(contact.Addresses),
		Phones:    dtaPhones(contact.Phones),

		Birthday:     contact.Birthday,
		Note:         contact.Note,
		Organization: contact.Organization,
	}
}

// Email
func atdEmails(emails []*ApiEmail) []*DbEmail {
	var dbMails []*DbEmail
	for _, mail := range emails {
		dbMails = append(dbMails, atdEmail(mail))
	}

	return dbMails
}

func atdEmail(email *ApiEmail) *DbEmail {
	return &DbEmail{
		Type:  email.Type,
		Email: email.Email,
	}
}

func dtaEmails(emails []*DbEmail) []*ApiEmail {
	var dbMails []*ApiEmail
	for _, mail := range emails {
		dbMails = append(dbMails, dtaEmail(mail))
	}

	return dbMails
}

func dtaEmail(email *DbEmail) *ApiEmail {
	return &ApiEmail{
		Type:  email.Type,
		Email: email.Email,
	}
}

// Address
func atdAddresses(addresses []*ApiAddress) []*DbAddress {
	var dbMails []*DbAddress
	for _, mail := range addresses {
		dbMails = append(dbMails, atdAddress(mail))
	}

	return dbMails
}

func atdAddress(address *ApiAddress) *DbAddress {
	return &DbAddress{
		Type:    address.Type,
		Address: address.Address,
	}
}

func dtaAddresses(addresses []*DbAddress) []*ApiAddress {
	var dbMails []*ApiAddress
	for _, mail := range addresses {
		dbMails = append(dbMails, dtaAddress(mail))
	}

	return dbMails
}

func dtaAddress(address *DbAddress) *ApiAddress {
	return &ApiAddress{
		Type:    address.Type,
		Address: address.Address,
	}
}

// Phone
func atdPhones(phones []*ApiPhone) []*DbPhone {
	var dbMails []*DbPhone
	for _, mail := range phones {
		dbMails = append(dbMails, atdPhone(mail))
	}

	return dbMails
}

func atdPhone(phone *ApiPhone) *DbPhone {
	return &DbPhone{
		Type:        phone.Type,
		PhoneNumber: phone.PhoneNumber,
	}
}

func dtaPhones(phones []*DbPhone) []*ApiPhone {
	var dbMails []*ApiPhone
	for _, mail := range phones {
		dbMails = append(dbMails, dtaPhone(mail))
	}

	return dbMails
}

func dtaPhone(phone *DbPhone) *ApiPhone {
	return &ApiPhone{
		Type:        phone.Type,
		PhoneNumber: phone.PhoneNumber,
	}
}
