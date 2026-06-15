package device

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"time"
)

type DeviceCertificate struct {
	ID           string    `json:"id"`
	TenantID     string    `json:"tenant_id"`
	DeviceID     string    `json:"device_id"`
	SerialNumber string    `json:"serial_number"`
	Issuer       string    `json:"issuer"`
	Subject      string    `json:"subject"`
	Fingerprint  string    `json:"fingerprint"`
	NotBefore    time.Time `json:"not_before"`
	NotAfter     time.Time `json:"not_after"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CertificateEnrollment struct {
	ID           string    `json:"id"`
	TenantID     string    `json:"tenant_id"`
	DeviceID     string    `json:"device_id"`
	CSR          string    `json:"csr"`
	Status       string    `json:"status"`
	Certificate  string    `json:"certificate,omitempty"`
	IssuedAt     *time.Time `json:"issued_at,omitempty"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

type CertificateRevocation struct {
	ID              string    `json:"id"`
	CertificateID   string    `json:"certificate_id"`
	Reason          string    `json:"reason"`
	RevokedAt       time.Time `json:"revoked_at"`
	RevokedBy       string    `json:"revoked_by"`
}

type DeviceCertificateRepository interface {
	Create(c *DeviceCertificate) error
	GetByID(id string) (*DeviceCertificate, error)
	GetByDeviceID(deviceID string) (*DeviceCertificate, error)
	GetByFingerprint(fingerprint string) (*DeviceCertificate, error)
	Update(c *DeviceCertificate) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*DeviceCertificate, error)
	ListExpiringBefore(tenantID string, before time.Time) ([]*DeviceCertificate, error)
}

type CertificateEnrollmentRepository interface {
	Create(e *CertificateEnrollment) error
	GetByID(id string) (*CertificateEnrollment, error)
	Update(e *CertificateEnrollment) error
	ListByTenant(tenantID string) ([]*CertificateEnrollment, error)
}

type CertificateRevocationRepository interface {
	Create(r *CertificateRevocation) error
	GetByID(id string) (*CertificateRevocation, error)
	GetByCertificateID(certID string) (*CertificateRevocation, error)
}

type DeviceCertificateService struct {
	certRepo     DeviceCertificateRepository
	enrollRepo   CertificateEnrollmentRepository
	revokeRepo   CertificateRevocationRepository
}

func NewDeviceCertificateService(cr DeviceCertificateRepository, er CertificateEnrollmentRepository, rr CertificateRevocationRepository) *DeviceCertificateService {
	return &DeviceCertificateService{
		certRepo:   cr,
		enrollRepo: er,
		revokeRepo: rr,
	}
}

func (s *DeviceCertificateService) EnrollDevice(tenantID, deviceID, csrPEM string) (*CertificateEnrollment, error) {
	enrollment := &CertificateEnrollment{
		TenantID: tenantID,
		DeviceID: deviceID,
		CSR:      csrPEM,
		Status:   "pending",
		CreatedAt: time.Now(),
	}

	if err := s.enrollRepo.Create(enrollment); err != nil {
		return nil, err
	}

	return enrollment, nil
}

func (s *DeviceCertificateService) ValidateCertificate(certPEM string) (bool, *DeviceCertificate, error) {
	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		return false, nil, nil
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return false, nil, err
	}

	fingerprint := sha256.Sum256(cert.Raw)
	fingerprintHex := hex.EncodeToString(fingerprint[:])

	existing, err := s.certRepo.GetByFingerprint(fingerprintHex)
	if err != nil {
		return false, nil, nil
	}

	if existing == nil {
		return false, nil, nil
	}

	if existing.Status != "active" {
		return false, nil, nil
	}

	now := time.Now()
	if now.Before(existing.NotBefore) || now.After(existing.NotAfter) {
		return false, nil, nil
	}

	return true, existing, nil
}

func (s *DeviceCertificateService) TrackExpiry(tenantID string) ([]*DeviceCertificate, error) {
	threshold := time.Now().AddDate(0, 30, 0)
	return s.certRepo.ListExpiringBefore(tenantID, threshold)
}

func (s *DeviceCertificateService) RevokeCertificate(certID, reason, revokedBy string) error {
	cert, err := s.certRepo.GetByID(certID)
	if err != nil {
		return err
	}

	revocation := &CertificateRevocation{
		CertificateID: certID,
		Reason:        reason,
		RevokedAt:     time.Now(),
		RevokedBy:     revokedBy,
	}

	if err := s.revokeRepo.Create(revocation); err != nil {
		return err
	}

	cert.Status = "revoked"
	cert.UpdatedAt = time.Now()

	return s.certRepo.Update(cert)
}

func (s *DeviceCertificateService) GetCertificate(id string) (*DeviceCertificate, error) {
	return s.certRepo.GetByID(id)
}

func (s *DeviceCertificateService) ListCertificates(tenantID string) ([]*DeviceCertificate, error) {
	return s.certRepo.ListByTenant(tenantID)
}

func (s *DeviceCertificateService) CheckCertificateTrust(cert *DeviceCertificate) (bool, string) {
	if cert.Status != "active" {
		return false, "certificate is not active"
	}

	now := time.Now()
	if now.Before(cert.NotBefore) {
		return false, "certificate not yet valid"
	}
	if now.After(cert.NotAfter) {
		return false, "certificate has expired"
	}

	return true, "certificate is trusted"
}
