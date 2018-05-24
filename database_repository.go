package main

type RepositoryManager struct {
	db *DB
}

type repo struct {
	repositoryManager *RepositoryManager
}

func (r *repo) db() *DB {
	return r.repositoryManager.db
}

func NewRepositoryManager(db *DB) *RepositoryManager {
	return &RepositoryManager{db}
}

func (rm *RepositoryManager) Contact() ContactRepository {
	return ContactRepository{repo{rm}}
}

type ContactRepository struct {
	repo
}

type contactActiveRow struct {
	ContactRepository
	entity *DbContact
}

func (r ContactRepository) WithEntity(entity *DbContact) *contactActiveRow {
	return &contactActiveRow{r, entity}
}

func (r ContactRepository) withPreload() *DB {
	return &DB{r.db().Set("gorm:auto_preload", true)}
}

func (r ContactRepository) FindWithUUID(userId string, uuid string) *contactActiveRow {
	var contact DbContact
	r.withPreload().First(&contact, "user_id = ? AND uuid = ?", userId, uuid)
	return r.WithEntity(&contact)
}

func (r ContactRepository) Find(id uint) *contactActiveRow {
	var contact DbContact
	r.withPreload().First(&contact, id)
	return r.WithEntity(&contact)
}

func (r ContactRepository) ListAll(userId string) []*DbContact {
	var contacts []*DbContact
	r.withPreload().Where("user_id = ?", userId).Find(&contacts)
	return contacts
}

func (r *contactActiveRow) IsValid() bool {
	return r.entity != nil && r.entity.ID != 0
}

func (r *contactActiveRow) Get() *DbContact {
	if r.IsValid() {
		return r.entity
	}

	return nil
}

func (r *contactActiveRow) Create() {
	r.db().Create(r.entity)
}

func (r *contactActiveRow) Delete() {
	r.db().Delete(&r.entity)
}

func (r *contactActiveRow) CleanRelated() {
	r.db().Delete(&DbEmail{}, "contact_id = ?", r.entity.ID)
	r.db().Delete(&DbAddress{}, "contact_id = ?", r.entity.ID)
	r.db().Delete(&DbPhone{}, "contact_id = ?", r.entity.ID)

	r.entity.Emails = nil
	r.entity.Addresses = nil
	r.entity.Phones = nil
}

func (r *contactActiveRow) Save() {
	if r.entity.ID == 0 {
		r.db().Create(r.entity)
	} else {
		r.db().Save(r.entity)
	}
}
