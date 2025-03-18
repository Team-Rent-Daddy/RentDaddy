// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type AccountStatus string

const (
	AccountStatusActive    AccountStatus = "active"
	AccountStatusInactive  AccountStatus = "inactive"
	AccountStatusSuspended AccountStatus = "suspended"
)

func (e *AccountStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AccountStatus(s)
	case string:
		*e = AccountStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for AccountStatus: %T", src)
	}
	return nil
}

type NullAccountStatus struct {
	AccountStatus AccountStatus `json:"Account_Status"`
	Valid         bool          `json:"valid"` // Valid is true if AccountStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAccountStatus) Scan(value interface{}) error {
	if value == nil {
		ns.AccountStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AccountStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAccountStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AccountStatus), nil
}

type ComplaintCategory string

const (
	ComplaintCategoryMaintenance     ComplaintCategory = "maintenance"
	ComplaintCategoryNoise           ComplaintCategory = "noise"
	ComplaintCategorySecurity        ComplaintCategory = "security"
	ComplaintCategoryParking         ComplaintCategory = "parking"
	ComplaintCategoryNeighbor        ComplaintCategory = "neighbor"
	ComplaintCategoryTrash           ComplaintCategory = "trash"
	ComplaintCategoryInternet        ComplaintCategory = "internet"
	ComplaintCategoryLease           ComplaintCategory = "lease"
	ComplaintCategoryNaturalDisaster ComplaintCategory = "natural_disaster"
	ComplaintCategoryOther           ComplaintCategory = "other"
)

func (e *ComplaintCategory) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ComplaintCategory(s)
	case string:
		*e = ComplaintCategory(s)
	default:
		return fmt.Errorf("unsupported scan type for ComplaintCategory: %T", src)
	}
	return nil
}

type NullComplaintCategory struct {
	ComplaintCategory ComplaintCategory `json:"Complaint_Category"`
	Valid             bool              `json:"valid"` // Valid is true if ComplaintCategory is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullComplaintCategory) Scan(value interface{}) error {
	if value == nil {
		ns.ComplaintCategory, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ComplaintCategory.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullComplaintCategory) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ComplaintCategory), nil
}

type ComplianceStatus string

const (
	ComplianceStatusPendingReview ComplianceStatus = "pending_review"
	ComplianceStatusCompliant     ComplianceStatus = "compliant"
	ComplianceStatusNonCompliant  ComplianceStatus = "non_compliant"
	ComplianceStatusExempted      ComplianceStatus = "exempted"
)

func (e *ComplianceStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ComplianceStatus(s)
	case string:
		*e = ComplianceStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for ComplianceStatus: %T", src)
	}
	return nil
}

type NullComplianceStatus struct {
	ComplianceStatus ComplianceStatus `json:"Compliance_Status"`
	Valid            bool             `json:"valid"` // Valid is true if ComplianceStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullComplianceStatus) Scan(value interface{}) error {
	if value == nil {
		ns.ComplianceStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ComplianceStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullComplianceStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ComplianceStatus), nil
}

type LeaseStatus string

const (
	LeaseStatusDraft           LeaseStatus = "draft"
	LeaseStatusPendingApproval LeaseStatus = "pending_approval"
	LeaseStatusActive          LeaseStatus = "active"
	LeaseStatusExpired         LeaseStatus = "expired"
	LeaseStatusTerminated      LeaseStatus = "terminated"
	LeaseStatusRenewed         LeaseStatus = "renewed"
)

func (e *LeaseStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = LeaseStatus(s)
	case string:
		*e = LeaseStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for LeaseStatus: %T", src)
	}
	return nil
}

type NullLeaseStatus struct {
	LeaseStatus LeaseStatus `json:"Lease_Status"`
	Valid       bool        `json:"valid"` // Valid is true if LeaseStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullLeaseStatus) Scan(value interface{}) error {
	if value == nil {
		ns.LeaseStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.LeaseStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullLeaseStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.LeaseStatus), nil
}

type Role string

const (
	RoleTenant Role = "tenant"
	RoleAdmin  Role = "admin"
)

func (e *Role) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Role(s)
	case string:
		*e = Role(s)
	default:
		return fmt.Errorf("unsupported scan type for Role: %T", src)
	}
	return nil
}

type NullRole struct {
	Role  Role `json:"Role"`
	Valid bool `json:"valid"` // Valid is true if Role is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRole) Scan(value interface{}) error {
	if value == nil {
		ns.Role, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Role.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Role), nil
}

type Status string

const (
	StatusOpen       Status = "open"
	StatusInProgress Status = "in_progress"
	StatusResolved   Status = "resolved"
	StatusClosed     Status = "closed"
)

func (e *Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Status(s)
	case string:
		*e = Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Status: %T", src)
	}
	return nil
}

type NullStatus struct {
	Status Status `json:"Status"`
	Valid  bool   `json:"valid"` // Valid is true if Status is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatus) Scan(value interface{}) error {
	if value == nil {
		ns.Status, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Status.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Status), nil
}

type Type string

const (
	TypeLeaseAgreement Type = "lease_agreement"
	TypeAmendment      Type = "amendment"
	TypeExtension      Type = "extension"
	TypeTermination    Type = "termination"
	TypeAddendum       Type = "addendum"
)

func (e *Type) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Type(s)
	case string:
		*e = Type(s)
	default:
		return fmt.Errorf("unsupported scan type for Type: %T", src)
	}
	return nil
}

type NullType struct {
	Type  Type `json:"Type"`
	Valid bool `json:"valid"` // Valid is true if Type is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullType) Scan(value interface{}) error {
	if value == nil {
		ns.Type, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Type.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Type), nil
}

type WorkCategory string

const (
	WorkCategoryPlumbing  WorkCategory = "plumbing"
	WorkCategoryElectric  WorkCategory = "electric"
	WorkCategoryCarpentry WorkCategory = "carpentry"
	WorkCategoryHvac      WorkCategory = "hvac"
	WorkCategoryOther     WorkCategory = "other"
)

func (e *WorkCategory) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WorkCategory(s)
	case string:
		*e = WorkCategory(s)
	default:
		return fmt.Errorf("unsupported scan type for WorkCategory: %T", src)
	}
	return nil
}

type NullWorkCategory struct {
	WorkCategory WorkCategory `json:"Work_Category"`
	Valid        bool         `json:"valid"` // Valid is true if WorkCategory is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWorkCategory) Scan(value interface{}) error {
	if value == nil {
		ns.WorkCategory, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.WorkCategory.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWorkCategory) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.WorkCategory), nil
}

type Apartment struct {
	ID int64 `json:"id"`
	// describes as <building><floor><door> -> 2145
	UnitNumber     int16            `json:"unit_number"`
	Price          pgtype.Numeric   `json:"price"`
	Size           int16            `json:"size"`
	ManagementID   int64            `json:"management_id"`
	Availability   bool             `json:"availability"`
	LeaseID        int64            `json:"lease_id"`
	LeaseStartDate pgtype.Date      `json:"lease_start_date"`
	LeaseEndDate   pgtype.Date      `json:"lease_end_date"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
}

type ApartmentTenant struct {
	ApartmentID int64 `json:"apartment_id"`
	TenantID    int64 `json:"tenant_id"`
}

type Complaint struct {
	ID              int64             `json:"id"`
	ComplaintNumber int64             `json:"complaint_number"`
	CreatedBy       int64             `json:"created_by"`
	Category        ComplaintCategory `json:"category"`
	Title           string            `json:"title"`
	Description     string            `json:"description"`
	UnitNumber      int16             `json:"unit_number"`
	Status          Status            `json:"status"`
	UpdatedAt       pgtype.Timestamp  `json:"updated_at"`
	CreatedAt       pgtype.Timestamp  `json:"created_at"`
}

type Lease struct {
	ID             int64            `json:"id"`
	LeaseNumber    int64            `json:"lease_number"`
	ExternalDocID  string           `json:"external_doc_id"`
	TenantID       int64            `json:"tenant_id"`
	LandlordID     int64            `json:"landlord_id"`
	ApartmentID    pgtype.Int8      `json:"apartment_id"`
	LeaseStartDate pgtype.Date      `json:"lease_start_date"`
	LeaseEndDate   pgtype.Date      `json:"lease_end_date"`
	RentAmount     pgtype.Numeric   `json:"rent_amount"`
	LeaseStatus    LeaseStatus      `json:"lease_status"`
	CreatedBy      int64            `json:"created_by"`
	UpdatedBy      int64            `json:"updated_by"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
}

type LeaseTenant struct {
	LeaseID  int64 `json:"lease_id"`
	TenantID int64 `json:"tenant_id"`
}

type Locker struct {
	ID         int64       `json:"id"`
	AccessCode pgtype.Text `json:"access_code"`
	InUse      bool        `json:"in_use"`
	UserID     pgtype.Int8 `json:"user_id"`
}

type ParkingPermit struct {
	ID           int64            `json:"id"`
	PermitNumber int64            `json:"permit_number"`
	CreatedBy    int64            `json:"created_by"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	// 5 days long
	ExpiresAt pgtype.Timestamp `json:"expires_at"`
}

type User struct {
	ID int64 `json:"id"`
	// provided by Clerk
	ClerkID    string           `json:"clerk_id"`
	FirstName  string           `json:"first_name"`
	LastName   string           `json:"last_name"`
	Email      string           `json:"email"`
	Phone      pgtype.Text      `json:"phone"`
	UnitNumber pgtype.Int2      `json:"unit_number"`
	Role       Role             `json:"role"`
	Status     AccountStatus    `json:"status"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

type WorkOrder struct {
	ID          int64            `json:"id"`
	CreatedBy   int64            `json:"created_by"`
	OrderNumber int64            `json:"order_number"`
	Category    WorkCategory     `json:"category"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	UnitNumber  int16            `json:"unit_number"`
	Status      Status           `json:"status"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}
