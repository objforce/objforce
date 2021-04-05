package dtos

type EncryptionScheme string

const(
	ProbabilisticEncryption EncryptionScheme = "ProbabilisticEncryption"
	CaseSensitiveDeterministicEncryption EncryptionScheme = "CaseSensitiveDeterministicEncryption"
	CaseInsensitiveDeterministicEncryption EncryptionScheme = "CaseInsensitiveDeterministicEncryption"
)